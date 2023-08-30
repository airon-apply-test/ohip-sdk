/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the StagedProfileCompanyType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StagedProfileCompanyType{}

// StagedProfileCompanyType struct for StagedProfileCompanyType
type StagedProfileCompanyType struct {
	// Name of the company.
	CompanyName *string `json:"companyName,omitempty"`
	// Second Name of the company.
	CompanyName2 *string `json:"companyName2,omitempty"`
	// Third Name of the company.
	CompanyName3 *string `json:"companyName3,omitempty"`
	// Alternate Name of the Company. Mainly, it's the name of the company written in the Alternate Language.
	AlternateName *string `json:"alternateName,omitempty"`
	// Alternate Language of the company.
	AlternateLanguage *string `json:"alternateLanguage,omitempty"`
	// A jurisdiction(Territory) in which a company is authorized to do business.
	BusinessLocale *string `json:"businessLocale,omitempty"`
	IATAInfo *IATAInfoType `json:"iATAInfo,omitempty"`
	// Business Title.
	BusinessTitle *string `json:"businessTitle,omitempty"`
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
	// Commission code of the company.
	CommissionCode *string `json:"commissionCode,omitempty"`
	// Credit Rating of the company.
	CreditRating *string `json:"creditRating,omitempty"`
	// The type of corporate ID
	CorporateIdType *string `json:"corporateIdType,omitempty"`
	// VIP status of the company.
	VipStatus *string `json:"vipStatus,omitempty"`
	// Description of the VIP status.
	VipDescription *string `json:"vipDescription,omitempty"`
	// This element tells if profile is blacklisted or not.
	Blacklist *bool `json:"blacklist,omitempty"`
	Errors *StagedProfileCustomerTypeErrors `json:"errors,omitempty"`
}

// NewStagedProfileCompanyType instantiates a new StagedProfileCompanyType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagedProfileCompanyType() *StagedProfileCompanyType {
	this := StagedProfileCompanyType{}
	return &this
}

// NewStagedProfileCompanyTypeWithDefaults instantiates a new StagedProfileCompanyType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagedProfileCompanyTypeWithDefaults() *StagedProfileCompanyType {
	this := StagedProfileCompanyType{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *StagedProfileCompanyType) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCompanyName2 returns the CompanyName2 field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCompanyName2() string {
	if o == nil || IsNil(o.CompanyName2) {
		var ret string
		return ret
	}
	return *o.CompanyName2
}

// GetCompanyName2Ok returns a tuple with the CompanyName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCompanyName2Ok() (*string, bool) {
	if o == nil || IsNil(o.CompanyName2) {
		return nil, false
	}
	return o.CompanyName2, true
}

// HasCompanyName2 returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCompanyName2() bool {
	if o != nil && !IsNil(o.CompanyName2) {
		return true
	}

	return false
}

// SetCompanyName2 gets a reference to the given string and assigns it to the CompanyName2 field.
func (o *StagedProfileCompanyType) SetCompanyName2(v string) {
	o.CompanyName2 = &v
}

// GetCompanyName3 returns the CompanyName3 field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCompanyName3() string {
	if o == nil || IsNil(o.CompanyName3) {
		var ret string
		return ret
	}
	return *o.CompanyName3
}

// GetCompanyName3Ok returns a tuple with the CompanyName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCompanyName3Ok() (*string, bool) {
	if o == nil || IsNil(o.CompanyName3) {
		return nil, false
	}
	return o.CompanyName3, true
}

// HasCompanyName3 returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCompanyName3() bool {
	if o != nil && !IsNil(o.CompanyName3) {
		return true
	}

	return false
}

// SetCompanyName3 gets a reference to the given string and assigns it to the CompanyName3 field.
func (o *StagedProfileCompanyType) SetCompanyName3(v string) {
	o.CompanyName3 = &v
}

// GetAlternateName returns the AlternateName field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetAlternateName() string {
	if o == nil || IsNil(o.AlternateName) {
		var ret string
		return ret
	}
	return *o.AlternateName
}

// GetAlternateNameOk returns a tuple with the AlternateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetAlternateNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateName) {
		return nil, false
	}
	return o.AlternateName, true
}

// HasAlternateName returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasAlternateName() bool {
	if o != nil && !IsNil(o.AlternateName) {
		return true
	}

	return false
}

// SetAlternateName gets a reference to the given string and assigns it to the AlternateName field.
func (o *StagedProfileCompanyType) SetAlternateName(v string) {
	o.AlternateName = &v
}

// GetAlternateLanguage returns the AlternateLanguage field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetAlternateLanguage() string {
	if o == nil || IsNil(o.AlternateLanguage) {
		var ret string
		return ret
	}
	return *o.AlternateLanguage
}

// GetAlternateLanguageOk returns a tuple with the AlternateLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetAlternateLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateLanguage) {
		return nil, false
	}
	return o.AlternateLanguage, true
}

// HasAlternateLanguage returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasAlternateLanguage() bool {
	if o != nil && !IsNil(o.AlternateLanguage) {
		return true
	}

	return false
}

// SetAlternateLanguage gets a reference to the given string and assigns it to the AlternateLanguage field.
func (o *StagedProfileCompanyType) SetAlternateLanguage(v string) {
	o.AlternateLanguage = &v
}

// GetBusinessLocale returns the BusinessLocale field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetBusinessLocale() string {
	if o == nil || IsNil(o.BusinessLocale) {
		var ret string
		return ret
	}
	return *o.BusinessLocale
}

// GetBusinessLocaleOk returns a tuple with the BusinessLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetBusinessLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLocale) {
		return nil, false
	}
	return o.BusinessLocale, true
}

// HasBusinessLocale returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasBusinessLocale() bool {
	if o != nil && !IsNil(o.BusinessLocale) {
		return true
	}

	return false
}

// SetBusinessLocale gets a reference to the given string and assigns it to the BusinessLocale field.
func (o *StagedProfileCompanyType) SetBusinessLocale(v string) {
	o.BusinessLocale = &v
}

// GetIATAInfo returns the IATAInfo field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetIATAInfo() IATAInfoType {
	if o == nil || IsNil(o.IATAInfo) {
		var ret IATAInfoType
		return ret
	}
	return *o.IATAInfo
}

// GetIATAInfoOk returns a tuple with the IATAInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetIATAInfoOk() (*IATAInfoType, bool) {
	if o == nil || IsNil(o.IATAInfo) {
		return nil, false
	}
	return o.IATAInfo, true
}

// HasIATAInfo returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasIATAInfo() bool {
	if o != nil && !IsNil(o.IATAInfo) {
		return true
	}

	return false
}

// SetIATAInfo gets a reference to the given IATAInfoType and assigns it to the IATAInfo field.
func (o *StagedProfileCompanyType) SetIATAInfo(v IATAInfoType) {
	o.IATAInfo = &v
}

// GetBusinessTitle returns the BusinessTitle field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetBusinessTitle() string {
	if o == nil || IsNil(o.BusinessTitle) {
		var ret string
		return ret
	}
	return *o.BusinessTitle
}

// GetBusinessTitleOk returns a tuple with the BusinessTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetBusinessTitleOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessTitle) {
		return nil, false
	}
	return o.BusinessTitle, true
}

// HasBusinessTitle returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasBusinessTitle() bool {
	if o != nil && !IsNil(o.BusinessTitle) {
		return true
	}

	return false
}

// SetBusinessTitle gets a reference to the given string and assigns it to the BusinessTitle field.
func (o *StagedProfileCompanyType) SetBusinessTitle(v string) {
	o.BusinessTitle = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *StagedProfileCompanyType) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *StagedProfileCompanyType) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetDecimalPlaces returns the DecimalPlaces field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetDecimalPlaces() int32 {
	if o == nil || IsNil(o.DecimalPlaces) {
		var ret int32
		return ret
	}
	return *o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil || IsNil(o.DecimalPlaces) {
		return nil, false
	}
	return o.DecimalPlaces, true
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasDecimalPlaces() bool {
	if o != nil && !IsNil(o.DecimalPlaces) {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given int32 and assigns it to the DecimalPlaces field.
func (o *StagedProfileCompanyType) SetDecimalPlaces(v int32) {
	o.DecimalPlaces = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *StagedProfileCompanyType) SetLanguage(v string) {
	o.Language = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *StagedProfileCompanyType) SetNationality(v string) {
	o.Nationality = &v
}

// GetNationalityDescription returns the NationalityDescription field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetNationalityDescription() string {
	if o == nil || IsNil(o.NationalityDescription) {
		var ret string
		return ret
	}
	return *o.NationalityDescription
}

// GetNationalityDescriptionOk returns a tuple with the NationalityDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetNationalityDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.NationalityDescription) {
		return nil, false
	}
	return o.NationalityDescription, true
}

// HasNationalityDescription returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasNationalityDescription() bool {
	if o != nil && !IsNil(o.NationalityDescription) {
		return true
	}

	return false
}

// SetNationalityDescription gets a reference to the given string and assigns it to the NationalityDescription field.
func (o *StagedProfileCompanyType) SetNationalityDescription(v string) {
	o.NationalityDescription = &v
}

// GetCommissionCode returns the CommissionCode field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCommissionCode() string {
	if o == nil || IsNil(o.CommissionCode) {
		var ret string
		return ret
	}
	return *o.CommissionCode
}

// GetCommissionCodeOk returns a tuple with the CommissionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCommissionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CommissionCode) {
		return nil, false
	}
	return o.CommissionCode, true
}

// HasCommissionCode returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCommissionCode() bool {
	if o != nil && !IsNil(o.CommissionCode) {
		return true
	}

	return false
}

// SetCommissionCode gets a reference to the given string and assigns it to the CommissionCode field.
func (o *StagedProfileCompanyType) SetCommissionCode(v string) {
	o.CommissionCode = &v
}

// GetCreditRating returns the CreditRating field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCreditRating() string {
	if o == nil || IsNil(o.CreditRating) {
		var ret string
		return ret
	}
	return *o.CreditRating
}

// GetCreditRatingOk returns a tuple with the CreditRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCreditRatingOk() (*string, bool) {
	if o == nil || IsNil(o.CreditRating) {
		return nil, false
	}
	return o.CreditRating, true
}

// HasCreditRating returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCreditRating() bool {
	if o != nil && !IsNil(o.CreditRating) {
		return true
	}

	return false
}

// SetCreditRating gets a reference to the given string and assigns it to the CreditRating field.
func (o *StagedProfileCompanyType) SetCreditRating(v string) {
	o.CreditRating = &v
}

// GetCorporateIdType returns the CorporateIdType field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetCorporateIdType() string {
	if o == nil || IsNil(o.CorporateIdType) {
		var ret string
		return ret
	}
	return *o.CorporateIdType
}

// GetCorporateIdTypeOk returns a tuple with the CorporateIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetCorporateIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CorporateIdType) {
		return nil, false
	}
	return o.CorporateIdType, true
}

// HasCorporateIdType returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasCorporateIdType() bool {
	if o != nil && !IsNil(o.CorporateIdType) {
		return true
	}

	return false
}

// SetCorporateIdType gets a reference to the given string and assigns it to the CorporateIdType field.
func (o *StagedProfileCompanyType) SetCorporateIdType(v string) {
	o.CorporateIdType = &v
}

// GetVipStatus returns the VipStatus field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetVipStatus() string {
	if o == nil || IsNil(o.VipStatus) {
		var ret string
		return ret
	}
	return *o.VipStatus
}

// GetVipStatusOk returns a tuple with the VipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetVipStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VipStatus) {
		return nil, false
	}
	return o.VipStatus, true
}

// HasVipStatus returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasVipStatus() bool {
	if o != nil && !IsNil(o.VipStatus) {
		return true
	}

	return false
}

// SetVipStatus gets a reference to the given string and assigns it to the VipStatus field.
func (o *StagedProfileCompanyType) SetVipStatus(v string) {
	o.VipStatus = &v
}

// GetVipDescription returns the VipDescription field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetVipDescription() string {
	if o == nil || IsNil(o.VipDescription) {
		var ret string
		return ret
	}
	return *o.VipDescription
}

// GetVipDescriptionOk returns a tuple with the VipDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetVipDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.VipDescription) {
		return nil, false
	}
	return o.VipDescription, true
}

// HasVipDescription returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasVipDescription() bool {
	if o != nil && !IsNil(o.VipDescription) {
		return true
	}

	return false
}

// SetVipDescription gets a reference to the given string and assigns it to the VipDescription field.
func (o *StagedProfileCompanyType) SetVipDescription(v string) {
	o.VipDescription = &v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetBlacklist() bool {
	if o == nil || IsNil(o.Blacklist) {
		var ret bool
		return ret
	}
	return *o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetBlacklistOk() (*bool, bool) {
	if o == nil || IsNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasBlacklist() bool {
	if o != nil && !IsNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given bool and assigns it to the Blacklist field.
func (o *StagedProfileCompanyType) SetBlacklist(v bool) {
	o.Blacklist = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *StagedProfileCompanyType) GetErrors() StagedProfileCustomerTypeErrors {
	if o == nil || IsNil(o.Errors) {
		var ret StagedProfileCustomerTypeErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileCompanyType) GetErrorsOk() (*StagedProfileCustomerTypeErrors, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *StagedProfileCompanyType) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given StagedProfileCustomerTypeErrors and assigns it to the Errors field.
func (o *StagedProfileCompanyType) SetErrors(v StagedProfileCustomerTypeErrors) {
	o.Errors = &v
}

func (o StagedProfileCompanyType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StagedProfileCompanyType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.CompanyName2) {
		toSerialize["companyName2"] = o.CompanyName2
	}
	if !IsNil(o.CompanyName3) {
		toSerialize["companyName3"] = o.CompanyName3
	}
	if !IsNil(o.AlternateName) {
		toSerialize["alternateName"] = o.AlternateName
	}
	if !IsNil(o.AlternateLanguage) {
		toSerialize["alternateLanguage"] = o.AlternateLanguage
	}
	if !IsNil(o.BusinessLocale) {
		toSerialize["businessLocale"] = o.BusinessLocale
	}
	if !IsNil(o.IATAInfo) {
		toSerialize["iATAInfo"] = o.IATAInfo
	}
	if !IsNil(o.BusinessTitle) {
		toSerialize["businessTitle"] = o.BusinessTitle
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
	if !IsNil(o.CommissionCode) {
		toSerialize["commissionCode"] = o.CommissionCode
	}
	if !IsNil(o.CreditRating) {
		toSerialize["creditRating"] = o.CreditRating
	}
	if !IsNil(o.CorporateIdType) {
		toSerialize["corporateIdType"] = o.CorporateIdType
	}
	if !IsNil(o.VipStatus) {
		toSerialize["vipStatus"] = o.VipStatus
	}
	if !IsNil(o.VipDescription) {
		toSerialize["vipDescription"] = o.VipDescription
	}
	if !IsNil(o.Blacklist) {
		toSerialize["blacklist"] = o.Blacklist
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableStagedProfileCompanyType struct {
	value *StagedProfileCompanyType
	isSet bool
}

func (v NullableStagedProfileCompanyType) Get() *StagedProfileCompanyType {
	return v.value
}

func (v *NullableStagedProfileCompanyType) Set(val *StagedProfileCompanyType) {
	v.value = val
	v.isSet = true
}

func (v NullableStagedProfileCompanyType) IsSet() bool {
	return v.isSet
}

func (v *NullableStagedProfileCompanyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagedProfileCompanyType(val *StagedProfileCompanyType) *NullableStagedProfileCompanyType {
	return &NullableStagedProfileCompanyType{value: val, isSet: true}
}

func (v NullableStagedProfileCompanyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagedProfileCompanyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


