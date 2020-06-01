PROTO_FILES=$(shell find proto -type f -name "*.proto")
PROTO_SRVFILES=$(shell find proto/service -type f -name "*.proto")
PROTO_GOFILES=$(shell find proto -type f -name "*.go")
TESTCMD=go test
MOCKGENCMD=mockgen

.PHONY: all
all: clean-proto proto-gen grpcui

.PHONY: proto-gen
proto-gen:
	@for pb in $(PROTO_FILES); do \
        protoc -I=$(GOPATH)/src:. \
        --go_out=paths=source_relative:. \
        --go-grpc_out=requireUnimplementedServers=false,paths=source_relative:. \
        --validate_out=lang=go,paths=source_relative:. $$pb; \
	done

.PHONY: clean-proto
clean-proto:
	@for go in $(PROTO_GOFILES); do \
        echo "remove" $$go; \
        rm -rf $$go; \
	done

.PHONY: grpcui
grpcui:
	grpcui -plaintext -port 60001 $(addprefix -proto ,$(PROTO_SRVFILES)) -import-path $(GOPATH)/src -import-path . localhost:50001

.PHONY: test
test:
	@$(TESTCMD) -count 1 ./... | grep -v 'no test files'

.PHONY: coverage
coverage:
	@go test ./... -coverprofile=/tmp/gf_coverage.out | grep -v 'no test files' && \
        go tool cover -html=/tmp/gf_coverage.out -o /tmp/gf_coverage.html && \
        open /tmp/gf_coverage.html

.PHONY: mock-all
mock-all: mock-dto mock-model mock-service

.PHONY: mock-dto
mock-dto:
	$(MOCKGENCMD) -source=app/dto/book/book.go -destination app/dto/book/mock/book.go

.PHONY: mock-model
mock-model:
	$(MOCKGENCMD) -source=app/model/book/book.go -destination app/model/book/mock/book.go

.PHONY: mock-service
mock-service:
	$(MOCKGENCMD) -source=app/service/book/book.go -destination app/service/book/mock/book.go
