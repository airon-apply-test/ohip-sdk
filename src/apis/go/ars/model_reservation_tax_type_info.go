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

// checks if the ReservationTaxTypeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationTaxTypeInfo{}

// ReservationTaxTypeInfo Provides information about the Tax Type.
type ReservationTaxTypeInfo struct {
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
	// Tax exempt number on the profile.
	TaxExemptNo *string `json:"taxExemptNo,omitempty"`
}

// NewReservationTaxTypeInfo instantiates a new ReservationTaxTypeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationTaxTypeInfo() *ReservationTaxTypeInfo {
	this := ReservationTaxTypeInfo{}
	return &this
}

// NewReservationTaxTypeInfoWithDefaults instantiates a new ReservationTaxTypeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationTaxTypeInfoWithDefaults() *ReservationTaxTypeInfo {
	this := ReservationTaxTypeInfo{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ReservationTaxTypeInfo) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ReservationTaxTypeInfo) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReservationTaxTypeInfo) SetDescription(v string) {
	o.Description = &v
}

// GetCollectingAgentTax returns the CollectingAgentTax field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetCollectingAgentTax() bool {
	if o == nil || IsNil(o.CollectingAgentTax) {
		var ret bool
		return ret
	}
	return *o.CollectingAgentTax
}

// GetCollectingAgentTaxOk returns a tuple with the CollectingAgentTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetCollectingAgentTaxOk() (*bool, bool) {
	if o == nil || IsNil(o.CollectingAgentTax) {
		return nil, false
	}
	return o.CollectingAgentTax, true
}

// HasCollectingAgentTax returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasCollectingAgentTax() bool {
	if o != nil && !IsNil(o.CollectingAgentTax) {
		return true
	}

	return false
}

// SetCollectingAgentTax gets a reference to the given bool and assigns it to the CollectingAgentTax field.
func (o *ReservationTaxTypeInfo) SetCollectingAgentTax(v bool) {
	o.CollectingAgentTax = &v
}

// GetPrintAutoAdjust returns the PrintAutoAdjust field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetPrintAutoAdjust() bool {
	if o == nil || IsNil(o.PrintAutoAdjust) {
		var ret bool
		return ret
	}
	return *o.PrintAutoAdjust
}

// GetPrintAutoAdjustOk returns a tuple with the PrintAutoAdjust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetPrintAutoAdjustOk() (*bool, bool) {
	if o == nil || IsNil(o.PrintAutoAdjust) {
		return nil, false
	}
	return o.PrintAutoAdjust, true
}

// HasPrintAutoAdjust returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasPrintAutoAdjust() bool {
	if o != nil && !IsNil(o.PrintAutoAdjust) {
		return true
	}

	return false
}

// SetPrintAutoAdjust gets a reference to the given bool and assigns it to the PrintAutoAdjust field.
func (o *ReservationTaxTypeInfo) SetPrintAutoAdjust(v bool) {
	o.PrintAutoAdjust = &v
}

// GetReportExemptDays returns the ReportExemptDays field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetReportExemptDays() int32 {
	if o == nil || IsNil(o.ReportExemptDays) {
		var ret int32
		return ret
	}
	return *o.ReportExemptDays
}

// GetReportExemptDaysOk returns a tuple with the ReportExemptDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetReportExemptDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportExemptDays) {
		return nil, false
	}
	return o.ReportExemptDays, true
}

// HasReportExemptDays returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasReportExemptDays() bool {
	if o != nil && !IsNil(o.ReportExemptDays) {
		return true
	}

	return false
}

// SetReportExemptDays gets a reference to the given int32 and assigns it to the ReportExemptDays field.
func (o *ReservationTaxTypeInfo) SetReportExemptDays(v int32) {
	o.ReportExemptDays = &v
}

// GetReportTaxPercentage returns the ReportTaxPercentage field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetReportTaxPercentage() float32 {
	if o == nil || IsNil(o.ReportTaxPercentage) {
		var ret float32
		return ret
	}
	return *o.ReportTaxPercentage
}

// GetReportTaxPercentageOk returns a tuple with the ReportTaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetReportTaxPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.ReportTaxPercentage) {
		return nil, false
	}
	return o.ReportTaxPercentage, true
}

// HasReportTaxPercentage returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasReportTaxPercentage() bool {
	if o != nil && !IsNil(o.ReportTaxPercentage) {
		return true
	}

	return false
}

// SetReportTaxPercentage gets a reference to the given float32 and assigns it to the ReportTaxPercentage field.
func (o *ReservationTaxTypeInfo) SetReportTaxPercentage(v float32) {
	o.ReportTaxPercentage = &v
}

// GetMinimumLengthOfStay returns the MinimumLengthOfStay field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetMinimumLengthOfStay() int32 {
	if o == nil || IsNil(o.MinimumLengthOfStay) {
		var ret int32
		return ret
	}
	return *o.MinimumLengthOfStay
}

// GetMinimumLengthOfStayOk returns a tuple with the MinimumLengthOfStay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetMinimumLengthOfStayOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumLengthOfStay) {
		return nil, false
	}
	return o.MinimumLengthOfStay, true
}

// HasMinimumLengthOfStay returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasMinimumLengthOfStay() bool {
	if o != nil && !IsNil(o.MinimumLengthOfStay) {
		return true
	}

	return false
}

// SetMinimumLengthOfStay gets a reference to the given int32 and assigns it to the MinimumLengthOfStay field.
func (o *ReservationTaxTypeInfo) SetMinimumLengthOfStay(v int32) {
	o.MinimumLengthOfStay = &v
}

// GetTaxExemptNo returns the TaxExemptNo field value if set, zero value otherwise.
func (o *ReservationTaxTypeInfo) GetTaxExemptNo() string {
	if o == nil || IsNil(o.TaxExemptNo) {
		var ret string
		return ret
	}
	return *o.TaxExemptNo
}

// GetTaxExemptNoOk returns a tuple with the TaxExemptNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationTaxTypeInfo) GetTaxExemptNoOk() (*string, bool) {
	if o == nil || IsNil(o.TaxExemptNo) {
		return nil, false
	}
	return o.TaxExemptNo, true
}

// HasTaxExemptNo returns a boolean if a field has been set.
func (o *ReservationTaxTypeInfo) HasTaxExemptNo() bool {
	if o != nil && !IsNil(o.TaxExemptNo) {
		return true
	}

	return false
}

// SetTaxExemptNo gets a reference to the given string and assigns it to the TaxExemptNo field.
func (o *ReservationTaxTypeInfo) SetTaxExemptNo(v string) {
	o.TaxExemptNo = &v
}

func (o ReservationTaxTypeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationTaxTypeInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TaxExemptNo) {
		toSerialize["taxExemptNo"] = o.TaxExemptNo
	}
	return toSerialize, nil
}

type NullableReservationTaxTypeInfo struct {
	value *ReservationTaxTypeInfo
	isSet bool
}

func (v NullableReservationTaxTypeInfo) Get() *ReservationTaxTypeInfo {
	return v.value
}

func (v *NullableReservationTaxTypeInfo) Set(val *ReservationTaxTypeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationTaxTypeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationTaxTypeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationTaxTypeInfo(val *ReservationTaxTypeInfo) *NullableReservationTaxTypeInfo {
	return &NullableReservationTaxTypeInfo{value: val, isSet: true}
}

func (v NullableReservationTaxTypeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationTaxTypeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


