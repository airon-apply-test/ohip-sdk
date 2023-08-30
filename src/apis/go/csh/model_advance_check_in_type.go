/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
	"time"
)

// checks if the AdvanceCheckInType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvanceCheckInType{}

// AdvanceCheckInType Information relating to Reservation's Advance Checked In state and Expected Time of Return
type AdvanceCheckInType struct {
	// Indicates if the reservation is marked as Advance Checked In
	AdvanceCheckedIn *bool `json:"advanceCheckedIn,omitempty"`
	// Expected Return Time
	ExpectedReturnTime *time.Time `json:"expectedReturnTime,omitempty"`
	// ETR Comments
	ETRComments *string `json:"eTRComments,omitempty"`
}

// NewAdvanceCheckInType instantiates a new AdvanceCheckInType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvanceCheckInType() *AdvanceCheckInType {
	this := AdvanceCheckInType{}
	return &this
}

// NewAdvanceCheckInTypeWithDefaults instantiates a new AdvanceCheckInType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvanceCheckInTypeWithDefaults() *AdvanceCheckInType {
	this := AdvanceCheckInType{}
	return &this
}

// GetAdvanceCheckedIn returns the AdvanceCheckedIn field value if set, zero value otherwise.
func (o *AdvanceCheckInType) GetAdvanceCheckedIn() bool {
	if o == nil || IsNil(o.AdvanceCheckedIn) {
		var ret bool
		return ret
	}
	return *o.AdvanceCheckedIn
}

// GetAdvanceCheckedInOk returns a tuple with the AdvanceCheckedIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceCheckInType) GetAdvanceCheckedInOk() (*bool, bool) {
	if o == nil || IsNil(o.AdvanceCheckedIn) {
		return nil, false
	}
	return o.AdvanceCheckedIn, true
}

// HasAdvanceCheckedIn returns a boolean if a field has been set.
func (o *AdvanceCheckInType) HasAdvanceCheckedIn() bool {
	if o != nil && !IsNil(o.AdvanceCheckedIn) {
		return true
	}

	return false
}

// SetAdvanceCheckedIn gets a reference to the given bool and assigns it to the AdvanceCheckedIn field.
func (o *AdvanceCheckInType) SetAdvanceCheckedIn(v bool) {
	o.AdvanceCheckedIn = &v
}

// GetExpectedReturnTime returns the ExpectedReturnTime field value if set, zero value otherwise.
func (o *AdvanceCheckInType) GetExpectedReturnTime() time.Time {
	if o == nil || IsNil(o.ExpectedReturnTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpectedReturnTime
}

// GetExpectedReturnTimeOk returns a tuple with the ExpectedReturnTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceCheckInType) GetExpectedReturnTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpectedReturnTime) {
		return nil, false
	}
	return o.ExpectedReturnTime, true
}

// HasExpectedReturnTime returns a boolean if a field has been set.
func (o *AdvanceCheckInType) HasExpectedReturnTime() bool {
	if o != nil && !IsNil(o.ExpectedReturnTime) {
		return true
	}

	return false
}

// SetExpectedReturnTime gets a reference to the given time.Time and assigns it to the ExpectedReturnTime field.
func (o *AdvanceCheckInType) SetExpectedReturnTime(v time.Time) {
	o.ExpectedReturnTime = &v
}

// GetETRComments returns the ETRComments field value if set, zero value otherwise.
func (o *AdvanceCheckInType) GetETRComments() string {
	if o == nil || IsNil(o.ETRComments) {
		var ret string
		return ret
	}
	return *o.ETRComments
}

// GetETRCommentsOk returns a tuple with the ETRComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceCheckInType) GetETRCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.ETRComments) {
		return nil, false
	}
	return o.ETRComments, true
}

// HasETRComments returns a boolean if a field has been set.
func (o *AdvanceCheckInType) HasETRComments() bool {
	if o != nil && !IsNil(o.ETRComments) {
		return true
	}

	return false
}

// SetETRComments gets a reference to the given string and assigns it to the ETRComments field.
func (o *AdvanceCheckInType) SetETRComments(v string) {
	o.ETRComments = &v
}

func (o AdvanceCheckInType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvanceCheckInType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvanceCheckedIn) {
		toSerialize["advanceCheckedIn"] = o.AdvanceCheckedIn
	}
	if !IsNil(o.ExpectedReturnTime) {
		toSerialize["expectedReturnTime"] = o.ExpectedReturnTime
	}
	if !IsNil(o.ETRComments) {
		toSerialize["eTRComments"] = o.ETRComments
	}
	return toSerialize, nil
}

type NullableAdvanceCheckInType struct {
	value *AdvanceCheckInType
	isSet bool
}

func (v NullableAdvanceCheckInType) Get() *AdvanceCheckInType {
	return v.value
}

func (v *NullableAdvanceCheckInType) Set(val *AdvanceCheckInType) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvanceCheckInType) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvanceCheckInType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvanceCheckInType(val *AdvanceCheckInType) *NullableAdvanceCheckInType {
	return &NullableAdvanceCheckInType{value: val, isSet: true}
}

func (v NullableAdvanceCheckInType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvanceCheckInType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


