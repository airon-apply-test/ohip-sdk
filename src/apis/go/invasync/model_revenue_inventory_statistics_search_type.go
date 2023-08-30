/*
Opera Cloud Inventory Asynchronous API

APIs to cater for Inventory related asynchronous functionality in OPERA Cloud. This includes viewing inventory data along with its revenue and updating inventory&apos;s sell limits for date ranges. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.4.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invasync

import (
	"encoding/json"
)

// checks if the RevenueInventoryStatisticsSearchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevenueInventoryStatisticsSearchType{}

// RevenueInventoryStatisticsSearchType Fetch criteria type for fetching revenue inventory statistics.
type RevenueInventoryStatisticsSearchType struct {
	// The starting value of the date range.
	DateRangeStart *string `json:"dateRangeStart,omitempty"`
	// The ending value of the date range.
	DateRangeEnd *string `json:"dateRangeEnd,omitempty"`
	// The optional room types by which revenue and inventory statistics will be grouped.
	RoomTypes []string `json:"roomTypes,omitempty"`
	// The market codes for which revenue and inventory statistics will be gathered.
	MarketCodes []string `json:"marketCodes,omitempty"`
	// The reservation guarantee codes for which revenue and inventory statistics will be gathered.
	GuaranteeCodes []string `json:"guaranteeCodes,omitempty"`
	// Group by instructions for revenue inventory statistics that can be used in request operation. Response data will be grouped according to the values provided in this array. MarketCode and RoomType grouping is used for past and future revenue and inventory statistics while GuaranteeType is used only for future revenue and inventory statistics.
	GroupBy []string `json:"groupBy,omitempty"`
}

// NewRevenueInventoryStatisticsSearchType instantiates a new RevenueInventoryStatisticsSearchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevenueInventoryStatisticsSearchType() *RevenueInventoryStatisticsSearchType {
	this := RevenueInventoryStatisticsSearchType{}
	return &this
}

// NewRevenueInventoryStatisticsSearchTypeWithDefaults instantiates a new RevenueInventoryStatisticsSearchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevenueInventoryStatisticsSearchTypeWithDefaults() *RevenueInventoryStatisticsSearchType {
	this := RevenueInventoryStatisticsSearchType{}
	return &this
}

// GetDateRangeStart returns the DateRangeStart field value if set, zero value otherwise.
func (o *RevenueInventoryStatisticsSearchType) GetDateRangeStart() string {
	if o == nil || IsNil(o.DateRangeStart) {
		var ret string
		return ret
	}
	return *o.DateRangeStart
}

// GetDateRangeStartOk returns a tuple with the DateRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueInventoryStatisticsSearchType) GetDateRangeStartOk() (*string, bool) {
	if o == nil || IsNil(o.DateRangeStart) {
		return nil, false
	}
	return o.DateRangeStart, true
}

// HasDateRangeStart returns a boolean if a field has been set.
func (o *RevenueInventoryStatisticsSearchType) HasDateRangeStart() bool {
	if o != nil && !IsNil(o.DateRangeStart) {
		return true
	}

	return false
}

// SetDateRangeStart gets a reference to the given string and assigns it to the DateRangeStart field.
func (o *RevenueInventoryStatisticsSearchType) SetDateRangeStart(v string) {
	o.DateRangeStart = &v
}

// GetDateRangeEnd returns the DateRangeEnd field value if set, zero value otherwise.
func (o *RevenueInventoryStatisticsSearchType) GetDateRangeEnd() string {
	if o == nil || IsNil(o.DateRangeEnd) {
		var ret string
		return ret
	}
	return *o.DateRangeEnd
}

// GetDateRangeEndOk returns a tuple with the DateRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueInventoryStatisticsSearchType) GetDateRangeEndOk() (*string, bool) {
	if o == nil || IsNil(o.DateRangeEnd) {
		return nil, false
	}
	return o.DateRangeEnd, true
}

// HasDateRangeEnd returns a boolean if a field has been set.
func (o *RevenueInventoryStatisticsSearchType) HasDateRangeEnd() bool {
	if o != nil && !IsNil(o.DateRangeEnd) {
		return true
	}

	return false
}

// SetDateRangeEnd gets a reference to the given string and assigns it to the DateRangeEnd field.
func (o *RevenueInventoryStatisticsSearchType) SetDateRangeEnd(v string) {
	o.DateRangeEnd = &v
}

// GetRoomTypes returns the RoomTypes field value if set, zero value otherwise.
func (o *RevenueInventoryStatisticsSearchType) GetRoomTypes() []string {
	if o == nil || IsNil(o.RoomTypes) {
		var ret []string
		return ret
	}
	return o.RoomTypes
}

// GetRoomTypesOk returns a tuple with the RoomTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueInventoryStatisticsSearchType) GetRoomTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.RoomTypes) {
		return nil, false
	}
	return o.RoomTypes, true
}

// HasRoomTypes returns a boolean if a field has been set.
func (o *RevenueInventoryStatisticsSearchType) HasRoomTypes() bool {
	if o != nil && !IsNil(o.RoomTypes) {
		return true
	}

	return false
}

// SetRoomTypes gets a reference to the given []string and assigns it to the RoomTypes field.
func (o *RevenueInventoryStatisticsSearchType) SetRoomTypes(v []string) {
	o.RoomTypes = v
}

// GetMarketCodes returns the MarketCodes field value if set, zero value otherwise.
func (o *RevenueInventoryStatisticsSearchType) GetMarketCodes() []string {
	if o == nil || IsNil(o.MarketCodes) {
		var ret []string
		return ret
	}
	return o.MarketCodes
}

// GetMarketCodesOk returns a tuple with the MarketCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueInventoryStatisticsSearchType) GetMarketCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.MarketCodes) {
		return nil, false
	}
	return o.MarketCodes, true
}

// HasMarketCodes returns a boolean if a field has been set.
func (o *RevenueInventoryStatisticsSearchType) HasMarketCodes() bool {
	if o != nil && !IsNil(o.MarketCodes) {
		return true
	}

	return false
}

// SetMarketCodes gets a reference to the given []string and assigns it to the MarketCodes field.
func (o *RevenueInventoryStatisticsSearchType) SetMarketCodes(v []string) {
	o.MarketCodes = v
}

// GetGuaranteeCodes returns the GuaranteeCodes field value if set, zero value otherwise.
func (o *RevenueInventoryStatisticsSearchType) GetGuaranteeCodes() []string {
	if o == nil || IsNil(o.GuaranteeCodes) {
		var ret []string
		return ret
	}
	return o.GuaranteeCodes
}

// GetGuaranteeCodesOk returns a tuple with the GuaranteeCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueInventoryStatisticsSearchType) GetGuaranteeCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.GuaranteeCodes) {
		return nil, false
	}
	return o.GuaranteeCodes, true
}

// HasGuaranteeCodes returns a boolean if a field has been set.
func (o *RevenueInventoryStatisticsSearchType) HasGuaranteeCodes() bool {
	if o != nil && !IsNil(o.GuaranteeCodes) {
		return true
	}

	return false
}

// SetGuaranteeCodes gets a reference to the given []string and assigns it to the GuaranteeCodes field.
func (o *RevenueInventoryStatisticsSearchType) SetGuaranteeCodes(v []string) {
	o.GuaranteeCodes = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RevenueInventoryStatisticsSearchType) GetGroupBy() []string {
	if o == nil || IsNil(o.GroupBy) {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueInventoryStatisticsSearchType) GetGroupByOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RevenueInventoryStatisticsSearchType) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *RevenueInventoryStatisticsSearchType) SetGroupBy(v []string) {
	o.GroupBy = v
}

func (o RevenueInventoryStatisticsSearchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevenueInventoryStatisticsSearchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateRangeStart) {
		toSerialize["dateRangeStart"] = o.DateRangeStart
	}
	if !IsNil(o.DateRangeEnd) {
		toSerialize["dateRangeEnd"] = o.DateRangeEnd
	}
	if !IsNil(o.RoomTypes) {
		toSerialize["roomTypes"] = o.RoomTypes
	}
	if !IsNil(o.MarketCodes) {
		toSerialize["marketCodes"] = o.MarketCodes
	}
	if !IsNil(o.GuaranteeCodes) {
		toSerialize["guaranteeCodes"] = o.GuaranteeCodes
	}
	if !IsNil(o.GroupBy) {
		toSerialize["groupBy"] = o.GroupBy
	}
	return toSerialize, nil
}

type NullableRevenueInventoryStatisticsSearchType struct {
	value *RevenueInventoryStatisticsSearchType
	isSet bool
}

func (v NullableRevenueInventoryStatisticsSearchType) Get() *RevenueInventoryStatisticsSearchType {
	return v.value
}

func (v *NullableRevenueInventoryStatisticsSearchType) Set(val *RevenueInventoryStatisticsSearchType) {
	v.value = val
	v.isSet = true
}

func (v NullableRevenueInventoryStatisticsSearchType) IsSet() bool {
	return v.isSet
}

func (v *NullableRevenueInventoryStatisticsSearchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevenueInventoryStatisticsSearchType(val *RevenueInventoryStatisticsSearchType) *NullableRevenueInventoryStatisticsSearchType {
	return &NullableRevenueInventoryStatisticsSearchType{value: val, isSet: true}
}

func (v NullableRevenueInventoryStatisticsSearchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevenueInventoryStatisticsSearchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


