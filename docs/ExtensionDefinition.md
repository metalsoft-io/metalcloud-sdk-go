# ExtensionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of extension. | 
**SchemaVersion** | **string** | Schema version of the extension. | 
**Name** | **string** | Name of the extension. | 
**Label** | **string** | Label of the extension. | 
**ExtensionType** | **string** | Type of the extension. | 
**Vendor** | **string** | Vendor of the extension. | 
**ExtensionVersion** | **string** | Version of the extension. | 
**Description** | Pointer to **string** | Description of the extension. | [optional] 
**Icon** | **string** | Icon of the extension. | 
**Dependencies** | [**ExtensionDependency**](ExtensionDependency.md) |  | 
**Inputs** | [**[]ExtensionInput**](ExtensionInput.md) | List of inputs for the platform service. | 
**Outputs** | [**[]ExtensionOutput**](ExtensionOutput.md) | List of outputs for the platform service. | 
**Infrastructure** | Pointer to [**ExtensionInfrastructure**](ExtensionInfrastructure.md) |  | [optional] 
**Assets** | [**[]ExtensionAsset**](ExtensionAsset.md) | List of assets for the platform service. | 
**OnCreate** | Pointer to [**ExtensionActions**](ExtensionActions.md) |  | [optional] 
**OnEdit** | Pointer to [**ExtensionActions**](ExtensionActions.md) |  | [optional] 
**OnDelete** | Pointer to [**ExtensionActions**](ExtensionActions.md) |  | [optional] 
**Tasks** | Pointer to [**[]ExtensionTask**](ExtensionTask.md) | Tasks. Only for extensions of kind action | [optional] 

## Methods

### NewExtensionDefinition

`func NewExtensionDefinition(kind string, schemaVersion string, name string, label string, extensionType string, vendor string, extensionVersion string, icon string, dependencies ExtensionDependency, inputs []ExtensionInput, outputs []ExtensionOutput, assets []ExtensionAsset, ) *ExtensionDefinition`

NewExtensionDefinition instantiates a new ExtensionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionDefinitionWithDefaults

`func NewExtensionDefinitionWithDefaults() *ExtensionDefinition`

NewExtensionDefinitionWithDefaults instantiates a new ExtensionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ExtensionDefinition) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExtensionDefinition) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExtensionDefinition) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSchemaVersion

`func (o *ExtensionDefinition) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ExtensionDefinition) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ExtensionDefinition) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetName

`func (o *ExtensionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ExtensionDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetExtensionType

`func (o *ExtensionDefinition) GetExtensionType() string`

GetExtensionType returns the ExtensionType field if non-nil, zero value otherwise.

### GetExtensionTypeOk

`func (o *ExtensionDefinition) GetExtensionTypeOk() (*string, bool)`

GetExtensionTypeOk returns a tuple with the ExtensionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionType

`func (o *ExtensionDefinition) SetExtensionType(v string)`

SetExtensionType sets ExtensionType field to given value.


### GetVendor

`func (o *ExtensionDefinition) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ExtensionDefinition) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ExtensionDefinition) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetExtensionVersion

`func (o *ExtensionDefinition) GetExtensionVersion() string`

GetExtensionVersion returns the ExtensionVersion field if non-nil, zero value otherwise.

### GetExtensionVersionOk

`func (o *ExtensionDefinition) GetExtensionVersionOk() (*string, bool)`

GetExtensionVersionOk returns a tuple with the ExtensionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionVersion

`func (o *ExtensionDefinition) SetExtensionVersion(v string)`

SetExtensionVersion sets ExtensionVersion field to given value.


### GetDescription

`func (o *ExtensionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtensionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtensionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtensionDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *ExtensionDefinition) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ExtensionDefinition) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ExtensionDefinition) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetDependencies

`func (o *ExtensionDefinition) GetDependencies() ExtensionDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *ExtensionDefinition) GetDependenciesOk() (*ExtensionDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *ExtensionDefinition) SetDependencies(v ExtensionDependency)`

SetDependencies sets Dependencies field to given value.


### GetInputs

`func (o *ExtensionDefinition) GetInputs() []ExtensionInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ExtensionDefinition) GetInputsOk() (*[]ExtensionInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ExtensionDefinition) SetInputs(v []ExtensionInput)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *ExtensionDefinition) GetOutputs() []ExtensionOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ExtensionDefinition) GetOutputsOk() (*[]ExtensionOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ExtensionDefinition) SetOutputs(v []ExtensionOutput)`

SetOutputs sets Outputs field to given value.


### GetInfrastructure

`func (o *ExtensionDefinition) GetInfrastructure() ExtensionInfrastructure`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ExtensionDefinition) GetInfrastructureOk() (*ExtensionInfrastructure, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ExtensionDefinition) SetInfrastructure(v ExtensionInfrastructure)`

SetInfrastructure sets Infrastructure field to given value.

### HasInfrastructure

`func (o *ExtensionDefinition) HasInfrastructure() bool`

HasInfrastructure returns a boolean if a field has been set.

### GetAssets

`func (o *ExtensionDefinition) GetAssets() []ExtensionAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ExtensionDefinition) GetAssetsOk() (*[]ExtensionAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ExtensionDefinition) SetAssets(v []ExtensionAsset)`

SetAssets sets Assets field to given value.


### GetOnCreate

`func (o *ExtensionDefinition) GetOnCreate() ExtensionActions`

GetOnCreate returns the OnCreate field if non-nil, zero value otherwise.

### GetOnCreateOk

`func (o *ExtensionDefinition) GetOnCreateOk() (*ExtensionActions, bool)`

GetOnCreateOk returns a tuple with the OnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCreate

`func (o *ExtensionDefinition) SetOnCreate(v ExtensionActions)`

SetOnCreate sets OnCreate field to given value.

### HasOnCreate

`func (o *ExtensionDefinition) HasOnCreate() bool`

HasOnCreate returns a boolean if a field has been set.

### GetOnEdit

`func (o *ExtensionDefinition) GetOnEdit() ExtensionActions`

GetOnEdit returns the OnEdit field if non-nil, zero value otherwise.

### GetOnEditOk

`func (o *ExtensionDefinition) GetOnEditOk() (*ExtensionActions, bool)`

GetOnEditOk returns a tuple with the OnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnEdit

`func (o *ExtensionDefinition) SetOnEdit(v ExtensionActions)`

SetOnEdit sets OnEdit field to given value.

### HasOnEdit

`func (o *ExtensionDefinition) HasOnEdit() bool`

HasOnEdit returns a boolean if a field has been set.

### GetOnDelete

`func (o *ExtensionDefinition) GetOnDelete() ExtensionActions`

GetOnDelete returns the OnDelete field if non-nil, zero value otherwise.

### GetOnDeleteOk

`func (o *ExtensionDefinition) GetOnDeleteOk() (*ExtensionActions, bool)`

GetOnDeleteOk returns a tuple with the OnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelete

`func (o *ExtensionDefinition) SetOnDelete(v ExtensionActions)`

SetOnDelete sets OnDelete field to given value.

### HasOnDelete

`func (o *ExtensionDefinition) HasOnDelete() bool`

HasOnDelete returns a boolean if a field has been set.

### GetTasks

`func (o *ExtensionDefinition) GetTasks() []ExtensionTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ExtensionDefinition) GetTasksOk() (*[]ExtensionTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ExtensionDefinition) SetTasks(v []ExtensionTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ExtensionDefinition) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


