/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the NumericUDFType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NumericUDFType{}

// NumericUDFType Used to hold user defined field of Numeric Type. It is highly recommended to use UDFN01, UDFN02,...UDFN40 (Total 40) as Numeric UDF names(commonly used on Reservation, Profile etc.). Name is not restricted using enumeration, to provide flexibility of different name usage if required.
type NumericUDFType struct {
	// Name of user defined field.
	Name *string `json:"name,omitempty"`
	// Value of user defined field.
	Value *float32 `json:"value,omitempty"`
	// Label of user defined field used by vendors or customers.
	Altname *string `json:"altname,omitempty"`
}

// NewNumericUDFType instantiates a new NumericUDFType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumericUDFType() *NumericUDFType {
	this := NumericUDFType{}
	return &this
}

// NewNumericUDFTypeWithDefaults instantiates a new NumericUDFType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumericUDFTypeWithDefaults() *NumericUDFType {
	this := NumericUDFType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NumericUDFType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumericUDFType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NumericUDFType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NumericUDFType) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NumericUDFType) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumericUDFType) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NumericUDFType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *NumericUDFType) SetValue(v float32) {
	o.Value = &v
}

// GetAltname returns the Altname field value if set, zero value otherwise.
func (o *NumericUDFType) GetAltname() string {
	if o == nil || IsNil(o.Altname) {
		var ret string
		return ret
	}
	return *o.Altname
}

// GetAltnameOk returns a tuple with the Altname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumericUDFType) GetAltnameOk() (*string, bool) {
	if o == nil || IsNil(o.Altname) {
		return nil, false
	}
	return o.Altname, true
}

// HasAltname returns a boolean if a field has been set.
func (o *NumericUDFType) HasAltname() bool {
	if o != nil && !IsNil(o.Altname) {
		return true
	}

	return false
}

// SetAltname gets a reference to the given string and assigns it to the Altname field.
func (o *NumericUDFType) SetAltname(v string) {
	o.Altname = &v
}

func (o NumericUDFType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumericUDFType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Altname) {
		toSerialize["altname"] = o.Altname
	}
	return toSerialize, nil
}

type NullableNumericUDFType struct {
	value *NumericUDFType
	isSet bool
}

func (v NullableNumericUDFType) Get() *NumericUDFType {
	return v.value
}

func (v *NullableNumericUDFType) Set(val *NumericUDFType) {
	v.value = val
	v.isSet = true
}

func (v NullableNumericUDFType) IsSet() bool {
	return v.isSet
}

func (v *NullableNumericUDFType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumericUDFType(val *NumericUDFType) *NullableNumericUDFType {
	return &NullableNumericUDFType{value: val, isSet: true}
}

func (v NullableNumericUDFType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumericUDFType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


