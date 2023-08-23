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

// checks if the ChannelRoomMappingMasterInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelRoomMappingMasterInfoType{}

// ChannelRoomMappingMasterInfoType Additional details about the booking channels and source descriptions for a hotel room type referenced within the fetched results.
type ChannelRoomMappingMasterInfoType struct {
	// Additional detail about booking channel.
	BookingChannelsInfo []BookingChannelInfoType `json:"bookingChannelsInfo,omitempty"`
	// This type holds hotel-channel room type mapping source descriptions.
	SourceDescriptions []ChannelRoomMappingSourceDescriptionsType `json:"sourceDescriptions,omitempty"`
}

// NewChannelRoomMappingMasterInfoType instantiates a new ChannelRoomMappingMasterInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelRoomMappingMasterInfoType() *ChannelRoomMappingMasterInfoType {
	this := ChannelRoomMappingMasterInfoType{}
	return &this
}

// NewChannelRoomMappingMasterInfoTypeWithDefaults instantiates a new ChannelRoomMappingMasterInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelRoomMappingMasterInfoTypeWithDefaults() *ChannelRoomMappingMasterInfoType {
	this := ChannelRoomMappingMasterInfoType{}
	return &this
}

// GetBookingChannelsInfo returns the BookingChannelsInfo field value if set, zero value otherwise.
func (o *ChannelRoomMappingMasterInfoType) GetBookingChannelsInfo() []BookingChannelInfoType {
	if o == nil || IsNil(o.BookingChannelsInfo) {
		var ret []BookingChannelInfoType
		return ret
	}
	return o.BookingChannelsInfo
}

// GetBookingChannelsInfoOk returns a tuple with the BookingChannelsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRoomMappingMasterInfoType) GetBookingChannelsInfoOk() ([]BookingChannelInfoType, bool) {
	if o == nil || IsNil(o.BookingChannelsInfo) {
		return nil, false
	}
	return o.BookingChannelsInfo, true
}

// HasBookingChannelsInfo returns a boolean if a field has been set.
func (o *ChannelRoomMappingMasterInfoType) HasBookingChannelsInfo() bool {
	if o != nil && !IsNil(o.BookingChannelsInfo) {
		return true
	}

	return false
}

// SetBookingChannelsInfo gets a reference to the given []BookingChannelInfoType and assigns it to the BookingChannelsInfo field.
func (o *ChannelRoomMappingMasterInfoType) SetBookingChannelsInfo(v []BookingChannelInfoType) {
	o.BookingChannelsInfo = v
}

// GetSourceDescriptions returns the SourceDescriptions field value if set, zero value otherwise.
func (o *ChannelRoomMappingMasterInfoType) GetSourceDescriptions() []ChannelRoomMappingSourceDescriptionsType {
	if o == nil || IsNil(o.SourceDescriptions) {
		var ret []ChannelRoomMappingSourceDescriptionsType
		return ret
	}
	return o.SourceDescriptions
}

// GetSourceDescriptionsOk returns a tuple with the SourceDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelRoomMappingMasterInfoType) GetSourceDescriptionsOk() ([]ChannelRoomMappingSourceDescriptionsType, bool) {
	if o == nil || IsNil(o.SourceDescriptions) {
		return nil, false
	}
	return o.SourceDescriptions, true
}

// HasSourceDescriptions returns a boolean if a field has been set.
func (o *ChannelRoomMappingMasterInfoType) HasSourceDescriptions() bool {
	if o != nil && !IsNil(o.SourceDescriptions) {
		return true
	}

	return false
}

// SetSourceDescriptions gets a reference to the given []ChannelRoomMappingSourceDescriptionsType and assigns it to the SourceDescriptions field.
func (o *ChannelRoomMappingMasterInfoType) SetSourceDescriptions(v []ChannelRoomMappingSourceDescriptionsType) {
	o.SourceDescriptions = v
}

func (o ChannelRoomMappingMasterInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelRoomMappingMasterInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BookingChannelsInfo) {
		toSerialize["bookingChannelsInfo"] = o.BookingChannelsInfo
	}
	if !IsNil(o.SourceDescriptions) {
		toSerialize["sourceDescriptions"] = o.SourceDescriptions
	}
	return toSerialize, nil
}

type NullableChannelRoomMappingMasterInfoType struct {
	value *ChannelRoomMappingMasterInfoType
	isSet bool
}

func (v NullableChannelRoomMappingMasterInfoType) Get() *ChannelRoomMappingMasterInfoType {
	return v.value
}

func (v *NullableChannelRoomMappingMasterInfoType) Set(val *ChannelRoomMappingMasterInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelRoomMappingMasterInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelRoomMappingMasterInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelRoomMappingMasterInfoType(val *ChannelRoomMappingMasterInfoType) *NullableChannelRoomMappingMasterInfoType {
	return &NullableChannelRoomMappingMasterInfoType{value: val, isSet: true}
}

func (v NullableChannelRoomMappingMasterInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelRoomMappingMasterInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


