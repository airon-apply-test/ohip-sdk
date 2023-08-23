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

// checks if the PostBatchAutoSettlementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBatchAutoSettlementRequest{}

// PostBatchAutoSettlementRequest struct for PostBatchAutoSettlementRequest
type PostBatchAutoSettlementRequest struct {
	Criteria *AutoSettlementType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostBatchAutoSettlementRequest instantiates a new PostBatchAutoSettlementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBatchAutoSettlementRequest() *PostBatchAutoSettlementRequest {
	this := PostBatchAutoSettlementRequest{}
	return &this
}

// NewPostBatchAutoSettlementRequestWithDefaults instantiates a new PostBatchAutoSettlementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBatchAutoSettlementRequestWithDefaults() *PostBatchAutoSettlementRequest {
	this := PostBatchAutoSettlementRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PostBatchAutoSettlementRequest) GetCriteria() AutoSettlementType {
	if o == nil || IsNil(o.Criteria) {
		var ret AutoSettlementType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBatchAutoSettlementRequest) GetCriteriaOk() (*AutoSettlementType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PostBatchAutoSettlementRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AutoSettlementType and assigns it to the Criteria field.
func (o *PostBatchAutoSettlementRequest) SetCriteria(v AutoSettlementType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostBatchAutoSettlementRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBatchAutoSettlementRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostBatchAutoSettlementRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostBatchAutoSettlementRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostBatchAutoSettlementRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBatchAutoSettlementRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostBatchAutoSettlementRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostBatchAutoSettlementRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostBatchAutoSettlementRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBatchAutoSettlementRequest) ToMap() (map[string]interface{}, error) {
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

type NullablePostBatchAutoSettlementRequest struct {
	value *PostBatchAutoSettlementRequest
	isSet bool
}

func (v NullablePostBatchAutoSettlementRequest) Get() *PostBatchAutoSettlementRequest {
	return v.value
}

func (v *NullablePostBatchAutoSettlementRequest) Set(val *PostBatchAutoSettlementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBatchAutoSettlementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBatchAutoSettlementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBatchAutoSettlementRequest(val *PostBatchAutoSettlementRequest) *NullablePostBatchAutoSettlementRequest {
	return &NullablePostBatchAutoSettlementRequest{value: val, isSet: true}
}

func (v NullablePostBatchAutoSettlementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBatchAutoSettlementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


