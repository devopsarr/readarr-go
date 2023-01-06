# NamingConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RenameBooks** | Pointer to **bool** |  | [optional] 
**ReplaceIllegalCharacters** | Pointer to **bool** |  | [optional] 
**StandardBookFormat** | Pointer to **NullableString** |  | [optional] 
**AuthorFolderFormat** | Pointer to **NullableString** |  | [optional] 
**IncludeAuthorName** | Pointer to **bool** |  | [optional] 
**IncludeBookTitle** | Pointer to **bool** |  | [optional] 
**IncludeQuality** | Pointer to **bool** |  | [optional] 
**ReplaceSpaces** | Pointer to **bool** |  | [optional] 
**Separator** | Pointer to **NullableString** |  | [optional] 
**NumberStyle** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNamingConfigResource

`func NewNamingConfigResource() *NamingConfigResource`

NewNamingConfigResource instantiates a new NamingConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingConfigResourceWithDefaults

`func NewNamingConfigResourceWithDefaults() *NamingConfigResource`

NewNamingConfigResourceWithDefaults instantiates a new NamingConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NamingConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRenameBooks

`func (o *NamingConfigResource) GetRenameBooks() bool`

GetRenameBooks returns the RenameBooks field if non-nil, zero value otherwise.

### GetRenameBooksOk

`func (o *NamingConfigResource) GetRenameBooksOk() (*bool, bool)`

GetRenameBooksOk returns a tuple with the RenameBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameBooks

`func (o *NamingConfigResource) SetRenameBooks(v bool)`

SetRenameBooks sets RenameBooks field to given value.

### HasRenameBooks

`func (o *NamingConfigResource) HasRenameBooks() bool`

HasRenameBooks returns a boolean if a field has been set.

### GetReplaceIllegalCharacters

`func (o *NamingConfigResource) GetReplaceIllegalCharacters() bool`

GetReplaceIllegalCharacters returns the ReplaceIllegalCharacters field if non-nil, zero value otherwise.

### GetReplaceIllegalCharactersOk

`func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool)`

GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceIllegalCharacters

`func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool)`

SetReplaceIllegalCharacters sets ReplaceIllegalCharacters field to given value.

### HasReplaceIllegalCharacters

`func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool`

HasReplaceIllegalCharacters returns a boolean if a field has been set.

### GetStandardBookFormat

`func (o *NamingConfigResource) GetStandardBookFormat() string`

GetStandardBookFormat returns the StandardBookFormat field if non-nil, zero value otherwise.

### GetStandardBookFormatOk

`func (o *NamingConfigResource) GetStandardBookFormatOk() (*string, bool)`

GetStandardBookFormatOk returns a tuple with the StandardBookFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardBookFormat

`func (o *NamingConfigResource) SetStandardBookFormat(v string)`

SetStandardBookFormat sets StandardBookFormat field to given value.

### HasStandardBookFormat

`func (o *NamingConfigResource) HasStandardBookFormat() bool`

HasStandardBookFormat returns a boolean if a field has been set.

### SetStandardBookFormatNil

`func (o *NamingConfigResource) SetStandardBookFormatNil(b bool)`

 SetStandardBookFormatNil sets the value for StandardBookFormat to be an explicit nil

### UnsetStandardBookFormat
`func (o *NamingConfigResource) UnsetStandardBookFormat()`

UnsetStandardBookFormat ensures that no value is present for StandardBookFormat, not even an explicit nil
### GetAuthorFolderFormat

`func (o *NamingConfigResource) GetAuthorFolderFormat() string`

GetAuthorFolderFormat returns the AuthorFolderFormat field if non-nil, zero value otherwise.

### GetAuthorFolderFormatOk

`func (o *NamingConfigResource) GetAuthorFolderFormatOk() (*string, bool)`

GetAuthorFolderFormatOk returns a tuple with the AuthorFolderFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorFolderFormat

`func (o *NamingConfigResource) SetAuthorFolderFormat(v string)`

SetAuthorFolderFormat sets AuthorFolderFormat field to given value.

### HasAuthorFolderFormat

`func (o *NamingConfigResource) HasAuthorFolderFormat() bool`

HasAuthorFolderFormat returns a boolean if a field has been set.

### SetAuthorFolderFormatNil

`func (o *NamingConfigResource) SetAuthorFolderFormatNil(b bool)`

 SetAuthorFolderFormatNil sets the value for AuthorFolderFormat to be an explicit nil

### UnsetAuthorFolderFormat
`func (o *NamingConfigResource) UnsetAuthorFolderFormat()`

UnsetAuthorFolderFormat ensures that no value is present for AuthorFolderFormat, not even an explicit nil
### GetIncludeAuthorName

`func (o *NamingConfigResource) GetIncludeAuthorName() bool`

GetIncludeAuthorName returns the IncludeAuthorName field if non-nil, zero value otherwise.

### GetIncludeAuthorNameOk

`func (o *NamingConfigResource) GetIncludeAuthorNameOk() (*bool, bool)`

GetIncludeAuthorNameOk returns a tuple with the IncludeAuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAuthorName

`func (o *NamingConfigResource) SetIncludeAuthorName(v bool)`

SetIncludeAuthorName sets IncludeAuthorName field to given value.

### HasIncludeAuthorName

`func (o *NamingConfigResource) HasIncludeAuthorName() bool`

HasIncludeAuthorName returns a boolean if a field has been set.

### GetIncludeBookTitle

`func (o *NamingConfigResource) GetIncludeBookTitle() bool`

GetIncludeBookTitle returns the IncludeBookTitle field if non-nil, zero value otherwise.

### GetIncludeBookTitleOk

`func (o *NamingConfigResource) GetIncludeBookTitleOk() (*bool, bool)`

GetIncludeBookTitleOk returns a tuple with the IncludeBookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBookTitle

`func (o *NamingConfigResource) SetIncludeBookTitle(v bool)`

SetIncludeBookTitle sets IncludeBookTitle field to given value.

### HasIncludeBookTitle

`func (o *NamingConfigResource) HasIncludeBookTitle() bool`

HasIncludeBookTitle returns a boolean if a field has been set.

### GetIncludeQuality

`func (o *NamingConfigResource) GetIncludeQuality() bool`

GetIncludeQuality returns the IncludeQuality field if non-nil, zero value otherwise.

### GetIncludeQualityOk

`func (o *NamingConfigResource) GetIncludeQualityOk() (*bool, bool)`

GetIncludeQualityOk returns a tuple with the IncludeQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeQuality

`func (o *NamingConfigResource) SetIncludeQuality(v bool)`

SetIncludeQuality sets IncludeQuality field to given value.

### HasIncludeQuality

`func (o *NamingConfigResource) HasIncludeQuality() bool`

HasIncludeQuality returns a boolean if a field has been set.

### GetReplaceSpaces

`func (o *NamingConfigResource) GetReplaceSpaces() bool`

GetReplaceSpaces returns the ReplaceSpaces field if non-nil, zero value otherwise.

### GetReplaceSpacesOk

`func (o *NamingConfigResource) GetReplaceSpacesOk() (*bool, bool)`

GetReplaceSpacesOk returns a tuple with the ReplaceSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceSpaces

`func (o *NamingConfigResource) SetReplaceSpaces(v bool)`

SetReplaceSpaces sets ReplaceSpaces field to given value.

### HasReplaceSpaces

`func (o *NamingConfigResource) HasReplaceSpaces() bool`

HasReplaceSpaces returns a boolean if a field has been set.

### GetSeparator

`func (o *NamingConfigResource) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *NamingConfigResource) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *NamingConfigResource) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *NamingConfigResource) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *NamingConfigResource) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *NamingConfigResource) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
### GetNumberStyle

`func (o *NamingConfigResource) GetNumberStyle() string`

GetNumberStyle returns the NumberStyle field if non-nil, zero value otherwise.

### GetNumberStyleOk

`func (o *NamingConfigResource) GetNumberStyleOk() (*string, bool)`

GetNumberStyleOk returns a tuple with the NumberStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberStyle

`func (o *NamingConfigResource) SetNumberStyle(v string)`

SetNumberStyle sets NumberStyle field to given value.

### HasNumberStyle

`func (o *NamingConfigResource) HasNumberStyle() bool`

HasNumberStyle returns a boolean if a field has been set.

### SetNumberStyleNil

`func (o *NamingConfigResource) SetNumberStyleNil(b bool)`

 SetNumberStyleNil sets the value for NumberStyle to be an explicit nil

### UnsetNumberStyle
`func (o *NamingConfigResource) UnsetNumberStyle()`

UnsetNumberStyle ensures that no value is present for NumberStyle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


