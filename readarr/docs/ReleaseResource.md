# ReleaseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**QualityWeight** | Pointer to **int32** |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 
**AgeHours** | Pointer to **float64** |  | [optional] 
**AgeMinutes** | Pointer to **float64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**IndexerId** | Pointer to **int32** |  | [optional] 
**Indexer** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**SubGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Discography** | Pointer to **bool** |  | [optional] 
**SceneSource** | Pointer to **bool** |  | [optional] 
**AirDate** | Pointer to **NullableString** |  | [optional] 
**AuthorName** | Pointer to **NullableString** |  | [optional] 
**BookTitle** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**TemporarilyRejected** | Pointer to **bool** |  | [optional] 
**Rejected** | Pointer to **bool** |  | [optional] 
**Rejections** | Pointer to **[]string** |  | [optional] 
**PublishDate** | Pointer to **time.Time** |  | [optional] 
**CommentUrl** | Pointer to **NullableString** |  | [optional] 
**DownloadUrl** | Pointer to **NullableString** |  | [optional] 
**InfoUrl** | Pointer to **NullableString** |  | [optional] 
**DownloadAllowed** | Pointer to **bool** |  | [optional] 
**ReleaseWeight** | Pointer to **int32** |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**CustomFormatScore** | Pointer to **int32** |  | [optional] 
**MagnetUrl** | Pointer to **NullableString** |  | [optional] 
**InfoHash** | Pointer to **NullableString** |  | [optional] 
**Seeders** | Pointer to **NullableInt32** |  | [optional] 
**Leechers** | Pointer to **NullableInt32** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**AuthorId** | Pointer to **NullableInt32** |  | [optional] 
**BookId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewReleaseResource

`func NewReleaseResource() *ReleaseResource`

NewReleaseResource instantiates a new ReleaseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseResourceWithDefaults

`func NewReleaseResourceWithDefaults() *ReleaseResource`

NewReleaseResourceWithDefaults instantiates a new ReleaseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGuid

`func (o *ReleaseResource) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ReleaseResource) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ReleaseResource) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ReleaseResource) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *ReleaseResource) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *ReleaseResource) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetQuality

`func (o *ReleaseResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ReleaseResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ReleaseResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ReleaseResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetQualityWeight

`func (o *ReleaseResource) GetQualityWeight() int32`

GetQualityWeight returns the QualityWeight field if non-nil, zero value otherwise.

### GetQualityWeightOk

`func (o *ReleaseResource) GetQualityWeightOk() (*int32, bool)`

GetQualityWeightOk returns a tuple with the QualityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityWeight

`func (o *ReleaseResource) SetQualityWeight(v int32)`

SetQualityWeight sets QualityWeight field to given value.

### HasQualityWeight

`func (o *ReleaseResource) HasQualityWeight() bool`

HasQualityWeight returns a boolean if a field has been set.

### GetAge

`func (o *ReleaseResource) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ReleaseResource) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ReleaseResource) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ReleaseResource) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAgeHours

`func (o *ReleaseResource) GetAgeHours() float64`

GetAgeHours returns the AgeHours field if non-nil, zero value otherwise.

### GetAgeHoursOk

`func (o *ReleaseResource) GetAgeHoursOk() (*float64, bool)`

GetAgeHoursOk returns a tuple with the AgeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeHours

`func (o *ReleaseResource) SetAgeHours(v float64)`

SetAgeHours sets AgeHours field to given value.

### HasAgeHours

`func (o *ReleaseResource) HasAgeHours() bool`

HasAgeHours returns a boolean if a field has been set.

### GetAgeMinutes

`func (o *ReleaseResource) GetAgeMinutes() float64`

GetAgeMinutes returns the AgeMinutes field if non-nil, zero value otherwise.

### GetAgeMinutesOk

`func (o *ReleaseResource) GetAgeMinutesOk() (*float64, bool)`

GetAgeMinutesOk returns a tuple with the AgeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMinutes

`func (o *ReleaseResource) SetAgeMinutes(v float64)`

SetAgeMinutes sets AgeMinutes field to given value.

### HasAgeMinutes

`func (o *ReleaseResource) HasAgeMinutes() bool`

HasAgeMinutes returns a boolean if a field has been set.

### GetSize

`func (o *ReleaseResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReleaseResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReleaseResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ReleaseResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIndexerId

`func (o *ReleaseResource) GetIndexerId() int32`

GetIndexerId returns the IndexerId field if non-nil, zero value otherwise.

### GetIndexerIdOk

`func (o *ReleaseResource) GetIndexerIdOk() (*int32, bool)`

GetIndexerIdOk returns a tuple with the IndexerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerId

`func (o *ReleaseResource) SetIndexerId(v int32)`

SetIndexerId sets IndexerId field to given value.

### HasIndexerId

`func (o *ReleaseResource) HasIndexerId() bool`

HasIndexerId returns a boolean if a field has been set.

### GetIndexer

`func (o *ReleaseResource) GetIndexer() string`

GetIndexer returns the Indexer field if non-nil, zero value otherwise.

### GetIndexerOk

`func (o *ReleaseResource) GetIndexerOk() (*string, bool)`

GetIndexerOk returns a tuple with the Indexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexer

`func (o *ReleaseResource) SetIndexer(v string)`

SetIndexer sets Indexer field to given value.

### HasIndexer

`func (o *ReleaseResource) HasIndexer() bool`

HasIndexer returns a boolean if a field has been set.

### SetIndexerNil

`func (o *ReleaseResource) SetIndexerNil(b bool)`

 SetIndexerNil sets the value for Indexer to be an explicit nil

### UnsetIndexer
`func (o *ReleaseResource) UnsetIndexer()`

UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
### GetReleaseGroup

`func (o *ReleaseResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ReleaseResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ReleaseResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ReleaseResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ReleaseResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ReleaseResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetSubGroup

`func (o *ReleaseResource) GetSubGroup() string`

GetSubGroup returns the SubGroup field if non-nil, zero value otherwise.

### GetSubGroupOk

`func (o *ReleaseResource) GetSubGroupOk() (*string, bool)`

GetSubGroupOk returns a tuple with the SubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGroup

`func (o *ReleaseResource) SetSubGroup(v string)`

SetSubGroup sets SubGroup field to given value.

### HasSubGroup

`func (o *ReleaseResource) HasSubGroup() bool`

HasSubGroup returns a boolean if a field has been set.

### SetSubGroupNil

`func (o *ReleaseResource) SetSubGroupNil(b bool)`

 SetSubGroupNil sets the value for SubGroup to be an explicit nil

### UnsetSubGroup
`func (o *ReleaseResource) UnsetSubGroup()`

UnsetSubGroup ensures that no value is present for SubGroup, not even an explicit nil
### GetReleaseHash

`func (o *ReleaseResource) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ReleaseResource) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ReleaseResource) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ReleaseResource) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ReleaseResource) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ReleaseResource) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
### GetTitle

`func (o *ReleaseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ReleaseResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ReleaseResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDiscography

`func (o *ReleaseResource) GetDiscography() bool`

GetDiscography returns the Discography field if non-nil, zero value otherwise.

### GetDiscographyOk

`func (o *ReleaseResource) GetDiscographyOk() (*bool, bool)`

GetDiscographyOk returns a tuple with the Discography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscography

`func (o *ReleaseResource) SetDiscography(v bool)`

SetDiscography sets Discography field to given value.

### HasDiscography

`func (o *ReleaseResource) HasDiscography() bool`

HasDiscography returns a boolean if a field has been set.

### GetSceneSource

`func (o *ReleaseResource) GetSceneSource() bool`

GetSceneSource returns the SceneSource field if non-nil, zero value otherwise.

### GetSceneSourceOk

`func (o *ReleaseResource) GetSceneSourceOk() (*bool, bool)`

GetSceneSourceOk returns a tuple with the SceneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneSource

`func (o *ReleaseResource) SetSceneSource(v bool)`

SetSceneSource sets SceneSource field to given value.

### HasSceneSource

`func (o *ReleaseResource) HasSceneSource() bool`

HasSceneSource returns a boolean if a field has been set.

### GetAirDate

`func (o *ReleaseResource) GetAirDate() string`

GetAirDate returns the AirDate field if non-nil, zero value otherwise.

### GetAirDateOk

`func (o *ReleaseResource) GetAirDateOk() (*string, bool)`

GetAirDateOk returns a tuple with the AirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDate

`func (o *ReleaseResource) SetAirDate(v string)`

SetAirDate sets AirDate field to given value.

### HasAirDate

`func (o *ReleaseResource) HasAirDate() bool`

HasAirDate returns a boolean if a field has been set.

### SetAirDateNil

`func (o *ReleaseResource) SetAirDateNil(b bool)`

 SetAirDateNil sets the value for AirDate to be an explicit nil

### UnsetAirDate
`func (o *ReleaseResource) UnsetAirDate()`

UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
### GetAuthorName

`func (o *ReleaseResource) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *ReleaseResource) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *ReleaseResource) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *ReleaseResource) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *ReleaseResource) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *ReleaseResource) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetBookTitle

`func (o *ReleaseResource) GetBookTitle() string`

GetBookTitle returns the BookTitle field if non-nil, zero value otherwise.

### GetBookTitleOk

`func (o *ReleaseResource) GetBookTitleOk() (*string, bool)`

GetBookTitleOk returns a tuple with the BookTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTitle

`func (o *ReleaseResource) SetBookTitle(v string)`

SetBookTitle sets BookTitle field to given value.

### HasBookTitle

`func (o *ReleaseResource) HasBookTitle() bool`

HasBookTitle returns a boolean if a field has been set.

### SetBookTitleNil

`func (o *ReleaseResource) SetBookTitleNil(b bool)`

 SetBookTitleNil sets the value for BookTitle to be an explicit nil

### UnsetBookTitle
`func (o *ReleaseResource) UnsetBookTitle()`

UnsetBookTitle ensures that no value is present for BookTitle, not even an explicit nil
### GetApproved

`func (o *ReleaseResource) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ReleaseResource) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ReleaseResource) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *ReleaseResource) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetTemporarilyRejected

`func (o *ReleaseResource) GetTemporarilyRejected() bool`

GetTemporarilyRejected returns the TemporarilyRejected field if non-nil, zero value otherwise.

### GetTemporarilyRejectedOk

`func (o *ReleaseResource) GetTemporarilyRejectedOk() (*bool, bool)`

GetTemporarilyRejectedOk returns a tuple with the TemporarilyRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarilyRejected

`func (o *ReleaseResource) SetTemporarilyRejected(v bool)`

SetTemporarilyRejected sets TemporarilyRejected field to given value.

### HasTemporarilyRejected

`func (o *ReleaseResource) HasTemporarilyRejected() bool`

HasTemporarilyRejected returns a boolean if a field has been set.

### GetRejected

`func (o *ReleaseResource) GetRejected() bool`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *ReleaseResource) GetRejectedOk() (*bool, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *ReleaseResource) SetRejected(v bool)`

SetRejected sets Rejected field to given value.

### HasRejected

`func (o *ReleaseResource) HasRejected() bool`

HasRejected returns a boolean if a field has been set.

### GetRejections

`func (o *ReleaseResource) GetRejections() []string`

GetRejections returns the Rejections field if non-nil, zero value otherwise.

### GetRejectionsOk

`func (o *ReleaseResource) GetRejectionsOk() (*[]string, bool)`

GetRejectionsOk returns a tuple with the Rejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejections

`func (o *ReleaseResource) SetRejections(v []string)`

SetRejections sets Rejections field to given value.

### HasRejections

`func (o *ReleaseResource) HasRejections() bool`

HasRejections returns a boolean if a field has been set.

### SetRejectionsNil

`func (o *ReleaseResource) SetRejectionsNil(b bool)`

 SetRejectionsNil sets the value for Rejections to be an explicit nil

### UnsetRejections
`func (o *ReleaseResource) UnsetRejections()`

UnsetRejections ensures that no value is present for Rejections, not even an explicit nil
### GetPublishDate

`func (o *ReleaseResource) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *ReleaseResource) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *ReleaseResource) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *ReleaseResource) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### GetCommentUrl

`func (o *ReleaseResource) GetCommentUrl() string`

GetCommentUrl returns the CommentUrl field if non-nil, zero value otherwise.

### GetCommentUrlOk

`func (o *ReleaseResource) GetCommentUrlOk() (*string, bool)`

GetCommentUrlOk returns a tuple with the CommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentUrl

`func (o *ReleaseResource) SetCommentUrl(v string)`

SetCommentUrl sets CommentUrl field to given value.

### HasCommentUrl

`func (o *ReleaseResource) HasCommentUrl() bool`

HasCommentUrl returns a boolean if a field has been set.

### SetCommentUrlNil

`func (o *ReleaseResource) SetCommentUrlNil(b bool)`

 SetCommentUrlNil sets the value for CommentUrl to be an explicit nil

### UnsetCommentUrl
`func (o *ReleaseResource) UnsetCommentUrl()`

UnsetCommentUrl ensures that no value is present for CommentUrl, not even an explicit nil
### GetDownloadUrl

`func (o *ReleaseResource) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ReleaseResource) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ReleaseResource) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *ReleaseResource) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### SetDownloadUrlNil

`func (o *ReleaseResource) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *ReleaseResource) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetInfoUrl

`func (o *ReleaseResource) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *ReleaseResource) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *ReleaseResource) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *ReleaseResource) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### SetInfoUrlNil

`func (o *ReleaseResource) SetInfoUrlNil(b bool)`

 SetInfoUrlNil sets the value for InfoUrl to be an explicit nil

### UnsetInfoUrl
`func (o *ReleaseResource) UnsetInfoUrl()`

UnsetInfoUrl ensures that no value is present for InfoUrl, not even an explicit nil
### GetDownloadAllowed

`func (o *ReleaseResource) GetDownloadAllowed() bool`

GetDownloadAllowed returns the DownloadAllowed field if non-nil, zero value otherwise.

### GetDownloadAllowedOk

`func (o *ReleaseResource) GetDownloadAllowedOk() (*bool, bool)`

GetDownloadAllowedOk returns a tuple with the DownloadAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadAllowed

`func (o *ReleaseResource) SetDownloadAllowed(v bool)`

SetDownloadAllowed sets DownloadAllowed field to given value.

### HasDownloadAllowed

`func (o *ReleaseResource) HasDownloadAllowed() bool`

HasDownloadAllowed returns a boolean if a field has been set.

### GetReleaseWeight

`func (o *ReleaseResource) GetReleaseWeight() int32`

GetReleaseWeight returns the ReleaseWeight field if non-nil, zero value otherwise.

### GetReleaseWeightOk

`func (o *ReleaseResource) GetReleaseWeightOk() (*int32, bool)`

GetReleaseWeightOk returns a tuple with the ReleaseWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseWeight

`func (o *ReleaseResource) SetReleaseWeight(v int32)`

SetReleaseWeight sets ReleaseWeight field to given value.

### HasReleaseWeight

`func (o *ReleaseResource) HasReleaseWeight() bool`

HasReleaseWeight returns a boolean if a field has been set.

### GetCustomFormats

`func (o *ReleaseResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *ReleaseResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *ReleaseResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *ReleaseResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *ReleaseResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *ReleaseResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetCustomFormatScore

`func (o *ReleaseResource) GetCustomFormatScore() int32`

GetCustomFormatScore returns the CustomFormatScore field if non-nil, zero value otherwise.

### GetCustomFormatScoreOk

`func (o *ReleaseResource) GetCustomFormatScoreOk() (*int32, bool)`

GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormatScore

`func (o *ReleaseResource) SetCustomFormatScore(v int32)`

SetCustomFormatScore sets CustomFormatScore field to given value.

### HasCustomFormatScore

`func (o *ReleaseResource) HasCustomFormatScore() bool`

HasCustomFormatScore returns a boolean if a field has been set.

### GetMagnetUrl

`func (o *ReleaseResource) GetMagnetUrl() string`

GetMagnetUrl returns the MagnetUrl field if non-nil, zero value otherwise.

### GetMagnetUrlOk

`func (o *ReleaseResource) GetMagnetUrlOk() (*string, bool)`

GetMagnetUrlOk returns a tuple with the MagnetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetUrl

`func (o *ReleaseResource) SetMagnetUrl(v string)`

SetMagnetUrl sets MagnetUrl field to given value.

### HasMagnetUrl

`func (o *ReleaseResource) HasMagnetUrl() bool`

HasMagnetUrl returns a boolean if a field has been set.

### SetMagnetUrlNil

`func (o *ReleaseResource) SetMagnetUrlNil(b bool)`

 SetMagnetUrlNil sets the value for MagnetUrl to be an explicit nil

### UnsetMagnetUrl
`func (o *ReleaseResource) UnsetMagnetUrl()`

UnsetMagnetUrl ensures that no value is present for MagnetUrl, not even an explicit nil
### GetInfoHash

`func (o *ReleaseResource) GetInfoHash() string`

GetInfoHash returns the InfoHash field if non-nil, zero value otherwise.

### GetInfoHashOk

`func (o *ReleaseResource) GetInfoHashOk() (*string, bool)`

GetInfoHashOk returns a tuple with the InfoHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoHash

`func (o *ReleaseResource) SetInfoHash(v string)`

SetInfoHash sets InfoHash field to given value.

### HasInfoHash

`func (o *ReleaseResource) HasInfoHash() bool`

HasInfoHash returns a boolean if a field has been set.

### SetInfoHashNil

`func (o *ReleaseResource) SetInfoHashNil(b bool)`

 SetInfoHashNil sets the value for InfoHash to be an explicit nil

### UnsetInfoHash
`func (o *ReleaseResource) UnsetInfoHash()`

UnsetInfoHash ensures that no value is present for InfoHash, not even an explicit nil
### GetSeeders

`func (o *ReleaseResource) GetSeeders() int32`

GetSeeders returns the Seeders field if non-nil, zero value otherwise.

### GetSeedersOk

`func (o *ReleaseResource) GetSeedersOk() (*int32, bool)`

GetSeedersOk returns a tuple with the Seeders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeders

`func (o *ReleaseResource) SetSeeders(v int32)`

SetSeeders sets Seeders field to given value.

### HasSeeders

`func (o *ReleaseResource) HasSeeders() bool`

HasSeeders returns a boolean if a field has been set.

### SetSeedersNil

`func (o *ReleaseResource) SetSeedersNil(b bool)`

 SetSeedersNil sets the value for Seeders to be an explicit nil

### UnsetSeeders
`func (o *ReleaseResource) UnsetSeeders()`

UnsetSeeders ensures that no value is present for Seeders, not even an explicit nil
### GetLeechers

`func (o *ReleaseResource) GetLeechers() int32`

GetLeechers returns the Leechers field if non-nil, zero value otherwise.

### GetLeechersOk

`func (o *ReleaseResource) GetLeechersOk() (*int32, bool)`

GetLeechersOk returns a tuple with the Leechers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeechers

`func (o *ReleaseResource) SetLeechers(v int32)`

SetLeechers sets Leechers field to given value.

### HasLeechers

`func (o *ReleaseResource) HasLeechers() bool`

HasLeechers returns a boolean if a field has been set.

### SetLeechersNil

`func (o *ReleaseResource) SetLeechersNil(b bool)`

 SetLeechersNil sets the value for Leechers to be an explicit nil

### UnsetLeechers
`func (o *ReleaseResource) UnsetLeechers()`

UnsetLeechers ensures that no value is present for Leechers, not even an explicit nil
### GetProtocol

`func (o *ReleaseResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ReleaseResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ReleaseResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ReleaseResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAuthorId

`func (o *ReleaseResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *ReleaseResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *ReleaseResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *ReleaseResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### SetAuthorIdNil

`func (o *ReleaseResource) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *ReleaseResource) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
### GetBookId

`func (o *ReleaseResource) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *ReleaseResource) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *ReleaseResource) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *ReleaseResource) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### SetBookIdNil

`func (o *ReleaseResource) SetBookIdNil(b bool)`

 SetBookIdNil sets the value for BookId to be an explicit nil

### UnsetBookId
`func (o *ReleaseResource) UnsetBookId()`

UnsetBookId ensures that no value is present for BookId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


