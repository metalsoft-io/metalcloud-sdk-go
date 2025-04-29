# ServerRegistrationBiosProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | BIOS attribute name | 
**AttributeValue** | **string** | BIOS attribute value | 
**ServerVendor** | **string** | Server vendor type | 
**ServerModel** | Pointer to **string** | Server model | [optional] 

## Methods

### NewServerRegistrationBiosProfile

`func NewServerRegistrationBiosProfile(attributeName string, attributeValue string, serverVendor string, ) *ServerRegistrationBiosProfile`

NewServerRegistrationBiosProfile instantiates a new ServerRegistrationBiosProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRegistrationBiosProfileWithDefaults

`func NewServerRegistrationBiosProfileWithDefaults() *ServerRegistrationBiosProfile`

NewServerRegistrationBiosProfileWithDefaults instantiates a new ServerRegistrationBiosProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *ServerRegistrationBiosProfile) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ServerRegistrationBiosProfile) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ServerRegistrationBiosProfile) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetAttributeValue

`func (o *ServerRegistrationBiosProfile) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *ServerRegistrationBiosProfile) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *ServerRegistrationBiosProfile) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.


### GetServerVendor

`func (o *ServerRegistrationBiosProfile) GetServerVendor() string`

GetServerVendor returns the ServerVendor field if non-nil, zero value otherwise.

### GetServerVendorOk

`func (o *ServerRegistrationBiosProfile) GetServerVendorOk() (*string, bool)`

GetServerVendorOk returns a tuple with the ServerVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVendor

`func (o *ServerRegistrationBiosProfile) SetServerVendor(v string)`

SetServerVendor sets ServerVendor field to given value.


### GetServerModel

`func (o *ServerRegistrationBiosProfile) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *ServerRegistrationBiosProfile) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *ServerRegistrationBiosProfile) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *ServerRegistrationBiosProfile) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


