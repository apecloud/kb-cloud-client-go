#!/bin/bash

set -e
set -x

cd .generator
poetry install
rm -rf ../api/*
poetry run python -m generator ./schemas/* -o ../api
goimports -w ../api