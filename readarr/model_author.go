/*
Readarr

Readarr API docs

API version: v0.4.10.2734
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
	"time"
)

// checks if the Author type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Author{}

// Author struct for Author
type Author struct {
	Id *int32 `json:"id,omitempty"`
	AuthorMetadataId *int32 `json:"authorMetadataId,omitempty"`
	CleanName NullableString `json:"cleanName,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	MonitorNewItems *NewItemMonitorTypes `json:"monitorNewItems,omitempty"`
	LastInfoSync NullableTime `json:"lastInfoSync,omitempty"`
	Path NullableString `json:"path,omitempty"`
	RootFolderPath NullableString `json:"rootFolderPath,omitempty"`
	Added *time.Time `json:"added,omitempty"`
	QualityProfileId *int32 `json:"qualityProfileId,omitempty"`
	MetadataProfileId *int32 `json:"metadataProfileId,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	AddOptions *AddAuthorOptions `json:"addOptions,omitempty"`
	Metadata *AuthorMetadataLazyLoaded `json:"metadata,omitempty"`
	QualityProfile *QualityProfileLazyLoaded `json:"qualityProfile,omitempty"`
	MetadataProfile *MetadataProfileLazyLoaded `json:"metadataProfile,omitempty"`
	Books *BookListLazyLoaded `json:"books,omitempty"`
	Series *SeriesListLazyLoaded `json:"series,omitempty"`
	Name NullableString `json:"name,omitempty"`
	ForeignAuthorId NullableString `json:"foreignAuthorId,omitempty"`
}

// NewAuthor instantiates a new Author object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthor() *Author {
	this := Author{}
	return &this
}

// NewAuthorWithDefaults instantiates a new Author object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorWithDefaults() *Author {
	this := Author{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Author) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Author) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Author) SetId(v int32) {
	o.Id = &v
}

// GetAuthorMetadataId returns the AuthorMetadataId field value if set, zero value otherwise.
func (o *Author) GetAuthorMetadataId() int32 {
	if o == nil || IsNil(o.AuthorMetadataId) {
		var ret int32
		return ret
	}
	return *o.AuthorMetadataId
}

// GetAuthorMetadataIdOk returns a tuple with the AuthorMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetAuthorMetadataIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuthorMetadataId) {
		return nil, false
	}
	return o.AuthorMetadataId, true
}

// HasAuthorMetadataId returns a boolean if a field has been set.
func (o *Author) HasAuthorMetadataId() bool {
	if o != nil && !IsNil(o.AuthorMetadataId) {
		return true
	}

	return false
}

// SetAuthorMetadataId gets a reference to the given int32 and assigns it to the AuthorMetadataId field.
func (o *Author) SetAuthorMetadataId(v int32) {
	o.AuthorMetadataId = &v
}

// GetCleanName returns the CleanName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Author) GetCleanName() string {
	if o == nil || IsNil(o.CleanName.Get()) {
		var ret string
		return ret
	}
	return *o.CleanName.Get()
}

// GetCleanNameOk returns a tuple with the CleanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Author) GetCleanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CleanName.Get(), o.CleanName.IsSet()
}

// HasCleanName returns a boolean if a field has been set.
func (o *Author) HasCleanName() bool {
	if o != nil && o.CleanName.IsSet() {
		return true
	}

	return false
}

// SetCleanName gets a reference to the given NullableString and assigns it to the CleanName field.
func (o *Author) SetCleanName(v string) {
	o.CleanName.Set(&v)
}
// SetCleanNameNil sets the value for CleanName to be an explicit nil
func (o *Author) SetCleanNameNil() {
	o.CleanName.Set(nil)
}

// UnsetCleanName ensures that no value is present for CleanName, not even an explicit nil
func (o *Author) UnsetCleanName() {
	o.CleanName.Unset()
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *Author) GetMonitored() bool {
	if o == nil || IsNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *Author) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *Author) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetMonitorNewItems returns the MonitorNewItems field value if set, zero value otherwise.
func (o *Author) GetMonitorNewItems() NewItemMonitorTypes {
	if o == nil || IsNil(o.MonitorNewItems) {
		var ret NewItemMonitorTypes
		return ret
	}
	return *o.MonitorNewItems
}

// GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool) {
	if o == nil || IsNil(o.MonitorNewItems) {
		return nil, false
	}
	return o.MonitorNewItems, true
}

// HasMonitorNewItems returns a boolean if a field has been set.
func (o *Author) HasMonitorNewItems() bool {
	if o != nil && !IsNil(o.MonitorNewItems) {
		return true
	}

	return false
}

// SetMonitorNewItems gets a reference to the given NewItemMonitorTypes and assigns it to the MonitorNewItems field.
func (o *Author) SetMonitorNewItems(v NewItemMonitorTypes) {
	o.MonitorNewItems = &v
}

// GetLastInfoSync returns the LastInfoSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Author) GetLastInfoSync() time.Time {
	if o == nil || IsNil(o.LastInfoSync.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastInfoSync.Get()
}

// GetLastInfoSyncOk returns a tuple with the LastInfoSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Author) GetLastInfoSyncOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastInfoSync.Get(), o.LastInfoSync.IsSet()
}

// HasLastInfoSync returns a boolean if a field has been set.
func (o *Author) HasLastInfoSync() bool {
	if o != nil && o.LastInfoSync.IsSet() {
		return true
	}

	return false
}

// SetLastInfoSync gets a reference to the given NullableTime and assigns it to the LastInfoSync field.
func (o *Author) SetLastInfoSync(v time.Time) {
	o.LastInfoSync.Set(&v)
}
// SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil
func (o *Author) SetLastInfoSyncNil() {
	o.LastInfoSync.Set(nil)
}

// UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
func (o *Author) UnsetLastInfoSync() {
	o.LastInfoSync.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Author) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Author) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *Author) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *Author) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *Author) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *Author) UnsetPath() {
	o.Path.Unset()
}

// GetRootFolderPath returns the RootFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Author) GetRootFolderPath() string {
	if o == nil || IsNil(o.RootFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderPath.Get()
}

// GetRootFolderPathOk returns a tuple with the RootFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Author) GetRootFolderPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootFolderPath.Get(), o.RootFolderPath.IsSet()
}

// HasRootFolderPath returns a boolean if a field has been set.
func (o *Author) HasRootFolderPath() bool {
	if o != nil && o.RootFolderPath.IsSet() {
		return true
	}

	return false
}

// SetRootFolderPath gets a reference to the given NullableString and assigns it to the RootFolderPath field.
func (o *Author) SetRootFolderPath(v string) {
	o.RootFolderPath.Set(&v)
}
// SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil
func (o *Author) SetRootFolderPathNil() {
	o.RootFolderPath.Set(nil)
}

// UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
func (o *Author) UnsetRootFolderPath() {
	o.RootFolderPath.Unset()
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *Author) GetAdded() time.Time {
	if o == nil || IsNil(o.Added) {
		var ret time.Time
		return ret
	}
	return *o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetAddedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Added) {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *Author) HasAdded() bool {
	if o != nil && !IsNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given time.Time and assigns it to the Added field.
func (o *Author) SetAdded(v time.Time) {
	o.Added = &v
}

// GetQualityProfileId returns the QualityProfileId field value if set, zero value otherwise.
func (o *Author) GetQualityProfileId() int32 {
	if o == nil || IsNil(o.QualityProfileId) {
		var ret int32
		return ret
	}
	return *o.QualityProfileId
}

// GetQualityProfileIdOk returns a tuple with the QualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetQualityProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.QualityProfileId) {
		return nil, false
	}
	return o.QualityProfileId, true
}

// HasQualityProfileId returns a boolean if a field has been set.
func (o *Author) HasQualityProfileId() bool {
	if o != nil && !IsNil(o.QualityProfileId) {
		return true
	}

	return false
}

// SetQualityProfileId gets a reference to the given int32 and assigns it to the QualityProfileId field.
func (o *Author) SetQualityProfileId(v int32) {
	o.QualityProfileId = &v
}

// GetMetadataProfileId returns the MetadataProfileId field value if set, zero value otherwise.
func (o *Author) GetMetadataProfileId() int32 {
	if o == nil || IsNil(o.MetadataProfileId) {
		var ret int32
		return ret
	}
	return *o.MetadataProfileId
}

// GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetMetadataProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MetadataProfileId) {
		return nil, false
	}
	return o.MetadataProfileId, true
}

// HasMetadataProfileId returns a boolean if a field has been set.
func (o *Author) HasMetadataProfileId() bool {
	if o != nil && !IsNil(o.MetadataProfileId) {
		return true
	}

	return false
}

// SetMetadataProfileId gets a reference to the given int32 and assigns it to the MetadataProfileId field.
func (o *Author) SetMetadataProfileId(v int32) {
	o.MetadataProfileId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Author) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Author) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Author) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *Author) SetTags(v []int32) {
	o.Tags = v
}

// GetAddOptions returns the AddOptions field value if set, zero value otherwise.
func (o *Author) GetAddOptions() AddAuthorOptions {
	if o == nil || IsNil(o.AddOptions) {
		var ret AddAuthorOptions
		return ret
	}
	return *o.AddOptions
}

// GetAddOptionsOk returns a tuple with the AddOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetAddOptionsOk() (*AddAuthorOptions, bool) {
	if o == nil || IsNil(o.AddOptions) {
		return nil, false
	}
	return o.AddOptions, true
}

// HasAddOptions returns a boolean if a field has been set.
func (o *Author) HasAddOptions() bool {
	if o != nil && !IsNil(o.AddOptions) {
		return true
	}

	return false
}

// SetAddOptions gets a reference to the given AddAuthorOptions and assigns it to the AddOptions field.
func (o *Author) SetAddOptions(v AddAuthorOptions) {
	o.AddOptions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Author) GetMetadata() AuthorMetadataLazyLoaded {
	if o == nil || IsNil(o.Metadata) {
		var ret AuthorMetadataLazyLoaded
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetMetadataOk() (*AuthorMetadataLazyLoaded, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Author) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given AuthorMetadataLazyLoaded and assigns it to the Metadata field.
func (o *Author) SetMetadata(v AuthorMetadataLazyLoaded) {
	o.Metadata = &v
}

// GetQualityProfile returns the QualityProfile field value if set, zero value otherwise.
func (o *Author) GetQualityProfile() QualityProfileLazyLoaded {
	if o == nil || IsNil(o.QualityProfile) {
		var ret QualityProfileLazyLoaded
		return ret
	}
	return *o.QualityProfile
}

// GetQualityProfileOk returns a tuple with the QualityProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetQualityProfileOk() (*QualityProfileLazyLoaded, bool) {
	if o == nil || IsNil(o.QualityProfile) {
		return nil, false
	}
	return o.QualityProfile, true
}

// HasQualityProfile returns a boolean if a field has been set.
func (o *Author) HasQualityProfile() bool {
	if o != nil && !IsNil(o.QualityProfile) {
		return true
	}

	return false
}

// SetQualityProfile gets a reference to the given QualityProfileLazyLoaded and assigns it to the QualityProfile field.
func (o *Author) SetQualityProfile(v QualityProfileLazyLoaded) {
	o.QualityProfile = &v
}

// GetMetadataProfile returns the MetadataProfile field value if set, zero value otherwise.
func (o *Author) GetMetadataProfile() MetadataProfileLazyLoaded {
	if o == nil || IsNil(o.MetadataProfile) {
		var ret MetadataProfileLazyLoaded
		return ret
	}
	return *o.MetadataProfile
}

// GetMetadataProfileOk returns a tuple with the MetadataProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetMetadataProfileOk() (*MetadataProfileLazyLoaded, bool) {
	if o == nil || IsNil(o.MetadataProfile) {
		return nil, false
	}
	return o.MetadataProfile, true
}

// HasMetadataProfile returns a boolean if a field has been set.
func (o *Author) HasMetadataProfile() bool {
	if o != nil && !IsNil(o.MetadataProfile) {
		return true
	}

	return false
}

// SetMetadataProfile gets a reference to the given MetadataProfileLazyLoaded and assigns it to the MetadataProfile field.
func (o *Author) SetMetadataProfile(v MetadataProfileLazyLoaded) {
	o.MetadataProfile = &v
}

// GetBooks returns the Books field value if set, zero value otherwise.
func (o *Author) GetBooks() BookListLazyLoaded {
	if o == nil || IsNil(o.Books) {
		var ret BookListLazyLoaded
		return ret
	}
	return *o.Books
}

// GetBooksOk returns a tuple with the Books field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetBooksOk() (*BookListLazyLoaded, bool) {
	if o == nil || IsNil(o.Books) {
		return nil, false
	}
	return o.Books, true
}

// HasBooks returns a boolean if a field has been set.
func (o *Author) HasBooks() bool {
	if o != nil && !IsNil(o.Books) {
		return true
	}

	return false
}

// SetBooks gets a reference to the given BookListLazyLoaded and assigns it to the Books field.
func (o *Author) SetBooks(v BookListLazyLoaded) {
	o.Books = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *Author) GetSeries() SeriesListLazyLoaded {
	if o == nil || IsNil(o.Series) {
		var ret SeriesListLazyLoaded
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Author) GetSeriesOk() (*SeriesListLazyLoaded, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *Author) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given SeriesListLazyLoaded and assigns it to the Series field.
func (o *Author) SetSeries(v SeriesListLazyLoaded) {
	o.Series = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Author) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Author) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Author) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Author) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Author) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Author) UnsetName() {
	o.Name.Unset()
}

// GetForeignAuthorId returns the ForeignAuthorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Author) GetForeignAuthorId() string {
	if o == nil || IsNil(o.ForeignAuthorId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignAuthorId.Get()
}

// GetForeignAuthorIdOk returns a tuple with the ForeignAuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Author) GetForeignAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForeignAuthorId.Get(), o.ForeignAuthorId.IsSet()
}

// HasForeignAuthorId returns a boolean if a field has been set.
func (o *Author) HasForeignAuthorId() bool {
	if o != nil && o.ForeignAuthorId.IsSet() {
		return true
	}

	return false
}

// SetForeignAuthorId gets a reference to the given NullableString and assigns it to the ForeignAuthorId field.
func (o *Author) SetForeignAuthorId(v string) {
	o.ForeignAuthorId.Set(&v)
}
// SetForeignAuthorIdNil sets the value for ForeignAuthorId to be an explicit nil
func (o *Author) SetForeignAuthorIdNil() {
	o.ForeignAuthorId.Set(nil)
}

// UnsetForeignAuthorId ensures that no value is present for ForeignAuthorId, not even an explicit nil
func (o *Author) UnsetForeignAuthorId() {
	o.ForeignAuthorId.Unset()
}

func (o Author) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Author) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AuthorMetadataId) {
		toSerialize["authorMetadataId"] = o.AuthorMetadataId
	}
	if o.CleanName.IsSet() {
		toSerialize["cleanName"] = o.CleanName.Get()
	}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !IsNil(o.MonitorNewItems) {
		toSerialize["monitorNewItems"] = o.MonitorNewItems
	}
	if o.LastInfoSync.IsSet() {
		toSerialize["lastInfoSync"] = o.LastInfoSync.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.RootFolderPath.IsSet() {
		toSerialize["rootFolderPath"] = o.RootFolderPath.Get()
	}
	if !IsNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !IsNil(o.QualityProfileId) {
		toSerialize["qualityProfileId"] = o.QualityProfileId
	}
	if !IsNil(o.MetadataProfileId) {
		toSerialize["metadataProfileId"] = o.MetadataProfileId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.AddOptions) {
		toSerialize["addOptions"] = o.AddOptions
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.QualityProfile) {
		toSerialize["qualityProfile"] = o.QualityProfile
	}
	if !IsNil(o.MetadataProfile) {
		toSerialize["metadataProfile"] = o.MetadataProfile
	}
	if !IsNil(o.Books) {
		toSerialize["books"] = o.Books
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ForeignAuthorId.IsSet() {
		toSerialize["foreignAuthorId"] = o.ForeignAuthorId.Get()
	}
	return toSerialize, nil
}

type NullableAuthor struct {
	value *Author
	isSet bool
}

func (v NullableAuthor) Get() *Author {
	return v.value
}

func (v *NullableAuthor) Set(val *Author) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthor(val *Author) *NullableAuthor {
	return &NullableAuthor{value: val, isSet: true}
}

func (v NullableAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


