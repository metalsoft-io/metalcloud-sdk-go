# SNMPOIDGroupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the SNMP OID group | 
**Oids** | **map[string]interface{}** | A map of SNMP OIDs to Label in the group { OID &#x3D;&gt; Label } | 
**Mapping** | [**SNMPOIDToMetricMapping**](SNMPOIDToMetricMapping.md) | Mapping configuration to convert SNMP OID values into monitoring metrics | 

## Methods

### NewSNMPOIDGroupConfig

`func NewSNMPOIDGroupConfig(name string, oids map[string]interface{}, mapping SNMPOIDToMetricMapping, ) *SNMPOIDGroupConfig`

NewSNMPOIDGroupConfig instantiates a new SNMPOIDGroupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPOIDGroupConfigWithDefaults

`func NewSNMPOIDGroupConfigWithDefaults() *SNMPOIDGroupConfig`

NewSNMPOIDGroupConfigWithDefaults instantiates a new SNMPOIDGroupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SNMPOIDGroupConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SNMPOIDGroupConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SNMPOIDGroupConfig) SetName(v string)`

SetName sets Name field to given value.


### GetOids

`func (o *SNMPOIDGroupConfig) GetOids() map[string]interface{}`

GetOids returns the Oids field if non-nil, zero value otherwise.

### GetOidsOk

`func (o *SNMPOIDGroupConfig) GetOidsOk() (*map[string]interface{}, bool)`

GetOidsOk returns a tuple with the Oids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOids

`func (o *SNMPOIDGroupConfig) SetOids(v map[string]interface{})`

SetOids sets Oids field to given value.


### GetMapping

`func (o *SNMPOIDGroupConfig) GetMapping() SNMPOIDToMetricMapping`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *SNMPOIDGroupConfig) GetMappingOk() (*SNMPOIDToMetricMapping, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *SNMPOIDGroupConfig) SetMapping(v SNMPOIDToMetricMapping)`

SetMapping sets Mapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


