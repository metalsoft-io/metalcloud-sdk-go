# ServerInstanceGroupOSInstallationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Product Instance ID. | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Product Instance. | [optional] 
**ServerGroupName** | Pointer to **string** |  | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Object containing custom variables and variable overrides. | [optional] 
**IsVmGroup** | **int32** | Flag to indicate if the Server Instance Group is belongs to a VM. | 

## Methods

### NewServerInstanceGroupOSInstallationData

`func NewServerInstanceGroupOSInstallationData(id int32, label string, isVmGroup int32, ) *ServerInstanceGroupOSInstallationData`

NewServerInstanceGroupOSInstallationData instantiates a new ServerInstanceGroupOSInstallationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupOSInstallationDataWithDefaults

`func NewServerInstanceGroupOSInstallationDataWithDefaults() *ServerInstanceGroupOSInstallationData`

NewServerInstanceGroupOSInstallationDataWithDefaults instantiates a new ServerInstanceGroupOSInstallationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceGroupOSInstallationData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceGroupOSInstallationData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceGroupOSInstallationData) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *ServerInstanceGroupOSInstallationData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupOSInstallationData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupOSInstallationData) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomainPermanent

`func (o *ServerInstanceGroupOSInstallationData) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstanceGroupOSInstallationData) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstanceGroupOSInstallationData) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstanceGroupOSInstallationData) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetServerGroupName

`func (o *ServerInstanceGroupOSInstallationData) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *ServerInstanceGroupOSInstallationData) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *ServerInstanceGroupOSInstallationData) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.

### HasServerGroupName

`func (o *ServerInstanceGroupOSInstallationData) HasServerGroupName() bool`

HasServerGroupName returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceGroupOSInstallationData) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceGroupOSInstallationData) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceGroupOSInstallationData) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceGroupOSInstallationData) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetIsVmGroup

`func (o *ServerInstanceGroupOSInstallationData) GetIsVmGroup() int32`

GetIsVmGroup returns the IsVmGroup field if non-nil, zero value otherwise.

### GetIsVmGroupOk

`func (o *ServerInstanceGroupOSInstallationData) GetIsVmGroupOk() (*int32, bool)`

GetIsVmGroupOk returns a tuple with the IsVmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVmGroup

`func (o *ServerInstanceGroupOSInstallationData) SetIsVmGroup(v int32)`

SetIsVmGroup sets IsVmGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


