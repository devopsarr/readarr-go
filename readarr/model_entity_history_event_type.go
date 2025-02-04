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

// EntityHistoryEventType the model 'EntityHistoryEventType'
type EntityHistoryEventType string

// List of EntityHistoryEventType
const (
	ENTITYHISTORYEVENTTYPE_UNKNOWN EntityHistoryEventType = "unknown"
	ENTITYHISTORYEVENTTYPE_GRABBED EntityHistoryEventType = "grabbed"
	ENTITYHISTORYEVENTTYPE_BOOK_FILE_IMPORTED EntityHistoryEventType = "bookFileImported"
	ENTITYHISTORYEVENTTYPE_DOWNLOAD_FAILED EntityHistoryEventType = "downloadFailed"
	ENTITYHISTORYEVENTTYPE_BOOK_FILE_DELETED EntityHistoryEventType = "bookFileDeleted"
	ENTITYHISTORYEVENTTYPE_BOOK_FILE_RENAMED EntityHistoryEventType = "bookFileRenamed"
	ENTITYHISTORYEVENTTYPE_BOOK_IMPORT_INCOMPLETE EntityHistoryEventType = "bookImportIncomplete"
	ENTITYHISTORYEVENTTYPE_DOWNLOAD_IMPORTED EntityHistoryEventType = "downloadImported"
	ENTITYHISTORYEVENTTYPE_BOOK_FILE_RETAGGED EntityHistoryEventType = "bookFileRetagged"
	ENTITYHISTORYEVENTTYPE_DOWNLOAD_IGNORED EntityHistoryEventType = "downloadIgnored"
)

// All allowed values of EntityHistoryEventType enum
var AllowedEntityHistoryEventTypeEnumValues = []EntityHistoryEventType{
	"unknown",
	"grabbed",
	"bookFileImported",
	"downloadFailed",
	"bookFileDeleted",
	"bookFileRenamed",
	"bookImportIncomplete",
	"downloadImported",
	"bookFileRetagged",
	"downloadIgnored",
}

func (v *EntityHistoryEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntityHistoryEventType(value)
	for _, existing := range AllowedEntityHistoryEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntityHistoryEventType", value)
}

// NewEntityHistoryEventTypeFromValue returns a pointer to a valid EntityHistoryEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntityHistoryEventTypeFromValue(v string) (*EntityHistoryEventType, error) {
	ev := EntityHistoryEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntityHistoryEventType: valid values are %v", v, AllowedEntityHistoryEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntityHistoryEventType) IsValid() bool {
	for _, existing := range AllowedEntityHistoryEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityHistoryEventType value
func (v EntityHistoryEventType) Ptr() *EntityHistoryEventType {
	return &v
}

type NullableEntityHistoryEventType struct {
	value *EntityHistoryEventType
	isSet bool
}

func (v NullableEntityHistoryEventType) Get() *EntityHistoryEventType {
	return v.value
}

func (v *NullableEntityHistoryEventType) Set(val *EntityHistoryEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityHistoryEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityHistoryEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityHistoryEventType(val *EntityHistoryEventType) *NullableEntityHistoryEventType {
	return &NullableEntityHistoryEventType{value: val, isSet: true}
}

func (v NullableEntityHistoryEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityHistoryEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

