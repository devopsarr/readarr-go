# SeriesBookLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **NullableString** |  | [optional] 
**SeriesId** | Pointer to **int32** |  | [optional] 
**BookId** | Pointer to **int32** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**Series** | Pointer to [**SeriesLazyLoaded**](SeriesLazyLoaded.md) |  | [optional] 
**Book** | Pointer to [**BookLazyLoaded**](BookLazyLoaded.md) |  | [optional] 

## Methods

### NewSeriesBookLink

`func NewSeriesBookLink() *SeriesBookLink`

NewSeriesBookLink instantiates a new SeriesBookLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesBookLinkWithDefaults

`func NewSeriesBookLinkWithDefaults() *SeriesBookLink`

NewSeriesBookLinkWithDefaults instantiates a new SeriesBookLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SeriesBookLink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeriesBookLink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeriesBookLink) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SeriesBookLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *SeriesBookLink) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SeriesBookLink) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SeriesBookLink) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SeriesBookLink) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *SeriesBookLink) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *SeriesBookLink) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetSeriesId

`func (o *SeriesBookLink) GetSeriesId() int32`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SeriesBookLink) GetSeriesIdOk() (*int32, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SeriesBookLink) SetSeriesId(v int32)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *SeriesBookLink) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetBookId

`func (o *SeriesBookLink) GetBookId() int32`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *SeriesBookLink) GetBookIdOk() (*int32, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *SeriesBookLink) SetBookId(v int32)`

SetBookId sets BookId field to given value.

### HasBookId

`func (o *SeriesBookLink) HasBookId() bool`

HasBookId returns a boolean if a field has been set.

### GetIsPrimary

`func (o *SeriesBookLink) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *SeriesBookLink) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *SeriesBookLink) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *SeriesBookLink) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetSeries

`func (o *SeriesBookLink) GetSeries() SeriesLazyLoaded`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SeriesBookLink) GetSeriesOk() (*SeriesLazyLoaded, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SeriesBookLink) SetSeries(v SeriesLazyLoaded)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *SeriesBookLink) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetBook

`func (o *SeriesBookLink) GetBook() BookLazyLoaded`

GetBook returns the Book field if non-nil, zero value otherwise.

### GetBookOk

`func (o *SeriesBookLink) GetBookOk() (*BookLazyLoaded, bool)`

GetBookOk returns a tuple with the Book field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBook

`func (o *SeriesBookLink) SetBook(v BookLazyLoaded)`

SetBook sets Book field to given value.

### HasBook

`func (o *SeriesBookLink) HasBook() bool`

HasBook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


