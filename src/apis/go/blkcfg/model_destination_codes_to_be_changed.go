/*
OPERA Cloud Block Configuration API

APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blkcfg

import (
	"encoding/json"
)

// checks if the DestinationCodesToBeChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationCodesToBeChanged{}

// DestinationCodesToBeChanged Request object for changing Destination Codes.
type DestinationCodesToBeChanged struct {
	// List of Destination Codes.
	DestinationCodes []DestinationCodeType `json:"destinationCodes,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewDestinationCodesToBeChanged instantiates a new DestinationCodesToBeChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationCodesToBeChanged() *DestinationCodesToBeChanged {
	this := DestinationCodesToBeChanged{}
	return &this
}

// NewDestinationCodesToBeChangedWithDefaults instantiates a new DestinationCodesToBeChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationCodesToBeChangedWithDefaults() *DestinationCodesToBeChanged {
	this := DestinationCodesToBeChanged{}
	return &this
}

// GetDestinationCodes returns the DestinationCodes field value if set, zero value otherwise.
func (o *DestinationCodesToBeChanged) GetDestinationCodes() []DestinationCodeType {
	if o == nil || IsNil(o.DestinationCodes) {
		var ret []DestinationCodeType
		return ret
	}
	return o.DestinationCodes
}

// GetDestinationCodesOk returns a tuple with the DestinationCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationCodesToBeChanged) GetDestinationCodesOk() ([]DestinationCodeType, bool) {
	if o == nil || IsNil(o.DestinationCodes) {
		return nil, false
	}
	return o.DestinationCodes, true
}

// HasDestinationCodes returns a boolean if a field has been set.
func (o *DestinationCodesToBeChanged) HasDestinationCodes() bool {
	if o != nil && !IsNil(o.DestinationCodes) {
		return true
	}

	return false
}

// SetDestinationCodes gets a reference to the given []DestinationCodeType and assigns it to the DestinationCodes field.
func (o *DestinationCodesToBeChanged) SetDestinationCodes(v []DestinationCodeType) {
	o.DestinationCodes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DestinationCodesToBeChanged) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationCodesToBeChanged) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DestinationCodesToBeChanged) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *DestinationCodesToBeChanged) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DestinationCodesToBeChanged) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationCodesToBeChanged) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DestinationCodesToBeChanged) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *DestinationCodesToBeChanged) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o DestinationCodesToBeChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationCodesToBeChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestinationCodes) {
		toSerialize["destinationCodes"] = o.DestinationCodes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableDestinationCodesToBeChanged struct {
	value *DestinationCodesToBeChanged
	isSet bool
}

func (v NullableDestinationCodesToBeChanged) Get() *DestinationCodesToBeChanged {
	return v.value
}

func (v *NullableDestinationCodesToBeChanged) Set(val *DestinationCodesToBeChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationCodesToBeChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationCodesToBeChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationCodesToBeChanged(val *DestinationCodesToBeChanged) *NullableDestinationCodesToBeChanged {
	return &NullableDestinationCodesToBeChanged{value: val, isSet: true}
}

func (v NullableDestinationCodesToBeChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationCodesToBeChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


