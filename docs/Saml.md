# Saml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryPoint** | **map[string]interface{}** | SAML Identity Provider SSO endpoint URL | 
**Issuer** | **map[string]interface{}** | Service provider entity ID / issuer sent in SAML requests | 
**CallbackUrl** | **map[string]interface{}** | Assertion Consumer Service (ACS) URL where the IdP posts the SAML response | 
**IdpCert** | **map[string]interface{}** | Identity provider X.509 certificate in PEM format | 
**SpCert** | Pointer to **string** | Service provider X.509 certificate in PEM format (auto-generated if not set) | [optional] 
**SpPrivateKey** | Pointer to **string** | Service provider private key in PEM format (auto-generated if not set) | [optional] 
**LogoutURL** | **map[string]interface{}** | Single Logout Service (SLS) URL | 
**ProfileMapping** | [**SamlProfileMapping**](SamlProfileMapping.md) | Mapping of SAML assertion attributes to user profile fields | 
**GroupsMapping** | [**[]GroupMapping**](GroupMapping.md) | Mapping of SAML group attribute values to application roles | 

## Methods

### NewSaml

`func NewSaml(entryPoint map[string]interface{}, issuer map[string]interface{}, callbackUrl map[string]interface{}, idpCert map[string]interface{}, logoutURL map[string]interface{}, profileMapping SamlProfileMapping, groupsMapping []GroupMapping, ) *Saml`

NewSaml instantiates a new Saml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlWithDefaults

`func NewSamlWithDefaults() *Saml`

NewSamlWithDefaults instantiates a new Saml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryPoint

`func (o *Saml) GetEntryPoint() map[string]interface{}`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *Saml) GetEntryPointOk() (*map[string]interface{}, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *Saml) SetEntryPoint(v map[string]interface{})`

SetEntryPoint sets EntryPoint field to given value.


### GetIssuer

`func (o *Saml) GetIssuer() map[string]interface{}`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Saml) GetIssuerOk() (*map[string]interface{}, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Saml) SetIssuer(v map[string]interface{})`

SetIssuer sets Issuer field to given value.


### GetCallbackUrl

`func (o *Saml) GetCallbackUrl() map[string]interface{}`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *Saml) GetCallbackUrlOk() (*map[string]interface{}, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *Saml) SetCallbackUrl(v map[string]interface{})`

SetCallbackUrl sets CallbackUrl field to given value.


### GetIdpCert

`func (o *Saml) GetIdpCert() map[string]interface{}`

GetIdpCert returns the IdpCert field if non-nil, zero value otherwise.

### GetIdpCertOk

`func (o *Saml) GetIdpCertOk() (*map[string]interface{}, bool)`

GetIdpCertOk returns a tuple with the IdpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCert

`func (o *Saml) SetIdpCert(v map[string]interface{})`

SetIdpCert sets IdpCert field to given value.


### GetSpCert

`func (o *Saml) GetSpCert() string`

GetSpCert returns the SpCert field if non-nil, zero value otherwise.

### GetSpCertOk

`func (o *Saml) GetSpCertOk() (*string, bool)`

GetSpCertOk returns a tuple with the SpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCert

`func (o *Saml) SetSpCert(v string)`

SetSpCert sets SpCert field to given value.

### HasSpCert

`func (o *Saml) HasSpCert() bool`

HasSpCert returns a boolean if a field has been set.

### GetSpPrivateKey

`func (o *Saml) GetSpPrivateKey() string`

GetSpPrivateKey returns the SpPrivateKey field if non-nil, zero value otherwise.

### GetSpPrivateKeyOk

`func (o *Saml) GetSpPrivateKeyOk() (*string, bool)`

GetSpPrivateKeyOk returns a tuple with the SpPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpPrivateKey

`func (o *Saml) SetSpPrivateKey(v string)`

SetSpPrivateKey sets SpPrivateKey field to given value.

### HasSpPrivateKey

`func (o *Saml) HasSpPrivateKey() bool`

HasSpPrivateKey returns a boolean if a field has been set.

### GetLogoutURL

`func (o *Saml) GetLogoutURL() map[string]interface{}`

GetLogoutURL returns the LogoutURL field if non-nil, zero value otherwise.

### GetLogoutURLOk

`func (o *Saml) GetLogoutURLOk() (*map[string]interface{}, bool)`

GetLogoutURLOk returns a tuple with the LogoutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutURL

`func (o *Saml) SetLogoutURL(v map[string]interface{})`

SetLogoutURL sets LogoutURL field to given value.


### GetProfileMapping

`func (o *Saml) GetProfileMapping() SamlProfileMapping`

GetProfileMapping returns the ProfileMapping field if non-nil, zero value otherwise.

### GetProfileMappingOk

`func (o *Saml) GetProfileMappingOk() (*SamlProfileMapping, bool)`

GetProfileMappingOk returns a tuple with the ProfileMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMapping

`func (o *Saml) SetProfileMapping(v SamlProfileMapping)`

SetProfileMapping sets ProfileMapping field to given value.


### GetGroupsMapping

`func (o *Saml) GetGroupsMapping() []GroupMapping`

GetGroupsMapping returns the GroupsMapping field if non-nil, zero value otherwise.

### GetGroupsMappingOk

`func (o *Saml) GetGroupsMappingOk() (*[]GroupMapping, bool)`

GetGroupsMappingOk returns a tuple with the GroupsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsMapping

`func (o *Saml) SetGroupsMapping(v []GroupMapping)`

SetGroupsMapping sets GroupsMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


