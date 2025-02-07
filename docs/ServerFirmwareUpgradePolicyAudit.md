# ServerFirmwareUpgradePolicyAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | **map[string]interface{}** | An object having as key the server component id and as value the target firmware version or null if not set | 

## Methods

### NewServerFirmwareUpgradePolicyAudit

`func NewServerFirmwareUpgradePolicyAudit(audit map[string]interface{}, ) *ServerFirmwareUpgradePolicyAudit`

NewServerFirmwareUpgradePolicyAudit instantiates a new ServerFirmwareUpgradePolicyAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareUpgradePolicyAuditWithDefaults

`func NewServerFirmwareUpgradePolicyAuditWithDefaults() *ServerFirmwareUpgradePolicyAudit`

NewServerFirmwareUpgradePolicyAuditWithDefaults instantiates a new ServerFirmwareUpgradePolicyAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ServerFirmwareUpgradePolicyAudit) GetAudit() map[string]interface{}`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ServerFirmwareUpgradePolicyAudit) GetAuditOk() (*map[string]interface{}, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ServerFirmwareUpgradePolicyAudit) SetAudit(v map[string]interface{})`

SetAudit sets Audit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


