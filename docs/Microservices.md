# Microservices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsi** | [**MicroservicesBsi**](MicroservicesBsi.md) | Platform service connection settings | 
**Auth** | [**MicroservicesAuth**](MicroservicesAuth.md) | Authentication microservice settings | 

## Methods

### NewMicroservices

`func NewMicroservices(bsi MicroservicesBsi, auth MicroservicesAuth, ) *Microservices`

NewMicroservices instantiates a new Microservices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicroservicesWithDefaults

`func NewMicroservicesWithDefaults() *Microservices`

NewMicroservicesWithDefaults instantiates a new Microservices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsi

`func (o *Microservices) GetBsi() MicroservicesBsi`

GetBsi returns the Bsi field if non-nil, zero value otherwise.

### GetBsiOk

`func (o *Microservices) GetBsiOk() (*MicroservicesBsi, bool)`

GetBsiOk returns a tuple with the Bsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsi

`func (o *Microservices) SetBsi(v MicroservicesBsi)`

SetBsi sets Bsi field to given value.


### GetAuth

`func (o *Microservices) GetAuth() MicroservicesAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Microservices) GetAuthOk() (*MicroservicesAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Microservices) SetAuth(v MicroservicesAuth)`

SetAuth sets Auth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


