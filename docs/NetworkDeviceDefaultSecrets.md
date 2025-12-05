# NetworkDeviceDefaultSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The ID of the network device default secrets | 
**SiteId** | **float32** | The site ID of the network device default secrets | 
**MacAddressOrSerialNumber** | **string** | The MAC address or serial number of the network device | 
**SecretName** | **string** | The name of the secret | 
**CreatedTimestamp** | **string** | The timestamp when the secret was created | 
**UpdatedTimestamp** | **string** | The timestamp when the secret was updated | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceDefaultSecrets

`func NewNetworkDeviceDefaultSecrets(id float32, siteId float32, macAddressOrSerialNumber string, secretName string, createdTimestamp string, updatedTimestamp string, ) *NetworkDeviceDefaultSecrets`

NewNetworkDeviceDefaultSecrets instantiates a new NetworkDeviceDefaultSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDefaultSecretsWithDefaults

`func NewNetworkDeviceDefaultSecretsWithDefaults() *NetworkDeviceDefaultSecrets`

NewNetworkDeviceDefaultSecretsWithDefaults instantiates a new NetworkDeviceDefaultSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDeviceDefaultSecrets) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceDefaultSecrets) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceDefaultSecrets) SetId(v float32)`

SetId sets Id field to given value.


### GetSiteId

`func (o *NetworkDeviceDefaultSecrets) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *NetworkDeviceDefaultSecrets) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *NetworkDeviceDefaultSecrets) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetMacAddressOrSerialNumber

`func (o *NetworkDeviceDefaultSecrets) GetMacAddressOrSerialNumber() string`

GetMacAddressOrSerialNumber returns the MacAddressOrSerialNumber field if non-nil, zero value otherwise.

### GetMacAddressOrSerialNumberOk

`func (o *NetworkDeviceDefaultSecrets) GetMacAddressOrSerialNumberOk() (*string, bool)`

GetMacAddressOrSerialNumberOk returns a tuple with the MacAddressOrSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressOrSerialNumber

`func (o *NetworkDeviceDefaultSecrets) SetMacAddressOrSerialNumber(v string)`

SetMacAddressOrSerialNumber sets MacAddressOrSerialNumber field to given value.


### GetSecretName

`func (o *NetworkDeviceDefaultSecrets) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *NetworkDeviceDefaultSecrets) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *NetworkDeviceDefaultSecrets) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetCreatedTimestamp

`func (o *NetworkDeviceDefaultSecrets) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkDeviceDefaultSecrets) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkDeviceDefaultSecrets) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkDeviceDefaultSecrets) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkDeviceDefaultSecrets) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkDeviceDefaultSecrets) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *NetworkDeviceDefaultSecrets) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceDefaultSecrets) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceDefaultSecrets) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceDefaultSecrets) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


