BIN        := http-here
BUILD      := $(shell git rev-parse --short HEAD)
VERSION    := $(shell git describe --tags $(shell git rev-list --tags --max-count=1))
MAKEFILE   := $(word $(words $(MAKEFILE_LIST)), $(MAKEFILE_LIST))
BASE_DIR   := $(shell cd $(dir $(MAKEFILE)); pwd)
SOURCES    := $(shell find . -type f -name '*.go')
PKGS       := $(shell go list ./...)
COVER_OUT  := coverage.out
COVER_HTML := coverage.html
TMP_COVER  := tmp_cover.out
ECHO       := /bin/echo
LINT       := $(GOBIN)/golint
GOTOOLDIR  := $(shell go env GOTOOLDIR)
VET        := $(GOTOOLDIR)/vet
COVER      := $(GOTOOLDIR)/cover


.PHONY: build
build: check_gopath $(BIN)


.PHONY: all
all: clean cover lint vet build


.PHONY: check_gopath
check_gopath:
ifndef GOPATH
	@echo "ERROR!! GOPATH must be declared. Check http://golang.org/doc/code.html#GOPATH"
	@exit 1
endif
ifeq ($(shell go list ./... | grep -q '^_'; echo $$?), 0)
	@echo "ERROR!! This directory is in the wrong place"
	@exit 1
endif
	@exit 0


.PHONY: check_gobin
check_gobin:
ifndef GOBIN
	@echo "ERROR!! GOBIN must be declared. Check http://golang.org/doc/code.html#GOBIN"
	@exit 1
endif
	@exit 0


$(BIN): $(SOURCES) 
	@$(ECHO) "Building $(BIN) $(VERSION) $(BUILD)..."
	@go build -ldflags "-X main.Build $(BUILD) -X main.Version $(VERSION)" -o $(BIN) 


.PHONY: clean
clean: check_gopath
	@$(ECHO) "Removing temp files..."
	@rm -fv $(BIN) *.cover *.out 
	@find . -name '.*.swp' -exec rm -fv {} \;
	@go clean -v


.PHONY: test
test: check_gopath
	@for pkg in $(PKGS); do go test -v -race $$pkg || exit 1; done


$(COVER): check_gopath check_gobin
	@go get code.google.com/p/go.tools/cmd/cover || exit 0

.PHONY: cover
cover: check_gopath $(COVER)
	@echo 'mode: set' > $(COVER_OUT)
	@touch $(TMP_COVER)
	@for pkg in $(PKGS); do \
		go test -v -coverprofile=$(TMP_COVER) $$pkg || exit 1; \
		grep -v 'mode: set' $(TMP_COVER) >> $(COVER_OUT); \
	done
	@go tool cover -html=$(COVER_OUT) -o $(COVER_HTML)
	@(which gnome-open && gnome-open $(COVER_HTML)) || (which -s open && open $(COVER_HTML)) || (exit 0)
	@echo Generated HTML report in $(COVER_HTML)...


$(LINT): check_gopath check_gobin
	@go get github.com/golang/lint/golint

.PHONY: lint
lint: $(LINT)
	@for src in $(SOURCES); do golint $$src || exit 1; done


$(VET): check_gopath check_gobin
	@go get code.google.com/p/go.tools/cmd/vet || exit 0

.PHONY: vet
vet: check_gopath $(VET)
	@for src in $(SOURCES); do go tool vet $$src; done


.PHONY: fmt
fmt: $(SOURCES)
	@for src in $(SOURCES); do gofmt -w $$src; done
