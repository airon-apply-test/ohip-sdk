/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
)

// checks if the PostBatchChargesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBatchChargesRequest{}

// PostBatchChargesRequest struct for PostBatchChargesRequest
type PostBatchChargesRequest struct {
	Criteria *ARPostChargesInBatchCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostBatchChargesRequest instantiates a new PostBatchChargesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBatchChargesRequest() *PostBatchChargesRequest {
	this := PostBatchChargesRequest{}
	return &this
}

// NewPostBatchChargesRequestWithDefaults instantiates a new PostBatchChargesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBatchChargesRequestWithDefaults() *PostBatchChargesRequest {
	this := PostBatchChargesRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PostBatchChargesRequest) GetCriteria() ARPostChargesInBatchCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret ARPostChargesInBatchCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBatchChargesRequest) GetCriteriaOk() (*ARPostChargesInBatchCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PostBatchChargesRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ARPostChargesInBatchCriteriaType and assigns it to the Criteria field.
func (o *PostBatchChargesRequest) SetCriteria(v ARPostChargesInBatchCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostBatchChargesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBatchChargesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostBatchChargesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostBatchChargesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostBatchChargesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBatchChargesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostBatchChargesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostBatchChargesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostBatchChargesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBatchChargesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostBatchChargesRequest struct {
	value *PostBatchChargesRequest
	isSet bool
}

func (v NullablePostBatchChargesRequest) Get() *PostBatchChargesRequest {
	return v.value
}

func (v *NullablePostBatchChargesRequest) Set(val *PostBatchChargesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBatchChargesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBatchChargesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBatchChargesRequest(val *PostBatchChargesRequest) *NullablePostBatchChargesRequest {
	return &NullablePostBatchChargesRequest{value: val, isSet: true}
}

func (v NullablePostBatchChargesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBatchChargesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


