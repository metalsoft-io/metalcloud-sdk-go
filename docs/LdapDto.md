# LdapDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**LdapServerDto**](LdapServerDto.md) | LDAP server connection settings | 
**ProfileMapping** | [**ProfileMappingDto**](ProfileMappingDto.md) | Mapping of LDAP attributes to user profile fields | 
**GroupsMapping** | [**[]GroupMappingDto**](GroupMappingDto.md) | Mapping of LDAP groups to application roles | 

## Methods

### NewLdapDto

`func NewLdapDto(server LdapServerDto, profileMapping ProfileMappingDto, groupsMapping []GroupMappingDto, ) *LdapDto`

NewLdapDto instantiates a new LdapDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDtoWithDefaults

`func NewLdapDtoWithDefaults() *LdapDto`

NewLdapDtoWithDefaults instantiates a new LdapDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *LdapDto) GetServer() LdapServerDto`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapDto) GetServerOk() (*LdapServerDto, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapDto) SetServer(v LdapServerDto)`

SetServer sets Server field to given value.


### GetProfileMapping

`func (o *LdapDto) GetProfileMapping() ProfileMappingDto`

GetProfileMapping returns the ProfileMapping field if non-nil, zero value otherwise.

### GetProfileMappingOk

`func (o *LdapDto) GetProfileMappingOk() (*ProfileMappingDto, bool)`

GetProfileMappingOk returns a tuple with the ProfileMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMapping

`func (o *LdapDto) SetProfileMapping(v ProfileMappingDto)`

SetProfileMapping sets ProfileMapping field to given value.


### GetGroupsMapping

`func (o *LdapDto) GetGroupsMapping() []GroupMappingDto`

GetGroupsMapping returns the GroupsMapping field if non-nil, zero value otherwise.

### GetGroupsMappingOk

`func (o *LdapDto) GetGroupsMappingOk() (*[]GroupMappingDto, bool)`

GetGroupsMappingOk returns a tuple with the GroupsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsMapping

`func (o *LdapDto) SetGroupsMapping(v []GroupMappingDto)`

SetGroupsMapping sets GroupsMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


