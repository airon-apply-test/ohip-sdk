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

// checks if the TransferFolioToWindowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferFolioToWindowRequest{}

// TransferFolioToWindowRequest struct for TransferFolioToWindowRequest
type TransferFolioToWindowRequest struct {
	Criteria *FolioTransferCriteriaType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewTransferFolioToWindowRequest instantiates a new TransferFolioToWindowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferFolioToWindowRequest() *TransferFolioToWindowRequest {
	this := TransferFolioToWindowRequest{}
	return &this
}

// NewTransferFolioToWindowRequestWithDefaults instantiates a new TransferFolioToWindowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferFolioToWindowRequestWithDefaults() *TransferFolioToWindowRequest {
	this := TransferFolioToWindowRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *TransferFolioToWindowRequest) GetCriteria() FolioTransferCriteriaType {
	if o == nil || IsNil(o.Criteria) {
		var ret FolioTransferCriteriaType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferFolioToWindowRequest) GetCriteriaOk() (*FolioTransferCriteriaType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *TransferFolioToWindowRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given FolioTransferCriteriaType and assigns it to the Criteria field.
func (o *TransferFolioToWindowRequest) SetCriteria(v FolioTransferCriteriaType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TransferFolioToWindowRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferFolioToWindowRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TransferFolioToWindowRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *TransferFolioToWindowRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TransferFolioToWindowRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferFolioToWindowRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TransferFolioToWindowRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *TransferFolioToWindowRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o TransferFolioToWindowRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferFolioToWindowRequest) ToMap() (map[string]interface{}, error) {
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

type NullableTransferFolioToWindowRequest struct {
	value *TransferFolioToWindowRequest
	isSet bool
}

func (v NullableTransferFolioToWindowRequest) Get() *TransferFolioToWindowRequest {
	return v.value
}

func (v *NullableTransferFolioToWindowRequest) Set(val *TransferFolioToWindowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferFolioToWindowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferFolioToWindowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferFolioToWindowRequest(val *TransferFolioToWindowRequest) *NullableTransferFolioToWindowRequest {
	return &NullableTransferFolioToWindowRequest{value: val, isSet: true}
}

func (v NullableTransferFolioToWindowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferFolioToWindowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


