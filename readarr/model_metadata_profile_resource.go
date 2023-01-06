/*
Readarr

Readarr API docs

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// MetadataProfileResource struct for MetadataProfileResource
type MetadataProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	MinPopularity *float64 `json:"minPopularity,omitempty"`
	SkipMissingDate *bool `json:"skipMissingDate,omitempty"`
	SkipMissingIsbn *bool `json:"skipMissingIsbn,omitempty"`
	SkipPartsAndSets *bool `json:"skipPartsAndSets,omitempty"`
	SkipSeriesSecondary *bool `json:"skipSeriesSecondary,omitempty"`
	AllowedLanguages NullableString `json:"allowedLanguages,omitempty"`
	MinPages *int32 `json:"minPages,omitempty"`
	Ignored NullableString `json:"ignored,omitempty"`
}

// NewMetadataProfileResource instantiates a new MetadataProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataProfileResource() *MetadataProfileResource {
	this := MetadataProfileResource{}
	return &this
}

// NewMetadataProfileResourceWithDefaults instantiates a new MetadataProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataProfileResourceWithDefaults() *MetadataProfileResource {
	this := MetadataProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MetadataProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProfileResource) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProfileResource) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetadataProfileResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetadataProfileResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetadataProfileResource) UnsetName() {
	o.Name.Unset()
}

// GetMinPopularity returns the MinPopularity field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetMinPopularity() float64 {
	if o == nil || isNil(o.MinPopularity) {
		var ret float64
		return ret
	}
	return *o.MinPopularity
}

// GetMinPopularityOk returns a tuple with the MinPopularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetMinPopularityOk() (*float64, bool) {
	if o == nil || isNil(o.MinPopularity) {
    return nil, false
	}
	return o.MinPopularity, true
}

// HasMinPopularity returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasMinPopularity() bool {
	if o != nil && !isNil(o.MinPopularity) {
		return true
	}

	return false
}

// SetMinPopularity gets a reference to the given float64 and assigns it to the MinPopularity field.
func (o *MetadataProfileResource) SetMinPopularity(v float64) {
	o.MinPopularity = &v
}

// GetSkipMissingDate returns the SkipMissingDate field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetSkipMissingDate() bool {
	if o == nil || isNil(o.SkipMissingDate) {
		var ret bool
		return ret
	}
	return *o.SkipMissingDate
}

// GetSkipMissingDateOk returns a tuple with the SkipMissingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetSkipMissingDateOk() (*bool, bool) {
	if o == nil || isNil(o.SkipMissingDate) {
    return nil, false
	}
	return o.SkipMissingDate, true
}

// HasSkipMissingDate returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasSkipMissingDate() bool {
	if o != nil && !isNil(o.SkipMissingDate) {
		return true
	}

	return false
}

// SetSkipMissingDate gets a reference to the given bool and assigns it to the SkipMissingDate field.
func (o *MetadataProfileResource) SetSkipMissingDate(v bool) {
	o.SkipMissingDate = &v
}

// GetSkipMissingIsbn returns the SkipMissingIsbn field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetSkipMissingIsbn() bool {
	if o == nil || isNil(o.SkipMissingIsbn) {
		var ret bool
		return ret
	}
	return *o.SkipMissingIsbn
}

// GetSkipMissingIsbnOk returns a tuple with the SkipMissingIsbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetSkipMissingIsbnOk() (*bool, bool) {
	if o == nil || isNil(o.SkipMissingIsbn) {
    return nil, false
	}
	return o.SkipMissingIsbn, true
}

// HasSkipMissingIsbn returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasSkipMissingIsbn() bool {
	if o != nil && !isNil(o.SkipMissingIsbn) {
		return true
	}

	return false
}

// SetSkipMissingIsbn gets a reference to the given bool and assigns it to the SkipMissingIsbn field.
func (o *MetadataProfileResource) SetSkipMissingIsbn(v bool) {
	o.SkipMissingIsbn = &v
}

// GetSkipPartsAndSets returns the SkipPartsAndSets field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetSkipPartsAndSets() bool {
	if o == nil || isNil(o.SkipPartsAndSets) {
		var ret bool
		return ret
	}
	return *o.SkipPartsAndSets
}

// GetSkipPartsAndSetsOk returns a tuple with the SkipPartsAndSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetSkipPartsAndSetsOk() (*bool, bool) {
	if o == nil || isNil(o.SkipPartsAndSets) {
    return nil, false
	}
	return o.SkipPartsAndSets, true
}

// HasSkipPartsAndSets returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasSkipPartsAndSets() bool {
	if o != nil && !isNil(o.SkipPartsAndSets) {
		return true
	}

	return false
}

// SetSkipPartsAndSets gets a reference to the given bool and assigns it to the SkipPartsAndSets field.
func (o *MetadataProfileResource) SetSkipPartsAndSets(v bool) {
	o.SkipPartsAndSets = &v
}

// GetSkipSeriesSecondary returns the SkipSeriesSecondary field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetSkipSeriesSecondary() bool {
	if o == nil || isNil(o.SkipSeriesSecondary) {
		var ret bool
		return ret
	}
	return *o.SkipSeriesSecondary
}

// GetSkipSeriesSecondaryOk returns a tuple with the SkipSeriesSecondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetSkipSeriesSecondaryOk() (*bool, bool) {
	if o == nil || isNil(o.SkipSeriesSecondary) {
    return nil, false
	}
	return o.SkipSeriesSecondary, true
}

// HasSkipSeriesSecondary returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasSkipSeriesSecondary() bool {
	if o != nil && !isNil(o.SkipSeriesSecondary) {
		return true
	}

	return false
}

// SetSkipSeriesSecondary gets a reference to the given bool and assigns it to the SkipSeriesSecondary field.
func (o *MetadataProfileResource) SetSkipSeriesSecondary(v bool) {
	o.SkipSeriesSecondary = &v
}

// GetAllowedLanguages returns the AllowedLanguages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProfileResource) GetAllowedLanguages() string {
	if o == nil || isNil(o.AllowedLanguages.Get()) {
		var ret string
		return ret
	}
	return *o.AllowedLanguages.Get()
}

// GetAllowedLanguagesOk returns a tuple with the AllowedLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProfileResource) GetAllowedLanguagesOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedLanguages.Get(), o.AllowedLanguages.IsSet()
}

// HasAllowedLanguages returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasAllowedLanguages() bool {
	if o != nil && o.AllowedLanguages.IsSet() {
		return true
	}

	return false
}

// SetAllowedLanguages gets a reference to the given NullableString and assigns it to the AllowedLanguages field.
func (o *MetadataProfileResource) SetAllowedLanguages(v string) {
	o.AllowedLanguages.Set(&v)
}
// SetAllowedLanguagesNil sets the value for AllowedLanguages to be an explicit nil
func (o *MetadataProfileResource) SetAllowedLanguagesNil() {
	o.AllowedLanguages.Set(nil)
}

// UnsetAllowedLanguages ensures that no value is present for AllowedLanguages, not even an explicit nil
func (o *MetadataProfileResource) UnsetAllowedLanguages() {
	o.AllowedLanguages.Unset()
}

// GetMinPages returns the MinPages field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetMinPages() int32 {
	if o == nil || isNil(o.MinPages) {
		var ret int32
		return ret
	}
	return *o.MinPages
}

// GetMinPagesOk returns a tuple with the MinPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetMinPagesOk() (*int32, bool) {
	if o == nil || isNil(o.MinPages) {
    return nil, false
	}
	return o.MinPages, true
}

// HasMinPages returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasMinPages() bool {
	if o != nil && !isNil(o.MinPages) {
		return true
	}

	return false
}

// SetMinPages gets a reference to the given int32 and assigns it to the MinPages field.
func (o *MetadataProfileResource) SetMinPages(v int32) {
	o.MinPages = &v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProfileResource) GetIgnored() string {
	if o == nil || isNil(o.Ignored.Get()) {
		var ret string
		return ret
	}
	return *o.Ignored.Get()
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProfileResource) GetIgnoredOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ignored.Get(), o.Ignored.IsSet()
}

// HasIgnored returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasIgnored() bool {
	if o != nil && o.Ignored.IsSet() {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given NullableString and assigns it to the Ignored field.
func (o *MetadataProfileResource) SetIgnored(v string) {
	o.Ignored.Set(&v)
}
// SetIgnoredNil sets the value for Ignored to be an explicit nil
func (o *MetadataProfileResource) SetIgnoredNil() {
	o.Ignored.Set(nil)
}

// UnsetIgnored ensures that no value is present for Ignored, not even an explicit nil
func (o *MetadataProfileResource) UnsetIgnored() {
	o.Ignored.Unset()
}

func (o MetadataProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.MinPopularity) {
		toSerialize["minPopularity"] = o.MinPopularity
	}
	if !isNil(o.SkipMissingDate) {
		toSerialize["skipMissingDate"] = o.SkipMissingDate
	}
	if !isNil(o.SkipMissingIsbn) {
		toSerialize["skipMissingIsbn"] = o.SkipMissingIsbn
	}
	if !isNil(o.SkipPartsAndSets) {
		toSerialize["skipPartsAndSets"] = o.SkipPartsAndSets
	}
	if !isNil(o.SkipSeriesSecondary) {
		toSerialize["skipSeriesSecondary"] = o.SkipSeriesSecondary
	}
	if o.AllowedLanguages.IsSet() {
		toSerialize["allowedLanguages"] = o.AllowedLanguages.Get()
	}
	if !isNil(o.MinPages) {
		toSerialize["minPages"] = o.MinPages
	}
	if o.Ignored.IsSet() {
		toSerialize["ignored"] = o.Ignored.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataProfileResource struct {
	value *MetadataProfileResource
	isSet bool
}

func (v NullableMetadataProfileResource) Get() *MetadataProfileResource {
	return v.value
}

func (v *NullableMetadataProfileResource) Set(val *MetadataProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataProfileResource(val *MetadataProfileResource) *NullableMetadataProfileResource {
	return &NullableMetadataProfileResource{value: val, isSet: true}
}

func (v NullableMetadataProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


