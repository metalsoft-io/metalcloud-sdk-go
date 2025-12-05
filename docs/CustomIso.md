# CustomIso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The ID of the custom ISO | 
**UserIdOwner** | Pointer to **float32** | The ID of the user who owns the custom ISO. If public then null | [optional] 
**ModifiedBy** | Pointer to **float32** | The ID of the user who last edited the custom ISO | [optional] 
**Label** | **string** | The label of the custom ISO | 
**Name** | **string** | The name of the custom ISO | 
**Type** | **string** | The type of the custom ISO | 
**IsPublic** | **float32** | Flag to indicate if the custom ISO is public | 
**AccessUrl** | **string** | The access URL of the custom ISO | 
**Username** | Pointer to **string** | The username used to access the custom ISO | [optional] 
**CreatedTimestamp** | **string** | The timestamp when the custom ISO was created | 
**UpdatedTimestamp** | **string** | The timestamp when the custom ISO was last updated | 
**ImageUrl** | Pointer to **string** | The URL to the image of the custom ISO for mounting on the server | [optional] 

## Methods

### NewCustomIso

`func NewCustomIso(id float32, label string, name string, type_ string, isPublic float32, accessUrl string, createdTimestamp string, updatedTimestamp string, ) *CustomIso`

NewCustomIso instantiates a new CustomIso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomIsoWithDefaults

`func NewCustomIsoWithDefaults() *CustomIso`

NewCustomIsoWithDefaults instantiates a new CustomIso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomIso) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomIso) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomIso) SetId(v float32)`

SetId sets Id field to given value.


### GetUserIdOwner

`func (o *CustomIso) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *CustomIso) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *CustomIso) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *CustomIso) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetModifiedBy

`func (o *CustomIso) GetModifiedBy() float32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CustomIso) GetModifiedByOk() (*float32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CustomIso) SetModifiedBy(v float32)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CustomIso) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetLabel

`func (o *CustomIso) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomIso) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomIso) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CustomIso) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomIso) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomIso) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CustomIso) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomIso) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomIso) SetType(v string)`

SetType sets Type field to given value.


### GetIsPublic

`func (o *CustomIso) GetIsPublic() float32`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CustomIso) GetIsPublicOk() (*float32, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CustomIso) SetIsPublic(v float32)`

SetIsPublic sets IsPublic field to given value.


### GetAccessUrl

`func (o *CustomIso) GetAccessUrl() string`

GetAccessUrl returns the AccessUrl field if non-nil, zero value otherwise.

### GetAccessUrlOk

`func (o *CustomIso) GetAccessUrlOk() (*string, bool)`

GetAccessUrlOk returns a tuple with the AccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrl

`func (o *CustomIso) SetAccessUrl(v string)`

SetAccessUrl sets AccessUrl field to given value.


### GetUsername

`func (o *CustomIso) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CustomIso) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CustomIso) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CustomIso) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *CustomIso) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *CustomIso) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *CustomIso) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *CustomIso) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *CustomIso) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *CustomIso) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetImageUrl

`func (o *CustomIso) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CustomIso) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CustomIso) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CustomIso) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


