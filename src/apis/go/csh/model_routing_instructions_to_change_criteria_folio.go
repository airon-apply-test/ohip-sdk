/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the RoutingInstructionsToChangeCriteriaFolio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingInstructionsToChangeCriteriaFolio{}

// RoutingInstructionsToChangeCriteriaFolio Folio routing type.
type RoutingInstructionsToChangeCriteriaFolio struct {
	GuestInfo *ResvRoutingInfoTypeFolioGuestInfo `json:"guestInfo,omitempty"`
	PayeeInfo *PayeeInfoType `json:"payeeInfo,omitempty"`
	// Accounts Receivable.
	ARNumber *string `json:"aRNumber,omitempty"`
	// Set of routing instructions associated to this routing type.
	Instructions []RoutingInstructionType `json:"instructions,omitempty"`
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	FolioWindowNo *int32 `json:"folioWindowNo,omitempty"`
}

// NewRoutingInstructionsToChangeCriteriaFolio instantiates a new RoutingInstructionsToChangeCriteriaFolio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingInstructionsToChangeCriteriaFolio() *RoutingInstructionsToChangeCriteriaFolio {
	this := RoutingInstructionsToChangeCriteriaFolio{}
	return &this
}

// NewRoutingInstructionsToChangeCriteriaFolioWithDefaults instantiates a new RoutingInstructionsToChangeCriteriaFolio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingInstructionsToChangeCriteriaFolioWithDefaults() *RoutingInstructionsToChangeCriteriaFolio {
	this := RoutingInstructionsToChangeCriteriaFolio{}
	return &this
}

// GetGuestInfo returns the GuestInfo field value if set, zero value otherwise.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetGuestInfo() ResvRoutingInfoTypeFolioGuestInfo {
	if o == nil || IsNil(o.GuestInfo) {
		var ret ResvRoutingInfoTypeFolioGuestInfo
		return ret
	}
	return *o.GuestInfo
}

// GetGuestInfoOk returns a tuple with the GuestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetGuestInfoOk() (*ResvRoutingInfoTypeFolioGuestInfo, bool) {
	if o == nil || IsNil(o.GuestInfo) {
		return nil, false
	}
	return o.GuestInfo, true
}

// HasGuestInfo returns a boolean if a field has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) HasGuestInfo() bool {
	if o != nil && !IsNil(o.GuestInfo) {
		return true
	}

	return false
}

// SetGuestInfo gets a reference to the given ResvRoutingInfoTypeFolioGuestInfo and assigns it to the GuestInfo field.
func (o *RoutingInstructionsToChangeCriteriaFolio) SetGuestInfo(v ResvRoutingInfoTypeFolioGuestInfo) {
	o.GuestInfo = &v
}

// GetPayeeInfo returns the PayeeInfo field value if set, zero value otherwise.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetPayeeInfo() PayeeInfoType {
	if o == nil || IsNil(o.PayeeInfo) {
		var ret PayeeInfoType
		return ret
	}
	return *o.PayeeInfo
}

// GetPayeeInfoOk returns a tuple with the PayeeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetPayeeInfoOk() (*PayeeInfoType, bool) {
	if o == nil || IsNil(o.PayeeInfo) {
		return nil, false
	}
	return o.PayeeInfo, true
}

// HasPayeeInfo returns a boolean if a field has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) HasPayeeInfo() bool {
	if o != nil && !IsNil(o.PayeeInfo) {
		return true
	}

	return false
}

// SetPayeeInfo gets a reference to the given PayeeInfoType and assigns it to the PayeeInfo field.
func (o *RoutingInstructionsToChangeCriteriaFolio) SetPayeeInfo(v PayeeInfoType) {
	o.PayeeInfo = &v
}

// GetARNumber returns the ARNumber field value if set, zero value otherwise.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetARNumber() string {
	if o == nil || IsNil(o.ARNumber) {
		var ret string
		return ret
	}
	return *o.ARNumber
}

// GetARNumberOk returns a tuple with the ARNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetARNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ARNumber) {
		return nil, false
	}
	return o.ARNumber, true
}

// HasARNumber returns a boolean if a field has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) HasARNumber() bool {
	if o != nil && !IsNil(o.ARNumber) {
		return true
	}

	return false
}

// SetARNumber gets a reference to the given string and assigns it to the ARNumber field.
func (o *RoutingInstructionsToChangeCriteriaFolio) SetARNumber(v string) {
	o.ARNumber = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetInstructions() []RoutingInstructionType {
	if o == nil || IsNil(o.Instructions) {
		var ret []RoutingInstructionType
		return ret
	}
	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetInstructionsOk() ([]RoutingInstructionType, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []RoutingInstructionType and assigns it to the Instructions field.
func (o *RoutingInstructionsToChangeCriteriaFolio) SetInstructions(v []RoutingInstructionType) {
	o.Instructions = v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *RoutingInstructionsToChangeCriteriaFolio) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetFolioWindowNo returns the FolioWindowNo field value if set, zero value otherwise.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetFolioWindowNo() int32 {
	if o == nil || IsNil(o.FolioWindowNo) {
		var ret int32
		return ret
	}
	return *o.FolioWindowNo
}

// GetFolioWindowNoOk returns a tuple with the FolioWindowNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) GetFolioWindowNoOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioWindowNo) {
		return nil, false
	}
	return o.FolioWindowNo, true
}

// HasFolioWindowNo returns a boolean if a field has been set.
func (o *RoutingInstructionsToChangeCriteriaFolio) HasFolioWindowNo() bool {
	if o != nil && !IsNil(o.FolioWindowNo) {
		return true
	}

	return false
}

// SetFolioWindowNo gets a reference to the given int32 and assigns it to the FolioWindowNo field.
func (o *RoutingInstructionsToChangeCriteriaFolio) SetFolioWindowNo(v int32) {
	o.FolioWindowNo = &v
}

func (o RoutingInstructionsToChangeCriteriaFolio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingInstructionsToChangeCriteriaFolio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestInfo) {
		toSerialize["guestInfo"] = o.GuestInfo
	}
	if !IsNil(o.PayeeInfo) {
		toSerialize["payeeInfo"] = o.PayeeInfo
	}
	if !IsNil(o.ARNumber) {
		toSerialize["aRNumber"] = o.ARNumber
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.FolioWindowNo) {
		toSerialize["folioWindowNo"] = o.FolioWindowNo
	}
	return toSerialize, nil
}

type NullableRoutingInstructionsToChangeCriteriaFolio struct {
	value *RoutingInstructionsToChangeCriteriaFolio
	isSet bool
}

func (v NullableRoutingInstructionsToChangeCriteriaFolio) Get() *RoutingInstructionsToChangeCriteriaFolio {
	return v.value
}

func (v *NullableRoutingInstructionsToChangeCriteriaFolio) Set(val *RoutingInstructionsToChangeCriteriaFolio) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingInstructionsToChangeCriteriaFolio) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingInstructionsToChangeCriteriaFolio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingInstructionsToChangeCriteriaFolio(val *RoutingInstructionsToChangeCriteriaFolio) *NullableRoutingInstructionsToChangeCriteriaFolio {
	return &NullableRoutingInstructionsToChangeCriteriaFolio{value: val, isSet: true}
}

func (v NullableRoutingInstructionsToChangeCriteriaFolio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingInstructionsToChangeCriteriaFolio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


