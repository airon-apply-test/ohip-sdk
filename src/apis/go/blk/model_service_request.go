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

// checks if the ServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceRequest{}

// ServiceRequest Service request.
type ServiceRequest struct {
	ServiceRequestId *UniqueIDType `json:"serviceRequestId,omitempty"`
	// Hotel Code of the service request.
	HotelId *string `json:"hotelId,omitempty"`
	// Service request code.
	Code *string `json:"code,omitempty"`
	Status *ServiceRequestStatusType `json:"status,omitempty"`
	// The priority level of the service request.
	Priority *string `json:"priority,omitempty"`
	Department *CodeDescriptionType `json:"department,omitempty"`
	// Unique Id that references an object uniquely in the system.
	ReservationIdList []UniqueIDType `json:"reservationIdList,omitempty"`
	ProfileId *ProfileId `json:"profileId,omitempty"`
	// The guest name.
	GuestName *string `json:"guestName,omitempty"`
	// The room number in which the service request was created.
	Room *string `json:"room,omitempty"`
	// The date-time in which the service request was opened.
	OpenDate *time.Time `json:"openDate,omitempty"`
	// The user who guaranteed the completion of the service request.
	GuaranteedBy *string `json:"guaranteedBy,omitempty"`
	// Description of the service request.
	Comment *string `json:"comment,omitempty"`
	// Description of the action taken to complete the request.
	Action *string `json:"action,omitempty"`
	// Communication method picked from guest profile.
	GuestContactMethod *string `json:"guestContactMethod,omitempty"`
	// The date-time in which the service request was completed.
	CompletionDate *time.Time `json:"completionDate,omitempty"`
	// The date-time in which the service request was closed.
	ClosedDate *time.Time `json:"closedDate,omitempty"`
	// The user who closed the service request.
	ClosedBy *string `json:"closedBy,omitempty"`
	// The follow up description.
	CloseDescription *string `json:"closeDescription,omitempty"`
}

// NewServiceRequest instantiates a new ServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRequest() *ServiceRequest {
	this := ServiceRequest{}
	return &this
}

// NewServiceRequestWithDefaults instantiates a new ServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRequestWithDefaults() *ServiceRequest {
	this := ServiceRequest{}
	return &this
}

// GetServiceRequestId returns the ServiceRequestId field value if set, zero value otherwise.
func (o *ServiceRequest) GetServiceRequestId() UniqueIDType {
	if o == nil || IsNil(o.ServiceRequestId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ServiceRequestId
}

// GetServiceRequestIdOk returns a tuple with the ServiceRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetServiceRequestIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ServiceRequestId) {
		return nil, false
	}
	return o.ServiceRequestId, true
}

// HasServiceRequestId returns a boolean if a field has been set.
func (o *ServiceRequest) HasServiceRequestId() bool {
	if o != nil && !IsNil(o.ServiceRequestId) {
		return true
	}

	return false
}

// SetServiceRequestId gets a reference to the given UniqueIDType and assigns it to the ServiceRequestId field.
func (o *ServiceRequest) SetServiceRequestId(v UniqueIDType) {
	o.ServiceRequestId = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ServiceRequest) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ServiceRequest) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ServiceRequest) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ServiceRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ServiceRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ServiceRequest) SetCode(v string) {
	o.Code = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceRequest) GetStatus() ServiceRequestStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ServiceRequestStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetStatusOk() (*ServiceRequestStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ServiceRequestStatusType and assigns it to the Status field.
func (o *ServiceRequest) SetStatus(v ServiceRequestStatusType) {
	o.Status = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ServiceRequest) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ServiceRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ServiceRequest) SetPriority(v string) {
	o.Priority = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *ServiceRequest) GetDepartment() CodeDescriptionType {
	if o == nil || IsNil(o.Department) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetDepartmentOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *ServiceRequest) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given CodeDescriptionType and assigns it to the Department field.
func (o *ServiceRequest) SetDepartment(v CodeDescriptionType) {
	o.Department = &v
}

// GetReservationIdList returns the ReservationIdList field value if set, zero value otherwise.
func (o *ServiceRequest) GetReservationIdList() []UniqueIDType {
	if o == nil || IsNil(o.ReservationIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.ReservationIdList
}

// GetReservationIdListOk returns a tuple with the ReservationIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetReservationIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationIdList) {
		return nil, false
	}
	return o.ReservationIdList, true
}

// HasReservationIdList returns a boolean if a field has been set.
func (o *ServiceRequest) HasReservationIdList() bool {
	if o != nil && !IsNil(o.ReservationIdList) {
		return true
	}

	return false
}

// SetReservationIdList gets a reference to the given []UniqueIDType and assigns it to the ReservationIdList field.
func (o *ServiceRequest) SetReservationIdList(v []UniqueIDType) {
	o.ReservationIdList = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ServiceRequest) GetProfileId() ProfileId {
	if o == nil || IsNil(o.ProfileId) {
		var ret ProfileId
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetProfileIdOk() (*ProfileId, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ServiceRequest) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given ProfileId and assigns it to the ProfileId field.
func (o *ServiceRequest) SetProfileId(v ProfileId) {
	o.ProfileId = &v
}

// GetGuestName returns the GuestName field value if set, zero value otherwise.
func (o *ServiceRequest) GetGuestName() string {
	if o == nil || IsNil(o.GuestName) {
		var ret string
		return ret
	}
	return *o.GuestName
}

// GetGuestNameOk returns a tuple with the GuestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetGuestNameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestName) {
		return nil, false
	}
	return o.GuestName, true
}

// HasGuestName returns a boolean if a field has been set.
func (o *ServiceRequest) HasGuestName() bool {
	if o != nil && !IsNil(o.GuestName) {
		return true
	}

	return false
}

// SetGuestName gets a reference to the given string and assigns it to the GuestName field.
func (o *ServiceRequest) SetGuestName(v string) {
	o.GuestName = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *ServiceRequest) GetRoom() string {
	if o == nil || IsNil(o.Room) {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetRoomOk() (*string, bool) {
	if o == nil || IsNil(o.Room) {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *ServiceRequest) HasRoom() bool {
	if o != nil && !IsNil(o.Room) {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *ServiceRequest) SetRoom(v string) {
	o.Room = &v
}

// GetOpenDate returns the OpenDate field value if set, zero value otherwise.
func (o *ServiceRequest) GetOpenDate() time.Time {
	if o == nil || IsNil(o.OpenDate) {
		var ret time.Time
		return ret
	}
	return *o.OpenDate
}

// GetOpenDateOk returns a tuple with the OpenDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetOpenDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpenDate) {
		return nil, false
	}
	return o.OpenDate, true
}

// HasOpenDate returns a boolean if a field has been set.
func (o *ServiceRequest) HasOpenDate() bool {
	if o != nil && !IsNil(o.OpenDate) {
		return true
	}

	return false
}

// SetOpenDate gets a reference to the given time.Time and assigns it to the OpenDate field.
func (o *ServiceRequest) SetOpenDate(v time.Time) {
	o.OpenDate = &v
}

// GetGuaranteedBy returns the GuaranteedBy field value if set, zero value otherwise.
func (o *ServiceRequest) GetGuaranteedBy() string {
	if o == nil || IsNil(o.GuaranteedBy) {
		var ret string
		return ret
	}
	return *o.GuaranteedBy
}

// GetGuaranteedByOk returns a tuple with the GuaranteedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetGuaranteedByOk() (*string, bool) {
	if o == nil || IsNil(o.GuaranteedBy) {
		return nil, false
	}
	return o.GuaranteedBy, true
}

// HasGuaranteedBy returns a boolean if a field has been set.
func (o *ServiceRequest) HasGuaranteedBy() bool {
	if o != nil && !IsNil(o.GuaranteedBy) {
		return true
	}

	return false
}

// SetGuaranteedBy gets a reference to the given string and assigns it to the GuaranteedBy field.
func (o *ServiceRequest) SetGuaranteedBy(v string) {
	o.GuaranteedBy = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ServiceRequest) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ServiceRequest) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ServiceRequest) SetComment(v string) {
	o.Comment = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ServiceRequest) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ServiceRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ServiceRequest) SetAction(v string) {
	o.Action = &v
}

// GetGuestContactMethod returns the GuestContactMethod field value if set, zero value otherwise.
func (o *ServiceRequest) GetGuestContactMethod() string {
	if o == nil || IsNil(o.GuestContactMethod) {
		var ret string
		return ret
	}
	return *o.GuestContactMethod
}

// GetGuestContactMethodOk returns a tuple with the GuestContactMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetGuestContactMethodOk() (*string, bool) {
	if o == nil || IsNil(o.GuestContactMethod) {
		return nil, false
	}
	return o.GuestContactMethod, true
}

// HasGuestContactMethod returns a boolean if a field has been set.
func (o *ServiceRequest) HasGuestContactMethod() bool {
	if o != nil && !IsNil(o.GuestContactMethod) {
		return true
	}

	return false
}

// SetGuestContactMethod gets a reference to the given string and assigns it to the GuestContactMethod field.
func (o *ServiceRequest) SetGuestContactMethod(v string) {
	o.GuestContactMethod = &v
}

// GetCompletionDate returns the CompletionDate field value if set, zero value otherwise.
func (o *ServiceRequest) GetCompletionDate() time.Time {
	if o == nil || IsNil(o.CompletionDate) {
		var ret time.Time
		return ret
	}
	return *o.CompletionDate
}

// GetCompletionDateOk returns a tuple with the CompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetCompletionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletionDate) {
		return nil, false
	}
	return o.CompletionDate, true
}

// HasCompletionDate returns a boolean if a field has been set.
func (o *ServiceRequest) HasCompletionDate() bool {
	if o != nil && !IsNil(o.CompletionDate) {
		return true
	}

	return false
}

// SetCompletionDate gets a reference to the given time.Time and assigns it to the CompletionDate field.
func (o *ServiceRequest) SetCompletionDate(v time.Time) {
	o.CompletionDate = &v
}

// GetClosedDate returns the ClosedDate field value if set, zero value otherwise.
func (o *ServiceRequest) GetClosedDate() time.Time {
	if o == nil || IsNil(o.ClosedDate) {
		var ret time.Time
		return ret
	}
	return *o.ClosedDate
}

// GetClosedDateOk returns a tuple with the ClosedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetClosedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClosedDate) {
		return nil, false
	}
	return o.ClosedDate, true
}

// HasClosedDate returns a boolean if a field has been set.
func (o *ServiceRequest) HasClosedDate() bool {
	if o != nil && !IsNil(o.ClosedDate) {
		return true
	}

	return false
}

// SetClosedDate gets a reference to the given time.Time and assigns it to the ClosedDate field.
func (o *ServiceRequest) SetClosedDate(v time.Time) {
	o.ClosedDate = &v
}

// GetClosedBy returns the ClosedBy field value if set, zero value otherwise.
func (o *ServiceRequest) GetClosedBy() string {
	if o == nil || IsNil(o.ClosedBy) {
		var ret string
		return ret
	}
	return *o.ClosedBy
}

// GetClosedByOk returns a tuple with the ClosedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetClosedByOk() (*string, bool) {
	if o == nil || IsNil(o.ClosedBy) {
		return nil, false
	}
	return o.ClosedBy, true
}

// HasClosedBy returns a boolean if a field has been set.
func (o *ServiceRequest) HasClosedBy() bool {
	if o != nil && !IsNil(o.ClosedBy) {
		return true
	}

	return false
}

// SetClosedBy gets a reference to the given string and assigns it to the ClosedBy field.
func (o *ServiceRequest) SetClosedBy(v string) {
	o.ClosedBy = &v
}

// GetCloseDescription returns the CloseDescription field value if set, zero value otherwise.
func (o *ServiceRequest) GetCloseDescription() string {
	if o == nil || IsNil(o.CloseDescription) {
		var ret string
		return ret
	}
	return *o.CloseDescription
}

// GetCloseDescriptionOk returns a tuple with the CloseDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetCloseDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.CloseDescription) {
		return nil, false
	}
	return o.CloseDescription, true
}

// HasCloseDescription returns a boolean if a field has been set.
func (o *ServiceRequest) HasCloseDescription() bool {
	if o != nil && !IsNil(o.CloseDescription) {
		return true
	}

	return false
}

// SetCloseDescription gets a reference to the given string and assigns it to the CloseDescription field.
func (o *ServiceRequest) SetCloseDescription(v string) {
	o.CloseDescription = &v
}

func (o ServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceRequestId) {
		toSerialize["serviceRequestId"] = o.ServiceRequestId
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.ReservationIdList) {
		toSerialize["reservationIdList"] = o.ReservationIdList
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.GuestName) {
		toSerialize["guestName"] = o.GuestName
	}
	if !IsNil(o.Room) {
		toSerialize["room"] = o.Room
	}
	if !IsNil(o.OpenDate) {
		toSerialize["openDate"] = o.OpenDate
	}
	if !IsNil(o.GuaranteedBy) {
		toSerialize["guaranteedBy"] = o.GuaranteedBy
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.GuestContactMethod) {
		toSerialize["guestContactMethod"] = o.GuestContactMethod
	}
	if !IsNil(o.CompletionDate) {
		toSerialize["completionDate"] = o.CompletionDate
	}
	if !IsNil(o.ClosedDate) {
		toSerialize["closedDate"] = o.ClosedDate
	}
	if !IsNil(o.ClosedBy) {
		toSerialize["closedBy"] = o.ClosedBy
	}
	if !IsNil(o.CloseDescription) {
		toSerialize["closeDescription"] = o.CloseDescription
	}
	return toSerialize, nil
}

type NullableServiceRequest struct {
	value *ServiceRequest
	isSet bool
}

func (v NullableServiceRequest) Get() *ServiceRequest {
	return v.value
}

func (v *NullableServiceRequest) Set(val *ServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRequest(val *ServiceRequest) *NullableServiceRequest {
	return &NullableServiceRequest{value: val, isSet: true}
}

func (v NullableServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


