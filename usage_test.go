// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/onkernel/kernel-go-sdk"
	"github.com/onkernel/kernel-go-sdk/internal/testutil"
	"github.com/onkernel/kernel-go-sdk/option"
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
	deployment, err := client.Apps.Deployments.New(context.TODO(), kernel.AppDeploymentNewParams{
		EntrypointRelPath: "main.ts",
		File:              io.Reader(bytes.NewBuffer([]byte("REPLACE_ME"))),
		Version:           kernel.String("1.0.0"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", deployment.Apps)
}
