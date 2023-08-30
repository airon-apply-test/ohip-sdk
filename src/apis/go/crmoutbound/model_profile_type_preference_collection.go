/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the ProfileTypePreferenceCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileTypePreferenceCollection{}

// ProfileTypePreferenceCollection List of customer preferences.
type ProfileTypePreferenceCollection struct {
	// Collection of Detailed information on preferences of the customer.
	PreferenceType []PreferenceTypeType `json:"preferenceType,omitempty"`
	// Total number of rows queried
	TotalRows *int32 `json:"totalRows,omitempty"`
}

// NewProfileTypePreferenceCollection instantiates a new ProfileTypePreferenceCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileTypePreferenceCollection() *ProfileTypePreferenceCollection {
	this := ProfileTypePreferenceCollection{}
	return &this
}

// NewProfileTypePreferenceCollectionWithDefaults instantiates a new ProfileTypePreferenceCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTypePreferenceCollectionWithDefaults() *ProfileTypePreferenceCollection {
	this := ProfileTypePreferenceCollection{}
	return &this
}

// GetPreferenceType returns the PreferenceType field value if set, zero value otherwise.
func (o *ProfileTypePreferenceCollection) GetPreferenceType() []PreferenceTypeType {
	if o == nil || IsNil(o.PreferenceType) {
		var ret []PreferenceTypeType
		return ret
	}
	return o.PreferenceType
}

// GetPreferenceTypeOk returns a tuple with the PreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypePreferenceCollection) GetPreferenceTypeOk() ([]PreferenceTypeType, bool) {
	if o == nil || IsNil(o.PreferenceType) {
		return nil, false
	}
	return o.PreferenceType, true
}

// HasPreferenceType returns a boolean if a field has been set.
func (o *ProfileTypePreferenceCollection) HasPreferenceType() bool {
	if o != nil && !IsNil(o.PreferenceType) {
		return true
	}

	return false
}

// SetPreferenceType gets a reference to the given []PreferenceTypeType and assigns it to the PreferenceType field.
func (o *ProfileTypePreferenceCollection) SetPreferenceType(v []PreferenceTypeType) {
	o.PreferenceType = v
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise.
func (o *ProfileTypePreferenceCollection) GetTotalRows() int32 {
	if o == nil || IsNil(o.TotalRows) {
		var ret int32
		return ret
	}
	return *o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypePreferenceCollection) GetTotalRowsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRows) {
		return nil, false
	}
	return o.TotalRows, true
}

// HasTotalRows returns a boolean if a field has been set.
func (o *ProfileTypePreferenceCollection) HasTotalRows() bool {
	if o != nil && !IsNil(o.TotalRows) {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given int32 and assigns it to the TotalRows field.
func (o *ProfileTypePreferenceCollection) SetTotalRows(v int32) {
	o.TotalRows = &v
}

func (o ProfileTypePreferenceCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileTypePreferenceCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferenceType) {
		toSerialize["preferenceType"] = o.PreferenceType
	}
	if !IsNil(o.TotalRows) {
		toSerialize["totalRows"] = o.TotalRows
	}
	return toSerialize, nil
}

type NullableProfileTypePreferenceCollection struct {
	value *ProfileTypePreferenceCollection
	isSet bool
}

func (v NullableProfileTypePreferenceCollection) Get() *ProfileTypePreferenceCollection {
	return v.value
}

func (v *NullableProfileTypePreferenceCollection) Set(val *ProfileTypePreferenceCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTypePreferenceCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTypePreferenceCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTypePreferenceCollection(val *ProfileTypePreferenceCollection) *NullableProfileTypePreferenceCollection {
	return &NullableProfileTypePreferenceCollection{value: val, isSet: true}
}

func (v NullableProfileTypePreferenceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTypePreferenceCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


