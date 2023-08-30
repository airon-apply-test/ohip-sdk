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

// checks if the GetBestAvailableRatesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBestAvailableRatesDetails{}

// GetBestAvailableRatesDetails struct for GetBestAvailableRatesDetails
type GetBestAvailableRatesDetails struct {
	BestAvailableRatesList *BestAvailableRatesListType `json:"bestAvailableRatesList,omitempty"`
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
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewGetBestAvailableRatesDetails instantiates a new GetBestAvailableRatesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBestAvailableRatesDetails() *GetBestAvailableRatesDetails {
	this := GetBestAvailableRatesDetails{}
	return &this
}

// NewGetBestAvailableRatesDetailsWithDefaults instantiates a new GetBestAvailableRatesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBestAvailableRatesDetailsWithDefaults() *GetBestAvailableRatesDetails {
	this := GetBestAvailableRatesDetails{}
	return &this
}

// GetBestAvailableRatesList returns the BestAvailableRatesList field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetBestAvailableRatesList() BestAvailableRatesListType {
	if o == nil || IsNil(o.BestAvailableRatesList) {
		var ret BestAvailableRatesListType
		return ret
	}
	return *o.BestAvailableRatesList
}

// GetBestAvailableRatesListOk returns a tuple with the BestAvailableRatesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetBestAvailableRatesListOk() (*BestAvailableRatesListType, bool) {
	if o == nil || IsNil(o.BestAvailableRatesList) {
		return nil, false
	}
	return o.BestAvailableRatesList, true
}

// HasBestAvailableRatesList returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasBestAvailableRatesList() bool {
	if o != nil && !IsNil(o.BestAvailableRatesList) {
		return true
	}

	return false
}

// SetBestAvailableRatesList gets a reference to the given BestAvailableRatesListType and assigns it to the BestAvailableRatesList field.
func (o *GetBestAvailableRatesDetails) SetBestAvailableRatesList(v BestAvailableRatesListType) {
	o.BestAvailableRatesList = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *GetBestAvailableRatesDetails) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetBestAvailableRatesDetails) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetBestAvailableRatesDetails) SetLimit(v int32) {
	o.Limit = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *GetBestAvailableRatesDetails) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *GetBestAvailableRatesDetails) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetBestAvailableRatesDetails) SetCount(v int32) {
	o.Count = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *GetBestAvailableRatesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBestAvailableRatesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *GetBestAvailableRatesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *GetBestAvailableRatesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o GetBestAvailableRatesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBestAvailableRatesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BestAvailableRatesList) {
		toSerialize["bestAvailableRatesList"] = o.BestAvailableRatesList
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
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableGetBestAvailableRatesDetails struct {
	value *GetBestAvailableRatesDetails
	isSet bool
}

func (v NullableGetBestAvailableRatesDetails) Get() *GetBestAvailableRatesDetails {
	return v.value
}

func (v *NullableGetBestAvailableRatesDetails) Set(val *GetBestAvailableRatesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBestAvailableRatesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBestAvailableRatesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBestAvailableRatesDetails(val *GetBestAvailableRatesDetails) *NullableGetBestAvailableRatesDetails {
	return &NullableGetBestAvailableRatesDetails{value: val, isSet: true}
}

func (v NullableGetBestAvailableRatesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBestAvailableRatesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


