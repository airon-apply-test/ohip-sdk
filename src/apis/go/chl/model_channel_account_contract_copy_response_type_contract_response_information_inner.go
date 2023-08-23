/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelAccountContractCopyResponseTypeContractResponseInformationInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelAccountContractCopyResponseTypeContractResponseInformationInner{}

// ChannelAccountContractCopyResponseTypeContractResponseInformationInner struct for ChannelAccountContractCopyResponseTypeContractResponseInformationInner
type ChannelAccountContractCopyResponseTypeContractResponseInformationInner struct {
	ContractId *UniqueIDType `json:"contractId,omitempty"`
	// Holds contract No of the contract which is copied.
	SourceContractNo *string `json:"sourceContractNo,omitempty"`
	// Holds contract No of the new contract created.
	TargetContractNo *string `json:"targetContractNo,omitempty"`
	ErrorMessage *ErrorType `json:"errorMessage,omitempty"`
	// Flag to identify contract is copied successfully.
	Success *bool `json:"success,omitempty"`
}

// NewChannelAccountContractCopyResponseTypeContractResponseInformationInner instantiates a new ChannelAccountContractCopyResponseTypeContractResponseInformationInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelAccountContractCopyResponseTypeContractResponseInformationInner() *ChannelAccountContractCopyResponseTypeContractResponseInformationInner {
	this := ChannelAccountContractCopyResponseTypeContractResponseInformationInner{}
	return &this
}

// NewChannelAccountContractCopyResponseTypeContractResponseInformationInnerWithDefaults instantiates a new ChannelAccountContractCopyResponseTypeContractResponseInformationInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelAccountContractCopyResponseTypeContractResponseInformationInnerWithDefaults() *ChannelAccountContractCopyResponseTypeContractResponseInformationInner {
	this := ChannelAccountContractCopyResponseTypeContractResponseInformationInner{}
	return &this
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetContractId() UniqueIDType {
	if o == nil || IsNil(o.ContractId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetContractIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given UniqueIDType and assigns it to the ContractId field.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) SetContractId(v UniqueIDType) {
	o.ContractId = &v
}

// GetSourceContractNo returns the SourceContractNo field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetSourceContractNo() string {
	if o == nil || IsNil(o.SourceContractNo) {
		var ret string
		return ret
	}
	return *o.SourceContractNo
}

// GetSourceContractNoOk returns a tuple with the SourceContractNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetSourceContractNoOk() (*string, bool) {
	if o == nil || IsNil(o.SourceContractNo) {
		return nil, false
	}
	return o.SourceContractNo, true
}

// HasSourceContractNo returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) HasSourceContractNo() bool {
	if o != nil && !IsNil(o.SourceContractNo) {
		return true
	}

	return false
}

// SetSourceContractNo gets a reference to the given string and assigns it to the SourceContractNo field.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) SetSourceContractNo(v string) {
	o.SourceContractNo = &v
}

// GetTargetContractNo returns the TargetContractNo field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetTargetContractNo() string {
	if o == nil || IsNil(o.TargetContractNo) {
		var ret string
		return ret
	}
	return *o.TargetContractNo
}

// GetTargetContractNoOk returns a tuple with the TargetContractNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetTargetContractNoOk() (*string, bool) {
	if o == nil || IsNil(o.TargetContractNo) {
		return nil, false
	}
	return o.TargetContractNo, true
}

// HasTargetContractNo returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) HasTargetContractNo() bool {
	if o != nil && !IsNil(o.TargetContractNo) {
		return true
	}

	return false
}

// SetTargetContractNo gets a reference to the given string and assigns it to the TargetContractNo field.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) SetTargetContractNo(v string) {
	o.TargetContractNo = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetErrorMessage() ErrorType {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret ErrorType
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetErrorMessageOk() (*ErrorType, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given ErrorType and assigns it to the ErrorMessage field.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) SetErrorMessage(v ErrorType) {
	o.ErrorMessage = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) SetSuccess(v bool) {
	o.Success = &v
}

func (o ChannelAccountContractCopyResponseTypeContractResponseInformationInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelAccountContractCopyResponseTypeContractResponseInformationInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	if !IsNil(o.SourceContractNo) {
		toSerialize["sourceContractNo"] = o.SourceContractNo
	}
	if !IsNil(o.TargetContractNo) {
		toSerialize["targetContractNo"] = o.TargetContractNo
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner struct {
	value *ChannelAccountContractCopyResponseTypeContractResponseInformationInner
	isSet bool
}

func (v NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner) Get() *ChannelAccountContractCopyResponseTypeContractResponseInformationInner {
	return v.value
}

func (v *NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner) Set(val *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelAccountContractCopyResponseTypeContractResponseInformationInner(val *ChannelAccountContractCopyResponseTypeContractResponseInformationInner) *NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner {
	return &NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner{value: val, isSet: true}
}

func (v NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelAccountContractCopyResponseTypeContractResponseInformationInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


