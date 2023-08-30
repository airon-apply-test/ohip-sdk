/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package act

import (
	"encoding/json"
)

// checks if the ActivityDetailsTypeLinkedAccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityDetailsTypeLinkedAccounts{}

// ActivityDetailsTypeLinkedAccounts Provides information about the accounts linked to an activity. Please note that during a change operation this performs a full overlay if the attribute FullOverlay is set to true. In a full overlay, all the accounts that should be associated to the activity should be provided during a change operation. Any accounts not provided will be detached from this activity. By default the full overlay is considered false if this property is left blank. If values are provided for accounts, only the full overlay functionality is provided at this time.
type ActivityDetailsTypeLinkedAccounts struct {
	// The list of accounts associated with an activity.
	ActivityAccount []ActivityLinkedProfilesType `json:"activityAccount,omitempty"`
	// Indicates whether to perform a full overlay for the accounts.
	FullOverlay *bool `json:"fullOverlay,omitempty"`
}

// NewActivityDetailsTypeLinkedAccounts instantiates a new ActivityDetailsTypeLinkedAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityDetailsTypeLinkedAccounts() *ActivityDetailsTypeLinkedAccounts {
	this := ActivityDetailsTypeLinkedAccounts{}
	return &this
}

// NewActivityDetailsTypeLinkedAccountsWithDefaults instantiates a new ActivityDetailsTypeLinkedAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityDetailsTypeLinkedAccountsWithDefaults() *ActivityDetailsTypeLinkedAccounts {
	this := ActivityDetailsTypeLinkedAccounts{}
	return &this
}

// GetActivityAccount returns the ActivityAccount field value if set, zero value otherwise.
func (o *ActivityDetailsTypeLinkedAccounts) GetActivityAccount() []ActivityLinkedProfilesType {
	if o == nil || IsNil(o.ActivityAccount) {
		var ret []ActivityLinkedProfilesType
		return ret
	}
	return o.ActivityAccount
}

// GetActivityAccountOk returns a tuple with the ActivityAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsTypeLinkedAccounts) GetActivityAccountOk() ([]ActivityLinkedProfilesType, bool) {
	if o == nil || IsNil(o.ActivityAccount) {
		return nil, false
	}
	return o.ActivityAccount, true
}

// HasActivityAccount returns a boolean if a field has been set.
func (o *ActivityDetailsTypeLinkedAccounts) HasActivityAccount() bool {
	if o != nil && !IsNil(o.ActivityAccount) {
		return true
	}

	return false
}

// SetActivityAccount gets a reference to the given []ActivityLinkedProfilesType and assigns it to the ActivityAccount field.
func (o *ActivityDetailsTypeLinkedAccounts) SetActivityAccount(v []ActivityLinkedProfilesType) {
	o.ActivityAccount = v
}

// GetFullOverlay returns the FullOverlay field value if set, zero value otherwise.
func (o *ActivityDetailsTypeLinkedAccounts) GetFullOverlay() bool {
	if o == nil || IsNil(o.FullOverlay) {
		var ret bool
		return ret
	}
	return *o.FullOverlay
}

// GetFullOverlayOk returns a tuple with the FullOverlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsTypeLinkedAccounts) GetFullOverlayOk() (*bool, bool) {
	if o == nil || IsNil(o.FullOverlay) {
		return nil, false
	}
	return o.FullOverlay, true
}

// HasFullOverlay returns a boolean if a field has been set.
func (o *ActivityDetailsTypeLinkedAccounts) HasFullOverlay() bool {
	if o != nil && !IsNil(o.FullOverlay) {
		return true
	}

	return false
}

// SetFullOverlay gets a reference to the given bool and assigns it to the FullOverlay field.
func (o *ActivityDetailsTypeLinkedAccounts) SetFullOverlay(v bool) {
	o.FullOverlay = &v
}

func (o ActivityDetailsTypeLinkedAccounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityDetailsTypeLinkedAccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityAccount) {
		toSerialize["activityAccount"] = o.ActivityAccount
	}
	if !IsNil(o.FullOverlay) {
		toSerialize["fullOverlay"] = o.FullOverlay
	}
	return toSerialize, nil
}

type NullableActivityDetailsTypeLinkedAccounts struct {
	value *ActivityDetailsTypeLinkedAccounts
	isSet bool
}

func (v NullableActivityDetailsTypeLinkedAccounts) Get() *ActivityDetailsTypeLinkedAccounts {
	return v.value
}

func (v *NullableActivityDetailsTypeLinkedAccounts) Set(val *ActivityDetailsTypeLinkedAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityDetailsTypeLinkedAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityDetailsTypeLinkedAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityDetailsTypeLinkedAccounts(val *ActivityDetailsTypeLinkedAccounts) *NullableActivityDetailsTypeLinkedAccounts {
	return &NullableActivityDetailsTypeLinkedAccounts{value: val, isSet: true}
}

func (v NullableActivityDetailsTypeLinkedAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityDetailsTypeLinkedAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


