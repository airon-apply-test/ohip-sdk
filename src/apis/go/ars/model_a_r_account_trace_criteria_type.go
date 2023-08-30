/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
	"time"
)

// checks if the ARAccountTraceCriteriaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARAccountTraceCriteriaType{}

// ARAccountTraceCriteriaType The traces on the AR Account.
type ARAccountTraceCriteriaType struct {
	TimeInfo *TraceTimeInfoType `json:"timeInfo,omitempty"`
	ReservationId *ReservationId `json:"reservationId,omitempty"`
	// Indicates the Department code.
	DepartmentId *string `json:"departmentId,omitempty"`
	// The information this trace contains.
	TraceText *string `json:"traceText,omitempty"`
	ResolveInfo *TraceResolveType `json:"resolveInfo,omitempty"`
	// URL that identifies the location associated with the record identified by the UniqueID.
	Url *string `json:"url,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
	// The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
	Instance *string `json:"instance,omitempty"`
	// Used to identify the source of the identifier (e.g., IATA, ABTA).
	IdContext *string `json:"idContext,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// Additional identifying value assigned by the creating system.
	IdExtension *int32 `json:"idExtension,omitempty"`
	// Time stamp of the creation.
	CreateDateTime *time.Time `json:"createDateTime,omitempty"`
	// ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
	CreatorId *string `json:"creatorId,omitempty"`
	// Time stamp of last modification.
	LastModifyDateTime *time.Time `json:"lastModifyDateTime,omitempty"`
	// Identifies the last software system or person to modify a record.
	LastModifierId *string `json:"lastModifierId,omitempty"`
	// Date an item will be purged from a database (e.g., from a live database to an archive).
	PurgeDate *string `json:"purgeDate,omitempty"`
	// The resort where the AR Account exists.
	HotelId *string `json:"hotelId,omitempty"`
	AccountId *UniqueIDType `json:"accountId,omitempty"`
}

// NewARAccountTraceCriteriaType instantiates a new ARAccountTraceCriteriaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARAccountTraceCriteriaType() *ARAccountTraceCriteriaType {
	this := ARAccountTraceCriteriaType{}
	return &this
}

// NewARAccountTraceCriteriaTypeWithDefaults instantiates a new ARAccountTraceCriteriaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARAccountTraceCriteriaTypeWithDefaults() *ARAccountTraceCriteriaType {
	this := ARAccountTraceCriteriaType{}
	return &this
}

// GetTimeInfo returns the TimeInfo field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetTimeInfo() TraceTimeInfoType {
	if o == nil || IsNil(o.TimeInfo) {
		var ret TraceTimeInfoType
		return ret
	}
	return *o.TimeInfo
}

// GetTimeInfoOk returns a tuple with the TimeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetTimeInfoOk() (*TraceTimeInfoType, bool) {
	if o == nil || IsNil(o.TimeInfo) {
		return nil, false
	}
	return o.TimeInfo, true
}

// HasTimeInfo returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasTimeInfo() bool {
	if o != nil && !IsNil(o.TimeInfo) {
		return true
	}

	return false
}

// SetTimeInfo gets a reference to the given TraceTimeInfoType and assigns it to the TimeInfo field.
func (o *ARAccountTraceCriteriaType) SetTimeInfo(v TraceTimeInfoType) {
	o.TimeInfo = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetReservationId() ReservationId {
	if o == nil || IsNil(o.ReservationId) {
		var ret ReservationId
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetReservationIdOk() (*ReservationId, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given ReservationId and assigns it to the ReservationId field.
func (o *ARAccountTraceCriteriaType) SetReservationId(v ReservationId) {
	o.ReservationId = &v
}

// GetDepartmentId returns the DepartmentId field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetDepartmentId() string {
	if o == nil || IsNil(o.DepartmentId) {
		var ret string
		return ret
	}
	return *o.DepartmentId
}

// GetDepartmentIdOk returns a tuple with the DepartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetDepartmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DepartmentId) {
		return nil, false
	}
	return o.DepartmentId, true
}

// HasDepartmentId returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasDepartmentId() bool {
	if o != nil && !IsNil(o.DepartmentId) {
		return true
	}

	return false
}

// SetDepartmentId gets a reference to the given string and assigns it to the DepartmentId field.
func (o *ARAccountTraceCriteriaType) SetDepartmentId(v string) {
	o.DepartmentId = &v
}

// GetTraceText returns the TraceText field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetTraceText() string {
	if o == nil || IsNil(o.TraceText) {
		var ret string
		return ret
	}
	return *o.TraceText
}

// GetTraceTextOk returns a tuple with the TraceText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetTraceTextOk() (*string, bool) {
	if o == nil || IsNil(o.TraceText) {
		return nil, false
	}
	return o.TraceText, true
}

// HasTraceText returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasTraceText() bool {
	if o != nil && !IsNil(o.TraceText) {
		return true
	}

	return false
}

// SetTraceText gets a reference to the given string and assigns it to the TraceText field.
func (o *ARAccountTraceCriteriaType) SetTraceText(v string) {
	o.TraceText = &v
}

// GetResolveInfo returns the ResolveInfo field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetResolveInfo() TraceResolveType {
	if o == nil || IsNil(o.ResolveInfo) {
		var ret TraceResolveType
		return ret
	}
	return *o.ResolveInfo
}

// GetResolveInfoOk returns a tuple with the ResolveInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetResolveInfoOk() (*TraceResolveType, bool) {
	if o == nil || IsNil(o.ResolveInfo) {
		return nil, false
	}
	return o.ResolveInfo, true
}

// HasResolveInfo returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasResolveInfo() bool {
	if o != nil && !IsNil(o.ResolveInfo) {
		return true
	}

	return false
}

// SetResolveInfo gets a reference to the given TraceResolveType and assigns it to the ResolveInfo field.
func (o *ARAccountTraceCriteriaType) SetResolveInfo(v TraceResolveType) {
	o.ResolveInfo = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ARAccountTraceCriteriaType) SetUrl(v string) {
	o.Url = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ARAccountTraceCriteriaType) SetType(v string) {
	o.Type = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ARAccountTraceCriteriaType) SetInstance(v string) {
	o.Instance = &v
}

// GetIdContext returns the IdContext field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetIdContext() string {
	if o == nil || IsNil(o.IdContext) {
		var ret string
		return ret
	}
	return *o.IdContext
}

// GetIdContextOk returns a tuple with the IdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetIdContextOk() (*string, bool) {
	if o == nil || IsNil(o.IdContext) {
		return nil, false
	}
	return o.IdContext, true
}

// HasIdContext returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasIdContext() bool {
	if o != nil && !IsNil(o.IdContext) {
		return true
	}

	return false
}

// SetIdContext gets a reference to the given string and assigns it to the IdContext field.
func (o *ARAccountTraceCriteriaType) SetIdContext(v string) {
	o.IdContext = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ARAccountTraceCriteriaType) SetId(v string) {
	o.Id = &v
}

// GetIdExtension returns the IdExtension field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetIdExtension() int32 {
	if o == nil || IsNil(o.IdExtension) {
		var ret int32
		return ret
	}
	return *o.IdExtension
}

// GetIdExtensionOk returns a tuple with the IdExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetIdExtensionOk() (*int32, bool) {
	if o == nil || IsNil(o.IdExtension) {
		return nil, false
	}
	return o.IdExtension, true
}

// HasIdExtension returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasIdExtension() bool {
	if o != nil && !IsNil(o.IdExtension) {
		return true
	}

	return false
}

// SetIdExtension gets a reference to the given int32 and assigns it to the IdExtension field.
func (o *ARAccountTraceCriteriaType) SetIdExtension(v int32) {
	o.IdExtension = &v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetCreateDateTime() time.Time {
	if o == nil || IsNil(o.CreateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetCreateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDateTime) {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasCreateDateTime() bool {
	if o != nil && !IsNil(o.CreateDateTime) {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given time.Time and assigns it to the CreateDateTime field.
func (o *ARAccountTraceCriteriaType) SetCreateDateTime(v time.Time) {
	o.CreateDateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *ARAccountTraceCriteriaType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModifyDateTime returns the LastModifyDateTime field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetLastModifyDateTime() time.Time {
	if o == nil || IsNil(o.LastModifyDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyDateTime
}

// GetLastModifyDateTimeOk returns a tuple with the LastModifyDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetLastModifyDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyDateTime) {
		return nil, false
	}
	return o.LastModifyDateTime, true
}

// HasLastModifyDateTime returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasLastModifyDateTime() bool {
	if o != nil && !IsNil(o.LastModifyDateTime) {
		return true
	}

	return false
}

// SetLastModifyDateTime gets a reference to the given time.Time and assigns it to the LastModifyDateTime field.
func (o *ARAccountTraceCriteriaType) SetLastModifyDateTime(v time.Time) {
	o.LastModifyDateTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *ARAccountTraceCriteriaType) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetPurgeDate returns the PurgeDate field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetPurgeDate() string {
	if o == nil || IsNil(o.PurgeDate) {
		var ret string
		return ret
	}
	return *o.PurgeDate
}

// GetPurgeDateOk returns a tuple with the PurgeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetPurgeDateOk() (*string, bool) {
	if o == nil || IsNil(o.PurgeDate) {
		return nil, false
	}
	return o.PurgeDate, true
}

// HasPurgeDate returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasPurgeDate() bool {
	if o != nil && !IsNil(o.PurgeDate) {
		return true
	}

	return false
}

// SetPurgeDate gets a reference to the given string and assigns it to the PurgeDate field.
func (o *ARAccountTraceCriteriaType) SetPurgeDate(v string) {
	o.PurgeDate = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ARAccountTraceCriteriaType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ARAccountTraceCriteriaType) GetAccountId() UniqueIDType {
	if o == nil || IsNil(o.AccountId) {
		var ret UniqueIDType
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTraceCriteriaType) GetAccountIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ARAccountTraceCriteriaType) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given UniqueIDType and assigns it to the AccountId field.
func (o *ARAccountTraceCriteriaType) SetAccountId(v UniqueIDType) {
	o.AccountId = &v
}

func (o ARAccountTraceCriteriaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARAccountTraceCriteriaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeInfo) {
		toSerialize["timeInfo"] = o.TimeInfo
	}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.DepartmentId) {
		toSerialize["departmentId"] = o.DepartmentId
	}
	if !IsNil(o.TraceText) {
		toSerialize["traceText"] = o.TraceText
	}
	if !IsNil(o.ResolveInfo) {
		toSerialize["resolveInfo"] = o.ResolveInfo
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.IdContext) {
		toSerialize["idContext"] = o.IdContext
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IdExtension) {
		toSerialize["idExtension"] = o.IdExtension
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
	if !IsNil(o.PurgeDate) {
		toSerialize["purgeDate"] = o.PurgeDate
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	return toSerialize, nil
}

type NullableARAccountTraceCriteriaType struct {
	value *ARAccountTraceCriteriaType
	isSet bool
}

func (v NullableARAccountTraceCriteriaType) Get() *ARAccountTraceCriteriaType {
	return v.value
}

func (v *NullableARAccountTraceCriteriaType) Set(val *ARAccountTraceCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableARAccountTraceCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableARAccountTraceCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARAccountTraceCriteriaType(val *ARAccountTraceCriteriaType) *NullableARAccountTraceCriteriaType {
	return &NullableARAccountTraceCriteriaType{value: val, isSet: true}
}

func (v NullableARAccountTraceCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARAccountTraceCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


