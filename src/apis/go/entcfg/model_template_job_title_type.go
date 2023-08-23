/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TemplateJobTitleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateJobTitleType{}

// TemplateJobTitleType Base details common between both template and hotel job titles.
type TemplateJobTitleType struct {
	// Job Title Code.
	Code *string `json:"code,omitempty"`
	Description *TranslationTextType80 `json:"description,omitempty"`
	// Flag to indicate if display reservation closing script.
	DisplayClosingScript *bool `json:"displayClosingScript,omitempty"`
}

// NewTemplateJobTitleType instantiates a new TemplateJobTitleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateJobTitleType() *TemplateJobTitleType {
	this := TemplateJobTitleType{}
	return &this
}

// NewTemplateJobTitleTypeWithDefaults instantiates a new TemplateJobTitleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateJobTitleTypeWithDefaults() *TemplateJobTitleType {
	this := TemplateJobTitleType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TemplateJobTitleType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateJobTitleType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TemplateJobTitleType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TemplateJobTitleType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateJobTitleType) GetDescription() TranslationTextType80 {
	if o == nil || IsNil(o.Description) {
		var ret TranslationTextType80
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateJobTitleType) GetDescriptionOk() (*TranslationTextType80, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateJobTitleType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given TranslationTextType80 and assigns it to the Description field.
func (o *TemplateJobTitleType) SetDescription(v TranslationTextType80) {
	o.Description = &v
}

// GetDisplayClosingScript returns the DisplayClosingScript field value if set, zero value otherwise.
func (o *TemplateJobTitleType) GetDisplayClosingScript() bool {
	if o == nil || IsNil(o.DisplayClosingScript) {
		var ret bool
		return ret
	}
	return *o.DisplayClosingScript
}

// GetDisplayClosingScriptOk returns a tuple with the DisplayClosingScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateJobTitleType) GetDisplayClosingScriptOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayClosingScript) {
		return nil, false
	}
	return o.DisplayClosingScript, true
}

// HasDisplayClosingScript returns a boolean if a field has been set.
func (o *TemplateJobTitleType) HasDisplayClosingScript() bool {
	if o != nil && !IsNil(o.DisplayClosingScript) {
		return true
	}

	return false
}

// SetDisplayClosingScript gets a reference to the given bool and assigns it to the DisplayClosingScript field.
func (o *TemplateJobTitleType) SetDisplayClosingScript(v bool) {
	o.DisplayClosingScript = &v
}

func (o TemplateJobTitleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateJobTitleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayClosingScript) {
		toSerialize["displayClosingScript"] = o.DisplayClosingScript
	}
	return toSerialize, nil
}

type NullableTemplateJobTitleType struct {
	value *TemplateJobTitleType
	isSet bool
}

func (v NullableTemplateJobTitleType) Get() *TemplateJobTitleType {
	return v.value
}

func (v *NullableTemplateJobTitleType) Set(val *TemplateJobTitleType) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateJobTitleType) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateJobTitleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateJobTitleType(val *TemplateJobTitleType) *NullableTemplateJobTitleType {
	return &NullableTemplateJobTitleType{value: val, isSet: true}
}

func (v NullableTemplateJobTitleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateJobTitleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


