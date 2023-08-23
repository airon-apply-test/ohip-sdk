/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransferCompTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferCompTransactionsRequest{}

// TransferCompTransactionsRequest struct for TransferCompTransactionsRequest
type TransferCompTransactionsRequest struct {
	Criteria *TransferCompTransactionsCriteria `json:"criteria,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewTransferCompTransactionsRequest instantiates a new TransferCompTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCompTransactionsRequest() *TransferCompTransactionsRequest {
	this := TransferCompTransactionsRequest{}
	return &this
}

// NewTransferCompTransactionsRequestWithDefaults instantiates a new TransferCompTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCompTransactionsRequestWithDefaults() *TransferCompTransactionsRequest {
	this := TransferCompTransactionsRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *TransferCompTransactionsRequest) GetCriteria() TransferCompTransactionsCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret TransferCompTransactionsCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCompTransactionsRequest) GetCriteriaOk() (*TransferCompTransactionsCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *TransferCompTransactionsRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given TransferCompTransactionsCriteria and assigns it to the Criteria field.
func (o *TransferCompTransactionsRequest) SetCriteria(v TransferCompTransactionsCriteria) {
	o.Criteria = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TransferCompTransactionsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCompTransactionsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TransferCompTransactionsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *TransferCompTransactionsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o TransferCompTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferCompTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableTransferCompTransactionsRequest struct {
	value *TransferCompTransactionsRequest
	isSet bool
}

func (v NullableTransferCompTransactionsRequest) Get() *TransferCompTransactionsRequest {
	return v.value
}

func (v *NullableTransferCompTransactionsRequest) Set(val *TransferCompTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCompTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCompTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCompTransactionsRequest(val *TransferCompTransactionsRequest) *NullableTransferCompTransactionsRequest {
	return &NullableTransferCompTransactionsRequest{value: val, isSet: true}
}

func (v NullableTransferCompTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCompTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


