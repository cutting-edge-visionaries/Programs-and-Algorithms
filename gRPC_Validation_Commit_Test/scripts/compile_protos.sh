#!/bin/bash

set -eux

# To set a proto root for a set of protos, create a .protoroot file in one of the parent directories
# which you wish to use as the proto root.  If no .protoroot file exists within fabric/.../<your_proto>
# then the proto root for that proto is inferred to be its containing directory.

# Find explicit proto roots

if [[ `uname` == "Darwin" ]]; then
  PROTO_ROOT_FILES="$(find protos/ -name ".protoroot" -exec stat -f {} \;)"
else
  PROTO_ROOT_FILES="$(find protos/ -name ".protoroot" -exec readlink -f {} \;)"
fi

PROTO_ROOT_DIRS="$(dirname $PROTO_ROOT_FILES)"

# Find all proto files to be compiled, excluding any which are in a proto root or in the vendor folder

if [[ `uname` == "Darwin" ]]; then
  ROOTLESS_PROTO_FILES="$(find $PWD/protos/ \
                             $(for dir in $PROTO_ROOT_DIRS ; do echo "-path $dir -prune -o " ; done) \
                             -path $PWD/vendor -prune -o \
                             -name "*.proto" -exec stat -f {} \;)"
else
  ROOTLESS_PROTO_FILES="$(find $PWD/protos/ \
                             $(for dir in $PROTO_ROOT_DIRS ; do echo "-path $dir -prune -o " ; done) \
                             -path $PWD/vendor -prune -o \
                             -name "*.proto" -exec readlink -f {} \;)"
fi

ROOTLESS_PROTO_DIRS="$(dirname $ROOTLESS_PROTO_FILES | sort | uniq)"


for dir in $ROOTLESS_PROTO_DIRS $PROTO_ROOT_DIRS; do
echo Working on dir $dir
  # As this is a proto root, and there may be subdirectories with protos, compile the protos for each sub-directory which contains them
  for protos in $(find "$dir" -name '*.proto' -exec dirname {} \; | sort | uniq) ; do
    protoc --proto_path="$dir" --proto_path="$PWD/protos" --go_out=plugins=grpc:$GOPATH/src/ "$protos"/*.proto
  done
done
