# Book

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorMetadataId** | Pointer to **int32** |  | [optional] 
**ForeignBookId** | Pointer to **NullableString** |  | [optional] 
**ForeignEditionId** | Pointer to **NullableString** |  | [optional] 
**TitleSlug** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **NullableTime** |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**RelatedBooks** | Pointer to **[]int32** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**LastSearchTime** | Pointer to **NullableTime** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**AnyEditionOk** | Pointer to **bool** |  | [optional] 
**LastInfoSync** | Pointer to **NullableTime** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**AddOptions** | Pointer to [**AddBookOptions**](AddBookOptions.md) |  | [optional] 
**AuthorMetadata** | Pointer to [**AuthorMetadataLazyLoaded**](AuthorMetadataLazyLoaded.md) |  | [optional] 
**Author** | Pointer to [**AuthorLazyLoaded**](AuthorLazyLoaded.md) |  | [optional] 
**Editions** | Pointer to [**EditionListLazyLoaded**](EditionListLazyLoaded.md) |  | [optional] 
**BookFiles** | Pointer to [**BookFileListLazyLoaded**](BookFileListLazyLoaded.md) |  | [optional] 
**SeriesLinks** | Pointer to [**SeriesBookLinkListLazyLoaded**](SeriesBookLinkListLazyLoaded.md) |  | [optional] 

## Methods

### NewBook

`func NewBook() *Book`

NewBook instantiates a new Book object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookWithDefaults

`func NewBookWithDefaults() *Book`

NewBookWithDefaults instantiates a new Book object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Book) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Book) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Book) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Book) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorMetadataId

`func (o *Book) GetAuthorMetadataId() int32`

GetAuthorMetadataId returns the AuthorMetadataId field if non-nil, zero value otherwise.

### GetAuthorMetadataIdOk

`func (o *Book) GetAuthorMetadataIdOk() (*int32, bool)`

GetAuthorMetadataIdOk returns a tuple with the AuthorMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMetadataId

`func (o *Book) SetAuthorMetadataId(v int32)`

SetAuthorMetadataId sets AuthorMetadataId field to given value.

### HasAuthorMetadataId

`func (o *Book) HasAuthorMetadataId() bool`

HasAuthorMetadataId returns a boolean if a field has been set.

### GetForeignBookId

`func (o *Book) GetForeignBookId() string`

GetForeignBookId returns the ForeignBookId field if non-nil, zero value otherwise.

### GetForeignBookIdOk

`func (o *Book) GetForeignBookIdOk() (*string, bool)`

GetForeignBookIdOk returns a tuple with the ForeignBookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignBookId

`func (o *Book) SetForeignBookId(v string)`

SetForeignBookId sets ForeignBookId field to given value.

### HasForeignBookId

`func (o *Book) HasForeignBookId() bool`

HasForeignBookId returns a boolean if a field has been set.

### SetForeignBookIdNil

`func (o *Book) SetForeignBookIdNil(b bool)`

 SetForeignBookIdNil sets the value for ForeignBookId to be an explicit nil

### UnsetForeignBookId
`func (o *Book) UnsetForeignBookId()`

UnsetForeignBookId ensures that no value is present for ForeignBookId, not even an explicit nil
### GetForeignEditionId

`func (o *Book) GetForeignEditionId() string`

GetForeignEditionId returns the ForeignEditionId field if non-nil, zero value otherwise.

### GetForeignEditionIdOk

`func (o *Book) GetForeignEditionIdOk() (*string, bool)`

GetForeignEditionIdOk returns a tuple with the ForeignEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignEditionId

`func (o *Book) SetForeignEditionId(v string)`

SetForeignEditionId sets ForeignEditionId field to given value.

### HasForeignEditionId

`func (o *Book) HasForeignEditionId() bool`

HasForeignEditionId returns a boolean if a field has been set.

### SetForeignEditionIdNil

`func (o *Book) SetForeignEditionIdNil(b bool)`

 SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil

### UnsetForeignEditionId
`func (o *Book) UnsetForeignEditionId()`

UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
### GetTitleSlug

`func (o *Book) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *Book) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *Book) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *Book) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *Book) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *Book) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetTitle

`func (o *Book) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Book) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Book) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Book) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Book) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Book) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetReleaseDate

`func (o *Book) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Book) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Book) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *Book) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *Book) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *Book) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetLinks

`func (o *Book) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Book) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Book) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Book) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *Book) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *Book) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetGenres

`func (o *Book) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *Book) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *Book) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *Book) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *Book) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *Book) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetRelatedBooks

`func (o *Book) GetRelatedBooks() []int32`

GetRelatedBooks returns the RelatedBooks field if non-nil, zero value otherwise.

### GetRelatedBooksOk

`func (o *Book) GetRelatedBooksOk() (*[]int32, bool)`

GetRelatedBooksOk returns a tuple with the RelatedBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedBooks

`func (o *Book) SetRelatedBooks(v []int32)`

SetRelatedBooks sets RelatedBooks field to given value.

### HasRelatedBooks

`func (o *Book) HasRelatedBooks() bool`

HasRelatedBooks returns a boolean if a field has been set.

### SetRelatedBooksNil

`func (o *Book) SetRelatedBooksNil(b bool)`

 SetRelatedBooksNil sets the value for RelatedBooks to be an explicit nil

### UnsetRelatedBooks
`func (o *Book) UnsetRelatedBooks()`

UnsetRelatedBooks ensures that no value is present for RelatedBooks, not even an explicit nil
### GetRatings

`func (o *Book) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *Book) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *Book) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *Book) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetLastSearchTime

`func (o *Book) GetLastSearchTime() time.Time`

GetLastSearchTime returns the LastSearchTime field if non-nil, zero value otherwise.

### GetLastSearchTimeOk

`func (o *Book) GetLastSearchTimeOk() (*time.Time, bool)`

GetLastSearchTimeOk returns a tuple with the LastSearchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSearchTime

`func (o *Book) SetLastSearchTime(v time.Time)`

SetLastSearchTime sets LastSearchTime field to given value.

### HasLastSearchTime

`func (o *Book) HasLastSearchTime() bool`

HasLastSearchTime returns a boolean if a field has been set.

### SetLastSearchTimeNil

`func (o *Book) SetLastSearchTimeNil(b bool)`

 SetLastSearchTimeNil sets the value for LastSearchTime to be an explicit nil

### UnsetLastSearchTime
`func (o *Book) UnsetLastSearchTime()`

UnsetLastSearchTime ensures that no value is present for LastSearchTime, not even an explicit nil
### GetCleanTitle

`func (o *Book) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *Book) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *Book) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *Book) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *Book) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *Book) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetMonitored

`func (o *Book) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *Book) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *Book) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *Book) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetAnyEditionOk

`func (o *Book) GetAnyEditionOk() bool`

GetAnyEditionOk returns the AnyEditionOk field if non-nil, zero value otherwise.

### GetAnyEditionOkOk

`func (o *Book) GetAnyEditionOkOk() (*bool, bool)`

GetAnyEditionOkOk returns a tuple with the AnyEditionOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyEditionOk

`func (o *Book) SetAnyEditionOk(v bool)`

SetAnyEditionOk sets AnyEditionOk field to given value.

### HasAnyEditionOk

`func (o *Book) HasAnyEditionOk() bool`

HasAnyEditionOk returns a boolean if a field has been set.

### GetLastInfoSync

`func (o *Book) GetLastInfoSync() time.Time`

GetLastInfoSync returns the LastInfoSync field if non-nil, zero value otherwise.

### GetLastInfoSyncOk

`func (o *Book) GetLastInfoSyncOk() (*time.Time, bool)`

GetLastInfoSyncOk returns a tuple with the LastInfoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInfoSync

`func (o *Book) SetLastInfoSync(v time.Time)`

SetLastInfoSync sets LastInfoSync field to given value.

### HasLastInfoSync

`func (o *Book) HasLastInfoSync() bool`

HasLastInfoSync returns a boolean if a field has been set.

### SetLastInfoSyncNil

`func (o *Book) SetLastInfoSyncNil(b bool)`

 SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil

### UnsetLastInfoSync
`func (o *Book) UnsetLastInfoSync()`

UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
### GetAdded

`func (o *Book) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *Book) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *Book) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *Book) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetAddOptions

`func (o *Book) GetAddOptions() AddBookOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *Book) GetAddOptionsOk() (*AddBookOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *Book) SetAddOptions(v AddBookOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *Book) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetAuthorMetadata

`func (o *Book) GetAuthorMetadata() AuthorMetadataLazyLoaded`

GetAuthorMetadata returns the AuthorMetadata field if non-nil, zero value otherwise.

### GetAuthorMetadataOk

`func (o *Book) GetAuthorMetadataOk() (*AuthorMetadataLazyLoaded, bool)`

GetAuthorMetadataOk returns a tuple with the AuthorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMetadata

`func (o *Book) SetAuthorMetadata(v AuthorMetadataLazyLoaded)`

SetAuthorMetadata sets AuthorMetadata field to given value.

### HasAuthorMetadata

`func (o *Book) HasAuthorMetadata() bool`

HasAuthorMetadata returns a boolean if a field has been set.

### GetAuthor

`func (o *Book) GetAuthor() AuthorLazyLoaded`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Book) GetAuthorOk() (*AuthorLazyLoaded, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Book) SetAuthor(v AuthorLazyLoaded)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Book) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetEditions

`func (o *Book) GetEditions() EditionListLazyLoaded`

GetEditions returns the Editions field if non-nil, zero value otherwise.

### GetEditionsOk

`func (o *Book) GetEditionsOk() (*EditionListLazyLoaded, bool)`

GetEditionsOk returns a tuple with the Editions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditions

`func (o *Book) SetEditions(v EditionListLazyLoaded)`

SetEditions sets Editions field to given value.

### HasEditions

`func (o *Book) HasEditions() bool`

HasEditions returns a boolean if a field has been set.

### GetBookFiles

`func (o *Book) GetBookFiles() BookFileListLazyLoaded`

GetBookFiles returns the BookFiles field if non-nil, zero value otherwise.

### GetBookFilesOk

`func (o *Book) GetBookFilesOk() (*BookFileListLazyLoaded, bool)`

GetBookFilesOk returns a tuple with the BookFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookFiles

`func (o *Book) SetBookFiles(v BookFileListLazyLoaded)`

SetBookFiles sets BookFiles field to given value.

### HasBookFiles

`func (o *Book) HasBookFiles() bool`

HasBookFiles returns a boolean if a field has been set.

### GetSeriesLinks

`func (o *Book) GetSeriesLinks() SeriesBookLinkListLazyLoaded`

GetSeriesLinks returns the SeriesLinks field if non-nil, zero value otherwise.

### GetSeriesLinksOk

`func (o *Book) GetSeriesLinksOk() (*SeriesBookLinkListLazyLoaded, bool)`

GetSeriesLinksOk returns a tuple with the SeriesLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLinks

`func (o *Book) SetSeriesLinks(v SeriesBookLinkListLazyLoaded)`

SetSeriesLinks sets SeriesLinks field to given value.

### HasSeriesLinks

`func (o *Book) HasSeriesLinks() bool`

HasSeriesLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


