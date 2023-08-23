/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PutCashierRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutCashierRequest{}

// PutCashierRequest struct for PutCashierRequest
type PutCashierRequest struct {
	Criteria *CashierClosureType `json:"criteria,omitempty"`
	// The Cashier Lock Handle to pass along with operation which required cashier to be locked.
	CashierLockHandle *float32 `json:"cashierLockHandle,omitempty"`
	// Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
	HotelId *string `json:"hotelId,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPutCashierRequest instantiates a new PutCashierRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutCashierRequest() *PutCashierRequest {
	this := PutCashierRequest{}
	return &this
}

// NewPutCashierRequestWithDefaults instantiates a new PutCashierRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutCashierRequestWithDefaults() *PutCashierRequest {
	this := PutCashierRequest{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PutCashierRequest) GetCriteria() CashierClosureType {
	if o == nil || IsNil(o.Criteria) {
		var ret CashierClosureType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutCashierRequest) GetCriteriaOk() (*CashierClosureType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PutCashierRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given CashierClosureType and assigns it to the Criteria field.
func (o *PutCashierRequest) SetCriteria(v CashierClosureType) {
	o.Criteria = &v
}

// GetCashierLockHandle returns the CashierLockHandle field value if set, zero value otherwise.
func (o *PutCashierRequest) GetCashierLockHandle() float32 {
	if o == nil || IsNil(o.CashierLockHandle) {
		var ret float32
		return ret
	}
	return *o.CashierLockHandle
}

// GetCashierLockHandleOk returns a tuple with the CashierLockHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutCashierRequest) GetCashierLockHandleOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierLockHandle) {
		return nil, false
	}
	return o.CashierLockHandle, true
}

// HasCashierLockHandle returns a boolean if a field has been set.
func (o *PutCashierRequest) HasCashierLockHandle() bool {
	if o != nil && !IsNil(o.CashierLockHandle) {
		return true
	}

	return false
}

// SetCashierLockHandle gets a reference to the given float32 and assigns it to the CashierLockHandle field.
func (o *PutCashierRequest) SetCashierLockHandle(v float32) {
	o.CashierLockHandle = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *PutCashierRequest) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutCashierRequest) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *PutCashierRequest) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *PutCashierRequest) SetHotelId(v string) {
	o.HotelId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutCashierRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutCashierRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutCashierRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PutCashierRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PutCashierRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutCashierRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PutCashierRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PutCashierRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PutCashierRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutCashierRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.CashierLockHandle) {
		toSerialize["cashierLockHandle"] = o.CashierLockHandle
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePutCashierRequest struct {
	value *PutCashierRequest
	isSet bool
}

func (v NullablePutCashierRequest) Get() *PutCashierRequest {
	return v.value
}

func (v *NullablePutCashierRequest) Set(val *PutCashierRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutCashierRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutCashierRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutCashierRequest(val *PutCashierRequest) *NullablePutCashierRequest {
	return &NullablePutCashierRequest{value: val, isSet: true}
}

func (v NullablePutCashierRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutCashierRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


