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

// checks if the FolioActivityDetailsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolioActivityDetailsType{}

// FolioActivityDetailsType Fiscal Folio Activity information
type FolioActivityDetailsType struct {
	// Details of Fiscal Folio Activity made.
	FolioActivityDetailInfo []FolioActivityDetailType `json:"folioActivityDetailInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFolioActivityDetailsType instantiates a new FolioActivityDetailsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolioActivityDetailsType() *FolioActivityDetailsType {
	this := FolioActivityDetailsType{}
	return &this
}

// NewFolioActivityDetailsTypeWithDefaults instantiates a new FolioActivityDetailsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolioActivityDetailsTypeWithDefaults() *FolioActivityDetailsType {
	this := FolioActivityDetailsType{}
	return &this
}

// GetFolioActivityDetailInfo returns the FolioActivityDetailInfo field value if set, zero value otherwise.
func (o *FolioActivityDetailsType) GetFolioActivityDetailInfo() []FolioActivityDetailType {
	if o == nil || IsNil(o.FolioActivityDetailInfo) {
		var ret []FolioActivityDetailType
		return ret
	}
	return o.FolioActivityDetailInfo
}

// GetFolioActivityDetailInfoOk returns a tuple with the FolioActivityDetailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityDetailsType) GetFolioActivityDetailInfoOk() ([]FolioActivityDetailType, bool) {
	if o == nil || IsNil(o.FolioActivityDetailInfo) {
		return nil, false
	}
	return o.FolioActivityDetailInfo, true
}

// HasFolioActivityDetailInfo returns a boolean if a field has been set.
func (o *FolioActivityDetailsType) HasFolioActivityDetailInfo() bool {
	if o != nil && !IsNil(o.FolioActivityDetailInfo) {
		return true
	}

	return false
}

// SetFolioActivityDetailInfo gets a reference to the given []FolioActivityDetailType and assigns it to the FolioActivityDetailInfo field.
func (o *FolioActivityDetailsType) SetFolioActivityDetailInfo(v []FolioActivityDetailType) {
	o.FolioActivityDetailInfo = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FolioActivityDetailsType) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityDetailsType) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FolioActivityDetailsType) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FolioActivityDetailsType) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FolioActivityDetailsType) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityDetailsType) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FolioActivityDetailsType) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FolioActivityDetailsType) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FolioActivityDetailsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolioActivityDetailsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolioActivityDetailInfo) {
		toSerialize["folioActivityDetailInfo"] = o.FolioActivityDetailInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFolioActivityDetailsType struct {
	value *FolioActivityDetailsType
	isSet bool
}

func (v NullableFolioActivityDetailsType) Get() *FolioActivityDetailsType {
	return v.value
}

func (v *NullableFolioActivityDetailsType) Set(val *FolioActivityDetailsType) {
	v.value = val
	v.isSet = true
}

func (v NullableFolioActivityDetailsType) IsSet() bool {
	return v.isSet
}

func (v *NullableFolioActivityDetailsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolioActivityDetailsType(val *FolioActivityDetailsType) *NullableFolioActivityDetailsType {
	return &NullableFolioActivityDetailsType{value: val, isSet: true}
}

func (v NullableFolioActivityDetailsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolioActivityDetailsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


