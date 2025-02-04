/*
Readarr

Readarr API docs

API version: v0.4.10.2734
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// checks if the ManualImportUpdateResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualImportUpdateResource{}

// ManualImportUpdateResource struct for ManualImportUpdateResource
type ManualImportUpdateResource struct {
	Id *int32 `json:"id,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Name NullableString `json:"name,omitempty"`
	AuthorId NullableInt32 `json:"authorId,omitempty"`
	BookId NullableInt32 `json:"bookId,omitempty"`
	ForeignEditionId NullableString `json:"foreignEditionId,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	IndexerFlags *int32 `json:"indexerFlags,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	AdditionalFile *bool `json:"additionalFile,omitempty"`
	ReplaceExistingFiles *bool `json:"replaceExistingFiles,omitempty"`
	DisableReleaseSwitching *bool `json:"disableReleaseSwitching,omitempty"`
	Rejections []Rejection `json:"rejections,omitempty"`
}

// NewManualImportUpdateResource instantiates a new ManualImportUpdateResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualImportUpdateResource() *ManualImportUpdateResource {
	this := ManualImportUpdateResource{}
	return &this
}

// NewManualImportUpdateResourceWithDefaults instantiates a new ManualImportUpdateResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualImportUpdateResourceWithDefaults() *ManualImportUpdateResource {
	this := ManualImportUpdateResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManualImportUpdateResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportUpdateResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManualImportUpdateResource) SetId(v int32) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ManualImportUpdateResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *ManualImportUpdateResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ManualImportUpdateResource) UnsetPath() {
	o.Path.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ManualImportUpdateResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ManualImportUpdateResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ManualImportUpdateResource) UnsetName() {
	o.Name.Unset()
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetAuthorId() int32 {
	if o == nil || IsNil(o.AuthorId.Get()) {
		var ret int32
		return ret
	}
	return *o.AuthorId.Get()
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetAuthorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorId.Get(), o.AuthorId.IsSet()
}

// HasAuthorId returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasAuthorId() bool {
	if o != nil && o.AuthorId.IsSet() {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given NullableInt32 and assigns it to the AuthorId field.
func (o *ManualImportUpdateResource) SetAuthorId(v int32) {
	o.AuthorId.Set(&v)
}
// SetAuthorIdNil sets the value for AuthorId to be an explicit nil
func (o *ManualImportUpdateResource) SetAuthorIdNil() {
	o.AuthorId.Set(nil)
}

// UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
func (o *ManualImportUpdateResource) UnsetAuthorId() {
	o.AuthorId.Unset()
}

// GetBookId returns the BookId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetBookId() int32 {
	if o == nil || IsNil(o.BookId.Get()) {
		var ret int32
		return ret
	}
	return *o.BookId.Get()
}

// GetBookIdOk returns a tuple with the BookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetBookIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BookId.Get(), o.BookId.IsSet()
}

// HasBookId returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasBookId() bool {
	if o != nil && o.BookId.IsSet() {
		return true
	}

	return false
}

// SetBookId gets a reference to the given NullableInt32 and assigns it to the BookId field.
func (o *ManualImportUpdateResource) SetBookId(v int32) {
	o.BookId.Set(&v)
}
// SetBookIdNil sets the value for BookId to be an explicit nil
func (o *ManualImportUpdateResource) SetBookIdNil() {
	o.BookId.Set(nil)
}

// UnsetBookId ensures that no value is present for BookId, not even an explicit nil
func (o *ManualImportUpdateResource) UnsetBookId() {
	o.BookId.Unset()
}

// GetForeignEditionId returns the ForeignEditionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetForeignEditionId() string {
	if o == nil || IsNil(o.ForeignEditionId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignEditionId.Get()
}

// GetForeignEditionIdOk returns a tuple with the ForeignEditionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetForeignEditionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForeignEditionId.Get(), o.ForeignEditionId.IsSet()
}

// HasForeignEditionId returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasForeignEditionId() bool {
	if o != nil && o.ForeignEditionId.IsSet() {
		return true
	}

	return false
}

// SetForeignEditionId gets a reference to the given NullableString and assigns it to the ForeignEditionId field.
func (o *ManualImportUpdateResource) SetForeignEditionId(v string) {
	o.ForeignEditionId.Set(&v)
}
// SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil
func (o *ManualImportUpdateResource) SetForeignEditionIdNil() {
	o.ForeignEditionId.Set(nil)
}

// UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
func (o *ManualImportUpdateResource) UnsetForeignEditionId() {
	o.ForeignEditionId.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ManualImportUpdateResource) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportUpdateResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ManualImportUpdateResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetReleaseGroup() string {
	if o == nil || IsNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ManualImportUpdateResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ManualImportUpdateResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ManualImportUpdateResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetIndexerFlags returns the IndexerFlags field value if set, zero value otherwise.
func (o *ManualImportUpdateResource) GetIndexerFlags() int32 {
	if o == nil || IsNil(o.IndexerFlags) {
		var ret int32
		return ret
	}
	return *o.IndexerFlags
}

// GetIndexerFlagsOk returns a tuple with the IndexerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportUpdateResource) GetIndexerFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.IndexerFlags) {
		return nil, false
	}
	return o.IndexerFlags, true
}

// HasIndexerFlags returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasIndexerFlags() bool {
	if o != nil && !IsNil(o.IndexerFlags) {
		return true
	}

	return false
}

// SetIndexerFlags gets a reference to the given int32 and assigns it to the IndexerFlags field.
func (o *ManualImportUpdateResource) SetIndexerFlags(v int32) {
	o.IndexerFlags = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetDownloadId() string {
	if o == nil || IsNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *ManualImportUpdateResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *ManualImportUpdateResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *ManualImportUpdateResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetAdditionalFile returns the AdditionalFile field value if set, zero value otherwise.
func (o *ManualImportUpdateResource) GetAdditionalFile() bool {
	if o == nil || IsNil(o.AdditionalFile) {
		var ret bool
		return ret
	}
	return *o.AdditionalFile
}

// GetAdditionalFileOk returns a tuple with the AdditionalFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportUpdateResource) GetAdditionalFileOk() (*bool, bool) {
	if o == nil || IsNil(o.AdditionalFile) {
		return nil, false
	}
	return o.AdditionalFile, true
}

// HasAdditionalFile returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasAdditionalFile() bool {
	if o != nil && !IsNil(o.AdditionalFile) {
		return true
	}

	return false
}

// SetAdditionalFile gets a reference to the given bool and assigns it to the AdditionalFile field.
func (o *ManualImportUpdateResource) SetAdditionalFile(v bool) {
	o.AdditionalFile = &v
}

// GetReplaceExistingFiles returns the ReplaceExistingFiles field value if set, zero value otherwise.
func (o *ManualImportUpdateResource) GetReplaceExistingFiles() bool {
	if o == nil || IsNil(o.ReplaceExistingFiles) {
		var ret bool
		return ret
	}
	return *o.ReplaceExistingFiles
}

// GetReplaceExistingFilesOk returns a tuple with the ReplaceExistingFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportUpdateResource) GetReplaceExistingFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceExistingFiles) {
		return nil, false
	}
	return o.ReplaceExistingFiles, true
}

// HasReplaceExistingFiles returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasReplaceExistingFiles() bool {
	if o != nil && !IsNil(o.ReplaceExistingFiles) {
		return true
	}

	return false
}

// SetReplaceExistingFiles gets a reference to the given bool and assigns it to the ReplaceExistingFiles field.
func (o *ManualImportUpdateResource) SetReplaceExistingFiles(v bool) {
	o.ReplaceExistingFiles = &v
}

// GetDisableReleaseSwitching returns the DisableReleaseSwitching field value if set, zero value otherwise.
func (o *ManualImportUpdateResource) GetDisableReleaseSwitching() bool {
	if o == nil || IsNil(o.DisableReleaseSwitching) {
		var ret bool
		return ret
	}
	return *o.DisableReleaseSwitching
}

// GetDisableReleaseSwitchingOk returns a tuple with the DisableReleaseSwitching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportUpdateResource) GetDisableReleaseSwitchingOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableReleaseSwitching) {
		return nil, false
	}
	return o.DisableReleaseSwitching, true
}

// HasDisableReleaseSwitching returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasDisableReleaseSwitching() bool {
	if o != nil && !IsNil(o.DisableReleaseSwitching) {
		return true
	}

	return false
}

// SetDisableReleaseSwitching gets a reference to the given bool and assigns it to the DisableReleaseSwitching field.
func (o *ManualImportUpdateResource) SetDisableReleaseSwitching(v bool) {
	o.DisableReleaseSwitching = &v
}

// GetRejections returns the Rejections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportUpdateResource) GetRejections() []Rejection {
	if o == nil {
		var ret []Rejection
		return ret
	}
	return o.Rejections
}

// GetRejectionsOk returns a tuple with the Rejections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportUpdateResource) GetRejectionsOk() ([]Rejection, bool) {
	if o == nil || IsNil(o.Rejections) {
		return nil, false
	}
	return o.Rejections, true
}

// HasRejections returns a boolean if a field has been set.
func (o *ManualImportUpdateResource) HasRejections() bool {
	if o != nil && !IsNil(o.Rejections) {
		return true
	}

	return false
}

// SetRejections gets a reference to the given []Rejection and assigns it to the Rejections field.
func (o *ManualImportUpdateResource) SetRejections(v []Rejection) {
	o.Rejections = v
}

func (o ManualImportUpdateResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualImportUpdateResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AuthorId.IsSet() {
		toSerialize["authorId"] = o.AuthorId.Get()
	}
	if o.BookId.IsSet() {
		toSerialize["bookId"] = o.BookId.Get()
	}
	if o.ForeignEditionId.IsSet() {
		toSerialize["foreignEditionId"] = o.ForeignEditionId.Get()
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if !IsNil(o.IndexerFlags) {
		toSerialize["indexerFlags"] = o.IndexerFlags
	}
	if o.DownloadId.IsSet() {
		toSerialize["downloadId"] = o.DownloadId.Get()
	}
	if !IsNil(o.AdditionalFile) {
		toSerialize["additionalFile"] = o.AdditionalFile
	}
	if !IsNil(o.ReplaceExistingFiles) {
		toSerialize["replaceExistingFiles"] = o.ReplaceExistingFiles
	}
	if !IsNil(o.DisableReleaseSwitching) {
		toSerialize["disableReleaseSwitching"] = o.DisableReleaseSwitching
	}
	if o.Rejections != nil {
		toSerialize["rejections"] = o.Rejections
	}
	return toSerialize, nil
}

type NullableManualImportUpdateResource struct {
	value *ManualImportUpdateResource
	isSet bool
}

func (v NullableManualImportUpdateResource) Get() *ManualImportUpdateResource {
	return v.value
}

func (v *NullableManualImportUpdateResource) Set(val *ManualImportUpdateResource) {
	v.value = val
	v.isSet = true
}

func (v NullableManualImportUpdateResource) IsSet() bool {
	return v.isSet
}

func (v *NullableManualImportUpdateResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualImportUpdateResource(val *ManualImportUpdateResource) *NullableManualImportUpdateResource {
	return &NullableManualImportUpdateResource{value: val, isSet: true}
}

func (v NullableManualImportUpdateResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualImportUpdateResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


