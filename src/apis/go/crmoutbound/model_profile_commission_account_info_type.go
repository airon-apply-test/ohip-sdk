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

// checks if the ProfileCommissionAccountInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileCommissionAccountInfoType{}

// ProfileCommissionAccountInfoType Profile commission info which contains bank account and commission code details
type ProfileCommissionAccountInfoType struct {
	ProfileId *ProfileId `json:"profileId,omitempty"`
	BankAccount *BankAccountType `json:"bankAccount,omitempty"`
	CommissionCode *CodeDescriptionType `json:"commissionCode,omitempty"`
}

// NewProfileCommissionAccountInfoType instantiates a new ProfileCommissionAccountInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileCommissionAccountInfoType() *ProfileCommissionAccountInfoType {
	this := ProfileCommissionAccountInfoType{}
	return &this
}

// NewProfileCommissionAccountInfoTypeWithDefaults instantiates a new ProfileCommissionAccountInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileCommissionAccountInfoTypeWithDefaults() *ProfileCommissionAccountInfoType {
	this := ProfileCommissionAccountInfoType{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ProfileCommissionAccountInfoType) GetProfileId() ProfileId {
	if o == nil || IsNil(o.ProfileId) {
		var ret ProfileId
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCommissionAccountInfoType) GetProfileIdOk() (*ProfileId, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ProfileCommissionAccountInfoType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given ProfileId and assigns it to the ProfileId field.
func (o *ProfileCommissionAccountInfoType) SetProfileId(v ProfileId) {
	o.ProfileId = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *ProfileCommissionAccountInfoType) GetBankAccount() BankAccountType {
	if o == nil || IsNil(o.BankAccount) {
		var ret BankAccountType
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCommissionAccountInfoType) GetBankAccountOk() (*BankAccountType, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *ProfileCommissionAccountInfoType) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given BankAccountType and assigns it to the BankAccount field.
func (o *ProfileCommissionAccountInfoType) SetBankAccount(v BankAccountType) {
	o.BankAccount = &v
}

// GetCommissionCode returns the CommissionCode field value if set, zero value otherwise.
func (o *ProfileCommissionAccountInfoType) GetCommissionCode() CodeDescriptionType {
	if o == nil || IsNil(o.CommissionCode) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.CommissionCode
}

// GetCommissionCodeOk returns a tuple with the CommissionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCommissionAccountInfoType) GetCommissionCodeOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.CommissionCode) {
		return nil, false
	}
	return o.CommissionCode, true
}

// HasCommissionCode returns a boolean if a field has been set.
func (o *ProfileCommissionAccountInfoType) HasCommissionCode() bool {
	if o != nil && !IsNil(o.CommissionCode) {
		return true
	}

	return false
}

// SetCommissionCode gets a reference to the given CodeDescriptionType and assigns it to the CommissionCode field.
func (o *ProfileCommissionAccountInfoType) SetCommissionCode(v CodeDescriptionType) {
	o.CommissionCode = &v
}

func (o ProfileCommissionAccountInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileCommissionAccountInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !IsNil(o.CommissionCode) {
		toSerialize["commissionCode"] = o.CommissionCode
	}
	return toSerialize, nil
}

type NullableProfileCommissionAccountInfoType struct {
	value *ProfileCommissionAccountInfoType
	isSet bool
}

func (v NullableProfileCommissionAccountInfoType) Get() *ProfileCommissionAccountInfoType {
	return v.value
}

func (v *NullableProfileCommissionAccountInfoType) Set(val *ProfileCommissionAccountInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileCommissionAccountInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileCommissionAccountInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileCommissionAccountInfoType(val *ProfileCommissionAccountInfoType) *NullableProfileCommissionAccountInfoType {
	return &NullableProfileCommissionAccountInfoType{value: val, isSet: true}
}

func (v NullableProfileCommissionAccountInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileCommissionAccountInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


