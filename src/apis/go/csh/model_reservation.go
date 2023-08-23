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

// checks if the Reservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reservation{}

// Reservation Checkout request can be used to verify a reservation for checkout and do an actual checkout. In case the verificationOnly attribute is sent false, the operation will perform an actual checkout. In case the verificationOnly attribute is sent true, the operation goes through the reservation in question and verifies if it's Ok to checkout, otherwise the verification status element will provide you the verification code. The verification codes are described in the documentation of verificationOnly attribute.
type Reservation struct {
	Reservation *CheckoutReservationType `json:"reservation,omitempty"`
	// VerificationOnly \"true\" validates the hotel code and reservation id supplied in the request. VerificationOnly \"false\" validates and then check-out the guest using the request details. Following codes might be returned during the verification: FOF00065 - The hotel code and/or reservation id is missing. FOF00066 - Reservation can't be found with the supplied hotel code and reservation id. FOF00107 - The guest's departure is not scheduled for today. Check-out not possible. FOF00109 - The guest's reservation is not in a valid status for check-out. FOF00108 - The guest has a balance. Check-out not possible.
	VerificationOnly *bool `json:"verificationOnly,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewReservation instantiates a new Reservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservation() *Reservation {
	this := Reservation{}
	return &this
}

// NewReservationWithDefaults instantiates a new Reservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationWithDefaults() *Reservation {
	this := Reservation{}
	return &this
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *Reservation) GetReservation() CheckoutReservationType {
	if o == nil || IsNil(o.Reservation) {
		var ret CheckoutReservationType
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetReservationOk() (*CheckoutReservationType, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *Reservation) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given CheckoutReservationType and assigns it to the Reservation field.
func (o *Reservation) SetReservation(v CheckoutReservationType) {
	o.Reservation = &v
}

// GetVerificationOnly returns the VerificationOnly field value if set, zero value otherwise.
func (o *Reservation) GetVerificationOnly() bool {
	if o == nil || IsNil(o.VerificationOnly) {
		var ret bool
		return ret
	}
	return *o.VerificationOnly
}

// GetVerificationOnlyOk returns a tuple with the VerificationOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetVerificationOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.VerificationOnly) {
		return nil, false
	}
	return o.VerificationOnly, true
}

// HasVerificationOnly returns a boolean if a field has been set.
func (o *Reservation) HasVerificationOnly() bool {
	if o != nil && !IsNil(o.VerificationOnly) {
		return true
	}

	return false
}

// SetVerificationOnly gets a reference to the given bool and assigns it to the VerificationOnly field.
func (o *Reservation) SetVerificationOnly(v bool) {
	o.VerificationOnly = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Reservation) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Reservation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *Reservation) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Reservation) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reservation) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Reservation) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *Reservation) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o Reservation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reservation) {
		toSerialize["reservation"] = o.Reservation
	}
	if !IsNil(o.VerificationOnly) {
		toSerialize["verificationOnly"] = o.VerificationOnly
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableReservation struct {
	value *Reservation
	isSet bool
}

func (v NullableReservation) Get() *Reservation {
	return v.value
}

func (v *NullableReservation) Set(val *Reservation) {
	v.value = val
	v.isSet = true
}

func (v NullableReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservation(val *Reservation) *NullableReservation {
	return &NullableReservation{value: val, isSet: true}
}

func (v NullableReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


