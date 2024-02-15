# BookFileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **int32** |  | [optional] 
**BookId** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**DateAdded** | Pointer to **time.Time** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**QualityWeight** | Pointer to **int32** |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoResource**](MediaInfoResource.md) |  | [optional] 
**QualityCutoffNotMet** | Pointer to **bool** |  | [optional] 
**AudioTags** | Pointer to [**ParsedTrackInfo**](ParsedTrackInfo.md) |  | [optional] 

## Methods

### NewBookFileResource

`func NewBookFileResource() *BookFileResource`

NewBookFileResource instantiates a new BookFileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookFileResourceWithDefaults

`func NewBookFileResourceWithDefaults() *BookFileResource`

NewBookFileResourceWithDefaults instantiates a new BookFileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookFileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookFileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookFileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BookFileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorId

`func (o *BookFileResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *BookFileResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *BookFileResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *BookFileResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetBookId

`func (o *BookFileResource) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *BookFileResource) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *BookFileResource) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *BookFileResource) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### GetPath

`func (o *BookFileResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BookFileResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BookFileResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BookFileResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BookFileResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BookFileResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSize

`func (o *BookFileResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BookFileResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BookFileResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BookFileResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDateAdded

`func (o *BookFileResource) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *BookFileResource) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *BookFileResource) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *BookFileResource) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetQuality

`func (o *BookFileResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BookFileResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BookFileResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BookFileResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetQualityWeight

`func (o *BookFileResource) GetQualityWeight() int32`

GetQualityWeight returns the QualityWeight field if non-nil, zero value otherwise.

### GetQualityWeightOk

`func (o *BookFileResource) GetQualityWeightOk() (*int32, bool)`

GetQualityWeightOk returns a tuple with the QualityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityWeight

`func (o *BookFileResource) SetQualityWeight(v int32)`

SetQualityWeight sets QualityWeight field to given value.

### HasQualityWeight

`func (o *BookFileResource) HasQualityWeight() bool`

HasQualityWeight returns a boolean if a field has been set.

### GetMediaInfo

`func (o *BookFileResource) GetMediaInfo() MediaInfoResource`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *BookFileResource) GetMediaInfoOk() (*MediaInfoResource, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *BookFileResource) SetMediaInfo(v MediaInfoResource)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *BookFileResource) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetQualityCutoffNotMet

`func (o *BookFileResource) GetQualityCutoffNotMet() bool`

GetQualityCutoffNotMet returns the QualityCutoffNotMet field if non-nil, zero value otherwise.

### GetQualityCutoffNotMetOk

`func (o *BookFileResource) GetQualityCutoffNotMetOk() (*bool, bool)`

GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityCutoffNotMet

`func (o *BookFileResource) SetQualityCutoffNotMet(v bool)`

SetQualityCutoffNotMet sets QualityCutoffNotMet field to given value.

### HasQualityCutoffNotMet

`func (o *BookFileResource) HasQualityCutoffNotMet() bool`

HasQualityCutoffNotMet returns a boolean if a field has been set.

### GetAudioTags

`func (o *BookFileResource) GetAudioTags() ParsedTrackInfo`

GetAudioTags returns the AudioTags field if non-nil, zero value otherwise.

### GetAudioTagsOk

`func (o *BookFileResource) GetAudioTagsOk() (*ParsedTrackInfo, bool)`

GetAudioTagsOk returns a tuple with the AudioTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioTags

`func (o *BookFileResource) SetAudioTags(v ParsedTrackInfo)`

SetAudioTags sets AudioTags field to given value.

### HasAudioTags

`func (o *BookFileResource) HasAudioTags() bool`

HasAudioTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


