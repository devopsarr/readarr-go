# ProfileFormatItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**CustomFormat**](CustomFormat.md) |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 

## Methods

### NewProfileFormatItem

`func NewProfileFormatItem() *ProfileFormatItem`

NewProfileFormatItem instantiates a new ProfileFormatItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileFormatItemWithDefaults

`func NewProfileFormatItemWithDefaults() *ProfileFormatItem`

NewProfileFormatItemWithDefaults instantiates a new ProfileFormatItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ProfileFormatItem) GetFormat() CustomFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ProfileFormatItem) GetFormatOk() (*CustomFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ProfileFormatItem) SetFormat(v CustomFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ProfileFormatItem) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetScore

`func (o *ProfileFormatItem) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ProfileFormatItem) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ProfileFormatItem) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ProfileFormatItem) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


