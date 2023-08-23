/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CompPostingsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompPostingsType{}

// CompPostingsType Collection of comp postings.
type CompPostingsType struct {
	// Authorizer name of the Comp Account.
	Authorizer *string `json:"authorizer,omitempty"`
	// Approval status of the comp account.
	ApprovalStatus *string `json:"approvalStatus,omitempty"`
}

// NewCompPostingsType instantiates a new CompPostingsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompPostingsType() *CompPostingsType {
	this := CompPostingsType{}
	return &this
}

// NewCompPostingsTypeWithDefaults instantiates a new CompPostingsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompPostingsTypeWithDefaults() *CompPostingsType {
	this := CompPostingsType{}
	return &this
}

// GetAuthorizer returns the Authorizer field value if set, zero value otherwise.
func (o *CompPostingsType) GetAuthorizer() string {
	if o == nil || IsNil(o.Authorizer) {
		var ret string
		return ret
	}
	return *o.Authorizer
}

// GetAuthorizerOk returns a tuple with the Authorizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompPostingsType) GetAuthorizerOk() (*string, bool) {
	if o == nil || IsNil(o.Authorizer) {
		return nil, false
	}
	return o.Authorizer, true
}

// HasAuthorizer returns a boolean if a field has been set.
func (o *CompPostingsType) HasAuthorizer() bool {
	if o != nil && !IsNil(o.Authorizer) {
		return true
	}

	return false
}

// SetAuthorizer gets a reference to the given string and assigns it to the Authorizer field.
func (o *CompPostingsType) SetAuthorizer(v string) {
	o.Authorizer = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *CompPostingsType) GetApprovalStatus() string {
	if o == nil || IsNil(o.ApprovalStatus) {
		var ret string
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompPostingsType) GetApprovalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *CompPostingsType) HasApprovalStatus() bool {
	if o != nil && !IsNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given string and assigns it to the ApprovalStatus field.
func (o *CompPostingsType) SetApprovalStatus(v string) {
	o.ApprovalStatus = &v
}

func (o CompPostingsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompPostingsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authorizer) {
		toSerialize["authorizer"] = o.Authorizer
	}
	if !IsNil(o.ApprovalStatus) {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	return toSerialize, nil
}

type NullableCompPostingsType struct {
	value *CompPostingsType
	isSet bool
}

func (v NullableCompPostingsType) Get() *CompPostingsType {
	return v.value
}

func (v *NullableCompPostingsType) Set(val *CompPostingsType) {
	v.value = val
	v.isSet = true
}

func (v NullableCompPostingsType) IsSet() bool {
	return v.isSet
}

func (v *NullableCompPostingsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompPostingsType(val *CompPostingsType) *NullableCompPostingsType {
	return &NullableCompPostingsType{value: val, isSet: true}
}

func (v NullableCompPostingsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompPostingsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


