.PHONY: generate.mock
generate.mock:
	@go generate ./domain/...

.PHONY: generate.wire
generate.wire:
	@go run github.com/google/wire/cmd/wire ./...

.PHONY: install.playwright
install.playwright:
	@go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps
