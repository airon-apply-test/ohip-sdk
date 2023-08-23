/*
OPERA Cloud Customer Relationship Management API

APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProfileSummaryTypeFormerName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileSummaryTypeFormerName{}

// ProfileSummaryTypeFormerName This provides name information for a person.
type ProfileSummaryTypeFormerName struct {
	// Family name, last name or Company Name.
	Name *string `json:"name,omitempty"`
	// Full display Name.
	FullName *string `json:"fullName,omitempty"`
	// Salutation of honorific (e.g. Mr., Mrs., Ms., Miss, Dr.)
	NamePrefix *string `json:"namePrefix,omitempty"`
	// Given name, first name or names.
	GivenName *string `json:"givenName,omitempty"`
	// The middle name of the person name.
	MiddleName *string `json:"middleName,omitempty"`
	// Hold various name suffixes and letters (e.g. Jr., Sr., III, Ret., Esq.)
	NameSuffix *string `json:"nameSuffix,omitempty"`
	// Degree or honors (e.g., Ph.D., M.D.)
	NameTitle *string `json:"nameTitle,omitempty"`
	NameType *PersonNameTypeType `json:"nameType,omitempty"`
	// Identifies the gender.
	Gender *string `json:"gender,omitempty"`
	// Indicates the date of birth as indicated in the document, in ISO 8601 prescribed format.
	BirthDate *string `json:"birthDate,omitempty"`
	// Indicates the date of birth as masked.
	BirthDateMasked *string `json:"birthDateMasked,omitempty"`
	// The code specifying a monetary unit. Use ISO 4217, three alpha code.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The symbol for the currency, e.g, for currencyCode USD the symbol is $.
	CurrencySymbol *string `json:"currencySymbol,omitempty"`
	// Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard \"minor unit\". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces=\"2\" to represent $85).
	DecimalPlaces *int32 `json:"decimalPlaces,omitempty"`
	// Language identification.
	Language *string `json:"language,omitempty"`
	// The supplier's ranking of the customer (e.g., VIP, numerical ranking).
	CustomerValue *string `json:"customerValue,omitempty"`
}

// NewProfileSummaryTypeFormerName instantiates a new ProfileSummaryTypeFormerName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSummaryTypeFormerName() *ProfileSummaryTypeFormerName {
	this := ProfileSummaryTypeFormerName{}
	return &this
}

// NewProfileSummaryTypeFormerNameWithDefaults instantiates a new ProfileSummaryTypeFormerName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSummaryTypeFormerNameWithDefaults() *ProfileSummaryTypeFormerName {
	this := ProfileSummaryTypeFormerName{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProfileSummaryTypeFormerName) SetName(v string) {
	o.Name = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ProfileSummaryTypeFormerName) SetFullName(v string) {
	o.FullName = &v
}

// GetNamePrefix returns the NamePrefix field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetNamePrefix() string {
	if o == nil || IsNil(o.NamePrefix) {
		var ret string
		return ret
	}
	return *o.NamePrefix
}

// GetNamePrefixOk returns a tuple with the NamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetNamePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.NamePrefix) {
		return nil, false
	}
	return o.NamePrefix, true
}

// HasNamePrefix returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasNamePrefix() bool {
	if o != nil && !IsNil(o.NamePrefix) {
		return true
	}

	return false
}

// SetNamePrefix gets a reference to the given string and assigns it to the NamePrefix field.
func (o *ProfileSummaryTypeFormerName) SetNamePrefix(v string) {
	o.NamePrefix = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *ProfileSummaryTypeFormerName) SetGivenName(v string) {
	o.GivenName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *ProfileSummaryTypeFormerName) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetNameSuffix returns the NameSuffix field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetNameSuffix() string {
	if o == nil || IsNil(o.NameSuffix) {
		var ret string
		return ret
	}
	return *o.NameSuffix
}

// GetNameSuffixOk returns a tuple with the NameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetNameSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.NameSuffix) {
		return nil, false
	}
	return o.NameSuffix, true
}

// HasNameSuffix returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasNameSuffix() bool {
	if o != nil && !IsNil(o.NameSuffix) {
		return true
	}

	return false
}

// SetNameSuffix gets a reference to the given string and assigns it to the NameSuffix field.
func (o *ProfileSummaryTypeFormerName) SetNameSuffix(v string) {
	o.NameSuffix = &v
}

// GetNameTitle returns the NameTitle field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetNameTitle() string {
	if o == nil || IsNil(o.NameTitle) {
		var ret string
		return ret
	}
	return *o.NameTitle
}

// GetNameTitleOk returns a tuple with the NameTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetNameTitleOk() (*string, bool) {
	if o == nil || IsNil(o.NameTitle) {
		return nil, false
	}
	return o.NameTitle, true
}

// HasNameTitle returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasNameTitle() bool {
	if o != nil && !IsNil(o.NameTitle) {
		return true
	}

	return false
}

// SetNameTitle gets a reference to the given string and assigns it to the NameTitle field.
func (o *ProfileSummaryTypeFormerName) SetNameTitle(v string) {
	o.NameTitle = &v
}

// GetNameType returns the NameType field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetNameType() PersonNameTypeType {
	if o == nil || IsNil(o.NameType) {
		var ret PersonNameTypeType
		return ret
	}
	return *o.NameType
}

// GetNameTypeOk returns a tuple with the NameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetNameTypeOk() (*PersonNameTypeType, bool) {
	if o == nil || IsNil(o.NameType) {
		return nil, false
	}
	return o.NameType, true
}

// HasNameType returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasNameType() bool {
	if o != nil && !IsNil(o.NameType) {
		return true
	}

	return false
}

// SetNameType gets a reference to the given PersonNameTypeType and assigns it to the NameType field.
func (o *ProfileSummaryTypeFormerName) SetNameType(v PersonNameTypeType) {
	o.NameType = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *ProfileSummaryTypeFormerName) SetGender(v string) {
	o.Gender = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *ProfileSummaryTypeFormerName) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetBirthDateMasked returns the BirthDateMasked field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetBirthDateMasked() string {
	if o == nil || IsNil(o.BirthDateMasked) {
		var ret string
		return ret
	}
	return *o.BirthDateMasked
}

// GetBirthDateMaskedOk returns a tuple with the BirthDateMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetBirthDateMaskedOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDateMasked) {
		return nil, false
	}
	return o.BirthDateMasked, true
}

// HasBirthDateMasked returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasBirthDateMasked() bool {
	if o != nil && !IsNil(o.BirthDateMasked) {
		return true
	}

	return false
}

// SetBirthDateMasked gets a reference to the given string and assigns it to the BirthDateMasked field.
func (o *ProfileSummaryTypeFormerName) SetBirthDateMasked(v string) {
	o.BirthDateMasked = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ProfileSummaryTypeFormerName) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *ProfileSummaryTypeFormerName) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetDecimalPlaces returns the DecimalPlaces field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetDecimalPlaces() int32 {
	if o == nil || IsNil(o.DecimalPlaces) {
		var ret int32
		return ret
	}
	return *o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil || IsNil(o.DecimalPlaces) {
		return nil, false
	}
	return o.DecimalPlaces, true
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasDecimalPlaces() bool {
	if o != nil && !IsNil(o.DecimalPlaces) {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given int32 and assigns it to the DecimalPlaces field.
func (o *ProfileSummaryTypeFormerName) SetDecimalPlaces(v int32) {
	o.DecimalPlaces = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ProfileSummaryTypeFormerName) SetLanguage(v string) {
	o.Language = &v
}

// GetCustomerValue returns the CustomerValue field value if set, zero value otherwise.
func (o *ProfileSummaryTypeFormerName) GetCustomerValue() string {
	if o == nil || IsNil(o.CustomerValue) {
		var ret string
		return ret
	}
	return *o.CustomerValue
}

// GetCustomerValueOk returns a tuple with the CustomerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSummaryTypeFormerName) GetCustomerValueOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerValue) {
		return nil, false
	}
	return o.CustomerValue, true
}

// HasCustomerValue returns a boolean if a field has been set.
func (o *ProfileSummaryTypeFormerName) HasCustomerValue() bool {
	if o != nil && !IsNil(o.CustomerValue) {
		return true
	}

	return false
}

// SetCustomerValue gets a reference to the given string and assigns it to the CustomerValue field.
func (o *ProfileSummaryTypeFormerName) SetCustomerValue(v string) {
	o.CustomerValue = &v
}

func (o ProfileSummaryTypeFormerName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileSummaryTypeFormerName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.NamePrefix) {
		toSerialize["namePrefix"] = o.NamePrefix
	}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.NameSuffix) {
		toSerialize["nameSuffix"] = o.NameSuffix
	}
	if !IsNil(o.NameTitle) {
		toSerialize["nameTitle"] = o.NameTitle
	}
	if !IsNil(o.NameType) {
		toSerialize["nameType"] = o.NameType
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.BirthDate) {
		toSerialize["birthDate"] = o.BirthDate
	}
	if !IsNil(o.BirthDateMasked) {
		toSerialize["birthDateMasked"] = o.BirthDateMasked
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencySymbol) {
		toSerialize["currencySymbol"] = o.CurrencySymbol
	}
	if !IsNil(o.DecimalPlaces) {
		toSerialize["decimalPlaces"] = o.DecimalPlaces
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.CustomerValue) {
		toSerialize["customerValue"] = o.CustomerValue
	}
	return toSerialize, nil
}

type NullableProfileSummaryTypeFormerName struct {
	value *ProfileSummaryTypeFormerName
	isSet bool
}

func (v NullableProfileSummaryTypeFormerName) Get() *ProfileSummaryTypeFormerName {
	return v.value
}

func (v *NullableProfileSummaryTypeFormerName) Set(val *ProfileSummaryTypeFormerName) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSummaryTypeFormerName) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSummaryTypeFormerName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSummaryTypeFormerName(val *ProfileSummaryTypeFormerName) *NullableProfileSummaryTypeFormerName {
	return &NullableProfileSummaryTypeFormerName{value: val, isSet: true}
}

func (v NullableProfileSummaryTypeFormerName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSummaryTypeFormerName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


