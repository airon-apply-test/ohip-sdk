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

// checks if the AddressSearchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressSearchType{}

// AddressSearchType Address Details such as city, state, country, postal code etc.
type AddressSearchType struct {
	// City (e.g., Dublin), town, or postal station (i.e., a postal service territory, often used in a military address).
	CityName *string `json:"cityName,omitempty"`
	// Post Office Code number.
	PostalCode *string `json:"postalCode,omitempty"`
	// State or Province name (e.g., Texas).
	State *string `json:"state,omitempty"`
	Country *CountryNameType `json:"country,omitempty"`
	// First Line of Street Address. For profile search it matches the first Address line.
	StreetAddress *string `json:"streetAddress,omitempty"`
	// When true indicates that only profiles with city will be fetched.
	ExcludeNoCity *bool `json:"excludeNoCity,omitempty"`
}

// NewAddressSearchType instantiates a new AddressSearchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressSearchType() *AddressSearchType {
	this := AddressSearchType{}
	return &this
}

// NewAddressSearchTypeWithDefaults instantiates a new AddressSearchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressSearchTypeWithDefaults() *AddressSearchType {
	this := AddressSearchType{}
	return &this
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *AddressSearchType) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchType) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *AddressSearchType) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *AddressSearchType) SetCityName(v string) {
	o.CityName = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressSearchType) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchType) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressSearchType) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressSearchType) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressSearchType) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchType) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressSearchType) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressSearchType) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressSearchType) GetCountry() CountryNameType {
	if o == nil || IsNil(o.Country) {
		var ret CountryNameType
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchType) GetCountryOk() (*CountryNameType, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressSearchType) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given CountryNameType and assigns it to the Country field.
func (o *AddressSearchType) SetCountry(v CountryNameType) {
	o.Country = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *AddressSearchType) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchType) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *AddressSearchType) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *AddressSearchType) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetExcludeNoCity returns the ExcludeNoCity field value if set, zero value otherwise.
func (o *AddressSearchType) GetExcludeNoCity() bool {
	if o == nil || IsNil(o.ExcludeNoCity) {
		var ret bool
		return ret
	}
	return *o.ExcludeNoCity
}

// GetExcludeNoCityOk returns a tuple with the ExcludeNoCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressSearchType) GetExcludeNoCityOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeNoCity) {
		return nil, false
	}
	return o.ExcludeNoCity, true
}

// HasExcludeNoCity returns a boolean if a field has been set.
func (o *AddressSearchType) HasExcludeNoCity() bool {
	if o != nil && !IsNil(o.ExcludeNoCity) {
		return true
	}

	return false
}

// SetExcludeNoCity gets a reference to the given bool and assigns it to the ExcludeNoCity field.
func (o *AddressSearchType) SetExcludeNoCity(v bool) {
	o.ExcludeNoCity = &v
}

func (o AddressSearchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressSearchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CityName) {
		toSerialize["cityName"] = o.CityName
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if !IsNil(o.ExcludeNoCity) {
		toSerialize["excludeNoCity"] = o.ExcludeNoCity
	}
	return toSerialize, nil
}

type NullableAddressSearchType struct {
	value *AddressSearchType
	isSet bool
}

func (v NullableAddressSearchType) Get() *AddressSearchType {
	return v.value
}

func (v *NullableAddressSearchType) Set(val *AddressSearchType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressSearchType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressSearchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressSearchType(val *AddressSearchType) *NullableAddressSearchType {
	return &NullableAddressSearchType{value: val, isSet: true}
}

func (v NullableAddressSearchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressSearchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


