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
	"time"
)

// checks if the TraceTimeInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceTimeInfoType{}

// TraceTimeInfoType struct for TraceTimeInfoType
type TraceTimeInfoType struct {
	DateTimeSpan *DateTimeSpanType `json:"dateTimeSpan,omitempty"`
	// Date of the trace.
	TraceOn *time.Time `json:"traceOn,omitempty"`
	// Time of the trace
	TraceTime *string `json:"traceTime,omitempty"`
	// User that entered this trace.
	EnteredBy *string `json:"enteredBy,omitempty"`
}

// NewTraceTimeInfoType instantiates a new TraceTimeInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceTimeInfoType() *TraceTimeInfoType {
	this := TraceTimeInfoType{}
	return &this
}

// NewTraceTimeInfoTypeWithDefaults instantiates a new TraceTimeInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceTimeInfoTypeWithDefaults() *TraceTimeInfoType {
	this := TraceTimeInfoType{}
	return &this
}

// GetDateTimeSpan returns the DateTimeSpan field value if set, zero value otherwise.
func (o *TraceTimeInfoType) GetDateTimeSpan() DateTimeSpanType {
	if o == nil || IsNil(o.DateTimeSpan) {
		var ret DateTimeSpanType
		return ret
	}
	return *o.DateTimeSpan
}

// GetDateTimeSpanOk returns a tuple with the DateTimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceTimeInfoType) GetDateTimeSpanOk() (*DateTimeSpanType, bool) {
	if o == nil || IsNil(o.DateTimeSpan) {
		return nil, false
	}
	return o.DateTimeSpan, true
}

// HasDateTimeSpan returns a boolean if a field has been set.
func (o *TraceTimeInfoType) HasDateTimeSpan() bool {
	if o != nil && !IsNil(o.DateTimeSpan) {
		return true
	}

	return false
}

// SetDateTimeSpan gets a reference to the given DateTimeSpanType and assigns it to the DateTimeSpan field.
func (o *TraceTimeInfoType) SetDateTimeSpan(v DateTimeSpanType) {
	o.DateTimeSpan = &v
}

// GetTraceOn returns the TraceOn field value if set, zero value otherwise.
func (o *TraceTimeInfoType) GetTraceOn() time.Time {
	if o == nil || IsNil(o.TraceOn) {
		var ret time.Time
		return ret
	}
	return *o.TraceOn
}

// GetTraceOnOk returns a tuple with the TraceOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceTimeInfoType) GetTraceOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TraceOn) {
		return nil, false
	}
	return o.TraceOn, true
}

// HasTraceOn returns a boolean if a field has been set.
func (o *TraceTimeInfoType) HasTraceOn() bool {
	if o != nil && !IsNil(o.TraceOn) {
		return true
	}

	return false
}

// SetTraceOn gets a reference to the given time.Time and assigns it to the TraceOn field.
func (o *TraceTimeInfoType) SetTraceOn(v time.Time) {
	o.TraceOn = &v
}

// GetTraceTime returns the TraceTime field value if set, zero value otherwise.
func (o *TraceTimeInfoType) GetTraceTime() string {
	if o == nil || IsNil(o.TraceTime) {
		var ret string
		return ret
	}
	return *o.TraceTime
}

// GetTraceTimeOk returns a tuple with the TraceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceTimeInfoType) GetTraceTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TraceTime) {
		return nil, false
	}
	return o.TraceTime, true
}

// HasTraceTime returns a boolean if a field has been set.
func (o *TraceTimeInfoType) HasTraceTime() bool {
	if o != nil && !IsNil(o.TraceTime) {
		return true
	}

	return false
}

// SetTraceTime gets a reference to the given string and assigns it to the TraceTime field.
func (o *TraceTimeInfoType) SetTraceTime(v string) {
	o.TraceTime = &v
}

// GetEnteredBy returns the EnteredBy field value if set, zero value otherwise.
func (o *TraceTimeInfoType) GetEnteredBy() string {
	if o == nil || IsNil(o.EnteredBy) {
		var ret string
		return ret
	}
	return *o.EnteredBy
}

// GetEnteredByOk returns a tuple with the EnteredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceTimeInfoType) GetEnteredByOk() (*string, bool) {
	if o == nil || IsNil(o.EnteredBy) {
		return nil, false
	}
	return o.EnteredBy, true
}

// HasEnteredBy returns a boolean if a field has been set.
func (o *TraceTimeInfoType) HasEnteredBy() bool {
	if o != nil && !IsNil(o.EnteredBy) {
		return true
	}

	return false
}

// SetEnteredBy gets a reference to the given string and assigns it to the EnteredBy field.
func (o *TraceTimeInfoType) SetEnteredBy(v string) {
	o.EnteredBy = &v
}

func (o TraceTimeInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceTimeInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateTimeSpan) {
		toSerialize["dateTimeSpan"] = o.DateTimeSpan
	}
	if !IsNil(o.TraceOn) {
		toSerialize["traceOn"] = o.TraceOn
	}
	if !IsNil(o.TraceTime) {
		toSerialize["traceTime"] = o.TraceTime
	}
	if !IsNil(o.EnteredBy) {
		toSerialize["enteredBy"] = o.EnteredBy
	}
	return toSerialize, nil
}

type NullableTraceTimeInfoType struct {
	value *TraceTimeInfoType
	isSet bool
}

func (v NullableTraceTimeInfoType) Get() *TraceTimeInfoType {
	return v.value
}

func (v *NullableTraceTimeInfoType) Set(val *TraceTimeInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceTimeInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceTimeInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceTimeInfoType(val *TraceTimeInfoType) *NullableTraceTimeInfoType {
	return &NullableTraceTimeInfoType{value: val, isSet: true}
}

func (v NullableTraceTimeInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceTimeInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


