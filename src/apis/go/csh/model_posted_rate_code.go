/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the PostedRateCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostedRateCode{}

// PostedRateCode Response for the operation that posts a Rate Code amount on the reservation.
type PostedRateCode struct {
	// List of postings.
	Postings []SummaryPostingType `json:"postings,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostedRateCode instantiates a new PostedRateCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostedRateCode() *PostedRateCode {
	this := PostedRateCode{}
	return &this
}

// NewPostedRateCodeWithDefaults instantiates a new PostedRateCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostedRateCodeWithDefaults() *PostedRateCode {
	this := PostedRateCode{}
	return &this
}

// GetPostings returns the Postings field value if set, zero value otherwise.
func (o *PostedRateCode) GetPostings() []SummaryPostingType {
	if o == nil || IsNil(o.Postings) {
		var ret []SummaryPostingType
		return ret
	}
	return o.Postings
}

// GetPostingsOk returns a tuple with the Postings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostedRateCode) GetPostingsOk() ([]SummaryPostingType, bool) {
	if o == nil || IsNil(o.Postings) {
		return nil, false
	}
	return o.Postings, true
}

// HasPostings returns a boolean if a field has been set.
func (o *PostedRateCode) HasPostings() bool {
	if o != nil && !IsNil(o.Postings) {
		return true
	}

	return false
}

// SetPostings gets a reference to the given []SummaryPostingType and assigns it to the Postings field.
func (o *PostedRateCode) SetPostings(v []SummaryPostingType) {
	o.Postings = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostedRateCode) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostedRateCode) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostedRateCode) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostedRateCode) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostedRateCode) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostedRateCode) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostedRateCode) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostedRateCode) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostedRateCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostedRateCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Postings) {
		toSerialize["postings"] = o.Postings
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostedRateCode struct {
	value *PostedRateCode
	isSet bool
}

func (v NullablePostedRateCode) Get() *PostedRateCode {
	return v.value
}

func (v *NullablePostedRateCode) Set(val *PostedRateCode) {
	v.value = val
	v.isSet = true
}

func (v NullablePostedRateCode) IsSet() bool {
	return v.isSet
}

func (v *NullablePostedRateCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostedRateCode(val *PostedRateCode) *NullablePostedRateCode {
	return &NullablePostedRateCode{value: val, isSet: true}
}

func (v NullablePostedRateCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostedRateCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


