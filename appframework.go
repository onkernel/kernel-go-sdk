package kernel

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

// KernelContext contains metadata that is propagated through a Go context.Context.
// At the moment it only exposes the invocation ID of the running action, but this
// struct can be extended in the future as we add more metadata.
//
// Use kernel.WithInvocationID to attach a KernelContext to an existing
// context.Context and kernel.Context to retrieve it back.
//
//	ctx := kernel.WithInvocationID(context.Background(), "inv-123")
//	kctx := kernel.Context(ctx) // => KernelContext{InvocationID: "inv-123"}
//
// End-users typically won't build the context themselves – the runtime that
// executes actions will do it – but providing the helpers here allows for unit
// testing and for advanced users to craft their own contexts when needed.
//
// NOTE: we intentionally keep the struct exported so that client code can
// inspect / copy it if they need to.
type KernelContext struct {
	InvocationID string
}

type kernelCtxKey struct{}

// WithInvocationID returns a child context that contains a KernelContext with
// the provided invocation id.
func WithInvocationID(ctx context.Context, invocationID string) context.Context {
	return context.WithValue(ctx, kernelCtxKey{}, KernelContext{InvocationID: invocationID})
}

// Context extracts the KernelContext stored inside the passed in ctx. If the
// context was not carrying any KernelContext value, the zero value is
// returned.
func Context(ctx context.Context) KernelContext {
	if v, ok := ctx.Value(kernelCtxKey{}).(KernelContext); ok {
		return v
	}
	return KernelContext{}
}

// -----------------------------------------------------------------------------
// Global App Registry
// -----------------------------------------------------------------------------

type appRegistry struct {
	mu   sync.RWMutex
	apps map[string]*KernelApp
}

func (r *appRegistry) registerApp(app *KernelApp) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.apps[app.Name] = app
}

func (r *appRegistry) getApp(name string) (*KernelApp, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	app, ok := r.apps[name]
	return app, ok
}

func (r *appRegistry) getApps() []*KernelApp {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*KernelApp, 0, len(r.apps))
	for _, app := range r.apps {
		out = append(out, app)
	}
	return out
}

var globalRegistry = &appRegistry{apps: make(map[string]*KernelApp)}

// App returns a *KernelApp with the given name, creating it if necessary. The
// returned app is automatically registered in the global registry so that the
// runtime can discover it later.
func App(name string) *KernelApp {
	if existing, ok := globalRegistry.getApp(name); ok {
		return existing
	}
	app := &KernelApp{
		Name:    name,
		actions: make(map[string]KernelAction),
	}
	globalRegistry.registerApp(app)
	return app
}

// Apps returns a slice with all the apps currently registered in the global registry.
func Apps() []*KernelApp { return globalRegistry.getApps() }

// GetApp retrieves an app by name from the global registry. It returns nil if
// the app could not be found.
func GetApp(name string) *KernelApp {
	if app, ok := globalRegistry.getApp(name); ok {
		return app
	}
	return nil
}

// Export returns a serialisable representation of the current registry. It is
// primarily useful for debugging and testing.
func Export() KernelJSON {
	apps := make([]KernelAppJSON, 0, len(globalRegistry.apps))
	for _, a := range Apps() {
		apps = append(apps, a.toJSON())
	}
	return KernelJSON{Apps: apps}
}

// ExportJSON exports the registry to a JSON string. The JSON is indented with
// two spaces to make it human-readable.
func ExportJSON() string {
	raw, _ := json.MarshalIndent(Export(), "", "  ")
	return string(raw)
}

// -----------------------------------------------------------------------------
// Kernel App & Actions
// -----------------------------------------------------------------------------

// KernelAction wraps the user-provided handler together with some metadata that
// the runtime may need when invoking the action.
//
// The internal exec function brings all handlers (with or without inputs /
// outputs) to a single canonical signature so that the runtime can treat them
// uniformly.
type KernelAction struct {
	Name string

	// handler is the original user-provided function so that we can surface it
	// back when needed (for example when generating reflection based stubs).
	handler any

	// exec adapts every supported signature to func(ctx, input) (output, err).
	exec func(ctx context.Context, input any) (any, error)
}

// Exec executes the adapted handler. It is **not** normally called by end-users
// but is exported to simplify unit testing.
func (a KernelAction) Exec(ctx context.Context, input any) (any, error) { // nolint: revive // false positive on stutter
	return a.exec(ctx, input)
}

// KernelApp represents a collection of actions that belong together.
//
// Users obtain an instance by calling kernel.App("my-app").
//
// All methods on KernelApp are safe for concurrent use.
type KernelApp struct {
	Name string

	mu      sync.RWMutex
	actions map[string]KernelAction
}

// GetActions returns a copy of all registered actions.
func (a *KernelApp) GetActions() []KernelAction {
	a.mu.RLock()
	defer a.mu.RUnlock()
	out := make([]KernelAction, 0, len(a.actions))
	for _, act := range a.actions {
		out = append(out, act)
	}
	return out
}

// GetAction retrieves a single action by name. The second return value is
// false when the action doesn't exist.
func (a *KernelApp) GetAction(name string) (KernelAction, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	act, ok := a.actions[name]
	return act, ok
}

// --- Action registration helpers ------------------------------------------------

// Action registers an action in the app. The handler must have one of the
// following signatures:
//
//  1. func(context.Context) error
//  2. func(context.Context) (Out, error)
//  3. func(context.Context, In) error
//  4. func(context.Context, In) (Out, error)
//
// "In" and "Out" can be any types (including struct{}).
//
// The method panics if the provided handler does not match any of the expected
// signatures. We choose to panic instead of returning an error because action
// registration happens at init-time and panicking provides immediate feedback
// to the developer.
func (a *KernelApp) Action(name string, handler any) {
	// Validate handler signature via reflection and create a wrapper that
	// normalises to func(ctx, input) (output, error)
	wrapper := buildActionWrapper(name, handler)
	a.addAction(name, handler, wrapper)
}

// buildActionWrapper analyses the handler's type and returns a wrapper with the
// canonical signature used internally by the runtime.
func buildActionWrapper(name string, handler any) func(context.Context, any) (any, error) {
	hv := reflect.ValueOf(handler)
	ht := hv.Type()

	if ht.Kind() != reflect.Func {
		panic(fmt.Sprintf("action %s: handler must be a function", name))
	}

	// All allowed signatures start with a context.Context parameter.
	if ht.NumIn() == 0 {
		panic(fmt.Sprintf("action %s: handler must accept context.Context", name))
	}
	firstParam := ht.In(0)
	if !firstParam.Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		panic(fmt.Sprintf("action %s: first parameter must be context.Context", name))
	}

	switch {
	// 1. func(ctx) error
	case ht.NumIn() == 1 && ht.NumOut() == 1 && ht.Out(0) == reflect.TypeOf((*error)(nil)).Elem():
		return func(ctx context.Context, _ any) (any, error) {
			out := hv.Call([]reflect.Value{reflect.ValueOf(ctx)})
			err, _ := out[0].Interface().(error)
			return nil, err
		}

	// 2. func(ctx) (Out, error)
	case ht.NumIn() == 1 && ht.NumOut() == 2 && ht.Out(1) == reflect.TypeOf((*error)(nil)).Elem():
		return func(ctx context.Context, _ any) (any, error) {
			outs := hv.Call([]reflect.Value{reflect.ValueOf(ctx)})
			result := outs[0].Interface()
			err, _ := outs[1].Interface().(error)
			return result, err
		}

	// 3. func(ctx, In) error
	case ht.NumIn() == 2 && ht.NumOut() == 1 && ht.Out(0) == reflect.TypeOf((*error)(nil)).Elem():
		inType := ht.In(1)
		return func(ctx context.Context, raw any) (any, error) {
			// Validate input type at runtime.
			if raw == nil {
				// Zero value of the expected type.
				raw = reflect.Zero(inType).Interface()
			}
			rv := reflect.ValueOf(raw)
			if !rv.IsValid() || !rv.Type().AssignableTo(inType) {
				return nil, fmt.Errorf("action %s: input type mismatch", name)
			}
			outs := hv.Call([]reflect.Value{reflect.ValueOf(ctx), rv})
			err, _ := outs[0].Interface().(error)
			return nil, err
		}

	// 4. func(ctx, In) (Out, error)
	case ht.NumIn() == 2 && ht.NumOut() == 2 && ht.Out(1) == reflect.TypeOf((*error)(nil)).Elem():
		inType := ht.In(1)
		return func(ctx context.Context, raw any) (any, error) {
			if raw == nil {
				raw = reflect.Zero(inType).Interface()
			}
			rv := reflect.ValueOf(raw)
			if !rv.IsValid() || !rv.Type().AssignableTo(inType) {
				return nil, fmt.Errorf("action %s: input type mismatch", name)
			}
			outs := hv.Call([]reflect.Value{reflect.ValueOf(ctx), rv})
			result := outs[0].Interface()
			err, _ := outs[1].Interface().(error)
			return result, err
		}
	default:
		panic(fmt.Sprintf("action %s: handler has an unsupported signature", name))
	}
}

// addAction centralises the logic of mutating the app.
func (a *KernelApp) addAction(name string, handler any, exec func(context.Context, any) (any, error)) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, exists := a.actions[name]; exists {
		panic(fmt.Sprintf("action with name %q already registered in app %q", name, a.Name))
	}
	a.actions[name] = KernelAction{Name: name, handler: handler, exec: exec}
}

// --- JSON helpers (for tests / debug) -------------------------------------------

type KernelActionJSON struct {
	Name string `json:"name"`
}

type KernelAppJSON struct {
	Name    string             `json:"name"`
	Actions []KernelActionJSON `json:"actions"`
}

type KernelJSON struct {
	Apps []KernelAppJSON `json:"apps"`
}

func (a *KernelApp) toJSON() KernelAppJSON {
	acts := make([]KernelActionJSON, 0, len(a.actions))
	for _, act := range a.GetActions() {
		acts = append(acts, KernelActionJSON{Name: act.Name})
	}
	return KernelAppJSON{Name: a.Name, Actions: acts}
}
