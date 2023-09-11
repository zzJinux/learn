all: build
test: unittest

PLATFORM=local

.PHONY: build unittest clean

build:
	@docker build . --target bin --output _out --platform ${PLATFORM}
unittest:
	@docker build . --target unittest
clean:
	@rm -rf _out
