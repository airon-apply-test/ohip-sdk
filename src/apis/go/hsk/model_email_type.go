/*
OPERA Cloud Housekeeping Service API

APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the EmailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailType{}

// EmailType Information on an email for the customer.
type EmailType struct {
	// Defines the e-mail address.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Defines the purpose of the e-mail address (e.g. personal, business, listserve).
	Type *string `json:"type,omitempty"`
	// Describes the Type code
	TypeDescription *string `json:"typeDescription,omitempty"`
	// Supported Email format.
	EmailFormat *string `json:"emailFormat,omitempty"`
	// When true, indicates a primary information.
	PrimaryInd *bool `json:"primaryInd,omitempty"`
	// Display Order sequence.
	OrderSequence *float32 `json:"orderSequence,omitempty"`
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

// NewEmailType instantiates a new EmailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailType() *EmailType {
	this := EmailType{}
	return &this
}

// NewEmailTypeWithDefaults instantiates a new EmailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTypeWithDefaults() *EmailType {
	this := EmailType{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *EmailType) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *EmailType) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *EmailType) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EmailType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EmailType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EmailType) SetType(v string) {
	o.Type = &v
}

// GetTypeDescription returns the TypeDescription field value if set, zero value otherwise.
func (o *EmailType) GetTypeDescription() string {
	if o == nil || IsNil(o.TypeDescription) {
		var ret string
		return ret
	}
	return *o.TypeDescription
}

// GetTypeDescriptionOk returns a tuple with the TypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TypeDescription) {
		return nil, false
	}
	return o.TypeDescription, true
}

// HasTypeDescription returns a boolean if a field has been set.
func (o *EmailType) HasTypeDescription() bool {
	if o != nil && !IsNil(o.TypeDescription) {
		return true
	}

	return false
}

// SetTypeDescription gets a reference to the given string and assigns it to the TypeDescription field.
func (o *EmailType) SetTypeDescription(v string) {
	o.TypeDescription = &v
}

// GetEmailFormat returns the EmailFormat field value if set, zero value otherwise.
func (o *EmailType) GetEmailFormat() string {
	if o == nil || IsNil(o.EmailFormat) {
		var ret string
		return ret
	}
	return *o.EmailFormat
}

// GetEmailFormatOk returns a tuple with the EmailFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetEmailFormatOk() (*string, bool) {
	if o == nil || IsNil(o.EmailFormat) {
		return nil, false
	}
	return o.EmailFormat, true
}

// HasEmailFormat returns a boolean if a field has been set.
func (o *EmailType) HasEmailFormat() bool {
	if o != nil && !IsNil(o.EmailFormat) {
		return true
	}

	return false
}

// SetEmailFormat gets a reference to the given string and assigns it to the EmailFormat field.
func (o *EmailType) SetEmailFormat(v string) {
	o.EmailFormat = &v
}

// GetPrimaryInd returns the PrimaryInd field value if set, zero value otherwise.
func (o *EmailType) GetPrimaryInd() bool {
	if o == nil || IsNil(o.PrimaryInd) {
		var ret bool
		return ret
	}
	return *o.PrimaryInd
}

// GetPrimaryIndOk returns a tuple with the PrimaryInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetPrimaryIndOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryInd) {
		return nil, false
	}
	return o.PrimaryInd, true
}

// HasPrimaryInd returns a boolean if a field has been set.
func (o *EmailType) HasPrimaryInd() bool {
	if o != nil && !IsNil(o.PrimaryInd) {
		return true
	}

	return false
}

// SetPrimaryInd gets a reference to the given bool and assigns it to the PrimaryInd field.
func (o *EmailType) SetPrimaryInd(v bool) {
	o.PrimaryInd = &v
}

// GetOrderSequence returns the OrderSequence field value if set, zero value otherwise.
func (o *EmailType) GetOrderSequence() float32 {
	if o == nil || IsNil(o.OrderSequence) {
		var ret float32
		return ret
	}
	return *o.OrderSequence
}

// GetOrderSequenceOk returns a tuple with the OrderSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetOrderSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderSequence) {
		return nil, false
	}
	return o.OrderSequence, true
}

// HasOrderSequence returns a boolean if a field has been set.
func (o *EmailType) HasOrderSequence() bool {
	if o != nil && !IsNil(o.OrderSequence) {
		return true
	}

	return false
}

// SetOrderSequence gets a reference to the given float32 and assigns it to the OrderSequence field.
func (o *EmailType) SetOrderSequence(v float32) {
	o.OrderSequence = &v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *EmailType) GetCreateDateTime() time.Time {
	if o == nil || IsNil(o.CreateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetCreateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDateTime) {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *EmailType) HasCreateDateTime() bool {
	if o != nil && !IsNil(o.CreateDateTime) {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given time.Time and assigns it to the CreateDateTime field.
func (o *EmailType) SetCreateDateTime(v time.Time) {
	o.CreateDateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *EmailType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *EmailType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *EmailType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModifyDateTime returns the LastModifyDateTime field value if set, zero value otherwise.
func (o *EmailType) GetLastModifyDateTime() time.Time {
	if o == nil || IsNil(o.LastModifyDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyDateTime
}

// GetLastModifyDateTimeOk returns a tuple with the LastModifyDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetLastModifyDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyDateTime) {
		return nil, false
	}
	return o.LastModifyDateTime, true
}

// HasLastModifyDateTime returns a boolean if a field has been set.
func (o *EmailType) HasLastModifyDateTime() bool {
	if o != nil && !IsNil(o.LastModifyDateTime) {
		return true
	}

	return false
}

// SetLastModifyDateTime gets a reference to the given time.Time and assigns it to the LastModifyDateTime field.
func (o *EmailType) SetLastModifyDateTime(v time.Time) {
	o.LastModifyDateTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *EmailType) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *EmailType) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *EmailType) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetPurgeDate returns the PurgeDate field value if set, zero value otherwise.
func (o *EmailType) GetPurgeDate() string {
	if o == nil || IsNil(o.PurgeDate) {
		var ret string
		return ret
	}
	return *o.PurgeDate
}

// GetPurgeDateOk returns a tuple with the PurgeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailType) GetPurgeDateOk() (*string, bool) {
	if o == nil || IsNil(o.PurgeDate) {
		return nil, false
	}
	return o.PurgeDate, true
}

// HasPurgeDate returns a boolean if a field has been set.
func (o *EmailType) HasPurgeDate() bool {
	if o != nil && !IsNil(o.PurgeDate) {
		return true
	}

	return false
}

// SetPurgeDate gets a reference to the given string and assigns it to the PurgeDate field.
func (o *EmailType) SetPurgeDate(v string) {
	o.PurgeDate = &v
}

func (o EmailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeDescription) {
		toSerialize["typeDescription"] = o.TypeDescription
	}
	if !IsNil(o.EmailFormat) {
		toSerialize["emailFormat"] = o.EmailFormat
	}
	if !IsNil(o.PrimaryInd) {
		toSerialize["primaryInd"] = o.PrimaryInd
	}
	if !IsNil(o.OrderSequence) {
		toSerialize["orderSequence"] = o.OrderSequence
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

type NullableEmailType struct {
	value *EmailType
	isSet bool
}

func (v NullableEmailType) Get() *EmailType {
	return v.value
}

func (v *NullableEmailType) Set(val *EmailType) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailType) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailType(val *EmailType) *NullableEmailType {
	return &NullableEmailType{value: val, isSet: true}
}

func (v NullableEmailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


