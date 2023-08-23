/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RatePackageCodeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatePackageCodeType{}

// RatePackageCodeType Package code details applied to a rate plan.
type RatePackageCodeType struct {
	Header *PackageCodeHeaderType `json:"header,omitempty"`
	// Package code price schedule details.
	Schedules []ConfigPackageScheduleType `json:"schedules,omitempty"`
	// Hotel code for the packages.
	HotelId *string `json:"hotelId,omitempty"`
	// Package Code.
	Code *string `json:"code,omitempty"`
	// Package Code Description.
	Description *string `json:"description,omitempty"`
	// Indicates if it is a Package Group or not.
	Group *bool `json:"group,omitempty"`
	// Package Code specific to a rate plan code.
	RatePlanCode *string `json:"ratePlanCode,omitempty"`
	// Flag to adjust the overlapping dates automatically. True will allow the system to adjust the overlapping dates automatically . False will not allow overlapping dates and throws error messages if overlapping dates are found
	AdjustOverlappingRange *bool `json:"adjustOverlappingRange,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
	// Package code schedule price exception for the rate plan.
	ScheduleExceptions []ConfigPackageScheduleType `json:"scheduleExceptions,omitempty"`
	// Quantity of the package code included in the rate plan.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewRatePackageCodeType instantiates a new RatePackageCodeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatePackageCodeType() *RatePackageCodeType {
	this := RatePackageCodeType{}
	return &this
}

// NewRatePackageCodeTypeWithDefaults instantiates a new RatePackageCodeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatePackageCodeTypeWithDefaults() *RatePackageCodeType {
	this := RatePackageCodeType{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetHeader() PackageCodeHeaderType {
	if o == nil || IsNil(o.Header) {
		var ret PackageCodeHeaderType
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetHeaderOk() (*PackageCodeHeaderType, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given PackageCodeHeaderType and assigns it to the Header field.
func (o *RatePackageCodeType) SetHeader(v PackageCodeHeaderType) {
	o.Header = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetSchedules() []ConfigPackageScheduleType {
	if o == nil || IsNil(o.Schedules) {
		var ret []ConfigPackageScheduleType
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetSchedulesOk() ([]ConfigPackageScheduleType, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []ConfigPackageScheduleType and assigns it to the Schedules field.
func (o *RatePackageCodeType) SetSchedules(v []ConfigPackageScheduleType) {
	o.Schedules = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *RatePackageCodeType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RatePackageCodeType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RatePackageCodeType) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetGroup() bool {
	if o == nil || IsNil(o.Group) {
		var ret bool
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given bool and assigns it to the Group field.
func (o *RatePackageCodeType) SetGroup(v bool) {
	o.Group = &v
}

// GetRatePlanCode returns the RatePlanCode field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetRatePlanCode() string {
	if o == nil || IsNil(o.RatePlanCode) {
		var ret string
		return ret
	}
	return *o.RatePlanCode
}

// GetRatePlanCodeOk returns a tuple with the RatePlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetRatePlanCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RatePlanCode) {
		return nil, false
	}
	return o.RatePlanCode, true
}

// HasRatePlanCode returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasRatePlanCode() bool {
	if o != nil && !IsNil(o.RatePlanCode) {
		return true
	}

	return false
}

// SetRatePlanCode gets a reference to the given string and assigns it to the RatePlanCode field.
func (o *RatePackageCodeType) SetRatePlanCode(v string) {
	o.RatePlanCode = &v
}

// GetAdjustOverlappingRange returns the AdjustOverlappingRange field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetAdjustOverlappingRange() bool {
	if o == nil || IsNil(o.AdjustOverlappingRange) {
		var ret bool
		return ret
	}
	return *o.AdjustOverlappingRange
}

// GetAdjustOverlappingRangeOk returns a tuple with the AdjustOverlappingRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetAdjustOverlappingRangeOk() (*bool, bool) {
	if o == nil || IsNil(o.AdjustOverlappingRange) {
		return nil, false
	}
	return o.AdjustOverlappingRange, true
}

// HasAdjustOverlappingRange returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasAdjustOverlappingRange() bool {
	if o != nil && !IsNil(o.AdjustOverlappingRange) {
		return true
	}

	return false
}

// SetAdjustOverlappingRange gets a reference to the given bool and assigns it to the AdjustOverlappingRange field.
func (o *RatePackageCodeType) SetAdjustOverlappingRange(v bool) {
	o.AdjustOverlappingRange = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *RatePackageCodeType) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *RatePackageCodeType) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *RatePackageCodeType) SetCount(v int32) {
	o.Count = &v
}

// GetScheduleExceptions returns the ScheduleExceptions field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetScheduleExceptions() []ConfigPackageScheduleType {
	if o == nil || IsNil(o.ScheduleExceptions) {
		var ret []ConfigPackageScheduleType
		return ret
	}
	return o.ScheduleExceptions
}

// GetScheduleExceptionsOk returns a tuple with the ScheduleExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetScheduleExceptionsOk() ([]ConfigPackageScheduleType, bool) {
	if o == nil || IsNil(o.ScheduleExceptions) {
		return nil, false
	}
	return o.ScheduleExceptions, true
}

// HasScheduleExceptions returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasScheduleExceptions() bool {
	if o != nil && !IsNil(o.ScheduleExceptions) {
		return true
	}

	return false
}

// SetScheduleExceptions gets a reference to the given []ConfigPackageScheduleType and assigns it to the ScheduleExceptions field.
func (o *RatePackageCodeType) SetScheduleExceptions(v []ConfigPackageScheduleType) {
	o.ScheduleExceptions = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *RatePackageCodeType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePackageCodeType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *RatePackageCodeType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *RatePackageCodeType) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o RatePackageCodeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatePackageCodeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.RatePlanCode) {
		toSerialize["ratePlanCode"] = o.RatePlanCode
	}
	if !IsNil(o.AdjustOverlappingRange) {
		toSerialize["adjustOverlappingRange"] = o.AdjustOverlappingRange
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
	if !IsNil(o.ScheduleExceptions) {
		toSerialize["scheduleExceptions"] = o.ScheduleExceptions
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableRatePackageCodeType struct {
	value *RatePackageCodeType
	isSet bool
}

func (v NullableRatePackageCodeType) Get() *RatePackageCodeType {
	return v.value
}

func (v *NullableRatePackageCodeType) Set(val *RatePackageCodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePackageCodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePackageCodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePackageCodeType(val *RatePackageCodeType) *NullableRatePackageCodeType {
	return &NullableRatePackageCodeType{value: val, isSet: true}
}

func (v NullableRatePackageCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePackageCodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


