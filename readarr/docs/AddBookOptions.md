# AddBookOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddType** | Pointer to [**BookAddType**](BookAddType.md) |  | [optional] 
**SearchForNewBook** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddBookOptions

`func NewAddBookOptions() *AddBookOptions`

NewAddBookOptions instantiates a new AddBookOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBookOptionsWithDefaults

`func NewAddBookOptionsWithDefaults() *AddBookOptions`

NewAddBookOptionsWithDefaults instantiates a new AddBookOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddType

`func (o *AddBookOptions) GetAddType() BookAddType`

GetAddType returns the AddType field if non-nil, zero value otherwise.

### GetAddTypeOk

`func (o *AddBookOptions) GetAddTypeOk() (*BookAddType, bool)`

GetAddTypeOk returns a tuple with the AddType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddType

`func (o *AddBookOptions) SetAddType(v BookAddType)`

SetAddType sets AddType field to given value.

### HasAddType

`func (o *AddBookOptions) HasAddType() bool`

HasAddType returns a boolean if a field has been set.

### GetSearchForNewBook

`func (o *AddBookOptions) GetSearchForNewBook() bool`

GetSearchForNewBook returns the SearchForNewBook field if non-nil, zero value otherwise.

### GetSearchForNewBookOk

`func (o *AddBookOptions) GetSearchForNewBookOk() (*bool, bool)`

GetSearchForNewBookOk returns a tuple with the SearchForNewBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForNewBook

`func (o *AddBookOptions) SetSearchForNewBook(v bool)`

SetSearchForNewBook sets SearchForNewBook field to given value.

### HasSearchForNewBook

`func (o *AddBookOptions) HasSearchForNewBook() bool`

HasSearchForNewBook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


