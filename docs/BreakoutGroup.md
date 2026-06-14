# BreakoutGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfInterfaces** | **int32** | Number of child interfaces this group produces (OpenConfig num-breakouts). | 
**Speed** | Pointer to **NullableString** | Per-child-interface speed (OpenConfig breakout-speed). Optional: omit or set null to split the port by count only (Cumulus \&quot;4x\&quot; form) and let the device derive the child speed from the port’s lane capability. When provided, the per-child speed is forced. | [optional] 

## Methods

### NewBreakoutGroup

`func NewBreakoutGroup(numberOfInterfaces int32, ) *BreakoutGroup`

NewBreakoutGroup instantiates a new BreakoutGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakoutGroupWithDefaults

`func NewBreakoutGroupWithDefaults() *BreakoutGroup`

NewBreakoutGroupWithDefaults instantiates a new BreakoutGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfInterfaces

`func (o *BreakoutGroup) GetNumberOfInterfaces() int32`

GetNumberOfInterfaces returns the NumberOfInterfaces field if non-nil, zero value otherwise.

### GetNumberOfInterfacesOk

`func (o *BreakoutGroup) GetNumberOfInterfacesOk() (*int32, bool)`

GetNumberOfInterfacesOk returns a tuple with the NumberOfInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfInterfaces

`func (o *BreakoutGroup) SetNumberOfInterfaces(v int32)`

SetNumberOfInterfaces sets NumberOfInterfaces field to given value.


### GetSpeed

`func (o *BreakoutGroup) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *BreakoutGroup) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *BreakoutGroup) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *BreakoutGroup) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *BreakoutGroup) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *BreakoutGroup) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


