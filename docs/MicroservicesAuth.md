# MicroservicesAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamlHostname** | Pointer to **map[string]interface{}** | Base URL of the SAML identity provider | [optional] 

## Methods

### NewMicroservicesAuth

`func NewMicroservicesAuth() *MicroservicesAuth`

NewMicroservicesAuth instantiates a new MicroservicesAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicroservicesAuthWithDefaults

`func NewMicroservicesAuthWithDefaults() *MicroservicesAuth`

NewMicroservicesAuthWithDefaults instantiates a new MicroservicesAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamlHostname

`func (o *MicroservicesAuth) GetSamlHostname() map[string]interface{}`

GetSamlHostname returns the SamlHostname field if non-nil, zero value otherwise.

### GetSamlHostnameOk

`func (o *MicroservicesAuth) GetSamlHostnameOk() (*map[string]interface{}, bool)`

GetSamlHostnameOk returns a tuple with the SamlHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlHostname

`func (o *MicroservicesAuth) SetSamlHostname(v map[string]interface{})`

SetSamlHostname sets SamlHostname field to given value.

### HasSamlHostname

`func (o *MicroservicesAuth) HasSamlHostname() bool`

HasSamlHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


