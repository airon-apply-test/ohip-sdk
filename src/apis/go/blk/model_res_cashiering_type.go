/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the ResCashieringType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResCashieringType{}

// ResCashieringType Cashiering Information for the reservation.
type ResCashieringType struct {
	RevenuesAndBalances *ResRevenueBalanceType `json:"revenuesAndBalances,omitempty"`
	BillingPrivileges *BillingPrivilegesType `json:"billingPrivileges,omitempty"`
	TaxType *ReservationTaxTypeInfo `json:"taxType,omitempty"`
	BedTaxReporting *BedTaxReportingType `json:"bedTaxReporting,omitempty"`
	// This stores the description for the type of tax calculation especially with tax exemption, etc.
	FolioTexts []FolioTextsTypeInner `json:"folioTexts,omitempty"`
	PeriodicFolio *ResPeriodicFolioType `json:"periodicFolio,omitempty"`
	CompAccounting *ResCompAccountingType `json:"compAccounting,omitempty"`
	ReservationPreConfiguredRoutingInstruction *ResPreConfiguredRoutingInstrType `json:"reservationPreConfiguredRoutingInstruction,omitempty"`
	// The guest from whom payment has to be recovered (direct guest).
	FinanciallyResponsible *bool `json:"financiallyResponsible,omitempty"`
	// In case of Appartment style billing indicates whether a prorated amount should be used for an Apartment Style Billing rate.
	ProratedBilling *bool `json:"proratedBilling,omitempty"`
	// Date of the last Room And Tax posting. Used primarily to know the date in case of Advance Billing.
	LastRoomAndTaxPostedDate *string `json:"lastRoomAndTaxPostedDate,omitempty"`
	// This attribute is to verify if reverse check-in is allowed for the reservation.
	ReverseCheckInAllowed *bool `json:"reverseCheckInAllowed,omitempty"`
	// This attribute is to verify if reverse advance check-in is allowed for the reservation.
	ReverseAdvanceCheckInAllowed *bool `json:"reverseAdvanceCheckInAllowed,omitempty"`
	// Specifies whether reservation has a financial transaction associated with it.
	TransactionsPosted *bool `json:"transactionsPosted,omitempty"`
}

// NewResCashieringType instantiates a new ResCashieringType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResCashieringType() *ResCashieringType {
	this := ResCashieringType{}
	return &this
}

// NewResCashieringTypeWithDefaults instantiates a new ResCashieringType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResCashieringTypeWithDefaults() *ResCashieringType {
	this := ResCashieringType{}
	return &this
}

// GetRevenuesAndBalances returns the RevenuesAndBalances field value if set, zero value otherwise.
func (o *ResCashieringType) GetRevenuesAndBalances() ResRevenueBalanceType {
	if o == nil || IsNil(o.RevenuesAndBalances) {
		var ret ResRevenueBalanceType
		return ret
	}
	return *o.RevenuesAndBalances
}

// GetRevenuesAndBalancesOk returns a tuple with the RevenuesAndBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetRevenuesAndBalancesOk() (*ResRevenueBalanceType, bool) {
	if o == nil || IsNil(o.RevenuesAndBalances) {
		return nil, false
	}
	return o.RevenuesAndBalances, true
}

// HasRevenuesAndBalances returns a boolean if a field has been set.
func (o *ResCashieringType) HasRevenuesAndBalances() bool {
	if o != nil && !IsNil(o.RevenuesAndBalances) {
		return true
	}

	return false
}

// SetRevenuesAndBalances gets a reference to the given ResRevenueBalanceType and assigns it to the RevenuesAndBalances field.
func (o *ResCashieringType) SetRevenuesAndBalances(v ResRevenueBalanceType) {
	o.RevenuesAndBalances = &v
}

// GetBillingPrivileges returns the BillingPrivileges field value if set, zero value otherwise.
func (o *ResCashieringType) GetBillingPrivileges() BillingPrivilegesType {
	if o == nil || IsNil(o.BillingPrivileges) {
		var ret BillingPrivilegesType
		return ret
	}
	return *o.BillingPrivileges
}

// GetBillingPrivilegesOk returns a tuple with the BillingPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetBillingPrivilegesOk() (*BillingPrivilegesType, bool) {
	if o == nil || IsNil(o.BillingPrivileges) {
		return nil, false
	}
	return o.BillingPrivileges, true
}

// HasBillingPrivileges returns a boolean if a field has been set.
func (o *ResCashieringType) HasBillingPrivileges() bool {
	if o != nil && !IsNil(o.BillingPrivileges) {
		return true
	}

	return false
}

// SetBillingPrivileges gets a reference to the given BillingPrivilegesType and assigns it to the BillingPrivileges field.
func (o *ResCashieringType) SetBillingPrivileges(v BillingPrivilegesType) {
	o.BillingPrivileges = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *ResCashieringType) GetTaxType() ReservationTaxTypeInfo {
	if o == nil || IsNil(o.TaxType) {
		var ret ReservationTaxTypeInfo
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetTaxTypeOk() (*ReservationTaxTypeInfo, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *ResCashieringType) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given ReservationTaxTypeInfo and assigns it to the TaxType field.
func (o *ResCashieringType) SetTaxType(v ReservationTaxTypeInfo) {
	o.TaxType = &v
}

// GetBedTaxReporting returns the BedTaxReporting field value if set, zero value otherwise.
func (o *ResCashieringType) GetBedTaxReporting() BedTaxReportingType {
	if o == nil || IsNil(o.BedTaxReporting) {
		var ret BedTaxReportingType
		return ret
	}
	return *o.BedTaxReporting
}

// GetBedTaxReportingOk returns a tuple with the BedTaxReporting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetBedTaxReportingOk() (*BedTaxReportingType, bool) {
	if o == nil || IsNil(o.BedTaxReporting) {
		return nil, false
	}
	return o.BedTaxReporting, true
}

// HasBedTaxReporting returns a boolean if a field has been set.
func (o *ResCashieringType) HasBedTaxReporting() bool {
	if o != nil && !IsNil(o.BedTaxReporting) {
		return true
	}

	return false
}

// SetBedTaxReporting gets a reference to the given BedTaxReportingType and assigns it to the BedTaxReporting field.
func (o *ResCashieringType) SetBedTaxReporting(v BedTaxReportingType) {
	o.BedTaxReporting = &v
}

// GetFolioTexts returns the FolioTexts field value if set, zero value otherwise.
func (o *ResCashieringType) GetFolioTexts() []FolioTextsTypeInner {
	if o == nil || IsNil(o.FolioTexts) {
		var ret []FolioTextsTypeInner
		return ret
	}
	return o.FolioTexts
}

// GetFolioTextsOk returns a tuple with the FolioTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetFolioTextsOk() ([]FolioTextsTypeInner, bool) {
	if o == nil || IsNil(o.FolioTexts) {
		return nil, false
	}
	return o.FolioTexts, true
}

// HasFolioTexts returns a boolean if a field has been set.
func (o *ResCashieringType) HasFolioTexts() bool {
	if o != nil && !IsNil(o.FolioTexts) {
		return true
	}

	return false
}

// SetFolioTexts gets a reference to the given []FolioTextsTypeInner and assigns it to the FolioTexts field.
func (o *ResCashieringType) SetFolioTexts(v []FolioTextsTypeInner) {
	o.FolioTexts = v
}

// GetPeriodicFolio returns the PeriodicFolio field value if set, zero value otherwise.
func (o *ResCashieringType) GetPeriodicFolio() ResPeriodicFolioType {
	if o == nil || IsNil(o.PeriodicFolio) {
		var ret ResPeriodicFolioType
		return ret
	}
	return *o.PeriodicFolio
}

// GetPeriodicFolioOk returns a tuple with the PeriodicFolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetPeriodicFolioOk() (*ResPeriodicFolioType, bool) {
	if o == nil || IsNil(o.PeriodicFolio) {
		return nil, false
	}
	return o.PeriodicFolio, true
}

// HasPeriodicFolio returns a boolean if a field has been set.
func (o *ResCashieringType) HasPeriodicFolio() bool {
	if o != nil && !IsNil(o.PeriodicFolio) {
		return true
	}

	return false
}

// SetPeriodicFolio gets a reference to the given ResPeriodicFolioType and assigns it to the PeriodicFolio field.
func (o *ResCashieringType) SetPeriodicFolio(v ResPeriodicFolioType) {
	o.PeriodicFolio = &v
}

// GetCompAccounting returns the CompAccounting field value if set, zero value otherwise.
func (o *ResCashieringType) GetCompAccounting() ResCompAccountingType {
	if o == nil || IsNil(o.CompAccounting) {
		var ret ResCompAccountingType
		return ret
	}
	return *o.CompAccounting
}

// GetCompAccountingOk returns a tuple with the CompAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetCompAccountingOk() (*ResCompAccountingType, bool) {
	if o == nil || IsNil(o.CompAccounting) {
		return nil, false
	}
	return o.CompAccounting, true
}

// HasCompAccounting returns a boolean if a field has been set.
func (o *ResCashieringType) HasCompAccounting() bool {
	if o != nil && !IsNil(o.CompAccounting) {
		return true
	}

	return false
}

// SetCompAccounting gets a reference to the given ResCompAccountingType and assigns it to the CompAccounting field.
func (o *ResCashieringType) SetCompAccounting(v ResCompAccountingType) {
	o.CompAccounting = &v
}

// GetReservationPreConfiguredRoutingInstruction returns the ReservationPreConfiguredRoutingInstruction field value if set, zero value otherwise.
func (o *ResCashieringType) GetReservationPreConfiguredRoutingInstruction() ResPreConfiguredRoutingInstrType {
	if o == nil || IsNil(o.ReservationPreConfiguredRoutingInstruction) {
		var ret ResPreConfiguredRoutingInstrType
		return ret
	}
	return *o.ReservationPreConfiguredRoutingInstruction
}

// GetReservationPreConfiguredRoutingInstructionOk returns a tuple with the ReservationPreConfiguredRoutingInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetReservationPreConfiguredRoutingInstructionOk() (*ResPreConfiguredRoutingInstrType, bool) {
	if o == nil || IsNil(o.ReservationPreConfiguredRoutingInstruction) {
		return nil, false
	}
	return o.ReservationPreConfiguredRoutingInstruction, true
}

// HasReservationPreConfiguredRoutingInstruction returns a boolean if a field has been set.
func (o *ResCashieringType) HasReservationPreConfiguredRoutingInstruction() bool {
	if o != nil && !IsNil(o.ReservationPreConfiguredRoutingInstruction) {
		return true
	}

	return false
}

// SetReservationPreConfiguredRoutingInstruction gets a reference to the given ResPreConfiguredRoutingInstrType and assigns it to the ReservationPreConfiguredRoutingInstruction field.
func (o *ResCashieringType) SetReservationPreConfiguredRoutingInstruction(v ResPreConfiguredRoutingInstrType) {
	o.ReservationPreConfiguredRoutingInstruction = &v
}

// GetFinanciallyResponsible returns the FinanciallyResponsible field value if set, zero value otherwise.
func (o *ResCashieringType) GetFinanciallyResponsible() bool {
	if o == nil || IsNil(o.FinanciallyResponsible) {
		var ret bool
		return ret
	}
	return *o.FinanciallyResponsible
}

// GetFinanciallyResponsibleOk returns a tuple with the FinanciallyResponsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetFinanciallyResponsibleOk() (*bool, bool) {
	if o == nil || IsNil(o.FinanciallyResponsible) {
		return nil, false
	}
	return o.FinanciallyResponsible, true
}

// HasFinanciallyResponsible returns a boolean if a field has been set.
func (o *ResCashieringType) HasFinanciallyResponsible() bool {
	if o != nil && !IsNil(o.FinanciallyResponsible) {
		return true
	}

	return false
}

// SetFinanciallyResponsible gets a reference to the given bool and assigns it to the FinanciallyResponsible field.
func (o *ResCashieringType) SetFinanciallyResponsible(v bool) {
	o.FinanciallyResponsible = &v
}

// GetProratedBilling returns the ProratedBilling field value if set, zero value otherwise.
func (o *ResCashieringType) GetProratedBilling() bool {
	if o == nil || IsNil(o.ProratedBilling) {
		var ret bool
		return ret
	}
	return *o.ProratedBilling
}

// GetProratedBillingOk returns a tuple with the ProratedBilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetProratedBillingOk() (*bool, bool) {
	if o == nil || IsNil(o.ProratedBilling) {
		return nil, false
	}
	return o.ProratedBilling, true
}

// HasProratedBilling returns a boolean if a field has been set.
func (o *ResCashieringType) HasProratedBilling() bool {
	if o != nil && !IsNil(o.ProratedBilling) {
		return true
	}

	return false
}

// SetProratedBilling gets a reference to the given bool and assigns it to the ProratedBilling field.
func (o *ResCashieringType) SetProratedBilling(v bool) {
	o.ProratedBilling = &v
}

// GetLastRoomAndTaxPostedDate returns the LastRoomAndTaxPostedDate field value if set, zero value otherwise.
func (o *ResCashieringType) GetLastRoomAndTaxPostedDate() string {
	if o == nil || IsNil(o.LastRoomAndTaxPostedDate) {
		var ret string
		return ret
	}
	return *o.LastRoomAndTaxPostedDate
}

// GetLastRoomAndTaxPostedDateOk returns a tuple with the LastRoomAndTaxPostedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetLastRoomAndTaxPostedDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastRoomAndTaxPostedDate) {
		return nil, false
	}
	return o.LastRoomAndTaxPostedDate, true
}

// HasLastRoomAndTaxPostedDate returns a boolean if a field has been set.
func (o *ResCashieringType) HasLastRoomAndTaxPostedDate() bool {
	if o != nil && !IsNil(o.LastRoomAndTaxPostedDate) {
		return true
	}

	return false
}

// SetLastRoomAndTaxPostedDate gets a reference to the given string and assigns it to the LastRoomAndTaxPostedDate field.
func (o *ResCashieringType) SetLastRoomAndTaxPostedDate(v string) {
	o.LastRoomAndTaxPostedDate = &v
}

// GetReverseCheckInAllowed returns the ReverseCheckInAllowed field value if set, zero value otherwise.
func (o *ResCashieringType) GetReverseCheckInAllowed() bool {
	if o == nil || IsNil(o.ReverseCheckInAllowed) {
		var ret bool
		return ret
	}
	return *o.ReverseCheckInAllowed
}

// GetReverseCheckInAllowedOk returns a tuple with the ReverseCheckInAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetReverseCheckInAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.ReverseCheckInAllowed) {
		return nil, false
	}
	return o.ReverseCheckInAllowed, true
}

// HasReverseCheckInAllowed returns a boolean if a field has been set.
func (o *ResCashieringType) HasReverseCheckInAllowed() bool {
	if o != nil && !IsNil(o.ReverseCheckInAllowed) {
		return true
	}

	return false
}

// SetReverseCheckInAllowed gets a reference to the given bool and assigns it to the ReverseCheckInAllowed field.
func (o *ResCashieringType) SetReverseCheckInAllowed(v bool) {
	o.ReverseCheckInAllowed = &v
}

// GetReverseAdvanceCheckInAllowed returns the ReverseAdvanceCheckInAllowed field value if set, zero value otherwise.
func (o *ResCashieringType) GetReverseAdvanceCheckInAllowed() bool {
	if o == nil || IsNil(o.ReverseAdvanceCheckInAllowed) {
		var ret bool
		return ret
	}
	return *o.ReverseAdvanceCheckInAllowed
}

// GetReverseAdvanceCheckInAllowedOk returns a tuple with the ReverseAdvanceCheckInAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetReverseAdvanceCheckInAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.ReverseAdvanceCheckInAllowed) {
		return nil, false
	}
	return o.ReverseAdvanceCheckInAllowed, true
}

// HasReverseAdvanceCheckInAllowed returns a boolean if a field has been set.
func (o *ResCashieringType) HasReverseAdvanceCheckInAllowed() bool {
	if o != nil && !IsNil(o.ReverseAdvanceCheckInAllowed) {
		return true
	}

	return false
}

// SetReverseAdvanceCheckInAllowed gets a reference to the given bool and assigns it to the ReverseAdvanceCheckInAllowed field.
func (o *ResCashieringType) SetReverseAdvanceCheckInAllowed(v bool) {
	o.ReverseAdvanceCheckInAllowed = &v
}

// GetTransactionsPosted returns the TransactionsPosted field value if set, zero value otherwise.
func (o *ResCashieringType) GetTransactionsPosted() bool {
	if o == nil || IsNil(o.TransactionsPosted) {
		var ret bool
		return ret
	}
	return *o.TransactionsPosted
}

// GetTransactionsPostedOk returns a tuple with the TransactionsPosted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCashieringType) GetTransactionsPostedOk() (*bool, bool) {
	if o == nil || IsNil(o.TransactionsPosted) {
		return nil, false
	}
	return o.TransactionsPosted, true
}

// HasTransactionsPosted returns a boolean if a field has been set.
func (o *ResCashieringType) HasTransactionsPosted() bool {
	if o != nil && !IsNil(o.TransactionsPosted) {
		return true
	}

	return false
}

// SetTransactionsPosted gets a reference to the given bool and assigns it to the TransactionsPosted field.
func (o *ResCashieringType) SetTransactionsPosted(v bool) {
	o.TransactionsPosted = &v
}

func (o ResCashieringType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResCashieringType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RevenuesAndBalances) {
		toSerialize["revenuesAndBalances"] = o.RevenuesAndBalances
	}
	if !IsNil(o.BillingPrivileges) {
		toSerialize["billingPrivileges"] = o.BillingPrivileges
	}
	if !IsNil(o.TaxType) {
		toSerialize["taxType"] = o.TaxType
	}
	if !IsNil(o.BedTaxReporting) {
		toSerialize["bedTaxReporting"] = o.BedTaxReporting
	}
	if !IsNil(o.FolioTexts) {
		toSerialize["folioTexts"] = o.FolioTexts
	}
	if !IsNil(o.PeriodicFolio) {
		toSerialize["periodicFolio"] = o.PeriodicFolio
	}
	if !IsNil(o.CompAccounting) {
		toSerialize["compAccounting"] = o.CompAccounting
	}
	if !IsNil(o.ReservationPreConfiguredRoutingInstruction) {
		toSerialize["reservationPreConfiguredRoutingInstruction"] = o.ReservationPreConfiguredRoutingInstruction
	}
	if !IsNil(o.FinanciallyResponsible) {
		toSerialize["financiallyResponsible"] = o.FinanciallyResponsible
	}
	if !IsNil(o.ProratedBilling) {
		toSerialize["proratedBilling"] = o.ProratedBilling
	}
	if !IsNil(o.LastRoomAndTaxPostedDate) {
		toSerialize["lastRoomAndTaxPostedDate"] = o.LastRoomAndTaxPostedDate
	}
	if !IsNil(o.ReverseCheckInAllowed) {
		toSerialize["reverseCheckInAllowed"] = o.ReverseCheckInAllowed
	}
	if !IsNil(o.ReverseAdvanceCheckInAllowed) {
		toSerialize["reverseAdvanceCheckInAllowed"] = o.ReverseAdvanceCheckInAllowed
	}
	if !IsNil(o.TransactionsPosted) {
		toSerialize["transactionsPosted"] = o.TransactionsPosted
	}
	return toSerialize, nil
}

type NullableResCashieringType struct {
	value *ResCashieringType
	isSet bool
}

func (v NullableResCashieringType) Get() *ResCashieringType {
	return v.value
}

func (v *NullableResCashieringType) Set(val *ResCashieringType) {
	v.value = val
	v.isSet = true
}

func (v NullableResCashieringType) IsSet() bool {
	return v.isSet
}

func (v *NullableResCashieringType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResCashieringType(val *ResCashieringType) *NullableResCashieringType {
	return &NullableResCashieringType{value: val, isSet: true}
}

func (v NullableResCashieringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResCashieringType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


