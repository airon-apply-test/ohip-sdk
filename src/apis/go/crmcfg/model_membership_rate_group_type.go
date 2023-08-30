/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmcfg

import (
	"encoding/json"
)

// checks if the MembershipRateGroupType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipRateGroupType{}

// MembershipRateGroupType Base type provides information about Membership Market/Property Groups Configuration.
type MembershipRateGroupType struct {
	// Code is used to identify a Membership Market/Resort Group.
	Code *string `json:"code,omitempty"`
	// Description of the Membership Market/Propety Group.
	Description *string `json:"description,omitempty"`
	// Membership Market/Property Groups display sequence Number
	DisplaySequence *float32 `json:"displaySequence,omitempty"`
	// Membership Rates code and Description.
	RateCodes []CodeDescriptionType `json:"rateCodes,omitempty"`
}

// NewMembershipRateGroupType instantiates a new MembershipRateGroupType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipRateGroupType() *MembershipRateGroupType {
	this := MembershipRateGroupType{}
	return &this
}

// NewMembershipRateGroupTypeWithDefaults instantiates a new MembershipRateGroupType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipRateGroupTypeWithDefaults() *MembershipRateGroupType {
	this := MembershipRateGroupType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MembershipRateGroupType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRateGroupType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MembershipRateGroupType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *MembershipRateGroupType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MembershipRateGroupType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRateGroupType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MembershipRateGroupType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MembershipRateGroupType) SetDescription(v string) {
	o.Description = &v
}

// GetDisplaySequence returns the DisplaySequence field value if set, zero value otherwise.
func (o *MembershipRateGroupType) GetDisplaySequence() float32 {
	if o == nil || IsNil(o.DisplaySequence) {
		var ret float32
		return ret
	}
	return *o.DisplaySequence
}

// GetDisplaySequenceOk returns a tuple with the DisplaySequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRateGroupType) GetDisplaySequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.DisplaySequence) {
		return nil, false
	}
	return o.DisplaySequence, true
}

// HasDisplaySequence returns a boolean if a field has been set.
func (o *MembershipRateGroupType) HasDisplaySequence() bool {
	if o != nil && !IsNil(o.DisplaySequence) {
		return true
	}

	return false
}

// SetDisplaySequence gets a reference to the given float32 and assigns it to the DisplaySequence field.
func (o *MembershipRateGroupType) SetDisplaySequence(v float32) {
	o.DisplaySequence = &v
}

// GetRateCodes returns the RateCodes field value if set, zero value otherwise.
func (o *MembershipRateGroupType) GetRateCodes() []CodeDescriptionType {
	if o == nil || IsNil(o.RateCodes) {
		var ret []CodeDescriptionType
		return ret
	}
	return o.RateCodes
}

// GetRateCodesOk returns a tuple with the RateCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipRateGroupType) GetRateCodesOk() ([]CodeDescriptionType, bool) {
	if o == nil || IsNil(o.RateCodes) {
		return nil, false
	}
	return o.RateCodes, true
}

// HasRateCodes returns a boolean if a field has been set.
func (o *MembershipRateGroupType) HasRateCodes() bool {
	if o != nil && !IsNil(o.RateCodes) {
		return true
	}

	return false
}

// SetRateCodes gets a reference to the given []CodeDescriptionType and assigns it to the RateCodes field.
func (o *MembershipRateGroupType) SetRateCodes(v []CodeDescriptionType) {
	o.RateCodes = v
}

func (o MembershipRateGroupType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipRateGroupType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplaySequence) {
		toSerialize["displaySequence"] = o.DisplaySequence
	}
	if !IsNil(o.RateCodes) {
		toSerialize["rateCodes"] = o.RateCodes
	}
	return toSerialize, nil
}

type NullableMembershipRateGroupType struct {
	value *MembershipRateGroupType
	isSet bool
}

func (v NullableMembershipRateGroupType) Get() *MembershipRateGroupType {
	return v.value
}

func (v *NullableMembershipRateGroupType) Set(val *MembershipRateGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipRateGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipRateGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipRateGroupType(val *MembershipRateGroupType) *NullableMembershipRateGroupType {
	return &NullableMembershipRateGroupType{value: val, isSet: true}
}

func (v NullableMembershipRateGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipRateGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


