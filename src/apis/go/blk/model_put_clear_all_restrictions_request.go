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

// checks if the PutClearAllRestrictionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutClearAllRestrictionsRequest{}

// PutClearAllRestrictionsRequest struct for PutClearAllRestrictionsRequest
type PutClearAllRestrictionsRequest struct {
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	HotelId *string `json:"hotelId,omitempty"`
	BlockId *BlockId `json:"blockId,omitempty"`
	// Date for which restrictions should be cleared.
	Date *string `json:"date,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutClearAllRestrictionsRequest instantiates a new PutClearAllRestrictionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutClearAllRestrictionsRequest() *PutClearAllRestrictionsRequest {
	this := PutClearAllRestrictionsRequest{}
	return &this
}

// NewPutClearAllRestrictionsRequestWithDefaults instantiates a new PutClearAllRestrictionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutClearAllRestrictionsRequestWithDefaults() *PutClearAllRestrictionsRequest {
	this := PutClearAllRestrictionsRequest{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *PutClearAllRestrictionsRequest) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClearAllRestrictionsRequest) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *PutClearAllRestrictionsRequest) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *PutClearAllRestrictionsRequest) SetHotelId(v string) {
	o.HotelId = &v
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *PutClearAllRestrictionsRequest) GetBlockId() BlockId {
	if o == nil || IsNil(o.BlockId) {
		var ret BlockId
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClearAllRestrictionsRequest) GetBlockIdOk() (*BlockId, bool) {
	if o == nil || IsNil(o.BlockId) {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *PutClearAllRestrictionsRequest) HasBlockId() bool {
	if o != nil && !IsNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given BlockId and assigns it to the BlockId field.
func (o *PutClearAllRestrictionsRequest) SetBlockId(v BlockId) {
	o.BlockId = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *PutClearAllRestrictionsRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClearAllRestrictionsRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *PutClearAllRestrictionsRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *PutClearAllRestrictionsRequest) SetDate(v string) {
	o.Date = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutClearAllRestrictionsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClearAllRestrictionsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutClearAllRestrictionsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutClearAllRestrictionsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutClearAllRestrictionsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClearAllRestrictionsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutClearAllRestrictionsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutClearAllRestrictionsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutClearAllRestrictionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutClearAllRestrictionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.BlockId) {
		toSerialize["blockId"] = o.BlockId
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutClearAllRestrictionsRequest struct {
	value *PutClearAllRestrictionsRequest
	isSet bool
}

func (v NullablePutClearAllRestrictionsRequest) Get() *PutClearAllRestrictionsRequest {
	return v.value
}

func (v *NullablePutClearAllRestrictionsRequest) Set(val *PutClearAllRestrictionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutClearAllRestrictionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutClearAllRestrictionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutClearAllRestrictionsRequest(val *PutClearAllRestrictionsRequest) *NullablePutClearAllRestrictionsRequest {
	return &NullablePutClearAllRestrictionsRequest{value: val, isSet: true}
}

func (v NullablePutClearAllRestrictionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutClearAllRestrictionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


