/*
Readarr

Readarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// BookEditorResource struct for BookEditorResource
type BookEditorResource struct {
	BookIds []*int32 `json:"bookIds,omitempty"`
	Monitored NullableBool `json:"monitored,omitempty"`
	DeleteFiles NullableBool `json:"deleteFiles,omitempty"`
	AddImportListExclusion NullableBool `json:"addImportListExclusion,omitempty"`
}

// NewBookEditorResource instantiates a new BookEditorResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookEditorResource() *BookEditorResource {
	this := BookEditorResource{}
	return &this
}

// NewBookEditorResourceWithDefaults instantiates a new BookEditorResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookEditorResourceWithDefaults() *BookEditorResource {
	this := BookEditorResource{}
	return &this
}

// GetBookIds returns the BookIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookEditorResource) GetBookIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.BookIds
}

// GetBookIdsOk returns a tuple with the BookIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookEditorResource) GetBookIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.BookIds) {
    return nil, false
	}
	return o.BookIds, true
}

// HasBookIds returns a boolean if a field has been set.
func (o *BookEditorResource) HasBookIds() bool {
	if o != nil && isNil(o.BookIds) {
		return true
	}

	return false
}

// SetBookIds gets a reference to the given []int32 and assigns it to the BookIds field.
func (o *BookEditorResource) SetBookIds(v []*int32) {
	o.BookIds = v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookEditorResource) GetMonitored() bool {
	if o == nil || isNil(o.Monitored.Get()) {
		var ret bool
		return ret
	}
	return *o.Monitored.Get()
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookEditorResource) GetMonitoredOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.Monitored.Get(), o.Monitored.IsSet()
}

// HasMonitored returns a boolean if a field has been set.
func (o *BookEditorResource) HasMonitored() bool {
	if o != nil && o.Monitored.IsSet() {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given NullableBool and assigns it to the Monitored field.
func (o *BookEditorResource) SetMonitored(v bool) {
	o.Monitored.Set(&v)
}
// SetMonitoredNil sets the value for Monitored to be an explicit nil
func (o *BookEditorResource) SetMonitoredNil() {
	o.Monitored.Set(nil)
}

// UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
func (o *BookEditorResource) UnsetMonitored() {
	o.Monitored.Unset()
}

// GetDeleteFiles returns the DeleteFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookEditorResource) GetDeleteFiles() bool {
	if o == nil || isNil(o.DeleteFiles.Get()) {
		var ret bool
		return ret
	}
	return *o.DeleteFiles.Get()
}

// GetDeleteFilesOk returns a tuple with the DeleteFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookEditorResource) GetDeleteFilesOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.DeleteFiles.Get(), o.DeleteFiles.IsSet()
}

// HasDeleteFiles returns a boolean if a field has been set.
func (o *BookEditorResource) HasDeleteFiles() bool {
	if o != nil && o.DeleteFiles.IsSet() {
		return true
	}

	return false
}

// SetDeleteFiles gets a reference to the given NullableBool and assigns it to the DeleteFiles field.
func (o *BookEditorResource) SetDeleteFiles(v bool) {
	o.DeleteFiles.Set(&v)
}
// SetDeleteFilesNil sets the value for DeleteFiles to be an explicit nil
func (o *BookEditorResource) SetDeleteFilesNil() {
	o.DeleteFiles.Set(nil)
}

// UnsetDeleteFiles ensures that no value is present for DeleteFiles, not even an explicit nil
func (o *BookEditorResource) UnsetDeleteFiles() {
	o.DeleteFiles.Unset()
}

// GetAddImportListExclusion returns the AddImportListExclusion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookEditorResource) GetAddImportListExclusion() bool {
	if o == nil || isNil(o.AddImportListExclusion.Get()) {
		var ret bool
		return ret
	}
	return *o.AddImportListExclusion.Get()
}

// GetAddImportListExclusionOk returns a tuple with the AddImportListExclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookEditorResource) GetAddImportListExclusionOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.AddImportListExclusion.Get(), o.AddImportListExclusion.IsSet()
}

// HasAddImportListExclusion returns a boolean if a field has been set.
func (o *BookEditorResource) HasAddImportListExclusion() bool {
	if o != nil && o.AddImportListExclusion.IsSet() {
		return true
	}

	return false
}

// SetAddImportListExclusion gets a reference to the given NullableBool and assigns it to the AddImportListExclusion field.
func (o *BookEditorResource) SetAddImportListExclusion(v bool) {
	o.AddImportListExclusion.Set(&v)
}
// SetAddImportListExclusionNil sets the value for AddImportListExclusion to be an explicit nil
func (o *BookEditorResource) SetAddImportListExclusionNil() {
	o.AddImportListExclusion.Set(nil)
}

// UnsetAddImportListExclusion ensures that no value is present for AddImportListExclusion, not even an explicit nil
func (o *BookEditorResource) UnsetAddImportListExclusion() {
	o.AddImportListExclusion.Unset()
}

func (o BookEditorResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BookIds != nil {
		toSerialize["bookIds"] = o.BookIds
	}
	if o.Monitored.IsSet() {
		toSerialize["monitored"] = o.Monitored.Get()
	}
	if o.DeleteFiles.IsSet() {
		toSerialize["deleteFiles"] = o.DeleteFiles.Get()
	}
	if o.AddImportListExclusion.IsSet() {
		toSerialize["addImportListExclusion"] = o.AddImportListExclusion.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBookEditorResource struct {
	value *BookEditorResource
	isSet bool
}

func (v NullableBookEditorResource) Get() *BookEditorResource {
	return v.value
}

func (v *NullableBookEditorResource) Set(val *BookEditorResource) {
	v.value = val
	v.isSet = true
}

func (v NullableBookEditorResource) IsSet() bool {
	return v.isSet
}

func (v *NullableBookEditorResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookEditorResource(val *BookEditorResource) *NullableBookEditorResource {
	return &NullableBookEditorResource{value: val, isSet: true}
}

func (v NullableBookEditorResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookEditorResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

