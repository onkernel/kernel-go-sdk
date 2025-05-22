// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/requestconfig"
	"github.com/onkernel/kernel-go-sdk/option"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/packages/respjson"
)

// BrowserService contains methods and other services that help with interacting
// with the kernel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserService] method instead.
type BrowserService struct {
	Options []option.RequestOption
}

// NewBrowserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrowserService(opts ...option.RequestOption) (r BrowserService) {
	r = BrowserService{}
	r.Options = opts
	return
}

// Create Browser Session
func (r *BrowserService) New(ctx context.Context, body BrowserNewParams, opts ...option.RequestOption) (res *BrowserNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Browser Session by ID
func (r *BrowserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BrowserGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("browsers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BrowserNewResponse struct {
	// Remote URL for live viewing the browser session
	BrowserLiveViewURL string `json:"browser_live_view_url,required"`
	// Websocket URL for Chrome DevTools Protocol connections to the browser session
	CdpWsURL string `json:"cdp_ws_url,required"`
	// Unique identifier for the browser session
	SessionID string `json:"session_id,required"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserNewResponsePersistence `json:"persistence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrowserLiveViewURL respjson.Field
		CdpWsURL           respjson.Field
		SessionID          respjson.Field
		Persistence        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional persistence configuration for the browser session.
type BrowserNewResponsePersistence struct {
	// Unique identifier for the persistent browser session.
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserNewResponsePersistence) RawJSON() string { return r.JSON.raw }
func (r *BrowserNewResponsePersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserGetResponse struct {
	// Remote URL for live viewing the browser session
	BrowserLiveViewURL string `json:"browser_live_view_url,required"`
	// Websocket URL for Chrome DevTools Protocol connections to the browser session
	CdpWsURL string `json:"cdp_ws_url,required"`
	// Unique identifier for the browser session
	SessionID string `json:"session_id,required"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserGetResponsePersistence `json:"persistence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrowserLiveViewURL respjson.Field
		CdpWsURL           respjson.Field
		SessionID          respjson.Field
		Persistence        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional persistence configuration for the browser session.
type BrowserGetResponsePersistence struct {
	// Unique identifier for the persistent browser session.
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserGetResponsePersistence) RawJSON() string { return r.JSON.raw }
func (r *BrowserGetResponsePersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserNewParams struct {
	// action invocation ID
	InvocationID string `json:"invocation_id,required"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserNewParamsPersistence `json:"persistence,omitzero"`
	paramObj
}

func (r BrowserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional persistence configuration for the browser session.
//
// The property ID is required.
type BrowserNewParamsPersistence struct {
	// Unique identifier for the persistent browser session.
	ID string `json:"id,required"`
	paramObj
}

func (r BrowserNewParamsPersistence) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewParamsPersistence
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewParamsPersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
