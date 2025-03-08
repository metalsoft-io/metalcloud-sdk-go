# ServerInstanceStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerStatus** | **map[string]interface{}** | An object containing server status as key and count as value | 
**Site** | **map[string]interface{}** | An object containing server site name as key and count as value | 

## Methods

### NewServerInstanceStatistics

`func NewServerInstanceStatistics(serverStatus map[string]interface{}, site map[string]interface{}, ) *ServerInstanceStatistics`

NewServerInstanceStatistics instantiates a new ServerInstanceStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceStatisticsWithDefaults

`func NewServerInstanceStatisticsWithDefaults() *ServerInstanceStatistics`

NewServerInstanceStatisticsWithDefaults instantiates a new ServerInstanceStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerStatus

`func (o *ServerInstanceStatistics) GetServerStatus() map[string]interface{}`

GetServerStatus returns the ServerStatus field if non-nil, zero value otherwise.

### GetServerStatusOk

`func (o *ServerInstanceStatistics) GetServerStatusOk() (*map[string]interface{}, bool)`

GetServerStatusOk returns a tuple with the ServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatus

`func (o *ServerInstanceStatistics) SetServerStatus(v map[string]interface{})`

SetServerStatus sets ServerStatus field to given value.


### GetSite

`func (o *ServerInstanceStatistics) GetSite() map[string]interface{}`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ServerInstanceStatistics) GetSiteOk() (*map[string]interface{}, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ServerInstanceStatistics) SetSite(v map[string]interface{})`

SetSite sets Site field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


