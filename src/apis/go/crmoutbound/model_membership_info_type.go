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

// checks if the MembershipInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipInfoType{}

// MembershipInfoType The Membership object identifies the frequent customer reward program.
type MembershipInfoType struct {
	// Membership ID (Unique ID from the memberships table).
	MembershipId *float32 `json:"membershipId,omitempty"`
	// The code or name of the membership program ('Hertz', 'AAdvantage', etc.).
	ProgramCode *string `json:"programCode,omitempty"`
	// The code or name of the bonus program. BonusCode can be used to indicate the level of membership (Gold Club, Platinum member, etc.)
	BonusCode *string `json:"bonusCode,omitempty"`
	// The description of the ProgramCode.(Delta Previlige for code DP)
	MembershipTypeDesc *string `json:"membershipTypeDesc,omitempty"`
	// The description of the Bonus Code.(Platinum for code P)
	MembershipLevelDesc *string `json:"membershipLevelDesc,omitempty"`
	// The account identification number for this particular member in this particular program.
	AccountId *string `json:"accountId,omitempty"`
	// The code or name of the membership level and indicates the level of membership (Gold Club, Platinum member, etc.). This is same as the BonusCode.
	MembershipLevel *string `json:"membershipLevel,omitempty"`
	// Ranking assigned to the Player Profile by the Gaming system.
	PlayerRanking *int32 `json:"playerRanking,omitempty"`
}

// NewMembershipInfoType instantiates a new MembershipInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipInfoType() *MembershipInfoType {
	this := MembershipInfoType{}
	return &this
}

// NewMembershipInfoTypeWithDefaults instantiates a new MembershipInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipInfoTypeWithDefaults() *MembershipInfoType {
	this := MembershipInfoType{}
	return &this
}

// GetMembershipId returns the MembershipId field value if set, zero value otherwise.
func (o *MembershipInfoType) GetMembershipId() float32 {
	if o == nil || IsNil(o.MembershipId) {
		var ret float32
		return ret
	}
	return *o.MembershipId
}

// GetMembershipIdOk returns a tuple with the MembershipId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetMembershipIdOk() (*float32, bool) {
	if o == nil || IsNil(o.MembershipId) {
		return nil, false
	}
	return o.MembershipId, true
}

// HasMembershipId returns a boolean if a field has been set.
func (o *MembershipInfoType) HasMembershipId() bool {
	if o != nil && !IsNil(o.MembershipId) {
		return true
	}

	return false
}

// SetMembershipId gets a reference to the given float32 and assigns it to the MembershipId field.
func (o *MembershipInfoType) SetMembershipId(v float32) {
	o.MembershipId = &v
}

// GetProgramCode returns the ProgramCode field value if set, zero value otherwise.
func (o *MembershipInfoType) GetProgramCode() string {
	if o == nil || IsNil(o.ProgramCode) {
		var ret string
		return ret
	}
	return *o.ProgramCode
}

// GetProgramCodeOk returns a tuple with the ProgramCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetProgramCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramCode) {
		return nil, false
	}
	return o.ProgramCode, true
}

// HasProgramCode returns a boolean if a field has been set.
func (o *MembershipInfoType) HasProgramCode() bool {
	if o != nil && !IsNil(o.ProgramCode) {
		return true
	}

	return false
}

// SetProgramCode gets a reference to the given string and assigns it to the ProgramCode field.
func (o *MembershipInfoType) SetProgramCode(v string) {
	o.ProgramCode = &v
}

// GetBonusCode returns the BonusCode field value if set, zero value otherwise.
func (o *MembershipInfoType) GetBonusCode() string {
	if o == nil || IsNil(o.BonusCode) {
		var ret string
		return ret
	}
	return *o.BonusCode
}

// GetBonusCodeOk returns a tuple with the BonusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetBonusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BonusCode) {
		return nil, false
	}
	return o.BonusCode, true
}

// HasBonusCode returns a boolean if a field has been set.
func (o *MembershipInfoType) HasBonusCode() bool {
	if o != nil && !IsNil(o.BonusCode) {
		return true
	}

	return false
}

// SetBonusCode gets a reference to the given string and assigns it to the BonusCode field.
func (o *MembershipInfoType) SetBonusCode(v string) {
	o.BonusCode = &v
}

// GetMembershipTypeDesc returns the MembershipTypeDesc field value if set, zero value otherwise.
func (o *MembershipInfoType) GetMembershipTypeDesc() string {
	if o == nil || IsNil(o.MembershipTypeDesc) {
		var ret string
		return ret
	}
	return *o.MembershipTypeDesc
}

// GetMembershipTypeDescOk returns a tuple with the MembershipTypeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetMembershipTypeDescOk() (*string, bool) {
	if o == nil || IsNil(o.MembershipTypeDesc) {
		return nil, false
	}
	return o.MembershipTypeDesc, true
}

// HasMembershipTypeDesc returns a boolean if a field has been set.
func (o *MembershipInfoType) HasMembershipTypeDesc() bool {
	if o != nil && !IsNil(o.MembershipTypeDesc) {
		return true
	}

	return false
}

// SetMembershipTypeDesc gets a reference to the given string and assigns it to the MembershipTypeDesc field.
func (o *MembershipInfoType) SetMembershipTypeDesc(v string) {
	o.MembershipTypeDesc = &v
}

// GetMembershipLevelDesc returns the MembershipLevelDesc field value if set, zero value otherwise.
func (o *MembershipInfoType) GetMembershipLevelDesc() string {
	if o == nil || IsNil(o.MembershipLevelDesc) {
		var ret string
		return ret
	}
	return *o.MembershipLevelDesc
}

// GetMembershipLevelDescOk returns a tuple with the MembershipLevelDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetMembershipLevelDescOk() (*string, bool) {
	if o == nil || IsNil(o.MembershipLevelDesc) {
		return nil, false
	}
	return o.MembershipLevelDesc, true
}

// HasMembershipLevelDesc returns a boolean if a field has been set.
func (o *MembershipInfoType) HasMembershipLevelDesc() bool {
	if o != nil && !IsNil(o.MembershipLevelDesc) {
		return true
	}

	return false
}

// SetMembershipLevelDesc gets a reference to the given string and assigns it to the MembershipLevelDesc field.
func (o *MembershipInfoType) SetMembershipLevelDesc(v string) {
	o.MembershipLevelDesc = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *MembershipInfoType) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *MembershipInfoType) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *MembershipInfoType) SetAccountId(v string) {
	o.AccountId = &v
}

// GetMembershipLevel returns the MembershipLevel field value if set, zero value otherwise.
func (o *MembershipInfoType) GetMembershipLevel() string {
	if o == nil || IsNil(o.MembershipLevel) {
		var ret string
		return ret
	}
	return *o.MembershipLevel
}

// GetMembershipLevelOk returns a tuple with the MembershipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetMembershipLevelOk() (*string, bool) {
	if o == nil || IsNil(o.MembershipLevel) {
		return nil, false
	}
	return o.MembershipLevel, true
}

// HasMembershipLevel returns a boolean if a field has been set.
func (o *MembershipInfoType) HasMembershipLevel() bool {
	if o != nil && !IsNil(o.MembershipLevel) {
		return true
	}

	return false
}

// SetMembershipLevel gets a reference to the given string and assigns it to the MembershipLevel field.
func (o *MembershipInfoType) SetMembershipLevel(v string) {
	o.MembershipLevel = &v
}

// GetPlayerRanking returns the PlayerRanking field value if set, zero value otherwise.
func (o *MembershipInfoType) GetPlayerRanking() int32 {
	if o == nil || IsNil(o.PlayerRanking) {
		var ret int32
		return ret
	}
	return *o.PlayerRanking
}

// GetPlayerRankingOk returns a tuple with the PlayerRanking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipInfoType) GetPlayerRankingOk() (*int32, bool) {
	if o == nil || IsNil(o.PlayerRanking) {
		return nil, false
	}
	return o.PlayerRanking, true
}

// HasPlayerRanking returns a boolean if a field has been set.
func (o *MembershipInfoType) HasPlayerRanking() bool {
	if o != nil && !IsNil(o.PlayerRanking) {
		return true
	}

	return false
}

// SetPlayerRanking gets a reference to the given int32 and assigns it to the PlayerRanking field.
func (o *MembershipInfoType) SetPlayerRanking(v int32) {
	o.PlayerRanking = &v
}

func (o MembershipInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MembershipId) {
		toSerialize["membershipId"] = o.MembershipId
	}
	if !IsNil(o.ProgramCode) {
		toSerialize["programCode"] = o.ProgramCode
	}
	if !IsNil(o.BonusCode) {
		toSerialize["bonusCode"] = o.BonusCode
	}
	if !IsNil(o.MembershipTypeDesc) {
		toSerialize["membershipTypeDesc"] = o.MembershipTypeDesc
	}
	if !IsNil(o.MembershipLevelDesc) {
		toSerialize["membershipLevelDesc"] = o.MembershipLevelDesc
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.MembershipLevel) {
		toSerialize["membershipLevel"] = o.MembershipLevel
	}
	if !IsNil(o.PlayerRanking) {
		toSerialize["playerRanking"] = o.PlayerRanking
	}
	return toSerialize, nil
}

type NullableMembershipInfoType struct {
	value *MembershipInfoType
	isSet bool
}

func (v NullableMembershipInfoType) Get() *MembershipInfoType {
	return v.value
}

func (v *NullableMembershipInfoType) Set(val *MembershipInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipInfoType(val *MembershipInfoType) *NullableMembershipInfoType {
	return &NullableMembershipInfoType{value: val, isSet: true}
}

func (v NullableMembershipInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


