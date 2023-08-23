/*
OPERA Cloud Room Configuration API

APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RoomTypeTemplatesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomTypeTemplatesDetails{}

// RoomTypeTemplatesDetails Response object for information regarding room type template of a property.
type RoomTypeTemplatesDetails struct {
	// This type holds collection of room type.
	RoomTypeTemplatesSummary []RoomTypeSummaryType `json:"roomTypeTemplatesSummary,omitempty"`
	// This type holds collection of room type.
	RoomTypeTemplates []RoomTypeType `json:"roomTypeTemplates,omitempty"`
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

// NewRoomTypeTemplatesDetails instantiates a new RoomTypeTemplatesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomTypeTemplatesDetails() *RoomTypeTemplatesDetails {
	this := RoomTypeTemplatesDetails{}
	return &this
}

// NewRoomTypeTemplatesDetailsWithDefaults instantiates a new RoomTypeTemplatesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomTypeTemplatesDetailsWithDefaults() *RoomTypeTemplatesDetails {
	this := RoomTypeTemplatesDetails{}
	return &this
}

// GetRoomTypeTemplatesSummary returns the RoomTypeTemplatesSummary field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetRoomTypeTemplatesSummary() []RoomTypeSummaryType {
	if o == nil || IsNil(o.RoomTypeTemplatesSummary) {
		var ret []RoomTypeSummaryType
		return ret
	}
	return o.RoomTypeTemplatesSummary
}

// GetRoomTypeTemplatesSummaryOk returns a tuple with the RoomTypeTemplatesSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetRoomTypeTemplatesSummaryOk() ([]RoomTypeSummaryType, bool) {
	if o == nil || IsNil(o.RoomTypeTemplatesSummary) {
		return nil, false
	}
	return o.RoomTypeTemplatesSummary, true
}

// HasRoomTypeTemplatesSummary returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasRoomTypeTemplatesSummary() bool {
	if o != nil && !IsNil(o.RoomTypeTemplatesSummary) {
		return true
	}

	return false
}

// SetRoomTypeTemplatesSummary gets a reference to the given []RoomTypeSummaryType and assigns it to the RoomTypeTemplatesSummary field.
func (o *RoomTypeTemplatesDetails) SetRoomTypeTemplatesSummary(v []RoomTypeSummaryType) {
	o.RoomTypeTemplatesSummary = v
}

// GetRoomTypeTemplates returns the RoomTypeTemplates field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetRoomTypeTemplates() []RoomTypeType {
	if o == nil || IsNil(o.RoomTypeTemplates) {
		var ret []RoomTypeType
		return ret
	}
	return o.RoomTypeTemplates
}

// GetRoomTypeTemplatesOk returns a tuple with the RoomTypeTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetRoomTypeTemplatesOk() ([]RoomTypeType, bool) {
	if o == nil || IsNil(o.RoomTypeTemplates) {
		return nil, false
	}
	return o.RoomTypeTemplates, true
}

// HasRoomTypeTemplates returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasRoomTypeTemplates() bool {
	if o != nil && !IsNil(o.RoomTypeTemplates) {
		return true
	}

	return false
}

// SetRoomTypeTemplates gets a reference to the given []RoomTypeType and assigns it to the RoomTypeTemplates field.
func (o *RoomTypeTemplatesDetails) SetRoomTypeTemplates(v []RoomTypeType) {
	o.RoomTypeTemplates = v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *RoomTypeTemplatesDetails) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *RoomTypeTemplatesDetails) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *RoomTypeTemplatesDetails) SetLimit(v int32) {
	o.Limit = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *RoomTypeTemplatesDetails) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *RoomTypeTemplatesDetails) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *RoomTypeTemplatesDetails) SetCount(v int32) {
	o.Count = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *RoomTypeTemplatesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RoomTypeTemplatesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTypeTemplatesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RoomTypeTemplatesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *RoomTypeTemplatesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o RoomTypeTemplatesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomTypeTemplatesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomTypeTemplatesSummary) {
		toSerialize["roomTypeTemplatesSummary"] = o.RoomTypeTemplatesSummary
	}
	if !IsNil(o.RoomTypeTemplates) {
		toSerialize["roomTypeTemplates"] = o.RoomTypeTemplates
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

type NullableRoomTypeTemplatesDetails struct {
	value *RoomTypeTemplatesDetails
	isSet bool
}

func (v NullableRoomTypeTemplatesDetails) Get() *RoomTypeTemplatesDetails {
	return v.value
}

func (v *NullableRoomTypeTemplatesDetails) Set(val *RoomTypeTemplatesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTypeTemplatesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTypeTemplatesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTypeTemplatesDetails(val *RoomTypeTemplatesDetails) *NullableRoomTypeTemplatesDetails {
	return &NullableRoomTypeTemplatesDetails{value: val, isSet: true}
}

func (v NullableRoomTypeTemplatesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTypeTemplatesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


