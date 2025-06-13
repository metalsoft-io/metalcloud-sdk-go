# CreateNetworkDeviceDefaultSecretsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **float32** | The site ID of the network device default secrets | 
**MacAddressOrSerialNumber** | **string** | The MAC address or serial number of the network device | 
**SecretName** | **string** | The name of the secret | 
**SecretValue** | **string** | The value of the secret | 

## Methods

### NewCreateNetworkDeviceDefaultSecretsDto

`func NewCreateNetworkDeviceDefaultSecretsDto(siteId float32, macAddressOrSerialNumber string, secretName string, secretValue string, ) *CreateNetworkDeviceDefaultSecretsDto`

NewCreateNetworkDeviceDefaultSecretsDto instantiates a new CreateNetworkDeviceDefaultSecretsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDeviceDefaultSecretsDtoWithDefaults

`func NewCreateNetworkDeviceDefaultSecretsDtoWithDefaults() *CreateNetworkDeviceDefaultSecretsDto`

NewCreateNetworkDeviceDefaultSecretsDtoWithDefaults instantiates a new CreateNetworkDeviceDefaultSecretsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateNetworkDeviceDefaultSecretsDto) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetMacAddressOrSerialNumber

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetMacAddressOrSerialNumber() string`

GetMacAddressOrSerialNumber returns the MacAddressOrSerialNumber field if non-nil, zero value otherwise.

### GetMacAddressOrSerialNumberOk

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetMacAddressOrSerialNumberOk() (*string, bool)`

GetMacAddressOrSerialNumberOk returns a tuple with the MacAddressOrSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressOrSerialNumber

`func (o *CreateNetworkDeviceDefaultSecretsDto) SetMacAddressOrSerialNumber(v string)`

SetMacAddressOrSerialNumber sets MacAddressOrSerialNumber field to given value.


### GetSecretName

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *CreateNetworkDeviceDefaultSecretsDto) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetSecretValue

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *CreateNetworkDeviceDefaultSecretsDto) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *CreateNetworkDeviceDefaultSecretsDto) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


