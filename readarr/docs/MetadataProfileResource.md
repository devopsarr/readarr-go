# MetadataProfileResource

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
**Ignored** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetadataProfileResource

`func NewMetadataProfileResource() *MetadataProfileResource`

NewMetadataProfileResource instantiates a new MetadataProfileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProfileResourceWithDefaults

`func NewMetadataProfileResourceWithDefaults() *MetadataProfileResource`

NewMetadataProfileResourceWithDefaults instantiates a new MetadataProfileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataProfileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataProfileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataProfileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataProfileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetadataProfileResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataProfileResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataProfileResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataProfileResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataProfileResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataProfileResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMinPopularity

`func (o *MetadataProfileResource) GetMinPopularity() float64`

GetMinPopularity returns the MinPopularity field if non-nil, zero value otherwise.

### GetMinPopularityOk

`func (o *MetadataProfileResource) GetMinPopularityOk() (*float64, bool)`

GetMinPopularityOk returns a tuple with the MinPopularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPopularity

`func (o *MetadataProfileResource) SetMinPopularity(v float64)`

SetMinPopularity sets MinPopularity field to given value.

### HasMinPopularity

`func (o *MetadataProfileResource) HasMinPopularity() bool`

HasMinPopularity returns a boolean if a field has been set.

### GetSkipMissingDate

`func (o *MetadataProfileResource) GetSkipMissingDate() bool`

GetSkipMissingDate returns the SkipMissingDate field if non-nil, zero value otherwise.

### GetSkipMissingDateOk

`func (o *MetadataProfileResource) GetSkipMissingDateOk() (*bool, bool)`

GetSkipMissingDateOk returns a tuple with the SkipMissingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMissingDate

`func (o *MetadataProfileResource) SetSkipMissingDate(v bool)`

SetSkipMissingDate sets SkipMissingDate field to given value.

### HasSkipMissingDate

`func (o *MetadataProfileResource) HasSkipMissingDate() bool`

HasSkipMissingDate returns a boolean if a field has been set.

### GetSkipMissingIsbn

`func (o *MetadataProfileResource) GetSkipMissingIsbn() bool`

GetSkipMissingIsbn returns the SkipMissingIsbn field if non-nil, zero value otherwise.

### GetSkipMissingIsbnOk

`func (o *MetadataProfileResource) GetSkipMissingIsbnOk() (*bool, bool)`

GetSkipMissingIsbnOk returns a tuple with the SkipMissingIsbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMissingIsbn

`func (o *MetadataProfileResource) SetSkipMissingIsbn(v bool)`

SetSkipMissingIsbn sets SkipMissingIsbn field to given value.

### HasSkipMissingIsbn

`func (o *MetadataProfileResource) HasSkipMissingIsbn() bool`

HasSkipMissingIsbn returns a boolean if a field has been set.

### GetSkipPartsAndSets

`func (o *MetadataProfileResource) GetSkipPartsAndSets() bool`

GetSkipPartsAndSets returns the SkipPartsAndSets field if non-nil, zero value otherwise.

### GetSkipPartsAndSetsOk

`func (o *MetadataProfileResource) GetSkipPartsAndSetsOk() (*bool, bool)`

GetSkipPartsAndSetsOk returns a tuple with the SkipPartsAndSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPartsAndSets

`func (o *MetadataProfileResource) SetSkipPartsAndSets(v bool)`

SetSkipPartsAndSets sets SkipPartsAndSets field to given value.

### HasSkipPartsAndSets

`func (o *MetadataProfileResource) HasSkipPartsAndSets() bool`

HasSkipPartsAndSets returns a boolean if a field has been set.

### GetSkipSeriesSecondary

`func (o *MetadataProfileResource) GetSkipSeriesSecondary() bool`

GetSkipSeriesSecondary returns the SkipSeriesSecondary field if non-nil, zero value otherwise.

### GetSkipSeriesSecondaryOk

`func (o *MetadataProfileResource) GetSkipSeriesSecondaryOk() (*bool, bool)`

GetSkipSeriesSecondaryOk returns a tuple with the SkipSeriesSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSeriesSecondary

`func (o *MetadataProfileResource) SetSkipSeriesSecondary(v bool)`

SetSkipSeriesSecondary sets SkipSeriesSecondary field to given value.

### HasSkipSeriesSecondary

`func (o *MetadataProfileResource) HasSkipSeriesSecondary() bool`

HasSkipSeriesSecondary returns a boolean if a field has been set.

### GetAllowedLanguages

`func (o *MetadataProfileResource) GetAllowedLanguages() string`

GetAllowedLanguages returns the AllowedLanguages field if non-nil, zero value otherwise.

### GetAllowedLanguagesOk

`func (o *MetadataProfileResource) GetAllowedLanguagesOk() (*string, bool)`

GetAllowedLanguagesOk returns a tuple with the AllowedLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLanguages

`func (o *MetadataProfileResource) SetAllowedLanguages(v string)`

SetAllowedLanguages sets AllowedLanguages field to given value.

### HasAllowedLanguages

`func (o *MetadataProfileResource) HasAllowedLanguages() bool`

HasAllowedLanguages returns a boolean if a field has been set.

### SetAllowedLanguagesNil

`func (o *MetadataProfileResource) SetAllowedLanguagesNil(b bool)`

 SetAllowedLanguagesNil sets the value for AllowedLanguages to be an explicit nil

### UnsetAllowedLanguages
`func (o *MetadataProfileResource) UnsetAllowedLanguages()`

UnsetAllowedLanguages ensures that no value is present for AllowedLanguages, not even an explicit nil
### GetMinPages

`func (o *MetadataProfileResource) GetMinPages() int32`

GetMinPages returns the MinPages field if non-nil, zero value otherwise.

### GetMinPagesOk

`func (o *MetadataProfileResource) GetMinPagesOk() (*int32, bool)`

GetMinPagesOk returns a tuple with the MinPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPages

`func (o *MetadataProfileResource) SetMinPages(v int32)`

SetMinPages sets MinPages field to given value.

### HasMinPages

`func (o *MetadataProfileResource) HasMinPages() bool`

HasMinPages returns a boolean if a field has been set.

### GetIgnored

`func (o *MetadataProfileResource) GetIgnored() []string`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *MetadataProfileResource) GetIgnoredOk() (*[]string, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *MetadataProfileResource) SetIgnored(v []string)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *MetadataProfileResource) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### SetIgnoredNil

`func (o *MetadataProfileResource) SetIgnoredNil(b bool)`

 SetIgnoredNil sets the value for Ignored to be an explicit nil

### UnsetIgnored
`func (o *MetadataProfileResource) UnsetIgnored()`

UnsetIgnored ensures that no value is present for Ignored, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


