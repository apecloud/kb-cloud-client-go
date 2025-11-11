#!/bin/bash

set -e
set -x

# Add Go bin directory to PATH if not already present
export PATH="${PATH}:$(go env GOPATH)/bin"

cd .generator
poetry install
rm -rf ../api/*
poetry run python -m generator ./schemas/* -o ../api
goimports -w ../api