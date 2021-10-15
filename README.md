# appstore-connect-sales-reporter

## Setup

Copy .env.example to .env and fill in the values.

```sh
cp .env.example .env
```

## Run

```sh
$ go run .
```

## Deploy to AWS

```sh
$ cd aws/cdk
$ cdk deploy
```

## Re-run oapi-codegen

```sh
$ cd openapi
$ oapi-codegen -generate "types" -package openapi app_store_connect_api_1.5.1_openapi_fixed.json > ./types.gen.go
$ oapi-codegen -generate "client" -package openapi app_store_connect_api_1.5.1_openapi_fixed.json > ./client.gen.go
```
