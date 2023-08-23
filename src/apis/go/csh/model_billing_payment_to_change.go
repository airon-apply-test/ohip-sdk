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

// checks if the BillingPaymentToChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPaymentToChange{}

// BillingPaymentToChange Request to change a payment posting.
type BillingPaymentToChange struct {
	Criteria *ChangePaymentCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBillingPaymentToChange instantiates a new BillingPaymentToChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPaymentToChange() *BillingPaymentToChange {
	this := BillingPaymentToChange{}
	return &this
}

// NewBillingPaymentToChangeWithDefaults instantiates a new BillingPaymentToChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPaymentToChangeWithDefaults() *BillingPaymentToChange {
	this := BillingPaymentToChange{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *BillingPaymentToChange) GetCriteria() ChangePaymentCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret ChangePaymentCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentToChange) GetCriteriaOk() (*ChangePaymentCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *BillingPaymentToChange) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ChangePaymentCriteriaType and assigns it to the Criteria field.
func (o *BillingPaymentToChange) SetCriteria(v ChangePaymentCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BillingPaymentToChange) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentToChange) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BillingPaymentToChange) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BillingPaymentToChange) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BillingPaymentToChange) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentToChange) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BillingPaymentToChange) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BillingPaymentToChange) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BillingPaymentToChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPaymentToChange) ToMap() (map[string]interface{}, error) {
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

type NullableBillingPaymentToChange struct {
	value *BillingPaymentToChange
	isSet bool
}

func (v NullableBillingPaymentToChange) Get() *BillingPaymentToChange {
	return v.value
}

func (v *NullableBillingPaymentToChange) Set(val *BillingPaymentToChange) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPaymentToChange) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPaymentToChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPaymentToChange(val *BillingPaymentToChange) *NullableBillingPaymentToChange {
	return &NullableBillingPaymentToChange{value: val, isSet: true}
}

func (v NullableBillingPaymentToChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPaymentToChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


