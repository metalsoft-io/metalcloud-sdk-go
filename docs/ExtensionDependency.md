# ExtensionDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerVersion** | **string** | Metalsoft Controller version required by the extension. | 
**OsTemplates** | **[]string** | List of OS templates required by the extension. | 

## Methods

### NewExtensionDependency

`func NewExtensionDependency(controllerVersion string, osTemplates []string, ) *ExtensionDependency`

NewExtensionDependency instantiates a new ExtensionDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionDependencyWithDefaults

`func NewExtensionDependencyWithDefaults() *ExtensionDependency`

NewExtensionDependencyWithDefaults instantiates a new ExtensionDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerVersion

`func (o *ExtensionDependency) GetControllerVersion() string`

GetControllerVersion returns the ControllerVersion field if non-nil, zero value otherwise.

### GetControllerVersionOk

`func (o *ExtensionDependency) GetControllerVersionOk() (*string, bool)`

GetControllerVersionOk returns a tuple with the ControllerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVersion

`func (o *ExtensionDependency) SetControllerVersion(v string)`

SetControllerVersion sets ControllerVersion field to given value.


### GetOsTemplates

`func (o *ExtensionDependency) GetOsTemplates() []string`

GetOsTemplates returns the OsTemplates field if non-nil, zero value otherwise.

### GetOsTemplatesOk

`func (o *ExtensionDependency) GetOsTemplatesOk() (*[]string, bool)`

GetOsTemplatesOk returns a tuple with the OsTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplates

`func (o *ExtensionDependency) SetOsTemplates(v []string)`

SetOsTemplates sets OsTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


