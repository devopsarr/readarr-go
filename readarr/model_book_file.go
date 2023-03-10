/*
Readarr

Readarr API docs

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
	"time"
)

// BookFile struct for BookFile
type BookFile struct {
	Id *int32 `json:"id,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
	DateAdded *time.Time `json:"dateAdded,omitempty"`
	SceneName NullableString `json:"sceneName,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	MediaInfo *MediaInfoModel `json:"mediaInfo,omitempty"`
	EditionId *int32 `json:"editionId,omitempty"`
	CalibreId *int32 `json:"calibreId,omitempty"`
	Part *int32 `json:"part,omitempty"`
	Author *AuthorLazyLoaded `json:"author,omitempty"`
	Edition *EditionLazyLoaded `json:"edition,omitempty"`
	PartCount *int32 `json:"partCount,omitempty"`
}

// NewBookFile instantiates a new BookFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookFile() *BookFile {
	this := BookFile{}
	return &this
}

// NewBookFileWithDefaults instantiates a new BookFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookFileWithDefaults() *BookFile {
	this := BookFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BookFile) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BookFile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BookFile) SetId(v int32) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookFile) GetPath() string {
	if o == nil || isNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookFile) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *BookFile) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *BookFile) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *BookFile) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *BookFile) UnsetPath() {
	o.Path.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BookFile) GetSize() int64 {
	if o == nil || isNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetSizeOk() (*int64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BookFile) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *BookFile) SetSize(v int64) {
	o.Size = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *BookFile) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
    return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *BookFile) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *BookFile) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *BookFile) GetDateAdded() time.Time {
	if o == nil || isNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateAdded) {
    return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *BookFile) HasDateAdded() bool {
	if o != nil && !isNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *BookFile) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookFile) GetSceneName() string {
	if o == nil || isNil(o.SceneName.Get()) {
		var ret string
		return ret
	}
	return *o.SceneName.Get()
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookFile) GetSceneNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SceneName.Get(), o.SceneName.IsSet()
}

// HasSceneName returns a boolean if a field has been set.
func (o *BookFile) HasSceneName() bool {
	if o != nil && o.SceneName.IsSet() {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given NullableString and assigns it to the SceneName field.
func (o *BookFile) SetSceneName(v string) {
	o.SceneName.Set(&v)
}
// SetSceneNameNil sets the value for SceneName to be an explicit nil
func (o *BookFile) SetSceneNameNil() {
	o.SceneName.Set(nil)
}

// UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
func (o *BookFile) UnsetSceneName() {
	o.SceneName.Unset()
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookFile) GetReleaseGroup() string {
	if o == nil || isNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookFile) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *BookFile) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *BookFile) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *BookFile) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *BookFile) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *BookFile) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *BookFile) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *BookFile) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *BookFile) GetMediaInfo() MediaInfoModel {
	if o == nil || isNil(o.MediaInfo) {
		var ret MediaInfoModel
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetMediaInfoOk() (*MediaInfoModel, bool) {
	if o == nil || isNil(o.MediaInfo) {
    return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *BookFile) HasMediaInfo() bool {
	if o != nil && !isNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfoModel and assigns it to the MediaInfo field.
func (o *BookFile) SetMediaInfo(v MediaInfoModel) {
	o.MediaInfo = &v
}

// GetEditionId returns the EditionId field value if set, zero value otherwise.
func (o *BookFile) GetEditionId() int32 {
	if o == nil || isNil(o.EditionId) {
		var ret int32
		return ret
	}
	return *o.EditionId
}

// GetEditionIdOk returns a tuple with the EditionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetEditionIdOk() (*int32, bool) {
	if o == nil || isNil(o.EditionId) {
    return nil, false
	}
	return o.EditionId, true
}

// HasEditionId returns a boolean if a field has been set.
func (o *BookFile) HasEditionId() bool {
	if o != nil && !isNil(o.EditionId) {
		return true
	}

	return false
}

// SetEditionId gets a reference to the given int32 and assigns it to the EditionId field.
func (o *BookFile) SetEditionId(v int32) {
	o.EditionId = &v
}

// GetCalibreId returns the CalibreId field value if set, zero value otherwise.
func (o *BookFile) GetCalibreId() int32 {
	if o == nil || isNil(o.CalibreId) {
		var ret int32
		return ret
	}
	return *o.CalibreId
}

// GetCalibreIdOk returns a tuple with the CalibreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetCalibreIdOk() (*int32, bool) {
	if o == nil || isNil(o.CalibreId) {
    return nil, false
	}
	return o.CalibreId, true
}

// HasCalibreId returns a boolean if a field has been set.
func (o *BookFile) HasCalibreId() bool {
	if o != nil && !isNil(o.CalibreId) {
		return true
	}

	return false
}

// SetCalibreId gets a reference to the given int32 and assigns it to the CalibreId field.
func (o *BookFile) SetCalibreId(v int32) {
	o.CalibreId = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *BookFile) GetPart() int32 {
	if o == nil || isNil(o.Part) {
		var ret int32
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetPartOk() (*int32, bool) {
	if o == nil || isNil(o.Part) {
    return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *BookFile) HasPart() bool {
	if o != nil && !isNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given int32 and assigns it to the Part field.
func (o *BookFile) SetPart(v int32) {
	o.Part = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *BookFile) GetAuthor() AuthorLazyLoaded {
	if o == nil || isNil(o.Author) {
		var ret AuthorLazyLoaded
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetAuthorOk() (*AuthorLazyLoaded, bool) {
	if o == nil || isNil(o.Author) {
    return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *BookFile) HasAuthor() bool {
	if o != nil && !isNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given AuthorLazyLoaded and assigns it to the Author field.
func (o *BookFile) SetAuthor(v AuthorLazyLoaded) {
	o.Author = &v
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *BookFile) GetEdition() EditionLazyLoaded {
	if o == nil || isNil(o.Edition) {
		var ret EditionLazyLoaded
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetEditionOk() (*EditionLazyLoaded, bool) {
	if o == nil || isNil(o.Edition) {
    return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *BookFile) HasEdition() bool {
	if o != nil && !isNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given EditionLazyLoaded and assigns it to the Edition field.
func (o *BookFile) SetEdition(v EditionLazyLoaded) {
	o.Edition = &v
}

// GetPartCount returns the PartCount field value if set, zero value otherwise.
func (o *BookFile) GetPartCount() int32 {
	if o == nil || isNil(o.PartCount) {
		var ret int32
		return ret
	}
	return *o.PartCount
}

// GetPartCountOk returns a tuple with the PartCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookFile) GetPartCountOk() (*int32, bool) {
	if o == nil || isNil(o.PartCount) {
    return nil, false
	}
	return o.PartCount, true
}

// HasPartCount returns a boolean if a field has been set.
func (o *BookFile) HasPartCount() bool {
	if o != nil && !isNil(o.PartCount) {
		return true
	}

	return false
}

// SetPartCount gets a reference to the given int32 and assigns it to the PartCount field.
func (o *BookFile) SetPartCount(v int32) {
	o.PartCount = &v
}

func (o BookFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.DateAdded) {
		toSerialize["dateAdded"] = o.DateAdded
	}
	if o.SceneName.IsSet() {
		toSerialize["sceneName"] = o.SceneName.Get()
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !isNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	if !isNil(o.EditionId) {
		toSerialize["editionId"] = o.EditionId
	}
	if !isNil(o.CalibreId) {
		toSerialize["calibreId"] = o.CalibreId
	}
	if !isNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	if !isNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !isNil(o.Edition) {
		toSerialize["edition"] = o.Edition
	}
	if !isNil(o.PartCount) {
		toSerialize["partCount"] = o.PartCount
	}
	return json.Marshal(toSerialize)
}

type NullableBookFile struct {
	value *BookFile
	isSet bool
}

func (v NullableBookFile) Get() *BookFile {
	return v.value
}

func (v *NullableBookFile) Set(val *BookFile) {
	v.value = val
	v.isSet = true
}

func (v NullableBookFile) IsSet() bool {
	return v.isSet
}

func (v *NullableBookFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookFile(val *BookFile) *NullableBookFile {
	return &NullableBookFile{value: val, isSet: true}
}

func (v NullableBookFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


