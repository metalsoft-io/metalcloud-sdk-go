# NetworkFabricInterconnectLinkValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | **int64** | The interconnect link ID | 
**NetworkDeviceId** | **int64** | The network device ID (switch) on this link | 
**CanActivate** | **bool** | Whether this link can be activated based on the switch configuration | 
**ResolvedVariables** | [**NetworkFabricInterconnectLinkValidationResolvedVariables**](NetworkFabricInterconnectLinkValidationResolvedVariables.md) | BGP variables resolved for the link&#39;s switch (custom_variables_json with column fallback) | 
**Errors** | **[]map[string]interface{}** | Raw class-validator errors against the combined switch DTO. Empty when canActivate is true. See class-validator ValidationError for shape. | 

## Methods

### NewNetworkFabricInterconnectLinkValidation

`func NewNetworkFabricInterconnectLinkValidation(linkId int64, networkDeviceId int64, canActivate bool, resolvedVariables NetworkFabricInterconnectLinkValidationResolvedVariables, errors []map[string]interface{}, ) *NetworkFabricInterconnectLinkValidation`

NewNetworkFabricInterconnectLinkValidation instantiates a new NetworkFabricInterconnectLinkValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectLinkValidationWithDefaults

`func NewNetworkFabricInterconnectLinkValidationWithDefaults() *NetworkFabricInterconnectLinkValidation`

NewNetworkFabricInterconnectLinkValidationWithDefaults instantiates a new NetworkFabricInterconnectLinkValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *NetworkFabricInterconnectLinkValidation) GetLinkId() int64`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *NetworkFabricInterconnectLinkValidation) GetLinkIdOk() (*int64, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *NetworkFabricInterconnectLinkValidation) SetLinkId(v int64)`

SetLinkId sets LinkId field to given value.


### GetNetworkDeviceId

`func (o *NetworkFabricInterconnectLinkValidation) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkFabricInterconnectLinkValidation) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkFabricInterconnectLinkValidation) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetCanActivate

`func (o *NetworkFabricInterconnectLinkValidation) GetCanActivate() bool`

GetCanActivate returns the CanActivate field if non-nil, zero value otherwise.

### GetCanActivateOk

`func (o *NetworkFabricInterconnectLinkValidation) GetCanActivateOk() (*bool, bool)`

GetCanActivateOk returns a tuple with the CanActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanActivate

`func (o *NetworkFabricInterconnectLinkValidation) SetCanActivate(v bool)`

SetCanActivate sets CanActivate field to given value.


### GetResolvedVariables

`func (o *NetworkFabricInterconnectLinkValidation) GetResolvedVariables() NetworkFabricInterconnectLinkValidationResolvedVariables`

GetResolvedVariables returns the ResolvedVariables field if non-nil, zero value otherwise.

### GetResolvedVariablesOk

`func (o *NetworkFabricInterconnectLinkValidation) GetResolvedVariablesOk() (*NetworkFabricInterconnectLinkValidationResolvedVariables, bool)`

GetResolvedVariablesOk returns a tuple with the ResolvedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedVariables

`func (o *NetworkFabricInterconnectLinkValidation) SetResolvedVariables(v NetworkFabricInterconnectLinkValidationResolvedVariables)`

SetResolvedVariables sets ResolvedVariables field to given value.


### GetErrors

`func (o *NetworkFabricInterconnectLinkValidation) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *NetworkFabricInterconnectLinkValidation) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *NetworkFabricInterconnectLinkValidation) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


