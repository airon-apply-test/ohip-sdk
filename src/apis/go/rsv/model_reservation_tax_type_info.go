/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ReservationTaxTypeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationTaxTypeInfo{}

// ReservationTaxTypeInfo Provides information about the Tax Type.
type ReservationTaxTypeInfo struct {
	// Code of the Tax Type.
	Code *string `json:"code,omitempty"`
	// Description of the Tax Type.
	Description *string `json:"description,omitempty"`
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
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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


