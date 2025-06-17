// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
	"github.com/onkernel/kernel-go-sdk/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error struct {
	// Application-specific error code (machine-readable)
	Code string `json:"code,required"`
	// Human-readable error description for debugging
	Message string `json:"message,required"`
	// Additional error details (for multiple errors)
	Details    []ErrorDetail `json:"details"`
	InnerError ErrorDetail   `json:"inner_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		InnerError  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ErrorDetail struct {
	// Lower-level error code providing more specific detail
	Code string `json:"code"`
	// Further detail about the error
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorDetail) RawJSON() string { return r.JSON.raw }
func (r *ErrorDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error event from the application.
type ErrorEvent struct {
	Error Error `json:"error,required"`
	// Event type identifier (always "error").
	Event constant.Error `json:"event,required"`
	// Time the error occurred.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Event       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorEvent) RawJSON() string { return r.JSON.raw }
func (r *ErrorEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A log entry from the application.
type LogEvent struct {
	// Event type identifier (always "log").
	Event constant.Log `json:"event,required"`
	// Log message text.
	Message string `json:"message,required"`
	// Time the log entry was produced.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
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
func (r LogEvent) RawJSON() string { return r.JSON.raw }
func (r *LogEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (LogEvent) ImplAppDeploymentFollowResponseUnion() {}
