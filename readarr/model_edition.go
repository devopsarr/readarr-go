/*
Readarr

Readarr API docs

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
	"time"
)

// Edition struct for Edition
type Edition struct {
	Id *int32 `json:"id,omitempty"`
	BookId *int32 `json:"bookId,omitempty"`
	ForeignEditionId NullableString `json:"foreignEditionId,omitempty"`
	TitleSlug NullableString `json:"titleSlug,omitempty"`
	Isbn13 NullableString `json:"isbn13,omitempty"`
	Asin NullableString `json:"asin,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Language NullableString `json:"language,omitempty"`
	Overview NullableString `json:"overview,omitempty"`
	Format NullableString `json:"format,omitempty"`
	IsEbook *bool `json:"isEbook,omitempty"`
	Disambiguation NullableString `json:"disambiguation,omitempty"`
	Publisher NullableString `json:"publisher,omitempty"`
	PageCount *int32 `json:"pageCount,omitempty"`
	ReleaseDate NullableTime `json:"releaseDate,omitempty"`
	Images []*MediaCover `json:"images,omitempty"`
	Links []*Links `json:"links,omitempty"`
	Ratings *Ratings `json:"ratings,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	ManualAdd *bool `json:"manualAdd,omitempty"`
	Book *BookLazyLoaded `json:"book,omitempty"`
	BookFiles *BookFileListLazyLoaded `json:"bookFiles,omitempty"`
}

// NewEdition instantiates a new Edition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdition() *Edition {
	this := Edition{}
	return &this
}

// NewEditionWithDefaults instantiates a new Edition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditionWithDefaults() *Edition {
	this := Edition{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Edition) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Edition) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Edition) SetId(v int32) {
	o.Id = &v
}

// GetBookId returns the BookId field value if set, zero value otherwise.
func (o *Edition) GetBookId() int32 {
	if o == nil || isNil(o.BookId) {
		var ret int32
		return ret
	}
	return *o.BookId
}

// GetBookIdOk returns a tuple with the BookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetBookIdOk() (*int32, bool) {
	if o == nil || isNil(o.BookId) {
    return nil, false
	}
	return o.BookId, true
}

// HasBookId returns a boolean if a field has been set.
func (o *Edition) HasBookId() bool {
	if o != nil && !isNil(o.BookId) {
		return true
	}

	return false
}

// SetBookId gets a reference to the given int32 and assigns it to the BookId field.
func (o *Edition) SetBookId(v int32) {
	o.BookId = &v
}

// GetForeignEditionId returns the ForeignEditionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetForeignEditionId() string {
	if o == nil || isNil(o.ForeignEditionId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignEditionId.Get()
}

// GetForeignEditionIdOk returns a tuple with the ForeignEditionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetForeignEditionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForeignEditionId.Get(), o.ForeignEditionId.IsSet()
}

// HasForeignEditionId returns a boolean if a field has been set.
func (o *Edition) HasForeignEditionId() bool {
	if o != nil && o.ForeignEditionId.IsSet() {
		return true
	}

	return false
}

// SetForeignEditionId gets a reference to the given NullableString and assigns it to the ForeignEditionId field.
func (o *Edition) SetForeignEditionId(v string) {
	o.ForeignEditionId.Set(&v)
}
// SetForeignEditionIdNil sets the value for ForeignEditionId to be an explicit nil
func (o *Edition) SetForeignEditionIdNil() {
	o.ForeignEditionId.Set(nil)
}

// UnsetForeignEditionId ensures that no value is present for ForeignEditionId, not even an explicit nil
func (o *Edition) UnsetForeignEditionId() {
	o.ForeignEditionId.Unset()
}

// GetTitleSlug returns the TitleSlug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetTitleSlug() string {
	if o == nil || isNil(o.TitleSlug.Get()) {
		var ret string
		return ret
	}
	return *o.TitleSlug.Get()
}

// GetTitleSlugOk returns a tuple with the TitleSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetTitleSlugOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TitleSlug.Get(), o.TitleSlug.IsSet()
}

// HasTitleSlug returns a boolean if a field has been set.
func (o *Edition) HasTitleSlug() bool {
	if o != nil && o.TitleSlug.IsSet() {
		return true
	}

	return false
}

// SetTitleSlug gets a reference to the given NullableString and assigns it to the TitleSlug field.
func (o *Edition) SetTitleSlug(v string) {
	o.TitleSlug.Set(&v)
}
// SetTitleSlugNil sets the value for TitleSlug to be an explicit nil
func (o *Edition) SetTitleSlugNil() {
	o.TitleSlug.Set(nil)
}

// UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
func (o *Edition) UnsetTitleSlug() {
	o.TitleSlug.Unset()
}

// GetIsbn13 returns the Isbn13 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetIsbn13() string {
	if o == nil || isNil(o.Isbn13.Get()) {
		var ret string
		return ret
	}
	return *o.Isbn13.Get()
}

// GetIsbn13Ok returns a tuple with the Isbn13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetIsbn13Ok() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Isbn13.Get(), o.Isbn13.IsSet()
}

// HasIsbn13 returns a boolean if a field has been set.
func (o *Edition) HasIsbn13() bool {
	if o != nil && o.Isbn13.IsSet() {
		return true
	}

	return false
}

// SetIsbn13 gets a reference to the given NullableString and assigns it to the Isbn13 field.
func (o *Edition) SetIsbn13(v string) {
	o.Isbn13.Set(&v)
}
// SetIsbn13Nil sets the value for Isbn13 to be an explicit nil
func (o *Edition) SetIsbn13Nil() {
	o.Isbn13.Set(nil)
}

// UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
func (o *Edition) UnsetIsbn13() {
	o.Isbn13.Unset()
}

// GetAsin returns the Asin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetAsin() string {
	if o == nil || isNil(o.Asin.Get()) {
		var ret string
		return ret
	}
	return *o.Asin.Get()
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetAsinOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Asin.Get(), o.Asin.IsSet()
}

// HasAsin returns a boolean if a field has been set.
func (o *Edition) HasAsin() bool {
	if o != nil && o.Asin.IsSet() {
		return true
	}

	return false
}

// SetAsin gets a reference to the given NullableString and assigns it to the Asin field.
func (o *Edition) SetAsin(v string) {
	o.Asin.Set(&v)
}
// SetAsinNil sets the value for Asin to be an explicit nil
func (o *Edition) SetAsinNil() {
	o.Asin.Set(nil)
}

// UnsetAsin ensures that no value is present for Asin, not even an explicit nil
func (o *Edition) UnsetAsin() {
	o.Asin.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Edition) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Edition) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Edition) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Edition) UnsetTitle() {
	o.Title.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetLanguage() string {
	if o == nil || isNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetLanguageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *Edition) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *Edition) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *Edition) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *Edition) UnsetLanguage() {
	o.Language.Unset()
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetOverview() string {
	if o == nil || isNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetOverviewOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *Edition) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *Edition) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *Edition) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *Edition) UnsetOverview() {
	o.Overview.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetFormat() string {
	if o == nil || isNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *Edition) HasFormat() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *Edition) SetFormat(v string) {
	o.Format.Set(&v)
}
// SetFormatNil sets the value for Format to be an explicit nil
func (o *Edition) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *Edition) UnsetFormat() {
	o.Format.Unset()
}

// GetIsEbook returns the IsEbook field value if set, zero value otherwise.
func (o *Edition) GetIsEbook() bool {
	if o == nil || isNil(o.IsEbook) {
		var ret bool
		return ret
	}
	return *o.IsEbook
}

// GetIsEbookOk returns a tuple with the IsEbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetIsEbookOk() (*bool, bool) {
	if o == nil || isNil(o.IsEbook) {
    return nil, false
	}
	return o.IsEbook, true
}

// HasIsEbook returns a boolean if a field has been set.
func (o *Edition) HasIsEbook() bool {
	if o != nil && !isNil(o.IsEbook) {
		return true
	}

	return false
}

// SetIsEbook gets a reference to the given bool and assigns it to the IsEbook field.
func (o *Edition) SetIsEbook(v bool) {
	o.IsEbook = &v
}

// GetDisambiguation returns the Disambiguation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetDisambiguation() string {
	if o == nil || isNil(o.Disambiguation.Get()) {
		var ret string
		return ret
	}
	return *o.Disambiguation.Get()
}

// GetDisambiguationOk returns a tuple with the Disambiguation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetDisambiguationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Disambiguation.Get(), o.Disambiguation.IsSet()
}

// HasDisambiguation returns a boolean if a field has been set.
func (o *Edition) HasDisambiguation() bool {
	if o != nil && o.Disambiguation.IsSet() {
		return true
	}

	return false
}

// SetDisambiguation gets a reference to the given NullableString and assigns it to the Disambiguation field.
func (o *Edition) SetDisambiguation(v string) {
	o.Disambiguation.Set(&v)
}
// SetDisambiguationNil sets the value for Disambiguation to be an explicit nil
func (o *Edition) SetDisambiguationNil() {
	o.Disambiguation.Set(nil)
}

// UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
func (o *Edition) UnsetDisambiguation() {
	o.Disambiguation.Unset()
}

// GetPublisher returns the Publisher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetPublisher() string {
	if o == nil || isNil(o.Publisher.Get()) {
		var ret string
		return ret
	}
	return *o.Publisher.Get()
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetPublisherOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Publisher.Get(), o.Publisher.IsSet()
}

// HasPublisher returns a boolean if a field has been set.
func (o *Edition) HasPublisher() bool {
	if o != nil && o.Publisher.IsSet() {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given NullableString and assigns it to the Publisher field.
func (o *Edition) SetPublisher(v string) {
	o.Publisher.Set(&v)
}
// SetPublisherNil sets the value for Publisher to be an explicit nil
func (o *Edition) SetPublisherNil() {
	o.Publisher.Set(nil)
}

// UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
func (o *Edition) UnsetPublisher() {
	o.Publisher.Unset()
}

// GetPageCount returns the PageCount field value if set, zero value otherwise.
func (o *Edition) GetPageCount() int32 {
	if o == nil || isNil(o.PageCount) {
		var ret int32
		return ret
	}
	return *o.PageCount
}

// GetPageCountOk returns a tuple with the PageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetPageCountOk() (*int32, bool) {
	if o == nil || isNil(o.PageCount) {
    return nil, false
	}
	return o.PageCount, true
}

// HasPageCount returns a boolean if a field has been set.
func (o *Edition) HasPageCount() bool {
	if o != nil && !isNil(o.PageCount) {
		return true
	}

	return false
}

// SetPageCount gets a reference to the given int32 and assigns it to the PageCount field.
func (o *Edition) SetPageCount(v int32) {
	o.PageCount = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetReleaseDate() time.Time {
	if o == nil || isNil(o.ReleaseDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *Edition) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate.IsSet() {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given NullableTime and assigns it to the ReleaseDate field.
func (o *Edition) SetReleaseDate(v time.Time) {
	o.ReleaseDate.Set(&v)
}
// SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil
func (o *Edition) SetReleaseDateNil() {
	o.ReleaseDate.Set(nil)
}

// UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
func (o *Edition) UnsetReleaseDate() {
	o.ReleaseDate.Unset()
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetImages() []*MediaCover {
	if o == nil {
		var ret []*MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetImagesOk() ([]*MediaCover, bool) {
	if o == nil || isNil(o.Images) {
    return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Edition) HasImages() bool {
	if o != nil && isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *Edition) SetImages(v []*MediaCover) {
	o.Images = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Edition) GetLinks() []*Links {
	if o == nil {
		var ret []*Links
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Edition) GetLinksOk() ([]*Links, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Edition) HasLinks() bool {
	if o != nil && isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Links and assigns it to the Links field.
func (o *Edition) SetLinks(v []*Links) {
	o.Links = v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *Edition) GetRatings() Ratings {
	if o == nil || isNil(o.Ratings) {
		var ret Ratings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetRatingsOk() (*Ratings, bool) {
	if o == nil || isNil(o.Ratings) {
    return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *Edition) HasRatings() bool {
	if o != nil && !isNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given Ratings and assigns it to the Ratings field.
func (o *Edition) SetRatings(v Ratings) {
	o.Ratings = &v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *Edition) GetMonitored() bool {
	if o == nil || isNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetMonitoredOk() (*bool, bool) {
	if o == nil || isNil(o.Monitored) {
    return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *Edition) HasMonitored() bool {
	if o != nil && !isNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *Edition) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetManualAdd returns the ManualAdd field value if set, zero value otherwise.
func (o *Edition) GetManualAdd() bool {
	if o == nil || isNil(o.ManualAdd) {
		var ret bool
		return ret
	}
	return *o.ManualAdd
}

// GetManualAddOk returns a tuple with the ManualAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetManualAddOk() (*bool, bool) {
	if o == nil || isNil(o.ManualAdd) {
    return nil, false
	}
	return o.ManualAdd, true
}

// HasManualAdd returns a boolean if a field has been set.
func (o *Edition) HasManualAdd() bool {
	if o != nil && !isNil(o.ManualAdd) {
		return true
	}

	return false
}

// SetManualAdd gets a reference to the given bool and assigns it to the ManualAdd field.
func (o *Edition) SetManualAdd(v bool) {
	o.ManualAdd = &v
}

// GetBook returns the Book field value if set, zero value otherwise.
func (o *Edition) GetBook() BookLazyLoaded {
	if o == nil || isNil(o.Book) {
		var ret BookLazyLoaded
		return ret
	}
	return *o.Book
}

// GetBookOk returns a tuple with the Book field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetBookOk() (*BookLazyLoaded, bool) {
	if o == nil || isNil(o.Book) {
    return nil, false
	}
	return o.Book, true
}

// HasBook returns a boolean if a field has been set.
func (o *Edition) HasBook() bool {
	if o != nil && !isNil(o.Book) {
		return true
	}

	return false
}

// SetBook gets a reference to the given BookLazyLoaded and assigns it to the Book field.
func (o *Edition) SetBook(v BookLazyLoaded) {
	o.Book = &v
}

// GetBookFiles returns the BookFiles field value if set, zero value otherwise.
func (o *Edition) GetBookFiles() BookFileListLazyLoaded {
	if o == nil || isNil(o.BookFiles) {
		var ret BookFileListLazyLoaded
		return ret
	}
	return *o.BookFiles
}

// GetBookFilesOk returns a tuple with the BookFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edition) GetBookFilesOk() (*BookFileListLazyLoaded, bool) {
	if o == nil || isNil(o.BookFiles) {
    return nil, false
	}
	return o.BookFiles, true
}

// HasBookFiles returns a boolean if a field has been set.
func (o *Edition) HasBookFiles() bool {
	if o != nil && !isNil(o.BookFiles) {
		return true
	}

	return false
}

// SetBookFiles gets a reference to the given BookFileListLazyLoaded and assigns it to the BookFiles field.
func (o *Edition) SetBookFiles(v BookFileListLazyLoaded) {
	o.BookFiles = &v
}

func (o Edition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.BookId) {
		toSerialize["bookId"] = o.BookId
	}
	if o.ForeignEditionId.IsSet() {
		toSerialize["foreignEditionId"] = o.ForeignEditionId.Get()
	}
	if o.TitleSlug.IsSet() {
		toSerialize["titleSlug"] = o.TitleSlug.Get()
	}
	if o.Isbn13.IsSet() {
		toSerialize["isbn13"] = o.Isbn13.Get()
	}
	if o.Asin.IsSet() {
		toSerialize["asin"] = o.Asin.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.Overview.IsSet() {
		toSerialize["overview"] = o.Overview.Get()
	}
	if o.Format.IsSet() {
		toSerialize["format"] = o.Format.Get()
	}
	if !isNil(o.IsEbook) {
		toSerialize["isEbook"] = o.IsEbook
	}
	if o.Disambiguation.IsSet() {
		toSerialize["disambiguation"] = o.Disambiguation.Get()
	}
	if o.Publisher.IsSet() {
		toSerialize["publisher"] = o.Publisher.Get()
	}
	if !isNil(o.PageCount) {
		toSerialize["pageCount"] = o.PageCount
	}
	if o.ReleaseDate.IsSet() {
		toSerialize["releaseDate"] = o.ReleaseDate.Get()
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if !isNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !isNil(o.ManualAdd) {
		toSerialize["manualAdd"] = o.ManualAdd
	}
	if !isNil(o.Book) {
		toSerialize["book"] = o.Book
	}
	if !isNil(o.BookFiles) {
		toSerialize["bookFiles"] = o.BookFiles
	}
	return json.Marshal(toSerialize)
}

type NullableEdition struct {
	value *Edition
	isSet bool
}

func (v NullableEdition) Get() *Edition {
	return v.value
}

func (v *NullableEdition) Set(val *Edition) {
	v.value = val
	v.isSet = true
}

func (v NullableEdition) IsSet() bool {
	return v.isSet
}

func (v *NullableEdition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdition(val *Edition) *NullableEdition {
	return &NullableEdition{value: val, isSet: true}
}

func (v NullableEdition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


