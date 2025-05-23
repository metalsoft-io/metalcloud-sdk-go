# VMInstanceGroupNetworkConnectionsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VMInstanceGroupNetworkConnection**](VMInstanceGroupNetworkConnection.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewVMInstanceGroupNetworkConnectionsList

`func NewVMInstanceGroupNetworkConnectionsList(data []VMInstanceGroupNetworkConnection, ) *VMInstanceGroupNetworkConnectionsList`

NewVMInstanceGroupNetworkConnectionsList instantiates a new VMInstanceGroupNetworkConnectionsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceGroupNetworkConnectionsListWithDefaults

`func NewVMInstanceGroupNetworkConnectionsListWithDefaults() *VMInstanceGroupNetworkConnectionsList`

NewVMInstanceGroupNetworkConnectionsListWithDefaults instantiates a new VMInstanceGroupNetworkConnectionsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VMInstanceGroupNetworkConnectionsList) GetData() []VMInstanceGroupNetworkConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VMInstanceGroupNetworkConnectionsList) GetDataOk() (*[]VMInstanceGroupNetworkConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VMInstanceGroupNetworkConnectionsList) SetData(v []VMInstanceGroupNetworkConnection)`

SetData sets Data field to given value.


### GetLinks

`func (o *VMInstanceGroupNetworkConnectionsList) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMInstanceGroupNetworkConnectionsList) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMInstanceGroupNetworkConnectionsList) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VMInstanceGroupNetworkConnectionsList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


