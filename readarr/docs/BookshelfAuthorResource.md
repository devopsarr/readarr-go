# BookshelfAuthorResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**Books** | Pointer to [**[]BookResource**](BookResource.md) |  | [optional] 

## Methods

### NewBookshelfAuthorResource

`func NewBookshelfAuthorResource() *BookshelfAuthorResource`

NewBookshelfAuthorResource instantiates a new BookshelfAuthorResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookshelfAuthorResourceWithDefaults

`func NewBookshelfAuthorResourceWithDefaults() *BookshelfAuthorResource`

NewBookshelfAuthorResourceWithDefaults instantiates a new BookshelfAuthorResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookshelfAuthorResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookshelfAuthorResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookshelfAuthorResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BookshelfAuthorResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonitored

`func (o *BookshelfAuthorResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *BookshelfAuthorResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *BookshelfAuthorResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *BookshelfAuthorResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *BookshelfAuthorResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *BookshelfAuthorResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetBooks

`func (o *BookshelfAuthorResource) GetBooks() []BookResource`

GetBooks returns the Books field if non-nil, zero value otherwise.

### GetBooksOk

`func (o *BookshelfAuthorResource) GetBooksOk() (*[]BookResource, bool)`

GetBooksOk returns a tuple with the Books field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooks

`func (o *BookshelfAuthorResource) SetBooks(v []BookResource)`

SetBooks sets Books field to given value.

### HasBooks

`func (o *BookshelfAuthorResource) HasBooks() bool`

HasBooks returns a boolean if a field has been set.

### SetBooksNil

`func (o *BookshelfAuthorResource) SetBooksNil(b bool)`

 SetBooksNil sets the value for Books to be an explicit nil

### UnsetBooks
`func (o *BookshelfAuthorResource) UnsetBooks()`

UnsetBooks ensures that no value is present for Books, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


