

.PHONY: b
b:
	go build cmd/main.go

.PHONY: br
br: build-run

.PHONY: build-run
build-run:
	go build cmd/main.go && ./main

.PHONY: test
test: _test
	-@docker compose -f docker-compose-test.yml down

