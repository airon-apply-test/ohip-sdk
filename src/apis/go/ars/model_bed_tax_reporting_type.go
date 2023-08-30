/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
)

// checks if the BedTaxReportingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BedTaxReportingType{}

// BedTaxReportingType This stores the information for Bed Tax Reporting. Mainly used in Maldives.
type BedTaxReportingType struct {
	// Tax Registration Number for Maldives Bed Tax Reporting.
	TaxRegistrationNo *float32 `json:"taxRegistrationNo,omitempty"`
	// Visa number used for Maldives Bed Tax Reporting
	VisaNumber *string `json:"visaNumber,omitempty"`
	// Visa Issue Date used for Maldives Bed Tax Reporting.
	VisaIssueDate *string `json:"visaIssueDate,omitempty"`
	// Visa Expiration Date used for Maldives Bed Tax Reporting
	VisaExpiryDate *string `json:"visaExpiryDate,omitempty"`
	// Number of days for which the Maldives tax is applicable.
	TaxableDays *int32 `json:"taxableDays,omitempty"`
}

// NewBedTaxReportingType instantiates a new BedTaxReportingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBedTaxReportingType() *BedTaxReportingType {
	this := BedTaxReportingType{}
	return &this
}

// NewBedTaxReportingTypeWithDefaults instantiates a new BedTaxReportingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBedTaxReportingTypeWithDefaults() *BedTaxReportingType {
	this := BedTaxReportingType{}
	return &this
}

// GetTaxRegistrationNo returns the TaxRegistrationNo field value if set, zero value otherwise.
func (o *BedTaxReportingType) GetTaxRegistrationNo() float32 {
	if o == nil || IsNil(o.TaxRegistrationNo) {
		var ret float32
		return ret
	}
	return *o.TaxRegistrationNo
}

// GetTaxRegistrationNoOk returns a tuple with the TaxRegistrationNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxReportingType) GetTaxRegistrationNoOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxRegistrationNo) {
		return nil, false
	}
	return o.TaxRegistrationNo, true
}

// HasTaxRegistrationNo returns a boolean if a field has been set.
func (o *BedTaxReportingType) HasTaxRegistrationNo() bool {
	if o != nil && !IsNil(o.TaxRegistrationNo) {
		return true
	}

	return false
}

// SetTaxRegistrationNo gets a reference to the given float32 and assigns it to the TaxRegistrationNo field.
func (o *BedTaxReportingType) SetTaxRegistrationNo(v float32) {
	o.TaxRegistrationNo = &v
}

// GetVisaNumber returns the VisaNumber field value if set, zero value otherwise.
func (o *BedTaxReportingType) GetVisaNumber() string {
	if o == nil || IsNil(o.VisaNumber) {
		var ret string
		return ret
	}
	return *o.VisaNumber
}

// GetVisaNumberOk returns a tuple with the VisaNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxReportingType) GetVisaNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VisaNumber) {
		return nil, false
	}
	return o.VisaNumber, true
}

// HasVisaNumber returns a boolean if a field has been set.
func (o *BedTaxReportingType) HasVisaNumber() bool {
	if o != nil && !IsNil(o.VisaNumber) {
		return true
	}

	return false
}

// SetVisaNumber gets a reference to the given string and assigns it to the VisaNumber field.
func (o *BedTaxReportingType) SetVisaNumber(v string) {
	o.VisaNumber = &v
}

// GetVisaIssueDate returns the VisaIssueDate field value if set, zero value otherwise.
func (o *BedTaxReportingType) GetVisaIssueDate() string {
	if o == nil || IsNil(o.VisaIssueDate) {
		var ret string
		return ret
	}
	return *o.VisaIssueDate
}

// GetVisaIssueDateOk returns a tuple with the VisaIssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxReportingType) GetVisaIssueDateOk() (*string, bool) {
	if o == nil || IsNil(o.VisaIssueDate) {
		return nil, false
	}
	return o.VisaIssueDate, true
}

// HasVisaIssueDate returns a boolean if a field has been set.
func (o *BedTaxReportingType) HasVisaIssueDate() bool {
	if o != nil && !IsNil(o.VisaIssueDate) {
		return true
	}

	return false
}

// SetVisaIssueDate gets a reference to the given string and assigns it to the VisaIssueDate field.
func (o *BedTaxReportingType) SetVisaIssueDate(v string) {
	o.VisaIssueDate = &v
}

// GetVisaExpiryDate returns the VisaExpiryDate field value if set, zero value otherwise.
func (o *BedTaxReportingType) GetVisaExpiryDate() string {
	if o == nil || IsNil(o.VisaExpiryDate) {
		var ret string
		return ret
	}
	return *o.VisaExpiryDate
}

// GetVisaExpiryDateOk returns a tuple with the VisaExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxReportingType) GetVisaExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.VisaExpiryDate) {
		return nil, false
	}
	return o.VisaExpiryDate, true
}

// HasVisaExpiryDate returns a boolean if a field has been set.
func (o *BedTaxReportingType) HasVisaExpiryDate() bool {
	if o != nil && !IsNil(o.VisaExpiryDate) {
		return true
	}

	return false
}

// SetVisaExpiryDate gets a reference to the given string and assigns it to the VisaExpiryDate field.
func (o *BedTaxReportingType) SetVisaExpiryDate(v string) {
	o.VisaExpiryDate = &v
}

// GetTaxableDays returns the TaxableDays field value if set, zero value otherwise.
func (o *BedTaxReportingType) GetTaxableDays() int32 {
	if o == nil || IsNil(o.TaxableDays) {
		var ret int32
		return ret
	}
	return *o.TaxableDays
}

// GetTaxableDaysOk returns a tuple with the TaxableDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxReportingType) GetTaxableDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxableDays) {
		return nil, false
	}
	return o.TaxableDays, true
}

// HasTaxableDays returns a boolean if a field has been set.
func (o *BedTaxReportingType) HasTaxableDays() bool {
	if o != nil && !IsNil(o.TaxableDays) {
		return true
	}

	return false
}

// SetTaxableDays gets a reference to the given int32 and assigns it to the TaxableDays field.
func (o *BedTaxReportingType) SetTaxableDays(v int32) {
	o.TaxableDays = &v
}

func (o BedTaxReportingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BedTaxReportingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxRegistrationNo) {
		toSerialize["taxRegistrationNo"] = o.TaxRegistrationNo
	}
	if !IsNil(o.VisaNumber) {
		toSerialize["visaNumber"] = o.VisaNumber
	}
	if !IsNil(o.VisaIssueDate) {
		toSerialize["visaIssueDate"] = o.VisaIssueDate
	}
	if !IsNil(o.VisaExpiryDate) {
		toSerialize["visaExpiryDate"] = o.VisaExpiryDate
	}
	if !IsNil(o.TaxableDays) {
		toSerialize["taxableDays"] = o.TaxableDays
	}
	return toSerialize, nil
}

type NullableBedTaxReportingType struct {
	value *BedTaxReportingType
	isSet bool
}

func (v NullableBedTaxReportingType) Get() *BedTaxReportingType {
	return v.value
}

func (v *NullableBedTaxReportingType) Set(val *BedTaxReportingType) {
	v.value = val
	v.isSet = true
}

func (v NullableBedTaxReportingType) IsSet() bool {
	return v.isSet
}

func (v *NullableBedTaxReportingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBedTaxReportingType(val *BedTaxReportingType) *NullableBedTaxReportingType {
	return &NullableBedTaxReportingType{value: val, isSet: true}
}

func (v NullableBedTaxReportingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBedTaxReportingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


