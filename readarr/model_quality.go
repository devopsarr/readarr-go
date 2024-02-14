/*
Readarr

Readarr API docs

API version: v0.3.18.2411
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// Quality struct for Quality
type Quality struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewQuality instantiates a new Quality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuality() *Quality {
	this := Quality{}
	return &this
}

// NewQualityWithDefaults instantiates a new Quality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityWithDefaults() *Quality {
	this := Quality{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Quality) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quality) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Quality) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Quality) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Quality) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Quality) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Quality) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Quality) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Quality) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Quality) UnsetName() {
	o.Name.Unset()
}

func (o Quality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableQuality struct {
	value *Quality
	isSet bool
}

func (v NullableQuality) Get() *Quality {
	return v.value
}

func (v *NullableQuality) Set(val *Quality) {
	v.value = val
	v.isSet = true
}

func (v NullableQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuality(val *Quality) *NullableQuality {
	return &NullableQuality{value: val, isSet: true}
}

func (v NullableQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


