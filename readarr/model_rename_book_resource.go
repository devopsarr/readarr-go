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

// RenameBookResource struct for RenameBookResource
type RenameBookResource struct {
	Id *int32 `json:"id,omitempty"`
	AuthorId *int32 `json:"authorId,omitempty"`
	BookId *int32 `json:"bookId,omitempty"`
	BookFileId *int32 `json:"bookFileId,omitempty"`
	ExistingPath NullableString `json:"existingPath,omitempty"`
	NewPath NullableString `json:"newPath,omitempty"`
}

// NewRenameBookResource instantiates a new RenameBookResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameBookResource() *RenameBookResource {
	this := RenameBookResource{}
	return &this
}

// NewRenameBookResourceWithDefaults instantiates a new RenameBookResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameBookResourceWithDefaults() *RenameBookResource {
	this := RenameBookResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RenameBookResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameBookResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RenameBookResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RenameBookResource) SetId(v int32) {
	o.Id = &v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *RenameBookResource) GetAuthorId() int32 {
	if o == nil || isNil(o.AuthorId) {
		var ret int32
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameBookResource) GetAuthorIdOk() (*int32, bool) {
	if o == nil || isNil(o.AuthorId) {
    return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *RenameBookResource) HasAuthorId() bool {
	if o != nil && !isNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given int32 and assigns it to the AuthorId field.
func (o *RenameBookResource) SetAuthorId(v int32) {
	o.AuthorId = &v
}

// GetBookId returns the BookId field value if set, zero value otherwise.
func (o *RenameBookResource) GetBookId() int32 {
	if o == nil || isNil(o.BookId) {
		var ret int32
		return ret
	}
	return *o.BookId
}

// GetBookIdOk returns a tuple with the BookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameBookResource) GetBookIdOk() (*int32, bool) {
	if o == nil || isNil(o.BookId) {
    return nil, false
	}
	return o.BookId, true
}

// HasBookId returns a boolean if a field has been set.
func (o *RenameBookResource) HasBookId() bool {
	if o != nil && !isNil(o.BookId) {
		return true
	}

	return false
}

// SetBookId gets a reference to the given int32 and assigns it to the BookId field.
func (o *RenameBookResource) SetBookId(v int32) {
	o.BookId = &v
}

// GetBookFileId returns the BookFileId field value if set, zero value otherwise.
func (o *RenameBookResource) GetBookFileId() int32 {
	if o == nil || isNil(o.BookFileId) {
		var ret int32
		return ret
	}
	return *o.BookFileId
}

// GetBookFileIdOk returns a tuple with the BookFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameBookResource) GetBookFileIdOk() (*int32, bool) {
	if o == nil || isNil(o.BookFileId) {
    return nil, false
	}
	return o.BookFileId, true
}

// HasBookFileId returns a boolean if a field has been set.
func (o *RenameBookResource) HasBookFileId() bool {
	if o != nil && !isNil(o.BookFileId) {
		return true
	}

	return false
}

// SetBookFileId gets a reference to the given int32 and assigns it to the BookFileId field.
func (o *RenameBookResource) SetBookFileId(v int32) {
	o.BookFileId = &v
}

// GetExistingPath returns the ExistingPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RenameBookResource) GetExistingPath() string {
	if o == nil || isNil(o.ExistingPath.Get()) {
		var ret string
		return ret
	}
	return *o.ExistingPath.Get()
}

// GetExistingPathOk returns a tuple with the ExistingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameBookResource) GetExistingPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExistingPath.Get(), o.ExistingPath.IsSet()
}

// HasExistingPath returns a boolean if a field has been set.
func (o *RenameBookResource) HasExistingPath() bool {
	if o != nil && o.ExistingPath.IsSet() {
		return true
	}

	return false
}

// SetExistingPath gets a reference to the given NullableString and assigns it to the ExistingPath field.
func (o *RenameBookResource) SetExistingPath(v string) {
	o.ExistingPath.Set(&v)
}
// SetExistingPathNil sets the value for ExistingPath to be an explicit nil
func (o *RenameBookResource) SetExistingPathNil() {
	o.ExistingPath.Set(nil)
}

// UnsetExistingPath ensures that no value is present for ExistingPath, not even an explicit nil
func (o *RenameBookResource) UnsetExistingPath() {
	o.ExistingPath.Unset()
}

// GetNewPath returns the NewPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RenameBookResource) GetNewPath() string {
	if o == nil || isNil(o.NewPath.Get()) {
		var ret string
		return ret
	}
	return *o.NewPath.Get()
}

// GetNewPathOk returns a tuple with the NewPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameBookResource) GetNewPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NewPath.Get(), o.NewPath.IsSet()
}

// HasNewPath returns a boolean if a field has been set.
func (o *RenameBookResource) HasNewPath() bool {
	if o != nil && o.NewPath.IsSet() {
		return true
	}

	return false
}

// SetNewPath gets a reference to the given NullableString and assigns it to the NewPath field.
func (o *RenameBookResource) SetNewPath(v string) {
	o.NewPath.Set(&v)
}
// SetNewPathNil sets the value for NewPath to be an explicit nil
func (o *RenameBookResource) SetNewPathNil() {
	o.NewPath.Set(nil)
}

// UnsetNewPath ensures that no value is present for NewPath, not even an explicit nil
func (o *RenameBookResource) UnsetNewPath() {
	o.NewPath.Unset()
}

func (o RenameBookResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AuthorId) {
		toSerialize["authorId"] = o.AuthorId
	}
	if !isNil(o.BookId) {
		toSerialize["bookId"] = o.BookId
	}
	if !isNil(o.BookFileId) {
		toSerialize["bookFileId"] = o.BookFileId
	}
	if o.ExistingPath.IsSet() {
		toSerialize["existingPath"] = o.ExistingPath.Get()
	}
	if o.NewPath.IsSet() {
		toSerialize["newPath"] = o.NewPath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRenameBookResource struct {
	value *RenameBookResource
	isSet bool
}

func (v NullableRenameBookResource) Get() *RenameBookResource {
	return v.value
}

func (v *NullableRenameBookResource) Set(val *RenameBookResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameBookResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameBookResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameBookResource(val *RenameBookResource) *NullableRenameBookResource {
	return &NullableRenameBookResource{value: val, isSet: true}
}

func (v NullableRenameBookResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameBookResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


