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

// checks if the CreatedRoutingInstructions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedRoutingInstructions{}

// CreatedRoutingInstructions Response when creating a routing instruction. It may optionally return a list of postings which are eligible for transfer as per the new routing instruction in case of room routing instruction.
type CreatedRoutingInstructions struct {
	PostingsForRoomRouting *PostingsInfoType `json:"postingsForRoomRouting,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCreatedRoutingInstructions instantiates a new CreatedRoutingInstructions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedRoutingInstructions() *CreatedRoutingInstructions {
	this := CreatedRoutingInstructions{}
	return &this
}

// NewCreatedRoutingInstructionsWithDefaults instantiates a new CreatedRoutingInstructions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedRoutingInstructionsWithDefaults() *CreatedRoutingInstructions {
	this := CreatedRoutingInstructions{}
	return &this
}

// GetPostingsForRoomRouting returns the PostingsForRoomRouting field value if set, zero value otherwise.
func (o *CreatedRoutingInstructions) GetPostingsForRoomRouting() PostingsInfoType {
	if o == nil || IsNil(o.PostingsForRoomRouting) {
		var ret PostingsInfoType
		return ret
	}
	return *o.PostingsForRoomRouting
}

// GetPostingsForRoomRoutingOk returns a tuple with the PostingsForRoomRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedRoutingInstructions) GetPostingsForRoomRoutingOk() (*PostingsInfoType, bool) {
	if o == nil || IsNil(o.PostingsForRoomRouting) {
		return nil, false
	}
	return o.PostingsForRoomRouting, true
}

// HasPostingsForRoomRouting returns a boolean if a field has been set.
func (o *CreatedRoutingInstructions) HasPostingsForRoomRouting() bool {
	if o != nil && !IsNil(o.PostingsForRoomRouting) {
		return true
	}

	return false
}

// SetPostingsForRoomRouting gets a reference to the given PostingsInfoType and assigns it to the PostingsForRoomRouting field.
func (o *CreatedRoutingInstructions) SetPostingsForRoomRouting(v PostingsInfoType) {
	o.PostingsForRoomRouting = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreatedRoutingInstructions) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedRoutingInstructions) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreatedRoutingInstructions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CreatedRoutingInstructions) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreatedRoutingInstructions) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedRoutingInstructions) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreatedRoutingInstructions) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CreatedRoutingInstructions) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CreatedRoutingInstructions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedRoutingInstructions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostingsForRoomRouting) {
		toSerialize["postingsForRoomRouting"] = o.PostingsForRoomRouting
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCreatedRoutingInstructions struct {
	value *CreatedRoutingInstructions
	isSet bool
}

func (v NullableCreatedRoutingInstructions) Get() *CreatedRoutingInstructions {
	return v.value
}

func (v *NullableCreatedRoutingInstructions) Set(val *CreatedRoutingInstructions) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedRoutingInstructions) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedRoutingInstructions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedRoutingInstructions(val *CreatedRoutingInstructions) *NullableCreatedRoutingInstructions {
	return &NullableCreatedRoutingInstructions{value: val, isSet: true}
}

func (v NullableCreatedRoutingInstructions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedRoutingInstructions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


