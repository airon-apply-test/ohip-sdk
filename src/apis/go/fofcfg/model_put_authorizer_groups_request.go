/*
OPERA Cloud Front Desk Configuration API

APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fofcfg

import (
	"encoding/json"
)

// checks if the PutAuthorizerGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutAuthorizerGroupsRequest{}

// PutAuthorizerGroupsRequest struct for PutAuthorizerGroupsRequest
type PutAuthorizerGroupsRequest struct {
	Criteria *AuthorizerGroupsCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutAuthorizerGroupsRequest instantiates a new PutAuthorizerGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutAuthorizerGroupsRequest() *PutAuthorizerGroupsRequest {
	this := PutAuthorizerGroupsRequest{}
	return &this
}

// NewPutAuthorizerGroupsRequestWithDefaults instantiates a new PutAuthorizerGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutAuthorizerGroupsRequestWithDefaults() *PutAuthorizerGroupsRequest {
	this := PutAuthorizerGroupsRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PutAuthorizerGroupsRequest) GetCriteria() AuthorizerGroupsCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret AuthorizerGroupsCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthorizerGroupsRequest) GetCriteriaOk() (*AuthorizerGroupsCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PutAuthorizerGroupsRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AuthorizerGroupsCriteriaType and assigns it to the Criteria field.
func (o *PutAuthorizerGroupsRequest) SetCriteria(v AuthorizerGroupsCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutAuthorizerGroupsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthorizerGroupsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutAuthorizerGroupsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutAuthorizerGroupsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutAuthorizerGroupsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthorizerGroupsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutAuthorizerGroupsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutAuthorizerGroupsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutAuthorizerGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutAuthorizerGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutAuthorizerGroupsRequest struct {
	value *PutAuthorizerGroupsRequest
	isSet bool
}

func (v NullablePutAuthorizerGroupsRequest) Get() *PutAuthorizerGroupsRequest {
	return v.value
}

func (v *NullablePutAuthorizerGroupsRequest) Set(val *PutAuthorizerGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutAuthorizerGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutAuthorizerGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutAuthorizerGroupsRequest(val *PutAuthorizerGroupsRequest) *NullablePutAuthorizerGroupsRequest {
	return &NullablePutAuthorizerGroupsRequest{value: val, isSet: true}
}

func (v NullablePutAuthorizerGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutAuthorizerGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


