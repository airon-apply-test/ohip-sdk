/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateSalesManagerGoals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSalesManagerGoals{}

// CreateSalesManagerGoals Request object for Creating Sales Manager Goal(s).
type CreateSalesManagerGoals struct {
	// Detail Information about Sales Manager's goal.
	SalesManagerGoals []SalesManagerGoalType `json:"salesManagerGoals,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreateSalesManagerGoals instantiates a new CreateSalesManagerGoals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSalesManagerGoals() *CreateSalesManagerGoals {
	this := CreateSalesManagerGoals{}
	return &this
}

// NewCreateSalesManagerGoalsWithDefaults instantiates a new CreateSalesManagerGoals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSalesManagerGoalsWithDefaults() *CreateSalesManagerGoals {
	this := CreateSalesManagerGoals{}
	return &this
}

// GetSalesManagerGoals returns the SalesManagerGoals field value if set, zero value otherwise.
func (o *CreateSalesManagerGoals) GetSalesManagerGoals() []SalesManagerGoalType {
	if o == nil || IsNil(o.SalesManagerGoals) {
		var ret []SalesManagerGoalType
		return ret
	}
	return o.SalesManagerGoals
}

// GetSalesManagerGoalsOk returns a tuple with the SalesManagerGoals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSalesManagerGoals) GetSalesManagerGoalsOk() ([]SalesManagerGoalType, bool) {
	if o == nil || IsNil(o.SalesManagerGoals) {
		return nil, false
	}
	return o.SalesManagerGoals, true
}

// HasSalesManagerGoals returns a boolean if a field has been set.
func (o *CreateSalesManagerGoals) HasSalesManagerGoals() bool {
	if o != nil && !IsNil(o.SalesManagerGoals) {
		return true
	}

	return false
}

// SetSalesManagerGoals gets a reference to the given []SalesManagerGoalType and assigns it to the SalesManagerGoals field.
func (o *CreateSalesManagerGoals) SetSalesManagerGoals(v []SalesManagerGoalType) {
	o.SalesManagerGoals = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreateSalesManagerGoals) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSalesManagerGoals) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreateSalesManagerGoals) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreateSalesManagerGoals) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreateSalesManagerGoals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSalesManagerGoals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SalesManagerGoals) {
		toSerialize["salesManagerGoals"] = o.SalesManagerGoals
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCreateSalesManagerGoals struct {
	value *CreateSalesManagerGoals
	isSet bool
}

func (v NullableCreateSalesManagerGoals) Get() *CreateSalesManagerGoals {
	return v.value
}

func (v *NullableCreateSalesManagerGoals) Set(val *CreateSalesManagerGoals) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSalesManagerGoals) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSalesManagerGoals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSalesManagerGoals(val *CreateSalesManagerGoals) *NullableCreateSalesManagerGoals {
	return &NullableCreateSalesManagerGoals{value: val, isSet: true}
}

func (v NullableCreateSalesManagerGoals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSalesManagerGoals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


