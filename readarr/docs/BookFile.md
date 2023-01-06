# BookFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**DateAdded** | Pointer to **time.Time** |  | [optional] 
**SceneName** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoModel**](MediaInfoModel.md) |  | [optional] 
**EditionId** | Pointer to **int32** |  | [optional] 
**CalibreId** | Pointer to **int32** |  | [optional] 
**Part** | Pointer to **int32** |  | [optional] 
**Author** | Pointer to [**AuthorLazyLoaded**](AuthorLazyLoaded.md) |  | [optional] 
**Edition** | Pointer to [**EditionLazyLoaded**](EditionLazyLoaded.md) |  | [optional] 
**PartCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewBookFile

`func NewBookFile() *BookFile`

NewBookFile instantiates a new BookFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookFileWithDefaults

`func NewBookFileWithDefaults() *BookFile`

NewBookFileWithDefaults instantiates a new BookFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookFile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BookFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *BookFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BookFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BookFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BookFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BookFile) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BookFile) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSize

`func (o *BookFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BookFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BookFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BookFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetModified

`func (o *BookFile) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BookFile) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BookFile) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *BookFile) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDateAdded

`func (o *BookFile) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *BookFile) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *BookFile) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *BookFile) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetSceneName

`func (o *BookFile) GetSceneName() string`

GetSceneName returns the SceneName field if non-nil, zero value otherwise.

### GetSceneNameOk

`func (o *BookFile) GetSceneNameOk() (*string, bool)`

GetSceneNameOk returns a tuple with the SceneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneName

`func (o *BookFile) SetSceneName(v string)`

SetSceneName sets SceneName field to given value.

### HasSceneName

`func (o *BookFile) HasSceneName() bool`

HasSceneName returns a boolean if a field has been set.

### SetSceneNameNil

`func (o *BookFile) SetSceneNameNil(b bool)`

 SetSceneNameNil sets the value for SceneName to be an explicit nil

### UnsetSceneName
`func (o *BookFile) UnsetSceneName()`

UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
### GetReleaseGroup

`func (o *BookFile) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *BookFile) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *BookFile) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *BookFile) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *BookFile) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *BookFile) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetQuality

`func (o *BookFile) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BookFile) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BookFile) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BookFile) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetMediaInfo

`func (o *BookFile) GetMediaInfo() MediaInfoModel`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *BookFile) GetMediaInfoOk() (*MediaInfoModel, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *BookFile) SetMediaInfo(v MediaInfoModel)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *BookFile) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetEditionId

`func (o *BookFile) GetEditionId() int32`

GetEditionId returns the EditionId field if non-nil, zero value otherwise.

### GetEditionIdOk

`func (o *BookFile) GetEditionIdOk() (*int32, bool)`

GetEditionIdOk returns a tuple with the EditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionId

`func (o *BookFile) SetEditionId(v int32)`

SetEditionId sets EditionId field to given value.

### HasEditionId

`func (o *BookFile) HasEditionId() bool`

HasEditionId returns a boolean if a field has been set.

### GetCalibreId

`func (o *BookFile) GetCalibreId() int32`

GetCalibreId returns the CalibreId field if non-nil, zero value otherwise.

### GetCalibreIdOk

`func (o *BookFile) GetCalibreIdOk() (*int32, bool)`

GetCalibreIdOk returns a tuple with the CalibreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalibreId

`func (o *BookFile) SetCalibreId(v int32)`

SetCalibreId sets CalibreId field to given value.

### HasCalibreId

`func (o *BookFile) HasCalibreId() bool`

HasCalibreId returns a boolean if a field has been set.

### GetPart

`func (o *BookFile) GetPart() int32`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *BookFile) GetPartOk() (*int32, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *BookFile) SetPart(v int32)`

SetPart sets Part field to given value.

### HasPart

`func (o *BookFile) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetAuthor

`func (o *BookFile) GetAuthor() AuthorLazyLoaded`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BookFile) GetAuthorOk() (*AuthorLazyLoaded, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BookFile) SetAuthor(v AuthorLazyLoaded)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *BookFile) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetEdition

`func (o *BookFile) GetEdition() EditionLazyLoaded`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *BookFile) GetEditionOk() (*EditionLazyLoaded, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *BookFile) SetEdition(v EditionLazyLoaded)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *BookFile) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetPartCount

`func (o *BookFile) GetPartCount() int32`

GetPartCount returns the PartCount field if non-nil, zero value otherwise.

### GetPartCountOk

`func (o *BookFile) GetPartCountOk() (*int32, bool)`

GetPartCountOk returns a tuple with the PartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartCount

`func (o *BookFile) SetPartCount(v int32)`

SetPartCount sets PartCount field to given value.

### HasPartCount

`func (o *BookFile) HasPartCount() bool`

HasPartCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


