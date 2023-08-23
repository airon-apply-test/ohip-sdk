/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StagedProfileCustomerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StagedProfileCustomerType{}

// StagedProfileCustomerType Contains basic data on the customer's identity, location, relationships, finances, memberships, etc.
type StagedProfileCustomerType struct {
	// Detailed name information for the customer.
	PersonName []PersonNameType `json:"personName,omitempty"`
	Anonymization *AnonymizationType `json:"anonymization,omitempty"`
	CitizenCountry *CountryNameType `json:"citizenCountry,omitempty"`
	Identifications *CustomerTypeIdentifications `json:"identifications,omitempty"`
	// Profession of a person.
	Profession *string `json:"profession,omitempty"`
	AlienInfo *AlienInfoType `json:"alienInfo,omitempty"`
	BirthCountry *CountryNameType `json:"birthCountry,omitempty"`
	// Name Of the company the individual is associated with.
	LegalCompany *string `json:"legalCompany,omitempty"`
	// Business Title.
	BusinessTitle *string `json:"businessTitle,omitempty"`
	// Identifies the profile gender code selected from Gender types List of values. Gender types LOV provides the values configured at gender configuration.
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
	// Nationality code identification
	Nationality *string `json:"nationality,omitempty"`
	// Nationality code description
	NationalityDescription *string `json:"nationalityDescription,omitempty"`
	// The supplier's ranking of the customer (e.g., VIP, numerical ranking).
	CustomerValue *string `json:"customerValue,omitempty"`
	// Credit Rating of the customer.
	CreditRating *string `json:"creditRating,omitempty"`
	// VIP status of the customer.
	VipStatus *string `json:"vipStatus,omitempty"`
	// Description of the VIP status.
	VipDescription *string `json:"vipDescription,omitempty"`
	// Place of birth.
	BirthPlace *string `json:"birthPlace,omitempty"`
	// This element tells profile is property exclusive or not.
	PrivateProfile *bool `json:"privateProfile,omitempty"`
	// This element tells if profile is blacklisted or not.
	Blacklist *bool `json:"blacklist,omitempty"`
	Errors *StagedProfileCustomerTypeErrors `json:"errors,omitempty"`
	// ALternate language identification.
	AlternateLanguage *string `json:"alternateLanguage,omitempty"`
}

// NewStagedProfileCustomerType instantiates a new StagedProfileCustomerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagedProfileCustomerType() *StagedProfileCustomerType {
	this := StagedProfileCustomerType{}
	return &this
}

// NewStagedProfileCustomerTypeWithDefaults instantiates a new StagedProfileCustomerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagedProfileCustomerTypeWithDefaults() *StagedProfileCustomerType {
	this := StagedProfileCustomerType{}
	return &this
}

// GetPersonName returns the PersonName field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetPersonName() []PersonNameType {
	if o == nil || IsNil(o.PersonName) {
		var ret []PersonNameType
		return ret
	}
	return o.PersonName
}

// GetPersonNameOk returns a tuple with the PersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetPersonNameOk() ([]PersonNameType, bool) {
	if o == nil || IsNil(o.PersonName) {
		return nil, false
	}
	return o.PersonName, true
}

// HasPersonName returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasPersonName() bool {
	if o != nil && !IsNil(o.PersonName) {
		return true
	}

	return false
}

// SetPersonName gets a reference to the given []PersonNameType and assigns it to the PersonName field.
func (o *StagedProfileCustomerType) SetPersonName(v []PersonNameType) {
	o.PersonName = v
}

// GetAnonymization returns the Anonymization field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetAnonymization() AnonymizationType {
	if o == nil || IsNil(o.Anonymization) {
		var ret AnonymizationType
		return ret
	}
	return *o.Anonymization
}

// GetAnonymizationOk returns a tuple with the Anonymization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetAnonymizationOk() (*AnonymizationType, bool) {
	if o == nil || IsNil(o.Anonymization) {
		return nil, false
	}
	return o.Anonymization, true
}

// HasAnonymization returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasAnonymization() bool {
	if o != nil && !IsNil(o.Anonymization) {
		return true
	}

	return false
}

// SetAnonymization gets a reference to the given AnonymizationType and assigns it to the Anonymization field.
func (o *StagedProfileCustomerType) SetAnonymization(v AnonymizationType) {
	o.Anonymization = &v
}

// GetCitizenCountry returns the CitizenCountry field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetCitizenCountry() CountryNameType {
	if o == nil || IsNil(o.CitizenCountry) {
		var ret CountryNameType
		return ret
	}
	return *o.CitizenCountry
}

// GetCitizenCountryOk returns a tuple with the CitizenCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetCitizenCountryOk() (*CountryNameType, bool) {
	if o == nil || IsNil(o.CitizenCountry) {
		return nil, false
	}
	return o.CitizenCountry, true
}

// HasCitizenCountry returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasCitizenCountry() bool {
	if o != nil && !IsNil(o.CitizenCountry) {
		return true
	}

	return false
}

// SetCitizenCountry gets a reference to the given CountryNameType and assigns it to the CitizenCountry field.
func (o *StagedProfileCustomerType) SetCitizenCountry(v CountryNameType) {
	o.CitizenCountry = &v
}

// GetIdentifications returns the Identifications field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetIdentifications() CustomerTypeIdentifications {
	if o == nil || IsNil(o.Identifications) {
		var ret CustomerTypeIdentifications
		return ret
	}
	return *o.Identifications
}

// GetIdentificationsOk returns a tuple with the Identifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetIdentificationsOk() (*CustomerTypeIdentifications, bool) {
	if o == nil || IsNil(o.Identifications) {
		return nil, false
	}
	return o.Identifications, true
}

// HasIdentifications returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasIdentifications() bool {
	if o != nil && !IsNil(o.Identifications) {
		return true
	}

	return false
}

// SetIdentifications gets a reference to the given CustomerTypeIdentifications and assigns it to the Identifications field.
func (o *StagedProfileCustomerType) SetIdentifications(v CustomerTypeIdentifications) {
	o.Identifications = &v
}

// GetProfession returns the Profession field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetProfession() string {
	if o == nil || IsNil(o.Profession) {
		var ret string
		return ret
	}
	return *o.Profession
}

// GetProfessionOk returns a tuple with the Profession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetProfessionOk() (*string, bool) {
	if o == nil || IsNil(o.Profession) {
		return nil, false
	}
	return o.Profession, true
}

// HasProfession returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasProfession() bool {
	if o != nil && !IsNil(o.Profession) {
		return true
	}

	return false
}

// SetProfession gets a reference to the given string and assigns it to the Profession field.
func (o *StagedProfileCustomerType) SetProfession(v string) {
	o.Profession = &v
}

// GetAlienInfo returns the AlienInfo field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetAlienInfo() AlienInfoType {
	if o == nil || IsNil(o.AlienInfo) {
		var ret AlienInfoType
		return ret
	}
	return *o.AlienInfo
}

// GetAlienInfoOk returns a tuple with the AlienInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetAlienInfoOk() (*AlienInfoType, bool) {
	if o == nil || IsNil(o.AlienInfo) {
		return nil, false
	}
	return o.AlienInfo, true
}

// HasAlienInfo returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasAlienInfo() bool {
	if o != nil && !IsNil(o.AlienInfo) {
		return true
	}

	return false
}

// SetAlienInfo gets a reference to the given AlienInfoType and assigns it to the AlienInfo field.
func (o *StagedProfileCustomerType) SetAlienInfo(v AlienInfoType) {
	o.AlienInfo = &v
}

// GetBirthCountry returns the BirthCountry field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetBirthCountry() CountryNameType {
	if o == nil || IsNil(o.BirthCountry) {
		var ret CountryNameType
		return ret
	}
	return *o.BirthCountry
}

// GetBirthCountryOk returns a tuple with the BirthCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetBirthCountryOk() (*CountryNameType, bool) {
	if o == nil || IsNil(o.BirthCountry) {
		return nil, false
	}
	return o.BirthCountry, true
}

// HasBirthCountry returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasBirthCountry() bool {
	if o != nil && !IsNil(o.BirthCountry) {
		return true
	}

	return false
}

// SetBirthCountry gets a reference to the given CountryNameType and assigns it to the BirthCountry field.
func (o *StagedProfileCustomerType) SetBirthCountry(v CountryNameType) {
	o.BirthCountry = &v
}

// GetLegalCompany returns the LegalCompany field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetLegalCompany() string {
	if o == nil || IsNil(o.LegalCompany) {
		var ret string
		return ret
	}
	return *o.LegalCompany
}

// GetLegalCompanyOk returns a tuple with the LegalCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetLegalCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.LegalCompany) {
		return nil, false
	}
	return o.LegalCompany, true
}

// HasLegalCompany returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasLegalCompany() bool {
	if o != nil && !IsNil(o.LegalCompany) {
		return true
	}

	return false
}

// SetLegalCompany gets a reference to the given string and assigns it to the LegalCompany field.
func (o *StagedProfileCustomerType) SetLegalCompany(v string) {
	o.LegalCompany = &v
}

// GetBusinessTitle returns the BusinessTitle field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetBusinessTitle() string {
	if o == nil || IsNil(o.BusinessTitle) {
		var ret string
		return ret
	}
	return *o.BusinessTitle
}

// GetBusinessTitleOk returns a tuple with the BusinessTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetBusinessTitleOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessTitle) {
		return nil, false
	}
	return o.BusinessTitle, true
}

// HasBusinessTitle returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasBusinessTitle() bool {
	if o != nil && !IsNil(o.BusinessTitle) {
		return true
	}

	return false
}

// SetBusinessTitle gets a reference to the given string and assigns it to the BusinessTitle field.
func (o *StagedProfileCustomerType) SetBusinessTitle(v string) {
	o.BusinessTitle = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *StagedProfileCustomerType) SetGender(v string) {
	o.Gender = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *StagedProfileCustomerType) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetBirthDateMasked returns the BirthDateMasked field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetBirthDateMasked() string {
	if o == nil || IsNil(o.BirthDateMasked) {
		var ret string
		return ret
	}
	return *o.BirthDateMasked
}

// GetBirthDateMaskedOk returns a tuple with the BirthDateMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetBirthDateMaskedOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDateMasked) {
		return nil, false
	}
	return o.BirthDateMasked, true
}

// HasBirthDateMasked returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasBirthDateMasked() bool {
	if o != nil && !IsNil(o.BirthDateMasked) {
		return true
	}

	return false
}

// SetBirthDateMasked gets a reference to the given string and assigns it to the BirthDateMasked field.
func (o *StagedProfileCustomerType) SetBirthDateMasked(v string) {
	o.BirthDateMasked = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *StagedProfileCustomerType) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *StagedProfileCustomerType) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetDecimalPlaces returns the DecimalPlaces field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetDecimalPlaces() int32 {
	if o == nil || IsNil(o.DecimalPlaces) {
		var ret int32
		return ret
	}
	return *o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil || IsNil(o.DecimalPlaces) {
		return nil, false
	}
	return o.DecimalPlaces, true
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasDecimalPlaces() bool {
	if o != nil && !IsNil(o.DecimalPlaces) {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given int32 and assigns it to the DecimalPlaces field.
func (o *StagedProfileCustomerType) SetDecimalPlaces(v int32) {
	o.DecimalPlaces = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *StagedProfileCustomerType) SetLanguage(v string) {
	o.Language = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *StagedProfileCustomerType) SetNationality(v string) {
	o.Nationality = &v
}

// GetNationalityDescription returns the NationalityDescription field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetNationalityDescription() string {
	if o == nil || IsNil(o.NationalityDescription) {
		var ret string
		return ret
	}
	return *o.NationalityDescription
}

// GetNationalityDescriptionOk returns a tuple with the NationalityDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetNationalityDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.NationalityDescription) {
		return nil, false
	}
	return o.NationalityDescription, true
}

// HasNationalityDescription returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasNationalityDescription() bool {
	if o != nil && !IsNil(o.NationalityDescription) {
		return true
	}

	return false
}

// SetNationalityDescription gets a reference to the given string and assigns it to the NationalityDescription field.
func (o *StagedProfileCustomerType) SetNationalityDescription(v string) {
	o.NationalityDescription = &v
}

// GetCustomerValue returns the CustomerValue field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetCustomerValue() string {
	if o == nil || IsNil(o.CustomerValue) {
		var ret string
		return ret
	}
	return *o.CustomerValue
}

// GetCustomerValueOk returns a tuple with the CustomerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetCustomerValueOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerValue) {
		return nil, false
	}
	return o.CustomerValue, true
}

// HasCustomerValue returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasCustomerValue() bool {
	if o != nil && !IsNil(o.CustomerValue) {
		return true
	}

	return false
}

// SetCustomerValue gets a reference to the given string and assigns it to the CustomerValue field.
func (o *StagedProfileCustomerType) SetCustomerValue(v string) {
	o.CustomerValue = &v
}

// GetCreditRating returns the CreditRating field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetCreditRating() string {
	if o == nil || IsNil(o.CreditRating) {
		var ret string
		return ret
	}
	return *o.CreditRating
}

// GetCreditRatingOk returns a tuple with the CreditRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetCreditRatingOk() (*string, bool) {
	if o == nil || IsNil(o.CreditRating) {
		return nil, false
	}
	return o.CreditRating, true
}

// HasCreditRating returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasCreditRating() bool {
	if o != nil && !IsNil(o.CreditRating) {
		return true
	}

	return false
}

// SetCreditRating gets a reference to the given string and assigns it to the CreditRating field.
func (o *StagedProfileCustomerType) SetCreditRating(v string) {
	o.CreditRating = &v
}

// GetVipStatus returns the VipStatus field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetVipStatus() string {
	if o == nil || IsNil(o.VipStatus) {
		var ret string
		return ret
	}
	return *o.VipStatus
}

// GetVipStatusOk returns a tuple with the VipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetVipStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VipStatus) {
		return nil, false
	}
	return o.VipStatus, true
}

// HasVipStatus returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasVipStatus() bool {
	if o != nil && !IsNil(o.VipStatus) {
		return true
	}

	return false
}

// SetVipStatus gets a reference to the given string and assigns it to the VipStatus field.
func (o *StagedProfileCustomerType) SetVipStatus(v string) {
	o.VipStatus = &v
}

// GetVipDescription returns the VipDescription field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetVipDescription() string {
	if o == nil || IsNil(o.VipDescription) {
		var ret string
		return ret
	}
	return *o.VipDescription
}

// GetVipDescriptionOk returns a tuple with the VipDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetVipDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.VipDescription) {
		return nil, false
	}
	return o.VipDescription, true
}

// HasVipDescription returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasVipDescription() bool {
	if o != nil && !IsNil(o.VipDescription) {
		return true
	}

	return false
}

// SetVipDescription gets a reference to the given string and assigns it to the VipDescription field.
func (o *StagedProfileCustomerType) SetVipDescription(v string) {
	o.VipDescription = &v
}

// GetBirthPlace returns the BirthPlace field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetBirthPlace() string {
	if o == nil || IsNil(o.BirthPlace) {
		var ret string
		return ret
	}
	return *o.BirthPlace
}

// GetBirthPlaceOk returns a tuple with the BirthPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetBirthPlaceOk() (*string, bool) {
	if o == nil || IsNil(o.BirthPlace) {
		return nil, false
	}
	return o.BirthPlace, true
}

// HasBirthPlace returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasBirthPlace() bool {
	if o != nil && !IsNil(o.BirthPlace) {
		return true
	}

	return false
}

// SetBirthPlace gets a reference to the given string and assigns it to the BirthPlace field.
func (o *StagedProfileCustomerType) SetBirthPlace(v string) {
	o.BirthPlace = &v
}

// GetPrivateProfile returns the PrivateProfile field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetPrivateProfile() bool {
	if o == nil || IsNil(o.PrivateProfile) {
		var ret bool
		return ret
	}
	return *o.PrivateProfile
}

// GetPrivateProfileOk returns a tuple with the PrivateProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetPrivateProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateProfile) {
		return nil, false
	}
	return o.PrivateProfile, true
}

// HasPrivateProfile returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasPrivateProfile() bool {
	if o != nil && !IsNil(o.PrivateProfile) {
		return true
	}

	return false
}

// SetPrivateProfile gets a reference to the given bool and assigns it to the PrivateProfile field.
func (o *StagedProfileCustomerType) SetPrivateProfile(v bool) {
	o.PrivateProfile = &v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetBlacklist() bool {
	if o == nil || IsNil(o.Blacklist) {
		var ret bool
		return ret
	}
	return *o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetBlacklistOk() (*bool, bool) {
	if o == nil || IsNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasBlacklist() bool {
	if o != nil && !IsNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given bool and assigns it to the Blacklist field.
func (o *StagedProfileCustomerType) SetBlacklist(v bool) {
	o.Blacklist = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetErrors() StagedProfileCustomerTypeErrors {
	if o == nil || IsNil(o.Errors) {
		var ret StagedProfileCustomerTypeErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetErrorsOk() (*StagedProfileCustomerTypeErrors, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given StagedProfileCustomerTypeErrors and assigns it to the Errors field.
func (o *StagedProfileCustomerType) SetErrors(v StagedProfileCustomerTypeErrors) {
	o.Errors = &v
}

// GetAlternateLanguage returns the AlternateLanguage field value if set, zero value otherwise.
func (o *StagedProfileCustomerType) GetAlternateLanguage() string {
	if o == nil || IsNil(o.AlternateLanguage) {
		var ret string
		return ret
	}
	return *o.AlternateLanguage
}

// GetAlternateLanguageOk returns a tuple with the AlternateLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCustomerType) GetAlternateLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateLanguage) {
		return nil, false
	}
	return o.AlternateLanguage, true
}

// HasAlternateLanguage returns a boolean if a field has been set.
func (o *StagedProfileCustomerType) HasAlternateLanguage() bool {
	if o != nil && !IsNil(o.AlternateLanguage) {
		return true
	}

	return false
}

// SetAlternateLanguage gets a reference to the given string and assigns it to the AlternateLanguage field.
func (o *StagedProfileCustomerType) SetAlternateLanguage(v string) {
	o.AlternateLanguage = &v
}

func (o StagedProfileCustomerType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StagedProfileCustomerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PersonName) {
		toSerialize["personName"] = o.PersonName
	}
	if !IsNil(o.Anonymization) {
		toSerialize["anonymization"] = o.Anonymization
	}
	if !IsNil(o.CitizenCountry) {
		toSerialize["citizenCountry"] = o.CitizenCountry
	}
	if !IsNil(o.Identifications) {
		toSerialize["identifications"] = o.Identifications
	}
	if !IsNil(o.Profession) {
		toSerialize["profession"] = o.Profession
	}
	if !IsNil(o.AlienInfo) {
		toSerialize["alienInfo"] = o.AlienInfo
	}
	if !IsNil(o.BirthCountry) {
		toSerialize["birthCountry"] = o.BirthCountry
	}
	if !IsNil(o.LegalCompany) {
		toSerialize["legalCompany"] = o.LegalCompany
	}
	if !IsNil(o.BusinessTitle) {
		toSerialize["businessTitle"] = o.BusinessTitle
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
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.NationalityDescription) {
		toSerialize["nationalityDescription"] = o.NationalityDescription
	}
	if !IsNil(o.CustomerValue) {
		toSerialize["customerValue"] = o.CustomerValue
	}
	if !IsNil(o.CreditRating) {
		toSerialize["creditRating"] = o.CreditRating
	}
	if !IsNil(o.VipStatus) {
		toSerialize["vipStatus"] = o.VipStatus
	}
	if !IsNil(o.VipDescription) {
		toSerialize["vipDescription"] = o.VipDescription
	}
	if !IsNil(o.BirthPlace) {
		toSerialize["birthPlace"] = o.BirthPlace
	}
	if !IsNil(o.PrivateProfile) {
		toSerialize["privateProfile"] = o.PrivateProfile
	}
	if !IsNil(o.Blacklist) {
		toSerialize["blacklist"] = o.Blacklist
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.AlternateLanguage) {
		toSerialize["alternateLanguage"] = o.AlternateLanguage
	}
	return toSerialize, nil
}

type NullableStagedProfileCustomerType struct {
	value *StagedProfileCustomerType
	isSet bool
}

func (v NullableStagedProfileCustomerType) Get() *StagedProfileCustomerType {
	return v.value
}

func (v *NullableStagedProfileCustomerType) Set(val *StagedProfileCustomerType) {
	v.value = val
	v.isSet = true
}

func (v NullableStagedProfileCustomerType) IsSet() bool {
	return v.isSet
}

func (v *NullableStagedProfileCustomerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagedProfileCustomerType(val *StagedProfileCustomerType) *NullableStagedProfileCustomerType {
	return &NullableStagedProfileCustomerType{value: val, isSet: true}
}

func (v NullableStagedProfileCustomerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagedProfileCustomerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


