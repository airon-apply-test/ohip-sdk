/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
	"time"
)

// checks if the MessageStatisticsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageStatisticsType{}

// MessageStatisticsType Type for Message statistics Details.
type MessageStatisticsType struct {
	// InterfaceId of the Messages.
	InterfaceId *string `json:"interfaceId,omitempty"`
	// Property of the Messages.
	HotelId *string `json:"hotelId,omitempty"`
	// Module name of messages.
	Module *string `json:"module,omitempty"`
	// Action Type of the messages.
	ActionType *string `json:"actionType,omitempty"`
	MessageStatus *OXIMessageStatusType `json:"messageStatus,omitempty"`
	// Number of the messages
	MessageCount *int32 `json:"messageCount,omitempty"`
	// Last date on which messages are processed for given criteria.
	LastProcessedDate *time.Time `json:"lastProcessedDate,omitempty"`
}

// NewMessageStatisticsType instantiates a new MessageStatisticsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageStatisticsType() *MessageStatisticsType {
	this := MessageStatisticsType{}
	return &this
}

// NewMessageStatisticsTypeWithDefaults instantiates a new MessageStatisticsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageStatisticsTypeWithDefaults() *MessageStatisticsType {
	this := MessageStatisticsType{}
	return &this
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *MessageStatisticsType) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatisticsType) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *MessageStatisticsType) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *MessageStatisticsType) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *MessageStatisticsType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatisticsType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *MessageStatisticsType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *MessageStatisticsType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *MessageStatisticsType) GetModule() string {
	if o == nil || IsNil(o.Module) {
		var ret string
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatisticsType) GetModuleOk() (*string, bool) {
	if o == nil || IsNil(o.Module) {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *MessageStatisticsType) HasModule() bool {
	if o != nil && !IsNil(o.Module) {
		return true
	}

	return false
}

// SetModule gets a reference to the given string and assigns it to the Module field.
func (o *MessageStatisticsType) SetModule(v string) {
	o.Module = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *MessageStatisticsType) GetActionType() string {
	if o == nil || IsNil(o.ActionType) {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatisticsType) GetActionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *MessageStatisticsType) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *MessageStatisticsType) SetActionType(v string) {
	o.ActionType = &v
}

// GetMessageStatus returns the MessageStatus field value if set, zero value otherwise.
func (o *MessageStatisticsType) GetMessageStatus() OXIMessageStatusType {
	if o == nil || IsNil(o.MessageStatus) {
		var ret OXIMessageStatusType
		return ret
	}
	return *o.MessageStatus
}

// GetMessageStatusOk returns a tuple with the MessageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatisticsType) GetMessageStatusOk() (*OXIMessageStatusType, bool) {
	if o == nil || IsNil(o.MessageStatus) {
		return nil, false
	}
	return o.MessageStatus, true
}

// HasMessageStatus returns a boolean if a field has been set.
func (o *MessageStatisticsType) HasMessageStatus() bool {
	if o != nil && !IsNil(o.MessageStatus) {
		return true
	}

	return false
}

// SetMessageStatus gets a reference to the given OXIMessageStatusType and assigns it to the MessageStatus field.
func (o *MessageStatisticsType) SetMessageStatus(v OXIMessageStatusType) {
	o.MessageStatus = &v
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *MessageStatisticsType) GetMessageCount() int32 {
	if o == nil || IsNil(o.MessageCount) {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatisticsType) GetMessageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageCount) {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *MessageStatisticsType) HasMessageCount() bool {
	if o != nil && !IsNil(o.MessageCount) {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *MessageStatisticsType) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetLastProcessedDate returns the LastProcessedDate field value if set, zero value otherwise.
func (o *MessageStatisticsType) GetLastProcessedDate() time.Time {
	if o == nil || IsNil(o.LastProcessedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastProcessedDate
}

// GetLastProcessedDateOk returns a tuple with the LastProcessedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatisticsType) GetLastProcessedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastProcessedDate) {
		return nil, false
	}
	return o.LastProcessedDate, true
}

// HasLastProcessedDate returns a boolean if a field has been set.
func (o *MessageStatisticsType) HasLastProcessedDate() bool {
	if o != nil && !IsNil(o.LastProcessedDate) {
		return true
	}

	return false
}

// SetLastProcessedDate gets a reference to the given time.Time and assigns it to the LastProcessedDate field.
func (o *MessageStatisticsType) SetLastProcessedDate(v time.Time) {
	o.LastProcessedDate = &v
}

func (o MessageStatisticsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageStatisticsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Module) {
		toSerialize["module"] = o.Module
	}
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !IsNil(o.MessageStatus) {
		toSerialize["messageStatus"] = o.MessageStatus
	}
	if !IsNil(o.MessageCount) {
		toSerialize["messageCount"] = o.MessageCount
	}
	if !IsNil(o.LastProcessedDate) {
		toSerialize["lastProcessedDate"] = o.LastProcessedDate
	}
	return toSerialize, nil
}

type NullableMessageStatisticsType struct {
	value *MessageStatisticsType
	isSet bool
}

func (v NullableMessageStatisticsType) Get() *MessageStatisticsType {
	return v.value
}

func (v *NullableMessageStatisticsType) Set(val *MessageStatisticsType) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageStatisticsType) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageStatisticsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageStatisticsType(val *MessageStatisticsType) *NullableMessageStatisticsType {
	return &NullableMessageStatisticsType{value: val, isSet: true}
}

func (v NullableMessageStatisticsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageStatisticsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


