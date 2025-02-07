# ServerTypeUtilizationReportGrouped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registering** | **[]string** | Ids of the servers having status registering. | 
**Available** | **[]string** | Ids of the servers having status available. | 
**CleaningRequired** | **[]string** | Ids of the servers having status cleaning_required. | 
**Cleaning** | **[]string** | Ids of the servers having status cleaning. | 
**Used** | **[]string** | Ids of the servers having status used. | 
**UsedRegistering** | **[]string** | Ids of the servers having status used_registering. | 
**Defective** | **[]string** | Ids of the servers having status defective. | 
**RemovedFromRack** | **[]string** | Ids of the servers having status removed_from_rack. | 
**Decommissioned** | **[]string** | Ids of the servers having status decommissioned. | 
**UpdatingFirmware** | **[]string** | Ids of the servers having status updating_firmware. | 
**UsedDiagnostics** | **[]string** | Ids of the servers having status used_diagnostics. | 
**PendingRegistration** | **[]string** | Ids of the servers having status pending_registration. | 
**Unavailable** | **[]string** | Ids of the servers having status unavailable. | 

## Methods

### NewServerTypeUtilizationReportGrouped

`func NewServerTypeUtilizationReportGrouped(registering []string, available []string, cleaningRequired []string, cleaning []string, used []string, usedRegistering []string, defective []string, removedFromRack []string, decommissioned []string, updatingFirmware []string, usedDiagnostics []string, pendingRegistration []string, unavailable []string, ) *ServerTypeUtilizationReportGrouped`

NewServerTypeUtilizationReportGrouped instantiates a new ServerTypeUtilizationReportGrouped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeUtilizationReportGroupedWithDefaults

`func NewServerTypeUtilizationReportGroupedWithDefaults() *ServerTypeUtilizationReportGrouped`

NewServerTypeUtilizationReportGroupedWithDefaults instantiates a new ServerTypeUtilizationReportGrouped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistering

`func (o *ServerTypeUtilizationReportGrouped) GetRegistering() []string`

GetRegistering returns the Registering field if non-nil, zero value otherwise.

### GetRegisteringOk

`func (o *ServerTypeUtilizationReportGrouped) GetRegisteringOk() (*[]string, bool)`

GetRegisteringOk returns a tuple with the Registering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistering

`func (o *ServerTypeUtilizationReportGrouped) SetRegistering(v []string)`

SetRegistering sets Registering field to given value.


### GetAvailable

`func (o *ServerTypeUtilizationReportGrouped) GetAvailable() []string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ServerTypeUtilizationReportGrouped) GetAvailableOk() (*[]string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ServerTypeUtilizationReportGrouped) SetAvailable(v []string)`

SetAvailable sets Available field to given value.


### GetCleaningRequired

`func (o *ServerTypeUtilizationReportGrouped) GetCleaningRequired() []string`

GetCleaningRequired returns the CleaningRequired field if non-nil, zero value otherwise.

### GetCleaningRequiredOk

`func (o *ServerTypeUtilizationReportGrouped) GetCleaningRequiredOk() (*[]string, bool)`

GetCleaningRequiredOk returns a tuple with the CleaningRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleaningRequired

`func (o *ServerTypeUtilizationReportGrouped) SetCleaningRequired(v []string)`

SetCleaningRequired sets CleaningRequired field to given value.


### GetCleaning

`func (o *ServerTypeUtilizationReportGrouped) GetCleaning() []string`

GetCleaning returns the Cleaning field if non-nil, zero value otherwise.

### GetCleaningOk

`func (o *ServerTypeUtilizationReportGrouped) GetCleaningOk() (*[]string, bool)`

GetCleaningOk returns a tuple with the Cleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleaning

`func (o *ServerTypeUtilizationReportGrouped) SetCleaning(v []string)`

SetCleaning sets Cleaning field to given value.


### GetUsed

`func (o *ServerTypeUtilizationReportGrouped) GetUsed() []string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *ServerTypeUtilizationReportGrouped) GetUsedOk() (*[]string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *ServerTypeUtilizationReportGrouped) SetUsed(v []string)`

SetUsed sets Used field to given value.


### GetUsedRegistering

`func (o *ServerTypeUtilizationReportGrouped) GetUsedRegistering() []string`

GetUsedRegistering returns the UsedRegistering field if non-nil, zero value otherwise.

### GetUsedRegisteringOk

`func (o *ServerTypeUtilizationReportGrouped) GetUsedRegisteringOk() (*[]string, bool)`

GetUsedRegisteringOk returns a tuple with the UsedRegistering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedRegistering

`func (o *ServerTypeUtilizationReportGrouped) SetUsedRegistering(v []string)`

SetUsedRegistering sets UsedRegistering field to given value.


### GetDefective

`func (o *ServerTypeUtilizationReportGrouped) GetDefective() []string`

GetDefective returns the Defective field if non-nil, zero value otherwise.

### GetDefectiveOk

`func (o *ServerTypeUtilizationReportGrouped) GetDefectiveOk() (*[]string, bool)`

GetDefectiveOk returns a tuple with the Defective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefective

`func (o *ServerTypeUtilizationReportGrouped) SetDefective(v []string)`

SetDefective sets Defective field to given value.


### GetRemovedFromRack

`func (o *ServerTypeUtilizationReportGrouped) GetRemovedFromRack() []string`

GetRemovedFromRack returns the RemovedFromRack field if non-nil, zero value otherwise.

### GetRemovedFromRackOk

`func (o *ServerTypeUtilizationReportGrouped) GetRemovedFromRackOk() (*[]string, bool)`

GetRemovedFromRackOk returns a tuple with the RemovedFromRack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedFromRack

`func (o *ServerTypeUtilizationReportGrouped) SetRemovedFromRack(v []string)`

SetRemovedFromRack sets RemovedFromRack field to given value.


### GetDecommissioned

`func (o *ServerTypeUtilizationReportGrouped) GetDecommissioned() []string`

GetDecommissioned returns the Decommissioned field if non-nil, zero value otherwise.

### GetDecommissionedOk

`func (o *ServerTypeUtilizationReportGrouped) GetDecommissionedOk() (*[]string, bool)`

GetDecommissionedOk returns a tuple with the Decommissioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissioned

`func (o *ServerTypeUtilizationReportGrouped) SetDecommissioned(v []string)`

SetDecommissioned sets Decommissioned field to given value.


### GetUpdatingFirmware

`func (o *ServerTypeUtilizationReportGrouped) GetUpdatingFirmware() []string`

GetUpdatingFirmware returns the UpdatingFirmware field if non-nil, zero value otherwise.

### GetUpdatingFirmwareOk

`func (o *ServerTypeUtilizationReportGrouped) GetUpdatingFirmwareOk() (*[]string, bool)`

GetUpdatingFirmwareOk returns a tuple with the UpdatingFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingFirmware

`func (o *ServerTypeUtilizationReportGrouped) SetUpdatingFirmware(v []string)`

SetUpdatingFirmware sets UpdatingFirmware field to given value.


### GetUsedDiagnostics

`func (o *ServerTypeUtilizationReportGrouped) GetUsedDiagnostics() []string`

GetUsedDiagnostics returns the UsedDiagnostics field if non-nil, zero value otherwise.

### GetUsedDiagnosticsOk

`func (o *ServerTypeUtilizationReportGrouped) GetUsedDiagnosticsOk() (*[]string, bool)`

GetUsedDiagnosticsOk returns a tuple with the UsedDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDiagnostics

`func (o *ServerTypeUtilizationReportGrouped) SetUsedDiagnostics(v []string)`

SetUsedDiagnostics sets UsedDiagnostics field to given value.


### GetPendingRegistration

`func (o *ServerTypeUtilizationReportGrouped) GetPendingRegistration() []string`

GetPendingRegistration returns the PendingRegistration field if non-nil, zero value otherwise.

### GetPendingRegistrationOk

`func (o *ServerTypeUtilizationReportGrouped) GetPendingRegistrationOk() (*[]string, bool)`

GetPendingRegistrationOk returns a tuple with the PendingRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRegistration

`func (o *ServerTypeUtilizationReportGrouped) SetPendingRegistration(v []string)`

SetPendingRegistration sets PendingRegistration field to given value.


### GetUnavailable

`func (o *ServerTypeUtilizationReportGrouped) GetUnavailable() []string`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *ServerTypeUtilizationReportGrouped) GetUnavailableOk() (*[]string, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *ServerTypeUtilizationReportGrouped) SetUnavailable(v []string)`

SetUnavailable sets Unavailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


