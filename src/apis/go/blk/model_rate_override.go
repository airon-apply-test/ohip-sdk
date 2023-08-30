/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the RateOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateOverride{}

// RateOverride Request object of change block rate override operation.
type RateOverride struct {
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	HotelId *string `json:"hotelId,omitempty"`
	// Unique Id that references an object uniquely in the system.
	BlockIdList []UniqueIDType `json:"blockIdList,omitempty"`
	// Indicates whether rates of a block can be overridden. Applicable only for blocks with a Rate Code.
	AllowRateOverride *bool `json:"allowRateOverride,omitempty"`
	RateOverrideReason *CodeDescriptionType `json:"rateOverrideReason,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewRateOverride instantiates a new RateOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateOverride() *RateOverride {
	this := RateOverride{}
	return &this
}

// NewRateOverrideWithDefaults instantiates a new RateOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateOverrideWithDefaults() *RateOverride {
	this := RateOverride{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *RateOverride) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateOverride) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *RateOverride) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *RateOverride) SetHotelId(v string) {
	o.HotelId = &v
}

// GetBlockIdList returns the BlockIdList field value if set, zero value otherwise.
func (o *RateOverride) GetBlockIdList() []UniqueIDType {
	if o == nil || IsNil(o.BlockIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.BlockIdList
}

// GetBlockIdListOk returns a tuple with the BlockIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateOverride) GetBlockIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.BlockIdList) {
		return nil, false
	}
	return o.BlockIdList, true
}

// HasBlockIdList returns a boolean if a field has been set.
func (o *RateOverride) HasBlockIdList() bool {
	if o != nil && !IsNil(o.BlockIdList) {
		return true
	}

	return false
}

// SetBlockIdList gets a reference to the given []UniqueIDType and assigns it to the BlockIdList field.
func (o *RateOverride) SetBlockIdList(v []UniqueIDType) {
	o.BlockIdList = v
}

// GetAllowRateOverride returns the AllowRateOverride field value if set, zero value otherwise.
func (o *RateOverride) GetAllowRateOverride() bool {
	if o == nil || IsNil(o.AllowRateOverride) {
		var ret bool
		return ret
	}
	return *o.AllowRateOverride
}

// GetAllowRateOverrideOk returns a tuple with the AllowRateOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateOverride) GetAllowRateOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRateOverride) {
		return nil, false
	}
	return o.AllowRateOverride, true
}

// HasAllowRateOverride returns a boolean if a field has been set.
func (o *RateOverride) HasAllowRateOverride() bool {
	if o != nil && !IsNil(o.AllowRateOverride) {
		return true
	}

	return false
}

// SetAllowRateOverride gets a reference to the given bool and assigns it to the AllowRateOverride field.
func (o *RateOverride) SetAllowRateOverride(v bool) {
	o.AllowRateOverride = &v
}

// GetRateOverrideReason returns the RateOverrideReason field value if set, zero value otherwise.
func (o *RateOverride) GetRateOverrideReason() CodeDescriptionType {
	if o == nil || IsNil(o.RateOverrideReason) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.RateOverrideReason
}

// GetRateOverrideReasonOk returns a tuple with the RateOverrideReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateOverride) GetRateOverrideReasonOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.RateOverrideReason) {
		return nil, false
	}
	return o.RateOverrideReason, true
}

// HasRateOverrideReason returns a boolean if a field has been set.
func (o *RateOverride) HasRateOverrideReason() bool {
	if o != nil && !IsNil(o.RateOverrideReason) {
		return true
	}

	return false
}

// SetRateOverrideReason gets a reference to the given CodeDescriptionType and assigns it to the RateOverrideReason field.
func (o *RateOverride) SetRateOverrideReason(v CodeDescriptionType) {
	o.RateOverrideReason = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RateOverride) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateOverride) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RateOverride) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *RateOverride) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RateOverride) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateOverride) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RateOverride) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *RateOverride) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o RateOverride) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.BlockIdList) {
		toSerialize["blockIdList"] = o.BlockIdList
	}
	if !IsNil(o.AllowRateOverride) {
		toSerialize["allowRateOverride"] = o.AllowRateOverride
	}
	if !IsNil(o.RateOverrideReason) {
		toSerialize["rateOverrideReason"] = o.RateOverrideReason
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableRateOverride struct {
	value *RateOverride
	isSet bool
}

func (v NullableRateOverride) Get() *RateOverride {
	return v.value
}

func (v *NullableRateOverride) Set(val *RateOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableRateOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableRateOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateOverride(val *RateOverride) *NullableRateOverride {
	return &NullableRateOverride{value: val, isSet: true}
}

func (v NullableRateOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


