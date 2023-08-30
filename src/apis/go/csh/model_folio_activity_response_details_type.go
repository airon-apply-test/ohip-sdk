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

// checks if the FolioActivityResponseDetailsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolioActivityResponseDetailsType{}

// FolioActivityResponseDetailsType Details of response for the Fiscal Folio Activity made.
type FolioActivityResponseDetailsType struct {
	// Business Date.
	BusinessDate *string `json:"businessDate,omitempty"`
	// Type of the response, possible values: ERROR, WARNING, RESPONSE.
	ResponseType *string `json:"responseType,omitempty"`
	// Name of the response element.
	ResponseName *string `json:"responseName,omitempty"`
	// Value of the response element.
	ResponseValue *string `json:"responseValue,omitempty"`
	// Number of fiscal response attempt made for the folio
	ResponseAttemptNo *int32 `json:"responseAttemptNo,omitempty"`
}

// NewFolioActivityResponseDetailsType instantiates a new FolioActivityResponseDetailsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolioActivityResponseDetailsType() *FolioActivityResponseDetailsType {
	this := FolioActivityResponseDetailsType{}
	return &this
}

// NewFolioActivityResponseDetailsTypeWithDefaults instantiates a new FolioActivityResponseDetailsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolioActivityResponseDetailsTypeWithDefaults() *FolioActivityResponseDetailsType {
	this := FolioActivityResponseDetailsType{}
	return &this
}

// GetBusinessDate returns the BusinessDate field value if set, zero value otherwise.
func (o *FolioActivityResponseDetailsType) GetBusinessDate() string {
	if o == nil || IsNil(o.BusinessDate) {
		var ret string
		return ret
	}
	return *o.BusinessDate
}

// GetBusinessDateOk returns a tuple with the BusinessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityResponseDetailsType) GetBusinessDateOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessDate) {
		return nil, false
	}
	return o.BusinessDate, true
}

// HasBusinessDate returns a boolean if a field has been set.
func (o *FolioActivityResponseDetailsType) HasBusinessDate() bool {
	if o != nil && !IsNil(o.BusinessDate) {
		return true
	}

	return false
}

// SetBusinessDate gets a reference to the given string and assigns it to the BusinessDate field.
func (o *FolioActivityResponseDetailsType) SetBusinessDate(v string) {
	o.BusinessDate = &v
}

// GetResponseType returns the ResponseType field value if set, zero value otherwise.
func (o *FolioActivityResponseDetailsType) GetResponseType() string {
	if o == nil || IsNil(o.ResponseType) {
		var ret string
		return ret
	}
	return *o.ResponseType
}

// GetResponseTypeOk returns a tuple with the ResponseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityResponseDetailsType) GetResponseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseType) {
		return nil, false
	}
	return o.ResponseType, true
}

// HasResponseType returns a boolean if a field has been set.
func (o *FolioActivityResponseDetailsType) HasResponseType() bool {
	if o != nil && !IsNil(o.ResponseType) {
		return true
	}

	return false
}

// SetResponseType gets a reference to the given string and assigns it to the ResponseType field.
func (o *FolioActivityResponseDetailsType) SetResponseType(v string) {
	o.ResponseType = &v
}

// GetResponseName returns the ResponseName field value if set, zero value otherwise.
func (o *FolioActivityResponseDetailsType) GetResponseName() string {
	if o == nil || IsNil(o.ResponseName) {
		var ret string
		return ret
	}
	return *o.ResponseName
}

// GetResponseNameOk returns a tuple with the ResponseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityResponseDetailsType) GetResponseNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseName) {
		return nil, false
	}
	return o.ResponseName, true
}

// HasResponseName returns a boolean if a field has been set.
func (o *FolioActivityResponseDetailsType) HasResponseName() bool {
	if o != nil && !IsNil(o.ResponseName) {
		return true
	}

	return false
}

// SetResponseName gets a reference to the given string and assigns it to the ResponseName field.
func (o *FolioActivityResponseDetailsType) SetResponseName(v string) {
	o.ResponseName = &v
}

// GetResponseValue returns the ResponseValue field value if set, zero value otherwise.
func (o *FolioActivityResponseDetailsType) GetResponseValue() string {
	if o == nil || IsNil(o.ResponseValue) {
		var ret string
		return ret
	}
	return *o.ResponseValue
}

// GetResponseValueOk returns a tuple with the ResponseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityResponseDetailsType) GetResponseValueOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseValue) {
		return nil, false
	}
	return o.ResponseValue, true
}

// HasResponseValue returns a boolean if a field has been set.
func (o *FolioActivityResponseDetailsType) HasResponseValue() bool {
	if o != nil && !IsNil(o.ResponseValue) {
		return true
	}

	return false
}

// SetResponseValue gets a reference to the given string and assigns it to the ResponseValue field.
func (o *FolioActivityResponseDetailsType) SetResponseValue(v string) {
	o.ResponseValue = &v
}

// GetResponseAttemptNo returns the ResponseAttemptNo field value if set, zero value otherwise.
func (o *FolioActivityResponseDetailsType) GetResponseAttemptNo() int32 {
	if o == nil || IsNil(o.ResponseAttemptNo) {
		var ret int32
		return ret
	}
	return *o.ResponseAttemptNo
}

// GetResponseAttemptNoOk returns a tuple with the ResponseAttemptNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioActivityResponseDetailsType) GetResponseAttemptNoOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponseAttemptNo) {
		return nil, false
	}
	return o.ResponseAttemptNo, true
}

// HasResponseAttemptNo returns a boolean if a field has been set.
func (o *FolioActivityResponseDetailsType) HasResponseAttemptNo() bool {
	if o != nil && !IsNil(o.ResponseAttemptNo) {
		return true
	}

	return false
}

// SetResponseAttemptNo gets a reference to the given int32 and assigns it to the ResponseAttemptNo field.
func (o *FolioActivityResponseDetailsType) SetResponseAttemptNo(v int32) {
	o.ResponseAttemptNo = &v
}

func (o FolioActivityResponseDetailsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolioActivityResponseDetailsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessDate) {
		toSerialize["businessDate"] = o.BusinessDate
	}
	if !IsNil(o.ResponseType) {
		toSerialize["responseType"] = o.ResponseType
	}
	if !IsNil(o.ResponseName) {
		toSerialize["responseName"] = o.ResponseName
	}
	if !IsNil(o.ResponseValue) {
		toSerialize["responseValue"] = o.ResponseValue
	}
	if !IsNil(o.ResponseAttemptNo) {
		toSerialize["responseAttemptNo"] = o.ResponseAttemptNo
	}
	return toSerialize, nil
}

type NullableFolioActivityResponseDetailsType struct {
	value *FolioActivityResponseDetailsType
	isSet bool
}

func (v NullableFolioActivityResponseDetailsType) Get() *FolioActivityResponseDetailsType {
	return v.value
}

func (v *NullableFolioActivityResponseDetailsType) Set(val *FolioActivityResponseDetailsType) {
	v.value = val
	v.isSet = true
}

func (v NullableFolioActivityResponseDetailsType) IsSet() bool {
	return v.isSet
}

func (v *NullableFolioActivityResponseDetailsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolioActivityResponseDetailsType(val *FolioActivityResponseDetailsType) *NullableFolioActivityResponseDetailsType {
	return &NullableFolioActivityResponseDetailsType{value: val, isSet: true}
}

func (v NullableFolioActivityResponseDetailsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolioActivityResponseDetailsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


