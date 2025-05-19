// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"bytes"
	"context"
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
