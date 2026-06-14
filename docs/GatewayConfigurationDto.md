# GatewayConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branding** | Pointer to [**BrandingDto**](BrandingDto.md) | UI branding configuration (logos, colors, company name) | [optional] 
**LicenseMicroservice** | Pointer to [**LicenseMicroserviceDto**](LicenseMicroserviceDto.md) | License microservice connection settings | [optional] 
**Microservices** | Pointer to [**MicroservicesDto**](MicroservicesDto.md) | Connected microservices configuration | [optional] 
**RateLimiting** | Pointer to [**RateLimitingConfigDto**](RateLimitingConfigDto.md) | API rate limiting configuration | [optional] 

## Methods

### NewGatewayConfigurationDto

`func NewGatewayConfigurationDto() *GatewayConfigurationDto`

NewGatewayConfigurationDto instantiates a new GatewayConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayConfigurationDtoWithDefaults

`func NewGatewayConfigurationDtoWithDefaults() *GatewayConfigurationDto`

NewGatewayConfigurationDtoWithDefaults instantiates a new GatewayConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranding

`func (o *GatewayConfigurationDto) GetBranding() BrandingDto`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *GatewayConfigurationDto) GetBrandingOk() (*BrandingDto, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *GatewayConfigurationDto) SetBranding(v BrandingDto)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *GatewayConfigurationDto) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetLicenseMicroservice

`func (o *GatewayConfigurationDto) GetLicenseMicroservice() LicenseMicroserviceDto`

GetLicenseMicroservice returns the LicenseMicroservice field if non-nil, zero value otherwise.

### GetLicenseMicroserviceOk

`func (o *GatewayConfigurationDto) GetLicenseMicroserviceOk() (*LicenseMicroserviceDto, bool)`

GetLicenseMicroserviceOk returns a tuple with the LicenseMicroservice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseMicroservice

`func (o *GatewayConfigurationDto) SetLicenseMicroservice(v LicenseMicroserviceDto)`

SetLicenseMicroservice sets LicenseMicroservice field to given value.

### HasLicenseMicroservice

`func (o *GatewayConfigurationDto) HasLicenseMicroservice() bool`

HasLicenseMicroservice returns a boolean if a field has been set.

### GetMicroservices

`func (o *GatewayConfigurationDto) GetMicroservices() MicroservicesDto`

GetMicroservices returns the Microservices field if non-nil, zero value otherwise.

### GetMicroservicesOk

`func (o *GatewayConfigurationDto) GetMicroservicesOk() (*MicroservicesDto, bool)`

GetMicroservicesOk returns a tuple with the Microservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroservices

`func (o *GatewayConfigurationDto) SetMicroservices(v MicroservicesDto)`

SetMicroservices sets Microservices field to given value.

### HasMicroservices

`func (o *GatewayConfigurationDto) HasMicroservices() bool`

HasMicroservices returns a boolean if a field has been set.

### GetRateLimiting

`func (o *GatewayConfigurationDto) GetRateLimiting() RateLimitingConfigDto`

GetRateLimiting returns the RateLimiting field if non-nil, zero value otherwise.

### GetRateLimitingOk

`func (o *GatewayConfigurationDto) GetRateLimitingOk() (*RateLimitingConfigDto, bool)`

GetRateLimitingOk returns a tuple with the RateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimiting

`func (o *GatewayConfigurationDto) SetRateLimiting(v RateLimitingConfigDto)`

SetRateLimiting sets RateLimiting field to given value.

### HasRateLimiting

`func (o *GatewayConfigurationDto) HasRateLimiting() bool`

HasRateLimiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


