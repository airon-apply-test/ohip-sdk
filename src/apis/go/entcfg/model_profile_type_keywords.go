/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProfileTypeKeywords type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileTypeKeywords{}

// ProfileTypeKeywords List of customer keywords.
type ProfileTypeKeywords struct {
	// Collection of keywords attached to the profile.
	Keyword []KeywordType `json:"keyword,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewProfileTypeKeywords instantiates a new ProfileTypeKeywords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileTypeKeywords() *ProfileTypeKeywords {
	this := ProfileTypeKeywords{}
	return &this
}

// NewProfileTypeKeywordsWithDefaults instantiates a new ProfileTypeKeywords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTypeKeywordsWithDefaults() *ProfileTypeKeywords {
	this := ProfileTypeKeywords{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *ProfileTypeKeywords) GetKeyword() []KeywordType {
	if o == nil || IsNil(o.Keyword) {
		var ret []KeywordType
		return ret
	}
	return o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeKeywords) GetKeywordOk() ([]KeywordType, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *ProfileTypeKeywords) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given []KeywordType and assigns it to the Keyword field.
func (o *ProfileTypeKeywords) SetKeyword(v []KeywordType) {
	o.Keyword = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ProfileTypeKeywords) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeKeywords) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ProfileTypeKeywords) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ProfileTypeKeywords) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ProfileTypeKeywords) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeKeywords) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ProfileTypeKeywords) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ProfileTypeKeywords) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProfileTypeKeywords) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeKeywords) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProfileTypeKeywords) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProfileTypeKeywords) SetCount(v int32) {
	o.Count = &v
}

func (o ProfileTypeKeywords) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileTypeKeywords) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
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

type NullableProfileTypeKeywords struct {
	value *ProfileTypeKeywords
	isSet bool
}

func (v NullableProfileTypeKeywords) Get() *ProfileTypeKeywords {
	return v.value
}

func (v *NullableProfileTypeKeywords) Set(val *ProfileTypeKeywords) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTypeKeywords) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTypeKeywords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTypeKeywords(val *ProfileTypeKeywords) *NullableProfileTypeKeywords {
	return &NullableProfileTypeKeywords{value: val, isSet: true}
}

func (v NullableProfileTypeKeywords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTypeKeywords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


