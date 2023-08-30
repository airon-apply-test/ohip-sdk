/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
)

// checks if the ReservationPaymentMethodType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservationPaymentMethodType{}

// ReservationPaymentMethodType struct for ReservationPaymentMethodType
type ReservationPaymentMethodType struct {
	PaymentCard *ResPaymentCardType `json:"paymentCard,omitempty"`
	Balance *CurrencyAmountType `json:"balance,omitempty"`
	AuthorizationRule *AuthorizationRuleType `json:"authorizationRule,omitempty"`
	EmailFolioInfo *ReservationPaymentMethodTypeEmailFolioInfo `json:"emailFolioInfo,omitempty"`
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	Description *string `json:"description,omitempty"`
	FolioView *int32 `json:"folioView,omitempty"`
}

// NewReservationPaymentMethodType instantiates a new ReservationPaymentMethodType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationPaymentMethodType() *ReservationPaymentMethodType {
	this := ReservationPaymentMethodType{}
	return &this
}

// NewReservationPaymentMethodTypeWithDefaults instantiates a new ReservationPaymentMethodType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationPaymentMethodTypeWithDefaults() *ReservationPaymentMethodType {
	this := ReservationPaymentMethodType{}
	return &this
}

// GetPaymentCard returns the PaymentCard field value if set, zero value otherwise.
func (o *ReservationPaymentMethodType) GetPaymentCard() ResPaymentCardType {
	if o == nil || IsNil(o.PaymentCard) {
		var ret ResPaymentCardType
		return ret
	}
	return *o.PaymentCard
}

// GetPaymentCardOk returns a tuple with the PaymentCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPaymentMethodType) GetPaymentCardOk() (*ResPaymentCardType, bool) {
	if o == nil || IsNil(o.PaymentCard) {
		return nil, false
	}
	return o.PaymentCard, true
}

// HasPaymentCard returns a boolean if a field has been set.
func (o *ReservationPaymentMethodType) HasPaymentCard() bool {
	if o != nil && !IsNil(o.PaymentCard) {
		return true
	}

	return false
}

// SetPaymentCard gets a reference to the given ResPaymentCardType and assigns it to the PaymentCard field.
func (o *ReservationPaymentMethodType) SetPaymentCard(v ResPaymentCardType) {
	o.PaymentCard = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *ReservationPaymentMethodType) GetBalance() CurrencyAmountType {
	if o == nil || IsNil(o.Balance) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPaymentMethodType) GetBalanceOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *ReservationPaymentMethodType) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given CurrencyAmountType and assigns it to the Balance field.
func (o *ReservationPaymentMethodType) SetBalance(v CurrencyAmountType) {
	o.Balance = &v
}

// GetAuthorizationRule returns the AuthorizationRule field value if set, zero value otherwise.
func (o *ReservationPaymentMethodType) GetAuthorizationRule() AuthorizationRuleType {
	if o == nil || IsNil(o.AuthorizationRule) {
		var ret AuthorizationRuleType
		return ret
	}
	return *o.AuthorizationRule
}

// GetAuthorizationRuleOk returns a tuple with the AuthorizationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPaymentMethodType) GetAuthorizationRuleOk() (*AuthorizationRuleType, bool) {
	if o == nil || IsNil(o.AuthorizationRule) {
		return nil, false
	}
	return o.AuthorizationRule, true
}

// HasAuthorizationRule returns a boolean if a field has been set.
func (o *ReservationPaymentMethodType) HasAuthorizationRule() bool {
	if o != nil && !IsNil(o.AuthorizationRule) {
		return true
	}

	return false
}

// SetAuthorizationRule gets a reference to the given AuthorizationRuleType and assigns it to the AuthorizationRule field.
func (o *ReservationPaymentMethodType) SetAuthorizationRule(v AuthorizationRuleType) {
	o.AuthorizationRule = &v
}

// GetEmailFolioInfo returns the EmailFolioInfo field value if set, zero value otherwise.
func (o *ReservationPaymentMethodType) GetEmailFolioInfo() ReservationPaymentMethodTypeEmailFolioInfo {
	if o == nil || IsNil(o.EmailFolioInfo) {
		var ret ReservationPaymentMethodTypeEmailFolioInfo
		return ret
	}
	return *o.EmailFolioInfo
}

// GetEmailFolioInfoOk returns a tuple with the EmailFolioInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPaymentMethodType) GetEmailFolioInfoOk() (*ReservationPaymentMethodTypeEmailFolioInfo, bool) {
	if o == nil || IsNil(o.EmailFolioInfo) {
		return nil, false
	}
	return o.EmailFolioInfo, true
}

// HasEmailFolioInfo returns a boolean if a field has been set.
func (o *ReservationPaymentMethodType) HasEmailFolioInfo() bool {
	if o != nil && !IsNil(o.EmailFolioInfo) {
		return true
	}

	return false
}

// SetEmailFolioInfo gets a reference to the given ReservationPaymentMethodTypeEmailFolioInfo and assigns it to the EmailFolioInfo field.
func (o *ReservationPaymentMethodType) SetEmailFolioInfo(v ReservationPaymentMethodTypeEmailFolioInfo) {
	o.EmailFolioInfo = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ReservationPaymentMethodType) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPaymentMethodType) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ReservationPaymentMethodType) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *ReservationPaymentMethodType) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReservationPaymentMethodType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPaymentMethodType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReservationPaymentMethodType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReservationPaymentMethodType) SetDescription(v string) {
	o.Description = &v
}

// GetFolioView returns the FolioView field value if set, zero value otherwise.
func (o *ReservationPaymentMethodType) GetFolioView() int32 {
	if o == nil || IsNil(o.FolioView) {
		var ret int32
		return ret
	}
	return *o.FolioView
}

// GetFolioViewOk returns a tuple with the FolioView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservationPaymentMethodType) GetFolioViewOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioView) {
		return nil, false
	}
	return o.FolioView, true
}

// HasFolioView returns a boolean if a field has been set.
func (o *ReservationPaymentMethodType) HasFolioView() bool {
	if o != nil && !IsNil(o.FolioView) {
		return true
	}

	return false
}

// SetFolioView gets a reference to the given int32 and assigns it to the FolioView field.
func (o *ReservationPaymentMethodType) SetFolioView(v int32) {
	o.FolioView = &v
}

func (o ReservationPaymentMethodType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservationPaymentMethodType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentCard) {
		toSerialize["paymentCard"] = o.PaymentCard
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.AuthorizationRule) {
		toSerialize["authorizationRule"] = o.AuthorizationRule
	}
	if !IsNil(o.EmailFolioInfo) {
		toSerialize["emailFolioInfo"] = o.EmailFolioInfo
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FolioView) {
		toSerialize["folioView"] = o.FolioView
	}
	return toSerialize, nil
}

type NullableReservationPaymentMethodType struct {
	value *ReservationPaymentMethodType
	isSet bool
}

func (v NullableReservationPaymentMethodType) Get() *ReservationPaymentMethodType {
	return v.value
}

func (v *NullableReservationPaymentMethodType) Set(val *ReservationPaymentMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationPaymentMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationPaymentMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationPaymentMethodType(val *ReservationPaymentMethodType) *NullableReservationPaymentMethodType {
	return &NullableReservationPaymentMethodType{value: val, isSet: true}
}

func (v NullableReservationPaymentMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationPaymentMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


