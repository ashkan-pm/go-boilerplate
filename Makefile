GO ?= go
GOFMT ?= gofmt "-s"
GOPACKAGES ?= $(shell $(GO) list ./...)
GOFILES := $(shell find . -name "*.go")

.PHONY: install
install:
	$(GO) mod tidy

.PHONY: run
run:
	$(GO) run src/cmd/go-boilerplate.go

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: vet
vet:
	$(GO) vet $(GOPACKAGES)

.PHONY: test
test:
	$(GO) test $(GOPACKAGES)

.PHONY: build
build:
	$(GO) build -o build/go-boilerplate src/cmd/go-boilerplate.go
