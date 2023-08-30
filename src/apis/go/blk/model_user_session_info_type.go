/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the UserSessionInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSessionInfoType{}

// UserSessionInfoType struct for UserSessionInfoType
type UserSessionInfoType struct {
	BusinessDate *string `json:"businessDate,omitempty"`
	SystemDate *string `json:"systemDate,omitempty"`
	Terminal *string `json:"terminal,omitempty"`
	RunningApp *string `json:"runningApp,omitempty"`
	ShareProfiles *bool `json:"shareProfiles,omitempty"`
	Hotel *CodeDescriptionType `json:"hotel,omitempty"`
	Cro *CodeDescriptionType `json:"cro,omitempty"`
	Chain *string `json:"chain,omitempty"`
	CROCountryCode *string `json:"cROCountryCode,omitempty"`
	SessionDefaults *UserSessionDefaultsType `json:"sessionDefaults,omitempty"`
	// Collection of generic Name-Value-Pair parameters.
	Parameters []ParameterType `json:"parameters,omitempty"`
}

// NewUserSessionInfoType instantiates a new UserSessionInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSessionInfoType() *UserSessionInfoType {
	this := UserSessionInfoType{}
	return &this
}

// NewUserSessionInfoTypeWithDefaults instantiates a new UserSessionInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSessionInfoTypeWithDefaults() *UserSessionInfoType {
	this := UserSessionInfoType{}
	return &this
}

// GetBusinessDate returns the BusinessDate field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetBusinessDate() string {
	if o == nil || IsNil(o.BusinessDate) {
		var ret string
		return ret
	}
	return *o.BusinessDate
}

// GetBusinessDateOk returns a tuple with the BusinessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetBusinessDateOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessDate) {
		return nil, false
	}
	return o.BusinessDate, true
}

// HasBusinessDate returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasBusinessDate() bool {
	if o != nil && !IsNil(o.BusinessDate) {
		return true
	}

	return false
}

// SetBusinessDate gets a reference to the given string and assigns it to the BusinessDate field.
func (o *UserSessionInfoType) SetBusinessDate(v string) {
	o.BusinessDate = &v
}

// GetSystemDate returns the SystemDate field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetSystemDate() string {
	if o == nil || IsNil(o.SystemDate) {
		var ret string
		return ret
	}
	return *o.SystemDate
}

// GetSystemDateOk returns a tuple with the SystemDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetSystemDateOk() (*string, bool) {
	if o == nil || IsNil(o.SystemDate) {
		return nil, false
	}
	return o.SystemDate, true
}

// HasSystemDate returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasSystemDate() bool {
	if o != nil && !IsNil(o.SystemDate) {
		return true
	}

	return false
}

// SetSystemDate gets a reference to the given string and assigns it to the SystemDate field.
func (o *UserSessionInfoType) SetSystemDate(v string) {
	o.SystemDate = &v
}

// GetTerminal returns the Terminal field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetTerminal() string {
	if o == nil || IsNil(o.Terminal) {
		var ret string
		return ret
	}
	return *o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetTerminalOk() (*string, bool) {
	if o == nil || IsNil(o.Terminal) {
		return nil, false
	}
	return o.Terminal, true
}

// HasTerminal returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasTerminal() bool {
	if o != nil && !IsNil(o.Terminal) {
		return true
	}

	return false
}

// SetTerminal gets a reference to the given string and assigns it to the Terminal field.
func (o *UserSessionInfoType) SetTerminal(v string) {
	o.Terminal = &v
}

// GetRunningApp returns the RunningApp field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetRunningApp() string {
	if o == nil || IsNil(o.RunningApp) {
		var ret string
		return ret
	}
	return *o.RunningApp
}

// GetRunningAppOk returns a tuple with the RunningApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetRunningAppOk() (*string, bool) {
	if o == nil || IsNil(o.RunningApp) {
		return nil, false
	}
	return o.RunningApp, true
}

// HasRunningApp returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasRunningApp() bool {
	if o != nil && !IsNil(o.RunningApp) {
		return true
	}

	return false
}

// SetRunningApp gets a reference to the given string and assigns it to the RunningApp field.
func (o *UserSessionInfoType) SetRunningApp(v string) {
	o.RunningApp = &v
}

// GetShareProfiles returns the ShareProfiles field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetShareProfiles() bool {
	if o == nil || IsNil(o.ShareProfiles) {
		var ret bool
		return ret
	}
	return *o.ShareProfiles
}

// GetShareProfilesOk returns a tuple with the ShareProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetShareProfilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ShareProfiles) {
		return nil, false
	}
	return o.ShareProfiles, true
}

// HasShareProfiles returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasShareProfiles() bool {
	if o != nil && !IsNil(o.ShareProfiles) {
		return true
	}

	return false
}

// SetShareProfiles gets a reference to the given bool and assigns it to the ShareProfiles field.
func (o *UserSessionInfoType) SetShareProfiles(v bool) {
	o.ShareProfiles = &v
}

// GetHotel returns the Hotel field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetHotel() CodeDescriptionType {
	if o == nil || IsNil(o.Hotel) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Hotel
}

// GetHotelOk returns a tuple with the Hotel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetHotelOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Hotel) {
		return nil, false
	}
	return o.Hotel, true
}

// HasHotel returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasHotel() bool {
	if o != nil && !IsNil(o.Hotel) {
		return true
	}

	return false
}

// SetHotel gets a reference to the given CodeDescriptionType and assigns it to the Hotel field.
func (o *UserSessionInfoType) SetHotel(v CodeDescriptionType) {
	o.Hotel = &v
}

// GetCro returns the Cro field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetCro() CodeDescriptionType {
	if o == nil || IsNil(o.Cro) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Cro
}

// GetCroOk returns a tuple with the Cro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetCroOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Cro) {
		return nil, false
	}
	return o.Cro, true
}

// HasCro returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasCro() bool {
	if o != nil && !IsNil(o.Cro) {
		return true
	}

	return false
}

// SetCro gets a reference to the given CodeDescriptionType and assigns it to the Cro field.
func (o *UserSessionInfoType) SetCro(v CodeDescriptionType) {
	o.Cro = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetChain() string {
	if o == nil || IsNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetChainOk() (*string, bool) {
	if o == nil || IsNil(o.Chain) {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasChain() bool {
	if o != nil && !IsNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *UserSessionInfoType) SetChain(v string) {
	o.Chain = &v
}

// GetCROCountryCode returns the CROCountryCode field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetCROCountryCode() string {
	if o == nil || IsNil(o.CROCountryCode) {
		var ret string
		return ret
	}
	return *o.CROCountryCode
}

// GetCROCountryCodeOk returns a tuple with the CROCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetCROCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CROCountryCode) {
		return nil, false
	}
	return o.CROCountryCode, true
}

// HasCROCountryCode returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasCROCountryCode() bool {
	if o != nil && !IsNil(o.CROCountryCode) {
		return true
	}

	return false
}

// SetCROCountryCode gets a reference to the given string and assigns it to the CROCountryCode field.
func (o *UserSessionInfoType) SetCROCountryCode(v string) {
	o.CROCountryCode = &v
}

// GetSessionDefaults returns the SessionDefaults field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetSessionDefaults() UserSessionDefaultsType {
	if o == nil || IsNil(o.SessionDefaults) {
		var ret UserSessionDefaultsType
		return ret
	}
	return *o.SessionDefaults
}

// GetSessionDefaultsOk returns a tuple with the SessionDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetSessionDefaultsOk() (*UserSessionDefaultsType, bool) {
	if o == nil || IsNil(o.SessionDefaults) {
		return nil, false
	}
	return o.SessionDefaults, true
}

// HasSessionDefaults returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasSessionDefaults() bool {
	if o != nil && !IsNil(o.SessionDefaults) {
		return true
	}

	return false
}

// SetSessionDefaults gets a reference to the given UserSessionDefaultsType and assigns it to the SessionDefaults field.
func (o *UserSessionInfoType) SetSessionDefaults(v UserSessionDefaultsType) {
	o.SessionDefaults = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *UserSessionInfoType) GetParameters() []ParameterType {
	if o == nil || IsNil(o.Parameters) {
		var ret []ParameterType
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionInfoType) GetParametersOk() ([]ParameterType, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *UserSessionInfoType) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ParameterType and assigns it to the Parameters field.
func (o *UserSessionInfoType) SetParameters(v []ParameterType) {
	o.Parameters = v
}

func (o UserSessionInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSessionInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessDate) {
		toSerialize["businessDate"] = o.BusinessDate
	}
	if !IsNil(o.SystemDate) {
		toSerialize["systemDate"] = o.SystemDate
	}
	if !IsNil(o.Terminal) {
		toSerialize["terminal"] = o.Terminal
	}
	if !IsNil(o.RunningApp) {
		toSerialize["runningApp"] = o.RunningApp
	}
	if !IsNil(o.ShareProfiles) {
		toSerialize["shareProfiles"] = o.ShareProfiles
	}
	if !IsNil(o.Hotel) {
		toSerialize["hotel"] = o.Hotel
	}
	if !IsNil(o.Cro) {
		toSerialize["cro"] = o.Cro
	}
	if !IsNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	if !IsNil(o.CROCountryCode) {
		toSerialize["cROCountryCode"] = o.CROCountryCode
	}
	if !IsNil(o.SessionDefaults) {
		toSerialize["sessionDefaults"] = o.SessionDefaults
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableUserSessionInfoType struct {
	value *UserSessionInfoType
	isSet bool
}

func (v NullableUserSessionInfoType) Get() *UserSessionInfoType {
	return v.value
}

func (v *NullableUserSessionInfoType) Set(val *UserSessionInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSessionInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSessionInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSessionInfoType(val *UserSessionInfoType) *NullableUserSessionInfoType {
	return &NullableUserSessionInfoType{value: val, isSet: true}
}

func (v NullableUserSessionInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSessionInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


