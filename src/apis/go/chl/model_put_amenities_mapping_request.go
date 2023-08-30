/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the PutAmenitiesMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutAmenitiesMappingRequest{}

// PutAmenitiesMappingRequest struct for PutAmenitiesMappingRequest
type PutAmenitiesMappingRequest struct {
	// Information about an external system amenity mapping.
	Amenities []AmenityMappingType `json:"amenities,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutAmenitiesMappingRequest instantiates a new PutAmenitiesMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutAmenitiesMappingRequest() *PutAmenitiesMappingRequest {
	this := PutAmenitiesMappingRequest{}
	return &this
}

// NewPutAmenitiesMappingRequestWithDefaults instantiates a new PutAmenitiesMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutAmenitiesMappingRequestWithDefaults() *PutAmenitiesMappingRequest {
	this := PutAmenitiesMappingRequest{}
	return &this
}

// GetAmenities returns the Amenities field value if set, zero value otherwise.
func (o *PutAmenitiesMappingRequest) GetAmenities() []AmenityMappingType {
	if o == nil || IsNil(o.Amenities) {
		var ret []AmenityMappingType
		return ret
	}
	return o.Amenities
}

// GetAmenitiesOk returns a tuple with the Amenities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAmenitiesMappingRequest) GetAmenitiesOk() ([]AmenityMappingType, bool) {
	if o == nil || IsNil(o.Amenities) {
		return nil, false
	}
	return o.Amenities, true
}

// HasAmenities returns a boolean if a field has been set.
func (o *PutAmenitiesMappingRequest) HasAmenities() bool {
	if o != nil && !IsNil(o.Amenities) {
		return true
	}

	return false
}

// SetAmenities gets a reference to the given []AmenityMappingType and assigns it to the Amenities field.
func (o *PutAmenitiesMappingRequest) SetAmenities(v []AmenityMappingType) {
	o.Amenities = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutAmenitiesMappingRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAmenitiesMappingRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutAmenitiesMappingRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutAmenitiesMappingRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutAmenitiesMappingRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAmenitiesMappingRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutAmenitiesMappingRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutAmenitiesMappingRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutAmenitiesMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutAmenitiesMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amenities) {
		toSerialize["amenities"] = o.Amenities
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutAmenitiesMappingRequest struct {
	value *PutAmenitiesMappingRequest
	isSet bool
}

func (v NullablePutAmenitiesMappingRequest) Get() *PutAmenitiesMappingRequest {
	return v.value
}

func (v *NullablePutAmenitiesMappingRequest) Set(val *PutAmenitiesMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutAmenitiesMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutAmenitiesMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutAmenitiesMappingRequest(val *PutAmenitiesMappingRequest) *NullablePutAmenitiesMappingRequest {
	return &NullablePutAmenitiesMappingRequest{value: val, isSet: true}
}

func (v NullablePutAmenitiesMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutAmenitiesMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


