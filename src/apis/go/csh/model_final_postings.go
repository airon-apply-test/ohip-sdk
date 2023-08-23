/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FinalPostings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinalPostings{}

// FinalPostings Request to apply any final charges or payments to a reservation prior to checkout. This operation should be called prior to the guest settlement which would then reflect the balance the guest has to pay.
type FinalPostings struct {
	Reservation *CheckoutReservationType `json:"reservation,omitempty"`
	ResponseInstruction *ResponseInstructionType `json:"responseInstruction,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFinalPostings instantiates a new FinalPostings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinalPostings() *FinalPostings {
	this := FinalPostings{}
	return &this
}

// NewFinalPostingsWithDefaults instantiates a new FinalPostings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinalPostingsWithDefaults() *FinalPostings {
	this := FinalPostings{}
	return &this
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *FinalPostings) GetReservation() CheckoutReservationType {
	if o == nil || IsNil(o.Reservation) {
		var ret CheckoutReservationType
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalPostings) GetReservationOk() (*CheckoutReservationType, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *FinalPostings) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given CheckoutReservationType and assigns it to the Reservation field.
func (o *FinalPostings) SetReservation(v CheckoutReservationType) {
	o.Reservation = &v
}

// GetResponseInstruction returns the ResponseInstruction field value if set, zero value otherwise.
func (o *FinalPostings) GetResponseInstruction() ResponseInstructionType {
	if o == nil || IsNil(o.ResponseInstruction) {
		var ret ResponseInstructionType
		return ret
	}
	return *o.ResponseInstruction
}

// GetResponseInstructionOk returns a tuple with the ResponseInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalPostings) GetResponseInstructionOk() (*ResponseInstructionType, bool) {
	if o == nil || IsNil(o.ResponseInstruction) {
		return nil, false
	}
	return o.ResponseInstruction, true
}

// HasResponseInstruction returns a boolean if a field has been set.
func (o *FinalPostings) HasResponseInstruction() bool {
	if o != nil && !IsNil(o.ResponseInstruction) {
		return true
	}

	return false
}

// SetResponseInstruction gets a reference to the given ResponseInstructionType and assigns it to the ResponseInstruction field.
func (o *FinalPostings) SetResponseInstruction(v ResponseInstructionType) {
	o.ResponseInstruction = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FinalPostings) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalPostings) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FinalPostings) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FinalPostings) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FinalPostings) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalPostings) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FinalPostings) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FinalPostings) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FinalPostings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinalPostings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reservation) {
		toSerialize["reservation"] = o.Reservation
	}
	if !IsNil(o.ResponseInstruction) {
		toSerialize["responseInstruction"] = o.ResponseInstruction
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFinalPostings struct {
	value *FinalPostings
	isSet bool
}

func (v NullableFinalPostings) Get() *FinalPostings {
	return v.value
}

func (v *NullableFinalPostings) Set(val *FinalPostings) {
	v.value = val
	v.isSet = true
}

func (v NullableFinalPostings) IsSet() bool {
	return v.isSet
}

func (v *NullableFinalPostings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinalPostings(val *FinalPostings) *NullableFinalPostings {
	return &NullableFinalPostings{value: val, isSet: true}
}

func (v NullableFinalPostings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinalPostings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


