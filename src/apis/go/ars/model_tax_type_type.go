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

// checks if the TaxTypeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxTypeType{}

// TaxTypeType Provides information about the Tax Type.
type TaxTypeType struct {
	// Code of the Hotel.
	HotelId *string `json:"hotelId,omitempty"`
	// Code of the Tax Type.
	Code *string `json:"code,omitempty"`
	// Description of the Tax Type.
	Description *string `json:"description,omitempty"`
	// A boolean flag for Collecting Agent Tax
	CollectingAgentTax *bool `json:"collectingAgentTax,omitempty"`
	// Print auto adjust information for this tax type on the tax exempt report.
	PrintAutoAdjust *bool `json:"printAutoAdjust,omitempty"`
	// Number of days after which the guest will be tax exempt. Only used for tax exempt report.
	ReportExemptDays *int32 `json:"reportExemptDays,omitempty"`
	// Tax percentage. Only used for tax exempt report.
	ReportTaxPercentage *float32 `json:"reportTaxPercentage,omitempty"`
	// Minimun Length of Stay.
	MinimumLengthOfStay *int32 `json:"minimumLengthOfStay,omitempty"`
}

// NewTaxTypeType instantiates a new TaxTypeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxTypeType() *TaxTypeType {
	this := TaxTypeType{}
	return &this
}

// NewTaxTypeTypeWithDefaults instantiates a new TaxTypeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxTypeTypeWithDefaults() *TaxTypeType {
	this := TaxTypeType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *TaxTypeType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *TaxTypeType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *TaxTypeType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TaxTypeType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TaxTypeType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TaxTypeType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TaxTypeType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TaxTypeType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TaxTypeType) SetDescription(v string) {
	o.Description = &v
}

// GetCollectingAgentTax returns the CollectingAgentTax field value if set, zero value otherwise.
func (o *TaxTypeType) GetCollectingAgentTax() bool {
	if o == nil || IsNil(o.CollectingAgentTax) {
		var ret bool
		return ret
	}
	return *o.CollectingAgentTax
}

// GetCollectingAgentTaxOk returns a tuple with the CollectingAgentTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetCollectingAgentTaxOk() (*bool, bool) {
	if o == nil || IsNil(o.CollectingAgentTax) {
		return nil, false
	}
	return o.CollectingAgentTax, true
}

// HasCollectingAgentTax returns a boolean if a field has been set.
func (o *TaxTypeType) HasCollectingAgentTax() bool {
	if o != nil && !IsNil(o.CollectingAgentTax) {
		return true
	}

	return false
}

// SetCollectingAgentTax gets a reference to the given bool and assigns it to the CollectingAgentTax field.
func (o *TaxTypeType) SetCollectingAgentTax(v bool) {
	o.CollectingAgentTax = &v
}

// GetPrintAutoAdjust returns the PrintAutoAdjust field value if set, zero value otherwise.
func (o *TaxTypeType) GetPrintAutoAdjust() bool {
	if o == nil || IsNil(o.PrintAutoAdjust) {
		var ret bool
		return ret
	}
	return *o.PrintAutoAdjust
}

// GetPrintAutoAdjustOk returns a tuple with the PrintAutoAdjust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetPrintAutoAdjustOk() (*bool, bool) {
	if o == nil || IsNil(o.PrintAutoAdjust) {
		return nil, false
	}
	return o.PrintAutoAdjust, true
}

// HasPrintAutoAdjust returns a boolean if a field has been set.
func (o *TaxTypeType) HasPrintAutoAdjust() bool {
	if o != nil && !IsNil(o.PrintAutoAdjust) {
		return true
	}

	return false
}

// SetPrintAutoAdjust gets a reference to the given bool and assigns it to the PrintAutoAdjust field.
func (o *TaxTypeType) SetPrintAutoAdjust(v bool) {
	o.PrintAutoAdjust = &v
}

// GetReportExemptDays returns the ReportExemptDays field value if set, zero value otherwise.
func (o *TaxTypeType) GetReportExemptDays() int32 {
	if o == nil || IsNil(o.ReportExemptDays) {
		var ret int32
		return ret
	}
	return *o.ReportExemptDays
}

// GetReportExemptDaysOk returns a tuple with the ReportExemptDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetReportExemptDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportExemptDays) {
		return nil, false
	}
	return o.ReportExemptDays, true
}

// HasReportExemptDays returns a boolean if a field has been set.
func (o *TaxTypeType) HasReportExemptDays() bool {
	if o != nil && !IsNil(o.ReportExemptDays) {
		return true
	}

	return false
}

// SetReportExemptDays gets a reference to the given int32 and assigns it to the ReportExemptDays field.
func (o *TaxTypeType) SetReportExemptDays(v int32) {
	o.ReportExemptDays = &v
}

// GetReportTaxPercentage returns the ReportTaxPercentage field value if set, zero value otherwise.
func (o *TaxTypeType) GetReportTaxPercentage() float32 {
	if o == nil || IsNil(o.ReportTaxPercentage) {
		var ret float32
		return ret
	}
	return *o.ReportTaxPercentage
}

// GetReportTaxPercentageOk returns a tuple with the ReportTaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetReportTaxPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.ReportTaxPercentage) {
		return nil, false
	}
	return o.ReportTaxPercentage, true
}

// HasReportTaxPercentage returns a boolean if a field has been set.
func (o *TaxTypeType) HasReportTaxPercentage() bool {
	if o != nil && !IsNil(o.ReportTaxPercentage) {
		return true
	}

	return false
}

// SetReportTaxPercentage gets a reference to the given float32 and assigns it to the ReportTaxPercentage field.
func (o *TaxTypeType) SetReportTaxPercentage(v float32) {
	o.ReportTaxPercentage = &v
}

// GetMinimumLengthOfStay returns the MinimumLengthOfStay field value if set, zero value otherwise.
func (o *TaxTypeType) GetMinimumLengthOfStay() int32 {
	if o == nil || IsNil(o.MinimumLengthOfStay) {
		var ret int32
		return ret
	}
	return *o.MinimumLengthOfStay
}

// GetMinimumLengthOfStayOk returns a tuple with the MinimumLengthOfStay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTypeType) GetMinimumLengthOfStayOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumLengthOfStay) {
		return nil, false
	}
	return o.MinimumLengthOfStay, true
}

// HasMinimumLengthOfStay returns a boolean if a field has been set.
func (o *TaxTypeType) HasMinimumLengthOfStay() bool {
	if o != nil && !IsNil(o.MinimumLengthOfStay) {
		return true
	}

	return false
}

// SetMinimumLengthOfStay gets a reference to the given int32 and assigns it to the MinimumLengthOfStay field.
func (o *TaxTypeType) SetMinimumLengthOfStay(v int32) {
	o.MinimumLengthOfStay = &v
}

func (o TaxTypeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxTypeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CollectingAgentTax) {
		toSerialize["collectingAgentTax"] = o.CollectingAgentTax
	}
	if !IsNil(o.PrintAutoAdjust) {
		toSerialize["printAutoAdjust"] = o.PrintAutoAdjust
	}
	if !IsNil(o.ReportExemptDays) {
		toSerialize["reportExemptDays"] = o.ReportExemptDays
	}
	if !IsNil(o.ReportTaxPercentage) {
		toSerialize["reportTaxPercentage"] = o.ReportTaxPercentage
	}
	if !IsNil(o.MinimumLengthOfStay) {
		toSerialize["minimumLengthOfStay"] = o.MinimumLengthOfStay
	}
	return toSerialize, nil
}

type NullableTaxTypeType struct {
	value *TaxTypeType
	isSet bool
}

func (v NullableTaxTypeType) Get() *TaxTypeType {
	return v.value
}

func (v *NullableTaxTypeType) Set(val *TaxTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxTypeType(val *TaxTypeType) *NullableTaxTypeType {
	return &NullableTaxTypeType{value: val, isSet: true}
}

func (v NullableTaxTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


