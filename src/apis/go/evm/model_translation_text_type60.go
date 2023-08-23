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

// checks if the TranslationTextType60 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranslationTextType60{}

// TranslationTextType60 Contains Multiple translated texts and language codes.
type TranslationTextType60 struct {
	// Default text with Character length from 0 to 60.
	DefaultText *string `json:"defaultText,omitempty"`
	// Language code for the translation.
	Translations []TranslationsTextTypeInner `json:"translations,omitempty"`
}

// NewTranslationTextType60 instantiates a new TranslationTextType60 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationTextType60() *TranslationTextType60 {
	this := TranslationTextType60{}
	return &this
}

// NewTranslationTextType60WithDefaults instantiates a new TranslationTextType60 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationTextType60WithDefaults() *TranslationTextType60 {
	this := TranslationTextType60{}
	return &this
}

// GetDefaultText returns the DefaultText field value if set, zero value otherwise.
func (o *TranslationTextType60) GetDefaultText() string {
	if o == nil || IsNil(o.DefaultText) {
		var ret string
		return ret
	}
	return *o.DefaultText
}

// GetDefaultTextOk returns a tuple with the DefaultText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationTextType60) GetDefaultTextOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultText) {
		return nil, false
	}
	return o.DefaultText, true
}

// HasDefaultText returns a boolean if a field has been set.
func (o *TranslationTextType60) HasDefaultText() bool {
	if o != nil && !IsNil(o.DefaultText) {
		return true
	}

	return false
}

// SetDefaultText gets a reference to the given string and assigns it to the DefaultText field.
func (o *TranslationTextType60) SetDefaultText(v string) {
	o.DefaultText = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *TranslationTextType60) GetTranslations() []TranslationsTextTypeInner {
	if o == nil || IsNil(o.Translations) {
		var ret []TranslationsTextTypeInner
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationTextType60) GetTranslationsOk() ([]TranslationsTextTypeInner, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *TranslationTextType60) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []TranslationsTextTypeInner and assigns it to the Translations field.
func (o *TranslationTextType60) SetTranslations(v []TranslationsTextTypeInner) {
	o.Translations = v
}

func (o TranslationTextType60) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranslationTextType60) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultText) {
		toSerialize["defaultText"] = o.DefaultText
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

type NullableTranslationTextType60 struct {
	value *TranslationTextType60
	isSet bool
}

func (v NullableTranslationTextType60) Get() *TranslationTextType60 {
	return v.value
}

func (v *NullableTranslationTextType60) Set(val *TranslationTextType60) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationTextType60) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationTextType60) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationTextType60(val *TranslationTextType60) *NullableTranslationTextType60 {
	return &NullableTranslationTextType60{value: val, isSet: true}
}

func (v NullableTranslationTextType60) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationTextType60) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


