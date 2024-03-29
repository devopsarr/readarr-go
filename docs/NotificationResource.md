# NotificationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**ImplementationName** | Pointer to **NullableString** |  | [optional] 
**Implementation** | Pointer to **NullableString** |  | [optional] 
**ConfigContract** | Pointer to **NullableString** |  | [optional] 
**InfoLink** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to [**ProviderMessage**](ProviderMessage.md) |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Presets** | Pointer to [**[]NotificationResource**](NotificationResource.md) |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**OnGrab** | Pointer to **bool** |  | [optional] 
**OnReleaseImport** | Pointer to **bool** |  | [optional] 
**OnUpgrade** | Pointer to **bool** |  | [optional] 
**OnRename** | Pointer to **bool** |  | [optional] 
**OnAuthorAdded** | Pointer to **bool** |  | [optional] 
**OnAuthorDelete** | Pointer to **bool** |  | [optional] 
**OnBookDelete** | Pointer to **bool** |  | [optional] 
**OnBookFileDelete** | Pointer to **bool** |  | [optional] 
**OnBookFileDeleteForUpgrade** | Pointer to **bool** |  | [optional] 
**OnHealthIssue** | Pointer to **bool** |  | [optional] 
**OnDownloadFailure** | Pointer to **bool** |  | [optional] 
**OnImportFailure** | Pointer to **bool** |  | [optional] 
**OnBookRetag** | Pointer to **bool** |  | [optional] 
**OnApplicationUpdate** | Pointer to **bool** |  | [optional] 
**SupportsOnGrab** | Pointer to **bool** |  | [optional] 
**SupportsOnReleaseImport** | Pointer to **bool** |  | [optional] 
**SupportsOnUpgrade** | Pointer to **bool** |  | [optional] 
**SupportsOnRename** | Pointer to **bool** |  | [optional] 
**SupportsOnAuthorAdded** | Pointer to **bool** |  | [optional] 
**SupportsOnAuthorDelete** | Pointer to **bool** |  | [optional] 
**SupportsOnBookDelete** | Pointer to **bool** |  | [optional] 
**SupportsOnBookFileDelete** | Pointer to **bool** |  | [optional] 
**SupportsOnBookFileDeleteForUpgrade** | Pointer to **bool** |  | [optional] 
**SupportsOnHealthIssue** | Pointer to **bool** |  | [optional] 
**IncludeHealthWarnings** | Pointer to **bool** |  | [optional] 
**SupportsOnDownloadFailure** | Pointer to **bool** |  | [optional] 
**SupportsOnImportFailure** | Pointer to **bool** |  | [optional] 
**SupportsOnBookRetag** | Pointer to **bool** |  | [optional] 
**SupportsOnApplicationUpdate** | Pointer to **bool** |  | [optional] 
**TestCommand** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNotificationResource

`func NewNotificationResource() *NotificationResource`

NewNotificationResource instantiates a new NotificationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationResourceWithDefaults

`func NewNotificationResourceWithDefaults() *NotificationResource`

NewNotificationResourceWithDefaults instantiates a new NotificationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NotificationResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NotificationResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *NotificationResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *NotificationResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *NotificationResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *NotificationResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *NotificationResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *NotificationResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *NotificationResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *NotificationResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *NotificationResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *NotificationResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *NotificationResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *NotificationResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *NotificationResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *NotificationResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *NotificationResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *NotificationResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *NotificationResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *NotificationResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *NotificationResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *NotificationResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *NotificationResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *NotificationResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *NotificationResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *NotificationResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *NotificationResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *NotificationResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *NotificationResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *NotificationResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *NotificationResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *NotificationResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *NotificationResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotificationResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *NotificationResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NotificationResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NotificationResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NotificationResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *NotificationResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NotificationResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *NotificationResource) GetPresets() []NotificationResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *NotificationResource) GetPresetsOk() (*[]NotificationResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *NotificationResource) SetPresets(v []NotificationResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *NotificationResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *NotificationResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *NotificationResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetLink

`func (o *NotificationResource) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *NotificationResource) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *NotificationResource) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *NotificationResource) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *NotificationResource) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *NotificationResource) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetOnGrab

`func (o *NotificationResource) GetOnGrab() bool`

GetOnGrab returns the OnGrab field if non-nil, zero value otherwise.

### GetOnGrabOk

`func (o *NotificationResource) GetOnGrabOk() (*bool, bool)`

GetOnGrabOk returns a tuple with the OnGrab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnGrab

`func (o *NotificationResource) SetOnGrab(v bool)`

SetOnGrab sets OnGrab field to given value.

### HasOnGrab

`func (o *NotificationResource) HasOnGrab() bool`

HasOnGrab returns a boolean if a field has been set.

### GetOnReleaseImport

`func (o *NotificationResource) GetOnReleaseImport() bool`

GetOnReleaseImport returns the OnReleaseImport field if non-nil, zero value otherwise.

### GetOnReleaseImportOk

`func (o *NotificationResource) GetOnReleaseImportOk() (*bool, bool)`

GetOnReleaseImportOk returns a tuple with the OnReleaseImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnReleaseImport

`func (o *NotificationResource) SetOnReleaseImport(v bool)`

SetOnReleaseImport sets OnReleaseImport field to given value.

### HasOnReleaseImport

`func (o *NotificationResource) HasOnReleaseImport() bool`

HasOnReleaseImport returns a boolean if a field has been set.

### GetOnUpgrade

`func (o *NotificationResource) GetOnUpgrade() bool`

GetOnUpgrade returns the OnUpgrade field if non-nil, zero value otherwise.

### GetOnUpgradeOk

`func (o *NotificationResource) GetOnUpgradeOk() (*bool, bool)`

GetOnUpgradeOk returns a tuple with the OnUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnUpgrade

`func (o *NotificationResource) SetOnUpgrade(v bool)`

SetOnUpgrade sets OnUpgrade field to given value.

### HasOnUpgrade

`func (o *NotificationResource) HasOnUpgrade() bool`

HasOnUpgrade returns a boolean if a field has been set.

### GetOnRename

`func (o *NotificationResource) GetOnRename() bool`

GetOnRename returns the OnRename field if non-nil, zero value otherwise.

### GetOnRenameOk

`func (o *NotificationResource) GetOnRenameOk() (*bool, bool)`

GetOnRenameOk returns a tuple with the OnRename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRename

`func (o *NotificationResource) SetOnRename(v bool)`

SetOnRename sets OnRename field to given value.

### HasOnRename

`func (o *NotificationResource) HasOnRename() bool`

HasOnRename returns a boolean if a field has been set.

### GetOnAuthorAdded

`func (o *NotificationResource) GetOnAuthorAdded() bool`

GetOnAuthorAdded returns the OnAuthorAdded field if non-nil, zero value otherwise.

### GetOnAuthorAddedOk

`func (o *NotificationResource) GetOnAuthorAddedOk() (*bool, bool)`

GetOnAuthorAddedOk returns a tuple with the OnAuthorAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnAuthorAdded

`func (o *NotificationResource) SetOnAuthorAdded(v bool)`

SetOnAuthorAdded sets OnAuthorAdded field to given value.

### HasOnAuthorAdded

`func (o *NotificationResource) HasOnAuthorAdded() bool`

HasOnAuthorAdded returns a boolean if a field has been set.

### GetOnAuthorDelete

`func (o *NotificationResource) GetOnAuthorDelete() bool`

GetOnAuthorDelete returns the OnAuthorDelete field if non-nil, zero value otherwise.

### GetOnAuthorDeleteOk

`func (o *NotificationResource) GetOnAuthorDeleteOk() (*bool, bool)`

GetOnAuthorDeleteOk returns a tuple with the OnAuthorDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnAuthorDelete

`func (o *NotificationResource) SetOnAuthorDelete(v bool)`

SetOnAuthorDelete sets OnAuthorDelete field to given value.

### HasOnAuthorDelete

`func (o *NotificationResource) HasOnAuthorDelete() bool`

HasOnAuthorDelete returns a boolean if a field has been set.

### GetOnBookDelete

`func (o *NotificationResource) GetOnBookDelete() bool`

GetOnBookDelete returns the OnBookDelete field if non-nil, zero value otherwise.

### GetOnBookDeleteOk

`func (o *NotificationResource) GetOnBookDeleteOk() (*bool, bool)`

GetOnBookDeleteOk returns a tuple with the OnBookDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBookDelete

`func (o *NotificationResource) SetOnBookDelete(v bool)`

SetOnBookDelete sets OnBookDelete field to given value.

### HasOnBookDelete

`func (o *NotificationResource) HasOnBookDelete() bool`

HasOnBookDelete returns a boolean if a field has been set.

### GetOnBookFileDelete

`func (o *NotificationResource) GetOnBookFileDelete() bool`

GetOnBookFileDelete returns the OnBookFileDelete field if non-nil, zero value otherwise.

### GetOnBookFileDeleteOk

`func (o *NotificationResource) GetOnBookFileDeleteOk() (*bool, bool)`

GetOnBookFileDeleteOk returns a tuple with the OnBookFileDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBookFileDelete

`func (o *NotificationResource) SetOnBookFileDelete(v bool)`

SetOnBookFileDelete sets OnBookFileDelete field to given value.

### HasOnBookFileDelete

`func (o *NotificationResource) HasOnBookFileDelete() bool`

HasOnBookFileDelete returns a boolean if a field has been set.

### GetOnBookFileDeleteForUpgrade

`func (o *NotificationResource) GetOnBookFileDeleteForUpgrade() bool`

GetOnBookFileDeleteForUpgrade returns the OnBookFileDeleteForUpgrade field if non-nil, zero value otherwise.

### GetOnBookFileDeleteForUpgradeOk

`func (o *NotificationResource) GetOnBookFileDeleteForUpgradeOk() (*bool, bool)`

GetOnBookFileDeleteForUpgradeOk returns a tuple with the OnBookFileDeleteForUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBookFileDeleteForUpgrade

`func (o *NotificationResource) SetOnBookFileDeleteForUpgrade(v bool)`

SetOnBookFileDeleteForUpgrade sets OnBookFileDeleteForUpgrade field to given value.

### HasOnBookFileDeleteForUpgrade

`func (o *NotificationResource) HasOnBookFileDeleteForUpgrade() bool`

HasOnBookFileDeleteForUpgrade returns a boolean if a field has been set.

### GetOnHealthIssue

`func (o *NotificationResource) GetOnHealthIssue() bool`

GetOnHealthIssue returns the OnHealthIssue field if non-nil, zero value otherwise.

### GetOnHealthIssueOk

`func (o *NotificationResource) GetOnHealthIssueOk() (*bool, bool)`

GetOnHealthIssueOk returns a tuple with the OnHealthIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHealthIssue

`func (o *NotificationResource) SetOnHealthIssue(v bool)`

SetOnHealthIssue sets OnHealthIssue field to given value.

### HasOnHealthIssue

`func (o *NotificationResource) HasOnHealthIssue() bool`

HasOnHealthIssue returns a boolean if a field has been set.

### GetOnDownloadFailure

`func (o *NotificationResource) GetOnDownloadFailure() bool`

GetOnDownloadFailure returns the OnDownloadFailure field if non-nil, zero value otherwise.

### GetOnDownloadFailureOk

`func (o *NotificationResource) GetOnDownloadFailureOk() (*bool, bool)`

GetOnDownloadFailureOk returns a tuple with the OnDownloadFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDownloadFailure

`func (o *NotificationResource) SetOnDownloadFailure(v bool)`

SetOnDownloadFailure sets OnDownloadFailure field to given value.

### HasOnDownloadFailure

`func (o *NotificationResource) HasOnDownloadFailure() bool`

HasOnDownloadFailure returns a boolean if a field has been set.

### GetOnImportFailure

`func (o *NotificationResource) GetOnImportFailure() bool`

GetOnImportFailure returns the OnImportFailure field if non-nil, zero value otherwise.

### GetOnImportFailureOk

`func (o *NotificationResource) GetOnImportFailureOk() (*bool, bool)`

GetOnImportFailureOk returns a tuple with the OnImportFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnImportFailure

`func (o *NotificationResource) SetOnImportFailure(v bool)`

SetOnImportFailure sets OnImportFailure field to given value.

### HasOnImportFailure

`func (o *NotificationResource) HasOnImportFailure() bool`

HasOnImportFailure returns a boolean if a field has been set.

### GetOnBookRetag

`func (o *NotificationResource) GetOnBookRetag() bool`

GetOnBookRetag returns the OnBookRetag field if non-nil, zero value otherwise.

### GetOnBookRetagOk

`func (o *NotificationResource) GetOnBookRetagOk() (*bool, bool)`

GetOnBookRetagOk returns a tuple with the OnBookRetag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBookRetag

`func (o *NotificationResource) SetOnBookRetag(v bool)`

SetOnBookRetag sets OnBookRetag field to given value.

### HasOnBookRetag

`func (o *NotificationResource) HasOnBookRetag() bool`

HasOnBookRetag returns a boolean if a field has been set.

### GetOnApplicationUpdate

`func (o *NotificationResource) GetOnApplicationUpdate() bool`

GetOnApplicationUpdate returns the OnApplicationUpdate field if non-nil, zero value otherwise.

### GetOnApplicationUpdateOk

`func (o *NotificationResource) GetOnApplicationUpdateOk() (*bool, bool)`

GetOnApplicationUpdateOk returns a tuple with the OnApplicationUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApplicationUpdate

`func (o *NotificationResource) SetOnApplicationUpdate(v bool)`

SetOnApplicationUpdate sets OnApplicationUpdate field to given value.

### HasOnApplicationUpdate

`func (o *NotificationResource) HasOnApplicationUpdate() bool`

HasOnApplicationUpdate returns a boolean if a field has been set.

### GetSupportsOnGrab

`func (o *NotificationResource) GetSupportsOnGrab() bool`

GetSupportsOnGrab returns the SupportsOnGrab field if non-nil, zero value otherwise.

### GetSupportsOnGrabOk

`func (o *NotificationResource) GetSupportsOnGrabOk() (*bool, bool)`

GetSupportsOnGrabOk returns a tuple with the SupportsOnGrab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnGrab

`func (o *NotificationResource) SetSupportsOnGrab(v bool)`

SetSupportsOnGrab sets SupportsOnGrab field to given value.

### HasSupportsOnGrab

`func (o *NotificationResource) HasSupportsOnGrab() bool`

HasSupportsOnGrab returns a boolean if a field has been set.

### GetSupportsOnReleaseImport

`func (o *NotificationResource) GetSupportsOnReleaseImport() bool`

GetSupportsOnReleaseImport returns the SupportsOnReleaseImport field if non-nil, zero value otherwise.

### GetSupportsOnReleaseImportOk

`func (o *NotificationResource) GetSupportsOnReleaseImportOk() (*bool, bool)`

GetSupportsOnReleaseImportOk returns a tuple with the SupportsOnReleaseImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnReleaseImport

`func (o *NotificationResource) SetSupportsOnReleaseImport(v bool)`

SetSupportsOnReleaseImport sets SupportsOnReleaseImport field to given value.

### HasSupportsOnReleaseImport

`func (o *NotificationResource) HasSupportsOnReleaseImport() bool`

HasSupportsOnReleaseImport returns a boolean if a field has been set.

### GetSupportsOnUpgrade

`func (o *NotificationResource) GetSupportsOnUpgrade() bool`

GetSupportsOnUpgrade returns the SupportsOnUpgrade field if non-nil, zero value otherwise.

### GetSupportsOnUpgradeOk

`func (o *NotificationResource) GetSupportsOnUpgradeOk() (*bool, bool)`

GetSupportsOnUpgradeOk returns a tuple with the SupportsOnUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnUpgrade

`func (o *NotificationResource) SetSupportsOnUpgrade(v bool)`

SetSupportsOnUpgrade sets SupportsOnUpgrade field to given value.

### HasSupportsOnUpgrade

`func (o *NotificationResource) HasSupportsOnUpgrade() bool`

HasSupportsOnUpgrade returns a boolean if a field has been set.

### GetSupportsOnRename

`func (o *NotificationResource) GetSupportsOnRename() bool`

GetSupportsOnRename returns the SupportsOnRename field if non-nil, zero value otherwise.

### GetSupportsOnRenameOk

`func (o *NotificationResource) GetSupportsOnRenameOk() (*bool, bool)`

GetSupportsOnRenameOk returns a tuple with the SupportsOnRename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnRename

`func (o *NotificationResource) SetSupportsOnRename(v bool)`

SetSupportsOnRename sets SupportsOnRename field to given value.

### HasSupportsOnRename

`func (o *NotificationResource) HasSupportsOnRename() bool`

HasSupportsOnRename returns a boolean if a field has been set.

### GetSupportsOnAuthorAdded

`func (o *NotificationResource) GetSupportsOnAuthorAdded() bool`

GetSupportsOnAuthorAdded returns the SupportsOnAuthorAdded field if non-nil, zero value otherwise.

### GetSupportsOnAuthorAddedOk

`func (o *NotificationResource) GetSupportsOnAuthorAddedOk() (*bool, bool)`

GetSupportsOnAuthorAddedOk returns a tuple with the SupportsOnAuthorAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnAuthorAdded

`func (o *NotificationResource) SetSupportsOnAuthorAdded(v bool)`

SetSupportsOnAuthorAdded sets SupportsOnAuthorAdded field to given value.

### HasSupportsOnAuthorAdded

`func (o *NotificationResource) HasSupportsOnAuthorAdded() bool`

HasSupportsOnAuthorAdded returns a boolean if a field has been set.

### GetSupportsOnAuthorDelete

`func (o *NotificationResource) GetSupportsOnAuthorDelete() bool`

GetSupportsOnAuthorDelete returns the SupportsOnAuthorDelete field if non-nil, zero value otherwise.

### GetSupportsOnAuthorDeleteOk

`func (o *NotificationResource) GetSupportsOnAuthorDeleteOk() (*bool, bool)`

GetSupportsOnAuthorDeleteOk returns a tuple with the SupportsOnAuthorDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnAuthorDelete

`func (o *NotificationResource) SetSupportsOnAuthorDelete(v bool)`

SetSupportsOnAuthorDelete sets SupportsOnAuthorDelete field to given value.

### HasSupportsOnAuthorDelete

`func (o *NotificationResource) HasSupportsOnAuthorDelete() bool`

HasSupportsOnAuthorDelete returns a boolean if a field has been set.

### GetSupportsOnBookDelete

`func (o *NotificationResource) GetSupportsOnBookDelete() bool`

GetSupportsOnBookDelete returns the SupportsOnBookDelete field if non-nil, zero value otherwise.

### GetSupportsOnBookDeleteOk

`func (o *NotificationResource) GetSupportsOnBookDeleteOk() (*bool, bool)`

GetSupportsOnBookDeleteOk returns a tuple with the SupportsOnBookDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnBookDelete

`func (o *NotificationResource) SetSupportsOnBookDelete(v bool)`

SetSupportsOnBookDelete sets SupportsOnBookDelete field to given value.

### HasSupportsOnBookDelete

`func (o *NotificationResource) HasSupportsOnBookDelete() bool`

HasSupportsOnBookDelete returns a boolean if a field has been set.

### GetSupportsOnBookFileDelete

`func (o *NotificationResource) GetSupportsOnBookFileDelete() bool`

GetSupportsOnBookFileDelete returns the SupportsOnBookFileDelete field if non-nil, zero value otherwise.

### GetSupportsOnBookFileDeleteOk

`func (o *NotificationResource) GetSupportsOnBookFileDeleteOk() (*bool, bool)`

GetSupportsOnBookFileDeleteOk returns a tuple with the SupportsOnBookFileDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnBookFileDelete

`func (o *NotificationResource) SetSupportsOnBookFileDelete(v bool)`

SetSupportsOnBookFileDelete sets SupportsOnBookFileDelete field to given value.

### HasSupportsOnBookFileDelete

`func (o *NotificationResource) HasSupportsOnBookFileDelete() bool`

HasSupportsOnBookFileDelete returns a boolean if a field has been set.

### GetSupportsOnBookFileDeleteForUpgrade

`func (o *NotificationResource) GetSupportsOnBookFileDeleteForUpgrade() bool`

GetSupportsOnBookFileDeleteForUpgrade returns the SupportsOnBookFileDeleteForUpgrade field if non-nil, zero value otherwise.

### GetSupportsOnBookFileDeleteForUpgradeOk

`func (o *NotificationResource) GetSupportsOnBookFileDeleteForUpgradeOk() (*bool, bool)`

GetSupportsOnBookFileDeleteForUpgradeOk returns a tuple with the SupportsOnBookFileDeleteForUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnBookFileDeleteForUpgrade

`func (o *NotificationResource) SetSupportsOnBookFileDeleteForUpgrade(v bool)`

SetSupportsOnBookFileDeleteForUpgrade sets SupportsOnBookFileDeleteForUpgrade field to given value.

### HasSupportsOnBookFileDeleteForUpgrade

`func (o *NotificationResource) HasSupportsOnBookFileDeleteForUpgrade() bool`

HasSupportsOnBookFileDeleteForUpgrade returns a boolean if a field has been set.

### GetSupportsOnHealthIssue

`func (o *NotificationResource) GetSupportsOnHealthIssue() bool`

GetSupportsOnHealthIssue returns the SupportsOnHealthIssue field if non-nil, zero value otherwise.

### GetSupportsOnHealthIssueOk

`func (o *NotificationResource) GetSupportsOnHealthIssueOk() (*bool, bool)`

GetSupportsOnHealthIssueOk returns a tuple with the SupportsOnHealthIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnHealthIssue

`func (o *NotificationResource) SetSupportsOnHealthIssue(v bool)`

SetSupportsOnHealthIssue sets SupportsOnHealthIssue field to given value.

### HasSupportsOnHealthIssue

`func (o *NotificationResource) HasSupportsOnHealthIssue() bool`

HasSupportsOnHealthIssue returns a boolean if a field has been set.

### GetIncludeHealthWarnings

`func (o *NotificationResource) GetIncludeHealthWarnings() bool`

GetIncludeHealthWarnings returns the IncludeHealthWarnings field if non-nil, zero value otherwise.

### GetIncludeHealthWarningsOk

`func (o *NotificationResource) GetIncludeHealthWarningsOk() (*bool, bool)`

GetIncludeHealthWarningsOk returns a tuple with the IncludeHealthWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHealthWarnings

`func (o *NotificationResource) SetIncludeHealthWarnings(v bool)`

SetIncludeHealthWarnings sets IncludeHealthWarnings field to given value.

### HasIncludeHealthWarnings

`func (o *NotificationResource) HasIncludeHealthWarnings() bool`

HasIncludeHealthWarnings returns a boolean if a field has been set.

### GetSupportsOnDownloadFailure

`func (o *NotificationResource) GetSupportsOnDownloadFailure() bool`

GetSupportsOnDownloadFailure returns the SupportsOnDownloadFailure field if non-nil, zero value otherwise.

### GetSupportsOnDownloadFailureOk

`func (o *NotificationResource) GetSupportsOnDownloadFailureOk() (*bool, bool)`

GetSupportsOnDownloadFailureOk returns a tuple with the SupportsOnDownloadFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnDownloadFailure

`func (o *NotificationResource) SetSupportsOnDownloadFailure(v bool)`

SetSupportsOnDownloadFailure sets SupportsOnDownloadFailure field to given value.

### HasSupportsOnDownloadFailure

`func (o *NotificationResource) HasSupportsOnDownloadFailure() bool`

HasSupportsOnDownloadFailure returns a boolean if a field has been set.

### GetSupportsOnImportFailure

`func (o *NotificationResource) GetSupportsOnImportFailure() bool`

GetSupportsOnImportFailure returns the SupportsOnImportFailure field if non-nil, zero value otherwise.

### GetSupportsOnImportFailureOk

`func (o *NotificationResource) GetSupportsOnImportFailureOk() (*bool, bool)`

GetSupportsOnImportFailureOk returns a tuple with the SupportsOnImportFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnImportFailure

`func (o *NotificationResource) SetSupportsOnImportFailure(v bool)`

SetSupportsOnImportFailure sets SupportsOnImportFailure field to given value.

### HasSupportsOnImportFailure

`func (o *NotificationResource) HasSupportsOnImportFailure() bool`

HasSupportsOnImportFailure returns a boolean if a field has been set.

### GetSupportsOnBookRetag

`func (o *NotificationResource) GetSupportsOnBookRetag() bool`

GetSupportsOnBookRetag returns the SupportsOnBookRetag field if non-nil, zero value otherwise.

### GetSupportsOnBookRetagOk

`func (o *NotificationResource) GetSupportsOnBookRetagOk() (*bool, bool)`

GetSupportsOnBookRetagOk returns a tuple with the SupportsOnBookRetag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnBookRetag

`func (o *NotificationResource) SetSupportsOnBookRetag(v bool)`

SetSupportsOnBookRetag sets SupportsOnBookRetag field to given value.

### HasSupportsOnBookRetag

`func (o *NotificationResource) HasSupportsOnBookRetag() bool`

HasSupportsOnBookRetag returns a boolean if a field has been set.

### GetSupportsOnApplicationUpdate

`func (o *NotificationResource) GetSupportsOnApplicationUpdate() bool`

GetSupportsOnApplicationUpdate returns the SupportsOnApplicationUpdate field if non-nil, zero value otherwise.

### GetSupportsOnApplicationUpdateOk

`func (o *NotificationResource) GetSupportsOnApplicationUpdateOk() (*bool, bool)`

GetSupportsOnApplicationUpdateOk returns a tuple with the SupportsOnApplicationUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnApplicationUpdate

`func (o *NotificationResource) SetSupportsOnApplicationUpdate(v bool)`

SetSupportsOnApplicationUpdate sets SupportsOnApplicationUpdate field to given value.

### HasSupportsOnApplicationUpdate

`func (o *NotificationResource) HasSupportsOnApplicationUpdate() bool`

HasSupportsOnApplicationUpdate returns a boolean if a field has been set.

### GetTestCommand

`func (o *NotificationResource) GetTestCommand() string`

GetTestCommand returns the TestCommand field if non-nil, zero value otherwise.

### GetTestCommandOk

`func (o *NotificationResource) GetTestCommandOk() (*string, bool)`

GetTestCommandOk returns a tuple with the TestCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCommand

`func (o *NotificationResource) SetTestCommand(v string)`

SetTestCommand sets TestCommand field to given value.

### HasTestCommand

`func (o *NotificationResource) HasTestCommand() bool`

HasTestCommand returns a boolean if a field has been set.

### SetTestCommandNil

`func (o *NotificationResource) SetTestCommandNil(b bool)`

 SetTestCommandNil sets the value for TestCommand to be an explicit nil

### UnsetTestCommand
`func (o *NotificationResource) UnsetTestCommand()`

UnsetTestCommand ensures that no value is present for TestCommand, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


