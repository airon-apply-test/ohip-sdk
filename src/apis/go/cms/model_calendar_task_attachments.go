/*
OPERA Cloud API for Customer Management Service

This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CalendarTaskAttachments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalendarTaskAttachments{}

// CalendarTaskAttachments Response for fetching calendar task attachments.
type CalendarTaskAttachments struct {
	// Attachment List.
	CalendarTaskAttachments []AttachmentType `json:"calendarTaskAttachments,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCalendarTaskAttachments instantiates a new CalendarTaskAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarTaskAttachments() *CalendarTaskAttachments {
	this := CalendarTaskAttachments{}
	return &this
}

// NewCalendarTaskAttachmentsWithDefaults instantiates a new CalendarTaskAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarTaskAttachmentsWithDefaults() *CalendarTaskAttachments {
	this := CalendarTaskAttachments{}
	return &this
}

// GetCalendarTaskAttachments returns the CalendarTaskAttachments field value if set, zero value otherwise.
func (o *CalendarTaskAttachments) GetCalendarTaskAttachments() []AttachmentType {
	if o == nil || IsNil(o.CalendarTaskAttachments) {
		var ret []AttachmentType
		return ret
	}
	return o.CalendarTaskAttachments
}

// GetCalendarTaskAttachmentsOk returns a tuple with the CalendarTaskAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTaskAttachments) GetCalendarTaskAttachmentsOk() ([]AttachmentType, bool) {
	if o == nil || IsNil(o.CalendarTaskAttachments) {
		return nil, false
	}
	return o.CalendarTaskAttachments, true
}

// HasCalendarTaskAttachments returns a boolean if a field has been set.
func (o *CalendarTaskAttachments) HasCalendarTaskAttachments() bool {
	if o != nil && !IsNil(o.CalendarTaskAttachments) {
		return true
	}

	return false
}

// SetCalendarTaskAttachments gets a reference to the given []AttachmentType and assigns it to the CalendarTaskAttachments field.
func (o *CalendarTaskAttachments) SetCalendarTaskAttachments(v []AttachmentType) {
	o.CalendarTaskAttachments = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CalendarTaskAttachments) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTaskAttachments) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CalendarTaskAttachments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *CalendarTaskAttachments) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CalendarTaskAttachments) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarTaskAttachments) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CalendarTaskAttachments) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *CalendarTaskAttachments) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o CalendarTaskAttachments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarTaskAttachments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CalendarTaskAttachments) {
		toSerialize["calendarTaskAttachments"] = o.CalendarTaskAttachments
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCalendarTaskAttachments struct {
	value *CalendarTaskAttachments
	isSet bool
}

func (v NullableCalendarTaskAttachments) Get() *CalendarTaskAttachments {
	return v.value
}

func (v *NullableCalendarTaskAttachments) Set(val *CalendarTaskAttachments) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarTaskAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarTaskAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarTaskAttachments(val *CalendarTaskAttachments) *NullableCalendarTaskAttachments {
	return &NullableCalendarTaskAttachments{value: val, isSet: true}
}

func (v NullableCalendarTaskAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarTaskAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


