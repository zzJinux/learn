all: build
test: lint unittest

PLATFORM=local

.PHONY: build unittest lint clean

build:
	@docker build . --target bin --output _out --platform ${PLATFORM}
unittest:
	@docker build . --target unittest
lint:
	@docker build . --target lint
clean:
	@rm -rf _out
