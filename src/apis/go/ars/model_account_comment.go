/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
)

// checks if the AccountComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountComment{}

// AccountComment Request to create a Comment for an Account.
type AccountComment struct {
	CommentInfo *ARAccountCommentCriteriaType `json:"commentInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewAccountComment instantiates a new AccountComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountComment() *AccountComment {
	this := AccountComment{}
	return &this
}

// NewAccountCommentWithDefaults instantiates a new AccountComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommentWithDefaults() *AccountComment {
	this := AccountComment{}
	return &this
}

// GetCommentInfo returns the CommentInfo field value if set, zero value otherwise.
func (o *AccountComment) GetCommentInfo() ARAccountCommentCriteriaType {
	if o == nil || IsNil(o.CommentInfo) {
		var ret ARAccountCommentCriteriaType
		return ret
	}
	return *o.CommentInfo
}

// GetCommentInfoOk returns a tuple with the CommentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountComment) GetCommentInfoOk() (*ARAccountCommentCriteriaType, bool) {
	if o == nil || IsNil(o.CommentInfo) {
		return nil, false
	}
	return o.CommentInfo, true
}

// HasCommentInfo returns a boolean if a field has been set.
func (o *AccountComment) HasCommentInfo() bool {
	if o != nil && !IsNil(o.CommentInfo) {
		return true
	}

	return false
}

// SetCommentInfo gets a reference to the given ARAccountCommentCriteriaType and assigns it to the CommentInfo field.
func (o *AccountComment) SetCommentInfo(v ARAccountCommentCriteriaType) {
	o.CommentInfo = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AccountComment) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountComment) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AccountComment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *AccountComment) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AccountComment) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountComment) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AccountComment) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *AccountComment) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o AccountComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommentInfo) {
		toSerialize["commentInfo"] = o.CommentInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAccountComment struct {
	value *AccountComment
	isSet bool
}

func (v NullableAccountComment) Get() *AccountComment {
	return v.value
}

func (v *NullableAccountComment) Set(val *AccountComment) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountComment) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountComment(val *AccountComment) *NullableAccountComment {
	return &NullableAccountComment{value: val, isSet: true}
}

func (v NullableAccountComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


