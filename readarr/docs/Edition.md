# Edition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BookId** | Pointer to **int32** |  | [optional] 
**ForeignEditionId** | Pointer to **NullableString** |  | [optional] 
**TitleSlug** | Pointer to **NullableString** |  | [optional] 
**Isbn13** | Pointer to **NullableString** |  | [optional] 
**Asin** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 
**IsEbook** | Pointer to **bool** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Publisher** | Pointer to **NullableString** |  | [optional] 
**PageCount** | Pointer to **int32** |  | [optional] 
**ReleaseDate** | Pointer to **NullableTime** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**ManualAdd** | Pointer to **bool** |  | [optional] 
**Book** | Pointer to [**BookLazyLoaded**](BookLazyLoaded.md) |  | [optional] 
**BookFiles** | Pointer to [**BookFileListLazyLoaded**](BookFileListLazyLoaded.md) |  | [optional] 

## Methods

### NewEdition

`func NewEdition() *Edition`

NewEdition instantiates a new Edition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditionWithDefaults

`func NewEditionWithDefaults() *Edition`

NewEditionWithDefaults instantiates a new Edition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Edition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Edition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Edition) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Edition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBookId

`func (o *Edition) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *Edition) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *Edition) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *Edition) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### GetForeignEditionId

`func (o *Edition) GetForeignEditionId() string`

GetForeignEditionId returns the ForeignEditionId field if non-nil, zero value otherwise.

### GetForeignEditionIdOk

`func (o *Edition) GetForeignEditionIdOk() (*string, bool)`

GetForeignEditionIdOk returns a tuple with the ForeignEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignEditionId

`func (o *Edition) SetForeignEditionId(v string)`

SetForeignEditionId sets ForeignEditionId field to given value.

### HasForeignEditionId

`func (o *Edition) HasForeignEditionId() bool`

HasForeignEditionId returns a boolean if a field has been set.

### SetForeignEditionIdNil

`func (o *Edition) SetForeignEditionIdNil(b bool)`

 SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil

### UnsetForeignEditionId
`func (o *Edition) UnsetForeignEditionId()`

UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
### GetTitleSlug

`func (o *Edition) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *Edition) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *Edition) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *Edition) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *Edition) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *Edition) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetIsbn13

`func (o *Edition) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *Edition) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *Edition) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *Edition) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *Edition) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *Edition) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetAsin

`func (o *Edition) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *Edition) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *Edition) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *Edition) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### SetAsinNil

`func (o *Edition) SetAsinNil(b bool)`

 SetAsinNil sets the value for Asin to be an explicit nil

### UnsetAsin
`func (o *Edition) UnsetAsin()`

UnsetAsin ensures that no value is present for Asin, not even an explicit nil
### GetTitle

`func (o *Edition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Edition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Edition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Edition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Edition) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Edition) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLanguage

`func (o *Edition) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Edition) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Edition) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Edition) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *Edition) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *Edition) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetOverview

`func (o *Edition) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *Edition) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *Edition) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *Edition) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *Edition) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *Edition) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetFormat

`func (o *Edition) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Edition) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Edition) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Edition) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *Edition) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *Edition) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetIsEbook

`func (o *Edition) GetIsEbook() bool`

GetIsEbook returns the IsEbook field if non-nil, zero value otherwise.

### GetIsEbookOk

`func (o *Edition) GetIsEbookOk() (*bool, bool)`

GetIsEbookOk returns a tuple with the IsEbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEbook

`func (o *Edition) SetIsEbook(v bool)`

SetIsEbook sets IsEbook field to given value.

### HasIsEbook

`func (o *Edition) HasIsEbook() bool`

HasIsEbook returns a boolean if a field has been set.

### GetDisambiguation

`func (o *Edition) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *Edition) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *Edition) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *Edition) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *Edition) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *Edition) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetPublisher

`func (o *Edition) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *Edition) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *Edition) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *Edition) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *Edition) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *Edition) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetPageCount

`func (o *Edition) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *Edition) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *Edition) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *Edition) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetReleaseDate

`func (o *Edition) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Edition) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Edition) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *Edition) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *Edition) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *Edition) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetImages

`func (o *Edition) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Edition) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Edition) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *Edition) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *Edition) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *Edition) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetLinks

`func (o *Edition) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Edition) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Edition) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Edition) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *Edition) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *Edition) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRatings

`func (o *Edition) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *Edition) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *Edition) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *Edition) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetMonitored

`func (o *Edition) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *Edition) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *Edition) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *Edition) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetManualAdd

`func (o *Edition) GetManualAdd() bool`

GetManualAdd returns the ManualAdd field if non-nil, zero value otherwise.

### GetManualAddOk

`func (o *Edition) GetManualAddOk() (*bool, bool)`

GetManualAddOk returns a tuple with the ManualAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualAdd

`func (o *Edition) SetManualAdd(v bool)`

SetManualAdd sets ManualAdd field to given value.

### HasManualAdd

`func (o *Edition) HasManualAdd() bool`

HasManualAdd returns a boolean if a field has been set.

### GetBook

`func (o *Edition) GetBook() BookLazyLoaded`

GetBook returns the Book field if non-nil, zero value otherwise.

### GetBookOk

`func (o *Edition) GetBookOk() (*BookLazyLoaded, bool)`

GetBookOk returns a tuple with the Book field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBook

`func (o *Edition) SetBook(v BookLazyLoaded)`

SetBook sets Book field to given value.

### HasBook

`func (o *Edition) HasBook() bool`

HasBook returns a boolean if a field has been set.

### GetBookFiles

`func (o *Edition) GetBookFiles() BookFileListLazyLoaded`

GetBookFiles returns the BookFiles field if non-nil, zero value otherwise.

### GetBookFilesOk

`func (o *Edition) GetBookFilesOk() (*BookFileListLazyLoaded, bool)`

GetBookFilesOk returns a tuple with the BookFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookFiles

`func (o *Edition) SetBookFiles(v BookFileListLazyLoaded)`

SetBookFiles sets BookFiles field to given value.

### HasBookFiles

`func (o *Edition) HasBookFiles() bool`

HasBookFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


