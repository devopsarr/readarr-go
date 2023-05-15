# ManualImportUpdateResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**AuthorId** | Pointer to **NullableInt32** |  | [optional] 
**BookId** | Pointer to **NullableInt32** |  | [optional] 
**ForeignEditionId** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**AdditionalFile** | Pointer to **bool** |  | [optional] 
**ReplaceExistingFiles** | Pointer to **bool** |  | [optional] 
**DisableReleaseSwitching** | Pointer to **bool** |  | [optional] 
**Rejections** | Pointer to [**[]Rejection**](Rejection.md) |  | [optional] 

## Methods

### NewManualImportUpdateResource

`func NewManualImportUpdateResource() *ManualImportUpdateResource`

NewManualImportUpdateResource instantiates a new ManualImportUpdateResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualImportUpdateResourceWithDefaults

`func NewManualImportUpdateResourceWithDefaults() *ManualImportUpdateResource`

NewManualImportUpdateResourceWithDefaults instantiates a new ManualImportUpdateResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualImportUpdateResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualImportUpdateResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualImportUpdateResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManualImportUpdateResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ManualImportUpdateResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManualImportUpdateResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManualImportUpdateResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManualImportUpdateResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ManualImportUpdateResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ManualImportUpdateResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetName

`func (o *ManualImportUpdateResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManualImportUpdateResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManualImportUpdateResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManualImportUpdateResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ManualImportUpdateResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ManualImportUpdateResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAuthorId

`func (o *ManualImportUpdateResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *ManualImportUpdateResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *ManualImportUpdateResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *ManualImportUpdateResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### SetAuthorIdNil

`func (o *ManualImportUpdateResource) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *ManualImportUpdateResource) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
### GetBookId

`func (o *ManualImportUpdateResource) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *ManualImportUpdateResource) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *ManualImportUpdateResource) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *ManualImportUpdateResource) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### SetBookIdNil

`func (o *ManualImportUpdateResource) SetBookIdNil(b bool)`

 SetBookIdNil sets the value for BookId to be an explicit nil

### UnsetBookId
`func (o *ManualImportUpdateResource) UnsetBookId()`

UnsetBookId ensures that no value is present for BookId, not even an explicit nil
### GetForeignEditionId

`func (o *ManualImportUpdateResource) GetForeignEditionId() string`

GetForeignEditionId returns the ForeignEditionId field if non-nil, zero value otherwise.

### GetForeignEditionIdOk

`func (o *ManualImportUpdateResource) GetForeignEditionIdOk() (*string, bool)`

GetForeignEditionIdOk returns a tuple with the ForeignEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignEditionId

`func (o *ManualImportUpdateResource) SetForeignEditionId(v string)`

SetForeignEditionId sets ForeignEditionId field to given value.

### HasForeignEditionId

`func (o *ManualImportUpdateResource) HasForeignEditionId() bool`

HasForeignEditionId returns a boolean if a field has been set.

### SetForeignEditionIdNil

`func (o *ManualImportUpdateResource) SetForeignEditionIdNil(b bool)`

 SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil

### UnsetForeignEditionId
`func (o *ManualImportUpdateResource) UnsetForeignEditionId()`

UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
### GetQuality

`func (o *ManualImportUpdateResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ManualImportUpdateResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ManualImportUpdateResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ManualImportUpdateResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetReleaseGroup

`func (o *ManualImportUpdateResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ManualImportUpdateResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ManualImportUpdateResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ManualImportUpdateResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ManualImportUpdateResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ManualImportUpdateResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetDownloadId

`func (o *ManualImportUpdateResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *ManualImportUpdateResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *ManualImportUpdateResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *ManualImportUpdateResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *ManualImportUpdateResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *ManualImportUpdateResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetAdditionalFile

`func (o *ManualImportUpdateResource) GetAdditionalFile() bool`

GetAdditionalFile returns the AdditionalFile field if non-nil, zero value otherwise.

### GetAdditionalFileOk

`func (o *ManualImportUpdateResource) GetAdditionalFileOk() (*bool, bool)`

GetAdditionalFileOk returns a tuple with the AdditionalFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFile

`func (o *ManualImportUpdateResource) SetAdditionalFile(v bool)`

SetAdditionalFile sets AdditionalFile field to given value.

### HasAdditionalFile

`func (o *ManualImportUpdateResource) HasAdditionalFile() bool`

HasAdditionalFile returns a boolean if a field has been set.

### GetReplaceExistingFiles

`func (o *ManualImportUpdateResource) GetReplaceExistingFiles() bool`

GetReplaceExistingFiles returns the ReplaceExistingFiles field if non-nil, zero value otherwise.

### GetReplaceExistingFilesOk

`func (o *ManualImportUpdateResource) GetReplaceExistingFilesOk() (*bool, bool)`

GetReplaceExistingFilesOk returns a tuple with the ReplaceExistingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceExistingFiles

`func (o *ManualImportUpdateResource) SetReplaceExistingFiles(v bool)`

SetReplaceExistingFiles sets ReplaceExistingFiles field to given value.

### HasReplaceExistingFiles

`func (o *ManualImportUpdateResource) HasReplaceExistingFiles() bool`

HasReplaceExistingFiles returns a boolean if a field has been set.

### GetDisableReleaseSwitching

`func (o *ManualImportUpdateResource) GetDisableReleaseSwitching() bool`

GetDisableReleaseSwitching returns the DisableReleaseSwitching field if non-nil, zero value otherwise.

### GetDisableReleaseSwitchingOk

`func (o *ManualImportUpdateResource) GetDisableReleaseSwitchingOk() (*bool, bool)`

GetDisableReleaseSwitchingOk returns a tuple with the DisableReleaseSwitching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReleaseSwitching

`func (o *ManualImportUpdateResource) SetDisableReleaseSwitching(v bool)`

SetDisableReleaseSwitching sets DisableReleaseSwitching field to given value.

### HasDisableReleaseSwitching

`func (o *ManualImportUpdateResource) HasDisableReleaseSwitching() bool`

HasDisableReleaseSwitching returns a boolean if a field has been set.

### GetRejections

`func (o *ManualImportUpdateResource) GetRejections() []Rejection`

GetRejections returns the Rejections field if non-nil, zero value otherwise.

### GetRejectionsOk

`func (o *ManualImportUpdateResource) GetRejectionsOk() (*[]Rejection, bool)`

GetRejectionsOk returns a tuple with the Rejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejections

`func (o *ManualImportUpdateResource) SetRejections(v []Rejection)`

SetRejections sets Rejections field to given value.

### HasRejections

`func (o *ManualImportUpdateResource) HasRejections() bool`

HasRejections returns a boolean if a field has been set.

### SetRejectionsNil

`func (o *ManualImportUpdateResource) SetRejectionsNil(b bool)`

 SetRejectionsNil sets the value for Rejections to be an explicit nil

### UnsetRejections
`func (o *ManualImportUpdateResource) UnsetRejections()`

UnsetRejections ensures that no value is present for Rejections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


