/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CompTransactionCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompTransactionCriteria{}

// CompTransactionCriteria Request object to submit a comp transaction to gaming system
type CompTransactionCriteria struct {
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	HotelId *string `json:"hotelId,omitempty"`
	// Transaction Number for which request is being submitted.
	TrxNo *float32 `json:"trxNo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCompTransactionCriteria instantiates a new CompTransactionCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompTransactionCriteria() *CompTransactionCriteria {
	this := CompTransactionCriteria{}
	return &this
}

// NewCompTransactionCriteriaWithDefaults instantiates a new CompTransactionCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompTransactionCriteriaWithDefaults() *CompTransactionCriteria {
	this := CompTransactionCriteria{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *CompTransactionCriteria) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompTransactionCriteria) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *CompTransactionCriteria) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *CompTransactionCriteria) SetHotelId(v string) {
	o.HotelId = &v
}

// GetTrxNo returns the TrxNo field value if set, zero value otherwise.
func (o *CompTransactionCriteria) GetTrxNo() float32 {
	if o == nil || IsNil(o.TrxNo) {
		var ret float32
		return ret
	}
	return *o.TrxNo
}

// GetTrxNoOk returns a tuple with the TrxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompTransactionCriteria) GetTrxNoOk() (*float32, bool) {
	if o == nil || IsNil(o.TrxNo) {
		return nil, false
	}
	return o.TrxNo, true
}

// HasTrxNo returns a boolean if a field has been set.
func (o *CompTransactionCriteria) HasTrxNo() bool {
	if o != nil && !IsNil(o.TrxNo) {
		return true
	}

	return false
}

// SetTrxNo gets a reference to the given float32 and assigns it to the TrxNo field.
func (o *CompTransactionCriteria) SetTrxNo(v float32) {
	o.TrxNo = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CompTransactionCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompTransactionCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CompTransactionCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CompTransactionCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CompTransactionCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompTransactionCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CompTransactionCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CompTransactionCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CompTransactionCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompTransactionCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.TrxNo) {
		toSerialize["trxNo"] = o.TrxNo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCompTransactionCriteria struct {
	value *CompTransactionCriteria
	isSet bool
}

func (v NullableCompTransactionCriteria) Get() *CompTransactionCriteria {
	return v.value
}

func (v *NullableCompTransactionCriteria) Set(val *CompTransactionCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCompTransactionCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCompTransactionCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompTransactionCriteria(val *CompTransactionCriteria) *NullableCompTransactionCriteria {
	return &NullableCompTransactionCriteria{value: val, isSet: true}
}

func (v NullableCompTransactionCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompTransactionCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


