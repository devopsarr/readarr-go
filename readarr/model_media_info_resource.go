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

// checks if the MediaInfoResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaInfoResource{}

// MediaInfoResource struct for MediaInfoResource
type MediaInfoResource struct {
	Id *int32 `json:"id,omitempty"`
	AudioChannels *float64 `json:"audioChannels,omitempty"`
	AudioBitRate NullableString `json:"audioBitRate,omitempty"`
	AudioCodec NullableString `json:"audioCodec,omitempty"`
	AudioBits NullableString `json:"audioBits,omitempty"`
	AudioSampleRate NullableString `json:"audioSampleRate,omitempty"`
}

// NewMediaInfoResource instantiates a new MediaInfoResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaInfoResource() *MediaInfoResource {
	this := MediaInfoResource{}
	return &this
}

// NewMediaInfoResourceWithDefaults instantiates a new MediaInfoResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaInfoResourceWithDefaults() *MediaInfoResource {
	this := MediaInfoResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MediaInfoResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MediaInfoResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MediaInfoResource) SetId(v int32) {
	o.Id = &v
}

// GetAudioChannels returns the AudioChannels field value if set, zero value otherwise.
func (o *MediaInfoResource) GetAudioChannels() float64 {
	if o == nil || IsNil(o.AudioChannels) {
		var ret float64
		return ret
	}
	return *o.AudioChannels
}

// GetAudioChannelsOk returns a tuple with the AudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetAudioChannelsOk() (*float64, bool) {
	if o == nil || IsNil(o.AudioChannels) {
		return nil, false
	}
	return o.AudioChannels, true
}

// HasAudioChannels returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioChannels() bool {
	if o != nil && !IsNil(o.AudioChannels) {
		return true
	}

	return false
}

// SetAudioChannels gets a reference to the given float64 and assigns it to the AudioChannels field.
func (o *MediaInfoResource) SetAudioChannels(v float64) {
	o.AudioChannels = &v
}

// GetAudioBitRate returns the AudioBitRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetAudioBitRate() string {
	if o == nil || IsNil(o.AudioBitRate.Get()) {
		var ret string
		return ret
	}
	return *o.AudioBitRate.Get()
}

// GetAudioBitRateOk returns a tuple with the AudioBitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetAudioBitRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioBitRate.Get(), o.AudioBitRate.IsSet()
}

// HasAudioBitRate returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioBitRate() bool {
	if o != nil && o.AudioBitRate.IsSet() {
		return true
	}

	return false
}

// SetAudioBitRate gets a reference to the given NullableString and assigns it to the AudioBitRate field.
func (o *MediaInfoResource) SetAudioBitRate(v string) {
	o.AudioBitRate.Set(&v)
}
// SetAudioBitRateNil sets the value for AudioBitRate to be an explicit nil
func (o *MediaInfoResource) SetAudioBitRateNil() {
	o.AudioBitRate.Set(nil)
}

// UnsetAudioBitRate ensures that no value is present for AudioBitRate, not even an explicit nil
func (o *MediaInfoResource) UnsetAudioBitRate() {
	o.AudioBitRate.Unset()
}

// GetAudioCodec returns the AudioCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetAudioCodec() string {
	if o == nil || IsNil(o.AudioCodec.Get()) {
		var ret string
		return ret
	}
	return *o.AudioCodec.Get()
}

// GetAudioCodecOk returns a tuple with the AudioCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetAudioCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioCodec.Get(), o.AudioCodec.IsSet()
}

// HasAudioCodec returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioCodec() bool {
	if o != nil && o.AudioCodec.IsSet() {
		return true
	}

	return false
}

// SetAudioCodec gets a reference to the given NullableString and assigns it to the AudioCodec field.
func (o *MediaInfoResource) SetAudioCodec(v string) {
	o.AudioCodec.Set(&v)
}
// SetAudioCodecNil sets the value for AudioCodec to be an explicit nil
func (o *MediaInfoResource) SetAudioCodecNil() {
	o.AudioCodec.Set(nil)
}

// UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
func (o *MediaInfoResource) UnsetAudioCodec() {
	o.AudioCodec.Unset()
}

// GetAudioBits returns the AudioBits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetAudioBits() string {
	if o == nil || IsNil(o.AudioBits.Get()) {
		var ret string
		return ret
	}
	return *o.AudioBits.Get()
}

// GetAudioBitsOk returns a tuple with the AudioBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetAudioBitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioBits.Get(), o.AudioBits.IsSet()
}

// HasAudioBits returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioBits() bool {
	if o != nil && o.AudioBits.IsSet() {
		return true
	}

	return false
}

// SetAudioBits gets a reference to the given NullableString and assigns it to the AudioBits field.
func (o *MediaInfoResource) SetAudioBits(v string) {
	o.AudioBits.Set(&v)
}
// SetAudioBitsNil sets the value for AudioBits to be an explicit nil
func (o *MediaInfoResource) SetAudioBitsNil() {
	o.AudioBits.Set(nil)
}

// UnsetAudioBits ensures that no value is present for AudioBits, not even an explicit nil
func (o *MediaInfoResource) UnsetAudioBits() {
	o.AudioBits.Unset()
}

// GetAudioSampleRate returns the AudioSampleRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetAudioSampleRate() string {
	if o == nil || IsNil(o.AudioSampleRate.Get()) {
		var ret string
		return ret
	}
	return *o.AudioSampleRate.Get()
}

// GetAudioSampleRateOk returns a tuple with the AudioSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetAudioSampleRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioSampleRate.Get(), o.AudioSampleRate.IsSet()
}

// HasAudioSampleRate returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioSampleRate() bool {
	if o != nil && o.AudioSampleRate.IsSet() {
		return true
	}

	return false
}

// SetAudioSampleRate gets a reference to the given NullableString and assigns it to the AudioSampleRate field.
func (o *MediaInfoResource) SetAudioSampleRate(v string) {
	o.AudioSampleRate.Set(&v)
}
// SetAudioSampleRateNil sets the value for AudioSampleRate to be an explicit nil
func (o *MediaInfoResource) SetAudioSampleRateNil() {
	o.AudioSampleRate.Set(nil)
}

// UnsetAudioSampleRate ensures that no value is present for AudioSampleRate, not even an explicit nil
func (o *MediaInfoResource) UnsetAudioSampleRate() {
	o.AudioSampleRate.Unset()
}

func (o MediaInfoResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaInfoResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AudioChannels) {
		toSerialize["audioChannels"] = o.AudioChannels
	}
	if o.AudioBitRate.IsSet() {
		toSerialize["audioBitRate"] = o.AudioBitRate.Get()
	}
	if o.AudioCodec.IsSet() {
		toSerialize["audioCodec"] = o.AudioCodec.Get()
	}
	if o.AudioBits.IsSet() {
		toSerialize["audioBits"] = o.AudioBits.Get()
	}
	if o.AudioSampleRate.IsSet() {
		toSerialize["audioSampleRate"] = o.AudioSampleRate.Get()
	}
	return toSerialize, nil
}

type NullableMediaInfoResource struct {
	value *MediaInfoResource
	isSet bool
}

func (v NullableMediaInfoResource) Get() *MediaInfoResource {
	return v.value
}

func (v *NullableMediaInfoResource) Set(val *MediaInfoResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInfoResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInfoResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInfoResource(val *MediaInfoResource) *NullableMediaInfoResource {
	return &NullableMediaInfoResource{value: val, isSet: true}
}

func (v NullableMediaInfoResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInfoResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


