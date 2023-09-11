all: build

PLATFORM=local

.PHONY: build clean
build:
	@docker build . --target bin --output _out --platform ${PLATFORM}
clean:
	@rm -rf _out