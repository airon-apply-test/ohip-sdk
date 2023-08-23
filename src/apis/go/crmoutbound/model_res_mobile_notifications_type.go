/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResMobileNotificationsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResMobileNotificationsType{}

// ResMobileNotificationsType Contains the status of Room Ready and Key Ready messages.
type ResMobileNotificationsType struct {
	RoomReady *CommunicationStatusType `json:"roomReady,omitempty"`
	KeyReady *CommunicationStatusType `json:"keyReady,omitempty"`
	// Indicates if the Mobile checkout message is received.
	CheckoutMessageReceived *bool `json:"checkoutMessageReceived,omitempty"`
	// Indicates if user action is required. The action could be for sending Room Ready or Key Ready Notification. It could also be for initiating Checkout
	RequiresAction *bool `json:"requiresAction,omitempty"`
	// Error message when Mobile Notification has failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewResMobileNotificationsType instantiates a new ResMobileNotificationsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResMobileNotificationsType() *ResMobileNotificationsType {
	this := ResMobileNotificationsType{}
	return &this
}

// NewResMobileNotificationsTypeWithDefaults instantiates a new ResMobileNotificationsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResMobileNotificationsTypeWithDefaults() *ResMobileNotificationsType {
	this := ResMobileNotificationsType{}
	return &this
}

// GetRoomReady returns the RoomReady field value if set, zero value otherwise.
func (o *ResMobileNotificationsType) GetRoomReady() CommunicationStatusType {
	if o == nil || IsNil(o.RoomReady) {
		var ret CommunicationStatusType
		return ret
	}
	return *o.RoomReady
}

// GetRoomReadyOk returns a tuple with the RoomReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResMobileNotificationsType) GetRoomReadyOk() (*CommunicationStatusType, bool) {
	if o == nil || IsNil(o.RoomReady) {
		return nil, false
	}
	return o.RoomReady, true
}

// HasRoomReady returns a boolean if a field has been set.
func (o *ResMobileNotificationsType) HasRoomReady() bool {
	if o != nil && !IsNil(o.RoomReady) {
		return true
	}

	return false
}

// SetRoomReady gets a reference to the given CommunicationStatusType and assigns it to the RoomReady field.
func (o *ResMobileNotificationsType) SetRoomReady(v CommunicationStatusType) {
	o.RoomReady = &v
}

// GetKeyReady returns the KeyReady field value if set, zero value otherwise.
func (o *ResMobileNotificationsType) GetKeyReady() CommunicationStatusType {
	if o == nil || IsNil(o.KeyReady) {
		var ret CommunicationStatusType
		return ret
	}
	return *o.KeyReady
}

// GetKeyReadyOk returns a tuple with the KeyReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResMobileNotificationsType) GetKeyReadyOk() (*CommunicationStatusType, bool) {
	if o == nil || IsNil(o.KeyReady) {
		return nil, false
	}
	return o.KeyReady, true
}

// HasKeyReady returns a boolean if a field has been set.
func (o *ResMobileNotificationsType) HasKeyReady() bool {
	if o != nil && !IsNil(o.KeyReady) {
		return true
	}

	return false
}

// SetKeyReady gets a reference to the given CommunicationStatusType and assigns it to the KeyReady field.
func (o *ResMobileNotificationsType) SetKeyReady(v CommunicationStatusType) {
	o.KeyReady = &v
}

// GetCheckoutMessageReceived returns the CheckoutMessageReceived field value if set, zero value otherwise.
func (o *ResMobileNotificationsType) GetCheckoutMessageReceived() bool {
	if o == nil || IsNil(o.CheckoutMessageReceived) {
		var ret bool
		return ret
	}
	return *o.CheckoutMessageReceived
}

// GetCheckoutMessageReceivedOk returns a tuple with the CheckoutMessageReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResMobileNotificationsType) GetCheckoutMessageReceivedOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckoutMessageReceived) {
		return nil, false
	}
	return o.CheckoutMessageReceived, true
}

// HasCheckoutMessageReceived returns a boolean if a field has been set.
func (o *ResMobileNotificationsType) HasCheckoutMessageReceived() bool {
	if o != nil && !IsNil(o.CheckoutMessageReceived) {
		return true
	}

	return false
}

// SetCheckoutMessageReceived gets a reference to the given bool and assigns it to the CheckoutMessageReceived field.
func (o *ResMobileNotificationsType) SetCheckoutMessageReceived(v bool) {
	o.CheckoutMessageReceived = &v
}

// GetRequiresAction returns the RequiresAction field value if set, zero value otherwise.
func (o *ResMobileNotificationsType) GetRequiresAction() bool {
	if o == nil || IsNil(o.RequiresAction) {
		var ret bool
		return ret
	}
	return *o.RequiresAction
}

// GetRequiresActionOk returns a tuple with the RequiresAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResMobileNotificationsType) GetRequiresActionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresAction) {
		return nil, false
	}
	return o.RequiresAction, true
}

// HasRequiresAction returns a boolean if a field has been set.
func (o *ResMobileNotificationsType) HasRequiresAction() bool {
	if o != nil && !IsNil(o.RequiresAction) {
		return true
	}

	return false
}

// SetRequiresAction gets a reference to the given bool and assigns it to the RequiresAction field.
func (o *ResMobileNotificationsType) SetRequiresAction(v bool) {
	o.RequiresAction = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ResMobileNotificationsType) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResMobileNotificationsType) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ResMobileNotificationsType) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ResMobileNotificationsType) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o ResMobileNotificationsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResMobileNotificationsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoomReady) {
		toSerialize["roomReady"] = o.RoomReady
	}
	if !IsNil(o.KeyReady) {
		toSerialize["keyReady"] = o.KeyReady
	}
	if !IsNil(o.CheckoutMessageReceived) {
		toSerialize["checkoutMessageReceived"] = o.CheckoutMessageReceived
	}
	if !IsNil(o.RequiresAction) {
		toSerialize["requiresAction"] = o.RequiresAction
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableResMobileNotificationsType struct {
	value *ResMobileNotificationsType
	isSet bool
}

func (v NullableResMobileNotificationsType) Get() *ResMobileNotificationsType {
	return v.value
}

func (v *NullableResMobileNotificationsType) Set(val *ResMobileNotificationsType) {
	v.value = val
	v.isSet = true
}

func (v NullableResMobileNotificationsType) IsSet() bool {
	return v.isSet
}

func (v *NullableResMobileNotificationsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResMobileNotificationsType(val *ResMobileNotificationsType) *NullableResMobileNotificationsType {
	return &NullableResMobileNotificationsType{value: val, isSet: true}
}

func (v NullableResMobileNotificationsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResMobileNotificationsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


