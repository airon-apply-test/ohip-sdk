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

// checks if the TimeSpanDaysOfWeekType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSpanDaysOfWeekType{}

// TimeSpanDaysOfWeekType Container for Time span with days of week.
type TimeSpanDaysOfWeekType struct {
	TimeSpan *TimeSpanType `json:"timeSpan,omitempty"`
	Sunday *bool `json:"sunday,omitempty"`
	Monday *bool `json:"monday,omitempty"`
	Tuesday *bool `json:"tuesday,omitempty"`
	Wednesday *bool `json:"wednesday,omitempty"`
	Thursday *bool `json:"thursday,omitempty"`
	Friday *bool `json:"friday,omitempty"`
	Saturday *bool `json:"saturday,omitempty"`
}

// NewTimeSpanDaysOfWeekType instantiates a new TimeSpanDaysOfWeekType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSpanDaysOfWeekType() *TimeSpanDaysOfWeekType {
	this := TimeSpanDaysOfWeekType{}
	return &this
}

// NewTimeSpanDaysOfWeekTypeWithDefaults instantiates a new TimeSpanDaysOfWeekType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSpanDaysOfWeekTypeWithDefaults() *TimeSpanDaysOfWeekType {
	this := TimeSpanDaysOfWeekType{}
	return &this
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetTimeSpan() TimeSpanType {
	if o == nil || IsNil(o.TimeSpan) {
		var ret TimeSpanType
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetTimeSpanOk() (*TimeSpanType, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given TimeSpanType and assigns it to the TimeSpan field.
func (o *TimeSpanDaysOfWeekType) SetTimeSpan(v TimeSpanType) {
	o.TimeSpan = &v
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetSunday() bool {
	if o == nil || IsNil(o.Sunday) {
		var ret bool
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetSundayOk() (*bool, bool) {
	if o == nil || IsNil(o.Sunday) {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasSunday() bool {
	if o != nil && !IsNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given bool and assigns it to the Sunday field.
func (o *TimeSpanDaysOfWeekType) SetSunday(v bool) {
	o.Sunday = &v
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetMonday() bool {
	if o == nil || IsNil(o.Monday) {
		var ret bool
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetMondayOk() (*bool, bool) {
	if o == nil || IsNil(o.Monday) {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasMonday() bool {
	if o != nil && !IsNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given bool and assigns it to the Monday field.
func (o *TimeSpanDaysOfWeekType) SetMonday(v bool) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetTuesday() bool {
	if o == nil || IsNil(o.Tuesday) {
		var ret bool
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetTuesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Tuesday) {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasTuesday() bool {
	if o != nil && !IsNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given bool and assigns it to the Tuesday field.
func (o *TimeSpanDaysOfWeekType) SetTuesday(v bool) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetWednesday() bool {
	if o == nil || IsNil(o.Wednesday) {
		var ret bool
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetWednesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Wednesday) {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasWednesday() bool {
	if o != nil && !IsNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given bool and assigns it to the Wednesday field.
func (o *TimeSpanDaysOfWeekType) SetWednesday(v bool) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetThursday() bool {
	if o == nil || IsNil(o.Thursday) {
		var ret bool
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetThursdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Thursday) {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasThursday() bool {
	if o != nil && !IsNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given bool and assigns it to the Thursday field.
func (o *TimeSpanDaysOfWeekType) SetThursday(v bool) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetFriday() bool {
	if o == nil || IsNil(o.Friday) {
		var ret bool
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetFridayOk() (*bool, bool) {
	if o == nil || IsNil(o.Friday) {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasFriday() bool {
	if o != nil && !IsNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given bool and assigns it to the Friday field.
func (o *TimeSpanDaysOfWeekType) SetFriday(v bool) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *TimeSpanDaysOfWeekType) GetSaturday() bool {
	if o == nil || IsNil(o.Saturday) {
		var ret bool
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSpanDaysOfWeekType) GetSaturdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Saturday) {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *TimeSpanDaysOfWeekType) HasSaturday() bool {
	if o != nil && !IsNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given bool and assigns it to the Saturday field.
func (o *TimeSpanDaysOfWeekType) SetSaturday(v bool) {
	o.Saturday = &v
}

func (o TimeSpanDaysOfWeekType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSpanDaysOfWeekType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	if !IsNil(o.Sunday) {
		toSerialize["sunday"] = o.Sunday
	}
	if !IsNil(o.Monday) {
		toSerialize["monday"] = o.Monday
	}
	if !IsNil(o.Tuesday) {
		toSerialize["tuesday"] = o.Tuesday
	}
	if !IsNil(o.Wednesday) {
		toSerialize["wednesday"] = o.Wednesday
	}
	if !IsNil(o.Thursday) {
		toSerialize["thursday"] = o.Thursday
	}
	if !IsNil(o.Friday) {
		toSerialize["friday"] = o.Friday
	}
	if !IsNil(o.Saturday) {
		toSerialize["saturday"] = o.Saturday
	}
	return toSerialize, nil
}

type NullableTimeSpanDaysOfWeekType struct {
	value *TimeSpanDaysOfWeekType
	isSet bool
}

func (v NullableTimeSpanDaysOfWeekType) Get() *TimeSpanDaysOfWeekType {
	return v.value
}

func (v *NullableTimeSpanDaysOfWeekType) Set(val *TimeSpanDaysOfWeekType) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSpanDaysOfWeekType) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSpanDaysOfWeekType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSpanDaysOfWeekType(val *TimeSpanDaysOfWeekType) *NullableTimeSpanDaysOfWeekType {
	return &NullableTimeSpanDaysOfWeekType{value: val, isSet: true}
}

func (v NullableTimeSpanDaysOfWeekType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSpanDaysOfWeekType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


