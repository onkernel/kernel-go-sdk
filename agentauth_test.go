// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/onkernel/kernel-go-sdk"
	"github.com/onkernel/kernel-go-sdk/internal/testutil"
	"github.com/onkernel/kernel-go-sdk/option"
)

func TestAgentAuthGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := kernel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Auth.Get(context.TODO(), "id")
	if err != nil {
		var apierr *kernel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentAuthStartWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := kernel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Auth.Start(context.TODO(), kernel.AgentAuthStartParams{
		ProfileName:  "auth-abc123",
		TargetDomain: "doordash.com",
		AppLogoURL:   kernel.String("https://example.com/logo.png"),
		LoginURL:     kernel.String("https://doordash.com/account/login"),
		Proxy: kernel.AgentAuthStartParamsProxy{
			ProxyID: kernel.String("proxy_id"),
		},
	})
	if err != nil {
		var apierr *kernel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
