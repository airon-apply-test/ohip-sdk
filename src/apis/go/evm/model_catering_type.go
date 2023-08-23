/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CateringType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CateringType{}

// CateringType Contains catering related information for the block.
type CateringType struct {
	CateringStatus *BookingStatusDetailType `json:"cateringStatus,omitempty"`
	// Internal status for the catering.
	CateringInternalStatus *string `json:"cateringInternalStatus,omitempty"`
	// Next catering status of the business block.
	CateringNextStatusList []BookingStatusDetailType `json:"cateringNextStatusList,omitempty"`
	// Collection of catering status history.
	CateringStatusChangeHistory []BookingStatusHistoryType `json:"cateringStatusChangeHistory,omitempty"`
	// Catering Owner of the block.
	CateringOwner *string `json:"cateringOwner,omitempty"`
	EventAttendees *EventAttendeesType `json:"eventAttendees,omitempty"`
	// The name by which the group wishes to be identified in the hotel.
	BoardInfo *string `json:"boardInfo,omitempty"`
	// The name of the customer's in-house representative or contact on the day of the catering event.
	OnSiteName *string `json:"onSiteName,omitempty"`
	// Catering contract number for the block.
	ContractNumber *string `json:"contractNumber,omitempty"`
	// This provides more detail the type of function for which the event is held. For example, Trade Show, Meeting, or Annual Convention.
	FunctionInfo *string `json:"functionInfo,omitempty"`
	// Indicates whether the catering change logging for the booking is active or not. Any change made to catering events and resources will be logged when this is true.
	TrackChanges *bool `json:"trackChanges,omitempty"`
	EventOrder *EventOrderType `json:"eventOrder,omitempty"`
	CateringRevenue *CateringRevenueType `json:"cateringRevenue,omitempty"`
	// Date used by catering manager or coordinator to follow up on the event.
	FollowUpDate *string `json:"followUpDate,omitempty"`
	// Date by which event group must make a decision on the booking.
	DecisionDate *string `json:"decisionDate,omitempty"`
	// Indicates whether event packages or templates can be applied to the block.
	PkgsTmplt *bool `json:"pkgsTmplt,omitempty"`
	CancellationDetails *CancellationDetailsType `json:"cancellationDetails,omitempty"`
	// Indicates the resource discount percentage to be applied to resource items associated with the catering event.
	ResourceDiscountPercentage *float32 `json:"resourceDiscountPercentage,omitempty"`
	// Indicates if the block has any package events.
	HasPackageEvents *bool `json:"hasPackageEvents,omitempty"`
	// Indicates whether BoardInfo changes can be applied to all events of the block.
	ApplyBoardInfoToAllEvents *bool `json:"applyBoardInfoToAllEvents,omitempty"`
	// Indicates whether to ignore any warning during applying the changes to the events associated with the current block.
	OverrideEventsProcessingWarnings *bool `json:"overrideEventsProcessingWarnings,omitempty"`
	// Indicates whether Guarantee changes can be applied to all events of the block.
	ApplyEventsGuaranteeToAllEvents *bool `json:"applyEventsGuaranteeToAllEvents,omitempty"`
	ApplyEventAttendeesChangesToEvents *ApplyEventAttendeesChangesToEventsType `json:"applyEventAttendeesChangesToEvents,omitempty"`
	ResourceDiscountType *ResourceDiscountTypeType `json:"resourceDiscountType,omitempty"`
}

// NewCateringType instantiates a new CateringType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCateringType() *CateringType {
	this := CateringType{}
	return &this
}

// NewCateringTypeWithDefaults instantiates a new CateringType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCateringTypeWithDefaults() *CateringType {
	this := CateringType{}
	return &this
}

// GetCateringStatus returns the CateringStatus field value if set, zero value otherwise.
func (o *CateringType) GetCateringStatus() BookingStatusDetailType {
	if o == nil || IsNil(o.CateringStatus) {
		var ret BookingStatusDetailType
		return ret
	}
	return *o.CateringStatus
}

// GetCateringStatusOk returns a tuple with the CateringStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetCateringStatusOk() (*BookingStatusDetailType, bool) {
	if o == nil || IsNil(o.CateringStatus) {
		return nil, false
	}
	return o.CateringStatus, true
}

// HasCateringStatus returns a boolean if a field has been set.
func (o *CateringType) HasCateringStatus() bool {
	if o != nil && !IsNil(o.CateringStatus) {
		return true
	}

	return false
}

// SetCateringStatus gets a reference to the given BookingStatusDetailType and assigns it to the CateringStatus field.
func (o *CateringType) SetCateringStatus(v BookingStatusDetailType) {
	o.CateringStatus = &v
}

// GetCateringInternalStatus returns the CateringInternalStatus field value if set, zero value otherwise.
func (o *CateringType) GetCateringInternalStatus() string {
	if o == nil || IsNil(o.CateringInternalStatus) {
		var ret string
		return ret
	}
	return *o.CateringInternalStatus
}

// GetCateringInternalStatusOk returns a tuple with the CateringInternalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetCateringInternalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CateringInternalStatus) {
		return nil, false
	}
	return o.CateringInternalStatus, true
}

// HasCateringInternalStatus returns a boolean if a field has been set.
func (o *CateringType) HasCateringInternalStatus() bool {
	if o != nil && !IsNil(o.CateringInternalStatus) {
		return true
	}

	return false
}

// SetCateringInternalStatus gets a reference to the given string and assigns it to the CateringInternalStatus field.
func (o *CateringType) SetCateringInternalStatus(v string) {
	o.CateringInternalStatus = &v
}

// GetCateringNextStatusList returns the CateringNextStatusList field value if set, zero value otherwise.
func (o *CateringType) GetCateringNextStatusList() []BookingStatusDetailType {
	if o == nil || IsNil(o.CateringNextStatusList) {
		var ret []BookingStatusDetailType
		return ret
	}
	return o.CateringNextStatusList
}

// GetCateringNextStatusListOk returns a tuple with the CateringNextStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetCateringNextStatusListOk() ([]BookingStatusDetailType, bool) {
	if o == nil || IsNil(o.CateringNextStatusList) {
		return nil, false
	}
	return o.CateringNextStatusList, true
}

// HasCateringNextStatusList returns a boolean if a field has been set.
func (o *CateringType) HasCateringNextStatusList() bool {
	if o != nil && !IsNil(o.CateringNextStatusList) {
		return true
	}

	return false
}

// SetCateringNextStatusList gets a reference to the given []BookingStatusDetailType and assigns it to the CateringNextStatusList field.
func (o *CateringType) SetCateringNextStatusList(v []BookingStatusDetailType) {
	o.CateringNextStatusList = v
}

// GetCateringStatusChangeHistory returns the CateringStatusChangeHistory field value if set, zero value otherwise.
func (o *CateringType) GetCateringStatusChangeHistory() []BookingStatusHistoryType {
	if o == nil || IsNil(o.CateringStatusChangeHistory) {
		var ret []BookingStatusHistoryType
		return ret
	}
	return o.CateringStatusChangeHistory
}

// GetCateringStatusChangeHistoryOk returns a tuple with the CateringStatusChangeHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetCateringStatusChangeHistoryOk() ([]BookingStatusHistoryType, bool) {
	if o == nil || IsNil(o.CateringStatusChangeHistory) {
		return nil, false
	}
	return o.CateringStatusChangeHistory, true
}

// HasCateringStatusChangeHistory returns a boolean if a field has been set.
func (o *CateringType) HasCateringStatusChangeHistory() bool {
	if o != nil && !IsNil(o.CateringStatusChangeHistory) {
		return true
	}

	return false
}

// SetCateringStatusChangeHistory gets a reference to the given []BookingStatusHistoryType and assigns it to the CateringStatusChangeHistory field.
func (o *CateringType) SetCateringStatusChangeHistory(v []BookingStatusHistoryType) {
	o.CateringStatusChangeHistory = v
}

// GetCateringOwner returns the CateringOwner field value if set, zero value otherwise.
func (o *CateringType) GetCateringOwner() string {
	if o == nil || IsNil(o.CateringOwner) {
		var ret string
		return ret
	}
	return *o.CateringOwner
}

// GetCateringOwnerOk returns a tuple with the CateringOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetCateringOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.CateringOwner) {
		return nil, false
	}
	return o.CateringOwner, true
}

// HasCateringOwner returns a boolean if a field has been set.
func (o *CateringType) HasCateringOwner() bool {
	if o != nil && !IsNil(o.CateringOwner) {
		return true
	}

	return false
}

// SetCateringOwner gets a reference to the given string and assigns it to the CateringOwner field.
func (o *CateringType) SetCateringOwner(v string) {
	o.CateringOwner = &v
}

// GetEventAttendees returns the EventAttendees field value if set, zero value otherwise.
func (o *CateringType) GetEventAttendees() EventAttendeesType {
	if o == nil || IsNil(o.EventAttendees) {
		var ret EventAttendeesType
		return ret
	}
	return *o.EventAttendees
}

// GetEventAttendeesOk returns a tuple with the EventAttendees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetEventAttendeesOk() (*EventAttendeesType, bool) {
	if o == nil || IsNil(o.EventAttendees) {
		return nil, false
	}
	return o.EventAttendees, true
}

// HasEventAttendees returns a boolean if a field has been set.
func (o *CateringType) HasEventAttendees() bool {
	if o != nil && !IsNil(o.EventAttendees) {
		return true
	}

	return false
}

// SetEventAttendees gets a reference to the given EventAttendeesType and assigns it to the EventAttendees field.
func (o *CateringType) SetEventAttendees(v EventAttendeesType) {
	o.EventAttendees = &v
}

// GetBoardInfo returns the BoardInfo field value if set, zero value otherwise.
func (o *CateringType) GetBoardInfo() string {
	if o == nil || IsNil(o.BoardInfo) {
		var ret string
		return ret
	}
	return *o.BoardInfo
}

// GetBoardInfoOk returns a tuple with the BoardInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetBoardInfoOk() (*string, bool) {
	if o == nil || IsNil(o.BoardInfo) {
		return nil, false
	}
	return o.BoardInfo, true
}

// HasBoardInfo returns a boolean if a field has been set.
func (o *CateringType) HasBoardInfo() bool {
	if o != nil && !IsNil(o.BoardInfo) {
		return true
	}

	return false
}

// SetBoardInfo gets a reference to the given string and assigns it to the BoardInfo field.
func (o *CateringType) SetBoardInfo(v string) {
	o.BoardInfo = &v
}

// GetOnSiteName returns the OnSiteName field value if set, zero value otherwise.
func (o *CateringType) GetOnSiteName() string {
	if o == nil || IsNil(o.OnSiteName) {
		var ret string
		return ret
	}
	return *o.OnSiteName
}

// GetOnSiteNameOk returns a tuple with the OnSiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetOnSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.OnSiteName) {
		return nil, false
	}
	return o.OnSiteName, true
}

// HasOnSiteName returns a boolean if a field has been set.
func (o *CateringType) HasOnSiteName() bool {
	if o != nil && !IsNil(o.OnSiteName) {
		return true
	}

	return false
}

// SetOnSiteName gets a reference to the given string and assigns it to the OnSiteName field.
func (o *CateringType) SetOnSiteName(v string) {
	o.OnSiteName = &v
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *CateringType) GetContractNumber() string {
	if o == nil || IsNil(o.ContractNumber) {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetContractNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContractNumber) {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *CateringType) HasContractNumber() bool {
	if o != nil && !IsNil(o.ContractNumber) {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *CateringType) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetFunctionInfo returns the FunctionInfo field value if set, zero value otherwise.
func (o *CateringType) GetFunctionInfo() string {
	if o == nil || IsNil(o.FunctionInfo) {
		var ret string
		return ret
	}
	return *o.FunctionInfo
}

// GetFunctionInfoOk returns a tuple with the FunctionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetFunctionInfoOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionInfo) {
		return nil, false
	}
	return o.FunctionInfo, true
}

// HasFunctionInfo returns a boolean if a field has been set.
func (o *CateringType) HasFunctionInfo() bool {
	if o != nil && !IsNil(o.FunctionInfo) {
		return true
	}

	return false
}

// SetFunctionInfo gets a reference to the given string and assigns it to the FunctionInfo field.
func (o *CateringType) SetFunctionInfo(v string) {
	o.FunctionInfo = &v
}

// GetTrackChanges returns the TrackChanges field value if set, zero value otherwise.
func (o *CateringType) GetTrackChanges() bool {
	if o == nil || IsNil(o.TrackChanges) {
		var ret bool
		return ret
	}
	return *o.TrackChanges
}

// GetTrackChangesOk returns a tuple with the TrackChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetTrackChangesOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackChanges) {
		return nil, false
	}
	return o.TrackChanges, true
}

// HasTrackChanges returns a boolean if a field has been set.
func (o *CateringType) HasTrackChanges() bool {
	if o != nil && !IsNil(o.TrackChanges) {
		return true
	}

	return false
}

// SetTrackChanges gets a reference to the given bool and assigns it to the TrackChanges field.
func (o *CateringType) SetTrackChanges(v bool) {
	o.TrackChanges = &v
}

// GetEventOrder returns the EventOrder field value if set, zero value otherwise.
func (o *CateringType) GetEventOrder() EventOrderType {
	if o == nil || IsNil(o.EventOrder) {
		var ret EventOrderType
		return ret
	}
	return *o.EventOrder
}

// GetEventOrderOk returns a tuple with the EventOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetEventOrderOk() (*EventOrderType, bool) {
	if o == nil || IsNil(o.EventOrder) {
		return nil, false
	}
	return o.EventOrder, true
}

// HasEventOrder returns a boolean if a field has been set.
func (o *CateringType) HasEventOrder() bool {
	if o != nil && !IsNil(o.EventOrder) {
		return true
	}

	return false
}

// SetEventOrder gets a reference to the given EventOrderType and assigns it to the EventOrder field.
func (o *CateringType) SetEventOrder(v EventOrderType) {
	o.EventOrder = &v
}

// GetCateringRevenue returns the CateringRevenue field value if set, zero value otherwise.
func (o *CateringType) GetCateringRevenue() CateringRevenueType {
	if o == nil || IsNil(o.CateringRevenue) {
		var ret CateringRevenueType
		return ret
	}
	return *o.CateringRevenue
}

// GetCateringRevenueOk returns a tuple with the CateringRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetCateringRevenueOk() (*CateringRevenueType, bool) {
	if o == nil || IsNil(o.CateringRevenue) {
		return nil, false
	}
	return o.CateringRevenue, true
}

// HasCateringRevenue returns a boolean if a field has been set.
func (o *CateringType) HasCateringRevenue() bool {
	if o != nil && !IsNil(o.CateringRevenue) {
		return true
	}

	return false
}

// SetCateringRevenue gets a reference to the given CateringRevenueType and assigns it to the CateringRevenue field.
func (o *CateringType) SetCateringRevenue(v CateringRevenueType) {
	o.CateringRevenue = &v
}

// GetFollowUpDate returns the FollowUpDate field value if set, zero value otherwise.
func (o *CateringType) GetFollowUpDate() string {
	if o == nil || IsNil(o.FollowUpDate) {
		var ret string
		return ret
	}
	return *o.FollowUpDate
}

// GetFollowUpDateOk returns a tuple with the FollowUpDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetFollowUpDateOk() (*string, bool) {
	if o == nil || IsNil(o.FollowUpDate) {
		return nil, false
	}
	return o.FollowUpDate, true
}

// HasFollowUpDate returns a boolean if a field has been set.
func (o *CateringType) HasFollowUpDate() bool {
	if o != nil && !IsNil(o.FollowUpDate) {
		return true
	}

	return false
}

// SetFollowUpDate gets a reference to the given string and assigns it to the FollowUpDate field.
func (o *CateringType) SetFollowUpDate(v string) {
	o.FollowUpDate = &v
}

// GetDecisionDate returns the DecisionDate field value if set, zero value otherwise.
func (o *CateringType) GetDecisionDate() string {
	if o == nil || IsNil(o.DecisionDate) {
		var ret string
		return ret
	}
	return *o.DecisionDate
}

// GetDecisionDateOk returns a tuple with the DecisionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetDecisionDateOk() (*string, bool) {
	if o == nil || IsNil(o.DecisionDate) {
		return nil, false
	}
	return o.DecisionDate, true
}

// HasDecisionDate returns a boolean if a field has been set.
func (o *CateringType) HasDecisionDate() bool {
	if o != nil && !IsNil(o.DecisionDate) {
		return true
	}

	return false
}

// SetDecisionDate gets a reference to the given string and assigns it to the DecisionDate field.
func (o *CateringType) SetDecisionDate(v string) {
	o.DecisionDate = &v
}

// GetPkgsTmplt returns the PkgsTmplt field value if set, zero value otherwise.
func (o *CateringType) GetPkgsTmplt() bool {
	if o == nil || IsNil(o.PkgsTmplt) {
		var ret bool
		return ret
	}
	return *o.PkgsTmplt
}

// GetPkgsTmpltOk returns a tuple with the PkgsTmplt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetPkgsTmpltOk() (*bool, bool) {
	if o == nil || IsNil(o.PkgsTmplt) {
		return nil, false
	}
	return o.PkgsTmplt, true
}

// HasPkgsTmplt returns a boolean if a field has been set.
func (o *CateringType) HasPkgsTmplt() bool {
	if o != nil && !IsNil(o.PkgsTmplt) {
		return true
	}

	return false
}

// SetPkgsTmplt gets a reference to the given bool and assigns it to the PkgsTmplt field.
func (o *CateringType) SetPkgsTmplt(v bool) {
	o.PkgsTmplt = &v
}

// GetCancellationDetails returns the CancellationDetails field value if set, zero value otherwise.
func (o *CateringType) GetCancellationDetails() CancellationDetailsType {
	if o == nil || IsNil(o.CancellationDetails) {
		var ret CancellationDetailsType
		return ret
	}
	return *o.CancellationDetails
}

// GetCancellationDetailsOk returns a tuple with the CancellationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetCancellationDetailsOk() (*CancellationDetailsType, bool) {
	if o == nil || IsNil(o.CancellationDetails) {
		return nil, false
	}
	return o.CancellationDetails, true
}

// HasCancellationDetails returns a boolean if a field has been set.
func (o *CateringType) HasCancellationDetails() bool {
	if o != nil && !IsNil(o.CancellationDetails) {
		return true
	}

	return false
}

// SetCancellationDetails gets a reference to the given CancellationDetailsType and assigns it to the CancellationDetails field.
func (o *CateringType) SetCancellationDetails(v CancellationDetailsType) {
	o.CancellationDetails = &v
}

// GetResourceDiscountPercentage returns the ResourceDiscountPercentage field value if set, zero value otherwise.
func (o *CateringType) GetResourceDiscountPercentage() float32 {
	if o == nil || IsNil(o.ResourceDiscountPercentage) {
		var ret float32
		return ret
	}
	return *o.ResourceDiscountPercentage
}

// GetResourceDiscountPercentageOk returns a tuple with the ResourceDiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetResourceDiscountPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.ResourceDiscountPercentage) {
		return nil, false
	}
	return o.ResourceDiscountPercentage, true
}

// HasResourceDiscountPercentage returns a boolean if a field has been set.
func (o *CateringType) HasResourceDiscountPercentage() bool {
	if o != nil && !IsNil(o.ResourceDiscountPercentage) {
		return true
	}

	return false
}

// SetResourceDiscountPercentage gets a reference to the given float32 and assigns it to the ResourceDiscountPercentage field.
func (o *CateringType) SetResourceDiscountPercentage(v float32) {
	o.ResourceDiscountPercentage = &v
}

// GetHasPackageEvents returns the HasPackageEvents field value if set, zero value otherwise.
func (o *CateringType) GetHasPackageEvents() bool {
	if o == nil || IsNil(o.HasPackageEvents) {
		var ret bool
		return ret
	}
	return *o.HasPackageEvents
}

// GetHasPackageEventsOk returns a tuple with the HasPackageEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetHasPackageEventsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPackageEvents) {
		return nil, false
	}
	return o.HasPackageEvents, true
}

// HasHasPackageEvents returns a boolean if a field has been set.
func (o *CateringType) HasHasPackageEvents() bool {
	if o != nil && !IsNil(o.HasPackageEvents) {
		return true
	}

	return false
}

// SetHasPackageEvents gets a reference to the given bool and assigns it to the HasPackageEvents field.
func (o *CateringType) SetHasPackageEvents(v bool) {
	o.HasPackageEvents = &v
}

// GetApplyBoardInfoToAllEvents returns the ApplyBoardInfoToAllEvents field value if set, zero value otherwise.
func (o *CateringType) GetApplyBoardInfoToAllEvents() bool {
	if o == nil || IsNil(o.ApplyBoardInfoToAllEvents) {
		var ret bool
		return ret
	}
	return *o.ApplyBoardInfoToAllEvents
}

// GetApplyBoardInfoToAllEventsOk returns a tuple with the ApplyBoardInfoToAllEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetApplyBoardInfoToAllEventsOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyBoardInfoToAllEvents) {
		return nil, false
	}
	return o.ApplyBoardInfoToAllEvents, true
}

// HasApplyBoardInfoToAllEvents returns a boolean if a field has been set.
func (o *CateringType) HasApplyBoardInfoToAllEvents() bool {
	if o != nil && !IsNil(o.ApplyBoardInfoToAllEvents) {
		return true
	}

	return false
}

// SetApplyBoardInfoToAllEvents gets a reference to the given bool and assigns it to the ApplyBoardInfoToAllEvents field.
func (o *CateringType) SetApplyBoardInfoToAllEvents(v bool) {
	o.ApplyBoardInfoToAllEvents = &v
}

// GetOverrideEventsProcessingWarnings returns the OverrideEventsProcessingWarnings field value if set, zero value otherwise.
func (o *CateringType) GetOverrideEventsProcessingWarnings() bool {
	if o == nil || IsNil(o.OverrideEventsProcessingWarnings) {
		var ret bool
		return ret
	}
	return *o.OverrideEventsProcessingWarnings
}

// GetOverrideEventsProcessingWarningsOk returns a tuple with the OverrideEventsProcessingWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetOverrideEventsProcessingWarningsOk() (*bool, bool) {
	if o == nil || IsNil(o.OverrideEventsProcessingWarnings) {
		return nil, false
	}
	return o.OverrideEventsProcessingWarnings, true
}

// HasOverrideEventsProcessingWarnings returns a boolean if a field has been set.
func (o *CateringType) HasOverrideEventsProcessingWarnings() bool {
	if o != nil && !IsNil(o.OverrideEventsProcessingWarnings) {
		return true
	}

	return false
}

// SetOverrideEventsProcessingWarnings gets a reference to the given bool and assigns it to the OverrideEventsProcessingWarnings field.
func (o *CateringType) SetOverrideEventsProcessingWarnings(v bool) {
	o.OverrideEventsProcessingWarnings = &v
}

// GetApplyEventsGuaranteeToAllEvents returns the ApplyEventsGuaranteeToAllEvents field value if set, zero value otherwise.
func (o *CateringType) GetApplyEventsGuaranteeToAllEvents() bool {
	if o == nil || IsNil(o.ApplyEventsGuaranteeToAllEvents) {
		var ret bool
		return ret
	}
	return *o.ApplyEventsGuaranteeToAllEvents
}

// GetApplyEventsGuaranteeToAllEventsOk returns a tuple with the ApplyEventsGuaranteeToAllEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetApplyEventsGuaranteeToAllEventsOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyEventsGuaranteeToAllEvents) {
		return nil, false
	}
	return o.ApplyEventsGuaranteeToAllEvents, true
}

// HasApplyEventsGuaranteeToAllEvents returns a boolean if a field has been set.
func (o *CateringType) HasApplyEventsGuaranteeToAllEvents() bool {
	if o != nil && !IsNil(o.ApplyEventsGuaranteeToAllEvents) {
		return true
	}

	return false
}

// SetApplyEventsGuaranteeToAllEvents gets a reference to the given bool and assigns it to the ApplyEventsGuaranteeToAllEvents field.
func (o *CateringType) SetApplyEventsGuaranteeToAllEvents(v bool) {
	o.ApplyEventsGuaranteeToAllEvents = &v
}

// GetApplyEventAttendeesChangesToEvents returns the ApplyEventAttendeesChangesToEvents field value if set, zero value otherwise.
func (o *CateringType) GetApplyEventAttendeesChangesToEvents() ApplyEventAttendeesChangesToEventsType {
	if o == nil || IsNil(o.ApplyEventAttendeesChangesToEvents) {
		var ret ApplyEventAttendeesChangesToEventsType
		return ret
	}
	return *o.ApplyEventAttendeesChangesToEvents
}

// GetApplyEventAttendeesChangesToEventsOk returns a tuple with the ApplyEventAttendeesChangesToEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetApplyEventAttendeesChangesToEventsOk() (*ApplyEventAttendeesChangesToEventsType, bool) {
	if o == nil || IsNil(o.ApplyEventAttendeesChangesToEvents) {
		return nil, false
	}
	return o.ApplyEventAttendeesChangesToEvents, true
}

// HasApplyEventAttendeesChangesToEvents returns a boolean if a field has been set.
func (o *CateringType) HasApplyEventAttendeesChangesToEvents() bool {
	if o != nil && !IsNil(o.ApplyEventAttendeesChangesToEvents) {
		return true
	}

	return false
}

// SetApplyEventAttendeesChangesToEvents gets a reference to the given ApplyEventAttendeesChangesToEventsType and assigns it to the ApplyEventAttendeesChangesToEvents field.
func (o *CateringType) SetApplyEventAttendeesChangesToEvents(v ApplyEventAttendeesChangesToEventsType) {
	o.ApplyEventAttendeesChangesToEvents = &v
}

// GetResourceDiscountType returns the ResourceDiscountType field value if set, zero value otherwise.
func (o *CateringType) GetResourceDiscountType() ResourceDiscountTypeType {
	if o == nil || IsNil(o.ResourceDiscountType) {
		var ret ResourceDiscountTypeType
		return ret
	}
	return *o.ResourceDiscountType
}

// GetResourceDiscountTypeOk returns a tuple with the ResourceDiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringType) GetResourceDiscountTypeOk() (*ResourceDiscountTypeType, bool) {
	if o == nil || IsNil(o.ResourceDiscountType) {
		return nil, false
	}
	return o.ResourceDiscountType, true
}

// HasResourceDiscountType returns a boolean if a field has been set.
func (o *CateringType) HasResourceDiscountType() bool {
	if o != nil && !IsNil(o.ResourceDiscountType) {
		return true
	}

	return false
}

// SetResourceDiscountType gets a reference to the given ResourceDiscountTypeType and assigns it to the ResourceDiscountType field.
func (o *CateringType) SetResourceDiscountType(v ResourceDiscountTypeType) {
	o.ResourceDiscountType = &v
}

func (o CateringType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CateringType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CateringStatus) {
		toSerialize["cateringStatus"] = o.CateringStatus
	}
	if !IsNil(o.CateringInternalStatus) {
		toSerialize["cateringInternalStatus"] = o.CateringInternalStatus
	}
	if !IsNil(o.CateringNextStatusList) {
		toSerialize["cateringNextStatusList"] = o.CateringNextStatusList
	}
	if !IsNil(o.CateringStatusChangeHistory) {
		toSerialize["cateringStatusChangeHistory"] = o.CateringStatusChangeHistory
	}
	if !IsNil(o.CateringOwner) {
		toSerialize["cateringOwner"] = o.CateringOwner
	}
	if !IsNil(o.EventAttendees) {
		toSerialize["eventAttendees"] = o.EventAttendees
	}
	if !IsNil(o.BoardInfo) {
		toSerialize["boardInfo"] = o.BoardInfo
	}
	if !IsNil(o.OnSiteName) {
		toSerialize["onSiteName"] = o.OnSiteName
	}
	if !IsNil(o.ContractNumber) {
		toSerialize["contractNumber"] = o.ContractNumber
	}
	if !IsNil(o.FunctionInfo) {
		toSerialize["functionInfo"] = o.FunctionInfo
	}
	if !IsNil(o.TrackChanges) {
		toSerialize["trackChanges"] = o.TrackChanges
	}
	if !IsNil(o.EventOrder) {
		toSerialize["eventOrder"] = o.EventOrder
	}
	if !IsNil(o.CateringRevenue) {
		toSerialize["cateringRevenue"] = o.CateringRevenue
	}
	if !IsNil(o.FollowUpDate) {
		toSerialize["followUpDate"] = o.FollowUpDate
	}
	if !IsNil(o.DecisionDate) {
		toSerialize["decisionDate"] = o.DecisionDate
	}
	if !IsNil(o.PkgsTmplt) {
		toSerialize["pkgsTmplt"] = o.PkgsTmplt
	}
	if !IsNil(o.CancellationDetails) {
		toSerialize["cancellationDetails"] = o.CancellationDetails
	}
	if !IsNil(o.ResourceDiscountPercentage) {
		toSerialize["resourceDiscountPercentage"] = o.ResourceDiscountPercentage
	}
	if !IsNil(o.HasPackageEvents) {
		toSerialize["hasPackageEvents"] = o.HasPackageEvents
	}
	if !IsNil(o.ApplyBoardInfoToAllEvents) {
		toSerialize["applyBoardInfoToAllEvents"] = o.ApplyBoardInfoToAllEvents
	}
	if !IsNil(o.OverrideEventsProcessingWarnings) {
		toSerialize["overrideEventsProcessingWarnings"] = o.OverrideEventsProcessingWarnings
	}
	if !IsNil(o.ApplyEventsGuaranteeToAllEvents) {
		toSerialize["applyEventsGuaranteeToAllEvents"] = o.ApplyEventsGuaranteeToAllEvents
	}
	if !IsNil(o.ApplyEventAttendeesChangesToEvents) {
		toSerialize["applyEventAttendeesChangesToEvents"] = o.ApplyEventAttendeesChangesToEvents
	}
	if !IsNil(o.ResourceDiscountType) {
		toSerialize["resourceDiscountType"] = o.ResourceDiscountType
	}
	return toSerialize, nil
}

type NullableCateringType struct {
	value *CateringType
	isSet bool
}

func (v NullableCateringType) Get() *CateringType {
	return v.value
}

func (v *NullableCateringType) Set(val *CateringType) {
	v.value = val
	v.isSet = true
}

func (v NullableCateringType) IsSet() bool {
	return v.isSet
}

func (v *NullableCateringType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCateringType(val *CateringType) *NullableCateringType {
	return &NullableCateringType{value: val, isSet: true}
}

func (v NullableCateringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCateringType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


