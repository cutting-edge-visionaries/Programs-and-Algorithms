# Parameters
GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build

BUILD_DIR=build/bin

all: validation commit

.PHONY: dep
dep:
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go
	$(GOGET) -u github.com/golang/dep/cmd/dep
	dep init
	dep ensure

.PHONY: protos
protos:
	./scripts/compile_protos.sh

.PHONY: validation
validation:
	$(GOBUILD) -o $(BUILD_DIR)/validation validation/main.go

.PHONY: commit
commit:
	$(GOBUILD) -o $(BUILD_DIR)/commit commit/main.go

clean:
	@rm -rf build/
