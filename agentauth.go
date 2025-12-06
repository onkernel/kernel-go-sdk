// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/requestconfig"
	"github.com/onkernel/kernel-go-sdk/option"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
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

// Creates a browser session and returns a handoff code for the hosted flow. Uses
// standard API key or JWT authentication (not the JWT returned by the exchange
// endpoint).
func (r *AgentAuthService) Start(ctx context.Context, body AgentAuthStartParams, opts ...option.RequestOption) (res *AgentAuthStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agents/auth/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// Response from starting an agent authentication invocation
type AgentAuthStartResponse struct {
	// Unique identifier for the auth agent managing this domain/profile
	AuthAgentID string `json:"auth_agent_id,required"`
	// When the handoff code expires
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// One-time code for handoff
	HandoffCode string `json:"handoff_code,required"`
	// URL to redirect user to
	HostedURL string `json:"hosted_url,required" format:"uri"`
	// Unique identifier for the invocation
	InvocationID string `json:"invocation_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthAgentID  respjson.Field
		ExpiresAt    respjson.Field
		HandoffCode  respjson.Field
		HostedURL    respjson.Field
		InvocationID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAuthStartResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentAuthStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Domain      respjson.Field
		ProfileName respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type AgentAuthStartParams struct {
	// Name of the profile to use for this flow
	ProfileName string `json:"profile_name,required"`
	// Target domain for authentication
	TargetDomain string `json:"target_domain,required"`
	// Optional logo URL for the application
	AppLogoURL param.Opt[string] `json:"app_logo_url,omitzero" format:"uri"`
	// Optional login page URL. If provided, will be stored on the agent and used to
	// skip Phase 1 discovery in future invocations.
	LoginURL param.Opt[string] `json:"login_url,omitzero" format:"uri"`
	// Optional proxy configuration
	Proxy AgentAuthStartParamsProxy `json:"proxy,omitzero"`
	paramObj
}

func (r AgentAuthStartParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentAuthStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAuthStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional proxy configuration
type AgentAuthStartParamsProxy struct {
	// ID of the proxy to use
	ProxyID param.Opt[string] `json:"proxy_id,omitzero"`
	paramObj
}

func (r AgentAuthStartParamsProxy) MarshalJSON() (data []byte, err error) {
	type shadow AgentAuthStartParamsProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAuthStartParamsProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
