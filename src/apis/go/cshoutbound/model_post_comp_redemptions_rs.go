/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cshoutbound

import (
	"encoding/json"
)

// checks if the PostCompRedemptionsRS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCompRedemptionsRS{}

// PostCompRedemptionsRS Response type of Complimentary Redemptions for posting.
type PostCompRedemptionsRS struct {
	// Collection of Complimentary Redemption codes and their respective Approval Code.
	CompRedemptions []PostCompRedemptionsRSCompRedemptionType `json:"compRedemptions,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostCompRedemptionsRS instantiates a new PostCompRedemptionsRS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCompRedemptionsRS() *PostCompRedemptionsRS {
	this := PostCompRedemptionsRS{}
	return &this
}

// NewPostCompRedemptionsRSWithDefaults instantiates a new PostCompRedemptionsRS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCompRedemptionsRSWithDefaults() *PostCompRedemptionsRS {
	this := PostCompRedemptionsRS{}
	return &this
}

// GetCompRedemptions returns the CompRedemptions field value if set, zero value otherwise.
func (o *PostCompRedemptionsRS) GetCompRedemptions() []PostCompRedemptionsRSCompRedemptionType {
	if o == nil || IsNil(o.CompRedemptions) {
		var ret []PostCompRedemptionsRSCompRedemptionType
		return ret
	}
	return o.CompRedemptions
}

// GetCompRedemptionsOk returns a tuple with the CompRedemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompRedemptionsRS) GetCompRedemptionsOk() ([]PostCompRedemptionsRSCompRedemptionType, bool) {
	if o == nil || IsNil(o.CompRedemptions) {
		return nil, false
	}
	return o.CompRedemptions, true
}

// HasCompRedemptions returns a boolean if a field has been set.
func (o *PostCompRedemptionsRS) HasCompRedemptions() bool {
	if o != nil && !IsNil(o.CompRedemptions) {
		return true
	}

	return false
}

// SetCompRedemptions gets a reference to the given []PostCompRedemptionsRSCompRedemptionType and assigns it to the CompRedemptions field.
func (o *PostCompRedemptionsRS) SetCompRedemptions(v []PostCompRedemptionsRSCompRedemptionType) {
	o.CompRedemptions = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostCompRedemptionsRS) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompRedemptionsRS) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostCompRedemptionsRS) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostCompRedemptionsRS) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostCompRedemptionsRS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCompRedemptionsRS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompRedemptions) {
		toSerialize["compRedemptions"] = o.CompRedemptions
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostCompRedemptionsRS struct {
	value *PostCompRedemptionsRS
	isSet bool
}

func (v NullablePostCompRedemptionsRS) Get() *PostCompRedemptionsRS {
	return v.value
}

func (v *NullablePostCompRedemptionsRS) Set(val *PostCompRedemptionsRS) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCompRedemptionsRS) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCompRedemptionsRS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCompRedemptionsRS(val *PostCompRedemptionsRS) *NullablePostCompRedemptionsRS {
	return &NullablePostCompRedemptionsRS{value: val, isSet: true}
}

func (v NullablePostCompRedemptionsRS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCompRedemptionsRS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


