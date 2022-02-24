# All source code and documents, used when checking for misspellings
ALLDOC := $(shell find . \( -name "*.md" -o -name "*.yaml" \) \
                                -type f | sort)
ALL_MODULES := $(shell find . -type f -name "go.mod" -exec dirname {} \; | sort )

# All source code files
ALL_SRC := $(shell find . -name '*.go' -o -name '*.sh' -o -name 'Dockerfile' -type f | sort)

OUTDIR=./dist
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

ifeq ($(GOOS), windows)
EXT?=.exe
else
EXT?=
endif

# Default build target; making this should build for the current os/arch
.PHONY: collector
collector:
	go build -o $(OUTDIR)/collector_$(GOOS)_$(GOARCH)$(EXT) ./cmd/collector

.PHONY: build-all
build-all: build-linux build-darwin build-windows

.PHONY: build-linux
build_linux: build-linux-amd64 build-linux-arm64 build-linux-arm

.PHONY: build-darwin
build-linux: build-darwin-amd64 build-darwin-arm64

.PHONY: build-windows
build-linux: build-windows-amd64

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(MAKE) collector

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(MAKE) collector

.PHONY: build-linux-arm
build-linux-arm:
	GOOS=linux GOARCH=arm $(MAKE) collector

.PHONY: build-darwin-amd64
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(MAKE) collector

.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(MAKE) collector

.PHONY: build-windows-amd64
build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(MAKE) collector

# tool-related commands
TOOLS_MOD_DIR := ./internal/tools
.PHONY: install-tools
install-tools:
	cd $(TOOLS_MOD_DIR) && go install github.com/mgechev/revive@latest 
	cd $(TOOLS_MOD_DIR) && go install github.com/google/addlicense
	cd $(TOOLS_MOD_DIR) && go install golang.org/x/tools/cmd/goimports	
	cd $(TOOLS_MOD_DIR) && go install github.com/client9/misspell/cmd/misspell
	cd $(TOOLS_MOD_DIR) && go install github.com/sigstore/cosign/cmd/cosign
	cd $(TOOLS_MOD_DIR) && go install github.com/goreleaser/goreleaser@v1.3.1

.PHONY: lint
lint:
	revive -config revive/config.toml -formatter friendly ./...

.PHONY: misspell
misspell:
	misspell $(ALLDOC)

.PHONY: misspell-fix
misspell-fix:
	misspell -w $(ALLDOC)

.PHONY: test
test:
	go test -race ./...

.PHONY: test-with-cover
test-with-cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: bench
bench:
	go test -benchmem -run=^$$ -bench ^* ./...

.PHONY: check-fmt
check-fmt:
	goimports -d ./ | diff -u /dev/null -

.PHONY: fmt
fmt:
	goimports -w .

.PHONY: tidy
tidy:
	$(MAKE) for-all CMD="go mod tidy -go=1.17"

# This target performs all checks that CI will do (excluding the build itself)
.PHONY: ci-checks
ci-checks: check-fmt misspell lint test

# This target checks that license copyright header is on every source file
.PHONY: check-license
check-license:
	@ADDLICENSEOUT=`addlicense -check $(ALL_SRC) 2>&1`; \
		if [ "$$ADDLICENSEOUT" ]; then \
			echo "addlicense FAILED => add License errors:\n"; \
			echo "$$ADDLICENSEOUT\n"; \
			echo "Use 'make add-license' to fix this."; \
			exit 1; \
		else \
			echo "Check License finished successfully"; \
		fi

# This target adds a license copyright header is on every source file that is missing one
.PHONY: add-license
add-license:
	@ADDLICENSEOUT=`addlicense -y "" -c "observIQ, Inc." $(ALL_SRC) 2>&1`; \
		if [ "$$ADDLICENSEOUT" ]; then \
			echo "addlicense FAILED => add License errors:\n"; \
			echo "$$ADDLICENSEOUT\n"; \
			exit 1; \
		else \
			echo "Add License finished successfully"; \
		fi

# Build, sign, and release
.PHONY: release
release:
	goreleaser release --parallelism 4 --rm-dist

# Build and sign, skip release and ignore dirty git tree
.PHONY: release-test
release-test:
	goreleaser release --parallelism 4 --skip-validate --skip-publish --rm-dist

.PHONY: for-all
for-all:
	@echo "running $${CMD} in root"
	@$${CMD}
	@set -e; for dir in $(ALL_MODULES); do \
	  (cd "$${dir}" && \
	  	echo "running $${CMD} in $${dir}" && \
	 	$${CMD} ); \
	done

.PHONY: clean
clean:
	rm -rf $(OUTDIR)
