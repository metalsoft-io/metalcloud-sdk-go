# AuthBsiConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **map[string]interface{}** | Platform service base URL | 
**Token** | **map[string]interface{}** | Authentication token for the platform service | 

## Methods

### NewAuthBsiConfigDto

`func NewAuthBsiConfigDto(host map[string]interface{}, token map[string]interface{}, ) *AuthBsiConfigDto`

NewAuthBsiConfigDto instantiates a new AuthBsiConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthBsiConfigDtoWithDefaults

`func NewAuthBsiConfigDtoWithDefaults() *AuthBsiConfigDto`

NewAuthBsiConfigDtoWithDefaults instantiates a new AuthBsiConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *AuthBsiConfigDto) GetHost() map[string]interface{}`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AuthBsiConfigDto) GetHostOk() (*map[string]interface{}, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AuthBsiConfigDto) SetHost(v map[string]interface{})`

SetHost sets Host field to given value.


### GetToken

`func (o *AuthBsiConfigDto) GetToken() map[string]interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthBsiConfigDto) GetTokenOk() (*map[string]interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthBsiConfigDto) SetToken(v map[string]interface{})`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


