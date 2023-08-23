/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ActivityLocationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLocationType{}

// ActivityLocationType Activity Location detail information.
type ActivityLocationType struct {
	// Property to which the activity Location belongs to.
	HotelId *string `json:"hotelId,omitempty"`
	// Code for the activity Location.
	Code *string `json:"code,omitempty"`
	// Description for the Activity Location.
	Description *string `json:"description,omitempty"`
}

// NewActivityLocationType instantiates a new ActivityLocationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLocationType() *ActivityLocationType {
	this := ActivityLocationType{}
	return &this
}

// NewActivityLocationTypeWithDefaults instantiates a new ActivityLocationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLocationTypeWithDefaults() *ActivityLocationType {
	this := ActivityLocationType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ActivityLocationType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLocationType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ActivityLocationType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ActivityLocationType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ActivityLocationType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLocationType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ActivityLocationType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ActivityLocationType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ActivityLocationType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLocationType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ActivityLocationType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ActivityLocationType) SetDescription(v string) {
	o.Description = &v
}

func (o ActivityLocationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLocationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableActivityLocationType struct {
	value *ActivityLocationType
	isSet bool
}

func (v NullableActivityLocationType) Get() *ActivityLocationType {
	return v.value
}

func (v *NullableActivityLocationType) Set(val *ActivityLocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLocationType(val *ActivityLocationType) *NullableActivityLocationType {
	return &NullableActivityLocationType{value: val, isSet: true}
}

func (v NullableActivityLocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


