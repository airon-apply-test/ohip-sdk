/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
	"time"
)

// checks if the CommentType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentType{}

// CommentType An indication of a new paragraph for a sub-section of a formatted text message.
type CommentType struct {
	Text *FormattedTextTextType `json:"text,omitempty"`
	// An image for this paragraph.
	Image *string `json:"image,omitempty"`
	// A URL for this paragraph.
	Url *string `json:"url,omitempty"`
	// Specifies Comment's Title.
	CommentTitle *string `json:"commentTitle,omitempty"`
	// Notification Location associated with the Note.
	NotificationLocation *string `json:"notificationLocation,omitempty"`
	// Specifies type of the comment.
	Type *string `json:"type,omitempty"`
	// Comment type Description.
	TypeDescription *string `json:"typeDescription,omitempty"`
	// When true, the comment may not be shown to the consumer. When false, the comment may be shown to the consumer.
	Internal *bool `json:"internal,omitempty"`
	// When true, the comment may be confidential.
	Confidential *bool `json:"confidential,omitempty"`
	// When true, the note internal could be modified.
	OverrideInternal *bool `json:"overrideInternal,omitempty"`
	// When true, the note title will be populated from the note type description and couldn't be modified.
	ProtectDescription *bool `json:"protectDescription,omitempty"`
	// If specified comment belongs to the Hotel, otherwise it is a global comment.
	HotelId *string `json:"hotelId,omitempty"`
	// Specifies type of action described in the comments.
	ActionType *string `json:"actionType,omitempty"`
	// Indicates at which date an action described in the comment must be taken.
	ActionDate *string `json:"actionDate,omitempty"`
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

// NewCommentType instantiates a new CommentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentType() *CommentType {
	this := CommentType{}
	return &this
}

// NewCommentTypeWithDefaults instantiates a new CommentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentTypeWithDefaults() *CommentType {
	this := CommentType{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CommentType) GetText() FormattedTextTextType {
	if o == nil || IsNil(o.Text) {
		var ret FormattedTextTextType
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetTextOk() (*FormattedTextTextType, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CommentType) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given FormattedTextTextType and assigns it to the Text field.
func (o *CommentType) SetText(v FormattedTextTextType) {
	o.Text = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CommentType) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CommentType) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CommentType) SetImage(v string) {
	o.Image = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CommentType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CommentType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CommentType) SetUrl(v string) {
	o.Url = &v
}

// GetCommentTitle returns the CommentTitle field value if set, zero value otherwise.
func (o *CommentType) GetCommentTitle() string {
	if o == nil || IsNil(o.CommentTitle) {
		var ret string
		return ret
	}
	return *o.CommentTitle
}

// GetCommentTitleOk returns a tuple with the CommentTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetCommentTitleOk() (*string, bool) {
	if o == nil || IsNil(o.CommentTitle) {
		return nil, false
	}
	return o.CommentTitle, true
}

// HasCommentTitle returns a boolean if a field has been set.
func (o *CommentType) HasCommentTitle() bool {
	if o != nil && !IsNil(o.CommentTitle) {
		return true
	}

	return false
}

// SetCommentTitle gets a reference to the given string and assigns it to the CommentTitle field.
func (o *CommentType) SetCommentTitle(v string) {
	o.CommentTitle = &v
}

// GetNotificationLocation returns the NotificationLocation field value if set, zero value otherwise.
func (o *CommentType) GetNotificationLocation() string {
	if o == nil || IsNil(o.NotificationLocation) {
		var ret string
		return ret
	}
	return *o.NotificationLocation
}

// GetNotificationLocationOk returns a tuple with the NotificationLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetNotificationLocationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationLocation) {
		return nil, false
	}
	return o.NotificationLocation, true
}

// HasNotificationLocation returns a boolean if a field has been set.
func (o *CommentType) HasNotificationLocation() bool {
	if o != nil && !IsNil(o.NotificationLocation) {
		return true
	}

	return false
}

// SetNotificationLocation gets a reference to the given string and assigns it to the NotificationLocation field.
func (o *CommentType) SetNotificationLocation(v string) {
	o.NotificationLocation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CommentType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CommentType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CommentType) SetType(v string) {
	o.Type = &v
}

// GetTypeDescription returns the TypeDescription field value if set, zero value otherwise.
func (o *CommentType) GetTypeDescription() string {
	if o == nil || IsNil(o.TypeDescription) {
		var ret string
		return ret
	}
	return *o.TypeDescription
}

// GetTypeDescriptionOk returns a tuple with the TypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TypeDescription) {
		return nil, false
	}
	return o.TypeDescription, true
}

// HasTypeDescription returns a boolean if a field has been set.
func (o *CommentType) HasTypeDescription() bool {
	if o != nil && !IsNil(o.TypeDescription) {
		return true
	}

	return false
}

// SetTypeDescription gets a reference to the given string and assigns it to the TypeDescription field.
func (o *CommentType) SetTypeDescription(v string) {
	o.TypeDescription = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *CommentType) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *CommentType) HasInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *CommentType) SetInternal(v bool) {
	o.Internal = &v
}

// GetConfidential returns the Confidential field value if set, zero value otherwise.
func (o *CommentType) GetConfidential() bool {
	if o == nil || IsNil(o.Confidential) {
		var ret bool
		return ret
	}
	return *o.Confidential
}

// GetConfidentialOk returns a tuple with the Confidential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetConfidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Confidential) {
		return nil, false
	}
	return o.Confidential, true
}

// HasConfidential returns a boolean if a field has been set.
func (o *CommentType) HasConfidential() bool {
	if o != nil && !IsNil(o.Confidential) {
		return true
	}

	return false
}

// SetConfidential gets a reference to the given bool and assigns it to the Confidential field.
func (o *CommentType) SetConfidential(v bool) {
	o.Confidential = &v
}

// GetOverrideInternal returns the OverrideInternal field value if set, zero value otherwise.
func (o *CommentType) GetOverrideInternal() bool {
	if o == nil || IsNil(o.OverrideInternal) {
		var ret bool
		return ret
	}
	return *o.OverrideInternal
}

// GetOverrideInternalOk returns a tuple with the OverrideInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetOverrideInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.OverrideInternal) {
		return nil, false
	}
	return o.OverrideInternal, true
}

// HasOverrideInternal returns a boolean if a field has been set.
func (o *CommentType) HasOverrideInternal() bool {
	if o != nil && !IsNil(o.OverrideInternal) {
		return true
	}

	return false
}

// SetOverrideInternal gets a reference to the given bool and assigns it to the OverrideInternal field.
func (o *CommentType) SetOverrideInternal(v bool) {
	o.OverrideInternal = &v
}

// GetProtectDescription returns the ProtectDescription field value if set, zero value otherwise.
func (o *CommentType) GetProtectDescription() bool {
	if o == nil || IsNil(o.ProtectDescription) {
		var ret bool
		return ret
	}
	return *o.ProtectDescription
}

// GetProtectDescriptionOk returns a tuple with the ProtectDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetProtectDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectDescription) {
		return nil, false
	}
	return o.ProtectDescription, true
}

// HasProtectDescription returns a boolean if a field has been set.
func (o *CommentType) HasProtectDescription() bool {
	if o != nil && !IsNil(o.ProtectDescription) {
		return true
	}

	return false
}

// SetProtectDescription gets a reference to the given bool and assigns it to the ProtectDescription field.
func (o *CommentType) SetProtectDescription(v bool) {
	o.ProtectDescription = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *CommentType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *CommentType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *CommentType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *CommentType) GetActionType() string {
	if o == nil || IsNil(o.ActionType) {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetActionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *CommentType) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *CommentType) SetActionType(v string) {
	o.ActionType = &v
}

// GetActionDate returns the ActionDate field value if set, zero value otherwise.
func (o *CommentType) GetActionDate() string {
	if o == nil || IsNil(o.ActionDate) {
		var ret string
		return ret
	}
	return *o.ActionDate
}

// GetActionDateOk returns a tuple with the ActionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetActionDateOk() (*string, bool) {
	if o == nil || IsNil(o.ActionDate) {
		return nil, false
	}
	return o.ActionDate, true
}

// HasActionDate returns a boolean if a field has been set.
func (o *CommentType) HasActionDate() bool {
	if o != nil && !IsNil(o.ActionDate) {
		return true
	}

	return false
}

// SetActionDate gets a reference to the given string and assigns it to the ActionDate field.
func (o *CommentType) SetActionDate(v string) {
	o.ActionDate = &v
}

// GetCreateDateTime returns the CreateDateTime field value if set, zero value otherwise.
func (o *CommentType) GetCreateDateTime() time.Time {
	if o == nil || IsNil(o.CreateDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateDateTime
}

// GetCreateDateTimeOk returns a tuple with the CreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetCreateDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDateTime) {
		return nil, false
	}
	return o.CreateDateTime, true
}

// HasCreateDateTime returns a boolean if a field has been set.
func (o *CommentType) HasCreateDateTime() bool {
	if o != nil && !IsNil(o.CreateDateTime) {
		return true
	}

	return false
}

// SetCreateDateTime gets a reference to the given time.Time and assigns it to the CreateDateTime field.
func (o *CommentType) SetCreateDateTime(v time.Time) {
	o.CreateDateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *CommentType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *CommentType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *CommentType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModifyDateTime returns the LastModifyDateTime field value if set, zero value otherwise.
func (o *CommentType) GetLastModifyDateTime() time.Time {
	if o == nil || IsNil(o.LastModifyDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyDateTime
}

// GetLastModifyDateTimeOk returns a tuple with the LastModifyDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetLastModifyDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyDateTime) {
		return nil, false
	}
	return o.LastModifyDateTime, true
}

// HasLastModifyDateTime returns a boolean if a field has been set.
func (o *CommentType) HasLastModifyDateTime() bool {
	if o != nil && !IsNil(o.LastModifyDateTime) {
		return true
	}

	return false
}

// SetLastModifyDateTime gets a reference to the given time.Time and assigns it to the LastModifyDateTime field.
func (o *CommentType) SetLastModifyDateTime(v time.Time) {
	o.LastModifyDateTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *CommentType) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *CommentType) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *CommentType) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetPurgeDate returns the PurgeDate field value if set, zero value otherwise.
func (o *CommentType) GetPurgeDate() string {
	if o == nil || IsNil(o.PurgeDate) {
		var ret string
		return ret
	}
	return *o.PurgeDate
}

// GetPurgeDateOk returns a tuple with the PurgeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentType) GetPurgeDateOk() (*string, bool) {
	if o == nil || IsNil(o.PurgeDate) {
		return nil, false
	}
	return o.PurgeDate, true
}

// HasPurgeDate returns a boolean if a field has been set.
func (o *CommentType) HasPurgeDate() bool {
	if o != nil && !IsNil(o.PurgeDate) {
		return true
	}

	return false
}

// SetPurgeDate gets a reference to the given string and assigns it to the PurgeDate field.
func (o *CommentType) SetPurgeDate(v string) {
	o.PurgeDate = &v
}

func (o CommentType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CommentTitle) {
		toSerialize["commentTitle"] = o.CommentTitle
	}
	if !IsNil(o.NotificationLocation) {
		toSerialize["notificationLocation"] = o.NotificationLocation
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeDescription) {
		toSerialize["typeDescription"] = o.TypeDescription
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.Confidential) {
		toSerialize["confidential"] = o.Confidential
	}
	if !IsNil(o.OverrideInternal) {
		toSerialize["overrideInternal"] = o.OverrideInternal
	}
	if !IsNil(o.ProtectDescription) {
		toSerialize["protectDescription"] = o.ProtectDescription
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !IsNil(o.ActionDate) {
		toSerialize["actionDate"] = o.ActionDate
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

type NullableCommentType struct {
	value *CommentType
	isSet bool
}

func (v NullableCommentType) Get() *CommentType {
	return v.value
}

func (v *NullableCommentType) Set(val *CommentType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentType(val *CommentType) *NullableCommentType {
	return &NullableCommentType{value: val, isSet: true}
}

func (v NullableCommentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


