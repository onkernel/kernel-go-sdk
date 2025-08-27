// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/onkernel/kernel-go-sdk/internal/apijson"
	"github.com/onkernel/kernel-go-sdk/internal/apiquery"
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
	Replays BrowserReplayService
	Fs      BrowserFService
	Process BrowserProcessService
	Logs    BrowserLogService
}

// NewBrowserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrowserService(opts ...option.RequestOption) (r BrowserService) {
	r = BrowserService{}
	r.Options = opts
	r.Replays = NewBrowserReplayService(opts...)
	r.Fs = NewBrowserFService(opts...)
	r.Process = NewBrowserProcessService(opts...)
	r.Logs = NewBrowserLogService(opts...)
	return
}

// Create a new browser session from within an action.
func (r *BrowserService) New(ctx context.Context, body BrowserNewParams, opts ...option.RequestOption) (res *BrowserNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a browser session.
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

// List active browser sessions
func (r *BrowserService) List(ctx context.Context, opts ...option.RequestOption) (res *[]BrowserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a persistent browser session by its persistent_id.
func (r *BrowserService) Delete(ctx context.Context, body BrowserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Delete a browser session by ID
func (r *BrowserService) DeleteByID(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("browsers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Optional persistence configuration for the browser session.
type BrowserPersistence struct {
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
func (r BrowserPersistence) RawJSON() string { return r.JSON.raw }
func (r *BrowserPersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrowserPersistence to a BrowserPersistenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrowserPersistenceParam.Overrides()
func (r BrowserPersistence) ToParam() BrowserPersistenceParam {
	return param.Override[BrowserPersistenceParam](json.RawMessage(r.RawJSON()))
}

// Optional persistence configuration for the browser session.
//
// The property ID is required.
type BrowserPersistenceParam struct {
	// Unique identifier for the persistent browser session.
	ID string `json:"id,required"`
	paramObj
}

func (r BrowserPersistenceParam) MarshalJSON() (data []byte, err error) {
	type shadow BrowserPersistenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserPersistenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserNewResponse struct {
	// Websocket URL for Chrome DevTools Protocol connections to the browser session
	CdpWsURL string `json:"cdp_ws_url,required"`
	// When the browser session was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the browser session is running in headless mode.
	Headless bool `json:"headless,required"`
	// Unique identifier for the browser session
	SessionID string `json:"session_id,required"`
	// Whether the browser session is running in stealth mode.
	Stealth bool `json:"stealth,required"`
	// The number of seconds of inactivity before the browser session is terminated.
	TimeoutSeconds int64 `json:"timeout_seconds,required"`
	// Remote URL for live viewing the browser session. Only available for non-headless
	// browsers.
	BrowserLiveViewURL string `json:"browser_live_view_url"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistence `json:"persistence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpWsURL           respjson.Field
		CreatedAt          respjson.Field
		Headless           respjson.Field
		SessionID          respjson.Field
		Stealth            respjson.Field
		TimeoutSeconds     respjson.Field
		BrowserLiveViewURL respjson.Field
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

type BrowserGetResponse struct {
	// Websocket URL for Chrome DevTools Protocol connections to the browser session
	CdpWsURL string `json:"cdp_ws_url,required"`
	// When the browser session was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the browser session is running in headless mode.
	Headless bool `json:"headless,required"`
	// Unique identifier for the browser session
	SessionID string `json:"session_id,required"`
	// Whether the browser session is running in stealth mode.
	Stealth bool `json:"stealth,required"`
	// The number of seconds of inactivity before the browser session is terminated.
	TimeoutSeconds int64 `json:"timeout_seconds,required"`
	// Remote URL for live viewing the browser session. Only available for non-headless
	// browsers.
	BrowserLiveViewURL string `json:"browser_live_view_url"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistence `json:"persistence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpWsURL           respjson.Field
		CreatedAt          respjson.Field
		Headless           respjson.Field
		SessionID          respjson.Field
		Stealth            respjson.Field
		TimeoutSeconds     respjson.Field
		BrowserLiveViewURL respjson.Field
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

type BrowserListResponse struct {
	// Websocket URL for Chrome DevTools Protocol connections to the browser session
	CdpWsURL string `json:"cdp_ws_url,required"`
	// When the browser session was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Whether the browser session is running in headless mode.
	Headless bool `json:"headless,required"`
	// Unique identifier for the browser session
	SessionID string `json:"session_id,required"`
	// Whether the browser session is running in stealth mode.
	Stealth bool `json:"stealth,required"`
	// The number of seconds of inactivity before the browser session is terminated.
	TimeoutSeconds int64 `json:"timeout_seconds,required"`
	// Remote URL for live viewing the browser session. Only available for non-headless
	// browsers.
	BrowserLiveViewURL string `json:"browser_live_view_url"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistence `json:"persistence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpWsURL           respjson.Field
		CreatedAt          respjson.Field
		Headless           respjson.Field
		SessionID          respjson.Field
		Stealth            respjson.Field
		TimeoutSeconds     respjson.Field
		BrowserLiveViewURL respjson.Field
		Persistence        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserNewParams struct {
	// If true, launches the browser using a headless image (no VNC/GUI). Defaults to
	// false.
	Headless param.Opt[bool] `json:"headless,omitzero"`
	// action invocation ID
	InvocationID param.Opt[string] `json:"invocation_id,omitzero"`
	// If true, launches the browser in stealth mode to reduce detection by anti-bot
	// mechanisms.
	Stealth param.Opt[bool] `json:"stealth,omitzero"`
	// The number of seconds of inactivity before the browser session is terminated.
	// Only applicable to non-persistent browsers. Activity includes CDP connections
	// and live view connections. Defaults to 60 seconds.
	TimeoutSeconds param.Opt[int64] `json:"timeout_seconds,omitzero"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistenceParam `json:"persistence,omitzero"`
	paramObj
}

func (r BrowserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserDeleteParams struct {
	// Persistent browser identifier
	PersistentID string `query:"persistent_id,required" json:"-"`
	paramObj
}

// URLQuery serializes [BrowserDeleteParams]'s query parameters as `url.Values`.
func (r BrowserDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
