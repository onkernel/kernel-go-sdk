// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
	"net/http"
	"net/url"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/apiquery"
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
	Options     []option.RequestOption
	Deployments AppDeploymentService
	Invocations AppInvocationService
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.Options = opts
	r.Deployments = NewAppDeploymentService(opts...)
	r.Invocations = NewAppInvocationService(opts...)
	return
}

// List applications. Optionally filter by app name and/or version label.
func (r *AppService) List(ctx context.Context, query AppListParams, opts ...option.RequestOption) (res *[]AppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "apps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Summary of an application version.
type AppListResponse struct {
	// Unique identifier for the app version
	ID string `json:"id,required"`
	// Name of the application
	AppName string `json:"app_name,required"`
	// Deployment region code
	Region string `json:"region,required"`
	// Version label for the application
	Version string `json:"version,required"`
	// Environment variables configured for this app version
	EnvVars map[string]string `json:"env_vars"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppName     respjson.Field
		Region      respjson.Field
		Version     respjson.Field
		EnvVars     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppListResponse) RawJSON() string { return r.JSON.raw }
func (r *AppListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppListParams struct {
	// Filter results by application name.
	AppName param.Opt[string] `query:"app_name,omitzero" json:"-"`
	// Filter results by version label.
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppListParams]'s query parameters as `url.Values`.
func (r AppListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
