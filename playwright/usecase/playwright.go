package usecase

import (
	"context"
	"fmt"
	"os"

	"github.com/Pranc1ngPegasus/golang-lab/playwright/domain/tracer"
	domain "github.com/Pranc1ngPegasus/golang-lab/playwright/domain/usecase"
	"github.com/google/wire"
	playwright "github.com/playwright-community/playwright-go"
	"github.com/volatiletech/null/v8"
)

var _ domain.Playwright = (*Playwright)(nil)

var NewPlaywrightSet = wire.NewSet(
	wire.Bind(new(domain.Playwright), new(*Playwright)),
	NewPlaywright,
)

type Playwright struct {
	tracer  tracer.Tracer
	browser playwright.Browser
}

func NewPlaywright(
	tracer tracer.Tracer,
) (*Playwright, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to start playwright: %w", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: null.BoolFrom(true).Ptr(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to launch browser: %w", err)
	}

	return &Playwright{
		tracer:  tracer,
		browser: browser,
	}, nil
}

func (u *Playwright) Do(ctx context.Context, input domain.PlaywrightInput) (*domain.PlaywrightOutput, error) {
	ctx, span := u.tracer.Tracer().Start(ctx, "usecase.Playwright")
	defer span.End()

	browserContext, err := u.browser.NewContext()
	if err != nil {
		return nil, fmt.Errorf("failed to create new context: %w", err)
	}

	if err := browserContext.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: null.BoolFrom(true).Ptr(),
		Snapshots:   null.BoolFrom(true).Ptr(),
	}); err != nil {
		return nil, fmt.Errorf("failed to start tracing: %w", err)
	}
	defer browserContext.Tracing().Stop(playwright.TracingStopOptions{
		Path: null.StringFrom("trace.zip").Ptr(),
	})

	page, err := browserContext.NewPage()
	if err != nil {
		return nil, fmt.Errorf("failed to create new page: %w", err)
	}

	if _, err := page.Goto(input.URL); err != nil {
		return nil, fmt.Errorf("failed to open(URL: %s): %w", input.URL, err)
	}

	buf, err := page.Screenshot(playwright.PageScreenshotOptions{
		FullPage: null.BoolFrom(true).Ptr(),
		Size:     playwright.ScreenshotSizeDevice,
		Type:     playwright.ScreenshotTypePng,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to take screenshot: %w", err)
	}

	file, err := os.Create("screenshot.png")
	if err != nil {
		return nil, fmt.Errorf("failed to create screenshot file: %w", err)
	}

	if _, err := file.Write(buf); err != nil {
		return nil, fmt.Errorf("failed to write screenshot: %w", err)
	}

	if err := page.Close(); err != nil {
		return nil, fmt.Errorf("failed to close page: %w", err)
	}

	return &domain.PlaywrightOutput{}, nil
}
