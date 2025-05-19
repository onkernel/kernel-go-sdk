// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/onkernel/kernel-go-sdk/internal/apiform"
	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/requestconfig"
	"github.com/onkernel/kernel-go-sdk/option"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
	"github.com/onkernel/kernel-go-sdk/packages/ssestream"
	"github.com/onkernel/kernel-go-sdk/shared/constant"
)

// AppDeploymentService contains methods and other services that help with
// interacting with the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppDeploymentService] method instead.
type AppDeploymentService struct {
	Options []option.RequestOption
}

// NewAppDeploymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppDeploymentService(opts ...option.RequestOption) (r AppDeploymentService) {
	r = AppDeploymentService{}
	r.Options = opts
	return
}

// Deploy a new application
func (r *AppDeploymentService) New(ctx context.Context, body AppDeploymentNewParams, opts ...option.RequestOption) (res *AppDeploymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "deploy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Establishes a Server-Sent Events (SSE) stream that delivers real-time logs and
// status updates for a deployed application. The stream terminates automatically
// once the application reaches a terminal state.
func (r *AppDeploymentService) FollowStreaming(ctx context.Context, id string, opts ...option.RequestOption) (stream *ssestream.Stream[[]AppDeploymentFollowResponseUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("apps/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[[]AppDeploymentFollowResponseUnion](ssestream.NewDecoder(raw), err)
}

type AppDeploymentNewResponse struct {
	// List of apps deployed
	Apps []AppDeploymentNewResponseApp `json:"apps,required"`
	// Current status of the deployment
	//
	// Any of "queued", "deploying", "succeeded", "failed".
	Status AppDeploymentNewResponseStatus `json:"status,required"`
	// Status reason
	StatusReason string `json:"status_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Apps         respjson.Field
		Status       respjson.Field
		StatusReason respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeploymentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AppDeploymentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppDeploymentNewResponseApp struct {
	// ID for the app version deployed
	ID string `json:"id,required"`
	// List of actions available on the app
	Actions []AppDeploymentNewResponseAppAction `json:"actions,required"`
	// Name of the app
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeploymentNewResponseApp) RawJSON() string { return r.JSON.raw }
func (r *AppDeploymentNewResponseApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppDeploymentNewResponseAppAction struct {
	// Name of the action
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeploymentNewResponseAppAction) RawJSON() string { return r.JSON.raw }
func (r *AppDeploymentNewResponseAppAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the deployment
type AppDeploymentNewResponseStatus string

const (
	AppDeploymentNewResponseStatusQueued    AppDeploymentNewResponseStatus = "queued"
	AppDeploymentNewResponseStatusDeploying AppDeploymentNewResponseStatus = "deploying"
	AppDeploymentNewResponseStatusSucceeded AppDeploymentNewResponseStatus = "succeeded"
	AppDeploymentNewResponseStatusFailed    AppDeploymentNewResponseStatus = "failed"
)

// AppDeploymentFollowResponseUnion contains all possible properties and values
// from [AppDeploymentFollowResponseState],
// [AppDeploymentFollowResponseStateUpdate], [AppDeploymentFollowResponseLog].
//
// Use the [AppDeploymentFollowResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AppDeploymentFollowResponseUnion struct {
	// Any of "state", "state_update", "log".
	Event     string    `json:"event"`
	State     string    `json:"state"`
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [AppDeploymentFollowResponseLog].
	Message string `json:"message"`
	JSON    struct {
		Event     respjson.Field
		State     respjson.Field
		Timestamp respjson.Field
		Message   respjson.Field
		raw       string
	} `json:"-"`
}

// anyAppDeploymentFollowResponse is implemented by each variant of
// [AppDeploymentFollowResponseUnion] to add type safety for the return type of
// [AppDeploymentFollowResponseUnion.AsAny]
type anyAppDeploymentFollowResponse interface {
	implAppDeploymentFollowResponseUnion()
}

func (AppDeploymentFollowResponseState) implAppDeploymentFollowResponseUnion()       {}
func (AppDeploymentFollowResponseStateUpdate) implAppDeploymentFollowResponseUnion() {}
func (AppDeploymentFollowResponseLog) implAppDeploymentFollowResponseUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AppDeploymentFollowResponseUnion.AsAny().(type) {
//	case kernel.AppDeploymentFollowResponseState:
//	case kernel.AppDeploymentFollowResponseStateUpdate:
//	case kernel.AppDeploymentFollowResponseLog:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AppDeploymentFollowResponseUnion) AsAny() anyAppDeploymentFollowResponse {
	switch u.Event {
	case "state":
		return u.AsState()
	case "state_update":
		return u.AsStateUpdate()
	case "log":
		return u.AsLog()
	}
	return nil
}

func (u AppDeploymentFollowResponseUnion) AsState() (v AppDeploymentFollowResponseState) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppDeploymentFollowResponseUnion) AsStateUpdate() (v AppDeploymentFollowResponseStateUpdate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppDeploymentFollowResponseUnion) AsLog() (v AppDeploymentFollowResponseLog) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AppDeploymentFollowResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AppDeploymentFollowResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial state of the application, emitted once when subscribing.
type AppDeploymentFollowResponseState struct {
	// Event type identifier (always "state").
	Event constant.State `json:"event,required"`
	// Current application state (e.g., "deploying", "running", "succeeded", "failed").
	State string `json:"state,required"`
	// Time the state was reported.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		State       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeploymentFollowResponseState) RawJSON() string { return r.JSON.raw }
func (r *AppDeploymentFollowResponseState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An update emitted when the application's state changes.
type AppDeploymentFollowResponseStateUpdate struct {
	// Event type identifier (always "state_update").
	Event constant.StateUpdate `json:"event,required"`
	// New application state (e.g., "running", "succeeded", "failed").
	State string `json:"state,required"`
	// Time the state change occurred.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		State       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeploymentFollowResponseStateUpdate) RawJSON() string { return r.JSON.raw }
func (r *AppDeploymentFollowResponseStateUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A log entry from the application.
type AppDeploymentFollowResponseLog struct {
	// Event type identifier (always "log").
	Event constant.Log `json:"event,required"`
	// Log message text.
	Message string `json:"message,required"`
	// Time the log entry was produced.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeploymentFollowResponseLog) RawJSON() string { return r.JSON.raw }
func (r *AppDeploymentFollowResponseLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppDeploymentNewParams struct {
	// Relative path to the entrypoint of the application
	EntrypointRelPath string `json:"entrypoint_rel_path,required"`
	// ZIP file containing the application source directory
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Allow overwriting an existing app version
	Force param.Opt[bool] `json:"force,omitzero"`
	// Version of the application. Can be any string.
	Version param.Opt[string] `json:"version,omitzero"`
	// Map of environment variables to set for the deployed application. Each key-value
	// pair represents an environment variable.
	EnvVars map[string]string `json:"env_vars,omitzero"`
	// Region for deployment. Currently we only support "aws.us-east-1a"
	//
	// Any of "aws.us-east-1a".
	Region AppDeploymentNewParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r AppDeploymentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Region for deployment. Currently we only support "aws.us-east-1a"
type AppDeploymentNewParamsRegion string

const (
	AppDeploymentNewParamsRegionAwsUsEast1a AppDeploymentNewParamsRegion = "aws.us-east-1a"
)
