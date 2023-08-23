/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelHotelCodeMappingBucketsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelHotelCodeMappingBucketsType{}

// ChannelHotelCodeMappingBucketsType Channel-hotel mapping buckets information.
type ChannelHotelCodeMappingBucketsType struct {
	// Default rate code revenue thresholds.
	ChannelHotelCodeMappingBucket []ChannelHotelCodeMappingBucketType `json:"channelHotelCodeMappingBucket,omitempty"`
	// Default rate code to be used to calculate the total revenue.
	DefaultRateCode *string `json:"defaultRateCode,omitempty"`
	// Flag indicating if channel resort mapping is available.
	Available *bool `json:"available,omitempty"`
}

// NewChannelHotelCodeMappingBucketsType instantiates a new ChannelHotelCodeMappingBucketsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelHotelCodeMappingBucketsType() *ChannelHotelCodeMappingBucketsType {
	this := ChannelHotelCodeMappingBucketsType{}
	return &this
}

// NewChannelHotelCodeMappingBucketsTypeWithDefaults instantiates a new ChannelHotelCodeMappingBucketsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelHotelCodeMappingBucketsTypeWithDefaults() *ChannelHotelCodeMappingBucketsType {
	this := ChannelHotelCodeMappingBucketsType{}
	return &this
}

// GetChannelHotelCodeMappingBucket returns the ChannelHotelCodeMappingBucket field value if set, zero value otherwise.
func (o *ChannelHotelCodeMappingBucketsType) GetChannelHotelCodeMappingBucket() []ChannelHotelCodeMappingBucketType {
	if o == nil || IsNil(o.ChannelHotelCodeMappingBucket) {
		var ret []ChannelHotelCodeMappingBucketType
		return ret
	}
	return o.ChannelHotelCodeMappingBucket
}

// GetChannelHotelCodeMappingBucketOk returns a tuple with the ChannelHotelCodeMappingBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelHotelCodeMappingBucketsType) GetChannelHotelCodeMappingBucketOk() ([]ChannelHotelCodeMappingBucketType, bool) {
	if o == nil || IsNil(o.ChannelHotelCodeMappingBucket) {
		return nil, false
	}
	return o.ChannelHotelCodeMappingBucket, true
}

// HasChannelHotelCodeMappingBucket returns a boolean if a field has been set.
func (o *ChannelHotelCodeMappingBucketsType) HasChannelHotelCodeMappingBucket() bool {
	if o != nil && !IsNil(o.ChannelHotelCodeMappingBucket) {
		return true
	}

	return false
}

// SetChannelHotelCodeMappingBucket gets a reference to the given []ChannelHotelCodeMappingBucketType and assigns it to the ChannelHotelCodeMappingBucket field.
func (o *ChannelHotelCodeMappingBucketsType) SetChannelHotelCodeMappingBucket(v []ChannelHotelCodeMappingBucketType) {
	o.ChannelHotelCodeMappingBucket = v
}

// GetDefaultRateCode returns the DefaultRateCode field value if set, zero value otherwise.
func (o *ChannelHotelCodeMappingBucketsType) GetDefaultRateCode() string {
	if o == nil || IsNil(o.DefaultRateCode) {
		var ret string
		return ret
	}
	return *o.DefaultRateCode
}

// GetDefaultRateCodeOk returns a tuple with the DefaultRateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelHotelCodeMappingBucketsType) GetDefaultRateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRateCode) {
		return nil, false
	}
	return o.DefaultRateCode, true
}

// HasDefaultRateCode returns a boolean if a field has been set.
func (o *ChannelHotelCodeMappingBucketsType) HasDefaultRateCode() bool {
	if o != nil && !IsNil(o.DefaultRateCode) {
		return true
	}

	return false
}

// SetDefaultRateCode gets a reference to the given string and assigns it to the DefaultRateCode field.
func (o *ChannelHotelCodeMappingBucketsType) SetDefaultRateCode(v string) {
	o.DefaultRateCode = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *ChannelHotelCodeMappingBucketsType) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelHotelCodeMappingBucketsType) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *ChannelHotelCodeMappingBucketsType) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *ChannelHotelCodeMappingBucketsType) SetAvailable(v bool) {
	o.Available = &v
}

func (o ChannelHotelCodeMappingBucketsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelHotelCodeMappingBucketsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelHotelCodeMappingBucket) {
		toSerialize["channelHotelCodeMappingBucket"] = o.ChannelHotelCodeMappingBucket
	}
	if !IsNil(o.DefaultRateCode) {
		toSerialize["defaultRateCode"] = o.DefaultRateCode
	}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	return toSerialize, nil
}

type NullableChannelHotelCodeMappingBucketsType struct {
	value *ChannelHotelCodeMappingBucketsType
	isSet bool
}

func (v NullableChannelHotelCodeMappingBucketsType) Get() *ChannelHotelCodeMappingBucketsType {
	return v.value
}

func (v *NullableChannelHotelCodeMappingBucketsType) Set(val *ChannelHotelCodeMappingBucketsType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelHotelCodeMappingBucketsType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelHotelCodeMappingBucketsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelHotelCodeMappingBucketsType(val *ChannelHotelCodeMappingBucketsType) *NullableChannelHotelCodeMappingBucketsType {
	return &NullableChannelHotelCodeMappingBucketsType{value: val, isSet: true}
}

func (v NullableChannelHotelCodeMappingBucketsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelHotelCodeMappingBucketsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


