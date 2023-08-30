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

// checks if the StagedProfileBillingInstructionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StagedProfileBillingInstructionType{}

// StagedProfileBillingInstructionType Configured Billing Instruction which represents a set of Transaction Codes.
type StagedProfileBillingInstructionType struct {
	// Billing Instruction code description.
	Desc *string `json:"desc,omitempty"`
	// This is the Routing Instruction Id attached with Reservation. It is only used for internal purpose. It should not be used by external vendor or consumer.
	RoutingInstructionsId *float32 `json:"routingInstructionsId,omitempty"`
	// Unique identifier for the Billing Instruction.
	BillingCode *string `json:"billingCode,omitempty"`
	// Hotel context of the Billing Instruction.
	HotelId *string `json:"hotelId,omitempty"`
	// The error in billing instruction information in a staged profile with an invalid status
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
}

// NewStagedProfileBillingInstructionType instantiates a new StagedProfileBillingInstructionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagedProfileBillingInstructionType() *StagedProfileBillingInstructionType {
	this := StagedProfileBillingInstructionType{}
	return &this
}

// NewStagedProfileBillingInstructionTypeWithDefaults instantiates a new StagedProfileBillingInstructionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagedProfileBillingInstructionTypeWithDefaults() *StagedProfileBillingInstructionType {
	this := StagedProfileBillingInstructionType{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *StagedProfileBillingInstructionType) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileBillingInstructionType) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *StagedProfileBillingInstructionType) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *StagedProfileBillingInstructionType) SetDesc(v string) {
	o.Desc = &v
}

// GetRoutingInstructionsId returns the RoutingInstructionsId field value if set, zero value otherwise.
func (o *StagedProfileBillingInstructionType) GetRoutingInstructionsId() float32 {
	if o == nil || IsNil(o.RoutingInstructionsId) {
		var ret float32
		return ret
	}
	return *o.RoutingInstructionsId
}

// GetRoutingInstructionsIdOk returns a tuple with the RoutingInstructionsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileBillingInstructionType) GetRoutingInstructionsIdOk() (*float32, bool) {
	if o == nil || IsNil(o.RoutingInstructionsId) {
		return nil, false
	}
	return o.RoutingInstructionsId, true
}

// HasRoutingInstructionsId returns a boolean if a field has been set.
func (o *StagedProfileBillingInstructionType) HasRoutingInstructionsId() bool {
	if o != nil && !IsNil(o.RoutingInstructionsId) {
		return true
	}

	return false
}

// SetRoutingInstructionsId gets a reference to the given float32 and assigns it to the RoutingInstructionsId field.
func (o *StagedProfileBillingInstructionType) SetRoutingInstructionsId(v float32) {
	o.RoutingInstructionsId = &v
}

// GetBillingCode returns the BillingCode field value if set, zero value otherwise.
func (o *StagedProfileBillingInstructionType) GetBillingCode() string {
	if o == nil || IsNil(o.BillingCode) {
		var ret string
		return ret
	}
	return *o.BillingCode
}

// GetBillingCodeOk returns a tuple with the BillingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileBillingInstructionType) GetBillingCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BillingCode) {
		return nil, false
	}
	return o.BillingCode, true
}

// HasBillingCode returns a boolean if a field has been set.
func (o *StagedProfileBillingInstructionType) HasBillingCode() bool {
	if o != nil && !IsNil(o.BillingCode) {
		return true
	}

	return false
}

// SetBillingCode gets a reference to the given string and assigns it to the BillingCode field.
func (o *StagedProfileBillingInstructionType) SetBillingCode(v string) {
	o.BillingCode = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *StagedProfileBillingInstructionType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileBillingInstructionType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *StagedProfileBillingInstructionType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *StagedProfileBillingInstructionType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *StagedProfileBillingInstructionType) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileBillingInstructionType) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *StagedProfileBillingInstructionType) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *StagedProfileBillingInstructionType) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StagedProfileBillingInstructionType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileBillingInstructionType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StagedProfileBillingInstructionType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StagedProfileBillingInstructionType) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StagedProfileBillingInstructionType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileBillingInstructionType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StagedProfileBillingInstructionType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StagedProfileBillingInstructionType) SetType(v string) {
	o.Type = &v
}

func (o StagedProfileBillingInstructionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StagedProfileBillingInstructionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.RoutingInstructionsId) {
		toSerialize["routingInstructionsId"] = o.RoutingInstructionsId
	}
	if !IsNil(o.BillingCode) {
		toSerialize["billingCode"] = o.BillingCode
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableStagedProfileBillingInstructionType struct {
	value *StagedProfileBillingInstructionType
	isSet bool
}

func (v NullableStagedProfileBillingInstructionType) Get() *StagedProfileBillingInstructionType {
	return v.value
}

func (v *NullableStagedProfileBillingInstructionType) Set(val *StagedProfileBillingInstructionType) {
	v.value = val
	v.isSet = true
}

func (v NullableStagedProfileBillingInstructionType) IsSet() bool {
	return v.isSet
}

func (v *NullableStagedProfileBillingInstructionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagedProfileBillingInstructionType(val *StagedProfileBillingInstructionType) *NullableStagedProfileBillingInstructionType {
	return &NullableStagedProfileBillingInstructionType{value: val, isSet: true}
}

func (v NullableStagedProfileBillingInstructionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagedProfileBillingInstructionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


