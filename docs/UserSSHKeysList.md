# UserSSHKeysList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]UserSSHKeys**](UserSSHKeys.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewUserSSHKeysList

`func NewUserSSHKeysList(data []UserSSHKeys, ) *UserSSHKeysList`

NewUserSSHKeysList instantiates a new UserSSHKeysList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSSHKeysListWithDefaults

`func NewUserSSHKeysListWithDefaults() *UserSSHKeysList`

NewUserSSHKeysListWithDefaults instantiates a new UserSSHKeysList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserSSHKeysList) GetData() []UserSSHKeys`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserSSHKeysList) GetDataOk() (*[]UserSSHKeys, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserSSHKeysList) SetData(v []UserSSHKeys)`

SetData sets Data field to given value.


### GetLinks

`func (o *UserSSHKeysList) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSSHKeysList) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSSHKeysList) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserSSHKeysList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


