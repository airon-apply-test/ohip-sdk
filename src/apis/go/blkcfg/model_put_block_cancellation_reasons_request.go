/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blkcfg

import (
	"encoding/json"
)

// checks if the PutBlockCancellationReasonsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBlockCancellationReasonsRequest{}

// PutBlockCancellationReasonsRequest struct for PutBlockCancellationReasonsRequest
type PutBlockCancellationReasonsRequest struct {
	// List of Block Cancellation Reasons.
	BlockCancellationReasons []BlockCancellationReasonType `json:"blockCancellationReasons,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutBlockCancellationReasonsRequest instantiates a new PutBlockCancellationReasonsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBlockCancellationReasonsRequest() *PutBlockCancellationReasonsRequest {
	this := PutBlockCancellationReasonsRequest{}
	return &this
}

// NewPutBlockCancellationReasonsRequestWithDefaults instantiates a new PutBlockCancellationReasonsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBlockCancellationReasonsRequestWithDefaults() *PutBlockCancellationReasonsRequest {
	this := PutBlockCancellationReasonsRequest{}
	return &this
}

// GetBlockCancellationReasons returns the BlockCancellationReasons field value if set, zero value otherwise.
func (o *PutBlockCancellationReasonsRequest) GetBlockCancellationReasons() []BlockCancellationReasonType {
	if o == nil || IsNil(o.BlockCancellationReasons) {
		var ret []BlockCancellationReasonType
		return ret
	}
	return o.BlockCancellationReasons
}

// GetBlockCancellationReasonsOk returns a tuple with the BlockCancellationReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBlockCancellationReasonsRequest) GetBlockCancellationReasonsOk() ([]BlockCancellationReasonType, bool) {
	if o == nil || IsNil(o.BlockCancellationReasons) {
		return nil, false
	}
	return o.BlockCancellationReasons, true
}

// HasBlockCancellationReasons returns a boolean if a field has been set.
func (o *PutBlockCancellationReasonsRequest) HasBlockCancellationReasons() bool {
	if o != nil && !IsNil(o.BlockCancellationReasons) {
		return true
	}

	return false
}

// SetBlockCancellationReasons gets a reference to the given []BlockCancellationReasonType and assigns it to the BlockCancellationReasons field.
func (o *PutBlockCancellationReasonsRequest) SetBlockCancellationReasons(v []BlockCancellationReasonType) {
	o.BlockCancellationReasons = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutBlockCancellationReasonsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBlockCancellationReasonsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutBlockCancellationReasonsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutBlockCancellationReasonsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutBlockCancellationReasonsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBlockCancellationReasonsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutBlockCancellationReasonsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutBlockCancellationReasonsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutBlockCancellationReasonsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBlockCancellationReasonsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockCancellationReasons) {
		toSerialize["blockCancellationReasons"] = o.BlockCancellationReasons
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutBlockCancellationReasonsRequest struct {
	value *PutBlockCancellationReasonsRequest
	isSet bool
}

func (v NullablePutBlockCancellationReasonsRequest) Get() *PutBlockCancellationReasonsRequest {
	return v.value
}

func (v *NullablePutBlockCancellationReasonsRequest) Set(val *PutBlockCancellationReasonsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBlockCancellationReasonsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBlockCancellationReasonsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBlockCancellationReasonsRequest(val *PutBlockCancellationReasonsRequest) *NullablePutBlockCancellationReasonsRequest {
	return &NullablePutBlockCancellationReasonsRequest{value: val, isSet: true}
}

func (v NullablePutBlockCancellationReasonsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBlockCancellationReasonsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


