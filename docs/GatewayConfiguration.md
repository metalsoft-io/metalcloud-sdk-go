# GatewayConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branding** | Pointer to [**Branding**](Branding.md) | UI branding configuration (logos, colors, company name) | [optional] 
**LicenseMicroservice** | Pointer to [**LicenseMicroservice**](LicenseMicroservice.md) | License microservice connection settings | [optional] 
**Microservices** | Pointer to [**Microservices**](Microservices.md) | Connected microservices configuration | [optional] 
**RateLimiting** | Pointer to [**RateLimitingConfig**](RateLimitingConfig.md) | API rate limiting configuration | [optional] 

## Methods

### NewGatewayConfiguration

`func NewGatewayConfiguration() *GatewayConfiguration`

NewGatewayConfiguration instantiates a new GatewayConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayConfigurationWithDefaults

`func NewGatewayConfigurationWithDefaults() *GatewayConfiguration`

NewGatewayConfigurationWithDefaults instantiates a new GatewayConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranding

`func (o *GatewayConfiguration) GetBranding() Branding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *GatewayConfiguration) GetBrandingOk() (*Branding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *GatewayConfiguration) SetBranding(v Branding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *GatewayConfiguration) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetLicenseMicroservice

`func (o *GatewayConfiguration) GetLicenseMicroservice() LicenseMicroservice`

GetLicenseMicroservice returns the LicenseMicroservice field if non-nil, zero value otherwise.

### GetLicenseMicroserviceOk

`func (o *GatewayConfiguration) GetLicenseMicroserviceOk() (*LicenseMicroservice, bool)`

GetLicenseMicroserviceOk returns a tuple with the LicenseMicroservice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseMicroservice

`func (o *GatewayConfiguration) SetLicenseMicroservice(v LicenseMicroservice)`

SetLicenseMicroservice sets LicenseMicroservice field to given value.

### HasLicenseMicroservice

`func (o *GatewayConfiguration) HasLicenseMicroservice() bool`

HasLicenseMicroservice returns a boolean if a field has been set.

### GetMicroservices

`func (o *GatewayConfiguration) GetMicroservices() Microservices`

GetMicroservices returns the Microservices field if non-nil, zero value otherwise.

### GetMicroservicesOk

`func (o *GatewayConfiguration) GetMicroservicesOk() (*Microservices, bool)`

GetMicroservicesOk returns a tuple with the Microservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroservices

`func (o *GatewayConfiguration) SetMicroservices(v Microservices)`

SetMicroservices sets Microservices field to given value.

### HasMicroservices

`func (o *GatewayConfiguration) HasMicroservices() bool`

HasMicroservices returns a boolean if a field has been set.

### GetRateLimiting

`func (o *GatewayConfiguration) GetRateLimiting() RateLimitingConfig`

GetRateLimiting returns the RateLimiting field if non-nil, zero value otherwise.

### GetRateLimitingOk

`func (o *GatewayConfiguration) GetRateLimitingOk() (*RateLimitingConfig, bool)`

GetRateLimitingOk returns a tuple with the RateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimiting

`func (o *GatewayConfiguration) SetRateLimiting(v RateLimitingConfig)`

SetRateLimiting sets RateLimiting field to given value.

### HasRateLimiting

`func (o *GatewayConfiguration) HasRateLimiting() bool`

HasRateLimiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


