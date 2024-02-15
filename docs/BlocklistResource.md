# BlocklistResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **int32** |  | [optional] 
**BookIds** | Pointer to **[]int32** |  | [optional] 
**SourceTitle** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**Indexer** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to [**AuthorResource**](AuthorResource.md) |  | [optional] 

## Methods

### NewBlocklistResource

`func NewBlocklistResource() *BlocklistResource`

NewBlocklistResource instantiates a new BlocklistResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistResourceWithDefaults

`func NewBlocklistResourceWithDefaults() *BlocklistResource`

NewBlocklistResourceWithDefaults instantiates a new BlocklistResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlocklistResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlocklistResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlocklistResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BlocklistResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorId

`func (o *BlocklistResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *BlocklistResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *BlocklistResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *BlocklistResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetBookIds

`func (o *BlocklistResource) GetBookIds() []int32`

GetBookIds returns the BookIds field if non-nil, zero value otherwise.

### GetBookIdsOk

`func (o *BlocklistResource) GetBookIdsOk() (*[]int32, bool)`

GetBookIdsOk returns a tuple with the BookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookIds

`func (o *BlocklistResource) SetBookIds(v []int32)`

SetBookIds sets BookIds field to given value.

### HasBookIds

`func (o *BlocklistResource) HasBookIds() bool`

HasBookIds returns a boolean if a field has been set.

### SetBookIdsNil

`func (o *BlocklistResource) SetBookIdsNil(b bool)`

 SetBookIdsNil sets the value for BookIds to be an explicit nil

### UnsetBookIds
`func (o *BlocklistResource) UnsetBookIds()`

UnsetBookIds ensures that no value is present for BookIds, not even an explicit nil
### GetSourceTitle

`func (o *BlocklistResource) GetSourceTitle() string`

GetSourceTitle returns the SourceTitle field if non-nil, zero value otherwise.

### GetSourceTitleOk

`func (o *BlocklistResource) GetSourceTitleOk() (*string, bool)`

GetSourceTitleOk returns a tuple with the SourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTitle

`func (o *BlocklistResource) SetSourceTitle(v string)`

SetSourceTitle sets SourceTitle field to given value.

### HasSourceTitle

`func (o *BlocklistResource) HasSourceTitle() bool`

HasSourceTitle returns a boolean if a field has been set.

### SetSourceTitleNil

`func (o *BlocklistResource) SetSourceTitleNil(b bool)`

 SetSourceTitleNil sets the value for SourceTitle to be an explicit nil

### UnsetSourceTitle
`func (o *BlocklistResource) UnsetSourceTitle()`

UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
### GetQuality

`func (o *BlocklistResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BlocklistResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BlocklistResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BlocklistResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCustomFormats

`func (o *BlocklistResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *BlocklistResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *BlocklistResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *BlocklistResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *BlocklistResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *BlocklistResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetDate

`func (o *BlocklistResource) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BlocklistResource) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BlocklistResource) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BlocklistResource) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetProtocol

`func (o *BlocklistResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BlocklistResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BlocklistResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BlocklistResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetIndexer

`func (o *BlocklistResource) GetIndexer() string`

GetIndexer returns the Indexer field if non-nil, zero value otherwise.

### GetIndexerOk

`func (o *BlocklistResource) GetIndexerOk() (*string, bool)`

GetIndexerOk returns a tuple with the Indexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexer

`func (o *BlocklistResource) SetIndexer(v string)`

SetIndexer sets Indexer field to given value.

### HasIndexer

`func (o *BlocklistResource) HasIndexer() bool`

HasIndexer returns a boolean if a field has been set.

### SetIndexerNil

`func (o *BlocklistResource) SetIndexerNil(b bool)`

 SetIndexerNil sets the value for Indexer to be an explicit nil

### UnsetIndexer
`func (o *BlocklistResource) UnsetIndexer()`

UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
### GetMessage

`func (o *BlocklistResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BlocklistResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BlocklistResource) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BlocklistResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BlocklistResource) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BlocklistResource) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetAuthor

`func (o *BlocklistResource) GetAuthor() AuthorResource`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BlocklistResource) GetAuthorOk() (*AuthorResource, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BlocklistResource) SetAuthor(v AuthorResource)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *BlocklistResource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


