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

// ParsedBookInfo struct for ParsedBookInfo
type ParsedBookInfo struct {
	BookTitle NullableString `json:"bookTitle,omitempty"`
	AuthorName NullableString `json:"authorName,omitempty"`
	AuthorTitleInfo *AuthorTitleInfo `json:"authorTitleInfo,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	ReleaseDate NullableString `json:"releaseDate,omitempty"`
	Discography *bool `json:"discography,omitempty"`
	DiscographyStart *int32 `json:"discographyStart,omitempty"`
	DiscographyEnd *int32 `json:"discographyEnd,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	ReleaseHash NullableString `json:"releaseHash,omitempty"`
	ReleaseVersion NullableString `json:"releaseVersion,omitempty"`
}

// NewParsedBookInfo instantiates a new ParsedBookInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParsedBookInfo() *ParsedBookInfo {
	this := ParsedBookInfo{}
	return &this
}

// NewParsedBookInfoWithDefaults instantiates a new ParsedBookInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParsedBookInfoWithDefaults() *ParsedBookInfo {
	this := ParsedBookInfo{}
	return &this
}

// GetBookTitle returns the BookTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedBookInfo) GetBookTitle() string {
	if o == nil || isNil(o.BookTitle.Get()) {
		var ret string
		return ret
	}
	return *o.BookTitle.Get()
}

// GetBookTitleOk returns a tuple with the BookTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedBookInfo) GetBookTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BookTitle.Get(), o.BookTitle.IsSet()
}

// HasBookTitle returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasBookTitle() bool {
	if o != nil && o.BookTitle.IsSet() {
		return true
	}

	return false
}

// SetBookTitle gets a reference to the given NullableString and assigns it to the BookTitle field.
func (o *ParsedBookInfo) SetBookTitle(v string) {
	o.BookTitle.Set(&v)
}
// SetBookTitleNil sets the value for BookTitle to be an explicit nil
func (o *ParsedBookInfo) SetBookTitleNil() {
	o.BookTitle.Set(nil)
}

// UnsetBookTitle ensures that no value is present for BookTitle, not even an explicit nil
func (o *ParsedBookInfo) UnsetBookTitle() {
	o.BookTitle.Unset()
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedBookInfo) GetAuthorName() string {
	if o == nil || isNil(o.AuthorName.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorName.Get()
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedBookInfo) GetAuthorNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AuthorName.Get(), o.AuthorName.IsSet()
}

// HasAuthorName returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasAuthorName() bool {
	if o != nil && o.AuthorName.IsSet() {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given NullableString and assigns it to the AuthorName field.
func (o *ParsedBookInfo) SetAuthorName(v string) {
	o.AuthorName.Set(&v)
}
// SetAuthorNameNil sets the value for AuthorName to be an explicit nil
func (o *ParsedBookInfo) SetAuthorNameNil() {
	o.AuthorName.Set(nil)
}

// UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
func (o *ParsedBookInfo) UnsetAuthorName() {
	o.AuthorName.Unset()
}

// GetAuthorTitleInfo returns the AuthorTitleInfo field value if set, zero value otherwise.
func (o *ParsedBookInfo) GetAuthorTitleInfo() AuthorTitleInfo {
	if o == nil || isNil(o.AuthorTitleInfo) {
		var ret AuthorTitleInfo
		return ret
	}
	return *o.AuthorTitleInfo
}

// GetAuthorTitleInfoOk returns a tuple with the AuthorTitleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedBookInfo) GetAuthorTitleInfoOk() (*AuthorTitleInfo, bool) {
	if o == nil || isNil(o.AuthorTitleInfo) {
    return nil, false
	}
	return o.AuthorTitleInfo, true
}

// HasAuthorTitleInfo returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasAuthorTitleInfo() bool {
	if o != nil && !isNil(o.AuthorTitleInfo) {
		return true
	}

	return false
}

// SetAuthorTitleInfo gets a reference to the given AuthorTitleInfo and assigns it to the AuthorTitleInfo field.
func (o *ParsedBookInfo) SetAuthorTitleInfo(v AuthorTitleInfo) {
	o.AuthorTitleInfo = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ParsedBookInfo) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedBookInfo) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ParsedBookInfo) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedBookInfo) GetReleaseDate() string {
	if o == nil || isNil(o.ReleaseDate.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedBookInfo) GetReleaseDateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate.IsSet() {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given NullableString and assigns it to the ReleaseDate field.
func (o *ParsedBookInfo) SetReleaseDate(v string) {
	o.ReleaseDate.Set(&v)
}
// SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil
func (o *ParsedBookInfo) SetReleaseDateNil() {
	o.ReleaseDate.Set(nil)
}

// UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
func (o *ParsedBookInfo) UnsetReleaseDate() {
	o.ReleaseDate.Unset()
}

// GetDiscography returns the Discography field value if set, zero value otherwise.
func (o *ParsedBookInfo) GetDiscography() bool {
	if o == nil || isNil(o.Discography) {
		var ret bool
		return ret
	}
	return *o.Discography
}

// GetDiscographyOk returns a tuple with the Discography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedBookInfo) GetDiscographyOk() (*bool, bool) {
	if o == nil || isNil(o.Discography) {
    return nil, false
	}
	return o.Discography, true
}

// HasDiscography returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasDiscography() bool {
	if o != nil && !isNil(o.Discography) {
		return true
	}

	return false
}

// SetDiscography gets a reference to the given bool and assigns it to the Discography field.
func (o *ParsedBookInfo) SetDiscography(v bool) {
	o.Discography = &v
}

// GetDiscographyStart returns the DiscographyStart field value if set, zero value otherwise.
func (o *ParsedBookInfo) GetDiscographyStart() int32 {
	if o == nil || isNil(o.DiscographyStart) {
		var ret int32
		return ret
	}
	return *o.DiscographyStart
}

// GetDiscographyStartOk returns a tuple with the DiscographyStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedBookInfo) GetDiscographyStartOk() (*int32, bool) {
	if o == nil || isNil(o.DiscographyStart) {
    return nil, false
	}
	return o.DiscographyStart, true
}

// HasDiscographyStart returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasDiscographyStart() bool {
	if o != nil && !isNil(o.DiscographyStart) {
		return true
	}

	return false
}

// SetDiscographyStart gets a reference to the given int32 and assigns it to the DiscographyStart field.
func (o *ParsedBookInfo) SetDiscographyStart(v int32) {
	o.DiscographyStart = &v
}

// GetDiscographyEnd returns the DiscographyEnd field value if set, zero value otherwise.
func (o *ParsedBookInfo) GetDiscographyEnd() int32 {
	if o == nil || isNil(o.DiscographyEnd) {
		var ret int32
		return ret
	}
	return *o.DiscographyEnd
}

// GetDiscographyEndOk returns a tuple with the DiscographyEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedBookInfo) GetDiscographyEndOk() (*int32, bool) {
	if o == nil || isNil(o.DiscographyEnd) {
    return nil, false
	}
	return o.DiscographyEnd, true
}

// HasDiscographyEnd returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasDiscographyEnd() bool {
	if o != nil && !isNil(o.DiscographyEnd) {
		return true
	}

	return false
}

// SetDiscographyEnd gets a reference to the given int32 and assigns it to the DiscographyEnd field.
func (o *ParsedBookInfo) SetDiscographyEnd(v int32) {
	o.DiscographyEnd = &v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedBookInfo) GetReleaseGroup() string {
	if o == nil || isNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedBookInfo) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ParsedBookInfo) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ParsedBookInfo) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ParsedBookInfo) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetReleaseHash returns the ReleaseHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedBookInfo) GetReleaseHash() string {
	if o == nil || isNil(o.ReleaseHash.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseHash.Get()
}

// GetReleaseHashOk returns a tuple with the ReleaseHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedBookInfo) GetReleaseHashOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseHash.Get(), o.ReleaseHash.IsSet()
}

// HasReleaseHash returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasReleaseHash() bool {
	if o != nil && o.ReleaseHash.IsSet() {
		return true
	}

	return false
}

// SetReleaseHash gets a reference to the given NullableString and assigns it to the ReleaseHash field.
func (o *ParsedBookInfo) SetReleaseHash(v string) {
	o.ReleaseHash.Set(&v)
}
// SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil
func (o *ParsedBookInfo) SetReleaseHashNil() {
	o.ReleaseHash.Set(nil)
}

// UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
func (o *ParsedBookInfo) UnsetReleaseHash() {
	o.ReleaseHash.Unset()
}

// GetReleaseVersion returns the ReleaseVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedBookInfo) GetReleaseVersion() string {
	if o == nil || isNil(o.ReleaseVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseVersion.Get()
}

// GetReleaseVersionOk returns a tuple with the ReleaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedBookInfo) GetReleaseVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseVersion.Get(), o.ReleaseVersion.IsSet()
}

// HasReleaseVersion returns a boolean if a field has been set.
func (o *ParsedBookInfo) HasReleaseVersion() bool {
	if o != nil && o.ReleaseVersion.IsSet() {
		return true
	}

	return false
}

// SetReleaseVersion gets a reference to the given NullableString and assigns it to the ReleaseVersion field.
func (o *ParsedBookInfo) SetReleaseVersion(v string) {
	o.ReleaseVersion.Set(&v)
}
// SetReleaseVersionNil sets the value for ReleaseVersion to be an explicit nil
func (o *ParsedBookInfo) SetReleaseVersionNil() {
	o.ReleaseVersion.Set(nil)
}

// UnsetReleaseVersion ensures that no value is present for ReleaseVersion, not even an explicit nil
func (o *ParsedBookInfo) UnsetReleaseVersion() {
	o.ReleaseVersion.Unset()
}

func (o ParsedBookInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BookTitle.IsSet() {
		toSerialize["bookTitle"] = o.BookTitle.Get()
	}
	if o.AuthorName.IsSet() {
		toSerialize["authorName"] = o.AuthorName.Get()
	}
	if !isNil(o.AuthorTitleInfo) {
		toSerialize["authorTitleInfo"] = o.AuthorTitleInfo
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.ReleaseDate.IsSet() {
		toSerialize["releaseDate"] = o.ReleaseDate.Get()
	}
	if !isNil(o.Discography) {
		toSerialize["discography"] = o.Discography
	}
	if !isNil(o.DiscographyStart) {
		toSerialize["discographyStart"] = o.DiscographyStart
	}
	if !isNil(o.DiscographyEnd) {
		toSerialize["discographyEnd"] = o.DiscographyEnd
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if o.ReleaseHash.IsSet() {
		toSerialize["releaseHash"] = o.ReleaseHash.Get()
	}
	if o.ReleaseVersion.IsSet() {
		toSerialize["releaseVersion"] = o.ReleaseVersion.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableParsedBookInfo struct {
	value *ParsedBookInfo
	isSet bool
}

func (v NullableParsedBookInfo) Get() *ParsedBookInfo {
	return v.value
}

func (v *NullableParsedBookInfo) Set(val *ParsedBookInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableParsedBookInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableParsedBookInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParsedBookInfo(val *ParsedBookInfo) *NullableParsedBookInfo {
	return &NullableParsedBookInfo{value: val, isSet: true}
}

func (v NullableParsedBookInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParsedBookInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


