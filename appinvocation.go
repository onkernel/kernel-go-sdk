// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/requestconfig"
	"github.com/onkernel/kernel-go-sdk/option"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
)

// AppInvocationService contains methods and other services that help with
// interacting with the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppInvocationService] method instead.
type AppInvocationService struct {
	Options []option.RequestOption
}

// NewAppInvocationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppInvocationService(opts ...option.RequestOption) (r AppInvocationService) {
	r = AppInvocationService{}
	r.Options = opts
	return
}

// Invoke an action.
func (r *AppInvocationService) New(ctx context.Context, body AppInvocationNewParams, opts ...option.RequestOption) (res *AppInvocationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "invocations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details about an invocation's status and output.
func (r *AppInvocationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AppInvocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("invocations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an invocation's status or output.
func (r *AppInvocationService) Update(ctx context.Context, id string, body AppInvocationUpdateParams, opts ...option.RequestOption) (res *AppInvocationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("invocations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AppInvocationNewResponse struct {
	// ID of the invocation
	ID string `json:"id,required"`
	// Status of the invocation
	//
	// Any of "queued", "running", "succeeded", "failed".
	Status AppInvocationNewResponseStatus `json:"status,required"`
	// The return value of the action that was invoked, rendered as a JSON string. This
	// could be: string, number, boolean, array, object, or null.
	Output string `json:"output"`
	// Status reason
	StatusReason string `json:"status_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Status       respjson.Field
		Output       respjson.Field
		StatusReason respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInvocationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AppInvocationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the invocation
type AppInvocationNewResponseStatus string

const (
	AppInvocationNewResponseStatusQueued    AppInvocationNewResponseStatus = "queued"
	AppInvocationNewResponseStatusRunning   AppInvocationNewResponseStatus = "running"
	AppInvocationNewResponseStatusSucceeded AppInvocationNewResponseStatus = "succeeded"
	AppInvocationNewResponseStatusFailed    AppInvocationNewResponseStatus = "failed"
)

type AppInvocationGetResponse struct {
	// ID of the invocation
	ID string `json:"id,required"`
	// Name of the action invoked
	ActionName string `json:"action_name,required"`
	// Name of the application
	AppName string `json:"app_name,required"`
	// RFC 3339 Nanoseconds timestamp when the invocation started
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Status of the invocation
	//
	// Any of "queued", "running", "succeeded", "failed".
	Status AppInvocationGetResponseStatus `json:"status,required"`
	// RFC 3339 Nanoseconds timestamp when the invocation finished (null if still
	// running)
	FinishedAt time.Time `json:"finished_at,nullable" format:"date-time"`
	// Output produced by the action, rendered as a JSON string. This could be: string,
	// number, boolean, array, object, or null.
	Output string `json:"output"`
	// Payload provided to the invocation. This is a string that can be parsed as JSON.
	Payload string `json:"payload"`
	// Status reason
	StatusReason string `json:"status_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ActionName   respjson.Field
		AppName      respjson.Field
		StartedAt    respjson.Field
		Status       respjson.Field
		FinishedAt   respjson.Field
		Output       respjson.Field
		Payload      respjson.Field
		StatusReason respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInvocationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AppInvocationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the invocation
type AppInvocationGetResponseStatus string

const (
	AppInvocationGetResponseStatusQueued    AppInvocationGetResponseStatus = "queued"
	AppInvocationGetResponseStatusRunning   AppInvocationGetResponseStatus = "running"
	AppInvocationGetResponseStatusSucceeded AppInvocationGetResponseStatus = "succeeded"
	AppInvocationGetResponseStatusFailed    AppInvocationGetResponseStatus = "failed"
)

type AppInvocationUpdateResponse struct {
	// ID of the invocation
	ID string `json:"id,required"`
	// Name of the action invoked
	ActionName string `json:"action_name,required"`
	// Name of the application
	AppName string `json:"app_name,required"`
	// RFC 3339 Nanoseconds timestamp when the invocation started
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Status of the invocation
	//
	// Any of "queued", "running", "succeeded", "failed".
	Status AppInvocationUpdateResponseStatus `json:"status,required"`
	// RFC 3339 Nanoseconds timestamp when the invocation finished (null if still
	// running)
	FinishedAt time.Time `json:"finished_at,nullable" format:"date-time"`
	// Output produced by the action, rendered as a JSON string. This could be: string,
	// number, boolean, array, object, or null.
	Output string `json:"output"`
	// Payload provided to the invocation. This is a string that can be parsed as JSON.
	Payload string `json:"payload"`
	// Status reason
	StatusReason string `json:"status_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ActionName   respjson.Field
		AppName      respjson.Field
		StartedAt    respjson.Field
		Status       respjson.Field
		FinishedAt   respjson.Field
		Output       respjson.Field
		Payload      respjson.Field
		StatusReason respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInvocationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AppInvocationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the invocation
type AppInvocationUpdateResponseStatus string

const (
	AppInvocationUpdateResponseStatusQueued    AppInvocationUpdateResponseStatus = "queued"
	AppInvocationUpdateResponseStatusRunning   AppInvocationUpdateResponseStatus = "running"
	AppInvocationUpdateResponseStatusSucceeded AppInvocationUpdateResponseStatus = "succeeded"
	AppInvocationUpdateResponseStatusFailed    AppInvocationUpdateResponseStatus = "failed"
)

type AppInvocationNewParams struct {
	// Name of the action to invoke
	ActionName string `json:"action_name,required"`
	// Name of the application
	AppName string `json:"app_name,required"`
	// Version of the application
	Version string `json:"version,required"`
	// If true, invoke asynchronously. When set, the API responds 202 Accepted with
	// status "queued".
	Async param.Opt[bool] `json:"async,omitzero"`
	// Input data for the action, sent as a JSON string.
	Payload param.Opt[string] `json:"payload,omitzero"`
	paramObj
}

func (r AppInvocationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppInvocationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppInvocationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInvocationUpdateParams struct {
	// New status for the invocation.
	//
	// Any of "succeeded", "failed".
	Status AppInvocationUpdateParamsStatus `json:"status,omitzero,required"`
	// Updated output of the invocation rendered as JSON string.
	Output param.Opt[string] `json:"output,omitzero"`
	paramObj
}

func (r AppInvocationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppInvocationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppInvocationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New status for the invocation.
type AppInvocationUpdateParamsStatus string

const (
	AppInvocationUpdateParamsStatusSucceeded AppInvocationUpdateParamsStatus = "succeeded"
	AppInvocationUpdateParamsStatusFailed    AppInvocationUpdateParamsStatus = "failed"
)
