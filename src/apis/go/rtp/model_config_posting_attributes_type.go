/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtp

import (
	"encoding/json"
)

// checks if the ConfigPostingAttributesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigPostingAttributesType{}

// ConfigPostingAttributesType A config Package posting attributes type.
type ConfigPostingAttributesType struct {
	// The package price is added to the room rate.
	AddToRate *bool `json:"addToRate,omitempty"`
	// The package price is printed on separate line of the folio.
	PrintSeparateLine *bool `json:"printSeparateLine,omitempty"`
	// Can the package be sold separate from rate plan code?
	SellSeparate *bool `json:"sellSeparate,omitempty"`
	// package charges will be posted next business day.
	PostNextDay *bool `json:"postNextDay,omitempty"`
	// Package will be forecasted for consumption the next business day.
	ForecastNextDay *bool `json:"forecastNextDay,omitempty"`
	// Indicates whether a delivery time is required for the package.
	DeliveryTimeRequired *bool `json:"deliveryTimeRequired,omitempty"`
	// Can package be sold via Web channel.
	WebBookable *bool `json:"webBookable,omitempty"`
	// The custom formula used for this package, if any.
	Formula *string `json:"formula,omitempty"`
	// The custom formula function name used for this package, if any.
	FormulaFunctionName *string `json:"formulaFunctionName,omitempty"`
	// Collection of function arguments and their corresponding values.
	FormulaFunctionArguments []FunctionArgumentType `json:"formulaFunctionArguments,omitempty"`
	// Start time the package is valid.
	StartTime *string `json:"startTime,omitempty"`
	// End time the package is valid.
	EndTime *string `json:"endTime,omitempty"`
	// Is package used for catering?
	Catering *bool `json:"catering,omitempty"`
	PostingRhythm *PackagePostingRhythmType `json:"postingRhythm,omitempty"`
	PriceCalculationRule *PackageCalculationRuleType `json:"priceCalculationRule,omitempty"`
	// Indicates whether a package is configured as a ticket or not.
	Ticket *bool `json:"ticket,omitempty"`
	// Package Code Inventory Items type.
	InventoryItems []PkgInventoryItemType `json:"inventoryItems,omitempty"`
	// Calculated Package Price based from Number of Adults, Children and Calculation Rule.
	CalculatedPrice *float32 `json:"calculatedPrice,omitempty"`
}

// NewConfigPostingAttributesType instantiates a new ConfigPostingAttributesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigPostingAttributesType() *ConfigPostingAttributesType {
	this := ConfigPostingAttributesType{}
	return &this
}

// NewConfigPostingAttributesTypeWithDefaults instantiates a new ConfigPostingAttributesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigPostingAttributesTypeWithDefaults() *ConfigPostingAttributesType {
	this := ConfigPostingAttributesType{}
	return &this
}

// GetAddToRate returns the AddToRate field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetAddToRate() bool {
	if o == nil || IsNil(o.AddToRate) {
		var ret bool
		return ret
	}
	return *o.AddToRate
}

// GetAddToRateOk returns a tuple with the AddToRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetAddToRateOk() (*bool, bool) {
	if o == nil || IsNil(o.AddToRate) {
		return nil, false
	}
	return o.AddToRate, true
}

// HasAddToRate returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasAddToRate() bool {
	if o != nil && !IsNil(o.AddToRate) {
		return true
	}

	return false
}

// SetAddToRate gets a reference to the given bool and assigns it to the AddToRate field.
func (o *ConfigPostingAttributesType) SetAddToRate(v bool) {
	o.AddToRate = &v
}

// GetPrintSeparateLine returns the PrintSeparateLine field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetPrintSeparateLine() bool {
	if o == nil || IsNil(o.PrintSeparateLine) {
		var ret bool
		return ret
	}
	return *o.PrintSeparateLine
}

// GetPrintSeparateLineOk returns a tuple with the PrintSeparateLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetPrintSeparateLineOk() (*bool, bool) {
	if o == nil || IsNil(o.PrintSeparateLine) {
		return nil, false
	}
	return o.PrintSeparateLine, true
}

// HasPrintSeparateLine returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasPrintSeparateLine() bool {
	if o != nil && !IsNil(o.PrintSeparateLine) {
		return true
	}

	return false
}

// SetPrintSeparateLine gets a reference to the given bool and assigns it to the PrintSeparateLine field.
func (o *ConfigPostingAttributesType) SetPrintSeparateLine(v bool) {
	o.PrintSeparateLine = &v
}

// GetSellSeparate returns the SellSeparate field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetSellSeparate() bool {
	if o == nil || IsNil(o.SellSeparate) {
		var ret bool
		return ret
	}
	return *o.SellSeparate
}

// GetSellSeparateOk returns a tuple with the SellSeparate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetSellSeparateOk() (*bool, bool) {
	if o == nil || IsNil(o.SellSeparate) {
		return nil, false
	}
	return o.SellSeparate, true
}

// HasSellSeparate returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasSellSeparate() bool {
	if o != nil && !IsNil(o.SellSeparate) {
		return true
	}

	return false
}

// SetSellSeparate gets a reference to the given bool and assigns it to the SellSeparate field.
func (o *ConfigPostingAttributesType) SetSellSeparate(v bool) {
	o.SellSeparate = &v
}

// GetPostNextDay returns the PostNextDay field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetPostNextDay() bool {
	if o == nil || IsNil(o.PostNextDay) {
		var ret bool
		return ret
	}
	return *o.PostNextDay
}

// GetPostNextDayOk returns a tuple with the PostNextDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetPostNextDayOk() (*bool, bool) {
	if o == nil || IsNil(o.PostNextDay) {
		return nil, false
	}
	return o.PostNextDay, true
}

// HasPostNextDay returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasPostNextDay() bool {
	if o != nil && !IsNil(o.PostNextDay) {
		return true
	}

	return false
}

// SetPostNextDay gets a reference to the given bool and assigns it to the PostNextDay field.
func (o *ConfigPostingAttributesType) SetPostNextDay(v bool) {
	o.PostNextDay = &v
}

// GetForecastNextDay returns the ForecastNextDay field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetForecastNextDay() bool {
	if o == nil || IsNil(o.ForecastNextDay) {
		var ret bool
		return ret
	}
	return *o.ForecastNextDay
}

// GetForecastNextDayOk returns a tuple with the ForecastNextDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetForecastNextDayOk() (*bool, bool) {
	if o == nil || IsNil(o.ForecastNextDay) {
		return nil, false
	}
	return o.ForecastNextDay, true
}

// HasForecastNextDay returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasForecastNextDay() bool {
	if o != nil && !IsNil(o.ForecastNextDay) {
		return true
	}

	return false
}

// SetForecastNextDay gets a reference to the given bool and assigns it to the ForecastNextDay field.
func (o *ConfigPostingAttributesType) SetForecastNextDay(v bool) {
	o.ForecastNextDay = &v
}

// GetDeliveryTimeRequired returns the DeliveryTimeRequired field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetDeliveryTimeRequired() bool {
	if o == nil || IsNil(o.DeliveryTimeRequired) {
		var ret bool
		return ret
	}
	return *o.DeliveryTimeRequired
}

// GetDeliveryTimeRequiredOk returns a tuple with the DeliveryTimeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetDeliveryTimeRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.DeliveryTimeRequired) {
		return nil, false
	}
	return o.DeliveryTimeRequired, true
}

// HasDeliveryTimeRequired returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasDeliveryTimeRequired() bool {
	if o != nil && !IsNil(o.DeliveryTimeRequired) {
		return true
	}

	return false
}

// SetDeliveryTimeRequired gets a reference to the given bool and assigns it to the DeliveryTimeRequired field.
func (o *ConfigPostingAttributesType) SetDeliveryTimeRequired(v bool) {
	o.DeliveryTimeRequired = &v
}

// GetWebBookable returns the WebBookable field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetWebBookable() bool {
	if o == nil || IsNil(o.WebBookable) {
		var ret bool
		return ret
	}
	return *o.WebBookable
}

// GetWebBookableOk returns a tuple with the WebBookable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetWebBookableOk() (*bool, bool) {
	if o == nil || IsNil(o.WebBookable) {
		return nil, false
	}
	return o.WebBookable, true
}

// HasWebBookable returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasWebBookable() bool {
	if o != nil && !IsNil(o.WebBookable) {
		return true
	}

	return false
}

// SetWebBookable gets a reference to the given bool and assigns it to the WebBookable field.
func (o *ConfigPostingAttributesType) SetWebBookable(v bool) {
	o.WebBookable = &v
}

// GetFormula returns the Formula field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetFormula() string {
	if o == nil || IsNil(o.Formula) {
		var ret string
		return ret
	}
	return *o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetFormulaOk() (*string, bool) {
	if o == nil || IsNil(o.Formula) {
		return nil, false
	}
	return o.Formula, true
}

// HasFormula returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasFormula() bool {
	if o != nil && !IsNil(o.Formula) {
		return true
	}

	return false
}

// SetFormula gets a reference to the given string and assigns it to the Formula field.
func (o *ConfigPostingAttributesType) SetFormula(v string) {
	o.Formula = &v
}

// GetFormulaFunctionName returns the FormulaFunctionName field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetFormulaFunctionName() string {
	if o == nil || IsNil(o.FormulaFunctionName) {
		var ret string
		return ret
	}
	return *o.FormulaFunctionName
}

// GetFormulaFunctionNameOk returns a tuple with the FormulaFunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetFormulaFunctionNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormulaFunctionName) {
		return nil, false
	}
	return o.FormulaFunctionName, true
}

// HasFormulaFunctionName returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasFormulaFunctionName() bool {
	if o != nil && !IsNil(o.FormulaFunctionName) {
		return true
	}

	return false
}

// SetFormulaFunctionName gets a reference to the given string and assigns it to the FormulaFunctionName field.
func (o *ConfigPostingAttributesType) SetFormulaFunctionName(v string) {
	o.FormulaFunctionName = &v
}

// GetFormulaFunctionArguments returns the FormulaFunctionArguments field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetFormulaFunctionArguments() []FunctionArgumentType {
	if o == nil || IsNil(o.FormulaFunctionArguments) {
		var ret []FunctionArgumentType
		return ret
	}
	return o.FormulaFunctionArguments
}

// GetFormulaFunctionArgumentsOk returns a tuple with the FormulaFunctionArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetFormulaFunctionArgumentsOk() ([]FunctionArgumentType, bool) {
	if o == nil || IsNil(o.FormulaFunctionArguments) {
		return nil, false
	}
	return o.FormulaFunctionArguments, true
}

// HasFormulaFunctionArguments returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasFormulaFunctionArguments() bool {
	if o != nil && !IsNil(o.FormulaFunctionArguments) {
		return true
	}

	return false
}

// SetFormulaFunctionArguments gets a reference to the given []FunctionArgumentType and assigns it to the FormulaFunctionArguments field.
func (o *ConfigPostingAttributesType) SetFormulaFunctionArguments(v []FunctionArgumentType) {
	o.FormulaFunctionArguments = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ConfigPostingAttributesType) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *ConfigPostingAttributesType) SetEndTime(v string) {
	o.EndTime = &v
}

// GetCatering returns the Catering field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetCatering() bool {
	if o == nil || IsNil(o.Catering) {
		var ret bool
		return ret
	}
	return *o.Catering
}

// GetCateringOk returns a tuple with the Catering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetCateringOk() (*bool, bool) {
	if o == nil || IsNil(o.Catering) {
		return nil, false
	}
	return o.Catering, true
}

// HasCatering returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasCatering() bool {
	if o != nil && !IsNil(o.Catering) {
		return true
	}

	return false
}

// SetCatering gets a reference to the given bool and assigns it to the Catering field.
func (o *ConfigPostingAttributesType) SetCatering(v bool) {
	o.Catering = &v
}

// GetPostingRhythm returns the PostingRhythm field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetPostingRhythm() PackagePostingRhythmType {
	if o == nil || IsNil(o.PostingRhythm) {
		var ret PackagePostingRhythmType
		return ret
	}
	return *o.PostingRhythm
}

// GetPostingRhythmOk returns a tuple with the PostingRhythm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetPostingRhythmOk() (*PackagePostingRhythmType, bool) {
	if o == nil || IsNil(o.PostingRhythm) {
		return nil, false
	}
	return o.PostingRhythm, true
}

// HasPostingRhythm returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasPostingRhythm() bool {
	if o != nil && !IsNil(o.PostingRhythm) {
		return true
	}

	return false
}

// SetPostingRhythm gets a reference to the given PackagePostingRhythmType and assigns it to the PostingRhythm field.
func (o *ConfigPostingAttributesType) SetPostingRhythm(v PackagePostingRhythmType) {
	o.PostingRhythm = &v
}

// GetPriceCalculationRule returns the PriceCalculationRule field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetPriceCalculationRule() PackageCalculationRuleType {
	if o == nil || IsNil(o.PriceCalculationRule) {
		var ret PackageCalculationRuleType
		return ret
	}
	return *o.PriceCalculationRule
}

// GetPriceCalculationRuleOk returns a tuple with the PriceCalculationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetPriceCalculationRuleOk() (*PackageCalculationRuleType, bool) {
	if o == nil || IsNil(o.PriceCalculationRule) {
		return nil, false
	}
	return o.PriceCalculationRule, true
}

// HasPriceCalculationRule returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasPriceCalculationRule() bool {
	if o != nil && !IsNil(o.PriceCalculationRule) {
		return true
	}

	return false
}

// SetPriceCalculationRule gets a reference to the given PackageCalculationRuleType and assigns it to the PriceCalculationRule field.
func (o *ConfigPostingAttributesType) SetPriceCalculationRule(v PackageCalculationRuleType) {
	o.PriceCalculationRule = &v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetTicket() bool {
	if o == nil || IsNil(o.Ticket) {
		var ret bool
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetTicketOk() (*bool, bool) {
	if o == nil || IsNil(o.Ticket) {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasTicket() bool {
	if o != nil && !IsNil(o.Ticket) {
		return true
	}

	return false
}

// SetTicket gets a reference to the given bool and assigns it to the Ticket field.
func (o *ConfigPostingAttributesType) SetTicket(v bool) {
	o.Ticket = &v
}

// GetInventoryItems returns the InventoryItems field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetInventoryItems() []PkgInventoryItemType {
	if o == nil || IsNil(o.InventoryItems) {
		var ret []PkgInventoryItemType
		return ret
	}
	return o.InventoryItems
}

// GetInventoryItemsOk returns a tuple with the InventoryItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetInventoryItemsOk() ([]PkgInventoryItemType, bool) {
	if o == nil || IsNil(o.InventoryItems) {
		return nil, false
	}
	return o.InventoryItems, true
}

// HasInventoryItems returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasInventoryItems() bool {
	if o != nil && !IsNil(o.InventoryItems) {
		return true
	}

	return false
}

// SetInventoryItems gets a reference to the given []PkgInventoryItemType and assigns it to the InventoryItems field.
func (o *ConfigPostingAttributesType) SetInventoryItems(v []PkgInventoryItemType) {
	o.InventoryItems = v
}

// GetCalculatedPrice returns the CalculatedPrice field value if set, zero value otherwise.
func (o *ConfigPostingAttributesType) GetCalculatedPrice() float32 {
	if o == nil || IsNil(o.CalculatedPrice) {
		var ret float32
		return ret
	}
	return *o.CalculatedPrice
}

// GetCalculatedPriceOk returns a tuple with the CalculatedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPostingAttributesType) GetCalculatedPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.CalculatedPrice) {
		return nil, false
	}
	return o.CalculatedPrice, true
}

// HasCalculatedPrice returns a boolean if a field has been set.
func (o *ConfigPostingAttributesType) HasCalculatedPrice() bool {
	if o != nil && !IsNil(o.CalculatedPrice) {
		return true
	}

	return false
}

// SetCalculatedPrice gets a reference to the given float32 and assigns it to the CalculatedPrice field.
func (o *ConfigPostingAttributesType) SetCalculatedPrice(v float32) {
	o.CalculatedPrice = &v
}

func (o ConfigPostingAttributesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigPostingAttributesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddToRate) {
		toSerialize["addToRate"] = o.AddToRate
	}
	if !IsNil(o.PrintSeparateLine) {
		toSerialize["printSeparateLine"] = o.PrintSeparateLine
	}
	if !IsNil(o.SellSeparate) {
		toSerialize["sellSeparate"] = o.SellSeparate
	}
	if !IsNil(o.PostNextDay) {
		toSerialize["postNextDay"] = o.PostNextDay
	}
	if !IsNil(o.ForecastNextDay) {
		toSerialize["forecastNextDay"] = o.ForecastNextDay
	}
	if !IsNil(o.DeliveryTimeRequired) {
		toSerialize["deliveryTimeRequired"] = o.DeliveryTimeRequired
	}
	if !IsNil(o.WebBookable) {
		toSerialize["webBookable"] = o.WebBookable
	}
	if !IsNil(o.Formula) {
		toSerialize["formula"] = o.Formula
	}
	if !IsNil(o.FormulaFunctionName) {
		toSerialize["formulaFunctionName"] = o.FormulaFunctionName
	}
	if !IsNil(o.FormulaFunctionArguments) {
		toSerialize["formulaFunctionArguments"] = o.FormulaFunctionArguments
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Catering) {
		toSerialize["catering"] = o.Catering
	}
	if !IsNil(o.PostingRhythm) {
		toSerialize["postingRhythm"] = o.PostingRhythm
	}
	if !IsNil(o.PriceCalculationRule) {
		toSerialize["priceCalculationRule"] = o.PriceCalculationRule
	}
	if !IsNil(o.Ticket) {
		toSerialize["ticket"] = o.Ticket
	}
	if !IsNil(o.InventoryItems) {
		toSerialize["inventoryItems"] = o.InventoryItems
	}
	if !IsNil(o.CalculatedPrice) {
		toSerialize["calculatedPrice"] = o.CalculatedPrice
	}
	return toSerialize, nil
}

type NullableConfigPostingAttributesType struct {
	value *ConfigPostingAttributesType
	isSet bool
}

func (v NullableConfigPostingAttributesType) Get() *ConfigPostingAttributesType {
	return v.value
}

func (v *NullableConfigPostingAttributesType) Set(val *ConfigPostingAttributesType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigPostingAttributesType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigPostingAttributesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigPostingAttributesType(val *ConfigPostingAttributesType) *NullableConfigPostingAttributesType {
	return &NullableConfigPostingAttributesType{value: val, isSet: true}
}

func (v NullableConfigPostingAttributesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigPostingAttributesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


