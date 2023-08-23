/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AwardVouchersTypeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwardVouchersTypeInner{}

// AwardVouchersTypeInner struct for AwardVouchersTypeInner
type AwardVouchersTypeInner struct {
	// Membership Award code applied on the reservation.
	AwardCode *string `json:"awardCode,omitempty"`
	// Membership Award number.
	VoucherNo *string `json:"voucherNo,omitempty"`
}

// NewAwardVouchersTypeInner instantiates a new AwardVouchersTypeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwardVouchersTypeInner() *AwardVouchersTypeInner {
	this := AwardVouchersTypeInner{}
	return &this
}

// NewAwardVouchersTypeInnerWithDefaults instantiates a new AwardVouchersTypeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwardVouchersTypeInnerWithDefaults() *AwardVouchersTypeInner {
	this := AwardVouchersTypeInner{}
	return &this
}

// GetAwardCode returns the AwardCode field value if set, zero value otherwise.
func (o *AwardVouchersTypeInner) GetAwardCode() string {
	if o == nil || IsNil(o.AwardCode) {
		var ret string
		return ret
	}
	return *o.AwardCode
}

// GetAwardCodeOk returns a tuple with the AwardCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwardVouchersTypeInner) GetAwardCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AwardCode) {
		return nil, false
	}
	return o.AwardCode, true
}

// HasAwardCode returns a boolean if a field has been set.
func (o *AwardVouchersTypeInner) HasAwardCode() bool {
	if o != nil && !IsNil(o.AwardCode) {
		return true
	}

	return false
}

// SetAwardCode gets a reference to the given string and assigns it to the AwardCode field.
func (o *AwardVouchersTypeInner) SetAwardCode(v string) {
	o.AwardCode = &v
}

// GetVoucherNo returns the VoucherNo field value if set, zero value otherwise.
func (o *AwardVouchersTypeInner) GetVoucherNo() string {
	if o == nil || IsNil(o.VoucherNo) {
		var ret string
		return ret
	}
	return *o.VoucherNo
}

// GetVoucherNoOk returns a tuple with the VoucherNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwardVouchersTypeInner) GetVoucherNoOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherNo) {
		return nil, false
	}
	return o.VoucherNo, true
}

// HasVoucherNo returns a boolean if a field has been set.
func (o *AwardVouchersTypeInner) HasVoucherNo() bool {
	if o != nil && !IsNil(o.VoucherNo) {
		return true
	}

	return false
}

// SetVoucherNo gets a reference to the given string and assigns it to the VoucherNo field.
func (o *AwardVouchersTypeInner) SetVoucherNo(v string) {
	o.VoucherNo = &v
}

func (o AwardVouchersTypeInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwardVouchersTypeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwardCode) {
		toSerialize["awardCode"] = o.AwardCode
	}
	if !IsNil(o.VoucherNo) {
		toSerialize["voucherNo"] = o.VoucherNo
	}
	return toSerialize, nil
}

type NullableAwardVouchersTypeInner struct {
	value *AwardVouchersTypeInner
	isSet bool
}

func (v NullableAwardVouchersTypeInner) Get() *AwardVouchersTypeInner {
	return v.value
}

func (v *NullableAwardVouchersTypeInner) Set(val *AwardVouchersTypeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAwardVouchersTypeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAwardVouchersTypeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwardVouchersTypeInner(val *AwardVouchersTypeInner) *NullableAwardVouchersTypeInner {
	return &NullableAwardVouchersTypeInner{value: val, isSet: true}
}

func (v NullableAwardVouchersTypeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwardVouchersTypeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


