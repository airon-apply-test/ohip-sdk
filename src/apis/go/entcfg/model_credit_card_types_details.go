/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
)

// checks if the CreditCardTypesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditCardTypesDetails{}

// CreditCardTypesDetails Response object for fetching Credit Card Types.
type CreditCardTypesDetails struct {
	// List of Credit Card Types.
	CreditCardTypes []CreditCardTypeType `json:"creditCardTypes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreditCardTypesDetails instantiates a new CreditCardTypesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardTypesDetails() *CreditCardTypesDetails {
	this := CreditCardTypesDetails{}
	return &this
}

// NewCreditCardTypesDetailsWithDefaults instantiates a new CreditCardTypesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardTypesDetailsWithDefaults() *CreditCardTypesDetails {
	this := CreditCardTypesDetails{}
	return &this
}

// GetCreditCardTypes returns the CreditCardTypes field value if set, zero value otherwise.
func (o *CreditCardTypesDetails) GetCreditCardTypes() []CreditCardTypeType {
	if o == nil || IsNil(o.CreditCardTypes) {
		var ret []CreditCardTypeType
		return ret
	}
	return o.CreditCardTypes
}

// GetCreditCardTypesOk returns a tuple with the CreditCardTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardTypesDetails) GetCreditCardTypesOk() ([]CreditCardTypeType, bool) {
	if o == nil || IsNil(o.CreditCardTypes) {
		return nil, false
	}
	return o.CreditCardTypes, true
}

// HasCreditCardTypes returns a boolean if a field has been set.
func (o *CreditCardTypesDetails) HasCreditCardTypes() bool {
	if o != nil && !IsNil(o.CreditCardTypes) {
		return true
	}

	return false
}

// SetCreditCardTypes gets a reference to the given []CreditCardTypeType and assigns it to the CreditCardTypes field.
func (o *CreditCardTypesDetails) SetCreditCardTypes(v []CreditCardTypeType) {
	o.CreditCardTypes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreditCardTypesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardTypesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreditCardTypesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CreditCardTypesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreditCardTypesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardTypesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreditCardTypesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreditCardTypesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreditCardTypesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditCardTypesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditCardTypes) {
		toSerialize["creditCardTypes"] = o.CreditCardTypes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCreditCardTypesDetails struct {
	value *CreditCardTypesDetails
	isSet bool
}

func (v NullableCreditCardTypesDetails) Get() *CreditCardTypesDetails {
	return v.value
}

func (v *NullableCreditCardTypesDetails) Set(val *CreditCardTypesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardTypesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardTypesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardTypesDetails(val *CreditCardTypesDetails) *NullableCreditCardTypesDetails {
	return &NullableCreditCardTypesDetails{value: val, isSet: true}
}

func (v NullableCreditCardTypesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardTypesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


