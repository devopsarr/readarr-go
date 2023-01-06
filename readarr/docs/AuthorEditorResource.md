# AuthorEditorResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorIds** | Pointer to **[]int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**QualityProfileId** | Pointer to **NullableInt32** |  | [optional] 
**MetadataProfileId** | Pointer to **NullableInt32** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**MoveFiles** | Pointer to **bool** |  | [optional] 
**DeleteFiles** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthorEditorResource

`func NewAuthorEditorResource() *AuthorEditorResource`

NewAuthorEditorResource instantiates a new AuthorEditorResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorEditorResourceWithDefaults

`func NewAuthorEditorResourceWithDefaults() *AuthorEditorResource`

NewAuthorEditorResourceWithDefaults instantiates a new AuthorEditorResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorIds

`func (o *AuthorEditorResource) GetAuthorIds() []int32`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *AuthorEditorResource) GetAuthorIdsOk() (*[]int32, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *AuthorEditorResource) SetAuthorIds(v []int32)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *AuthorEditorResource) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### SetAuthorIdsNil

`func (o *AuthorEditorResource) SetAuthorIdsNil(b bool)`

 SetAuthorIdsNil sets the value for AuthorIds to be an explicit nil

### UnsetAuthorIds
`func (o *AuthorEditorResource) UnsetAuthorIds()`

UnsetAuthorIds ensures that no value is present for AuthorIds, not even an explicit nil
### GetMonitored

`func (o *AuthorEditorResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AuthorEditorResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AuthorEditorResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AuthorEditorResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *AuthorEditorResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *AuthorEditorResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetQualityProfileId

`func (o *AuthorEditorResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *AuthorEditorResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *AuthorEditorResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *AuthorEditorResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### SetQualityProfileIdNil

`func (o *AuthorEditorResource) SetQualityProfileIdNil(b bool)`

 SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil

### UnsetQualityProfileId
`func (o *AuthorEditorResource) UnsetQualityProfileId()`

UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
### GetMetadataProfileId

`func (o *AuthorEditorResource) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *AuthorEditorResource) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *AuthorEditorResource) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *AuthorEditorResource) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### SetMetadataProfileIdNil

`func (o *AuthorEditorResource) SetMetadataProfileIdNil(b bool)`

 SetMetadataProfileIdNil sets the value for MetadataProfileId to be an explicit nil

### UnsetMetadataProfileId
`func (o *AuthorEditorResource) UnsetMetadataProfileId()`

UnsetMetadataProfileId ensures that no value is present for MetadataProfileId, not even an explicit nil
### GetRootFolderPath

`func (o *AuthorEditorResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *AuthorEditorResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *AuthorEditorResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *AuthorEditorResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *AuthorEditorResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *AuthorEditorResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetTags

`func (o *AuthorEditorResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AuthorEditorResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AuthorEditorResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AuthorEditorResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AuthorEditorResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AuthorEditorResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *AuthorEditorResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *AuthorEditorResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *AuthorEditorResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *AuthorEditorResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetMoveFiles

`func (o *AuthorEditorResource) GetMoveFiles() bool`

GetMoveFiles returns the MoveFiles field if non-nil, zero value otherwise.

### GetMoveFilesOk

`func (o *AuthorEditorResource) GetMoveFilesOk() (*bool, bool)`

GetMoveFilesOk returns a tuple with the MoveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveFiles

`func (o *AuthorEditorResource) SetMoveFiles(v bool)`

SetMoveFiles sets MoveFiles field to given value.

### HasMoveFiles

`func (o *AuthorEditorResource) HasMoveFiles() bool`

HasMoveFiles returns a boolean if a field has been set.

### GetDeleteFiles

`func (o *AuthorEditorResource) GetDeleteFiles() bool`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *AuthorEditorResource) GetDeleteFilesOk() (*bool, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *AuthorEditorResource) SetDeleteFiles(v bool)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *AuthorEditorResource) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


