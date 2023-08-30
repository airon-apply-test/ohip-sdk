/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
)

// checks if the AuthorizationInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationInfoType{}

// AuthorizationInfoType struct for AuthorizationInfoType
type AuthorizationInfoType struct {
	ApprovalAmount *CurrencyAmountType `json:"approvalAmount,omitempty"`
	// The approval code authenticates the authorization.
	ApprovalCode *string `json:"approvalCode,omitempty"`
	// Unique Authorization Sequence for the authorization and settlement.
	OriginalAuthSequence *int32 `json:"originalAuthSequence,omitempty"`
	// Vendor transaction id for the authorization.
	VendorTranId *string `json:"vendorTranId,omitempty"`
}

// NewAuthorizationInfoType instantiates a new AuthorizationInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationInfoType() *AuthorizationInfoType {
	this := AuthorizationInfoType{}
	return &this
}

// NewAuthorizationInfoTypeWithDefaults instantiates a new AuthorizationInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationInfoTypeWithDefaults() *AuthorizationInfoType {
	this := AuthorizationInfoType{}
	return &this
}

// GetApprovalAmount returns the ApprovalAmount field value if set, zero value otherwise.
func (o *AuthorizationInfoType) GetApprovalAmount() CurrencyAmountType {
	if o == nil || IsNil(o.ApprovalAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.ApprovalAmount
}

// GetApprovalAmountOk returns a tuple with the ApprovalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfoType) GetApprovalAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.ApprovalAmount) {
		return nil, false
	}
	return o.ApprovalAmount, true
}

// HasApprovalAmount returns a boolean if a field has been set.
func (o *AuthorizationInfoType) HasApprovalAmount() bool {
	if o != nil && !IsNil(o.ApprovalAmount) {
		return true
	}

	return false
}

// SetApprovalAmount gets a reference to the given CurrencyAmountType and assigns it to the ApprovalAmount field.
func (o *AuthorizationInfoType) SetApprovalAmount(v CurrencyAmountType) {
	o.ApprovalAmount = &v
}

// GetApprovalCode returns the ApprovalCode field value if set, zero value otherwise.
func (o *AuthorizationInfoType) GetApprovalCode() string {
	if o == nil || IsNil(o.ApprovalCode) {
		var ret string
		return ret
	}
	return *o.ApprovalCode
}

// GetApprovalCodeOk returns a tuple with the ApprovalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfoType) GetApprovalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalCode) {
		return nil, false
	}
	return o.ApprovalCode, true
}

// HasApprovalCode returns a boolean if a field has been set.
func (o *AuthorizationInfoType) HasApprovalCode() bool {
	if o != nil && !IsNil(o.ApprovalCode) {
		return true
	}

	return false
}

// SetApprovalCode gets a reference to the given string and assigns it to the ApprovalCode field.
func (o *AuthorizationInfoType) SetApprovalCode(v string) {
	o.ApprovalCode = &v
}

// GetOriginalAuthSequence returns the OriginalAuthSequence field value if set, zero value otherwise.
func (o *AuthorizationInfoType) GetOriginalAuthSequence() int32 {
	if o == nil || IsNil(o.OriginalAuthSequence) {
		var ret int32
		return ret
	}
	return *o.OriginalAuthSequence
}

// GetOriginalAuthSequenceOk returns a tuple with the OriginalAuthSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfoType) GetOriginalAuthSequenceOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalAuthSequence) {
		return nil, false
	}
	return o.OriginalAuthSequence, true
}

// HasOriginalAuthSequence returns a boolean if a field has been set.
func (o *AuthorizationInfoType) HasOriginalAuthSequence() bool {
	if o != nil && !IsNil(o.OriginalAuthSequence) {
		return true
	}

	return false
}

// SetOriginalAuthSequence gets a reference to the given int32 and assigns it to the OriginalAuthSequence field.
func (o *AuthorizationInfoType) SetOriginalAuthSequence(v int32) {
	o.OriginalAuthSequence = &v
}

// GetVendorTranId returns the VendorTranId field value if set, zero value otherwise.
func (o *AuthorizationInfoType) GetVendorTranId() string {
	if o == nil || IsNil(o.VendorTranId) {
		var ret string
		return ret
	}
	return *o.VendorTranId
}

// GetVendorTranIdOk returns a tuple with the VendorTranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfoType) GetVendorTranIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorTranId) {
		return nil, false
	}
	return o.VendorTranId, true
}

// HasVendorTranId returns a boolean if a field has been set.
func (o *AuthorizationInfoType) HasVendorTranId() bool {
	if o != nil && !IsNil(o.VendorTranId) {
		return true
	}

	return false
}

// SetVendorTranId gets a reference to the given string and assigns it to the VendorTranId field.
func (o *AuthorizationInfoType) SetVendorTranId(v string) {
	o.VendorTranId = &v
}

func (o AuthorizationInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApprovalAmount) {
		toSerialize["approvalAmount"] = o.ApprovalAmount
	}
	if !IsNil(o.ApprovalCode) {
		toSerialize["approvalCode"] = o.ApprovalCode
	}
	if !IsNil(o.OriginalAuthSequence) {
		toSerialize["originalAuthSequence"] = o.OriginalAuthSequence
	}
	if !IsNil(o.VendorTranId) {
		toSerialize["vendorTranId"] = o.VendorTranId
	}
	return toSerialize, nil
}

type NullableAuthorizationInfoType struct {
	value *AuthorizationInfoType
	isSet bool
}

func (v NullableAuthorizationInfoType) Get() *AuthorizationInfoType {
	return v.value
}

func (v *NullableAuthorizationInfoType) Set(val *AuthorizationInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationInfoType(val *AuthorizationInfoType) *NullableAuthorizationInfoType {
	return &NullableAuthorizationInfoType{value: val, isSet: true}
}

func (v NullableAuthorizationInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


