/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SalesInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesInfoType{}

// SalesInfoType Provides sales information about the profiles of type company, travel agent, source and contact.
type SalesInfoType struct {
	// Defines the scope.
	Scope *string `json:"scope,omitempty"`
	// Defines the scope city.
	ScopeCity *string `json:"scopeCity,omitempty"`
	// Defines the account type.
	AccountType *string `json:"accountType,omitempty"`
	// Defines the account source.
	AccountSource *string `json:"accountSource,omitempty"`
	// Defines the industry code.
	IndustryCode *string `json:"industryCode,omitempty"`
	// Defines the Business segments.
	BusinessSegments *string `json:"businessSegments,omitempty"`
	// Defines the priority.
	Priority *string `json:"priority,omitempty"`
	// Defines the rooms potential.
	RoomsPotential *string `json:"roomsPotential,omitempty"`
	// Defines the action code.
	ActionCode *string `json:"actionCode,omitempty"`
	// Defines the competition code.
	CompetitionCode *string `json:"competitionCode,omitempty"`
	// Defines the influence for the contact profile.
	Influence *string `json:"influence,omitempty"`
	// Defines the Preferred Room for profile.
	PreferredRoom *string `json:"preferredRoom,omitempty"`
	// Hotel Code used to filter the sales information.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewSalesInfoType instantiates a new SalesInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesInfoType() *SalesInfoType {
	this := SalesInfoType{}
	return &this
}

// NewSalesInfoTypeWithDefaults instantiates a new SalesInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesInfoTypeWithDefaults() *SalesInfoType {
	this := SalesInfoType{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SalesInfoType) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SalesInfoType) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *SalesInfoType) SetScope(v string) {
	o.Scope = &v
}

// GetScopeCity returns the ScopeCity field value if set, zero value otherwise.
func (o *SalesInfoType) GetScopeCity() string {
	if o == nil || IsNil(o.ScopeCity) {
		var ret string
		return ret
	}
	return *o.ScopeCity
}

// GetScopeCityOk returns a tuple with the ScopeCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetScopeCityOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeCity) {
		return nil, false
	}
	return o.ScopeCity, true
}

// HasScopeCity returns a boolean if a field has been set.
func (o *SalesInfoType) HasScopeCity() bool {
	if o != nil && !IsNil(o.ScopeCity) {
		return true
	}

	return false
}

// SetScopeCity gets a reference to the given string and assigns it to the ScopeCity field.
func (o *SalesInfoType) SetScopeCity(v string) {
	o.ScopeCity = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *SalesInfoType) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *SalesInfoType) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *SalesInfoType) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAccountSource returns the AccountSource field value if set, zero value otherwise.
func (o *SalesInfoType) GetAccountSource() string {
	if o == nil || IsNil(o.AccountSource) {
		var ret string
		return ret
	}
	return *o.AccountSource
}

// GetAccountSourceOk returns a tuple with the AccountSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetAccountSourceOk() (*string, bool) {
	if o == nil || IsNil(o.AccountSource) {
		return nil, false
	}
	return o.AccountSource, true
}

// HasAccountSource returns a boolean if a field has been set.
func (o *SalesInfoType) HasAccountSource() bool {
	if o != nil && !IsNil(o.AccountSource) {
		return true
	}

	return false
}

// SetAccountSource gets a reference to the given string and assigns it to the AccountSource field.
func (o *SalesInfoType) SetAccountSource(v string) {
	o.AccountSource = &v
}

// GetIndustryCode returns the IndustryCode field value if set, zero value otherwise.
func (o *SalesInfoType) GetIndustryCode() string {
	if o == nil || IsNil(o.IndustryCode) {
		var ret string
		return ret
	}
	return *o.IndustryCode
}

// GetIndustryCodeOk returns a tuple with the IndustryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetIndustryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.IndustryCode) {
		return nil, false
	}
	return o.IndustryCode, true
}

// HasIndustryCode returns a boolean if a field has been set.
func (o *SalesInfoType) HasIndustryCode() bool {
	if o != nil && !IsNil(o.IndustryCode) {
		return true
	}

	return false
}

// SetIndustryCode gets a reference to the given string and assigns it to the IndustryCode field.
func (o *SalesInfoType) SetIndustryCode(v string) {
	o.IndustryCode = &v
}

// GetBusinessSegments returns the BusinessSegments field value if set, zero value otherwise.
func (o *SalesInfoType) GetBusinessSegments() string {
	if o == nil || IsNil(o.BusinessSegments) {
		var ret string
		return ret
	}
	return *o.BusinessSegments
}

// GetBusinessSegmentsOk returns a tuple with the BusinessSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetBusinessSegmentsOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessSegments) {
		return nil, false
	}
	return o.BusinessSegments, true
}

// HasBusinessSegments returns a boolean if a field has been set.
func (o *SalesInfoType) HasBusinessSegments() bool {
	if o != nil && !IsNil(o.BusinessSegments) {
		return true
	}

	return false
}

// SetBusinessSegments gets a reference to the given string and assigns it to the BusinessSegments field.
func (o *SalesInfoType) SetBusinessSegments(v string) {
	o.BusinessSegments = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SalesInfoType) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SalesInfoType) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *SalesInfoType) SetPriority(v string) {
	o.Priority = &v
}

// GetRoomsPotential returns the RoomsPotential field value if set, zero value otherwise.
func (o *SalesInfoType) GetRoomsPotential() string {
	if o == nil || IsNil(o.RoomsPotential) {
		var ret string
		return ret
	}
	return *o.RoomsPotential
}

// GetRoomsPotentialOk returns a tuple with the RoomsPotential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetRoomsPotentialOk() (*string, bool) {
	if o == nil || IsNil(o.RoomsPotential) {
		return nil, false
	}
	return o.RoomsPotential, true
}

// HasRoomsPotential returns a boolean if a field has been set.
func (o *SalesInfoType) HasRoomsPotential() bool {
	if o != nil && !IsNil(o.RoomsPotential) {
		return true
	}

	return false
}

// SetRoomsPotential gets a reference to the given string and assigns it to the RoomsPotential field.
func (o *SalesInfoType) SetRoomsPotential(v string) {
	o.RoomsPotential = &v
}

// GetActionCode returns the ActionCode field value if set, zero value otherwise.
func (o *SalesInfoType) GetActionCode() string {
	if o == nil || IsNil(o.ActionCode) {
		var ret string
		return ret
	}
	return *o.ActionCode
}

// GetActionCodeOk returns a tuple with the ActionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetActionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionCode) {
		return nil, false
	}
	return o.ActionCode, true
}

// HasActionCode returns a boolean if a field has been set.
func (o *SalesInfoType) HasActionCode() bool {
	if o != nil && !IsNil(o.ActionCode) {
		return true
	}

	return false
}

// SetActionCode gets a reference to the given string and assigns it to the ActionCode field.
func (o *SalesInfoType) SetActionCode(v string) {
	o.ActionCode = &v
}

// GetCompetitionCode returns the CompetitionCode field value if set, zero value otherwise.
func (o *SalesInfoType) GetCompetitionCode() string {
	if o == nil || IsNil(o.CompetitionCode) {
		var ret string
		return ret
	}
	return *o.CompetitionCode
}

// GetCompetitionCodeOk returns a tuple with the CompetitionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetCompetitionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CompetitionCode) {
		return nil, false
	}
	return o.CompetitionCode, true
}

// HasCompetitionCode returns a boolean if a field has been set.
func (o *SalesInfoType) HasCompetitionCode() bool {
	if o != nil && !IsNil(o.CompetitionCode) {
		return true
	}

	return false
}

// SetCompetitionCode gets a reference to the given string and assigns it to the CompetitionCode field.
func (o *SalesInfoType) SetCompetitionCode(v string) {
	o.CompetitionCode = &v
}

// GetInfluence returns the Influence field value if set, zero value otherwise.
func (o *SalesInfoType) GetInfluence() string {
	if o == nil || IsNil(o.Influence) {
		var ret string
		return ret
	}
	return *o.Influence
}

// GetInfluenceOk returns a tuple with the Influence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetInfluenceOk() (*string, bool) {
	if o == nil || IsNil(o.Influence) {
		return nil, false
	}
	return o.Influence, true
}

// HasInfluence returns a boolean if a field has been set.
func (o *SalesInfoType) HasInfluence() bool {
	if o != nil && !IsNil(o.Influence) {
		return true
	}

	return false
}

// SetInfluence gets a reference to the given string and assigns it to the Influence field.
func (o *SalesInfoType) SetInfluence(v string) {
	o.Influence = &v
}

// GetPreferredRoom returns the PreferredRoom field value if set, zero value otherwise.
func (o *SalesInfoType) GetPreferredRoom() string {
	if o == nil || IsNil(o.PreferredRoom) {
		var ret string
		return ret
	}
	return *o.PreferredRoom
}

// GetPreferredRoomOk returns a tuple with the PreferredRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetPreferredRoomOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredRoom) {
		return nil, false
	}
	return o.PreferredRoom, true
}

// HasPreferredRoom returns a boolean if a field has been set.
func (o *SalesInfoType) HasPreferredRoom() bool {
	if o != nil && !IsNil(o.PreferredRoom) {
		return true
	}

	return false
}

// SetPreferredRoom gets a reference to the given string and assigns it to the PreferredRoom field.
func (o *SalesInfoType) SetPreferredRoom(v string) {
	o.PreferredRoom = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *SalesInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *SalesInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *SalesInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o SalesInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.ScopeCity) {
		toSerialize["scopeCity"] = o.ScopeCity
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.AccountSource) {
		toSerialize["accountSource"] = o.AccountSource
	}
	if !IsNil(o.IndustryCode) {
		toSerialize["industryCode"] = o.IndustryCode
	}
	if !IsNil(o.BusinessSegments) {
		toSerialize["businessSegments"] = o.BusinessSegments
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.RoomsPotential) {
		toSerialize["roomsPotential"] = o.RoomsPotential
	}
	if !IsNil(o.ActionCode) {
		toSerialize["actionCode"] = o.ActionCode
	}
	if !IsNil(o.CompetitionCode) {
		toSerialize["competitionCode"] = o.CompetitionCode
	}
	if !IsNil(o.Influence) {
		toSerialize["influence"] = o.Influence
	}
	if !IsNil(o.PreferredRoom) {
		toSerialize["preferredRoom"] = o.PreferredRoom
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableSalesInfoType struct {
	value *SalesInfoType
	isSet bool
}

func (v NullableSalesInfoType) Get() *SalesInfoType {
	return v.value
}

func (v *NullableSalesInfoType) Set(val *SalesInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesInfoType(val *SalesInfoType) *NullableSalesInfoType {
	return &NullableSalesInfoType{value: val, isSet: true}
}

func (v NullableSalesInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


