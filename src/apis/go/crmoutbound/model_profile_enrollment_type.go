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

// checks if the ProfileEnrollmentType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileEnrollmentType{}

// ProfileEnrollmentType Type provides the detailed information about the profile and its children.
type ProfileEnrollmentType struct {
	Customer *CustomerType `json:"customer,omitempty"`
	Addresses *GuestProfileTypeAddresses `json:"addresses,omitempty"`
	Telephones *ResCommunicationTypeTelephones `json:"telephones,omitempty"`
	Emails *ResCommunicationTypeEmails `json:"emails,omitempty"`
	Urls *ProfileEnrollmentTypeUrls `json:"urls,omitempty"`
	ProfileMemberships *ProfileEnrollmentTypeProfileMemberships `json:"profileMemberships,omitempty"`
	Keywords *ProfileTypeKeywords `json:"keywords,omitempty"`
	ProfileType *ProfileTypeType `json:"profileType,omitempty"`
	StatusCode *ProfileStatusType `json:"statusCode,omitempty"`
	// Hotel which this profile is registered with. This attribute is not used for configuration.
	RegisteredProperty *string `json:"registeredProperty,omitempty"`
	// Hotel which this profile is to be registered. This attribute is only used during creation of profile.
	RequestForHotel *string `json:"requestForHotel,omitempty"`
	// Time stamp of the creation.
	CreateDateTime *time.Time `json:"createDateTime,omitempty"`
	// ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
	CreatorId *string `json:"creatorId,omitempty"`
	// Time stamp of last modification.
	LastModifyDateTime *time.Time `json:"lastModifyDateTime,omitempty"`
	// Identifies the last software system or person to modify a record.
	LastModifierId *string `json:"lastModifierId,omitempty"`
	ProfileId *UniqueIDType `json:"profileId,omitempty"`
	ReservationId *UniqueIDType `json:"reservationId,omitempty"`
	EnrollmentCode *CodeDescriptionType `json:"enrollmentCode,omitempty"`
	// Enrollment details will be fetched from this External database.
	ExternalDatabaseID *string `json:"externalDatabaseID,omitempty"`
	// Hotel Code, It is used to filter hotel specific children to this specific hotel code.
	HotelId *string `json:"hotelId,omitempty"`
	// UserID/LoginID of the user who is enrolling the Guest.
	UserId *string `json:"userId,omitempty"`
	// EmployeeID of the user who is enrolling the Guest.
	EmployeeId *string `json:"employeeId,omitempty"`
}

// NewProfileEnrollmentType instantiates a new ProfileEnrollmentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentType() *ProfileEnrollmentType {
	this := ProfileEnrollmentType{}
	return &this
}

// NewProfileEnrollmentTypeWithDefaults instantiates a new ProfileEnrollmentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentTypeWithDefaults() *ProfileEnrollmentType {
	this := ProfileEnrollmentType{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetCustomer() CustomerType {
	if o == nil || IsNil(o.Customer) {
		var ret CustomerType
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetCustomerOk() (*CustomerType, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CustomerType and assigns it to the Customer field.
func (o *ProfileEnrollmentType) SetCustomer(v CustomerType) {
	o.Customer = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetAddresses() GuestProfileTypeAddresses {
	if o == nil || IsNil(o.Addresses) {
		var ret GuestProfileTypeAddresses
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetAddressesOk() (*GuestProfileTypeAddresses, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given GuestProfileTypeAddresses and assigns it to the Addresses field.
func (o *ProfileEnrollmentType) SetAddresses(v GuestProfileTypeAddresses) {
	o.Addresses = &v
}

// GetTelephones returns the Telephones field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetTelephones() ResCommunicationTypeTelephones {
	if o == nil || IsNil(o.Telephones) {
		var ret ResCommunicationTypeTelephones
		return ret
	}
	return *o.Telephones
}

// GetTelephonesOk returns a tuple with the Telephones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetTelephonesOk() (*ResCommunicationTypeTelephones, bool) {
	if o == nil || IsNil(o.Telephones) {
		return nil, false
	}
	return o.Telephones, true
}

// HasTelephones returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasTelephones() bool {
	if o != nil && !IsNil(o.Telephones) {
		return true
	}

	return false
}

// SetTelephones gets a reference to the given ResCommunicationTypeTelephones and assigns it to the Telephones field.
func (o *ProfileEnrollmentType) SetTelephones(v ResCommunicationTypeTelephones) {
	o.Telephones = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetEmails() ResCommunicationTypeEmails {
	if o == nil || IsNil(o.Emails) {
		var ret ResCommunicationTypeEmails
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetEmailsOk() (*ResCommunicationTypeEmails, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given ResCommunicationTypeEmails and assigns it to the Emails field.
func (o *ProfileEnrollmentType) SetEmails(v ResCommunicationTypeEmails) {
	o.Emails = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetUrls() ProfileEnrollmentTypeUrls {
	if o == nil || IsNil(o.Urls) {
		var ret ProfileEnrollmentTypeUrls
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetUrlsOk() (*ProfileEnrollmentTypeUrls, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given ProfileEnrollmentTypeUrls and assigns it to the Urls field.
func (o *ProfileEnrollmentType) SetUrls(v ProfileEnrollmentTypeUrls) {
	o.Urls = &v
}

// GetProfileMemberships returns the ProfileMemberships field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetProfileMemberships() ProfileEnrollmentTypeProfileMemberships {
	if o == nil || IsNil(o.ProfileMemberships) {
		var ret ProfileEnrollmentTypeProfileMemberships
		return ret
	}
	return *o.ProfileMemberships
}

// GetProfileMembershipsOk returns a tuple with the ProfileMemberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetProfileMembershipsOk() (*ProfileEnrollmentTypeProfileMemberships, bool) {
	if o == nil || IsNil(o.ProfileMemberships) {
		return nil, false
	}
	return o.ProfileMemberships, true
}

// HasProfileMemberships returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasProfileMemberships() bool {
	if o != nil && !IsNil(o.ProfileMemberships) {
		return true
	}

	return false
}

// SetProfileMemberships gets a reference to the given ProfileEnrollmentTypeProfileMemberships and assigns it to the ProfileMemberships field.
func (o *ProfileEnrollmentType) SetProfileMemberships(v ProfileEnrollmentTypeProfileMemberships) {
	o.ProfileMemberships = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetKeywords() ProfileTypeKeywords {
	if o == nil || IsNil(o.Keywords) {
		var ret ProfileTypeKeywords
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetKeywordsOk() (*ProfileTypeKeywords, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given ProfileTypeKeywords and assigns it to the Keywords field.
func (o *ProfileEnrollmentType) SetKeywords(v ProfileTypeKeywords) {
	o.Keywords = &v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetProfileType() ProfileTypeType {
	if o == nil || IsNil(o.ProfileType) {
		var ret ProfileTypeType
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetProfileTypeOk() (*ProfileTypeType, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given ProfileTypeType and assigns it to the ProfileType field.
func (o *ProfileEnrollmentType) SetProfileType(v ProfileTypeType) {
	o.ProfileType = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetStatusCode() ProfileStatusType {
	if o == nil || IsNil(o.StatusCode) {
		var ret ProfileStatusType
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetStatusCodeOk() (*ProfileStatusType, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given ProfileStatusType and assigns it to the StatusCode field.
func (o *ProfileEnrollmentType) SetStatusCode(v ProfileStatusType) {
	o.StatusCode = &v
}

// GetRegisteredProperty returns the RegisteredProperty field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetRegisteredProperty() string {
	if o == nil || IsNil(o.RegisteredProperty) {
		var ret string
		return ret
	}
	return *o.RegisteredProperty
}

// GetRegisteredPropertyOk returns a tuple with the RegisteredProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetRegisteredPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.RegisteredProperty) {
		return nil, false
	}
	return o.RegisteredProperty, true
}

// HasRegisteredProperty returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasRegisteredProperty() bool {
	if o != nil && !IsNil(o.RegisteredProperty) {
		return true
	}

	return false
}

// SetRegisteredProperty gets a reference to the given string and assigns it to the RegisteredProperty field.
func (o *ProfileEnrollmentType) SetRegisteredProperty(v string) {
	o.RegisteredProperty = &v
}

// GetRequestForHotel returns the RequestForHotel field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetRequestForHotel() string {
	if o == nil || IsNil(o.RequestForHotel) {
		var ret string
		return ret
	}
	return *o.RequestForHotel
}

// GetRequestForHotelOk returns a tuple with the RequestForHotel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetRequestForHotelOk() (*string, bool) {
	if o == nil || IsNil(o.RequestForHotel) {
		return nil, false
	}
	return o.RequestForHotel, true
}

// HasRequestForHotel returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasRequestForHotel() bool {
	if o != nil && !IsNil(o.RequestForHotel) {
		return true
	}

	return false
}

// SetRequestForHotel gets a reference to the given string and assigns it to the RequestForHotel field.
func (o *ProfileEnrollmentType) SetRequestForHotel(v string) {
	o.RequestForHotel = &v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetCreateDateTime() time.Time {
	if o == nil || IsNil(o.CreateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetCreateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDateTime) {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasCreateDateTime() bool {
	if o != nil && !IsNil(o.CreateDateTime) {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given time.Time and assigns it to the CreateDateTime field.
func (o *ProfileEnrollmentType) SetCreateDateTime(v time.Time) {
	o.CreateDateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *ProfileEnrollmentType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModifyDateTime returns the LastModifyDateTime field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetLastModifyDateTime() time.Time {
	if o == nil || IsNil(o.LastModifyDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyDateTime
}

// GetLastModifyDateTimeOk returns a tuple with the LastModifyDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetLastModifyDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyDateTime) {
		return nil, false
	}
	return o.LastModifyDateTime, true
}

// HasLastModifyDateTime returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasLastModifyDateTime() bool {
	if o != nil && !IsNil(o.LastModifyDateTime) {
		return true
	}

	return false
}

// SetLastModifyDateTime gets a reference to the given time.Time and assigns it to the LastModifyDateTime field.
func (o *ProfileEnrollmentType) SetLastModifyDateTime(v time.Time) {
	o.LastModifyDateTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *ProfileEnrollmentType) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetProfileId() UniqueIDType {
	if o == nil || IsNil(o.ProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given UniqueIDType and assigns it to the ProfileId field.
func (o *ProfileEnrollmentType) SetProfileId(v UniqueIDType) {
	o.ProfileId = &v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetReservationId() UniqueIDType {
	if o == nil || IsNil(o.ReservationId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetReservationIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given UniqueIDType and assigns it to the ReservationId field.
func (o *ProfileEnrollmentType) SetReservationId(v UniqueIDType) {
	o.ReservationId = &v
}

// GetEnrollmentCode returns the EnrollmentCode field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetEnrollmentCode() CodeDescriptionType {
	if o == nil || IsNil(o.EnrollmentCode) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.EnrollmentCode
}

// GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetEnrollmentCodeOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.EnrollmentCode) {
		return nil, false
	}
	return o.EnrollmentCode, true
}

// HasEnrollmentCode returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasEnrollmentCode() bool {
	if o != nil && !IsNil(o.EnrollmentCode) {
		return true
	}

	return false
}

// SetEnrollmentCode gets a reference to the given CodeDescriptionType and assigns it to the EnrollmentCode field.
func (o *ProfileEnrollmentType) SetEnrollmentCode(v CodeDescriptionType) {
	o.EnrollmentCode = &v
}

// GetExternalDatabaseID returns the ExternalDatabaseID field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetExternalDatabaseID() string {
	if o == nil || IsNil(o.ExternalDatabaseID) {
		var ret string
		return ret
	}
	return *o.ExternalDatabaseID
}

// GetExternalDatabaseIDOk returns a tuple with the ExternalDatabaseID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetExternalDatabaseIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalDatabaseID) {
		return nil, false
	}
	return o.ExternalDatabaseID, true
}

// HasExternalDatabaseID returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasExternalDatabaseID() bool {
	if o != nil && !IsNil(o.ExternalDatabaseID) {
		return true
	}

	return false
}

// SetExternalDatabaseID gets a reference to the given string and assigns it to the ExternalDatabaseID field.
func (o *ProfileEnrollmentType) SetExternalDatabaseID(v string) {
	o.ExternalDatabaseID = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ProfileEnrollmentType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ProfileEnrollmentType) SetUserId(v string) {
	o.UserId = &v
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *ProfileEnrollmentType) GetEmployeeId() string {
	if o == nil || IsNil(o.EmployeeId) {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentType) GetEmployeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeId) {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *ProfileEnrollmentType) HasEmployeeId() bool {
	if o != nil && !IsNil(o.EmployeeId) {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *ProfileEnrollmentType) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

func (o ProfileEnrollmentType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileEnrollmentType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Telephones) {
		toSerialize["telephones"] = o.Telephones
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.ProfileMemberships) {
		toSerialize["profileMemberships"] = o.ProfileMemberships
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !IsNil(o.RegisteredProperty) {
		toSerialize["registeredProperty"] = o.RegisteredProperty
	}
	if !IsNil(o.RequestForHotel) {
		toSerialize["requestForHotel"] = o.RequestForHotel
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
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.EnrollmentCode) {
		toSerialize["enrollmentCode"] = o.EnrollmentCode
	}
	if !IsNil(o.ExternalDatabaseID) {
		toSerialize["externalDatabaseID"] = o.ExternalDatabaseID
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.EmployeeId) {
		toSerialize["employeeId"] = o.EmployeeId
	}
	return toSerialize, nil
}

type NullableProfileEnrollmentType struct {
	value *ProfileEnrollmentType
	isSet bool
}

func (v NullableProfileEnrollmentType) Get() *ProfileEnrollmentType {
	return v.value
}

func (v *NullableProfileEnrollmentType) Set(val *ProfileEnrollmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileEnrollmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileEnrollmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileEnrollmentType(val *ProfileEnrollmentType) *NullableProfileEnrollmentType {
	return &NullableProfileEnrollmentType{value: val, isSet: true}
}

func (v NullableProfileEnrollmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileEnrollmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


