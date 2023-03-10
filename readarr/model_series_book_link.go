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

// SeriesBookLink struct for SeriesBookLink
type SeriesBookLink struct {
	Id *int32 `json:"id,omitempty"`
	Position NullableString `json:"position,omitempty"`
	SeriesId *int32 `json:"seriesId,omitempty"`
	BookId *int32 `json:"bookId,omitempty"`
	IsPrimary *bool `json:"isPrimary,omitempty"`
	Series *SeriesLazyLoaded `json:"series,omitempty"`
	Book *BookLazyLoaded `json:"book,omitempty"`
}

// NewSeriesBookLink instantiates a new SeriesBookLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesBookLink() *SeriesBookLink {
	this := SeriesBookLink{}
	return &this
}

// NewSeriesBookLinkWithDefaults instantiates a new SeriesBookLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesBookLinkWithDefaults() *SeriesBookLink {
	this := SeriesBookLink{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SeriesBookLink) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesBookLink) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SeriesBookLink) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SeriesBookLink) SetId(v int32) {
	o.Id = &v
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesBookLink) GetPosition() string {
	if o == nil || isNil(o.Position.Get()) {
		var ret string
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesBookLink) GetPositionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *SeriesBookLink) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableString and assigns it to the Position field.
func (o *SeriesBookLink) SetPosition(v string) {
	o.Position.Set(&v)
}
// SetPositionNil sets the value for Position to be an explicit nil
func (o *SeriesBookLink) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *SeriesBookLink) UnsetPosition() {
	o.Position.Unset()
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *SeriesBookLink) GetSeriesId() int32 {
	if o == nil || isNil(o.SeriesId) {
		var ret int32
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesBookLink) GetSeriesIdOk() (*int32, bool) {
	if o == nil || isNil(o.SeriesId) {
    return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *SeriesBookLink) HasSeriesId() bool {
	if o != nil && !isNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given int32 and assigns it to the SeriesId field.
func (o *SeriesBookLink) SetSeriesId(v int32) {
	o.SeriesId = &v
}

// GetBookId returns the BookId field value if set, zero value otherwise.
func (o *SeriesBookLink) GetBookId() int32 {
	if o == nil || isNil(o.BookId) {
		var ret int32
		return ret
	}
	return *o.BookId
}

// GetBookIdOk returns a tuple with the BookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesBookLink) GetBookIdOk() (*int32, bool) {
	if o == nil || isNil(o.BookId) {
    return nil, false
	}
	return o.BookId, true
}

// HasBookId returns a boolean if a field has been set.
func (o *SeriesBookLink) HasBookId() bool {
	if o != nil && !isNil(o.BookId) {
		return true
	}

	return false
}

// SetBookId gets a reference to the given int32 and assigns it to the BookId field.
func (o *SeriesBookLink) SetBookId(v int32) {
	o.BookId = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *SeriesBookLink) GetIsPrimary() bool {
	if o == nil || isNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesBookLink) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrimary) {
    return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *SeriesBookLink) HasIsPrimary() bool {
	if o != nil && !isNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *SeriesBookLink) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *SeriesBookLink) GetSeries() SeriesLazyLoaded {
	if o == nil || isNil(o.Series) {
		var ret SeriesLazyLoaded
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesBookLink) GetSeriesOk() (*SeriesLazyLoaded, bool) {
	if o == nil || isNil(o.Series) {
    return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *SeriesBookLink) HasSeries() bool {
	if o != nil && !isNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given SeriesLazyLoaded and assigns it to the Series field.
func (o *SeriesBookLink) SetSeries(v SeriesLazyLoaded) {
	o.Series = &v
}

// GetBook returns the Book field value if set, zero value otherwise.
func (o *SeriesBookLink) GetBook() BookLazyLoaded {
	if o == nil || isNil(o.Book) {
		var ret BookLazyLoaded
		return ret
	}
	return *o.Book
}

// GetBookOk returns a tuple with the Book field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesBookLink) GetBookOk() (*BookLazyLoaded, bool) {
	if o == nil || isNil(o.Book) {
    return nil, false
	}
	return o.Book, true
}

// HasBook returns a boolean if a field has been set.
func (o *SeriesBookLink) HasBook() bool {
	if o != nil && !isNil(o.Book) {
		return true
	}

	return false
}

// SetBook gets a reference to the given BookLazyLoaded and assigns it to the Book field.
func (o *SeriesBookLink) SetBook(v BookLazyLoaded) {
	o.Book = &v
}

func (o SeriesBookLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if !isNil(o.SeriesId) {
		toSerialize["seriesId"] = o.SeriesId
	}
	if !isNil(o.BookId) {
		toSerialize["bookId"] = o.BookId
	}
	if !isNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if !isNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !isNil(o.Book) {
		toSerialize["book"] = o.Book
	}
	return json.Marshal(toSerialize)
}

type NullableSeriesBookLink struct {
	value *SeriesBookLink
	isSet bool
}

func (v NullableSeriesBookLink) Get() *SeriesBookLink {
	return v.value
}

func (v *NullableSeriesBookLink) Set(val *SeriesBookLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesBookLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesBookLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesBookLink(val *SeriesBookLink) *NullableSeriesBookLink {
	return &NullableSeriesBookLink{value: val, isSet: true}
}

func (v NullableSeriesBookLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesBookLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


