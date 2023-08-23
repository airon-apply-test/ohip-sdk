/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FolioPrinterType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolioPrinterType{}

// FolioPrinterType Information about a Printer which can be used to print a Folio.
type FolioPrinterType struct {
	// Description of the Printer
	PrinterName *string `json:"printerName,omitempty"`
	// The Printer Device name.
	Device *string `json:"device,omitempty"`
	// The Folio Type for which this printer is being used.
	FolioType *string `json:"folioType,omitempty"`
	// Folio Queue name for which this printer is being used.
	FolioQueueName *string `json:"folioQueueName,omitempty"`
}

// NewFolioPrinterType instantiates a new FolioPrinterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolioPrinterType() *FolioPrinterType {
	this := FolioPrinterType{}
	return &this
}

// NewFolioPrinterTypeWithDefaults instantiates a new FolioPrinterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolioPrinterTypeWithDefaults() *FolioPrinterType {
	this := FolioPrinterType{}
	return &this
}

// GetPrinterName returns the PrinterName field value if set, zero value otherwise.
func (o *FolioPrinterType) GetPrinterName() string {
	if o == nil || IsNil(o.PrinterName) {
		var ret string
		return ret
	}
	return *o.PrinterName
}

// GetPrinterNameOk returns a tuple with the PrinterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioPrinterType) GetPrinterNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrinterName) {
		return nil, false
	}
	return o.PrinterName, true
}

// HasPrinterName returns a boolean if a field has been set.
func (o *FolioPrinterType) HasPrinterName() bool {
	if o != nil && !IsNil(o.PrinterName) {
		return true
	}

	return false
}

// SetPrinterName gets a reference to the given string and assigns it to the PrinterName field.
func (o *FolioPrinterType) SetPrinterName(v string) {
	o.PrinterName = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FolioPrinterType) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioPrinterType) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FolioPrinterType) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *FolioPrinterType) SetDevice(v string) {
	o.Device = &v
}

// GetFolioType returns the FolioType field value if set, zero value otherwise.
func (o *FolioPrinterType) GetFolioType() string {
	if o == nil || IsNil(o.FolioType) {
		var ret string
		return ret
	}
	return *o.FolioType
}

// GetFolioTypeOk returns a tuple with the FolioType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioPrinterType) GetFolioTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FolioType) {
		return nil, false
	}
	return o.FolioType, true
}

// HasFolioType returns a boolean if a field has been set.
func (o *FolioPrinterType) HasFolioType() bool {
	if o != nil && !IsNil(o.FolioType) {
		return true
	}

	return false
}

// SetFolioType gets a reference to the given string and assigns it to the FolioType field.
func (o *FolioPrinterType) SetFolioType(v string) {
	o.FolioType = &v
}

// GetFolioQueueName returns the FolioQueueName field value if set, zero value otherwise.
func (o *FolioPrinterType) GetFolioQueueName() string {
	if o == nil || IsNil(o.FolioQueueName) {
		var ret string
		return ret
	}
	return *o.FolioQueueName
}

// GetFolioQueueNameOk returns a tuple with the FolioQueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioPrinterType) GetFolioQueueNameOk() (*string, bool) {
	if o == nil || IsNil(o.FolioQueueName) {
		return nil, false
	}
	return o.FolioQueueName, true
}

// HasFolioQueueName returns a boolean if a field has been set.
func (o *FolioPrinterType) HasFolioQueueName() bool {
	if o != nil && !IsNil(o.FolioQueueName) {
		return true
	}

	return false
}

// SetFolioQueueName gets a reference to the given string and assigns it to the FolioQueueName field.
func (o *FolioPrinterType) SetFolioQueueName(v string) {
	o.FolioQueueName = &v
}

func (o FolioPrinterType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolioPrinterType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrinterName) {
		toSerialize["printerName"] = o.PrinterName
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.FolioType) {
		toSerialize["folioType"] = o.FolioType
	}
	if !IsNil(o.FolioQueueName) {
		toSerialize["folioQueueName"] = o.FolioQueueName
	}
	return toSerialize, nil
}

type NullableFolioPrinterType struct {
	value *FolioPrinterType
	isSet bool
}

func (v NullableFolioPrinterType) Get() *FolioPrinterType {
	return v.value
}

func (v *NullableFolioPrinterType) Set(val *FolioPrinterType) {
	v.value = val
	v.isSet = true
}

func (v NullableFolioPrinterType) IsSet() bool {
	return v.isSet
}

func (v *NullableFolioPrinterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolioPrinterType(val *FolioPrinterType) *NullableFolioPrinterType {
	return &NullableFolioPrinterType{value: val, isSet: true}
}

func (v NullableFolioPrinterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolioPrinterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


