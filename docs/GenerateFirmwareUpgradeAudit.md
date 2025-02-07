# GenerateFirmwareUpgradeAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIds** | **[]float32** | The list of server ids for which firmware upgrade audit should be generated. | 
**BaselineId** | Pointer to **float32** | Filter the available firmware upgrades using the specified baseline id. | [optional] 
**OsTemplateLabel** | Pointer to **string** | Filter the available firmware upgrades using the specified os template label. | [optional] 
**ServerTypeName** | Pointer to **string** | Filter the available firmware upgrades using the specified server type name. | [optional] 
**SiteId** | Pointer to **float32** | Filter the available firmware upgrades using the specified site id. | [optional] 

## Methods

### NewGenerateFirmwareUpgradeAudit

`func NewGenerateFirmwareUpgradeAudit(serverIds []float32, ) *GenerateFirmwareUpgradeAudit`

NewGenerateFirmwareUpgradeAudit instantiates a new GenerateFirmwareUpgradeAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateFirmwareUpgradeAuditWithDefaults

`func NewGenerateFirmwareUpgradeAuditWithDefaults() *GenerateFirmwareUpgradeAudit`

NewGenerateFirmwareUpgradeAuditWithDefaults instantiates a new GenerateFirmwareUpgradeAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerIds

`func (o *GenerateFirmwareUpgradeAudit) GetServerIds() []float32`

GetServerIds returns the ServerIds field if non-nil, zero value otherwise.

### GetServerIdsOk

`func (o *GenerateFirmwareUpgradeAudit) GetServerIdsOk() (*[]float32, bool)`

GetServerIdsOk returns a tuple with the ServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIds

`func (o *GenerateFirmwareUpgradeAudit) SetServerIds(v []float32)`

SetServerIds sets ServerIds field to given value.


### GetBaselineId

`func (o *GenerateFirmwareUpgradeAudit) GetBaselineId() float32`

GetBaselineId returns the BaselineId field if non-nil, zero value otherwise.

### GetBaselineIdOk

`func (o *GenerateFirmwareUpgradeAudit) GetBaselineIdOk() (*float32, bool)`

GetBaselineIdOk returns a tuple with the BaselineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineId

`func (o *GenerateFirmwareUpgradeAudit) SetBaselineId(v float32)`

SetBaselineId sets BaselineId field to given value.

### HasBaselineId

`func (o *GenerateFirmwareUpgradeAudit) HasBaselineId() bool`

HasBaselineId returns a boolean if a field has been set.

### GetOsTemplateLabel

`func (o *GenerateFirmwareUpgradeAudit) GetOsTemplateLabel() string`

GetOsTemplateLabel returns the OsTemplateLabel field if non-nil, zero value otherwise.

### GetOsTemplateLabelOk

`func (o *GenerateFirmwareUpgradeAudit) GetOsTemplateLabelOk() (*string, bool)`

GetOsTemplateLabelOk returns a tuple with the OsTemplateLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateLabel

`func (o *GenerateFirmwareUpgradeAudit) SetOsTemplateLabel(v string)`

SetOsTemplateLabel sets OsTemplateLabel field to given value.

### HasOsTemplateLabel

`func (o *GenerateFirmwareUpgradeAudit) HasOsTemplateLabel() bool`

HasOsTemplateLabel returns a boolean if a field has been set.

### GetServerTypeName

`func (o *GenerateFirmwareUpgradeAudit) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *GenerateFirmwareUpgradeAudit) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *GenerateFirmwareUpgradeAudit) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.

### HasServerTypeName

`func (o *GenerateFirmwareUpgradeAudit) HasServerTypeName() bool`

HasServerTypeName returns a boolean if a field has been set.

### GetSiteId

`func (o *GenerateFirmwareUpgradeAudit) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GenerateFirmwareUpgradeAudit) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GenerateFirmwareUpgradeAudit) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GenerateFirmwareUpgradeAudit) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


