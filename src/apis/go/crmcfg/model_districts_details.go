/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DistrictsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistrictsDetails{}

// DistrictsDetails Response object for fetching Districts.
type DistrictsDetails struct {
	// List of Districts.
	Districts []DistrictType `json:"districts,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewDistrictsDetails instantiates a new DistrictsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistrictsDetails() *DistrictsDetails {
	this := DistrictsDetails{}
	return &this
}

// NewDistrictsDetailsWithDefaults instantiates a new DistrictsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistrictsDetailsWithDefaults() *DistrictsDetails {
	this := DistrictsDetails{}
	return &this
}

// GetDistricts returns the Districts field value if set, zero value otherwise.
func (o *DistrictsDetails) GetDistricts() []DistrictType {
	if o == nil || IsNil(o.Districts) {
		var ret []DistrictType
		return ret
	}
	return o.Districts
}

// GetDistrictsOk returns a tuple with the Districts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistrictsDetails) GetDistrictsOk() ([]DistrictType, bool) {
	if o == nil || IsNil(o.Districts) {
		return nil, false
	}
	return o.Districts, true
}

// HasDistricts returns a boolean if a field has been set.
func (o *DistrictsDetails) HasDistricts() bool {
	if o != nil && !IsNil(o.Districts) {
		return true
	}

	return false
}

// SetDistricts gets a reference to the given []DistrictType and assigns it to the Districts field.
func (o *DistrictsDetails) SetDistricts(v []DistrictType) {
	o.Districts = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DistrictsDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistrictsDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DistrictsDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *DistrictsDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DistrictsDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistrictsDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DistrictsDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *DistrictsDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o DistrictsDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistrictsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Districts) {
		toSerialize["districts"] = o.Districts
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableDistrictsDetails struct {
	value *DistrictsDetails
	isSet bool
}

func (v NullableDistrictsDetails) Get() *DistrictsDetails {
	return v.value
}

func (v *NullableDistrictsDetails) Set(val *DistrictsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDistrictsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDistrictsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistrictsDetails(val *DistrictsDetails) *NullableDistrictsDetails {
	return &NullableDistrictsDetails{value: val, isSet: true}
}

func (v NullableDistrictsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistrictsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


