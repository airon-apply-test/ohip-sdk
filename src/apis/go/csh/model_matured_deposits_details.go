/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the MaturedDepositsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaturedDepositsDetails{}

// MaturedDepositsDetails Response of matured deposits.
type MaturedDepositsDetails struct {
	// Detailed information of matured deposits.
	MaturedDeposits []MaturedDepositType `json:"maturedDeposits,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewMaturedDepositsDetails instantiates a new MaturedDepositsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaturedDepositsDetails() *MaturedDepositsDetails {
	this := MaturedDepositsDetails{}
	return &this
}

// NewMaturedDepositsDetailsWithDefaults instantiates a new MaturedDepositsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaturedDepositsDetailsWithDefaults() *MaturedDepositsDetails {
	this := MaturedDepositsDetails{}
	return &this
}

// GetMaturedDeposits returns the MaturedDeposits field value if set, zero value otherwise.
func (o *MaturedDepositsDetails) GetMaturedDeposits() []MaturedDepositType {
	if o == nil || IsNil(o.MaturedDeposits) {
		var ret []MaturedDepositType
		return ret
	}
	return o.MaturedDeposits
}

// GetMaturedDepositsOk returns a tuple with the MaturedDeposits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaturedDepositsDetails) GetMaturedDepositsOk() ([]MaturedDepositType, bool) {
	if o == nil || IsNil(o.MaturedDeposits) {
		return nil, false
	}
	return o.MaturedDeposits, true
}

// HasMaturedDeposits returns a boolean if a field has been set.
func (o *MaturedDepositsDetails) HasMaturedDeposits() bool {
	if o != nil && !IsNil(o.MaturedDeposits) {
		return true
	}

	return false
}

// SetMaturedDeposits gets a reference to the given []MaturedDepositType and assigns it to the MaturedDeposits field.
func (o *MaturedDepositsDetails) SetMaturedDeposits(v []MaturedDepositType) {
	o.MaturedDeposits = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MaturedDepositsDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaturedDepositsDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MaturedDepositsDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *MaturedDepositsDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MaturedDepositsDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaturedDepositsDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MaturedDepositsDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *MaturedDepositsDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o MaturedDepositsDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaturedDepositsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaturedDeposits) {
		toSerialize["maturedDeposits"] = o.MaturedDeposits
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableMaturedDepositsDetails struct {
	value *MaturedDepositsDetails
	isSet bool
}

func (v NullableMaturedDepositsDetails) Get() *MaturedDepositsDetails {
	return v.value
}

func (v *NullableMaturedDepositsDetails) Set(val *MaturedDepositsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMaturedDepositsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMaturedDepositsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaturedDepositsDetails(val *MaturedDepositsDetails) *NullableMaturedDepositsDetails {
	return &NullableMaturedDepositsDetails{value: val, isSet: true}
}

func (v NullableMaturedDepositsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaturedDepositsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


