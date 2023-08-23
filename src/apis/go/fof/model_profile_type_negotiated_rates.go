/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProfileTypeNegotiatedRates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileTypeNegotiatedRates{}

// ProfileTypeNegotiatedRates List of profile negotiated rates.
type ProfileTypeNegotiatedRates struct {
	// Collection of Detailed information on profile negotiated rates.
	NegotiatedRate []NegotiatedType `json:"negotiatedRate,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewProfileTypeNegotiatedRates instantiates a new ProfileTypeNegotiatedRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileTypeNegotiatedRates() *ProfileTypeNegotiatedRates {
	this := ProfileTypeNegotiatedRates{}
	return &this
}

// NewProfileTypeNegotiatedRatesWithDefaults instantiates a new ProfileTypeNegotiatedRates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTypeNegotiatedRatesWithDefaults() *ProfileTypeNegotiatedRates {
	this := ProfileTypeNegotiatedRates{}
	return &this
}

// GetNegotiatedRate returns the NegotiatedRate field value if set, zero value otherwise.
func (o *ProfileTypeNegotiatedRates) GetNegotiatedRate() []NegotiatedType {
	if o == nil || IsNil(o.NegotiatedRate) {
		var ret []NegotiatedType
		return ret
	}
	return o.NegotiatedRate
}

// GetNegotiatedRateOk returns a tuple with the NegotiatedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeNegotiatedRates) GetNegotiatedRateOk() ([]NegotiatedType, bool) {
	if o == nil || IsNil(o.NegotiatedRate) {
		return nil, false
	}
	return o.NegotiatedRate, true
}

// HasNegotiatedRate returns a boolean if a field has been set.
func (o *ProfileTypeNegotiatedRates) HasNegotiatedRate() bool {
	if o != nil && !IsNil(o.NegotiatedRate) {
		return true
	}

	return false
}

// SetNegotiatedRate gets a reference to the given []NegotiatedType and assigns it to the NegotiatedRate field.
func (o *ProfileTypeNegotiatedRates) SetNegotiatedRate(v []NegotiatedType) {
	o.NegotiatedRate = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ProfileTypeNegotiatedRates) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeNegotiatedRates) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ProfileTypeNegotiatedRates) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ProfileTypeNegotiatedRates) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ProfileTypeNegotiatedRates) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeNegotiatedRates) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ProfileTypeNegotiatedRates) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ProfileTypeNegotiatedRates) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProfileTypeNegotiatedRates) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeNegotiatedRates) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProfileTypeNegotiatedRates) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProfileTypeNegotiatedRates) SetCount(v int32) {
	o.Count = &v
}

func (o ProfileTypeNegotiatedRates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileTypeNegotiatedRates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegotiatedRate) {
		toSerialize["negotiatedRate"] = o.NegotiatedRate
	}
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableProfileTypeNegotiatedRates struct {
	value *ProfileTypeNegotiatedRates
	isSet bool
}

func (v NullableProfileTypeNegotiatedRates) Get() *ProfileTypeNegotiatedRates {
	return v.value
}

func (v *NullableProfileTypeNegotiatedRates) Set(val *ProfileTypeNegotiatedRates) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTypeNegotiatedRates) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTypeNegotiatedRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTypeNegotiatedRates(val *ProfileTypeNegotiatedRates) *NullableProfileTypeNegotiatedRates {
	return &NullableProfileTypeNegotiatedRates{value: val, isSet: true}
}

func (v NullableProfileTypeNegotiatedRates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTypeNegotiatedRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


