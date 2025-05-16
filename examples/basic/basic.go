package basic

import (
	"context"

	kernel "github.com/onkernel/kernel-go-sdk"
)

// PageTitleInput represents the input expected by the get-page-title action.
type PageTitleInput struct {
	URL string `json:"url"`
}

// PageTitleOutput represents the output returned by the get-page-title action.
type PageTitleOutput struct {
	Title string `json:"title"`
}

var app = kernel.App("go-basic")

func init() {
	// Signature 1: ctx only => error
	app.Action("ping", func(ctx context.Context) error {
		// do nothing
		return nil
	})

	// Signature 2: ctx only => (Out, err)
	app.Action("get-random-number", func(ctx context.Context) (int, error) {
		return 42, nil
	})

	// Signature 3: ctx, In => error
	app.Action("print-url", func(ctx context.Context, input PageTitleInput) error {
		_ = input.URL // pretend to do something
		return nil
	})

	// Signature 4: ctx, In => (Out, err)
	app.Action("get-page-title", func(ctx context.Context, input PageTitleInput) (PageTitleOutput, error) {
		// This is just a stub implementation used for demonstration purposes.
		return PageTitleOutput{Title: "Example Title for " + input.URL}, nil
	})
}
