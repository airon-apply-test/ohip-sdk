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
)

// checks if the PackagePostingRhythmTypeCertainNightsOfTheWeek type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackagePostingRhythmTypeCertainNightsOfTheWeek{}

// PackagePostingRhythmTypeCertainNightsOfTheWeek Post the package on certain nights of the week.
type PackagePostingRhythmTypeCertainNightsOfTheWeek struct {
	Sunday *bool `json:"sunday,omitempty"`
	Monday *bool `json:"monday,omitempty"`
	Tuesday *bool `json:"tuesday,omitempty"`
	Wednesday *bool `json:"wednesday,omitempty"`
	Thursday *bool `json:"thursday,omitempty"`
	Friday *bool `json:"friday,omitempty"`
	Saturday *bool `json:"saturday,omitempty"`
}

// NewPackagePostingRhythmTypeCertainNightsOfTheWeek instantiates a new PackagePostingRhythmTypeCertainNightsOfTheWeek object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagePostingRhythmTypeCertainNightsOfTheWeek() *PackagePostingRhythmTypeCertainNightsOfTheWeek {
	this := PackagePostingRhythmTypeCertainNightsOfTheWeek{}
	return &this
}

// NewPackagePostingRhythmTypeCertainNightsOfTheWeekWithDefaults instantiates a new PackagePostingRhythmTypeCertainNightsOfTheWeek object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagePostingRhythmTypeCertainNightsOfTheWeekWithDefaults() *PackagePostingRhythmTypeCertainNightsOfTheWeek {
	this := PackagePostingRhythmTypeCertainNightsOfTheWeek{}
	return &this
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetSunday() bool {
	if o == nil || IsNil(o.Sunday) {
		var ret bool
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetSundayOk() (*bool, bool) {
	if o == nil || IsNil(o.Sunday) {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) HasSunday() bool {
	if o != nil && !IsNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given bool and assigns it to the Sunday field.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) SetSunday(v bool) {
	o.Sunday = &v
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetMonday() bool {
	if o == nil || IsNil(o.Monday) {
		var ret bool
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetMondayOk() (*bool, bool) {
	if o == nil || IsNil(o.Monday) {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) HasMonday() bool {
	if o != nil && !IsNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given bool and assigns it to the Monday field.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) SetMonday(v bool) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetTuesday() bool {
	if o == nil || IsNil(o.Tuesday) {
		var ret bool
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetTuesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Tuesday) {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) HasTuesday() bool {
	if o != nil && !IsNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given bool and assigns it to the Tuesday field.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) SetTuesday(v bool) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetWednesday() bool {
	if o == nil || IsNil(o.Wednesday) {
		var ret bool
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetWednesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Wednesday) {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) HasWednesday() bool {
	if o != nil && !IsNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given bool and assigns it to the Wednesday field.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) SetWednesday(v bool) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetThursday() bool {
	if o == nil || IsNil(o.Thursday) {
		var ret bool
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetThursdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Thursday) {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) HasThursday() bool {
	if o != nil && !IsNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given bool and assigns it to the Thursday field.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) SetThursday(v bool) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetFriday() bool {
	if o == nil || IsNil(o.Friday) {
		var ret bool
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetFridayOk() (*bool, bool) {
	if o == nil || IsNil(o.Friday) {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) HasFriday() bool {
	if o != nil && !IsNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given bool and assigns it to the Friday field.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) SetFriday(v bool) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetSaturday() bool {
	if o == nil || IsNil(o.Saturday) {
		var ret bool
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) GetSaturdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Saturday) {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) HasSaturday() bool {
	if o != nil && !IsNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given bool and assigns it to the Saturday field.
func (o *PackagePostingRhythmTypeCertainNightsOfTheWeek) SetSaturday(v bool) {
	o.Saturday = &v
}

func (o PackagePostingRhythmTypeCertainNightsOfTheWeek) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackagePostingRhythmTypeCertainNightsOfTheWeek) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullablePackagePostingRhythmTypeCertainNightsOfTheWeek struct {
	value *PackagePostingRhythmTypeCertainNightsOfTheWeek
	isSet bool
}

func (v NullablePackagePostingRhythmTypeCertainNightsOfTheWeek) Get() *PackagePostingRhythmTypeCertainNightsOfTheWeek {
	return v.value
}

func (v *NullablePackagePostingRhythmTypeCertainNightsOfTheWeek) Set(val *PackagePostingRhythmTypeCertainNightsOfTheWeek) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagePostingRhythmTypeCertainNightsOfTheWeek) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagePostingRhythmTypeCertainNightsOfTheWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagePostingRhythmTypeCertainNightsOfTheWeek(val *PackagePostingRhythmTypeCertainNightsOfTheWeek) *NullablePackagePostingRhythmTypeCertainNightsOfTheWeek {
	return &NullablePackagePostingRhythmTypeCertainNightsOfTheWeek{value: val, isSet: true}
}

func (v NullablePackagePostingRhythmTypeCertainNightsOfTheWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagePostingRhythmTypeCertainNightsOfTheWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


