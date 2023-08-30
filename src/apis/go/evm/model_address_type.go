/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
)

// checks if the AddressType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressType{}

// AddressType Provides address information.
type AddressType struct {
	// Indicator to define if the Address is validated by the Address Validation System.
	IsValidated *bool `json:"isValidated,omitempty"`
	// When the address is unformatted (FormattedInd=\"false\") these lines will contain free form address details. When the address is formatted and street number and street name must be sent independently, the street number will be sent using StreetNmbr, and the street name will be sent in the first AddressLine occurrence.
	AddressLine []string `json:"addressLine,omitempty"`
	// City (e.g., Dublin), town, or postal station (i.e., a postal service territory, often used in a military address).
	CityName *string `json:"cityName,omitempty"`
	// Post Office Code number.
	PostalCode *string `json:"postalCode,omitempty"`
	// Post Office City Extension Code number. City Extension mainly used for UK addresses.
	CityExtension *string `json:"cityExtension,omitempty"`
	// County or District Name (e.g., Fairfax). This is read only.
	County *string `json:"county,omitempty"`
	// State or Province name (e.g., Texas).
	State *string `json:"state,omitempty"`
	Country *CountryNameType `json:"country,omitempty"`
	// Defines the type of address (e.g. home, business, other).
	Type *string `json:"type,omitempty"`
	// Describes the type code
	TypeDescription *string `json:"typeDescription,omitempty"`
}

// NewAddressType instantiates a new AddressType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressType() *AddressType {
	this := AddressType{}
	return &this
}

// NewAddressTypeWithDefaults instantiates a new AddressType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTypeWithDefaults() *AddressType {
	this := AddressType{}
	return &this
}

// GetIsValidated returns the IsValidated field value if set, zero value otherwise.
func (o *AddressType) GetIsValidated() bool {
	if o == nil || IsNil(o.IsValidated) {
		var ret bool
		return ret
	}
	return *o.IsValidated
}

// GetIsValidatedOk returns a tuple with the IsValidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetIsValidatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValidated) {
		return nil, false
	}
	return o.IsValidated, true
}

// HasIsValidated returns a boolean if a field has been set.
func (o *AddressType) HasIsValidated() bool {
	if o != nil && !IsNil(o.IsValidated) {
		return true
	}

	return false
}

// SetIsValidated gets a reference to the given bool and assigns it to the IsValidated field.
func (o *AddressType) SetIsValidated(v bool) {
	o.IsValidated = &v
}

// GetAddressLine returns the AddressLine field value if set, zero value otherwise.
func (o *AddressType) GetAddressLine() []string {
	if o == nil || IsNil(o.AddressLine) {
		var ret []string
		return ret
	}
	return o.AddressLine
}

// GetAddressLineOk returns a tuple with the AddressLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetAddressLineOk() ([]string, bool) {
	if o == nil || IsNil(o.AddressLine) {
		return nil, false
	}
	return o.AddressLine, true
}

// HasAddressLine returns a boolean if a field has been set.
func (o *AddressType) HasAddressLine() bool {
	if o != nil && !IsNil(o.AddressLine) {
		return true
	}

	return false
}

// SetAddressLine gets a reference to the given []string and assigns it to the AddressLine field.
func (o *AddressType) SetAddressLine(v []string) {
	o.AddressLine = v
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *AddressType) GetCityName() string {
	if o == nil || IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetCityNameOk() (*string, bool) {
	if o == nil || IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *AddressType) HasCityName() bool {
	if o != nil && !IsNil(o.CityName) {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *AddressType) SetCityName(v string) {
	o.CityName = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressType) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressType) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressType) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCityExtension returns the CityExtension field value if set, zero value otherwise.
func (o *AddressType) GetCityExtension() string {
	if o == nil || IsNil(o.CityExtension) {
		var ret string
		return ret
	}
	return *o.CityExtension
}

// GetCityExtensionOk returns a tuple with the CityExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetCityExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.CityExtension) {
		return nil, false
	}
	return o.CityExtension, true
}

// HasCityExtension returns a boolean if a field has been set.
func (o *AddressType) HasCityExtension() bool {
	if o != nil && !IsNil(o.CityExtension) {
		return true
	}

	return false
}

// SetCityExtension gets a reference to the given string and assigns it to the CityExtension field.
func (o *AddressType) SetCityExtension(v string) {
	o.CityExtension = &v
}

// GetCounty returns the County field value if set, zero value otherwise.
func (o *AddressType) GetCounty() string {
	if o == nil || IsNil(o.County) {
		var ret string
		return ret
	}
	return *o.County
}

// GetCountyOk returns a tuple with the County field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetCountyOk() (*string, bool) {
	if o == nil || IsNil(o.County) {
		return nil, false
	}
	return o.County, true
}

// HasCounty returns a boolean if a field has been set.
func (o *AddressType) HasCounty() bool {
	if o != nil && !IsNil(o.County) {
		return true
	}

	return false
}

// SetCounty gets a reference to the given string and assigns it to the County field.
func (o *AddressType) SetCounty(v string) {
	o.County = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressType) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressType) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressType) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressType) GetCountry() CountryNameType {
	if o == nil || IsNil(o.Country) {
		var ret CountryNameType
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetCountryOk() (*CountryNameType, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressType) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given CountryNameType and assigns it to the Country field.
func (o *AddressType) SetCountry(v CountryNameType) {
	o.Country = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddressType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AddressType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AddressType) SetType(v string) {
	o.Type = &v
}

// GetTypeDescription returns the TypeDescription field value if set, zero value otherwise.
func (o *AddressType) GetTypeDescription() string {
	if o == nil || IsNil(o.TypeDescription) {
		var ret string
		return ret
	}
	return *o.TypeDescription
}

// GetTypeDescriptionOk returns a tuple with the TypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressType) GetTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TypeDescription) {
		return nil, false
	}
	return o.TypeDescription, true
}

// HasTypeDescription returns a boolean if a field has been set.
func (o *AddressType) HasTypeDescription() bool {
	if o != nil && !IsNil(o.TypeDescription) {
		return true
	}

	return false
}

// SetTypeDescription gets a reference to the given string and assigns it to the TypeDescription field.
func (o *AddressType) SetTypeDescription(v string) {
	o.TypeDescription = &v
}

func (o AddressType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsValidated) {
		toSerialize["isValidated"] = o.IsValidated
	}
	if !IsNil(o.AddressLine) {
		toSerialize["addressLine"] = o.AddressLine
	}
	if !IsNil(o.CityName) {
		toSerialize["cityName"] = o.CityName
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.CityExtension) {
		toSerialize["cityExtension"] = o.CityExtension
	}
	if !IsNil(o.County) {
		toSerialize["county"] = o.County
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeDescription) {
		toSerialize["typeDescription"] = o.TypeDescription
	}
	return toSerialize, nil
}

type NullableAddressType struct {
	value *AddressType
	isSet bool
}

func (v NullableAddressType) Get() *AddressType {
	return v.value
}

func (v *NullableAddressType) Set(val *AddressType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressType(val *AddressType) *NullableAddressType {
	return &NullableAddressType{value: val, isSet: true}
}

func (v NullableAddressType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


