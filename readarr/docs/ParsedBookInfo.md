# ParsedBookInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookTitle** | Pointer to **NullableString** |  | [optional] 
**AuthorName** | Pointer to **NullableString** |  | [optional] 
**AuthorTitleInfo** | Pointer to [**AuthorTitleInfo**](AuthorTitleInfo.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**Discography** | Pointer to **bool** |  | [optional] 
**DiscographyStart** | Pointer to **int32** |  | [optional] 
**DiscographyEnd** | Pointer to **int32** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 
**ReleaseVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewParsedBookInfo

`func NewParsedBookInfo() *ParsedBookInfo`

NewParsedBookInfo instantiates a new ParsedBookInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsedBookInfoWithDefaults

`func NewParsedBookInfoWithDefaults() *ParsedBookInfo`

NewParsedBookInfoWithDefaults instantiates a new ParsedBookInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookTitle

`func (o *ParsedBookInfo) GetBookTitle() string`

GetBookTitle returns the BookTitle field if non-nil, zero value otherwise.

### GetBookTitleOk

`func (o *ParsedBookInfo) GetBookTitleOk() (*string, bool)`

GetBookTitleOk returns a tuple with the BookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTitle

`func (o *ParsedBookInfo) SetBookTitle(v string)`

SetBookTitle sets BookTitle field to given value.

### HasBookTitle

`func (o *ParsedBookInfo) HasBookTitle() bool`

HasBookTitle returns a boolean if a field has been set.

### SetBookTitleNil

`func (o *ParsedBookInfo) SetBookTitleNil(b bool)`

 SetBookTitleNil sets the value for BookTitle to be an explicit nil

### UnsetBookTitle
`func (o *ParsedBookInfo) UnsetBookTitle()`

UnsetBookTitle ensures that no value is present for BookTitle, not even an explicit nil
### GetAuthorName

`func (o *ParsedBookInfo) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *ParsedBookInfo) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *ParsedBookInfo) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *ParsedBookInfo) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *ParsedBookInfo) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *ParsedBookInfo) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetAuthorTitleInfo

`func (o *ParsedBookInfo) GetAuthorTitleInfo() AuthorTitleInfo`

GetAuthorTitleInfo returns the AuthorTitleInfo field if non-nil, zero value otherwise.

### GetAuthorTitleInfoOk

`func (o *ParsedBookInfo) GetAuthorTitleInfoOk() (*AuthorTitleInfo, bool)`

GetAuthorTitleInfoOk returns a tuple with the AuthorTitleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTitleInfo

`func (o *ParsedBookInfo) SetAuthorTitleInfo(v AuthorTitleInfo)`

SetAuthorTitleInfo sets AuthorTitleInfo field to given value.

### HasAuthorTitleInfo

`func (o *ParsedBookInfo) HasAuthorTitleInfo() bool`

HasAuthorTitleInfo returns a boolean if a field has been set.

### GetQuality

`func (o *ParsedBookInfo) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ParsedBookInfo) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ParsedBookInfo) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ParsedBookInfo) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ParsedBookInfo) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ParsedBookInfo) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ParsedBookInfo) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ParsedBookInfo) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *ParsedBookInfo) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *ParsedBookInfo) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetDiscography

`func (o *ParsedBookInfo) GetDiscography() bool`

GetDiscography returns the Discography field if non-nil, zero value otherwise.

### GetDiscographyOk

`func (o *ParsedBookInfo) GetDiscographyOk() (*bool, bool)`

GetDiscographyOk returns a tuple with the Discography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscography

`func (o *ParsedBookInfo) SetDiscography(v bool)`

SetDiscography sets Discography field to given value.

### HasDiscography

`func (o *ParsedBookInfo) HasDiscography() bool`

HasDiscography returns a boolean if a field has been set.

### GetDiscographyStart

`func (o *ParsedBookInfo) GetDiscographyStart() int32`

GetDiscographyStart returns the DiscographyStart field if non-nil, zero value otherwise.

### GetDiscographyStartOk

`func (o *ParsedBookInfo) GetDiscographyStartOk() (*int32, bool)`

GetDiscographyStartOk returns a tuple with the DiscographyStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscographyStart

`func (o *ParsedBookInfo) SetDiscographyStart(v int32)`

SetDiscographyStart sets DiscographyStart field to given value.

### HasDiscographyStart

`func (o *ParsedBookInfo) HasDiscographyStart() bool`

HasDiscographyStart returns a boolean if a field has been set.

### GetDiscographyEnd

`func (o *ParsedBookInfo) GetDiscographyEnd() int32`

GetDiscographyEnd returns the DiscographyEnd field if non-nil, zero value otherwise.

### GetDiscographyEndOk

`func (o *ParsedBookInfo) GetDiscographyEndOk() (*int32, bool)`

GetDiscographyEndOk returns a tuple with the DiscographyEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscographyEnd

`func (o *ParsedBookInfo) SetDiscographyEnd(v int32)`

SetDiscographyEnd sets DiscographyEnd field to given value.

### HasDiscographyEnd

`func (o *ParsedBookInfo) HasDiscographyEnd() bool`

HasDiscographyEnd returns a boolean if a field has been set.

### GetReleaseGroup

`func (o *ParsedBookInfo) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ParsedBookInfo) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ParsedBookInfo) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ParsedBookInfo) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ParsedBookInfo) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ParsedBookInfo) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetReleaseHash

`func (o *ParsedBookInfo) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ParsedBookInfo) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ParsedBookInfo) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ParsedBookInfo) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ParsedBookInfo) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ParsedBookInfo) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
### GetReleaseVersion

`func (o *ParsedBookInfo) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *ParsedBookInfo) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *ParsedBookInfo) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *ParsedBookInfo) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.

### SetReleaseVersionNil

`func (o *ParsedBookInfo) SetReleaseVersionNil(b bool)`

 SetReleaseVersionNil sets the value for ReleaseVersion to be an explicit nil

### UnsetReleaseVersion
`func (o *ParsedBookInfo) UnsetReleaseVersion()`

UnsetReleaseVersion ensures that no value is present for ReleaseVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


