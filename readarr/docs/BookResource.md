# BookResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**AuthorTitle** | Pointer to **NullableString** |  | [optional] 
**SeriesTitle** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**AuthorId** | Pointer to **int32** |  | [optional] 
**ForeignBookId** | Pointer to **NullableString** |  | [optional] 
**ForeignEditionId** | Pointer to **NullableString** |  | [optional] 
**TitleSlug** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**AnyEditionOk** | Pointer to **bool** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**ReleaseDate** | Pointer to **NullableTime** |  | [optional] 
**PageCount** | Pointer to **int32** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Author** | Pointer to [**AuthorResource**](AuthorResource.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**Statistics** | Pointer to [**BookStatisticsResource**](BookStatisticsResource.md) |  | [optional] 
**Added** | Pointer to **NullableTime** |  | [optional] 
**AddOptions** | Pointer to [**AddBookOptions**](AddBookOptions.md) |  | [optional] 
**RemoteCover** | Pointer to **NullableString** |  | [optional] 
**Editions** | Pointer to [**[]EditionResource**](EditionResource.md) |  | [optional] 
**Grabbed** | Pointer to **bool** |  | [optional] 

## Methods

### NewBookResource

`func NewBookResource() *BookResource`

NewBookResource instantiates a new BookResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookResourceWithDefaults

`func NewBookResourceWithDefaults() *BookResource`

NewBookResourceWithDefaults instantiates a new BookResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BookResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *BookResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BookResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BookResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthorTitle

`func (o *BookResource) GetAuthorTitle() string`

GetAuthorTitle returns the AuthorTitle field if non-nil, zero value otherwise.

### GetAuthorTitleOk

`func (o *BookResource) GetAuthorTitleOk() (*string, bool)`

GetAuthorTitleOk returns a tuple with the AuthorTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTitle

`func (o *BookResource) SetAuthorTitle(v string)`

SetAuthorTitle sets AuthorTitle field to given value.

### HasAuthorTitle

`func (o *BookResource) HasAuthorTitle() bool`

HasAuthorTitle returns a boolean if a field has been set.

### SetAuthorTitleNil

`func (o *BookResource) SetAuthorTitleNil(b bool)`

 SetAuthorTitleNil sets the value for AuthorTitle to be an explicit nil

### UnsetAuthorTitle
`func (o *BookResource) UnsetAuthorTitle()`

UnsetAuthorTitle ensures that no value is present for AuthorTitle, not even an explicit nil
### GetSeriesTitle

`func (o *BookResource) GetSeriesTitle() string`

GetSeriesTitle returns the SeriesTitle field if non-nil, zero value otherwise.

### GetSeriesTitleOk

`func (o *BookResource) GetSeriesTitleOk() (*string, bool)`

GetSeriesTitleOk returns a tuple with the SeriesTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTitle

`func (o *BookResource) SetSeriesTitle(v string)`

SetSeriesTitle sets SeriesTitle field to given value.

### HasSeriesTitle

`func (o *BookResource) HasSeriesTitle() bool`

HasSeriesTitle returns a boolean if a field has been set.

### SetSeriesTitleNil

`func (o *BookResource) SetSeriesTitleNil(b bool)`

 SetSeriesTitleNil sets the value for SeriesTitle to be an explicit nil

### UnsetSeriesTitle
`func (o *BookResource) UnsetSeriesTitle()`

UnsetSeriesTitle ensures that no value is present for SeriesTitle, not even an explicit nil
### GetDisambiguation

`func (o *BookResource) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *BookResource) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *BookResource) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *BookResource) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *BookResource) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *BookResource) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetOverview

`func (o *BookResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *BookResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *BookResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *BookResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *BookResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *BookResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetAuthorId

`func (o *BookResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *BookResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *BookResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *BookResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetForeignBookId

`func (o *BookResource) GetForeignBookId() string`

GetForeignBookId returns the ForeignBookId field if non-nil, zero value otherwise.

### GetForeignBookIdOk

`func (o *BookResource) GetForeignBookIdOk() (*string, bool)`

GetForeignBookIdOk returns a tuple with the ForeignBookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignBookId

`func (o *BookResource) SetForeignBookId(v string)`

SetForeignBookId sets ForeignBookId field to given value.

### HasForeignBookId

`func (o *BookResource) HasForeignBookId() bool`

HasForeignBookId returns a boolean if a field has been set.

### SetForeignBookIdNil

`func (o *BookResource) SetForeignBookIdNil(b bool)`

 SetForeignBookIdNil sets the value for ForeignBookId to be an explicit nil

### UnsetForeignBookId
`func (o *BookResource) UnsetForeignBookId()`

UnsetForeignBookId ensures that no value is present for ForeignBookId, not even an explicit nil
### GetForeignEditionId

`func (o *BookResource) GetForeignEditionId() string`

GetForeignEditionId returns the ForeignEditionId field if non-nil, zero value otherwise.

### GetForeignEditionIdOk

`func (o *BookResource) GetForeignEditionIdOk() (*string, bool)`

GetForeignEditionIdOk returns a tuple with the ForeignEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignEditionId

`func (o *BookResource) SetForeignEditionId(v string)`

SetForeignEditionId sets ForeignEditionId field to given value.

### HasForeignEditionId

`func (o *BookResource) HasForeignEditionId() bool`

HasForeignEditionId returns a boolean if a field has been set.

### SetForeignEditionIdNil

`func (o *BookResource) SetForeignEditionIdNil(b bool)`

 SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil

### UnsetForeignEditionId
`func (o *BookResource) UnsetForeignEditionId()`

UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
### GetTitleSlug

`func (o *BookResource) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *BookResource) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *BookResource) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *BookResource) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *BookResource) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *BookResource) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetMonitored

`func (o *BookResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *BookResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *BookResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *BookResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetAnyEditionOk

`func (o *BookResource) GetAnyEditionOk() bool`

GetAnyEditionOk returns the AnyEditionOk field if non-nil, zero value otherwise.

### GetAnyEditionOkOk

`func (o *BookResource) GetAnyEditionOkOk() (*bool, bool)`

GetAnyEditionOkOk returns a tuple with the AnyEditionOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyEditionOk

`func (o *BookResource) SetAnyEditionOk(v bool)`

SetAnyEditionOk sets AnyEditionOk field to given value.

### HasAnyEditionOk

`func (o *BookResource) HasAnyEditionOk() bool`

HasAnyEditionOk returns a boolean if a field has been set.

### GetRatings

`func (o *BookResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *BookResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *BookResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *BookResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetReleaseDate

`func (o *BookResource) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *BookResource) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *BookResource) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *BookResource) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *BookResource) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *BookResource) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetPageCount

`func (o *BookResource) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *BookResource) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *BookResource) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *BookResource) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetGenres

`func (o *BookResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *BookResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *BookResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *BookResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *BookResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *BookResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetAuthor

`func (o *BookResource) GetAuthor() AuthorResource`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BookResource) GetAuthorOk() (*AuthorResource, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BookResource) SetAuthor(v AuthorResource)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *BookResource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetImages

`func (o *BookResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *BookResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *BookResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *BookResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *BookResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *BookResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetLinks

`func (o *BookResource) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BookResource) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BookResource) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BookResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *BookResource) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *BookResource) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStatistics

`func (o *BookResource) GetStatistics() BookStatisticsResource`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *BookResource) GetStatisticsOk() (*BookStatisticsResource, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *BookResource) SetStatistics(v BookStatisticsResource)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *BookResource) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetAdded

`func (o *BookResource) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *BookResource) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *BookResource) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *BookResource) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### SetAddedNil

`func (o *BookResource) SetAddedNil(b bool)`

 SetAddedNil sets the value for Added to be an explicit nil

### UnsetAdded
`func (o *BookResource) UnsetAdded()`

UnsetAdded ensures that no value is present for Added, not even an explicit nil
### GetAddOptions

`func (o *BookResource) GetAddOptions() AddBookOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *BookResource) GetAddOptionsOk() (*AddBookOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *BookResource) SetAddOptions(v AddBookOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *BookResource) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetRemoteCover

`func (o *BookResource) GetRemoteCover() string`

GetRemoteCover returns the RemoteCover field if non-nil, zero value otherwise.

### GetRemoteCoverOk

`func (o *BookResource) GetRemoteCoverOk() (*string, bool)`

GetRemoteCoverOk returns a tuple with the RemoteCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCover

`func (o *BookResource) SetRemoteCover(v string)`

SetRemoteCover sets RemoteCover field to given value.

### HasRemoteCover

`func (o *BookResource) HasRemoteCover() bool`

HasRemoteCover returns a boolean if a field has been set.

### SetRemoteCoverNil

`func (o *BookResource) SetRemoteCoverNil(b bool)`

 SetRemoteCoverNil sets the value for RemoteCover to be an explicit nil

### UnsetRemoteCover
`func (o *BookResource) UnsetRemoteCover()`

UnsetRemoteCover ensures that no value is present for RemoteCover, not even an explicit nil
### GetEditions

`func (o *BookResource) GetEditions() []EditionResource`

GetEditions returns the Editions field if non-nil, zero value otherwise.

### GetEditionsOk

`func (o *BookResource) GetEditionsOk() (*[]EditionResource, bool)`

GetEditionsOk returns a tuple with the Editions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditions

`func (o *BookResource) SetEditions(v []EditionResource)`

SetEditions sets Editions field to given value.

### HasEditions

`func (o *BookResource) HasEditions() bool`

HasEditions returns a boolean if a field has been set.

### SetEditionsNil

`func (o *BookResource) SetEditionsNil(b bool)`

 SetEditionsNil sets the value for Editions to be an explicit nil

### UnsetEditions
`func (o *BookResource) UnsetEditions()`

UnsetEditions ensures that no value is present for Editions, not even an explicit nil
### GetGrabbed

`func (o *BookResource) GetGrabbed() bool`

GetGrabbed returns the Grabbed field if non-nil, zero value otherwise.

### GetGrabbedOk

`func (o *BookResource) GetGrabbedOk() (*bool, bool)`

GetGrabbedOk returns a tuple with the Grabbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabbed

`func (o *BookResource) SetGrabbed(v bool)`

SetGrabbed sets Grabbed field to given value.

### HasGrabbed

`func (o *BookResource) HasGrabbed() bool`

HasGrabbed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


