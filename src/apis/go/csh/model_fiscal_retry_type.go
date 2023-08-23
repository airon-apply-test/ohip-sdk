/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FiscalRetryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiscalRetryType{}

// FiscalRetryType Generate Fiscal Folio Retry Criteria type to be used for fiscal folio generation
type FiscalRetryType struct {
	// Hotel where the transaction belongs.
	HotelId *string `json:"hotelId,omitempty"`
	// Fiscal Folio sequence ID stored in queue table.
	FolioSeqId *int32 `json:"folioSeqId,omitempty"`
	VoidFolioModes *VoidFolioModes `json:"voidFolioModes,omitempty"`
	FiscalFolioInstruction *FiscalFolioInstruction `json:"fiscalFolioInstruction,omitempty"`
}

// NewFiscalRetryType instantiates a new FiscalRetryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiscalRetryType() *FiscalRetryType {
	this := FiscalRetryType{}
	return &this
}

// NewFiscalRetryTypeWithDefaults instantiates a new FiscalRetryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiscalRetryTypeWithDefaults() *FiscalRetryType {
	this := FiscalRetryType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *FiscalRetryType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalRetryType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *FiscalRetryType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *FiscalRetryType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetFolioSeqId returns the FolioSeqId field value if set, zero value otherwise.
func (o *FiscalRetryType) GetFolioSeqId() int32 {
	if o == nil || IsNil(o.FolioSeqId) {
		var ret int32
		return ret
	}
	return *o.FolioSeqId
}

// GetFolioSeqIdOk returns a tuple with the FolioSeqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalRetryType) GetFolioSeqIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioSeqId) {
		return nil, false
	}
	return o.FolioSeqId, true
}

// HasFolioSeqId returns a boolean if a field has been set.
func (o *FiscalRetryType) HasFolioSeqId() bool {
	if o != nil && !IsNil(o.FolioSeqId) {
		return true
	}

	return false
}

// SetFolioSeqId gets a reference to the given int32 and assigns it to the FolioSeqId field.
func (o *FiscalRetryType) SetFolioSeqId(v int32) {
	o.FolioSeqId = &v
}

// GetVoidFolioModes returns the VoidFolioModes field value if set, zero value otherwise.
func (o *FiscalRetryType) GetVoidFolioModes() VoidFolioModes {
	if o == nil || IsNil(o.VoidFolioModes) {
		var ret VoidFolioModes
		return ret
	}
	return *o.VoidFolioModes
}

// GetVoidFolioModesOk returns a tuple with the VoidFolioModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalRetryType) GetVoidFolioModesOk() (*VoidFolioModes, bool) {
	if o == nil || IsNil(o.VoidFolioModes) {
		return nil, false
	}
	return o.VoidFolioModes, true
}

// HasVoidFolioModes returns a boolean if a field has been set.
func (o *FiscalRetryType) HasVoidFolioModes() bool {
	if o != nil && !IsNil(o.VoidFolioModes) {
		return true
	}

	return false
}

// SetVoidFolioModes gets a reference to the given VoidFolioModes and assigns it to the VoidFolioModes field.
func (o *FiscalRetryType) SetVoidFolioModes(v VoidFolioModes) {
	o.VoidFolioModes = &v
}

// GetFiscalFolioInstruction returns the FiscalFolioInstruction field value if set, zero value otherwise.
func (o *FiscalRetryType) GetFiscalFolioInstruction() FiscalFolioInstruction {
	if o == nil || IsNil(o.FiscalFolioInstruction) {
		var ret FiscalFolioInstruction
		return ret
	}
	return *o.FiscalFolioInstruction
}

// GetFiscalFolioInstructionOk returns a tuple with the FiscalFolioInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalRetryType) GetFiscalFolioInstructionOk() (*FiscalFolioInstruction, bool) {
	if o == nil || IsNil(o.FiscalFolioInstruction) {
		return nil, false
	}
	return o.FiscalFolioInstruction, true
}

// HasFiscalFolioInstruction returns a boolean if a field has been set.
func (o *FiscalRetryType) HasFiscalFolioInstruction() bool {
	if o != nil && !IsNil(o.FiscalFolioInstruction) {
		return true
	}

	return false
}

// SetFiscalFolioInstruction gets a reference to the given FiscalFolioInstruction and assigns it to the FiscalFolioInstruction field.
func (o *FiscalRetryType) SetFiscalFolioInstruction(v FiscalFolioInstruction) {
	o.FiscalFolioInstruction = &v
}

func (o FiscalRetryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiscalRetryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.FolioSeqId) {
		toSerialize["folioSeqId"] = o.FolioSeqId
	}
	if !IsNil(o.VoidFolioModes) {
		toSerialize["voidFolioModes"] = o.VoidFolioModes
	}
	if !IsNil(o.FiscalFolioInstruction) {
		toSerialize["fiscalFolioInstruction"] = o.FiscalFolioInstruction
	}
	return toSerialize, nil
}

type NullableFiscalRetryType struct {
	value *FiscalRetryType
	isSet bool
}

func (v NullableFiscalRetryType) Get() *FiscalRetryType {
	return v.value
}

func (v *NullableFiscalRetryType) Set(val *FiscalRetryType) {
	v.value = val
	v.isSet = true
}

func (v NullableFiscalRetryType) IsSet() bool {
	return v.isSet
}

func (v *NullableFiscalRetryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiscalRetryType(val *FiscalRetryType) *NullableFiscalRetryType {
	return &NullableFiscalRetryType{value: val, isSet: true}
}

func (v NullableFiscalRetryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiscalRetryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


