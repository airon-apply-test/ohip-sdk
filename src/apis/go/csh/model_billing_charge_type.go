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

// checks if the BillingChargeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingChargeType{}

// BillingChargeType struct for BillingChargeType
type BillingChargeType struct {
	// Transaction number of the posting being corrected.
	TransactionNo *float32 `json:"transactionNo,omitempty"`
	Price *CurrencyAmountType `json:"price,omitempty"`
	// Corrected posting quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// Corrected user-defined posting reference.
	Reference *string `json:"reference,omitempty"`
	// Corrected user-defined posting remark.
	Remark *string `json:"remark,omitempty"`
	// Corrected Cheque number.
	CheckNumber *string `json:"checkNumber,omitempty"`
	// Corrected Revenue Date.
	RevenueDate *string `json:"revenueDate,omitempty"`
	// Corrected POS covers - number of copies of receipts that got printed for that particular receipt.
	Covers *string `json:"covers,omitempty"`
	// Corrected arrangement code from the package associated to this transaction.
	ArrangementCode *string `json:"arrangementCode,omitempty"`
	// Approval code of the posting.
	ApprovalCode *string `json:"approvalCode,omitempty"`
	// Approval status of the posting.
	ApprovalStatus *string `json:"approvalStatus,omitempty"`
	// Approval date of the posting.
	ApprovalDate *string `json:"approvalDate,omitempty"`
	// The Cashier ID of the Cashier who is currently processing the transaction(s).
	CashierId *float32 `json:"cashierId,omitempty"`
}

// NewBillingChargeType instantiates a new BillingChargeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingChargeType() *BillingChargeType {
	this := BillingChargeType{}
	return &this
}

// NewBillingChargeTypeWithDefaults instantiates a new BillingChargeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingChargeTypeWithDefaults() *BillingChargeType {
	this := BillingChargeType{}
	return &this
}

// GetTransactionNo returns the TransactionNo field value if set, zero value otherwise.
func (o *BillingChargeType) GetTransactionNo() float32 {
	if o == nil || IsNil(o.TransactionNo) {
		var ret float32
		return ret
	}
	return *o.TransactionNo
}

// GetTransactionNoOk returns a tuple with the TransactionNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetTransactionNoOk() (*float32, bool) {
	if o == nil || IsNil(o.TransactionNo) {
		return nil, false
	}
	return o.TransactionNo, true
}

// HasTransactionNo returns a boolean if a field has been set.
func (o *BillingChargeType) HasTransactionNo() bool {
	if o != nil && !IsNil(o.TransactionNo) {
		return true
	}

	return false
}

// SetTransactionNo gets a reference to the given float32 and assigns it to the TransactionNo field.
func (o *BillingChargeType) SetTransactionNo(v float32) {
	o.TransactionNo = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingChargeType) GetPrice() CurrencyAmountType {
	if o == nil || IsNil(o.Price) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetPriceOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingChargeType) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given CurrencyAmountType and assigns it to the Price field.
func (o *BillingChargeType) SetPrice(v CurrencyAmountType) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BillingChargeType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BillingChargeType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *BillingChargeType) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *BillingChargeType) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *BillingChargeType) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *BillingChargeType) SetReference(v string) {
	o.Reference = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *BillingChargeType) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *BillingChargeType) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *BillingChargeType) SetRemark(v string) {
	o.Remark = &v
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise.
func (o *BillingChargeType) GetCheckNumber() string {
	if o == nil || IsNil(o.CheckNumber) {
		var ret string
		return ret
	}
	return *o.CheckNumber
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetCheckNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CheckNumber) {
		return nil, false
	}
	return o.CheckNumber, true
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *BillingChargeType) HasCheckNumber() bool {
	if o != nil && !IsNil(o.CheckNumber) {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given string and assigns it to the CheckNumber field.
func (o *BillingChargeType) SetCheckNumber(v string) {
	o.CheckNumber = &v
}

// GetRevenueDate returns the RevenueDate field value if set, zero value otherwise.
func (o *BillingChargeType) GetRevenueDate() string {
	if o == nil || IsNil(o.RevenueDate) {
		var ret string
		return ret
	}
	return *o.RevenueDate
}

// GetRevenueDateOk returns a tuple with the RevenueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetRevenueDateOk() (*string, bool) {
	if o == nil || IsNil(o.RevenueDate) {
		return nil, false
	}
	return o.RevenueDate, true
}

// HasRevenueDate returns a boolean if a field has been set.
func (o *BillingChargeType) HasRevenueDate() bool {
	if o != nil && !IsNil(o.RevenueDate) {
		return true
	}

	return false
}

// SetRevenueDate gets a reference to the given string and assigns it to the RevenueDate field.
func (o *BillingChargeType) SetRevenueDate(v string) {
	o.RevenueDate = &v
}

// GetCovers returns the Covers field value if set, zero value otherwise.
func (o *BillingChargeType) GetCovers() string {
	if o == nil || IsNil(o.Covers) {
		var ret string
		return ret
	}
	return *o.Covers
}

// GetCoversOk returns a tuple with the Covers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetCoversOk() (*string, bool) {
	if o == nil || IsNil(o.Covers) {
		return nil, false
	}
	return o.Covers, true
}

// HasCovers returns a boolean if a field has been set.
func (o *BillingChargeType) HasCovers() bool {
	if o != nil && !IsNil(o.Covers) {
		return true
	}

	return false
}

// SetCovers gets a reference to the given string and assigns it to the Covers field.
func (o *BillingChargeType) SetCovers(v string) {
	o.Covers = &v
}

// GetArrangementCode returns the ArrangementCode field value if set, zero value otherwise.
func (o *BillingChargeType) GetArrangementCode() string {
	if o == nil || IsNil(o.ArrangementCode) {
		var ret string
		return ret
	}
	return *o.ArrangementCode
}

// GetArrangementCodeOk returns a tuple with the ArrangementCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetArrangementCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ArrangementCode) {
		return nil, false
	}
	return o.ArrangementCode, true
}

// HasArrangementCode returns a boolean if a field has been set.
func (o *BillingChargeType) HasArrangementCode() bool {
	if o != nil && !IsNil(o.ArrangementCode) {
		return true
	}

	return false
}

// SetArrangementCode gets a reference to the given string and assigns it to the ArrangementCode field.
func (o *BillingChargeType) SetArrangementCode(v string) {
	o.ArrangementCode = &v
}

// GetApprovalCode returns the ApprovalCode field value if set, zero value otherwise.
func (o *BillingChargeType) GetApprovalCode() string {
	if o == nil || IsNil(o.ApprovalCode) {
		var ret string
		return ret
	}
	return *o.ApprovalCode
}

// GetApprovalCodeOk returns a tuple with the ApprovalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetApprovalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalCode) {
		return nil, false
	}
	return o.ApprovalCode, true
}

// HasApprovalCode returns a boolean if a field has been set.
func (o *BillingChargeType) HasApprovalCode() bool {
	if o != nil && !IsNil(o.ApprovalCode) {
		return true
	}

	return false
}

// SetApprovalCode gets a reference to the given string and assigns it to the ApprovalCode field.
func (o *BillingChargeType) SetApprovalCode(v string) {
	o.ApprovalCode = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *BillingChargeType) GetApprovalStatus() string {
	if o == nil || IsNil(o.ApprovalStatus) {
		var ret string
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetApprovalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalStatus) {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *BillingChargeType) HasApprovalStatus() bool {
	if o != nil && !IsNil(o.ApprovalStatus) {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given string and assigns it to the ApprovalStatus field.
func (o *BillingChargeType) SetApprovalStatus(v string) {
	o.ApprovalStatus = &v
}

// GetApprovalDate returns the ApprovalDate field value if set, zero value otherwise.
func (o *BillingChargeType) GetApprovalDate() string {
	if o == nil || IsNil(o.ApprovalDate) {
		var ret string
		return ret
	}
	return *o.ApprovalDate
}

// GetApprovalDateOk returns a tuple with the ApprovalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetApprovalDateOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalDate) {
		return nil, false
	}
	return o.ApprovalDate, true
}

// HasApprovalDate returns a boolean if a field has been set.
func (o *BillingChargeType) HasApprovalDate() bool {
	if o != nil && !IsNil(o.ApprovalDate) {
		return true
	}

	return false
}

// SetApprovalDate gets a reference to the given string and assigns it to the ApprovalDate field.
func (o *BillingChargeType) SetApprovalDate(v string) {
	o.ApprovalDate = &v
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *BillingChargeType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingChargeType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *BillingChargeType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *BillingChargeType) SetCashierId(v float32) {
	o.CashierId = &v
}

func (o BillingChargeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingChargeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionNo) {
		toSerialize["transactionNo"] = o.TransactionNo
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.CheckNumber) {
		toSerialize["checkNumber"] = o.CheckNumber
	}
	if !IsNil(o.RevenueDate) {
		toSerialize["revenueDate"] = o.RevenueDate
	}
	if !IsNil(o.Covers) {
		toSerialize["covers"] = o.Covers
	}
	if !IsNil(o.ArrangementCode) {
		toSerialize["arrangementCode"] = o.ArrangementCode
	}
	if !IsNil(o.ApprovalCode) {
		toSerialize["approvalCode"] = o.ApprovalCode
	}
	if !IsNil(o.ApprovalStatus) {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if !IsNil(o.ApprovalDate) {
		toSerialize["approvalDate"] = o.ApprovalDate
	}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	return toSerialize, nil
}

type NullableBillingChargeType struct {
	value *BillingChargeType
	isSet bool
}

func (v NullableBillingChargeType) Get() *BillingChargeType {
	return v.value
}

func (v *NullableBillingChargeType) Set(val *BillingChargeType) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingChargeType) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingChargeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingChargeType(val *BillingChargeType) *NullableBillingChargeType {
	return &NullableBillingChargeType{value: val, isSet: true}
}

func (v NullableBillingChargeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingChargeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


