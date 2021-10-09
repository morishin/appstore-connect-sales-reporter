# appstore-connect-sales-reporter

## Run

```sh
$ go run .
```

## Re-run oapi-codegen

```sh
$ cd openapi
$ oapi-codegen -generate "types" -package openapi app_store_connect_api_1.5.1_openapi_fixed.json > ./types.gen.go
$ oapi-codegen -generate "client" -package openapi app_store_connect_api_1.5.1_openapi_fixed.json > ./client.gen.go
```
