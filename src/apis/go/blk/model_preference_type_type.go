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
	"time"
)

// checks if the PreferenceTypeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreferenceTypeType{}

// PreferenceTypeType Preference details for the profile.
type PreferenceTypeType struct {
	// Collection of Preferences for the profile.
	Preference []PreferenceType `json:"preference,omitempty"`
	// Preference group code.
	PreferenceType *string `json:"preferenceType,omitempty"`
	// Preference group description.
	PreferenceTypeDescription *string `json:"preferenceTypeDescription,omitempty"`
	// Preference Sequence.
	Sequence *string `json:"sequence,omitempty"`
	// Maximum quantity of preferences allowed per preference group.
	MaxQuantity *int32 `json:"maxQuantity,omitempty"`
	// Available quantity of preferences (maximum quantity - Existing preferences)per preference group.
	AvailableQuantity *int32 `json:"availableQuantity,omitempty"`
	// Maximum quantity of preferences used by any resort per preference group.
	MaxResortUsedQuantity *int32 `json:"maxResortUsedQuantity,omitempty"`
	// Whether this preference is reservation preference or not.
	ReservationPreference *bool `json:"reservationPreference,omitempty"`
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
}

// NewPreferenceTypeType instantiates a new PreferenceTypeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferenceTypeType() *PreferenceTypeType {
	this := PreferenceTypeType{}
	return &this
}

// NewPreferenceTypeTypeWithDefaults instantiates a new PreferenceTypeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceTypeTypeWithDefaults() *PreferenceTypeType {
	this := PreferenceTypeType{}
	return &this
}

// GetPreference returns the Preference field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetPreference() []PreferenceType {
	if o == nil || IsNil(o.Preference) {
		var ret []PreferenceType
		return ret
	}
	return o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetPreferenceOk() ([]PreferenceType, bool) {
	if o == nil || IsNil(o.Preference) {
		return nil, false
	}
	return o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasPreference() bool {
	if o != nil && !IsNil(o.Preference) {
		return true
	}

	return false
}

// SetPreference gets a reference to the given []PreferenceType and assigns it to the Preference field.
func (o *PreferenceTypeType) SetPreference(v []PreferenceType) {
	o.Preference = v
}

// GetPreferenceType returns the PreferenceType field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetPreferenceType() string {
	if o == nil || IsNil(o.PreferenceType) {
		var ret string
		return ret
	}
	return *o.PreferenceType
}

// GetPreferenceTypeOk returns a tuple with the PreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetPreferenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PreferenceType) {
		return nil, false
	}
	return o.PreferenceType, true
}

// HasPreferenceType returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasPreferenceType() bool {
	if o != nil && !IsNil(o.PreferenceType) {
		return true
	}

	return false
}

// SetPreferenceType gets a reference to the given string and assigns it to the PreferenceType field.
func (o *PreferenceTypeType) SetPreferenceType(v string) {
	o.PreferenceType = &v
}

// GetPreferenceTypeDescription returns the PreferenceTypeDescription field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetPreferenceTypeDescription() string {
	if o == nil || IsNil(o.PreferenceTypeDescription) {
		var ret string
		return ret
	}
	return *o.PreferenceTypeDescription
}

// GetPreferenceTypeDescriptionOk returns a tuple with the PreferenceTypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetPreferenceTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PreferenceTypeDescription) {
		return nil, false
	}
	return o.PreferenceTypeDescription, true
}

// HasPreferenceTypeDescription returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasPreferenceTypeDescription() bool {
	if o != nil && !IsNil(o.PreferenceTypeDescription) {
		return true
	}

	return false
}

// SetPreferenceTypeDescription gets a reference to the given string and assigns it to the PreferenceTypeDescription field.
func (o *PreferenceTypeType) SetPreferenceTypeDescription(v string) {
	o.PreferenceTypeDescription = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetSequence() string {
	if o == nil || IsNil(o.Sequence) {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetSequenceOk() (*string, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *PreferenceTypeType) SetSequence(v string) {
	o.Sequence = &v
}

// GetMaxQuantity returns the MaxQuantity field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetMaxQuantity() int32 {
	if o == nil || IsNil(o.MaxQuantity) {
		var ret int32
		return ret
	}
	return *o.MaxQuantity
}

// GetMaxQuantityOk returns a tuple with the MaxQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetMaxQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxQuantity) {
		return nil, false
	}
	return o.MaxQuantity, true
}

// HasMaxQuantity returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasMaxQuantity() bool {
	if o != nil && !IsNil(o.MaxQuantity) {
		return true
	}

	return false
}

// SetMaxQuantity gets a reference to the given int32 and assigns it to the MaxQuantity field.
func (o *PreferenceTypeType) SetMaxQuantity(v int32) {
	o.MaxQuantity = &v
}

// GetAvailableQuantity returns the AvailableQuantity field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetAvailableQuantity() int32 {
	if o == nil || IsNil(o.AvailableQuantity) {
		var ret int32
		return ret
	}
	return *o.AvailableQuantity
}

// GetAvailableQuantityOk returns a tuple with the AvailableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetAvailableQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableQuantity) {
		return nil, false
	}
	return o.AvailableQuantity, true
}

// HasAvailableQuantity returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasAvailableQuantity() bool {
	if o != nil && !IsNil(o.AvailableQuantity) {
		return true
	}

	return false
}

// SetAvailableQuantity gets a reference to the given int32 and assigns it to the AvailableQuantity field.
func (o *PreferenceTypeType) SetAvailableQuantity(v int32) {
	o.AvailableQuantity = &v
}

// GetMaxResortUsedQuantity returns the MaxResortUsedQuantity field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetMaxResortUsedQuantity() int32 {
	if o == nil || IsNil(o.MaxResortUsedQuantity) {
		var ret int32
		return ret
	}
	return *o.MaxResortUsedQuantity
}

// GetMaxResortUsedQuantityOk returns a tuple with the MaxResortUsedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetMaxResortUsedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResortUsedQuantity) {
		return nil, false
	}
	return o.MaxResortUsedQuantity, true
}

// HasMaxResortUsedQuantity returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasMaxResortUsedQuantity() bool {
	if o != nil && !IsNil(o.MaxResortUsedQuantity) {
		return true
	}

	return false
}

// SetMaxResortUsedQuantity gets a reference to the given int32 and assigns it to the MaxResortUsedQuantity field.
func (o *PreferenceTypeType) SetMaxResortUsedQuantity(v int32) {
	o.MaxResortUsedQuantity = &v
}

// GetReservationPreference returns the ReservationPreference field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetReservationPreference() bool {
	if o == nil || IsNil(o.ReservationPreference) {
		var ret bool
		return ret
	}
	return *o.ReservationPreference
}

// GetReservationPreferenceOk returns a tuple with the ReservationPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetReservationPreferenceOk() (*bool, bool) {
	if o == nil || IsNil(o.ReservationPreference) {
		return nil, false
	}
	return o.ReservationPreference, true
}

// HasReservationPreference returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasReservationPreference() bool {
	if o != nil && !IsNil(o.ReservationPreference) {
		return true
	}

	return false
}

// SetReservationPreference gets a reference to the given bool and assigns it to the ReservationPreference field.
func (o *PreferenceTypeType) SetReservationPreference(v bool) {
	o.ReservationPreference = &v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetCreateDateTime() time.Time {
	if o == nil || IsNil(o.CreateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetCreateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDateTime) {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasCreateDateTime() bool {
	if o != nil && !IsNil(o.CreateDateTime) {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given time.Time and assigns it to the CreateDateTime field.
func (o *PreferenceTypeType) SetCreateDateTime(v time.Time) {
	o.CreateDateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *PreferenceTypeType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModifyDateTime returns the LastModifyDateTime field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetLastModifyDateTime() time.Time {
	if o == nil || IsNil(o.LastModifyDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyDateTime
}

// GetLastModifyDateTimeOk returns a tuple with the LastModifyDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetLastModifyDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyDateTime) {
		return nil, false
	}
	return o.LastModifyDateTime, true
}

// HasLastModifyDateTime returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasLastModifyDateTime() bool {
	if o != nil && !IsNil(o.LastModifyDateTime) {
		return true
	}

	return false
}

// SetLastModifyDateTime gets a reference to the given time.Time and assigns it to the LastModifyDateTime field.
func (o *PreferenceTypeType) SetLastModifyDateTime(v time.Time) {
	o.LastModifyDateTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *PreferenceTypeType) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetPurgeDate returns the PurgeDate field value if set, zero value otherwise.
func (o *PreferenceTypeType) GetPurgeDate() string {
	if o == nil || IsNil(o.PurgeDate) {
		var ret string
		return ret
	}
	return *o.PurgeDate
}

// GetPurgeDateOk returns a tuple with the PurgeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferenceTypeType) GetPurgeDateOk() (*string, bool) {
	if o == nil || IsNil(o.PurgeDate) {
		return nil, false
	}
	return o.PurgeDate, true
}

// HasPurgeDate returns a boolean if a field has been set.
func (o *PreferenceTypeType) HasPurgeDate() bool {
	if o != nil && !IsNil(o.PurgeDate) {
		return true
	}

	return false
}

// SetPurgeDate gets a reference to the given string and assigns it to the PurgeDate field.
func (o *PreferenceTypeType) SetPurgeDate(v string) {
	o.PurgeDate = &v
}

func (o PreferenceTypeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreferenceTypeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Preference) {
		toSerialize["preference"] = o.Preference
	}
	if !IsNil(o.PreferenceType) {
		toSerialize["preferenceType"] = o.PreferenceType
	}
	if !IsNil(o.PreferenceTypeDescription) {
		toSerialize["preferenceTypeDescription"] = o.PreferenceTypeDescription
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.MaxQuantity) {
		toSerialize["maxQuantity"] = o.MaxQuantity
	}
	if !IsNil(o.AvailableQuantity) {
		toSerialize["availableQuantity"] = o.AvailableQuantity
	}
	if !IsNil(o.MaxResortUsedQuantity) {
		toSerialize["maxResortUsedQuantity"] = o.MaxResortUsedQuantity
	}
	if !IsNil(o.ReservationPreference) {
		toSerialize["reservationPreference"] = o.ReservationPreference
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
	return toSerialize, nil
}

type NullablePreferenceTypeType struct {
	value *PreferenceTypeType
	isSet bool
}

func (v NullablePreferenceTypeType) Get() *PreferenceTypeType {
	return v.value
}

func (v *NullablePreferenceTypeType) Set(val *PreferenceTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferenceTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferenceTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferenceTypeType(val *PreferenceTypeType) *NullablePreferenceTypeType {
	return &NullablePreferenceTypeType{value: val, isSet: true}
}

func (v NullablePreferenceTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferenceTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


