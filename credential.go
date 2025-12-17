// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/apiquery"
	shimjson "github.com/onkernel/kernel-go-sdk/internal/encoding/json"
	"github.com/onkernel/kernel-go-sdk/internal/requestconfig"
	"github.com/onkernel/kernel-go-sdk/option"
	"github.com/onkernel/kernel-go-sdk/packages/pagination"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
)

// CredentialService contains methods and other services that help with interacting
// with the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialService] method instead.
type CredentialService struct {
	Options []option.RequestOption
}

// NewCredentialService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCredentialService(opts ...option.RequestOption) (r CredentialService) {
	r = CredentialService{}
	r.Options = opts
	return
}

// Create a new credential for storing login information. Values are encrypted at
// rest.
func (r *CredentialService) New(ctx context.Context, body CredentialNewParams, opts ...option.RequestOption) (res *Credential, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a credential by its ID. Credential values are not returned.
func (r *CredentialService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Credential, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credentials/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a credential's name or values. Values are encrypted at rest.
func (r *CredentialService) Update(ctx context.Context, id string, body CredentialUpdateParams, opts ...option.RequestOption) (res *Credential, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credentials/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List credentials owned by the caller's organization. Credential values are not
// returned.
func (r *CredentialService) List(ctx context.Context, query CredentialListParams, opts ...option.RequestOption) (res *pagination.OffsetPagination[Credential], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credentials"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List credentials owned by the caller's organization. Credential values are not
// returned.
func (r *CredentialService) ListAutoPaging(ctx context.Context, query CredentialListParams, opts ...option.RequestOption) *pagination.OffsetPaginationAutoPager[Credential] {
	return pagination.NewOffsetPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a credential by its ID.
func (r *CredentialService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credentials/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Request to create a new credential
//
// The properties Domain, Name, Values are required.
type CreateCredentialRequestParam struct {
	// Target domain this credential is for
	Domain string `json:"domain,required"`
	// Unique name for the credential within the organization
	Name string `json:"name,required"`
	// Field name to value mapping (e.g., username, password)
	Values map[string]string `json:"values,omitzero,required"`
	paramObj
}

func (r CreateCredentialRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCredentialRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCredentialRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A stored credential for automatic re-authentication
type Credential struct {
	// Unique identifier for the credential
	ID string `json:"id,required"`
	// When the credential was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Target domain this credential is for
	Domain string `json:"domain,required"`
	// Unique name for the credential within the organization
	Name string `json:"name,required"`
	// When the credential was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Domain      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Credential) RawJSON() string { return r.JSON.raw }
func (r *Credential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to update an existing credential
type UpdateCredentialRequestParam struct {
	// New name for the credential
	Name param.Opt[string] `json:"name,omitzero"`
	// Field name to value mapping (e.g., username, password). Replaces all existing
	// values.
	Values map[string]string `json:"values,omitzero"`
	paramObj
}

func (r UpdateCredentialRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCredentialRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCredentialRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialNewParams struct {
	// Request to create a new credential
	CreateCredentialRequest CreateCredentialRequestParam
	paramObj
}

func (r CredentialNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCredentialRequest)
}
func (r *CredentialNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateCredentialRequest)
}

type CredentialUpdateParams struct {
	// Request to update an existing credential
	UpdateCredentialRequest UpdateCredentialRequestParam
	paramObj
}

func (r CredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCredentialRequest)
}
func (r *CredentialUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateCredentialRequest)
}

type CredentialListParams struct {
	// Filter by domain
	Domain param.Opt[string] `query:"domain,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of results to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CredentialListParams]'s query parameters as `url.Values`.
func (r CredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
