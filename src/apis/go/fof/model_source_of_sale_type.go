/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the SourceOfSaleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceOfSaleType{}

// SourceOfSaleType Point of Sale of reservation. Identifies the entity/channel who made the reservation.
type SourceOfSaleType struct {
	// Type of entity/channel who made the reservation..
	SourceType *string `json:"sourceType,omitempty"`
	// The entity/channel who made the reservation.
	SourceCode *string `json:"sourceCode,omitempty"`
}

// NewSourceOfSaleType instantiates a new SourceOfSaleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceOfSaleType() *SourceOfSaleType {
	this := SourceOfSaleType{}
	return &this
}

// NewSourceOfSaleTypeWithDefaults instantiates a new SourceOfSaleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceOfSaleTypeWithDefaults() *SourceOfSaleType {
	this := SourceOfSaleType{}
	return &this
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *SourceOfSaleType) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOfSaleType) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *SourceOfSaleType) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *SourceOfSaleType) SetSourceType(v string) {
	o.SourceType = &v
}

// GetSourceCode returns the SourceCode field value if set, zero value otherwise.
func (o *SourceOfSaleType) GetSourceCode() string {
	if o == nil || IsNil(o.SourceCode) {
		var ret string
		return ret
	}
	return *o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOfSaleType) GetSourceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCode) {
		return nil, false
	}
	return o.SourceCode, true
}

// HasSourceCode returns a boolean if a field has been set.
func (o *SourceOfSaleType) HasSourceCode() bool {
	if o != nil && !IsNil(o.SourceCode) {
		return true
	}

	return false
}

// SetSourceCode gets a reference to the given string and assigns it to the SourceCode field.
func (o *SourceOfSaleType) SetSourceCode(v string) {
	o.SourceCode = &v
}

func (o SourceOfSaleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceOfSaleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !IsNil(o.SourceCode) {
		toSerialize["sourceCode"] = o.SourceCode
	}
	return toSerialize, nil
}

type NullableSourceOfSaleType struct {
	value *SourceOfSaleType
	isSet bool
}

func (v NullableSourceOfSaleType) Get() *SourceOfSaleType {
	return v.value
}

func (v *NullableSourceOfSaleType) Set(val *SourceOfSaleType) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceOfSaleType) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceOfSaleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceOfSaleType(val *SourceOfSaleType) *NullableSourceOfSaleType {
	return &NullableSourceOfSaleType{value: val, isSet: true}
}

func (v NullableSourceOfSaleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceOfSaleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


