/*
OPERA Cloud Inventory API

APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inv

import (
	"encoding/json"
)

// checks if the SellMessagesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SellMessagesType{}

// SellMessagesType The SellMessagesType is the list of message and attributes returned by web service.
type SellMessagesType struct {
	// This is the message text.
	SellMessage []SellMessageType `json:"sellMessage,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewSellMessagesType instantiates a new SellMessagesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellMessagesType() *SellMessagesType {
	this := SellMessagesType{}
	return &this
}

// NewSellMessagesTypeWithDefaults instantiates a new SellMessagesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellMessagesTypeWithDefaults() *SellMessagesType {
	this := SellMessagesType{}
	return &this
}

// GetSellMessage returns the SellMessage field value if set, zero value otherwise.
func (o *SellMessagesType) GetSellMessage() []SellMessageType {
	if o == nil || IsNil(o.SellMessage) {
		var ret []SellMessageType
		return ret
	}
	return o.SellMessage
}

// GetSellMessageOk returns a tuple with the SellMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellMessagesType) GetSellMessageOk() ([]SellMessageType, bool) {
	if o == nil || IsNil(o.SellMessage) {
		return nil, false
	}
	return o.SellMessage, true
}

// HasSellMessage returns a boolean if a field has been set.
func (o *SellMessagesType) HasSellMessage() bool {
	if o != nil && !IsNil(o.SellMessage) {
		return true
	}

	return false
}

// SetSellMessage gets a reference to the given []SellMessageType and assigns it to the SellMessage field.
func (o *SellMessagesType) SetSellMessage(v []SellMessageType) {
	o.SellMessage = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *SellMessagesType) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellMessagesType) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *SellMessagesType) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *SellMessagesType) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SellMessagesType) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellMessagesType) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SellMessagesType) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *SellMessagesType) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SellMessagesType) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellMessagesType) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SellMessagesType) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SellMessagesType) SetCount(v int32) {
	o.Count = &v
}

func (o SellMessagesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SellMessagesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SellMessage) {
		toSerialize["sellMessage"] = o.SellMessage
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

type NullableSellMessagesType struct {
	value *SellMessagesType
	isSet bool
}

func (v NullableSellMessagesType) Get() *SellMessagesType {
	return v.value
}

func (v *NullableSellMessagesType) Set(val *SellMessagesType) {
	v.value = val
	v.isSet = true
}

func (v NullableSellMessagesType) IsSet() bool {
	return v.isSet
}

func (v *NullableSellMessagesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellMessagesType(val *SellMessagesType) *NullableSellMessagesType {
	return &NullableSellMessagesType{value: val, isSet: true}
}

func (v NullableSellMessagesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSellMessagesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


