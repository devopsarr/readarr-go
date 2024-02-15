# AuthorResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorMetadataId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**AuthorStatusType**](AuthorStatusType.md) |  | [optional] 
**Ended** | Pointer to **bool** |  | [optional] [readonly] 
**AuthorName** | Pointer to **NullableString** |  | [optional] 
**AuthorNameLastFirst** | Pointer to **NullableString** |  | [optional] 
**ForeignAuthorId** | Pointer to **NullableString** |  | [optional] 
**TitleSlug** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**NextBook** | Pointer to [**Book**](Book.md) |  | [optional] 
**LastBook** | Pointer to [**Book**](Book.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**RemotePoster** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**MonitorNewItems** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**CleanName** | Pointer to **NullableString** |  | [optional] 
**SortName** | Pointer to **NullableString** |  | [optional] 
**SortNameLastFirst** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**AddOptions** | Pointer to [**AddAuthorOptions**](AddAuthorOptions.md) |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**Statistics** | Pointer to [**AuthorStatisticsResource**](AuthorStatisticsResource.md) |  | [optional] 

## Methods

### NewAuthorResource

`func NewAuthorResource() *AuthorResource`

NewAuthorResource instantiates a new AuthorResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorResourceWithDefaults

`func NewAuthorResourceWithDefaults() *AuthorResource`

NewAuthorResourceWithDefaults instantiates a new AuthorResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorMetadataId

`func (o *AuthorResource) GetAuthorMetadataId() int32`

GetAuthorMetadataId returns the AuthorMetadataId field if non-nil, zero value otherwise.

### GetAuthorMetadataIdOk

`func (o *AuthorResource) GetAuthorMetadataIdOk() (*int32, bool)`

GetAuthorMetadataIdOk returns a tuple with the AuthorMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMetadataId

`func (o *AuthorResource) SetAuthorMetadataId(v int32)`

SetAuthorMetadataId sets AuthorMetadataId field to given value.

### HasAuthorMetadataId

`func (o *AuthorResource) HasAuthorMetadataId() bool`

HasAuthorMetadataId returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorResource) GetStatus() AuthorStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorResource) GetStatusOk() (*AuthorStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorResource) SetStatus(v AuthorStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnded

`func (o *AuthorResource) GetEnded() bool`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *AuthorResource) GetEndedOk() (*bool, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *AuthorResource) SetEnded(v bool)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *AuthorResource) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetAuthorName

`func (o *AuthorResource) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *AuthorResource) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *AuthorResource) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *AuthorResource) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *AuthorResource) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *AuthorResource) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetAuthorNameLastFirst

`func (o *AuthorResource) GetAuthorNameLastFirst() string`

GetAuthorNameLastFirst returns the AuthorNameLastFirst field if non-nil, zero value otherwise.

### GetAuthorNameLastFirstOk

`func (o *AuthorResource) GetAuthorNameLastFirstOk() (*string, bool)`

GetAuthorNameLastFirstOk returns a tuple with the AuthorNameLastFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorNameLastFirst

`func (o *AuthorResource) SetAuthorNameLastFirst(v string)`

SetAuthorNameLastFirst sets AuthorNameLastFirst field to given value.

### HasAuthorNameLastFirst

`func (o *AuthorResource) HasAuthorNameLastFirst() bool`

HasAuthorNameLastFirst returns a boolean if a field has been set.

### SetAuthorNameLastFirstNil

`func (o *AuthorResource) SetAuthorNameLastFirstNil(b bool)`

 SetAuthorNameLastFirstNil sets the value for AuthorNameLastFirst to be an explicit nil

### UnsetAuthorNameLastFirst
`func (o *AuthorResource) UnsetAuthorNameLastFirst()`

UnsetAuthorNameLastFirst ensures that no value is present for AuthorNameLastFirst, not even an explicit nil
### GetForeignAuthorId

`func (o *AuthorResource) GetForeignAuthorId() string`

GetForeignAuthorId returns the ForeignAuthorId field if non-nil, zero value otherwise.

### GetForeignAuthorIdOk

`func (o *AuthorResource) GetForeignAuthorIdOk() (*string, bool)`

GetForeignAuthorIdOk returns a tuple with the ForeignAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAuthorId

`func (o *AuthorResource) SetForeignAuthorId(v string)`

SetForeignAuthorId sets ForeignAuthorId field to given value.

### HasForeignAuthorId

`func (o *AuthorResource) HasForeignAuthorId() bool`

HasForeignAuthorId returns a boolean if a field has been set.

### SetForeignAuthorIdNil

`func (o *AuthorResource) SetForeignAuthorIdNil(b bool)`

 SetForeignAuthorIdNil sets the value for ForeignAuthorId to be an explicit nil

### UnsetForeignAuthorId
`func (o *AuthorResource) UnsetForeignAuthorId()`

UnsetForeignAuthorId ensures that no value is present for ForeignAuthorId, not even an explicit nil
### GetTitleSlug

`func (o *AuthorResource) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *AuthorResource) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *AuthorResource) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *AuthorResource) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *AuthorResource) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *AuthorResource) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetOverview

`func (o *AuthorResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *AuthorResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *AuthorResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *AuthorResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *AuthorResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *AuthorResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetDisambiguation

`func (o *AuthorResource) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *AuthorResource) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *AuthorResource) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *AuthorResource) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *AuthorResource) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *AuthorResource) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetLinks

`func (o *AuthorResource) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorResource) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorResource) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AuthorResource) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AuthorResource) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNextBook

`func (o *AuthorResource) GetNextBook() Book`

GetNextBook returns the NextBook field if non-nil, zero value otherwise.

### GetNextBookOk

`func (o *AuthorResource) GetNextBookOk() (*Book, bool)`

GetNextBookOk returns a tuple with the NextBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBook

`func (o *AuthorResource) SetNextBook(v Book)`

SetNextBook sets NextBook field to given value.

### HasNextBook

`func (o *AuthorResource) HasNextBook() bool`

HasNextBook returns a boolean if a field has been set.

### GetLastBook

`func (o *AuthorResource) GetLastBook() Book`

GetLastBook returns the LastBook field if non-nil, zero value otherwise.

### GetLastBookOk

`func (o *AuthorResource) GetLastBookOk() (*Book, bool)`

GetLastBookOk returns a tuple with the LastBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBook

`func (o *AuthorResource) SetLastBook(v Book)`

SetLastBook sets LastBook field to given value.

### HasLastBook

`func (o *AuthorResource) HasLastBook() bool`

HasLastBook returns a boolean if a field has been set.

### GetImages

`func (o *AuthorResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AuthorResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AuthorResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *AuthorResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *AuthorResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *AuthorResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetRemotePoster

`func (o *AuthorResource) GetRemotePoster() string`

GetRemotePoster returns the RemotePoster field if non-nil, zero value otherwise.

### GetRemotePosterOk

`func (o *AuthorResource) GetRemotePosterOk() (*string, bool)`

GetRemotePosterOk returns a tuple with the RemotePoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePoster

`func (o *AuthorResource) SetRemotePoster(v string)`

SetRemotePoster sets RemotePoster field to given value.

### HasRemotePoster

`func (o *AuthorResource) HasRemotePoster() bool`

HasRemotePoster returns a boolean if a field has been set.

### SetRemotePosterNil

`func (o *AuthorResource) SetRemotePosterNil(b bool)`

 SetRemotePosterNil sets the value for RemotePoster to be an explicit nil

### UnsetRemotePoster
`func (o *AuthorResource) UnsetRemotePoster()`

UnsetRemotePoster ensures that no value is present for RemotePoster, not even an explicit nil
### GetPath

`func (o *AuthorResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AuthorResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AuthorResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AuthorResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AuthorResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AuthorResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetQualityProfileId

`func (o *AuthorResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *AuthorResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *AuthorResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *AuthorResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *AuthorResource) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *AuthorResource) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *AuthorResource) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *AuthorResource) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetMonitored

`func (o *AuthorResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AuthorResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AuthorResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AuthorResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetMonitorNewItems

`func (o *AuthorResource) GetMonitorNewItems() NewItemMonitorTypes`

GetMonitorNewItems returns the MonitorNewItems field if non-nil, zero value otherwise.

### GetMonitorNewItemsOk

`func (o *AuthorResource) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool)`

GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorNewItems

`func (o *AuthorResource) SetMonitorNewItems(v NewItemMonitorTypes)`

SetMonitorNewItems sets MonitorNewItems field to given value.

### HasMonitorNewItems

`func (o *AuthorResource) HasMonitorNewItems() bool`

HasMonitorNewItems returns a boolean if a field has been set.

### GetRootFolderPath

`func (o *AuthorResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *AuthorResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *AuthorResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *AuthorResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *AuthorResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *AuthorResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetGenres

`func (o *AuthorResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AuthorResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AuthorResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *AuthorResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *AuthorResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *AuthorResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetCleanName

`func (o *AuthorResource) GetCleanName() string`

GetCleanName returns the CleanName field if non-nil, zero value otherwise.

### GetCleanNameOk

`func (o *AuthorResource) GetCleanNameOk() (*string, bool)`

GetCleanNameOk returns a tuple with the CleanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanName

`func (o *AuthorResource) SetCleanName(v string)`

SetCleanName sets CleanName field to given value.

### HasCleanName

`func (o *AuthorResource) HasCleanName() bool`

HasCleanName returns a boolean if a field has been set.

### SetCleanNameNil

`func (o *AuthorResource) SetCleanNameNil(b bool)`

 SetCleanNameNil sets the value for CleanName to be an explicit nil

### UnsetCleanName
`func (o *AuthorResource) UnsetCleanName()`

UnsetCleanName ensures that no value is present for CleanName, not even an explicit nil
### GetSortName

`func (o *AuthorResource) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *AuthorResource) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *AuthorResource) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *AuthorResource) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### SetSortNameNil

`func (o *AuthorResource) SetSortNameNil(b bool)`

 SetSortNameNil sets the value for SortName to be an explicit nil

### UnsetSortName
`func (o *AuthorResource) UnsetSortName()`

UnsetSortName ensures that no value is present for SortName, not even an explicit nil
### GetSortNameLastFirst

`func (o *AuthorResource) GetSortNameLastFirst() string`

GetSortNameLastFirst returns the SortNameLastFirst field if non-nil, zero value otherwise.

### GetSortNameLastFirstOk

`func (o *AuthorResource) GetSortNameLastFirstOk() (*string, bool)`

GetSortNameLastFirstOk returns a tuple with the SortNameLastFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortNameLastFirst

`func (o *AuthorResource) SetSortNameLastFirst(v string)`

SetSortNameLastFirst sets SortNameLastFirst field to given value.

### HasSortNameLastFirst

`func (o *AuthorResource) HasSortNameLastFirst() bool`

HasSortNameLastFirst returns a boolean if a field has been set.

### SetSortNameLastFirstNil

`func (o *AuthorResource) SetSortNameLastFirstNil(b bool)`

 SetSortNameLastFirstNil sets the value for SortNameLastFirst to be an explicit nil

### UnsetSortNameLastFirst
`func (o *AuthorResource) UnsetSortNameLastFirst()`

UnsetSortNameLastFirst ensures that no value is present for SortNameLastFirst, not even an explicit nil
### GetTags

`func (o *AuthorResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AuthorResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AuthorResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AuthorResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AuthorResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AuthorResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAdded

`func (o *AuthorResource) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *AuthorResource) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *AuthorResource) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *AuthorResource) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetAddOptions

`func (o *AuthorResource) GetAddOptions() AddAuthorOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *AuthorResource) GetAddOptionsOk() (*AddAuthorOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *AuthorResource) SetAddOptions(v AddAuthorOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *AuthorResource) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetRatings

`func (o *AuthorResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *AuthorResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *AuthorResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *AuthorResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetStatistics

`func (o *AuthorResource) GetStatistics() AuthorStatisticsResource`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *AuthorResource) GetStatisticsOk() (*AuthorStatisticsResource, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *AuthorResource) SetStatistics(v AuthorStatisticsResource)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *AuthorResource) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


