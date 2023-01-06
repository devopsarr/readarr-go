# BlacklistResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **int32** |  | [optional] 
**BookIds** | Pointer to **[]int32** |  | [optional] 
**SourceTitle** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**Indexer** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to [**AuthorResource**](AuthorResource.md) |  | [optional] 

## Methods

### NewBlacklistResource

`func NewBlacklistResource() *BlacklistResource`

NewBlacklistResource instantiates a new BlacklistResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlacklistResourceWithDefaults

`func NewBlacklistResourceWithDefaults() *BlacklistResource`

NewBlacklistResourceWithDefaults instantiates a new BlacklistResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlacklistResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlacklistResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlacklistResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BlacklistResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorId

`func (o *BlacklistResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *BlacklistResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *BlacklistResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *BlacklistResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetBookIds

`func (o *BlacklistResource) GetBookIds() []int32`

GetBookIds returns the BookIds field if non-nil, zero value otherwise.

### GetBookIdsOk

`func (o *BlacklistResource) GetBookIdsOk() (*[]int32, bool)`

GetBookIdsOk returns a tuple with the BookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookIds

`func (o *BlacklistResource) SetBookIds(v []int32)`

SetBookIds sets BookIds field to given value.

### HasBookIds

`func (o *BlacklistResource) HasBookIds() bool`

HasBookIds returns a boolean if a field has been set.

### SetBookIdsNil

`func (o *BlacklistResource) SetBookIdsNil(b bool)`

 SetBookIdsNil sets the value for BookIds to be an explicit nil

### UnsetBookIds
`func (o *BlacklistResource) UnsetBookIds()`

UnsetBookIds ensures that no value is present for BookIds, not even an explicit nil
### GetSourceTitle

`func (o *BlacklistResource) GetSourceTitle() string`

GetSourceTitle returns the SourceTitle field if non-nil, zero value otherwise.

### GetSourceTitleOk

`func (o *BlacklistResource) GetSourceTitleOk() (*string, bool)`

GetSourceTitleOk returns a tuple with the SourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTitle

`func (o *BlacklistResource) SetSourceTitle(v string)`

SetSourceTitle sets SourceTitle field to given value.

### HasSourceTitle

`func (o *BlacklistResource) HasSourceTitle() bool`

HasSourceTitle returns a boolean if a field has been set.

### SetSourceTitleNil

`func (o *BlacklistResource) SetSourceTitleNil(b bool)`

 SetSourceTitleNil sets the value for SourceTitle to be an explicit nil

### UnsetSourceTitle
`func (o *BlacklistResource) UnsetSourceTitle()`

UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
### GetQuality

`func (o *BlacklistResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BlacklistResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BlacklistResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BlacklistResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetDate

`func (o *BlacklistResource) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BlacklistResource) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BlacklistResource) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BlacklistResource) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetProtocol

`func (o *BlacklistResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BlacklistResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BlacklistResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BlacklistResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetIndexer

`func (o *BlacklistResource) GetIndexer() string`

GetIndexer returns the Indexer field if non-nil, zero value otherwise.

### GetIndexerOk

`func (o *BlacklistResource) GetIndexerOk() (*string, bool)`

GetIndexerOk returns a tuple with the Indexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexer

`func (o *BlacklistResource) SetIndexer(v string)`

SetIndexer sets Indexer field to given value.

### HasIndexer

`func (o *BlacklistResource) HasIndexer() bool`

HasIndexer returns a boolean if a field has been set.

### SetIndexerNil

`func (o *BlacklistResource) SetIndexerNil(b bool)`

 SetIndexerNil sets the value for Indexer to be an explicit nil

### UnsetIndexer
`func (o *BlacklistResource) UnsetIndexer()`

UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
### GetMessage

`func (o *BlacklistResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BlacklistResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BlacklistResource) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BlacklistResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BlacklistResource) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BlacklistResource) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetAuthor

`func (o *BlacklistResource) GetAuthor() AuthorResource`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BlacklistResource) GetAuthorOk() (*AuthorResource, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BlacklistResource) SetAuthor(v AuthorResource)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *BlacklistResource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


