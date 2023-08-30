/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the MemberAwardDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberAwardDetailType{}

// MemberAwardDetailType Details related to member award like award type, stay date, rate code, etc.
type MemberAwardDetailType struct {
	// The award type or code.
	AwardType *string `json:"awardType,omitempty"`
	// Indicates if award is a Stay, Package Element or other.
	AwardBasedOn *string `json:"awardBasedOn,omitempty"`
	// Number days before arrival to apply penalty for cancellation.
	CancelPenaltyDays *int32 `json:"cancelPenaltyDays,omitempty"`
	// Number points deducted if award is cancelled.
	CancelPenaltyCharge *int32 `json:"cancelPenaltyCharge,omitempty"`
	// Type of cancel penalty like Points, etc.
	CancelPenaltyType *string `json:"cancelPenaltyType,omitempty"`
	// Number of penalty points if cancelled.
	CancelPenaltyPoints *int32 `json:"cancelPenaltyPoints,omitempty"`
	// Date of stay.
	StayDate *string `json:"stayDate,omitempty"`
	// Rate code associated with the reservation.
	RateCode *string `json:"rateCode,omitempty"`
	// Room type label associated with the reservation.
	RoomType *string `json:"roomType,omitempty"`
	// Product code for which the award was issued, in case of a product award.
	Product *string `json:"product,omitempty"`
	// Room type label before the upgrade in case of an upgrade award.
	FromRoomType *string `json:"fromRoomType,omitempty"`
	// Room type label after the upgrade for an upgrade award.
	ToRoomType *string `json:"toRoomType,omitempty"`
	// Total Local Amount on bill in Hotel Currency.
	TotalLocalAmount *float32 `json:"totalLocalAmount,omitempty"`
	// Redeemed Local Amount on bill in Hotel Currency.
	RedeemedLocalAmount *float32 `json:"redeemedLocalAmount,omitempty"`
	// Total Central Amount on bill in External System Currency.
	TotalCentralAmount *float32 `json:"totalCentralAmount,omitempty"`
	// Redeemed Central Amount on bill in External System Currency.
	RedeemedCentralAmount *float32 `json:"redeemedCentralAmount,omitempty"`
	// The Payment Transaction Code for which the Surcharge Applies.
	TransactionCode *string `json:"transactionCode,omitempty"`
	// Unique Transaction Identifier.
	TransactionNo *float32 `json:"transactionNo,omitempty"`
	// Exchange Rate Type for the Currency Exchange.
	ExchangeRateType *string `json:"exchangeRateType,omitempty"`
	// Award Voucher Number.
	AwardVoucherNo *string `json:"awardVoucherNo,omitempty"`
	AwardCancellationNo *UniqueIDType `json:"awardCancellationNo,omitempty"`
	// If the award detail is inactive.
	Inactive *bool `json:"inactive,omitempty"`
	// Points required for the stay date.
	PointsRequired *int32 `json:"pointsRequired,omitempty"`
}

// NewMemberAwardDetailType instantiates a new MemberAwardDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberAwardDetailType() *MemberAwardDetailType {
	this := MemberAwardDetailType{}
	return &this
}

// NewMemberAwardDetailTypeWithDefaults instantiates a new MemberAwardDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberAwardDetailTypeWithDefaults() *MemberAwardDetailType {
	this := MemberAwardDetailType{}
	return &this
}

// GetAwardType returns the AwardType field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetAwardType() string {
	if o == nil || IsNil(o.AwardType) {
		var ret string
		return ret
	}
	return *o.AwardType
}

// GetAwardTypeOk returns a tuple with the AwardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetAwardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AwardType) {
		return nil, false
	}
	return o.AwardType, true
}

// HasAwardType returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasAwardType() bool {
	if o != nil && !IsNil(o.AwardType) {
		return true
	}

	return false
}

// SetAwardType gets a reference to the given string and assigns it to the AwardType field.
func (o *MemberAwardDetailType) SetAwardType(v string) {
	o.AwardType = &v
}

// GetAwardBasedOn returns the AwardBasedOn field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetAwardBasedOn() string {
	if o == nil || IsNil(o.AwardBasedOn) {
		var ret string
		return ret
	}
	return *o.AwardBasedOn
}

// GetAwardBasedOnOk returns a tuple with the AwardBasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetAwardBasedOnOk() (*string, bool) {
	if o == nil || IsNil(o.AwardBasedOn) {
		return nil, false
	}
	return o.AwardBasedOn, true
}

// HasAwardBasedOn returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasAwardBasedOn() bool {
	if o != nil && !IsNil(o.AwardBasedOn) {
		return true
	}

	return false
}

// SetAwardBasedOn gets a reference to the given string and assigns it to the AwardBasedOn field.
func (o *MemberAwardDetailType) SetAwardBasedOn(v string) {
	o.AwardBasedOn = &v
}

// GetCancelPenaltyDays returns the CancelPenaltyDays field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetCancelPenaltyDays() int32 {
	if o == nil || IsNil(o.CancelPenaltyDays) {
		var ret int32
		return ret
	}
	return *o.CancelPenaltyDays
}

// GetCancelPenaltyDaysOk returns a tuple with the CancelPenaltyDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetCancelPenaltyDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.CancelPenaltyDays) {
		return nil, false
	}
	return o.CancelPenaltyDays, true
}

// HasCancelPenaltyDays returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasCancelPenaltyDays() bool {
	if o != nil && !IsNil(o.CancelPenaltyDays) {
		return true
	}

	return false
}

// SetCancelPenaltyDays gets a reference to the given int32 and assigns it to the CancelPenaltyDays field.
func (o *MemberAwardDetailType) SetCancelPenaltyDays(v int32) {
	o.CancelPenaltyDays = &v
}

// GetCancelPenaltyCharge returns the CancelPenaltyCharge field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetCancelPenaltyCharge() int32 {
	if o == nil || IsNil(o.CancelPenaltyCharge) {
		var ret int32
		return ret
	}
	return *o.CancelPenaltyCharge
}

// GetCancelPenaltyChargeOk returns a tuple with the CancelPenaltyCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetCancelPenaltyChargeOk() (*int32, bool) {
	if o == nil || IsNil(o.CancelPenaltyCharge) {
		return nil, false
	}
	return o.CancelPenaltyCharge, true
}

// HasCancelPenaltyCharge returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasCancelPenaltyCharge() bool {
	if o != nil && !IsNil(o.CancelPenaltyCharge) {
		return true
	}

	return false
}

// SetCancelPenaltyCharge gets a reference to the given int32 and assigns it to the CancelPenaltyCharge field.
func (o *MemberAwardDetailType) SetCancelPenaltyCharge(v int32) {
	o.CancelPenaltyCharge = &v
}

// GetCancelPenaltyType returns the CancelPenaltyType field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetCancelPenaltyType() string {
	if o == nil || IsNil(o.CancelPenaltyType) {
		var ret string
		return ret
	}
	return *o.CancelPenaltyType
}

// GetCancelPenaltyTypeOk returns a tuple with the CancelPenaltyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetCancelPenaltyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CancelPenaltyType) {
		return nil, false
	}
	return o.CancelPenaltyType, true
}

// HasCancelPenaltyType returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasCancelPenaltyType() bool {
	if o != nil && !IsNil(o.CancelPenaltyType) {
		return true
	}

	return false
}

// SetCancelPenaltyType gets a reference to the given string and assigns it to the CancelPenaltyType field.
func (o *MemberAwardDetailType) SetCancelPenaltyType(v string) {
	o.CancelPenaltyType = &v
}

// GetCancelPenaltyPoints returns the CancelPenaltyPoints field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetCancelPenaltyPoints() int32 {
	if o == nil || IsNil(o.CancelPenaltyPoints) {
		var ret int32
		return ret
	}
	return *o.CancelPenaltyPoints
}

// GetCancelPenaltyPointsOk returns a tuple with the CancelPenaltyPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetCancelPenaltyPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.CancelPenaltyPoints) {
		return nil, false
	}
	return o.CancelPenaltyPoints, true
}

// HasCancelPenaltyPoints returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasCancelPenaltyPoints() bool {
	if o != nil && !IsNil(o.CancelPenaltyPoints) {
		return true
	}

	return false
}

// SetCancelPenaltyPoints gets a reference to the given int32 and assigns it to the CancelPenaltyPoints field.
func (o *MemberAwardDetailType) SetCancelPenaltyPoints(v int32) {
	o.CancelPenaltyPoints = &v
}

// GetStayDate returns the StayDate field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetStayDate() string {
	if o == nil || IsNil(o.StayDate) {
		var ret string
		return ret
	}
	return *o.StayDate
}

// GetStayDateOk returns a tuple with the StayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetStayDateOk() (*string, bool) {
	if o == nil || IsNil(o.StayDate) {
		return nil, false
	}
	return o.StayDate, true
}

// HasStayDate returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasStayDate() bool {
	if o != nil && !IsNil(o.StayDate) {
		return true
	}

	return false
}

// SetStayDate gets a reference to the given string and assigns it to the StayDate field.
func (o *MemberAwardDetailType) SetStayDate(v string) {
	o.StayDate = &v
}

// GetRateCode returns the RateCode field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetRateCode() string {
	if o == nil || IsNil(o.RateCode) {
		var ret string
		return ret
	}
	return *o.RateCode
}

// GetRateCodeOk returns a tuple with the RateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetRateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RateCode) {
		return nil, false
	}
	return o.RateCode, true
}

// HasRateCode returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasRateCode() bool {
	if o != nil && !IsNil(o.RateCode) {
		return true
	}

	return false
}

// SetRateCode gets a reference to the given string and assigns it to the RateCode field.
func (o *MemberAwardDetailType) SetRateCode(v string) {
	o.RateCode = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *MemberAwardDetailType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *MemberAwardDetailType) SetProduct(v string) {
	o.Product = &v
}

// GetFromRoomType returns the FromRoomType field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetFromRoomType() string {
	if o == nil || IsNil(o.FromRoomType) {
		var ret string
		return ret
	}
	return *o.FromRoomType
}

// GetFromRoomTypeOk returns a tuple with the FromRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetFromRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FromRoomType) {
		return nil, false
	}
	return o.FromRoomType, true
}

// HasFromRoomType returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasFromRoomType() bool {
	if o != nil && !IsNil(o.FromRoomType) {
		return true
	}

	return false
}

// SetFromRoomType gets a reference to the given string and assigns it to the FromRoomType field.
func (o *MemberAwardDetailType) SetFromRoomType(v string) {
	o.FromRoomType = &v
}

// GetToRoomType returns the ToRoomType field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetToRoomType() string {
	if o == nil || IsNil(o.ToRoomType) {
		var ret string
		return ret
	}
	return *o.ToRoomType
}

// GetToRoomTypeOk returns a tuple with the ToRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetToRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ToRoomType) {
		return nil, false
	}
	return o.ToRoomType, true
}

// HasToRoomType returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasToRoomType() bool {
	if o != nil && !IsNil(o.ToRoomType) {
		return true
	}

	return false
}

// SetToRoomType gets a reference to the given string and assigns it to the ToRoomType field.
func (o *MemberAwardDetailType) SetToRoomType(v string) {
	o.ToRoomType = &v
}

// GetTotalLocalAmount returns the TotalLocalAmount field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetTotalLocalAmount() float32 {
	if o == nil || IsNil(o.TotalLocalAmount) {
		var ret float32
		return ret
	}
	return *o.TotalLocalAmount
}

// GetTotalLocalAmountOk returns a tuple with the TotalLocalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetTotalLocalAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalLocalAmount) {
		return nil, false
	}
	return o.TotalLocalAmount, true
}

// HasTotalLocalAmount returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasTotalLocalAmount() bool {
	if o != nil && !IsNil(o.TotalLocalAmount) {
		return true
	}

	return false
}

// SetTotalLocalAmount gets a reference to the given float32 and assigns it to the TotalLocalAmount field.
func (o *MemberAwardDetailType) SetTotalLocalAmount(v float32) {
	o.TotalLocalAmount = &v
}

// GetRedeemedLocalAmount returns the RedeemedLocalAmount field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetRedeemedLocalAmount() float32 {
	if o == nil || IsNil(o.RedeemedLocalAmount) {
		var ret float32
		return ret
	}
	return *o.RedeemedLocalAmount
}

// GetRedeemedLocalAmountOk returns a tuple with the RedeemedLocalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetRedeemedLocalAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RedeemedLocalAmount) {
		return nil, false
	}
	return o.RedeemedLocalAmount, true
}

// HasRedeemedLocalAmount returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasRedeemedLocalAmount() bool {
	if o != nil && !IsNil(o.RedeemedLocalAmount) {
		return true
	}

	return false
}

// SetRedeemedLocalAmount gets a reference to the given float32 and assigns it to the RedeemedLocalAmount field.
func (o *MemberAwardDetailType) SetRedeemedLocalAmount(v float32) {
	o.RedeemedLocalAmount = &v
}

// GetTotalCentralAmount returns the TotalCentralAmount field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetTotalCentralAmount() float32 {
	if o == nil || IsNil(o.TotalCentralAmount) {
		var ret float32
		return ret
	}
	return *o.TotalCentralAmount
}

// GetTotalCentralAmountOk returns a tuple with the TotalCentralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetTotalCentralAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCentralAmount) {
		return nil, false
	}
	return o.TotalCentralAmount, true
}

// HasTotalCentralAmount returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasTotalCentralAmount() bool {
	if o != nil && !IsNil(o.TotalCentralAmount) {
		return true
	}

	return false
}

// SetTotalCentralAmount gets a reference to the given float32 and assigns it to the TotalCentralAmount field.
func (o *MemberAwardDetailType) SetTotalCentralAmount(v float32) {
	o.TotalCentralAmount = &v
}

// GetRedeemedCentralAmount returns the RedeemedCentralAmount field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetRedeemedCentralAmount() float32 {
	if o == nil || IsNil(o.RedeemedCentralAmount) {
		var ret float32
		return ret
	}
	return *o.RedeemedCentralAmount
}

// GetRedeemedCentralAmountOk returns a tuple with the RedeemedCentralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetRedeemedCentralAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RedeemedCentralAmount) {
		return nil, false
	}
	return o.RedeemedCentralAmount, true
}

// HasRedeemedCentralAmount returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasRedeemedCentralAmount() bool {
	if o != nil && !IsNil(o.RedeemedCentralAmount) {
		return true
	}

	return false
}

// SetRedeemedCentralAmount gets a reference to the given float32 and assigns it to the RedeemedCentralAmount field.
func (o *MemberAwardDetailType) SetRedeemedCentralAmount(v float32) {
	o.RedeemedCentralAmount = &v
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode) {
		var ret string
		return ret
	}
	return *o.TransactionCode
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetTransactionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionCode) {
		return nil, false
	}
	return o.TransactionCode, true
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasTransactionCode() bool {
	if o != nil && !IsNil(o.TransactionCode) {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given string and assigns it to the TransactionCode field.
func (o *MemberAwardDetailType) SetTransactionCode(v string) {
	o.TransactionCode = &v
}

// GetTransactionNo returns the TransactionNo field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetTransactionNo() float32 {
	if o == nil || IsNil(o.TransactionNo) {
		var ret float32
		return ret
	}
	return *o.TransactionNo
}

// GetTransactionNoOk returns a tuple with the TransactionNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetTransactionNoOk() (*float32, bool) {
	if o == nil || IsNil(o.TransactionNo) {
		return nil, false
	}
	return o.TransactionNo, true
}

// HasTransactionNo returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasTransactionNo() bool {
	if o != nil && !IsNil(o.TransactionNo) {
		return true
	}

	return false
}

// SetTransactionNo gets a reference to the given float32 and assigns it to the TransactionNo field.
func (o *MemberAwardDetailType) SetTransactionNo(v float32) {
	o.TransactionNo = &v
}

// GetExchangeRateType returns the ExchangeRateType field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetExchangeRateType() string {
	if o == nil || IsNil(o.ExchangeRateType) {
		var ret string
		return ret
	}
	return *o.ExchangeRateType
}

// GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetExchangeRateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeRateType) {
		return nil, false
	}
	return o.ExchangeRateType, true
}

// HasExchangeRateType returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasExchangeRateType() bool {
	if o != nil && !IsNil(o.ExchangeRateType) {
		return true
	}

	return false
}

// SetExchangeRateType gets a reference to the given string and assigns it to the ExchangeRateType field.
func (o *MemberAwardDetailType) SetExchangeRateType(v string) {
	o.ExchangeRateType = &v
}

// GetAwardVoucherNo returns the AwardVoucherNo field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetAwardVoucherNo() string {
	if o == nil || IsNil(o.AwardVoucherNo) {
		var ret string
		return ret
	}
	return *o.AwardVoucherNo
}

// GetAwardVoucherNoOk returns a tuple with the AwardVoucherNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetAwardVoucherNoOk() (*string, bool) {
	if o == nil || IsNil(o.AwardVoucherNo) {
		return nil, false
	}
	return o.AwardVoucherNo, true
}

// HasAwardVoucherNo returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasAwardVoucherNo() bool {
	if o != nil && !IsNil(o.AwardVoucherNo) {
		return true
	}

	return false
}

// SetAwardVoucherNo gets a reference to the given string and assigns it to the AwardVoucherNo field.
func (o *MemberAwardDetailType) SetAwardVoucherNo(v string) {
	o.AwardVoucherNo = &v
}

// GetAwardCancellationNo returns the AwardCancellationNo field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetAwardCancellationNo() UniqueIDType {
	if o == nil || IsNil(o.AwardCancellationNo) {
		var ret UniqueIDType
		return ret
	}
	return *o.AwardCancellationNo
}

// GetAwardCancellationNoOk returns a tuple with the AwardCancellationNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetAwardCancellationNoOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.AwardCancellationNo) {
		return nil, false
	}
	return o.AwardCancellationNo, true
}

// HasAwardCancellationNo returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasAwardCancellationNo() bool {
	if o != nil && !IsNil(o.AwardCancellationNo) {
		return true
	}

	return false
}

// SetAwardCancellationNo gets a reference to the given UniqueIDType and assigns it to the AwardCancellationNo field.
func (o *MemberAwardDetailType) SetAwardCancellationNo(v UniqueIDType) {
	o.AwardCancellationNo = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *MemberAwardDetailType) SetInactive(v bool) {
	o.Inactive = &v
}

// GetPointsRequired returns the PointsRequired field value if set, zero value otherwise.
func (o *MemberAwardDetailType) GetPointsRequired() int32 {
	if o == nil || IsNil(o.PointsRequired) {
		var ret int32
		return ret
	}
	return *o.PointsRequired
}

// GetPointsRequiredOk returns a tuple with the PointsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAwardDetailType) GetPointsRequiredOk() (*int32, bool) {
	if o == nil || IsNil(o.PointsRequired) {
		return nil, false
	}
	return o.PointsRequired, true
}

// HasPointsRequired returns a boolean if a field has been set.
func (o *MemberAwardDetailType) HasPointsRequired() bool {
	if o != nil && !IsNil(o.PointsRequired) {
		return true
	}

	return false
}

// SetPointsRequired gets a reference to the given int32 and assigns it to the PointsRequired field.
func (o *MemberAwardDetailType) SetPointsRequired(v int32) {
	o.PointsRequired = &v
}

func (o MemberAwardDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberAwardDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwardType) {
		toSerialize["awardType"] = o.AwardType
	}
	if !IsNil(o.AwardBasedOn) {
		toSerialize["awardBasedOn"] = o.AwardBasedOn
	}
	if !IsNil(o.CancelPenaltyDays) {
		toSerialize["cancelPenaltyDays"] = o.CancelPenaltyDays
	}
	if !IsNil(o.CancelPenaltyCharge) {
		toSerialize["cancelPenaltyCharge"] = o.CancelPenaltyCharge
	}
	if !IsNil(o.CancelPenaltyType) {
		toSerialize["cancelPenaltyType"] = o.CancelPenaltyType
	}
	if !IsNil(o.CancelPenaltyPoints) {
		toSerialize["cancelPenaltyPoints"] = o.CancelPenaltyPoints
	}
	if !IsNil(o.StayDate) {
		toSerialize["stayDate"] = o.StayDate
	}
	if !IsNil(o.RateCode) {
		toSerialize["rateCode"] = o.RateCode
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.FromRoomType) {
		toSerialize["fromRoomType"] = o.FromRoomType
	}
	if !IsNil(o.ToRoomType) {
		toSerialize["toRoomType"] = o.ToRoomType
	}
	if !IsNil(o.TotalLocalAmount) {
		toSerialize["totalLocalAmount"] = o.TotalLocalAmount
	}
	if !IsNil(o.RedeemedLocalAmount) {
		toSerialize["redeemedLocalAmount"] = o.RedeemedLocalAmount
	}
	if !IsNil(o.TotalCentralAmount) {
		toSerialize["totalCentralAmount"] = o.TotalCentralAmount
	}
	if !IsNil(o.RedeemedCentralAmount) {
		toSerialize["redeemedCentralAmount"] = o.RedeemedCentralAmount
	}
	if !IsNil(o.TransactionCode) {
		toSerialize["transactionCode"] = o.TransactionCode
	}
	if !IsNil(o.TransactionNo) {
		toSerialize["transactionNo"] = o.TransactionNo
	}
	if !IsNil(o.ExchangeRateType) {
		toSerialize["exchangeRateType"] = o.ExchangeRateType
	}
	if !IsNil(o.AwardVoucherNo) {
		toSerialize["awardVoucherNo"] = o.AwardVoucherNo
	}
	if !IsNil(o.AwardCancellationNo) {
		toSerialize["awardCancellationNo"] = o.AwardCancellationNo
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.PointsRequired) {
		toSerialize["pointsRequired"] = o.PointsRequired
	}
	return toSerialize, nil
}

type NullableMemberAwardDetailType struct {
	value *MemberAwardDetailType
	isSet bool
}

func (v NullableMemberAwardDetailType) Get() *MemberAwardDetailType {
	return v.value
}

func (v *NullableMemberAwardDetailType) Set(val *MemberAwardDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberAwardDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberAwardDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberAwardDetailType(val *MemberAwardDetailType) *NullableMemberAwardDetailType {
	return &NullableMemberAwardDetailType{value: val, isSet: true}
}

func (v NullableMemberAwardDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberAwardDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


