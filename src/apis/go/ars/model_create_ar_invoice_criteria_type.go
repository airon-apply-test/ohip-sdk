/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateARInvoiceCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateARInvoiceCriteriaType{}

// CreateARInvoiceCriteriaType Criteria to Create a new AR Invoice.
type CreateARInvoiceCriteriaType struct {
	// Custom Folio Name Value Information to be saved
	FolioNameValue []NameValueHeaderDetailType `json:"folioNameValue,omitempty"`
	Account *ARAccountCriteriaType `json:"account,omitempty"`
	GuestProfileId *UniqueIDType `json:"guestProfileId,omitempty"`
	// Collection of Charges to be posted.
	Charges []ChargeCriteriaType `json:"charges,omitempty"`
	// User-defined invoice reference.
	Reference *string `json:"reference,omitempty"`
	// User-defined invoice remark.
	Remark *string `json:"remark,omitempty"`
	// Invoice market code.
	Market *string `json:"market,omitempty"`
	// Invoice room class code.
	RoomClass *string `json:"roomClass,omitempty"`
	// Invoice source code.
	Source *string `json:"source,omitempty"`
	FiscalFolioInfo *FiscalServiceType `json:"fiscalFolioInfo,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
	OverrideCreditHoldCheck *bool `json:"overrideCreditHoldCheck,omitempty"`
}

// NewCreateARInvoiceCriteriaType instantiates a new CreateARInvoiceCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateARInvoiceCriteriaType() *CreateARInvoiceCriteriaType {
	this := CreateARInvoiceCriteriaType{}
	return &this
}

// NewCreateARInvoiceCriteriaTypeWithDefaults instantiates a new CreateARInvoiceCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateARInvoiceCriteriaTypeWithDefaults() *CreateARInvoiceCriteriaType {
	this := CreateARInvoiceCriteriaType{}
	return &this
}

// GetFolioNameValue returns the FolioNameValue field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetFolioNameValue() []NameValueHeaderDetailType {
	if o == nil || IsNil(o.FolioNameValue) {
		var ret []NameValueHeaderDetailType
		return ret
	}
	return o.FolioNameValue
}

// GetFolioNameValueOk returns a tuple with the FolioNameValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetFolioNameValueOk() ([]NameValueHeaderDetailType, bool) {
	if o == nil || IsNil(o.FolioNameValue) {
		return nil, false
	}
	return o.FolioNameValue, true
}

// HasFolioNameValue returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasFolioNameValue() bool {
	if o != nil && !IsNil(o.FolioNameValue) {
		return true
	}

	return false
}

// SetFolioNameValue gets a reference to the given []NameValueHeaderDetailType and assigns it to the FolioNameValue field.
func (o *CreateARInvoiceCriteriaType) SetFolioNameValue(v []NameValueHeaderDetailType) {
	o.FolioNameValue = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetAccount() ARAccountCriteriaType {
	if o == nil || IsNil(o.Account) {
		var ret ARAccountCriteriaType
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetAccountOk() (*ARAccountCriteriaType, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ARAccountCriteriaType and assigns it to the Account field.
func (o *CreateARInvoiceCriteriaType) SetAccount(v ARAccountCriteriaType) {
	o.Account = &v
}

// GetGuestProfileId returns the GuestProfileId field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetGuestProfileId() UniqueIDType {
	if o == nil || IsNil(o.GuestProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.GuestProfileId
}

// GetGuestProfileIdOk returns a tuple with the GuestProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetGuestProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.GuestProfileId) {
		return nil, false
	}
	return o.GuestProfileId, true
}

// HasGuestProfileId returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasGuestProfileId() bool {
	if o != nil && !IsNil(o.GuestProfileId) {
		return true
	}

	return false
}

// SetGuestProfileId gets a reference to the given UniqueIDType and assigns it to the GuestProfileId field.
func (o *CreateARInvoiceCriteriaType) SetGuestProfileId(v UniqueIDType) {
	o.GuestProfileId = &v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetCharges() []ChargeCriteriaType {
	if o == nil || IsNil(o.Charges) {
		var ret []ChargeCriteriaType
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetChargesOk() ([]ChargeCriteriaType, bool) {
	if o == nil || IsNil(o.Charges) {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasCharges() bool {
	if o != nil && !IsNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given []ChargeCriteriaType and assigns it to the Charges field.
func (o *CreateARInvoiceCriteriaType) SetCharges(v []ChargeCriteriaType) {
	o.Charges = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CreateARInvoiceCriteriaType) SetReference(v string) {
	o.Reference = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *CreateARInvoiceCriteriaType) SetRemark(v string) {
	o.Remark = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetMarket() string {
	if o == nil || IsNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetMarketOk() (*string, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CreateARInvoiceCriteriaType) SetMarket(v string) {
	o.Market = &v
}

// GetRoomClass returns the RoomClass field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetRoomClass() string {
	if o == nil || IsNil(o.RoomClass) {
		var ret string
		return ret
	}
	return *o.RoomClass
}

// GetRoomClassOk returns a tuple with the RoomClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetRoomClassOk() (*string, bool) {
	if o == nil || IsNil(o.RoomClass) {
		return nil, false
	}
	return o.RoomClass, true
}

// HasRoomClass returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasRoomClass() bool {
	if o != nil && !IsNil(o.RoomClass) {
		return true
	}

	return false
}

// SetRoomClass gets a reference to the given string and assigns it to the RoomClass field.
func (o *CreateARInvoiceCriteriaType) SetRoomClass(v string) {
	o.RoomClass = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CreateARInvoiceCriteriaType) SetSource(v string) {
	o.Source = &v
}

// GetFiscalFolioInfo returns the FiscalFolioInfo field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetFiscalFolioInfo() FiscalServiceType {
	if o == nil || IsNil(o.FiscalFolioInfo) {
		var ret FiscalServiceType
		return ret
	}
	return *o.FiscalFolioInfo
}

// GetFiscalFolioInfoOk returns a tuple with the FiscalFolioInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetFiscalFolioInfoOk() (*FiscalServiceType, bool) {
	if o == nil || IsNil(o.FiscalFolioInfo) {
		return nil, false
	}
	return o.FiscalFolioInfo, true
}

// HasFiscalFolioInfo returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasFiscalFolioInfo() bool {
	if o != nil && !IsNil(o.FiscalFolioInfo) {
		return true
	}

	return false
}

// SetFiscalFolioInfo gets a reference to the given FiscalServiceType and assigns it to the FiscalFolioInfo field.
func (o *CreateARInvoiceCriteriaType) SetFiscalFolioInfo(v FiscalServiceType) {
	o.FiscalFolioInfo = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *CreateARInvoiceCriteriaType) SetCashierId(v float32) {
	o.CashierId = &v
}

// GetOverrideCreditHoldCheck returns the OverrideCreditHoldCheck field value if set, zero value otherwise.
func (o *CreateARInvoiceCriteriaType) GetOverrideCreditHoldCheck() bool {
	if o == nil || IsNil(o.OverrideCreditHoldCheck) {
		var ret bool
		return ret
	}
	return *o.OverrideCreditHoldCheck
}

// GetOverrideCreditHoldCheckOk returns a tuple with the OverrideCreditHoldCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateARInvoiceCriteriaType) GetOverrideCreditHoldCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.OverrideCreditHoldCheck) {
		return nil, false
	}
	return o.OverrideCreditHoldCheck, true
}

// HasOverrideCreditHoldCheck returns a boolean if a field has been set.
func (o *CreateARInvoiceCriteriaType) HasOverrideCreditHoldCheck() bool {
	if o != nil && !IsNil(o.OverrideCreditHoldCheck) {
		return true
	}

	return false
}

// SetOverrideCreditHoldCheck gets a reference to the given bool and assigns it to the OverrideCreditHoldCheck field.
func (o *CreateARInvoiceCriteriaType) SetOverrideCreditHoldCheck(v bool) {
	o.OverrideCreditHoldCheck = &v
}

func (o CreateARInvoiceCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateARInvoiceCriteriaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolioNameValue) {
		toSerialize["folioNameValue"] = o.FolioNameValue
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.GuestProfileId) {
		toSerialize["guestProfileId"] = o.GuestProfileId
	}
	if !IsNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.RoomClass) {
		toSerialize["roomClass"] = o.RoomClass
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.FiscalFolioInfo) {
		toSerialize["fiscalFolioInfo"] = o.FiscalFolioInfo
	}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	if !IsNil(o.OverrideCreditHoldCheck) {
		toSerialize["overrideCreditHoldCheck"] = o.OverrideCreditHoldCheck
	}
	return toSerialize, nil
}

type NullableCreateARInvoiceCriteriaType struct {
	value *CreateARInvoiceCriteriaType
	isSet bool
}

func (v NullableCreateARInvoiceCriteriaType) Get() *CreateARInvoiceCriteriaType {
	return v.value
}

func (v *NullableCreateARInvoiceCriteriaType) Set(val *CreateARInvoiceCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateARInvoiceCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateARInvoiceCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateARInvoiceCriteriaType(val *CreateARInvoiceCriteriaType) *NullableCreateARInvoiceCriteriaType {
	return &NullableCreateARInvoiceCriteriaType{value: val, isSet: true}
}

func (v NullableCreateARInvoiceCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateARInvoiceCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


