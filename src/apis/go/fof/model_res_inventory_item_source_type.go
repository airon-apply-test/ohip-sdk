/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResInventoryItemSourceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResInventoryItemSourceType{}

// ResInventoryItemSourceType Defines whether the item is setup due to a Rate Plan, Package or a Block.
type ResInventoryItemSourceType struct {
	// Rate Plan Code, If populated specifies that the item is setup due to a Rate Plan.
	RatePlanCode *string `json:"ratePlanCode,omitempty"`
	// Package Code, If populated specifies that the item is setup due to a Package.
	PackageCode *string `json:"packageCode,omitempty"`
	BlockId *BlockId `json:"blockId,omitempty"`
	// If true, it implies that the item has been attached to the reservation as part of a Welcome Offer
	WelcomeOffer *bool `json:"welcomeOffer,omitempty"`
	// Source Reservation Package Opera Internal Unique Id. This is the unique Id used for the reservation package associated with this item.
	PackageInternalID *float32 `json:"packageInternalID,omitempty"`
}

// NewResInventoryItemSourceType instantiates a new ResInventoryItemSourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResInventoryItemSourceType() *ResInventoryItemSourceType {
	this := ResInventoryItemSourceType{}
	return &this
}

// NewResInventoryItemSourceTypeWithDefaults instantiates a new ResInventoryItemSourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResInventoryItemSourceTypeWithDefaults() *ResInventoryItemSourceType {
	this := ResInventoryItemSourceType{}
	return &this
}

// GetRatePlanCode returns the RatePlanCode field value if set, zero value otherwise.
func (o *ResInventoryItemSourceType) GetRatePlanCode() string {
	if o == nil || IsNil(o.RatePlanCode) {
		var ret string
		return ret
	}
	return *o.RatePlanCode
}

// GetRatePlanCodeOk returns a tuple with the RatePlanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResInventoryItemSourceType) GetRatePlanCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RatePlanCode) {
		return nil, false
	}
	return o.RatePlanCode, true
}

// HasRatePlanCode returns a boolean if a field has been set.
func (o *ResInventoryItemSourceType) HasRatePlanCode() bool {
	if o != nil && !IsNil(o.RatePlanCode) {
		return true
	}

	return false
}

// SetRatePlanCode gets a reference to the given string and assigns it to the RatePlanCode field.
func (o *ResInventoryItemSourceType) SetRatePlanCode(v string) {
	o.RatePlanCode = &v
}

// GetPackageCode returns the PackageCode field value if set, zero value otherwise.
func (o *ResInventoryItemSourceType) GetPackageCode() string {
	if o == nil || IsNil(o.PackageCode) {
		var ret string
		return ret
	}
	return *o.PackageCode
}

// GetPackageCodeOk returns a tuple with the PackageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResInventoryItemSourceType) GetPackageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageCode) {
		return nil, false
	}
	return o.PackageCode, true
}

// HasPackageCode returns a boolean if a field has been set.
func (o *ResInventoryItemSourceType) HasPackageCode() bool {
	if o != nil && !IsNil(o.PackageCode) {
		return true
	}

	return false
}

// SetPackageCode gets a reference to the given string and assigns it to the PackageCode field.
func (o *ResInventoryItemSourceType) SetPackageCode(v string) {
	o.PackageCode = &v
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *ResInventoryItemSourceType) GetBlockId() BlockId {
	if o == nil || IsNil(o.BlockId) {
		var ret BlockId
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResInventoryItemSourceType) GetBlockIdOk() (*BlockId, bool) {
	if o == nil || IsNil(o.BlockId) {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *ResInventoryItemSourceType) HasBlockId() bool {
	if o != nil && !IsNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given BlockId and assigns it to the BlockId field.
func (o *ResInventoryItemSourceType) SetBlockId(v BlockId) {
	o.BlockId = &v
}

// GetWelcomeOffer returns the WelcomeOffer field value if set, zero value otherwise.
func (o *ResInventoryItemSourceType) GetWelcomeOffer() bool {
	if o == nil || IsNil(o.WelcomeOffer) {
		var ret bool
		return ret
	}
	return *o.WelcomeOffer
}

// GetWelcomeOfferOk returns a tuple with the WelcomeOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResInventoryItemSourceType) GetWelcomeOfferOk() (*bool, bool) {
	if o == nil || IsNil(o.WelcomeOffer) {
		return nil, false
	}
	return o.WelcomeOffer, true
}

// HasWelcomeOffer returns a boolean if a field has been set.
func (o *ResInventoryItemSourceType) HasWelcomeOffer() bool {
	if o != nil && !IsNil(o.WelcomeOffer) {
		return true
	}

	return false
}

// SetWelcomeOffer gets a reference to the given bool and assigns it to the WelcomeOffer field.
func (o *ResInventoryItemSourceType) SetWelcomeOffer(v bool) {
	o.WelcomeOffer = &v
}

// GetPackageInternalID returns the PackageInternalID field value if set, zero value otherwise.
func (o *ResInventoryItemSourceType) GetPackageInternalID() float32 {
	if o == nil || IsNil(o.PackageInternalID) {
		var ret float32
		return ret
	}
	return *o.PackageInternalID
}

// GetPackageInternalIDOk returns a tuple with the PackageInternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResInventoryItemSourceType) GetPackageInternalIDOk() (*float32, bool) {
	if o == nil || IsNil(o.PackageInternalID) {
		return nil, false
	}
	return o.PackageInternalID, true
}

// HasPackageInternalID returns a boolean if a field has been set.
func (o *ResInventoryItemSourceType) HasPackageInternalID() bool {
	if o != nil && !IsNil(o.PackageInternalID) {
		return true
	}

	return false
}

// SetPackageInternalID gets a reference to the given float32 and assigns it to the PackageInternalID field.
func (o *ResInventoryItemSourceType) SetPackageInternalID(v float32) {
	o.PackageInternalID = &v
}

func (o ResInventoryItemSourceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResInventoryItemSourceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RatePlanCode) {
		toSerialize["ratePlanCode"] = o.RatePlanCode
	}
	if !IsNil(o.PackageCode) {
		toSerialize["packageCode"] = o.PackageCode
	}
	if !IsNil(o.BlockId) {
		toSerialize["blockId"] = o.BlockId
	}
	if !IsNil(o.WelcomeOffer) {
		toSerialize["welcomeOffer"] = o.WelcomeOffer
	}
	if !IsNil(o.PackageInternalID) {
		toSerialize["packageInternalID"] = o.PackageInternalID
	}
	return toSerialize, nil
}

type NullableResInventoryItemSourceType struct {
	value *ResInventoryItemSourceType
	isSet bool
}

func (v NullableResInventoryItemSourceType) Get() *ResInventoryItemSourceType {
	return v.value
}

func (v *NullableResInventoryItemSourceType) Set(val *ResInventoryItemSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableResInventoryItemSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableResInventoryItemSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResInventoryItemSourceType(val *ResInventoryItemSourceType) *NullableResInventoryItemSourceType {
	return &NullableResInventoryItemSourceType{value: val, isSet: true}
}

func (v NullableResInventoryItemSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResInventoryItemSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


