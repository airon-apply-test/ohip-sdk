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

// checks if the GamingRequestStatusInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GamingRequestStatusInfo{}

// GamingRequestStatusInfo Response object for fetch gaming info
type GamingRequestStatusInfo struct {
	// List of guest request status details.
	StatusInfo []GamingRequestStatusType `json:"statusInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewGamingRequestStatusInfo instantiates a new GamingRequestStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGamingRequestStatusInfo() *GamingRequestStatusInfo {
	this := GamingRequestStatusInfo{}
	return &this
}

// NewGamingRequestStatusInfoWithDefaults instantiates a new GamingRequestStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGamingRequestStatusInfoWithDefaults() *GamingRequestStatusInfo {
	this := GamingRequestStatusInfo{}
	return &this
}

// GetStatusInfo returns the StatusInfo field value if set, zero value otherwise.
func (o *GamingRequestStatusInfo) GetStatusInfo() []GamingRequestStatusType {
	if o == nil || IsNil(o.StatusInfo) {
		var ret []GamingRequestStatusType
		return ret
	}
	return o.StatusInfo
}

// GetStatusInfoOk returns a tuple with the StatusInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GamingRequestStatusInfo) GetStatusInfoOk() ([]GamingRequestStatusType, bool) {
	if o == nil || IsNil(o.StatusInfo) {
		return nil, false
	}
	return o.StatusInfo, true
}

// HasStatusInfo returns a boolean if a field has been set.
func (o *GamingRequestStatusInfo) HasStatusInfo() bool {
	if o != nil && !IsNil(o.StatusInfo) {
		return true
	}

	return false
}

// SetStatusInfo gets a reference to the given []GamingRequestStatusType and assigns it to the StatusInfo field.
func (o *GamingRequestStatusInfo) SetStatusInfo(v []GamingRequestStatusType) {
	o.StatusInfo = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GamingRequestStatusInfo) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GamingRequestStatusInfo) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GamingRequestStatusInfo) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *GamingRequestStatusInfo) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *GamingRequestStatusInfo) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GamingRequestStatusInfo) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *GamingRequestStatusInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *GamingRequestStatusInfo) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o GamingRequestStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GamingRequestStatusInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusInfo) {
		toSerialize["statusInfo"] = o.StatusInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableGamingRequestStatusInfo struct {
	value *GamingRequestStatusInfo
	isSet bool
}

func (v NullableGamingRequestStatusInfo) Get() *GamingRequestStatusInfo {
	return v.value
}

func (v *NullableGamingRequestStatusInfo) Set(val *GamingRequestStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGamingRequestStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGamingRequestStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGamingRequestStatusInfo(val *GamingRequestStatusInfo) *NullableGamingRequestStatusInfo {
	return &NullableGamingRequestStatusInfo{value: val, isSet: true}
}

func (v NullableGamingRequestStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGamingRequestStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


