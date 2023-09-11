ROOT_DIR:=$(realpath $(shell dirname $(firstword $(MAKEFILE_LIST))))

.PHONY: download install-tools

download:
	@echo "Downloading dependencies..."
	@go mod download

install-tools: download
	@echo "Installing tools... from tools.go"
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | GOBIN=$(ROOT_DIR)/bin xargs -tI % go install %

codegen: install-tools
	@echo "Generating code..."
	@mkdir -p api && bin/oapi-codegen --config server.cfg.yaml petstore-expanded.yaml