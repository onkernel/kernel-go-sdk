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
	"github.com/onkernel/kernel-go-sdk/shared/constant"
)

// AgentAuthService contains methods and other services that help with interacting
// with the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentAuthService] method instead.
type AgentAuthService struct {
	Options     []option.RequestOption
	Invocations AgentAuthInvocationService
}

// NewAgentAuthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentAuthService(opts ...option.RequestOption) (r AgentAuthService) {
	r = AgentAuthService{}
	r.Options = opts
	r.Invocations = NewAgentAuthInvocationService(opts...)
	return
}

// Creates a new auth agent for the specified domain and profile combination, or
// returns an existing one if it already exists. This is idempotent - calling with
// the same domain and profile will return the same agent. Does NOT start an
// invocation - use POST /agents/auth/invocations to start an auth flow.
func (r *AgentAuthService) New(ctx context.Context, body AgentAuthNewParams, opts ...option.RequestOption) (res *AuthAgent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agents/auth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an auth agent by its ID. Returns the current authentication status of
// the managed profile.
func (r *AgentAuthService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AuthAgent, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agents/auth/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List auth agents with optional filters for profile_name and target_domain.
func (r *AgentAuthService) List(ctx context.Context, query AgentAuthListParams, opts ...option.RequestOption) (res *pagination.OffsetPagination[AuthAgent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "agents/auth"
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

// List auth agents with optional filters for profile_name and target_domain.
func (r *AgentAuthService) ListAutoPaging(ctx context.Context, query AgentAuthListParams, opts ...option.RequestOption) *pagination.OffsetPaginationAutoPager[AuthAgent] {
	return pagination.NewOffsetPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an auth agent and terminates its workflow. This will:
//
// - Soft delete the auth agent record
// - Gracefully terminate the agent's Temporal workflow
// - Cancel any in-progress invocations
func (r *AgentAuthService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agents/auth/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Triggers automatic re-authentication for an auth agent using stored credentials.
// Requires the auth agent to have a linked credential, stored selectors, and
// login_url. Returns immediately with status indicating whether re-auth was
// started.
func (r *AgentAuthService) Reauth(ctx context.Context, id string, opts ...option.RequestOption) (res *ReauthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agents/auth/%s/reauth", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response from discover endpoint matching AuthBlueprint schema
type AgentAuthDiscoverResponse struct {
	// Whether discovery succeeded
	Success bool `json:"success,required"`
	// Error message if discovery failed
	ErrorMessage string `json:"error_message"`
	// Discovered form fields (present when success is true)
	Fields []DiscoveredField `json:"fields"`
	// Whether user is already logged in
	LoggedIn bool `json:"logged_in"`
	// URL of the discovered login page
	LoginURL string `json:"login_url" format:"uri"`
	// Title of the login page
	PageTitle string `json:"page_title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success      respjson.Field
		ErrorMessage respjson.Field
		Fields       respjson.Field
		LoggedIn     respjson.Field
		LoginURL     respjson.Field
		PageTitle    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAuthDiscoverResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentAuthDiscoverResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from get invocation endpoint
type AgentAuthInvocationResponse struct {
	// App name (org name at time of invocation creation)
	AppName string `json:"app_name,required"`
	// When the handoff code expires
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Invocation status
	//
	// Any of "IN_PROGRESS", "SUCCESS", "EXPIRED", "CANCELED".
	Status AgentAuthInvocationResponseStatus `json:"status,required"`
	// Target domain for authentication
	TargetDomain string `json:"target_domain,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppName      respjson.Field
		ExpiresAt    respjson.Field
		Status       respjson.Field
		TargetDomain respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAuthInvocationResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentAuthInvocationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invocation status
type AgentAuthInvocationResponseStatus string

const (
	AgentAuthInvocationResponseStatusInProgress AgentAuthInvocationResponseStatus = "IN_PROGRESS"
	AgentAuthInvocationResponseStatusSuccess    AgentAuthInvocationResponseStatus = "SUCCESS"
	AgentAuthInvocationResponseStatusExpired    AgentAuthInvocationResponseStatus = "EXPIRED"
	AgentAuthInvocationResponseStatusCanceled   AgentAuthInvocationResponseStatus = "CANCELED"
)

// Response from submit endpoint matching SubmitResult schema
type AgentAuthSubmitResponse struct {
	// Whether submission succeeded
	Success bool `json:"success,required"`
	// Additional fields needed (e.g., OTP) - present when needs_additional_auth is
	// true
	AdditionalFields []DiscoveredField `json:"additional_fields"`
	// App name (only present when logged_in is true)
	AppName string `json:"app_name"`
	// Error message if submission failed
	ErrorMessage string `json:"error_message"`
	// Whether user is now logged in
	LoggedIn bool `json:"logged_in"`
	// Whether additional authentication fields are needed
	NeedsAdditionalAuth bool `json:"needs_additional_auth"`
	// Target domain (only present when logged_in is true)
	TargetDomain string `json:"target_domain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success             respjson.Field
		AdditionalFields    respjson.Field
		AppName             respjson.Field
		ErrorMessage        respjson.Field
		LoggedIn            respjson.Field
		NeedsAdditionalAuth respjson.Field
		TargetDomain        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAuthSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentAuthSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An auth agent that manages authentication for a specific domain and profile
// combination
type AuthAgent struct {
	// Unique identifier for the auth agent
	ID string `json:"id,required"`
	// Target domain for authentication
	Domain string `json:"domain,required"`
	// Name of the profile associated with this auth agent
	ProfileName string `json:"profile_name,required"`
	// Current authentication status of the managed profile
	//
	// Any of "AUTHENTICATED", "NEEDS_AUTH".
	Status AuthAgentStatus `json:"status,required"`
	// Whether automatic re-authentication is possible (has credential_id, selectors,
	// and login_url)
	CanReauth bool `json:"can_reauth"`
	// ID of the linked credential for automatic re-authentication
	CredentialID string `json:"credential_id"`
	// Name of the linked credential for automatic re-authentication
	CredentialName string `json:"credential_name"`
	// Whether this auth agent has stored selectors for deterministic re-authentication
	HasSelectors bool `json:"has_selectors"`
	// When the last authentication check was performed
	LastAuthCheckAt time.Time `json:"last_auth_check_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Domain          respjson.Field
		ProfileName     respjson.Field
		Status          respjson.Field
		CanReauth       respjson.Field
		CredentialID    respjson.Field
		CredentialName  respjson.Field
		HasSelectors    respjson.Field
		LastAuthCheckAt respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthAgent) RawJSON() string { return r.JSON.raw }
func (r *AuthAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current authentication status of the managed profile
type AuthAgentStatus string

const (
	AuthAgentStatusAuthenticated AuthAgentStatus = "AUTHENTICATED"
	AuthAgentStatusNeedsAuth     AuthAgentStatus = "NEEDS_AUTH"
)

// Request to create or find an auth agent
//
// The properties ProfileName, TargetDomain are required.
type AuthAgentCreateRequestParam struct {
	// Name of the profile to use for this auth agent
	ProfileName string `json:"profile_name,required"`
	// Target domain for authentication
	TargetDomain string `json:"target_domain,required"`
	// Optional name of an existing credential to use for this auth agent. If provided,
	// the credential will be linked to the agent and its values will be used to
	// auto-fill the login form on invocation.
	CredentialName param.Opt[string] `json:"credential_name,omitzero"`
	// Optional login page URL. If provided, will be stored on the agent and used to
	// skip discovery in future invocations.
	LoginURL param.Opt[string] `json:"login_url,omitzero" format:"uri"`
	// Optional proxy configuration
	Proxy AuthAgentCreateRequestProxyParam `json:"proxy,omitzero"`
	paramObj
}

func (r AuthAgentCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthAgentCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthAgentCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional proxy configuration
type AuthAgentCreateRequestProxyParam struct {
	// ID of the proxy to use
	ProxyID param.Opt[string] `json:"proxy_id,omitzero"`
	paramObj
}

func (r AuthAgentCreateRequestProxyParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthAgentCreateRequestProxyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthAgentCreateRequestProxyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to create an invocation for an existing auth agent
//
// The property AuthAgentID is required.
type AuthAgentInvocationCreateRequestParam struct {
	// ID of the auth agent to create an invocation for
	AuthAgentID string `json:"auth_agent_id,required"`
	// If provided, saves the submitted credentials under this name upon successful
	// login. The credential will be linked to the auth agent for automatic
	// re-authentication.
	SaveCredentialAs param.Opt[string] `json:"save_credential_as,omitzero"`
	paramObj
}

func (r AuthAgentInvocationCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthAgentInvocationCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthAgentInvocationCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AuthAgentInvocationCreateResponseUnion contains all possible properties and
// values from [AuthAgentInvocationCreateResponseAlreadyAuthenticated],
// [AuthAgentInvocationCreateResponseInvocationCreated].
//
// Use the [AuthAgentInvocationCreateResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AuthAgentInvocationCreateResponseUnion struct {
	// Any of "already_authenticated", "invocation_created".
	Status string `json:"status"`
	// This field is from variant [AuthAgentInvocationCreateResponseInvocationCreated].
	ExpiresAt time.Time `json:"expires_at"`
	// This field is from variant [AuthAgentInvocationCreateResponseInvocationCreated].
	HandoffCode string `json:"handoff_code"`
	// This field is from variant [AuthAgentInvocationCreateResponseInvocationCreated].
	HostedURL string `json:"hosted_url"`
	// This field is from variant [AuthAgentInvocationCreateResponseInvocationCreated].
	InvocationID string `json:"invocation_id"`
	JSON         struct {
		Status       respjson.Field
		ExpiresAt    respjson.Field
		HandoffCode  respjson.Field
		HostedURL    respjson.Field
		InvocationID respjson.Field
		raw          string
	} `json:"-"`
}

// anyAuthAgentInvocationCreateResponse is implemented by each variant of
// [AuthAgentInvocationCreateResponseUnion] to add type safety for the return type
// of [AuthAgentInvocationCreateResponseUnion.AsAny]
type anyAuthAgentInvocationCreateResponse interface {
	implAuthAgentInvocationCreateResponseUnion()
}

func (AuthAgentInvocationCreateResponseAlreadyAuthenticated) implAuthAgentInvocationCreateResponseUnion() {
}
func (AuthAgentInvocationCreateResponseInvocationCreated) implAuthAgentInvocationCreateResponseUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := AuthAgentInvocationCreateResponseUnion.AsAny().(type) {
//	case kernel.AuthAgentInvocationCreateResponseAlreadyAuthenticated:
//	case kernel.AuthAgentInvocationCreateResponseInvocationCreated:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AuthAgentInvocationCreateResponseUnion) AsAny() anyAuthAgentInvocationCreateResponse {
	switch u.Status {
	case "already_authenticated":
		return u.AsAlreadyAuthenticated()
	case "invocation_created":
		return u.AsInvocationCreated()
	}
	return nil
}

func (u AuthAgentInvocationCreateResponseUnion) AsAlreadyAuthenticated() (v AuthAgentInvocationCreateResponseAlreadyAuthenticated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuthAgentInvocationCreateResponseUnion) AsInvocationCreated() (v AuthAgentInvocationCreateResponseInvocationCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AuthAgentInvocationCreateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AuthAgentInvocationCreateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when the agent is already authenticated.
type AuthAgentInvocationCreateResponseAlreadyAuthenticated struct {
	// Indicates the agent is already authenticated and no invocation was created.
	Status constant.AlreadyAuthenticated `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthAgentInvocationCreateResponseAlreadyAuthenticated) RawJSON() string { return r.JSON.raw }
func (r *AuthAgentInvocationCreateResponseAlreadyAuthenticated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when a new invocation was created.
type AuthAgentInvocationCreateResponseInvocationCreated struct {
	// When the handoff code expires.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// One-time code for handoff.
	HandoffCode string `json:"handoff_code,required"`
	// URL to redirect user to.
	HostedURL string `json:"hosted_url,required" format:"uri"`
	// Unique identifier for the invocation.
	InvocationID string `json:"invocation_id,required"`
	// Indicates an invocation was created.
	Status constant.InvocationCreated `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt    respjson.Field
		HandoffCode  respjson.Field
		HostedURL    respjson.Field
		InvocationID respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthAgentInvocationCreateResponseInvocationCreated) RawJSON() string { return r.JSON.raw }
func (r *AuthAgentInvocationCreateResponseInvocationCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A discovered form field
type DiscoveredField struct {
	// Field label
	Label string `json:"label,required"`
	// Field name
	Name string `json:"name,required"`
	// CSS selector for the field
	Selector string `json:"selector,required"`
	// Field type
	//
	// Any of "text", "email", "password", "tel", "number", "url", "code".
	Type DiscoveredFieldType `json:"type,required"`
	// Field placeholder
	Placeholder string `json:"placeholder"`
	// Whether field is required
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Name        respjson.Field
		Selector    respjson.Field
		Type        respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DiscoveredField) RawJSON() string { return r.JSON.raw }
func (r *DiscoveredField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Field type
type DiscoveredFieldType string

const (
	DiscoveredFieldTypeText     DiscoveredFieldType = "text"
	DiscoveredFieldTypeEmail    DiscoveredFieldType = "email"
	DiscoveredFieldTypePassword DiscoveredFieldType = "password"
	DiscoveredFieldTypeTel      DiscoveredFieldType = "tel"
	DiscoveredFieldTypeNumber   DiscoveredFieldType = "number"
	DiscoveredFieldTypeURL      DiscoveredFieldType = "url"
	DiscoveredFieldTypeCode     DiscoveredFieldType = "code"
)

// Response from triggering re-authentication
type ReauthResponse struct {
	// Result of the re-authentication attempt
	//
	// Any of "reauth_started", "already_authenticated", "cannot_reauth".
	Status ReauthResponseStatus `json:"status,required"`
	// ID of the re-auth invocation if one was created
	InvocationID string `json:"invocation_id"`
	// Human-readable description of the result
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status       respjson.Field
		InvocationID respjson.Field
		Message      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReauthResponse) RawJSON() string { return r.JSON.raw }
func (r *ReauthResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of the re-authentication attempt
type ReauthResponseStatus string

const (
	ReauthResponseStatusReauthStarted        ReauthResponseStatus = "reauth_started"
	ReauthResponseStatusAlreadyAuthenticated ReauthResponseStatus = "already_authenticated"
	ReauthResponseStatusCannotReauth         ReauthResponseStatus = "cannot_reauth"
)

type AgentAuthNewParams struct {
	// Request to create or find an auth agent
	AuthAgentCreateRequest AuthAgentCreateRequestParam
	paramObj
}

func (r AgentAuthNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AuthAgentCreateRequest)
}
func (r *AgentAuthNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AuthAgentCreateRequest)
}

type AgentAuthListParams struct {
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of results to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by profile name
	ProfileName param.Opt[string] `query:"profile_name,omitzero" json:"-"`
	// Filter by target domain
	TargetDomain param.Opt[string] `query:"target_domain,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentAuthListParams]'s query parameters as `url.Values`.
func (r AgentAuthListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
