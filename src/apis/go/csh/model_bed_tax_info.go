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

// checks if the BedTaxInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BedTaxInfo{}

// BedTaxInfo Response for bed transaction info
type BedTaxInfo struct {
	// List of Bed Tax info.
	BedTaxInfoTypes []BedTaxInfoType `json:"bedTaxInfoTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBedTaxInfo instantiates a new BedTaxInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBedTaxInfo() *BedTaxInfo {
	this := BedTaxInfo{}
	return &this
}

// NewBedTaxInfoWithDefaults instantiates a new BedTaxInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBedTaxInfoWithDefaults() *BedTaxInfo {
	this := BedTaxInfo{}
	return &this
}

// GetBedTaxInfoTypes returns the BedTaxInfoTypes field value if set, zero value otherwise.
func (o *BedTaxInfo) GetBedTaxInfoTypes() []BedTaxInfoType {
	if o == nil || IsNil(o.BedTaxInfoTypes) {
		var ret []BedTaxInfoType
		return ret
	}
	return o.BedTaxInfoTypes
}

// GetBedTaxInfoTypesOk returns a tuple with the BedTaxInfoTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfo) GetBedTaxInfoTypesOk() ([]BedTaxInfoType, bool) {
	if o == nil || IsNil(o.BedTaxInfoTypes) {
		return nil, false
	}
	return o.BedTaxInfoTypes, true
}

// HasBedTaxInfoTypes returns a boolean if a field has been set.
func (o *BedTaxInfo) HasBedTaxInfoTypes() bool {
	if o != nil && !IsNil(o.BedTaxInfoTypes) {
		return true
	}

	return false
}

// SetBedTaxInfoTypes gets a reference to the given []BedTaxInfoType and assigns it to the BedTaxInfoTypes field.
func (o *BedTaxInfo) SetBedTaxInfoTypes(v []BedTaxInfoType) {
	o.BedTaxInfoTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BedTaxInfo) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfo) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BedTaxInfo) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BedTaxInfo) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BedTaxInfo) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfo) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BedTaxInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BedTaxInfo) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BedTaxInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BedTaxInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BedTaxInfoTypes) {
		toSerialize["bedTaxInfoTypes"] = o.BedTaxInfoTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBedTaxInfo struct {
	value *BedTaxInfo
	isSet bool
}

func (v NullableBedTaxInfo) Get() *BedTaxInfo {
	return v.value
}

func (v *NullableBedTaxInfo) Set(val *BedTaxInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBedTaxInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBedTaxInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBedTaxInfo(val *BedTaxInfo) *NullableBedTaxInfo {
	return &NullableBedTaxInfo{value: val, isSet: true}
}

func (v NullableBedTaxInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBedTaxInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


