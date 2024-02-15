# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ForeignSeriesId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Numbered** | Pointer to **bool** |  | [optional] 
**WorkCount** | Pointer to **int32** |  | [optional] 
**PrimaryWorkCount** | Pointer to **int32** |  | [optional] 
**LinkItems** | Pointer to [**SeriesBookLinkListLazyLoaded**](SeriesBookLinkListLazyLoaded.md) |  | [optional] 
**Books** | Pointer to [**BookListLazyLoaded**](BookListLazyLoaded.md) |  | [optional] 
**ForeignAuthorId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSeries

`func NewSeries() *Series`

NewSeries instantiates a new Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesWithDefaults

`func NewSeriesWithDefaults() *Series`

NewSeriesWithDefaults instantiates a new Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Series) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Series) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Series) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Series) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignSeriesId

`func (o *Series) GetForeignSeriesId() string`

GetForeignSeriesId returns the ForeignSeriesId field if non-nil, zero value otherwise.

### GetForeignSeriesIdOk

`func (o *Series) GetForeignSeriesIdOk() (*string, bool)`

GetForeignSeriesIdOk returns a tuple with the ForeignSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignSeriesId

`func (o *Series) SetForeignSeriesId(v string)`

SetForeignSeriesId sets ForeignSeriesId field to given value.

### HasForeignSeriesId

`func (o *Series) HasForeignSeriesId() bool`

HasForeignSeriesId returns a boolean if a field has been set.

### SetForeignSeriesIdNil

`func (o *Series) SetForeignSeriesIdNil(b bool)`

 SetForeignSeriesIdNil sets the value for ForeignSeriesId to be an explicit nil

### UnsetForeignSeriesId
`func (o *Series) UnsetForeignSeriesId()`

UnsetForeignSeriesId ensures that no value is present for ForeignSeriesId, not even an explicit nil
### GetTitle

`func (o *Series) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Series) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Series) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Series) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Series) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Series) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *Series) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Series) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Series) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Series) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Series) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Series) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNumbered

`func (o *Series) GetNumbered() bool`

GetNumbered returns the Numbered field if non-nil, zero value otherwise.

### GetNumberedOk

`func (o *Series) GetNumberedOk() (*bool, bool)`

GetNumberedOk returns a tuple with the Numbered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbered

`func (o *Series) SetNumbered(v bool)`

SetNumbered sets Numbered field to given value.

### HasNumbered

`func (o *Series) HasNumbered() bool`

HasNumbered returns a boolean if a field has been set.

### GetWorkCount

`func (o *Series) GetWorkCount() int32`

GetWorkCount returns the WorkCount field if non-nil, zero value otherwise.

### GetWorkCountOk

`func (o *Series) GetWorkCountOk() (*int32, bool)`

GetWorkCountOk returns a tuple with the WorkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkCount

`func (o *Series) SetWorkCount(v int32)`

SetWorkCount sets WorkCount field to given value.

### HasWorkCount

`func (o *Series) HasWorkCount() bool`

HasWorkCount returns a boolean if a field has been set.

### GetPrimaryWorkCount

`func (o *Series) GetPrimaryWorkCount() int32`

GetPrimaryWorkCount returns the PrimaryWorkCount field if non-nil, zero value otherwise.

### GetPrimaryWorkCountOk

`func (o *Series) GetPrimaryWorkCountOk() (*int32, bool)`

GetPrimaryWorkCountOk returns a tuple with the PrimaryWorkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWorkCount

`func (o *Series) SetPrimaryWorkCount(v int32)`

SetPrimaryWorkCount sets PrimaryWorkCount field to given value.

### HasPrimaryWorkCount

`func (o *Series) HasPrimaryWorkCount() bool`

HasPrimaryWorkCount returns a boolean if a field has been set.

### GetLinkItems

`func (o *Series) GetLinkItems() SeriesBookLinkListLazyLoaded`

GetLinkItems returns the LinkItems field if non-nil, zero value otherwise.

### GetLinkItemsOk

`func (o *Series) GetLinkItemsOk() (*SeriesBookLinkListLazyLoaded, bool)`

GetLinkItemsOk returns a tuple with the LinkItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkItems

`func (o *Series) SetLinkItems(v SeriesBookLinkListLazyLoaded)`

SetLinkItems sets LinkItems field to given value.

### HasLinkItems

`func (o *Series) HasLinkItems() bool`

HasLinkItems returns a boolean if a field has been set.

### GetBooks

`func (o *Series) GetBooks() BookListLazyLoaded`

GetBooks returns the Books field if non-nil, zero value otherwise.

### GetBooksOk

`func (o *Series) GetBooksOk() (*BookListLazyLoaded, bool)`

GetBooksOk returns a tuple with the Books field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooks

`func (o *Series) SetBooks(v BookListLazyLoaded)`

SetBooks sets Books field to given value.

### HasBooks

`func (o *Series) HasBooks() bool`

HasBooks returns a boolean if a field has been set.

### GetForeignAuthorId

`func (o *Series) GetForeignAuthorId() string`

GetForeignAuthorId returns the ForeignAuthorId field if non-nil, zero value otherwise.

### GetForeignAuthorIdOk

`func (o *Series) GetForeignAuthorIdOk() (*string, bool)`

GetForeignAuthorIdOk returns a tuple with the ForeignAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAuthorId

`func (o *Series) SetForeignAuthorId(v string)`

SetForeignAuthorId sets ForeignAuthorId field to given value.

### HasForeignAuthorId

`func (o *Series) HasForeignAuthorId() bool`

HasForeignAuthorId returns a boolean if a field has been set.

### SetForeignAuthorIdNil

`func (o *Series) SetForeignAuthorIdNil(b bool)`

 SetForeignAuthorIdNil sets the value for ForeignAuthorId to be an explicit nil

### UnsetForeignAuthorId
`func (o *Series) UnsetForeignAuthorId()`

UnsetForeignAuthorId ensures that no value is present for ForeignAuthorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


