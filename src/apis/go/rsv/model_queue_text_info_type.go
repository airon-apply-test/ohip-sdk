/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the QueueTextInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueTextInfoType{}

// QueueTextInfoType Information regarding the message sent to guest.
type QueueTextInfoType struct {
	// Time the text was last sent to the user.
	SentTime *time.Time `json:"sentTime,omitempty"`
	// User name of the user who sent message.
	SentBy *string `json:"sentBy,omitempty"`
}

// NewQueueTextInfoType instantiates a new QueueTextInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueTextInfoType() *QueueTextInfoType {
	this := QueueTextInfoType{}
	return &this
}

// NewQueueTextInfoTypeWithDefaults instantiates a new QueueTextInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueTextInfoTypeWithDefaults() *QueueTextInfoType {
	this := QueueTextInfoType{}
	return &this
}

// GetSentTime returns the SentTime field value if set, zero value otherwise.
func (o *QueueTextInfoType) GetSentTime() time.Time {
	if o == nil || IsNil(o.SentTime) {
		var ret time.Time
		return ret
	}
	return *o.SentTime
}

// GetSentTimeOk returns a tuple with the SentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueTextInfoType) GetSentTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentTime) {
		return nil, false
	}
	return o.SentTime, true
}

// HasSentTime returns a boolean if a field has been set.
func (o *QueueTextInfoType) HasSentTime() bool {
	if o != nil && !IsNil(o.SentTime) {
		return true
	}

	return false
}

// SetSentTime gets a reference to the given time.Time and assigns it to the SentTime field.
func (o *QueueTextInfoType) SetSentTime(v time.Time) {
	o.SentTime = &v
}

// GetSentBy returns the SentBy field value if set, zero value otherwise.
func (o *QueueTextInfoType) GetSentBy() string {
	if o == nil || IsNil(o.SentBy) {
		var ret string
		return ret
	}
	return *o.SentBy
}

// GetSentByOk returns a tuple with the SentBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueTextInfoType) GetSentByOk() (*string, bool) {
	if o == nil || IsNil(o.SentBy) {
		return nil, false
	}
	return o.SentBy, true
}

// HasSentBy returns a boolean if a field has been set.
func (o *QueueTextInfoType) HasSentBy() bool {
	if o != nil && !IsNil(o.SentBy) {
		return true
	}

	return false
}

// SetSentBy gets a reference to the given string and assigns it to the SentBy field.
func (o *QueueTextInfoType) SetSentBy(v string) {
	o.SentBy = &v
}

func (o QueueTextInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueTextInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SentTime) {
		toSerialize["sentTime"] = o.SentTime
	}
	if !IsNil(o.SentBy) {
		toSerialize["sentBy"] = o.SentBy
	}
	return toSerialize, nil
}

type NullableQueueTextInfoType struct {
	value *QueueTextInfoType
	isSet bool
}

func (v NullableQueueTextInfoType) Get() *QueueTextInfoType {
	return v.value
}

func (v *NullableQueueTextInfoType) Set(val *QueueTextInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueTextInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueTextInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueTextInfoType(val *QueueTextInfoType) *NullableQueueTextInfoType {
	return &NullableQueueTextInfoType{value: val, isSet: true}
}

func (v NullableQueueTextInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueTextInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


