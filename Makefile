.PHONY: test build generate fmt

ginkgo := go run github.com/onsi/ginkgo/v2/ginkgo -r --race --cover --trace --timeout 2m -v

build:
	go build ./...

unit-test:
	$(ginkgo) --timeout 3m --skip-package integration

test: unit-test integration-test

integration-test:
	$(ginkgo) ./internal/integration

clean:
	find . -type f -wholename "*fakes*/fake_*go" -wholename "*internal*/fake_*go" -exec rm -v {} \;

generate: clean
	go generate ./...

fmt:
	go fmt -s -w .
