# Buiding the MetalCloud SDK for Go

## To generate the SDK follow these steps:

* Install the OpenAPI generator.

```bash
npm install @openapitools/openapi-generator-cli -g
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
