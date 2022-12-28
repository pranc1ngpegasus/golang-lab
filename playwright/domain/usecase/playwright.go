package usecase

import "context"

type Playwright interface {
	Do(ctx context.Context, input PlaywrightInput) (*PlaywrightOutput, error)
}

type (
	PlaywrightInput struct {
		URL string
	}

	PlaywrightOutput struct{}
)
