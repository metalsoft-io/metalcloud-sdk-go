# SiteStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteControllerSeenAliveStatus** | [**SiteControllerSeenAliveStatus**](SiteControllerSeenAliveStatus.md) | Status of site controllers categorized by their alive status | 
**SitesTotalCount** | **float32** | Total count of sites | 
**SitesActiveCount** | **float32** | Count of active sites | 
**SitesResourceCount** | [**[]SiteControllerSeenAliveStatus**](SiteControllerSeenAliveStatus.md) | List of resources counts for sites | 

## Methods

### NewSiteStatistics

`func NewSiteStatistics(siteControllerSeenAliveStatus SiteControllerSeenAliveStatus, sitesTotalCount float32, sitesActiveCount float32, sitesResourceCount []SiteControllerSeenAliveStatus, ) *SiteStatistics`

NewSiteStatistics instantiates a new SiteStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteStatisticsWithDefaults

`func NewSiteStatisticsWithDefaults() *SiteStatistics`

NewSiteStatisticsWithDefaults instantiates a new SiteStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteControllerSeenAliveStatus

`func (o *SiteStatistics) GetSiteControllerSeenAliveStatus() SiteControllerSeenAliveStatus`

GetSiteControllerSeenAliveStatus returns the SiteControllerSeenAliveStatus field if non-nil, zero value otherwise.

### GetSiteControllerSeenAliveStatusOk

`func (o *SiteStatistics) GetSiteControllerSeenAliveStatusOk() (*SiteControllerSeenAliveStatus, bool)`

GetSiteControllerSeenAliveStatusOk returns a tuple with the SiteControllerSeenAliveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteControllerSeenAliveStatus

`func (o *SiteStatistics) SetSiteControllerSeenAliveStatus(v SiteControllerSeenAliveStatus)`

SetSiteControllerSeenAliveStatus sets SiteControllerSeenAliveStatus field to given value.


### GetSitesTotalCount

`func (o *SiteStatistics) GetSitesTotalCount() float32`

GetSitesTotalCount returns the SitesTotalCount field if non-nil, zero value otherwise.

### GetSitesTotalCountOk

`func (o *SiteStatistics) GetSitesTotalCountOk() (*float32, bool)`

GetSitesTotalCountOk returns a tuple with the SitesTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesTotalCount

`func (o *SiteStatistics) SetSitesTotalCount(v float32)`

SetSitesTotalCount sets SitesTotalCount field to given value.


### GetSitesActiveCount

`func (o *SiteStatistics) GetSitesActiveCount() float32`

GetSitesActiveCount returns the SitesActiveCount field if non-nil, zero value otherwise.

### GetSitesActiveCountOk

`func (o *SiteStatistics) GetSitesActiveCountOk() (*float32, bool)`

GetSitesActiveCountOk returns a tuple with the SitesActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesActiveCount

`func (o *SiteStatistics) SetSitesActiveCount(v float32)`

SetSitesActiveCount sets SitesActiveCount field to given value.


### GetSitesResourceCount

`func (o *SiteStatistics) GetSitesResourceCount() []SiteControllerSeenAliveStatus`

GetSitesResourceCount returns the SitesResourceCount field if non-nil, zero value otherwise.

### GetSitesResourceCountOk

`func (o *SiteStatistics) GetSitesResourceCountOk() (*[]SiteControllerSeenAliveStatus, bool)`

GetSitesResourceCountOk returns a tuple with the SitesResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesResourceCount

`func (o *SiteStatistics) SetSitesResourceCount(v []SiteControllerSeenAliveStatus)`

SetSitesResourceCount sets SitesResourceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


