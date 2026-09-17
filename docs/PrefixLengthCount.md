# PrefixLengthCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrefixLength** | **int32** |  | 
**Count** | **int64** | How many free blocks of this prefix length there are | 

## Methods

### NewPrefixLengthCount

`func NewPrefixLengthCount(prefixLength int32, count int64, ) *PrefixLengthCount`

NewPrefixLengthCount instantiates a new PrefixLengthCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixLengthCountWithDefaults

`func NewPrefixLengthCountWithDefaults() *PrefixLengthCount`

NewPrefixLengthCountWithDefaults instantiates a new PrefixLengthCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefixLength

`func (o *PrefixLengthCount) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *PrefixLengthCount) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *PrefixLengthCount) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetCount

`func (o *PrefixLengthCount) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PrefixLengthCount) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PrefixLengthCount) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


