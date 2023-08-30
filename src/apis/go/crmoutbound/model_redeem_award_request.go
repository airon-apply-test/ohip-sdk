/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the RedeemAwardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedeemAwardRequest{}

// RedeemAwardRequest struct for RedeemAwardRequest
type RedeemAwardRequest struct {
	MemberAward *MemberAwardType `json:"memberAward,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
}

// NewRedeemAwardRequest instantiates a new RedeemAwardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemAwardRequest() *RedeemAwardRequest {
	this := RedeemAwardRequest{}
	return &this
}

// NewRedeemAwardRequestWithDefaults instantiates a new RedeemAwardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemAwardRequestWithDefaults() *RedeemAwardRequest {
	this := RedeemAwardRequest{}
	return &this
}

// GetMemberAward returns the MemberAward field value if set, zero value otherwise.
func (o *RedeemAwardRequest) GetMemberAward() MemberAwardType {
	if o == nil || IsNil(o.MemberAward) {
		var ret MemberAwardType
		return ret
	}
	return *o.MemberAward
}

// GetMemberAwardOk returns a tuple with the MemberAward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemAwardRequest) GetMemberAwardOk() (*MemberAwardType, bool) {
	if o == nil || IsNil(o.MemberAward) {
		return nil, false
	}
	return o.MemberAward, true
}

// HasMemberAward returns a boolean if a field has been set.
func (o *RedeemAwardRequest) HasMemberAward() bool {
	if o != nil && !IsNil(o.MemberAward) {
		return true
	}

	return false
}

// SetMemberAward gets a reference to the given MemberAwardType and assigns it to the MemberAward field.
func (o *RedeemAwardRequest) SetMemberAward(v MemberAwardType) {
	o.MemberAward = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RedeemAwardRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemAwardRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RedeemAwardRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *RedeemAwardRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

func (o RedeemAwardRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemAwardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberAward) {
		toSerialize["memberAward"] = o.MemberAward
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableRedeemAwardRequest struct {
	value *RedeemAwardRequest
	isSet bool
}

func (v NullableRedeemAwardRequest) Get() *RedeemAwardRequest {
	return v.value
}

func (v *NullableRedeemAwardRequest) Set(val *RedeemAwardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemAwardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemAwardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemAwardRequest(val *RedeemAwardRequest) *NullableRedeemAwardRequest {
	return &NullableRedeemAwardRequest{value: val, isSet: true}
}

func (v NullableRedeemAwardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemAwardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


