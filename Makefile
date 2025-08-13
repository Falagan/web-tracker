MAIN_PACKAGE_PATH := ./cmd
SERVER_PACKAGE_PATH := ./cmd/server
BINARY_NAME := web_tracker
PERF_PATH := ./perf
SHELL := /bin/bash

# ==================================================================================== #
# 	                                   RUN                                          	 #
# ==================================================================================== #

.PHONY: run
run:
	@echo -e "[RUN] Starting server..."
	go run $(MAIN_PACKAGE_PATH)/main.go

.PHONY: run/live
run/live:
	@echo -e "[RUN] Starting server with hot-reload..."
	~/go/bin/air \
		--build.cmd "make tidy format build/dev" \
		--build.bin "./dist/bin/$(BINARY_NAME)" \
		--build.exclude_dir "dist" \
		--build.stop_on_error "true" 

# ==================================================================================== #
# 	                                   TEST                                          	 #
# ==================================================================================== #

.PHONY: test
test:
	@echo -e "[TEST] Testing..."
	go test -v -race -buildvcs -vet=off ./...

.PHONY: test/coverage
test/coverage:
	@echo -e "[TEST] Testing with coverage..."
	go test -v -race -buildvcs -vet=off -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# 	                                  BUILD                                          	 #
# ==================================================================================== #

.PHONY: build/dev
build/dev: 
	@echo -e "[BUILD] Go build dev..."
	go build -o=./dist/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: build
build: 
	@echo -e "[BUILD] Go build..."
	go build -o=./dist/bin/${BINARY_NAME} -ldflags="-s -w" -trimpath ${MAIN_PACKAGE_PATH}

# ==================================================================================== #
# 	                                   LINT                                          	 #
# ==================================================================================== #

.PHONY: tidy
tidy:
	@echo -e "[DEPENDENCIES] Get Go dependencies...$(NC)"
	go mod tidy -v

.PHONY: format
format:
	@echo -e "[FORMAT] Formating go code..."
	go fmt ./...

# audit: dependencies, basic go linter, advance go linter and security vul.
.PHONY: audit
audit:
	@echo -e "[LINT] Audit go code..."
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: audit-soft
audit-soft:
	@echo -e "[LINT] Audit go code softly..."
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000 ./...

# ==================================================================================== #
# 	                                   PERF                                          	 #
# ==================================================================================== #

.PHONY: perf/mem
perf/mem:
	@echo -e "[PERF] Starting memory profiling..."
	@mkdir -p $(PERF_PATH)
	go tool pprof -http=:8081 http://localhost:6060/debug/pprof/heap

.PHONY: perf/clean
perf/clean:
	@echo -e "[PERF] Cleaning profile files..."
	rm -rf $(PERF_PATH)

# ==================================================================================== #
# 	                                   GIT                                          	 #
# ==================================================================================== #

## commit: commit changes with previous audit soft
.PHONY: commit
commit: audit-soft
	@read -p "[COMMIT] Message: " COMMIT_MSG; \
	git commit -m "$$COMMIT_MSG"

## push: push changes to the remote Git repository
.PHONY: push
push: audit
	@echo -e "[PUSH] Git push..."
	git push
