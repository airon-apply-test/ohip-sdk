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

// checks if the StartSetRatePlansHeadersProcessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartSetRatePlansHeadersProcessRequest{}

// StartSetRatePlansHeadersProcessRequest struct for StartSetRatePlansHeadersProcessRequest
type StartSetRatePlansHeadersProcessRequest struct {
	// Rate plan code details to be created.
	RatePlans []RatePlanType `json:"ratePlans,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewStartSetRatePlansHeadersProcessRequest instantiates a new StartSetRatePlansHeadersProcessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartSetRatePlansHeadersProcessRequest() *StartSetRatePlansHeadersProcessRequest {
	this := StartSetRatePlansHeadersProcessRequest{}
	return &this
}

// NewStartSetRatePlansHeadersProcessRequestWithDefaults instantiates a new StartSetRatePlansHeadersProcessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartSetRatePlansHeadersProcessRequestWithDefaults() *StartSetRatePlansHeadersProcessRequest {
	this := StartSetRatePlansHeadersProcessRequest{}
	return &this
}

// GetRatePlans returns the RatePlans field value if set, zero value otherwise.
func (o *StartSetRatePlansHeadersProcessRequest) GetRatePlans() []RatePlanType {
	if o == nil || IsNil(o.RatePlans) {
		var ret []RatePlanType
		return ret
	}
	return o.RatePlans
}

// GetRatePlansOk returns a tuple with the RatePlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartSetRatePlansHeadersProcessRequest) GetRatePlansOk() ([]RatePlanType, bool) {
	if o == nil || IsNil(o.RatePlans) {
		return nil, false
	}
	return o.RatePlans, true
}

// HasRatePlans returns a boolean if a field has been set.
func (o *StartSetRatePlansHeadersProcessRequest) HasRatePlans() bool {
	if o != nil && !IsNil(o.RatePlans) {
		return true
	}

	return false
}

// SetRatePlans gets a reference to the given []RatePlanType and assigns it to the RatePlans field.
func (o *StartSetRatePlansHeadersProcessRequest) SetRatePlans(v []RatePlanType) {
	o.RatePlans = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *StartSetRatePlansHeadersProcessRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartSetRatePlansHeadersProcessRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StartSetRatePlansHeadersProcessRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *StartSetRatePlansHeadersProcessRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *StartSetRatePlansHeadersProcessRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartSetRatePlansHeadersProcessRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *StartSetRatePlansHeadersProcessRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *StartSetRatePlansHeadersProcessRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o StartSetRatePlansHeadersProcessRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartSetRatePlansHeadersProcessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RatePlans) {
		toSerialize["ratePlans"] = o.RatePlans
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableStartSetRatePlansHeadersProcessRequest struct {
	value *StartSetRatePlansHeadersProcessRequest
	isSet bool
}

func (v NullableStartSetRatePlansHeadersProcessRequest) Get() *StartSetRatePlansHeadersProcessRequest {
	return v.value
}

func (v *NullableStartSetRatePlansHeadersProcessRequest) Set(val *StartSetRatePlansHeadersProcessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartSetRatePlansHeadersProcessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartSetRatePlansHeadersProcessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartSetRatePlansHeadersProcessRequest(val *StartSetRatePlansHeadersProcessRequest) *NullableStartSetRatePlansHeadersProcessRequest {
	return &NullableStartSetRatePlansHeadersProcessRequest{value: val, isSet: true}
}

func (v NullableStartSetRatePlansHeadersProcessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartSetRatePlansHeadersProcessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


