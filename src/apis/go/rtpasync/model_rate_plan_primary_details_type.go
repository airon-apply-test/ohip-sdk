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

// checks if the RatePlanPrimaryDetailsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatePlanPrimaryDetailsType{}

// RatePlanPrimaryDetailsType struct for RatePlanPrimaryDetailsType
type RatePlanPrimaryDetailsType struct {
	Description *TranslationTextType2000 `json:"description,omitempty"`
	// Start sell date of the rate plan.
	StartSellDate *string `json:"startSellDate,omitempty"`
	// End sell date of the rate plan.
	EndSellDate *string `json:"endSellDate,omitempty"`
	// The particular rate code is marked as privileged making it restrictive on who can update information.
	PrivilegedRate *bool `json:"privilegedRate,omitempty"`
	// The restrictions marked on the rate code are marked as privileged making it restrictive on who can update information.
	PrivilegedRateRestriction *bool `json:"privilegedRateRestriction,omitempty"`
	LockStatus *RateCodeLockStatusType `json:"lockStatus,omitempty"`
}

// NewRatePlanPrimaryDetailsType instantiates a new RatePlanPrimaryDetailsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatePlanPrimaryDetailsType() *RatePlanPrimaryDetailsType {
	this := RatePlanPrimaryDetailsType{}
	return &this
}

// NewRatePlanPrimaryDetailsTypeWithDefaults instantiates a new RatePlanPrimaryDetailsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatePlanPrimaryDetailsTypeWithDefaults() *RatePlanPrimaryDetailsType {
	this := RatePlanPrimaryDetailsType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RatePlanPrimaryDetailsType) GetDescription() TranslationTextType2000 {
	if o == nil || IsNil(o.Description) {
		var ret TranslationTextType2000
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanPrimaryDetailsType) GetDescriptionOk() (*TranslationTextType2000, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RatePlanPrimaryDetailsType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given TranslationTextType2000 and assigns it to the Description field.
func (o *RatePlanPrimaryDetailsType) SetDescription(v TranslationTextType2000) {
	o.Description = &v
}

// GetStartSellDate returns the StartSellDate field value if set, zero value otherwise.
func (o *RatePlanPrimaryDetailsType) GetStartSellDate() string {
	if o == nil || IsNil(o.StartSellDate) {
		var ret string
		return ret
	}
	return *o.StartSellDate
}

// GetStartSellDateOk returns a tuple with the StartSellDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanPrimaryDetailsType) GetStartSellDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartSellDate) {
		return nil, false
	}
	return o.StartSellDate, true
}

// HasStartSellDate returns a boolean if a field has been set.
func (o *RatePlanPrimaryDetailsType) HasStartSellDate() bool {
	if o != nil && !IsNil(o.StartSellDate) {
		return true
	}

	return false
}

// SetStartSellDate gets a reference to the given string and assigns it to the StartSellDate field.
func (o *RatePlanPrimaryDetailsType) SetStartSellDate(v string) {
	o.StartSellDate = &v
}

// GetEndSellDate returns the EndSellDate field value if set, zero value otherwise.
func (o *RatePlanPrimaryDetailsType) GetEndSellDate() string {
	if o == nil || IsNil(o.EndSellDate) {
		var ret string
		return ret
	}
	return *o.EndSellDate
}

// GetEndSellDateOk returns a tuple with the EndSellDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanPrimaryDetailsType) GetEndSellDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndSellDate) {
		return nil, false
	}
	return o.EndSellDate, true
}

// HasEndSellDate returns a boolean if a field has been set.
func (o *RatePlanPrimaryDetailsType) HasEndSellDate() bool {
	if o != nil && !IsNil(o.EndSellDate) {
		return true
	}

	return false
}

// SetEndSellDate gets a reference to the given string and assigns it to the EndSellDate field.
func (o *RatePlanPrimaryDetailsType) SetEndSellDate(v string) {
	o.EndSellDate = &v
}

// GetPrivilegedRate returns the PrivilegedRate field value if set, zero value otherwise.
func (o *RatePlanPrimaryDetailsType) GetPrivilegedRate() bool {
	if o == nil || IsNil(o.PrivilegedRate) {
		var ret bool
		return ret
	}
	return *o.PrivilegedRate
}

// GetPrivilegedRateOk returns a tuple with the PrivilegedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanPrimaryDetailsType) GetPrivilegedRateOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivilegedRate) {
		return nil, false
	}
	return o.PrivilegedRate, true
}

// HasPrivilegedRate returns a boolean if a field has been set.
func (o *RatePlanPrimaryDetailsType) HasPrivilegedRate() bool {
	if o != nil && !IsNil(o.PrivilegedRate) {
		return true
	}

	return false
}

// SetPrivilegedRate gets a reference to the given bool and assigns it to the PrivilegedRate field.
func (o *RatePlanPrimaryDetailsType) SetPrivilegedRate(v bool) {
	o.PrivilegedRate = &v
}

// GetPrivilegedRateRestriction returns the PrivilegedRateRestriction field value if set, zero value otherwise.
func (o *RatePlanPrimaryDetailsType) GetPrivilegedRateRestriction() bool {
	if o == nil || IsNil(o.PrivilegedRateRestriction) {
		var ret bool
		return ret
	}
	return *o.PrivilegedRateRestriction
}

// GetPrivilegedRateRestrictionOk returns a tuple with the PrivilegedRateRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanPrimaryDetailsType) GetPrivilegedRateRestrictionOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivilegedRateRestriction) {
		return nil, false
	}
	return o.PrivilegedRateRestriction, true
}

// HasPrivilegedRateRestriction returns a boolean if a field has been set.
func (o *RatePlanPrimaryDetailsType) HasPrivilegedRateRestriction() bool {
	if o != nil && !IsNil(o.PrivilegedRateRestriction) {
		return true
	}

	return false
}

// SetPrivilegedRateRestriction gets a reference to the given bool and assigns it to the PrivilegedRateRestriction field.
func (o *RatePlanPrimaryDetailsType) SetPrivilegedRateRestriction(v bool) {
	o.PrivilegedRateRestriction = &v
}

// GetLockStatus returns the LockStatus field value if set, zero value otherwise.
func (o *RatePlanPrimaryDetailsType) GetLockStatus() RateCodeLockStatusType {
	if o == nil || IsNil(o.LockStatus) {
		var ret RateCodeLockStatusType
		return ret
	}
	return *o.LockStatus
}

// GetLockStatusOk returns a tuple with the LockStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatePlanPrimaryDetailsType) GetLockStatusOk() (*RateCodeLockStatusType, bool) {
	if o == nil || IsNil(o.LockStatus) {
		return nil, false
	}
	return o.LockStatus, true
}

// HasLockStatus returns a boolean if a field has been set.
func (o *RatePlanPrimaryDetailsType) HasLockStatus() bool {
	if o != nil && !IsNil(o.LockStatus) {
		return true
	}

	return false
}

// SetLockStatus gets a reference to the given RateCodeLockStatusType and assigns it to the LockStatus field.
func (o *RatePlanPrimaryDetailsType) SetLockStatus(v RateCodeLockStatusType) {
	o.LockStatus = &v
}

func (o RatePlanPrimaryDetailsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatePlanPrimaryDetailsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.StartSellDate) {
		toSerialize["startSellDate"] = o.StartSellDate
	}
	if !IsNil(o.EndSellDate) {
		toSerialize["endSellDate"] = o.EndSellDate
	}
	if !IsNil(o.PrivilegedRate) {
		toSerialize["privilegedRate"] = o.PrivilegedRate
	}
	if !IsNil(o.PrivilegedRateRestriction) {
		toSerialize["privilegedRateRestriction"] = o.PrivilegedRateRestriction
	}
	if !IsNil(o.LockStatus) {
		toSerialize["lockStatus"] = o.LockStatus
	}
	return toSerialize, nil
}

type NullableRatePlanPrimaryDetailsType struct {
	value *RatePlanPrimaryDetailsType
	isSet bool
}

func (v NullableRatePlanPrimaryDetailsType) Get() *RatePlanPrimaryDetailsType {
	return v.value
}

func (v *NullableRatePlanPrimaryDetailsType) Set(val *RatePlanPrimaryDetailsType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatePlanPrimaryDetailsType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatePlanPrimaryDetailsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatePlanPrimaryDetailsType(val *RatePlanPrimaryDetailsType) *NullableRatePlanPrimaryDetailsType {
	return &NullableRatePlanPrimaryDetailsType{value: val, isSet: true}
}

func (v NullableRatePlanPrimaryDetailsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatePlanPrimaryDetailsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


