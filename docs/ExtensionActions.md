# ExtensionActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreDeploy** | Pointer to [**[]ExtensionTask**](ExtensionTask.md) | Pre-deploy tasks. | [optional] 
**PostDeploy** | Pointer to [**[]ExtensionTask**](ExtensionTask.md) | Post-deploy tasks. | [optional] 

## Methods

### NewExtensionActions

`func NewExtensionActions() *ExtensionActions`

NewExtensionActions instantiates a new ExtensionActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionActionsWithDefaults

`func NewExtensionActionsWithDefaults() *ExtensionActions`

NewExtensionActionsWithDefaults instantiates a new ExtensionActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreDeploy

`func (o *ExtensionActions) GetPreDeploy() []ExtensionTask`

GetPreDeploy returns the PreDeploy field if non-nil, zero value otherwise.

### GetPreDeployOk

`func (o *ExtensionActions) GetPreDeployOk() (*[]ExtensionTask, bool)`

GetPreDeployOk returns a tuple with the PreDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDeploy

`func (o *ExtensionActions) SetPreDeploy(v []ExtensionTask)`

SetPreDeploy sets PreDeploy field to given value.

### HasPreDeploy

`func (o *ExtensionActions) HasPreDeploy() bool`

HasPreDeploy returns a boolean if a field has been set.

### GetPostDeploy

`func (o *ExtensionActions) GetPostDeploy() []ExtensionTask`

GetPostDeploy returns the PostDeploy field if non-nil, zero value otherwise.

### GetPostDeployOk

`func (o *ExtensionActions) GetPostDeployOk() (*[]ExtensionTask, bool)`

GetPostDeployOk returns a tuple with the PostDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDeploy

`func (o *ExtensionActions) SetPostDeploy(v []ExtensionTask)`

SetPostDeploy sets PostDeploy field to given value.

### HasPostDeploy

`func (o *ExtensionActions) HasPostDeploy() bool`

HasPostDeploy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


