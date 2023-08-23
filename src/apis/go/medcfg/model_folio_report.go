/*
OPERA Cloud Content Service

Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FolioReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolioReport{}

// FolioReport Response object for folio report.
type FolioReport struct {
	Folio *FolioReportResultType `json:"folio,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewFolioReport instantiates a new FolioReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolioReport() *FolioReport {
	this := FolioReport{}
	return &this
}

// NewFolioReportWithDefaults instantiates a new FolioReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolioReportWithDefaults() *FolioReport {
	this := FolioReport{}
	return &this
}

// GetFolio returns the Folio field value if set, zero value otherwise.
func (o *FolioReport) GetFolio() FolioReportResultType {
	if o == nil || IsNil(o.Folio) {
		var ret FolioReportResultType
		return ret
	}
	return *o.Folio
}

// GetFolioOk returns a tuple with the Folio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReport) GetFolioOk() (*FolioReportResultType, bool) {
	if o == nil || IsNil(o.Folio) {
		return nil, false
	}
	return o.Folio, true
}

// HasFolio returns a boolean if a field has been set.
func (o *FolioReport) HasFolio() bool {
	if o != nil && !IsNil(o.Folio) {
		return true
	}

	return false
}

// SetFolio gets a reference to the given FolioReportResultType and assigns it to the Folio field.
func (o *FolioReport) SetFolio(v FolioReportResultType) {
	o.Folio = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FolioReport) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReport) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FolioReport) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *FolioReport) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FolioReport) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReport) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FolioReport) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *FolioReport) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o FolioReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolioReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Folio) {
		toSerialize["folio"] = o.Folio
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableFolioReport struct {
	value *FolioReport
	isSet bool
}

func (v NullableFolioReport) Get() *FolioReport {
	return v.value
}

func (v *NullableFolioReport) Set(val *FolioReport) {
	v.value = val
	v.isSet = true
}

func (v NullableFolioReport) IsSet() bool {
	return v.isSet
}

func (v *NullableFolioReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolioReport(val *FolioReport) *NullableFolioReport {
	return &NullableFolioReport{value: val, isSet: true}
}

func (v NullableFolioReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolioReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


