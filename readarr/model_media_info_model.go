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

// MediaInfoModel struct for MediaInfoModel
type MediaInfoModel struct {
	AudioFormat NullableString `json:"audioFormat,omitempty"`
	AudioBitrate *int32 `json:"audioBitrate,omitempty"`
	AudioChannels *int32 `json:"audioChannels,omitempty"`
	AudioBits *int32 `json:"audioBits,omitempty"`
	AudioSampleRate *int32 `json:"audioSampleRate,omitempty"`
}

// NewMediaInfoModel instantiates a new MediaInfoModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaInfoModel() *MediaInfoModel {
	this := MediaInfoModel{}
	return &this
}

// NewMediaInfoModelWithDefaults instantiates a new MediaInfoModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaInfoModelWithDefaults() *MediaInfoModel {
	this := MediaInfoModel{}
	return &this
}

// GetAudioFormat returns the AudioFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoModel) GetAudioFormat() string {
	if o == nil || isNil(o.AudioFormat.Get()) {
		var ret string
		return ret
	}
	return *o.AudioFormat.Get()
}

// GetAudioFormatOk returns a tuple with the AudioFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoModel) GetAudioFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AudioFormat.Get(), o.AudioFormat.IsSet()
}

// HasAudioFormat returns a boolean if a field has been set.
func (o *MediaInfoModel) HasAudioFormat() bool {
	if o != nil && o.AudioFormat.IsSet() {
		return true
	}

	return false
}

// SetAudioFormat gets a reference to the given NullableString and assigns it to the AudioFormat field.
func (o *MediaInfoModel) SetAudioFormat(v string) {
	o.AudioFormat.Set(&v)
}
// SetAudioFormatNil sets the value for AudioFormat to be an explicit nil
func (o *MediaInfoModel) SetAudioFormatNil() {
	o.AudioFormat.Set(nil)
}

// UnsetAudioFormat ensures that no value is present for AudioFormat, not even an explicit nil
func (o *MediaInfoModel) UnsetAudioFormat() {
	o.AudioFormat.Unset()
}

// GetAudioBitrate returns the AudioBitrate field value if set, zero value otherwise.
func (o *MediaInfoModel) GetAudioBitrate() int32 {
	if o == nil || isNil(o.AudioBitrate) {
		var ret int32
		return ret
	}
	return *o.AudioBitrate
}

// GetAudioBitrateOk returns a tuple with the AudioBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoModel) GetAudioBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.AudioBitrate) {
    return nil, false
	}
	return o.AudioBitrate, true
}

// HasAudioBitrate returns a boolean if a field has been set.
func (o *MediaInfoModel) HasAudioBitrate() bool {
	if o != nil && !isNil(o.AudioBitrate) {
		return true
	}

	return false
}

// SetAudioBitrate gets a reference to the given int32 and assigns it to the AudioBitrate field.
func (o *MediaInfoModel) SetAudioBitrate(v int32) {
	o.AudioBitrate = &v
}

// GetAudioChannels returns the AudioChannels field value if set, zero value otherwise.
func (o *MediaInfoModel) GetAudioChannels() int32 {
	if o == nil || isNil(o.AudioChannels) {
		var ret int32
		return ret
	}
	return *o.AudioChannels
}

// GetAudioChannelsOk returns a tuple with the AudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoModel) GetAudioChannelsOk() (*int32, bool) {
	if o == nil || isNil(o.AudioChannels) {
    return nil, false
	}
	return o.AudioChannels, true
}

// HasAudioChannels returns a boolean if a field has been set.
func (o *MediaInfoModel) HasAudioChannels() bool {
	if o != nil && !isNil(o.AudioChannels) {
		return true
	}

	return false
}

// SetAudioChannels gets a reference to the given int32 and assigns it to the AudioChannels field.
func (o *MediaInfoModel) SetAudioChannels(v int32) {
	o.AudioChannels = &v
}

// GetAudioBits returns the AudioBits field value if set, zero value otherwise.
func (o *MediaInfoModel) GetAudioBits() int32 {
	if o == nil || isNil(o.AudioBits) {
		var ret int32
		return ret
	}
	return *o.AudioBits
}

// GetAudioBitsOk returns a tuple with the AudioBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoModel) GetAudioBitsOk() (*int32, bool) {
	if o == nil || isNil(o.AudioBits) {
    return nil, false
	}
	return o.AudioBits, true
}

// HasAudioBits returns a boolean if a field has been set.
func (o *MediaInfoModel) HasAudioBits() bool {
	if o != nil && !isNil(o.AudioBits) {
		return true
	}

	return false
}

// SetAudioBits gets a reference to the given int32 and assigns it to the AudioBits field.
func (o *MediaInfoModel) SetAudioBits(v int32) {
	o.AudioBits = &v
}

// GetAudioSampleRate returns the AudioSampleRate field value if set, zero value otherwise.
func (o *MediaInfoModel) GetAudioSampleRate() int32 {
	if o == nil || isNil(o.AudioSampleRate) {
		var ret int32
		return ret
	}
	return *o.AudioSampleRate
}

// GetAudioSampleRateOk returns a tuple with the AudioSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoModel) GetAudioSampleRateOk() (*int32, bool) {
	if o == nil || isNil(o.AudioSampleRate) {
    return nil, false
	}
	return o.AudioSampleRate, true
}

// HasAudioSampleRate returns a boolean if a field has been set.
func (o *MediaInfoModel) HasAudioSampleRate() bool {
	if o != nil && !isNil(o.AudioSampleRate) {
		return true
	}

	return false
}

// SetAudioSampleRate gets a reference to the given int32 and assigns it to the AudioSampleRate field.
func (o *MediaInfoModel) SetAudioSampleRate(v int32) {
	o.AudioSampleRate = &v
}

func (o MediaInfoModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioFormat.IsSet() {
		toSerialize["audioFormat"] = o.AudioFormat.Get()
	}
	if !isNil(o.AudioBitrate) {
		toSerialize["audioBitrate"] = o.AudioBitrate
	}
	if !isNil(o.AudioChannels) {
		toSerialize["audioChannels"] = o.AudioChannels
	}
	if !isNil(o.AudioBits) {
		toSerialize["audioBits"] = o.AudioBits
	}
	if !isNil(o.AudioSampleRate) {
		toSerialize["audioSampleRate"] = o.AudioSampleRate
	}
	return json.Marshal(toSerialize)
}

type NullableMediaInfoModel struct {
	value *MediaInfoModel
	isSet bool
}

func (v NullableMediaInfoModel) Get() *MediaInfoModel {
	return v.value
}

func (v *NullableMediaInfoModel) Set(val *MediaInfoModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInfoModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInfoModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInfoModel(val *MediaInfoModel) *NullableMediaInfoModel {
	return &NullableMediaInfoModel{value: val, isSet: true}
}

func (v NullableMediaInfoModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInfoModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


