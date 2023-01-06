# QualityProfileQualityItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**Quality**](Quality.md) |  | [optional] 
**Items** | Pointer to [**[]QualityProfileQualityItem**](QualityProfileQualityItem.md) |  | [optional] 
**Allowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewQualityProfileQualityItem

`func NewQualityProfileQualityItem() *QualityProfileQualityItem`

NewQualityProfileQualityItem instantiates a new QualityProfileQualityItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityProfileQualityItemWithDefaults

`func NewQualityProfileQualityItemWithDefaults() *QualityProfileQualityItem`

NewQualityProfileQualityItemWithDefaults instantiates a new QualityProfileQualityItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QualityProfileQualityItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QualityProfileQualityItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QualityProfileQualityItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QualityProfileQualityItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *QualityProfileQualityItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QualityProfileQualityItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QualityProfileQualityItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QualityProfileQualityItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *QualityProfileQualityItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QualityProfileQualityItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuality

`func (o *QualityProfileQualityItem) GetQuality() Quality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *QualityProfileQualityItem) GetQualityOk() (*Quality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *QualityProfileQualityItem) SetQuality(v Quality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *QualityProfileQualityItem) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetItems

`func (o *QualityProfileQualityItem) GetItems() []QualityProfileQualityItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QualityProfileQualityItem) GetItemsOk() (*[]QualityProfileQualityItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QualityProfileQualityItem) SetItems(v []QualityProfileQualityItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *QualityProfileQualityItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *QualityProfileQualityItem) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *QualityProfileQualityItem) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetAllowed

`func (o *QualityProfileQualityItem) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *QualityProfileQualityItem) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *QualityProfileQualityItem) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *QualityProfileQualityItem) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


