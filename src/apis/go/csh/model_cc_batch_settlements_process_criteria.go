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

// checks if the CcBatchSettlementsProcessCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CcBatchSettlementsProcessCriteria{}

// CcBatchSettlementsProcessCriteria Request to process batch settlements
type CcBatchSettlementsProcessCriteria struct {
	Criteria *CCBatchSettlementsProcessType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCcBatchSettlementsProcessCriteria instantiates a new CcBatchSettlementsProcessCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCcBatchSettlementsProcessCriteria() *CcBatchSettlementsProcessCriteria {
	this := CcBatchSettlementsProcessCriteria{}
	return &this
}

// NewCcBatchSettlementsProcessCriteriaWithDefaults instantiates a new CcBatchSettlementsProcessCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCcBatchSettlementsProcessCriteriaWithDefaults() *CcBatchSettlementsProcessCriteria {
	this := CcBatchSettlementsProcessCriteria{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *CcBatchSettlementsProcessCriteria) GetCriteria() CCBatchSettlementsProcessType {
	if o == nil || IsNil(o.Criteria) {
		var ret CCBatchSettlementsProcessType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcBatchSettlementsProcessCriteria) GetCriteriaOk() (*CCBatchSettlementsProcessType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *CcBatchSettlementsProcessCriteria) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given CCBatchSettlementsProcessType and assigns it to the Criteria field.
func (o *CcBatchSettlementsProcessCriteria) SetCriteria(v CCBatchSettlementsProcessType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CcBatchSettlementsProcessCriteria) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcBatchSettlementsProcessCriteria) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CcBatchSettlementsProcessCriteria) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CcBatchSettlementsProcessCriteria) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CcBatchSettlementsProcessCriteria) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CcBatchSettlementsProcessCriteria) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CcBatchSettlementsProcessCriteria) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CcBatchSettlementsProcessCriteria) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CcBatchSettlementsProcessCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CcBatchSettlementsProcessCriteria) ToMap() (map[string]interface{}, error) {
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

type NullableCcBatchSettlementsProcessCriteria struct {
	value *CcBatchSettlementsProcessCriteria
	isSet bool
}

func (v NullableCcBatchSettlementsProcessCriteria) Get() *CcBatchSettlementsProcessCriteria {
	return v.value
}

func (v *NullableCcBatchSettlementsProcessCriteria) Set(val *CcBatchSettlementsProcessCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCcBatchSettlementsProcessCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCcBatchSettlementsProcessCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCcBatchSettlementsProcessCriteria(val *CcBatchSettlementsProcessCriteria) *NullableCcBatchSettlementsProcessCriteria {
	return &NullableCcBatchSettlementsProcessCriteria{value: val, isSet: true}
}

func (v NullableCcBatchSettlementsProcessCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCcBatchSettlementsProcessCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


