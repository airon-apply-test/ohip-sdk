/*
OPERA Cloud CRM Configuration API

APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmcfg

import (
	"encoding/json"
)

// checks if the MembershipHandlingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipHandlingType{}

// MembershipHandlingType Represents the way a membership type has to be processed.
type MembershipHandlingType struct {
	// Indicates database for external system (if applicable).
	ExternalDatabase *string `json:"externalDatabase,omitempty"`
	// Indicates whether to include membership level for validation or not , when profile membership information is send from External System to ORS/OCIS.
	LevelRequired *bool `json:"levelRequired,omitempty"`
	// True if you wish to be able to select this membership type when creating an export file for fulfillment.
	Fulfillment *bool `json:"fulfillment,omitempty"`
	// Indicates whether to perform profile merge or not.
	AllowCardNumberOverride *bool `json:"allowCardNumberOverride,omitempty"`
	// Indicates whether Enrollment Code required of not.
	EnrollmentCodeRequired *bool `json:"enrollmentCodeRequired,omitempty"`
	// Indicates whether to store discard membership number or not.Membership Number will be discarded during profile merge.
	SaveCardNumberHistory *bool `json:"saveCardNumberHistory,omitempty"`
	// Represents Membership Statement UDF Set. Selected Statement UDF template will be attached to Membership.
	StatementUDFSet *string `json:"statementUDFSet,omitempty"`
	// Represents Membership status.This status will be assigned to Guest's profile with membership type.
	DefaultMembershipStatus *string `json:"defaultMembershipStatus,omitempty"`
	// Prevent profile name change.
	NameProtected *bool `json:"nameProtected,omitempty"`
	// Prevent profile alternate name change.
	AlternateNameProtected *bool `json:"alternateNameProtected,omitempty"`
	// Automatically generate the reference number for the primary membership.
	AutoGenerateReferenceNo *bool `json:"autoGenerateReferenceNo,omitempty"`
	// Automatically populate number from name.
	AutoPopulateNumberFromName *bool `json:"autoPopulateNumberFromName,omitempty"`
	// Membership card name using the alternate name.
	NameOnCardFromAltName *bool `json:"nameOnCardFromAltName,omitempty"`
	// Move the membership promotions to the active membership.
	MoveMemPromotions *bool `json:"moveMemPromotions,omitempty"`
}

// NewMembershipHandlingType instantiates a new MembershipHandlingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipHandlingType() *MembershipHandlingType {
	this := MembershipHandlingType{}
	return &this
}

// NewMembershipHandlingTypeWithDefaults instantiates a new MembershipHandlingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipHandlingTypeWithDefaults() *MembershipHandlingType {
	this := MembershipHandlingType{}
	return &this
}

// GetExternalDatabase returns the ExternalDatabase field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetExternalDatabase() string {
	if o == nil || IsNil(o.ExternalDatabase) {
		var ret string
		return ret
	}
	return *o.ExternalDatabase
}

// GetExternalDatabaseOk returns a tuple with the ExternalDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetExternalDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalDatabase) {
		return nil, false
	}
	return o.ExternalDatabase, true
}

// HasExternalDatabase returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasExternalDatabase() bool {
	if o != nil && !IsNil(o.ExternalDatabase) {
		return true
	}

	return false
}

// SetExternalDatabase gets a reference to the given string and assigns it to the ExternalDatabase field.
func (o *MembershipHandlingType) SetExternalDatabase(v string) {
	o.ExternalDatabase = &v
}

// GetLevelRequired returns the LevelRequired field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetLevelRequired() bool {
	if o == nil || IsNil(o.LevelRequired) {
		var ret bool
		return ret
	}
	return *o.LevelRequired
}

// GetLevelRequiredOk returns a tuple with the LevelRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetLevelRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.LevelRequired) {
		return nil, false
	}
	return o.LevelRequired, true
}

// HasLevelRequired returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasLevelRequired() bool {
	if o != nil && !IsNil(o.LevelRequired) {
		return true
	}

	return false
}

// SetLevelRequired gets a reference to the given bool and assigns it to the LevelRequired field.
func (o *MembershipHandlingType) SetLevelRequired(v bool) {
	o.LevelRequired = &v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetFulfillment() bool {
	if o == nil || IsNil(o.Fulfillment) {
		var ret bool
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetFulfillmentOk() (*bool, bool) {
	if o == nil || IsNil(o.Fulfillment) {
		return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasFulfillment() bool {
	if o != nil && !IsNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given bool and assigns it to the Fulfillment field.
func (o *MembershipHandlingType) SetFulfillment(v bool) {
	o.Fulfillment = &v
}

// GetAllowCardNumberOverride returns the AllowCardNumberOverride field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetAllowCardNumberOverride() bool {
	if o == nil || IsNil(o.AllowCardNumberOverride) {
		var ret bool
		return ret
	}
	return *o.AllowCardNumberOverride
}

// GetAllowCardNumberOverrideOk returns a tuple with the AllowCardNumberOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetAllowCardNumberOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCardNumberOverride) {
		return nil, false
	}
	return o.AllowCardNumberOverride, true
}

// HasAllowCardNumberOverride returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasAllowCardNumberOverride() bool {
	if o != nil && !IsNil(o.AllowCardNumberOverride) {
		return true
	}

	return false
}

// SetAllowCardNumberOverride gets a reference to the given bool and assigns it to the AllowCardNumberOverride field.
func (o *MembershipHandlingType) SetAllowCardNumberOverride(v bool) {
	o.AllowCardNumberOverride = &v
}

// GetEnrollmentCodeRequired returns the EnrollmentCodeRequired field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetEnrollmentCodeRequired() bool {
	if o == nil || IsNil(o.EnrollmentCodeRequired) {
		var ret bool
		return ret
	}
	return *o.EnrollmentCodeRequired
}

// GetEnrollmentCodeRequiredOk returns a tuple with the EnrollmentCodeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetEnrollmentCodeRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.EnrollmentCodeRequired) {
		return nil, false
	}
	return o.EnrollmentCodeRequired, true
}

// HasEnrollmentCodeRequired returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasEnrollmentCodeRequired() bool {
	if o != nil && !IsNil(o.EnrollmentCodeRequired) {
		return true
	}

	return false
}

// SetEnrollmentCodeRequired gets a reference to the given bool and assigns it to the EnrollmentCodeRequired field.
func (o *MembershipHandlingType) SetEnrollmentCodeRequired(v bool) {
	o.EnrollmentCodeRequired = &v
}

// GetSaveCardNumberHistory returns the SaveCardNumberHistory field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetSaveCardNumberHistory() bool {
	if o == nil || IsNil(o.SaveCardNumberHistory) {
		var ret bool
		return ret
	}
	return *o.SaveCardNumberHistory
}

// GetSaveCardNumberHistoryOk returns a tuple with the SaveCardNumberHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetSaveCardNumberHistoryOk() (*bool, bool) {
	if o == nil || IsNil(o.SaveCardNumberHistory) {
		return nil, false
	}
	return o.SaveCardNumberHistory, true
}

// HasSaveCardNumberHistory returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasSaveCardNumberHistory() bool {
	if o != nil && !IsNil(o.SaveCardNumberHistory) {
		return true
	}

	return false
}

// SetSaveCardNumberHistory gets a reference to the given bool and assigns it to the SaveCardNumberHistory field.
func (o *MembershipHandlingType) SetSaveCardNumberHistory(v bool) {
	o.SaveCardNumberHistory = &v
}

// GetStatementUDFSet returns the StatementUDFSet field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetStatementUDFSet() string {
	if o == nil || IsNil(o.StatementUDFSet) {
		var ret string
		return ret
	}
	return *o.StatementUDFSet
}

// GetStatementUDFSetOk returns a tuple with the StatementUDFSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetStatementUDFSetOk() (*string, bool) {
	if o == nil || IsNil(o.StatementUDFSet) {
		return nil, false
	}
	return o.StatementUDFSet, true
}

// HasStatementUDFSet returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasStatementUDFSet() bool {
	if o != nil && !IsNil(o.StatementUDFSet) {
		return true
	}

	return false
}

// SetStatementUDFSet gets a reference to the given string and assigns it to the StatementUDFSet field.
func (o *MembershipHandlingType) SetStatementUDFSet(v string) {
	o.StatementUDFSet = &v
}

// GetDefaultMembershipStatus returns the DefaultMembershipStatus field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetDefaultMembershipStatus() string {
	if o == nil || IsNil(o.DefaultMembershipStatus) {
		var ret string
		return ret
	}
	return *o.DefaultMembershipStatus
}

// GetDefaultMembershipStatusOk returns a tuple with the DefaultMembershipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetDefaultMembershipStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultMembershipStatus) {
		return nil, false
	}
	return o.DefaultMembershipStatus, true
}

// HasDefaultMembershipStatus returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasDefaultMembershipStatus() bool {
	if o != nil && !IsNil(o.DefaultMembershipStatus) {
		return true
	}

	return false
}

// SetDefaultMembershipStatus gets a reference to the given string and assigns it to the DefaultMembershipStatus field.
func (o *MembershipHandlingType) SetDefaultMembershipStatus(v string) {
	o.DefaultMembershipStatus = &v
}

// GetNameProtected returns the NameProtected field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetNameProtected() bool {
	if o == nil || IsNil(o.NameProtected) {
		var ret bool
		return ret
	}
	return *o.NameProtected
}

// GetNameProtectedOk returns a tuple with the NameProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetNameProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.NameProtected) {
		return nil, false
	}
	return o.NameProtected, true
}

// HasNameProtected returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasNameProtected() bool {
	if o != nil && !IsNil(o.NameProtected) {
		return true
	}

	return false
}

// SetNameProtected gets a reference to the given bool and assigns it to the NameProtected field.
func (o *MembershipHandlingType) SetNameProtected(v bool) {
	o.NameProtected = &v
}

// GetAlternateNameProtected returns the AlternateNameProtected field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetAlternateNameProtected() bool {
	if o == nil || IsNil(o.AlternateNameProtected) {
		var ret bool
		return ret
	}
	return *o.AlternateNameProtected
}

// GetAlternateNameProtectedOk returns a tuple with the AlternateNameProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetAlternateNameProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.AlternateNameProtected) {
		return nil, false
	}
	return o.AlternateNameProtected, true
}

// HasAlternateNameProtected returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasAlternateNameProtected() bool {
	if o != nil && !IsNil(o.AlternateNameProtected) {
		return true
	}

	return false
}

// SetAlternateNameProtected gets a reference to the given bool and assigns it to the AlternateNameProtected field.
func (o *MembershipHandlingType) SetAlternateNameProtected(v bool) {
	o.AlternateNameProtected = &v
}

// GetAutoGenerateReferenceNo returns the AutoGenerateReferenceNo field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetAutoGenerateReferenceNo() bool {
	if o == nil || IsNil(o.AutoGenerateReferenceNo) {
		var ret bool
		return ret
	}
	return *o.AutoGenerateReferenceNo
}

// GetAutoGenerateReferenceNoOk returns a tuple with the AutoGenerateReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetAutoGenerateReferenceNoOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoGenerateReferenceNo) {
		return nil, false
	}
	return o.AutoGenerateReferenceNo, true
}

// HasAutoGenerateReferenceNo returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasAutoGenerateReferenceNo() bool {
	if o != nil && !IsNil(o.AutoGenerateReferenceNo) {
		return true
	}

	return false
}

// SetAutoGenerateReferenceNo gets a reference to the given bool and assigns it to the AutoGenerateReferenceNo field.
func (o *MembershipHandlingType) SetAutoGenerateReferenceNo(v bool) {
	o.AutoGenerateReferenceNo = &v
}

// GetAutoPopulateNumberFromName returns the AutoPopulateNumberFromName field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetAutoPopulateNumberFromName() bool {
	if o == nil || IsNil(o.AutoPopulateNumberFromName) {
		var ret bool
		return ret
	}
	return *o.AutoPopulateNumberFromName
}

// GetAutoPopulateNumberFromNameOk returns a tuple with the AutoPopulateNumberFromName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetAutoPopulateNumberFromNameOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPopulateNumberFromName) {
		return nil, false
	}
	return o.AutoPopulateNumberFromName, true
}

// HasAutoPopulateNumberFromName returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasAutoPopulateNumberFromName() bool {
	if o != nil && !IsNil(o.AutoPopulateNumberFromName) {
		return true
	}

	return false
}

// SetAutoPopulateNumberFromName gets a reference to the given bool and assigns it to the AutoPopulateNumberFromName field.
func (o *MembershipHandlingType) SetAutoPopulateNumberFromName(v bool) {
	o.AutoPopulateNumberFromName = &v
}

// GetNameOnCardFromAltName returns the NameOnCardFromAltName field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetNameOnCardFromAltName() bool {
	if o == nil || IsNil(o.NameOnCardFromAltName) {
		var ret bool
		return ret
	}
	return *o.NameOnCardFromAltName
}

// GetNameOnCardFromAltNameOk returns a tuple with the NameOnCardFromAltName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetNameOnCardFromAltNameOk() (*bool, bool) {
	if o == nil || IsNil(o.NameOnCardFromAltName) {
		return nil, false
	}
	return o.NameOnCardFromAltName, true
}

// HasNameOnCardFromAltName returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasNameOnCardFromAltName() bool {
	if o != nil && !IsNil(o.NameOnCardFromAltName) {
		return true
	}

	return false
}

// SetNameOnCardFromAltName gets a reference to the given bool and assigns it to the NameOnCardFromAltName field.
func (o *MembershipHandlingType) SetNameOnCardFromAltName(v bool) {
	o.NameOnCardFromAltName = &v
}

// GetMoveMemPromotions returns the MoveMemPromotions field value if set, zero value otherwise.
func (o *MembershipHandlingType) GetMoveMemPromotions() bool {
	if o == nil || IsNil(o.MoveMemPromotions) {
		var ret bool
		return ret
	}
	return *o.MoveMemPromotions
}

// GetMoveMemPromotionsOk returns a tuple with the MoveMemPromotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipHandlingType) GetMoveMemPromotionsOk() (*bool, bool) {
	if o == nil || IsNil(o.MoveMemPromotions) {
		return nil, false
	}
	return o.MoveMemPromotions, true
}

// HasMoveMemPromotions returns a boolean if a field has been set.
func (o *MembershipHandlingType) HasMoveMemPromotions() bool {
	if o != nil && !IsNil(o.MoveMemPromotions) {
		return true
	}

	return false
}

// SetMoveMemPromotions gets a reference to the given bool and assigns it to the MoveMemPromotions field.
func (o *MembershipHandlingType) SetMoveMemPromotions(v bool) {
	o.MoveMemPromotions = &v
}

func (o MembershipHandlingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipHandlingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalDatabase) {
		toSerialize["externalDatabase"] = o.ExternalDatabase
	}
	if !IsNil(o.LevelRequired) {
		toSerialize["levelRequired"] = o.LevelRequired
	}
	if !IsNil(o.Fulfillment) {
		toSerialize["fulfillment"] = o.Fulfillment
	}
	if !IsNil(o.AllowCardNumberOverride) {
		toSerialize["allowCardNumberOverride"] = o.AllowCardNumberOverride
	}
	if !IsNil(o.EnrollmentCodeRequired) {
		toSerialize["enrollmentCodeRequired"] = o.EnrollmentCodeRequired
	}
	if !IsNil(o.SaveCardNumberHistory) {
		toSerialize["saveCardNumberHistory"] = o.SaveCardNumberHistory
	}
	if !IsNil(o.StatementUDFSet) {
		toSerialize["statementUDFSet"] = o.StatementUDFSet
	}
	if !IsNil(o.DefaultMembershipStatus) {
		toSerialize["defaultMembershipStatus"] = o.DefaultMembershipStatus
	}
	if !IsNil(o.NameProtected) {
		toSerialize["nameProtected"] = o.NameProtected
	}
	if !IsNil(o.AlternateNameProtected) {
		toSerialize["alternateNameProtected"] = o.AlternateNameProtected
	}
	if !IsNil(o.AutoGenerateReferenceNo) {
		toSerialize["autoGenerateReferenceNo"] = o.AutoGenerateReferenceNo
	}
	if !IsNil(o.AutoPopulateNumberFromName) {
		toSerialize["autoPopulateNumberFromName"] = o.AutoPopulateNumberFromName
	}
	if !IsNil(o.NameOnCardFromAltName) {
		toSerialize["nameOnCardFromAltName"] = o.NameOnCardFromAltName
	}
	if !IsNil(o.MoveMemPromotions) {
		toSerialize["moveMemPromotions"] = o.MoveMemPromotions
	}
	return toSerialize, nil
}

type NullableMembershipHandlingType struct {
	value *MembershipHandlingType
	isSet bool
}

func (v NullableMembershipHandlingType) Get() *MembershipHandlingType {
	return v.value
}

func (v *NullableMembershipHandlingType) Set(val *MembershipHandlingType) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipHandlingType) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipHandlingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipHandlingType(val *MembershipHandlingType) *NullableMembershipHandlingType {
	return &NullableMembershipHandlingType{value: val, isSet: true}
}

func (v NullableMembershipHandlingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipHandlingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


