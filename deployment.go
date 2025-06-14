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

// DeploymentService contains methods and other services that help with interacting
// with the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeploymentService] method instead.
type DeploymentService struct {
	Options []option.RequestOption
}

// NewDeploymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeploymentService(opts ...option.RequestOption) (r DeploymentService) {
	r = DeploymentService{}
	r.Options = opts
	return
}

// Create a new deployment.
func (r *DeploymentService) New(ctx context.Context, body DeploymentNewParams, opts ...option.RequestOption) (res *DeploymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a deployment's status.
func (r *DeploymentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DeploymentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Establishes a Server-Sent Events (SSE) stream that delivers real-time logs and
// status updates for a deployment. The stream terminates automatically once the
// deployment reaches a terminal state.
func (r *DeploymentService) FollowStreaming(ctx context.Context, id string, opts ...option.RequestOption) (stream *ssestream.Stream[DeploymentFollowResponseUnion]) {
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
	path := fmt.Sprintf("deployments/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[DeploymentFollowResponseUnion](ssestream.NewDecoder(raw), err)
}

// Deployment record information.
type DeploymentNewResponse struct {
	// Unique identifier for the deployment
	ID string `json:"id,required"`
	// Timestamp when the deployment was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deployment region code
	Region string `json:"region,required"`
	// Current status of the deployment
	//
	// Any of "queued", "in_progress", "running", "failed", "stopped".
	Status DeploymentNewResponseStatus `json:"status,required"`
	// Relative path to the application entrypoint
	EntrypointRelPath string `json:"entrypoint_rel_path"`
	// Environment variables configured for this deployment
	EnvVars map[string]string `json:"env_vars"`
	// Status reason
	StatusReason string `json:"status_reason"`
	// Timestamp when the deployment was last updated
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Region            respjson.Field
		Status            respjson.Field
		EntrypointRelPath respjson.Field
		EnvVars           respjson.Field
		StatusReason      respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DeploymentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the deployment
type DeploymentNewResponseStatus string

const (
	DeploymentNewResponseStatusQueued     DeploymentNewResponseStatus = "queued"
	DeploymentNewResponseStatusInProgress DeploymentNewResponseStatus = "in_progress"
	DeploymentNewResponseStatusRunning    DeploymentNewResponseStatus = "running"
	DeploymentNewResponseStatusFailed     DeploymentNewResponseStatus = "failed"
	DeploymentNewResponseStatusStopped    DeploymentNewResponseStatus = "stopped"
)

// Deployment record information.
type DeploymentGetResponse struct {
	// Unique identifier for the deployment
	ID string `json:"id,required"`
	// Timestamp when the deployment was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deployment region code
	Region string `json:"region,required"`
	// Current status of the deployment
	//
	// Any of "queued", "in_progress", "running", "failed", "stopped".
	Status DeploymentGetResponseStatus `json:"status,required"`
	// Relative path to the application entrypoint
	EntrypointRelPath string `json:"entrypoint_rel_path"`
	// Environment variables configured for this deployment
	EnvVars map[string]string `json:"env_vars"`
	// Status reason
	StatusReason string `json:"status_reason"`
	// Timestamp when the deployment was last updated
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Region            respjson.Field
		Status            respjson.Field
		EntrypointRelPath respjson.Field
		EnvVars           respjson.Field
		StatusReason      respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DeploymentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the deployment
type DeploymentGetResponseStatus string

const (
	DeploymentGetResponseStatusQueued     DeploymentGetResponseStatus = "queued"
	DeploymentGetResponseStatusInProgress DeploymentGetResponseStatus = "in_progress"
	DeploymentGetResponseStatusRunning    DeploymentGetResponseStatus = "running"
	DeploymentGetResponseStatusFailed     DeploymentGetResponseStatus = "failed"
	DeploymentGetResponseStatusStopped    DeploymentGetResponseStatus = "stopped"
)

// DeploymentFollowResponseUnion contains all possible properties and values from
// [DeploymentFollowResponseLog], [DeploymentFollowResponseDeploymentState],
// [DeploymentFollowResponseAppVersionSummaryEvent],
// [DeploymentFollowResponseErrorEvent].
//
// Use the [DeploymentFollowResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DeploymentFollowResponseUnion struct {
	// Any of "log", "deployment_state", nil, nil.
	Event string `json:"event"`
	// This field is from variant [DeploymentFollowResponseLog].
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [DeploymentFollowResponseDeploymentState].
	Deployment DeploymentFollowResponseDeploymentStateDeployment `json:"deployment"`
	// This field is from variant [DeploymentFollowResponseAppVersionSummaryEvent].
	ID string `json:"id"`
	// This field is from variant [DeploymentFollowResponseAppVersionSummaryEvent].
	AppName string `json:"app_name"`
	// This field is from variant [DeploymentFollowResponseAppVersionSummaryEvent].
	Region string `json:"region"`
	// This field is from variant [DeploymentFollowResponseAppVersionSummaryEvent].
	Version string `json:"version"`
	// This field is from variant [DeploymentFollowResponseAppVersionSummaryEvent].
	EnvVars map[string]string `json:"env_vars"`
	// This field is from variant [DeploymentFollowResponseErrorEvent].
	Error DeploymentFollowResponseErrorEventError `json:"error"`
	JSON  struct {
		Event      respjson.Field
		Message    respjson.Field
		Timestamp  respjson.Field
		Deployment respjson.Field
		ID         respjson.Field
		AppName    respjson.Field
		Region     respjson.Field
		Version    respjson.Field
		EnvVars    respjson.Field
		Error      respjson.Field
		raw        string
	} `json:"-"`
}

func (u DeploymentFollowResponseUnion) AsLog() (v DeploymentFollowResponseLog) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeploymentFollowResponseUnion) AsDeploymentState() (v DeploymentFollowResponseDeploymentState) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeploymentFollowResponseUnion) AsDeploymentFollowResponseAppVersionSummaryEvent() (v DeploymentFollowResponseAppVersionSummaryEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeploymentFollowResponseUnion) AsDeploymentFollowResponseErrorEvent() (v DeploymentFollowResponseErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DeploymentFollowResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *DeploymentFollowResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A log entry from the application.
type DeploymentFollowResponseLog struct {
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
func (r DeploymentFollowResponseLog) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An event representing the current state of a deployment.
type DeploymentFollowResponseDeploymentState struct {
	// Deployment record information.
	Deployment DeploymentFollowResponseDeploymentStateDeployment `json:"deployment,required"`
	// Event type identifier (always "deployment_state").
	Event constant.DeploymentState `json:"event,required"`
	// Time the state was reported.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deployment  respjson.Field
		Event       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentFollowResponseDeploymentState) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseDeploymentState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deployment record information.
type DeploymentFollowResponseDeploymentStateDeployment struct {
	// Unique identifier for the deployment
	ID string `json:"id,required"`
	// Timestamp when the deployment was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deployment region code
	Region string `json:"region,required"`
	// Current status of the deployment
	//
	// Any of "queued", "in_progress", "running", "failed", "stopped".
	Status string `json:"status,required"`
	// Relative path to the application entrypoint
	EntrypointRelPath string `json:"entrypoint_rel_path"`
	// Environment variables configured for this deployment
	EnvVars map[string]string `json:"env_vars"`
	// Status reason
	StatusReason string `json:"status_reason"`
	// Timestamp when the deployment was last updated
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Region            respjson.Field
		Status            respjson.Field
		EntrypointRelPath respjson.Field
		EnvVars           respjson.Field
		StatusReason      respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentFollowResponseDeploymentStateDeployment) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseDeploymentStateDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of an application version.
type DeploymentFollowResponseAppVersionSummaryEvent struct {
	// Unique identifier for the app version
	ID string `json:"id,required"`
	// Name of the application
	AppName string `json:"app_name,required"`
	// Event type identifier (always "app_version_summary").
	Event constant.AppVersionSummary `json:"event,required"`
	// Deployment region code
	Region string `json:"region,required"`
	// Time the state was reported.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Version label for the application
	Version string `json:"version,required"`
	// Environment variables configured for this app version
	EnvVars map[string]string `json:"env_vars"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppName     respjson.Field
		Event       respjson.Field
		Region      respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		EnvVars     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentFollowResponseAppVersionSummaryEvent) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseAppVersionSummaryEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error event from the application.
type DeploymentFollowResponseErrorEvent struct {
	Error DeploymentFollowResponseErrorEventError `json:"error,required"`
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
func (r DeploymentFollowResponseErrorEvent) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseErrorEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentFollowResponseErrorEventError struct {
	// Application-specific error code (machine-readable)
	Code string `json:"code,required"`
	// Human-readable error description for debugging
	Message string `json:"message,required"`
	// Additional error details (for multiple errors)
	Details    []DeploymentFollowResponseErrorEventErrorDetail   `json:"details"`
	InnerError DeploymentFollowResponseErrorEventErrorInnerError `json:"inner_error"`
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
func (r DeploymentFollowResponseErrorEventError) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseErrorEventError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentFollowResponseErrorEventErrorDetail struct {
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
func (r DeploymentFollowResponseErrorEventErrorDetail) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseErrorEventErrorDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentFollowResponseErrorEventErrorInnerError struct {
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
func (r DeploymentFollowResponseErrorEventErrorInnerError) RawJSON() string { return r.JSON.raw }
func (r *DeploymentFollowResponseErrorEventErrorInnerError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentNewParams struct {
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
	Region DeploymentNewParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r DeploymentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
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
type DeploymentNewParamsRegion string

const (
	DeploymentNewParamsRegionAwsUsEast1a DeploymentNewParamsRegion = "aws.us-east-1a"
)
