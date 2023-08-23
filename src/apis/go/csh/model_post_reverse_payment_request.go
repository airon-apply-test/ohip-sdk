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

// checks if the PostReversePaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostReversePaymentRequest{}

// PostReversePaymentRequest struct for PostReversePaymentRequest
type PostReversePaymentRequest struct {
	Payment *PaymentReversalType `json:"payment,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostReversePaymentRequest instantiates a new PostReversePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostReversePaymentRequest() *PostReversePaymentRequest {
	this := PostReversePaymentRequest{}
	return &this
}

// NewPostReversePaymentRequestWithDefaults instantiates a new PostReversePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostReversePaymentRequestWithDefaults() *PostReversePaymentRequest {
	this := PostReversePaymentRequest{}
	return &this
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *PostReversePaymentRequest) GetPayment() PaymentReversalType {
	if o == nil || IsNil(o.Payment) {
		var ret PaymentReversalType
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostReversePaymentRequest) GetPaymentOk() (*PaymentReversalType, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *PostReversePaymentRequest) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given PaymentReversalType and assigns it to the Payment field.
func (o *PostReversePaymentRequest) SetPayment(v PaymentReversalType) {
	o.Payment = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostReversePaymentRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostReversePaymentRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostReversePaymentRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostReversePaymentRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostReversePaymentRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostReversePaymentRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostReversePaymentRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostReversePaymentRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostReversePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostReversePaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostReversePaymentRequest struct {
	value *PostReversePaymentRequest
	isSet bool
}

func (v NullablePostReversePaymentRequest) Get() *PostReversePaymentRequest {
	return v.value
}

func (v *NullablePostReversePaymentRequest) Set(val *PostReversePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostReversePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostReversePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostReversePaymentRequest(val *PostReversePaymentRequest) *NullablePostReversePaymentRequest {
	return &NullablePostReversePaymentRequest{value: val, isSet: true}
}

func (v NullablePostReversePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostReversePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


