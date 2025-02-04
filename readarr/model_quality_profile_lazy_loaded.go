/*
Readarr

Readarr API docs

API version: v0.4.10.2734
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// checks if the QualityProfileLazyLoaded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QualityProfileLazyLoaded{}

// QualityProfileLazyLoaded struct for QualityProfileLazyLoaded
type QualityProfileLazyLoaded struct {
	Value *QualityProfile `json:"value,omitempty"`
	IsLoaded *bool `json:"isLoaded,omitempty"`
}

// NewQualityProfileLazyLoaded instantiates a new QualityProfileLazyLoaded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityProfileLazyLoaded() *QualityProfileLazyLoaded {
	this := QualityProfileLazyLoaded{}
	return &this
}

// NewQualityProfileLazyLoadedWithDefaults instantiates a new QualityProfileLazyLoaded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityProfileLazyLoadedWithDefaults() *QualityProfileLazyLoaded {
	this := QualityProfileLazyLoaded{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *QualityProfileLazyLoaded) GetValue() QualityProfile {
	if o == nil || IsNil(o.Value) {
		var ret QualityProfile
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileLazyLoaded) GetValueOk() (*QualityProfile, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *QualityProfileLazyLoaded) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given QualityProfile and assigns it to the Value field.
func (o *QualityProfileLazyLoaded) SetValue(v QualityProfile) {
	o.Value = &v
}

// GetIsLoaded returns the IsLoaded field value if set, zero value otherwise.
func (o *QualityProfileLazyLoaded) GetIsLoaded() bool {
	if o == nil || IsNil(o.IsLoaded) {
		var ret bool
		return ret
	}
	return *o.IsLoaded
}

// GetIsLoadedOk returns a tuple with the IsLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileLazyLoaded) GetIsLoadedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLoaded) {
		return nil, false
	}
	return o.IsLoaded, true
}

// HasIsLoaded returns a boolean if a field has been set.
func (o *QualityProfileLazyLoaded) HasIsLoaded() bool {
	if o != nil && !IsNil(o.IsLoaded) {
		return true
	}

	return false
}

// SetIsLoaded gets a reference to the given bool and assigns it to the IsLoaded field.
func (o *QualityProfileLazyLoaded) SetIsLoaded(v bool) {
	o.IsLoaded = &v
}

func (o QualityProfileLazyLoaded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QualityProfileLazyLoaded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.IsLoaded) {
		toSerialize["isLoaded"] = o.IsLoaded
	}
	return toSerialize, nil
}

type NullableQualityProfileLazyLoaded struct {
	value *QualityProfileLazyLoaded
	isSet bool
}

func (v NullableQualityProfileLazyLoaded) Get() *QualityProfileLazyLoaded {
	return v.value
}

func (v *NullableQualityProfileLazyLoaded) Set(val *QualityProfileLazyLoaded) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityProfileLazyLoaded) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityProfileLazyLoaded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityProfileLazyLoaded(val *QualityProfileLazyLoaded) *NullableQualityProfileLazyLoaded {
	return &NullableQualityProfileLazyLoaded{value: val, isSet: true}
}

func (v NullableQualityProfileLazyLoaded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityProfileLazyLoaded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


