# VMPoolHosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | VM Pool Host ID | 
**Name** | **string** | Name of the VM Pool Host | 
**Address** | Pointer to **string** | Address of the VM Pool Host | [optional] 
**Port** | Pointer to **float32** | Port of the VM Pool Host | [optional] 
**PoolId** | Pointer to **float32** | VM Pool ID | [optional] 
**Roles** | Pointer to **[]string** | Roles of the VM Pool Host | [optional] 
**FailureDomain** | Pointer to **string** | Failure domain of the VM Pool Host | [optional] 
**Architecture** | Pointer to **string** | Architecture of the VM Pool Host | [optional] 
**Database** | Pointer to **float32** | Flag specifying if the VM Pool Host is database | [optional] 
**Status** | Pointer to **string** | Status of the VM Pool Host | [optional] 
**Description** | Pointer to **string** | Description of the VM Pool Host | [optional] 
**UpdatedTimestamp** | **string** | Timestamp when the VM Pool Host was updated | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewVMPoolHosts

`func NewVMPoolHosts(id float32, name string, updatedTimestamp string, links map[string]interface{}, ) *VMPoolHosts`

NewVMPoolHosts instantiates a new VMPoolHosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolHostsWithDefaults

`func NewVMPoolHostsWithDefaults() *VMPoolHosts`

NewVMPoolHostsWithDefaults instantiates a new VMPoolHosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMPoolHosts) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMPoolHosts) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMPoolHosts) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *VMPoolHosts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMPoolHosts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMPoolHosts) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *VMPoolHosts) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VMPoolHosts) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VMPoolHosts) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VMPoolHosts) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *VMPoolHosts) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VMPoolHosts) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VMPoolHosts) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VMPoolHosts) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPoolId

`func (o *VMPoolHosts) GetPoolId() float32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *VMPoolHosts) GetPoolIdOk() (*float32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *VMPoolHosts) SetPoolId(v float32)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *VMPoolHosts) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetRoles

`func (o *VMPoolHosts) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *VMPoolHosts) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *VMPoolHosts) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *VMPoolHosts) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetFailureDomain

`func (o *VMPoolHosts) GetFailureDomain() string`

GetFailureDomain returns the FailureDomain field if non-nil, zero value otherwise.

### GetFailureDomainOk

`func (o *VMPoolHosts) GetFailureDomainOk() (*string, bool)`

GetFailureDomainOk returns a tuple with the FailureDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomain

`func (o *VMPoolHosts) SetFailureDomain(v string)`

SetFailureDomain sets FailureDomain field to given value.

### HasFailureDomain

`func (o *VMPoolHosts) HasFailureDomain() bool`

HasFailureDomain returns a boolean if a field has been set.

### GetArchitecture

`func (o *VMPoolHosts) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *VMPoolHosts) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *VMPoolHosts) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *VMPoolHosts) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetDatabase

`func (o *VMPoolHosts) GetDatabase() float32`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *VMPoolHosts) GetDatabaseOk() (*float32, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *VMPoolHosts) SetDatabase(v float32)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *VMPoolHosts) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetStatus

`func (o *VMPoolHosts) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VMPoolHosts) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VMPoolHosts) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VMPoolHosts) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *VMPoolHosts) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VMPoolHosts) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VMPoolHosts) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VMPoolHosts) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *VMPoolHosts) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMPoolHosts) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMPoolHosts) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *VMPoolHosts) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMPoolHosts) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMPoolHosts) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


