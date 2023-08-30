/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lms

import (
	"encoding/json"
)

// checks if the PutActivityLocationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutActivityLocationsRequest{}

// PutActivityLocationsRequest struct for PutActivityLocationsRequest
type PutActivityLocationsRequest struct {
	// Collection of Activity Locations.
	ActivityLocations []ActivityLocationType `json:"activityLocations,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutActivityLocationsRequest instantiates a new PutActivityLocationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutActivityLocationsRequest() *PutActivityLocationsRequest {
	this := PutActivityLocationsRequest{}
	return &this
}

// NewPutActivityLocationsRequestWithDefaults instantiates a new PutActivityLocationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutActivityLocationsRequestWithDefaults() *PutActivityLocationsRequest {
	this := PutActivityLocationsRequest{}
	return &this
}

// GetActivityLocations returns the ActivityLocations field value if set, zero value otherwise.
func (o *PutActivityLocationsRequest) GetActivityLocations() []ActivityLocationType {
	if o == nil || IsNil(o.ActivityLocations) {
		var ret []ActivityLocationType
		return ret
	}
	return o.ActivityLocations
}

// GetActivityLocationsOk returns a tuple with the ActivityLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutActivityLocationsRequest) GetActivityLocationsOk() ([]ActivityLocationType, bool) {
	if o == nil || IsNil(o.ActivityLocations) {
		return nil, false
	}
	return o.ActivityLocations, true
}

// HasActivityLocations returns a boolean if a field has been set.
func (o *PutActivityLocationsRequest) HasActivityLocations() bool {
	if o != nil && !IsNil(o.ActivityLocations) {
		return true
	}

	return false
}

// SetActivityLocations gets a reference to the given []ActivityLocationType and assigns it to the ActivityLocations field.
func (o *PutActivityLocationsRequest) SetActivityLocations(v []ActivityLocationType) {
	o.ActivityLocations = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutActivityLocationsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutActivityLocationsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutActivityLocationsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutActivityLocationsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutActivityLocationsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutActivityLocationsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutActivityLocationsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutActivityLocationsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutActivityLocationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutActivityLocationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityLocations) {
		toSerialize["activityLocations"] = o.ActivityLocations
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutActivityLocationsRequest struct {
	value *PutActivityLocationsRequest
	isSet bool
}

func (v NullablePutActivityLocationsRequest) Get() *PutActivityLocationsRequest {
	return v.value
}

func (v *NullablePutActivityLocationsRequest) Set(val *PutActivityLocationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutActivityLocationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutActivityLocationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutActivityLocationsRequest(val *PutActivityLocationsRequest) *NullablePutActivityLocationsRequest {
	return &NullablePutActivityLocationsRequest{value: val, isSet: true}
}

func (v NullablePutActivityLocationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutActivityLocationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


