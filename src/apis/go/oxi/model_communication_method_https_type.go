/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the CommunicationMethodHTTPSType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationMethodHTTPSType{}

// CommunicationMethodHTTPSType Type represents the Communication Type Base Details.
type CommunicationMethodHTTPSType struct {
	// Attribute represents Username for Types HTTP
	UserName *string `json:"userName,omitempty"`
	// Attribute represents Password for Types HTTP
	Password *string `json:"password,omitempty"`
	// Attribute represents Idle Time(in Minutes) of HTTP,None
	SleepTime *float32 `json:"sleepTime,omitempty"`
	// Attribute is used to indentify the TimeOut(in Seconds) of HTTP,None
	TimeOut *float32 `json:"timeOut,omitempty"`
	// Attribute represents Address
	Address *string `json:"address,omitempty"`
	// Attribute that represents Proxy URL
	ProxyAddress *string `json:"proxyAddress,omitempty"`
	// Attribute that indicates whether data to compressed or not
	CompressData *bool `json:"compressData,omitempty"`
	// Attribute that indicates whether to use client certificate or not
	UseClientCertificate *bool `json:"useClientCertificate,omitempty"`
	System *SystemType `json:"system,omitempty"`
	// Locale(Holidex specific) Values like US,EU
	Locale *string `json:"locale,omitempty"`
	// Attribute that indicates whether to use Proxy URL or User configured URL.
	UseVaultProxy *bool `json:"useVaultProxy,omitempty"`
	// Attribute that indicates whether to allow compress data or not.
	AllowCompressData *bool `json:"allowCompressData,omitempty"`
}

// NewCommunicationMethodHTTPSType instantiates a new CommunicationMethodHTTPSType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationMethodHTTPSType() *CommunicationMethodHTTPSType {
	this := CommunicationMethodHTTPSType{}
	return &this
}

// NewCommunicationMethodHTTPSTypeWithDefaults instantiates a new CommunicationMethodHTTPSType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationMethodHTTPSTypeWithDefaults() *CommunicationMethodHTTPSType {
	this := CommunicationMethodHTTPSType{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *CommunicationMethodHTTPSType) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CommunicationMethodHTTPSType) SetPassword(v string) {
	o.Password = &v
}

// GetSleepTime returns the SleepTime field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetSleepTime() float32 {
	if o == nil || IsNil(o.SleepTime) {
		var ret float32
		return ret
	}
	return *o.SleepTime
}

// GetSleepTimeOk returns a tuple with the SleepTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetSleepTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.SleepTime) {
		return nil, false
	}
	return o.SleepTime, true
}

// HasSleepTime returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasSleepTime() bool {
	if o != nil && !IsNil(o.SleepTime) {
		return true
	}

	return false
}

// SetSleepTime gets a reference to the given float32 and assigns it to the SleepTime field.
func (o *CommunicationMethodHTTPSType) SetSleepTime(v float32) {
	o.SleepTime = &v
}

// GetTimeOut returns the TimeOut field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetTimeOut() float32 {
	if o == nil || IsNil(o.TimeOut) {
		var ret float32
		return ret
	}
	return *o.TimeOut
}

// GetTimeOutOk returns a tuple with the TimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetTimeOutOk() (*float32, bool) {
	if o == nil || IsNil(o.TimeOut) {
		return nil, false
	}
	return o.TimeOut, true
}

// HasTimeOut returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasTimeOut() bool {
	if o != nil && !IsNil(o.TimeOut) {
		return true
	}

	return false
}

// SetTimeOut gets a reference to the given float32 and assigns it to the TimeOut field.
func (o *CommunicationMethodHTTPSType) SetTimeOut(v float32) {
	o.TimeOut = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CommunicationMethodHTTPSType) SetAddress(v string) {
	o.Address = &v
}

// GetProxyAddress returns the ProxyAddress field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetProxyAddress() string {
	if o == nil || IsNil(o.ProxyAddress) {
		var ret string
		return ret
	}
	return *o.ProxyAddress
}

// GetProxyAddressOk returns a tuple with the ProxyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetProxyAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyAddress) {
		return nil, false
	}
	return o.ProxyAddress, true
}

// HasProxyAddress returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasProxyAddress() bool {
	if o != nil && !IsNil(o.ProxyAddress) {
		return true
	}

	return false
}

// SetProxyAddress gets a reference to the given string and assigns it to the ProxyAddress field.
func (o *CommunicationMethodHTTPSType) SetProxyAddress(v string) {
	o.ProxyAddress = &v
}

// GetCompressData returns the CompressData field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetCompressData() bool {
	if o == nil || IsNil(o.CompressData) {
		var ret bool
		return ret
	}
	return *o.CompressData
}

// GetCompressDataOk returns a tuple with the CompressData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetCompressDataOk() (*bool, bool) {
	if o == nil || IsNil(o.CompressData) {
		return nil, false
	}
	return o.CompressData, true
}

// HasCompressData returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasCompressData() bool {
	if o != nil && !IsNil(o.CompressData) {
		return true
	}

	return false
}

// SetCompressData gets a reference to the given bool and assigns it to the CompressData field.
func (o *CommunicationMethodHTTPSType) SetCompressData(v bool) {
	o.CompressData = &v
}

// GetUseClientCertificate returns the UseClientCertificate field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetUseClientCertificate() bool {
	if o == nil || IsNil(o.UseClientCertificate) {
		var ret bool
		return ret
	}
	return *o.UseClientCertificate
}

// GetUseClientCertificateOk returns a tuple with the UseClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetUseClientCertificateOk() (*bool, bool) {
	if o == nil || IsNil(o.UseClientCertificate) {
		return nil, false
	}
	return o.UseClientCertificate, true
}

// HasUseClientCertificate returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasUseClientCertificate() bool {
	if o != nil && !IsNil(o.UseClientCertificate) {
		return true
	}

	return false
}

// SetUseClientCertificate gets a reference to the given bool and assigns it to the UseClientCertificate field.
func (o *CommunicationMethodHTTPSType) SetUseClientCertificate(v bool) {
	o.UseClientCertificate = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetSystem() SystemType {
	if o == nil || IsNil(o.System) {
		var ret SystemType
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetSystemOk() (*SystemType, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given SystemType and assigns it to the System field.
func (o *CommunicationMethodHTTPSType) SetSystem(v SystemType) {
	o.System = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *CommunicationMethodHTTPSType) SetLocale(v string) {
	o.Locale = &v
}

// GetUseVaultProxy returns the UseVaultProxy field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetUseVaultProxy() bool {
	if o == nil || IsNil(o.UseVaultProxy) {
		var ret bool
		return ret
	}
	return *o.UseVaultProxy
}

// GetUseVaultProxyOk returns a tuple with the UseVaultProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetUseVaultProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseVaultProxy) {
		return nil, false
	}
	return o.UseVaultProxy, true
}

// HasUseVaultProxy returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasUseVaultProxy() bool {
	if o != nil && !IsNil(o.UseVaultProxy) {
		return true
	}

	return false
}

// SetUseVaultProxy gets a reference to the given bool and assigns it to the UseVaultProxy field.
func (o *CommunicationMethodHTTPSType) SetUseVaultProxy(v bool) {
	o.UseVaultProxy = &v
}

// GetAllowCompressData returns the AllowCompressData field value if set, zero value otherwise.
func (o *CommunicationMethodHTTPSType) GetAllowCompressData() bool {
	if o == nil || IsNil(o.AllowCompressData) {
		var ret bool
		return ret
	}
	return *o.AllowCompressData
}

// GetAllowCompressDataOk returns a tuple with the AllowCompressData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMethodHTTPSType) GetAllowCompressDataOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCompressData) {
		return nil, false
	}
	return o.AllowCompressData, true
}

// HasAllowCompressData returns a boolean if a field has been set.
func (o *CommunicationMethodHTTPSType) HasAllowCompressData() bool {
	if o != nil && !IsNil(o.AllowCompressData) {
		return true
	}

	return false
}

// SetAllowCompressData gets a reference to the given bool and assigns it to the AllowCompressData field.
func (o *CommunicationMethodHTTPSType) SetAllowCompressData(v bool) {
	o.AllowCompressData = &v
}

func (o CommunicationMethodHTTPSType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationMethodHTTPSType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.SleepTime) {
		toSerialize["sleepTime"] = o.SleepTime
	}
	if !IsNil(o.TimeOut) {
		toSerialize["timeOut"] = o.TimeOut
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ProxyAddress) {
		toSerialize["proxyAddress"] = o.ProxyAddress
	}
	if !IsNil(o.CompressData) {
		toSerialize["compressData"] = o.CompressData
	}
	if !IsNil(o.UseClientCertificate) {
		toSerialize["useClientCertificate"] = o.UseClientCertificate
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.UseVaultProxy) {
		toSerialize["useVaultProxy"] = o.UseVaultProxy
	}
	if !IsNil(o.AllowCompressData) {
		toSerialize["allowCompressData"] = o.AllowCompressData
	}
	return toSerialize, nil
}

type NullableCommunicationMethodHTTPSType struct {
	value *CommunicationMethodHTTPSType
	isSet bool
}

func (v NullableCommunicationMethodHTTPSType) Get() *CommunicationMethodHTTPSType {
	return v.value
}

func (v *NullableCommunicationMethodHTTPSType) Set(val *CommunicationMethodHTTPSType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationMethodHTTPSType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationMethodHTTPSType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationMethodHTTPSType(val *CommunicationMethodHTTPSType) *NullableCommunicationMethodHTTPSType {
	return &NullableCommunicationMethodHTTPSType{value: val, isSet: true}
}

func (v NullableCommunicationMethodHTTPSType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationMethodHTTPSType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


