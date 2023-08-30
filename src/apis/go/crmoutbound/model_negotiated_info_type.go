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

// checks if the NegotiatedInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegotiatedInfoType{}

// NegotiatedInfoType This provides information for a profile negotiated rate.
type NegotiatedInfoType struct {
	// The master identifier for multiple offices/locations under the same company profile. This is optional
	CorporateAgreementId *string `json:"corporateAgreementId,omitempty"`
	// Informational purposes only in numeric format.
	ComissionCode *string `json:"comissionCode,omitempty"`
	// The sell order.
	Order *int32 `json:"order,omitempty"`
	// Negotiated Rate is inactive or not
	Inactive *bool `json:"inactive,omitempty"`
	// The starting value of the date range.
	Start *string `json:"start,omitempty"`
	// The ending value of the date range.
	End *string `json:"end,omitempty"`
}

// NewNegotiatedInfoType instantiates a new NegotiatedInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegotiatedInfoType() *NegotiatedInfoType {
	this := NegotiatedInfoType{}
	return &this
}

// NewNegotiatedInfoTypeWithDefaults instantiates a new NegotiatedInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegotiatedInfoTypeWithDefaults() *NegotiatedInfoType {
	this := NegotiatedInfoType{}
	return &this
}

// GetCorporateAgreementId returns the CorporateAgreementId field value if set, zero value otherwise.
func (o *NegotiatedInfoType) GetCorporateAgreementId() string {
	if o == nil || IsNil(o.CorporateAgreementId) {
		var ret string
		return ret
	}
	return *o.CorporateAgreementId
}

// GetCorporateAgreementIdOk returns a tuple with the CorporateAgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiatedInfoType) GetCorporateAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorporateAgreementId) {
		return nil, false
	}
	return o.CorporateAgreementId, true
}

// HasCorporateAgreementId returns a boolean if a field has been set.
func (o *NegotiatedInfoType) HasCorporateAgreementId() bool {
	if o != nil && !IsNil(o.CorporateAgreementId) {
		return true
	}

	return false
}

// SetCorporateAgreementId gets a reference to the given string and assigns it to the CorporateAgreementId field.
func (o *NegotiatedInfoType) SetCorporateAgreementId(v string) {
	o.CorporateAgreementId = &v
}

// GetComissionCode returns the ComissionCode field value if set, zero value otherwise.
func (o *NegotiatedInfoType) GetComissionCode() string {
	if o == nil || IsNil(o.ComissionCode) {
		var ret string
		return ret
	}
	return *o.ComissionCode
}

// GetComissionCodeOk returns a tuple with the ComissionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiatedInfoType) GetComissionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ComissionCode) {
		return nil, false
	}
	return o.ComissionCode, true
}

// HasComissionCode returns a boolean if a field has been set.
func (o *NegotiatedInfoType) HasComissionCode() bool {
	if o != nil && !IsNil(o.ComissionCode) {
		return true
	}

	return false
}

// SetComissionCode gets a reference to the given string and assigns it to the ComissionCode field.
func (o *NegotiatedInfoType) SetComissionCode(v string) {
	o.ComissionCode = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *NegotiatedInfoType) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiatedInfoType) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *NegotiatedInfoType) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *NegotiatedInfoType) SetOrder(v int32) {
	o.Order = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *NegotiatedInfoType) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiatedInfoType) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *NegotiatedInfoType) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *NegotiatedInfoType) SetInactive(v bool) {
	o.Inactive = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *NegotiatedInfoType) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiatedInfoType) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *NegotiatedInfoType) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *NegotiatedInfoType) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *NegotiatedInfoType) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegotiatedInfoType) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *NegotiatedInfoType) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *NegotiatedInfoType) SetEnd(v string) {
	o.End = &v
}

func (o NegotiatedInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegotiatedInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorporateAgreementId) {
		toSerialize["corporateAgreementId"] = o.CorporateAgreementId
	}
	if !IsNil(o.ComissionCode) {
		toSerialize["comissionCode"] = o.ComissionCode
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableNegotiatedInfoType struct {
	value *NegotiatedInfoType
	isSet bool
}

func (v NullableNegotiatedInfoType) Get() *NegotiatedInfoType {
	return v.value
}

func (v *NullableNegotiatedInfoType) Set(val *NegotiatedInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableNegotiatedInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableNegotiatedInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegotiatedInfoType(val *NegotiatedInfoType) *NullableNegotiatedInfoType {
	return &NullableNegotiatedInfoType{value: val, isSet: true}
}

func (v NullableNegotiatedInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegotiatedInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


