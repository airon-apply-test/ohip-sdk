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
)

// checks if the ResCommunicationTypeTelephones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResCommunicationTypeTelephones{}

// ResCommunicationTypeTelephones List of Telephone Number Information
type ResCommunicationTypeTelephones struct {
	// Collection of Detailed information on telephone/fax for the customer.
	TelephoneInfo []TelephoneInfoType `json:"telephoneInfo,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewResCommunicationTypeTelephones instantiates a new ResCommunicationTypeTelephones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResCommunicationTypeTelephones() *ResCommunicationTypeTelephones {
	this := ResCommunicationTypeTelephones{}
	return &this
}

// NewResCommunicationTypeTelephonesWithDefaults instantiates a new ResCommunicationTypeTelephones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResCommunicationTypeTelephonesWithDefaults() *ResCommunicationTypeTelephones {
	this := ResCommunicationTypeTelephones{}
	return &this
}

// GetTelephoneInfo returns the TelephoneInfo field value if set, zero value otherwise.
func (o *ResCommunicationTypeTelephones) GetTelephoneInfo() []TelephoneInfoType {
	if o == nil || IsNil(o.TelephoneInfo) {
		var ret []TelephoneInfoType
		return ret
	}
	return o.TelephoneInfo
}

// GetTelephoneInfoOk returns a tuple with the TelephoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeTelephones) GetTelephoneInfoOk() ([]TelephoneInfoType, bool) {
	if o == nil || IsNil(o.TelephoneInfo) {
		return nil, false
	}
	return o.TelephoneInfo, true
}

// HasTelephoneInfo returns a boolean if a field has been set.
func (o *ResCommunicationTypeTelephones) HasTelephoneInfo() bool {
	if o != nil && !IsNil(o.TelephoneInfo) {
		return true
	}

	return false
}

// SetTelephoneInfo gets a reference to the given []TelephoneInfoType and assigns it to the TelephoneInfo field.
func (o *ResCommunicationTypeTelephones) SetTelephoneInfo(v []TelephoneInfoType) {
	o.TelephoneInfo = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ResCommunicationTypeTelephones) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeTelephones) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ResCommunicationTypeTelephones) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ResCommunicationTypeTelephones) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ResCommunicationTypeTelephones) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeTelephones) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ResCommunicationTypeTelephones) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ResCommunicationTypeTelephones) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ResCommunicationTypeTelephones) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeTelephones) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ResCommunicationTypeTelephones) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ResCommunicationTypeTelephones) SetCount(v int32) {
	o.Count = &v
}

func (o ResCommunicationTypeTelephones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResCommunicationTypeTelephones) ToMap() (map[string]interface{}, error) {
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

type NullableResCommunicationTypeTelephones struct {
	value *ResCommunicationTypeTelephones
	isSet bool
}

func (v NullableResCommunicationTypeTelephones) Get() *ResCommunicationTypeTelephones {
	return v.value
}

func (v *NullableResCommunicationTypeTelephones) Set(val *ResCommunicationTypeTelephones) {
	v.value = val
	v.isSet = true
}

func (v NullableResCommunicationTypeTelephones) IsSet() bool {
	return v.isSet
}

func (v *NullableResCommunicationTypeTelephones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResCommunicationTypeTelephones(val *ResCommunicationTypeTelephones) *NullableResCommunicationTypeTelephones {
	return &NullableResCommunicationTypeTelephones{value: val, isSet: true}
}

func (v NullableResCommunicationTypeTelephones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResCommunicationTypeTelephones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


