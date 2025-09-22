# ExtensionDnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | DNS name, allows the use of placeholder variables | 
**RecordType** | [**DNSRecordType**](DNSRecordType.md) | The type of DNS record (e.g., A, AAAA, CNAME, NS, PTR, TXT, SOA) to create | 
**GeneratePtrRecord** | Pointer to **bool** | Whether to generate a PTR record for this DNS record | [optional] [default to false]
**Aliases** | Pointer to **[]string** | CNAME aliases for the DNS record | [optional] 
**Ttl** | Pointer to **int32** | TTL (Time to Live) for the DNS Record Set. | [optional] [default to 3600]
**Tags** | Pointer to **map[string]string** | The tags associated with the DNS Record Set | [optional] 

## Methods

### NewExtensionDnsRecord

`func NewExtensionDnsRecord(name string, recordType DNSRecordType, ) *ExtensionDnsRecord`

NewExtensionDnsRecord instantiates a new ExtensionDnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionDnsRecordWithDefaults

`func NewExtensionDnsRecordWithDefaults() *ExtensionDnsRecord`

NewExtensionDnsRecordWithDefaults instantiates a new ExtensionDnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtensionDnsRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionDnsRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionDnsRecord) SetName(v string)`

SetName sets Name field to given value.


### GetRecordType

`func (o *ExtensionDnsRecord) GetRecordType() DNSRecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ExtensionDnsRecord) GetRecordTypeOk() (*DNSRecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ExtensionDnsRecord) SetRecordType(v DNSRecordType)`

SetRecordType sets RecordType field to given value.


### GetGeneratePtrRecord

`func (o *ExtensionDnsRecord) GetGeneratePtrRecord() bool`

GetGeneratePtrRecord returns the GeneratePtrRecord field if non-nil, zero value otherwise.

### GetGeneratePtrRecordOk

`func (o *ExtensionDnsRecord) GetGeneratePtrRecordOk() (*bool, bool)`

GetGeneratePtrRecordOk returns a tuple with the GeneratePtrRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratePtrRecord

`func (o *ExtensionDnsRecord) SetGeneratePtrRecord(v bool)`

SetGeneratePtrRecord sets GeneratePtrRecord field to given value.

### HasGeneratePtrRecord

`func (o *ExtensionDnsRecord) HasGeneratePtrRecord() bool`

HasGeneratePtrRecord returns a boolean if a field has been set.

### GetAliases

`func (o *ExtensionDnsRecord) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ExtensionDnsRecord) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ExtensionDnsRecord) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ExtensionDnsRecord) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetTtl

`func (o *ExtensionDnsRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ExtensionDnsRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ExtensionDnsRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ExtensionDnsRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTags

`func (o *ExtensionDnsRecord) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtensionDnsRecord) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtensionDnsRecord) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtensionDnsRecord) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


