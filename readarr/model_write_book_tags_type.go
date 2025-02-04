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

// WriteBookTagsType the model 'WriteBookTagsType'
type WriteBookTagsType string

// List of WriteBookTagsType
const (
	WRITEBOOKTAGSTYPE_NEW_FILES WriteBookTagsType = "newFiles"
	WRITEBOOKTAGSTYPE_ALL_FILES WriteBookTagsType = "allFiles"
	WRITEBOOKTAGSTYPE_SYNC WriteBookTagsType = "sync"
)

// All allowed values of WriteBookTagsType enum
var AllowedWriteBookTagsTypeEnumValues = []WriteBookTagsType{
	"newFiles",
	"allFiles",
	"sync",
}

func (v *WriteBookTagsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WriteBookTagsType(value)
	for _, existing := range AllowedWriteBookTagsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WriteBookTagsType", value)
}

// NewWriteBookTagsTypeFromValue returns a pointer to a valid WriteBookTagsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWriteBookTagsTypeFromValue(v string) (*WriteBookTagsType, error) {
	ev := WriteBookTagsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WriteBookTagsType: valid values are %v", v, AllowedWriteBookTagsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WriteBookTagsType) IsValid() bool {
	for _, existing := range AllowedWriteBookTagsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WriteBookTagsType value
func (v WriteBookTagsType) Ptr() *WriteBookTagsType {
	return &v
}

type NullableWriteBookTagsType struct {
	value *WriteBookTagsType
	isSet bool
}

func (v NullableWriteBookTagsType) Get() *WriteBookTagsType {
	return v.value
}

func (v *NullableWriteBookTagsType) Set(val *WriteBookTagsType) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteBookTagsType) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteBookTagsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteBookTagsType(val *WriteBookTagsType) *NullableWriteBookTagsType {
	return &NullableWriteBookTagsType{value: val, isSet: true}
}

func (v NullableWriteBookTagsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteBookTagsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

