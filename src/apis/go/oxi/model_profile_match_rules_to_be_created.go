/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the ProfileMatchRulesToBeCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileMatchRulesToBeCreated{}

// ProfileMatchRulesToBeCreated struct for ProfileMatchRulesToBeCreated
type ProfileMatchRulesToBeCreated struct {
	// Type that holds collection of exchange profile match rules.
	ExchangeProfileMatchRules []ExchangeProfileMatchRuleType `json:"exchangeProfileMatchRules,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewProfileMatchRulesToBeCreated instantiates a new ProfileMatchRulesToBeCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMatchRulesToBeCreated() *ProfileMatchRulesToBeCreated {
	this := ProfileMatchRulesToBeCreated{}
	return &this
}

// NewProfileMatchRulesToBeCreatedWithDefaults instantiates a new ProfileMatchRulesToBeCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMatchRulesToBeCreatedWithDefaults() *ProfileMatchRulesToBeCreated {
	this := ProfileMatchRulesToBeCreated{}
	return &this
}

// GetExchangeProfileMatchRules returns the ExchangeProfileMatchRules field value if set, zero value otherwise.
func (o *ProfileMatchRulesToBeCreated) GetExchangeProfileMatchRules() []ExchangeProfileMatchRuleType {
	if o == nil || IsNil(o.ExchangeProfileMatchRules) {
		var ret []ExchangeProfileMatchRuleType
		return ret
	}
	return o.ExchangeProfileMatchRules
}

// GetExchangeProfileMatchRulesOk returns a tuple with the ExchangeProfileMatchRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMatchRulesToBeCreated) GetExchangeProfileMatchRulesOk() ([]ExchangeProfileMatchRuleType, bool) {
	if o == nil || IsNil(o.ExchangeProfileMatchRules) {
		return nil, false
	}
	return o.ExchangeProfileMatchRules, true
}

// HasExchangeProfileMatchRules returns a boolean if a field has been set.
func (o *ProfileMatchRulesToBeCreated) HasExchangeProfileMatchRules() bool {
	if o != nil && !IsNil(o.ExchangeProfileMatchRules) {
		return true
	}

	return false
}

// SetExchangeProfileMatchRules gets a reference to the given []ExchangeProfileMatchRuleType and assigns it to the ExchangeProfileMatchRules field.
func (o *ProfileMatchRulesToBeCreated) SetExchangeProfileMatchRules(v []ExchangeProfileMatchRuleType) {
	o.ExchangeProfileMatchRules = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileMatchRulesToBeCreated) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMatchRulesToBeCreated) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileMatchRulesToBeCreated) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ProfileMatchRulesToBeCreated) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ProfileMatchRulesToBeCreated) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMatchRulesToBeCreated) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ProfileMatchRulesToBeCreated) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ProfileMatchRulesToBeCreated) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ProfileMatchRulesToBeCreated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileMatchRulesToBeCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExchangeProfileMatchRules) {
		toSerialize["exchangeProfileMatchRules"] = o.ExchangeProfileMatchRules
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableProfileMatchRulesToBeCreated struct {
	value *ProfileMatchRulesToBeCreated
	isSet bool
}

func (v NullableProfileMatchRulesToBeCreated) Get() *ProfileMatchRulesToBeCreated {
	return v.value
}

func (v *NullableProfileMatchRulesToBeCreated) Set(val *ProfileMatchRulesToBeCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMatchRulesToBeCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMatchRulesToBeCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMatchRulesToBeCreated(val *ProfileMatchRulesToBeCreated) *NullableProfileMatchRulesToBeCreated {
	return &NullableProfileMatchRulesToBeCreated{value: val, isSet: true}
}

func (v NullableProfileMatchRulesToBeCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMatchRulesToBeCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


