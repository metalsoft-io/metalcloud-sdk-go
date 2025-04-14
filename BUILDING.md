# Buiding the MetalCloud SDK for Go

## To generate the SDK follow these steps:

* Install the OpenAPI generator and goimports.

```bash
npm install @openapitools/openapi-generator-cli -g
go install golang.org/x/tools/cmd/goimports@latest
```

* Create a configuration file for the OpenAPI generator.

```bash
cat >> metalcloud-sdk-go-config.json << END
{
    "packageName": "sdk",
    "packageVersion": "7.0.0",
    "disallowAdditionalPropertiesIfNotPresent": false,
    "useOneOfDiscriminatorLookup": true,
    "structPrefix": true,
    "enumClassPrefix": true,
    "withGoMod": false
}
END
```

* Checkout the MetalCloud SDK for Go repository.

```bash
git clone https://github.com/metalsoft-io/metalcloud-sdk-go.git
```

* Download the OpenAPI definition from the MetalCloud API server.

```bash
wget https://my-metalcloud.com/api/v2/swagger-json -O ./metalcloud-api.json --no-check-certificate
```

* Generate the SDK using

```bash
openapi-generator-cli generate -g go -i ./metalcloud-api.json -o ./metalcloud-sdk-go -c ./metalcloud-sdk-go-config.json --git-user-id=metalsoft-io --git-repo-id=metalcloud-sdk-go
```

* Removing unused imports

This step is needed as the OpenAPI generator leaves unused imports when used with the `useOneOfDiscriminatorLookup` option.
The `goimports` tool is used to remove unused imports.

```bash
~/go/bin/goimports -w ./metalcloud-sdk-go/model_authentication_request_properties.go
~/go/bin/goimports -w ./metalcloud-sdk-go/model_extension_input_options.go
~/go/bin/goimports -w ./metalcloud-sdk-go/model_extension_task_options.go
~/go/bin/goimports -w ./metalcloud-sdk-go/model_network_fabric_fabric_configuration.go
```
