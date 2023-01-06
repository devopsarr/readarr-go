# TagDetailsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**DelayProfileIds** | Pointer to **[]int32** |  | [optional] 
**ImportListIds** | Pointer to **[]int32** |  | [optional] 
**NotificationIds** | Pointer to **[]int32** |  | [optional] 
**RestrictionIds** | Pointer to **[]int32** |  | [optional] 
**AuthorIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewTagDetailsResource

`func NewTagDetailsResource() *TagDetailsResource`

NewTagDetailsResource instantiates a new TagDetailsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagDetailsResourceWithDefaults

`func NewTagDetailsResourceWithDefaults() *TagDetailsResource`

NewTagDetailsResourceWithDefaults instantiates a new TagDetailsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagDetailsResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagDetailsResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagDetailsResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TagDetailsResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *TagDetailsResource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TagDetailsResource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TagDetailsResource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TagDetailsResource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *TagDetailsResource) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *TagDetailsResource) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDelayProfileIds

`func (o *TagDetailsResource) GetDelayProfileIds() []int32`

GetDelayProfileIds returns the DelayProfileIds field if non-nil, zero value otherwise.

### GetDelayProfileIdsOk

`func (o *TagDetailsResource) GetDelayProfileIdsOk() (*[]int32, bool)`

GetDelayProfileIdsOk returns a tuple with the DelayProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayProfileIds

`func (o *TagDetailsResource) SetDelayProfileIds(v []int32)`

SetDelayProfileIds sets DelayProfileIds field to given value.

### HasDelayProfileIds

`func (o *TagDetailsResource) HasDelayProfileIds() bool`

HasDelayProfileIds returns a boolean if a field has been set.

### SetDelayProfileIdsNil

`func (o *TagDetailsResource) SetDelayProfileIdsNil(b bool)`

 SetDelayProfileIdsNil sets the value for DelayProfileIds to be an explicit nil

### UnsetDelayProfileIds
`func (o *TagDetailsResource) UnsetDelayProfileIds()`

UnsetDelayProfileIds ensures that no value is present for DelayProfileIds, not even an explicit nil
### GetImportListIds

`func (o *TagDetailsResource) GetImportListIds() []int32`

GetImportListIds returns the ImportListIds field if non-nil, zero value otherwise.

### GetImportListIdsOk

`func (o *TagDetailsResource) GetImportListIdsOk() (*[]int32, bool)`

GetImportListIdsOk returns a tuple with the ImportListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportListIds

`func (o *TagDetailsResource) SetImportListIds(v []int32)`

SetImportListIds sets ImportListIds field to given value.

### HasImportListIds

`func (o *TagDetailsResource) HasImportListIds() bool`

HasImportListIds returns a boolean if a field has been set.

### SetImportListIdsNil

`func (o *TagDetailsResource) SetImportListIdsNil(b bool)`

 SetImportListIdsNil sets the value for ImportListIds to be an explicit nil

### UnsetImportListIds
`func (o *TagDetailsResource) UnsetImportListIds()`

UnsetImportListIds ensures that no value is present for ImportListIds, not even an explicit nil
### GetNotificationIds

`func (o *TagDetailsResource) GetNotificationIds() []int32`

GetNotificationIds returns the NotificationIds field if non-nil, zero value otherwise.

### GetNotificationIdsOk

`func (o *TagDetailsResource) GetNotificationIdsOk() (*[]int32, bool)`

GetNotificationIdsOk returns a tuple with the NotificationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIds

`func (o *TagDetailsResource) SetNotificationIds(v []int32)`

SetNotificationIds sets NotificationIds field to given value.

### HasNotificationIds

`func (o *TagDetailsResource) HasNotificationIds() bool`

HasNotificationIds returns a boolean if a field has been set.

### SetNotificationIdsNil

`func (o *TagDetailsResource) SetNotificationIdsNil(b bool)`

 SetNotificationIdsNil sets the value for NotificationIds to be an explicit nil

### UnsetNotificationIds
`func (o *TagDetailsResource) UnsetNotificationIds()`

UnsetNotificationIds ensures that no value is present for NotificationIds, not even an explicit nil
### GetRestrictionIds

`func (o *TagDetailsResource) GetRestrictionIds() []int32`

GetRestrictionIds returns the RestrictionIds field if non-nil, zero value otherwise.

### GetRestrictionIdsOk

`func (o *TagDetailsResource) GetRestrictionIdsOk() (*[]int32, bool)`

GetRestrictionIdsOk returns a tuple with the RestrictionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionIds

`func (o *TagDetailsResource) SetRestrictionIds(v []int32)`

SetRestrictionIds sets RestrictionIds field to given value.

### HasRestrictionIds

`func (o *TagDetailsResource) HasRestrictionIds() bool`

HasRestrictionIds returns a boolean if a field has been set.

### SetRestrictionIdsNil

`func (o *TagDetailsResource) SetRestrictionIdsNil(b bool)`

 SetRestrictionIdsNil sets the value for RestrictionIds to be an explicit nil

### UnsetRestrictionIds
`func (o *TagDetailsResource) UnsetRestrictionIds()`

UnsetRestrictionIds ensures that no value is present for RestrictionIds, not even an explicit nil
### GetAuthorIds

`func (o *TagDetailsResource) GetAuthorIds() []int32`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *TagDetailsResource) GetAuthorIdsOk() (*[]int32, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *TagDetailsResource) SetAuthorIds(v []int32)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *TagDetailsResource) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### SetAuthorIdsNil

`func (o *TagDetailsResource) SetAuthorIdsNil(b bool)`

 SetAuthorIdsNil sets the value for AuthorIds to be an explicit nil

### UnsetAuthorIds
`func (o *TagDetailsResource) UnsetAuthorIds()`

UnsetAuthorIds ensures that no value is present for AuthorIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


