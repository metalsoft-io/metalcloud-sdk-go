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

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event struct for Event
type Event struct {
	// The id of the event
	Id float32 `json:"id"`
	// The id of the user who triggered the event
	UserIdAuthenticated *float32 `json:"userIdAuthenticated,omitempty"`
	// The type of the event
	Type float32 `json:"type"`
	// The severity of the event
	Severity string `json:"severity"`
	// The visibility of the event
	Visibility string `json:"visibility"`
	// The id of the infrastructure linked to the event
	InfrastructureId *float32 `json:"infrastructureId,omitempty"`
	// The id of the user linked to the event
	UserId *float32 `json:"userId,omitempty"`
	// The id of the server linked to the event
	ServerId *float32 `json:"serverId,omitempty"`
	// The id of the job linked to the event
	JobId *string `json:"jobId,omitempty"`
	// The id of the site linked to the event
	SiteId *float32 `json:"siteId,omitempty"`
	// The title of the event
	Title string `json:"title"`
	// The message of the event
	Message string `json:"message"`
	// The email of the user who triggered the event
	UserEmailAuthenticated *string `json:"userEmailAuthenticated,omitempty"`
	// The context of the event
	Context map[string]interface{} `json:"context,omitempty"`
	// The timestamp the event occurred
	OccurredTimestamp string `json:"occurredTimestamp"`
	// The http user agent linked to the event
	HttpUserAgent *string `json:"httpUserAgent,omitempty"`
	// The remote ip address of the user who triggered the event
	RemoteIpAddress *string `json:"remoteIpAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Event Event

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(id float32, type_ float32, severity string, visibility string, title string, message string, occurredTimestamp string) *Event {
	this := Event{}
	this.Id = id
	this.Type = type_
	this.Severity = severity
	this.Visibility = visibility
	this.Title = title
	this.Message = message
	this.OccurredTimestamp = occurredTimestamp
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetId returns the Id field value
func (o *Event) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Event) SetId(v float32) {
	o.Id = v
}

// GetUserIdAuthenticated returns the UserIdAuthenticated field value if set, zero value otherwise.
func (o *Event) GetUserIdAuthenticated() float32 {
	if o == nil || IsNil(o.UserIdAuthenticated) {
		var ret float32
		return ret
	}
	return *o.UserIdAuthenticated
}

// GetUserIdAuthenticatedOk returns a tuple with the UserIdAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetUserIdAuthenticatedOk() (*float32, bool) {
	if o == nil || IsNil(o.UserIdAuthenticated) {
		return nil, false
	}
	return o.UserIdAuthenticated, true
}

// HasUserIdAuthenticated returns a boolean if a field has been set.
func (o *Event) HasUserIdAuthenticated() bool {
	if o != nil && !IsNil(o.UserIdAuthenticated) {
		return true
	}

	return false
}

// SetUserIdAuthenticated gets a reference to the given float32 and assigns it to the UserIdAuthenticated field.
func (o *Event) SetUserIdAuthenticated(v float32) {
	o.UserIdAuthenticated = &v
}

// GetType returns the Type field value
func (o *Event) GetType() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Event) SetType(v float32) {
	o.Type = v
}

// GetSeverity returns the Severity field value
func (o *Event) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Event) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Event) SetSeverity(v string) {
	o.Severity = v
}

// GetVisibility returns the Visibility field value
func (o *Event) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *Event) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *Event) SetVisibility(v string) {
	o.Visibility = v
}

// GetInfrastructureId returns the InfrastructureId field value if set, zero value otherwise.
func (o *Event) GetInfrastructureId() float32 {
	if o == nil || IsNil(o.InfrastructureId) {
		var ret float32
		return ret
	}
	return *o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetInfrastructureIdOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrastructureId) {
		return nil, false
	}
	return o.InfrastructureId, true
}

// HasInfrastructureId returns a boolean if a field has been set.
func (o *Event) HasInfrastructureId() bool {
	if o != nil && !IsNil(o.InfrastructureId) {
		return true
	}

	return false
}

// SetInfrastructureId gets a reference to the given float32 and assigns it to the InfrastructureId field.
func (o *Event) SetInfrastructureId(v float32) {
	o.InfrastructureId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Event) GetUserId() float32 {
	if o == nil || IsNil(o.UserId) {
		var ret float32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetUserIdOk() (*float32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Event) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given float32 and assigns it to the UserId field.
func (o *Event) SetUserId(v float32) {
	o.UserId = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *Event) GetServerId() float32 {
	if o == nil || IsNil(o.ServerId) {
		var ret float32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetServerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *Event) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given float32 and assigns it to the ServerId field.
func (o *Event) SetServerId(v float32) {
	o.ServerId = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *Event) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *Event) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *Event) SetJobId(v string) {
	o.JobId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Event) GetSiteId() float32 {
	if o == nil || IsNil(o.SiteId) {
		var ret float32
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSiteIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Event) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given float32 and assigns it to the SiteId field.
func (o *Event) SetSiteId(v float32) {
	o.SiteId = &v
}

// GetTitle returns the Title field value
func (o *Event) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Event) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Event) SetTitle(v string) {
	o.Title = v
}

// GetMessage returns the Message field value
func (o *Event) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Event) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Event) SetMessage(v string) {
	o.Message = v
}

// GetUserEmailAuthenticated returns the UserEmailAuthenticated field value if set, zero value otherwise.
func (o *Event) GetUserEmailAuthenticated() string {
	if o == nil || IsNil(o.UserEmailAuthenticated) {
		var ret string
		return ret
	}
	return *o.UserEmailAuthenticated
}

// GetUserEmailAuthenticatedOk returns a tuple with the UserEmailAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetUserEmailAuthenticatedOk() (*string, bool) {
	if o == nil || IsNil(o.UserEmailAuthenticated) {
		return nil, false
	}
	return o.UserEmailAuthenticated, true
}

// HasUserEmailAuthenticated returns a boolean if a field has been set.
func (o *Event) HasUserEmailAuthenticated() bool {
	if o != nil && !IsNil(o.UserEmailAuthenticated) {
		return true
	}

	return false
}

// SetUserEmailAuthenticated gets a reference to the given string and assigns it to the UserEmailAuthenticated field.
func (o *Event) SetUserEmailAuthenticated(v string) {
	o.UserEmailAuthenticated = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *Event) GetContext() map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *Event) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *Event) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetOccurredTimestamp returns the OccurredTimestamp field value
func (o *Event) GetOccurredTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OccurredTimestamp
}

// GetOccurredTimestampOk returns a tuple with the OccurredTimestamp field value
// and a boolean to check if the value has been set.
func (o *Event) GetOccurredTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OccurredTimestamp, true
}

// SetOccurredTimestamp sets field value
func (o *Event) SetOccurredTimestamp(v string) {
	o.OccurredTimestamp = v
}

// GetHttpUserAgent returns the HttpUserAgent field value if set, zero value otherwise.
func (o *Event) GetHttpUserAgent() string {
	if o == nil || IsNil(o.HttpUserAgent) {
		var ret string
		return ret
	}
	return *o.HttpUserAgent
}

// GetHttpUserAgentOk returns a tuple with the HttpUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetHttpUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.HttpUserAgent) {
		return nil, false
	}
	return o.HttpUserAgent, true
}

// HasHttpUserAgent returns a boolean if a field has been set.
func (o *Event) HasHttpUserAgent() bool {
	if o != nil && !IsNil(o.HttpUserAgent) {
		return true
	}

	return false
}

// SetHttpUserAgent gets a reference to the given string and assigns it to the HttpUserAgent field.
func (o *Event) SetHttpUserAgent(v string) {
	o.HttpUserAgent = &v
}

// GetRemoteIpAddress returns the RemoteIpAddress field value if set, zero value otherwise.
func (o *Event) GetRemoteIpAddress() string {
	if o == nil || IsNil(o.RemoteIpAddress) {
		var ret string
		return ret
	}
	return *o.RemoteIpAddress
}

// GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetRemoteIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteIpAddress) {
		return nil, false
	}
	return o.RemoteIpAddress, true
}

// HasRemoteIpAddress returns a boolean if a field has been set.
func (o *Event) HasRemoteIpAddress() bool {
	if o != nil && !IsNil(o.RemoteIpAddress) {
		return true
	}

	return false
}

// SetRemoteIpAddress gets a reference to the given string and assigns it to the RemoteIpAddress field.
func (o *Event) SetRemoteIpAddress(v string) {
	o.RemoteIpAddress = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.UserIdAuthenticated) {
		toSerialize["userIdAuthenticated"] = o.UserIdAuthenticated
	}
	toSerialize["type"] = o.Type
	toSerialize["severity"] = o.Severity
	toSerialize["visibility"] = o.Visibility
	if !IsNil(o.InfrastructureId) {
		toSerialize["infrastructureId"] = o.InfrastructureId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	toSerialize["title"] = o.Title
	toSerialize["message"] = o.Message
	if !IsNil(o.UserEmailAuthenticated) {
		toSerialize["userEmailAuthenticated"] = o.UserEmailAuthenticated
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	toSerialize["occurredTimestamp"] = o.OccurredTimestamp
	if !IsNil(o.HttpUserAgent) {
		toSerialize["httpUserAgent"] = o.HttpUserAgent
	}
	if !IsNil(o.RemoteIpAddress) {
		toSerialize["remoteIpAddress"] = o.RemoteIpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Event) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"severity",
		"visibility",
		"title",
		"message",
		"occurredTimestamp",
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

	varEvent := _Event{}

	err = json.Unmarshal(data, &varEvent)

	if err != nil {
		return err
	}

	*o = Event(varEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "userIdAuthenticated")
		delete(additionalProperties, "type")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "visibility")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "jobId")
		delete(additionalProperties, "siteId")
		delete(additionalProperties, "title")
		delete(additionalProperties, "message")
		delete(additionalProperties, "userEmailAuthenticated")
		delete(additionalProperties, "context")
		delete(additionalProperties, "occurredTimestamp")
		delete(additionalProperties, "httpUserAgent")
		delete(additionalProperties, "remoteIpAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


