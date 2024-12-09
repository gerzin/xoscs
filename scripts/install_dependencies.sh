#!/usr/bin/env bash
echo "Starting the installation of dependencies..."

# Check if protoc is installed
if ! command -v protoc &> /dev/null
then
    echo "protoc could not be found, installing it now..."
    sudo apt update
    sudo apt install -y protobuf-compiler
else
    echo "protoc is already installed."
fi

# Check if go is installed
if ! command -v xgen &> /dev/null
then
    echo "xgen could not be found, installing it now..."
    go install github.com/openscenario/xgen@latest
else
    echo "xgen is already installed."
fi

# Check if protoc-gen-go is installed
if ! command -v protoc-gen-go &> /dev/null
then
    echo "protoc-gen-go could not be found, installing it now..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
else
    echo "protoc-gen-go is already installed."
fi
