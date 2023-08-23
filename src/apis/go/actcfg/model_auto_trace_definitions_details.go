/*
OPERA Cloud Activity Management API

APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AutoTraceDefinitionsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTraceDefinitionsDetails{}

// AutoTraceDefinitionsDetails Response object after fetching Auto Trace Definitions.
type AutoTraceDefinitionsDetails struct {
	// Auto Trace Definition.
	AutoTraceDefinitions []AutoTraceDefinitionType `json:"autoTraceDefinitions,omitempty"`
	// Evaluated total page count based on the requested max fetch count.
	TotalPages *int32 `json:"totalPages,omitempty"`
	// Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
	Offset *int32 `json:"offset,omitempty"`
	// Indicates maximum number of records a Web Service should return.
	Limit *int32 `json:"limit,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewAutoTraceDefinitionsDetails instantiates a new AutoTraceDefinitionsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTraceDefinitionsDetails() *AutoTraceDefinitionsDetails {
	this := AutoTraceDefinitionsDetails{}
	return &this
}

// NewAutoTraceDefinitionsDetailsWithDefaults instantiates a new AutoTraceDefinitionsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTraceDefinitionsDetailsWithDefaults() *AutoTraceDefinitionsDetails {
	this := AutoTraceDefinitionsDetails{}
	return &this
}

// GetAutoTraceDefinitions returns the AutoTraceDefinitions field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetAutoTraceDefinitions() []AutoTraceDefinitionType {
	if o == nil || IsNil(o.AutoTraceDefinitions) {
		var ret []AutoTraceDefinitionType
		return ret
	}
	return o.AutoTraceDefinitions
}

// GetAutoTraceDefinitionsOk returns a tuple with the AutoTraceDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetAutoTraceDefinitionsOk() ([]AutoTraceDefinitionType, bool) {
	if o == nil || IsNil(o.AutoTraceDefinitions) {
		return nil, false
	}
	return o.AutoTraceDefinitions, true
}

// HasAutoTraceDefinitions returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasAutoTraceDefinitions() bool {
	if o != nil && !IsNil(o.AutoTraceDefinitions) {
		return true
	}

	return false
}

// SetAutoTraceDefinitions gets a reference to the given []AutoTraceDefinitionType and assigns it to the AutoTraceDefinitions field.
func (o *AutoTraceDefinitionsDetails) SetAutoTraceDefinitions(v []AutoTraceDefinitionType) {
	o.AutoTraceDefinitions = v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *AutoTraceDefinitionsDetails) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *AutoTraceDefinitionsDetails) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *AutoTraceDefinitionsDetails) SetLimit(v int32) {
	o.Limit = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *AutoTraceDefinitionsDetails) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *AutoTraceDefinitionsDetails) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AutoTraceDefinitionsDetails) SetCount(v int32) {
	o.Count = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *AutoTraceDefinitionsDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AutoTraceDefinitionsDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTraceDefinitionsDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AutoTraceDefinitionsDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *AutoTraceDefinitionsDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o AutoTraceDefinitionsDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTraceDefinitionsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoTraceDefinitions) {
		toSerialize["autoTraceDefinitions"] = o.AutoTraceDefinitions
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
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
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAutoTraceDefinitionsDetails struct {
	value *AutoTraceDefinitionsDetails
	isSet bool
}

func (v NullableAutoTraceDefinitionsDetails) Get() *AutoTraceDefinitionsDetails {
	return v.value
}

func (v *NullableAutoTraceDefinitionsDetails) Set(val *AutoTraceDefinitionsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTraceDefinitionsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTraceDefinitionsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTraceDefinitionsDetails(val *AutoTraceDefinitionsDetails) *NullableAutoTraceDefinitionsDetails {
	return &NullableAutoTraceDefinitionsDetails{value: val, isSet: true}
}

func (v NullableAutoTraceDefinitionsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTraceDefinitionsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


