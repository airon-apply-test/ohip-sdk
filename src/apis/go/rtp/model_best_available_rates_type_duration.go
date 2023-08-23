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

// checks if the BestAvailableRatesTypeDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BestAvailableRatesTypeDuration{}

// BestAvailableRatesTypeDuration Days for which best available rates will be considered
type BestAvailableRatesTypeDuration struct {
	Sunday *bool `json:"sunday,omitempty"`
	Monday *bool `json:"monday,omitempty"`
	Tuesday *bool `json:"tuesday,omitempty"`
	Wednesday *bool `json:"wednesday,omitempty"`
	Thursday *bool `json:"thursday,omitempty"`
	Friday *bool `json:"friday,omitempty"`
	Saturday *bool `json:"saturday,omitempty"`
}

// NewBestAvailableRatesTypeDuration instantiates a new BestAvailableRatesTypeDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBestAvailableRatesTypeDuration() *BestAvailableRatesTypeDuration {
	this := BestAvailableRatesTypeDuration{}
	return &this
}

// NewBestAvailableRatesTypeDurationWithDefaults instantiates a new BestAvailableRatesTypeDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBestAvailableRatesTypeDurationWithDefaults() *BestAvailableRatesTypeDuration {
	this := BestAvailableRatesTypeDuration{}
	return &this
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *BestAvailableRatesTypeDuration) GetSunday() bool {
	if o == nil || IsNil(o.Sunday) {
		var ret bool
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesTypeDuration) GetSundayOk() (*bool, bool) {
	if o == nil || IsNil(o.Sunday) {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *BestAvailableRatesTypeDuration) HasSunday() bool {
	if o != nil && !IsNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given bool and assigns it to the Sunday field.
func (o *BestAvailableRatesTypeDuration) SetSunday(v bool) {
	o.Sunday = &v
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *BestAvailableRatesTypeDuration) GetMonday() bool {
	if o == nil || IsNil(o.Monday) {
		var ret bool
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesTypeDuration) GetMondayOk() (*bool, bool) {
	if o == nil || IsNil(o.Monday) {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *BestAvailableRatesTypeDuration) HasMonday() bool {
	if o != nil && !IsNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given bool and assigns it to the Monday field.
func (o *BestAvailableRatesTypeDuration) SetMonday(v bool) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *BestAvailableRatesTypeDuration) GetTuesday() bool {
	if o == nil || IsNil(o.Tuesday) {
		var ret bool
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesTypeDuration) GetTuesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Tuesday) {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *BestAvailableRatesTypeDuration) HasTuesday() bool {
	if o != nil && !IsNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given bool and assigns it to the Tuesday field.
func (o *BestAvailableRatesTypeDuration) SetTuesday(v bool) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *BestAvailableRatesTypeDuration) GetWednesday() bool {
	if o == nil || IsNil(o.Wednesday) {
		var ret bool
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesTypeDuration) GetWednesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Wednesday) {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *BestAvailableRatesTypeDuration) HasWednesday() bool {
	if o != nil && !IsNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given bool and assigns it to the Wednesday field.
func (o *BestAvailableRatesTypeDuration) SetWednesday(v bool) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *BestAvailableRatesTypeDuration) GetThursday() bool {
	if o == nil || IsNil(o.Thursday) {
		var ret bool
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesTypeDuration) GetThursdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Thursday) {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *BestAvailableRatesTypeDuration) HasThursday() bool {
	if o != nil && !IsNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given bool and assigns it to the Thursday field.
func (o *BestAvailableRatesTypeDuration) SetThursday(v bool) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *BestAvailableRatesTypeDuration) GetFriday() bool {
	if o == nil || IsNil(o.Friday) {
		var ret bool
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesTypeDuration) GetFridayOk() (*bool, bool) {
	if o == nil || IsNil(o.Friday) {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *BestAvailableRatesTypeDuration) HasFriday() bool {
	if o != nil && !IsNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given bool and assigns it to the Friday field.
func (o *BestAvailableRatesTypeDuration) SetFriday(v bool) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *BestAvailableRatesTypeDuration) GetSaturday() bool {
	if o == nil || IsNil(o.Saturday) {
		var ret bool
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatesTypeDuration) GetSaturdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Saturday) {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *BestAvailableRatesTypeDuration) HasSaturday() bool {
	if o != nil && !IsNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given bool and assigns it to the Saturday field.
func (o *BestAvailableRatesTypeDuration) SetSaturday(v bool) {
	o.Saturday = &v
}

func (o BestAvailableRatesTypeDuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BestAvailableRatesTypeDuration) ToMap() (map[string]interface{}, error) {
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

type NullableBestAvailableRatesTypeDuration struct {
	value *BestAvailableRatesTypeDuration
	isSet bool
}

func (v NullableBestAvailableRatesTypeDuration) Get() *BestAvailableRatesTypeDuration {
	return v.value
}

func (v *NullableBestAvailableRatesTypeDuration) Set(val *BestAvailableRatesTypeDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableBestAvailableRatesTypeDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableBestAvailableRatesTypeDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBestAvailableRatesTypeDuration(val *BestAvailableRatesTypeDuration) *NullableBestAvailableRatesTypeDuration {
	return &NullableBestAvailableRatesTypeDuration{value: val, isSet: true}
}

func (v NullableBestAvailableRatesTypeDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBestAvailableRatesTypeDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


