.PHONY: build-image run.local run.image

build-image:
	@docker build . --target bin -t local/echocors:echo-v4.6.1 --platform linux/amd64

run.local:
	@echo "Running locally"
	@eval "$(<.env sed 's/^/export /')" && go run .

run.image:
	@echo "Running image"
	@docker run -t --rm -p 8989:8989 --env-file .env local/echocors:echo-v4.6.1