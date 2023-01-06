# ParseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ParsedBookInfo** | Pointer to [**ParsedBookInfo**](ParsedBookInfo.md) |  | [optional] 
**Author** | Pointer to [**AuthorResource**](AuthorResource.md) |  | [optional] 
**Books** | Pointer to [**[]BookResource**](BookResource.md) |  | [optional] 

## Methods

### NewParseResource

`func NewParseResource() *ParseResource`

NewParseResource instantiates a new ParseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseResourceWithDefaults

`func NewParseResourceWithDefaults() *ParseResource`

NewParseResourceWithDefaults instantiates a new ParseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParseResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParseResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParseResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ParseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ParseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ParseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ParseResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ParseResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetParsedBookInfo

`func (o *ParseResource) GetParsedBookInfo() ParsedBookInfo`

GetParsedBookInfo returns the ParsedBookInfo field if non-nil, zero value otherwise.

### GetParsedBookInfoOk

`func (o *ParseResource) GetParsedBookInfoOk() (*ParsedBookInfo, bool)`

GetParsedBookInfoOk returns a tuple with the ParsedBookInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedBookInfo

`func (o *ParseResource) SetParsedBookInfo(v ParsedBookInfo)`

SetParsedBookInfo sets ParsedBookInfo field to given value.

### HasParsedBookInfo

`func (o *ParseResource) HasParsedBookInfo() bool`

HasParsedBookInfo returns a boolean if a field has been set.

### GetAuthor

`func (o *ParseResource) GetAuthor() AuthorResource`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ParseResource) GetAuthorOk() (*AuthorResource, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ParseResource) SetAuthor(v AuthorResource)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ParseResource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBooks

`func (o *ParseResource) GetBooks() []BookResource`

GetBooks returns the Books field if non-nil, zero value otherwise.

### GetBooksOk

`func (o *ParseResource) GetBooksOk() (*[]BookResource, bool)`

GetBooksOk returns a tuple with the Books field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooks

`func (o *ParseResource) SetBooks(v []BookResource)`

SetBooks sets Books field to given value.

### HasBooks

`func (o *ParseResource) HasBooks() bool`

HasBooks returns a boolean if a field has been set.

### SetBooksNil

`func (o *ParseResource) SetBooksNil(b bool)`

 SetBooksNil sets the value for Books to be an explicit nil

### UnsetBooks
`func (o *ParseResource) UnsetBooks()`

UnsetBooks ensures that no value is present for Books, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


