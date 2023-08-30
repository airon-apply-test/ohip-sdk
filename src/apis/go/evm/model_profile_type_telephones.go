/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
)

// checks if the ProfileTypeTelephones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileTypeTelephones{}

// ProfileTypeTelephones List of Telephone Number Information
type ProfileTypeTelephones struct {
	// Collection of Detailed information on telephone/fax for the customer.
	TelephoneInfo []TelephoneInfoType `json:"telephoneInfo,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewProfileTypeTelephones instantiates a new ProfileTypeTelephones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileTypeTelephones() *ProfileTypeTelephones {
	this := ProfileTypeTelephones{}
	return &this
}

// NewProfileTypeTelephonesWithDefaults instantiates a new ProfileTypeTelephones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTypeTelephonesWithDefaults() *ProfileTypeTelephones {
	this := ProfileTypeTelephones{}
	return &this
}

// GetTelephoneInfo returns the TelephoneInfo field value if set, zero value otherwise.
func (o *ProfileTypeTelephones) GetTelephoneInfo() []TelephoneInfoType {
	if o == nil || IsNil(o.TelephoneInfo) {
		var ret []TelephoneInfoType
		return ret
	}
	return o.TelephoneInfo
}

// GetTelephoneInfoOk returns a tuple with the TelephoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeTelephones) GetTelephoneInfoOk() ([]TelephoneInfoType, bool) {
	if o == nil || IsNil(o.TelephoneInfo) {
		return nil, false
	}
	return o.TelephoneInfo, true
}

// HasTelephoneInfo returns a boolean if a field has been set.
func (o *ProfileTypeTelephones) HasTelephoneInfo() bool {
	if o != nil && !IsNil(o.TelephoneInfo) {
		return true
	}

	return false
}

// SetTelephoneInfo gets a reference to the given []TelephoneInfoType and assigns it to the TelephoneInfo field.
func (o *ProfileTypeTelephones) SetTelephoneInfo(v []TelephoneInfoType) {
	o.TelephoneInfo = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ProfileTypeTelephones) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeTelephones) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ProfileTypeTelephones) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ProfileTypeTelephones) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ProfileTypeTelephones) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeTelephones) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ProfileTypeTelephones) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ProfileTypeTelephones) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProfileTypeTelephones) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeTelephones) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProfileTypeTelephones) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProfileTypeTelephones) SetCount(v int32) {
	o.Count = &v
}

func (o ProfileTypeTelephones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileTypeTelephones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TelephoneInfo) {
		toSerialize["telephoneInfo"] = o.TelephoneInfo
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

type NullableProfileTypeTelephones struct {
	value *ProfileTypeTelephones
	isSet bool
}

func (v NullableProfileTypeTelephones) Get() *ProfileTypeTelephones {
	return v.value
}

func (v *NullableProfileTypeTelephones) Set(val *ProfileTypeTelephones) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTypeTelephones) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTypeTelephones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTypeTelephones(val *ProfileTypeTelephones) *NullableProfileTypeTelephones {
	return &NullableProfileTypeTelephones{value: val, isSet: true}
}

func (v NullableProfileTypeTelephones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTypeTelephones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


