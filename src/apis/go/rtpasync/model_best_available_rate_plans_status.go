/*
Opera Cloud Rate Plan Asynchronous Service API

APIs catering to the Rate Plan asynchronous related functionality in a hotel.  This includes adding/updating daily rates&apos; pricing schedules and best available rates by day or length of stay. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BestAvailableRatePlansStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BestAvailableRatePlansStatus{}

// BestAvailableRatePlansStatus Response for configured best available rate plans status.
type BestAvailableRatePlansStatus struct {
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
}

// NewBestAvailableRatePlansStatus instantiates a new BestAvailableRatePlansStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBestAvailableRatePlansStatus() *BestAvailableRatePlansStatus {
	this := BestAvailableRatePlansStatus{}
	return &this
}

// NewBestAvailableRatePlansStatusWithDefaults instantiates a new BestAvailableRatePlansStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBestAvailableRatePlansStatusWithDefaults() *BestAvailableRatePlansStatus {
	this := BestAvailableRatePlansStatus{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BestAvailableRatePlansStatus) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlansStatus) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BestAvailableRatePlansStatus) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *BestAvailableRatePlansStatus) SetWarnings(v []WarningType) {
	o.Warnings = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BestAvailableRatePlansStatus) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlansStatus) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BestAvailableRatePlansStatus) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *BestAvailableRatePlansStatus) SetLinks(v []InstanceLink) {
	o.Links = v
}

func (o BestAvailableRatePlansStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BestAvailableRatePlansStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBestAvailableRatePlansStatus struct {
	value *BestAvailableRatePlansStatus
	isSet bool
}

func (v NullableBestAvailableRatePlansStatus) Get() *BestAvailableRatePlansStatus {
	return v.value
}

func (v *NullableBestAvailableRatePlansStatus) Set(val *BestAvailableRatePlansStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBestAvailableRatePlansStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBestAvailableRatePlansStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBestAvailableRatePlansStatus(val *BestAvailableRatePlansStatus) *NullableBestAvailableRatePlansStatus {
	return &NullableBestAvailableRatePlansStatus{value: val, isSet: true}
}

func (v NullableBestAvailableRatePlansStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBestAvailableRatePlansStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


