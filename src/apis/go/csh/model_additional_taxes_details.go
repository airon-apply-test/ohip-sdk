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

// checks if the AdditionalTaxesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalTaxesDetails{}

// AdditionalTaxesDetails Response object containing additional tax information.
type AdditionalTaxesDetails struct {
	TaxAmountInfo *TaxAmountInfoType `json:"taxAmountInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewAdditionalTaxesDetails instantiates a new AdditionalTaxesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalTaxesDetails() *AdditionalTaxesDetails {
	this := AdditionalTaxesDetails{}
	return &this
}

// NewAdditionalTaxesDetailsWithDefaults instantiates a new AdditionalTaxesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalTaxesDetailsWithDefaults() *AdditionalTaxesDetails {
	this := AdditionalTaxesDetails{}
	return &this
}

// GetTaxAmountInfo returns the TaxAmountInfo field value if set, zero value otherwise.
func (o *AdditionalTaxesDetails) GetTaxAmountInfo() TaxAmountInfoType {
	if o == nil || IsNil(o.TaxAmountInfo) {
		var ret TaxAmountInfoType
		return ret
	}
	return *o.TaxAmountInfo
}

// GetTaxAmountInfoOk returns a tuple with the TaxAmountInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalTaxesDetails) GetTaxAmountInfoOk() (*TaxAmountInfoType, bool) {
	if o == nil || IsNil(o.TaxAmountInfo) {
		return nil, false
	}
	return o.TaxAmountInfo, true
}

// HasTaxAmountInfo returns a boolean if a field has been set.
func (o *AdditionalTaxesDetails) HasTaxAmountInfo() bool {
	if o != nil && !IsNil(o.TaxAmountInfo) {
		return true
	}

	return false
}

// SetTaxAmountInfo gets a reference to the given TaxAmountInfoType and assigns it to the TaxAmountInfo field.
func (o *AdditionalTaxesDetails) SetTaxAmountInfo(v TaxAmountInfoType) {
	o.TaxAmountInfo = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AdditionalTaxesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalTaxesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AdditionalTaxesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *AdditionalTaxesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AdditionalTaxesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalTaxesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AdditionalTaxesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *AdditionalTaxesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o AdditionalTaxesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalTaxesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxAmountInfo) {
		toSerialize["taxAmountInfo"] = o.TaxAmountInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAdditionalTaxesDetails struct {
	value *AdditionalTaxesDetails
	isSet bool
}

func (v NullableAdditionalTaxesDetails) Get() *AdditionalTaxesDetails {
	return v.value
}

func (v *NullableAdditionalTaxesDetails) Set(val *AdditionalTaxesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalTaxesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalTaxesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalTaxesDetails(val *AdditionalTaxesDetails) *NullableAdditionalTaxesDetails {
	return &NullableAdditionalTaxesDetails{value: val, isSet: true}
}

func (v NullableAdditionalTaxesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalTaxesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


