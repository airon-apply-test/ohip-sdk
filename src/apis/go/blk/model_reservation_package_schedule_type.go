/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ReservationPackageScheduleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationPackageScheduleType{}

// ReservationPackageScheduleType A HotelPackageSchedule type.
type ReservationPackageScheduleType struct {
	// The date the package was used or can be used.
	ConsumptionDate *string `json:"consumptionDate,omitempty"`
	// The price per unit of the package.
	UnitPrice *float32 `json:"unitPrice,omitempty"`
	// The total quantity of the package for this date, calculated based on the calculation rule as defined in the PackageHeaderType
	TotalQuantity *int32 `json:"totalQuantity,omitempty"`
	// Computed Reservation Price of the package. Calculation Will Be Performed Based On Other Parameters.
	ComputedResvPrice *float32 `json:"computedResvPrice,omitempty"`
	// The allowance per unit of the package.
	UnitAllowance *float32 `json:"unitAllowance,omitempty"`
	// The date of the Reservation when this package is applicable. This can be different from the date the package will be consumed. Example are next day packages. Reservation date is when the package is applied to the guest and Consumption date is when the guest can consume the package.
	ReservationDate *string `json:"reservationDate,omitempty"`
	// The original price per unit of the package if it has been changed.
	OriginalUnitPrice *float32 `json:"originalUnitPrice,omitempty"`
	// The original allowance per unit of the package if it has been changed.
	OriginalUnitAllowance *float32 `json:"originalUnitAllowance,omitempty"`
}

// NewReservationPackageScheduleType instantiates a new ReservationPackageScheduleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationPackageScheduleType() *ReservationPackageScheduleType {
	this := ReservationPackageScheduleType{}
	return &this
}

// NewReservationPackageScheduleTypeWithDefaults instantiates a new ReservationPackageScheduleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationPackageScheduleTypeWithDefaults() *ReservationPackageScheduleType {
	this := ReservationPackageScheduleType{}
	return &this
}

// GetConsumptionDate returns the ConsumptionDate field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetConsumptionDate() string {
	if o == nil || IsNil(o.ConsumptionDate) {
		var ret string
		return ret
	}
	return *o.ConsumptionDate
}

// GetConsumptionDateOk returns a tuple with the ConsumptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetConsumptionDateOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumptionDate) {
		return nil, false
	}
	return o.ConsumptionDate, true
}

// HasConsumptionDate returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasConsumptionDate() bool {
	if o != nil && !IsNil(o.ConsumptionDate) {
		return true
	}

	return false
}

// SetConsumptionDate gets a reference to the given string and assigns it to the ConsumptionDate field.
func (o *ReservationPackageScheduleType) SetConsumptionDate(v string) {
	o.ConsumptionDate = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetUnitPrice() float32 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret float32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetUnitPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float32 and assigns it to the UnitPrice field.
func (o *ReservationPackageScheduleType) SetUnitPrice(v float32) {
	o.UnitPrice = &v
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetTotalQuantity() int32 {
	if o == nil || IsNil(o.TotalQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetTotalQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalQuantity) {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasTotalQuantity() bool {
	if o != nil && !IsNil(o.TotalQuantity) {
		return true
	}

	return false
}

// SetTotalQuantity gets a reference to the given int32 and assigns it to the TotalQuantity field.
func (o *ReservationPackageScheduleType) SetTotalQuantity(v int32) {
	o.TotalQuantity = &v
}

// GetComputedResvPrice returns the ComputedResvPrice field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetComputedResvPrice() float32 {
	if o == nil || IsNil(o.ComputedResvPrice) {
		var ret float32
		return ret
	}
	return *o.ComputedResvPrice
}

// GetComputedResvPriceOk returns a tuple with the ComputedResvPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetComputedResvPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.ComputedResvPrice) {
		return nil, false
	}
	return o.ComputedResvPrice, true
}

// HasComputedResvPrice returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasComputedResvPrice() bool {
	if o != nil && !IsNil(o.ComputedResvPrice) {
		return true
	}

	return false
}

// SetComputedResvPrice gets a reference to the given float32 and assigns it to the ComputedResvPrice field.
func (o *ReservationPackageScheduleType) SetComputedResvPrice(v float32) {
	o.ComputedResvPrice = &v
}

// GetUnitAllowance returns the UnitAllowance field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetUnitAllowance() float32 {
	if o == nil || IsNil(o.UnitAllowance) {
		var ret float32
		return ret
	}
	return *o.UnitAllowance
}

// GetUnitAllowanceOk returns a tuple with the UnitAllowance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetUnitAllowanceOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitAllowance) {
		return nil, false
	}
	return o.UnitAllowance, true
}

// HasUnitAllowance returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasUnitAllowance() bool {
	if o != nil && !IsNil(o.UnitAllowance) {
		return true
	}

	return false
}

// SetUnitAllowance gets a reference to the given float32 and assigns it to the UnitAllowance field.
func (o *ReservationPackageScheduleType) SetUnitAllowance(v float32) {
	o.UnitAllowance = &v
}

// GetReservationDate returns the ReservationDate field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetReservationDate() string {
	if o == nil || IsNil(o.ReservationDate) {
		var ret string
		return ret
	}
	return *o.ReservationDate
}

// GetReservationDateOk returns a tuple with the ReservationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetReservationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReservationDate) {
		return nil, false
	}
	return o.ReservationDate, true
}

// HasReservationDate returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasReservationDate() bool {
	if o != nil && !IsNil(o.ReservationDate) {
		return true
	}

	return false
}

// SetReservationDate gets a reference to the given string and assigns it to the ReservationDate field.
func (o *ReservationPackageScheduleType) SetReservationDate(v string) {
	o.ReservationDate = &v
}

// GetOriginalUnitPrice returns the OriginalUnitPrice field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetOriginalUnitPrice() float32 {
	if o == nil || IsNil(o.OriginalUnitPrice) {
		var ret float32
		return ret
	}
	return *o.OriginalUnitPrice
}

// GetOriginalUnitPriceOk returns a tuple with the OriginalUnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetOriginalUnitPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.OriginalUnitPrice) {
		return nil, false
	}
	return o.OriginalUnitPrice, true
}

// HasOriginalUnitPrice returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasOriginalUnitPrice() bool {
	if o != nil && !IsNil(o.OriginalUnitPrice) {
		return true
	}

	return false
}

// SetOriginalUnitPrice gets a reference to the given float32 and assigns it to the OriginalUnitPrice field.
func (o *ReservationPackageScheduleType) SetOriginalUnitPrice(v float32) {
	o.OriginalUnitPrice = &v
}

// GetOriginalUnitAllowance returns the OriginalUnitAllowance field value if set, zero value otherwise.
func (o *ReservationPackageScheduleType) GetOriginalUnitAllowance() float32 {
	if o == nil || IsNil(o.OriginalUnitAllowance) {
		var ret float32
		return ret
	}
	return *o.OriginalUnitAllowance
}

// GetOriginalUnitAllowanceOk returns a tuple with the OriginalUnitAllowance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPackageScheduleType) GetOriginalUnitAllowanceOk() (*float32, bool) {
	if o == nil || IsNil(o.OriginalUnitAllowance) {
		return nil, false
	}
	return o.OriginalUnitAllowance, true
}

// HasOriginalUnitAllowance returns a boolean if a field has been set.
func (o *ReservationPackageScheduleType) HasOriginalUnitAllowance() bool {
	if o != nil && !IsNil(o.OriginalUnitAllowance) {
		return true
	}

	return false
}

// SetOriginalUnitAllowance gets a reference to the given float32 and assigns it to the OriginalUnitAllowance field.
func (o *ReservationPackageScheduleType) SetOriginalUnitAllowance(v float32) {
	o.OriginalUnitAllowance = &v
}

func (o ReservationPackageScheduleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationPackageScheduleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumptionDate) {
		toSerialize["consumptionDate"] = o.ConsumptionDate
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if !IsNil(o.TotalQuantity) {
		toSerialize["totalQuantity"] = o.TotalQuantity
	}
	if !IsNil(o.ComputedResvPrice) {
		toSerialize["computedResvPrice"] = o.ComputedResvPrice
	}
	if !IsNil(o.UnitAllowance) {
		toSerialize["unitAllowance"] = o.UnitAllowance
	}
	if !IsNil(o.ReservationDate) {
		toSerialize["reservationDate"] = o.ReservationDate
	}
	if !IsNil(o.OriginalUnitPrice) {
		toSerialize["originalUnitPrice"] = o.OriginalUnitPrice
	}
	if !IsNil(o.OriginalUnitAllowance) {
		toSerialize["originalUnitAllowance"] = o.OriginalUnitAllowance
	}
	return toSerialize, nil
}

type NullableReservationPackageScheduleType struct {
	value *ReservationPackageScheduleType
	isSet bool
}

func (v NullableReservationPackageScheduleType) Get() *ReservationPackageScheduleType {
	return v.value
}

func (v *NullableReservationPackageScheduleType) Set(val *ReservationPackageScheduleType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationPackageScheduleType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationPackageScheduleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationPackageScheduleType(val *ReservationPackageScheduleType) *NullableReservationPackageScheduleType {
	return &NullableReservationPackageScheduleType{value: val, isSet: true}
}

func (v NullableReservationPackageScheduleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationPackageScheduleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


