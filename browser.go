// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"context"
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
func (r *BrowserService) NewSession(ctx context.Context, body BrowserNewSessionParams, opts ...option.RequestOption) (res *BrowserNewSessionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BrowserNewSessionResponse struct {
	// Websocket URL for Chrome DevTools Protocol connections to the browser session
	CdpWsURL string `json:"cdp_ws_url,required"`
	// Remote URL for live viewing the browser session
	RemoteURL string `json:"remote_url,required"`
	// Unique identifier for the browser session
	SessionID string `json:"sessionId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpWsURL    respjson.Field
		RemoteURL   respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserNewSessionResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserNewSessionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserNewSessionParams struct {
	// Kernel App invocation ID
	InvocationID string `json:"invocationId,required"`
	paramObj
}

func (r BrowserNewSessionParams) MarshalJSON() (data []byte, err error) {
	type shadow BrowserNewSessionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserNewSessionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
