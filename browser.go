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
	"net/url"
	"slices"
	"time"

	"github.com/onkernel/kernel-go-sdk/internal/apiform"
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
	Options    []option.RequestOption
	Replays    BrowserReplayService
	Fs         BrowserFService
	Process    BrowserProcessService
	Logs       BrowserLogService
	Computer   BrowserComputerService
	Playwright BrowserPlaywrightService
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
	r.Computer = NewBrowserComputerService(opts...)
	r.Playwright = NewBrowserPlaywrightService(opts...)
	return
}

// Create a new browser session from within an action.
func (r *BrowserService) New(ctx context.Context, body BrowserNewParams, opts ...option.RequestOption) (res *BrowserNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a browser session.
func (r *BrowserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BrowserGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	path := "browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a persistent browser session by its persistent_id.
func (r *BrowserService) Delete(ctx context.Context, body BrowserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Delete a browser session by ID
func (r *BrowserService) DeleteByID(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("browsers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Loads one or more unpacked extensions and restarts Chromium on the browser
// instance.
func (r *BrowserService) LoadExtensions(ctx context.Context, id string, body BrowserLoadExtensionsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("browsers/%s/extensions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
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

// Browser profile metadata.
type Profile struct {
	// Unique identifier for the profile
	ID string `json:"id,required"`
	// Timestamp when the profile was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp when the profile was last used
	LastUsedAt time.Time `json:"last_used_at" format:"date-time"`
	// Optional, easier-to-reference name for the profile
	Name string `json:"name,nullable"`
	// Timestamp when the profile was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		LastUsedAt  respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Profile) RawJSON() string { return r.JSON.raw }
func (r *Profile) UnmarshalJSON(data []byte) error {
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
	// Whether the browser session is running in kiosk mode.
	KioskMode bool `json:"kiosk_mode"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistence `json:"persistence"`
	// Browser profile metadata.
	Profile Profile `json:"profile"`
	// ID of the proxy associated with this browser session, if any.
	ProxyID string `json:"proxy_id"`
	// Initial browser window size in pixels with optional refresh rate. If omitted,
	// image defaults apply (commonly 1024x768@60). Only specific viewport
	// configurations are supported. The server will reject unsupported combinations.
	// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
	// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
	// automatically determined from the width and height if they match a supported
	// configuration exactly. Note: Higher resolutions may affect the responsiveness of
	// live view browser
	Viewport BrowserNewResponseViewport `json:"viewport"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpWsURL           respjson.Field
		CreatedAt          respjson.Field
		Headless           respjson.Field
		SessionID          respjson.Field
		Stealth            respjson.Field
		TimeoutSeconds     respjson.Field
		BrowserLiveViewURL respjson.Field
		KioskMode          respjson.Field
		Persistence        respjson.Field
		Profile            respjson.Field
		ProxyID            respjson.Field
		Viewport           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial browser window size in pixels with optional refresh rate. If omitted,
// image defaults apply (commonly 1024x768@60). Only specific viewport
// configurations are supported. The server will reject unsupported combinations.
// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
// automatically determined from the width and height if they match a supported
// configuration exactly. Note: Higher resolutions may affect the responsiveness of
// live view browser
type BrowserNewResponseViewport struct {
	// Browser window height in pixels.
	Height int64 `json:"height,required"`
	// Browser window width in pixels.
	Width int64 `json:"width,required"`
	// Display refresh rate in Hz. If omitted, automatically determined from width and
	// height.
	RefreshRate int64 `json:"refresh_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		RefreshRate respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserNewResponseViewport) RawJSON() string { return r.JSON.raw }
func (r *BrowserNewResponseViewport) UnmarshalJSON(data []byte) error {
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
	// Whether the browser session is running in kiosk mode.
	KioskMode bool `json:"kiosk_mode"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistence `json:"persistence"`
	// Browser profile metadata.
	Profile Profile `json:"profile"`
	// ID of the proxy associated with this browser session, if any.
	ProxyID string `json:"proxy_id"`
	// Initial browser window size in pixels with optional refresh rate. If omitted,
	// image defaults apply (commonly 1024x768@60). Only specific viewport
	// configurations are supported. The server will reject unsupported combinations.
	// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
	// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
	// automatically determined from the width and height if they match a supported
	// configuration exactly. Note: Higher resolutions may affect the responsiveness of
	// live view browser
	Viewport BrowserGetResponseViewport `json:"viewport"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpWsURL           respjson.Field
		CreatedAt          respjson.Field
		Headless           respjson.Field
		SessionID          respjson.Field
		Stealth            respjson.Field
		TimeoutSeconds     respjson.Field
		BrowserLiveViewURL respjson.Field
		KioskMode          respjson.Field
		Persistence        respjson.Field
		Profile            respjson.Field
		ProxyID            respjson.Field
		Viewport           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial browser window size in pixels with optional refresh rate. If omitted,
// image defaults apply (commonly 1024x768@60). Only specific viewport
// configurations are supported. The server will reject unsupported combinations.
// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
// automatically determined from the width and height if they match a supported
// configuration exactly. Note: Higher resolutions may affect the responsiveness of
// live view browser
type BrowserGetResponseViewport struct {
	// Browser window height in pixels.
	Height int64 `json:"height,required"`
	// Browser window width in pixels.
	Width int64 `json:"width,required"`
	// Display refresh rate in Hz. If omitted, automatically determined from width and
	// height.
	RefreshRate int64 `json:"refresh_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		RefreshRate respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserGetResponseViewport) RawJSON() string { return r.JSON.raw }
func (r *BrowserGetResponseViewport) UnmarshalJSON(data []byte) error {
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
	// Whether the browser session is running in kiosk mode.
	KioskMode bool `json:"kiosk_mode"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistence `json:"persistence"`
	// Browser profile metadata.
	Profile Profile `json:"profile"`
	// ID of the proxy associated with this browser session, if any.
	ProxyID string `json:"proxy_id"`
	// Initial browser window size in pixels with optional refresh rate. If omitted,
	// image defaults apply (commonly 1024x768@60). Only specific viewport
	// configurations are supported. The server will reject unsupported combinations.
	// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
	// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
	// automatically determined from the width and height if they match a supported
	// configuration exactly. Note: Higher resolutions may affect the responsiveness of
	// live view browser
	Viewport BrowserListResponseViewport `json:"viewport"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpWsURL           respjson.Field
		CreatedAt          respjson.Field
		Headless           respjson.Field
		SessionID          respjson.Field
		Stealth            respjson.Field
		TimeoutSeconds     respjson.Field
		BrowserLiveViewURL respjson.Field
		KioskMode          respjson.Field
		Persistence        respjson.Field
		Profile            respjson.Field
		ProxyID            respjson.Field
		Viewport           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial browser window size in pixels with optional refresh rate. If omitted,
// image defaults apply (commonly 1024x768@60). Only specific viewport
// configurations are supported. The server will reject unsupported combinations.
// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
// automatically determined from the width and height if they match a supported
// configuration exactly. Note: Higher resolutions may affect the responsiveness of
// live view browser
type BrowserListResponseViewport struct {
	// Browser window height in pixels.
	Height int64 `json:"height,required"`
	// Browser window width in pixels.
	Width int64 `json:"width,required"`
	// Display refresh rate in Hz. If omitted, automatically determined from width and
	// height.
	RefreshRate int64 `json:"refresh_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		RefreshRate respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserListResponseViewport) RawJSON() string { return r.JSON.raw }
func (r *BrowserListResponseViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserNewParams struct {
	// If true, launches the browser using a headless image (no VNC/GUI). Defaults to
	// false.
	Headless param.Opt[bool] `json:"headless,omitzero"`
	// action invocation ID
	InvocationID param.Opt[string] `json:"invocation_id,omitzero"`
	// If true, launches the browser in kiosk mode to hide address bar and tabs in live
	// view.
	KioskMode param.Opt[bool] `json:"kiosk_mode,omitzero"`
	// Optional proxy to associate to the browser session. Must reference a proxy
	// belonging to the caller's org.
	ProxyID param.Opt[string] `json:"proxy_id,omitzero"`
	// If true, launches the browser in stealth mode to reduce detection by anti-bot
	// mechanisms.
	Stealth param.Opt[bool] `json:"stealth,omitzero"`
	// The number of seconds of inactivity before the browser session is terminated.
	// Only applicable to non-persistent browsers. Activity includes CDP connections
	// and live view connections. Defaults to 60 seconds. Minimum allowed is 10
	// seconds. Maximum allowed is 86400 (24 hours). We check for inactivity every 5
	// seconds, so the actual timeout behavior you will see is +/- 5 seconds around the
	// specified value.
	TimeoutSeconds param.Opt[int64] `json:"timeout_seconds,omitzero"`
	// List of browser extensions to load into the session. Provide each by id or name.
	Extensions []BrowserNewParamsExtension `json:"extensions,omitzero"`
	// Optional persistence configuration for the browser session.
	Persistence BrowserPersistenceParam `json:"persistence,omitzero"`
	// Profile selection for the browser session. Provide either id or name. If
	// specified, the matching profile will be loaded into the browser session.
	// Profiles must be created beforehand.
	Profile BrowserNewParamsProfile `json:"profile,omitzero"`
	// Initial browser window size in pixels with optional refresh rate. If omitted,
	// image defaults apply (commonly 1024x768@60). Only specific viewport
	// configurations are supported. The server will reject unsupported combinations.
	// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
	// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
	// automatically determined from the width and height if they match a supported
	// configuration exactly. Note: Higher resolutions may affect the responsiveness of
	// live view browser
	Viewport BrowserNewParamsViewport `json:"viewport,omitzero"`
	paramObj
}

func (r BrowserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extension selection for the browser session. Provide either id or name of an
// extension uploaded to Kernel.
type BrowserNewParamsExtension struct {
	// Extension ID to load for this browser session
	ID param.Opt[string] `json:"id,omitzero"`
	// Extension name to load for this browser session (instead of id). Must be 1-255
	// characters, using letters, numbers, dots, underscores, or hyphens.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r BrowserNewParamsExtension) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewParamsExtension
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewParamsExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile selection for the browser session. Provide either id or name. If
// specified, the matching profile will be loaded into the browser session.
// Profiles must be created beforehand.
type BrowserNewParamsProfile struct {
	// Profile ID to load for this browser session
	ID param.Opt[string] `json:"id,omitzero"`
	// Profile name to load for this browser session (instead of id). Must be 1-255
	// characters, using letters, numbers, dots, underscores, or hyphens.
	Name param.Opt[string] `json:"name,omitzero"`
	// If true, save changes made during the session back to the profile when the
	// session ends.
	SaveChanges param.Opt[bool] `json:"save_changes,omitzero"`
	paramObj
}

func (r BrowserNewParamsProfile) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewParamsProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewParamsProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial browser window size in pixels with optional refresh rate. If omitted,
// image defaults apply (commonly 1024x768@60). Only specific viewport
// configurations are supported. The server will reject unsupported combinations.
// Supported resolutions are: 2560x1440@10, 1920x1080@25, 1920x1200@25,
// 1440x900@25, 1024x768@60 If refresh_rate is not provided, it will be
// automatically determined from the width and height if they match a supported
// configuration exactly. Note: Higher resolutions may affect the responsiveness of
// live view browser
//
// The properties Height, Width are required.
type BrowserNewParamsViewport struct {
	// Browser window height in pixels.
	Height int64 `json:"height,required"`
	// Browser window width in pixels.
	Width int64 `json:"width,required"`
	// Display refresh rate in Hz. If omitted, automatically determined from width and
	// height.
	RefreshRate param.Opt[int64] `json:"refresh_rate,omitzero"`
	paramObj
}

func (r BrowserNewParamsViewport) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewParamsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewParamsViewport) UnmarshalJSON(data []byte) error {
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

type BrowserLoadExtensionsParams struct {
	// List of extensions to upload and activate
	Extensions []BrowserLoadExtensionsParamsExtension `json:"extensions,omitzero,required"`
	paramObj
}

func (r BrowserLoadExtensionsParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// The properties Name, ZipFile are required.
type BrowserLoadExtensionsParamsExtension struct {
	// Folder name to place the extension under /home/kernel/extensions/<name>
	Name string `json:"name,required"`
	// Zip archive containing an unpacked Chromium extension (must include
	// manifest.json)
	ZipFile io.Reader `json:"zip_file,omitzero,required" format:"binary"`
	paramObj
}

func (r BrowserLoadExtensionsParamsExtension) MarshalJSON() (data []byte, err error) {
	type shadow BrowserLoadExtensionsParamsExtension
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserLoadExtensionsParamsExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
