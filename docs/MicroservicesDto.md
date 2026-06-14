# MicroservicesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsi** | [**BsiDto**](BsiDto.md) | Platform service connection settings | 
**Auth** | [**AuthDto**](AuthDto.md) | Authentication microservice settings | 

## Methods

### NewMicroservicesDto

`func NewMicroservicesDto(bsi BsiDto, auth AuthDto, ) *MicroservicesDto`

NewMicroservicesDto instantiates a new MicroservicesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicroservicesDtoWithDefaults

`func NewMicroservicesDtoWithDefaults() *MicroservicesDto`

NewMicroservicesDtoWithDefaults instantiates a new MicroservicesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsi

`func (o *MicroservicesDto) GetBsi() BsiDto`

GetBsi returns the Bsi field if non-nil, zero value otherwise.

### GetBsiOk

`func (o *MicroservicesDto) GetBsiOk() (*BsiDto, bool)`

GetBsiOk returns a tuple with the Bsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsi

`func (o *MicroservicesDto) SetBsi(v BsiDto)`

SetBsi sets Bsi field to given value.


### GetAuth

`func (o *MicroservicesDto) GetAuth() AuthDto`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *MicroservicesDto) GetAuthOk() (*AuthDto, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *MicroservicesDto) SetAuth(v AuthDto)`

SetAuth sets Auth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


