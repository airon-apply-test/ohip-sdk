/*
OPERA Cloud Front Desk Configuration API

APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fofcfg

import (
	"encoding/json"
)

// checks if the CustomTaxTypesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTaxTypesDetails{}

// CustomTaxTypesDetails Response object for fetching Custom Tax Types.
type CustomTaxTypesDetails struct {
	// List of Custom Tax Types.
	CustomTaxTypes []CustomTaxTypeType `json:"customTaxTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCustomTaxTypesDetails instantiates a new CustomTaxTypesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTaxTypesDetails() *CustomTaxTypesDetails {
	this := CustomTaxTypesDetails{}
	return &this
}

// NewCustomTaxTypesDetailsWithDefaults instantiates a new CustomTaxTypesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTaxTypesDetailsWithDefaults() *CustomTaxTypesDetails {
	this := CustomTaxTypesDetails{}
	return &this
}

// GetCustomTaxTypes returns the CustomTaxTypes field value if set, zero value otherwise.
func (o *CustomTaxTypesDetails) GetCustomTaxTypes() []CustomTaxTypeType {
	if o == nil || IsNil(o.CustomTaxTypes) {
		var ret []CustomTaxTypeType
		return ret
	}
	return o.CustomTaxTypes
}

// GetCustomTaxTypesOk returns a tuple with the CustomTaxTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTaxTypesDetails) GetCustomTaxTypesOk() ([]CustomTaxTypeType, bool) {
	if o == nil || IsNil(o.CustomTaxTypes) {
		return nil, false
	}
	return o.CustomTaxTypes, true
}

// HasCustomTaxTypes returns a boolean if a field has been set.
func (o *CustomTaxTypesDetails) HasCustomTaxTypes() bool {
	if o != nil && !IsNil(o.CustomTaxTypes) {
		return true
	}

	return false
}

// SetCustomTaxTypes gets a reference to the given []CustomTaxTypeType and assigns it to the CustomTaxTypes field.
func (o *CustomTaxTypesDetails) SetCustomTaxTypes(v []CustomTaxTypeType) {
	o.CustomTaxTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomTaxTypesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTaxTypesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomTaxTypesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CustomTaxTypesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CustomTaxTypesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTaxTypesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CustomTaxTypesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CustomTaxTypesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CustomTaxTypesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTaxTypesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomTaxTypes) {
		toSerialize["customTaxTypes"] = o.CustomTaxTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCustomTaxTypesDetails struct {
	value *CustomTaxTypesDetails
	isSet bool
}

func (v NullableCustomTaxTypesDetails) Get() *CustomTaxTypesDetails {
	return v.value
}

func (v *NullableCustomTaxTypesDetails) Set(val *CustomTaxTypesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTaxTypesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTaxTypesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTaxTypesDetails(val *CustomTaxTypesDetails) *NullableCustomTaxTypesDetails {
	return &NullableCustomTaxTypesDetails{value: val, isSet: true}
}

func (v NullableCustomTaxTypesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTaxTypesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


