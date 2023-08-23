/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateBestAvailableRates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBestAvailableRates{}

// CreateBestAvailableRates Request to create best available rates for rate code(s)
type CreateBestAvailableRates struct {
	BestAvailableRates *BestAvailableRatesType `json:"bestAvailableRates,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreateBestAvailableRates instantiates a new CreateBestAvailableRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBestAvailableRates() *CreateBestAvailableRates {
	this := CreateBestAvailableRates{}
	return &this
}

// NewCreateBestAvailableRatesWithDefaults instantiates a new CreateBestAvailableRates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBestAvailableRatesWithDefaults() *CreateBestAvailableRates {
	this := CreateBestAvailableRates{}
	return &this
}

// GetBestAvailableRates returns the BestAvailableRates field value if set, zero value otherwise.
func (o *CreateBestAvailableRates) GetBestAvailableRates() BestAvailableRatesType {
	if o == nil || IsNil(o.BestAvailableRates) {
		var ret BestAvailableRatesType
		return ret
	}
	return *o.BestAvailableRates
}

// GetBestAvailableRatesOk returns a tuple with the BestAvailableRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBestAvailableRates) GetBestAvailableRatesOk() (*BestAvailableRatesType, bool) {
	if o == nil || IsNil(o.BestAvailableRates) {
		return nil, false
	}
	return o.BestAvailableRates, true
}

// HasBestAvailableRates returns a boolean if a field has been set.
func (o *CreateBestAvailableRates) HasBestAvailableRates() bool {
	if o != nil && !IsNil(o.BestAvailableRates) {
		return true
	}

	return false
}

// SetBestAvailableRates gets a reference to the given BestAvailableRatesType and assigns it to the BestAvailableRates field.
func (o *CreateBestAvailableRates) SetBestAvailableRates(v BestAvailableRatesType) {
	o.BestAvailableRates = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateBestAvailableRates) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBestAvailableRates) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateBestAvailableRates) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CreateBestAvailableRates) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreateBestAvailableRates) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBestAvailableRates) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreateBestAvailableRates) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreateBestAvailableRates) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreateBestAvailableRates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBestAvailableRates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BestAvailableRates) {
		toSerialize["bestAvailableRates"] = o.BestAvailableRates
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCreateBestAvailableRates struct {
	value *CreateBestAvailableRates
	isSet bool
}

func (v NullableCreateBestAvailableRates) Get() *CreateBestAvailableRates {
	return v.value
}

func (v *NullableCreateBestAvailableRates) Set(val *CreateBestAvailableRates) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBestAvailableRates) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBestAvailableRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBestAvailableRates(val *CreateBestAvailableRates) *NullableCreateBestAvailableRates {
	return &NullableCreateBestAvailableRates{value: val, isSet: true}
}

func (v NullableCreateBestAvailableRates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBestAvailableRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


