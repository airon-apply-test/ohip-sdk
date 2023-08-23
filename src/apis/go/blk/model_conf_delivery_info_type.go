/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ConfDeliveryInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfDeliveryInfoType{}

// ConfDeliveryInfoType struct for ConfDeliveryInfoType
type ConfDeliveryInfoType struct {
	CommunicationType *ConfDeliveryMethod `json:"communicationType,omitempty"`
	LastStatus *SentConfirmationStatus `json:"lastStatus,omitempty"`
	// Date of last attempt to send confirmation letter.
	LastAttempted *time.Time `json:"lastAttempted,omitempty"`
	// Number of success letter sent.
	SuccessfulTries *int32 `json:"successfulTries,omitempty"`
}

// NewConfDeliveryInfoType instantiates a new ConfDeliveryInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfDeliveryInfoType() *ConfDeliveryInfoType {
	this := ConfDeliveryInfoType{}
	return &this
}

// NewConfDeliveryInfoTypeWithDefaults instantiates a new ConfDeliveryInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfDeliveryInfoTypeWithDefaults() *ConfDeliveryInfoType {
	this := ConfDeliveryInfoType{}
	return &this
}

// GetCommunicationType returns the CommunicationType field value if set, zero value otherwise.
func (o *ConfDeliveryInfoType) GetCommunicationType() ConfDeliveryMethod {
	if o == nil || IsNil(o.CommunicationType) {
		var ret ConfDeliveryMethod
		return ret
	}
	return *o.CommunicationType
}

// GetCommunicationTypeOk returns a tuple with the CommunicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfDeliveryInfoType) GetCommunicationTypeOk() (*ConfDeliveryMethod, bool) {
	if o == nil || IsNil(o.CommunicationType) {
		return nil, false
	}
	return o.CommunicationType, true
}

// HasCommunicationType returns a boolean if a field has been set.
func (o *ConfDeliveryInfoType) HasCommunicationType() bool {
	if o != nil && !IsNil(o.CommunicationType) {
		return true
	}

	return false
}

// SetCommunicationType gets a reference to the given ConfDeliveryMethod and assigns it to the CommunicationType field.
func (o *ConfDeliveryInfoType) SetCommunicationType(v ConfDeliveryMethod) {
	o.CommunicationType = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *ConfDeliveryInfoType) GetLastStatus() SentConfirmationStatus {
	if o == nil || IsNil(o.LastStatus) {
		var ret SentConfirmationStatus
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfDeliveryInfoType) GetLastStatusOk() (*SentConfirmationStatus, bool) {
	if o == nil || IsNil(o.LastStatus) {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *ConfDeliveryInfoType) HasLastStatus() bool {
	if o != nil && !IsNil(o.LastStatus) {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given SentConfirmationStatus and assigns it to the LastStatus field.
func (o *ConfDeliveryInfoType) SetLastStatus(v SentConfirmationStatus) {
	o.LastStatus = &v
}

// GetLastAttempted returns the LastAttempted field value if set, zero value otherwise.
func (o *ConfDeliveryInfoType) GetLastAttempted() time.Time {
	if o == nil || IsNil(o.LastAttempted) {
		var ret time.Time
		return ret
	}
	return *o.LastAttempted
}

// GetLastAttemptedOk returns a tuple with the LastAttempted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfDeliveryInfoType) GetLastAttemptedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAttempted) {
		return nil, false
	}
	return o.LastAttempted, true
}

// HasLastAttempted returns a boolean if a field has been set.
func (o *ConfDeliveryInfoType) HasLastAttempted() bool {
	if o != nil && !IsNil(o.LastAttempted) {
		return true
	}

	return false
}

// SetLastAttempted gets a reference to the given time.Time and assigns it to the LastAttempted field.
func (o *ConfDeliveryInfoType) SetLastAttempted(v time.Time) {
	o.LastAttempted = &v
}

// GetSuccessfulTries returns the SuccessfulTries field value if set, zero value otherwise.
func (o *ConfDeliveryInfoType) GetSuccessfulTries() int32 {
	if o == nil || IsNil(o.SuccessfulTries) {
		var ret int32
		return ret
	}
	return *o.SuccessfulTries
}

// GetSuccessfulTriesOk returns a tuple with the SuccessfulTries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfDeliveryInfoType) GetSuccessfulTriesOk() (*int32, bool) {
	if o == nil || IsNil(o.SuccessfulTries) {
		return nil, false
	}
	return o.SuccessfulTries, true
}

// HasSuccessfulTries returns a boolean if a field has been set.
func (o *ConfDeliveryInfoType) HasSuccessfulTries() bool {
	if o != nil && !IsNil(o.SuccessfulTries) {
		return true
	}

	return false
}

// SetSuccessfulTries gets a reference to the given int32 and assigns it to the SuccessfulTries field.
func (o *ConfDeliveryInfoType) SetSuccessfulTries(v int32) {
	o.SuccessfulTries = &v
}

func (o ConfDeliveryInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfDeliveryInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommunicationType) {
		toSerialize["communicationType"] = o.CommunicationType
	}
	if !IsNil(o.LastStatus) {
		toSerialize["lastStatus"] = o.LastStatus
	}
	if !IsNil(o.LastAttempted) {
		toSerialize["lastAttempted"] = o.LastAttempted
	}
	if !IsNil(o.SuccessfulTries) {
		toSerialize["successfulTries"] = o.SuccessfulTries
	}
	return toSerialize, nil
}

type NullableConfDeliveryInfoType struct {
	value *ConfDeliveryInfoType
	isSet bool
}

func (v NullableConfDeliveryInfoType) Get() *ConfDeliveryInfoType {
	return v.value
}

func (v *NullableConfDeliveryInfoType) Set(val *ConfDeliveryInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfDeliveryInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfDeliveryInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfDeliveryInfoType(val *ConfDeliveryInfoType) *NullableConfDeliveryInfoType {
	return &NullableConfDeliveryInfoType{value: val, isSet: true}
}

func (v NullableConfDeliveryInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfDeliveryInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


