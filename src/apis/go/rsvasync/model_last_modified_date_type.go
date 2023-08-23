/*
OPERA Cloud Reservation Asynchronous API

APIs to cater for Reservation Asynchronous functionality in OPERA Cloud. This includes viewing reservation data along with its revenue. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.4.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the LastModifiedDateType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastModifiedDateType{}

// LastModifiedDateType Start and End Modified Dates of Reservations after which the results are to be fetched
type LastModifiedDateType struct {
	StartLastModifiedDate *time.Time `json:"startLastModifiedDate,omitempty"`
	EndLastModifiedDate *time.Time `json:"endLastModifiedDate,omitempty"`
}

// NewLastModifiedDateType instantiates a new LastModifiedDateType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastModifiedDateType() *LastModifiedDateType {
	this := LastModifiedDateType{}
	return &this
}

// NewLastModifiedDateTypeWithDefaults instantiates a new LastModifiedDateType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastModifiedDateTypeWithDefaults() *LastModifiedDateType {
	this := LastModifiedDateType{}
	return &this
}

// GetStartLastModifiedDate returns the StartLastModifiedDate field value if set, zero value otherwise.
func (o *LastModifiedDateType) GetStartLastModifiedDate() time.Time {
	if o == nil || IsNil(o.StartLastModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.StartLastModifiedDate
}

// GetStartLastModifiedDateOk returns a tuple with the StartLastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastModifiedDateType) GetStartLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartLastModifiedDate) {
		return nil, false
	}
	return o.StartLastModifiedDate, true
}

// HasStartLastModifiedDate returns a boolean if a field has been set.
func (o *LastModifiedDateType) HasStartLastModifiedDate() bool {
	if o != nil && !IsNil(o.StartLastModifiedDate) {
		return true
	}

	return false
}

// SetStartLastModifiedDate gets a reference to the given time.Time and assigns it to the StartLastModifiedDate field.
func (o *LastModifiedDateType) SetStartLastModifiedDate(v time.Time) {
	o.StartLastModifiedDate = &v
}

// GetEndLastModifiedDate returns the EndLastModifiedDate field value if set, zero value otherwise.
func (o *LastModifiedDateType) GetEndLastModifiedDate() time.Time {
	if o == nil || IsNil(o.EndLastModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.EndLastModifiedDate
}

// GetEndLastModifiedDateOk returns a tuple with the EndLastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastModifiedDateType) GetEndLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndLastModifiedDate) {
		return nil, false
	}
	return o.EndLastModifiedDate, true
}

// HasEndLastModifiedDate returns a boolean if a field has been set.
func (o *LastModifiedDateType) HasEndLastModifiedDate() bool {
	if o != nil && !IsNil(o.EndLastModifiedDate) {
		return true
	}

	return false
}

// SetEndLastModifiedDate gets a reference to the given time.Time and assigns it to the EndLastModifiedDate field.
func (o *LastModifiedDateType) SetEndLastModifiedDate(v time.Time) {
	o.EndLastModifiedDate = &v
}

func (o LastModifiedDateType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastModifiedDateType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartLastModifiedDate) {
		toSerialize["startLastModifiedDate"] = o.StartLastModifiedDate
	}
	if !IsNil(o.EndLastModifiedDate) {
		toSerialize["endLastModifiedDate"] = o.EndLastModifiedDate
	}
	return toSerialize, nil
}

type NullableLastModifiedDateType struct {
	value *LastModifiedDateType
	isSet bool
}

func (v NullableLastModifiedDateType) Get() *LastModifiedDateType {
	return v.value
}

func (v *NullableLastModifiedDateType) Set(val *LastModifiedDateType) {
	v.value = val
	v.isSet = true
}

func (v NullableLastModifiedDateType) IsSet() bool {
	return v.isSet
}

func (v *NullableLastModifiedDateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastModifiedDateType(val *LastModifiedDateType) *NullableLastModifiedDateType {
	return &NullableLastModifiedDateType{value: val, isSet: true}
}

func (v NullableLastModifiedDateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastModifiedDateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


