# NetworkFabricDeployOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireConfirmation** | **bool** | Flag to indicate if the deployment should be confirmed before proceeding with the execution of template on network devices. There will be a confirmation step before the deployment is executed, giving you the chance to review the changes. | 

## Methods

### NewNetworkFabricDeployOptions

`func NewNetworkFabricDeployOptions(requireConfirmation bool, ) *NetworkFabricDeployOptions`

NewNetworkFabricDeployOptions instantiates a new NetworkFabricDeployOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricDeployOptionsWithDefaults

`func NewNetworkFabricDeployOptionsWithDefaults() *NetworkFabricDeployOptions`

NewNetworkFabricDeployOptionsWithDefaults instantiates a new NetworkFabricDeployOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireConfirmation

`func (o *NetworkFabricDeployOptions) GetRequireConfirmation() bool`

GetRequireConfirmation returns the RequireConfirmation field if non-nil, zero value otherwise.

### GetRequireConfirmationOk

`func (o *NetworkFabricDeployOptions) GetRequireConfirmationOk() (*bool, bool)`

GetRequireConfirmationOk returns a tuple with the RequireConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireConfirmation

`func (o *NetworkFabricDeployOptions) SetRequireConfirmation(v bool)`

SetRequireConfirmation sets RequireConfirmation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


