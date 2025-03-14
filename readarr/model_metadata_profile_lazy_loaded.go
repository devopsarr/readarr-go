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

// checks if the MetadataProfileLazyLoaded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataProfileLazyLoaded{}

// MetadataProfileLazyLoaded struct for MetadataProfileLazyLoaded
type MetadataProfileLazyLoaded struct {
	Value *MetadataProfile `json:"value,omitempty"`
	IsLoaded *bool `json:"isLoaded,omitempty"`
}

// NewMetadataProfileLazyLoaded instantiates a new MetadataProfileLazyLoaded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataProfileLazyLoaded() *MetadataProfileLazyLoaded {
	this := MetadataProfileLazyLoaded{}
	return &this
}

// NewMetadataProfileLazyLoadedWithDefaults instantiates a new MetadataProfileLazyLoaded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataProfileLazyLoadedWithDefaults() *MetadataProfileLazyLoaded {
	this := MetadataProfileLazyLoaded{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetadataProfileLazyLoaded) GetValue() MetadataProfile {
	if o == nil || IsNil(o.Value) {
		var ret MetadataProfile
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileLazyLoaded) GetValueOk() (*MetadataProfile, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetadataProfileLazyLoaded) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given MetadataProfile and assigns it to the Value field.
func (o *MetadataProfileLazyLoaded) SetValue(v MetadataProfile) {
	o.Value = &v
}

// GetIsLoaded returns the IsLoaded field value if set, zero value otherwise.
func (o *MetadataProfileLazyLoaded) GetIsLoaded() bool {
	if o == nil || IsNil(o.IsLoaded) {
		var ret bool
		return ret
	}
	return *o.IsLoaded
}

// GetIsLoadedOk returns a tuple with the IsLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileLazyLoaded) GetIsLoadedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLoaded) {
		return nil, false
	}
	return o.IsLoaded, true
}

// HasIsLoaded returns a boolean if a field has been set.
func (o *MetadataProfileLazyLoaded) HasIsLoaded() bool {
	if o != nil && !IsNil(o.IsLoaded) {
		return true
	}

	return false
}

// SetIsLoaded gets a reference to the given bool and assigns it to the IsLoaded field.
func (o *MetadataProfileLazyLoaded) SetIsLoaded(v bool) {
	o.IsLoaded = &v
}

func (o MetadataProfileLazyLoaded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataProfileLazyLoaded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.IsLoaded) {
		toSerialize["isLoaded"] = o.IsLoaded
	}
	return toSerialize, nil
}

type NullableMetadataProfileLazyLoaded struct {
	value *MetadataProfileLazyLoaded
	isSet bool
}

func (v NullableMetadataProfileLazyLoaded) Get() *MetadataProfileLazyLoaded {
	return v.value
}

func (v *NullableMetadataProfileLazyLoaded) Set(val *MetadataProfileLazyLoaded) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataProfileLazyLoaded) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataProfileLazyLoaded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataProfileLazyLoaded(val *MetadataProfileLazyLoaded) *NullableMetadataProfileLazyLoaded {
	return &NullableMetadataProfileLazyLoaded{value: val, isSet: true}
}

func (v NullableMetadataProfileLazyLoaded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataProfileLazyLoaded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


