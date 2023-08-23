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

// checks if the CreditCards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditCards{}

// CreditCards Request object for changing existing external system credit cards.
type CreditCards struct {
	// Information about an external system credit card mapping.
	CreditCards []CreditCardMappingType `json:"creditCards,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreditCards instantiates a new CreditCards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCards() *CreditCards {
	this := CreditCards{}
	return &this
}

// NewCreditCardsWithDefaults instantiates a new CreditCards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardsWithDefaults() *CreditCards {
	this := CreditCards{}
	return &this
}

// GetCreditCards returns the CreditCards field value if set, zero value otherwise.
func (o *CreditCards) GetCreditCards() []CreditCardMappingType {
	if o == nil || IsNil(o.CreditCards) {
		var ret []CreditCardMappingType
		return ret
	}
	return o.CreditCards
}

// GetCreditCardsOk returns a tuple with the CreditCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCards) GetCreditCardsOk() ([]CreditCardMappingType, bool) {
	if o == nil || IsNil(o.CreditCards) {
		return nil, false
	}
	return o.CreditCards, true
}

// HasCreditCards returns a boolean if a field has been set.
func (o *CreditCards) HasCreditCards() bool {
	if o != nil && !IsNil(o.CreditCards) {
		return true
	}

	return false
}

// SetCreditCards gets a reference to the given []CreditCardMappingType and assigns it to the CreditCards field.
func (o *CreditCards) SetCreditCards(v []CreditCardMappingType) {
	o.CreditCards = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreditCards) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCards) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreditCards) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CreditCards) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreditCards) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCards) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreditCards) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreditCards) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreditCards) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditCards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditCards) {
		toSerialize["creditCards"] = o.CreditCards
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCreditCards struct {
	value *CreditCards
	isSet bool
}

func (v NullableCreditCards) Get() *CreditCards {
	return v.value
}

func (v *NullableCreditCards) Set(val *CreditCards) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCards) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCards(val *CreditCards) *NullableCreditCards {
	return &NullableCreditCards{value: val, isSet: true}
}

func (v NullableCreditCards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


