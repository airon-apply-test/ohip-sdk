/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelBillingStatement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelBillingStatement{}

// ChannelBillingStatement Response object of the channel billing statement fetch request.
type ChannelBillingStatement struct {
	ChannelBillingStatement *ChannelBillingStatementType `json:"channelBillingStatement,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChannelBillingStatement instantiates a new ChannelBillingStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelBillingStatement() *ChannelBillingStatement {
	this := ChannelBillingStatement{}
	return &this
}

// NewChannelBillingStatementWithDefaults instantiates a new ChannelBillingStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelBillingStatementWithDefaults() *ChannelBillingStatement {
	this := ChannelBillingStatement{}
	return &this
}

// GetChannelBillingStatement returns the ChannelBillingStatement field value if set, zero value otherwise.
func (o *ChannelBillingStatement) GetChannelBillingStatement() ChannelBillingStatementType {
	if o == nil || IsNil(o.ChannelBillingStatement) {
		var ret ChannelBillingStatementType
		return ret
	}
	return *o.ChannelBillingStatement
}

// GetChannelBillingStatementOk returns a tuple with the ChannelBillingStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatement) GetChannelBillingStatementOk() (*ChannelBillingStatementType, bool) {
	if o == nil || IsNil(o.ChannelBillingStatement) {
		return nil, false
	}
	return o.ChannelBillingStatement, true
}

// HasChannelBillingStatement returns a boolean if a field has been set.
func (o *ChannelBillingStatement) HasChannelBillingStatement() bool {
	if o != nil && !IsNil(o.ChannelBillingStatement) {
		return true
	}

	return false
}

// SetChannelBillingStatement gets a reference to the given ChannelBillingStatementType and assigns it to the ChannelBillingStatement field.
func (o *ChannelBillingStatement) SetChannelBillingStatement(v ChannelBillingStatementType) {
	o.ChannelBillingStatement = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChannelBillingStatement) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatement) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChannelBillingStatement) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChannelBillingStatement) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChannelBillingStatement) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatement) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChannelBillingStatement) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChannelBillingStatement) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChannelBillingStatement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelBillingStatement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelBillingStatement) {
		toSerialize["channelBillingStatement"] = o.ChannelBillingStatement
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChannelBillingStatement struct {
	value *ChannelBillingStatement
	isSet bool
}

func (v NullableChannelBillingStatement) Get() *ChannelBillingStatement {
	return v.value
}

func (v *NullableChannelBillingStatement) Set(val *ChannelBillingStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelBillingStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelBillingStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelBillingStatement(val *ChannelBillingStatement) *NullableChannelBillingStatement {
	return &NullableChannelBillingStatement{value: val, isSet: true}
}

func (v NullableChannelBillingStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelBillingStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


