# RetagBookResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **int32** |  | [optional] 
**BookId** | Pointer to **int32** |  | [optional] 
**TrackNumbers** | Pointer to **[]int32** |  | [optional] 
**BookFileId** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Changes** | Pointer to [**[]TagDifference**](TagDifference.md) |  | [optional] 

## Methods

### NewRetagBookResource

`func NewRetagBookResource() *RetagBookResource`

NewRetagBookResource instantiates a new RetagBookResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetagBookResourceWithDefaults

`func NewRetagBookResourceWithDefaults() *RetagBookResource`

NewRetagBookResourceWithDefaults instantiates a new RetagBookResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RetagBookResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RetagBookResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RetagBookResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RetagBookResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorId

`func (o *RetagBookResource) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *RetagBookResource) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *RetagBookResource) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *RetagBookResource) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetBookId

`func (o *RetagBookResource) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *RetagBookResource) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *RetagBookResource) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *RetagBookResource) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### GetTrackNumbers

`func (o *RetagBookResource) GetTrackNumbers() []int32`

GetTrackNumbers returns the TrackNumbers field if non-nil, zero value otherwise.

### GetTrackNumbersOk

`func (o *RetagBookResource) GetTrackNumbersOk() (*[]int32, bool)`

GetTrackNumbersOk returns a tuple with the TrackNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumbers

`func (o *RetagBookResource) SetTrackNumbers(v []int32)`

SetTrackNumbers sets TrackNumbers field to given value.

### HasTrackNumbers

`func (o *RetagBookResource) HasTrackNumbers() bool`

HasTrackNumbers returns a boolean if a field has been set.

### SetTrackNumbersNil

`func (o *RetagBookResource) SetTrackNumbersNil(b bool)`

 SetTrackNumbersNil sets the value for TrackNumbers to be an explicit nil

### UnsetTrackNumbers
`func (o *RetagBookResource) UnsetTrackNumbers()`

UnsetTrackNumbers ensures that no value is present for TrackNumbers, not even an explicit nil
### GetBookFileId

`func (o *RetagBookResource) GetBookFileId() int32`

GetBookFileId returns the BookFileId field if non-nil, zero value otherwise.

### GetBookFileIdOk

`func (o *RetagBookResource) GetBookFileIdOk() (*int32, bool)`

GetBookFileIdOk returns a tuple with the BookFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookFileId

`func (o *RetagBookResource) SetBookFileId(v int32)`

SetBookFileId sets BookFileId field to given value.

### HasBookFileId

`func (o *RetagBookResource) HasBookFileId() bool`

HasBookFileId returns a boolean if a field has been set.

### GetPath

`func (o *RetagBookResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RetagBookResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RetagBookResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RetagBookResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *RetagBookResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *RetagBookResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetChanges

`func (o *RetagBookResource) GetChanges() []TagDifference`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *RetagBookResource) GetChangesOk() (*[]TagDifference, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *RetagBookResource) SetChanges(v []TagDifference)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *RetagBookResource) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *RetagBookResource) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *RetagBookResource) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


