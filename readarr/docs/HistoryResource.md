# HistoryResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BookId** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **int32** |  | [optional] 
**SourceTitle** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**QualityCutoffNotMet** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**EventType** | Pointer to [**EntityHistoryEventType**](EntityHistoryEventType.md) |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 
**Book** | Pointer to [**BookResource**](BookResource.md) |  | [optional] 
**Author** | Pointer to [**AuthorResource**](AuthorResource.md) |  | [optional] 

## Methods

### NewHistoryResource

`func NewHistoryResource() *HistoryResource`

NewHistoryResource instantiates a new HistoryResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryResourceWithDefaults

`func NewHistoryResourceWithDefaults() *HistoryResource`

NewHistoryResourceWithDefaults instantiates a new HistoryResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HistoryResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBookId

`func (o *HistoryResource) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *HistoryResource) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *HistoryResource) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *HistoryResource) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### GetAuthorId

`func (o *HistoryResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *HistoryResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *HistoryResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *HistoryResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetSourceTitle

`func (o *HistoryResource) GetSourceTitle() string`

GetSourceTitle returns the SourceTitle field if non-nil, zero value otherwise.

### GetSourceTitleOk

`func (o *HistoryResource) GetSourceTitleOk() (*string, bool)`

GetSourceTitleOk returns a tuple with the SourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTitle

`func (o *HistoryResource) SetSourceTitle(v string)`

SetSourceTitle sets SourceTitle field to given value.

### HasSourceTitle

`func (o *HistoryResource) HasSourceTitle() bool`

HasSourceTitle returns a boolean if a field has been set.

### SetSourceTitleNil

`func (o *HistoryResource) SetSourceTitleNil(b bool)`

 SetSourceTitleNil sets the value for SourceTitle to be an explicit nil

### UnsetSourceTitle
`func (o *HistoryResource) UnsetSourceTitle()`

UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
### GetQuality

`func (o *HistoryResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *HistoryResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *HistoryResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *HistoryResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCustomFormats

`func (o *HistoryResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *HistoryResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *HistoryResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *HistoryResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *HistoryResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *HistoryResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetQualityCutoffNotMet

`func (o *HistoryResource) GetQualityCutoffNotMet() bool`

GetQualityCutoffNotMet returns the QualityCutoffNotMet field if non-nil, zero value otherwise.

### GetQualityCutoffNotMetOk

`func (o *HistoryResource) GetQualityCutoffNotMetOk() (*bool, bool)`

GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityCutoffNotMet

`func (o *HistoryResource) SetQualityCutoffNotMet(v bool)`

SetQualityCutoffNotMet sets QualityCutoffNotMet field to given value.

### HasQualityCutoffNotMet

`func (o *HistoryResource) HasQualityCutoffNotMet() bool`

HasQualityCutoffNotMet returns a boolean if a field has been set.

### GetDate

`func (o *HistoryResource) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistoryResource) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistoryResource) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *HistoryResource) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDownloadId

`func (o *HistoryResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *HistoryResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *HistoryResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *HistoryResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *HistoryResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *HistoryResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetEventType

`func (o *HistoryResource) GetEventType() EntityHistoryEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *HistoryResource) GetEventTypeOk() (*EntityHistoryEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *HistoryResource) SetEventType(v EntityHistoryEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *HistoryResource) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetData

`func (o *HistoryResource) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryResource) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryResource) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryResource) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *HistoryResource) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *HistoryResource) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetBook

`func (o *HistoryResource) GetBook() BookResource`

GetBook returns the Book field if non-nil, zero value otherwise.

### GetBookOk

`func (o *HistoryResource) GetBookOk() (*BookResource, bool)`

GetBookOk returns a tuple with the Book field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBook

`func (o *HistoryResource) SetBook(v BookResource)`

SetBook sets Book field to given value.

### HasBook

`func (o *HistoryResource) HasBook() bool`

HasBook returns a boolean if a field has been set.

### GetAuthor

`func (o *HistoryResource) GetAuthor() AuthorResource`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *HistoryResource) GetAuthorOk() (*AuthorResource, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *HistoryResource) SetAuthor(v AuthorResource)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *HistoryResource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


