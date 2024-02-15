# RootFolderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**DefaultMetadataProfileId** | Pointer to **int32** |  | [optional] 
**DefaultQualityProfileId** | Pointer to **int32** |  | [optional] 
**DefaultMonitorOption** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 
**DefaultNewItemMonitorOption** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 
**DefaultTags** | Pointer to **[]int32** |  | [optional] 
**IsCalibreLibrary** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**UrlBase** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Library** | Pointer to **NullableString** |  | [optional] 
**OutputFormat** | Pointer to **NullableString** |  | [optional] 
**OutputProfile** | Pointer to **NullableString** |  | [optional] 
**UseSsl** | Pointer to **bool** |  | [optional] 
**Accessible** | Pointer to **bool** |  | [optional] 
**FreeSpace** | Pointer to **NullableInt64** |  | [optional] 
**TotalSpace** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewRootFolderResource

`func NewRootFolderResource() *RootFolderResource`

NewRootFolderResource instantiates a new RootFolderResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootFolderResourceWithDefaults

`func NewRootFolderResourceWithDefaults() *RootFolderResource`

NewRootFolderResourceWithDefaults instantiates a new RootFolderResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RootFolderResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RootFolderResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RootFolderResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RootFolderResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RootFolderResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RootFolderResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RootFolderResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RootFolderResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RootFolderResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RootFolderResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *RootFolderResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RootFolderResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RootFolderResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RootFolderResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *RootFolderResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *RootFolderResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetDefaultMetadataProfileId

`func (o *RootFolderResource) GetDefaultMetadataProfileId() int32`

GetDefaultMetadataProfileId returns the DefaultMetadataProfileId field if non-nil, zero value otherwise.

### GetDefaultMetadataProfileIdOk

`func (o *RootFolderResource) GetDefaultMetadataProfileIdOk() (*int32, bool)`

GetDefaultMetadataProfileIdOk returns a tuple with the DefaultMetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMetadataProfileId

`func (o *RootFolderResource) SetDefaultMetadataProfileId(v int32)`

SetDefaultMetadataProfileId sets DefaultMetadataProfileId field to given value.

### HasDefaultMetadataProfileId

`func (o *RootFolderResource) HasDefaultMetadataProfileId() bool`

HasDefaultMetadataProfileId returns a boolean if a field has been set.

### GetDefaultQualityProfileId

`func (o *RootFolderResource) GetDefaultQualityProfileId() int32`

GetDefaultQualityProfileId returns the DefaultQualityProfileId field if non-nil, zero value otherwise.

### GetDefaultQualityProfileIdOk

`func (o *RootFolderResource) GetDefaultQualityProfileIdOk() (*int32, bool)`

GetDefaultQualityProfileIdOk returns a tuple with the DefaultQualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQualityProfileId

`func (o *RootFolderResource) SetDefaultQualityProfileId(v int32)`

SetDefaultQualityProfileId sets DefaultQualityProfileId field to given value.

### HasDefaultQualityProfileId

`func (o *RootFolderResource) HasDefaultQualityProfileId() bool`

HasDefaultQualityProfileId returns a boolean if a field has been set.

### GetDefaultMonitorOption

`func (o *RootFolderResource) GetDefaultMonitorOption() MonitorTypes`

GetDefaultMonitorOption returns the DefaultMonitorOption field if non-nil, zero value otherwise.

### GetDefaultMonitorOptionOk

`func (o *RootFolderResource) GetDefaultMonitorOptionOk() (*MonitorTypes, bool)`

GetDefaultMonitorOptionOk returns a tuple with the DefaultMonitorOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMonitorOption

`func (o *RootFolderResource) SetDefaultMonitorOption(v MonitorTypes)`

SetDefaultMonitorOption sets DefaultMonitorOption field to given value.

### HasDefaultMonitorOption

`func (o *RootFolderResource) HasDefaultMonitorOption() bool`

HasDefaultMonitorOption returns a boolean if a field has been set.

### GetDefaultNewItemMonitorOption

`func (o *RootFolderResource) GetDefaultNewItemMonitorOption() NewItemMonitorTypes`

GetDefaultNewItemMonitorOption returns the DefaultNewItemMonitorOption field if non-nil, zero value otherwise.

### GetDefaultNewItemMonitorOptionOk

`func (o *RootFolderResource) GetDefaultNewItemMonitorOptionOk() (*NewItemMonitorTypes, bool)`

GetDefaultNewItemMonitorOptionOk returns a tuple with the DefaultNewItemMonitorOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNewItemMonitorOption

`func (o *RootFolderResource) SetDefaultNewItemMonitorOption(v NewItemMonitorTypes)`

SetDefaultNewItemMonitorOption sets DefaultNewItemMonitorOption field to given value.

### HasDefaultNewItemMonitorOption

`func (o *RootFolderResource) HasDefaultNewItemMonitorOption() bool`

HasDefaultNewItemMonitorOption returns a boolean if a field has been set.

### GetDefaultTags

`func (o *RootFolderResource) GetDefaultTags() []int32`

GetDefaultTags returns the DefaultTags field if non-nil, zero value otherwise.

### GetDefaultTagsOk

`func (o *RootFolderResource) GetDefaultTagsOk() (*[]int32, bool)`

GetDefaultTagsOk returns a tuple with the DefaultTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTags

`func (o *RootFolderResource) SetDefaultTags(v []int32)`

SetDefaultTags sets DefaultTags field to given value.

### HasDefaultTags

`func (o *RootFolderResource) HasDefaultTags() bool`

HasDefaultTags returns a boolean if a field has been set.

### SetDefaultTagsNil

`func (o *RootFolderResource) SetDefaultTagsNil(b bool)`

 SetDefaultTagsNil sets the value for DefaultTags to be an explicit nil

### UnsetDefaultTags
`func (o *RootFolderResource) UnsetDefaultTags()`

UnsetDefaultTags ensures that no value is present for DefaultTags, not even an explicit nil
### GetIsCalibreLibrary

`func (o *RootFolderResource) GetIsCalibreLibrary() bool`

GetIsCalibreLibrary returns the IsCalibreLibrary field if non-nil, zero value otherwise.

### GetIsCalibreLibraryOk

`func (o *RootFolderResource) GetIsCalibreLibraryOk() (*bool, bool)`

GetIsCalibreLibraryOk returns a tuple with the IsCalibreLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCalibreLibrary

`func (o *RootFolderResource) SetIsCalibreLibrary(v bool)`

SetIsCalibreLibrary sets IsCalibreLibrary field to given value.

### HasIsCalibreLibrary

`func (o *RootFolderResource) HasIsCalibreLibrary() bool`

HasIsCalibreLibrary returns a boolean if a field has been set.

### GetHost

`func (o *RootFolderResource) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RootFolderResource) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RootFolderResource) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RootFolderResource) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *RootFolderResource) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RootFolderResource) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *RootFolderResource) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RootFolderResource) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RootFolderResource) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RootFolderResource) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUrlBase

`func (o *RootFolderResource) GetUrlBase() string`

GetUrlBase returns the UrlBase field if non-nil, zero value otherwise.

### GetUrlBaseOk

`func (o *RootFolderResource) GetUrlBaseOk() (*string, bool)`

GetUrlBaseOk returns a tuple with the UrlBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBase

`func (o *RootFolderResource) SetUrlBase(v string)`

SetUrlBase sets UrlBase field to given value.

### HasUrlBase

`func (o *RootFolderResource) HasUrlBase() bool`

HasUrlBase returns a boolean if a field has been set.

### SetUrlBaseNil

`func (o *RootFolderResource) SetUrlBaseNil(b bool)`

 SetUrlBaseNil sets the value for UrlBase to be an explicit nil

### UnsetUrlBase
`func (o *RootFolderResource) UnsetUrlBase()`

UnsetUrlBase ensures that no value is present for UrlBase, not even an explicit nil
### GetUsername

`func (o *RootFolderResource) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RootFolderResource) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RootFolderResource) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RootFolderResource) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RootFolderResource) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RootFolderResource) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *RootFolderResource) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RootFolderResource) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RootFolderResource) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RootFolderResource) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RootFolderResource) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RootFolderResource) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetLibrary

`func (o *RootFolderResource) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *RootFolderResource) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *RootFolderResource) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *RootFolderResource) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### SetLibraryNil

`func (o *RootFolderResource) SetLibraryNil(b bool)`

 SetLibraryNil sets the value for Library to be an explicit nil

### UnsetLibrary
`func (o *RootFolderResource) UnsetLibrary()`

UnsetLibrary ensures that no value is present for Library, not even an explicit nil
### GetOutputFormat

`func (o *RootFolderResource) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *RootFolderResource) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *RootFolderResource) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *RootFolderResource) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### SetOutputFormatNil

`func (o *RootFolderResource) SetOutputFormatNil(b bool)`

 SetOutputFormatNil sets the value for OutputFormat to be an explicit nil

### UnsetOutputFormat
`func (o *RootFolderResource) UnsetOutputFormat()`

UnsetOutputFormat ensures that no value is present for OutputFormat, not even an explicit nil
### GetOutputProfile

`func (o *RootFolderResource) GetOutputProfile() string`

GetOutputProfile returns the OutputProfile field if non-nil, zero value otherwise.

### GetOutputProfileOk

`func (o *RootFolderResource) GetOutputProfileOk() (*string, bool)`

GetOutputProfileOk returns a tuple with the OutputProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputProfile

`func (o *RootFolderResource) SetOutputProfile(v string)`

SetOutputProfile sets OutputProfile field to given value.

### HasOutputProfile

`func (o *RootFolderResource) HasOutputProfile() bool`

HasOutputProfile returns a boolean if a field has been set.

### SetOutputProfileNil

`func (o *RootFolderResource) SetOutputProfileNil(b bool)`

 SetOutputProfileNil sets the value for OutputProfile to be an explicit nil

### UnsetOutputProfile
`func (o *RootFolderResource) UnsetOutputProfile()`

UnsetOutputProfile ensures that no value is present for OutputProfile, not even an explicit nil
### GetUseSsl

`func (o *RootFolderResource) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *RootFolderResource) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *RootFolderResource) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *RootFolderResource) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetAccessible

`func (o *RootFolderResource) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *RootFolderResource) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *RootFolderResource) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *RootFolderResource) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### GetFreeSpace

`func (o *RootFolderResource) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *RootFolderResource) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *RootFolderResource) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *RootFolderResource) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *RootFolderResource) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *RootFolderResource) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetTotalSpace

`func (o *RootFolderResource) GetTotalSpace() int64`

GetTotalSpace returns the TotalSpace field if non-nil, zero value otherwise.

### GetTotalSpaceOk

`func (o *RootFolderResource) GetTotalSpaceOk() (*int64, bool)`

GetTotalSpaceOk returns a tuple with the TotalSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpace

`func (o *RootFolderResource) SetTotalSpace(v int64)`

SetTotalSpace sets TotalSpace field to given value.

### HasTotalSpace

`func (o *RootFolderResource) HasTotalSpace() bool`

HasTotalSpace returns a boolean if a field has been set.

### SetTotalSpaceNil

`func (o *RootFolderResource) SetTotalSpaceNil(b bool)`

 SetTotalSpaceNil sets the value for TotalSpace to be an explicit nil

### UnsetTotalSpace
`func (o *RootFolderResource) UnsetTotalSpace()`

UnsetTotalSpace ensures that no value is present for TotalSpace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


