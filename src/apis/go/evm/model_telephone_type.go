/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TelephoneType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelephoneType{}

// TelephoneType Information on a telephone number for the customer.
type TelephoneType struct {
	// Indicates type of technology associated with this telephone number, such as Voice, Data, Fax, Pager, Mobile, TTY, etc.
	PhoneTechType *string `json:"phoneTechType,omitempty"`
	// Describes the type of telephone number, in the context of its general use (e.g. Home, Business, Emergency Contact, Travel Arranger, Day, Evening).
	PhoneUseType *string `json:"phoneUseType,omitempty"`
	// Description of the PhoneUseType code
	PhoneUseTypeDescription *string `json:"phoneUseTypeDescription,omitempty"`
	// Telephone number assigned to a single location.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// NewTelephoneType instantiates a new TelephoneType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephoneType() *TelephoneType {
	this := TelephoneType{}
	return &this
}

// NewTelephoneTypeWithDefaults instantiates a new TelephoneType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephoneTypeWithDefaults() *TelephoneType {
	this := TelephoneType{}
	return &this
}

// GetPhoneTechType returns the PhoneTechType field value if set, zero value otherwise.
func (o *TelephoneType) GetPhoneTechType() string {
	if o == nil || IsNil(o.PhoneTechType) {
		var ret string
		return ret
	}
	return *o.PhoneTechType
}

// GetPhoneTechTypeOk returns a tuple with the PhoneTechType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneType) GetPhoneTechTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneTechType) {
		return nil, false
	}
	return o.PhoneTechType, true
}

// HasPhoneTechType returns a boolean if a field has been set.
func (o *TelephoneType) HasPhoneTechType() bool {
	if o != nil && !IsNil(o.PhoneTechType) {
		return true
	}

	return false
}

// SetPhoneTechType gets a reference to the given string and assigns it to the PhoneTechType field.
func (o *TelephoneType) SetPhoneTechType(v string) {
	o.PhoneTechType = &v
}

// GetPhoneUseType returns the PhoneUseType field value if set, zero value otherwise.
func (o *TelephoneType) GetPhoneUseType() string {
	if o == nil || IsNil(o.PhoneUseType) {
		var ret string
		return ret
	}
	return *o.PhoneUseType
}

// GetPhoneUseTypeOk returns a tuple with the PhoneUseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneType) GetPhoneUseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneUseType) {
		return nil, false
	}
	return o.PhoneUseType, true
}

// HasPhoneUseType returns a boolean if a field has been set.
func (o *TelephoneType) HasPhoneUseType() bool {
	if o != nil && !IsNil(o.PhoneUseType) {
		return true
	}

	return false
}

// SetPhoneUseType gets a reference to the given string and assigns it to the PhoneUseType field.
func (o *TelephoneType) SetPhoneUseType(v string) {
	o.PhoneUseType = &v
}

// GetPhoneUseTypeDescription returns the PhoneUseTypeDescription field value if set, zero value otherwise.
func (o *TelephoneType) GetPhoneUseTypeDescription() string {
	if o == nil || IsNil(o.PhoneUseTypeDescription) {
		var ret string
		return ret
	}
	return *o.PhoneUseTypeDescription
}

// GetPhoneUseTypeDescriptionOk returns a tuple with the PhoneUseTypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneType) GetPhoneUseTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneUseTypeDescription) {
		return nil, false
	}
	return o.PhoneUseTypeDescription, true
}

// HasPhoneUseTypeDescription returns a boolean if a field has been set.
func (o *TelephoneType) HasPhoneUseTypeDescription() bool {
	if o != nil && !IsNil(o.PhoneUseTypeDescription) {
		return true
	}

	return false
}

// SetPhoneUseTypeDescription gets a reference to the given string and assigns it to the PhoneUseTypeDescription field.
func (o *TelephoneType) SetPhoneUseTypeDescription(v string) {
	o.PhoneUseTypeDescription = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *TelephoneType) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephoneType) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *TelephoneType) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *TelephoneType) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o TelephoneType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelephoneType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhoneTechType) {
		toSerialize["phoneTechType"] = o.PhoneTechType
	}
	if !IsNil(o.PhoneUseType) {
		toSerialize["phoneUseType"] = o.PhoneUseType
	}
	if !IsNil(o.PhoneUseTypeDescription) {
		toSerialize["phoneUseTypeDescription"] = o.PhoneUseTypeDescription
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableTelephoneType struct {
	value *TelephoneType
	isSet bool
}

func (v NullableTelephoneType) Get() *TelephoneType {
	return v.value
}

func (v *NullableTelephoneType) Set(val *TelephoneType) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephoneType) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephoneType(val *TelephoneType) *NullableTelephoneType {
	return &NullableTelephoneType{value: val, isSet: true}
}

func (v NullableTelephoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


