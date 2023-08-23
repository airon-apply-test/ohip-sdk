/*
OPI Token Exchange Service API

Oracle Token Exchange Service Specification<br /><br /> Compatible with OPERA Cloud release 1.0.1.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 1.0.1
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OpenPaymentTokenExchange200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenPaymentTokenExchange200Response{}

// OpenPaymentTokenExchange200Response TokenResponse object
type OpenPaymentTokenExchange200Response struct {
	// OPERA Card Type
	CardType string `json:"cardType"`
	// Expiration Date, YYMM format
	ExpiryDate string `json:"expiryDate"`
	// Masked Primary Account Number (PAN)
	Pan string `json:"pan"`
	// Card Token
	Token string `json:"token"`
}

// NewOpenPaymentTokenExchange200Response instantiates a new OpenPaymentTokenExchange200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenPaymentTokenExchange200Response(cardType string, expiryDate string, pan string, token string) *OpenPaymentTokenExchange200Response {
	this := OpenPaymentTokenExchange200Response{}
	this.CardType = cardType
	this.ExpiryDate = expiryDate
	this.Pan = pan
	this.Token = token
	return &this
}

// NewOpenPaymentTokenExchange200ResponseWithDefaults instantiates a new OpenPaymentTokenExchange200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenPaymentTokenExchange200ResponseWithDefaults() *OpenPaymentTokenExchange200Response {
	this := OpenPaymentTokenExchange200Response{}
	return &this
}

// GetCardType returns the CardType field value
func (o *OpenPaymentTokenExchange200Response) GetCardType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange200Response) GetCardTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardType, true
}

// SetCardType sets field value
func (o *OpenPaymentTokenExchange200Response) SetCardType(v string) {
	o.CardType = v
}

// GetExpiryDate returns the ExpiryDate field value
func (o *OpenPaymentTokenExchange200Response) GetExpiryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange200Response) GetExpiryDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryDate, true
}

// SetExpiryDate sets field value
func (o *OpenPaymentTokenExchange200Response) SetExpiryDate(v string) {
	o.ExpiryDate = v
}

// GetPan returns the Pan field value
func (o *OpenPaymentTokenExchange200Response) GetPan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pan
}

// GetPanOk returns a tuple with the Pan field value
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange200Response) GetPanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pan, true
}

// SetPan sets field value
func (o *OpenPaymentTokenExchange200Response) SetPan(v string) {
	o.Pan = v
}

// GetToken returns the Token field value
func (o *OpenPaymentTokenExchange200Response) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OpenPaymentTokenExchange200Response) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *OpenPaymentTokenExchange200Response) SetToken(v string) {
	o.Token = v
}

func (o OpenPaymentTokenExchange200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenPaymentTokenExchange200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cardType"] = o.CardType
	toSerialize["expiryDate"] = o.ExpiryDate
	toSerialize["pan"] = o.Pan
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableOpenPaymentTokenExchange200Response struct {
	value *OpenPaymentTokenExchange200Response
	isSet bool
}

func (v NullableOpenPaymentTokenExchange200Response) Get() *OpenPaymentTokenExchange200Response {
	return v.value
}

func (v *NullableOpenPaymentTokenExchange200Response) Set(val *OpenPaymentTokenExchange200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenPaymentTokenExchange200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenPaymentTokenExchange200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenPaymentTokenExchange200Response(val *OpenPaymentTokenExchange200Response) *NullableOpenPaymentTokenExchange200Response {
	return &NullableOpenPaymentTokenExchange200Response{value: val, isSet: true}
}

func (v NullableOpenPaymentTokenExchange200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenPaymentTokenExchange200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


