/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GuestTitleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuestTitleType{}

// GuestTitleType Information representation of Guest Title.
type GuestTitleType struct {
	// Description of the Guest Title.
	Description *string `json:"description,omitempty"`
	// Business Title for advanced title configuration.
	Greeting *string `json:"greeting,omitempty"`
	// Guest Title record sequence number.
	DisplayOrder *float32 `json:"displayOrder,omitempty"`
	// Description of the Guest Title.
	NewTitleType *int32 `json:"newTitleType,omitempty"`
	// Description of the Guest Title.
	NewLanguageCode *string `json:"newLanguageCode,omitempty"`
	// Code of the Guest Title.
	Code *string `json:"code,omitempty"`
	// Title Type for advanced title configuration.
	TitleType *int32 `json:"titleType,omitempty"`
	// Language code of the Guest Title.
	LanguageCode *string `json:"languageCode,omitempty"`
}

// NewGuestTitleType instantiates a new GuestTitleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestTitleType() *GuestTitleType {
	this := GuestTitleType{}
	return &this
}

// NewGuestTitleTypeWithDefaults instantiates a new GuestTitleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestTitleTypeWithDefaults() *GuestTitleType {
	this := GuestTitleType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GuestTitleType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GuestTitleType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GuestTitleType) SetDescription(v string) {
	o.Description = &v
}

// GetGreeting returns the Greeting field value if set, zero value otherwise.
func (o *GuestTitleType) GetGreeting() string {
	if o == nil || IsNil(o.Greeting) {
		var ret string
		return ret
	}
	return *o.Greeting
}

// GetGreetingOk returns a tuple with the Greeting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetGreetingOk() (*string, bool) {
	if o == nil || IsNil(o.Greeting) {
		return nil, false
	}
	return o.Greeting, true
}

// HasGreeting returns a boolean if a field has been set.
func (o *GuestTitleType) HasGreeting() bool {
	if o != nil && !IsNil(o.Greeting) {
		return true
	}

	return false
}

// SetGreeting gets a reference to the given string and assigns it to the Greeting field.
func (o *GuestTitleType) SetGreeting(v string) {
	o.Greeting = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *GuestTitleType) GetDisplayOrder() float32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret float32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetDisplayOrderOk() (*float32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *GuestTitleType) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given float32 and assigns it to the DisplayOrder field.
func (o *GuestTitleType) SetDisplayOrder(v float32) {
	o.DisplayOrder = &v
}

// GetNewTitleType returns the NewTitleType field value if set, zero value otherwise.
func (o *GuestTitleType) GetNewTitleType() int32 {
	if o == nil || IsNil(o.NewTitleType) {
		var ret int32
		return ret
	}
	return *o.NewTitleType
}

// GetNewTitleTypeOk returns a tuple with the NewTitleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetNewTitleTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.NewTitleType) {
		return nil, false
	}
	return o.NewTitleType, true
}

// HasNewTitleType returns a boolean if a field has been set.
func (o *GuestTitleType) HasNewTitleType() bool {
	if o != nil && !IsNil(o.NewTitleType) {
		return true
	}

	return false
}

// SetNewTitleType gets a reference to the given int32 and assigns it to the NewTitleType field.
func (o *GuestTitleType) SetNewTitleType(v int32) {
	o.NewTitleType = &v
}

// GetNewLanguageCode returns the NewLanguageCode field value if set, zero value otherwise.
func (o *GuestTitleType) GetNewLanguageCode() string {
	if o == nil || IsNil(o.NewLanguageCode) {
		var ret string
		return ret
	}
	return *o.NewLanguageCode
}

// GetNewLanguageCodeOk returns a tuple with the NewLanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetNewLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.NewLanguageCode) {
		return nil, false
	}
	return o.NewLanguageCode, true
}

// HasNewLanguageCode returns a boolean if a field has been set.
func (o *GuestTitleType) HasNewLanguageCode() bool {
	if o != nil && !IsNil(o.NewLanguageCode) {
		return true
	}

	return false
}

// SetNewLanguageCode gets a reference to the given string and assigns it to the NewLanguageCode field.
func (o *GuestTitleType) SetNewLanguageCode(v string) {
	o.NewLanguageCode = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GuestTitleType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GuestTitleType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GuestTitleType) SetCode(v string) {
	o.Code = &v
}

// GetTitleType returns the TitleType field value if set, zero value otherwise.
func (o *GuestTitleType) GetTitleType() int32 {
	if o == nil || IsNil(o.TitleType) {
		var ret int32
		return ret
	}
	return *o.TitleType
}

// GetTitleTypeOk returns a tuple with the TitleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetTitleTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TitleType) {
		return nil, false
	}
	return o.TitleType, true
}

// HasTitleType returns a boolean if a field has been set.
func (o *GuestTitleType) HasTitleType() bool {
	if o != nil && !IsNil(o.TitleType) {
		return true
	}

	return false
}

// SetTitleType gets a reference to the given int32 and assigns it to the TitleType field.
func (o *GuestTitleType) SetTitleType(v int32) {
	o.TitleType = &v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *GuestTitleType) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestTitleType) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *GuestTitleType) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *GuestTitleType) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

func (o GuestTitleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuestTitleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Greeting) {
		toSerialize["greeting"] = o.Greeting
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.NewTitleType) {
		toSerialize["newTitleType"] = o.NewTitleType
	}
	if !IsNil(o.NewLanguageCode) {
		toSerialize["newLanguageCode"] = o.NewLanguageCode
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.TitleType) {
		toSerialize["titleType"] = o.TitleType
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["languageCode"] = o.LanguageCode
	}
	return toSerialize, nil
}

type NullableGuestTitleType struct {
	value *GuestTitleType
	isSet bool
}

func (v NullableGuestTitleType) Get() *GuestTitleType {
	return v.value
}

func (v *NullableGuestTitleType) Set(val *GuestTitleType) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestTitleType) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestTitleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestTitleType(val *GuestTitleType) *NullableGuestTitleType {
	return &NullableGuestTitleType{value: val, isSet: true}
}

func (v NullableGuestTitleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestTitleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


