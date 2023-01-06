/*
Readarr

Readarr API docs

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// NamingConfigResource struct for NamingConfigResource
type NamingConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	RenameBooks *bool `json:"renameBooks,omitempty"`
	ReplaceIllegalCharacters *bool `json:"replaceIllegalCharacters,omitempty"`
	StandardBookFormat NullableString `json:"standardBookFormat,omitempty"`
	AuthorFolderFormat NullableString `json:"authorFolderFormat,omitempty"`
	IncludeAuthorName *bool `json:"includeAuthorName,omitempty"`
	IncludeBookTitle *bool `json:"includeBookTitle,omitempty"`
	IncludeQuality *bool `json:"includeQuality,omitempty"`
	ReplaceSpaces *bool `json:"replaceSpaces,omitempty"`
	Separator NullableString `json:"separator,omitempty"`
	NumberStyle NullableString `json:"numberStyle,omitempty"`
}

// NewNamingConfigResource instantiates a new NamingConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamingConfigResource() *NamingConfigResource {
	this := NamingConfigResource{}
	return &this
}

// NewNamingConfigResourceWithDefaults instantiates a new NamingConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamingConfigResourceWithDefaults() *NamingConfigResource {
	this := NamingConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NamingConfigResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NamingConfigResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NamingConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetRenameBooks returns the RenameBooks field value if set, zero value otherwise.
func (o *NamingConfigResource) GetRenameBooks() bool {
	if o == nil || isNil(o.RenameBooks) {
		var ret bool
		return ret
	}
	return *o.RenameBooks
}

// GetRenameBooksOk returns a tuple with the RenameBooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetRenameBooksOk() (*bool, bool) {
	if o == nil || isNil(o.RenameBooks) {
    return nil, false
	}
	return o.RenameBooks, true
}

// HasRenameBooks returns a boolean if a field has been set.
func (o *NamingConfigResource) HasRenameBooks() bool {
	if o != nil && !isNil(o.RenameBooks) {
		return true
	}

	return false
}

// SetRenameBooks gets a reference to the given bool and assigns it to the RenameBooks field.
func (o *NamingConfigResource) SetRenameBooks(v bool) {
	o.RenameBooks = &v
}

// GetReplaceIllegalCharacters returns the ReplaceIllegalCharacters field value if set, zero value otherwise.
func (o *NamingConfigResource) GetReplaceIllegalCharacters() bool {
	if o == nil || isNil(o.ReplaceIllegalCharacters) {
		var ret bool
		return ret
	}
	return *o.ReplaceIllegalCharacters
}

// GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool) {
	if o == nil || isNil(o.ReplaceIllegalCharacters) {
    return nil, false
	}
	return o.ReplaceIllegalCharacters, true
}

// HasReplaceIllegalCharacters returns a boolean if a field has been set.
func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool {
	if o != nil && !isNil(o.ReplaceIllegalCharacters) {
		return true
	}

	return false
}

// SetReplaceIllegalCharacters gets a reference to the given bool and assigns it to the ReplaceIllegalCharacters field.
func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool) {
	o.ReplaceIllegalCharacters = &v
}

// GetStandardBookFormat returns the StandardBookFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetStandardBookFormat() string {
	if o == nil || isNil(o.StandardBookFormat.Get()) {
		var ret string
		return ret
	}
	return *o.StandardBookFormat.Get()
}

// GetStandardBookFormatOk returns a tuple with the StandardBookFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetStandardBookFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.StandardBookFormat.Get(), o.StandardBookFormat.IsSet()
}

// HasStandardBookFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasStandardBookFormat() bool {
	if o != nil && o.StandardBookFormat.IsSet() {
		return true
	}

	return false
}

// SetStandardBookFormat gets a reference to the given NullableString and assigns it to the StandardBookFormat field.
func (o *NamingConfigResource) SetStandardBookFormat(v string) {
	o.StandardBookFormat.Set(&v)
}
// SetStandardBookFormatNil sets the value for StandardBookFormat to be an explicit nil
func (o *NamingConfigResource) SetStandardBookFormatNil() {
	o.StandardBookFormat.Set(nil)
}

// UnsetStandardBookFormat ensures that no value is present for StandardBookFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetStandardBookFormat() {
	o.StandardBookFormat.Unset()
}

// GetAuthorFolderFormat returns the AuthorFolderFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetAuthorFolderFormat() string {
	if o == nil || isNil(o.AuthorFolderFormat.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorFolderFormat.Get()
}

// GetAuthorFolderFormatOk returns a tuple with the AuthorFolderFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetAuthorFolderFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AuthorFolderFormat.Get(), o.AuthorFolderFormat.IsSet()
}

// HasAuthorFolderFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasAuthorFolderFormat() bool {
	if o != nil && o.AuthorFolderFormat.IsSet() {
		return true
	}

	return false
}

// SetAuthorFolderFormat gets a reference to the given NullableString and assigns it to the AuthorFolderFormat field.
func (o *NamingConfigResource) SetAuthorFolderFormat(v string) {
	o.AuthorFolderFormat.Set(&v)
}
// SetAuthorFolderFormatNil sets the value for AuthorFolderFormat to be an explicit nil
func (o *NamingConfigResource) SetAuthorFolderFormatNil() {
	o.AuthorFolderFormat.Set(nil)
}

// UnsetAuthorFolderFormat ensures that no value is present for AuthorFolderFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetAuthorFolderFormat() {
	o.AuthorFolderFormat.Unset()
}

// GetIncludeAuthorName returns the IncludeAuthorName field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeAuthorName() bool {
	if o == nil || isNil(o.IncludeAuthorName) {
		var ret bool
		return ret
	}
	return *o.IncludeAuthorName
}

// GetIncludeAuthorNameOk returns a tuple with the IncludeAuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeAuthorNameOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeAuthorName) {
    return nil, false
	}
	return o.IncludeAuthorName, true
}

// HasIncludeAuthorName returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeAuthorName() bool {
	if o != nil && !isNil(o.IncludeAuthorName) {
		return true
	}

	return false
}

// SetIncludeAuthorName gets a reference to the given bool and assigns it to the IncludeAuthorName field.
func (o *NamingConfigResource) SetIncludeAuthorName(v bool) {
	o.IncludeAuthorName = &v
}

// GetIncludeBookTitle returns the IncludeBookTitle field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeBookTitle() bool {
	if o == nil || isNil(o.IncludeBookTitle) {
		var ret bool
		return ret
	}
	return *o.IncludeBookTitle
}

// GetIncludeBookTitleOk returns a tuple with the IncludeBookTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeBookTitleOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeBookTitle) {
    return nil, false
	}
	return o.IncludeBookTitle, true
}

// HasIncludeBookTitle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeBookTitle() bool {
	if o != nil && !isNil(o.IncludeBookTitle) {
		return true
	}

	return false
}

// SetIncludeBookTitle gets a reference to the given bool and assigns it to the IncludeBookTitle field.
func (o *NamingConfigResource) SetIncludeBookTitle(v bool) {
	o.IncludeBookTitle = &v
}

// GetIncludeQuality returns the IncludeQuality field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeQuality() bool {
	if o == nil || isNil(o.IncludeQuality) {
		var ret bool
		return ret
	}
	return *o.IncludeQuality
}

// GetIncludeQualityOk returns a tuple with the IncludeQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeQualityOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeQuality) {
    return nil, false
	}
	return o.IncludeQuality, true
}

// HasIncludeQuality returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeQuality() bool {
	if o != nil && !isNil(o.IncludeQuality) {
		return true
	}

	return false
}

// SetIncludeQuality gets a reference to the given bool and assigns it to the IncludeQuality field.
func (o *NamingConfigResource) SetIncludeQuality(v bool) {
	o.IncludeQuality = &v
}

// GetReplaceSpaces returns the ReplaceSpaces field value if set, zero value otherwise.
func (o *NamingConfigResource) GetReplaceSpaces() bool {
	if o == nil || isNil(o.ReplaceSpaces) {
		var ret bool
		return ret
	}
	return *o.ReplaceSpaces
}

// GetReplaceSpacesOk returns a tuple with the ReplaceSpaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetReplaceSpacesOk() (*bool, bool) {
	if o == nil || isNil(o.ReplaceSpaces) {
    return nil, false
	}
	return o.ReplaceSpaces, true
}

// HasReplaceSpaces returns a boolean if a field has been set.
func (o *NamingConfigResource) HasReplaceSpaces() bool {
	if o != nil && !isNil(o.ReplaceSpaces) {
		return true
	}

	return false
}

// SetReplaceSpaces gets a reference to the given bool and assigns it to the ReplaceSpaces field.
func (o *NamingConfigResource) SetReplaceSpaces(v bool) {
	o.ReplaceSpaces = &v
}

// GetSeparator returns the Separator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetSeparator() string {
	if o == nil || isNil(o.Separator.Get()) {
		var ret string
		return ret
	}
	return *o.Separator.Get()
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetSeparatorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Separator.Get(), o.Separator.IsSet()
}

// HasSeparator returns a boolean if a field has been set.
func (o *NamingConfigResource) HasSeparator() bool {
	if o != nil && o.Separator.IsSet() {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given NullableString and assigns it to the Separator field.
func (o *NamingConfigResource) SetSeparator(v string) {
	o.Separator.Set(&v)
}
// SetSeparatorNil sets the value for Separator to be an explicit nil
func (o *NamingConfigResource) SetSeparatorNil() {
	o.Separator.Set(nil)
}

// UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
func (o *NamingConfigResource) UnsetSeparator() {
	o.Separator.Unset()
}

// GetNumberStyle returns the NumberStyle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetNumberStyle() string {
	if o == nil || isNil(o.NumberStyle.Get()) {
		var ret string
		return ret
	}
	return *o.NumberStyle.Get()
}

// GetNumberStyleOk returns a tuple with the NumberStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetNumberStyleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumberStyle.Get(), o.NumberStyle.IsSet()
}

// HasNumberStyle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasNumberStyle() bool {
	if o != nil && o.NumberStyle.IsSet() {
		return true
	}

	return false
}

// SetNumberStyle gets a reference to the given NullableString and assigns it to the NumberStyle field.
func (o *NamingConfigResource) SetNumberStyle(v string) {
	o.NumberStyle.Set(&v)
}
// SetNumberStyleNil sets the value for NumberStyle to be an explicit nil
func (o *NamingConfigResource) SetNumberStyleNil() {
	o.NumberStyle.Set(nil)
}

// UnsetNumberStyle ensures that no value is present for NumberStyle, not even an explicit nil
func (o *NamingConfigResource) UnsetNumberStyle() {
	o.NumberStyle.Unset()
}

func (o NamingConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.RenameBooks) {
		toSerialize["renameBooks"] = o.RenameBooks
	}
	if !isNil(o.ReplaceIllegalCharacters) {
		toSerialize["replaceIllegalCharacters"] = o.ReplaceIllegalCharacters
	}
	if o.StandardBookFormat.IsSet() {
		toSerialize["standardBookFormat"] = o.StandardBookFormat.Get()
	}
	if o.AuthorFolderFormat.IsSet() {
		toSerialize["authorFolderFormat"] = o.AuthorFolderFormat.Get()
	}
	if !isNil(o.IncludeAuthorName) {
		toSerialize["includeAuthorName"] = o.IncludeAuthorName
	}
	if !isNil(o.IncludeBookTitle) {
		toSerialize["includeBookTitle"] = o.IncludeBookTitle
	}
	if !isNil(o.IncludeQuality) {
		toSerialize["includeQuality"] = o.IncludeQuality
	}
	if !isNil(o.ReplaceSpaces) {
		toSerialize["replaceSpaces"] = o.ReplaceSpaces
	}
	if o.Separator.IsSet() {
		toSerialize["separator"] = o.Separator.Get()
	}
	if o.NumberStyle.IsSet() {
		toSerialize["numberStyle"] = o.NumberStyle.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNamingConfigResource struct {
	value *NamingConfigResource
	isSet bool
}

func (v NullableNamingConfigResource) Get() *NamingConfigResource {
	return v.value
}

func (v *NullableNamingConfigResource) Set(val *NamingConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableNamingConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableNamingConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamingConfigResource(val *NamingConfigResource) *NullableNamingConfigResource {
	return &NullableNamingConfigResource{value: val, isSet: true}
}

func (v NullableNamingConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamingConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


