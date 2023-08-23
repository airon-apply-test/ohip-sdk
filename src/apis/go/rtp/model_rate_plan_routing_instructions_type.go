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

// checks if the RatePlanRoutingInstructionsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatePlanRoutingInstructionsType{}

// RatePlanRoutingInstructionsType Routing fetchInstructions details for a rate plan.
type RatePlanRoutingInstructionsType struct {
	ProfileType *RoutingProfileTypeType `json:"profileType,omitempty"`
	// List of Transaction codes info.
	TransactionCodes []TrxInfoType `json:"transactionCodes,omitempty"`
	// Set of Billing Instruction codes.
	BillingInstructions []BillingInstructionType `json:"billingInstructions,omitempty"`
}

// NewRatePlanRoutingInstructionsType instantiates a new RatePlanRoutingInstructionsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatePlanRoutingInstructionsType() *RatePlanRoutingInstructionsType {
	this := RatePlanRoutingInstructionsType{}
	return &this
}

// NewRatePlanRoutingInstructionsTypeWithDefaults instantiates a new RatePlanRoutingInstructionsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatePlanRoutingInstructionsTypeWithDefaults() *RatePlanRoutingInstructionsType {
	this := RatePlanRoutingInstructionsType{}
	return &this
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *RatePlanRoutingInstructionsType) GetProfileType() RoutingProfileTypeType {
	if o == nil || IsNil(o.ProfileType) {
		var ret RoutingProfileTypeType
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanRoutingInstructionsType) GetProfileTypeOk() (*RoutingProfileTypeType, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *RatePlanRoutingInstructionsType) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given RoutingProfileTypeType and assigns it to the ProfileType field.
func (o *RatePlanRoutingInstructionsType) SetProfileType(v RoutingProfileTypeType) {
	o.ProfileType = &v
}

// GetTransactionCodes returns the TransactionCodes field value if set, zero value otherwise.
func (o *RatePlanRoutingInstructionsType) GetTransactionCodes() []TrxInfoType {
	if o == nil || IsNil(o.TransactionCodes) {
		var ret []TrxInfoType
		return ret
	}
	return o.TransactionCodes
}

// GetTransactionCodesOk returns a tuple with the TransactionCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanRoutingInstructionsType) GetTransactionCodesOk() ([]TrxInfoType, bool) {
	if o == nil || IsNil(o.TransactionCodes) {
		return nil, false
	}
	return o.TransactionCodes, true
}

// HasTransactionCodes returns a boolean if a field has been set.
func (o *RatePlanRoutingInstructionsType) HasTransactionCodes() bool {
	if o != nil && !IsNil(o.TransactionCodes) {
		return true
	}

	return false
}

// SetTransactionCodes gets a reference to the given []TrxInfoType and assigns it to the TransactionCodes field.
func (o *RatePlanRoutingInstructionsType) SetTransactionCodes(v []TrxInfoType) {
	o.TransactionCodes = v
}

// GetBillingInstructions returns the BillingInstructions field value if set, zero value otherwise.
func (o *RatePlanRoutingInstructionsType) GetBillingInstructions() []BillingInstructionType {
	if o == nil || IsNil(o.BillingInstructions) {
		var ret []BillingInstructionType
		return ret
	}
	return o.BillingInstructions
}

// GetBillingInstructionsOk returns a tuple with the BillingInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanRoutingInstructionsType) GetBillingInstructionsOk() ([]BillingInstructionType, bool) {
	if o == nil || IsNil(o.BillingInstructions) {
		return nil, false
	}
	return o.BillingInstructions, true
}

// HasBillingInstructions returns a boolean if a field has been set.
func (o *RatePlanRoutingInstructionsType) HasBillingInstructions() bool {
	if o != nil && !IsNil(o.BillingInstructions) {
		return true
	}

	return false
}

// SetBillingInstructions gets a reference to the given []BillingInstructionType and assigns it to the BillingInstructions field.
func (o *RatePlanRoutingInstructionsType) SetBillingInstructions(v []BillingInstructionType) {
	o.BillingInstructions = v
}

func (o RatePlanRoutingInstructionsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatePlanRoutingInstructionsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	if !IsNil(o.TransactionCodes) {
		toSerialize["transactionCodes"] = o.TransactionCodes
	}
	if !IsNil(o.BillingInstructions) {
		toSerialize["billingInstructions"] = o.BillingInstructions
	}
	return toSerialize, nil
}

type NullableRatePlanRoutingInstructionsType struct {
	value *RatePlanRoutingInstructionsType
	isSet bool
}

func (v NullableRatePlanRoutingInstructionsType) Get() *RatePlanRoutingInstructionsType {
	return v.value
}

func (v *NullableRatePlanRoutingInstructionsType) Set(val *RatePlanRoutingInstructionsType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanRoutingInstructionsType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanRoutingInstructionsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanRoutingInstructionsType(val *RatePlanRoutingInstructionsType) *NullableRatePlanRoutingInstructionsType {
	return &NullableRatePlanRoutingInstructionsType{value: val, isSet: true}
}

func (v NullableRatePlanRoutingInstructionsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanRoutingInstructionsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


