/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtp

import (
	"encoding/json"
)

// checks if the RatePlanClassificationsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatePlanClassificationsType{}

// RatePlanClassificationsType struct for RatePlanClassificationsType
type RatePlanClassificationsType struct {
	// Rate Category for the rate plan.
	RateCategory *string `json:"rateCategory,omitempty"`
	// Display set for the rate plan.
	DisplaySet *string `json:"displaySet,omitempty"`
	// Market code for the rate plan.
	MarketCode *string `json:"marketCode,omitempty"`
	// Source code for the rate plan.
	SourceCode *string `json:"sourceCode,omitempty"`
	// Rate group for the rate plan.
	RateGroup *string `json:"rateGroup,omitempty"`
}

// NewRatePlanClassificationsType instantiates a new RatePlanClassificationsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatePlanClassificationsType() *RatePlanClassificationsType {
	this := RatePlanClassificationsType{}
	return &this
}

// NewRatePlanClassificationsTypeWithDefaults instantiates a new RatePlanClassificationsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatePlanClassificationsTypeWithDefaults() *RatePlanClassificationsType {
	this := RatePlanClassificationsType{}
	return &this
}

// GetRateCategory returns the RateCategory field value if set, zero value otherwise.
func (o *RatePlanClassificationsType) GetRateCategory() string {
	if o == nil || IsNil(o.RateCategory) {
		var ret string
		return ret
	}
	return *o.RateCategory
}

// GetRateCategoryOk returns a tuple with the RateCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanClassificationsType) GetRateCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.RateCategory) {
		return nil, false
	}
	return o.RateCategory, true
}

// HasRateCategory returns a boolean if a field has been set.
func (o *RatePlanClassificationsType) HasRateCategory() bool {
	if o != nil && !IsNil(o.RateCategory) {
		return true
	}

	return false
}

// SetRateCategory gets a reference to the given string and assigns it to the RateCategory field.
func (o *RatePlanClassificationsType) SetRateCategory(v string) {
	o.RateCategory = &v
}

// GetDisplaySet returns the DisplaySet field value if set, zero value otherwise.
func (o *RatePlanClassificationsType) GetDisplaySet() string {
	if o == nil || IsNil(o.DisplaySet) {
		var ret string
		return ret
	}
	return *o.DisplaySet
}

// GetDisplaySetOk returns a tuple with the DisplaySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanClassificationsType) GetDisplaySetOk() (*string, bool) {
	if o == nil || IsNil(o.DisplaySet) {
		return nil, false
	}
	return o.DisplaySet, true
}

// HasDisplaySet returns a boolean if a field has been set.
func (o *RatePlanClassificationsType) HasDisplaySet() bool {
	if o != nil && !IsNil(o.DisplaySet) {
		return true
	}

	return false
}

// SetDisplaySet gets a reference to the given string and assigns it to the DisplaySet field.
func (o *RatePlanClassificationsType) SetDisplaySet(v string) {
	o.DisplaySet = &v
}

// GetMarketCode returns the MarketCode field value if set, zero value otherwise.
func (o *RatePlanClassificationsType) GetMarketCode() string {
	if o == nil || IsNil(o.MarketCode) {
		var ret string
		return ret
	}
	return *o.MarketCode
}

// GetMarketCodeOk returns a tuple with the MarketCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanClassificationsType) GetMarketCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MarketCode) {
		return nil, false
	}
	return o.MarketCode, true
}

// HasMarketCode returns a boolean if a field has been set.
func (o *RatePlanClassificationsType) HasMarketCode() bool {
	if o != nil && !IsNil(o.MarketCode) {
		return true
	}

	return false
}

// SetMarketCode gets a reference to the given string and assigns it to the MarketCode field.
func (o *RatePlanClassificationsType) SetMarketCode(v string) {
	o.MarketCode = &v
}

// GetSourceCode returns the SourceCode field value if set, zero value otherwise.
func (o *RatePlanClassificationsType) GetSourceCode() string {
	if o == nil || IsNil(o.SourceCode) {
		var ret string
		return ret
	}
	return *o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanClassificationsType) GetSourceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCode) {
		return nil, false
	}
	return o.SourceCode, true
}

// HasSourceCode returns a boolean if a field has been set.
func (o *RatePlanClassificationsType) HasSourceCode() bool {
	if o != nil && !IsNil(o.SourceCode) {
		return true
	}

	return false
}

// SetSourceCode gets a reference to the given string and assigns it to the SourceCode field.
func (o *RatePlanClassificationsType) SetSourceCode(v string) {
	o.SourceCode = &v
}

// GetRateGroup returns the RateGroup field value if set, zero value otherwise.
func (o *RatePlanClassificationsType) GetRateGroup() string {
	if o == nil || IsNil(o.RateGroup) {
		var ret string
		return ret
	}
	return *o.RateGroup
}

// GetRateGroupOk returns a tuple with the RateGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanClassificationsType) GetRateGroupOk() (*string, bool) {
	if o == nil || IsNil(o.RateGroup) {
		return nil, false
	}
	return o.RateGroup, true
}

// HasRateGroup returns a boolean if a field has been set.
func (o *RatePlanClassificationsType) HasRateGroup() bool {
	if o != nil && !IsNil(o.RateGroup) {
		return true
	}

	return false
}

// SetRateGroup gets a reference to the given string and assigns it to the RateGroup field.
func (o *RatePlanClassificationsType) SetRateGroup(v string) {
	o.RateGroup = &v
}

func (o RatePlanClassificationsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatePlanClassificationsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RateCategory) {
		toSerialize["rateCategory"] = o.RateCategory
	}
	if !IsNil(o.DisplaySet) {
		toSerialize["displaySet"] = o.DisplaySet
	}
	if !IsNil(o.MarketCode) {
		toSerialize["marketCode"] = o.MarketCode
	}
	if !IsNil(o.SourceCode) {
		toSerialize["sourceCode"] = o.SourceCode
	}
	if !IsNil(o.RateGroup) {
		toSerialize["rateGroup"] = o.RateGroup
	}
	return toSerialize, nil
}

type NullableRatePlanClassificationsType struct {
	value *RatePlanClassificationsType
	isSet bool
}

func (v NullableRatePlanClassificationsType) Get() *RatePlanClassificationsType {
	return v.value
}

func (v *NullableRatePlanClassificationsType) Set(val *RatePlanClassificationsType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanClassificationsType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanClassificationsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanClassificationsType(val *RatePlanClassificationsType) *NullableRatePlanClassificationsType {
	return &NullableRatePlanClassificationsType{value: val, isSet: true}
}

func (v NullableRatePlanClassificationsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanClassificationsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


