/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
	"time"
)

// checks if the AnonymizationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnonymizationType{}

// AnonymizationType Provides information about guest's anonymization status.
type AnonymizationType struct {
	AnonymizationStatus *AnonymizationStatusType `json:"anonymizationStatus,omitempty"`
	// Date and Time when the guest was anonymized.
	AnonymizationDate *time.Time `json:"anonymizationDate,omitempty"`
}

// NewAnonymizationType instantiates a new AnonymizationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnonymizationType() *AnonymizationType {
	this := AnonymizationType{}
	return &this
}

// NewAnonymizationTypeWithDefaults instantiates a new AnonymizationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnonymizationTypeWithDefaults() *AnonymizationType {
	this := AnonymizationType{}
	return &this
}

// GetAnonymizationStatus returns the AnonymizationStatus field value if set, zero value otherwise.
func (o *AnonymizationType) GetAnonymizationStatus() AnonymizationStatusType {
	if o == nil || IsNil(o.AnonymizationStatus) {
		var ret AnonymizationStatusType
		return ret
	}
	return *o.AnonymizationStatus
}

// GetAnonymizationStatusOk returns a tuple with the AnonymizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnonymizationType) GetAnonymizationStatusOk() (*AnonymizationStatusType, bool) {
	if o == nil || IsNil(o.AnonymizationStatus) {
		return nil, false
	}
	return o.AnonymizationStatus, true
}

// HasAnonymizationStatus returns a boolean if a field has been set.
func (o *AnonymizationType) HasAnonymizationStatus() bool {
	if o != nil && !IsNil(o.AnonymizationStatus) {
		return true
	}

	return false
}

// SetAnonymizationStatus gets a reference to the given AnonymizationStatusType and assigns it to the AnonymizationStatus field.
func (o *AnonymizationType) SetAnonymizationStatus(v AnonymizationStatusType) {
	o.AnonymizationStatus = &v
}

// GetAnonymizationDate returns the AnonymizationDate field value if set, zero value otherwise.
func (o *AnonymizationType) GetAnonymizationDate() time.Time {
	if o == nil || IsNil(o.AnonymizationDate) {
		var ret time.Time
		return ret
	}
	return *o.AnonymizationDate
}

// GetAnonymizationDateOk returns a tuple with the AnonymizationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnonymizationType) GetAnonymizationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AnonymizationDate) {
		return nil, false
	}
	return o.AnonymizationDate, true
}

// HasAnonymizationDate returns a boolean if a field has been set.
func (o *AnonymizationType) HasAnonymizationDate() bool {
	if o != nil && !IsNil(o.AnonymizationDate) {
		return true
	}

	return false
}

// SetAnonymizationDate gets a reference to the given time.Time and assigns it to the AnonymizationDate field.
func (o *AnonymizationType) SetAnonymizationDate(v time.Time) {
	o.AnonymizationDate = &v
}

func (o AnonymizationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnonymizationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnonymizationStatus) {
		toSerialize["anonymizationStatus"] = o.AnonymizationStatus
	}
	if !IsNil(o.AnonymizationDate) {
		toSerialize["anonymizationDate"] = o.AnonymizationDate
	}
	return toSerialize, nil
}

type NullableAnonymizationType struct {
	value *AnonymizationType
	isSet bool
}

func (v NullableAnonymizationType) Get() *AnonymizationType {
	return v.value
}

func (v *NullableAnonymizationType) Set(val *AnonymizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnonymizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnonymizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnonymizationType(val *AnonymizationType) *NullableAnonymizationType {
	return &NullableAnonymizationType{value: val, isSet: true}
}

func (v NullableAnonymizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnonymizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


