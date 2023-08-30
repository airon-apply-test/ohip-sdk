/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
	"time"
)

// checks if the BlockType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockType{}

// BlockType struct for BlockType
type BlockType struct {
	// Unique Id that references an object uniquely in the system.
	BlockIdList []UniqueIDType `json:"blockIdList,omitempty"`
	// This type contains unique information of external reference.
	ExternalReferences []ExternalReferenceType `json:"externalReferences,omitempty"`
	BlockDetails *BlockDetailsType `json:"blockDetails,omitempty"`
	BlockOwners *BlockOwnersType `json:"blockOwners,omitempty"`
	Catering *CateringType `json:"catering,omitempty"`
	BlockProfiles *BlockTypeBlockProfiles `json:"blockProfiles,omitempty"`
	// Inventory item attached to a block.
	InventoryItems []BlockInventoryItemType `json:"inventoryItems,omitempty"`
	BlockStatistics *BlockStatisticsType `json:"blockStatistics,omitempty"`
	Comments *BlockTypeComments `json:"comments,omitempty"`
	// Collection of lamp indicators.
	BlockIndicators []IndicatorType `json:"blockIndicators,omitempty"`
	// List of Block traces.
	Traces []BlockTraceType `json:"traces,omitempty"`
	// Statistics summary information including Rooms Sold, Room Revenue, Food and Beverage Revenue, Average Room Rate,etc.
	Statistics []BlockStatisticsSummaryType `json:"statistics,omitempty"`
	BlockSecurity *BlockTypeBlockSecurity `json:"blockSecurity,omitempty"`
	SellMessages *SellMessageConfigsType `json:"sellMessages,omitempty"`
	// Attachment List.
	Attachments []AttachmentType `json:"attachments,omitempty"`
	// Wash Schedule attached to the block, which allows to release inventory based on selected wash schedule type.
	WashSchedules []BlockWashScheduleType `json:"washSchedules,omitempty"`
	Restrictions *BlockRestrictionsType `json:"restrictions,omitempty"`
	HotelId *string `json:"hotelId,omitempty"`
	// Mark this block as recently accessed.
	MarkAsRecentlyAccessed *bool `json:"markAsRecentlyAccessed,omitempty"`
	// Time stamp of the creation.
	CreateDateTime *time.Time `json:"createDateTime,omitempty"`
	// ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
	CreatorId *string `json:"creatorId,omitempty"`
	// Time stamp of last modification.
	LastModifyDateTime *time.Time `json:"lastModifyDateTime,omitempty"`
	// Identifies the last software system or person to modify a record.
	LastModifierId *string `json:"lastModifierId,omitempty"`
}

// NewBlockType instantiates a new BlockType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockType() *BlockType {
	this := BlockType{}
	return &this
}

// NewBlockTypeWithDefaults instantiates a new BlockType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTypeWithDefaults() *BlockType {
	this := BlockType{}
	return &this
}

// GetBlockIdList returns the BlockIdList field value if set, zero value otherwise.
func (o *BlockType) GetBlockIdList() []UniqueIDType {
	if o == nil || IsNil(o.BlockIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.BlockIdList
}

// GetBlockIdListOk returns a tuple with the BlockIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetBlockIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.BlockIdList) {
		return nil, false
	}
	return o.BlockIdList, true
}

// HasBlockIdList returns a boolean if a field has been set.
func (o *BlockType) HasBlockIdList() bool {
	if o != nil && !IsNil(o.BlockIdList) {
		return true
	}

	return false
}

// SetBlockIdList gets a reference to the given []UniqueIDType and assigns it to the BlockIdList field.
func (o *BlockType) SetBlockIdList(v []UniqueIDType) {
	o.BlockIdList = v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *BlockType) GetExternalReferences() []ExternalReferenceType {
	if o == nil || IsNil(o.ExternalReferences) {
		var ret []ExternalReferenceType
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetExternalReferencesOk() ([]ExternalReferenceType, bool) {
	if o == nil || IsNil(o.ExternalReferences) {
		return nil, false
	}
	return o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *BlockType) HasExternalReferences() bool {
	if o != nil && !IsNil(o.ExternalReferences) {
		return true
	}

	return false
}

// SetExternalReferences gets a reference to the given []ExternalReferenceType and assigns it to the ExternalReferences field.
func (o *BlockType) SetExternalReferences(v []ExternalReferenceType) {
	o.ExternalReferences = v
}

// GetBlockDetails returns the BlockDetails field value if set, zero value otherwise.
func (o *BlockType) GetBlockDetails() BlockDetailsType {
	if o == nil || IsNil(o.BlockDetails) {
		var ret BlockDetailsType
		return ret
	}
	return *o.BlockDetails
}

// GetBlockDetailsOk returns a tuple with the BlockDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetBlockDetailsOk() (*BlockDetailsType, bool) {
	if o == nil || IsNil(o.BlockDetails) {
		return nil, false
	}
	return o.BlockDetails, true
}

// HasBlockDetails returns a boolean if a field has been set.
func (o *BlockType) HasBlockDetails() bool {
	if o != nil && !IsNil(o.BlockDetails) {
		return true
	}

	return false
}

// SetBlockDetails gets a reference to the given BlockDetailsType and assigns it to the BlockDetails field.
func (o *BlockType) SetBlockDetails(v BlockDetailsType) {
	o.BlockDetails = &v
}

// GetBlockOwners returns the BlockOwners field value if set, zero value otherwise.
func (o *BlockType) GetBlockOwners() BlockOwnersType {
	if o == nil || IsNil(o.BlockOwners) {
		var ret BlockOwnersType
		return ret
	}
	return *o.BlockOwners
}

// GetBlockOwnersOk returns a tuple with the BlockOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetBlockOwnersOk() (*BlockOwnersType, bool) {
	if o == nil || IsNil(o.BlockOwners) {
		return nil, false
	}
	return o.BlockOwners, true
}

// HasBlockOwners returns a boolean if a field has been set.
func (o *BlockType) HasBlockOwners() bool {
	if o != nil && !IsNil(o.BlockOwners) {
		return true
	}

	return false
}

// SetBlockOwners gets a reference to the given BlockOwnersType and assigns it to the BlockOwners field.
func (o *BlockType) SetBlockOwners(v BlockOwnersType) {
	o.BlockOwners = &v
}

// GetCatering returns the Catering field value if set, zero value otherwise.
func (o *BlockType) GetCatering() CateringType {
	if o == nil || IsNil(o.Catering) {
		var ret CateringType
		return ret
	}
	return *o.Catering
}

// GetCateringOk returns a tuple with the Catering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetCateringOk() (*CateringType, bool) {
	if o == nil || IsNil(o.Catering) {
		return nil, false
	}
	return o.Catering, true
}

// HasCatering returns a boolean if a field has been set.
func (o *BlockType) HasCatering() bool {
	if o != nil && !IsNil(o.Catering) {
		return true
	}

	return false
}

// SetCatering gets a reference to the given CateringType and assigns it to the Catering field.
func (o *BlockType) SetCatering(v CateringType) {
	o.Catering = &v
}

// GetBlockProfiles returns the BlockProfiles field value if set, zero value otherwise.
func (o *BlockType) GetBlockProfiles() BlockTypeBlockProfiles {
	if o == nil || IsNil(o.BlockProfiles) {
		var ret BlockTypeBlockProfiles
		return ret
	}
	return *o.BlockProfiles
}

// GetBlockProfilesOk returns a tuple with the BlockProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetBlockProfilesOk() (*BlockTypeBlockProfiles, bool) {
	if o == nil || IsNil(o.BlockProfiles) {
		return nil, false
	}
	return o.BlockProfiles, true
}

// HasBlockProfiles returns a boolean if a field has been set.
func (o *BlockType) HasBlockProfiles() bool {
	if o != nil && !IsNil(o.BlockProfiles) {
		return true
	}

	return false
}

// SetBlockProfiles gets a reference to the given BlockTypeBlockProfiles and assigns it to the BlockProfiles field.
func (o *BlockType) SetBlockProfiles(v BlockTypeBlockProfiles) {
	o.BlockProfiles = &v
}

// GetInventoryItems returns the InventoryItems field value if set, zero value otherwise.
func (o *BlockType) GetInventoryItems() []BlockInventoryItemType {
	if o == nil || IsNil(o.InventoryItems) {
		var ret []BlockInventoryItemType
		return ret
	}
	return o.InventoryItems
}

// GetInventoryItemsOk returns a tuple with the InventoryItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetInventoryItemsOk() ([]BlockInventoryItemType, bool) {
	if o == nil || IsNil(o.InventoryItems) {
		return nil, false
	}
	return o.InventoryItems, true
}

// HasInventoryItems returns a boolean if a field has been set.
func (o *BlockType) HasInventoryItems() bool {
	if o != nil && !IsNil(o.InventoryItems) {
		return true
	}

	return false
}

// SetInventoryItems gets a reference to the given []BlockInventoryItemType and assigns it to the InventoryItems field.
func (o *BlockType) SetInventoryItems(v []BlockInventoryItemType) {
	o.InventoryItems = v
}

// GetBlockStatistics returns the BlockStatistics field value if set, zero value otherwise.
func (o *BlockType) GetBlockStatistics() BlockStatisticsType {
	if o == nil || IsNil(o.BlockStatistics) {
		var ret BlockStatisticsType
		return ret
	}
	return *o.BlockStatistics
}

// GetBlockStatisticsOk returns a tuple with the BlockStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetBlockStatisticsOk() (*BlockStatisticsType, bool) {
	if o == nil || IsNil(o.BlockStatistics) {
		return nil, false
	}
	return o.BlockStatistics, true
}

// HasBlockStatistics returns a boolean if a field has been set.
func (o *BlockType) HasBlockStatistics() bool {
	if o != nil && !IsNil(o.BlockStatistics) {
		return true
	}

	return false
}

// SetBlockStatistics gets a reference to the given BlockStatisticsType and assigns it to the BlockStatistics field.
func (o *BlockType) SetBlockStatistics(v BlockStatisticsType) {
	o.BlockStatistics = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BlockType) GetComments() BlockTypeComments {
	if o == nil || IsNil(o.Comments) {
		var ret BlockTypeComments
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetCommentsOk() (*BlockTypeComments, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BlockType) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given BlockTypeComments and assigns it to the Comments field.
func (o *BlockType) SetComments(v BlockTypeComments) {
	o.Comments = &v
}

// GetBlockIndicators returns the BlockIndicators field value if set, zero value otherwise.
func (o *BlockType) GetBlockIndicators() []IndicatorType {
	if o == nil || IsNil(o.BlockIndicators) {
		var ret []IndicatorType
		return ret
	}
	return o.BlockIndicators
}

// GetBlockIndicatorsOk returns a tuple with the BlockIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetBlockIndicatorsOk() ([]IndicatorType, bool) {
	if o == nil || IsNil(o.BlockIndicators) {
		return nil, false
	}
	return o.BlockIndicators, true
}

// HasBlockIndicators returns a boolean if a field has been set.
func (o *BlockType) HasBlockIndicators() bool {
	if o != nil && !IsNil(o.BlockIndicators) {
		return true
	}

	return false
}

// SetBlockIndicators gets a reference to the given []IndicatorType and assigns it to the BlockIndicators field.
func (o *BlockType) SetBlockIndicators(v []IndicatorType) {
	o.BlockIndicators = v
}

// GetTraces returns the Traces field value if set, zero value otherwise.
func (o *BlockType) GetTraces() []BlockTraceType {
	if o == nil || IsNil(o.Traces) {
		var ret []BlockTraceType
		return ret
	}
	return o.Traces
}

// GetTracesOk returns a tuple with the Traces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetTracesOk() ([]BlockTraceType, bool) {
	if o == nil || IsNil(o.Traces) {
		return nil, false
	}
	return o.Traces, true
}

// HasTraces returns a boolean if a field has been set.
func (o *BlockType) HasTraces() bool {
	if o != nil && !IsNil(o.Traces) {
		return true
	}

	return false
}

// SetTraces gets a reference to the given []BlockTraceType and assigns it to the Traces field.
func (o *BlockType) SetTraces(v []BlockTraceType) {
	o.Traces = v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *BlockType) GetStatistics() []BlockStatisticsSummaryType {
	if o == nil || IsNil(o.Statistics) {
		var ret []BlockStatisticsSummaryType
		return ret
	}
	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetStatisticsOk() ([]BlockStatisticsSummaryType, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *BlockType) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given []BlockStatisticsSummaryType and assigns it to the Statistics field.
func (o *BlockType) SetStatistics(v []BlockStatisticsSummaryType) {
	o.Statistics = v
}

// GetBlockSecurity returns the BlockSecurity field value if set, zero value otherwise.
func (o *BlockType) GetBlockSecurity() BlockTypeBlockSecurity {
	if o == nil || IsNil(o.BlockSecurity) {
		var ret BlockTypeBlockSecurity
		return ret
	}
	return *o.BlockSecurity
}

// GetBlockSecurityOk returns a tuple with the BlockSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetBlockSecurityOk() (*BlockTypeBlockSecurity, bool) {
	if o == nil || IsNil(o.BlockSecurity) {
		return nil, false
	}
	return o.BlockSecurity, true
}

// HasBlockSecurity returns a boolean if a field has been set.
func (o *BlockType) HasBlockSecurity() bool {
	if o != nil && !IsNil(o.BlockSecurity) {
		return true
	}

	return false
}

// SetBlockSecurity gets a reference to the given BlockTypeBlockSecurity and assigns it to the BlockSecurity field.
func (o *BlockType) SetBlockSecurity(v BlockTypeBlockSecurity) {
	o.BlockSecurity = &v
}

// GetSellMessages returns the SellMessages field value if set, zero value otherwise.
func (o *BlockType) GetSellMessages() SellMessageConfigsType {
	if o == nil || IsNil(o.SellMessages) {
		var ret SellMessageConfigsType
		return ret
	}
	return *o.SellMessages
}

// GetSellMessagesOk returns a tuple with the SellMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetSellMessagesOk() (*SellMessageConfigsType, bool) {
	if o == nil || IsNil(o.SellMessages) {
		return nil, false
	}
	return o.SellMessages, true
}

// HasSellMessages returns a boolean if a field has been set.
func (o *BlockType) HasSellMessages() bool {
	if o != nil && !IsNil(o.SellMessages) {
		return true
	}

	return false
}

// SetSellMessages gets a reference to the given SellMessageConfigsType and assigns it to the SellMessages field.
func (o *BlockType) SetSellMessages(v SellMessageConfigsType) {
	o.SellMessages = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *BlockType) GetAttachments() []AttachmentType {
	if o == nil || IsNil(o.Attachments) {
		var ret []AttachmentType
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetAttachmentsOk() ([]AttachmentType, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *BlockType) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentType and assigns it to the Attachments field.
func (o *BlockType) SetAttachments(v []AttachmentType) {
	o.Attachments = v
}

// GetWashSchedules returns the WashSchedules field value if set, zero value otherwise.
func (o *BlockType) GetWashSchedules() []BlockWashScheduleType {
	if o == nil || IsNil(o.WashSchedules) {
		var ret []BlockWashScheduleType
		return ret
	}
	return o.WashSchedules
}

// GetWashSchedulesOk returns a tuple with the WashSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetWashSchedulesOk() ([]BlockWashScheduleType, bool) {
	if o == nil || IsNil(o.WashSchedules) {
		return nil, false
	}
	return o.WashSchedules, true
}

// HasWashSchedules returns a boolean if a field has been set.
func (o *BlockType) HasWashSchedules() bool {
	if o != nil && !IsNil(o.WashSchedules) {
		return true
	}

	return false
}

// SetWashSchedules gets a reference to the given []BlockWashScheduleType and assigns it to the WashSchedules field.
func (o *BlockType) SetWashSchedules(v []BlockWashScheduleType) {
	o.WashSchedules = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *BlockType) GetRestrictions() BlockRestrictionsType {
	if o == nil || IsNil(o.Restrictions) {
		var ret BlockRestrictionsType
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetRestrictionsOk() (*BlockRestrictionsType, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *BlockType) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given BlockRestrictionsType and assigns it to the Restrictions field.
func (o *BlockType) SetRestrictions(v BlockRestrictionsType) {
	o.Restrictions = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *BlockType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *BlockType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *BlockType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetMarkAsRecentlyAccessed returns the MarkAsRecentlyAccessed field value if set, zero value otherwise.
func (o *BlockType) GetMarkAsRecentlyAccessed() bool {
	if o == nil || IsNil(o.MarkAsRecentlyAccessed) {
		var ret bool
		return ret
	}
	return *o.MarkAsRecentlyAccessed
}

// GetMarkAsRecentlyAccessedOk returns a tuple with the MarkAsRecentlyAccessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetMarkAsRecentlyAccessedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkAsRecentlyAccessed) {
		return nil, false
	}
	return o.MarkAsRecentlyAccessed, true
}

// HasMarkAsRecentlyAccessed returns a boolean if a field has been set.
func (o *BlockType) HasMarkAsRecentlyAccessed() bool {
	if o != nil && !IsNil(o.MarkAsRecentlyAccessed) {
		return true
	}

	return false
}

// SetMarkAsRecentlyAccessed gets a reference to the given bool and assigns it to the MarkAsRecentlyAccessed field.
func (o *BlockType) SetMarkAsRecentlyAccessed(v bool) {
	o.MarkAsRecentlyAccessed = &v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *BlockType) GetCreateDateTime() time.Time {
	if o == nil || IsNil(o.CreateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetCreateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDateTime) {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *BlockType) HasCreateDateTime() bool {
	if o != nil && !IsNil(o.CreateDateTime) {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given time.Time and assigns it to the CreateDateTime field.
func (o *BlockType) SetCreateDateTime(v time.Time) {
	o.CreateDateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *BlockType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *BlockType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *BlockType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModifyDateTime returns the LastModifyDateTime field value if set, zero value otherwise.
func (o *BlockType) GetLastModifyDateTime() time.Time {
	if o == nil || IsNil(o.LastModifyDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyDateTime
}

// GetLastModifyDateTimeOk returns a tuple with the LastModifyDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetLastModifyDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyDateTime) {
		return nil, false
	}
	return o.LastModifyDateTime, true
}

// HasLastModifyDateTime returns a boolean if a field has been set.
func (o *BlockType) HasLastModifyDateTime() bool {
	if o != nil && !IsNil(o.LastModifyDateTime) {
		return true
	}

	return false
}

// SetLastModifyDateTime gets a reference to the given time.Time and assigns it to the LastModifyDateTime field.
func (o *BlockType) SetLastModifyDateTime(v time.Time) {
	o.LastModifyDateTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *BlockType) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockType) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *BlockType) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *BlockType) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

func (o BlockType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockIdList) {
		toSerialize["blockIdList"] = o.BlockIdList
	}
	if !IsNil(o.ExternalReferences) {
		toSerialize["externalReferences"] = o.ExternalReferences
	}
	if !IsNil(o.BlockDetails) {
		toSerialize["blockDetails"] = o.BlockDetails
	}
	if !IsNil(o.BlockOwners) {
		toSerialize["blockOwners"] = o.BlockOwners
	}
	if !IsNil(o.Catering) {
		toSerialize["catering"] = o.Catering
	}
	if !IsNil(o.BlockProfiles) {
		toSerialize["blockProfiles"] = o.BlockProfiles
	}
	if !IsNil(o.InventoryItems) {
		toSerialize["inventoryItems"] = o.InventoryItems
	}
	if !IsNil(o.BlockStatistics) {
		toSerialize["blockStatistics"] = o.BlockStatistics
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.BlockIndicators) {
		toSerialize["blockIndicators"] = o.BlockIndicators
	}
	if !IsNil(o.Traces) {
		toSerialize["traces"] = o.Traces
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.BlockSecurity) {
		toSerialize["blockSecurity"] = o.BlockSecurity
	}
	if !IsNil(o.SellMessages) {
		toSerialize["sellMessages"] = o.SellMessages
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.WashSchedules) {
		toSerialize["washSchedules"] = o.WashSchedules
	}
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.MarkAsRecentlyAccessed) {
		toSerialize["markAsRecentlyAccessed"] = o.MarkAsRecentlyAccessed
	}
	if !IsNil(o.CreateDateTime) {
		toSerialize["createDateTime"] = o.CreateDateTime
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !IsNil(o.LastModifyDateTime) {
		toSerialize["lastModifyDateTime"] = o.LastModifyDateTime
	}
	if !IsNil(o.LastModifierId) {
		toSerialize["lastModifierId"] = o.LastModifierId
	}
	return toSerialize, nil
}

type NullableBlockType struct {
	value *BlockType
	isSet bool
}

func (v NullableBlockType) Get() *BlockType {
	return v.value
}

func (v *NullableBlockType) Set(val *BlockType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockType(val *BlockType) *NullableBlockType {
	return &NullableBlockType{value: val, isSet: true}
}

func (v NullableBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


