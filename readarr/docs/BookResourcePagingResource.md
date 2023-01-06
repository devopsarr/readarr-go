# BookResourcePagingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**SortKey** | Pointer to **NullableString** |  | [optional] 
**SortDirection** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] 
**Filters** | Pointer to [**[]PagingResourceFilter**](PagingResourceFilter.md) |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**Records** | Pointer to [**[]BookResource**](BookResource.md) |  | [optional] 

## Methods

### NewBookResourcePagingResource

`func NewBookResourcePagingResource() *BookResourcePagingResource`

NewBookResourcePagingResource instantiates a new BookResourcePagingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookResourcePagingResourceWithDefaults

`func NewBookResourcePagingResourceWithDefaults() *BookResourcePagingResource`

NewBookResourcePagingResourceWithDefaults instantiates a new BookResourcePagingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *BookResourcePagingResource) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BookResourcePagingResource) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BookResourcePagingResource) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *BookResourcePagingResource) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *BookResourcePagingResource) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *BookResourcePagingResource) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *BookResourcePagingResource) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *BookResourcePagingResource) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *BookResourcePagingResource) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *BookResourcePagingResource) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *BookResourcePagingResource) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *BookResourcePagingResource) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### SetSortKeyNil

`func (o *BookResourcePagingResource) SetSortKeyNil(b bool)`

 SetSortKeyNil sets the value for SortKey to be an explicit nil

### UnsetSortKey
`func (o *BookResourcePagingResource) UnsetSortKey()`

UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
### GetSortDirection

`func (o *BookResourcePagingResource) GetSortDirection() SortDirection`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *BookResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *BookResourcePagingResource) SetSortDirection(v SortDirection)`

SetSortDirection sets SortDirection field to given value.

### HasSortDirection

`func (o *BookResourcePagingResource) HasSortDirection() bool`

HasSortDirection returns a boolean if a field has been set.

### GetFilters

`func (o *BookResourcePagingResource) GetFilters() []PagingResourceFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BookResourcePagingResource) GetFiltersOk() (*[]PagingResourceFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BookResourcePagingResource) SetFilters(v []PagingResourceFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BookResourcePagingResource) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *BookResourcePagingResource) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *BookResourcePagingResource) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetTotalRecords

`func (o *BookResourcePagingResource) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *BookResourcePagingResource) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *BookResourcePagingResource) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *BookResourcePagingResource) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetRecords

`func (o *BookResourcePagingResource) GetRecords() []BookResource`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *BookResourcePagingResource) GetRecordsOk() (*[]BookResource, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *BookResourcePagingResource) SetRecords(v []BookResource)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *BookResourcePagingResource) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### SetRecordsNil

`func (o *BookResourcePagingResource) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *BookResourcePagingResource) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


