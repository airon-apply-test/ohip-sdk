/*
OPERA Cloud Content Service

Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package medcfg

import (
	"encoding/json"
)

// checks if the CustomizedLetter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomizedLetter{}

// CustomizedLetter Response object for retrieving customized letter.
type CustomizedLetter struct {
	CustomizedLetterDetails *CustomizedLetterType `json:"customizedLetterDetails,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCustomizedLetter instantiates a new CustomizedLetter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomizedLetter() *CustomizedLetter {
	this := CustomizedLetter{}
	return &this
}

// NewCustomizedLetterWithDefaults instantiates a new CustomizedLetter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomizedLetterWithDefaults() *CustomizedLetter {
	this := CustomizedLetter{}
	return &this
}

// GetCustomizedLetterDetails returns the CustomizedLetterDetails field value if set, zero value otherwise.
func (o *CustomizedLetter) GetCustomizedLetterDetails() CustomizedLetterType {
	if o == nil || IsNil(o.CustomizedLetterDetails) {
		var ret CustomizedLetterType
		return ret
	}
	return *o.CustomizedLetterDetails
}

// GetCustomizedLetterDetailsOk returns a tuple with the CustomizedLetterDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizedLetter) GetCustomizedLetterDetailsOk() (*CustomizedLetterType, bool) {
	if o == nil || IsNil(o.CustomizedLetterDetails) {
		return nil, false
	}
	return o.CustomizedLetterDetails, true
}

// HasCustomizedLetterDetails returns a boolean if a field has been set.
func (o *CustomizedLetter) HasCustomizedLetterDetails() bool {
	if o != nil && !IsNil(o.CustomizedLetterDetails) {
		return true
	}

	return false
}

// SetCustomizedLetterDetails gets a reference to the given CustomizedLetterType and assigns it to the CustomizedLetterDetails field.
func (o *CustomizedLetter) SetCustomizedLetterDetails(v CustomizedLetterType) {
	o.CustomizedLetterDetails = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomizedLetter) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizedLetter) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomizedLetter) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CustomizedLetter) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CustomizedLetter) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizedLetter) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CustomizedLetter) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CustomizedLetter) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CustomizedLetter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomizedLetter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomizedLetterDetails) {
		toSerialize["customizedLetterDetails"] = o.CustomizedLetterDetails
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCustomizedLetter struct {
	value *CustomizedLetter
	isSet bool
}

func (v NullableCustomizedLetter) Get() *CustomizedLetter {
	return v.value
}

func (v *NullableCustomizedLetter) Set(val *CustomizedLetter) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomizedLetter) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomizedLetter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomizedLetter(val *CustomizedLetter) *NullableCustomizedLetter {
	return &NullableCustomizedLetter{value: val, isSet: true}
}

func (v NullableCustomizedLetter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomizedLetter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


