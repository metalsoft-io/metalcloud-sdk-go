# UserPromotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promotion** | Pointer to **string** | The promotion for the user | [optional] 
**RemovePromotion** | Pointer to **bool** | Whether to remove the promotion | [optional] [default to false]

## Methods

### NewUserPromotion

`func NewUserPromotion() *UserPromotion`

NewUserPromotion instantiates a new UserPromotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPromotionWithDefaults

`func NewUserPromotionWithDefaults() *UserPromotion`

NewUserPromotionWithDefaults instantiates a new UserPromotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromotion

`func (o *UserPromotion) GetPromotion() string`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *UserPromotion) GetPromotionOk() (*string, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *UserPromotion) SetPromotion(v string)`

SetPromotion sets Promotion field to given value.

### HasPromotion

`func (o *UserPromotion) HasPromotion() bool`

HasPromotion returns a boolean if a field has been set.

### GetRemovePromotion

`func (o *UserPromotion) GetRemovePromotion() bool`

GetRemovePromotion returns the RemovePromotion field if non-nil, zero value otherwise.

### GetRemovePromotionOk

`func (o *UserPromotion) GetRemovePromotionOk() (*bool, bool)`

GetRemovePromotionOk returns a tuple with the RemovePromotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePromotion

`func (o *UserPromotion) SetRemovePromotion(v bool)`

SetRemovePromotion sets RemovePromotion field to given value.

### HasRemovePromotion

`func (o *UserPromotion) HasRemovePromotion() bool`

HasRemovePromotion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


