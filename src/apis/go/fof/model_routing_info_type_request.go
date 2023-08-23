/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RoutingInfoTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingInfoTypeRequest{}

// RoutingInfoTypeRequest Comp Accounting Request routing
type RoutingInfoTypeRequest struct {
	CompRequestInfo *CompRoutingRequestType `json:"compRequestInfo,omitempty"`
	// Set of routing instructions associated to this routing type.
	Instructions []RoutingInstructionType `json:"instructions,omitempty"`
}

// NewRoutingInfoTypeRequest instantiates a new RoutingInfoTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingInfoTypeRequest() *RoutingInfoTypeRequest {
	this := RoutingInfoTypeRequest{}
	return &this
}

// NewRoutingInfoTypeRequestWithDefaults instantiates a new RoutingInfoTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingInfoTypeRequestWithDefaults() *RoutingInfoTypeRequest {
	this := RoutingInfoTypeRequest{}
	return &this
}

// GetCompRequestInfo returns the CompRequestInfo field value if set, zero value otherwise.
func (o *RoutingInfoTypeRequest) GetCompRequestInfo() CompRoutingRequestType {
	if o == nil || IsNil(o.CompRequestInfo) {
		var ret CompRoutingRequestType
		return ret
	}
	return *o.CompRequestInfo
}

// GetCompRequestInfoOk returns a tuple with the CompRequestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoTypeRequest) GetCompRequestInfoOk() (*CompRoutingRequestType, bool) {
	if o == nil || IsNil(o.CompRequestInfo) {
		return nil, false
	}
	return o.CompRequestInfo, true
}

// HasCompRequestInfo returns a boolean if a field has been set.
func (o *RoutingInfoTypeRequest) HasCompRequestInfo() bool {
	if o != nil && !IsNil(o.CompRequestInfo) {
		return true
	}

	return false
}

// SetCompRequestInfo gets a reference to the given CompRoutingRequestType and assigns it to the CompRequestInfo field.
func (o *RoutingInfoTypeRequest) SetCompRequestInfo(v CompRoutingRequestType) {
	o.CompRequestInfo = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *RoutingInfoTypeRequest) GetInstructions() []RoutingInstructionType {
	if o == nil || IsNil(o.Instructions) {
		var ret []RoutingInstructionType
		return ret
	}
	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoTypeRequest) GetInstructionsOk() ([]RoutingInstructionType, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *RoutingInfoTypeRequest) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []RoutingInstructionType and assigns it to the Instructions field.
func (o *RoutingInfoTypeRequest) SetInstructions(v []RoutingInstructionType) {
	o.Instructions = v
}

func (o RoutingInfoTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingInfoTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompRequestInfo) {
		toSerialize["compRequestInfo"] = o.CompRequestInfo
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	return toSerialize, nil
}

type NullableRoutingInfoTypeRequest struct {
	value *RoutingInfoTypeRequest
	isSet bool
}

func (v NullableRoutingInfoTypeRequest) Get() *RoutingInfoTypeRequest {
	return v.value
}

func (v *NullableRoutingInfoTypeRequest) Set(val *RoutingInfoTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingInfoTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingInfoTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingInfoTypeRequest(val *RoutingInfoTypeRequest) *NullableRoutingInfoTypeRequest {
	return &NullableRoutingInfoTypeRequest{value: val, isSet: true}
}

func (v NullableRoutingInfoTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingInfoTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


