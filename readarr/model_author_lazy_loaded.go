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

// AuthorLazyLoaded struct for AuthorLazyLoaded
type AuthorLazyLoaded struct {
	Value *Author `json:"value,omitempty"`
	IsLoaded *bool `json:"isLoaded,omitempty"`
}

// NewAuthorLazyLoaded instantiates a new AuthorLazyLoaded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorLazyLoaded() *AuthorLazyLoaded {
	this := AuthorLazyLoaded{}
	return &this
}

// NewAuthorLazyLoadedWithDefaults instantiates a new AuthorLazyLoaded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorLazyLoadedWithDefaults() *AuthorLazyLoaded {
	this := AuthorLazyLoaded{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AuthorLazyLoaded) GetValue() Author {
	if o == nil || isNil(o.Value) {
		var ret Author
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorLazyLoaded) GetValueOk() (*Author, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AuthorLazyLoaded) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Author and assigns it to the Value field.
func (o *AuthorLazyLoaded) SetValue(v Author) {
	o.Value = &v
}

// GetIsLoaded returns the IsLoaded field value if set, zero value otherwise.
func (o *AuthorLazyLoaded) GetIsLoaded() bool {
	if o == nil || isNil(o.IsLoaded) {
		var ret bool
		return ret
	}
	return *o.IsLoaded
}

// GetIsLoadedOk returns a tuple with the IsLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorLazyLoaded) GetIsLoadedOk() (*bool, bool) {
	if o == nil || isNil(o.IsLoaded) {
    return nil, false
	}
	return o.IsLoaded, true
}

// HasIsLoaded returns a boolean if a field has been set.
func (o *AuthorLazyLoaded) HasIsLoaded() bool {
	if o != nil && !isNil(o.IsLoaded) {
		return true
	}

	return false
}

// SetIsLoaded gets a reference to the given bool and assigns it to the IsLoaded field.
func (o *AuthorLazyLoaded) SetIsLoaded(v bool) {
	o.IsLoaded = &v
}

func (o AuthorLazyLoaded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.IsLoaded) {
		toSerialize["isLoaded"] = o.IsLoaded
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorLazyLoaded struct {
	value *AuthorLazyLoaded
	isSet bool
}

func (v NullableAuthorLazyLoaded) Get() *AuthorLazyLoaded {
	return v.value
}

func (v *NullableAuthorLazyLoaded) Set(val *AuthorLazyLoaded) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorLazyLoaded) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorLazyLoaded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorLazyLoaded(val *AuthorLazyLoaded) *NullableAuthorLazyLoaded {
	return &NullableAuthorLazyLoaded{value: val, isSet: true}
}

func (v NullableAuthorLazyLoaded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorLazyLoaded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


