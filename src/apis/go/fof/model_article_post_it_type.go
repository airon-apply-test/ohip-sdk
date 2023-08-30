/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the ArticlePostItType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticlePostItType{}

// ArticlePostItType Post it information of an article.
type ArticlePostItType struct {
	// Indicates whether the article is available for post it.
	AvailableForPostIt *bool `json:"availableForPostIt,omitempty"`
	Color *ColorType `json:"color,omitempty"`
}

// NewArticlePostItType instantiates a new ArticlePostItType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticlePostItType() *ArticlePostItType {
	this := ArticlePostItType{}
	return &this
}

// NewArticlePostItTypeWithDefaults instantiates a new ArticlePostItType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticlePostItTypeWithDefaults() *ArticlePostItType {
	this := ArticlePostItType{}
	return &this
}

// GetAvailableForPostIt returns the AvailableForPostIt field value if set, zero value otherwise.
func (o *ArticlePostItType) GetAvailableForPostIt() bool {
	if o == nil || IsNil(o.AvailableForPostIt) {
		var ret bool
		return ret
	}
	return *o.AvailableForPostIt
}

// GetAvailableForPostItOk returns a tuple with the AvailableForPostIt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticlePostItType) GetAvailableForPostItOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailableForPostIt) {
		return nil, false
	}
	return o.AvailableForPostIt, true
}

// HasAvailableForPostIt returns a boolean if a field has been set.
func (o *ArticlePostItType) HasAvailableForPostIt() bool {
	if o != nil && !IsNil(o.AvailableForPostIt) {
		return true
	}

	return false
}

// SetAvailableForPostIt gets a reference to the given bool and assigns it to the AvailableForPostIt field.
func (o *ArticlePostItType) SetAvailableForPostIt(v bool) {
	o.AvailableForPostIt = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ArticlePostItType) GetColor() ColorType {
	if o == nil || IsNil(o.Color) {
		var ret ColorType
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticlePostItType) GetColorOk() (*ColorType, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ArticlePostItType) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given ColorType and assigns it to the Color field.
func (o *ArticlePostItType) SetColor(v ColorType) {
	o.Color = &v
}

func (o ArticlePostItType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticlePostItType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableForPostIt) {
		toSerialize["availableForPostIt"] = o.AvailableForPostIt
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	return toSerialize, nil
}

type NullableArticlePostItType struct {
	value *ArticlePostItType
	isSet bool
}

func (v NullableArticlePostItType) Get() *ArticlePostItType {
	return v.value
}

func (v *NullableArticlePostItType) Set(val *ArticlePostItType) {
	v.value = val
	v.isSet = true
}

func (v NullableArticlePostItType) IsSet() bool {
	return v.isSet
}

func (v *NullableArticlePostItType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticlePostItType(val *ArticlePostItType) *NullableArticlePostItType {
	return &NullableArticlePostItType{value: val, isSet: true}
}

func (v NullableArticlePostItType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticlePostItType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


