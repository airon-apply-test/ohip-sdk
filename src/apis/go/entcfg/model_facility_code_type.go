/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
)

// checks if the FacilityCodeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacilityCodeType{}

// FacilityCodeType Facility Housekeeping Code, its description and quantity.
type FacilityCodeType struct {
	// Facility Code.
	Description *string `json:"description,omitempty"`
	// Signifies the quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// Facility code value.
	Code *string `json:"code,omitempty"`
}

// NewFacilityCodeType instantiates a new FacilityCodeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacilityCodeType() *FacilityCodeType {
	this := FacilityCodeType{}
	return &this
}

// NewFacilityCodeTypeWithDefaults instantiates a new FacilityCodeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacilityCodeTypeWithDefaults() *FacilityCodeType {
	this := FacilityCodeType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FacilityCodeType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacilityCodeType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FacilityCodeType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FacilityCodeType) SetDescription(v string) {
	o.Description = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *FacilityCodeType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacilityCodeType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *FacilityCodeType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *FacilityCodeType) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FacilityCodeType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacilityCodeType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FacilityCodeType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *FacilityCodeType) SetCode(v string) {
	o.Code = &v
}

func (o FacilityCodeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacilityCodeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableFacilityCodeType struct {
	value *FacilityCodeType
	isSet bool
}

func (v NullableFacilityCodeType) Get() *FacilityCodeType {
	return v.value
}

func (v *NullableFacilityCodeType) Set(val *FacilityCodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableFacilityCodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableFacilityCodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacilityCodeType(val *FacilityCodeType) *NullableFacilityCodeType {
	return &NullableFacilityCodeType{value: val, isSet: true}
}

func (v NullableFacilityCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacilityCodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


