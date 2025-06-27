// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel

import (
	"github.com/onkernel/kernel-go-sdk/internal/apierror"
	"github.com/onkernel/kernel-go-sdk/packages/param"
	"github.com/onkernel/kernel-go-sdk/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// An action available on the app
//
// This is an alias to an internal type.
type AppAction = shared.AppAction

// This is an alias to an internal type.
type ErrorDetail = shared.ErrorDetail

// An error event from the application.
//
// This is an alias to an internal type.
type ErrorEvent = shared.ErrorEvent

// This is an alias to an internal type.
type ErrorModel = shared.ErrorModel

// Heartbeat event sent periodically to keep SSE connection alive.
//
// This is an alias to an internal type.
type HeartbeatEvent = shared.HeartbeatEvent

// A log entry from the application.
//
// This is an alias to an internal type.
type LogEvent = shared.LogEvent
