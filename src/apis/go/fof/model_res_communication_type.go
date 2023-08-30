/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the ResCommunicationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResCommunicationType{}

// ResCommunicationType Communication details for a reservation.
type ResCommunicationType struct {
	Telephones *ResCommunicationTypeTelephones `json:"telephones,omitempty"`
	Emails *ResCommunicationTypeEmails `json:"emails,omitempty"`
}

// NewResCommunicationType instantiates a new ResCommunicationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResCommunicationType() *ResCommunicationType {
	this := ResCommunicationType{}
	return &this
}

// NewResCommunicationTypeWithDefaults instantiates a new ResCommunicationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResCommunicationTypeWithDefaults() *ResCommunicationType {
	this := ResCommunicationType{}
	return &this
}

// GetTelephones returns the Telephones field value if set, zero value otherwise.
func (o *ResCommunicationType) GetTelephones() ResCommunicationTypeTelephones {
	if o == nil || IsNil(o.Telephones) {
		var ret ResCommunicationTypeTelephones
		return ret
	}
	return *o.Telephones
}

// GetTelephonesOk returns a tuple with the Telephones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationType) GetTelephonesOk() (*ResCommunicationTypeTelephones, bool) {
	if o == nil || IsNil(o.Telephones) {
		return nil, false
	}
	return o.Telephones, true
}

// HasTelephones returns a boolean if a field has been set.
func (o *ResCommunicationType) HasTelephones() bool {
	if o != nil && !IsNil(o.Telephones) {
		return true
	}

	return false
}

// SetTelephones gets a reference to the given ResCommunicationTypeTelephones and assigns it to the Telephones field.
func (o *ResCommunicationType) SetTelephones(v ResCommunicationTypeTelephones) {
	o.Telephones = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ResCommunicationType) GetEmails() ResCommunicationTypeEmails {
	if o == nil || IsNil(o.Emails) {
		var ret ResCommunicationTypeEmails
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationType) GetEmailsOk() (*ResCommunicationTypeEmails, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ResCommunicationType) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given ResCommunicationTypeEmails and assigns it to the Emails field.
func (o *ResCommunicationType) SetEmails(v ResCommunicationTypeEmails) {
	o.Emails = &v
}

func (o ResCommunicationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResCommunicationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Telephones) {
		toSerialize["telephones"] = o.Telephones
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	return toSerialize, nil
}

type NullableResCommunicationType struct {
	value *ResCommunicationType
	isSet bool
}

func (v NullableResCommunicationType) Get() *ResCommunicationType {
	return v.value
}

func (v *NullableResCommunicationType) Set(val *ResCommunicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableResCommunicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableResCommunicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResCommunicationType(val *ResCommunicationType) *NullableResCommunicationType {
	return &NullableResCommunicationType{value: val, isSet: true}
}

func (v NullableResCommunicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResCommunicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


