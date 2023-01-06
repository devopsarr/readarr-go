# BookListLazyLoaded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]Book**](Book.md) |  | [optional] [readonly] 
**IsLoaded** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewBookListLazyLoaded

`func NewBookListLazyLoaded() *BookListLazyLoaded`

NewBookListLazyLoaded instantiates a new BookListLazyLoaded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookListLazyLoadedWithDefaults

`func NewBookListLazyLoadedWithDefaults() *BookListLazyLoaded`

NewBookListLazyLoadedWithDefaults instantiates a new BookListLazyLoaded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *BookListLazyLoaded) GetValue() []Book`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BookListLazyLoaded) GetValueOk() (*[]Book, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BookListLazyLoaded) SetValue(v []Book)`

SetValue sets Value field to given value.

### HasValue

`func (o *BookListLazyLoaded) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BookListLazyLoaded) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BookListLazyLoaded) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsLoaded

`func (o *BookListLazyLoaded) GetIsLoaded() bool`

GetIsLoaded returns the IsLoaded field if non-nil, zero value otherwise.

### GetIsLoadedOk

`func (o *BookListLazyLoaded) GetIsLoadedOk() (*bool, bool)`

GetIsLoadedOk returns a tuple with the IsLoaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoaded

`func (o *BookListLazyLoaded) SetIsLoaded(v bool)`

SetIsLoaded sets IsLoaded field to given value.

### HasIsLoaded

`func (o *BookListLazyLoaded) HasIsLoaded() bool`

HasIsLoaded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


