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

// checks if the SubAllocationTypeComments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubAllocationTypeComments{}

// SubAllocationTypeComments List of notes for the Block.
type SubAllocationTypeComments struct {
	CommentInfo []CommentInfoType `json:"commentInfo,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewSubAllocationTypeComments instantiates a new SubAllocationTypeComments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAllocationTypeComments() *SubAllocationTypeComments {
	this := SubAllocationTypeComments{}
	return &this
}

// NewSubAllocationTypeCommentsWithDefaults instantiates a new SubAllocationTypeComments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAllocationTypeCommentsWithDefaults() *SubAllocationTypeComments {
	this := SubAllocationTypeComments{}
	return &this
}

// GetCommentInfo returns the CommentInfo field value if set, zero value otherwise.
func (o *SubAllocationTypeComments) GetCommentInfo() []CommentInfoType {
	if o == nil || IsNil(o.CommentInfo) {
		var ret []CommentInfoType
		return ret
	}
	return o.CommentInfo
}

// GetCommentInfoOk returns a tuple with the CommentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeComments) GetCommentInfoOk() ([]CommentInfoType, bool) {
	if o == nil || IsNil(o.CommentInfo) {
		return nil, false
	}
	return o.CommentInfo, true
}

// HasCommentInfo returns a boolean if a field has been set.
func (o *SubAllocationTypeComments) HasCommentInfo() bool {
	if o != nil && !IsNil(o.CommentInfo) {
		return true
	}

	return false
}

// SetCommentInfo gets a reference to the given []CommentInfoType and assigns it to the CommentInfo field.
func (o *SubAllocationTypeComments) SetCommentInfo(v []CommentInfoType) {
	o.CommentInfo = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *SubAllocationTypeComments) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeComments) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *SubAllocationTypeComments) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *SubAllocationTypeComments) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SubAllocationTypeComments) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeComments) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SubAllocationTypeComments) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *SubAllocationTypeComments) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SubAllocationTypeComments) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAllocationTypeComments) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SubAllocationTypeComments) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SubAllocationTypeComments) SetCount(v int32) {
	o.Count = &v
}

func (o SubAllocationTypeComments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAllocationTypeComments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommentInfo) {
		toSerialize["commentInfo"] = o.CommentInfo
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

type NullableSubAllocationTypeComments struct {
	value *SubAllocationTypeComments
	isSet bool
}

func (v NullableSubAllocationTypeComments) Get() *SubAllocationTypeComments {
	return v.value
}

func (v *NullableSubAllocationTypeComments) Set(val *SubAllocationTypeComments) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAllocationTypeComments) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAllocationTypeComments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAllocationTypeComments(val *SubAllocationTypeComments) *NullableSubAllocationTypeComments {
	return &NullableSubAllocationTypeComments{value: val, isSet: true}
}

func (v NullableSubAllocationTypeComments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAllocationTypeComments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


