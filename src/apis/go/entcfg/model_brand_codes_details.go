/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BrandCodesDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandCodesDetails{}

// BrandCodesDetails Response object for fetching Brand Codes.
type BrandCodesDetails struct {
	// List of Brand Codes.
	BrandCodes []BrandCodeType `json:"brandCodes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewBrandCodesDetails instantiates a new BrandCodesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandCodesDetails() *BrandCodesDetails {
	this := BrandCodesDetails{}
	return &this
}

// NewBrandCodesDetailsWithDefaults instantiates a new BrandCodesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandCodesDetailsWithDefaults() *BrandCodesDetails {
	this := BrandCodesDetails{}
	return &this
}

// GetBrandCodes returns the BrandCodes field value if set, zero value otherwise.
func (o *BrandCodesDetails) GetBrandCodes() []BrandCodeType {
	if o == nil || IsNil(o.BrandCodes) {
		var ret []BrandCodeType
		return ret
	}
	return o.BrandCodes
}

// GetBrandCodesOk returns a tuple with the BrandCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandCodesDetails) GetBrandCodesOk() ([]BrandCodeType, bool) {
	if o == nil || IsNil(o.BrandCodes) {
		return nil, false
	}
	return o.BrandCodes, true
}

// HasBrandCodes returns a boolean if a field has been set.
func (o *BrandCodesDetails) HasBrandCodes() bool {
	if o != nil && !IsNil(o.BrandCodes) {
		return true
	}

	return false
}

// SetBrandCodes gets a reference to the given []BrandCodeType and assigns it to the BrandCodes field.
func (o *BrandCodesDetails) SetBrandCodes(v []BrandCodeType) {
	o.BrandCodes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BrandCodesDetails) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandCodesDetails) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BrandCodesDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BrandCodesDetails) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BrandCodesDetails) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandCodesDetails) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BrandCodesDetails) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BrandCodesDetails) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o BrandCodesDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandCodesDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandCodes) {
		toSerialize["brandCodes"] = o.BrandCodes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBrandCodesDetails struct {
	value *BrandCodesDetails
	isSet bool
}

func (v NullableBrandCodesDetails) Get() *BrandCodesDetails {
	return v.value
}

func (v *NullableBrandCodesDetails) Set(val *BrandCodesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandCodesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandCodesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandCodesDetails(val *BrandCodesDetails) *NullableBrandCodesDetails {
	return &NullableBrandCodesDetails{value: val, isSet: true}
}

func (v NullableBrandCodesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandCodesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


