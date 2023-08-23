/*
OPERA Cloud Reservation Asynchronous API

APIs to cater for Reservation Asynchronous functionality in OPERA Cloud. This includes viewing reservation data along with its revenue. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.4.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the DailySummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailySummaryType{}

// DailySummaryType Statistic information related to a specific reservation stay date.
type DailySummaryType struct {
	// The rate plan code of the reservation.
	RateCode *string `json:"rateCode,omitempty"`
	// A monetary amount.
	RateAmount *float32 `json:"rateAmount,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	RateAmountCurrency *string `json:"rateAmountCurrency,omitempty"`
	// Holds the code that relates to the market being sold (e.g., the corporate market, packages).
	MarketCode *string `json:"marketCode,omitempty"`
	// Room type.
	RoomType *string `json:"roomType,omitempty"`
	// Booked room type.
	BookedRoomType *string `json:"bookedRoomType,omitempty"`
	// Room number.
	Room *string `json:"room,omitempty"`
	// Net Amount generated by the reservation.
	NetRateAmount *float32 `json:"netRateAmount,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	NetRateAmountCurrency *string `json:"netRateAmountCurrency,omitempty"`
	// Room Revenue generated by the reservation.
	RoomRevenue *float32 `json:"roomRevenue,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	RoomRevenueCurrency *string `json:"roomRevenueCurrency,omitempty"`
	// Food and Beverage Revenue generated by the reservation.
	FbRevenue *float32 `json:"fbRevenue,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	FbRevenueCurrency *string `json:"fbRevenueCurrency,omitempty"`
	// Other Revenue generated by the reservation.
	OtherRevenue *float32 `json:"otherRevenue,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	OtherRevenueCurrency *string `json:"otherRevenueCurrency,omitempty"`
	// Total Revenue generated by the reservation.
	TotalRevenue *float32 `json:"totalRevenue,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	TotalRevenueCurrency *string `json:"totalRevenueCurrency,omitempty"`
	// Package Revenue generated by the reservation.
	PackageRevenue *float32 `json:"packageRevenue,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	PackageRevenueCurrency *string `json:"packageRevenueCurrency,omitempty"`
	// Tax Revenue generated by the reservation.
	Tax *float32 `json:"tax,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	TaxCurrency *string `json:"taxCurrency,omitempty"`
	// Room type charged for the reservation.
	RoomTypeCharged *string `json:"roomTypeCharged,omitempty"`
	// Reservation stay date for which the daily statistics are projected.
	TrxDate *time.Time `json:"trxDate,omitempty"`
	// The entity/channel who made the reservation.
	SourceCode *string `json:"sourceCode,omitempty"`
	// Classifies the Channel field on reservation screen through which the reservation is made.
	Channel *string `json:"channel,omitempty"`
	// Number of adults of the reservation.
	Adults *int32 `json:"adults,omitempty"`
	// Number of children of the reservation.
	Children *int32 `json:"children,omitempty"`
}

// NewDailySummaryType instantiates a new DailySummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailySummaryType() *DailySummaryType {
	this := DailySummaryType{}
	return &this
}

// NewDailySummaryTypeWithDefaults instantiates a new DailySummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailySummaryTypeWithDefaults() *DailySummaryType {
	this := DailySummaryType{}
	return &this
}

// GetRateCode returns the RateCode field value if set, zero value otherwise.
func (o *DailySummaryType) GetRateCode() string {
	if o == nil || IsNil(o.RateCode) {
		var ret string
		return ret
	}
	return *o.RateCode
}

// GetRateCodeOk returns a tuple with the RateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RateCode) {
		return nil, false
	}
	return o.RateCode, true
}

// HasRateCode returns a boolean if a field has been set.
func (o *DailySummaryType) HasRateCode() bool {
	if o != nil && !IsNil(o.RateCode) {
		return true
	}

	return false
}

// SetRateCode gets a reference to the given string and assigns it to the RateCode field.
func (o *DailySummaryType) SetRateCode(v string) {
	o.RateCode = &v
}

// GetRateAmount returns the RateAmount field value if set, zero value otherwise.
func (o *DailySummaryType) GetRateAmount() float32 {
	if o == nil || IsNil(o.RateAmount) {
		var ret float32
		return ret
	}
	return *o.RateAmount
}

// GetRateAmountOk returns a tuple with the RateAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRateAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RateAmount) {
		return nil, false
	}
	return o.RateAmount, true
}

// HasRateAmount returns a boolean if a field has been set.
func (o *DailySummaryType) HasRateAmount() bool {
	if o != nil && !IsNil(o.RateAmount) {
		return true
	}

	return false
}

// SetRateAmount gets a reference to the given float32 and assigns it to the RateAmount field.
func (o *DailySummaryType) SetRateAmount(v float32) {
	o.RateAmount = &v
}

// GetRateAmountCurrency returns the RateAmountCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetRateAmountCurrency() string {
	if o == nil || IsNil(o.RateAmountCurrency) {
		var ret string
		return ret
	}
	return *o.RateAmountCurrency
}

// GetRateAmountCurrencyOk returns a tuple with the RateAmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRateAmountCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.RateAmountCurrency) {
		return nil, false
	}
	return o.RateAmountCurrency, true
}

// HasRateAmountCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasRateAmountCurrency() bool {
	if o != nil && !IsNil(o.RateAmountCurrency) {
		return true
	}

	return false
}

// SetRateAmountCurrency gets a reference to the given string and assigns it to the RateAmountCurrency field.
func (o *DailySummaryType) SetRateAmountCurrency(v string) {
	o.RateAmountCurrency = &v
}

// GetMarketCode returns the MarketCode field value if set, zero value otherwise.
func (o *DailySummaryType) GetMarketCode() string {
	if o == nil || IsNil(o.MarketCode) {
		var ret string
		return ret
	}
	return *o.MarketCode
}

// GetMarketCodeOk returns a tuple with the MarketCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetMarketCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MarketCode) {
		return nil, false
	}
	return o.MarketCode, true
}

// HasMarketCode returns a boolean if a field has been set.
func (o *DailySummaryType) HasMarketCode() bool {
	if o != nil && !IsNil(o.MarketCode) {
		return true
	}

	return false
}

// SetMarketCode gets a reference to the given string and assigns it to the MarketCode field.
func (o *DailySummaryType) SetMarketCode(v string) {
	o.MarketCode = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *DailySummaryType) GetRoomType() string {
	if o == nil || IsNil(o.RoomType) {
		var ret string
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *DailySummaryType) HasRoomType() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given string and assigns it to the RoomType field.
func (o *DailySummaryType) SetRoomType(v string) {
	o.RoomType = &v
}

// GetBookedRoomType returns the BookedRoomType field value if set, zero value otherwise.
func (o *DailySummaryType) GetBookedRoomType() string {
	if o == nil || IsNil(o.BookedRoomType) {
		var ret string
		return ret
	}
	return *o.BookedRoomType
}

// GetBookedRoomTypeOk returns a tuple with the BookedRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetBookedRoomTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BookedRoomType) {
		return nil, false
	}
	return o.BookedRoomType, true
}

// HasBookedRoomType returns a boolean if a field has been set.
func (o *DailySummaryType) HasBookedRoomType() bool {
	if o != nil && !IsNil(o.BookedRoomType) {
		return true
	}

	return false
}

// SetBookedRoomType gets a reference to the given string and assigns it to the BookedRoomType field.
func (o *DailySummaryType) SetBookedRoomType(v string) {
	o.BookedRoomType = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *DailySummaryType) GetRoom() string {
	if o == nil || IsNil(o.Room) {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRoomOk() (*string, bool) {
	if o == nil || IsNil(o.Room) {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *DailySummaryType) HasRoom() bool {
	if o != nil && !IsNil(o.Room) {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *DailySummaryType) SetRoom(v string) {
	o.Room = &v
}

// GetNetRateAmount returns the NetRateAmount field value if set, zero value otherwise.
func (o *DailySummaryType) GetNetRateAmount() float32 {
	if o == nil || IsNil(o.NetRateAmount) {
		var ret float32
		return ret
	}
	return *o.NetRateAmount
}

// GetNetRateAmountOk returns a tuple with the NetRateAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetNetRateAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.NetRateAmount) {
		return nil, false
	}
	return o.NetRateAmount, true
}

// HasNetRateAmount returns a boolean if a field has been set.
func (o *DailySummaryType) HasNetRateAmount() bool {
	if o != nil && !IsNil(o.NetRateAmount) {
		return true
	}

	return false
}

// SetNetRateAmount gets a reference to the given float32 and assigns it to the NetRateAmount field.
func (o *DailySummaryType) SetNetRateAmount(v float32) {
	o.NetRateAmount = &v
}

// GetNetRateAmountCurrency returns the NetRateAmountCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetNetRateAmountCurrency() string {
	if o == nil || IsNil(o.NetRateAmountCurrency) {
		var ret string
		return ret
	}
	return *o.NetRateAmountCurrency
}

// GetNetRateAmountCurrencyOk returns a tuple with the NetRateAmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetNetRateAmountCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.NetRateAmountCurrency) {
		return nil, false
	}
	return o.NetRateAmountCurrency, true
}

// HasNetRateAmountCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasNetRateAmountCurrency() bool {
	if o != nil && !IsNil(o.NetRateAmountCurrency) {
		return true
	}

	return false
}

// SetNetRateAmountCurrency gets a reference to the given string and assigns it to the NetRateAmountCurrency field.
func (o *DailySummaryType) SetNetRateAmountCurrency(v string) {
	o.NetRateAmountCurrency = &v
}

// GetRoomRevenue returns the RoomRevenue field value if set, zero value otherwise.
func (o *DailySummaryType) GetRoomRevenue() float32 {
	if o == nil || IsNil(o.RoomRevenue) {
		var ret float32
		return ret
	}
	return *o.RoomRevenue
}

// GetRoomRevenueOk returns a tuple with the RoomRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRoomRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.RoomRevenue) {
		return nil, false
	}
	return o.RoomRevenue, true
}

// HasRoomRevenue returns a boolean if a field has been set.
func (o *DailySummaryType) HasRoomRevenue() bool {
	if o != nil && !IsNil(o.RoomRevenue) {
		return true
	}

	return false
}

// SetRoomRevenue gets a reference to the given float32 and assigns it to the RoomRevenue field.
func (o *DailySummaryType) SetRoomRevenue(v float32) {
	o.RoomRevenue = &v
}

// GetRoomRevenueCurrency returns the RoomRevenueCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetRoomRevenueCurrency() string {
	if o == nil || IsNil(o.RoomRevenueCurrency) {
		var ret string
		return ret
	}
	return *o.RoomRevenueCurrency
}

// GetRoomRevenueCurrencyOk returns a tuple with the RoomRevenueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRoomRevenueCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.RoomRevenueCurrency) {
		return nil, false
	}
	return o.RoomRevenueCurrency, true
}

// HasRoomRevenueCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasRoomRevenueCurrency() bool {
	if o != nil && !IsNil(o.RoomRevenueCurrency) {
		return true
	}

	return false
}

// SetRoomRevenueCurrency gets a reference to the given string and assigns it to the RoomRevenueCurrency field.
func (o *DailySummaryType) SetRoomRevenueCurrency(v string) {
	o.RoomRevenueCurrency = &v
}

// GetFbRevenue returns the FbRevenue field value if set, zero value otherwise.
func (o *DailySummaryType) GetFbRevenue() float32 {
	if o == nil || IsNil(o.FbRevenue) {
		var ret float32
		return ret
	}
	return *o.FbRevenue
}

// GetFbRevenueOk returns a tuple with the FbRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetFbRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.FbRevenue) {
		return nil, false
	}
	return o.FbRevenue, true
}

// HasFbRevenue returns a boolean if a field has been set.
func (o *DailySummaryType) HasFbRevenue() bool {
	if o != nil && !IsNil(o.FbRevenue) {
		return true
	}

	return false
}

// SetFbRevenue gets a reference to the given float32 and assigns it to the FbRevenue field.
func (o *DailySummaryType) SetFbRevenue(v float32) {
	o.FbRevenue = &v
}

// GetFbRevenueCurrency returns the FbRevenueCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetFbRevenueCurrency() string {
	if o == nil || IsNil(o.FbRevenueCurrency) {
		var ret string
		return ret
	}
	return *o.FbRevenueCurrency
}

// GetFbRevenueCurrencyOk returns a tuple with the FbRevenueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetFbRevenueCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.FbRevenueCurrency) {
		return nil, false
	}
	return o.FbRevenueCurrency, true
}

// HasFbRevenueCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasFbRevenueCurrency() bool {
	if o != nil && !IsNil(o.FbRevenueCurrency) {
		return true
	}

	return false
}

// SetFbRevenueCurrency gets a reference to the given string and assigns it to the FbRevenueCurrency field.
func (o *DailySummaryType) SetFbRevenueCurrency(v string) {
	o.FbRevenueCurrency = &v
}

// GetOtherRevenue returns the OtherRevenue field value if set, zero value otherwise.
func (o *DailySummaryType) GetOtherRevenue() float32 {
	if o == nil || IsNil(o.OtherRevenue) {
		var ret float32
		return ret
	}
	return *o.OtherRevenue
}

// GetOtherRevenueOk returns a tuple with the OtherRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetOtherRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.OtherRevenue) {
		return nil, false
	}
	return o.OtherRevenue, true
}

// HasOtherRevenue returns a boolean if a field has been set.
func (o *DailySummaryType) HasOtherRevenue() bool {
	if o != nil && !IsNil(o.OtherRevenue) {
		return true
	}

	return false
}

// SetOtherRevenue gets a reference to the given float32 and assigns it to the OtherRevenue field.
func (o *DailySummaryType) SetOtherRevenue(v float32) {
	o.OtherRevenue = &v
}

// GetOtherRevenueCurrency returns the OtherRevenueCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetOtherRevenueCurrency() string {
	if o == nil || IsNil(o.OtherRevenueCurrency) {
		var ret string
		return ret
	}
	return *o.OtherRevenueCurrency
}

// GetOtherRevenueCurrencyOk returns a tuple with the OtherRevenueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetOtherRevenueCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.OtherRevenueCurrency) {
		return nil, false
	}
	return o.OtherRevenueCurrency, true
}

// HasOtherRevenueCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasOtherRevenueCurrency() bool {
	if o != nil && !IsNil(o.OtherRevenueCurrency) {
		return true
	}

	return false
}

// SetOtherRevenueCurrency gets a reference to the given string and assigns it to the OtherRevenueCurrency field.
func (o *DailySummaryType) SetOtherRevenueCurrency(v string) {
	o.OtherRevenueCurrency = &v
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *DailySummaryType) GetTotalRevenue() float32 {
	if o == nil || IsNil(o.TotalRevenue) {
		var ret float32
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetTotalRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRevenue) {
		return nil, false
	}
	return o.TotalRevenue, true
}

// HasTotalRevenue returns a boolean if a field has been set.
func (o *DailySummaryType) HasTotalRevenue() bool {
	if o != nil && !IsNil(o.TotalRevenue) {
		return true
	}

	return false
}

// SetTotalRevenue gets a reference to the given float32 and assigns it to the TotalRevenue field.
func (o *DailySummaryType) SetTotalRevenue(v float32) {
	o.TotalRevenue = &v
}

// GetTotalRevenueCurrency returns the TotalRevenueCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetTotalRevenueCurrency() string {
	if o == nil || IsNil(o.TotalRevenueCurrency) {
		var ret string
		return ret
	}
	return *o.TotalRevenueCurrency
}

// GetTotalRevenueCurrencyOk returns a tuple with the TotalRevenueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetTotalRevenueCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.TotalRevenueCurrency) {
		return nil, false
	}
	return o.TotalRevenueCurrency, true
}

// HasTotalRevenueCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasTotalRevenueCurrency() bool {
	if o != nil && !IsNil(o.TotalRevenueCurrency) {
		return true
	}

	return false
}

// SetTotalRevenueCurrency gets a reference to the given string and assigns it to the TotalRevenueCurrency field.
func (o *DailySummaryType) SetTotalRevenueCurrency(v string) {
	o.TotalRevenueCurrency = &v
}

// GetPackageRevenue returns the PackageRevenue field value if set, zero value otherwise.
func (o *DailySummaryType) GetPackageRevenue() float32 {
	if o == nil || IsNil(o.PackageRevenue) {
		var ret float32
		return ret
	}
	return *o.PackageRevenue
}

// GetPackageRevenueOk returns a tuple with the PackageRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetPackageRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.PackageRevenue) {
		return nil, false
	}
	return o.PackageRevenue, true
}

// HasPackageRevenue returns a boolean if a field has been set.
func (o *DailySummaryType) HasPackageRevenue() bool {
	if o != nil && !IsNil(o.PackageRevenue) {
		return true
	}

	return false
}

// SetPackageRevenue gets a reference to the given float32 and assigns it to the PackageRevenue field.
func (o *DailySummaryType) SetPackageRevenue(v float32) {
	o.PackageRevenue = &v
}

// GetPackageRevenueCurrency returns the PackageRevenueCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetPackageRevenueCurrency() string {
	if o == nil || IsNil(o.PackageRevenueCurrency) {
		var ret string
		return ret
	}
	return *o.PackageRevenueCurrency
}

// GetPackageRevenueCurrencyOk returns a tuple with the PackageRevenueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetPackageRevenueCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.PackageRevenueCurrency) {
		return nil, false
	}
	return o.PackageRevenueCurrency, true
}

// HasPackageRevenueCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasPackageRevenueCurrency() bool {
	if o != nil && !IsNil(o.PackageRevenueCurrency) {
		return true
	}

	return false
}

// SetPackageRevenueCurrency gets a reference to the given string and assigns it to the PackageRevenueCurrency field.
func (o *DailySummaryType) SetPackageRevenueCurrency(v string) {
	o.PackageRevenueCurrency = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *DailySummaryType) GetTax() float32 {
	if o == nil || IsNil(o.Tax) {
		var ret float32
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *DailySummaryType) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given float32 and assigns it to the Tax field.
func (o *DailySummaryType) SetTax(v float32) {
	o.Tax = &v
}

// GetTaxCurrency returns the TaxCurrency field value if set, zero value otherwise.
func (o *DailySummaryType) GetTaxCurrency() string {
	if o == nil || IsNil(o.TaxCurrency) {
		var ret string
		return ret
	}
	return *o.TaxCurrency
}

// GetTaxCurrencyOk returns a tuple with the TaxCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetTaxCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.TaxCurrency) {
		return nil, false
	}
	return o.TaxCurrency, true
}

// HasTaxCurrency returns a boolean if a field has been set.
func (o *DailySummaryType) HasTaxCurrency() bool {
	if o != nil && !IsNil(o.TaxCurrency) {
		return true
	}

	return false
}

// SetTaxCurrency gets a reference to the given string and assigns it to the TaxCurrency field.
func (o *DailySummaryType) SetTaxCurrency(v string) {
	o.TaxCurrency = &v
}

// GetRoomTypeCharged returns the RoomTypeCharged field value if set, zero value otherwise.
func (o *DailySummaryType) GetRoomTypeCharged() string {
	if o == nil || IsNil(o.RoomTypeCharged) {
		var ret string
		return ret
	}
	return *o.RoomTypeCharged
}

// GetRoomTypeChargedOk returns a tuple with the RoomTypeCharged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetRoomTypeChargedOk() (*string, bool) {
	if o == nil || IsNil(o.RoomTypeCharged) {
		return nil, false
	}
	return o.RoomTypeCharged, true
}

// HasRoomTypeCharged returns a boolean if a field has been set.
func (o *DailySummaryType) HasRoomTypeCharged() bool {
	if o != nil && !IsNil(o.RoomTypeCharged) {
		return true
	}

	return false
}

// SetRoomTypeCharged gets a reference to the given string and assigns it to the RoomTypeCharged field.
func (o *DailySummaryType) SetRoomTypeCharged(v string) {
	o.RoomTypeCharged = &v
}

// GetTrxDate returns the TrxDate field value if set, zero value otherwise.
func (o *DailySummaryType) GetTrxDate() time.Time {
	if o == nil || IsNil(o.TrxDate) {
		var ret time.Time
		return ret
	}
	return *o.TrxDate
}

// GetTrxDateOk returns a tuple with the TrxDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetTrxDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TrxDate) {
		return nil, false
	}
	return o.TrxDate, true
}

// HasTrxDate returns a boolean if a field has been set.
func (o *DailySummaryType) HasTrxDate() bool {
	if o != nil && !IsNil(o.TrxDate) {
		return true
	}

	return false
}

// SetTrxDate gets a reference to the given time.Time and assigns it to the TrxDate field.
func (o *DailySummaryType) SetTrxDate(v time.Time) {
	o.TrxDate = &v
}

// GetSourceCode returns the SourceCode field value if set, zero value otherwise.
func (o *DailySummaryType) GetSourceCode() string {
	if o == nil || IsNil(o.SourceCode) {
		var ret string
		return ret
	}
	return *o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetSourceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCode) {
		return nil, false
	}
	return o.SourceCode, true
}

// HasSourceCode returns a boolean if a field has been set.
func (o *DailySummaryType) HasSourceCode() bool {
	if o != nil && !IsNil(o.SourceCode) {
		return true
	}

	return false
}

// SetSourceCode gets a reference to the given string and assigns it to the SourceCode field.
func (o *DailySummaryType) SetSourceCode(v string) {
	o.SourceCode = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *DailySummaryType) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *DailySummaryType) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *DailySummaryType) SetChannel(v string) {
	o.Channel = &v
}

// GetAdults returns the Adults field value if set, zero value otherwise.
func (o *DailySummaryType) GetAdults() int32 {
	if o == nil || IsNil(o.Adults) {
		var ret int32
		return ret
	}
	return *o.Adults
}

// GetAdultsOk returns a tuple with the Adults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetAdultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Adults) {
		return nil, false
	}
	return o.Adults, true
}

// HasAdults returns a boolean if a field has been set.
func (o *DailySummaryType) HasAdults() bool {
	if o != nil && !IsNil(o.Adults) {
		return true
	}

	return false
}

// SetAdults gets a reference to the given int32 and assigns it to the Adults field.
func (o *DailySummaryType) SetAdults(v int32) {
	o.Adults = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *DailySummaryType) GetChildren() int32 {
	if o == nil || IsNil(o.Children) {
		var ret int32
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySummaryType) GetChildrenOk() (*int32, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *DailySummaryType) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given int32 and assigns it to the Children field.
func (o *DailySummaryType) SetChildren(v int32) {
	o.Children = &v
}

func (o DailySummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailySummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RateCode) {
		toSerialize["rateCode"] = o.RateCode
	}
	if !IsNil(o.RateAmount) {
		toSerialize["rateAmount"] = o.RateAmount
	}
	if !IsNil(o.RateAmountCurrency) {
		toSerialize["rateAmountCurrency"] = o.RateAmountCurrency
	}
	if !IsNil(o.MarketCode) {
		toSerialize["marketCode"] = o.MarketCode
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.BookedRoomType) {
		toSerialize["bookedRoomType"] = o.BookedRoomType
	}
	if !IsNil(o.Room) {
		toSerialize["room"] = o.Room
	}
	if !IsNil(o.NetRateAmount) {
		toSerialize["netRateAmount"] = o.NetRateAmount
	}
	if !IsNil(o.NetRateAmountCurrency) {
		toSerialize["netRateAmountCurrency"] = o.NetRateAmountCurrency
	}
	if !IsNil(o.RoomRevenue) {
		toSerialize["roomRevenue"] = o.RoomRevenue
	}
	if !IsNil(o.RoomRevenueCurrency) {
		toSerialize["roomRevenueCurrency"] = o.RoomRevenueCurrency
	}
	if !IsNil(o.FbRevenue) {
		toSerialize["fbRevenue"] = o.FbRevenue
	}
	if !IsNil(o.FbRevenueCurrency) {
		toSerialize["fbRevenueCurrency"] = o.FbRevenueCurrency
	}
	if !IsNil(o.OtherRevenue) {
		toSerialize["otherRevenue"] = o.OtherRevenue
	}
	if !IsNil(o.OtherRevenueCurrency) {
		toSerialize["otherRevenueCurrency"] = o.OtherRevenueCurrency
	}
	if !IsNil(o.TotalRevenue) {
		toSerialize["totalRevenue"] = o.TotalRevenue
	}
	if !IsNil(o.TotalRevenueCurrency) {
		toSerialize["totalRevenueCurrency"] = o.TotalRevenueCurrency
	}
	if !IsNil(o.PackageRevenue) {
		toSerialize["packageRevenue"] = o.PackageRevenue
	}
	if !IsNil(o.PackageRevenueCurrency) {
		toSerialize["packageRevenueCurrency"] = o.PackageRevenueCurrency
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	if !IsNil(o.TaxCurrency) {
		toSerialize["taxCurrency"] = o.TaxCurrency
	}
	if !IsNil(o.RoomTypeCharged) {
		toSerialize["roomTypeCharged"] = o.RoomTypeCharged
	}
	if !IsNil(o.TrxDate) {
		toSerialize["trxDate"] = o.TrxDate
	}
	if !IsNil(o.SourceCode) {
		toSerialize["sourceCode"] = o.SourceCode
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Adults) {
		toSerialize["adults"] = o.Adults
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableDailySummaryType struct {
	value *DailySummaryType
	isSet bool
}

func (v NullableDailySummaryType) Get() *DailySummaryType {
	return v.value
}

func (v *NullableDailySummaryType) Set(val *DailySummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableDailySummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableDailySummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailySummaryType(val *DailySummaryType) *NullableDailySummaryType {
	return &NullableDailySummaryType{value: val, isSet: true}
}

func (v NullableDailySummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailySummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


