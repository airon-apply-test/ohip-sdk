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

// checks if the TransactionCurrencyExchangeInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionCurrencyExchangeInfoType{}

// TransactionCurrencyExchangeInfoType Details of the Exchange Information for this transaction, if the posting was made in a currency different from the default currency.
type TransactionCurrencyExchangeInfoType struct {
	// Description or reference for this exchange.
	Description *string `json:"description,omitempty"`
	ExchangeRate *CurrencyAmountType `json:"exchangeRate,omitempty"`
	CurrencyAmount *CurrencyAmountType `json:"currencyAmount,omitempty"`
	Amount *CurrencyAmountType `json:"amount,omitempty"`
	// Percentage for commission used for the currency conversion.
	CommissionPercent *float32 `json:"commissionPercent,omitempty"`
	// Foreign Currency Code of the currency which was used to post this transaction.
	Code *string `json:"code,omitempty"`
	// The date when this exchange rate was set.
	ExchangeDate *string `json:"exchangeDate,omitempty"`
}

// NewTransactionCurrencyExchangeInfoType instantiates a new TransactionCurrencyExchangeInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCurrencyExchangeInfoType() *TransactionCurrencyExchangeInfoType {
	this := TransactionCurrencyExchangeInfoType{}
	return &this
}

// NewTransactionCurrencyExchangeInfoTypeWithDefaults instantiates a new TransactionCurrencyExchangeInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCurrencyExchangeInfoTypeWithDefaults() *TransactionCurrencyExchangeInfoType {
	this := TransactionCurrencyExchangeInfoType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionCurrencyExchangeInfoType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCurrencyExchangeInfoType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionCurrencyExchangeInfoType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionCurrencyExchangeInfoType) SetDescription(v string) {
	o.Description = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *TransactionCurrencyExchangeInfoType) GetExchangeRate() CurrencyAmountType {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCurrencyExchangeInfoType) GetExchangeRateOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *TransactionCurrencyExchangeInfoType) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given CurrencyAmountType and assigns it to the ExchangeRate field.
func (o *TransactionCurrencyExchangeInfoType) SetExchangeRate(v CurrencyAmountType) {
	o.ExchangeRate = &v
}

// GetCurrencyAmount returns the CurrencyAmount field value if set, zero value otherwise.
func (o *TransactionCurrencyExchangeInfoType) GetCurrencyAmount() CurrencyAmountType {
	if o == nil || IsNil(o.CurrencyAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.CurrencyAmount
}

// GetCurrencyAmountOk returns a tuple with the CurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCurrencyExchangeInfoType) GetCurrencyAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.CurrencyAmount) {
		return nil, false
	}
	return o.CurrencyAmount, true
}

// HasCurrencyAmount returns a boolean if a field has been set.
func (o *TransactionCurrencyExchangeInfoType) HasCurrencyAmount() bool {
	if o != nil && !IsNil(o.CurrencyAmount) {
		return true
	}

	return false
}

// SetCurrencyAmount gets a reference to the given CurrencyAmountType and assigns it to the CurrencyAmount field.
func (o *TransactionCurrencyExchangeInfoType) SetCurrencyAmount(v CurrencyAmountType) {
	o.CurrencyAmount = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionCurrencyExchangeInfoType) GetAmount() CurrencyAmountType {
	if o == nil || IsNil(o.Amount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCurrencyExchangeInfoType) GetAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionCurrencyExchangeInfoType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given CurrencyAmountType and assigns it to the Amount field.
func (o *TransactionCurrencyExchangeInfoType) SetAmount(v CurrencyAmountType) {
	o.Amount = &v
}

// GetCommissionPercent returns the CommissionPercent field value if set, zero value otherwise.
func (o *TransactionCurrencyExchangeInfoType) GetCommissionPercent() float32 {
	if o == nil || IsNil(o.CommissionPercent) {
		var ret float32
		return ret
	}
	return *o.CommissionPercent
}

// GetCommissionPercentOk returns a tuple with the CommissionPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCurrencyExchangeInfoType) GetCommissionPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.CommissionPercent) {
		return nil, false
	}
	return o.CommissionPercent, true
}

// HasCommissionPercent returns a boolean if a field has been set.
func (o *TransactionCurrencyExchangeInfoType) HasCommissionPercent() bool {
	if o != nil && !IsNil(o.CommissionPercent) {
		return true
	}

	return false
}

// SetCommissionPercent gets a reference to the given float32 and assigns it to the CommissionPercent field.
func (o *TransactionCurrencyExchangeInfoType) SetCommissionPercent(v float32) {
	o.CommissionPercent = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TransactionCurrencyExchangeInfoType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCurrencyExchangeInfoType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TransactionCurrencyExchangeInfoType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TransactionCurrencyExchangeInfoType) SetCode(v string) {
	o.Code = &v
}

// GetExchangeDate returns the ExchangeDate field value if set, zero value otherwise.
func (o *TransactionCurrencyExchangeInfoType) GetExchangeDate() string {
	if o == nil || IsNil(o.ExchangeDate) {
		var ret string
		return ret
	}
	return *o.ExchangeDate
}

// GetExchangeDateOk returns a tuple with the ExchangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCurrencyExchangeInfoType) GetExchangeDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeDate) {
		return nil, false
	}
	return o.ExchangeDate, true
}

// HasExchangeDate returns a boolean if a field has been set.
func (o *TransactionCurrencyExchangeInfoType) HasExchangeDate() bool {
	if o != nil && !IsNil(o.ExchangeDate) {
		return true
	}

	return false
}

// SetExchangeDate gets a reference to the given string and assigns it to the ExchangeDate field.
func (o *TransactionCurrencyExchangeInfoType) SetExchangeDate(v string) {
	o.ExchangeDate = &v
}

func (o TransactionCurrencyExchangeInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionCurrencyExchangeInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.CurrencyAmount) {
		toSerialize["currencyAmount"] = o.CurrencyAmount
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CommissionPercent) {
		toSerialize["commissionPercent"] = o.CommissionPercent
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ExchangeDate) {
		toSerialize["exchangeDate"] = o.ExchangeDate
	}
	return toSerialize, nil
}

type NullableTransactionCurrencyExchangeInfoType struct {
	value *TransactionCurrencyExchangeInfoType
	isSet bool
}

func (v NullableTransactionCurrencyExchangeInfoType) Get() *TransactionCurrencyExchangeInfoType {
	return v.value
}

func (v *NullableTransactionCurrencyExchangeInfoType) Set(val *TransactionCurrencyExchangeInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCurrencyExchangeInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCurrencyExchangeInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCurrencyExchangeInfoType(val *TransactionCurrencyExchangeInfoType) *NullableTransactionCurrencyExchangeInfoType {
	return &NullableTransactionCurrencyExchangeInfoType{value: val, isSet: true}
}

func (v NullableTransactionCurrencyExchangeInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCurrencyExchangeInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


