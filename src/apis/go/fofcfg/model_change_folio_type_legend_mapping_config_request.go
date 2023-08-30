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

// checks if the ChangeFolioTypeLegendMappingConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeFolioTypeLegendMappingConfigRequest{}

// ChangeFolioTypeLegendMappingConfigRequest struct for ChangeFolioTypeLegendMappingConfigRequest
type ChangeFolioTypeLegendMappingConfigRequest struct {
	Criteria *FolioTypeLegendMappingConfigType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChangeFolioTypeLegendMappingConfigRequest instantiates a new ChangeFolioTypeLegendMappingConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeFolioTypeLegendMappingConfigRequest() *ChangeFolioTypeLegendMappingConfigRequest {
	this := ChangeFolioTypeLegendMappingConfigRequest{}
	return &this
}

// NewChangeFolioTypeLegendMappingConfigRequestWithDefaults instantiates a new ChangeFolioTypeLegendMappingConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeFolioTypeLegendMappingConfigRequestWithDefaults() *ChangeFolioTypeLegendMappingConfigRequest {
	this := ChangeFolioTypeLegendMappingConfigRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ChangeFolioTypeLegendMappingConfigRequest) GetCriteria() FolioTypeLegendMappingConfigType {
	if o == nil || IsNil(o.Criteria) {
		var ret FolioTypeLegendMappingConfigType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeFolioTypeLegendMappingConfigRequest) GetCriteriaOk() (*FolioTypeLegendMappingConfigType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ChangeFolioTypeLegendMappingConfigRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given FolioTypeLegendMappingConfigType and assigns it to the Criteria field.
func (o *ChangeFolioTypeLegendMappingConfigRequest) SetCriteria(v FolioTypeLegendMappingConfigType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangeFolioTypeLegendMappingConfigRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeFolioTypeLegendMappingConfigRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangeFolioTypeLegendMappingConfigRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangeFolioTypeLegendMappingConfigRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangeFolioTypeLegendMappingConfigRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeFolioTypeLegendMappingConfigRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangeFolioTypeLegendMappingConfigRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangeFolioTypeLegendMappingConfigRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChangeFolioTypeLegendMappingConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeFolioTypeLegendMappingConfigRequest) ToMap() (map[string]interface{}, error) {
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

type NullableChangeFolioTypeLegendMappingConfigRequest struct {
	value *ChangeFolioTypeLegendMappingConfigRequest
	isSet bool
}

func (v NullableChangeFolioTypeLegendMappingConfigRequest) Get() *ChangeFolioTypeLegendMappingConfigRequest {
	return v.value
}

func (v *NullableChangeFolioTypeLegendMappingConfigRequest) Set(val *ChangeFolioTypeLegendMappingConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeFolioTypeLegendMappingConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeFolioTypeLegendMappingConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeFolioTypeLegendMappingConfigRequest(val *ChangeFolioTypeLegendMappingConfigRequest) *NullableChangeFolioTypeLegendMappingConfigRequest {
	return &NullableChangeFolioTypeLegendMappingConfigRequest{value: val, isSet: true}
}

func (v NullableChangeFolioTypeLegendMappingConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeFolioTypeLegendMappingConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


