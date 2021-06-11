.PHONY: format
format:
	bin/format.sh

.PHONY: check-import
check-import:
	bin/check-import.sh

.PHONY: lint
lint: cleanlintcache
	golangci-lint run ./...

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: pretty
pretty: tidy format lint