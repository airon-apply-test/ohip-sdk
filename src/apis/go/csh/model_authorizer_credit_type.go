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

// checks if the AuthorizerCreditType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizerCreditType{}

// AuthorizerCreditType Authorizer Information
type AuthorizerCreditType struct {
	AuthorizerId *UniqueIDType `json:"authorizerId,omitempty"`
	// Application user name of the authorizer
	AuthorizerUserName *string `json:"authorizerUserName,omitempty"`
	// Full name of the authorizer.
	AuthorizerName *string `json:"authorizerName,omitempty"`
	// Rate code of the authorizer.
	AuthorizerRateCode *string `json:"authorizerRateCode,omitempty"`
	// Indicates whether user has the choice to have reservation inherit rate code from the authorizer.
	InheritAuthorizerRateCode *bool `json:"inheritAuthorizerRateCode,omitempty"`
	// Identifies the hotel code.
	HotelId *string `json:"hotelId,omitempty"`
	CreditLimit *CurrencyAmountType `json:"creditLimit,omitempty"`
	ActualAmount *CurrencyAmountType `json:"actualAmount,omitempty"`
	// List of Comp Accounting Authorizers details
	AuthorizerCreditDetails []AuthorizerCreditDetailType `json:"authorizerCreditDetails,omitempty"`
	// Transaction Date associated with the transaction.
	TransactionDate *string `json:"transactionDate,omitempty"`
}

// NewAuthorizerCreditType instantiates a new AuthorizerCreditType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizerCreditType() *AuthorizerCreditType {
	this := AuthorizerCreditType{}
	return &this
}

// NewAuthorizerCreditTypeWithDefaults instantiates a new AuthorizerCreditType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizerCreditTypeWithDefaults() *AuthorizerCreditType {
	this := AuthorizerCreditType{}
	return &this
}

// GetAuthorizerId returns the AuthorizerId field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetAuthorizerId() UniqueIDType {
	if o == nil || IsNil(o.AuthorizerId) {
		var ret UniqueIDType
		return ret
	}
	return *o.AuthorizerId
}

// GetAuthorizerIdOk returns a tuple with the AuthorizerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetAuthorizerIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.AuthorizerId) {
		return nil, false
	}
	return o.AuthorizerId, true
}

// HasAuthorizerId returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasAuthorizerId() bool {
	if o != nil && !IsNil(o.AuthorizerId) {
		return true
	}

	return false
}

// SetAuthorizerId gets a reference to the given UniqueIDType and assigns it to the AuthorizerId field.
func (o *AuthorizerCreditType) SetAuthorizerId(v UniqueIDType) {
	o.AuthorizerId = &v
}

// GetAuthorizerUserName returns the AuthorizerUserName field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetAuthorizerUserName() string {
	if o == nil || IsNil(o.AuthorizerUserName) {
		var ret string
		return ret
	}
	return *o.AuthorizerUserName
}

// GetAuthorizerUserNameOk returns a tuple with the AuthorizerUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetAuthorizerUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizerUserName) {
		return nil, false
	}
	return o.AuthorizerUserName, true
}

// HasAuthorizerUserName returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasAuthorizerUserName() bool {
	if o != nil && !IsNil(o.AuthorizerUserName) {
		return true
	}

	return false
}

// SetAuthorizerUserName gets a reference to the given string and assigns it to the AuthorizerUserName field.
func (o *AuthorizerCreditType) SetAuthorizerUserName(v string) {
	o.AuthorizerUserName = &v
}

// GetAuthorizerName returns the AuthorizerName field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetAuthorizerName() string {
	if o == nil || IsNil(o.AuthorizerName) {
		var ret string
		return ret
	}
	return *o.AuthorizerName
}

// GetAuthorizerNameOk returns a tuple with the AuthorizerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetAuthorizerNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizerName) {
		return nil, false
	}
	return o.AuthorizerName, true
}

// HasAuthorizerName returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasAuthorizerName() bool {
	if o != nil && !IsNil(o.AuthorizerName) {
		return true
	}

	return false
}

// SetAuthorizerName gets a reference to the given string and assigns it to the AuthorizerName field.
func (o *AuthorizerCreditType) SetAuthorizerName(v string) {
	o.AuthorizerName = &v
}

// GetAuthorizerRateCode returns the AuthorizerRateCode field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetAuthorizerRateCode() string {
	if o == nil || IsNil(o.AuthorizerRateCode) {
		var ret string
		return ret
	}
	return *o.AuthorizerRateCode
}

// GetAuthorizerRateCodeOk returns a tuple with the AuthorizerRateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetAuthorizerRateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizerRateCode) {
		return nil, false
	}
	return o.AuthorizerRateCode, true
}

// HasAuthorizerRateCode returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasAuthorizerRateCode() bool {
	if o != nil && !IsNil(o.AuthorizerRateCode) {
		return true
	}

	return false
}

// SetAuthorizerRateCode gets a reference to the given string and assigns it to the AuthorizerRateCode field.
func (o *AuthorizerCreditType) SetAuthorizerRateCode(v string) {
	o.AuthorizerRateCode = &v
}

// GetInheritAuthorizerRateCode returns the InheritAuthorizerRateCode field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetInheritAuthorizerRateCode() bool {
	if o == nil || IsNil(o.InheritAuthorizerRateCode) {
		var ret bool
		return ret
	}
	return *o.InheritAuthorizerRateCode
}

// GetInheritAuthorizerRateCodeOk returns a tuple with the InheritAuthorizerRateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetInheritAuthorizerRateCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.InheritAuthorizerRateCode) {
		return nil, false
	}
	return o.InheritAuthorizerRateCode, true
}

// HasInheritAuthorizerRateCode returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasInheritAuthorizerRateCode() bool {
	if o != nil && !IsNil(o.InheritAuthorizerRateCode) {
		return true
	}

	return false
}

// SetInheritAuthorizerRateCode gets a reference to the given bool and assigns it to the InheritAuthorizerRateCode field.
func (o *AuthorizerCreditType) SetInheritAuthorizerRateCode(v bool) {
	o.InheritAuthorizerRateCode = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *AuthorizerCreditType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCreditLimit returns the CreditLimit field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetCreditLimit() CurrencyAmountType {
	if o == nil || IsNil(o.CreditLimit) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.CreditLimit
}

// GetCreditLimitOk returns a tuple with the CreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetCreditLimitOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.CreditLimit) {
		return nil, false
	}
	return o.CreditLimit, true
}

// HasCreditLimit returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasCreditLimit() bool {
	if o != nil && !IsNil(o.CreditLimit) {
		return true
	}

	return false
}

// SetCreditLimit gets a reference to the given CurrencyAmountType and assigns it to the CreditLimit field.
func (o *AuthorizerCreditType) SetCreditLimit(v CurrencyAmountType) {
	o.CreditLimit = &v
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetActualAmount() CurrencyAmountType {
	if o == nil || IsNil(o.ActualAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.ActualAmount
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetActualAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.ActualAmount) {
		return nil, false
	}
	return o.ActualAmount, true
}

// HasActualAmount returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasActualAmount() bool {
	if o != nil && !IsNil(o.ActualAmount) {
		return true
	}

	return false
}

// SetActualAmount gets a reference to the given CurrencyAmountType and assigns it to the ActualAmount field.
func (o *AuthorizerCreditType) SetActualAmount(v CurrencyAmountType) {
	o.ActualAmount = &v
}

// GetAuthorizerCreditDetails returns the AuthorizerCreditDetails field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetAuthorizerCreditDetails() []AuthorizerCreditDetailType {
	if o == nil || IsNil(o.AuthorizerCreditDetails) {
		var ret []AuthorizerCreditDetailType
		return ret
	}
	return o.AuthorizerCreditDetails
}

// GetAuthorizerCreditDetailsOk returns a tuple with the AuthorizerCreditDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetAuthorizerCreditDetailsOk() ([]AuthorizerCreditDetailType, bool) {
	if o == nil || IsNil(o.AuthorizerCreditDetails) {
		return nil, false
	}
	return o.AuthorizerCreditDetails, true
}

// HasAuthorizerCreditDetails returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasAuthorizerCreditDetails() bool {
	if o != nil && !IsNil(o.AuthorizerCreditDetails) {
		return true
	}

	return false
}

// SetAuthorizerCreditDetails gets a reference to the given []AuthorizerCreditDetailType and assigns it to the AuthorizerCreditDetails field.
func (o *AuthorizerCreditType) SetAuthorizerCreditDetails(v []AuthorizerCreditDetailType) {
	o.AuthorizerCreditDetails = v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *AuthorizerCreditType) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditType) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *AuthorizerCreditType) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *AuthorizerCreditType) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

func (o AuthorizerCreditType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizerCreditType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizerId) {
		toSerialize["authorizerId"] = o.AuthorizerId
	}
	if !IsNil(o.AuthorizerUserName) {
		toSerialize["authorizerUserName"] = o.AuthorizerUserName
	}
	if !IsNil(o.AuthorizerName) {
		toSerialize["authorizerName"] = o.AuthorizerName
	}
	if !IsNil(o.AuthorizerRateCode) {
		toSerialize["authorizerRateCode"] = o.AuthorizerRateCode
	}
	if !IsNil(o.InheritAuthorizerRateCode) {
		toSerialize["inheritAuthorizerRateCode"] = o.InheritAuthorizerRateCode
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.CreditLimit) {
		toSerialize["creditLimit"] = o.CreditLimit
	}
	if !IsNil(o.ActualAmount) {
		toSerialize["actualAmount"] = o.ActualAmount
	}
	if !IsNil(o.AuthorizerCreditDetails) {
		toSerialize["authorizerCreditDetails"] = o.AuthorizerCreditDetails
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	return toSerialize, nil
}

type NullableAuthorizerCreditType struct {
	value *AuthorizerCreditType
	isSet bool
}

func (v NullableAuthorizerCreditType) Get() *AuthorizerCreditType {
	return v.value
}

func (v *NullableAuthorizerCreditType) Set(val *AuthorizerCreditType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizerCreditType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizerCreditType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizerCreditType(val *AuthorizerCreditType) *NullableAuthorizerCreditType {
	return &NullableAuthorizerCreditType{value: val, isSet: true}
}

func (v NullableAuthorizerCreditType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizerCreditType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


