#!/usr/bin/env bash

# Check if running with root privileges
if [[ $EUID -eq 0 ]]; then
    echo "Error: It is not recommended to run this script with root privileges" >&2
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed, please install Go first" >&2
    exit 1
fi

echo "Detected Go version: $(go version)"

# Get GOPATH and GOBIN paths
GOPATH=$(go env GOPATH)
GOBIN=$(go env GOBIN)

# If GOBIN is not set, use the default path
if [[ -z "$GOBIN" ]]; then
    GOBIN="$GOPATH/bin"
fi

echo "GOBIN path: $GOBIN"

# Check if GOBIN is in PATH
if [[ ":$PATH:" != *":$GOBIN:"* ]]; then
    echo "Error: GOBIN path ($GOBIN) is not in PATH" >&2
    echo "Error: Please add the following line to your ~/.bashrc or ~/.zshrc file:" >&2
    echo "Error: export PATH=\$PATH:\$GOPATH/bin" >&2
    exit 1
fi

# install gs
echo "Installing gs ..."
go install github.com/go-spring/gs@main && echo "gs installed successfully" || { echo "Failed to install gs"; exit 1; }

## install gs-new
#echo "Installing gs-new ..."
#go install github.com/go-spring/gs-new@latest && echo "gs-new installed successfully" || { echo "Failed to install gs-new"; exit 1; }
#
## install gs-gen
#echo "Installing gs-gen ..."
#go install github.com/go-spring/gs-gen/cmd@latest && echo "gs-gen installed successfully" || { echo "Failed to install gs-gen"; exit 1; }

# install gs-mock
echo "Installing gs-mock ..."
go install github.com/go-spring/gs-mock@main && echo "gs-mock installed successfully" || { echo "Failed to install gs-mock"; exit 1; }

echo "All gs tools installed successfully!"
