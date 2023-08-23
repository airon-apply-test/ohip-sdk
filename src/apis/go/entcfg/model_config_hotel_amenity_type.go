/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ConfigHotelAmenityType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigHotelAmenityType{}

// ConfigHotelAmenityType Base details used for amenities.
type ConfigHotelAmenityType struct {
	// The description about amenity of the hotel.
	Description *string `json:"description,omitempty"`
	// The comments about amenity of the hotel.
	Comments *string `json:"comments,omitempty"`
	// Specifies the feature code (aka amenity code).
	FeatureCode *string `json:"featureCode,omitempty"`
	// Display Order sequence.
	OrderSequence *float32 `json:"orderSequence,omitempty"`
	AmenityType *AmenityTypeType `json:"amenityType,omitempty"`
	// The date the amenity is scheduled to become active.
	BeginDate *string `json:"beginDate,omitempty"`
	// The date the amenity is scheduled to become inactive.
	EndDate *string `json:"endDate,omitempty"`
	// The new amenity code which is used in the change method.
	NewAmenityCode *string `json:"newAmenityCode,omitempty"`
	// The new date the amenity is scheduled to become active.
	NewBeginDate *string `json:"newBeginDate,omitempty"`
	// The hours of operation of the amenity in the hotel.
	Hours *string `json:"hours,omitempty"`
	// The price range of the amenity in the hotel.
	PriceRange *string `json:"priceRange,omitempty"`
	// Specifies the hotel code for which the amenity is specified.
	HotelId *string `json:"hotelId,omitempty"`
}

// NewConfigHotelAmenityType instantiates a new ConfigHotelAmenityType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigHotelAmenityType() *ConfigHotelAmenityType {
	this := ConfigHotelAmenityType{}
	return &this
}

// NewConfigHotelAmenityTypeWithDefaults instantiates a new ConfigHotelAmenityType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigHotelAmenityTypeWithDefaults() *ConfigHotelAmenityType {
	this := ConfigHotelAmenityType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigHotelAmenityType) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ConfigHotelAmenityType) SetComments(v string) {
	o.Comments = &v
}

// GetFeatureCode returns the FeatureCode field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetFeatureCode() string {
	if o == nil || IsNil(o.FeatureCode) {
		var ret string
		return ret
	}
	return *o.FeatureCode
}

// GetFeatureCodeOk returns a tuple with the FeatureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetFeatureCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureCode) {
		return nil, false
	}
	return o.FeatureCode, true
}

// HasFeatureCode returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasFeatureCode() bool {
	if o != nil && !IsNil(o.FeatureCode) {
		return true
	}

	return false
}

// SetFeatureCode gets a reference to the given string and assigns it to the FeatureCode field.
func (o *ConfigHotelAmenityType) SetFeatureCode(v string) {
	o.FeatureCode = &v
}

// GetOrderSequence returns the OrderSequence field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetOrderSequence() float32 {
	if o == nil || IsNil(o.OrderSequence) {
		var ret float32
		return ret
	}
	return *o.OrderSequence
}

// GetOrderSequenceOk returns a tuple with the OrderSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetOrderSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderSequence) {
		return nil, false
	}
	return o.OrderSequence, true
}

// HasOrderSequence returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasOrderSequence() bool {
	if o != nil && !IsNil(o.OrderSequence) {
		return true
	}

	return false
}

// SetOrderSequence gets a reference to the given float32 and assigns it to the OrderSequence field.
func (o *ConfigHotelAmenityType) SetOrderSequence(v float32) {
	o.OrderSequence = &v
}

// GetAmenityType returns the AmenityType field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetAmenityType() AmenityTypeType {
	if o == nil || IsNil(o.AmenityType) {
		var ret AmenityTypeType
		return ret
	}
	return *o.AmenityType
}

// GetAmenityTypeOk returns a tuple with the AmenityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetAmenityTypeOk() (*AmenityTypeType, bool) {
	if o == nil || IsNil(o.AmenityType) {
		return nil, false
	}
	return o.AmenityType, true
}

// HasAmenityType returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasAmenityType() bool {
	if o != nil && !IsNil(o.AmenityType) {
		return true
	}

	return false
}

// SetAmenityType gets a reference to the given AmenityTypeType and assigns it to the AmenityType field.
func (o *ConfigHotelAmenityType) SetAmenityType(v AmenityTypeType) {
	o.AmenityType = &v
}

// GetBeginDate returns the BeginDate field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetBeginDate() string {
	if o == nil || IsNil(o.BeginDate) {
		var ret string
		return ret
	}
	return *o.BeginDate
}

// GetBeginDateOk returns a tuple with the BeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetBeginDateOk() (*string, bool) {
	if o == nil || IsNil(o.BeginDate) {
		return nil, false
	}
	return o.BeginDate, true
}

// HasBeginDate returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasBeginDate() bool {
	if o != nil && !IsNil(o.BeginDate) {
		return true
	}

	return false
}

// SetBeginDate gets a reference to the given string and assigns it to the BeginDate field.
func (o *ConfigHotelAmenityType) SetBeginDate(v string) {
	o.BeginDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ConfigHotelAmenityType) SetEndDate(v string) {
	o.EndDate = &v
}

// GetNewAmenityCode returns the NewAmenityCode field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetNewAmenityCode() string {
	if o == nil || IsNil(o.NewAmenityCode) {
		var ret string
		return ret
	}
	return *o.NewAmenityCode
}

// GetNewAmenityCodeOk returns a tuple with the NewAmenityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetNewAmenityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.NewAmenityCode) {
		return nil, false
	}
	return o.NewAmenityCode, true
}

// HasNewAmenityCode returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasNewAmenityCode() bool {
	if o != nil && !IsNil(o.NewAmenityCode) {
		return true
	}

	return false
}

// SetNewAmenityCode gets a reference to the given string and assigns it to the NewAmenityCode field.
func (o *ConfigHotelAmenityType) SetNewAmenityCode(v string) {
	o.NewAmenityCode = &v
}

// GetNewBeginDate returns the NewBeginDate field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetNewBeginDate() string {
	if o == nil || IsNil(o.NewBeginDate) {
		var ret string
		return ret
	}
	return *o.NewBeginDate
}

// GetNewBeginDateOk returns a tuple with the NewBeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetNewBeginDateOk() (*string, bool) {
	if o == nil || IsNil(o.NewBeginDate) {
		return nil, false
	}
	return o.NewBeginDate, true
}

// HasNewBeginDate returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasNewBeginDate() bool {
	if o != nil && !IsNil(o.NewBeginDate) {
		return true
	}

	return false
}

// SetNewBeginDate gets a reference to the given string and assigns it to the NewBeginDate field.
func (o *ConfigHotelAmenityType) SetNewBeginDate(v string) {
	o.NewBeginDate = &v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetHours() string {
	if o == nil || IsNil(o.Hours) {
		var ret string
		return ret
	}
	return *o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetHoursOk() (*string, bool) {
	if o == nil || IsNil(o.Hours) {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasHours() bool {
	if o != nil && !IsNil(o.Hours) {
		return true
	}

	return false
}

// SetHours gets a reference to the given string and assigns it to the Hours field.
func (o *ConfigHotelAmenityType) SetHours(v string) {
	o.Hours = &v
}

// GetPriceRange returns the PriceRange field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetPriceRange() string {
	if o == nil || IsNil(o.PriceRange) {
		var ret string
		return ret
	}
	return *o.PriceRange
}

// GetPriceRangeOk returns a tuple with the PriceRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetPriceRangeOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRange) {
		return nil, false
	}
	return o.PriceRange, true
}

// HasPriceRange returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasPriceRange() bool {
	if o != nil && !IsNil(o.PriceRange) {
		return true
	}

	return false
}

// SetPriceRange gets a reference to the given string and assigns it to the PriceRange field.
func (o *ConfigHotelAmenityType) SetPriceRange(v string) {
	o.PriceRange = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ConfigHotelAmenityType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHotelAmenityType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ConfigHotelAmenityType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ConfigHotelAmenityType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o ConfigHotelAmenityType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigHotelAmenityType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.FeatureCode) {
		toSerialize["featureCode"] = o.FeatureCode
	}
	if !IsNil(o.OrderSequence) {
		toSerialize["orderSequence"] = o.OrderSequence
	}
	if !IsNil(o.AmenityType) {
		toSerialize["amenityType"] = o.AmenityType
	}
	if !IsNil(o.BeginDate) {
		toSerialize["beginDate"] = o.BeginDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.NewAmenityCode) {
		toSerialize["newAmenityCode"] = o.NewAmenityCode
	}
	if !IsNil(o.NewBeginDate) {
		toSerialize["newBeginDate"] = o.NewBeginDate
	}
	if !IsNil(o.Hours) {
		toSerialize["hours"] = o.Hours
	}
	if !IsNil(o.PriceRange) {
		toSerialize["priceRange"] = o.PriceRange
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableConfigHotelAmenityType struct {
	value *ConfigHotelAmenityType
	isSet bool
}

func (v NullableConfigHotelAmenityType) Get() *ConfigHotelAmenityType {
	return v.value
}

func (v *NullableConfigHotelAmenityType) Set(val *ConfigHotelAmenityType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigHotelAmenityType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigHotelAmenityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigHotelAmenityType(val *ConfigHotelAmenityType) *NullableConfigHotelAmenityType {
	return &NullableConfigHotelAmenityType{value: val, isSet: true}
}

func (v NullableConfigHotelAmenityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigHotelAmenityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


