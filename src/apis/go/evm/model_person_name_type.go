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
	// Family name, last name. May also be used for full name if the sending system does not have the ability to separate a full name into its parts, e.g. the surname elementSpace may be used to pass the full name.
	Surname *string `json:"surname,omitempty"`
	// Hold various name suffixes and letters (e.g. Jr., Sr., III, Ret., Esq.)
	NameSuffix *string `json:"nameSuffix,omitempty"`
	// Degree or honors (e.g., Ph.D., M.D.)
	NameTitle *string `json:"nameTitle,omitempty"`
	// Title Suffix. Must be populated if ADVANCED_TITLE is on.
	NameTitleSuffix *int32 `json:"nameTitleSuffix,omitempty"`
	NameType *PersonNameTypeType `json:"nameType,omitempty"`
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

// GetNameSuffix returns the NameSuffix field value if set, zero value otherwise.
func (o *PersonNameType) GetNameSuffix() string {
	if o == nil || IsNil(o.NameSuffix) {
		var ret string
		return ret
	}
	return *o.NameSuffix
}

// GetNameSuffixOk returns a tuple with the NameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetNameSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.NameSuffix) {
		return nil, false
	}
	return o.NameSuffix, true
}

// HasNameSuffix returns a boolean if a field has been set.
func (o *PersonNameType) HasNameSuffix() bool {
	if o != nil && !IsNil(o.NameSuffix) {
		return true
	}

	return false
}

// SetNameSuffix gets a reference to the given string and assigns it to the NameSuffix field.
func (o *PersonNameType) SetNameSuffix(v string) {
	o.NameSuffix = &v
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

// GetNameTitleSuffix returns the NameTitleSuffix field value if set, zero value otherwise.
func (o *PersonNameType) GetNameTitleSuffix() int32 {
	if o == nil || IsNil(o.NameTitleSuffix) {
		var ret int32
		return ret
	}
	return *o.NameTitleSuffix
}

// GetNameTitleSuffixOk returns a tuple with the NameTitleSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonNameType) GetNameTitleSuffixOk() (*int32, bool) {
	if o == nil || IsNil(o.NameTitleSuffix) {
		return nil, false
	}
	return o.NameTitleSuffix, true
}

// HasNameTitleSuffix returns a boolean if a field has been set.
func (o *PersonNameType) HasNameTitleSuffix() bool {
	if o != nil && !IsNil(o.NameTitleSuffix) {
		return true
	}

	return false
}

// SetNameTitleSuffix gets a reference to the given int32 and assigns it to the NameTitleSuffix field.
func (o *PersonNameType) SetNameTitleSuffix(v int32) {
	o.NameTitleSuffix = &v
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
	if !IsNil(o.NameSuffix) {
		toSerialize["nameSuffix"] = o.NameSuffix
	}
	if !IsNil(o.NameTitle) {
		toSerialize["nameTitle"] = o.NameTitle
	}
	if !IsNil(o.NameTitleSuffix) {
		toSerialize["nameTitleSuffix"] = o.NameTitleSuffix
	}
	if !IsNil(o.NameType) {
		toSerialize["nameType"] = o.NameType
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


