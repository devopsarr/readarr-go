# MonitoringOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 
**BooksToMonitor** | Pointer to **[]string** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 

## Methods

### NewMonitoringOptions

`func NewMonitoringOptions() *MonitoringOptions`

NewMonitoringOptions instantiates a new MonitoringOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringOptionsWithDefaults

`func NewMonitoringOptionsWithDefaults() *MonitoringOptions`

NewMonitoringOptionsWithDefaults instantiates a new MonitoringOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *MonitoringOptions) GetMonitor() MonitorTypes`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *MonitoringOptions) GetMonitorOk() (*MonitorTypes, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *MonitoringOptions) SetMonitor(v MonitorTypes)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *MonitoringOptions) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetBooksToMonitor

`func (o *MonitoringOptions) GetBooksToMonitor() []string`

GetBooksToMonitor returns the BooksToMonitor field if non-nil, zero value otherwise.

### GetBooksToMonitorOk

`func (o *MonitoringOptions) GetBooksToMonitorOk() (*[]string, bool)`

GetBooksToMonitorOk returns a tuple with the BooksToMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooksToMonitor

`func (o *MonitoringOptions) SetBooksToMonitor(v []string)`

SetBooksToMonitor sets BooksToMonitor field to given value.

### HasBooksToMonitor

`func (o *MonitoringOptions) HasBooksToMonitor() bool`

HasBooksToMonitor returns a boolean if a field has been set.

### SetBooksToMonitorNil

`func (o *MonitoringOptions) SetBooksToMonitorNil(b bool)`

 SetBooksToMonitorNil sets the value for BooksToMonitor to be an explicit nil

### UnsetBooksToMonitor
`func (o *MonitoringOptions) UnsetBooksToMonitor()`

UnsetBooksToMonitor ensures that no value is present for BooksToMonitor, not even an explicit nil
### GetMonitored

`func (o *MonitoringOptions) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *MonitoringOptions) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *MonitoringOptions) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *MonitoringOptions) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


