#!/bin/bash

set -e
set -x

cd .generator
poetry install
poetry run python -m generator ./schemas/* -o ../api

cd ../api
goimports -w .