/*
OPERA Cloud Housekeeping Service API

APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CompleteOutOfServiceRoomsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteOutOfServiceRoomsRequest{}

// CompleteOutOfServiceRoomsRequest struct for CompleteOutOfServiceRoomsRequest
type CompleteOutOfServiceRoomsRequest struct {
	Criteria *CompleteRoomRepairCriteria `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCompleteOutOfServiceRoomsRequest instantiates a new CompleteOutOfServiceRoomsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteOutOfServiceRoomsRequest() *CompleteOutOfServiceRoomsRequest {
	this := CompleteOutOfServiceRoomsRequest{}
	return &this
}

// NewCompleteOutOfServiceRoomsRequestWithDefaults instantiates a new CompleteOutOfServiceRoomsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteOutOfServiceRoomsRequestWithDefaults() *CompleteOutOfServiceRoomsRequest {
	this := CompleteOutOfServiceRoomsRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *CompleteOutOfServiceRoomsRequest) GetCriteria() CompleteRoomRepairCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret CompleteRoomRepairCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteOutOfServiceRoomsRequest) GetCriteriaOk() (*CompleteRoomRepairCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *CompleteOutOfServiceRoomsRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given CompleteRoomRepairCriteria and assigns it to the Criteria field.
func (o *CompleteOutOfServiceRoomsRequest) SetCriteria(v CompleteRoomRepairCriteria) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CompleteOutOfServiceRoomsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteOutOfServiceRoomsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CompleteOutOfServiceRoomsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CompleteOutOfServiceRoomsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CompleteOutOfServiceRoomsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteOutOfServiceRoomsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CompleteOutOfServiceRoomsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CompleteOutOfServiceRoomsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CompleteOutOfServiceRoomsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteOutOfServiceRoomsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCompleteOutOfServiceRoomsRequest struct {
	value *CompleteOutOfServiceRoomsRequest
	isSet bool
}

func (v NullableCompleteOutOfServiceRoomsRequest) Get() *CompleteOutOfServiceRoomsRequest {
	return v.value
}

func (v *NullableCompleteOutOfServiceRoomsRequest) Set(val *CompleteOutOfServiceRoomsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteOutOfServiceRoomsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteOutOfServiceRoomsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteOutOfServiceRoomsRequest(val *CompleteOutOfServiceRoomsRequest) *NullableCompleteOutOfServiceRoomsRequest {
	return &NullableCompleteOutOfServiceRoomsRequest{value: val, isSet: true}
}

func (v NullableCompleteOutOfServiceRoomsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteOutOfServiceRoomsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


