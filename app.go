// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/onkernel/kernel-go-sdk/internal/apiform"
	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/requestconfig"
	"github.com/onkernel/kernel-go-sdk/option"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
)

// AppService contains methods and other services that help with interacting with
// the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.Options = opts
	return
}

// Deploy a new application
func (r *AppService) Deploy(ctx context.Context, body AppDeployParams, opts ...option.RequestOption) (res *AppDeployResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "apps/deploy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Invoke an application
func (r *AppService) Invoke(ctx context.Context, body AppInvokeParams, opts ...option.RequestOption) (res *AppInvokeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "apps/invoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an app invocation by id
func (r *AppService) GetInvocation(ctx context.Context, id string, opts ...option.RequestOption) (res *AppGetInvocationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("apps/invocations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AppDeployResponse struct {
	Apps []AppDeployResponseApp `json:"apps,required"`
	// Success message
	Message string `json:"message,required"`
	// Status of the deployment
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Apps        respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppDeployResponse) RawJSON() string { return r.JSON.raw }
func (r *AppDeployResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppDeployResponseApp struct {
	// ID for the app version deployed
	ID      string                       `json:"id,required"`
	Actions []AppDeployResponseAppAction `json:"actions,required"`
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
func (r AppDeployResponseApp) RawJSON() string { return r.JSON.raw }
func (r *AppDeployResponseApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppDeployResponseAppAction struct {
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
func (r AppDeployResponseAppAction) RawJSON() string { return r.JSON.raw }
func (r *AppDeployResponseAppAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInvokeResponse struct {
	// ID of the invocation
	ID string `json:"id,required"`
	// Status of the invocation
	//
	// Any of "QUEUED", "RUNNING", "SUCCEEDED", "FAILED".
	Status AppInvokeResponseStatus `json:"status,required"`
	// Output from the invocation (if available)
	Output string `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInvokeResponse) RawJSON() string { return r.JSON.raw }
func (r *AppInvokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the invocation
type AppInvokeResponseStatus string

const (
	AppInvokeResponseStatusQueued    AppInvokeResponseStatus = "QUEUED"
	AppInvokeResponseStatusRunning   AppInvokeResponseStatus = "RUNNING"
	AppInvokeResponseStatusSucceeded AppInvokeResponseStatus = "SUCCEEDED"
	AppInvokeResponseStatusFailed    AppInvokeResponseStatus = "FAILED"
)

type AppGetInvocationResponse struct {
	ID         string `json:"id,required"`
	AppName    string `json:"appName,required"`
	FinishedAt string `json:"finishedAt,required"`
	Input      string `json:"input,required"`
	Output     string `json:"output,required"`
	StartedAt  string `json:"startedAt,required"`
	Status     string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppName     respjson.Field
		FinishedAt  respjson.Field
		Input       respjson.Field
		Output      respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppGetInvocationResponse) RawJSON() string { return r.JSON.raw }
func (r *AppGetInvocationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppDeployParams struct {
	// Relative path to the entrypoint of the application
	EntrypointRelPath string `json:"entrypointRelPath,required"`
	// ZIP file containing the application source directory
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Version of the application. Can be any string.
	Version param.Opt[string] `json:"version,omitzero"`
	// Allow overwriting an existing app version
	//
	// Any of "true", "false".
	Force AppDeployParamsForce `json:"force,omitzero"`
	// Region for deployment. Currently we only support "aws.us-east-1a"
	//
	// Any of "aws.us-east-1a".
	Region AppDeployParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r AppDeployParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Allow overwriting an existing app version
type AppDeployParamsForce string

const (
	AppDeployParamsForceTrue  AppDeployParamsForce = "true"
	AppDeployParamsForceFalse AppDeployParamsForce = "false"
)

// Region for deployment. Currently we only support "aws.us-east-1a"
type AppDeployParamsRegion string

const (
	AppDeployParamsRegionAwsUsEast1a AppDeployParamsRegion = "aws.us-east-1a"
)

type AppInvokeParams struct {
	// Name of the action to invoke
	ActionName string `json:"actionName,required"`
	// Name of the application
	AppName string `json:"appName,required"`
	// Input data for the application
	Payload any `json:"payload,omitzero,required"`
	// Version of the application
	Version string `json:"version,required"`
	paramObj
}

func (r AppInvokeParams) MarshalJSON() (data []byte, err error) {
	type shadow AppInvokeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppInvokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
