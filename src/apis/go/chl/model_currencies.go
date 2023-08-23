/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Currencies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Currencies{}

// Currencies Request object for changing existing external system currencies.
type Currencies struct {
	// Information about an external currency mapping.
	Currencies []CurrencyMappingType `json:"currencies,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCurrencies instantiates a new Currencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencies() *Currencies {
	this := Currencies{}
	return &this
}

// NewCurrenciesWithDefaults instantiates a new Currencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrenciesWithDefaults() *Currencies {
	this := Currencies{}
	return &this
}

// GetCurrencies returns the Currencies field value if set, zero value otherwise.
func (o *Currencies) GetCurrencies() []CurrencyMappingType {
	if o == nil || IsNil(o.Currencies) {
		var ret []CurrencyMappingType
		return ret
	}
	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currencies) GetCurrenciesOk() ([]CurrencyMappingType, bool) {
	if o == nil || IsNil(o.Currencies) {
		return nil, false
	}
	return o.Currencies, true
}

// HasCurrencies returns a boolean if a field has been set.
func (o *Currencies) HasCurrencies() bool {
	if o != nil && !IsNil(o.Currencies) {
		return true
	}

	return false
}

// SetCurrencies gets a reference to the given []CurrencyMappingType and assigns it to the Currencies field.
func (o *Currencies) SetCurrencies(v []CurrencyMappingType) {
	o.Currencies = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Currencies) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currencies) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Currencies) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *Currencies) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Currencies) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currencies) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Currencies) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *Currencies) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o Currencies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Currencies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currencies) {
		toSerialize["currencies"] = o.Currencies
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCurrencies struct {
	value *Currencies
	isSet bool
}

func (v NullableCurrencies) Get() *Currencies {
	return v.value
}

func (v *NullableCurrencies) Set(val *Currencies) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencies) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencies(val *Currencies) *NullableCurrencies {
	return &NullableCurrencies{value: val, isSet: true}
}

func (v NullableCurrencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


