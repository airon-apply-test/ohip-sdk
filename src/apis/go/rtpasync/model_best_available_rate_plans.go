/*
Opera Cloud Rate Plan Asynchronous Service API

APIs catering to the Rate Plan asynchronous related functionality in a hotel.  This includes adding/updating daily rates&apos; pricing schedules and best available rates by day or length of stay. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtpasync

import (
	"encoding/json"
)

// checks if the BestAvailableRatePlans type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BestAvailableRatePlans{}

// BestAvailableRatePlans Request for configuring best available rate plans.
type BestAvailableRatePlans struct {
	// Collection of best available rate plans.
	BestAvailableRatePlans []BestAvailableRatePlanType `json:"bestAvailableRatePlans,omitempty"`
}

// NewBestAvailableRatePlans instantiates a new BestAvailableRatePlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBestAvailableRatePlans() *BestAvailableRatePlans {
	this := BestAvailableRatePlans{}
	return &this
}

// NewBestAvailableRatePlansWithDefaults instantiates a new BestAvailableRatePlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBestAvailableRatePlansWithDefaults() *BestAvailableRatePlans {
	this := BestAvailableRatePlans{}
	return &this
}

// GetBestAvailableRatePlans returns the BestAvailableRatePlans field value if set, zero value otherwise.
func (o *BestAvailableRatePlans) GetBestAvailableRatePlans() []BestAvailableRatePlanType {
	if o == nil || IsNil(o.BestAvailableRatePlans) {
		var ret []BestAvailableRatePlanType
		return ret
	}
	return o.BestAvailableRatePlans
}

// GetBestAvailableRatePlansOk returns a tuple with the BestAvailableRatePlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlans) GetBestAvailableRatePlansOk() ([]BestAvailableRatePlanType, bool) {
	if o == nil || IsNil(o.BestAvailableRatePlans) {
		return nil, false
	}
	return o.BestAvailableRatePlans, true
}

// HasBestAvailableRatePlans returns a boolean if a field has been set.
func (o *BestAvailableRatePlans) HasBestAvailableRatePlans() bool {
	if o != nil && !IsNil(o.BestAvailableRatePlans) {
		return true
	}

	return false
}

// SetBestAvailableRatePlans gets a reference to the given []BestAvailableRatePlanType and assigns it to the BestAvailableRatePlans field.
func (o *BestAvailableRatePlans) SetBestAvailableRatePlans(v []BestAvailableRatePlanType) {
	o.BestAvailableRatePlans = v
}

func (o BestAvailableRatePlans) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BestAvailableRatePlans) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BestAvailableRatePlans) {
		toSerialize["bestAvailableRatePlans"] = o.BestAvailableRatePlans
	}
	return toSerialize, nil
}

type NullableBestAvailableRatePlans struct {
	value *BestAvailableRatePlans
	isSet bool
}

func (v NullableBestAvailableRatePlans) Get() *BestAvailableRatePlans {
	return v.value
}

func (v *NullableBestAvailableRatePlans) Set(val *BestAvailableRatePlans) {
	v.value = val
	v.isSet = true
}

func (v NullableBestAvailableRatePlans) IsSet() bool {
	return v.isSet
}

func (v *NullableBestAvailableRatePlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBestAvailableRatePlans(val *BestAvailableRatePlans) *NullableBestAvailableRatePlans {
	return &NullableBestAvailableRatePlans{value: val, isSet: true}
}

func (v NullableBestAvailableRatePlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBestAvailableRatePlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


