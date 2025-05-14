// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/kernel-go"
	"github.com/stainless-sdks/kernel-go/internal/testutil"
	"github.com/stainless-sdks/kernel-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Apps.Deploy(context.TODO(), kernel.AppDeployParams{
		File:    io.Reader(bytes.NewBuffer([]byte("REPLACE_ME"))),
		Version: kernel.String("REPLACE_ME"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Apps)
}
