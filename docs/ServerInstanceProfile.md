# ServerInstanceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Server profile ID. | 
**Revision** | **float32** | Revision number | 
**OwnerId** | Pointer to **int32** | The Server profile owner ID. Can be null if the profile is public. Will be set to the currently logged user. | [optional] 
**Label** | Pointer to **string** | The Server profile label. Will be automatically generated if not provided. | [optional] 
**NetworkProfiles** | Pointer to [**[]ServerInstanceProfileNetworkProfilesInner**](ServerInstanceProfileNetworkProfilesInner.md) | Network profiles mapping for each network in this infrastructure. Changes to this configuration will be duplicated on each vm-instance of this group. | [optional] 
**NetworkInterfaces** | Pointer to [**ServerInstanceProfileNetworkInterfaces**](ServerInstanceProfileNetworkInterfaces.md) |  | [optional] 
**TemplateId** | Pointer to **int32** | The template id of the operating system to deploy on the Server. Can be null in which case no OS will be deployed but all operations will continue as normal. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerInstanceProfile

`func NewServerInstanceProfile(id int32, revision float32, ) *ServerInstanceProfile`

NewServerInstanceProfile instantiates a new ServerInstanceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceProfileWithDefaults

`func NewServerInstanceProfileWithDefaults() *ServerInstanceProfile`

NewServerInstanceProfileWithDefaults instantiates a new ServerInstanceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ServerInstanceProfile) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceProfile) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceProfile) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetOwnerId

`func (o *ServerInstanceProfile) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ServerInstanceProfile) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ServerInstanceProfile) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ServerInstanceProfile) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetLabel

`func (o *ServerInstanceProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNetworkProfiles

`func (o *ServerInstanceProfile) GetNetworkProfiles() []ServerInstanceProfileNetworkProfilesInner`

GetNetworkProfiles returns the NetworkProfiles field if non-nil, zero value otherwise.

### GetNetworkProfilesOk

`func (o *ServerInstanceProfile) GetNetworkProfilesOk() (*[]ServerInstanceProfileNetworkProfilesInner, bool)`

GetNetworkProfilesOk returns a tuple with the NetworkProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfiles

`func (o *ServerInstanceProfile) SetNetworkProfiles(v []ServerInstanceProfileNetworkProfilesInner)`

SetNetworkProfiles sets NetworkProfiles field to given value.

### HasNetworkProfiles

`func (o *ServerInstanceProfile) HasNetworkProfiles() bool`

HasNetworkProfiles returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ServerInstanceProfile) GetNetworkInterfaces() ServerInstanceProfileNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ServerInstanceProfile) GetNetworkInterfacesOk() (*ServerInstanceProfileNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ServerInstanceProfile) SetNetworkInterfaces(v ServerInstanceProfileNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ServerInstanceProfile) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetTemplateId

`func (o *ServerInstanceProfile) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ServerInstanceProfile) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ServerInstanceProfile) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ServerInstanceProfile) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetLinks

`func (o *ServerInstanceProfile) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstanceProfile) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstanceProfile) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerInstanceProfile) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


