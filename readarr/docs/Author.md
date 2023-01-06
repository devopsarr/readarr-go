# Author

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorMetadataId** | Pointer to **int32** |  | [optional] 
**CleanName** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**LastInfoSync** | Pointer to **NullableTime** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**AddOptions** | Pointer to [**AddAuthorOptions**](AddAuthorOptions.md) |  | [optional] 
**Metadata** | Pointer to [**AuthorMetadataLazyLoaded**](AuthorMetadataLazyLoaded.md) |  | [optional] 
**QualityProfile** | Pointer to [**QualityProfileLazyLoaded**](QualityProfileLazyLoaded.md) |  | [optional] 
**MetadataProfile** | Pointer to [**MetadataProfileLazyLoaded**](MetadataProfileLazyLoaded.md) |  | [optional] 
**Books** | Pointer to [**BookListLazyLoaded**](BookListLazyLoaded.md) |  | [optional] 
**Series** | Pointer to [**SeriesListLazyLoaded**](SeriesListLazyLoaded.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ForeignAuthorId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuthor

`func NewAuthor() *Author`

NewAuthor instantiates a new Author object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorWithDefaults

`func NewAuthorWithDefaults() *Author`

NewAuthorWithDefaults instantiates a new Author object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Author) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Author) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Author) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Author) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorMetadataId

`func (o *Author) GetAuthorMetadataId() int32`

GetAuthorMetadataId returns the AuthorMetadataId field if non-nil, zero value otherwise.

### GetAuthorMetadataIdOk

`func (o *Author) GetAuthorMetadataIdOk() (*int32, bool)`

GetAuthorMetadataIdOk returns a tuple with the AuthorMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMetadataId

`func (o *Author) SetAuthorMetadataId(v int32)`

SetAuthorMetadataId sets AuthorMetadataId field to given value.

### HasAuthorMetadataId

`func (o *Author) HasAuthorMetadataId() bool`

HasAuthorMetadataId returns a boolean if a field has been set.

### GetCleanName

`func (o *Author) GetCleanName() string`

GetCleanName returns the CleanName field if non-nil, zero value otherwise.

### GetCleanNameOk

`func (o *Author) GetCleanNameOk() (*string, bool)`

GetCleanNameOk returns a tuple with the CleanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanName

`func (o *Author) SetCleanName(v string)`

SetCleanName sets CleanName field to given value.

### HasCleanName

`func (o *Author) HasCleanName() bool`

HasCleanName returns a boolean if a field has been set.

### SetCleanNameNil

`func (o *Author) SetCleanNameNil(b bool)`

 SetCleanNameNil sets the value for CleanName to be an explicit nil

### UnsetCleanName
`func (o *Author) UnsetCleanName()`

UnsetCleanName ensures that no value is present for CleanName, not even an explicit nil
### GetMonitored

`func (o *Author) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *Author) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *Author) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *Author) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetLastInfoSync

`func (o *Author) GetLastInfoSync() time.Time`

GetLastInfoSync returns the LastInfoSync field if non-nil, zero value otherwise.

### GetLastInfoSyncOk

`func (o *Author) GetLastInfoSyncOk() (*time.Time, bool)`

GetLastInfoSyncOk returns a tuple with the LastInfoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInfoSync

`func (o *Author) SetLastInfoSync(v time.Time)`

SetLastInfoSync sets LastInfoSync field to given value.

### HasLastInfoSync

`func (o *Author) HasLastInfoSync() bool`

HasLastInfoSync returns a boolean if a field has been set.

### SetLastInfoSyncNil

`func (o *Author) SetLastInfoSyncNil(b bool)`

 SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil

### UnsetLastInfoSync
`func (o *Author) UnsetLastInfoSync()`

UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
### GetPath

`func (o *Author) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Author) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Author) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Author) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Author) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Author) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRootFolderPath

`func (o *Author) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *Author) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *Author) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *Author) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *Author) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *Author) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetAdded

`func (o *Author) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *Author) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *Author) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *Author) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetQualityProfileId

`func (o *Author) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *Author) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *Author) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *Author) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *Author) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *Author) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *Author) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *Author) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetTags

`func (o *Author) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Author) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Author) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Author) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Author) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Author) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAddOptions

`func (o *Author) GetAddOptions() AddAuthorOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *Author) GetAddOptionsOk() (*AddAuthorOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *Author) SetAddOptions(v AddAuthorOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *Author) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *Author) GetMetadata() AuthorMetadataLazyLoaded`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Author) GetMetadataOk() (*AuthorMetadataLazyLoaded, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Author) SetMetadata(v AuthorMetadataLazyLoaded)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Author) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQualityProfile

`func (o *Author) GetQualityProfile() QualityProfileLazyLoaded`

GetQualityProfile returns the QualityProfile field if non-nil, zero value otherwise.

### GetQualityProfileOk

`func (o *Author) GetQualityProfileOk() (*QualityProfileLazyLoaded, bool)`

GetQualityProfileOk returns a tuple with the QualityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfile

`func (o *Author) SetQualityProfile(v QualityProfileLazyLoaded)`

SetQualityProfile sets QualityProfile field to given value.

### HasQualityProfile

`func (o *Author) HasQualityProfile() bool`

HasQualityProfile returns a boolean if a field has been set.

### GetMetadataProfile

`func (o *Author) GetMetadataProfile() MetadataProfileLazyLoaded`

GetMetadataProfile returns the MetadataProfile field if non-nil, zero value otherwise.

### GetMetadataProfileOk

`func (o *Author) GetMetadataProfileOk() (*MetadataProfileLazyLoaded, bool)`

GetMetadataProfileOk returns a tuple with the MetadataProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfile

`func (o *Author) SetMetadataProfile(v MetadataProfileLazyLoaded)`

SetMetadataProfile sets MetadataProfile field to given value.

### HasMetadataProfile

`func (o *Author) HasMetadataProfile() bool`

HasMetadataProfile returns a boolean if a field has been set.

### GetBooks

`func (o *Author) GetBooks() BookListLazyLoaded`

GetBooks returns the Books field if non-nil, zero value otherwise.

### GetBooksOk

`func (o *Author) GetBooksOk() (*BookListLazyLoaded, bool)`

GetBooksOk returns a tuple with the Books field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooks

`func (o *Author) SetBooks(v BookListLazyLoaded)`

SetBooks sets Books field to given value.

### HasBooks

`func (o *Author) HasBooks() bool`

HasBooks returns a boolean if a field has been set.

### GetSeries

`func (o *Author) GetSeries() SeriesListLazyLoaded`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Author) GetSeriesOk() (*SeriesListLazyLoaded, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Author) SetSeries(v SeriesListLazyLoaded)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Author) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetName

`func (o *Author) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Author) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Author) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Author) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Author) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Author) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetForeignAuthorId

`func (o *Author) GetForeignAuthorId() string`

GetForeignAuthorId returns the ForeignAuthorId field if non-nil, zero value otherwise.

### GetForeignAuthorIdOk

`func (o *Author) GetForeignAuthorIdOk() (*string, bool)`

GetForeignAuthorIdOk returns a tuple with the ForeignAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAuthorId

`func (o *Author) SetForeignAuthorId(v string)`

SetForeignAuthorId sets ForeignAuthorId field to given value.

### HasForeignAuthorId

`func (o *Author) HasForeignAuthorId() bool`

HasForeignAuthorId returns a boolean if a field has been set.

### SetForeignAuthorIdNil

`func (o *Author) SetForeignAuthorIdNil(b bool)`

 SetForeignAuthorIdNil sets the value for ForeignAuthorId to be an explicit nil

### UnsetForeignAuthorId
`func (o *Author) UnsetForeignAuthorId()`

UnsetForeignAuthorId ensures that no value is present for ForeignAuthorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


