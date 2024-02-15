# SeriesResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]SeriesBookLinkResource**](SeriesBookLinkResource.md) |  | [optional] 

## Methods

### NewSeriesResource

`func NewSeriesResource() *SeriesResource`

NewSeriesResource instantiates a new SeriesResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesResourceWithDefaults

`func NewSeriesResourceWithDefaults() *SeriesResource`

NewSeriesResourceWithDefaults instantiates a new SeriesResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SeriesResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeriesResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeriesResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SeriesResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *SeriesResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeriesResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeriesResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeriesResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SeriesResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SeriesResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SeriesResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeriesResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeriesResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SeriesResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SeriesResource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SeriesResource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLinks

`func (o *SeriesResource) GetLinks() []SeriesBookLinkResource`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SeriesResource) GetLinksOk() (*[]SeriesBookLinkResource, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SeriesResource) SetLinks(v []SeriesBookLinkResource)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SeriesResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *SeriesResource) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *SeriesResource) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


