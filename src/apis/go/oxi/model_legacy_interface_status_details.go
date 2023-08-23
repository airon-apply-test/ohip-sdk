/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LegacyInterfaceStatusDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegacyInterfaceStatusDetails{}

// LegacyInterfaceStatusDetails Response object for fetch legacy interface status.
type LegacyInterfaceStatusDetails struct {
	// Collection which contains Interface ID and its details.
	InterfacesStatus []InterfaceStatusType `json:"interfacesStatus,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewLegacyInterfaceStatusDetails instantiates a new LegacyInterfaceStatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyInterfaceStatusDetails() *LegacyInterfaceStatusDetails {
	this := LegacyInterfaceStatusDetails{}
	return &this
}

// NewLegacyInterfaceStatusDetailsWithDefaults instantiates a new LegacyInterfaceStatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyInterfaceStatusDetailsWithDefaults() *LegacyInterfaceStatusDetails {
	this := LegacyInterfaceStatusDetails{}
	return &this
}

// GetInterfacesStatus returns the InterfacesStatus field value if set, zero value otherwise.
func (o *LegacyInterfaceStatusDetails) GetInterfacesStatus() []InterfaceStatusType {
	if o == nil || IsNil(o.InterfacesStatus) {
		var ret []InterfaceStatusType
		return ret
	}
	return o.InterfacesStatus
}

// GetInterfacesStatusOk returns a tuple with the InterfacesStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyInterfaceStatusDetails) GetInterfacesStatusOk() ([]InterfaceStatusType, bool) {
	if o == nil || IsNil(o.InterfacesStatus) {
		return nil, false
	}
	return o.InterfacesStatus, true
}

// HasInterfacesStatus returns a boolean if a field has been set.
func (o *LegacyInterfaceStatusDetails) HasInterfacesStatus() bool {
	if o != nil && !IsNil(o.InterfacesStatus) {
		return true
	}

	return false
}

// SetInterfacesStatus gets a reference to the given []InterfaceStatusType and assigns it to the InterfacesStatus field.
func (o *LegacyInterfaceStatusDetails) SetInterfacesStatus(v []InterfaceStatusType) {
	o.InterfacesStatus = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LegacyInterfaceStatusDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyInterfaceStatusDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LegacyInterfaceStatusDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *LegacyInterfaceStatusDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *LegacyInterfaceStatusDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyInterfaceStatusDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *LegacyInterfaceStatusDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *LegacyInterfaceStatusDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o LegacyInterfaceStatusDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegacyInterfaceStatusDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterfacesStatus) {
		toSerialize["interfacesStatus"] = o.InterfacesStatus
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableLegacyInterfaceStatusDetails struct {
	value *LegacyInterfaceStatusDetails
	isSet bool
}

func (v NullableLegacyInterfaceStatusDetails) Get() *LegacyInterfaceStatusDetails {
	return v.value
}

func (v *NullableLegacyInterfaceStatusDetails) Set(val *LegacyInterfaceStatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyInterfaceStatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyInterfaceStatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyInterfaceStatusDetails(val *LegacyInterfaceStatusDetails) *NullableLegacyInterfaceStatusDetails {
	return &NullableLegacyInterfaceStatusDetails{value: val, isSet: true}
}

func (v NullableLegacyInterfaceStatusDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyInterfaceStatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


