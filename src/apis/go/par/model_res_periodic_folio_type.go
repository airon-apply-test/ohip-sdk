/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResPeriodicFolioType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResPeriodicFolioType{}

// ResPeriodicFolioType Information regarding periodic folios set on the reservation.
type ResPeriodicFolioType struct {
	// Latest date when a direct bill settlement was automatically done using the \"Direct Bill Batch Folios\" option.
	LastSettlementDate *string `json:"lastSettlementDate,omitempty"`
	// Latest date when a folio was printed using the \"Periodic Batch Folios\" option
	LastFolioDate *string `json:"lastFolioDate,omitempty"`
	// Frequency in number of days when folios should be printed for this reservation.
	Frequency *int32 `json:"frequency,omitempty"`
}

// NewResPeriodicFolioType instantiates a new ResPeriodicFolioType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResPeriodicFolioType() *ResPeriodicFolioType {
	this := ResPeriodicFolioType{}
	return &this
}

// NewResPeriodicFolioTypeWithDefaults instantiates a new ResPeriodicFolioType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResPeriodicFolioTypeWithDefaults() *ResPeriodicFolioType {
	this := ResPeriodicFolioType{}
	return &this
}

// GetLastSettlementDate returns the LastSettlementDate field value if set, zero value otherwise.
func (o *ResPeriodicFolioType) GetLastSettlementDate() string {
	if o == nil || IsNil(o.LastSettlementDate) {
		var ret string
		return ret
	}
	return *o.LastSettlementDate
}

// GetLastSettlementDateOk returns a tuple with the LastSettlementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResPeriodicFolioType) GetLastSettlementDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastSettlementDate) {
		return nil, false
	}
	return o.LastSettlementDate, true
}

// HasLastSettlementDate returns a boolean if a field has been set.
func (o *ResPeriodicFolioType) HasLastSettlementDate() bool {
	if o != nil && !IsNil(o.LastSettlementDate) {
		return true
	}

	return false
}

// SetLastSettlementDate gets a reference to the given string and assigns it to the LastSettlementDate field.
func (o *ResPeriodicFolioType) SetLastSettlementDate(v string) {
	o.LastSettlementDate = &v
}

// GetLastFolioDate returns the LastFolioDate field value if set, zero value otherwise.
func (o *ResPeriodicFolioType) GetLastFolioDate() string {
	if o == nil || IsNil(o.LastFolioDate) {
		var ret string
		return ret
	}
	return *o.LastFolioDate
}

// GetLastFolioDateOk returns a tuple with the LastFolioDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResPeriodicFolioType) GetLastFolioDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastFolioDate) {
		return nil, false
	}
	return o.LastFolioDate, true
}

// HasLastFolioDate returns a boolean if a field has been set.
func (o *ResPeriodicFolioType) HasLastFolioDate() bool {
	if o != nil && !IsNil(o.LastFolioDate) {
		return true
	}

	return false
}

// SetLastFolioDate gets a reference to the given string and assigns it to the LastFolioDate field.
func (o *ResPeriodicFolioType) SetLastFolioDate(v string) {
	o.LastFolioDate = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *ResPeriodicFolioType) GetFrequency() int32 {
	if o == nil || IsNil(o.Frequency) {
		var ret int32
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResPeriodicFolioType) GetFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *ResPeriodicFolioType) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int32 and assigns it to the Frequency field.
func (o *ResPeriodicFolioType) SetFrequency(v int32) {
	o.Frequency = &v
}

func (o ResPeriodicFolioType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResPeriodicFolioType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastSettlementDate) {
		toSerialize["lastSettlementDate"] = o.LastSettlementDate
	}
	if !IsNil(o.LastFolioDate) {
		toSerialize["lastFolioDate"] = o.LastFolioDate
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	return toSerialize, nil
}

type NullableResPeriodicFolioType struct {
	value *ResPeriodicFolioType
	isSet bool
}

func (v NullableResPeriodicFolioType) Get() *ResPeriodicFolioType {
	return v.value
}

func (v *NullableResPeriodicFolioType) Set(val *ResPeriodicFolioType) {
	v.value = val
	v.isSet = true
}

func (v NullableResPeriodicFolioType) IsSet() bool {
	return v.isSet
}

func (v *NullableResPeriodicFolioType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResPeriodicFolioType(val *ResPeriodicFolioType) *NullableResPeriodicFolioType {
	return &NullableResPeriodicFolioType{value: val, isSet: true}
}

func (v NullableResPeriodicFolioType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResPeriodicFolioType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


