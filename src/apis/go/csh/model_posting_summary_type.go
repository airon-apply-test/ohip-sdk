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

// checks if the PostingSummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostingSummaryType{}

// PostingSummaryType summary of the postings by external system.
type PostingSummaryType struct {
	// Total count of postings.
	TotalPostings *int32 `json:"totalPostings,omitempty"`
	TotalAmount *CurrencyAmountType `json:"totalAmount,omitempty"`
	// Total count of checks posted.
	TotalCheckCount *int32 `json:"totalCheckCount,omitempty"`
}

// NewPostingSummaryType instantiates a new PostingSummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostingSummaryType() *PostingSummaryType {
	this := PostingSummaryType{}
	return &this
}

// NewPostingSummaryTypeWithDefaults instantiates a new PostingSummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostingSummaryTypeWithDefaults() *PostingSummaryType {
	this := PostingSummaryType{}
	return &this
}

// GetTotalPostings returns the TotalPostings field value if set, zero value otherwise.
func (o *PostingSummaryType) GetTotalPostings() int32 {
	if o == nil || IsNil(o.TotalPostings) {
		var ret int32
		return ret
	}
	return *o.TotalPostings
}

// GetTotalPostingsOk returns a tuple with the TotalPostings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostingSummaryType) GetTotalPostingsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPostings) {
		return nil, false
	}
	return o.TotalPostings, true
}

// HasTotalPostings returns a boolean if a field has been set.
func (o *PostingSummaryType) HasTotalPostings() bool {
	if o != nil && !IsNil(o.TotalPostings) {
		return true
	}

	return false
}

// SetTotalPostings gets a reference to the given int32 and assigns it to the TotalPostings field.
func (o *PostingSummaryType) SetTotalPostings(v int32) {
	o.TotalPostings = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *PostingSummaryType) GetTotalAmount() CurrencyAmountType {
	if o == nil || IsNil(o.TotalAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostingSummaryType) GetTotalAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *PostingSummaryType) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given CurrencyAmountType and assigns it to the TotalAmount field.
func (o *PostingSummaryType) SetTotalAmount(v CurrencyAmountType) {
	o.TotalAmount = &v
}

// GetTotalCheckCount returns the TotalCheckCount field value if set, zero value otherwise.
func (o *PostingSummaryType) GetTotalCheckCount() int32 {
	if o == nil || IsNil(o.TotalCheckCount) {
		var ret int32
		return ret
	}
	return *o.TotalCheckCount
}

// GetTotalCheckCountOk returns a tuple with the TotalCheckCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostingSummaryType) GetTotalCheckCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCheckCount) {
		return nil, false
	}
	return o.TotalCheckCount, true
}

// HasTotalCheckCount returns a boolean if a field has been set.
func (o *PostingSummaryType) HasTotalCheckCount() bool {
	if o != nil && !IsNil(o.TotalCheckCount) {
		return true
	}

	return false
}

// SetTotalCheckCount gets a reference to the given int32 and assigns it to the TotalCheckCount field.
func (o *PostingSummaryType) SetTotalCheckCount(v int32) {
	o.TotalCheckCount = &v
}

func (o PostingSummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostingSummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalPostings) {
		toSerialize["totalPostings"] = o.TotalPostings
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.TotalCheckCount) {
		toSerialize["totalCheckCount"] = o.TotalCheckCount
	}
	return toSerialize, nil
}

type NullablePostingSummaryType struct {
	value *PostingSummaryType
	isSet bool
}

func (v NullablePostingSummaryType) Get() *PostingSummaryType {
	return v.value
}

func (v *NullablePostingSummaryType) Set(val *PostingSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullablePostingSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullablePostingSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostingSummaryType(val *PostingSummaryType) *NullablePostingSummaryType {
	return &NullablePostingSummaryType{value: val, isSet: true}
}

func (v NullablePostingSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostingSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


