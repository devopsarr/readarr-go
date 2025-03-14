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

// AuthorStatusType the model 'AuthorStatusType'
type AuthorStatusType string

// List of AuthorStatusType
const (
	AUTHORSTATUSTYPE_CONTINUING AuthorStatusType = "continuing"
	AUTHORSTATUSTYPE_ENDED AuthorStatusType = "ended"
)

// All allowed values of AuthorStatusType enum
var AllowedAuthorStatusTypeEnumValues = []AuthorStatusType{
	"continuing",
	"ended",
}

func (v *AuthorStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthorStatusType(value)
	for _, existing := range AllowedAuthorStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthorStatusType", value)
}

// NewAuthorStatusTypeFromValue returns a pointer to a valid AuthorStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthorStatusTypeFromValue(v string) (*AuthorStatusType, error) {
	ev := AuthorStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthorStatusType: valid values are %v", v, AllowedAuthorStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthorStatusType) IsValid() bool {
	for _, existing := range AllowedAuthorStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthorStatusType value
func (v AuthorStatusType) Ptr() *AuthorStatusType {
	return &v
}

type NullableAuthorStatusType struct {
	value *AuthorStatusType
	isSet bool
}

func (v NullableAuthorStatusType) Get() *AuthorStatusType {
	return v.value
}

func (v *NullableAuthorStatusType) Set(val *AuthorStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorStatusType(val *AuthorStatusType) *NullableAuthorStatusType {
	return &NullableAuthorStatusType{value: val, isSet: true}
}

func (v NullableAuthorStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

