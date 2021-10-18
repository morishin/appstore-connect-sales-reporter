# appstore-connect-sales-reporter

Post App Store Connect Sales Report summary to Slack.

![screenshot](https://user-images.githubusercontent.com/1413408/137770668-a2e9891f-060d-474a-a87a-3b24b4f69bfe.png)

## Setup

1. Copy .env.example to .env and fill in the values.

    ```sh
    cp .env.example .env
    ```

1. Download your own AppStore Connect API key (See https://developer.apple.com/documentation/appstoreconnectapi/creating_api_keys_for_app_store_connect_api#3028598)
1. Rename it into `AuthKey.p8` and put it next to AuthKey.p8.example.


## Run

```sh
$ cd dev
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
