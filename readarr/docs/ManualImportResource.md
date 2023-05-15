# ManualImportResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Author** | Pointer to [**AuthorResource**](AuthorResource.md) |  | [optional] 
**Book** | Pointer to [**BookResource**](BookResource.md) |  | [optional] 
**ForeignEditionId** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**QualityWeight** | Pointer to **int32** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**Rejections** | Pointer to [**[]Rejection**](Rejection.md) |  | [optional] 
**AudioTags** | Pointer to [**ParsedTrackInfo**](ParsedTrackInfo.md) |  | [optional] 
**AdditionalFile** | Pointer to **bool** |  | [optional] 
**ReplaceExistingFiles** | Pointer to **bool** |  | [optional] 
**DisableReleaseSwitching** | Pointer to **bool** |  | [optional] 

## Methods

### NewManualImportResource

`func NewManualImportResource() *ManualImportResource`

NewManualImportResource instantiates a new ManualImportResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualImportResourceWithDefaults

`func NewManualImportResourceWithDefaults() *ManualImportResource`

NewManualImportResourceWithDefaults instantiates a new ManualImportResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualImportResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualImportResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualImportResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManualImportResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ManualImportResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManualImportResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManualImportResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManualImportResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ManualImportResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ManualImportResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetName

`func (o *ManualImportResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManualImportResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManualImportResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManualImportResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ManualImportResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ManualImportResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *ManualImportResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ManualImportResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ManualImportResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ManualImportResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetAuthor

`func (o *ManualImportResource) GetAuthor() AuthorResource`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ManualImportResource) GetAuthorOk() (*AuthorResource, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ManualImportResource) SetAuthor(v AuthorResource)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ManualImportResource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBook

`func (o *ManualImportResource) GetBook() BookResource`

GetBook returns the Book field if non-nil, zero value otherwise.

### GetBookOk

`func (o *ManualImportResource) GetBookOk() (*BookResource, bool)`

GetBookOk returns a tuple with the Book field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBook

`func (o *ManualImportResource) SetBook(v BookResource)`

SetBook sets Book field to given value.

### HasBook

`func (o *ManualImportResource) HasBook() bool`

HasBook returns a boolean if a field has been set.

### GetForeignEditionId

`func (o *ManualImportResource) GetForeignEditionId() string`

GetForeignEditionId returns the ForeignEditionId field if non-nil, zero value otherwise.

### GetForeignEditionIdOk

`func (o *ManualImportResource) GetForeignEditionIdOk() (*string, bool)`

GetForeignEditionIdOk returns a tuple with the ForeignEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignEditionId

`func (o *ManualImportResource) SetForeignEditionId(v string)`

SetForeignEditionId sets ForeignEditionId field to given value.

### HasForeignEditionId

`func (o *ManualImportResource) HasForeignEditionId() bool`

HasForeignEditionId returns a boolean if a field has been set.

### SetForeignEditionIdNil

`func (o *ManualImportResource) SetForeignEditionIdNil(b bool)`

 SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil

### UnsetForeignEditionId
`func (o *ManualImportResource) UnsetForeignEditionId()`

UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
### GetQuality

`func (o *ManualImportResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ManualImportResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ManualImportResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ManualImportResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetReleaseGroup

`func (o *ManualImportResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ManualImportResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ManualImportResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ManualImportResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ManualImportResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ManualImportResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetQualityWeight

`func (o *ManualImportResource) GetQualityWeight() int32`

GetQualityWeight returns the QualityWeight field if non-nil, zero value otherwise.

### GetQualityWeightOk

`func (o *ManualImportResource) GetQualityWeightOk() (*int32, bool)`

GetQualityWeightOk returns a tuple with the QualityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityWeight

`func (o *ManualImportResource) SetQualityWeight(v int32)`

SetQualityWeight sets QualityWeight field to given value.

### HasQualityWeight

`func (o *ManualImportResource) HasQualityWeight() bool`

HasQualityWeight returns a boolean if a field has been set.

### GetDownloadId

`func (o *ManualImportResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *ManualImportResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *ManualImportResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *ManualImportResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *ManualImportResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *ManualImportResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetRejections

`func (o *ManualImportResource) GetRejections() []Rejection`

GetRejections returns the Rejections field if non-nil, zero value otherwise.

### GetRejectionsOk

`func (o *ManualImportResource) GetRejectionsOk() (*[]Rejection, bool)`

GetRejectionsOk returns a tuple with the Rejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejections

`func (o *ManualImportResource) SetRejections(v []Rejection)`

SetRejections sets Rejections field to given value.

### HasRejections

`func (o *ManualImportResource) HasRejections() bool`

HasRejections returns a boolean if a field has been set.

### SetRejectionsNil

`func (o *ManualImportResource) SetRejectionsNil(b bool)`

 SetRejectionsNil sets the value for Rejections to be an explicit nil

### UnsetRejections
`func (o *ManualImportResource) UnsetRejections()`

UnsetRejections ensures that no value is present for Rejections, not even an explicit nil
### GetAudioTags

`func (o *ManualImportResource) GetAudioTags() ParsedTrackInfo`

GetAudioTags returns the AudioTags field if non-nil, zero value otherwise.

### GetAudioTagsOk

`func (o *ManualImportResource) GetAudioTagsOk() (*ParsedTrackInfo, bool)`

GetAudioTagsOk returns a tuple with the AudioTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioTags

`func (o *ManualImportResource) SetAudioTags(v ParsedTrackInfo)`

SetAudioTags sets AudioTags field to given value.

### HasAudioTags

`func (o *ManualImportResource) HasAudioTags() bool`

HasAudioTags returns a boolean if a field has been set.

### GetAdditionalFile

`func (o *ManualImportResource) GetAdditionalFile() bool`

GetAdditionalFile returns the AdditionalFile field if non-nil, zero value otherwise.

### GetAdditionalFileOk

`func (o *ManualImportResource) GetAdditionalFileOk() (*bool, bool)`

GetAdditionalFileOk returns a tuple with the AdditionalFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFile

`func (o *ManualImportResource) SetAdditionalFile(v bool)`

SetAdditionalFile sets AdditionalFile field to given value.

### HasAdditionalFile

`func (o *ManualImportResource) HasAdditionalFile() bool`

HasAdditionalFile returns a boolean if a field has been set.

### GetReplaceExistingFiles

`func (o *ManualImportResource) GetReplaceExistingFiles() bool`

GetReplaceExistingFiles returns the ReplaceExistingFiles field if non-nil, zero value otherwise.

### GetReplaceExistingFilesOk

`func (o *ManualImportResource) GetReplaceExistingFilesOk() (*bool, bool)`

GetReplaceExistingFilesOk returns a tuple with the ReplaceExistingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceExistingFiles

`func (o *ManualImportResource) SetReplaceExistingFiles(v bool)`

SetReplaceExistingFiles sets ReplaceExistingFiles field to given value.

### HasReplaceExistingFiles

`func (o *ManualImportResource) HasReplaceExistingFiles() bool`

HasReplaceExistingFiles returns a boolean if a field has been set.

### GetDisableReleaseSwitching

`func (o *ManualImportResource) GetDisableReleaseSwitching() bool`

GetDisableReleaseSwitching returns the DisableReleaseSwitching field if non-nil, zero value otherwise.

### GetDisableReleaseSwitchingOk

`func (o *ManualImportResource) GetDisableReleaseSwitchingOk() (*bool, bool)`

GetDisableReleaseSwitchingOk returns a tuple with the DisableReleaseSwitching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReleaseSwitching

`func (o *ManualImportResource) SetDisableReleaseSwitching(v bool)`

SetDisableReleaseSwitching sets DisableReleaseSwitching field to given value.

### HasDisableReleaseSwitching

`func (o *ManualImportResource) HasDisableReleaseSwitching() bool`

HasDisableReleaseSwitching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


