/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CateringEventLinkType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CateringEventLinkType{}

// CateringEventLinkType struct for CateringEventLinkType
type CateringEventLinkType struct {
	Id *UniqueIDType `json:"id,omitempty"`
	Type *CateringEventLinkTypeType `json:"type,omitempty"`
	MasterEventDates *DateTimeSpanType `json:"masterEventDates,omitempty"`
}

// NewCateringEventLinkType instantiates a new CateringEventLinkType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCateringEventLinkType() *CateringEventLinkType {
	this := CateringEventLinkType{}
	return &this
}

// NewCateringEventLinkTypeWithDefaults instantiates a new CateringEventLinkType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCateringEventLinkTypeWithDefaults() *CateringEventLinkType {
	this := CateringEventLinkType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CateringEventLinkType) GetId() UniqueIDType {
	if o == nil || IsNil(o.Id) {
		var ret UniqueIDType
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventLinkType) GetIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CateringEventLinkType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given UniqueIDType and assigns it to the Id field.
func (o *CateringEventLinkType) SetId(v UniqueIDType) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CateringEventLinkType) GetType() CateringEventLinkTypeType {
	if o == nil || IsNil(o.Type) {
		var ret CateringEventLinkTypeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventLinkType) GetTypeOk() (*CateringEventLinkTypeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CateringEventLinkType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CateringEventLinkTypeType and assigns it to the Type field.
func (o *CateringEventLinkType) SetType(v CateringEventLinkTypeType) {
	o.Type = &v
}

// GetMasterEventDates returns the MasterEventDates field value if set, zero value otherwise.
func (o *CateringEventLinkType) GetMasterEventDates() DateTimeSpanType {
	if o == nil || IsNil(o.MasterEventDates) {
		var ret DateTimeSpanType
		return ret
	}
	return *o.MasterEventDates
}

// GetMasterEventDatesOk returns a tuple with the MasterEventDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventLinkType) GetMasterEventDatesOk() (*DateTimeSpanType, bool) {
	if o == nil || IsNil(o.MasterEventDates) {
		return nil, false
	}
	return o.MasterEventDates, true
}

// HasMasterEventDates returns a boolean if a field has been set.
func (o *CateringEventLinkType) HasMasterEventDates() bool {
	if o != nil && !IsNil(o.MasterEventDates) {
		return true
	}

	return false
}

// SetMasterEventDates gets a reference to the given DateTimeSpanType and assigns it to the MasterEventDates field.
func (o *CateringEventLinkType) SetMasterEventDates(v DateTimeSpanType) {
	o.MasterEventDates = &v
}

func (o CateringEventLinkType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CateringEventLinkType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.MasterEventDates) {
		toSerialize["masterEventDates"] = o.MasterEventDates
	}
	return toSerialize, nil
}

type NullableCateringEventLinkType struct {
	value *CateringEventLinkType
	isSet bool
}

func (v NullableCateringEventLinkType) Get() *CateringEventLinkType {
	return v.value
}

func (v *NullableCateringEventLinkType) Set(val *CateringEventLinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableCateringEventLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableCateringEventLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCateringEventLinkType(val *CateringEventLinkType) *NullableCateringEventLinkType {
	return &NullableCateringEventLinkType{value: val, isSet: true}
}

func (v NullableCateringEventLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCateringEventLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


