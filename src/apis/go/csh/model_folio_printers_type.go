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

// checks if the FolioPrintersType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolioPrintersType{}

// FolioPrintersType List of Folio Printers.
type FolioPrintersType struct {
	// Name of the Folio Type.
	FolioTypeName *string `json:"folioTypeName,omitempty"`
	// Folio Printer Information.
	Printer []FolioPrinterType `json:"printer,omitempty"`
}

// NewFolioPrintersType instantiates a new FolioPrintersType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolioPrintersType() *FolioPrintersType {
	this := FolioPrintersType{}
	return &this
}

// NewFolioPrintersTypeWithDefaults instantiates a new FolioPrintersType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolioPrintersTypeWithDefaults() *FolioPrintersType {
	this := FolioPrintersType{}
	return &this
}

// GetFolioTypeName returns the FolioTypeName field value if set, zero value otherwise.
func (o *FolioPrintersType) GetFolioTypeName() string {
	if o == nil || IsNil(o.FolioTypeName) {
		var ret string
		return ret
	}
	return *o.FolioTypeName
}

// GetFolioTypeNameOk returns a tuple with the FolioTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioPrintersType) GetFolioTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.FolioTypeName) {
		return nil, false
	}
	return o.FolioTypeName, true
}

// HasFolioTypeName returns a boolean if a field has been set.
func (o *FolioPrintersType) HasFolioTypeName() bool {
	if o != nil && !IsNil(o.FolioTypeName) {
		return true
	}

	return false
}

// SetFolioTypeName gets a reference to the given string and assigns it to the FolioTypeName field.
func (o *FolioPrintersType) SetFolioTypeName(v string) {
	o.FolioTypeName = &v
}

// GetPrinter returns the Printer field value if set, zero value otherwise.
func (o *FolioPrintersType) GetPrinter() []FolioPrinterType {
	if o == nil || IsNil(o.Printer) {
		var ret []FolioPrinterType
		return ret
	}
	return o.Printer
}

// GetPrinterOk returns a tuple with the Printer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioPrintersType) GetPrinterOk() ([]FolioPrinterType, bool) {
	if o == nil || IsNil(o.Printer) {
		return nil, false
	}
	return o.Printer, true
}

// HasPrinter returns a boolean if a field has been set.
func (o *FolioPrintersType) HasPrinter() bool {
	if o != nil && !IsNil(o.Printer) {
		return true
	}

	return false
}

// SetPrinter gets a reference to the given []FolioPrinterType and assigns it to the Printer field.
func (o *FolioPrintersType) SetPrinter(v []FolioPrinterType) {
	o.Printer = v
}

func (o FolioPrintersType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolioPrintersType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolioTypeName) {
		toSerialize["folioTypeName"] = o.FolioTypeName
	}
	if !IsNil(o.Printer) {
		toSerialize["printer"] = o.Printer
	}
	return toSerialize, nil
}

type NullableFolioPrintersType struct {
	value *FolioPrintersType
	isSet bool
}

func (v NullableFolioPrintersType) Get() *FolioPrintersType {
	return v.value
}

func (v *NullableFolioPrintersType) Set(val *FolioPrintersType) {
	v.value = val
	v.isSet = true
}

func (v NullableFolioPrintersType) IsSet() bool {
	return v.isSet
}

func (v *NullableFolioPrintersType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolioPrintersType(val *FolioPrintersType) *NullableFolioPrintersType {
	return &NullableFolioPrintersType{value: val, isSet: true}
}

func (v NullableFolioPrintersType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolioPrintersType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


