# BookFileListResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookFileIds** | Pointer to **[]int32** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 

## Methods

### NewBookFileListResource

`func NewBookFileListResource() *BookFileListResource`

NewBookFileListResource instantiates a new BookFileListResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookFileListResourceWithDefaults

`func NewBookFileListResourceWithDefaults() *BookFileListResource`

NewBookFileListResourceWithDefaults instantiates a new BookFileListResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookFileIds

`func (o *BookFileListResource) GetBookFileIds() []int32`

GetBookFileIds returns the BookFileIds field if non-nil, zero value otherwise.

### GetBookFileIdsOk

`func (o *BookFileListResource) GetBookFileIdsOk() (*[]int32, bool)`

GetBookFileIdsOk returns a tuple with the BookFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookFileIds

`func (o *BookFileListResource) SetBookFileIds(v []int32)`

SetBookFileIds sets BookFileIds field to given value.

### HasBookFileIds

`func (o *BookFileListResource) HasBookFileIds() bool`

HasBookFileIds returns a boolean if a field has been set.

### SetBookFileIdsNil

`func (o *BookFileListResource) SetBookFileIdsNil(b bool)`

 SetBookFileIdsNil sets the value for BookFileIds to be an explicit nil

### UnsetBookFileIds
`func (o *BookFileListResource) UnsetBookFileIds()`

UnsetBookFileIds ensures that no value is present for BookFileIds, not even an explicit nil
### GetQuality

`func (o *BookFileListResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BookFileListResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BookFileListResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BookFileListResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


