/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the ProfileEnrollmentTypeUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileEnrollmentTypeUrls{}

// ProfileEnrollmentTypeUrls List of Information on a URL for the customer.
type ProfileEnrollmentTypeUrls struct {
	// Collection of Detailed information on web url/address for the customer.
	UrlInfo []URLInfoType `json:"urlInfo,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	AllRowsFetched *bool `json:"allRowsFetched,omitempty"`
	// Total number of rows queried
	TotalRows *int32 `json:"totalRows,omitempty"`
}

// NewProfileEnrollmentTypeUrls instantiates a new ProfileEnrollmentTypeUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentTypeUrls() *ProfileEnrollmentTypeUrls {
	this := ProfileEnrollmentTypeUrls{}
	return &this
}

// NewProfileEnrollmentTypeUrlsWithDefaults instantiates a new ProfileEnrollmentTypeUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentTypeUrlsWithDefaults() *ProfileEnrollmentTypeUrls {
	this := ProfileEnrollmentTypeUrls{}
	return &this
}

// GetUrlInfo returns the UrlInfo field value if set, zero value otherwise.
func (o *ProfileEnrollmentTypeUrls) GetUrlInfo() []URLInfoType {
	if o == nil || IsNil(o.UrlInfo) {
		var ret []URLInfoType
		return ret
	}
	return o.UrlInfo
}

// GetUrlInfoOk returns a tuple with the UrlInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentTypeUrls) GetUrlInfoOk() ([]URLInfoType, bool) {
	if o == nil || IsNil(o.UrlInfo) {
		return nil, false
	}
	return o.UrlInfo, true
}

// HasUrlInfo returns a boolean if a field has been set.
func (o *ProfileEnrollmentTypeUrls) HasUrlInfo() bool {
	if o != nil && !IsNil(o.UrlInfo) {
		return true
	}

	return false
}

// SetUrlInfo gets a reference to the given []URLInfoType and assigns it to the UrlInfo field.
func (o *ProfileEnrollmentTypeUrls) SetUrlInfo(v []URLInfoType) {
	o.UrlInfo = v
}

// GetAllRowsFetched returns the AllRowsFetched field value if set, zero value otherwise.
func (o *ProfileEnrollmentTypeUrls) GetAllRowsFetched() bool {
	if o == nil || IsNil(o.AllRowsFetched) {
		var ret bool
		return ret
	}
	return *o.AllRowsFetched
}

// GetAllRowsFetchedOk returns a tuple with the AllRowsFetched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentTypeUrls) GetAllRowsFetchedOk() (*bool, bool) {
	if o == nil || IsNil(o.AllRowsFetched) {
		return nil, false
	}
	return o.AllRowsFetched, true
}

// HasAllRowsFetched returns a boolean if a field has been set.
func (o *ProfileEnrollmentTypeUrls) HasAllRowsFetched() bool {
	if o != nil && !IsNil(o.AllRowsFetched) {
		return true
	}

	return false
}

// SetAllRowsFetched gets a reference to the given bool and assigns it to the AllRowsFetched field.
func (o *ProfileEnrollmentTypeUrls) SetAllRowsFetched(v bool) {
	o.AllRowsFetched = &v
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise.
func (o *ProfileEnrollmentTypeUrls) GetTotalRows() int32 {
	if o == nil || IsNil(o.TotalRows) {
		var ret int32
		return ret
	}
	return *o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentTypeUrls) GetTotalRowsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRows) {
		return nil, false
	}
	return o.TotalRows, true
}

// HasTotalRows returns a boolean if a field has been set.
func (o *ProfileEnrollmentTypeUrls) HasTotalRows() bool {
	if o != nil && !IsNil(o.TotalRows) {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given int32 and assigns it to the TotalRows field.
func (o *ProfileEnrollmentTypeUrls) SetTotalRows(v int32) {
	o.TotalRows = &v
}

func (o ProfileEnrollmentTypeUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileEnrollmentTypeUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UrlInfo) {
		toSerialize["urlInfo"] = o.UrlInfo
	}
	if !IsNil(o.AllRowsFetched) {
		toSerialize["allRowsFetched"] = o.AllRowsFetched
	}
	if !IsNil(o.TotalRows) {
		toSerialize["totalRows"] = o.TotalRows
	}
	return toSerialize, nil
}

type NullableProfileEnrollmentTypeUrls struct {
	value *ProfileEnrollmentTypeUrls
	isSet bool
}

func (v NullableProfileEnrollmentTypeUrls) Get() *ProfileEnrollmentTypeUrls {
	return v.value
}

func (v *NullableProfileEnrollmentTypeUrls) Set(val *ProfileEnrollmentTypeUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileEnrollmentTypeUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileEnrollmentTypeUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileEnrollmentTypeUrls(val *ProfileEnrollmentTypeUrls) *NullableProfileEnrollmentTypeUrls {
	return &NullableProfileEnrollmentTypeUrls{value: val, isSet: true}
}

func (v NullableProfileEnrollmentTypeUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileEnrollmentTypeUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


