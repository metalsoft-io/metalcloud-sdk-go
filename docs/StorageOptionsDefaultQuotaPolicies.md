# StorageOptionsDefaultQuotaPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realm** | Pointer to [**StorageOptionsDefaultQuotaPolicySettings**](StorageOptionsDefaultQuotaPolicySettings.md) | Default Quota policy settings for the realm entity (applicable to Pure Flash Array. 1 Pure FA realm corresponds to 1 MetalSoft infrastructure) | [optional] 

## Methods

### NewStorageOptionsDefaultQuotaPolicies

`func NewStorageOptionsDefaultQuotaPolicies() *StorageOptionsDefaultQuotaPolicies`

NewStorageOptionsDefaultQuotaPolicies instantiates a new StorageOptionsDefaultQuotaPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageOptionsDefaultQuotaPoliciesWithDefaults

`func NewStorageOptionsDefaultQuotaPoliciesWithDefaults() *StorageOptionsDefaultQuotaPolicies`

NewStorageOptionsDefaultQuotaPoliciesWithDefaults instantiates a new StorageOptionsDefaultQuotaPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealm

`func (o *StorageOptionsDefaultQuotaPolicies) GetRealm() StorageOptionsDefaultQuotaPolicySettings`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *StorageOptionsDefaultQuotaPolicies) GetRealmOk() (*StorageOptionsDefaultQuotaPolicySettings, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *StorageOptionsDefaultQuotaPolicies) SetRealm(v StorageOptionsDefaultQuotaPolicySettings)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *StorageOptionsDefaultQuotaPolicies) HasRealm() bool`

HasRealm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


