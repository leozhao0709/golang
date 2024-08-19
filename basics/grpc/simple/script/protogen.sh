#!/bin/sh

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"
GO_PROJECT="."
CPP_PROJECT="."
PYTHON_PROJECT="."

# Define output directories for different languages
GO_OUT_DIR=$BASE_PATH/$GO_PROJECT/protogen
CPP_OUT_DIR=$BASE_PATH/$CPP_PROJECT/protogen
PYTHON_OUT_DIR=$BASE_PATH/$PYTHON_PROJECT/protogen

# Create output directories if they don't exist
mkdir -p $GO_OUT_DIR
# mkdir -p $CPP_OUT_DIR
# mkdir -p $PYTHON_OUT_DIR

# Find all .proto files
PROTO_FILES=$(find $BASE_PATH/proto -name "*.proto")


# Generate Go files
protoc -I $BASE_PATH/proto \
--go_out=$GO_OUT_DIR --go_opt=paths=source_relative \
--go-grpc_out=$GO_OUT_DIR --go-grpc_opt=paths=source_relative \
$PROTO_FILES
# --cpp_out=$CPP_OUT_DIR \
# --python_out=$PYTHON_OUT_DIR \