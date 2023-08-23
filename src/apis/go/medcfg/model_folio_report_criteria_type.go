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

// checks if the FolioReportCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolioReportCriteriaType{}

// FolioReportCriteriaType Criteria required to generate or retrieve a folio for a reservation.
type FolioReportCriteriaType struct {
	// Hotel code for the folio.
	HotelId *string `json:"hotelId,omitempty"`
	ReservationId *UniqueIDType `json:"reservationId,omitempty"`
	// Folio window number to generate (defaults to view 1).
	FolioWindowNo *int32 `json:"folioWindowNo,omitempty"`
	// Optional bill number to generate.
	BillNumber *int32 `json:"billNumber,omitempty"`
	// Optional folio type.
	FolioType *string `json:"folioType,omitempty"`
	// Optional folio generation date - defaults to hotel business date.
	FolioDate *string `json:"folioDate,omitempty"`
	// Folio currency code.
	ReferenceCurrency *string `json:"referenceCurrency,omitempty"`
}

// NewFolioReportCriteriaType instantiates a new FolioReportCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolioReportCriteriaType() *FolioReportCriteriaType {
	this := FolioReportCriteriaType{}
	return &this
}

// NewFolioReportCriteriaTypeWithDefaults instantiates a new FolioReportCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolioReportCriteriaTypeWithDefaults() *FolioReportCriteriaType {
	this := FolioReportCriteriaType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *FolioReportCriteriaType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReportCriteriaType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *FolioReportCriteriaType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *FolioReportCriteriaType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *FolioReportCriteriaType) GetReservationId() UniqueIDType {
	if o == nil || IsNil(o.ReservationId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReportCriteriaType) GetReservationIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *FolioReportCriteriaType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given UniqueIDType and assigns it to the ReservationId field.
func (o *FolioReportCriteriaType) SetReservationId(v UniqueIDType) {
	o.ReservationId = &v
}

// GetFolioWindowNo returns the FolioWindowNo field value if set, zero value otherwise.
func (o *FolioReportCriteriaType) GetFolioWindowNo() int32 {
	if o == nil || IsNil(o.FolioWindowNo) {
		var ret int32
		return ret
	}
	return *o.FolioWindowNo
}

// GetFolioWindowNoOk returns a tuple with the FolioWindowNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReportCriteriaType) GetFolioWindowNoOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioWindowNo) {
		return nil, false
	}
	return o.FolioWindowNo, true
}

// HasFolioWindowNo returns a boolean if a field has been set.
func (o *FolioReportCriteriaType) HasFolioWindowNo() bool {
	if o != nil && !IsNil(o.FolioWindowNo) {
		return true
	}

	return false
}

// SetFolioWindowNo gets a reference to the given int32 and assigns it to the FolioWindowNo field.
func (o *FolioReportCriteriaType) SetFolioWindowNo(v int32) {
	o.FolioWindowNo = &v
}

// GetBillNumber returns the BillNumber field value if set, zero value otherwise.
func (o *FolioReportCriteriaType) GetBillNumber() int32 {
	if o == nil || IsNil(o.BillNumber) {
		var ret int32
		return ret
	}
	return *o.BillNumber
}

// GetBillNumberOk returns a tuple with the BillNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReportCriteriaType) GetBillNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.BillNumber) {
		return nil, false
	}
	return o.BillNumber, true
}

// HasBillNumber returns a boolean if a field has been set.
func (o *FolioReportCriteriaType) HasBillNumber() bool {
	if o != nil && !IsNil(o.BillNumber) {
		return true
	}

	return false
}

// SetBillNumber gets a reference to the given int32 and assigns it to the BillNumber field.
func (o *FolioReportCriteriaType) SetBillNumber(v int32) {
	o.BillNumber = &v
}

// GetFolioType returns the FolioType field value if set, zero value otherwise.
func (o *FolioReportCriteriaType) GetFolioType() string {
	if o == nil || IsNil(o.FolioType) {
		var ret string
		return ret
	}
	return *o.FolioType
}

// GetFolioTypeOk returns a tuple with the FolioType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReportCriteriaType) GetFolioTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FolioType) {
		return nil, false
	}
	return o.FolioType, true
}

// HasFolioType returns a boolean if a field has been set.
func (o *FolioReportCriteriaType) HasFolioType() bool {
	if o != nil && !IsNil(o.FolioType) {
		return true
	}

	return false
}

// SetFolioType gets a reference to the given string and assigns it to the FolioType field.
func (o *FolioReportCriteriaType) SetFolioType(v string) {
	o.FolioType = &v
}

// GetFolioDate returns the FolioDate field value if set, zero value otherwise.
func (o *FolioReportCriteriaType) GetFolioDate() string {
	if o == nil || IsNil(o.FolioDate) {
		var ret string
		return ret
	}
	return *o.FolioDate
}

// GetFolioDateOk returns a tuple with the FolioDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReportCriteriaType) GetFolioDateOk() (*string, bool) {
	if o == nil || IsNil(o.FolioDate) {
		return nil, false
	}
	return o.FolioDate, true
}

// HasFolioDate returns a boolean if a field has been set.
func (o *FolioReportCriteriaType) HasFolioDate() bool {
	if o != nil && !IsNil(o.FolioDate) {
		return true
	}

	return false
}

// SetFolioDate gets a reference to the given string and assigns it to the FolioDate field.
func (o *FolioReportCriteriaType) SetFolioDate(v string) {
	o.FolioDate = &v
}

// GetReferenceCurrency returns the ReferenceCurrency field value if set, zero value otherwise.
func (o *FolioReportCriteriaType) GetReferenceCurrency() string {
	if o == nil || IsNil(o.ReferenceCurrency) {
		var ret string
		return ret
	}
	return *o.ReferenceCurrency
}

// GetReferenceCurrencyOk returns a tuple with the ReferenceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolioReportCriteriaType) GetReferenceCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceCurrency) {
		return nil, false
	}
	return o.ReferenceCurrency, true
}

// HasReferenceCurrency returns a boolean if a field has been set.
func (o *FolioReportCriteriaType) HasReferenceCurrency() bool {
	if o != nil && !IsNil(o.ReferenceCurrency) {
		return true
	}

	return false
}

// SetReferenceCurrency gets a reference to the given string and assigns it to the ReferenceCurrency field.
func (o *FolioReportCriteriaType) SetReferenceCurrency(v string) {
	o.ReferenceCurrency = &v
}

func (o FolioReportCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolioReportCriteriaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.FolioWindowNo) {
		toSerialize["folioWindowNo"] = o.FolioWindowNo
	}
	if !IsNil(o.BillNumber) {
		toSerialize["billNumber"] = o.BillNumber
	}
	if !IsNil(o.FolioType) {
		toSerialize["folioType"] = o.FolioType
	}
	if !IsNil(o.FolioDate) {
		toSerialize["folioDate"] = o.FolioDate
	}
	if !IsNil(o.ReferenceCurrency) {
		toSerialize["referenceCurrency"] = o.ReferenceCurrency
	}
	return toSerialize, nil
}

type NullableFolioReportCriteriaType struct {
	value *FolioReportCriteriaType
	isSet bool
}

func (v NullableFolioReportCriteriaType) Get() *FolioReportCriteriaType {
	return v.value
}

func (v *NullableFolioReportCriteriaType) Set(val *FolioReportCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableFolioReportCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableFolioReportCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolioReportCriteriaType(val *FolioReportCriteriaType) *NullableFolioReportCriteriaType {
	return &NullableFolioReportCriteriaType{value: val, isSet: true}
}

func (v NullableFolioReportCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolioReportCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


