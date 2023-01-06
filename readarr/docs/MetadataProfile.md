# MetadataProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**MinPopularity** | Pointer to **float64** |  | [optional] 
**SkipMissingDate** | Pointer to **bool** |  | [optional] 
**SkipMissingIsbn** | Pointer to **bool** |  | [optional] 
**SkipPartsAndSets** | Pointer to **bool** |  | [optional] 
**SkipSeriesSecondary** | Pointer to **bool** |  | [optional] 
**AllowedLanguages** | Pointer to **NullableString** |  | [optional] 
**MinPages** | Pointer to **int32** |  | [optional] 
**Ignored** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMetadataProfile

`func NewMetadataProfile() *MetadataProfile`

NewMetadataProfile instantiates a new MetadataProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProfileWithDefaults

`func NewMetadataProfileWithDefaults() *MetadataProfile`

NewMetadataProfileWithDefaults instantiates a new MetadataProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetadataProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMinPopularity

`func (o *MetadataProfile) GetMinPopularity() float64`

GetMinPopularity returns the MinPopularity field if non-nil, zero value otherwise.

### GetMinPopularityOk

`func (o *MetadataProfile) GetMinPopularityOk() (*float64, bool)`

GetMinPopularityOk returns a tuple with the MinPopularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPopularity

`func (o *MetadataProfile) SetMinPopularity(v float64)`

SetMinPopularity sets MinPopularity field to given value.

### HasMinPopularity

`func (o *MetadataProfile) HasMinPopularity() bool`

HasMinPopularity returns a boolean if a field has been set.

### GetSkipMissingDate

`func (o *MetadataProfile) GetSkipMissingDate() bool`

GetSkipMissingDate returns the SkipMissingDate field if non-nil, zero value otherwise.

### GetSkipMissingDateOk

`func (o *MetadataProfile) GetSkipMissingDateOk() (*bool, bool)`

GetSkipMissingDateOk returns a tuple with the SkipMissingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMissingDate

`func (o *MetadataProfile) SetSkipMissingDate(v bool)`

SetSkipMissingDate sets SkipMissingDate field to given value.

### HasSkipMissingDate

`func (o *MetadataProfile) HasSkipMissingDate() bool`

HasSkipMissingDate returns a boolean if a field has been set.

### GetSkipMissingIsbn

`func (o *MetadataProfile) GetSkipMissingIsbn() bool`

GetSkipMissingIsbn returns the SkipMissingIsbn field if non-nil, zero value otherwise.

### GetSkipMissingIsbnOk

`func (o *MetadataProfile) GetSkipMissingIsbnOk() (*bool, bool)`

GetSkipMissingIsbnOk returns a tuple with the SkipMissingIsbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMissingIsbn

`func (o *MetadataProfile) SetSkipMissingIsbn(v bool)`

SetSkipMissingIsbn sets SkipMissingIsbn field to given value.

### HasSkipMissingIsbn

`func (o *MetadataProfile) HasSkipMissingIsbn() bool`

HasSkipMissingIsbn returns a boolean if a field has been set.

### GetSkipPartsAndSets

`func (o *MetadataProfile) GetSkipPartsAndSets() bool`

GetSkipPartsAndSets returns the SkipPartsAndSets field if non-nil, zero value otherwise.

### GetSkipPartsAndSetsOk

`func (o *MetadataProfile) GetSkipPartsAndSetsOk() (*bool, bool)`

GetSkipPartsAndSetsOk returns a tuple with the SkipPartsAndSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPartsAndSets

`func (o *MetadataProfile) SetSkipPartsAndSets(v bool)`

SetSkipPartsAndSets sets SkipPartsAndSets field to given value.

### HasSkipPartsAndSets

`func (o *MetadataProfile) HasSkipPartsAndSets() bool`

HasSkipPartsAndSets returns a boolean if a field has been set.

### GetSkipSeriesSecondary

`func (o *MetadataProfile) GetSkipSeriesSecondary() bool`

GetSkipSeriesSecondary returns the SkipSeriesSecondary field if non-nil, zero value otherwise.

### GetSkipSeriesSecondaryOk

`func (o *MetadataProfile) GetSkipSeriesSecondaryOk() (*bool, bool)`

GetSkipSeriesSecondaryOk returns a tuple with the SkipSeriesSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSeriesSecondary

`func (o *MetadataProfile) SetSkipSeriesSecondary(v bool)`

SetSkipSeriesSecondary sets SkipSeriesSecondary field to given value.

### HasSkipSeriesSecondary

`func (o *MetadataProfile) HasSkipSeriesSecondary() bool`

HasSkipSeriesSecondary returns a boolean if a field has been set.

### GetAllowedLanguages

`func (o *MetadataProfile) GetAllowedLanguages() string`

GetAllowedLanguages returns the AllowedLanguages field if non-nil, zero value otherwise.

### GetAllowedLanguagesOk

`func (o *MetadataProfile) GetAllowedLanguagesOk() (*string, bool)`

GetAllowedLanguagesOk returns a tuple with the AllowedLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLanguages

`func (o *MetadataProfile) SetAllowedLanguages(v string)`

SetAllowedLanguages sets AllowedLanguages field to given value.

### HasAllowedLanguages

`func (o *MetadataProfile) HasAllowedLanguages() bool`

HasAllowedLanguages returns a boolean if a field has been set.

### SetAllowedLanguagesNil

`func (o *MetadataProfile) SetAllowedLanguagesNil(b bool)`

 SetAllowedLanguagesNil sets the value for AllowedLanguages to be an explicit nil

### UnsetAllowedLanguages
`func (o *MetadataProfile) UnsetAllowedLanguages()`

UnsetAllowedLanguages ensures that no value is present for AllowedLanguages, not even an explicit nil
### GetMinPages

`func (o *MetadataProfile) GetMinPages() int32`

GetMinPages returns the MinPages field if non-nil, zero value otherwise.

### GetMinPagesOk

`func (o *MetadataProfile) GetMinPagesOk() (*int32, bool)`

GetMinPagesOk returns a tuple with the MinPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPages

`func (o *MetadataProfile) SetMinPages(v int32)`

SetMinPages sets MinPages field to given value.

### HasMinPages

`func (o *MetadataProfile) HasMinPages() bool`

HasMinPages returns a boolean if a field has been set.

### GetIgnored

`func (o *MetadataProfile) GetIgnored() string`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *MetadataProfile) GetIgnoredOk() (*string, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *MetadataProfile) SetIgnored(v string)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *MetadataProfile) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### SetIgnoredNil

`func (o *MetadataProfile) SetIgnoredNil(b bool)`

 SetIgnoredNil sets the value for Ignored to be an explicit nil

### UnsetIgnored
`func (o *MetadataProfile) UnsetIgnored()`

UnsetIgnored ensures that no value is present for Ignored, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


