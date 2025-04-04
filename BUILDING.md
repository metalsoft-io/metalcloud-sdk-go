# Buiding the MetalCloud SDK for Go

## To generate the SDK follow these steps:

* Install the OpenAPI generator.

```bash
npm install @openapitools/openapi-generator-cli -g
```

* Modify the OpenAPI templates for Go to handle polymorphic types.

Extract the Go templates:

```bash
openapi-generator-cli author template -g go -o ./openapi-templates-go
```

Then modify the `model_oneof.mustache` file - lines 92:101:

```diff
#     92:        if len(dst.{{#lambda.type-to-name}}{{{.}}}{{/lambda.type-to-name}}.AdditionalProperties) > 0 {
#     93:            dst.{{#lambda.type-to-name}}{{{.}}}{{/lambda.type-to-name}} = nil
#     94:        } else {
#    content of original lines 92:101 indented by 4 spaces
#    105:        }
```

* Create a configuration file for the OpenAPI generator.

```bash
cat >> metalcloud-sdk-go-config.json << END
{
    "packageName": "sdk",
    "packageVersion": "7.0.0",
    "git-host": "github.com",
    "git-user-id": "metalsoft-io",
    "git-repo-id": "metalcloud-sdk-go",
    "disallowAdditionalPropertiesIfNotPresent": false,
    "structPrefix": true,
    "enumClassPrefix": true,
    "withGoMod": false
}
END
```

* Download the OpenAPI definition from the MetalCloud API server.

```bash
wget https://my-metalcloud.com/api/v2/swagger-json -O ./metalcloud-api.json --no-check-certificate
```

* Checkout the MetalCloud SDK for Go repository.

```bash
git clone https://github.com/metalsoft-io/metalcloud-sdk-go.git
```

* Generate the SDK using the modified templates.

```bash
openapi-generator-cli generate -g go -i ./metalcloud-api.json -o ./metalcloud-sdk-go -c ./metalcloud-sdk-go-config.json --git-user-id=metalsoft-io --git-repo-id=metalcloud-sdk-go -t ./openapi-templates-go
```
