/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the MiscTraceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiscTraceType{}

// MiscTraceType A collection of retrieved blocks and reservations traces.
type MiscTraceType struct {
	Trace *TraceType `json:"trace,omitempty"`
	BlockInfo *TraceBlockInfoType `json:"blockInfo,omitempty"`
	ReservationInfo *TraceResvInfoType `json:"reservationInfo,omitempty"`
}

// NewMiscTraceType instantiates a new MiscTraceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiscTraceType() *MiscTraceType {
	this := MiscTraceType{}
	return &this
}

// NewMiscTraceTypeWithDefaults instantiates a new MiscTraceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiscTraceTypeWithDefaults() *MiscTraceType {
	this := MiscTraceType{}
	return &this
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *MiscTraceType) GetTrace() TraceType {
	if o == nil || IsNil(o.Trace) {
		var ret TraceType
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiscTraceType) GetTraceOk() (*TraceType, bool) {
	if o == nil || IsNil(o.Trace) {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *MiscTraceType) HasTrace() bool {
	if o != nil && !IsNil(o.Trace) {
		return true
	}

	return false
}

// SetTrace gets a reference to the given TraceType and assigns it to the Trace field.
func (o *MiscTraceType) SetTrace(v TraceType) {
	o.Trace = &v
}

// GetBlockInfo returns the BlockInfo field value if set, zero value otherwise.
func (o *MiscTraceType) GetBlockInfo() TraceBlockInfoType {
	if o == nil || IsNil(o.BlockInfo) {
		var ret TraceBlockInfoType
		return ret
	}
	return *o.BlockInfo
}

// GetBlockInfoOk returns a tuple with the BlockInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiscTraceType) GetBlockInfoOk() (*TraceBlockInfoType, bool) {
	if o == nil || IsNil(o.BlockInfo) {
		return nil, false
	}
	return o.BlockInfo, true
}

// HasBlockInfo returns a boolean if a field has been set.
func (o *MiscTraceType) HasBlockInfo() bool {
	if o != nil && !IsNil(o.BlockInfo) {
		return true
	}

	return false
}

// SetBlockInfo gets a reference to the given TraceBlockInfoType and assigns it to the BlockInfo field.
func (o *MiscTraceType) SetBlockInfo(v TraceBlockInfoType) {
	o.BlockInfo = &v
}

// GetReservationInfo returns the ReservationInfo field value if set, zero value otherwise.
func (o *MiscTraceType) GetReservationInfo() TraceResvInfoType {
	if o == nil || IsNil(o.ReservationInfo) {
		var ret TraceResvInfoType
		return ret
	}
	return *o.ReservationInfo
}

// GetReservationInfoOk returns a tuple with the ReservationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiscTraceType) GetReservationInfoOk() (*TraceResvInfoType, bool) {
	if o == nil || IsNil(o.ReservationInfo) {
		return nil, false
	}
	return o.ReservationInfo, true
}

// HasReservationInfo returns a boolean if a field has been set.
func (o *MiscTraceType) HasReservationInfo() bool {
	if o != nil && !IsNil(o.ReservationInfo) {
		return true
	}

	return false
}

// SetReservationInfo gets a reference to the given TraceResvInfoType and assigns it to the ReservationInfo field.
func (o *MiscTraceType) SetReservationInfo(v TraceResvInfoType) {
	o.ReservationInfo = &v
}

func (o MiscTraceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiscTraceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Trace) {
		toSerialize["trace"] = o.Trace
	}
	if !IsNil(o.BlockInfo) {
		toSerialize["blockInfo"] = o.BlockInfo
	}
	if !IsNil(o.ReservationInfo) {
		toSerialize["reservationInfo"] = o.ReservationInfo
	}
	return toSerialize, nil
}

type NullableMiscTraceType struct {
	value *MiscTraceType
	isSet bool
}

func (v NullableMiscTraceType) Get() *MiscTraceType {
	return v.value
}

func (v *NullableMiscTraceType) Set(val *MiscTraceType) {
	v.value = val
	v.isSet = true
}

func (v NullableMiscTraceType) IsSet() bool {
	return v.isSet
}

func (v *NullableMiscTraceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiscTraceType(val *MiscTraceType) *NullableMiscTraceType {
	return &NullableMiscTraceType{value: val, isSet: true}
}

func (v NullableMiscTraceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiscTraceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


