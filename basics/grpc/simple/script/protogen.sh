#!/bin/sh

# Define output directories for different languages
GO_OUT_DIR=gen/go
CPP_OUT_DIR=gen/cpp
PYTHON_OUT_DIR=gen/python

# Create output directories if they don't exist
mkdir -p $GO_OUT_DIR
mkdir -p $CPP_OUT_DIR
mkdir -p $PYTHON_OUT_DIR

# Find all .proto files
PROTO_FILES=$(find proto -name "*.proto")


# Generate Go files
protoc -I . \
--go_out=$GO_OUT_DIR --go_opt=paths=source_relative \
--go-grpc_out=$GO_OUT_DIR --go-grpc_opt=paths=source_relative \
--cpp_out=$CPP_OUT_DIR \
--python_out=$PYTHON_OUT_DIR \
$PROTO_FILES
