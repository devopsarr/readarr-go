# AuthorMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ForeignAuthorId** | Pointer to **NullableString** |  | [optional] 
**TitleSlug** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SortName** | Pointer to **NullableString** |  | [optional] 
**NameLastFirst** | Pointer to **NullableString** |  | [optional] 
**SortNameLastFirst** | Pointer to **NullableString** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Hometown** | Pointer to **NullableString** |  | [optional] 
**Born** | Pointer to **NullableTime** |  | [optional] 
**Died** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to [**AuthorStatusType**](AuthorStatusType.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 

## Methods

### NewAuthorMetadata

`func NewAuthorMetadata() *AuthorMetadata`

NewAuthorMetadata instantiates a new AuthorMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorMetadataWithDefaults

`func NewAuthorMetadataWithDefaults() *AuthorMetadata`

NewAuthorMetadataWithDefaults instantiates a new AuthorMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorMetadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorMetadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorMetadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignAuthorId

`func (o *AuthorMetadata) GetForeignAuthorId() string`

GetForeignAuthorId returns the ForeignAuthorId field if non-nil, zero value otherwise.

### GetForeignAuthorIdOk

`func (o *AuthorMetadata) GetForeignAuthorIdOk() (*string, bool)`

GetForeignAuthorIdOk returns a tuple with the ForeignAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAuthorId

`func (o *AuthorMetadata) SetForeignAuthorId(v string)`

SetForeignAuthorId sets ForeignAuthorId field to given value.

### HasForeignAuthorId

`func (o *AuthorMetadata) HasForeignAuthorId() bool`

HasForeignAuthorId returns a boolean if a field has been set.

### SetForeignAuthorIdNil

`func (o *AuthorMetadata) SetForeignAuthorIdNil(b bool)`

 SetForeignAuthorIdNil sets the value for ForeignAuthorId to be an explicit nil

### UnsetForeignAuthorId
`func (o *AuthorMetadata) UnsetForeignAuthorId()`

UnsetForeignAuthorId ensures that no value is present for ForeignAuthorId, not even an explicit nil
### GetTitleSlug

`func (o *AuthorMetadata) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *AuthorMetadata) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *AuthorMetadata) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *AuthorMetadata) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *AuthorMetadata) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *AuthorMetadata) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetName

`func (o *AuthorMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AuthorMetadata) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthorMetadata) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSortName

`func (o *AuthorMetadata) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *AuthorMetadata) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *AuthorMetadata) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *AuthorMetadata) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### SetSortNameNil

`func (o *AuthorMetadata) SetSortNameNil(b bool)`

 SetSortNameNil sets the value for SortName to be an explicit nil

### UnsetSortName
`func (o *AuthorMetadata) UnsetSortName()`

UnsetSortName ensures that no value is present for SortName, not even an explicit nil
### GetNameLastFirst

`func (o *AuthorMetadata) GetNameLastFirst() string`

GetNameLastFirst returns the NameLastFirst field if non-nil, zero value otherwise.

### GetNameLastFirstOk

`func (o *AuthorMetadata) GetNameLastFirstOk() (*string, bool)`

GetNameLastFirstOk returns a tuple with the NameLastFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLastFirst

`func (o *AuthorMetadata) SetNameLastFirst(v string)`

SetNameLastFirst sets NameLastFirst field to given value.

### HasNameLastFirst

`func (o *AuthorMetadata) HasNameLastFirst() bool`

HasNameLastFirst returns a boolean if a field has been set.

### SetNameLastFirstNil

`func (o *AuthorMetadata) SetNameLastFirstNil(b bool)`

 SetNameLastFirstNil sets the value for NameLastFirst to be an explicit nil

### UnsetNameLastFirst
`func (o *AuthorMetadata) UnsetNameLastFirst()`

UnsetNameLastFirst ensures that no value is present for NameLastFirst, not even an explicit nil
### GetSortNameLastFirst

`func (o *AuthorMetadata) GetSortNameLastFirst() string`

GetSortNameLastFirst returns the SortNameLastFirst field if non-nil, zero value otherwise.

### GetSortNameLastFirstOk

`func (o *AuthorMetadata) GetSortNameLastFirstOk() (*string, bool)`

GetSortNameLastFirstOk returns a tuple with the SortNameLastFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortNameLastFirst

`func (o *AuthorMetadata) SetSortNameLastFirst(v string)`

SetSortNameLastFirst sets SortNameLastFirst field to given value.

### HasSortNameLastFirst

`func (o *AuthorMetadata) HasSortNameLastFirst() bool`

HasSortNameLastFirst returns a boolean if a field has been set.

### SetSortNameLastFirstNil

`func (o *AuthorMetadata) SetSortNameLastFirstNil(b bool)`

 SetSortNameLastFirstNil sets the value for SortNameLastFirst to be an explicit nil

### UnsetSortNameLastFirst
`func (o *AuthorMetadata) UnsetSortNameLastFirst()`

UnsetSortNameLastFirst ensures that no value is present for SortNameLastFirst, not even an explicit nil
### GetAliases

`func (o *AuthorMetadata) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AuthorMetadata) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AuthorMetadata) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AuthorMetadata) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *AuthorMetadata) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *AuthorMetadata) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil
### GetOverview

`func (o *AuthorMetadata) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *AuthorMetadata) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *AuthorMetadata) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *AuthorMetadata) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *AuthorMetadata) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *AuthorMetadata) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetDisambiguation

`func (o *AuthorMetadata) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *AuthorMetadata) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *AuthorMetadata) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *AuthorMetadata) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *AuthorMetadata) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *AuthorMetadata) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetGender

`func (o *AuthorMetadata) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *AuthorMetadata) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *AuthorMetadata) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *AuthorMetadata) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *AuthorMetadata) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *AuthorMetadata) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHometown

`func (o *AuthorMetadata) GetHometown() string`

GetHometown returns the Hometown field if non-nil, zero value otherwise.

### GetHometownOk

`func (o *AuthorMetadata) GetHometownOk() (*string, bool)`

GetHometownOk returns a tuple with the Hometown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHometown

`func (o *AuthorMetadata) SetHometown(v string)`

SetHometown sets Hometown field to given value.

### HasHometown

`func (o *AuthorMetadata) HasHometown() bool`

HasHometown returns a boolean if a field has been set.

### SetHometownNil

`func (o *AuthorMetadata) SetHometownNil(b bool)`

 SetHometownNil sets the value for Hometown to be an explicit nil

### UnsetHometown
`func (o *AuthorMetadata) UnsetHometown()`

UnsetHometown ensures that no value is present for Hometown, not even an explicit nil
### GetBorn

`func (o *AuthorMetadata) GetBorn() time.Time`

GetBorn returns the Born field if non-nil, zero value otherwise.

### GetBornOk

`func (o *AuthorMetadata) GetBornOk() (*time.Time, bool)`

GetBornOk returns a tuple with the Born field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorn

`func (o *AuthorMetadata) SetBorn(v time.Time)`

SetBorn sets Born field to given value.

### HasBorn

`func (o *AuthorMetadata) HasBorn() bool`

HasBorn returns a boolean if a field has been set.

### SetBornNil

`func (o *AuthorMetadata) SetBornNil(b bool)`

 SetBornNil sets the value for Born to be an explicit nil

### UnsetBorn
`func (o *AuthorMetadata) UnsetBorn()`

UnsetBorn ensures that no value is present for Born, not even an explicit nil
### GetDied

`func (o *AuthorMetadata) GetDied() time.Time`

GetDied returns the Died field if non-nil, zero value otherwise.

### GetDiedOk

`func (o *AuthorMetadata) GetDiedOk() (*time.Time, bool)`

GetDiedOk returns a tuple with the Died field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDied

`func (o *AuthorMetadata) SetDied(v time.Time)`

SetDied sets Died field to given value.

### HasDied

`func (o *AuthorMetadata) HasDied() bool`

HasDied returns a boolean if a field has been set.

### SetDiedNil

`func (o *AuthorMetadata) SetDiedNil(b bool)`

 SetDiedNil sets the value for Died to be an explicit nil

### UnsetDied
`func (o *AuthorMetadata) UnsetDied()`

UnsetDied ensures that no value is present for Died, not even an explicit nil
### GetStatus

`func (o *AuthorMetadata) GetStatus() AuthorStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorMetadata) GetStatusOk() (*AuthorStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorMetadata) SetStatus(v AuthorStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImages

`func (o *AuthorMetadata) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AuthorMetadata) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AuthorMetadata) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *AuthorMetadata) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *AuthorMetadata) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *AuthorMetadata) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetLinks

`func (o *AuthorMetadata) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorMetadata) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorMetadata) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorMetadata) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AuthorMetadata) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AuthorMetadata) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetGenres

`func (o *AuthorMetadata) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AuthorMetadata) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AuthorMetadata) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *AuthorMetadata) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *AuthorMetadata) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *AuthorMetadata) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetRatings

`func (o *AuthorMetadata) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *AuthorMetadata) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *AuthorMetadata) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *AuthorMetadata) HasRatings() bool`

HasRatings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


