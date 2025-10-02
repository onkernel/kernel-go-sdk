package basic_test

import (
	"testing"

	kernel "github.com/onkernel/kernel-go-sdk"
	_ "github.com/onkernel/kernel-go-sdk/examples/basic" // import for side-effects (init)
)

func TestBasicAppRegistration(t *testing.T) {
	app := kernel.GetApp("go-basic")
	if app == nil {
		t.Fatalf("expected app to be registered, got nil")
	}

	actions := app.GetActions()
	if len(actions) != 4 {
		t.Fatalf("expected 4 actions registered, got %d", len(actions))
	}

	expected := map[string]bool{
		"ping":              true,
		"get-random-number": true,
		"print-url":         true,
		"get-page-title":    true,
	}
	for _, act := range actions {
		if !expected[act.Name] {
			t.Fatalf("unexpected action name %s", act.Name)
		}
	}
}
