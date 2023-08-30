/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
)

// checks if the ChangedDeliveryMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangedDeliveryMethods{}

// ChangedDeliveryMethods Existing Operations Responses will eventually be modified to be extended from this type.
type ChangedDeliveryMethods struct {
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChangedDeliveryMethods instantiates a new ChangedDeliveryMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangedDeliveryMethods() *ChangedDeliveryMethods {
	this := ChangedDeliveryMethods{}
	return &this
}

// NewChangedDeliveryMethodsWithDefaults instantiates a new ChangedDeliveryMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangedDeliveryMethodsWithDefaults() *ChangedDeliveryMethods {
	this := ChangedDeliveryMethods{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangedDeliveryMethods) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedDeliveryMethods) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangedDeliveryMethods) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangedDeliveryMethods) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangedDeliveryMethods) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedDeliveryMethods) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangedDeliveryMethods) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangedDeliveryMethods) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChangedDeliveryMethods) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangedDeliveryMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChangedDeliveryMethods struct {
	value *ChangedDeliveryMethods
	isSet bool
}

func (v NullableChangedDeliveryMethods) Get() *ChangedDeliveryMethods {
	return v.value
}

func (v *NullableChangedDeliveryMethods) Set(val *ChangedDeliveryMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableChangedDeliveryMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableChangedDeliveryMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangedDeliveryMethods(val *ChangedDeliveryMethods) *NullableChangedDeliveryMethods {
	return &NullableChangedDeliveryMethods{value: val, isSet: true}
}

func (v NullableChangedDeliveryMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangedDeliveryMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


