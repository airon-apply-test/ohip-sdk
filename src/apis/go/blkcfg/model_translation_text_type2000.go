/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TranslationTextType2000 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranslationTextType2000{}

// TranslationTextType2000 Contains Multiple translated texts and language codes.
type TranslationTextType2000 struct {
	// Default text with Character length from 0 to 2000.
	DefaultText *string `json:"defaultText,omitempty"`
	// Language code for the translation.
	TranslatedTexts []TranslationsTextTypeInner `json:"translatedTexts,omitempty"`
}

// NewTranslationTextType2000 instantiates a new TranslationTextType2000 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationTextType2000() *TranslationTextType2000 {
	this := TranslationTextType2000{}
	return &this
}

// NewTranslationTextType2000WithDefaults instantiates a new TranslationTextType2000 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationTextType2000WithDefaults() *TranslationTextType2000 {
	this := TranslationTextType2000{}
	return &this
}

// GetDefaultText returns the DefaultText field value if set, zero value otherwise.
func (o *TranslationTextType2000) GetDefaultText() string {
	if o == nil || IsNil(o.DefaultText) {
		var ret string
		return ret
	}
	return *o.DefaultText
}

// GetDefaultTextOk returns a tuple with the DefaultText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationTextType2000) GetDefaultTextOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultText) {
		return nil, false
	}
	return o.DefaultText, true
}

// HasDefaultText returns a boolean if a field has been set.
func (o *TranslationTextType2000) HasDefaultText() bool {
	if o != nil && !IsNil(o.DefaultText) {
		return true
	}

	return false
}

// SetDefaultText gets a reference to the given string and assigns it to the DefaultText field.
func (o *TranslationTextType2000) SetDefaultText(v string) {
	o.DefaultText = &v
}

// GetTranslatedTexts returns the TranslatedTexts field value if set, zero value otherwise.
func (o *TranslationTextType2000) GetTranslatedTexts() []TranslationsTextTypeInner {
	if o == nil || IsNil(o.TranslatedTexts) {
		var ret []TranslationsTextTypeInner
		return ret
	}
	return o.TranslatedTexts
}

// GetTranslatedTextsOk returns a tuple with the TranslatedTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationTextType2000) GetTranslatedTextsOk() ([]TranslationsTextTypeInner, bool) {
	if o == nil || IsNil(o.TranslatedTexts) {
		return nil, false
	}
	return o.TranslatedTexts, true
}

// HasTranslatedTexts returns a boolean if a field has been set.
func (o *TranslationTextType2000) HasTranslatedTexts() bool {
	if o != nil && !IsNil(o.TranslatedTexts) {
		return true
	}

	return false
}

// SetTranslatedTexts gets a reference to the given []TranslationsTextTypeInner and assigns it to the TranslatedTexts field.
func (o *TranslationTextType2000) SetTranslatedTexts(v []TranslationsTextTypeInner) {
	o.TranslatedTexts = v
}

func (o TranslationTextType2000) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranslationTextType2000) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultText) {
		toSerialize["defaultText"] = o.DefaultText
	}
	if !IsNil(o.TranslatedTexts) {
		toSerialize["translatedTexts"] = o.TranslatedTexts
	}
	return toSerialize, nil
}

type NullableTranslationTextType2000 struct {
	value *TranslationTextType2000
	isSet bool
}

func (v NullableTranslationTextType2000) Get() *TranslationTextType2000 {
	return v.value
}

func (v *NullableTranslationTextType2000) Set(val *TranslationTextType2000) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationTextType2000) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationTextType2000) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationTextType2000(val *TranslationTextType2000) *NullableTranslationTextType2000 {
	return &NullableTranslationTextType2000{value: val, isSet: true}
}

func (v NullableTranslationTextType2000) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationTextType2000) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


