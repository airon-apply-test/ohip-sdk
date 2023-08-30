/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the CommunicationMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationMethods{}

// CommunicationMethods struct for CommunicationMethods
type CommunicationMethods struct {
	// 
	CommunicationDetails []CommunicationMethodType `json:"communicationDetails,omitempty"`
	// Indicates the index of the next applicable set(page).
	Offset *int32 `json:"offset,omitempty"`
	// Indicates number of records the API can return as per the API request limit sent. A maximum of 200 records can be only returned at a time.
	Limit *int32 `json:"limit,omitempty"`
	// Indicates number of records the API has returned actually as per the API request criteria.
	Count *int32 `json:"count,omitempty"`
	// Indicates whether there are more records available to be returned as per the API request criteria or not.
	HasMore *bool `json:"hasMore,omitempty"`
	// Indicates total number of records available that can be returned as per the API request criteria.
	TotalResults *int32 `json:"totalResults,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCommunicationMethods instantiates a new CommunicationMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationMethods() *CommunicationMethods {
	this := CommunicationMethods{}
	return &this
}

// NewCommunicationMethodsWithDefaults instantiates a new CommunicationMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationMethodsWithDefaults() *CommunicationMethods {
	this := CommunicationMethods{}
	return &this
}

// GetCommunicationDetails returns the CommunicationDetails field value if set, zero value otherwise.
func (o *CommunicationMethods) GetCommunicationDetails() []CommunicationMethodType {
	if o == nil || IsNil(o.CommunicationDetails) {
		var ret []CommunicationMethodType
		return ret
	}
	return o.CommunicationDetails
}

// GetCommunicationDetailsOk returns a tuple with the CommunicationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetCommunicationDetailsOk() ([]CommunicationMethodType, bool) {
	if o == nil || IsNil(o.CommunicationDetails) {
		return nil, false
	}
	return o.CommunicationDetails, true
}

// HasCommunicationDetails returns a boolean if a field has been set.
func (o *CommunicationMethods) HasCommunicationDetails() bool {
	if o != nil && !IsNil(o.CommunicationDetails) {
		return true
	}

	return false
}

// SetCommunicationDetails gets a reference to the given []CommunicationMethodType and assigns it to the CommunicationDetails field.
func (o *CommunicationMethods) SetCommunicationDetails(v []CommunicationMethodType) {
	o.CommunicationDetails = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CommunicationMethods) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CommunicationMethods) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CommunicationMethods) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CommunicationMethods) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CommunicationMethods) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CommunicationMethods) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CommunicationMethods) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CommunicationMethods) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CommunicationMethods) SetCount(v int32) {
	o.Count = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *CommunicationMethods) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *CommunicationMethods) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *CommunicationMethods) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *CommunicationMethods) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *CommunicationMethods) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *CommunicationMethods) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CommunicationMethods) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CommunicationMethods) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CommunicationMethods) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CommunicationMethods) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethods) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CommunicationMethods) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CommunicationMethods) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CommunicationMethods) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommunicationDetails) {
		toSerialize["communicationDetails"] = o.CommunicationDetails
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCommunicationMethods struct {
	value *CommunicationMethods
	isSet bool
}

func (v NullableCommunicationMethods) Get() *CommunicationMethods {
	return v.value
}

func (v *NullableCommunicationMethods) Set(val *CommunicationMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationMethods(val *CommunicationMethods) *NullableCommunicationMethods {
	return &NullableCommunicationMethods{value: val, isSet: true}
}

func (v NullableCommunicationMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


