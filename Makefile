${GOPATH}/cmd/govrhub-ex: $(shell find cmd pkg)
	go install ./cmd/govrhub-ex

# Will default run to dev profile
run: ${GOPATH}/cmd/govrhub-ex
	${GOPATH}/bin/govrhub-ex

clean:
	@rm -f ${GOPATH}/bin/govrhub-ex

test:
	go test ./pkg/... -v -covermode=count

.PHONY: run clean test
