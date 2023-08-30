/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the SubAllocationTypeBlockSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubAllocationTypeBlockSecurity{}

// SubAllocationTypeBlockSecurity Block security information used for an external system. Only available if configured in OPERA Cloud.
type SubAllocationTypeBlockSecurity struct {
	// Secured from DI Display.
	SecuredFromDIdisplayYn *bool `json:"securedFromDIdisplayYn,omitempty"`
	// All Description DD Secured.
	AllDescriptionDDSecured *bool `json:"allDescriptionDDSecured,omitempty"`
	// Rates Secured from GNR.
	RatesSecuredfromGNR *bool `json:"ratesSecuredfromGNR,omitempty"`
	// Rates Secured from All Displays.
	RatesSecuredfromAllDisplays *bool `json:"ratesSecuredfromAllDisplays,omitempty"`
	// Housing Information Secured.
	HousingInformationSecured *bool `json:"housingInformationSecured,omitempty"`
	// Number of Days Deposit due to guarantee the guest booking.
	DaysDepositRequired *float32 `json:"daysDepositRequired,omitempty"`
	// Number of days before the arrival date a reservation can be canceled without losing the deposit.
	AdvanceCancelDays *float32 `json:"advanceCancelDays,omitempty"`
	// Return One Day at a time.
	ReturnOneDayAtTimeYn *bool `json:"returnOneDayAtTimeYn,omitempty"`
	// Determines if Travel Agent commission will be paid on reservations booked through the HOLIDEX Plus TACP program.
	CommissionableYn *bool `json:"commissionableYn,omitempty"`
}

// NewSubAllocationTypeBlockSecurity instantiates a new SubAllocationTypeBlockSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAllocationTypeBlockSecurity() *SubAllocationTypeBlockSecurity {
	this := SubAllocationTypeBlockSecurity{}
	return &this
}

// NewSubAllocationTypeBlockSecurityWithDefaults instantiates a new SubAllocationTypeBlockSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAllocationTypeBlockSecurityWithDefaults() *SubAllocationTypeBlockSecurity {
	this := SubAllocationTypeBlockSecurity{}
	return &this
}

// GetSecuredFromDIdisplayYn returns the SecuredFromDIdisplayYn field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetSecuredFromDIdisplayYn() bool {
	if o == nil || IsNil(o.SecuredFromDIdisplayYn) {
		var ret bool
		return ret
	}
	return *o.SecuredFromDIdisplayYn
}

// GetSecuredFromDIdisplayYnOk returns a tuple with the SecuredFromDIdisplayYn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetSecuredFromDIdisplayYnOk() (*bool, bool) {
	if o == nil || IsNil(o.SecuredFromDIdisplayYn) {
		return nil, false
	}
	return o.SecuredFromDIdisplayYn, true
}

// HasSecuredFromDIdisplayYn returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasSecuredFromDIdisplayYn() bool {
	if o != nil && !IsNil(o.SecuredFromDIdisplayYn) {
		return true
	}

	return false
}

// SetSecuredFromDIdisplayYn gets a reference to the given bool and assigns it to the SecuredFromDIdisplayYn field.
func (o *SubAllocationTypeBlockSecurity) SetSecuredFromDIdisplayYn(v bool) {
	o.SecuredFromDIdisplayYn = &v
}

// GetAllDescriptionDDSecured returns the AllDescriptionDDSecured field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetAllDescriptionDDSecured() bool {
	if o == nil || IsNil(o.AllDescriptionDDSecured) {
		var ret bool
		return ret
	}
	return *o.AllDescriptionDDSecured
}

// GetAllDescriptionDDSecuredOk returns a tuple with the AllDescriptionDDSecured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetAllDescriptionDDSecuredOk() (*bool, bool) {
	if o == nil || IsNil(o.AllDescriptionDDSecured) {
		return nil, false
	}
	return o.AllDescriptionDDSecured, true
}

// HasAllDescriptionDDSecured returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasAllDescriptionDDSecured() bool {
	if o != nil && !IsNil(o.AllDescriptionDDSecured) {
		return true
	}

	return false
}

// SetAllDescriptionDDSecured gets a reference to the given bool and assigns it to the AllDescriptionDDSecured field.
func (o *SubAllocationTypeBlockSecurity) SetAllDescriptionDDSecured(v bool) {
	o.AllDescriptionDDSecured = &v
}

// GetRatesSecuredfromGNR returns the RatesSecuredfromGNR field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetRatesSecuredfromGNR() bool {
	if o == nil || IsNil(o.RatesSecuredfromGNR) {
		var ret bool
		return ret
	}
	return *o.RatesSecuredfromGNR
}

// GetRatesSecuredfromGNROk returns a tuple with the RatesSecuredfromGNR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetRatesSecuredfromGNROk() (*bool, bool) {
	if o == nil || IsNil(o.RatesSecuredfromGNR) {
		return nil, false
	}
	return o.RatesSecuredfromGNR, true
}

// HasRatesSecuredfromGNR returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasRatesSecuredfromGNR() bool {
	if o != nil && !IsNil(o.RatesSecuredfromGNR) {
		return true
	}

	return false
}

// SetRatesSecuredfromGNR gets a reference to the given bool and assigns it to the RatesSecuredfromGNR field.
func (o *SubAllocationTypeBlockSecurity) SetRatesSecuredfromGNR(v bool) {
	o.RatesSecuredfromGNR = &v
}

// GetRatesSecuredfromAllDisplays returns the RatesSecuredfromAllDisplays field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetRatesSecuredfromAllDisplays() bool {
	if o == nil || IsNil(o.RatesSecuredfromAllDisplays) {
		var ret bool
		return ret
	}
	return *o.RatesSecuredfromAllDisplays
}

// GetRatesSecuredfromAllDisplaysOk returns a tuple with the RatesSecuredfromAllDisplays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetRatesSecuredfromAllDisplaysOk() (*bool, bool) {
	if o == nil || IsNil(o.RatesSecuredfromAllDisplays) {
		return nil, false
	}
	return o.RatesSecuredfromAllDisplays, true
}

// HasRatesSecuredfromAllDisplays returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasRatesSecuredfromAllDisplays() bool {
	if o != nil && !IsNil(o.RatesSecuredfromAllDisplays) {
		return true
	}

	return false
}

// SetRatesSecuredfromAllDisplays gets a reference to the given bool and assigns it to the RatesSecuredfromAllDisplays field.
func (o *SubAllocationTypeBlockSecurity) SetRatesSecuredfromAllDisplays(v bool) {
	o.RatesSecuredfromAllDisplays = &v
}

// GetHousingInformationSecured returns the HousingInformationSecured field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetHousingInformationSecured() bool {
	if o == nil || IsNil(o.HousingInformationSecured) {
		var ret bool
		return ret
	}
	return *o.HousingInformationSecured
}

// GetHousingInformationSecuredOk returns a tuple with the HousingInformationSecured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetHousingInformationSecuredOk() (*bool, bool) {
	if o == nil || IsNil(o.HousingInformationSecured) {
		return nil, false
	}
	return o.HousingInformationSecured, true
}

// HasHousingInformationSecured returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasHousingInformationSecured() bool {
	if o != nil && !IsNil(o.HousingInformationSecured) {
		return true
	}

	return false
}

// SetHousingInformationSecured gets a reference to the given bool and assigns it to the HousingInformationSecured field.
func (o *SubAllocationTypeBlockSecurity) SetHousingInformationSecured(v bool) {
	o.HousingInformationSecured = &v
}

// GetDaysDepositRequired returns the DaysDepositRequired field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetDaysDepositRequired() float32 {
	if o == nil || IsNil(o.DaysDepositRequired) {
		var ret float32
		return ret
	}
	return *o.DaysDepositRequired
}

// GetDaysDepositRequiredOk returns a tuple with the DaysDepositRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetDaysDepositRequiredOk() (*float32, bool) {
	if o == nil || IsNil(o.DaysDepositRequired) {
		return nil, false
	}
	return o.DaysDepositRequired, true
}

// HasDaysDepositRequired returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasDaysDepositRequired() bool {
	if o != nil && !IsNil(o.DaysDepositRequired) {
		return true
	}

	return false
}

// SetDaysDepositRequired gets a reference to the given float32 and assigns it to the DaysDepositRequired field.
func (o *SubAllocationTypeBlockSecurity) SetDaysDepositRequired(v float32) {
	o.DaysDepositRequired = &v
}

// GetAdvanceCancelDays returns the AdvanceCancelDays field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetAdvanceCancelDays() float32 {
	if o == nil || IsNil(o.AdvanceCancelDays) {
		var ret float32
		return ret
	}
	return *o.AdvanceCancelDays
}

// GetAdvanceCancelDaysOk returns a tuple with the AdvanceCancelDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetAdvanceCancelDaysOk() (*float32, bool) {
	if o == nil || IsNil(o.AdvanceCancelDays) {
		return nil, false
	}
	return o.AdvanceCancelDays, true
}

// HasAdvanceCancelDays returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasAdvanceCancelDays() bool {
	if o != nil && !IsNil(o.AdvanceCancelDays) {
		return true
	}

	return false
}

// SetAdvanceCancelDays gets a reference to the given float32 and assigns it to the AdvanceCancelDays field.
func (o *SubAllocationTypeBlockSecurity) SetAdvanceCancelDays(v float32) {
	o.AdvanceCancelDays = &v
}

// GetReturnOneDayAtTimeYn returns the ReturnOneDayAtTimeYn field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetReturnOneDayAtTimeYn() bool {
	if o == nil || IsNil(o.ReturnOneDayAtTimeYn) {
		var ret bool
		return ret
	}
	return *o.ReturnOneDayAtTimeYn
}

// GetReturnOneDayAtTimeYnOk returns a tuple with the ReturnOneDayAtTimeYn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetReturnOneDayAtTimeYnOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnOneDayAtTimeYn) {
		return nil, false
	}
	return o.ReturnOneDayAtTimeYn, true
}

// HasReturnOneDayAtTimeYn returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasReturnOneDayAtTimeYn() bool {
	if o != nil && !IsNil(o.ReturnOneDayAtTimeYn) {
		return true
	}

	return false
}

// SetReturnOneDayAtTimeYn gets a reference to the given bool and assigns it to the ReturnOneDayAtTimeYn field.
func (o *SubAllocationTypeBlockSecurity) SetReturnOneDayAtTimeYn(v bool) {
	o.ReturnOneDayAtTimeYn = &v
}

// GetCommissionableYn returns the CommissionableYn field value if set, zero value otherwise.
func (o *SubAllocationTypeBlockSecurity) GetCommissionableYn() bool {
	if o == nil || IsNil(o.CommissionableYn) {
		var ret bool
		return ret
	}
	return *o.CommissionableYn
}

// GetCommissionableYnOk returns a tuple with the CommissionableYn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeBlockSecurity) GetCommissionableYnOk() (*bool, bool) {
	if o == nil || IsNil(o.CommissionableYn) {
		return nil, false
	}
	return o.CommissionableYn, true
}

// HasCommissionableYn returns a boolean if a field has been set.
func (o *SubAllocationTypeBlockSecurity) HasCommissionableYn() bool {
	if o != nil && !IsNil(o.CommissionableYn) {
		return true
	}

	return false
}

// SetCommissionableYn gets a reference to the given bool and assigns it to the CommissionableYn field.
func (o *SubAllocationTypeBlockSecurity) SetCommissionableYn(v bool) {
	o.CommissionableYn = &v
}

func (o SubAllocationTypeBlockSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAllocationTypeBlockSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecuredFromDIdisplayYn) {
		toSerialize["securedFromDIdisplayYn"] = o.SecuredFromDIdisplayYn
	}
	if !IsNil(o.AllDescriptionDDSecured) {
		toSerialize["allDescriptionDDSecured"] = o.AllDescriptionDDSecured
	}
	if !IsNil(o.RatesSecuredfromGNR) {
		toSerialize["ratesSecuredfromGNR"] = o.RatesSecuredfromGNR
	}
	if !IsNil(o.RatesSecuredfromAllDisplays) {
		toSerialize["ratesSecuredfromAllDisplays"] = o.RatesSecuredfromAllDisplays
	}
	if !IsNil(o.HousingInformationSecured) {
		toSerialize["housingInformationSecured"] = o.HousingInformationSecured
	}
	if !IsNil(o.DaysDepositRequired) {
		toSerialize["daysDepositRequired"] = o.DaysDepositRequired
	}
	if !IsNil(o.AdvanceCancelDays) {
		toSerialize["advanceCancelDays"] = o.AdvanceCancelDays
	}
	if !IsNil(o.ReturnOneDayAtTimeYn) {
		toSerialize["returnOneDayAtTimeYn"] = o.ReturnOneDayAtTimeYn
	}
	if !IsNil(o.CommissionableYn) {
		toSerialize["commissionableYn"] = o.CommissionableYn
	}
	return toSerialize, nil
}

type NullableSubAllocationTypeBlockSecurity struct {
	value *SubAllocationTypeBlockSecurity
	isSet bool
}

func (v NullableSubAllocationTypeBlockSecurity) Get() *SubAllocationTypeBlockSecurity {
	return v.value
}

func (v *NullableSubAllocationTypeBlockSecurity) Set(val *SubAllocationTypeBlockSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAllocationTypeBlockSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAllocationTypeBlockSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAllocationTypeBlockSecurity(val *SubAllocationTypeBlockSecurity) *NullableSubAllocationTypeBlockSecurity {
	return &NullableSubAllocationTypeBlockSecurity{value: val, isSet: true}
}

func (v NullableSubAllocationTypeBlockSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAllocationTypeBlockSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


