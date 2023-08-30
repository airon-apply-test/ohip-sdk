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

// checks if the CreditBillCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditBillCriteriaType{}

// CreditBillCriteriaType Criteria for posting the Credit Bill. Includes charges and payments.
type CreditBillCriteriaType struct {
	// Property where the charges are to be posted.
	HotelId *string `json:"hotelId,omitempty"`
	// Collection of Charges to be posted.
	Charges []ChargeCriteriaType `json:"charges,omitempty"`
	// The payment information to be posted.
	Payments []PaymentCriteriaType `json:"payments,omitempty"`
	FiscalFolioInfo *FiscalServiceType `json:"fiscalFolioInfo,omitempty"`
	// Date of the Audit. This is used when postings are being created using the Income Audit functionality.
	IncomeAuditDate *string `json:"incomeAuditDate,omitempty"`
	// Applicable for Fiscal Terminal. The ID of the terminal where the fiscal device is connected.
	FiscalTerminalId *string `json:"fiscalTerminalId,omitempty"`
	// Custom Folio Name Value Informatoin to be saved
	FolioNameValue []NameValueHeaderDetailType `json:"folioNameValue,omitempty"`
	// Transaction service type which the Folio is being associated.
	TrxServiceType *string `json:"trxServiceType,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
	OriginalFolio *FolioType `json:"originalFolio,omitempty"`
	// Flag can be used when a Credit Bill is to be created for full set of transactions in the Original Bill.
	FullCredit *bool `json:"fullCredit,omitempty"`
}

// NewCreditBillCriteriaType instantiates a new CreditBillCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBillCriteriaType() *CreditBillCriteriaType {
	this := CreditBillCriteriaType{}
	return &this
}

// NewCreditBillCriteriaTypeWithDefaults instantiates a new CreditBillCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBillCriteriaTypeWithDefaults() *CreditBillCriteriaType {
	this := CreditBillCriteriaType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *CreditBillCriteriaType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetCharges() []ChargeCriteriaType {
	if o == nil || IsNil(o.Charges) {
		var ret []ChargeCriteriaType
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetChargesOk() ([]ChargeCriteriaType, bool) {
	if o == nil || IsNil(o.Charges) {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasCharges() bool {
	if o != nil && !IsNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given []ChargeCriteriaType and assigns it to the Charges field.
func (o *CreditBillCriteriaType) SetCharges(v []ChargeCriteriaType) {
	o.Charges = v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetPayments() []PaymentCriteriaType {
	if o == nil || IsNil(o.Payments) {
		var ret []PaymentCriteriaType
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetPaymentsOk() ([]PaymentCriteriaType, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []PaymentCriteriaType and assigns it to the Payments field.
func (o *CreditBillCriteriaType) SetPayments(v []PaymentCriteriaType) {
	o.Payments = v
}

// GetFiscalFolioInfo returns the FiscalFolioInfo field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetFiscalFolioInfo() FiscalServiceType {
	if o == nil || IsNil(o.FiscalFolioInfo) {
		var ret FiscalServiceType
		return ret
	}
	return *o.FiscalFolioInfo
}

// GetFiscalFolioInfoOk returns a tuple with the FiscalFolioInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetFiscalFolioInfoOk() (*FiscalServiceType, bool) {
	if o == nil || IsNil(o.FiscalFolioInfo) {
		return nil, false
	}
	return o.FiscalFolioInfo, true
}

// HasFiscalFolioInfo returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasFiscalFolioInfo() bool {
	if o != nil && !IsNil(o.FiscalFolioInfo) {
		return true
	}

	return false
}

// SetFiscalFolioInfo gets a reference to the given FiscalServiceType and assigns it to the FiscalFolioInfo field.
func (o *CreditBillCriteriaType) SetFiscalFolioInfo(v FiscalServiceType) {
	o.FiscalFolioInfo = &v
}

// GetIncomeAuditDate returns the IncomeAuditDate field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetIncomeAuditDate() string {
	if o == nil || IsNil(o.IncomeAuditDate) {
		var ret string
		return ret
	}
	return *o.IncomeAuditDate
}

// GetIncomeAuditDateOk returns a tuple with the IncomeAuditDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetIncomeAuditDateOk() (*string, bool) {
	if o == nil || IsNil(o.IncomeAuditDate) {
		return nil, false
	}
	return o.IncomeAuditDate, true
}

// HasIncomeAuditDate returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasIncomeAuditDate() bool {
	if o != nil && !IsNil(o.IncomeAuditDate) {
		return true
	}

	return false
}

// SetIncomeAuditDate gets a reference to the given string and assigns it to the IncomeAuditDate field.
func (o *CreditBillCriteriaType) SetIncomeAuditDate(v string) {
	o.IncomeAuditDate = &v
}

// GetFiscalTerminalId returns the FiscalTerminalId field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetFiscalTerminalId() string {
	if o == nil || IsNil(o.FiscalTerminalId) {
		var ret string
		return ret
	}
	return *o.FiscalTerminalId
}

// GetFiscalTerminalIdOk returns a tuple with the FiscalTerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetFiscalTerminalIdOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalTerminalId) {
		return nil, false
	}
	return o.FiscalTerminalId, true
}

// HasFiscalTerminalId returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasFiscalTerminalId() bool {
	if o != nil && !IsNil(o.FiscalTerminalId) {
		return true
	}

	return false
}

// SetFiscalTerminalId gets a reference to the given string and assigns it to the FiscalTerminalId field.
func (o *CreditBillCriteriaType) SetFiscalTerminalId(v string) {
	o.FiscalTerminalId = &v
}

// GetFolioNameValue returns the FolioNameValue field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetFolioNameValue() []NameValueHeaderDetailType {
	if o == nil || IsNil(o.FolioNameValue) {
		var ret []NameValueHeaderDetailType
		return ret
	}
	return o.FolioNameValue
}

// GetFolioNameValueOk returns a tuple with the FolioNameValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetFolioNameValueOk() ([]NameValueHeaderDetailType, bool) {
	if o == nil || IsNil(o.FolioNameValue) {
		return nil, false
	}
	return o.FolioNameValue, true
}

// HasFolioNameValue returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasFolioNameValue() bool {
	if o != nil && !IsNil(o.FolioNameValue) {
		return true
	}

	return false
}

// SetFolioNameValue gets a reference to the given []NameValueHeaderDetailType and assigns it to the FolioNameValue field.
func (o *CreditBillCriteriaType) SetFolioNameValue(v []NameValueHeaderDetailType) {
	o.FolioNameValue = v
}

// GetTrxServiceType returns the TrxServiceType field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetTrxServiceType() string {
	if o == nil || IsNil(o.TrxServiceType) {
		var ret string
		return ret
	}
	return *o.TrxServiceType
}

// GetTrxServiceTypeOk returns a tuple with the TrxServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetTrxServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrxServiceType) {
		return nil, false
	}
	return o.TrxServiceType, true
}

// HasTrxServiceType returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasTrxServiceType() bool {
	if o != nil && !IsNil(o.TrxServiceType) {
		return true
	}

	return false
}

// SetTrxServiceType gets a reference to the given string and assigns it to the TrxServiceType field.
func (o *CreditBillCriteriaType) SetTrxServiceType(v string) {
	o.TrxServiceType = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *CreditBillCriteriaType) SetCashierId(v float32) {
	o.CashierId = &v
}

// GetOriginalFolio returns the OriginalFolio field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetOriginalFolio() FolioType {
	if o == nil || IsNil(o.OriginalFolio) {
		var ret FolioType
		return ret
	}
	return *o.OriginalFolio
}

// GetOriginalFolioOk returns a tuple with the OriginalFolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetOriginalFolioOk() (*FolioType, bool) {
	if o == nil || IsNil(o.OriginalFolio) {
		return nil, false
	}
	return o.OriginalFolio, true
}

// HasOriginalFolio returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasOriginalFolio() bool {
	if o != nil && !IsNil(o.OriginalFolio) {
		return true
	}

	return false
}

// SetOriginalFolio gets a reference to the given FolioType and assigns it to the OriginalFolio field.
func (o *CreditBillCriteriaType) SetOriginalFolio(v FolioType) {
	o.OriginalFolio = &v
}

// GetFullCredit returns the FullCredit field value if set, zero value otherwise.
func (o *CreditBillCriteriaType) GetFullCredit() bool {
	if o == nil || IsNil(o.FullCredit) {
		var ret bool
		return ret
	}
	return *o.FullCredit
}

// GetFullCreditOk returns a tuple with the FullCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBillCriteriaType) GetFullCreditOk() (*bool, bool) {
	if o == nil || IsNil(o.FullCredit) {
		return nil, false
	}
	return o.FullCredit, true
}

// HasFullCredit returns a boolean if a field has been set.
func (o *CreditBillCriteriaType) HasFullCredit() bool {
	if o != nil && !IsNil(o.FullCredit) {
		return true
	}

	return false
}

// SetFullCredit gets a reference to the given bool and assigns it to the FullCredit field.
func (o *CreditBillCriteriaType) SetFullCredit(v bool) {
	o.FullCredit = &v
}

func (o CreditBillCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditBillCriteriaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !IsNil(o.Payments) {
		toSerialize["payments"] = o.Payments
	}
	if !IsNil(o.FiscalFolioInfo) {
		toSerialize["fiscalFolioInfo"] = o.FiscalFolioInfo
	}
	if !IsNil(o.IncomeAuditDate) {
		toSerialize["incomeAuditDate"] = o.IncomeAuditDate
	}
	if !IsNil(o.FiscalTerminalId) {
		toSerialize["fiscalTerminalId"] = o.FiscalTerminalId
	}
	if !IsNil(o.FolioNameValue) {
		toSerialize["folioNameValue"] = o.FolioNameValue
	}
	if !IsNil(o.TrxServiceType) {
		toSerialize["trxServiceType"] = o.TrxServiceType
	}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	if !IsNil(o.OriginalFolio) {
		toSerialize["originalFolio"] = o.OriginalFolio
	}
	if !IsNil(o.FullCredit) {
		toSerialize["fullCredit"] = o.FullCredit
	}
	return toSerialize, nil
}

type NullableCreditBillCriteriaType struct {
	value *CreditBillCriteriaType
	isSet bool
}

func (v NullableCreditBillCriteriaType) Get() *CreditBillCriteriaType {
	return v.value
}

func (v *NullableCreditBillCriteriaType) Set(val *CreditBillCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBillCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBillCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBillCriteriaType(val *CreditBillCriteriaType) *NullableCreditBillCriteriaType {
	return &NullableCreditBillCriteriaType{value: val, isSet: true}
}

func (v NullableCreditBillCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBillCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


