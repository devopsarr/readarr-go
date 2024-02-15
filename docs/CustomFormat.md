# CustomFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IncludeCustomFormatWhenRenaming** | Pointer to **bool** |  | [optional] 
**Specifications** | Pointer to [**[]ICustomFormatSpecification**](ICustomFormatSpecification.md) |  | [optional] 

## Methods

### NewCustomFormat

`func NewCustomFormat() *CustomFormat`

NewCustomFormat instantiates a new CustomFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFormatWithDefaults

`func NewCustomFormatWithDefaults() *CustomFormat`

NewCustomFormatWithDefaults instantiates a new CustomFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFormat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFormat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFormat) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomFormat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFormat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFormat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFormat) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomFormat) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomFormat) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIncludeCustomFormatWhenRenaming

`func (o *CustomFormat) GetIncludeCustomFormatWhenRenaming() bool`

GetIncludeCustomFormatWhenRenaming returns the IncludeCustomFormatWhenRenaming field if non-nil, zero value otherwise.

### GetIncludeCustomFormatWhenRenamingOk

`func (o *CustomFormat) GetIncludeCustomFormatWhenRenamingOk() (*bool, bool)`

GetIncludeCustomFormatWhenRenamingOk returns a tuple with the IncludeCustomFormatWhenRenaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCustomFormatWhenRenaming

`func (o *CustomFormat) SetIncludeCustomFormatWhenRenaming(v bool)`

SetIncludeCustomFormatWhenRenaming sets IncludeCustomFormatWhenRenaming field to given value.

### HasIncludeCustomFormatWhenRenaming

`func (o *CustomFormat) HasIncludeCustomFormatWhenRenaming() bool`

HasIncludeCustomFormatWhenRenaming returns a boolean if a field has been set.

### GetSpecifications

`func (o *CustomFormat) GetSpecifications() []ICustomFormatSpecification`

GetSpecifications returns the Specifications field if non-nil, zero value otherwise.

### GetSpecificationsOk

`func (o *CustomFormat) GetSpecificationsOk() (*[]ICustomFormatSpecification, bool)`

GetSpecificationsOk returns a tuple with the Specifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifications

`func (o *CustomFormat) SetSpecifications(v []ICustomFormatSpecification)`

SetSpecifications sets Specifications field to given value.

### HasSpecifications

`func (o *CustomFormat) HasSpecifications() bool`

HasSpecifications returns a boolean if a field has been set.

### SetSpecificationsNil

`func (o *CustomFormat) SetSpecificationsNil(b bool)`

 SetSpecificationsNil sets the value for Specifications to be an explicit nil

### UnsetSpecifications
`func (o *CustomFormat) UnsetSpecifications()`

UnsetSpecifications ensures that no value is present for Specifications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


