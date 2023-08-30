/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
)

// checks if the EventFunctionSpaceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventFunctionSpaceType{}

// EventFunctionSpaceType Pertain Information about the Function Space for events.
type EventFunctionSpaceType struct {
	// Function Space (event functionSpaceDetails) code.
	FunctionSpaceCode *string `json:"functionSpaceCode,omitempty"`
	// Function Space description.
	FunctionSpaceDescription *string `json:"functionSpaceDescription,omitempty"`
	// Setup style code.
	SetupCode *string `json:"setupCode,omitempty"`
	// Setup Style description.
	SetupDescription *string `json:"setupDescription,omitempty"`
	// Setup time in minutes.
	SetupTime *int32 `json:"setupTime,omitempty"`
	// Setdown time in minutes.
	SetdownTime *int32 `json:"setdownTime,omitempty"`
	// Rate Code for function space.
	RentalCode *string `json:"rentalCode,omitempty"`
	// Function Space Rate description.
	RateDesc *string `json:"rateDesc,omitempty"`
	RentalAmount *CurrencyAmountType `json:"rentalAmount,omitempty"`
	// Discount Percentage applied to Rent Amount.
	DiscountPercentage *float32 `json:"discountPercentage,omitempty"`
	// Minimum Occupancy for the SetupCode of Catring Event.
	MinimumOccupancy *int32 `json:"minimumOccupancy,omitempty"`
	// Maximum Occupancy for the SetupCode of Catring Event.
	MaximumOccupancy *int32 `json:"maximumOccupancy,omitempty"`
	// Flag to indicate if the event functionSpaceDetails is shareable.
	Shareable *bool `json:"shareable,omitempty"`
	// Flag to indicate if the function space is a combination functionSpaceDetails.
	ComboSpace *bool `json:"comboSpace,omitempty"`
	// Flag to indicate if the function space is an elementSpace of a comboSpace function space.
	ElementSpace *bool `json:"elementSpace,omitempty"`
}

// NewEventFunctionSpaceType instantiates a new EventFunctionSpaceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFunctionSpaceType() *EventFunctionSpaceType {
	this := EventFunctionSpaceType{}
	return &this
}

// NewEventFunctionSpaceTypeWithDefaults instantiates a new EventFunctionSpaceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFunctionSpaceTypeWithDefaults() *EventFunctionSpaceType {
	this := EventFunctionSpaceType{}
	return &this
}

// GetFunctionSpaceCode returns the FunctionSpaceCode field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetFunctionSpaceCode() string {
	if o == nil || IsNil(o.FunctionSpaceCode) {
		var ret string
		return ret
	}
	return *o.FunctionSpaceCode
}

// GetFunctionSpaceCodeOk returns a tuple with the FunctionSpaceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetFunctionSpaceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionSpaceCode) {
		return nil, false
	}
	return o.FunctionSpaceCode, true
}

// HasFunctionSpaceCode returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasFunctionSpaceCode() bool {
	if o != nil && !IsNil(o.FunctionSpaceCode) {
		return true
	}

	return false
}

// SetFunctionSpaceCode gets a reference to the given string and assigns it to the FunctionSpaceCode field.
func (o *EventFunctionSpaceType) SetFunctionSpaceCode(v string) {
	o.FunctionSpaceCode = &v
}

// GetFunctionSpaceDescription returns the FunctionSpaceDescription field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetFunctionSpaceDescription() string {
	if o == nil || IsNil(o.FunctionSpaceDescription) {
		var ret string
		return ret
	}
	return *o.FunctionSpaceDescription
}

// GetFunctionSpaceDescriptionOk returns a tuple with the FunctionSpaceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetFunctionSpaceDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionSpaceDescription) {
		return nil, false
	}
	return o.FunctionSpaceDescription, true
}

// HasFunctionSpaceDescription returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasFunctionSpaceDescription() bool {
	if o != nil && !IsNil(o.FunctionSpaceDescription) {
		return true
	}

	return false
}

// SetFunctionSpaceDescription gets a reference to the given string and assigns it to the FunctionSpaceDescription field.
func (o *EventFunctionSpaceType) SetFunctionSpaceDescription(v string) {
	o.FunctionSpaceDescription = &v
}

// GetSetupCode returns the SetupCode field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetSetupCode() string {
	if o == nil || IsNil(o.SetupCode) {
		var ret string
		return ret
	}
	return *o.SetupCode
}

// GetSetupCodeOk returns a tuple with the SetupCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetSetupCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SetupCode) {
		return nil, false
	}
	return o.SetupCode, true
}

// HasSetupCode returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasSetupCode() bool {
	if o != nil && !IsNil(o.SetupCode) {
		return true
	}

	return false
}

// SetSetupCode gets a reference to the given string and assigns it to the SetupCode field.
func (o *EventFunctionSpaceType) SetSetupCode(v string) {
	o.SetupCode = &v
}

// GetSetupDescription returns the SetupDescription field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetSetupDescription() string {
	if o == nil || IsNil(o.SetupDescription) {
		var ret string
		return ret
	}
	return *o.SetupDescription
}

// GetSetupDescriptionOk returns a tuple with the SetupDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetSetupDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SetupDescription) {
		return nil, false
	}
	return o.SetupDescription, true
}

// HasSetupDescription returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasSetupDescription() bool {
	if o != nil && !IsNil(o.SetupDescription) {
		return true
	}

	return false
}

// SetSetupDescription gets a reference to the given string and assigns it to the SetupDescription field.
func (o *EventFunctionSpaceType) SetSetupDescription(v string) {
	o.SetupDescription = &v
}

// GetSetupTime returns the SetupTime field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetSetupTime() int32 {
	if o == nil || IsNil(o.SetupTime) {
		var ret int32
		return ret
	}
	return *o.SetupTime
}

// GetSetupTimeOk returns a tuple with the SetupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetSetupTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.SetupTime) {
		return nil, false
	}
	return o.SetupTime, true
}

// HasSetupTime returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasSetupTime() bool {
	if o != nil && !IsNil(o.SetupTime) {
		return true
	}

	return false
}

// SetSetupTime gets a reference to the given int32 and assigns it to the SetupTime field.
func (o *EventFunctionSpaceType) SetSetupTime(v int32) {
	o.SetupTime = &v
}

// GetSetdownTime returns the SetdownTime field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetSetdownTime() int32 {
	if o == nil || IsNil(o.SetdownTime) {
		var ret int32
		return ret
	}
	return *o.SetdownTime
}

// GetSetdownTimeOk returns a tuple with the SetdownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetSetdownTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.SetdownTime) {
		return nil, false
	}
	return o.SetdownTime, true
}

// HasSetdownTime returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasSetdownTime() bool {
	if o != nil && !IsNil(o.SetdownTime) {
		return true
	}

	return false
}

// SetSetdownTime gets a reference to the given int32 and assigns it to the SetdownTime field.
func (o *EventFunctionSpaceType) SetSetdownTime(v int32) {
	o.SetdownTime = &v
}

// GetRentalCode returns the RentalCode field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetRentalCode() string {
	if o == nil || IsNil(o.RentalCode) {
		var ret string
		return ret
	}
	return *o.RentalCode
}

// GetRentalCodeOk returns a tuple with the RentalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetRentalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RentalCode) {
		return nil, false
	}
	return o.RentalCode, true
}

// HasRentalCode returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasRentalCode() bool {
	if o != nil && !IsNil(o.RentalCode) {
		return true
	}

	return false
}

// SetRentalCode gets a reference to the given string and assigns it to the RentalCode field.
func (o *EventFunctionSpaceType) SetRentalCode(v string) {
	o.RentalCode = &v
}

// GetRateDesc returns the RateDesc field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetRateDesc() string {
	if o == nil || IsNil(o.RateDesc) {
		var ret string
		return ret
	}
	return *o.RateDesc
}

// GetRateDescOk returns a tuple with the RateDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetRateDescOk() (*string, bool) {
	if o == nil || IsNil(o.RateDesc) {
		return nil, false
	}
	return o.RateDesc, true
}

// HasRateDesc returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasRateDesc() bool {
	if o != nil && !IsNil(o.RateDesc) {
		return true
	}

	return false
}

// SetRateDesc gets a reference to the given string and assigns it to the RateDesc field.
func (o *EventFunctionSpaceType) SetRateDesc(v string) {
	o.RateDesc = &v
}

// GetRentalAmount returns the RentalAmount field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetRentalAmount() CurrencyAmountType {
	if o == nil || IsNil(o.RentalAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.RentalAmount
}

// GetRentalAmountOk returns a tuple with the RentalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetRentalAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.RentalAmount) {
		return nil, false
	}
	return o.RentalAmount, true
}

// HasRentalAmount returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasRentalAmount() bool {
	if o != nil && !IsNil(o.RentalAmount) {
		return true
	}

	return false
}

// SetRentalAmount gets a reference to the given CurrencyAmountType and assigns it to the RentalAmount field.
func (o *EventFunctionSpaceType) SetRentalAmount(v CurrencyAmountType) {
	o.RentalAmount = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetDiscountPercentage() float32 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret float32
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetDiscountPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given float32 and assigns it to the DiscountPercentage field.
func (o *EventFunctionSpaceType) SetDiscountPercentage(v float32) {
	o.DiscountPercentage = &v
}

// GetMinimumOccupancy returns the MinimumOccupancy field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetMinimumOccupancy() int32 {
	if o == nil || IsNil(o.MinimumOccupancy) {
		var ret int32
		return ret
	}
	return *o.MinimumOccupancy
}

// GetMinimumOccupancyOk returns a tuple with the MinimumOccupancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetMinimumOccupancyOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumOccupancy) {
		return nil, false
	}
	return o.MinimumOccupancy, true
}

// HasMinimumOccupancy returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasMinimumOccupancy() bool {
	if o != nil && !IsNil(o.MinimumOccupancy) {
		return true
	}

	return false
}

// SetMinimumOccupancy gets a reference to the given int32 and assigns it to the MinimumOccupancy field.
func (o *EventFunctionSpaceType) SetMinimumOccupancy(v int32) {
	o.MinimumOccupancy = &v
}

// GetMaximumOccupancy returns the MaximumOccupancy field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetMaximumOccupancy() int32 {
	if o == nil || IsNil(o.MaximumOccupancy) {
		var ret int32
		return ret
	}
	return *o.MaximumOccupancy
}

// GetMaximumOccupancyOk returns a tuple with the MaximumOccupancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetMaximumOccupancyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumOccupancy) {
		return nil, false
	}
	return o.MaximumOccupancy, true
}

// HasMaximumOccupancy returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasMaximumOccupancy() bool {
	if o != nil && !IsNil(o.MaximumOccupancy) {
		return true
	}

	return false
}

// SetMaximumOccupancy gets a reference to the given int32 and assigns it to the MaximumOccupancy field.
func (o *EventFunctionSpaceType) SetMaximumOccupancy(v int32) {
	o.MaximumOccupancy = &v
}

// GetShareable returns the Shareable field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetShareable() bool {
	if o == nil || IsNil(o.Shareable) {
		var ret bool
		return ret
	}
	return *o.Shareable
}

// GetShareableOk returns a tuple with the Shareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetShareableOk() (*bool, bool) {
	if o == nil || IsNil(o.Shareable) {
		return nil, false
	}
	return o.Shareable, true
}

// HasShareable returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasShareable() bool {
	if o != nil && !IsNil(o.Shareable) {
		return true
	}

	return false
}

// SetShareable gets a reference to the given bool and assigns it to the Shareable field.
func (o *EventFunctionSpaceType) SetShareable(v bool) {
	o.Shareable = &v
}

// GetComboSpace returns the ComboSpace field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetComboSpace() bool {
	if o == nil || IsNil(o.ComboSpace) {
		var ret bool
		return ret
	}
	return *o.ComboSpace
}

// GetComboSpaceOk returns a tuple with the ComboSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetComboSpaceOk() (*bool, bool) {
	if o == nil || IsNil(o.ComboSpace) {
		return nil, false
	}
	return o.ComboSpace, true
}

// HasComboSpace returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasComboSpace() bool {
	if o != nil && !IsNil(o.ComboSpace) {
		return true
	}

	return false
}

// SetComboSpace gets a reference to the given bool and assigns it to the ComboSpace field.
func (o *EventFunctionSpaceType) SetComboSpace(v bool) {
	o.ComboSpace = &v
}

// GetElementSpace returns the ElementSpace field value if set, zero value otherwise.
func (o *EventFunctionSpaceType) GetElementSpace() bool {
	if o == nil || IsNil(o.ElementSpace) {
		var ret bool
		return ret
	}
	return *o.ElementSpace
}

// GetElementSpaceOk returns a tuple with the ElementSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFunctionSpaceType) GetElementSpaceOk() (*bool, bool) {
	if o == nil || IsNil(o.ElementSpace) {
		return nil, false
	}
	return o.ElementSpace, true
}

// HasElementSpace returns a boolean if a field has been set.
func (o *EventFunctionSpaceType) HasElementSpace() bool {
	if o != nil && !IsNil(o.ElementSpace) {
		return true
	}

	return false
}

// SetElementSpace gets a reference to the given bool and assigns it to the ElementSpace field.
func (o *EventFunctionSpaceType) SetElementSpace(v bool) {
	o.ElementSpace = &v
}

func (o EventFunctionSpaceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventFunctionSpaceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FunctionSpaceCode) {
		toSerialize["functionSpaceCode"] = o.FunctionSpaceCode
	}
	if !IsNil(o.FunctionSpaceDescription) {
		toSerialize["functionSpaceDescription"] = o.FunctionSpaceDescription
	}
	if !IsNil(o.SetupCode) {
		toSerialize["setupCode"] = o.SetupCode
	}
	if !IsNil(o.SetupDescription) {
		toSerialize["setupDescription"] = o.SetupDescription
	}
	if !IsNil(o.SetupTime) {
		toSerialize["setupTime"] = o.SetupTime
	}
	if !IsNil(o.SetdownTime) {
		toSerialize["setdownTime"] = o.SetdownTime
	}
	if !IsNil(o.RentalCode) {
		toSerialize["rentalCode"] = o.RentalCode
	}
	if !IsNil(o.RateDesc) {
		toSerialize["rateDesc"] = o.RateDesc
	}
	if !IsNil(o.RentalAmount) {
		toSerialize["rentalAmount"] = o.RentalAmount
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if !IsNil(o.MinimumOccupancy) {
		toSerialize["minimumOccupancy"] = o.MinimumOccupancy
	}
	if !IsNil(o.MaximumOccupancy) {
		toSerialize["maximumOccupancy"] = o.MaximumOccupancy
	}
	if !IsNil(o.Shareable) {
		toSerialize["shareable"] = o.Shareable
	}
	if !IsNil(o.ComboSpace) {
		toSerialize["comboSpace"] = o.ComboSpace
	}
	if !IsNil(o.ElementSpace) {
		toSerialize["elementSpace"] = o.ElementSpace
	}
	return toSerialize, nil
}

type NullableEventFunctionSpaceType struct {
	value *EventFunctionSpaceType
	isSet bool
}

func (v NullableEventFunctionSpaceType) Get() *EventFunctionSpaceType {
	return v.value
}

func (v *NullableEventFunctionSpaceType) Set(val *EventFunctionSpaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFunctionSpaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFunctionSpaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFunctionSpaceType(val *EventFunctionSpaceType) *NullableEventFunctionSpaceType {
	return &NullableEventFunctionSpaceType{value: val, isSet: true}
}

func (v NullableEventFunctionSpaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFunctionSpaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


