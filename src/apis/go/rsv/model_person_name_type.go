/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PersonNameType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonNameType{}

// PersonNameType This provides name information for a person.
type PersonNameType struct {
	// Salutation of honorific (e.g. Mr., Mrs., Ms., Miss, Dr.)
	NamePrefix *string `json:"namePrefix,omitempty"`
	// Given name, first name or names.
	GivenName *string `json:"givenName,omitempty"`
	// The middle name of the person name.
	MiddleName *string `json:"middleName,omitempty"`
	// Family name, last name. May also be used for full name if the sending system does not have the ability to separate a full name into its parts, e.g. the surname element may be used to pass the full name.
	Surname *string `json:"surname,omitempty"`
	// Degree or honors (e.g., Ph.D., M.D.)
	NameTitle *string `json:"nameTitle,omitempty"`
	// Salutation of the profile
	Salutation *string `json:"salutation,omitempty"`
	NameType *PersonNameTypeType `json:"nameType,omitempty"`
	// Language identification.
	Language *string `json:"language,omitempty"`
}

// NewPersonNameType instantiates a new PersonNameType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonNameType() *PersonNameType {
	this := PersonNameType{}
	return &this
}

// NewPersonNameTypeWithDefaults instantiates a new PersonNameType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonNameTypeWithDefaults() *PersonNameType {
	this := PersonNameType{}
	return &this
}

// GetNamePrefix returns the NamePrefix field value if set, zero value otherwise.
func (o *PersonNameType) GetNamePrefix() string {
	if o == nil || IsNil(o.NamePrefix) {
		var ret string
		return ret
	}
	return *o.NamePrefix
}

// GetNamePrefixOk returns a tuple with the NamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetNamePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.NamePrefix) {
		return nil, false
	}
	return o.NamePrefix, true
}

// HasNamePrefix returns a boolean if a field has been set.
func (o *PersonNameType) HasNamePrefix() bool {
	if o != nil && !IsNil(o.NamePrefix) {
		return true
	}

	return false
}

// SetNamePrefix gets a reference to the given string and assigns it to the NamePrefix field.
func (o *PersonNameType) SetNamePrefix(v string) {
	o.NamePrefix = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *PersonNameType) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *PersonNameType) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *PersonNameType) SetGivenName(v string) {
	o.GivenName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *PersonNameType) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *PersonNameType) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *PersonNameType) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *PersonNameType) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *PersonNameType) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *PersonNameType) SetSurname(v string) {
	o.Surname = &v
}

// GetNameTitle returns the NameTitle field value if set, zero value otherwise.
func (o *PersonNameType) GetNameTitle() string {
	if o == nil || IsNil(o.NameTitle) {
		var ret string
		return ret
	}
	return *o.NameTitle
}

// GetNameTitleOk returns a tuple with the NameTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetNameTitleOk() (*string, bool) {
	if o == nil || IsNil(o.NameTitle) {
		return nil, false
	}
	return o.NameTitle, true
}

// HasNameTitle returns a boolean if a field has been set.
func (o *PersonNameType) HasNameTitle() bool {
	if o != nil && !IsNil(o.NameTitle) {
		return true
	}

	return false
}

// SetNameTitle gets a reference to the given string and assigns it to the NameTitle field.
func (o *PersonNameType) SetNameTitle(v string) {
	o.NameTitle = &v
}

// GetSalutation returns the Salutation field value if set, zero value otherwise.
func (o *PersonNameType) GetSalutation() string {
	if o == nil || IsNil(o.Salutation) {
		var ret string
		return ret
	}
	return *o.Salutation
}

// GetSalutationOk returns a tuple with the Salutation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetSalutationOk() (*string, bool) {
	if o == nil || IsNil(o.Salutation) {
		return nil, false
	}
	return o.Salutation, true
}

// HasSalutation returns a boolean if a field has been set.
func (o *PersonNameType) HasSalutation() bool {
	if o != nil && !IsNil(o.Salutation) {
		return true
	}

	return false
}

// SetSalutation gets a reference to the given string and assigns it to the Salutation field.
func (o *PersonNameType) SetSalutation(v string) {
	o.Salutation = &v
}

// GetNameType returns the NameType field value if set, zero value otherwise.
func (o *PersonNameType) GetNameType() PersonNameTypeType {
	if o == nil || IsNil(o.NameType) {
		var ret PersonNameTypeType
		return ret
	}
	return *o.NameType
}

// GetNameTypeOk returns a tuple with the NameType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetNameTypeOk() (*PersonNameTypeType, bool) {
	if o == nil || IsNil(o.NameType) {
		return nil, false
	}
	return o.NameType, true
}

// HasNameType returns a boolean if a field has been set.
func (o *PersonNameType) HasNameType() bool {
	if o != nil && !IsNil(o.NameType) {
		return true
	}

	return false
}

// SetNameType gets a reference to the given PersonNameTypeType and assigns it to the NameType field.
func (o *PersonNameType) SetNameType(v PersonNameTypeType) {
	o.NameType = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PersonNameType) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PersonNameType) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PersonNameType) SetLanguage(v string) {
	o.Language = &v
}

func (o PersonNameType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonNameType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NamePrefix) {
		toSerialize["namePrefix"] = o.NamePrefix
	}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.NameTitle) {
		toSerialize["nameTitle"] = o.NameTitle
	}
	if !IsNil(o.Salutation) {
		toSerialize["salutation"] = o.Salutation
	}
	if !IsNil(o.NameType) {
		toSerialize["nameType"] = o.NameType
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullablePersonNameType struct {
	value *PersonNameType
	isSet bool
}

func (v NullablePersonNameType) Get() *PersonNameType {
	return v.value
}

func (v *NullablePersonNameType) Set(val *PersonNameType) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonNameType) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonNameType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonNameType(val *PersonNameType) *NullablePersonNameType {
	return &NullablePersonNameType{value: val, isSet: true}
}

func (v NullablePersonNameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonNameType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


