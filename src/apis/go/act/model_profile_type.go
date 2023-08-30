/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package act

import (
	"encoding/json"
)

// checks if the ProfileType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileType{}

// ProfileType Type provides the detailed information about the profile and its children.
type ProfileType struct {
	Customer *CustomerType `json:"customer,omitempty"`
	Company *CompanyType `json:"company,omitempty"`
	Addresses *ProfileTypeAddresses `json:"addresses,omitempty"`
	Telephones *ProfileTypeTelephones `json:"telephones,omitempty"`
	Emails *ProfileTypeEmails `json:"emails,omitempty"`
	URLs *ProfileTypeURLs `json:"uRLs,omitempty"`
	ProfileDeliveryMethods *ProfileTypeProfileDeliveryMethods `json:"profileDeliveryMethods,omitempty"`
	Relationships *ProfileTypeRelationships `json:"relationships,omitempty"`
	RelationshipsSummary *ProfileTypeRelationshipsSummary `json:"relationshipsSummary,omitempty"`
	StayReservationInfoList *ReservationStayHistoryFutureInfoType `json:"stayReservationInfoList,omitempty"`
	// Eligible for Fiscal Folio/Payload generation.
	EligibleForFiscalFolio *string `json:"eligibleForFiscalFolio,omitempty"`
	ProfileType *ProfileTypeType `json:"profileType,omitempty"`
}

// NewProfileType instantiates a new ProfileType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileType() *ProfileType {
	this := ProfileType{}
	return &this
}

// NewProfileTypeWithDefaults instantiates a new ProfileType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTypeWithDefaults() *ProfileType {
	this := ProfileType{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ProfileType) GetCustomer() CustomerType {
	if o == nil || IsNil(o.Customer) {
		var ret CustomerType
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetCustomerOk() (*CustomerType, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ProfileType) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CustomerType and assigns it to the Customer field.
func (o *ProfileType) SetCustomer(v CustomerType) {
	o.Customer = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ProfileType) GetCompany() CompanyType {
	if o == nil || IsNil(o.Company) {
		var ret CompanyType
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetCompanyOk() (*CompanyType, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ProfileType) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyType and assigns it to the Company field.
func (o *ProfileType) SetCompany(v CompanyType) {
	o.Company = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ProfileType) GetAddresses() ProfileTypeAddresses {
	if o == nil || IsNil(o.Addresses) {
		var ret ProfileTypeAddresses
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetAddressesOk() (*ProfileTypeAddresses, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ProfileType) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given ProfileTypeAddresses and assigns it to the Addresses field.
func (o *ProfileType) SetAddresses(v ProfileTypeAddresses) {
	o.Addresses = &v
}

// GetTelephones returns the Telephones field value if set, zero value otherwise.
func (o *ProfileType) GetTelephones() ProfileTypeTelephones {
	if o == nil || IsNil(o.Telephones) {
		var ret ProfileTypeTelephones
		return ret
	}
	return *o.Telephones
}

// GetTelephonesOk returns a tuple with the Telephones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetTelephonesOk() (*ProfileTypeTelephones, bool) {
	if o == nil || IsNil(o.Telephones) {
		return nil, false
	}
	return o.Telephones, true
}

// HasTelephones returns a boolean if a field has been set.
func (o *ProfileType) HasTelephones() bool {
	if o != nil && !IsNil(o.Telephones) {
		return true
	}

	return false
}

// SetTelephones gets a reference to the given ProfileTypeTelephones and assigns it to the Telephones field.
func (o *ProfileType) SetTelephones(v ProfileTypeTelephones) {
	o.Telephones = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ProfileType) GetEmails() ProfileTypeEmails {
	if o == nil || IsNil(o.Emails) {
		var ret ProfileTypeEmails
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetEmailsOk() (*ProfileTypeEmails, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ProfileType) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given ProfileTypeEmails and assigns it to the Emails field.
func (o *ProfileType) SetEmails(v ProfileTypeEmails) {
	o.Emails = &v
}

// GetURLs returns the URLs field value if set, zero value otherwise.
func (o *ProfileType) GetURLs() ProfileTypeURLs {
	if o == nil || IsNil(o.URLs) {
		var ret ProfileTypeURLs
		return ret
	}
	return *o.URLs
}

// GetURLsOk returns a tuple with the URLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetURLsOk() (*ProfileTypeURLs, bool) {
	if o == nil || IsNil(o.URLs) {
		return nil, false
	}
	return o.URLs, true
}

// HasURLs returns a boolean if a field has been set.
func (o *ProfileType) HasURLs() bool {
	if o != nil && !IsNil(o.URLs) {
		return true
	}

	return false
}

// SetURLs gets a reference to the given ProfileTypeURLs and assigns it to the URLs field.
func (o *ProfileType) SetURLs(v ProfileTypeURLs) {
	o.URLs = &v
}

// GetProfileDeliveryMethods returns the ProfileDeliveryMethods field value if set, zero value otherwise.
func (o *ProfileType) GetProfileDeliveryMethods() ProfileTypeProfileDeliveryMethods {
	if o == nil || IsNil(o.ProfileDeliveryMethods) {
		var ret ProfileTypeProfileDeliveryMethods
		return ret
	}
	return *o.ProfileDeliveryMethods
}

// GetProfileDeliveryMethodsOk returns a tuple with the ProfileDeliveryMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetProfileDeliveryMethodsOk() (*ProfileTypeProfileDeliveryMethods, bool) {
	if o == nil || IsNil(o.ProfileDeliveryMethods) {
		return nil, false
	}
	return o.ProfileDeliveryMethods, true
}

// HasProfileDeliveryMethods returns a boolean if a field has been set.
func (o *ProfileType) HasProfileDeliveryMethods() bool {
	if o != nil && !IsNil(o.ProfileDeliveryMethods) {
		return true
	}

	return false
}

// SetProfileDeliveryMethods gets a reference to the given ProfileTypeProfileDeliveryMethods and assigns it to the ProfileDeliveryMethods field.
func (o *ProfileType) SetProfileDeliveryMethods(v ProfileTypeProfileDeliveryMethods) {
	o.ProfileDeliveryMethods = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ProfileType) GetRelationships() ProfileTypeRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ProfileTypeRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetRelationshipsOk() (*ProfileTypeRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ProfileType) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ProfileTypeRelationships and assigns it to the Relationships field.
func (o *ProfileType) SetRelationships(v ProfileTypeRelationships) {
	o.Relationships = &v
}

// GetRelationshipsSummary returns the RelationshipsSummary field value if set, zero value otherwise.
func (o *ProfileType) GetRelationshipsSummary() ProfileTypeRelationshipsSummary {
	if o == nil || IsNil(o.RelationshipsSummary) {
		var ret ProfileTypeRelationshipsSummary
		return ret
	}
	return *o.RelationshipsSummary
}

// GetRelationshipsSummaryOk returns a tuple with the RelationshipsSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetRelationshipsSummaryOk() (*ProfileTypeRelationshipsSummary, bool) {
	if o == nil || IsNil(o.RelationshipsSummary) {
		return nil, false
	}
	return o.RelationshipsSummary, true
}

// HasRelationshipsSummary returns a boolean if a field has been set.
func (o *ProfileType) HasRelationshipsSummary() bool {
	if o != nil && !IsNil(o.RelationshipsSummary) {
		return true
	}

	return false
}

// SetRelationshipsSummary gets a reference to the given ProfileTypeRelationshipsSummary and assigns it to the RelationshipsSummary field.
func (o *ProfileType) SetRelationshipsSummary(v ProfileTypeRelationshipsSummary) {
	o.RelationshipsSummary = &v
}

// GetStayReservationInfoList returns the StayReservationInfoList field value if set, zero value otherwise.
func (o *ProfileType) GetStayReservationInfoList() ReservationStayHistoryFutureInfoType {
	if o == nil || IsNil(o.StayReservationInfoList) {
		var ret ReservationStayHistoryFutureInfoType
		return ret
	}
	return *o.StayReservationInfoList
}

// GetStayReservationInfoListOk returns a tuple with the StayReservationInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetStayReservationInfoListOk() (*ReservationStayHistoryFutureInfoType, bool) {
	if o == nil || IsNil(o.StayReservationInfoList) {
		return nil, false
	}
	return o.StayReservationInfoList, true
}

// HasStayReservationInfoList returns a boolean if a field has been set.
func (o *ProfileType) HasStayReservationInfoList() bool {
	if o != nil && !IsNil(o.StayReservationInfoList) {
		return true
	}

	return false
}

// SetStayReservationInfoList gets a reference to the given ReservationStayHistoryFutureInfoType and assigns it to the StayReservationInfoList field.
func (o *ProfileType) SetStayReservationInfoList(v ReservationStayHistoryFutureInfoType) {
	o.StayReservationInfoList = &v
}

// GetEligibleForFiscalFolio returns the EligibleForFiscalFolio field value if set, zero value otherwise.
func (o *ProfileType) GetEligibleForFiscalFolio() string {
	if o == nil || IsNil(o.EligibleForFiscalFolio) {
		var ret string
		return ret
	}
	return *o.EligibleForFiscalFolio
}

// GetEligibleForFiscalFolioOk returns a tuple with the EligibleForFiscalFolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetEligibleForFiscalFolioOk() (*string, bool) {
	if o == nil || IsNil(o.EligibleForFiscalFolio) {
		return nil, false
	}
	return o.EligibleForFiscalFolio, true
}

// HasEligibleForFiscalFolio returns a boolean if a field has been set.
func (o *ProfileType) HasEligibleForFiscalFolio() bool {
	if o != nil && !IsNil(o.EligibleForFiscalFolio) {
		return true
	}

	return false
}

// SetEligibleForFiscalFolio gets a reference to the given string and assigns it to the EligibleForFiscalFolio field.
func (o *ProfileType) SetEligibleForFiscalFolio(v string) {
	o.EligibleForFiscalFolio = &v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *ProfileType) GetProfileType() ProfileTypeType {
	if o == nil || IsNil(o.ProfileType) {
		var ret ProfileTypeType
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileType) GetProfileTypeOk() (*ProfileTypeType, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *ProfileType) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given ProfileTypeType and assigns it to the ProfileType field.
func (o *ProfileType) SetProfileType(v ProfileTypeType) {
	o.ProfileType = &v
}

func (o ProfileType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Telephones) {
		toSerialize["telephones"] = o.Telephones
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.URLs) {
		toSerialize["uRLs"] = o.URLs
	}
	if !IsNil(o.ProfileDeliveryMethods) {
		toSerialize["profileDeliveryMethods"] = o.ProfileDeliveryMethods
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.RelationshipsSummary) {
		toSerialize["relationshipsSummary"] = o.RelationshipsSummary
	}
	if !IsNil(o.StayReservationInfoList) {
		toSerialize["stayReservationInfoList"] = o.StayReservationInfoList
	}
	if !IsNil(o.EligibleForFiscalFolio) {
		toSerialize["eligibleForFiscalFolio"] = o.EligibleForFiscalFolio
	}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	return toSerialize, nil
}

type NullableProfileType struct {
	value *ProfileType
	isSet bool
}

func (v NullableProfileType) Get() *ProfileType {
	return v.value
}

func (v *NullableProfileType) Set(val *ProfileType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileType(val *ProfileType) *NullableProfileType {
	return &NullableProfileType{value: val, isSet: true}
}

func (v NullableProfileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


