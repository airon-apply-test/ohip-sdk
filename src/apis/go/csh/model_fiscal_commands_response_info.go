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

// checks if the FiscalCommandsResponseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiscalCommandsResponseInfo{}

// FiscalCommandsResponseInfo Response after generating a commands from Fiscal Terminals screen.
type FiscalCommandsResponseInfo struct {
	FiscalResponseInfo *FiscalResponseInfoType `json:"fiscalResponseInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFiscalCommandsResponseInfo instantiates a new FiscalCommandsResponseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiscalCommandsResponseInfo() *FiscalCommandsResponseInfo {
	this := FiscalCommandsResponseInfo{}
	return &this
}

// NewFiscalCommandsResponseInfoWithDefaults instantiates a new FiscalCommandsResponseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiscalCommandsResponseInfoWithDefaults() *FiscalCommandsResponseInfo {
	this := FiscalCommandsResponseInfo{}
	return &this
}

// GetFiscalResponseInfo returns the FiscalResponseInfo field value if set, zero value otherwise.
func (o *FiscalCommandsResponseInfo) GetFiscalResponseInfo() FiscalResponseInfoType {
	if o == nil || IsNil(o.FiscalResponseInfo) {
		var ret FiscalResponseInfoType
		return ret
	}
	return *o.FiscalResponseInfo
}

// GetFiscalResponseInfoOk returns a tuple with the FiscalResponseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalCommandsResponseInfo) GetFiscalResponseInfoOk() (*FiscalResponseInfoType, bool) {
	if o == nil || IsNil(o.FiscalResponseInfo) {
		return nil, false
	}
	return o.FiscalResponseInfo, true
}

// HasFiscalResponseInfo returns a boolean if a field has been set.
func (o *FiscalCommandsResponseInfo) HasFiscalResponseInfo() bool {
	if o != nil && !IsNil(o.FiscalResponseInfo) {
		return true
	}

	return false
}

// SetFiscalResponseInfo gets a reference to the given FiscalResponseInfoType and assigns it to the FiscalResponseInfo field.
func (o *FiscalCommandsResponseInfo) SetFiscalResponseInfo(v FiscalResponseInfoType) {
	o.FiscalResponseInfo = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FiscalCommandsResponseInfo) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalCommandsResponseInfo) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FiscalCommandsResponseInfo) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FiscalCommandsResponseInfo) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FiscalCommandsResponseInfo) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalCommandsResponseInfo) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FiscalCommandsResponseInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FiscalCommandsResponseInfo) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FiscalCommandsResponseInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiscalCommandsResponseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiscalResponseInfo) {
		toSerialize["fiscalResponseInfo"] = o.FiscalResponseInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFiscalCommandsResponseInfo struct {
	value *FiscalCommandsResponseInfo
	isSet bool
}

func (v NullableFiscalCommandsResponseInfo) Get() *FiscalCommandsResponseInfo {
	return v.value
}

func (v *NullableFiscalCommandsResponseInfo) Set(val *FiscalCommandsResponseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFiscalCommandsResponseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFiscalCommandsResponseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiscalCommandsResponseInfo(val *FiscalCommandsResponseInfo) *NullableFiscalCommandsResponseInfo {
	return &NullableFiscalCommandsResponseInfo{value: val, isSet: true}
}

func (v NullableFiscalCommandsResponseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiscalCommandsResponseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


