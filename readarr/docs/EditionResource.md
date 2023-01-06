# EditionResource

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
**RemoteCover** | Pointer to **NullableString** |  | [optional] 
**Grabbed** | Pointer to **bool** |  | [optional] 

## Methods

### NewEditionResource

`func NewEditionResource() *EditionResource`

NewEditionResource instantiates a new EditionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditionResourceWithDefaults

`func NewEditionResourceWithDefaults() *EditionResource`

NewEditionResourceWithDefaults instantiates a new EditionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditionResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditionResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditionResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EditionResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBookId

`func (o *EditionResource) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *EditionResource) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *EditionResource) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *EditionResource) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### GetForeignEditionId

`func (o *EditionResource) GetForeignEditionId() string`

GetForeignEditionId returns the ForeignEditionId field if non-nil, zero value otherwise.

### GetForeignEditionIdOk

`func (o *EditionResource) GetForeignEditionIdOk() (*string, bool)`

GetForeignEditionIdOk returns a tuple with the ForeignEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignEditionId

`func (o *EditionResource) SetForeignEditionId(v string)`

SetForeignEditionId sets ForeignEditionId field to given value.

### HasForeignEditionId

`func (o *EditionResource) HasForeignEditionId() bool`

HasForeignEditionId returns a boolean if a field has been set.

### SetForeignEditionIdNil

`func (o *EditionResource) SetForeignEditionIdNil(b bool)`

 SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil

### UnsetForeignEditionId
`func (o *EditionResource) UnsetForeignEditionId()`

UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
### GetTitleSlug

`func (o *EditionResource) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *EditionResource) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *EditionResource) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *EditionResource) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *EditionResource) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *EditionResource) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetIsbn13

`func (o *EditionResource) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *EditionResource) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *EditionResource) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *EditionResource) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *EditionResource) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *EditionResource) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetAsin

`func (o *EditionResource) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *EditionResource) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *EditionResource) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *EditionResource) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### SetAsinNil

`func (o *EditionResource) SetAsinNil(b bool)`

 SetAsinNil sets the value for Asin to be an explicit nil

### UnsetAsin
`func (o *EditionResource) UnsetAsin()`

UnsetAsin ensures that no value is present for Asin, not even an explicit nil
### GetTitle

`func (o *EditionResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EditionResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EditionResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EditionResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *EditionResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *EditionResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLanguage

`func (o *EditionResource) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EditionResource) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EditionResource) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *EditionResource) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *EditionResource) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *EditionResource) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetOverview

`func (o *EditionResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *EditionResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *EditionResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *EditionResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *EditionResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *EditionResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetFormat

`func (o *EditionResource) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *EditionResource) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *EditionResource) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *EditionResource) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *EditionResource) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *EditionResource) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetIsEbook

`func (o *EditionResource) GetIsEbook() bool`

GetIsEbook returns the IsEbook field if non-nil, zero value otherwise.

### GetIsEbookOk

`func (o *EditionResource) GetIsEbookOk() (*bool, bool)`

GetIsEbookOk returns a tuple with the IsEbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEbook

`func (o *EditionResource) SetIsEbook(v bool)`

SetIsEbook sets IsEbook field to given value.

### HasIsEbook

`func (o *EditionResource) HasIsEbook() bool`

HasIsEbook returns a boolean if a field has been set.

### GetDisambiguation

`func (o *EditionResource) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *EditionResource) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *EditionResource) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *EditionResource) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *EditionResource) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *EditionResource) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetPublisher

`func (o *EditionResource) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *EditionResource) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *EditionResource) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *EditionResource) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *EditionResource) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *EditionResource) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetPageCount

`func (o *EditionResource) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *EditionResource) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *EditionResource) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *EditionResource) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetReleaseDate

`func (o *EditionResource) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *EditionResource) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *EditionResource) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *EditionResource) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *EditionResource) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *EditionResource) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetImages

`func (o *EditionResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *EditionResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *EditionResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *EditionResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *EditionResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *EditionResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetLinks

`func (o *EditionResource) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EditionResource) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EditionResource) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EditionResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *EditionResource) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EditionResource) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRatings

`func (o *EditionResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *EditionResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *EditionResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *EditionResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetMonitored

`func (o *EditionResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *EditionResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *EditionResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *EditionResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetManualAdd

`func (o *EditionResource) GetManualAdd() bool`

GetManualAdd returns the ManualAdd field if non-nil, zero value otherwise.

### GetManualAddOk

`func (o *EditionResource) GetManualAddOk() (*bool, bool)`

GetManualAddOk returns a tuple with the ManualAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualAdd

`func (o *EditionResource) SetManualAdd(v bool)`

SetManualAdd sets ManualAdd field to given value.

### HasManualAdd

`func (o *EditionResource) HasManualAdd() bool`

HasManualAdd returns a boolean if a field has been set.

### GetRemoteCover

`func (o *EditionResource) GetRemoteCover() string`

GetRemoteCover returns the RemoteCover field if non-nil, zero value otherwise.

### GetRemoteCoverOk

`func (o *EditionResource) GetRemoteCoverOk() (*string, bool)`

GetRemoteCoverOk returns a tuple with the RemoteCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCover

`func (o *EditionResource) SetRemoteCover(v string)`

SetRemoteCover sets RemoteCover field to given value.

### HasRemoteCover

`func (o *EditionResource) HasRemoteCover() bool`

HasRemoteCover returns a boolean if a field has been set.

### SetRemoteCoverNil

`func (o *EditionResource) SetRemoteCoverNil(b bool)`

 SetRemoteCoverNil sets the value for RemoteCover to be an explicit nil

### UnsetRemoteCover
`func (o *EditionResource) UnsetRemoteCover()`

UnsetRemoteCover ensures that no value is present for RemoteCover, not even an explicit nil
### GetGrabbed

`func (o *EditionResource) GetGrabbed() bool`

GetGrabbed returns the Grabbed field if non-nil, zero value otherwise.

### GetGrabbedOk

`func (o *EditionResource) GetGrabbedOk() (*bool, bool)`

GetGrabbedOk returns a tuple with the Grabbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabbed

`func (o *EditionResource) SetGrabbed(v bool)`

SetGrabbed sets Grabbed field to given value.

### HasGrabbed

`func (o *EditionResource) HasGrabbed() bool`

HasGrabbed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


