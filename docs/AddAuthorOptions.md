# AddAuthorOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 
**BooksToMonitor** | Pointer to **[]string** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**SearchForMissingBooks** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddAuthorOptions

`func NewAddAuthorOptions() *AddAuthorOptions`

NewAddAuthorOptions instantiates a new AddAuthorOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAuthorOptionsWithDefaults

`func NewAddAuthorOptionsWithDefaults() *AddAuthorOptions`

NewAddAuthorOptionsWithDefaults instantiates a new AddAuthorOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *AddAuthorOptions) GetMonitor() MonitorTypes`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *AddAuthorOptions) GetMonitorOk() (*MonitorTypes, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *AddAuthorOptions) SetMonitor(v MonitorTypes)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *AddAuthorOptions) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetBooksToMonitor

`func (o *AddAuthorOptions) GetBooksToMonitor() []string`

GetBooksToMonitor returns the BooksToMonitor field if non-nil, zero value otherwise.

### GetBooksToMonitorOk

`func (o *AddAuthorOptions) GetBooksToMonitorOk() (*[]string, bool)`

GetBooksToMonitorOk returns a tuple with the BooksToMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooksToMonitor

`func (o *AddAuthorOptions) SetBooksToMonitor(v []string)`

SetBooksToMonitor sets BooksToMonitor field to given value.

### HasBooksToMonitor

`func (o *AddAuthorOptions) HasBooksToMonitor() bool`

HasBooksToMonitor returns a boolean if a field has been set.

### SetBooksToMonitorNil

`func (o *AddAuthorOptions) SetBooksToMonitorNil(b bool)`

 SetBooksToMonitorNil sets the value for BooksToMonitor to be an explicit nil

### UnsetBooksToMonitor
`func (o *AddAuthorOptions) UnsetBooksToMonitor()`

UnsetBooksToMonitor ensures that no value is present for BooksToMonitor, not even an explicit nil
### GetMonitored

`func (o *AddAuthorOptions) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AddAuthorOptions) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AddAuthorOptions) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AddAuthorOptions) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetSearchForMissingBooks

`func (o *AddAuthorOptions) GetSearchForMissingBooks() bool`

GetSearchForMissingBooks returns the SearchForMissingBooks field if non-nil, zero value otherwise.

### GetSearchForMissingBooksOk

`func (o *AddAuthorOptions) GetSearchForMissingBooksOk() (*bool, bool)`

GetSearchForMissingBooksOk returns a tuple with the SearchForMissingBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForMissingBooks

`func (o *AddAuthorOptions) SetSearchForMissingBooks(v bool)`

SetSearchForMissingBooks sets SearchForMissingBooks field to given value.

### HasSearchForMissingBooks

`func (o *AddAuthorOptions) HasSearchForMissingBooks() bool`

HasSearchForMissingBooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


