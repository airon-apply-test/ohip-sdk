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

// checks if the FiscalFolioCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiscalFolioCriteriaType{}

// FiscalFolioCriteriaType Criteria for calling fiscal folio service for the generation of invoices
type FiscalFolioCriteriaType struct {
	// Property where the folio is being generated.
	HotelId *string `json:"hotelId,omitempty"`
	FolioCommand *FolioCommandType `json:"folioCommand,omitempty"`
	// Effective date to run fiscal command.
	EffectiveDate *string `json:"effectiveDate,omitempty"`
	Folios *FiscalInvoiceSummaryType `json:"folios,omitempty"`
}

// NewFiscalFolioCriteriaType instantiates a new FiscalFolioCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiscalFolioCriteriaType() *FiscalFolioCriteriaType {
	this := FiscalFolioCriteriaType{}
	return &this
}

// NewFiscalFolioCriteriaTypeWithDefaults instantiates a new FiscalFolioCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiscalFolioCriteriaTypeWithDefaults() *FiscalFolioCriteriaType {
	this := FiscalFolioCriteriaType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *FiscalFolioCriteriaType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalFolioCriteriaType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *FiscalFolioCriteriaType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *FiscalFolioCriteriaType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetFolioCommand returns the FolioCommand field value if set, zero value otherwise.
func (o *FiscalFolioCriteriaType) GetFolioCommand() FolioCommandType {
	if o == nil || IsNil(o.FolioCommand) {
		var ret FolioCommandType
		return ret
	}
	return *o.FolioCommand
}

// GetFolioCommandOk returns a tuple with the FolioCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalFolioCriteriaType) GetFolioCommandOk() (*FolioCommandType, bool) {
	if o == nil || IsNil(o.FolioCommand) {
		return nil, false
	}
	return o.FolioCommand, true
}

// HasFolioCommand returns a boolean if a field has been set.
func (o *FiscalFolioCriteriaType) HasFolioCommand() bool {
	if o != nil && !IsNil(o.FolioCommand) {
		return true
	}

	return false
}

// SetFolioCommand gets a reference to the given FolioCommandType and assigns it to the FolioCommand field.
func (o *FiscalFolioCriteriaType) SetFolioCommand(v FolioCommandType) {
	o.FolioCommand = &v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *FiscalFolioCriteriaType) GetEffectiveDate() string {
	if o == nil || IsNil(o.EffectiveDate) {
		var ret string
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalFolioCriteriaType) GetEffectiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveDate) {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *FiscalFolioCriteriaType) HasEffectiveDate() bool {
	if o != nil && !IsNil(o.EffectiveDate) {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given string and assigns it to the EffectiveDate field.
func (o *FiscalFolioCriteriaType) SetEffectiveDate(v string) {
	o.EffectiveDate = &v
}

// GetFolios returns the Folios field value if set, zero value otherwise.
func (o *FiscalFolioCriteriaType) GetFolios() FiscalInvoiceSummaryType {
	if o == nil || IsNil(o.Folios) {
		var ret FiscalInvoiceSummaryType
		return ret
	}
	return *o.Folios
}

// GetFoliosOk returns a tuple with the Folios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiscalFolioCriteriaType) GetFoliosOk() (*FiscalInvoiceSummaryType, bool) {
	if o == nil || IsNil(o.Folios) {
		return nil, false
	}
	return o.Folios, true
}

// HasFolios returns a boolean if a field has been set.
func (o *FiscalFolioCriteriaType) HasFolios() bool {
	if o != nil && !IsNil(o.Folios) {
		return true
	}

	return false
}

// SetFolios gets a reference to the given FiscalInvoiceSummaryType and assigns it to the Folios field.
func (o *FiscalFolioCriteriaType) SetFolios(v FiscalInvoiceSummaryType) {
	o.Folios = &v
}

func (o FiscalFolioCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiscalFolioCriteriaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.FolioCommand) {
		toSerialize["folioCommand"] = o.FolioCommand
	}
	if !IsNil(o.EffectiveDate) {
		toSerialize["effectiveDate"] = o.EffectiveDate
	}
	if !IsNil(o.Folios) {
		toSerialize["folios"] = o.Folios
	}
	return toSerialize, nil
}

type NullableFiscalFolioCriteriaType struct {
	value *FiscalFolioCriteriaType
	isSet bool
}

func (v NullableFiscalFolioCriteriaType) Get() *FiscalFolioCriteriaType {
	return v.value
}

func (v *NullableFiscalFolioCriteriaType) Set(val *FiscalFolioCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableFiscalFolioCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableFiscalFolioCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiscalFolioCriteriaType(val *FiscalFolioCriteriaType) *NullableFiscalFolioCriteriaType {
	return &NullableFiscalFolioCriteriaType{value: val, isSet: true}
}

func (v NullableFiscalFolioCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiscalFolioCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


