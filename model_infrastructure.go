/*
MetalSoft REST API

MetalSoft REST API documentation

API version: 2.0
Contact: support@metalsoft.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// checks if the Infrastructure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Infrastructure{}

// Infrastructure struct for Infrastructure
type Infrastructure struct {
	Label string `json:"label"`
	// Custom variables in JSON format.
	CustomVariables map[string]interface{} `json:"customVariables,omitempty"`
	// User ID of the owner of the Infrastructure.
	UserIdOwner *float32 `json:"userIdOwner,omitempty"`
	// Subdomain associated with the Infrastructure.
	Subdomain *string `json:"subdomain,omitempty"`
	// Number of active instances.
	InstancesCountActive *float32 `json:"instancesCountActive,omitempty"`
	// Number of active drives.
	DrivesCountActive *float32 `json:"drivesCountActive,omitempty"`
	// Number of active IPv4 subnets.
	Ipv4SubnetsCountActive *float32 `json:"ipv4SubnetsCountActive,omitempty"`
	// Number of active IPv6 subnets.
	Ipv6SubnetsCountActive *float32 `json:"ipv6SubnetsCountActive,omitempty"`
	// Number of unused IPv4 addresses.
	Ipv4UnusedIpAddresses *float32 `json:"ipv4UnusedIpAddresses,omitempty"`
	// Description of the infrastructure.
	Description *string `json:"description,omitempty"`
	// Settings in JSON format.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// Whether the infrastructure API is private.
	IsApiPrivate *float32 `json:"isApiPrivate,omitempty"`
	// Experimental priority.
	ExperimentalPriority *string `json:"experimentalPriority,omitempty"`
	// Whether the infrastructure is a member of public designs.
	IsPublicDesignsMember *float32 `json:"isPublicDesignsMember,omitempty"`
	// Certificates in JSON format.
	CertificatesJson *string `json:"certificatesJson,omitempty"`
	// Deploy cookie jar JSON.
	DeployCookieJarJson map[string]interface{} `json:"deployCookieJarJson,omitempty"`
	// Last error of deferred deploy attempt.
	DeferredDeployAttemptLastErrorJson map[string]interface{} `json:"deferredDeployAttemptLastErrorJson,omitempty"`
	// Whether the infrastructure is automanaged.
	IsAutomanaged *float32 `json:"isAutomanaged,omitempty"`
	// Timestamp of the latest update for the Infrastructure.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// Infrastructure Id
	Id float32 `json:"id"`
	// Revision of the Infrastructure
	Revision float32 `json:"revision"`
	// Service status of the Infrastructure
	ServiceStatus GenericServiceStatus `json:"serviceStatus"`
	// Datacenter name where the Infrastructure is located.
	DatacenterName string `json:"datacenterName"`
	// The ID of the site where the Infrastructure is located.
	SiteId float32 `json:"siteId"`
	// Timestamp of the Infrastructure creation.
	CreatedTimestamp string `json:"createdTimestamp"`
	// Permanent subdomain associated with the Infrastructure.
	SubdomainPermanent *string `json:"subdomainPermanent,omitempty"`
	// DNS Subdomain ID.
	DnsSubdomainId *float32 `json:"dnsSubdomainId,omitempty"`
	// Permanent DNS Subdomain ID.
	DnsSubdomainPermanentId *float32 `json:"dnsSubdomainPermanentId,omitempty"`
	// Infrastructure design locked flag.
	DesignIsLocked float32 `json:"designIsLocked"`
	// The current changes to be deployed for the Infrastructure.
	Config InfrastructureConfig `json:"config"`
	// Meta information for the Infrastructure
	Meta *InfrastructureMeta `json:"meta,omitempty"`
	Statistics *JobGroupStatisticsWithoutId `json:"statistics,omitempty"`
	// Links to other resources
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Infrastructure Infrastructure

// NewInfrastructure instantiates a new Infrastructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfrastructure(label string, updatedTimestamp string, id float32, revision float32, serviceStatus GenericServiceStatus, datacenterName string, siteId float32, createdTimestamp string, designIsLocked float32, config InfrastructureConfig) *Infrastructure {
	this := Infrastructure{}
	this.Label = label
	this.UpdatedTimestamp = updatedTimestamp
	this.Id = id
	this.Revision = revision
	this.ServiceStatus = serviceStatus
	this.DatacenterName = datacenterName
	this.SiteId = siteId
	this.CreatedTimestamp = createdTimestamp
	this.DesignIsLocked = designIsLocked
	this.Config = config
	return &this
}

// NewInfrastructureWithDefaults instantiates a new Infrastructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfrastructureWithDefaults() *Infrastructure {
	this := Infrastructure{}
	return &this
}

// GetLabel returns the Label field value
func (o *Infrastructure) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Infrastructure) SetLabel(v string) {
	o.Label = v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *Infrastructure) GetCustomVariables() map[string]interface{} {
	if o == nil || IsNil(o.CustomVariables) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetCustomVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return map[string]interface{}{}, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *Infrastructure) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given map[string]interface{} and assigns it to the CustomVariables field.
func (o *Infrastructure) SetCustomVariables(v map[string]interface{}) {
	o.CustomVariables = v
}

// GetUserIdOwner returns the UserIdOwner field value if set, zero value otherwise.
func (o *Infrastructure) GetUserIdOwner() float32 {
	if o == nil || IsNil(o.UserIdOwner) {
		var ret float32
		return ret
	}
	return *o.UserIdOwner
}

// GetUserIdOwnerOk returns a tuple with the UserIdOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetUserIdOwnerOk() (*float32, bool) {
	if o == nil || IsNil(o.UserIdOwner) {
		return nil, false
	}
	return o.UserIdOwner, true
}

// HasUserIdOwner returns a boolean if a field has been set.
func (o *Infrastructure) HasUserIdOwner() bool {
	if o != nil && !IsNil(o.UserIdOwner) {
		return true
	}

	return false
}

// SetUserIdOwner gets a reference to the given float32 and assigns it to the UserIdOwner field.
func (o *Infrastructure) SetUserIdOwner(v float32) {
	o.UserIdOwner = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *Infrastructure) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *Infrastructure) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *Infrastructure) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetInstancesCountActive returns the InstancesCountActive field value if set, zero value otherwise.
func (o *Infrastructure) GetInstancesCountActive() float32 {
	if o == nil || IsNil(o.InstancesCountActive) {
		var ret float32
		return ret
	}
	return *o.InstancesCountActive
}

// GetInstancesCountActiveOk returns a tuple with the InstancesCountActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetInstancesCountActiveOk() (*float32, bool) {
	if o == nil || IsNil(o.InstancesCountActive) {
		return nil, false
	}
	return o.InstancesCountActive, true
}

// HasInstancesCountActive returns a boolean if a field has been set.
func (o *Infrastructure) HasInstancesCountActive() bool {
	if o != nil && !IsNil(o.InstancesCountActive) {
		return true
	}

	return false
}

// SetInstancesCountActive gets a reference to the given float32 and assigns it to the InstancesCountActive field.
func (o *Infrastructure) SetInstancesCountActive(v float32) {
	o.InstancesCountActive = &v
}

// GetDrivesCountActive returns the DrivesCountActive field value if set, zero value otherwise.
func (o *Infrastructure) GetDrivesCountActive() float32 {
	if o == nil || IsNil(o.DrivesCountActive) {
		var ret float32
		return ret
	}
	return *o.DrivesCountActive
}

// GetDrivesCountActiveOk returns a tuple with the DrivesCountActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDrivesCountActiveOk() (*float32, bool) {
	if o == nil || IsNil(o.DrivesCountActive) {
		return nil, false
	}
	return o.DrivesCountActive, true
}

// HasDrivesCountActive returns a boolean if a field has been set.
func (o *Infrastructure) HasDrivesCountActive() bool {
	if o != nil && !IsNil(o.DrivesCountActive) {
		return true
	}

	return false
}

// SetDrivesCountActive gets a reference to the given float32 and assigns it to the DrivesCountActive field.
func (o *Infrastructure) SetDrivesCountActive(v float32) {
	o.DrivesCountActive = &v
}

// GetIpv4SubnetsCountActive returns the Ipv4SubnetsCountActive field value if set, zero value otherwise.
func (o *Infrastructure) GetIpv4SubnetsCountActive() float32 {
	if o == nil || IsNil(o.Ipv4SubnetsCountActive) {
		var ret float32
		return ret
	}
	return *o.Ipv4SubnetsCountActive
}

// GetIpv4SubnetsCountActiveOk returns a tuple with the Ipv4SubnetsCountActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetIpv4SubnetsCountActiveOk() (*float32, bool) {
	if o == nil || IsNil(o.Ipv4SubnetsCountActive) {
		return nil, false
	}
	return o.Ipv4SubnetsCountActive, true
}

// HasIpv4SubnetsCountActive returns a boolean if a field has been set.
func (o *Infrastructure) HasIpv4SubnetsCountActive() bool {
	if o != nil && !IsNil(o.Ipv4SubnetsCountActive) {
		return true
	}

	return false
}

// SetIpv4SubnetsCountActive gets a reference to the given float32 and assigns it to the Ipv4SubnetsCountActive field.
func (o *Infrastructure) SetIpv4SubnetsCountActive(v float32) {
	o.Ipv4SubnetsCountActive = &v
}

// GetIpv6SubnetsCountActive returns the Ipv6SubnetsCountActive field value if set, zero value otherwise.
func (o *Infrastructure) GetIpv6SubnetsCountActive() float32 {
	if o == nil || IsNil(o.Ipv6SubnetsCountActive) {
		var ret float32
		return ret
	}
	return *o.Ipv6SubnetsCountActive
}

// GetIpv6SubnetsCountActiveOk returns a tuple with the Ipv6SubnetsCountActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetIpv6SubnetsCountActiveOk() (*float32, bool) {
	if o == nil || IsNil(o.Ipv6SubnetsCountActive) {
		return nil, false
	}
	return o.Ipv6SubnetsCountActive, true
}

// HasIpv6SubnetsCountActive returns a boolean if a field has been set.
func (o *Infrastructure) HasIpv6SubnetsCountActive() bool {
	if o != nil && !IsNil(o.Ipv6SubnetsCountActive) {
		return true
	}

	return false
}

// SetIpv6SubnetsCountActive gets a reference to the given float32 and assigns it to the Ipv6SubnetsCountActive field.
func (o *Infrastructure) SetIpv6SubnetsCountActive(v float32) {
	o.Ipv6SubnetsCountActive = &v
}

// GetIpv4UnusedIpAddresses returns the Ipv4UnusedIpAddresses field value if set, zero value otherwise.
func (o *Infrastructure) GetIpv4UnusedIpAddresses() float32 {
	if o == nil || IsNil(o.Ipv4UnusedIpAddresses) {
		var ret float32
		return ret
	}
	return *o.Ipv4UnusedIpAddresses
}

// GetIpv4UnusedIpAddressesOk returns a tuple with the Ipv4UnusedIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetIpv4UnusedIpAddressesOk() (*float32, bool) {
	if o == nil || IsNil(o.Ipv4UnusedIpAddresses) {
		return nil, false
	}
	return o.Ipv4UnusedIpAddresses, true
}

// HasIpv4UnusedIpAddresses returns a boolean if a field has been set.
func (o *Infrastructure) HasIpv4UnusedIpAddresses() bool {
	if o != nil && !IsNil(o.Ipv4UnusedIpAddresses) {
		return true
	}

	return false
}

// SetIpv4UnusedIpAddresses gets a reference to the given float32 and assigns it to the Ipv4UnusedIpAddresses field.
func (o *Infrastructure) SetIpv4UnusedIpAddresses(v float32) {
	o.Ipv4UnusedIpAddresses = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Infrastructure) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Infrastructure) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Infrastructure) SetDescription(v string) {
	o.Description = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Infrastructure) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Infrastructure) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *Infrastructure) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetIsApiPrivate returns the IsApiPrivate field value if set, zero value otherwise.
func (o *Infrastructure) GetIsApiPrivate() float32 {
	if o == nil || IsNil(o.IsApiPrivate) {
		var ret float32
		return ret
	}
	return *o.IsApiPrivate
}

// GetIsApiPrivateOk returns a tuple with the IsApiPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetIsApiPrivateOk() (*float32, bool) {
	if o == nil || IsNil(o.IsApiPrivate) {
		return nil, false
	}
	return o.IsApiPrivate, true
}

// HasIsApiPrivate returns a boolean if a field has been set.
func (o *Infrastructure) HasIsApiPrivate() bool {
	if o != nil && !IsNil(o.IsApiPrivate) {
		return true
	}

	return false
}

// SetIsApiPrivate gets a reference to the given float32 and assigns it to the IsApiPrivate field.
func (o *Infrastructure) SetIsApiPrivate(v float32) {
	o.IsApiPrivate = &v
}

// GetExperimentalPriority returns the ExperimentalPriority field value if set, zero value otherwise.
func (o *Infrastructure) GetExperimentalPriority() string {
	if o == nil || IsNil(o.ExperimentalPriority) {
		var ret string
		return ret
	}
	return *o.ExperimentalPriority
}

// GetExperimentalPriorityOk returns a tuple with the ExperimentalPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetExperimentalPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.ExperimentalPriority) {
		return nil, false
	}
	return o.ExperimentalPriority, true
}

// HasExperimentalPriority returns a boolean if a field has been set.
func (o *Infrastructure) HasExperimentalPriority() bool {
	if o != nil && !IsNil(o.ExperimentalPriority) {
		return true
	}

	return false
}

// SetExperimentalPriority gets a reference to the given string and assigns it to the ExperimentalPriority field.
func (o *Infrastructure) SetExperimentalPriority(v string) {
	o.ExperimentalPriority = &v
}

// GetIsPublicDesignsMember returns the IsPublicDesignsMember field value if set, zero value otherwise.
func (o *Infrastructure) GetIsPublicDesignsMember() float32 {
	if o == nil || IsNil(o.IsPublicDesignsMember) {
		var ret float32
		return ret
	}
	return *o.IsPublicDesignsMember
}

// GetIsPublicDesignsMemberOk returns a tuple with the IsPublicDesignsMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetIsPublicDesignsMemberOk() (*float32, bool) {
	if o == nil || IsNil(o.IsPublicDesignsMember) {
		return nil, false
	}
	return o.IsPublicDesignsMember, true
}

// HasIsPublicDesignsMember returns a boolean if a field has been set.
func (o *Infrastructure) HasIsPublicDesignsMember() bool {
	if o != nil && !IsNil(o.IsPublicDesignsMember) {
		return true
	}

	return false
}

// SetIsPublicDesignsMember gets a reference to the given float32 and assigns it to the IsPublicDesignsMember field.
func (o *Infrastructure) SetIsPublicDesignsMember(v float32) {
	o.IsPublicDesignsMember = &v
}

// GetCertificatesJson returns the CertificatesJson field value if set, zero value otherwise.
func (o *Infrastructure) GetCertificatesJson() string {
	if o == nil || IsNil(o.CertificatesJson) {
		var ret string
		return ret
	}
	return *o.CertificatesJson
}

// GetCertificatesJsonOk returns a tuple with the CertificatesJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetCertificatesJsonOk() (*string, bool) {
	if o == nil || IsNil(o.CertificatesJson) {
		return nil, false
	}
	return o.CertificatesJson, true
}

// HasCertificatesJson returns a boolean if a field has been set.
func (o *Infrastructure) HasCertificatesJson() bool {
	if o != nil && !IsNil(o.CertificatesJson) {
		return true
	}

	return false
}

// SetCertificatesJson gets a reference to the given string and assigns it to the CertificatesJson field.
func (o *Infrastructure) SetCertificatesJson(v string) {
	o.CertificatesJson = &v
}

// GetDeployCookieJarJson returns the DeployCookieJarJson field value if set, zero value otherwise.
func (o *Infrastructure) GetDeployCookieJarJson() map[string]interface{} {
	if o == nil || IsNil(o.DeployCookieJarJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.DeployCookieJarJson
}

// GetDeployCookieJarJsonOk returns a tuple with the DeployCookieJarJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDeployCookieJarJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DeployCookieJarJson) {
		return map[string]interface{}{}, false
	}
	return o.DeployCookieJarJson, true
}

// HasDeployCookieJarJson returns a boolean if a field has been set.
func (o *Infrastructure) HasDeployCookieJarJson() bool {
	if o != nil && !IsNil(o.DeployCookieJarJson) {
		return true
	}

	return false
}

// SetDeployCookieJarJson gets a reference to the given map[string]interface{} and assigns it to the DeployCookieJarJson field.
func (o *Infrastructure) SetDeployCookieJarJson(v map[string]interface{}) {
	o.DeployCookieJarJson = v
}

// GetDeferredDeployAttemptLastErrorJson returns the DeferredDeployAttemptLastErrorJson field value if set, zero value otherwise.
func (o *Infrastructure) GetDeferredDeployAttemptLastErrorJson() map[string]interface{} {
	if o == nil || IsNil(o.DeferredDeployAttemptLastErrorJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.DeferredDeployAttemptLastErrorJson
}

// GetDeferredDeployAttemptLastErrorJsonOk returns a tuple with the DeferredDeployAttemptLastErrorJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDeferredDeployAttemptLastErrorJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DeferredDeployAttemptLastErrorJson) {
		return map[string]interface{}{}, false
	}
	return o.DeferredDeployAttemptLastErrorJson, true
}

// HasDeferredDeployAttemptLastErrorJson returns a boolean if a field has been set.
func (o *Infrastructure) HasDeferredDeployAttemptLastErrorJson() bool {
	if o != nil && !IsNil(o.DeferredDeployAttemptLastErrorJson) {
		return true
	}

	return false
}

// SetDeferredDeployAttemptLastErrorJson gets a reference to the given map[string]interface{} and assigns it to the DeferredDeployAttemptLastErrorJson field.
func (o *Infrastructure) SetDeferredDeployAttemptLastErrorJson(v map[string]interface{}) {
	o.DeferredDeployAttemptLastErrorJson = v
}

// GetIsAutomanaged returns the IsAutomanaged field value if set, zero value otherwise.
func (o *Infrastructure) GetIsAutomanaged() float32 {
	if o == nil || IsNil(o.IsAutomanaged) {
		var ret float32
		return ret
	}
	return *o.IsAutomanaged
}

// GetIsAutomanagedOk returns a tuple with the IsAutomanaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetIsAutomanagedOk() (*float32, bool) {
	if o == nil || IsNil(o.IsAutomanaged) {
		return nil, false
	}
	return o.IsAutomanaged, true
}

// HasIsAutomanaged returns a boolean if a field has been set.
func (o *Infrastructure) HasIsAutomanaged() bool {
	if o != nil && !IsNil(o.IsAutomanaged) {
		return true
	}

	return false
}

// SetIsAutomanaged gets a reference to the given float32 and assigns it to the IsAutomanaged field.
func (o *Infrastructure) SetIsAutomanaged(v float32) {
	o.IsAutomanaged = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *Infrastructure) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *Infrastructure) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetId returns the Id field value
func (o *Infrastructure) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Infrastructure) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *Infrastructure) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *Infrastructure) SetRevision(v float32) {
	o.Revision = v
}

// GetServiceStatus returns the ServiceStatus field value
func (o *Infrastructure) GetServiceStatus() GenericServiceStatus {
	if o == nil {
		var ret GenericServiceStatus
		return ret
	}

	return o.ServiceStatus
}

// GetServiceStatusOk returns a tuple with the ServiceStatus field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetServiceStatusOk() (*GenericServiceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceStatus, true
}

// SetServiceStatus sets field value
func (o *Infrastructure) SetServiceStatus(v GenericServiceStatus) {
	o.ServiceStatus = v
}

// GetDatacenterName returns the DatacenterName field value
func (o *Infrastructure) GetDatacenterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatacenterName, true
}

// SetDatacenterName sets field value
func (o *Infrastructure) SetDatacenterName(v string) {
	o.DatacenterName = v
}

// GetSiteId returns the SiteId field value
func (o *Infrastructure) GetSiteId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetSiteIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *Infrastructure) SetSiteId(v float32) {
	o.SiteId = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *Infrastructure) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *Infrastructure) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetSubdomainPermanent returns the SubdomainPermanent field value if set, zero value otherwise.
func (o *Infrastructure) GetSubdomainPermanent() string {
	if o == nil || IsNil(o.SubdomainPermanent) {
		var ret string
		return ret
	}
	return *o.SubdomainPermanent
}

// GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetSubdomainPermanentOk() (*string, bool) {
	if o == nil || IsNil(o.SubdomainPermanent) {
		return nil, false
	}
	return o.SubdomainPermanent, true
}

// HasSubdomainPermanent returns a boolean if a field has been set.
func (o *Infrastructure) HasSubdomainPermanent() bool {
	if o != nil && !IsNil(o.SubdomainPermanent) {
		return true
	}

	return false
}

// SetSubdomainPermanent gets a reference to the given string and assigns it to the SubdomainPermanent field.
func (o *Infrastructure) SetSubdomainPermanent(v string) {
	o.SubdomainPermanent = &v
}

// GetDnsSubdomainId returns the DnsSubdomainId field value if set, zero value otherwise.
func (o *Infrastructure) GetDnsSubdomainId() float32 {
	if o == nil || IsNil(o.DnsSubdomainId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainId
}

// GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDnsSubdomainIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainId) {
		return nil, false
	}
	return o.DnsSubdomainId, true
}

// HasDnsSubdomainId returns a boolean if a field has been set.
func (o *Infrastructure) HasDnsSubdomainId() bool {
	if o != nil && !IsNil(o.DnsSubdomainId) {
		return true
	}

	return false
}

// SetDnsSubdomainId gets a reference to the given float32 and assigns it to the DnsSubdomainId field.
func (o *Infrastructure) SetDnsSubdomainId(v float32) {
	o.DnsSubdomainId = &v
}

// GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field value if set, zero value otherwise.
func (o *Infrastructure) GetDnsSubdomainPermanentId() float32 {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainPermanentId
}

// GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDnsSubdomainPermanentIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		return nil, false
	}
	return o.DnsSubdomainPermanentId, true
}

// HasDnsSubdomainPermanentId returns a boolean if a field has been set.
func (o *Infrastructure) HasDnsSubdomainPermanentId() bool {
	if o != nil && !IsNil(o.DnsSubdomainPermanentId) {
		return true
	}

	return false
}

// SetDnsSubdomainPermanentId gets a reference to the given float32 and assigns it to the DnsSubdomainPermanentId field.
func (o *Infrastructure) SetDnsSubdomainPermanentId(v float32) {
	o.DnsSubdomainPermanentId = &v
}

// GetDesignIsLocked returns the DesignIsLocked field value
func (o *Infrastructure) GetDesignIsLocked() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DesignIsLocked
}

// GetDesignIsLockedOk returns a tuple with the DesignIsLocked field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetDesignIsLockedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesignIsLocked, true
}

// SetDesignIsLocked sets field value
func (o *Infrastructure) SetDesignIsLocked(v float32) {
	o.DesignIsLocked = v
}

// GetConfig returns the Config field value
func (o *Infrastructure) GetConfig() InfrastructureConfig {
	if o == nil {
		var ret InfrastructureConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetConfigOk() (*InfrastructureConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *Infrastructure) SetConfig(v InfrastructureConfig) {
	o.Config = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Infrastructure) GetMeta() InfrastructureMeta {
	if o == nil || IsNil(o.Meta) {
		var ret InfrastructureMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetMetaOk() (*InfrastructureMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Infrastructure) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InfrastructureMeta and assigns it to the Meta field.
func (o *Infrastructure) SetMeta(v InfrastructureMeta) {
	o.Meta = &v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *Infrastructure) GetStatistics() JobGroupStatisticsWithoutId {
	if o == nil || IsNil(o.Statistics) {
		var ret JobGroupStatisticsWithoutId
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetStatisticsOk() (*JobGroupStatisticsWithoutId, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *Infrastructure) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given JobGroupStatisticsWithoutId and assigns it to the Statistics field.
func (o *Infrastructure) SetStatistics(v JobGroupStatisticsWithoutId) {
	o.Statistics = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Infrastructure) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Infrastructure) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Infrastructure) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Infrastructure) SetLinks(v []Link) {
	o.Links = v
}

func (o Infrastructure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Infrastructure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	if !IsNil(o.UserIdOwner) {
		toSerialize["userIdOwner"] = o.UserIdOwner
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.InstancesCountActive) {
		toSerialize["instancesCountActive"] = o.InstancesCountActive
	}
	if !IsNil(o.DrivesCountActive) {
		toSerialize["drivesCountActive"] = o.DrivesCountActive
	}
	if !IsNil(o.Ipv4SubnetsCountActive) {
		toSerialize["ipv4SubnetsCountActive"] = o.Ipv4SubnetsCountActive
	}
	if !IsNil(o.Ipv6SubnetsCountActive) {
		toSerialize["ipv6SubnetsCountActive"] = o.Ipv6SubnetsCountActive
	}
	if !IsNil(o.Ipv4UnusedIpAddresses) {
		toSerialize["ipv4UnusedIpAddresses"] = o.Ipv4UnusedIpAddresses
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.IsApiPrivate) {
		toSerialize["isApiPrivate"] = o.IsApiPrivate
	}
	if !IsNil(o.ExperimentalPriority) {
		toSerialize["experimentalPriority"] = o.ExperimentalPriority
	}
	if !IsNil(o.IsPublicDesignsMember) {
		toSerialize["isPublicDesignsMember"] = o.IsPublicDesignsMember
	}
	if !IsNil(o.CertificatesJson) {
		toSerialize["certificatesJson"] = o.CertificatesJson
	}
	if !IsNil(o.DeployCookieJarJson) {
		toSerialize["deployCookieJarJson"] = o.DeployCookieJarJson
	}
	if !IsNil(o.DeferredDeployAttemptLastErrorJson) {
		toSerialize["deferredDeployAttemptLastErrorJson"] = o.DeferredDeployAttemptLastErrorJson
	}
	if !IsNil(o.IsAutomanaged) {
		toSerialize["isAutomanaged"] = o.IsAutomanaged
	}
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	toSerialize["serviceStatus"] = o.ServiceStatus
	toSerialize["datacenterName"] = o.DatacenterName
	toSerialize["siteId"] = o.SiteId
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	if !IsNil(o.SubdomainPermanent) {
		toSerialize["subdomainPermanent"] = o.SubdomainPermanent
	}
	if !IsNil(o.DnsSubdomainId) {
		toSerialize["dnsSubdomainId"] = o.DnsSubdomainId
	}
	if !IsNil(o.DnsSubdomainPermanentId) {
		toSerialize["dnsSubdomainPermanentId"] = o.DnsSubdomainPermanentId
	}
	toSerialize["designIsLocked"] = o.DesignIsLocked
	toSerialize["config"] = o.Config
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Infrastructure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"updatedTimestamp",
		"id",
		"revision",
		"serviceStatus",
		"datacenterName",
		"siteId",
		"createdTimestamp",
		"designIsLocked",
		"config",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInfrastructure := _Infrastructure{}

	err = json.Unmarshal(data, &varInfrastructure)

	if err != nil {
		return err
	}

	*o = Infrastructure(varInfrastructure)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "customVariables")
		delete(additionalProperties, "userIdOwner")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "instancesCountActive")
		delete(additionalProperties, "drivesCountActive")
		delete(additionalProperties, "ipv4SubnetsCountActive")
		delete(additionalProperties, "ipv6SubnetsCountActive")
		delete(additionalProperties, "ipv4UnusedIpAddresses")
		delete(additionalProperties, "description")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "isApiPrivate")
		delete(additionalProperties, "experimentalPriority")
		delete(additionalProperties, "isPublicDesignsMember")
		delete(additionalProperties, "certificatesJson")
		delete(additionalProperties, "deployCookieJarJson")
		delete(additionalProperties, "deferredDeployAttemptLastErrorJson")
		delete(additionalProperties, "isAutomanaged")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "serviceStatus")
		delete(additionalProperties, "datacenterName")
		delete(additionalProperties, "siteId")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "subdomainPermanent")
		delete(additionalProperties, "dnsSubdomainId")
		delete(additionalProperties, "dnsSubdomainPermanentId")
		delete(additionalProperties, "designIsLocked")
		delete(additionalProperties, "config")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "statistics")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfrastructure struct {
	value *Infrastructure
	isSet bool
}

func (v NullableInfrastructure) Get() *Infrastructure {
	return v.value
}

func (v *NullableInfrastructure) Set(val *Infrastructure) {
	v.value = val
	v.isSet = true
}

func (v NullableInfrastructure) IsSet() bool {
	return v.isSet
}

func (v *NullableInfrastructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfrastructure(val *Infrastructure) *NullableInfrastructure {
	return &NullableInfrastructure{value: val, isSet: true}
}

func (v NullableInfrastructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfrastructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


