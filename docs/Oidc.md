# Oidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerUrl** | **map[string]interface{}** | OIDC issuer URL — the discovery document is fetched from &lt;issuerUrl&gt;/.well-known/openid-configuration | 
**ClientId** | **map[string]interface{}** | OAuth2/OIDC client identifier registered at the identity provider | 
**ClientSecret** | **string** | OAuth2/OIDC client secret (write-only; never returned by the API) | 
**Scopes** | **map[string]interface{}** | Space-separated OAuth2 scopes requested at login | 
**RedirectUrl** | **map[string]interface{}** | Redirect URI registered at the identity provider (the gateway OIDC callback URL) | 
**PostLogoutRedirectUrl** | Pointer to **string** | URL the identity provider redirects to after RP-initiated logout | [optional] 
**Prompt** | Pointer to **map[string]interface{}** | OIDC prompt parameter sent on login requests (&#39;login&#39; forces re-authentication at the IdP) | [optional] 
**ProfileMapping** | [**OidcProfileMapping**](OidcProfileMapping.md) | Mapping of ID token claims to user profile fields | 
**GroupsMapping** | [**[]GroupMapping**](GroupMapping.md) | Mapping of groups claim values to application roles | 

## Methods

### NewOidc

`func NewOidc(issuerUrl map[string]interface{}, clientId map[string]interface{}, clientSecret string, scopes map[string]interface{}, redirectUrl map[string]interface{}, profileMapping OidcProfileMapping, groupsMapping []GroupMapping, ) *Oidc`

NewOidc instantiates a new Oidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcWithDefaults

`func NewOidcWithDefaults() *Oidc`

NewOidcWithDefaults instantiates a new Oidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerUrl

`func (o *Oidc) GetIssuerUrl() map[string]interface{}`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *Oidc) GetIssuerUrlOk() (*map[string]interface{}, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *Oidc) SetIssuerUrl(v map[string]interface{})`

SetIssuerUrl sets IssuerUrl field to given value.


### GetClientId

`func (o *Oidc) GetClientId() map[string]interface{}`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Oidc) GetClientIdOk() (*map[string]interface{}, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Oidc) SetClientId(v map[string]interface{})`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *Oidc) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Oidc) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Oidc) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetScopes

`func (o *Oidc) GetScopes() map[string]interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Oidc) GetScopesOk() (*map[string]interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Oidc) SetScopes(v map[string]interface{})`

SetScopes sets Scopes field to given value.


### GetRedirectUrl

`func (o *Oidc) GetRedirectUrl() map[string]interface{}`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *Oidc) GetRedirectUrlOk() (*map[string]interface{}, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *Oidc) SetRedirectUrl(v map[string]interface{})`

SetRedirectUrl sets RedirectUrl field to given value.


### GetPostLogoutRedirectUrl

`func (o *Oidc) GetPostLogoutRedirectUrl() string`

GetPostLogoutRedirectUrl returns the PostLogoutRedirectUrl field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrlOk

`func (o *Oidc) GetPostLogoutRedirectUrlOk() (*string, bool)`

GetPostLogoutRedirectUrlOk returns a tuple with the PostLogoutRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUrl

`func (o *Oidc) SetPostLogoutRedirectUrl(v string)`

SetPostLogoutRedirectUrl sets PostLogoutRedirectUrl field to given value.

### HasPostLogoutRedirectUrl

`func (o *Oidc) HasPostLogoutRedirectUrl() bool`

HasPostLogoutRedirectUrl returns a boolean if a field has been set.

### GetPrompt

`func (o *Oidc) GetPrompt() map[string]interface{}`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *Oidc) GetPromptOk() (*map[string]interface{}, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *Oidc) SetPrompt(v map[string]interface{})`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *Oidc) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetProfileMapping

`func (o *Oidc) GetProfileMapping() OidcProfileMapping`

GetProfileMapping returns the ProfileMapping field if non-nil, zero value otherwise.

### GetProfileMappingOk

`func (o *Oidc) GetProfileMappingOk() (*OidcProfileMapping, bool)`

GetProfileMappingOk returns a tuple with the ProfileMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMapping

`func (o *Oidc) SetProfileMapping(v OidcProfileMapping)`

SetProfileMapping sets ProfileMapping field to given value.


### GetGroupsMapping

`func (o *Oidc) GetGroupsMapping() []GroupMapping`

GetGroupsMapping returns the GroupsMapping field if non-nil, zero value otherwise.

### GetGroupsMappingOk

`func (o *Oidc) GetGroupsMappingOk() (*[]GroupMapping, bool)`

GetGroupsMappingOk returns a tuple with the GroupsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsMapping

`func (o *Oidc) SetGroupsMapping(v []GroupMapping)`

SetGroupsMapping sets GroupsMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


