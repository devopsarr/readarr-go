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

// checks if the MetadataProviderConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataProviderConfigResource{}

// MetadataProviderConfigResource struct for MetadataProviderConfigResource
type MetadataProviderConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	WriteAudioTags *WriteAudioTagsType `json:"writeAudioTags,omitempty"`
	ScrubAudioTags *bool `json:"scrubAudioTags,omitempty"`
	WriteBookTags *WriteBookTagsType `json:"writeBookTags,omitempty"`
	UpdateCovers *bool `json:"updateCovers,omitempty"`
	EmbedMetadata *bool `json:"embedMetadata,omitempty"`
}

// NewMetadataProviderConfigResource instantiates a new MetadataProviderConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataProviderConfigResource() *MetadataProviderConfigResource {
	this := MetadataProviderConfigResource{}
	return &this
}

// NewMetadataProviderConfigResourceWithDefaults instantiates a new MetadataProviderConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataProviderConfigResourceWithDefaults() *MetadataProviderConfigResource {
	this := MetadataProviderConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MetadataProviderConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetWriteAudioTags returns the WriteAudioTags field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetWriteAudioTags() WriteAudioTagsType {
	if o == nil || IsNil(o.WriteAudioTags) {
		var ret WriteAudioTagsType
		return ret
	}
	return *o.WriteAudioTags
}

// GetWriteAudioTagsOk returns a tuple with the WriteAudioTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetWriteAudioTagsOk() (*WriteAudioTagsType, bool) {
	if o == nil || IsNil(o.WriteAudioTags) {
		return nil, false
	}
	return o.WriteAudioTags, true
}

// HasWriteAudioTags returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasWriteAudioTags() bool {
	if o != nil && !IsNil(o.WriteAudioTags) {
		return true
	}

	return false
}

// SetWriteAudioTags gets a reference to the given WriteAudioTagsType and assigns it to the WriteAudioTags field.
func (o *MetadataProviderConfigResource) SetWriteAudioTags(v WriteAudioTagsType) {
	o.WriteAudioTags = &v
}

// GetScrubAudioTags returns the ScrubAudioTags field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetScrubAudioTags() bool {
	if o == nil || IsNil(o.ScrubAudioTags) {
		var ret bool
		return ret
	}
	return *o.ScrubAudioTags
}

// GetScrubAudioTagsOk returns a tuple with the ScrubAudioTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetScrubAudioTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.ScrubAudioTags) {
		return nil, false
	}
	return o.ScrubAudioTags, true
}

// HasScrubAudioTags returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasScrubAudioTags() bool {
	if o != nil && !IsNil(o.ScrubAudioTags) {
		return true
	}

	return false
}

// SetScrubAudioTags gets a reference to the given bool and assigns it to the ScrubAudioTags field.
func (o *MetadataProviderConfigResource) SetScrubAudioTags(v bool) {
	o.ScrubAudioTags = &v
}

// GetWriteBookTags returns the WriteBookTags field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetWriteBookTags() WriteBookTagsType {
	if o == nil || IsNil(o.WriteBookTags) {
		var ret WriteBookTagsType
		return ret
	}
	return *o.WriteBookTags
}

// GetWriteBookTagsOk returns a tuple with the WriteBookTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetWriteBookTagsOk() (*WriteBookTagsType, bool) {
	if o == nil || IsNil(o.WriteBookTags) {
		return nil, false
	}
	return o.WriteBookTags, true
}

// HasWriteBookTags returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasWriteBookTags() bool {
	if o != nil && !IsNil(o.WriteBookTags) {
		return true
	}

	return false
}

// SetWriteBookTags gets a reference to the given WriteBookTagsType and assigns it to the WriteBookTags field.
func (o *MetadataProviderConfigResource) SetWriteBookTags(v WriteBookTagsType) {
	o.WriteBookTags = &v
}

// GetUpdateCovers returns the UpdateCovers field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetUpdateCovers() bool {
	if o == nil || IsNil(o.UpdateCovers) {
		var ret bool
		return ret
	}
	return *o.UpdateCovers
}

// GetUpdateCoversOk returns a tuple with the UpdateCovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetUpdateCoversOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateCovers) {
		return nil, false
	}
	return o.UpdateCovers, true
}

// HasUpdateCovers returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasUpdateCovers() bool {
	if o != nil && !IsNil(o.UpdateCovers) {
		return true
	}

	return false
}

// SetUpdateCovers gets a reference to the given bool and assigns it to the UpdateCovers field.
func (o *MetadataProviderConfigResource) SetUpdateCovers(v bool) {
	o.UpdateCovers = &v
}

// GetEmbedMetadata returns the EmbedMetadata field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetEmbedMetadata() bool {
	if o == nil || IsNil(o.EmbedMetadata) {
		var ret bool
		return ret
	}
	return *o.EmbedMetadata
}

// GetEmbedMetadataOk returns a tuple with the EmbedMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetEmbedMetadataOk() (*bool, bool) {
	if o == nil || IsNil(o.EmbedMetadata) {
		return nil, false
	}
	return o.EmbedMetadata, true
}

// HasEmbedMetadata returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasEmbedMetadata() bool {
	if o != nil && !IsNil(o.EmbedMetadata) {
		return true
	}

	return false
}

// SetEmbedMetadata gets a reference to the given bool and assigns it to the EmbedMetadata field.
func (o *MetadataProviderConfigResource) SetEmbedMetadata(v bool) {
	o.EmbedMetadata = &v
}

func (o MetadataProviderConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataProviderConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.WriteAudioTags) {
		toSerialize["writeAudioTags"] = o.WriteAudioTags
	}
	if !IsNil(o.ScrubAudioTags) {
		toSerialize["scrubAudioTags"] = o.ScrubAudioTags
	}
	if !IsNil(o.WriteBookTags) {
		toSerialize["writeBookTags"] = o.WriteBookTags
	}
	if !IsNil(o.UpdateCovers) {
		toSerialize["updateCovers"] = o.UpdateCovers
	}
	if !IsNil(o.EmbedMetadata) {
		toSerialize["embedMetadata"] = o.EmbedMetadata
	}
	return toSerialize, nil
}

type NullableMetadataProviderConfigResource struct {
	value *MetadataProviderConfigResource
	isSet bool
}

func (v NullableMetadataProviderConfigResource) Get() *MetadataProviderConfigResource {
	return v.value
}

func (v *NullableMetadataProviderConfigResource) Set(val *MetadataProviderConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataProviderConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataProviderConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataProviderConfigResource(val *MetadataProviderConfigResource) *NullableMetadataProviderConfigResource {
	return &NullableMetadataProviderConfigResource{value: val, isSet: true}
}

func (v NullableMetadataProviderConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataProviderConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


