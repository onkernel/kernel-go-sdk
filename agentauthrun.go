// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/requestconfig"
	"github.com/onkernel/kernel-go-sdk/option"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
)

// AgentAuthRunService contains methods and other services that help with
// interacting with the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentAuthRunService] method instead.
type AgentAuthRunService struct {
	Options []option.RequestOption
}

// NewAgentAuthRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentAuthRunService(opts ...option.RequestOption) (r AgentAuthRunService) {
	r = AgentAuthRunService{}
	r.Options = opts
	return
}

// Returns run details including app_name and target_domain. Uses the JWT returned
// by the exchange endpoint, or standard API key or JWT authentication.
func (r *AgentAuthRunService) Get(ctx context.Context, runID string, opts ...option.RequestOption) (res *AgentAuthRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("agents/auth/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Inspects the target site to detect logged-in state or discover required fields.
// Returns 200 with success: true when fields are found, or 4xx/5xx for failures.
// Requires the JWT returned by the exchange endpoint.
func (r *AgentAuthRunService) Discover(ctx context.Context, runID string, body AgentAuthRunDiscoverParams, opts ...option.RequestOption) (res *AgentAuthDiscoverResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("agents/auth/runs/%s/discover", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Validates the handoff code and returns a JWT token for subsequent requests. No
// authentication required (the handoff code serves as the credential).
func (r *AgentAuthRunService) Exchange(ctx context.Context, runID string, body AgentAuthRunExchangeParams, opts ...option.RequestOption) (res *AgentAuthRunExchangeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("agents/auth/runs/%s/exchange", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submits field values for the discovered login form and may return additional
// auth fields or success. Requires the JWT returned by the exchange endpoint.
func (r *AgentAuthRunService) Submit(ctx context.Context, runID string, body AgentAuthRunSubmitParams, opts ...option.RequestOption) (res *AgentAuthSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("agents/auth/runs/%s/submit", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from exchange endpoint
type AgentAuthRunExchangeResponse struct {
	// JWT token with run_id claim (30 minute TTL)
	Jwt string `json:"jwt,required"`
	// Run ID
	RunID string `json:"run_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Jwt         respjson.Field
		RunID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAuthRunExchangeResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentAuthRunExchangeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAuthRunDiscoverParams struct {
	paramObj
}

func (r AgentAuthRunDiscoverParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentAuthRunDiscoverParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAuthRunDiscoverParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAuthRunExchangeParams struct {
	// Handoff code from start endpoint
	Code string `json:"code,required"`
	paramObj
}

func (r AgentAuthRunExchangeParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentAuthRunExchangeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAuthRunExchangeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAuthRunSubmitParams struct {
	// Values for the discovered login fields
	FieldValues map[string]string `json:"field_values,omitzero,required"`
	paramObj
}

func (r AgentAuthRunSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentAuthRunSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAuthRunSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
