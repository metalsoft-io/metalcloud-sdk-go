# ServerInstanceReinstallOS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformAtNextDeploy** | **bool** | Whether to perform the reinstall at next deploy. (always true at the moment) | 
**ReinstallOS** | **bool** | Whether to reinstall the OS. | 

## Methods

### NewServerInstanceReinstallOS

`func NewServerInstanceReinstallOS(performAtNextDeploy bool, reinstallOS bool, ) *ServerInstanceReinstallOS`

NewServerInstanceReinstallOS instantiates a new ServerInstanceReinstallOS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceReinstallOSWithDefaults

`func NewServerInstanceReinstallOSWithDefaults() *ServerInstanceReinstallOS`

NewServerInstanceReinstallOSWithDefaults instantiates a new ServerInstanceReinstallOS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformAtNextDeploy

`func (o *ServerInstanceReinstallOS) GetPerformAtNextDeploy() bool`

GetPerformAtNextDeploy returns the PerformAtNextDeploy field if non-nil, zero value otherwise.

### GetPerformAtNextDeployOk

`func (o *ServerInstanceReinstallOS) GetPerformAtNextDeployOk() (*bool, bool)`

GetPerformAtNextDeployOk returns a tuple with the PerformAtNextDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformAtNextDeploy

`func (o *ServerInstanceReinstallOS) SetPerformAtNextDeploy(v bool)`

SetPerformAtNextDeploy sets PerformAtNextDeploy field to given value.


### GetReinstallOS

`func (o *ServerInstanceReinstallOS) GetReinstallOS() bool`

GetReinstallOS returns the ReinstallOS field if non-nil, zero value otherwise.

### GetReinstallOSOk

`func (o *ServerInstanceReinstallOS) GetReinstallOSOk() (*bool, bool)`

GetReinstallOSOk returns a tuple with the ReinstallOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinstallOS

`func (o *ServerInstanceReinstallOS) SetReinstallOS(v bool)`

SetReinstallOS sets ReinstallOS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


