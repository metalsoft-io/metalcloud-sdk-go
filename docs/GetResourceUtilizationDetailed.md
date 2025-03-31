# GetResourceUtilizationDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIdOwner** | **float32** | User ID of the owner | 
**StartTimestamp** | **time.Time** | Start timestamp for the resource utilization | 
**EndTimestamp** | **time.Time** | End timestamp for the resource utilization | 
**InfrastructureIds** | Pointer to **[]float32** | List of infrastructure IDs | [optional] 
**SiteIds** | Pointer to **[]float32** | List of site IDs | [optional] 

## Methods

### NewGetResourceUtilizationDetailed

`func NewGetResourceUtilizationDetailed(userIdOwner float32, startTimestamp time.Time, endTimestamp time.Time, ) *GetResourceUtilizationDetailed`

NewGetResourceUtilizationDetailed instantiates a new GetResourceUtilizationDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceUtilizationDetailedWithDefaults

`func NewGetResourceUtilizationDetailedWithDefaults() *GetResourceUtilizationDetailed`

NewGetResourceUtilizationDetailedWithDefaults instantiates a new GetResourceUtilizationDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIdOwner

`func (o *GetResourceUtilizationDetailed) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *GetResourceUtilizationDetailed) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *GetResourceUtilizationDetailed) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.


### GetStartTimestamp

`func (o *GetResourceUtilizationDetailed) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *GetResourceUtilizationDetailed) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *GetResourceUtilizationDetailed) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *GetResourceUtilizationDetailed) GetEndTimestamp() time.Time`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *GetResourceUtilizationDetailed) GetEndTimestampOk() (*time.Time, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *GetResourceUtilizationDetailed) SetEndTimestamp(v time.Time)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetInfrastructureIds

`func (o *GetResourceUtilizationDetailed) GetInfrastructureIds() []float32`

GetInfrastructureIds returns the InfrastructureIds field if non-nil, zero value otherwise.

### GetInfrastructureIdsOk

`func (o *GetResourceUtilizationDetailed) GetInfrastructureIdsOk() (*[]float32, bool)`

GetInfrastructureIdsOk returns a tuple with the InfrastructureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureIds

`func (o *GetResourceUtilizationDetailed) SetInfrastructureIds(v []float32)`

SetInfrastructureIds sets InfrastructureIds field to given value.

### HasInfrastructureIds

`func (o *GetResourceUtilizationDetailed) HasInfrastructureIds() bool`

HasInfrastructureIds returns a boolean if a field has been set.

### GetSiteIds

`func (o *GetResourceUtilizationDetailed) GetSiteIds() []float32`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *GetResourceUtilizationDetailed) GetSiteIdsOk() (*[]float32, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *GetResourceUtilizationDetailed) SetSiteIds(v []float32)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *GetResourceUtilizationDetailed) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


