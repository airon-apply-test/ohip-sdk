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
	"time"
)

// checks if the ExternalProfileSummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalProfileSummaryType{}

// ExternalProfileSummaryType Type provides the basic information about the external profile.
type ExternalProfileSummaryType struct {
	FormerName *ExternalProfileSummaryTypeFormerName `json:"formerName,omitempty"`
	AddressInfo *AddressInfoType `json:"addressInfo,omitempty"`
	TelephoneInfo *TelephoneInfoType `json:"telephoneInfo,omitempty"`
	EmailInfo *EmailInfoType `json:"emailInfo,omitempty"`
	ProfileMembership *ProfileMembershipType `json:"profileMembership,omitempty"`
	UrlInfo *URLInfoType `json:"urlInfo,omitempty"`
	// Generic type for a list of owners.
	Owners []OwnerType `json:"owners,omitempty"`
	ProfileType *ProfileTypeType `json:"profileType,omitempty"`
	StatusCode *ProfileStatusType `json:"statusCode,omitempty"`
	// Time stamp of the creation.
	CreateDateTime *time.Time `json:"createDateTime,omitempty"`
	// ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
	CreatorId *string `json:"creatorId,omitempty"`
	// Time stamp of last modification.
	LastModifyDateTime *time.Time `json:"lastModifyDateTime,omitempty"`
	// Identifies the last software system or person to modify a record.
	LastModifierId *string `json:"lastModifierId,omitempty"`
	// Property this profile is registered with.
	RegisteredProperty *string `json:"registeredProperty,omitempty"`
}

// NewExternalProfileSummaryType instantiates a new ExternalProfileSummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalProfileSummaryType() *ExternalProfileSummaryType {
	this := ExternalProfileSummaryType{}
	return &this
}

// NewExternalProfileSummaryTypeWithDefaults instantiates a new ExternalProfileSummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalProfileSummaryTypeWithDefaults() *ExternalProfileSummaryType {
	this := ExternalProfileSummaryType{}
	return &this
}

// GetFormerName returns the FormerName field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetFormerName() ExternalProfileSummaryTypeFormerName {
	if o == nil || IsNil(o.FormerName) {
		var ret ExternalProfileSummaryTypeFormerName
		return ret
	}
	return *o.FormerName
}

// GetFormerNameOk returns a tuple with the FormerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetFormerNameOk() (*ExternalProfileSummaryTypeFormerName, bool) {
	if o == nil || IsNil(o.FormerName) {
		return nil, false
	}
	return o.FormerName, true
}

// HasFormerName returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasFormerName() bool {
	if o != nil && !IsNil(o.FormerName) {
		return true
	}

	return false
}

// SetFormerName gets a reference to the given ExternalProfileSummaryTypeFormerName and assigns it to the FormerName field.
func (o *ExternalProfileSummaryType) SetFormerName(v ExternalProfileSummaryTypeFormerName) {
	o.FormerName = &v
}

// GetAddressInfo returns the AddressInfo field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetAddressInfo() AddressInfoType {
	if o == nil || IsNil(o.AddressInfo) {
		var ret AddressInfoType
		return ret
	}
	return *o.AddressInfo
}

// GetAddressInfoOk returns a tuple with the AddressInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetAddressInfoOk() (*AddressInfoType, bool) {
	if o == nil || IsNil(o.AddressInfo) {
		return nil, false
	}
	return o.AddressInfo, true
}

// HasAddressInfo returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasAddressInfo() bool {
	if o != nil && !IsNil(o.AddressInfo) {
		return true
	}

	return false
}

// SetAddressInfo gets a reference to the given AddressInfoType and assigns it to the AddressInfo field.
func (o *ExternalProfileSummaryType) SetAddressInfo(v AddressInfoType) {
	o.AddressInfo = &v
}

// GetTelephoneInfo returns the TelephoneInfo field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetTelephoneInfo() TelephoneInfoType {
	if o == nil || IsNil(o.TelephoneInfo) {
		var ret TelephoneInfoType
		return ret
	}
	return *o.TelephoneInfo
}

// GetTelephoneInfoOk returns a tuple with the TelephoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetTelephoneInfoOk() (*TelephoneInfoType, bool) {
	if o == nil || IsNil(o.TelephoneInfo) {
		return nil, false
	}
	return o.TelephoneInfo, true
}

// HasTelephoneInfo returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasTelephoneInfo() bool {
	if o != nil && !IsNil(o.TelephoneInfo) {
		return true
	}

	return false
}

// SetTelephoneInfo gets a reference to the given TelephoneInfoType and assigns it to the TelephoneInfo field.
func (o *ExternalProfileSummaryType) SetTelephoneInfo(v TelephoneInfoType) {
	o.TelephoneInfo = &v
}

// GetEmailInfo returns the EmailInfo field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetEmailInfo() EmailInfoType {
	if o == nil || IsNil(o.EmailInfo) {
		var ret EmailInfoType
		return ret
	}
	return *o.EmailInfo
}

// GetEmailInfoOk returns a tuple with the EmailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetEmailInfoOk() (*EmailInfoType, bool) {
	if o == nil || IsNil(o.EmailInfo) {
		return nil, false
	}
	return o.EmailInfo, true
}

// HasEmailInfo returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasEmailInfo() bool {
	if o != nil && !IsNil(o.EmailInfo) {
		return true
	}

	return false
}

// SetEmailInfo gets a reference to the given EmailInfoType and assigns it to the EmailInfo field.
func (o *ExternalProfileSummaryType) SetEmailInfo(v EmailInfoType) {
	o.EmailInfo = &v
}

// GetProfileMembership returns the ProfileMembership field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetProfileMembership() ProfileMembershipType {
	if o == nil || IsNil(o.ProfileMembership) {
		var ret ProfileMembershipType
		return ret
	}
	return *o.ProfileMembership
}

// GetProfileMembershipOk returns a tuple with the ProfileMembership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetProfileMembershipOk() (*ProfileMembershipType, bool) {
	if o == nil || IsNil(o.ProfileMembership) {
		return nil, false
	}
	return o.ProfileMembership, true
}

// HasProfileMembership returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasProfileMembership() bool {
	if o != nil && !IsNil(o.ProfileMembership) {
		return true
	}

	return false
}

// SetProfileMembership gets a reference to the given ProfileMembershipType and assigns it to the ProfileMembership field.
func (o *ExternalProfileSummaryType) SetProfileMembership(v ProfileMembershipType) {
	o.ProfileMembership = &v
}

// GetUrlInfo returns the UrlInfo field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetUrlInfo() URLInfoType {
	if o == nil || IsNil(o.UrlInfo) {
		var ret URLInfoType
		return ret
	}
	return *o.UrlInfo
}

// GetUrlInfoOk returns a tuple with the UrlInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetUrlInfoOk() (*URLInfoType, bool) {
	if o == nil || IsNil(o.UrlInfo) {
		return nil, false
	}
	return o.UrlInfo, true
}

// HasUrlInfo returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasUrlInfo() bool {
	if o != nil && !IsNil(o.UrlInfo) {
		return true
	}

	return false
}

// SetUrlInfo gets a reference to the given URLInfoType and assigns it to the UrlInfo field.
func (o *ExternalProfileSummaryType) SetUrlInfo(v URLInfoType) {
	o.UrlInfo = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetOwners() []OwnerType {
	if o == nil || IsNil(o.Owners) {
		var ret []OwnerType
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetOwnersOk() ([]OwnerType, bool) {
	if o == nil || IsNil(o.Owners) {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasOwners() bool {
	if o != nil && !IsNil(o.Owners) {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []OwnerType and assigns it to the Owners field.
func (o *ExternalProfileSummaryType) SetOwners(v []OwnerType) {
	o.Owners = v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetProfileType() ProfileTypeType {
	if o == nil || IsNil(o.ProfileType) {
		var ret ProfileTypeType
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetProfileTypeOk() (*ProfileTypeType, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given ProfileTypeType and assigns it to the ProfileType field.
func (o *ExternalProfileSummaryType) SetProfileType(v ProfileTypeType) {
	o.ProfileType = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetStatusCode() ProfileStatusType {
	if o == nil || IsNil(o.StatusCode) {
		var ret ProfileStatusType
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetStatusCodeOk() (*ProfileStatusType, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given ProfileStatusType and assigns it to the StatusCode field.
func (o *ExternalProfileSummaryType) SetStatusCode(v ProfileStatusType) {
	o.StatusCode = &v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetCreateDateTime() time.Time {
	if o == nil || IsNil(o.CreateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetCreateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDateTime) {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasCreateDateTime() bool {
	if o != nil && !IsNil(o.CreateDateTime) {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given time.Time and assigns it to the CreateDateTime field.
func (o *ExternalProfileSummaryType) SetCreateDateTime(v time.Time) {
	o.CreateDateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *ExternalProfileSummaryType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModifyDateTime returns the LastModifyDateTime field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetLastModifyDateTime() time.Time {
	if o == nil || IsNil(o.LastModifyDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyDateTime
}

// GetLastModifyDateTimeOk returns a tuple with the LastModifyDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetLastModifyDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyDateTime) {
		return nil, false
	}
	return o.LastModifyDateTime, true
}

// HasLastModifyDateTime returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasLastModifyDateTime() bool {
	if o != nil && !IsNil(o.LastModifyDateTime) {
		return true
	}

	return false
}

// SetLastModifyDateTime gets a reference to the given time.Time and assigns it to the LastModifyDateTime field.
func (o *ExternalProfileSummaryType) SetLastModifyDateTime(v time.Time) {
	o.LastModifyDateTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *ExternalProfileSummaryType) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetRegisteredProperty returns the RegisteredProperty field value if set, zero value otherwise.
func (o *ExternalProfileSummaryType) GetRegisteredProperty() string {
	if o == nil || IsNil(o.RegisteredProperty) {
		var ret string
		return ret
	}
	return *o.RegisteredProperty
}

// GetRegisteredPropertyOk returns a tuple with the RegisteredProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalProfileSummaryType) GetRegisteredPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.RegisteredProperty) {
		return nil, false
	}
	return o.RegisteredProperty, true
}

// HasRegisteredProperty returns a boolean if a field has been set.
func (o *ExternalProfileSummaryType) HasRegisteredProperty() bool {
	if o != nil && !IsNil(o.RegisteredProperty) {
		return true
	}

	return false
}

// SetRegisteredProperty gets a reference to the given string and assigns it to the RegisteredProperty field.
func (o *ExternalProfileSummaryType) SetRegisteredProperty(v string) {
	o.RegisteredProperty = &v
}

func (o ExternalProfileSummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalProfileSummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FormerName) {
		toSerialize["formerName"] = o.FormerName
	}
	if !IsNil(o.AddressInfo) {
		toSerialize["addressInfo"] = o.AddressInfo
	}
	if !IsNil(o.TelephoneInfo) {
		toSerialize["telephoneInfo"] = o.TelephoneInfo
	}
	if !IsNil(o.EmailInfo) {
		toSerialize["emailInfo"] = o.EmailInfo
	}
	if !IsNil(o.ProfileMembership) {
		toSerialize["profileMembership"] = o.ProfileMembership
	}
	if !IsNil(o.UrlInfo) {
		toSerialize["urlInfo"] = o.UrlInfo
	}
	if !IsNil(o.Owners) {
		toSerialize["owners"] = o.Owners
	}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
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
	if !IsNil(o.RegisteredProperty) {
		toSerialize["registeredProperty"] = o.RegisteredProperty
	}
	return toSerialize, nil
}

type NullableExternalProfileSummaryType struct {
	value *ExternalProfileSummaryType
	isSet bool
}

func (v NullableExternalProfileSummaryType) Get() *ExternalProfileSummaryType {
	return v.value
}

func (v *NullableExternalProfileSummaryType) Set(val *ExternalProfileSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalProfileSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalProfileSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalProfileSummaryType(val *ExternalProfileSummaryType) *NullableExternalProfileSummaryType {
	return &NullableExternalProfileSummaryType{value: val, isSet: true}
}

func (v NullableExternalProfileSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalProfileSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


