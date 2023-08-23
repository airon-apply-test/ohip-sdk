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

// checks if the PostCompRedemptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCompRedemptionsRequest{}

// PostCompRedemptionsRequest struct for PostCompRedemptionsRequest
type PostCompRedemptionsRequest struct {
	Criteria *PostCompRedemptionsCriteria `json:"criteria,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostCompRedemptionsRequest instantiates a new PostCompRedemptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCompRedemptionsRequest() *PostCompRedemptionsRequest {
	this := PostCompRedemptionsRequest{}
	return &this
}

// NewPostCompRedemptionsRequestWithDefaults instantiates a new PostCompRedemptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCompRedemptionsRequestWithDefaults() *PostCompRedemptionsRequest {
	this := PostCompRedemptionsRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PostCompRedemptionsRequest) GetCriteria() PostCompRedemptionsCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret PostCompRedemptionsCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompRedemptionsRequest) GetCriteriaOk() (*PostCompRedemptionsCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PostCompRedemptionsRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given PostCompRedemptionsCriteria and assigns it to the Criteria field.
func (o *PostCompRedemptionsRequest) SetCriteria(v PostCompRedemptionsCriteria) {
	o.Criteria = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostCompRedemptionsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCompRedemptionsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostCompRedemptionsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostCompRedemptionsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostCompRedemptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCompRedemptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostCompRedemptionsRequest struct {
	value *PostCompRedemptionsRequest
	isSet bool
}

func (v NullablePostCompRedemptionsRequest) Get() *PostCompRedemptionsRequest {
	return v.value
}

func (v *NullablePostCompRedemptionsRequest) Set(val *PostCompRedemptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCompRedemptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCompRedemptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCompRedemptionsRequest(val *PostCompRedemptionsRequest) *NullablePostCompRedemptionsRequest {
	return &NullablePostCompRedemptionsRequest{value: val, isSet: true}
}

func (v NullablePostCompRedemptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCompRedemptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


