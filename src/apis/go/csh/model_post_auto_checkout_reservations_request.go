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

// checks if the PostAutoCheckoutReservationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAutoCheckoutReservationsRequest{}

// PostAutoCheckoutReservationsRequest struct for PostAutoCheckoutReservationsRequest
type PostAutoCheckoutReservationsRequest struct {
	Criteria *CheckoutReservationType `json:"criteria,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostAutoCheckoutReservationsRequest instantiates a new PostAutoCheckoutReservationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAutoCheckoutReservationsRequest() *PostAutoCheckoutReservationsRequest {
	this := PostAutoCheckoutReservationsRequest{}
	return &this
}

// NewPostAutoCheckoutReservationsRequestWithDefaults instantiates a new PostAutoCheckoutReservationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAutoCheckoutReservationsRequestWithDefaults() *PostAutoCheckoutReservationsRequest {
	this := PostAutoCheckoutReservationsRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PostAutoCheckoutReservationsRequest) GetCriteria() CheckoutReservationType {
	if o == nil || IsNil(o.Criteria) {
		var ret CheckoutReservationType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAutoCheckoutReservationsRequest) GetCriteriaOk() (*CheckoutReservationType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PostAutoCheckoutReservationsRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given CheckoutReservationType and assigns it to the Criteria field.
func (o *PostAutoCheckoutReservationsRequest) SetCriteria(v CheckoutReservationType) {
	o.Criteria = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostAutoCheckoutReservationsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAutoCheckoutReservationsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostAutoCheckoutReservationsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostAutoCheckoutReservationsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostAutoCheckoutReservationsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAutoCheckoutReservationsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostAutoCheckoutReservationsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostAutoCheckoutReservationsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostAutoCheckoutReservationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAutoCheckoutReservationsRequest) ToMap() (map[string]interface{}, error) {
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

type NullablePostAutoCheckoutReservationsRequest struct {
	value *PostAutoCheckoutReservationsRequest
	isSet bool
}

func (v NullablePostAutoCheckoutReservationsRequest) Get() *PostAutoCheckoutReservationsRequest {
	return v.value
}

func (v *NullablePostAutoCheckoutReservationsRequest) Set(val *PostAutoCheckoutReservationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAutoCheckoutReservationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAutoCheckoutReservationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAutoCheckoutReservationsRequest(val *PostAutoCheckoutReservationsRequest) *NullablePostAutoCheckoutReservationsRequest {
	return &NullablePostAutoCheckoutReservationsRequest{value: val, isSet: true}
}

func (v NullablePostAutoCheckoutReservationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAutoCheckoutReservationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


