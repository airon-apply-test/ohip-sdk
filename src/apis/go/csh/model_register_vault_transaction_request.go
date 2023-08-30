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
)

// checks if the RegisterVaultTransactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterVaultTransactionRequest{}

// RegisterVaultTransactionRequest struct for RegisterVaultTransactionRequest
type RegisterVaultTransactionRequest struct {
	HttpTransactionMessage *VaultHTTPTransactionMessageType `json:"httpTransactionMessage,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewRegisterVaultTransactionRequest instantiates a new RegisterVaultTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterVaultTransactionRequest() *RegisterVaultTransactionRequest {
	this := RegisterVaultTransactionRequest{}
	return &this
}

// NewRegisterVaultTransactionRequestWithDefaults instantiates a new RegisterVaultTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterVaultTransactionRequestWithDefaults() *RegisterVaultTransactionRequest {
	this := RegisterVaultTransactionRequest{}
	return &this
}

// GetHttpTransactionMessage returns the HttpTransactionMessage field value if set, zero value otherwise.
func (o *RegisterVaultTransactionRequest) GetHttpTransactionMessage() VaultHTTPTransactionMessageType {
	if o == nil || IsNil(o.HttpTransactionMessage) {
		var ret VaultHTTPTransactionMessageType
		return ret
	}
	return *o.HttpTransactionMessage
}

// GetHttpTransactionMessageOk returns a tuple with the HttpTransactionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterVaultTransactionRequest) GetHttpTransactionMessageOk() (*VaultHTTPTransactionMessageType, bool) {
	if o == nil || IsNil(o.HttpTransactionMessage) {
		return nil, false
	}
	return o.HttpTransactionMessage, true
}

// HasHttpTransactionMessage returns a boolean if a field has been set.
func (o *RegisterVaultTransactionRequest) HasHttpTransactionMessage() bool {
	if o != nil && !IsNil(o.HttpTransactionMessage) {
		return true
	}

	return false
}

// SetHttpTransactionMessage gets a reference to the given VaultHTTPTransactionMessageType and assigns it to the HttpTransactionMessage field.
func (o *RegisterVaultTransactionRequest) SetHttpTransactionMessage(v VaultHTTPTransactionMessageType) {
	o.HttpTransactionMessage = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RegisterVaultTransactionRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterVaultTransactionRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RegisterVaultTransactionRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *RegisterVaultTransactionRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *RegisterVaultTransactionRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterVaultTransactionRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RegisterVaultTransactionRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *RegisterVaultTransactionRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o RegisterVaultTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterVaultTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpTransactionMessage) {
		toSerialize["httpTransactionMessage"] = o.HttpTransactionMessage
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableRegisterVaultTransactionRequest struct {
	value *RegisterVaultTransactionRequest
	isSet bool
}

func (v NullableRegisterVaultTransactionRequest) Get() *RegisterVaultTransactionRequest {
	return v.value
}

func (v *NullableRegisterVaultTransactionRequest) Set(val *RegisterVaultTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterVaultTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterVaultTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterVaultTransactionRequest(val *RegisterVaultTransactionRequest) *NullableRegisterVaultTransactionRequest {
	return &NullableRegisterVaultTransactionRequest{value: val, isSet: true}
}

func (v NullableRegisterVaultTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterVaultTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


