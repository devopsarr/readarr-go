/*
Readarr

Readarr API docs

API version: v0.3.18.2411
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
	"fmt"
)

// MonitorTypes the model 'MonitorTypes'
type MonitorTypes string

// List of MonitorTypes
const (
	MONITORTYPES_ALL MonitorTypes = "all"
	MONITORTYPES_FUTURE MonitorTypes = "future"
	MONITORTYPES_MISSING MonitorTypes = "missing"
	MONITORTYPES_EXISTING MonitorTypes = "existing"
	MONITORTYPES_LATEST MonitorTypes = "latest"
	MONITORTYPES_FIRST MonitorTypes = "first"
	MONITORTYPES_NONE MonitorTypes = "none"
	MONITORTYPES_UNKNOWN MonitorTypes = "unknown"
)

// All allowed values of MonitorTypes enum
var AllowedMonitorTypesEnumValues = []MonitorTypes{
	"all",
	"future",
	"missing",
	"existing",
	"latest",
	"first",
	"none",
	"unknown",
}

func (v *MonitorTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonitorTypes(value)
	for _, existing := range AllowedMonitorTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonitorTypes", value)
}

// NewMonitorTypesFromValue returns a pointer to a valid MonitorTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorTypesFromValue(v string) (*MonitorTypes, error) {
	ev := MonitorTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorTypes: valid values are %v", v, AllowedMonitorTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorTypes) IsValid() bool {
	for _, existing := range AllowedMonitorTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorTypes value
func (v MonitorTypes) Ptr() *MonitorTypes {
	return &v
}

type NullableMonitorTypes struct {
	value *MonitorTypes
	isSet bool
}

func (v NullableMonitorTypes) Get() *MonitorTypes {
	return v.value
}

func (v *NullableMonitorTypes) Set(val *MonitorTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTypes(val *MonitorTypes) *NullableMonitorTypes {
	return &NullableMonitorTypes{value: val, isSet: true}
}

func (v NullableMonitorTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

