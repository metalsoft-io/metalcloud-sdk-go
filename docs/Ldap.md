# Ldap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**LdapServer**](LdapServer.md) | LDAP server connection settings | 
**ProfileMapping** | [**LdapProfileMapping**](LdapProfileMapping.md) | Mapping of LDAP attributes to user profile fields | 
**GroupsMapping** | [**[]GroupMapping**](GroupMapping.md) | Mapping of LDAP groups to application roles | 

## Methods

### NewLdap

`func NewLdap(server LdapServer, profileMapping LdapProfileMapping, groupsMapping []GroupMapping, ) *Ldap`

NewLdap instantiates a new Ldap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapWithDefaults

`func NewLdapWithDefaults() *Ldap`

NewLdapWithDefaults instantiates a new Ldap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *Ldap) GetServer() LdapServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Ldap) GetServerOk() (*LdapServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Ldap) SetServer(v LdapServer)`

SetServer sets Server field to given value.


### GetProfileMapping

`func (o *Ldap) GetProfileMapping() LdapProfileMapping`

GetProfileMapping returns the ProfileMapping field if non-nil, zero value otherwise.

### GetProfileMappingOk

`func (o *Ldap) GetProfileMappingOk() (*LdapProfileMapping, bool)`

GetProfileMappingOk returns a tuple with the ProfileMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMapping

`func (o *Ldap) SetProfileMapping(v LdapProfileMapping)`

SetProfileMapping sets ProfileMapping field to given value.


### GetGroupsMapping

`func (o *Ldap) GetGroupsMapping() []GroupMapping`

GetGroupsMapping returns the GroupsMapping field if non-nil, zero value otherwise.

### GetGroupsMappingOk

`func (o *Ldap) GetGroupsMappingOk() (*[]GroupMapping, bool)`

GetGroupsMappingOk returns a tuple with the GroupsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsMapping

`func (o *Ldap) SetGroupsMapping(v []GroupMapping)`

SetGroupsMapping sets GroupsMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


