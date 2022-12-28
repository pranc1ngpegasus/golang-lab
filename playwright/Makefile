.PHONY: dev
dev:
	@go run github.com/cosmtrek/air -c .air.toml

.PHONY: generate.mock
generate.mock:
	@go generate ./domain/...
	$(call format)

.PHONY: generate.wire
generate.wire:
	@go run github.com/google/wire/cmd/wire ./...
	$(call format)

.PHONY: install.playwright
install.playwright:
	@go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps

define format
	@go run mvdan.cc/gofumpt -l -w .
	@go run golang.org/x/tools/cmd/goimports -w .
	@go mod tidy
endef
