// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kernel_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/onkernel/kernel-go-sdk"
	"github.com/onkernel/kernel-go-sdk/internal/testutil"
	"github.com/onkernel/kernel-go-sdk/option"
)

func TestAppDeploymentNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Apps.Deployments.New(context.TODO(), kernel.AppDeploymentNewParams{
		EntrypointRelPath: "src/app.py",
		File:              io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		EnvVars: map[string]string{
			"foo": "string",
		},
		Force:   kernel.Bool(false),
		Region:  kernel.AppDeploymentNewParamsRegionAwsUsEast1a,
		Version: kernel.String("1.0.0"),
	})
	if err != nil {
		var apierr *kernel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
