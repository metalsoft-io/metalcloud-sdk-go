# PartialLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address of the site | [optional] 
**Latitude** | Pointer to **float32** | Latitude of the site | [optional] 
**Longitude** | Pointer to **float32** | Longitude of the site | [optional] 

## Methods

### NewPartialLocation

`func NewPartialLocation() *PartialLocation`

NewPartialLocation instantiates a new PartialLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialLocationWithDefaults

`func NewPartialLocationWithDefaults() *PartialLocation`

NewPartialLocationWithDefaults instantiates a new PartialLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PartialLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartialLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartialLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PartialLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *PartialLocation) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PartialLocation) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PartialLocation) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PartialLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *PartialLocation) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PartialLocation) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PartialLocation) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PartialLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


