/*
Readarr

Readarr API docs

API version: v0.4.10.2734
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
	"fmt"
)

// BookAddType the model 'BookAddType'
type BookAddType string

// List of BookAddType
const (
	BOOKADDTYPE_AUTOMATIC BookAddType = "automatic"
	BOOKADDTYPE_MANUAL BookAddType = "manual"
)

// All allowed values of BookAddType enum
var AllowedBookAddTypeEnumValues = []BookAddType{
	"automatic",
	"manual",
}

func (v *BookAddType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BookAddType(value)
	for _, existing := range AllowedBookAddTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BookAddType", value)
}

// NewBookAddTypeFromValue returns a pointer to a valid BookAddType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBookAddTypeFromValue(v string) (*BookAddType, error) {
	ev := BookAddType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BookAddType: valid values are %v", v, AllowedBookAddTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BookAddType) IsValid() bool {
	for _, existing := range AllowedBookAddTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BookAddType value
func (v BookAddType) Ptr() *BookAddType {
	return &v
}

type NullableBookAddType struct {
	value *BookAddType
	isSet bool
}

func (v NullableBookAddType) Get() *BookAddType {
	return v.value
}

func (v *NullableBookAddType) Set(val *BookAddType) {
	v.value = val
	v.isSet = true
}

func (v NullableBookAddType) IsSet() bool {
	return v.isSet
}

func (v *NullableBookAddType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookAddType(val *BookAddType) *NullableBookAddType {
	return &NullableBookAddType{value: val, isSet: true}
}

func (v NullableBookAddType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookAddType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

