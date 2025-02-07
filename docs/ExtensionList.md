# ExtensionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExtensionInfo**](ExtensionInfo.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewExtensionList

`func NewExtensionList(data []ExtensionInfo, ) *ExtensionList`

NewExtensionList instantiates a new ExtensionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionListWithDefaults

`func NewExtensionListWithDefaults() *ExtensionList`

NewExtensionListWithDefaults instantiates a new ExtensionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExtensionList) GetData() []ExtensionInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExtensionList) GetDataOk() (*[]ExtensionInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExtensionList) SetData(v []ExtensionInfo)`

SetData sets Data field to given value.


### GetLinks

`func (o *ExtensionList) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtensionList) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtensionList) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExtensionList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


