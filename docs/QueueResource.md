# QueueResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **NullableInt32** |  | [optional] 
**BookId** | Pointer to **NullableInt32** |  | [optional] 
**Author** | Pointer to [**AuthorResource**](AuthorResource.md) |  | [optional] 
**Book** | Pointer to [**BookResource**](BookResource.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**CustomFormatScore** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **float64** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Sizeleft** | Pointer to **float64** |  | [optional] 
**Timeleft** | Pointer to **NullableString** |  | [optional] 
**EstimatedCompletionTime** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**TrackedDownloadStatus** | Pointer to [**TrackedDownloadStatus**](TrackedDownloadStatus.md) |  | [optional] 
**TrackedDownloadState** | Pointer to [**TrackedDownloadState**](TrackedDownloadState.md) |  | [optional] 
**StatusMessages** | Pointer to [**[]TrackedDownloadStatusMessage**](TrackedDownloadStatusMessage.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**DownloadClient** | Pointer to **NullableString** |  | [optional] 
**DownloadClientHasPostImportCategory** | Pointer to **bool** |  | [optional] 
**Indexer** | Pointer to **NullableString** |  | [optional] 
**OutputPath** | Pointer to **NullableString** |  | [optional] 
**DownloadForced** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueueResource

`func NewQueueResource() *QueueResource`

NewQueueResource instantiates a new QueueResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueResourceWithDefaults

`func NewQueueResourceWithDefaults() *QueueResource`

NewQueueResourceWithDefaults instantiates a new QueueResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueueResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QueueResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorId

`func (o *QueueResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *QueueResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *QueueResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *QueueResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### SetAuthorIdNil

`func (o *QueueResource) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *QueueResource) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
### GetBookId

`func (o *QueueResource) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *QueueResource) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *QueueResource) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *QueueResource) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### SetBookIdNil

`func (o *QueueResource) SetBookIdNil(b bool)`

 SetBookIdNil sets the value for BookId to be an explicit nil

### UnsetBookId
`func (o *QueueResource) UnsetBookId()`

UnsetBookId ensures that no value is present for BookId, not even an explicit nil
### GetAuthor

`func (o *QueueResource) GetAuthor() AuthorResource`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *QueueResource) GetAuthorOk() (*AuthorResource, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *QueueResource) SetAuthor(v AuthorResource)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *QueueResource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBook

`func (o *QueueResource) GetBook() BookResource`

GetBook returns the Book field if non-nil, zero value otherwise.

### GetBookOk

`func (o *QueueResource) GetBookOk() (*BookResource, bool)`

GetBookOk returns a tuple with the Book field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBook

`func (o *QueueResource) SetBook(v BookResource)`

SetBook sets Book field to given value.

### HasBook

`func (o *QueueResource) HasBook() bool`

HasBook returns a boolean if a field has been set.

### GetQuality

`func (o *QueueResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *QueueResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *QueueResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *QueueResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCustomFormats

`func (o *QueueResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *QueueResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *QueueResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *QueueResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *QueueResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *QueueResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetCustomFormatScore

`func (o *QueueResource) GetCustomFormatScore() int32`

GetCustomFormatScore returns the CustomFormatScore field if non-nil, zero value otherwise.

### GetCustomFormatScoreOk

`func (o *QueueResource) GetCustomFormatScoreOk() (*int32, bool)`

GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormatScore

`func (o *QueueResource) SetCustomFormatScore(v int32)`

SetCustomFormatScore sets CustomFormatScore field to given value.

### HasCustomFormatScore

`func (o *QueueResource) HasCustomFormatScore() bool`

HasCustomFormatScore returns a boolean if a field has been set.

### GetSize

`func (o *QueueResource) GetSize() float64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *QueueResource) GetSizeOk() (*float64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *QueueResource) SetSize(v float64)`

SetSize sets Size field to given value.

### HasSize

`func (o *QueueResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTitle

`func (o *QueueResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QueueResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QueueResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QueueResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *QueueResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *QueueResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSizeleft

`func (o *QueueResource) GetSizeleft() float64`

GetSizeleft returns the Sizeleft field if non-nil, zero value otherwise.

### GetSizeleftOk

`func (o *QueueResource) GetSizeleftOk() (*float64, bool)`

GetSizeleftOk returns a tuple with the Sizeleft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeleft

`func (o *QueueResource) SetSizeleft(v float64)`

SetSizeleft sets Sizeleft field to given value.

### HasSizeleft

`func (o *QueueResource) HasSizeleft() bool`

HasSizeleft returns a boolean if a field has been set.

### GetTimeleft

`func (o *QueueResource) GetTimeleft() string`

GetTimeleft returns the Timeleft field if non-nil, zero value otherwise.

### GetTimeleftOk

`func (o *QueueResource) GetTimeleftOk() (*string, bool)`

GetTimeleftOk returns a tuple with the Timeleft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeleft

`func (o *QueueResource) SetTimeleft(v string)`

SetTimeleft sets Timeleft field to given value.

### HasTimeleft

`func (o *QueueResource) HasTimeleft() bool`

HasTimeleft returns a boolean if a field has been set.

### SetTimeleftNil

`func (o *QueueResource) SetTimeleftNil(b bool)`

 SetTimeleftNil sets the value for Timeleft to be an explicit nil

### UnsetTimeleft
`func (o *QueueResource) UnsetTimeleft()`

UnsetTimeleft ensures that no value is present for Timeleft, not even an explicit nil
### GetEstimatedCompletionTime

`func (o *QueueResource) GetEstimatedCompletionTime() time.Time`

GetEstimatedCompletionTime returns the EstimatedCompletionTime field if non-nil, zero value otherwise.

### GetEstimatedCompletionTimeOk

`func (o *QueueResource) GetEstimatedCompletionTimeOk() (*time.Time, bool)`

GetEstimatedCompletionTimeOk returns a tuple with the EstimatedCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionTime

`func (o *QueueResource) SetEstimatedCompletionTime(v time.Time)`

SetEstimatedCompletionTime sets EstimatedCompletionTime field to given value.

### HasEstimatedCompletionTime

`func (o *QueueResource) HasEstimatedCompletionTime() bool`

HasEstimatedCompletionTime returns a boolean if a field has been set.

### SetEstimatedCompletionTimeNil

`func (o *QueueResource) SetEstimatedCompletionTimeNil(b bool)`

 SetEstimatedCompletionTimeNil sets the value for EstimatedCompletionTime to be an explicit nil

### UnsetEstimatedCompletionTime
`func (o *QueueResource) UnsetEstimatedCompletionTime()`

UnsetEstimatedCompletionTime ensures that no value is present for EstimatedCompletionTime, not even an explicit nil
### GetStatus

`func (o *QueueResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueueResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueueResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueueResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *QueueResource) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *QueueResource) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTrackedDownloadStatus

`func (o *QueueResource) GetTrackedDownloadStatus() TrackedDownloadStatus`

GetTrackedDownloadStatus returns the TrackedDownloadStatus field if non-nil, zero value otherwise.

### GetTrackedDownloadStatusOk

`func (o *QueueResource) GetTrackedDownloadStatusOk() (*TrackedDownloadStatus, bool)`

GetTrackedDownloadStatusOk returns a tuple with the TrackedDownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedDownloadStatus

`func (o *QueueResource) SetTrackedDownloadStatus(v TrackedDownloadStatus)`

SetTrackedDownloadStatus sets TrackedDownloadStatus field to given value.

### HasTrackedDownloadStatus

`func (o *QueueResource) HasTrackedDownloadStatus() bool`

HasTrackedDownloadStatus returns a boolean if a field has been set.

### GetTrackedDownloadState

`func (o *QueueResource) GetTrackedDownloadState() TrackedDownloadState`

GetTrackedDownloadState returns the TrackedDownloadState field if non-nil, zero value otherwise.

### GetTrackedDownloadStateOk

`func (o *QueueResource) GetTrackedDownloadStateOk() (*TrackedDownloadState, bool)`

GetTrackedDownloadStateOk returns a tuple with the TrackedDownloadState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedDownloadState

`func (o *QueueResource) SetTrackedDownloadState(v TrackedDownloadState)`

SetTrackedDownloadState sets TrackedDownloadState field to given value.

### HasTrackedDownloadState

`func (o *QueueResource) HasTrackedDownloadState() bool`

HasTrackedDownloadState returns a boolean if a field has been set.

### GetStatusMessages

`func (o *QueueResource) GetStatusMessages() []TrackedDownloadStatusMessage`

GetStatusMessages returns the StatusMessages field if non-nil, zero value otherwise.

### GetStatusMessagesOk

`func (o *QueueResource) GetStatusMessagesOk() (*[]TrackedDownloadStatusMessage, bool)`

GetStatusMessagesOk returns a tuple with the StatusMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessages

`func (o *QueueResource) SetStatusMessages(v []TrackedDownloadStatusMessage)`

SetStatusMessages sets StatusMessages field to given value.

### HasStatusMessages

`func (o *QueueResource) HasStatusMessages() bool`

HasStatusMessages returns a boolean if a field has been set.

### SetStatusMessagesNil

`func (o *QueueResource) SetStatusMessagesNil(b bool)`

 SetStatusMessagesNil sets the value for StatusMessages to be an explicit nil

### UnsetStatusMessages
`func (o *QueueResource) UnsetStatusMessages()`

UnsetStatusMessages ensures that no value is present for StatusMessages, not even an explicit nil
### GetErrorMessage

`func (o *QueueResource) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *QueueResource) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *QueueResource) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *QueueResource) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *QueueResource) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *QueueResource) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDownloadId

`func (o *QueueResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *QueueResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *QueueResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *QueueResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *QueueResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *QueueResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetProtocol

`func (o *QueueResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *QueueResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *QueueResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *QueueResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDownloadClient

`func (o *QueueResource) GetDownloadClient() string`

GetDownloadClient returns the DownloadClient field if non-nil, zero value otherwise.

### GetDownloadClientOk

`func (o *QueueResource) GetDownloadClientOk() (*string, bool)`

GetDownloadClientOk returns a tuple with the DownloadClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadClient

`func (o *QueueResource) SetDownloadClient(v string)`

SetDownloadClient sets DownloadClient field to given value.

### HasDownloadClient

`func (o *QueueResource) HasDownloadClient() bool`

HasDownloadClient returns a boolean if a field has been set.

### SetDownloadClientNil

`func (o *QueueResource) SetDownloadClientNil(b bool)`

 SetDownloadClientNil sets the value for DownloadClient to be an explicit nil

### UnsetDownloadClient
`func (o *QueueResource) UnsetDownloadClient()`

UnsetDownloadClient ensures that no value is present for DownloadClient, not even an explicit nil
### GetDownloadClientHasPostImportCategory

`func (o *QueueResource) GetDownloadClientHasPostImportCategory() bool`

GetDownloadClientHasPostImportCategory returns the DownloadClientHasPostImportCategory field if non-nil, zero value otherwise.

### GetDownloadClientHasPostImportCategoryOk

`func (o *QueueResource) GetDownloadClientHasPostImportCategoryOk() (*bool, bool)`

GetDownloadClientHasPostImportCategoryOk returns a tuple with the DownloadClientHasPostImportCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadClientHasPostImportCategory

`func (o *QueueResource) SetDownloadClientHasPostImportCategory(v bool)`

SetDownloadClientHasPostImportCategory sets DownloadClientHasPostImportCategory field to given value.

### HasDownloadClientHasPostImportCategory

`func (o *QueueResource) HasDownloadClientHasPostImportCategory() bool`

HasDownloadClientHasPostImportCategory returns a boolean if a field has been set.

### GetIndexer

`func (o *QueueResource) GetIndexer() string`

GetIndexer returns the Indexer field if non-nil, zero value otherwise.

### GetIndexerOk

`func (o *QueueResource) GetIndexerOk() (*string, bool)`

GetIndexerOk returns a tuple with the Indexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexer

`func (o *QueueResource) SetIndexer(v string)`

SetIndexer sets Indexer field to given value.

### HasIndexer

`func (o *QueueResource) HasIndexer() bool`

HasIndexer returns a boolean if a field has been set.

### SetIndexerNil

`func (o *QueueResource) SetIndexerNil(b bool)`

 SetIndexerNil sets the value for Indexer to be an explicit nil

### UnsetIndexer
`func (o *QueueResource) UnsetIndexer()`

UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
### GetOutputPath

`func (o *QueueResource) GetOutputPath() string`

GetOutputPath returns the OutputPath field if non-nil, zero value otherwise.

### GetOutputPathOk

`func (o *QueueResource) GetOutputPathOk() (*string, bool)`

GetOutputPathOk returns a tuple with the OutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPath

`func (o *QueueResource) SetOutputPath(v string)`

SetOutputPath sets OutputPath field to given value.

### HasOutputPath

`func (o *QueueResource) HasOutputPath() bool`

HasOutputPath returns a boolean if a field has been set.

### SetOutputPathNil

`func (o *QueueResource) SetOutputPathNil(b bool)`

 SetOutputPathNil sets the value for OutputPath to be an explicit nil

### UnsetOutputPath
`func (o *QueueResource) UnsetOutputPath()`

UnsetOutputPath ensures that no value is present for OutputPath, not even an explicit nil
### GetDownloadForced

`func (o *QueueResource) GetDownloadForced() bool`

GetDownloadForced returns the DownloadForced field if non-nil, zero value otherwise.

### GetDownloadForcedOk

`func (o *QueueResource) GetDownloadForcedOk() (*bool, bool)`

GetDownloadForcedOk returns a tuple with the DownloadForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadForced

`func (o *QueueResource) SetDownloadForced(v bool)`

SetDownloadForced sets DownloadForced field to given value.

### HasDownloadForced

`func (o *QueueResource) HasDownloadForced() bool`

HasDownloadForced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


