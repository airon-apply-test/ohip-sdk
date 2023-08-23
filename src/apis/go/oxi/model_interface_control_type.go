/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InterfaceControlType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceControlType{}

// InterfaceControlType Internal settings that can be of a certain valueType.
type InterfaceControlType struct {
	ShortDescription *string `json:"shortDescription,omitempty"`
	Description *string `json:"description,omitempty"`
	// A flag which indicate whether a wild card search should be made.
	WildCardMatch *bool `json:"wildCardMatch,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Type *ApplicationSettingTypeType `json:"type,omitempty"`
	// Each configuration item will come with a HotelCode which will help the configuration to specify what context the update has to be. Eg. _Global,ORS,'CRO', etc.
	HotelId *string `json:"hotelId,omitempty"`
	Sequence *float32 `json:"sequence,omitempty"`
	EditAllowed *bool `json:"editAllowed,omitempty"`
	Value *string `json:"value,omitempty"`
	Scope *string `json:"scope,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
	ConversionType *ApplicationSettingConversionType `json:"conversionType,omitempty"`
	// Indicator if the function does not count against the OPERA Control function count limit or not
	SubscriptionCountEligible *bool `json:"subscriptionCountEligible,omitempty"`
	Settings []BaseApplicationSettingType `json:"settings,omitempty"`
	LevelType *ApplicationSettingLevelType `json:"levelType,omitempty"`
	LevelCode *string `json:"levelCode,omitempty"`
	// Identifier of the Interface.
	InterfaceId *string `json:"interfaceId,omitempty"`
	// Flag indicating to use global inbound/outbound value for conversion code.
	UseGlobal *bool `json:"useGlobal,omitempty"`
	// V5 parameter group.
	OxiParameterGroup *string `json:"oxiParameterGroup,omitempty"`
}

// NewInterfaceControlType instantiates a new InterfaceControlType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceControlType() *InterfaceControlType {
	this := InterfaceControlType{}
	return &this
}

// NewInterfaceControlTypeWithDefaults instantiates a new InterfaceControlType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceControlTypeWithDefaults() *InterfaceControlType {
	this := InterfaceControlType{}
	return &this
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *InterfaceControlType) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *InterfaceControlType) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *InterfaceControlType) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterfaceControlType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceControlType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterfaceControlType) SetDescription(v string) {
	o.Description = &v
}

// GetWildCardMatch returns the WildCardMatch field value if set, zero value otherwise.
func (o *InterfaceControlType) GetWildCardMatch() bool {
	if o == nil || IsNil(o.WildCardMatch) {
		var ret bool
		return ret
	}
	return *o.WildCardMatch
}

// GetWildCardMatchOk returns a tuple with the WildCardMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetWildCardMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.WildCardMatch) {
		return nil, false
	}
	return o.WildCardMatch, true
}

// HasWildCardMatch returns a boolean if a field has been set.
func (o *InterfaceControlType) HasWildCardMatch() bool {
	if o != nil && !IsNil(o.WildCardMatch) {
		return true
	}

	return false
}

// SetWildCardMatch gets a reference to the given bool and assigns it to the WildCardMatch field.
func (o *InterfaceControlType) SetWildCardMatch(v bool) {
	o.WildCardMatch = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InterfaceControlType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InterfaceControlType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InterfaceControlType) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InterfaceControlType) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InterfaceControlType) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InterfaceControlType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InterfaceControlType) GetType() ApplicationSettingTypeType {
	if o == nil || IsNil(o.Type) {
		var ret ApplicationSettingTypeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetTypeOk() (*ApplicationSettingTypeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InterfaceControlType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ApplicationSettingTypeType and assigns it to the Type field.
func (o *InterfaceControlType) SetType(v ApplicationSettingTypeType) {
	o.Type = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *InterfaceControlType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *InterfaceControlType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *InterfaceControlType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *InterfaceControlType) GetSequence() float32 {
	if o == nil || IsNil(o.Sequence) {
		var ret float32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetSequenceOk() (*float32, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *InterfaceControlType) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given float32 and assigns it to the Sequence field.
func (o *InterfaceControlType) SetSequence(v float32) {
	o.Sequence = &v
}

// GetEditAllowed returns the EditAllowed field value if set, zero value otherwise.
func (o *InterfaceControlType) GetEditAllowed() bool {
	if o == nil || IsNil(o.EditAllowed) {
		var ret bool
		return ret
	}
	return *o.EditAllowed
}

// GetEditAllowedOk returns a tuple with the EditAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetEditAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.EditAllowed) {
		return nil, false
	}
	return o.EditAllowed, true
}

// HasEditAllowed returns a boolean if a field has been set.
func (o *InterfaceControlType) HasEditAllowed() bool {
	if o != nil && !IsNil(o.EditAllowed) {
		return true
	}

	return false
}

// SetEditAllowed gets a reference to the given bool and assigns it to the EditAllowed field.
func (o *InterfaceControlType) SetEditAllowed(v bool) {
	o.EditAllowed = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceControlType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceControlType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InterfaceControlType) SetValue(v string) {
	o.Value = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InterfaceControlType) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InterfaceControlType) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InterfaceControlType) SetScope(v string) {
	o.Scope = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *InterfaceControlType) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *InterfaceControlType) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *InterfaceControlType) SetValueType(v string) {
	o.ValueType = &v
}

// GetConversionType returns the ConversionType field value if set, zero value otherwise.
func (o *InterfaceControlType) GetConversionType() ApplicationSettingConversionType {
	if o == nil || IsNil(o.ConversionType) {
		var ret ApplicationSettingConversionType
		return ret
	}
	return *o.ConversionType
}

// GetConversionTypeOk returns a tuple with the ConversionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetConversionTypeOk() (*ApplicationSettingConversionType, bool) {
	if o == nil || IsNil(o.ConversionType) {
		return nil, false
	}
	return o.ConversionType, true
}

// HasConversionType returns a boolean if a field has been set.
func (o *InterfaceControlType) HasConversionType() bool {
	if o != nil && !IsNil(o.ConversionType) {
		return true
	}

	return false
}

// SetConversionType gets a reference to the given ApplicationSettingConversionType and assigns it to the ConversionType field.
func (o *InterfaceControlType) SetConversionType(v ApplicationSettingConversionType) {
	o.ConversionType = &v
}

// GetSubscriptionCountEligible returns the SubscriptionCountEligible field value if set, zero value otherwise.
func (o *InterfaceControlType) GetSubscriptionCountEligible() bool {
	if o == nil || IsNil(o.SubscriptionCountEligible) {
		var ret bool
		return ret
	}
	return *o.SubscriptionCountEligible
}

// GetSubscriptionCountEligibleOk returns a tuple with the SubscriptionCountEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetSubscriptionCountEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.SubscriptionCountEligible) {
		return nil, false
	}
	return o.SubscriptionCountEligible, true
}

// HasSubscriptionCountEligible returns a boolean if a field has been set.
func (o *InterfaceControlType) HasSubscriptionCountEligible() bool {
	if o != nil && !IsNil(o.SubscriptionCountEligible) {
		return true
	}

	return false
}

// SetSubscriptionCountEligible gets a reference to the given bool and assigns it to the SubscriptionCountEligible field.
func (o *InterfaceControlType) SetSubscriptionCountEligible(v bool) {
	o.SubscriptionCountEligible = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *InterfaceControlType) GetSettings() []BaseApplicationSettingType {
	if o == nil || IsNil(o.Settings) {
		var ret []BaseApplicationSettingType
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetSettingsOk() ([]BaseApplicationSettingType, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *InterfaceControlType) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []BaseApplicationSettingType and assigns it to the Settings field.
func (o *InterfaceControlType) SetSettings(v []BaseApplicationSettingType) {
	o.Settings = v
}

// GetLevelType returns the LevelType field value if set, zero value otherwise.
func (o *InterfaceControlType) GetLevelType() ApplicationSettingLevelType {
	if o == nil || IsNil(o.LevelType) {
		var ret ApplicationSettingLevelType
		return ret
	}
	return *o.LevelType
}

// GetLevelTypeOk returns a tuple with the LevelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetLevelTypeOk() (*ApplicationSettingLevelType, bool) {
	if o == nil || IsNil(o.LevelType) {
		return nil, false
	}
	return o.LevelType, true
}

// HasLevelType returns a boolean if a field has been set.
func (o *InterfaceControlType) HasLevelType() bool {
	if o != nil && !IsNil(o.LevelType) {
		return true
	}

	return false
}

// SetLevelType gets a reference to the given ApplicationSettingLevelType and assigns it to the LevelType field.
func (o *InterfaceControlType) SetLevelType(v ApplicationSettingLevelType) {
	o.LevelType = &v
}

// GetLevelCode returns the LevelCode field value if set, zero value otherwise.
func (o *InterfaceControlType) GetLevelCode() string {
	if o == nil || IsNil(o.LevelCode) {
		var ret string
		return ret
	}
	return *o.LevelCode
}

// GetLevelCodeOk returns a tuple with the LevelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetLevelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LevelCode) {
		return nil, false
	}
	return o.LevelCode, true
}

// HasLevelCode returns a boolean if a field has been set.
func (o *InterfaceControlType) HasLevelCode() bool {
	if o != nil && !IsNil(o.LevelCode) {
		return true
	}

	return false
}

// SetLevelCode gets a reference to the given string and assigns it to the LevelCode field.
func (o *InterfaceControlType) SetLevelCode(v string) {
	o.LevelCode = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *InterfaceControlType) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *InterfaceControlType) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *InterfaceControlType) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetUseGlobal returns the UseGlobal field value if set, zero value otherwise.
func (o *InterfaceControlType) GetUseGlobal() bool {
	if o == nil || IsNil(o.UseGlobal) {
		var ret bool
		return ret
	}
	return *o.UseGlobal
}

// GetUseGlobalOk returns a tuple with the UseGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetUseGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.UseGlobal) {
		return nil, false
	}
	return o.UseGlobal, true
}

// HasUseGlobal returns a boolean if a field has been set.
func (o *InterfaceControlType) HasUseGlobal() bool {
	if o != nil && !IsNil(o.UseGlobal) {
		return true
	}

	return false
}

// SetUseGlobal gets a reference to the given bool and assigns it to the UseGlobal field.
func (o *InterfaceControlType) SetUseGlobal(v bool) {
	o.UseGlobal = &v
}

// GetOxiParameterGroup returns the OxiParameterGroup field value if set, zero value otherwise.
func (o *InterfaceControlType) GetOxiParameterGroup() string {
	if o == nil || IsNil(o.OxiParameterGroup) {
		var ret string
		return ret
	}
	return *o.OxiParameterGroup
}

// GetOxiParameterGroupOk returns a tuple with the OxiParameterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceControlType) GetOxiParameterGroupOk() (*string, bool) {
	if o == nil || IsNil(o.OxiParameterGroup) {
		return nil, false
	}
	return o.OxiParameterGroup, true
}

// HasOxiParameterGroup returns a boolean if a field has been set.
func (o *InterfaceControlType) HasOxiParameterGroup() bool {
	if o != nil && !IsNil(o.OxiParameterGroup) {
		return true
	}

	return false
}

// SetOxiParameterGroup gets a reference to the given string and assigns it to the OxiParameterGroup field.
func (o *InterfaceControlType) SetOxiParameterGroup(v string) {
	o.OxiParameterGroup = &v
}

func (o InterfaceControlType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceControlType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.WildCardMatch) {
		toSerialize["wildCardMatch"] = o.WildCardMatch
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.EditAllowed) {
		toSerialize["editAllowed"] = o.EditAllowed
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.ConversionType) {
		toSerialize["conversionType"] = o.ConversionType
	}
	if !IsNil(o.SubscriptionCountEligible) {
		toSerialize["subscriptionCountEligible"] = o.SubscriptionCountEligible
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.LevelType) {
		toSerialize["levelType"] = o.LevelType
	}
	if !IsNil(o.LevelCode) {
		toSerialize["levelCode"] = o.LevelCode
	}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !IsNil(o.UseGlobal) {
		toSerialize["useGlobal"] = o.UseGlobal
	}
	if !IsNil(o.OxiParameterGroup) {
		toSerialize["oxiParameterGroup"] = o.OxiParameterGroup
	}
	return toSerialize, nil
}

type NullableInterfaceControlType struct {
	value *InterfaceControlType
	isSet bool
}

func (v NullableInterfaceControlType) Get() *InterfaceControlType {
	return v.value
}

func (v *NullableInterfaceControlType) Set(val *InterfaceControlType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceControlType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceControlType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceControlType(val *InterfaceControlType) *NullableInterfaceControlType {
	return &NullableInterfaceControlType{value: val, isSet: true}
}

func (v NullableInterfaceControlType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceControlType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


