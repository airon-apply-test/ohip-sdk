/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutInfluenceCodesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutInfluenceCodesRequest{}

// PutInfluenceCodesRequest struct for PutInfluenceCodesRequest
type PutInfluenceCodesRequest struct {
	// List of Influence Codes.
	InfluenceCodes []InfluenceCodeType `json:"influenceCodes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutInfluenceCodesRequest instantiates a new PutInfluenceCodesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutInfluenceCodesRequest() *PutInfluenceCodesRequest {
	this := PutInfluenceCodesRequest{}
	return &this
}

// NewPutInfluenceCodesRequestWithDefaults instantiates a new PutInfluenceCodesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutInfluenceCodesRequestWithDefaults() *PutInfluenceCodesRequest {
	this := PutInfluenceCodesRequest{}
	return &this
}

// GetInfluenceCodes returns the InfluenceCodes field value if set, zero value otherwise.
func (o *PutInfluenceCodesRequest) GetInfluenceCodes() []InfluenceCodeType {
	if o == nil || IsNil(o.InfluenceCodes) {
		var ret []InfluenceCodeType
		return ret
	}
	return o.InfluenceCodes
}

// GetInfluenceCodesOk returns a tuple with the InfluenceCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutInfluenceCodesRequest) GetInfluenceCodesOk() ([]InfluenceCodeType, bool) {
	if o == nil || IsNil(o.InfluenceCodes) {
		return nil, false
	}
	return o.InfluenceCodes, true
}

// HasInfluenceCodes returns a boolean if a field has been set.
func (o *PutInfluenceCodesRequest) HasInfluenceCodes() bool {
	if o != nil && !IsNil(o.InfluenceCodes) {
		return true
	}

	return false
}

// SetInfluenceCodes gets a reference to the given []InfluenceCodeType and assigns it to the InfluenceCodes field.
func (o *PutInfluenceCodesRequest) SetInfluenceCodes(v []InfluenceCodeType) {
	o.InfluenceCodes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutInfluenceCodesRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutInfluenceCodesRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutInfluenceCodesRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutInfluenceCodesRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutInfluenceCodesRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutInfluenceCodesRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutInfluenceCodesRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutInfluenceCodesRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutInfluenceCodesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutInfluenceCodesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InfluenceCodes) {
		toSerialize["influenceCodes"] = o.InfluenceCodes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutInfluenceCodesRequest struct {
	value *PutInfluenceCodesRequest
	isSet bool
}

func (v NullablePutInfluenceCodesRequest) Get() *PutInfluenceCodesRequest {
	return v.value
}

func (v *NullablePutInfluenceCodesRequest) Set(val *PutInfluenceCodesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutInfluenceCodesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutInfluenceCodesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutInfluenceCodesRequest(val *PutInfluenceCodesRequest) *NullablePutInfluenceCodesRequest {
	return &NullablePutInfluenceCodesRequest{value: val, isSet: true}
}

func (v NullablePutInfluenceCodesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutInfluenceCodesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


