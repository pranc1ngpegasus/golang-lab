.PHONY: generate.mock
generate.mock:
	@go generate ./domain/...

.PHONY: install.playwright
install.playwright:
	@go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps
