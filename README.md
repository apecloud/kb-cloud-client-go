# kb-cloud-api-client-go
Golang client for the KubeBlocks Cloud API


## Development

```bash
cd .generator
poetry install
poetry run python -m generator ./schemas/* -o ../api

cd ../api
goimports -w .
```

