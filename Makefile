include dev.env

OS=linux
ARCH=x86_64
PROTOC_VERSION=3.14.0
PROTOC_DIR=protoc-$(PROTOC_VERSION)-$(OS)-$(ARCH)
CGO_ENABLED=0

export

.PHONY: get-protoc
get-protoc:
	@if [ ! -d "$(PROTOC_DIR)" ]; then \
		make get-protoc-internal; \
	fi

.PHONY: get-protoc-internal
get-protoc-internal:
	curl -Ls https://github.com/protocolbuffers/protobuf/releases/download/v$$PROTOC_VERSION/protoc-$$PROTOC_VERSION-$$OS-$$ARCH.zip -o protoc.zip
	unzip -oq protoc.zip -d $$PROTOC_DIR
	rm protoc.zip

.PHONY: get-protoc
proto: get-protoc
	GO111MODULE=off \
		go get google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

	mkdir -p pkg

	PATH="$$PATH:$$(go env GOPATH)/bin" \
		$$PROTOC_DIR/bin/protoc \
		--go_out=pkg \
		--go-grpc_out=pkg \
		api/products.proto

.PHONY: protset
protoset: proto
	PATH="$$PATH:$$(go env GOPATH)/bin" \
		$$PROTOC_DIR/bin/protoc \
		--proto_path=./api \
		--descriptor_set_out=products.protoset \
		--include_imports \
		products.proto

.PHONY: lint
lint: proto
	# TBD

.PHONY: test
test: proto
	# TBD

.PHONY: run
run: proto
	go run cmd/products/main.go

.PHONY: build
build: proto
	go build -o bin/products cmd/products/main.go

.PHONY: clean
clean:
	rm -rf bin protoc-* pkg/productspb

.PHONY: up
up:
	sudo docker-compose up -d

.PHONY: down
down:
	sudo docker-compose down
