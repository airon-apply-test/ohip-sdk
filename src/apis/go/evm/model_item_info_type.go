/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ItemInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemInfoType{}

// ItemInfoType Basic information regarding an Item.
type ItemInfoType struct {
	// Detail description of an item.
	Description *string `json:"description,omitempty"`
	AvailabilityPeriod *TimeWindowType `json:"availabilityPeriod,omitempty"`
	BlockDates *TimeSpanType `json:"blockDates,omitempty"`
	// Quantity of hold Item
	Quantity *int32 `json:"quantity,omitempty"`
	// ID reference for the hold Item
	ItemHoldId *float32 `json:"itemHoldId,omitempty"`
	// Item Code.
	Code *string `json:"code,omitempty"`
	// Name of an item.
	Name *string `json:"name,omitempty"`
	// Indicates if it is an item pool. Not applicable for Item within the Item Pool.
	ItemPool *bool `json:"itemPool,omitempty"`
	// If true indicates that item is allowed to sell separately.
	SellSeparate *bool `json:"sellSeparate,omitempty"`
	// If true indicates that item can be sold in reservation.
	SellInReservation *bool `json:"sellInReservation,omitempty"`
	// If true indicates that item can be sold in event.
	SellInEvent *bool `json:"sellInEvent,omitempty"`
	// If true indicates that item is required for the reservation.
	RequiredForBooking *bool `json:"requiredForBooking,omitempty"`
	// If true indicates that item has fixed charge when it is attached to a reservation.
	FixedCharge *bool `json:"fixedCharge,omitempty"`
	// If true indicates that item could be held outside of the reservation stay days.
	OutsideStay *bool `json:"outsideStay,omitempty"`
	// Define the default duration in days when booking the item.
	DefaultDuration *int32 `json:"defaultDuration,omitempty"`
}

// NewItemInfoType instantiates a new ItemInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemInfoType() *ItemInfoType {
	this := ItemInfoType{}
	return &this
}

// NewItemInfoTypeWithDefaults instantiates a new ItemInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemInfoTypeWithDefaults() *ItemInfoType {
	this := ItemInfoType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ItemInfoType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ItemInfoType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ItemInfoType) SetDescription(v string) {
	o.Description = &v
}

// GetAvailabilityPeriod returns the AvailabilityPeriod field value if set, zero value otherwise.
func (o *ItemInfoType) GetAvailabilityPeriod() TimeWindowType {
	if o == nil || IsNil(o.AvailabilityPeriod) {
		var ret TimeWindowType
		return ret
	}
	return *o.AvailabilityPeriod
}

// GetAvailabilityPeriodOk returns a tuple with the AvailabilityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetAvailabilityPeriodOk() (*TimeWindowType, bool) {
	if o == nil || IsNil(o.AvailabilityPeriod) {
		return nil, false
	}
	return o.AvailabilityPeriod, true
}

// HasAvailabilityPeriod returns a boolean if a field has been set.
func (o *ItemInfoType) HasAvailabilityPeriod() bool {
	if o != nil && !IsNil(o.AvailabilityPeriod) {
		return true
	}

	return false
}

// SetAvailabilityPeriod gets a reference to the given TimeWindowType and assigns it to the AvailabilityPeriod field.
func (o *ItemInfoType) SetAvailabilityPeriod(v TimeWindowType) {
	o.AvailabilityPeriod = &v
}

// GetBlockDates returns the BlockDates field value if set, zero value otherwise.
func (o *ItemInfoType) GetBlockDates() TimeSpanType {
	if o == nil || IsNil(o.BlockDates) {
		var ret TimeSpanType
		return ret
	}
	return *o.BlockDates
}

// GetBlockDatesOk returns a tuple with the BlockDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetBlockDatesOk() (*TimeSpanType, bool) {
	if o == nil || IsNil(o.BlockDates) {
		return nil, false
	}
	return o.BlockDates, true
}

// HasBlockDates returns a boolean if a field has been set.
func (o *ItemInfoType) HasBlockDates() bool {
	if o != nil && !IsNil(o.BlockDates) {
		return true
	}

	return false
}

// SetBlockDates gets a reference to the given TimeSpanType and assigns it to the BlockDates field.
func (o *ItemInfoType) SetBlockDates(v TimeSpanType) {
	o.BlockDates = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ItemInfoType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ItemInfoType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ItemInfoType) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetItemHoldId returns the ItemHoldId field value if set, zero value otherwise.
func (o *ItemInfoType) GetItemHoldId() float32 {
	if o == nil || IsNil(o.ItemHoldId) {
		var ret float32
		return ret
	}
	return *o.ItemHoldId
}

// GetItemHoldIdOk returns a tuple with the ItemHoldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetItemHoldIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ItemHoldId) {
		return nil, false
	}
	return o.ItemHoldId, true
}

// HasItemHoldId returns a boolean if a field has been set.
func (o *ItemInfoType) HasItemHoldId() bool {
	if o != nil && !IsNil(o.ItemHoldId) {
		return true
	}

	return false
}

// SetItemHoldId gets a reference to the given float32 and assigns it to the ItemHoldId field.
func (o *ItemInfoType) SetItemHoldId(v float32) {
	o.ItemHoldId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ItemInfoType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ItemInfoType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ItemInfoType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ItemInfoType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ItemInfoType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ItemInfoType) SetName(v string) {
	o.Name = &v
}

// GetItemPool returns the ItemPool field value if set, zero value otherwise.
func (o *ItemInfoType) GetItemPool() bool {
	if o == nil || IsNil(o.ItemPool) {
		var ret bool
		return ret
	}
	return *o.ItemPool
}

// GetItemPoolOk returns a tuple with the ItemPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetItemPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.ItemPool) {
		return nil, false
	}
	return o.ItemPool, true
}

// HasItemPool returns a boolean if a field has been set.
func (o *ItemInfoType) HasItemPool() bool {
	if o != nil && !IsNil(o.ItemPool) {
		return true
	}

	return false
}

// SetItemPool gets a reference to the given bool and assigns it to the ItemPool field.
func (o *ItemInfoType) SetItemPool(v bool) {
	o.ItemPool = &v
}

// GetSellSeparate returns the SellSeparate field value if set, zero value otherwise.
func (o *ItemInfoType) GetSellSeparate() bool {
	if o == nil || IsNil(o.SellSeparate) {
		var ret bool
		return ret
	}
	return *o.SellSeparate
}

// GetSellSeparateOk returns a tuple with the SellSeparate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetSellSeparateOk() (*bool, bool) {
	if o == nil || IsNil(o.SellSeparate) {
		return nil, false
	}
	return o.SellSeparate, true
}

// HasSellSeparate returns a boolean if a field has been set.
func (o *ItemInfoType) HasSellSeparate() bool {
	if o != nil && !IsNil(o.SellSeparate) {
		return true
	}

	return false
}

// SetSellSeparate gets a reference to the given bool and assigns it to the SellSeparate field.
func (o *ItemInfoType) SetSellSeparate(v bool) {
	o.SellSeparate = &v
}

// GetSellInReservation returns the SellInReservation field value if set, zero value otherwise.
func (o *ItemInfoType) GetSellInReservation() bool {
	if o == nil || IsNil(o.SellInReservation) {
		var ret bool
		return ret
	}
	return *o.SellInReservation
}

// GetSellInReservationOk returns a tuple with the SellInReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetSellInReservationOk() (*bool, bool) {
	if o == nil || IsNil(o.SellInReservation) {
		return nil, false
	}
	return o.SellInReservation, true
}

// HasSellInReservation returns a boolean if a field has been set.
func (o *ItemInfoType) HasSellInReservation() bool {
	if o != nil && !IsNil(o.SellInReservation) {
		return true
	}

	return false
}

// SetSellInReservation gets a reference to the given bool and assigns it to the SellInReservation field.
func (o *ItemInfoType) SetSellInReservation(v bool) {
	o.SellInReservation = &v
}

// GetSellInEvent returns the SellInEvent field value if set, zero value otherwise.
func (o *ItemInfoType) GetSellInEvent() bool {
	if o == nil || IsNil(o.SellInEvent) {
		var ret bool
		return ret
	}
	return *o.SellInEvent
}

// GetSellInEventOk returns a tuple with the SellInEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetSellInEventOk() (*bool, bool) {
	if o == nil || IsNil(o.SellInEvent) {
		return nil, false
	}
	return o.SellInEvent, true
}

// HasSellInEvent returns a boolean if a field has been set.
func (o *ItemInfoType) HasSellInEvent() bool {
	if o != nil && !IsNil(o.SellInEvent) {
		return true
	}

	return false
}

// SetSellInEvent gets a reference to the given bool and assigns it to the SellInEvent field.
func (o *ItemInfoType) SetSellInEvent(v bool) {
	o.SellInEvent = &v
}

// GetRequiredForBooking returns the RequiredForBooking field value if set, zero value otherwise.
func (o *ItemInfoType) GetRequiredForBooking() bool {
	if o == nil || IsNil(o.RequiredForBooking) {
		var ret bool
		return ret
	}
	return *o.RequiredForBooking
}

// GetRequiredForBookingOk returns a tuple with the RequiredForBooking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetRequiredForBookingOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredForBooking) {
		return nil, false
	}
	return o.RequiredForBooking, true
}

// HasRequiredForBooking returns a boolean if a field has been set.
func (o *ItemInfoType) HasRequiredForBooking() bool {
	if o != nil && !IsNil(o.RequiredForBooking) {
		return true
	}

	return false
}

// SetRequiredForBooking gets a reference to the given bool and assigns it to the RequiredForBooking field.
func (o *ItemInfoType) SetRequiredForBooking(v bool) {
	o.RequiredForBooking = &v
}

// GetFixedCharge returns the FixedCharge field value if set, zero value otherwise.
func (o *ItemInfoType) GetFixedCharge() bool {
	if o == nil || IsNil(o.FixedCharge) {
		var ret bool
		return ret
	}
	return *o.FixedCharge
}

// GetFixedChargeOk returns a tuple with the FixedCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetFixedChargeOk() (*bool, bool) {
	if o == nil || IsNil(o.FixedCharge) {
		return nil, false
	}
	return o.FixedCharge, true
}

// HasFixedCharge returns a boolean if a field has been set.
func (o *ItemInfoType) HasFixedCharge() bool {
	if o != nil && !IsNil(o.FixedCharge) {
		return true
	}

	return false
}

// SetFixedCharge gets a reference to the given bool and assigns it to the FixedCharge field.
func (o *ItemInfoType) SetFixedCharge(v bool) {
	o.FixedCharge = &v
}

// GetOutsideStay returns the OutsideStay field value if set, zero value otherwise.
func (o *ItemInfoType) GetOutsideStay() bool {
	if o == nil || IsNil(o.OutsideStay) {
		var ret bool
		return ret
	}
	return *o.OutsideStay
}

// GetOutsideStayOk returns a tuple with the OutsideStay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetOutsideStayOk() (*bool, bool) {
	if o == nil || IsNil(o.OutsideStay) {
		return nil, false
	}
	return o.OutsideStay, true
}

// HasOutsideStay returns a boolean if a field has been set.
func (o *ItemInfoType) HasOutsideStay() bool {
	if o != nil && !IsNil(o.OutsideStay) {
		return true
	}

	return false
}

// SetOutsideStay gets a reference to the given bool and assigns it to the OutsideStay field.
func (o *ItemInfoType) SetOutsideStay(v bool) {
	o.OutsideStay = &v
}

// GetDefaultDuration returns the DefaultDuration field value if set, zero value otherwise.
func (o *ItemInfoType) GetDefaultDuration() int32 {
	if o == nil || IsNil(o.DefaultDuration) {
		var ret int32
		return ret
	}
	return *o.DefaultDuration
}

// GetDefaultDurationOk returns a tuple with the DefaultDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInfoType) GetDefaultDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultDuration) {
		return nil, false
	}
	return o.DefaultDuration, true
}

// HasDefaultDuration returns a boolean if a field has been set.
func (o *ItemInfoType) HasDefaultDuration() bool {
	if o != nil && !IsNil(o.DefaultDuration) {
		return true
	}

	return false
}

// SetDefaultDuration gets a reference to the given int32 and assigns it to the DefaultDuration field.
func (o *ItemInfoType) SetDefaultDuration(v int32) {
	o.DefaultDuration = &v
}

func (o ItemInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AvailabilityPeriod) {
		toSerialize["availabilityPeriod"] = o.AvailabilityPeriod
	}
	if !IsNil(o.BlockDates) {
		toSerialize["blockDates"] = o.BlockDates
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ItemHoldId) {
		toSerialize["itemHoldId"] = o.ItemHoldId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ItemPool) {
		toSerialize["itemPool"] = o.ItemPool
	}
	if !IsNil(o.SellSeparate) {
		toSerialize["sellSeparate"] = o.SellSeparate
	}
	if !IsNil(o.SellInReservation) {
		toSerialize["sellInReservation"] = o.SellInReservation
	}
	if !IsNil(o.SellInEvent) {
		toSerialize["sellInEvent"] = o.SellInEvent
	}
	if !IsNil(o.RequiredForBooking) {
		toSerialize["requiredForBooking"] = o.RequiredForBooking
	}
	if !IsNil(o.FixedCharge) {
		toSerialize["fixedCharge"] = o.FixedCharge
	}
	if !IsNil(o.OutsideStay) {
		toSerialize["outsideStay"] = o.OutsideStay
	}
	if !IsNil(o.DefaultDuration) {
		toSerialize["defaultDuration"] = o.DefaultDuration
	}
	return toSerialize, nil
}

type NullableItemInfoType struct {
	value *ItemInfoType
	isSet bool
}

func (v NullableItemInfoType) Get() *ItemInfoType {
	return v.value
}

func (v *NullableItemInfoType) Set(val *ItemInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableItemInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableItemInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemInfoType(val *ItemInfoType) *NullableItemInfoType {
	return &NullableItemInfoType{value: val, isSet: true}
}

func (v NullableItemInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


