/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the FetchIntegrationMessageAttachments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchIntegrationMessageAttachments{}

// FetchIntegrationMessageAttachments struct for FetchIntegrationMessageAttachments
type FetchIntegrationMessageAttachments struct {
	Message *IntegrationMessageAttachmentsType `json:"message,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFetchIntegrationMessageAttachments instantiates a new FetchIntegrationMessageAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchIntegrationMessageAttachments() *FetchIntegrationMessageAttachments {
	this := FetchIntegrationMessageAttachments{}
	return &this
}

// NewFetchIntegrationMessageAttachmentsWithDefaults instantiates a new FetchIntegrationMessageAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchIntegrationMessageAttachmentsWithDefaults() *FetchIntegrationMessageAttachments {
	this := FetchIntegrationMessageAttachments{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FetchIntegrationMessageAttachments) GetMessage() IntegrationMessageAttachmentsType {
	if o == nil || IsNil(o.Message) {
		var ret IntegrationMessageAttachmentsType
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchIntegrationMessageAttachments) GetMessageOk() (*IntegrationMessageAttachmentsType, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FetchIntegrationMessageAttachments) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given IntegrationMessageAttachmentsType and assigns it to the Message field.
func (o *FetchIntegrationMessageAttachments) SetMessage(v IntegrationMessageAttachmentsType) {
	o.Message = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FetchIntegrationMessageAttachments) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchIntegrationMessageAttachments) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FetchIntegrationMessageAttachments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FetchIntegrationMessageAttachments) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FetchIntegrationMessageAttachments) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchIntegrationMessageAttachments) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FetchIntegrationMessageAttachments) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FetchIntegrationMessageAttachments) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FetchIntegrationMessageAttachments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchIntegrationMessageAttachments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFetchIntegrationMessageAttachments struct {
	value *FetchIntegrationMessageAttachments
	isSet bool
}

func (v NullableFetchIntegrationMessageAttachments) Get() *FetchIntegrationMessageAttachments {
	return v.value
}

func (v *NullableFetchIntegrationMessageAttachments) Set(val *FetchIntegrationMessageAttachments) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchIntegrationMessageAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchIntegrationMessageAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchIntegrationMessageAttachments(val *FetchIntegrationMessageAttachments) *NullableFetchIntegrationMessageAttachments {
	return &NullableFetchIntegrationMessageAttachments{value: val, isSet: true}
}

func (v NullableFetchIntegrationMessageAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchIntegrationMessageAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


