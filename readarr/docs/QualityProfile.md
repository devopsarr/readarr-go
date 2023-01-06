# QualityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**UpgradeAllowed** | Pointer to **bool** |  | [optional] 
**Cutoff** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]QualityProfileQualityItem**](QualityProfileQualityItem.md) |  | [optional] 

## Methods

### NewQualityProfile

`func NewQualityProfile() *QualityProfile`

NewQualityProfile instantiates a new QualityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityProfileWithDefaults

`func NewQualityProfileWithDefaults() *QualityProfile`

NewQualityProfileWithDefaults instantiates a new QualityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QualityProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QualityProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QualityProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QualityProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *QualityProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QualityProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QualityProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QualityProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *QualityProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QualityProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUpgradeAllowed

`func (o *QualityProfile) GetUpgradeAllowed() bool`

GetUpgradeAllowed returns the UpgradeAllowed field if non-nil, zero value otherwise.

### GetUpgradeAllowedOk

`func (o *QualityProfile) GetUpgradeAllowedOk() (*bool, bool)`

GetUpgradeAllowedOk returns a tuple with the UpgradeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeAllowed

`func (o *QualityProfile) SetUpgradeAllowed(v bool)`

SetUpgradeAllowed sets UpgradeAllowed field to given value.

### HasUpgradeAllowed

`func (o *QualityProfile) HasUpgradeAllowed() bool`

HasUpgradeAllowed returns a boolean if a field has been set.

### GetCutoff

`func (o *QualityProfile) GetCutoff() int32`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *QualityProfile) GetCutoffOk() (*int32, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *QualityProfile) SetCutoff(v int32)`

SetCutoff sets Cutoff field to given value.

### HasCutoff

`func (o *QualityProfile) HasCutoff() bool`

HasCutoff returns a boolean if a field has been set.

### GetItems

`func (o *QualityProfile) GetItems() []QualityProfileQualityItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QualityProfile) GetItemsOk() (*[]QualityProfileQualityItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QualityProfile) SetItems(v []QualityProfileQualityItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *QualityProfile) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *QualityProfile) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *QualityProfile) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


