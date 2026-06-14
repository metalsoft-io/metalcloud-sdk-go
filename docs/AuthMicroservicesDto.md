# AuthMicroservicesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsi** | [**AuthBsiConfigDto**](AuthBsiConfigDto.md) | Platform service connection settings | 

## Methods

### NewAuthMicroservicesDto

`func NewAuthMicroservicesDto(bsi AuthBsiConfigDto, ) *AuthMicroservicesDto`

NewAuthMicroservicesDto instantiates a new AuthMicroservicesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMicroservicesDtoWithDefaults

`func NewAuthMicroservicesDtoWithDefaults() *AuthMicroservicesDto`

NewAuthMicroservicesDtoWithDefaults instantiates a new AuthMicroservicesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsi

`func (o *AuthMicroservicesDto) GetBsi() AuthBsiConfigDto`

GetBsi returns the Bsi field if non-nil, zero value otherwise.

### GetBsiOk

`func (o *AuthMicroservicesDto) GetBsiOk() (*AuthBsiConfigDto, bool)`

GetBsiOk returns a tuple with the Bsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsi

`func (o *AuthMicroservicesDto) SetBsi(v AuthBsiConfigDto)`

SetBsi sets Bsi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


