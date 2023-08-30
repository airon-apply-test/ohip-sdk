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

// checks if the PostGuestCheckDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGuestCheckDetailsRequest{}

// PostGuestCheckDetailsRequest struct for PostGuestCheckDetailsRequest
type PostGuestCheckDetailsRequest struct {
	ReservationId *ReservationId `json:"reservationId,omitempty"`
	CheckDetails *CheckDetailsType `json:"checkDetails,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostGuestCheckDetailsRequest instantiates a new PostGuestCheckDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGuestCheckDetailsRequest() *PostGuestCheckDetailsRequest {
	this := PostGuestCheckDetailsRequest{}
	return &this
}

// NewPostGuestCheckDetailsRequestWithDefaults instantiates a new PostGuestCheckDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGuestCheckDetailsRequestWithDefaults() *PostGuestCheckDetailsRequest {
	this := PostGuestCheckDetailsRequest{}
	return &this
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *PostGuestCheckDetailsRequest) GetReservationId() ReservationId {
	if o == nil || IsNil(o.ReservationId) {
		var ret ReservationId
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGuestCheckDetailsRequest) GetReservationIdOk() (*ReservationId, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *PostGuestCheckDetailsRequest) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given ReservationId and assigns it to the ReservationId field.
func (o *PostGuestCheckDetailsRequest) SetReservationId(v ReservationId) {
	o.ReservationId = &v
}

// GetCheckDetails returns the CheckDetails field value if set, zero value otherwise.
func (o *PostGuestCheckDetailsRequest) GetCheckDetails() CheckDetailsType {
	if o == nil || IsNil(o.CheckDetails) {
		var ret CheckDetailsType
		return ret
	}
	return *o.CheckDetails
}

// GetCheckDetailsOk returns a tuple with the CheckDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGuestCheckDetailsRequest) GetCheckDetailsOk() (*CheckDetailsType, bool) {
	if o == nil || IsNil(o.CheckDetails) {
		return nil, false
	}
	return o.CheckDetails, true
}

// HasCheckDetails returns a boolean if a field has been set.
func (o *PostGuestCheckDetailsRequest) HasCheckDetails() bool {
	if o != nil && !IsNil(o.CheckDetails) {
		return true
	}

	return false
}

// SetCheckDetails gets a reference to the given CheckDetailsType and assigns it to the CheckDetails field.
func (o *PostGuestCheckDetailsRequest) SetCheckDetails(v CheckDetailsType) {
	o.CheckDetails = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostGuestCheckDetailsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGuestCheckDetailsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostGuestCheckDetailsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostGuestCheckDetailsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostGuestCheckDetailsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGuestCheckDetailsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostGuestCheckDetailsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostGuestCheckDetailsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostGuestCheckDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGuestCheckDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.CheckDetails) {
		toSerialize["checkDetails"] = o.CheckDetails
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostGuestCheckDetailsRequest struct {
	value *PostGuestCheckDetailsRequest
	isSet bool
}

func (v NullablePostGuestCheckDetailsRequest) Get() *PostGuestCheckDetailsRequest {
	return v.value
}

func (v *NullablePostGuestCheckDetailsRequest) Set(val *PostGuestCheckDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGuestCheckDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGuestCheckDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGuestCheckDetailsRequest(val *PostGuestCheckDetailsRequest) *NullablePostGuestCheckDetailsRequest {
	return &NullablePostGuestCheckDetailsRequest{value: val, isSet: true}
}

func (v NullablePostGuestCheckDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGuestCheckDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


