export DOCKERHOST=$(shell ifconfig | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3}" | grep -v 127.0.0.1 | awk '{ print $$2 }' | cut -f2 -d: | head -n1)

.PHONY: cli
cli:
	docker-compose run --rm app sh

.PHONY: compile
compile:
	CGO_ENABLED=0 \
	GO111MODULES=on \
	go build \
		-a \
		-o ./bin/main \
		-ldflags '-extldflags -static'

.PHONY: build-image
build-image:
	docker build -t gopher-quizz:0.0.0 .
