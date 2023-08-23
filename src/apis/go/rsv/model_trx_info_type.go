/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TrxInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrxInfoType{}

// TrxInfoType Transaction codes info.
type TrxInfoType struct {
	// Transaction codes info.
	Description *string `json:"description,omitempty"`
	// Contains service type for transaction code.
	TrxServiceType *string `json:"trxServiceType,omitempty"`
	// Unique identifier for the Transaction code.
	TransactionCode *string `json:"transactionCode,omitempty"`
	// Hotel context of the Transaction code.
	HotelId *string `json:"hotelId,omitempty"`
	// Print receipt flag that tells whether the transaction receipt is to be printed or not. This is based on the transaction code.
	PrintTrxReceipt *bool `json:"printTrxReceipt,omitempty"`
}

// NewTrxInfoType instantiates a new TrxInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrxInfoType() *TrxInfoType {
	this := TrxInfoType{}
	return &this
}

// NewTrxInfoTypeWithDefaults instantiates a new TrxInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrxInfoTypeWithDefaults() *TrxInfoType {
	this := TrxInfoType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TrxInfoType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TrxInfoType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TrxInfoType) SetDescription(v string) {
	o.Description = &v
}

// GetTrxServiceType returns the TrxServiceType field value if set, zero value otherwise.
func (o *TrxInfoType) GetTrxServiceType() string {
	if o == nil || IsNil(o.TrxServiceType) {
		var ret string
		return ret
	}
	return *o.TrxServiceType
}

// GetTrxServiceTypeOk returns a tuple with the TrxServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetTrxServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrxServiceType) {
		return nil, false
	}
	return o.TrxServiceType, true
}

// HasTrxServiceType returns a boolean if a field has been set.
func (o *TrxInfoType) HasTrxServiceType() bool {
	if o != nil && !IsNil(o.TrxServiceType) {
		return true
	}

	return false
}

// SetTrxServiceType gets a reference to the given string and assigns it to the TrxServiceType field.
func (o *TrxInfoType) SetTrxServiceType(v string) {
	o.TrxServiceType = &v
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise.
func (o *TrxInfoType) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode) {
		var ret string
		return ret
	}
	return *o.TransactionCode
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetTransactionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionCode) {
		return nil, false
	}
	return o.TransactionCode, true
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *TrxInfoType) HasTransactionCode() bool {
	if o != nil && !IsNil(o.TransactionCode) {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given string and assigns it to the TransactionCode field.
func (o *TrxInfoType) SetTransactionCode(v string) {
	o.TransactionCode = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *TrxInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *TrxInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *TrxInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetPrintTrxReceipt returns the PrintTrxReceipt field value if set, zero value otherwise.
func (o *TrxInfoType) GetPrintTrxReceipt() bool {
	if o == nil || IsNil(o.PrintTrxReceipt) {
		var ret bool
		return ret
	}
	return *o.PrintTrxReceipt
}

// GetPrintTrxReceiptOk returns a tuple with the PrintTrxReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetPrintTrxReceiptOk() (*bool, bool) {
	if o == nil || IsNil(o.PrintTrxReceipt) {
		return nil, false
	}
	return o.PrintTrxReceipt, true
}

// HasPrintTrxReceipt returns a boolean if a field has been set.
func (o *TrxInfoType) HasPrintTrxReceipt() bool {
	if o != nil && !IsNil(o.PrintTrxReceipt) {
		return true
	}

	return false
}

// SetPrintTrxReceipt gets a reference to the given bool and assigns it to the PrintTrxReceipt field.
func (o *TrxInfoType) SetPrintTrxReceipt(v bool) {
	o.PrintTrxReceipt = &v
}

func (o TrxInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrxInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TrxServiceType) {
		toSerialize["trxServiceType"] = o.TrxServiceType
	}
	if !IsNil(o.TransactionCode) {
		toSerialize["transactionCode"] = o.TransactionCode
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.PrintTrxReceipt) {
		toSerialize["printTrxReceipt"] = o.PrintTrxReceipt
	}
	return toSerialize, nil
}

type NullableTrxInfoType struct {
	value *TrxInfoType
	isSet bool
}

func (v NullableTrxInfoType) Get() *TrxInfoType {
	return v.value
}

func (v *NullableTrxInfoType) Set(val *TrxInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTrxInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTrxInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrxInfoType(val *TrxInfoType) *NullableTrxInfoType {
	return &NullableTrxInfoType{value: val, isSet: true}
}

func (v NullableTrxInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrxInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


