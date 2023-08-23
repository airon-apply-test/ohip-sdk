/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PromotionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionType{}

// PromotionType Type to specify a rate promotion. Usually attached to a reservation to indicate a specific promotion is applied to the reservation.
type PromotionType struct {
	// Promotion code associated with the rate plan.
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Promotion code Name associated with the rate plan.
	PromotionName *string `json:"promotionName,omitempty"`
	// Promotion Coupon Code when promotion is setupCode to have a valid coupon code.
	CouponCode *string `json:"couponCode,omitempty"`
}

// NewPromotionType instantiates a new PromotionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionType() *PromotionType {
	this := PromotionType{}
	return &this
}

// NewPromotionTypeWithDefaults instantiates a new PromotionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionTypeWithDefaults() *PromotionType {
	this := PromotionType{}
	return &this
}

// GetPromotionCode returns the PromotionCode field value if set, zero value otherwise.
func (o *PromotionType) GetPromotionCode() string {
	if o == nil || IsNil(o.PromotionCode) {
		var ret string
		return ret
	}
	return *o.PromotionCode
}

// GetPromotionCodeOk returns a tuple with the PromotionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionType) GetPromotionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionCode) {
		return nil, false
	}
	return o.PromotionCode, true
}

// HasPromotionCode returns a boolean if a field has been set.
func (o *PromotionType) HasPromotionCode() bool {
	if o != nil && !IsNil(o.PromotionCode) {
		return true
	}

	return false
}

// SetPromotionCode gets a reference to the given string and assigns it to the PromotionCode field.
func (o *PromotionType) SetPromotionCode(v string) {
	o.PromotionCode = &v
}

// GetPromotionName returns the PromotionName field value if set, zero value otherwise.
func (o *PromotionType) GetPromotionName() string {
	if o == nil || IsNil(o.PromotionName) {
		var ret string
		return ret
	}
	return *o.PromotionName
}

// GetPromotionNameOk returns a tuple with the PromotionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionType) GetPromotionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionName) {
		return nil, false
	}
	return o.PromotionName, true
}

// HasPromotionName returns a boolean if a field has been set.
func (o *PromotionType) HasPromotionName() bool {
	if o != nil && !IsNil(o.PromotionName) {
		return true
	}

	return false
}

// SetPromotionName gets a reference to the given string and assigns it to the PromotionName field.
func (o *PromotionType) SetPromotionName(v string) {
	o.PromotionName = &v
}

// GetCouponCode returns the CouponCode field value if set, zero value otherwise.
func (o *PromotionType) GetCouponCode() string {
	if o == nil || IsNil(o.CouponCode) {
		var ret string
		return ret
	}
	return *o.CouponCode
}

// GetCouponCodeOk returns a tuple with the CouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionType) GetCouponCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CouponCode) {
		return nil, false
	}
	return o.CouponCode, true
}

// HasCouponCode returns a boolean if a field has been set.
func (o *PromotionType) HasCouponCode() bool {
	if o != nil && !IsNil(o.CouponCode) {
		return true
	}

	return false
}

// SetCouponCode gets a reference to the given string and assigns it to the CouponCode field.
func (o *PromotionType) SetCouponCode(v string) {
	o.CouponCode = &v
}

func (o PromotionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromotionCode) {
		toSerialize["promotionCode"] = o.PromotionCode
	}
	if !IsNil(o.PromotionName) {
		toSerialize["promotionName"] = o.PromotionName
	}
	if !IsNil(o.CouponCode) {
		toSerialize["couponCode"] = o.CouponCode
	}
	return toSerialize, nil
}

type NullablePromotionType struct {
	value *PromotionType
	isSet bool
}

func (v NullablePromotionType) Get() *PromotionType {
	return v.value
}

func (v *NullablePromotionType) Set(val *PromotionType) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionType) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionType(val *PromotionType) *NullablePromotionType {
	return &NullablePromotionType{value: val, isSet: true}
}

func (v NullablePromotionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


