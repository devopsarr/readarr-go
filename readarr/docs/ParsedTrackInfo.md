# ParsedTrackInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**Authors** | Pointer to **[]string** |  | [optional] 
**AuthorTitle** | Pointer to **NullableString** |  | [optional] [readonly] 
**BookTitle** | Pointer to **NullableString** |  | [optional] 
**SeriesTitle** | Pointer to **NullableString** |  | [optional] 
**SeriesIndex** | Pointer to **NullableString** |  | [optional] 
**Isbn** | Pointer to **NullableString** |  | [optional] 
**Asin** | Pointer to **NullableString** |  | [optional] 
**GoodreadsId** | Pointer to **NullableString** |  | [optional] 
**AuthorMBId** | Pointer to **NullableString** |  | [optional] 
**BookMBId** | Pointer to **NullableString** |  | [optional] 
**ReleaseMBId** | Pointer to **NullableString** |  | [optional] 
**RecordingMBId** | Pointer to **NullableString** |  | [optional] 
**TrackMBId** | Pointer to **NullableString** |  | [optional] 
**DiscNumber** | Pointer to **int32** |  | [optional] 
**DiscCount** | Pointer to **int32** |  | [optional] 
**Country** | Pointer to [**IsoCountry**](IsoCountry.md) |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Publisher** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**CatalogNumber** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to [**TimeSpan**](TimeSpan.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoModel**](MediaInfoModel.md) |  | [optional] 
**TrackNumbers** | Pointer to **[]int32** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewParsedTrackInfo

`func NewParsedTrackInfo() *ParsedTrackInfo`

NewParsedTrackInfo instantiates a new ParsedTrackInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsedTrackInfoWithDefaults

`func NewParsedTrackInfoWithDefaults() *ParsedTrackInfo`

NewParsedTrackInfoWithDefaults instantiates a new ParsedTrackInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ParsedTrackInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParsedTrackInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParsedTrackInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ParsedTrackInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ParsedTrackInfo) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ParsedTrackInfo) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCleanTitle

`func (o *ParsedTrackInfo) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *ParsedTrackInfo) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *ParsedTrackInfo) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *ParsedTrackInfo) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *ParsedTrackInfo) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *ParsedTrackInfo) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetAuthors

`func (o *ParsedTrackInfo) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ParsedTrackInfo) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ParsedTrackInfo) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *ParsedTrackInfo) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### SetAuthorsNil

`func (o *ParsedTrackInfo) SetAuthorsNil(b bool)`

 SetAuthorsNil sets the value for Authors to be an explicit nil

### UnsetAuthors
`func (o *ParsedTrackInfo) UnsetAuthors()`

UnsetAuthors ensures that no value is present for Authors, not even an explicit nil
### GetAuthorTitle

`func (o *ParsedTrackInfo) GetAuthorTitle() string`

GetAuthorTitle returns the AuthorTitle field if non-nil, zero value otherwise.

### GetAuthorTitleOk

`func (o *ParsedTrackInfo) GetAuthorTitleOk() (*string, bool)`

GetAuthorTitleOk returns a tuple with the AuthorTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTitle

`func (o *ParsedTrackInfo) SetAuthorTitle(v string)`

SetAuthorTitle sets AuthorTitle field to given value.

### HasAuthorTitle

`func (o *ParsedTrackInfo) HasAuthorTitle() bool`

HasAuthorTitle returns a boolean if a field has been set.

### SetAuthorTitleNil

`func (o *ParsedTrackInfo) SetAuthorTitleNil(b bool)`

 SetAuthorTitleNil sets the value for AuthorTitle to be an explicit nil

### UnsetAuthorTitle
`func (o *ParsedTrackInfo) UnsetAuthorTitle()`

UnsetAuthorTitle ensures that no value is present for AuthorTitle, not even an explicit nil
### GetBookTitle

`func (o *ParsedTrackInfo) GetBookTitle() string`

GetBookTitle returns the BookTitle field if non-nil, zero value otherwise.

### GetBookTitleOk

`func (o *ParsedTrackInfo) GetBookTitleOk() (*string, bool)`

GetBookTitleOk returns a tuple with the BookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTitle

`func (o *ParsedTrackInfo) SetBookTitle(v string)`

SetBookTitle sets BookTitle field to given value.

### HasBookTitle

`func (o *ParsedTrackInfo) HasBookTitle() bool`

HasBookTitle returns a boolean if a field has been set.

### SetBookTitleNil

`func (o *ParsedTrackInfo) SetBookTitleNil(b bool)`

 SetBookTitleNil sets the value for BookTitle to be an explicit nil

### UnsetBookTitle
`func (o *ParsedTrackInfo) UnsetBookTitle()`

UnsetBookTitle ensures that no value is present for BookTitle, not even an explicit nil
### GetSeriesTitle

`func (o *ParsedTrackInfo) GetSeriesTitle() string`

GetSeriesTitle returns the SeriesTitle field if non-nil, zero value otherwise.

### GetSeriesTitleOk

`func (o *ParsedTrackInfo) GetSeriesTitleOk() (*string, bool)`

GetSeriesTitleOk returns a tuple with the SeriesTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTitle

`func (o *ParsedTrackInfo) SetSeriesTitle(v string)`

SetSeriesTitle sets SeriesTitle field to given value.

### HasSeriesTitle

`func (o *ParsedTrackInfo) HasSeriesTitle() bool`

HasSeriesTitle returns a boolean if a field has been set.

### SetSeriesTitleNil

`func (o *ParsedTrackInfo) SetSeriesTitleNil(b bool)`

 SetSeriesTitleNil sets the value for SeriesTitle to be an explicit nil

### UnsetSeriesTitle
`func (o *ParsedTrackInfo) UnsetSeriesTitle()`

UnsetSeriesTitle ensures that no value is present for SeriesTitle, not even an explicit nil
### GetSeriesIndex

`func (o *ParsedTrackInfo) GetSeriesIndex() string`

GetSeriesIndex returns the SeriesIndex field if non-nil, zero value otherwise.

### GetSeriesIndexOk

`func (o *ParsedTrackInfo) GetSeriesIndexOk() (*string, bool)`

GetSeriesIndexOk returns a tuple with the SeriesIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesIndex

`func (o *ParsedTrackInfo) SetSeriesIndex(v string)`

SetSeriesIndex sets SeriesIndex field to given value.

### HasSeriesIndex

`func (o *ParsedTrackInfo) HasSeriesIndex() bool`

HasSeriesIndex returns a boolean if a field has been set.

### SetSeriesIndexNil

`func (o *ParsedTrackInfo) SetSeriesIndexNil(b bool)`

 SetSeriesIndexNil sets the value for SeriesIndex to be an explicit nil

### UnsetSeriesIndex
`func (o *ParsedTrackInfo) UnsetSeriesIndex()`

UnsetSeriesIndex ensures that no value is present for SeriesIndex, not even an explicit nil
### GetIsbn

`func (o *ParsedTrackInfo) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ParsedTrackInfo) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ParsedTrackInfo) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ParsedTrackInfo) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### SetIsbnNil

`func (o *ParsedTrackInfo) SetIsbnNil(b bool)`

 SetIsbnNil sets the value for Isbn to be an explicit nil

### UnsetIsbn
`func (o *ParsedTrackInfo) UnsetIsbn()`

UnsetIsbn ensures that no value is present for Isbn, not even an explicit nil
### GetAsin

`func (o *ParsedTrackInfo) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *ParsedTrackInfo) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *ParsedTrackInfo) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *ParsedTrackInfo) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### SetAsinNil

`func (o *ParsedTrackInfo) SetAsinNil(b bool)`

 SetAsinNil sets the value for Asin to be an explicit nil

### UnsetAsin
`func (o *ParsedTrackInfo) UnsetAsin()`

UnsetAsin ensures that no value is present for Asin, not even an explicit nil
### GetGoodreadsId

`func (o *ParsedTrackInfo) GetGoodreadsId() string`

GetGoodreadsId returns the GoodreadsId field if non-nil, zero value otherwise.

### GetGoodreadsIdOk

`func (o *ParsedTrackInfo) GetGoodreadsIdOk() (*string, bool)`

GetGoodreadsIdOk returns a tuple with the GoodreadsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodreadsId

`func (o *ParsedTrackInfo) SetGoodreadsId(v string)`

SetGoodreadsId sets GoodreadsId field to given value.

### HasGoodreadsId

`func (o *ParsedTrackInfo) HasGoodreadsId() bool`

HasGoodreadsId returns a boolean if a field has been set.

### SetGoodreadsIdNil

`func (o *ParsedTrackInfo) SetGoodreadsIdNil(b bool)`

 SetGoodreadsIdNil sets the value for GoodreadsId to be an explicit nil

### UnsetGoodreadsId
`func (o *ParsedTrackInfo) UnsetGoodreadsId()`

UnsetGoodreadsId ensures that no value is present for GoodreadsId, not even an explicit nil
### GetAuthorMBId

`func (o *ParsedTrackInfo) GetAuthorMBId() string`

GetAuthorMBId returns the AuthorMBId field if non-nil, zero value otherwise.

### GetAuthorMBIdOk

`func (o *ParsedTrackInfo) GetAuthorMBIdOk() (*string, bool)`

GetAuthorMBIdOk returns a tuple with the AuthorMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMBId

`func (o *ParsedTrackInfo) SetAuthorMBId(v string)`

SetAuthorMBId sets AuthorMBId field to given value.

### HasAuthorMBId

`func (o *ParsedTrackInfo) HasAuthorMBId() bool`

HasAuthorMBId returns a boolean if a field has been set.

### SetAuthorMBIdNil

`func (o *ParsedTrackInfo) SetAuthorMBIdNil(b bool)`

 SetAuthorMBIdNil sets the value for AuthorMBId to be an explicit nil

### UnsetAuthorMBId
`func (o *ParsedTrackInfo) UnsetAuthorMBId()`

UnsetAuthorMBId ensures that no value is present for AuthorMBId, not even an explicit nil
### GetBookMBId

`func (o *ParsedTrackInfo) GetBookMBId() string`

GetBookMBId returns the BookMBId field if non-nil, zero value otherwise.

### GetBookMBIdOk

`func (o *ParsedTrackInfo) GetBookMBIdOk() (*string, bool)`

GetBookMBIdOk returns a tuple with the BookMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookMBId

`func (o *ParsedTrackInfo) SetBookMBId(v string)`

SetBookMBId sets BookMBId field to given value.

### HasBookMBId

`func (o *ParsedTrackInfo) HasBookMBId() bool`

HasBookMBId returns a boolean if a field has been set.

### SetBookMBIdNil

`func (o *ParsedTrackInfo) SetBookMBIdNil(b bool)`

 SetBookMBIdNil sets the value for BookMBId to be an explicit nil

### UnsetBookMBId
`func (o *ParsedTrackInfo) UnsetBookMBId()`

UnsetBookMBId ensures that no value is present for BookMBId, not even an explicit nil
### GetReleaseMBId

`func (o *ParsedTrackInfo) GetReleaseMBId() string`

GetReleaseMBId returns the ReleaseMBId field if non-nil, zero value otherwise.

### GetReleaseMBIdOk

`func (o *ParsedTrackInfo) GetReleaseMBIdOk() (*string, bool)`

GetReleaseMBIdOk returns a tuple with the ReleaseMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseMBId

`func (o *ParsedTrackInfo) SetReleaseMBId(v string)`

SetReleaseMBId sets ReleaseMBId field to given value.

### HasReleaseMBId

`func (o *ParsedTrackInfo) HasReleaseMBId() bool`

HasReleaseMBId returns a boolean if a field has been set.

### SetReleaseMBIdNil

`func (o *ParsedTrackInfo) SetReleaseMBIdNil(b bool)`

 SetReleaseMBIdNil sets the value for ReleaseMBId to be an explicit nil

### UnsetReleaseMBId
`func (o *ParsedTrackInfo) UnsetReleaseMBId()`

UnsetReleaseMBId ensures that no value is present for ReleaseMBId, not even an explicit nil
### GetRecordingMBId

`func (o *ParsedTrackInfo) GetRecordingMBId() string`

GetRecordingMBId returns the RecordingMBId field if non-nil, zero value otherwise.

### GetRecordingMBIdOk

`func (o *ParsedTrackInfo) GetRecordingMBIdOk() (*string, bool)`

GetRecordingMBIdOk returns a tuple with the RecordingMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingMBId

`func (o *ParsedTrackInfo) SetRecordingMBId(v string)`

SetRecordingMBId sets RecordingMBId field to given value.

### HasRecordingMBId

`func (o *ParsedTrackInfo) HasRecordingMBId() bool`

HasRecordingMBId returns a boolean if a field has been set.

### SetRecordingMBIdNil

`func (o *ParsedTrackInfo) SetRecordingMBIdNil(b bool)`

 SetRecordingMBIdNil sets the value for RecordingMBId to be an explicit nil

### UnsetRecordingMBId
`func (o *ParsedTrackInfo) UnsetRecordingMBId()`

UnsetRecordingMBId ensures that no value is present for RecordingMBId, not even an explicit nil
### GetTrackMBId

`func (o *ParsedTrackInfo) GetTrackMBId() string`

GetTrackMBId returns the TrackMBId field if non-nil, zero value otherwise.

### GetTrackMBIdOk

`func (o *ParsedTrackInfo) GetTrackMBIdOk() (*string, bool)`

GetTrackMBIdOk returns a tuple with the TrackMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackMBId

`func (o *ParsedTrackInfo) SetTrackMBId(v string)`

SetTrackMBId sets TrackMBId field to given value.

### HasTrackMBId

`func (o *ParsedTrackInfo) HasTrackMBId() bool`

HasTrackMBId returns a boolean if a field has been set.

### SetTrackMBIdNil

`func (o *ParsedTrackInfo) SetTrackMBIdNil(b bool)`

 SetTrackMBIdNil sets the value for TrackMBId to be an explicit nil

### UnsetTrackMBId
`func (o *ParsedTrackInfo) UnsetTrackMBId()`

UnsetTrackMBId ensures that no value is present for TrackMBId, not even an explicit nil
### GetDiscNumber

`func (o *ParsedTrackInfo) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *ParsedTrackInfo) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *ParsedTrackInfo) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *ParsedTrackInfo) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDiscCount

`func (o *ParsedTrackInfo) GetDiscCount() int32`

GetDiscCount returns the DiscCount field if non-nil, zero value otherwise.

### GetDiscCountOk

`func (o *ParsedTrackInfo) GetDiscCountOk() (*int32, bool)`

GetDiscCountOk returns a tuple with the DiscCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscCount

`func (o *ParsedTrackInfo) SetDiscCount(v int32)`

SetDiscCount sets DiscCount field to given value.

### HasDiscCount

`func (o *ParsedTrackInfo) HasDiscCount() bool`

HasDiscCount returns a boolean if a field has been set.

### GetCountry

`func (o *ParsedTrackInfo) GetCountry() IsoCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ParsedTrackInfo) GetCountryOk() (*IsoCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ParsedTrackInfo) SetCountry(v IsoCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ParsedTrackInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetYear

`func (o *ParsedTrackInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ParsedTrackInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ParsedTrackInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ParsedTrackInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetPublisher

`func (o *ParsedTrackInfo) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ParsedTrackInfo) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ParsedTrackInfo) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ParsedTrackInfo) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *ParsedTrackInfo) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *ParsedTrackInfo) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetLabel

`func (o *ParsedTrackInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ParsedTrackInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ParsedTrackInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ParsedTrackInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ParsedTrackInfo) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ParsedTrackInfo) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetSource

`func (o *ParsedTrackInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ParsedTrackInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ParsedTrackInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ParsedTrackInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ParsedTrackInfo) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ParsedTrackInfo) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCatalogNumber

`func (o *ParsedTrackInfo) GetCatalogNumber() string`

GetCatalogNumber returns the CatalogNumber field if non-nil, zero value otherwise.

### GetCatalogNumberOk

`func (o *ParsedTrackInfo) GetCatalogNumberOk() (*string, bool)`

GetCatalogNumberOk returns a tuple with the CatalogNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogNumber

`func (o *ParsedTrackInfo) SetCatalogNumber(v string)`

SetCatalogNumber sets CatalogNumber field to given value.

### HasCatalogNumber

`func (o *ParsedTrackInfo) HasCatalogNumber() bool`

HasCatalogNumber returns a boolean if a field has been set.

### SetCatalogNumberNil

`func (o *ParsedTrackInfo) SetCatalogNumberNil(b bool)`

 SetCatalogNumberNil sets the value for CatalogNumber to be an explicit nil

### UnsetCatalogNumber
`func (o *ParsedTrackInfo) UnsetCatalogNumber()`

UnsetCatalogNumber ensures that no value is present for CatalogNumber, not even an explicit nil
### GetDisambiguation

`func (o *ParsedTrackInfo) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *ParsedTrackInfo) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *ParsedTrackInfo) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *ParsedTrackInfo) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *ParsedTrackInfo) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *ParsedTrackInfo) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetDuration

`func (o *ParsedTrackInfo) GetDuration() TimeSpan`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ParsedTrackInfo) GetDurationOk() (*TimeSpan, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ParsedTrackInfo) SetDuration(v TimeSpan)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ParsedTrackInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetQuality

`func (o *ParsedTrackInfo) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ParsedTrackInfo) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ParsedTrackInfo) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ParsedTrackInfo) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetMediaInfo

`func (o *ParsedTrackInfo) GetMediaInfo() MediaInfoModel`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *ParsedTrackInfo) GetMediaInfoOk() (*MediaInfoModel, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *ParsedTrackInfo) SetMediaInfo(v MediaInfoModel)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *ParsedTrackInfo) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetTrackNumbers

`func (o *ParsedTrackInfo) GetTrackNumbers() []int32`

GetTrackNumbers returns the TrackNumbers field if non-nil, zero value otherwise.

### GetTrackNumbersOk

`func (o *ParsedTrackInfo) GetTrackNumbersOk() (*[]int32, bool)`

GetTrackNumbersOk returns a tuple with the TrackNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumbers

`func (o *ParsedTrackInfo) SetTrackNumbers(v []int32)`

SetTrackNumbers sets TrackNumbers field to given value.

### HasTrackNumbers

`func (o *ParsedTrackInfo) HasTrackNumbers() bool`

HasTrackNumbers returns a boolean if a field has been set.

### SetTrackNumbersNil

`func (o *ParsedTrackInfo) SetTrackNumbersNil(b bool)`

 SetTrackNumbersNil sets the value for TrackNumbers to be an explicit nil

### UnsetTrackNumbers
`func (o *ParsedTrackInfo) UnsetTrackNumbers()`

UnsetTrackNumbers ensures that no value is present for TrackNumbers, not even an explicit nil
### GetLanguage

`func (o *ParsedTrackInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ParsedTrackInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ParsedTrackInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ParsedTrackInfo) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ParsedTrackInfo) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ParsedTrackInfo) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetReleaseGroup

`func (o *ParsedTrackInfo) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ParsedTrackInfo) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ParsedTrackInfo) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ParsedTrackInfo) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ParsedTrackInfo) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ParsedTrackInfo) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetReleaseHash

`func (o *ParsedTrackInfo) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ParsedTrackInfo) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ParsedTrackInfo) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ParsedTrackInfo) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ParsedTrackInfo) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ParsedTrackInfo) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


