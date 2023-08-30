/*
OPERA Cloud API for Customer Management Service

This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cms

import (
	"encoding/json"
)

// checks if the PutTrackItItemsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutTrackItItemsRequest{}

// PutTrackItItemsRequest struct for PutTrackItItemsRequest
type PutTrackItItemsRequest struct {
	// 
	TrackItItemsInfo []TrackItItemType `json:"trackItItemsInfo,omitempty"`
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	HotelId *string `json:"hotelId,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutTrackItItemsRequest instantiates a new PutTrackItItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutTrackItItemsRequest() *PutTrackItItemsRequest {
	this := PutTrackItItemsRequest{}
	return &this
}

// NewPutTrackItItemsRequestWithDefaults instantiates a new PutTrackItItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutTrackItItemsRequestWithDefaults() *PutTrackItItemsRequest {
	this := PutTrackItItemsRequest{}
	return &this
}

// GetTrackItItemsInfo returns the TrackItItemsInfo field value if set, zero value otherwise.
func (o *PutTrackItItemsRequest) GetTrackItItemsInfo() []TrackItItemType {
	if o == nil || IsNil(o.TrackItItemsInfo) {
		var ret []TrackItItemType
		return ret
	}
	return o.TrackItItemsInfo
}

// GetTrackItItemsInfoOk returns a tuple with the TrackItItemsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTrackItItemsRequest) GetTrackItItemsInfoOk() ([]TrackItItemType, bool) {
	if o == nil || IsNil(o.TrackItItemsInfo) {
		return nil, false
	}
	return o.TrackItItemsInfo, true
}

// HasTrackItItemsInfo returns a boolean if a field has been set.
func (o *PutTrackItItemsRequest) HasTrackItItemsInfo() bool {
	if o != nil && !IsNil(o.TrackItItemsInfo) {
		return true
	}

	return false
}

// SetTrackItItemsInfo gets a reference to the given []TrackItItemType and assigns it to the TrackItItemsInfo field.
func (o *PutTrackItItemsRequest) SetTrackItItemsInfo(v []TrackItItemType) {
	o.TrackItItemsInfo = v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *PutTrackItItemsRequest) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTrackItItemsRequest) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *PutTrackItItemsRequest) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *PutTrackItItemsRequest) SetHotelId(v string) {
	o.HotelId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutTrackItItemsRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTrackItItemsRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutTrackItItemsRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutTrackItItemsRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutTrackItItemsRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTrackItItemsRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutTrackItItemsRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutTrackItItemsRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutTrackItItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutTrackItItemsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackItItemsInfo) {
		toSerialize["trackItItemsInfo"] = o.TrackItItemsInfo
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutTrackItItemsRequest struct {
	value *PutTrackItItemsRequest
	isSet bool
}

func (v NullablePutTrackItItemsRequest) Get() *PutTrackItItemsRequest {
	return v.value
}

func (v *NullablePutTrackItItemsRequest) Set(val *PutTrackItItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutTrackItItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutTrackItItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutTrackItItemsRequest(val *PutTrackItItemsRequest) *NullablePutTrackItItemsRequest {
	return &NullablePutTrackItItemsRequest{value: val, isSet: true}
}

func (v NullablePutTrackItItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutTrackItItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


