SHELL := /bin/bash

include config.mk

# Builds all the gotest related images then brings them up
.PHONY: run
run: build up

# builds the go microservice
.PHONY: build
build: clean_service_artifact setup_service_config
	@set -e; \
	if [ $(os) ]; then \
		export GO_BUILD_OS=$(os); \
	fi; \
	if [ "$$GO_BUILD_OS" == "linux" ]; then \
		cd ../cmd/go-assessment; \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../build/$${DISQO_BUILD_BINARY_DIR}/$${DISQO_SERVICE_BINARY_DIR} .; \
	else \
		cd ../cmd/go-assessment; \
		go build -o ../../build/$${DISQO_BUILD_BINARY_DIR}/$${DISQO_SERVICE_BINARY_DIR} .; \
	fi
# brings up the go microservice
.PHONY: up
up:
	@cd ./$${DISQO_BUILD_BINARY_DIR}/$${DISQO_SERVICE_BINARY_DIR}; \
	./go-assessment -config-filepath ./config.json

# Sets up respective service config depending on build environment
setup_service_config: setup_artifact_dir
	@set -e; \
	if [ $(env) ]; then \
		export DISQO_ENV=$(env); \
	else \
	  	export DISQO_ENV=dev; \
	fi; \
	if [ -d ../config/cluster/go-assessment-$${DISQO_ENV} ]; then \
		cp -R ../config/cluster/go-assessment-$${DISQO_ENV}/* ./$${DISQO_BUILD_BINARY_DIR}/$${DISQO_SERVICE_BINARY_DIR}; \
	fi

# Sets up the artifact directories needed to build binaries
setup_artifact_dir:
	@set -e; \
	mkdir -p $${DISQO_BUILD_BINARY_DIR}/$${DISQO_SERVICE_BINARY_DIR}


# Deletes the artifact folder that holds the binaries / configs for gotest microservice
clean_service_artifact:
	@if [ -d $${DISQO_BUILD_BINARY_DIR}/$${DISQO_SERVICE_BINARY_DIR} ]; then \
		echo; \
		echo -e "$${DISQO_LIGHT_CYAN}cleaning microservice artifact directory...$${DISQO_NC}"; \
		rm -rf $${DISQO_BUILD_BINARY_DIR}/$${DISQO_SERVICE_BINARY_DIR}; \
		echo -e "$${DISQO_LIGHT_CYAN}done!$${DISQO_NC}"; \
		echo; \
	fi

.PHONY: setup
setup:
	@docker-compose -f docker-compose.yml build go-assessment-mysql
	@docker-compose -f docker-compose.yml up -d go-assessment-mysql
	@$(MAKE) heart_beat

.PHONY: kill
kill:
	@docker rm -f go-assessment-mysql
	@docker rm -f go-assessment-app

# Checks RDS heartbeat
.PHONY: heart_beat
heart_beat:
	@cd ../sql-setup; \
	go run .