# BookshelfResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authors** | Pointer to [**[]BookshelfAuthorResource**](BookshelfAuthorResource.md) |  | [optional] 
**MonitoringOptions** | Pointer to [**MonitoringOptions**](MonitoringOptions.md) |  | [optional] 
**MonitorNewItems** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 

## Methods

### NewBookshelfResource

`func NewBookshelfResource() *BookshelfResource`

NewBookshelfResource instantiates a new BookshelfResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookshelfResourceWithDefaults

`func NewBookshelfResourceWithDefaults() *BookshelfResource`

NewBookshelfResourceWithDefaults instantiates a new BookshelfResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthors

`func (o *BookshelfResource) GetAuthors() []BookshelfAuthorResource`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *BookshelfResource) GetAuthorsOk() (*[]BookshelfAuthorResource, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *BookshelfResource) SetAuthors(v []BookshelfAuthorResource)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *BookshelfResource) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### SetAuthorsNil

`func (o *BookshelfResource) SetAuthorsNil(b bool)`

 SetAuthorsNil sets the value for Authors to be an explicit nil

### UnsetAuthors
`func (o *BookshelfResource) UnsetAuthors()`

UnsetAuthors ensures that no value is present for Authors, not even an explicit nil
### GetMonitoringOptions

`func (o *BookshelfResource) GetMonitoringOptions() MonitoringOptions`

GetMonitoringOptions returns the MonitoringOptions field if non-nil, zero value otherwise.

### GetMonitoringOptionsOk

`func (o *BookshelfResource) GetMonitoringOptionsOk() (*MonitoringOptions, bool)`

GetMonitoringOptionsOk returns a tuple with the MonitoringOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringOptions

`func (o *BookshelfResource) SetMonitoringOptions(v MonitoringOptions)`

SetMonitoringOptions sets MonitoringOptions field to given value.

### HasMonitoringOptions

`func (o *BookshelfResource) HasMonitoringOptions() bool`

HasMonitoringOptions returns a boolean if a field has been set.

### GetMonitorNewItems

`func (o *BookshelfResource) GetMonitorNewItems() NewItemMonitorTypes`

GetMonitorNewItems returns the MonitorNewItems field if non-nil, zero value otherwise.

### GetMonitorNewItemsOk

`func (o *BookshelfResource) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool)`

GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorNewItems

`func (o *BookshelfResource) SetMonitorNewItems(v NewItemMonitorTypes)`

SetMonitorNewItems sets MonitorNewItems field to given value.

### HasMonitorNewItems

`func (o *BookshelfResource) HasMonitorNewItems() bool`

HasMonitorNewItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


