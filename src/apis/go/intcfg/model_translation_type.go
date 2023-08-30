/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
)

// checks if the TranslationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranslationType{}

// TranslationType struct for TranslationType
type TranslationType struct {
	// Guest number length of a translation setup.
	GuestNoLength *string `json:"guestNoLength,omitempty"`
	// Guest message id length of a translation setup.
	GuestMessageIdLength *string `json:"guestMessageIdLength,omitempty"`
	// Group number length of a translation setup.
	GroupNoLength *string `json:"groupNoLength,omitempty"`
	// Default charge of a translation setup.
	DefaultCharge *int32 `json:"defaultCharge,omitempty"`
	// Translation configuration of a hotel interface.
	Configuation []InterfaceControlCfgType `json:"configuation,omitempty"`
	// Translation article setup of a hotel interface.
	Articles []InterfaceControlArticleType `json:"articles,omitempty"`
	// Translation specifications setup of a hotel interface.
	Specifications []InterfaceControlSpecType `json:"specifications,omitempty"`
	// Translation languages setup of a hotel interface.
	Languages []InterfaceControlLangType `json:"languages,omitempty"`
}

// NewTranslationType instantiates a new TranslationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationType() *TranslationType {
	this := TranslationType{}
	return &this
}

// NewTranslationTypeWithDefaults instantiates a new TranslationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationTypeWithDefaults() *TranslationType {
	this := TranslationType{}
	return &this
}

// GetGuestNoLength returns the GuestNoLength field value if set, zero value otherwise.
func (o *TranslationType) GetGuestNoLength() string {
	if o == nil || IsNil(o.GuestNoLength) {
		var ret string
		return ret
	}
	return *o.GuestNoLength
}

// GetGuestNoLengthOk returns a tuple with the GuestNoLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetGuestNoLengthOk() (*string, bool) {
	if o == nil || IsNil(o.GuestNoLength) {
		return nil, false
	}
	return o.GuestNoLength, true
}

// HasGuestNoLength returns a boolean if a field has been set.
func (o *TranslationType) HasGuestNoLength() bool {
	if o != nil && !IsNil(o.GuestNoLength) {
		return true
	}

	return false
}

// SetGuestNoLength gets a reference to the given string and assigns it to the GuestNoLength field.
func (o *TranslationType) SetGuestNoLength(v string) {
	o.GuestNoLength = &v
}

// GetGuestMessageIdLength returns the GuestMessageIdLength field value if set, zero value otherwise.
func (o *TranslationType) GetGuestMessageIdLength() string {
	if o == nil || IsNil(o.GuestMessageIdLength) {
		var ret string
		return ret
	}
	return *o.GuestMessageIdLength
}

// GetGuestMessageIdLengthOk returns a tuple with the GuestMessageIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetGuestMessageIdLengthOk() (*string, bool) {
	if o == nil || IsNil(o.GuestMessageIdLength) {
		return nil, false
	}
	return o.GuestMessageIdLength, true
}

// HasGuestMessageIdLength returns a boolean if a field has been set.
func (o *TranslationType) HasGuestMessageIdLength() bool {
	if o != nil && !IsNil(o.GuestMessageIdLength) {
		return true
	}

	return false
}

// SetGuestMessageIdLength gets a reference to the given string and assigns it to the GuestMessageIdLength field.
func (o *TranslationType) SetGuestMessageIdLength(v string) {
	o.GuestMessageIdLength = &v
}

// GetGroupNoLength returns the GroupNoLength field value if set, zero value otherwise.
func (o *TranslationType) GetGroupNoLength() string {
	if o == nil || IsNil(o.GroupNoLength) {
		var ret string
		return ret
	}
	return *o.GroupNoLength
}

// GetGroupNoLengthOk returns a tuple with the GroupNoLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetGroupNoLengthOk() (*string, bool) {
	if o == nil || IsNil(o.GroupNoLength) {
		return nil, false
	}
	return o.GroupNoLength, true
}

// HasGroupNoLength returns a boolean if a field has been set.
func (o *TranslationType) HasGroupNoLength() bool {
	if o != nil && !IsNil(o.GroupNoLength) {
		return true
	}

	return false
}

// SetGroupNoLength gets a reference to the given string and assigns it to the GroupNoLength field.
func (o *TranslationType) SetGroupNoLength(v string) {
	o.GroupNoLength = &v
}

// GetDefaultCharge returns the DefaultCharge field value if set, zero value otherwise.
func (o *TranslationType) GetDefaultCharge() int32 {
	if o == nil || IsNil(o.DefaultCharge) {
		var ret int32
		return ret
	}
	return *o.DefaultCharge
}

// GetDefaultChargeOk returns a tuple with the DefaultCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetDefaultChargeOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultCharge) {
		return nil, false
	}
	return o.DefaultCharge, true
}

// HasDefaultCharge returns a boolean if a field has been set.
func (o *TranslationType) HasDefaultCharge() bool {
	if o != nil && !IsNil(o.DefaultCharge) {
		return true
	}

	return false
}

// SetDefaultCharge gets a reference to the given int32 and assigns it to the DefaultCharge field.
func (o *TranslationType) SetDefaultCharge(v int32) {
	o.DefaultCharge = &v
}

// GetConfiguation returns the Configuation field value if set, zero value otherwise.
func (o *TranslationType) GetConfiguation() []InterfaceControlCfgType {
	if o == nil || IsNil(o.Configuation) {
		var ret []InterfaceControlCfgType
		return ret
	}
	return o.Configuation
}

// GetConfiguationOk returns a tuple with the Configuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetConfiguationOk() ([]InterfaceControlCfgType, bool) {
	if o == nil || IsNil(o.Configuation) {
		return nil, false
	}
	return o.Configuation, true
}

// HasConfiguation returns a boolean if a field has been set.
func (o *TranslationType) HasConfiguation() bool {
	if o != nil && !IsNil(o.Configuation) {
		return true
	}

	return false
}

// SetConfiguation gets a reference to the given []InterfaceControlCfgType and assigns it to the Configuation field.
func (o *TranslationType) SetConfiguation(v []InterfaceControlCfgType) {
	o.Configuation = v
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *TranslationType) GetArticles() []InterfaceControlArticleType {
	if o == nil || IsNil(o.Articles) {
		var ret []InterfaceControlArticleType
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetArticlesOk() ([]InterfaceControlArticleType, bool) {
	if o == nil || IsNil(o.Articles) {
		return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *TranslationType) HasArticles() bool {
	if o != nil && !IsNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []InterfaceControlArticleType and assigns it to the Articles field.
func (o *TranslationType) SetArticles(v []InterfaceControlArticleType) {
	o.Articles = v
}

// GetSpecifications returns the Specifications field value if set, zero value otherwise.
func (o *TranslationType) GetSpecifications() []InterfaceControlSpecType {
	if o == nil || IsNil(o.Specifications) {
		var ret []InterfaceControlSpecType
		return ret
	}
	return o.Specifications
}

// GetSpecificationsOk returns a tuple with the Specifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetSpecificationsOk() ([]InterfaceControlSpecType, bool) {
	if o == nil || IsNil(o.Specifications) {
		return nil, false
	}
	return o.Specifications, true
}

// HasSpecifications returns a boolean if a field has been set.
func (o *TranslationType) HasSpecifications() bool {
	if o != nil && !IsNil(o.Specifications) {
		return true
	}

	return false
}

// SetSpecifications gets a reference to the given []InterfaceControlSpecType and assigns it to the Specifications field.
func (o *TranslationType) SetSpecifications(v []InterfaceControlSpecType) {
	o.Specifications = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *TranslationType) GetLanguages() []InterfaceControlLangType {
	if o == nil || IsNil(o.Languages) {
		var ret []InterfaceControlLangType
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationType) GetLanguagesOk() ([]InterfaceControlLangType, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *TranslationType) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []InterfaceControlLangType and assigns it to the Languages field.
func (o *TranslationType) SetLanguages(v []InterfaceControlLangType) {
	o.Languages = v
}

func (o TranslationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranslationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestNoLength) {
		toSerialize["guestNoLength"] = o.GuestNoLength
	}
	if !IsNil(o.GuestMessageIdLength) {
		toSerialize["guestMessageIdLength"] = o.GuestMessageIdLength
	}
	if !IsNil(o.GroupNoLength) {
		toSerialize["groupNoLength"] = o.GroupNoLength
	}
	if !IsNil(o.DefaultCharge) {
		toSerialize["defaultCharge"] = o.DefaultCharge
	}
	if !IsNil(o.Configuation) {
		toSerialize["configuation"] = o.Configuation
	}
	if !IsNil(o.Articles) {
		toSerialize["articles"] = o.Articles
	}
	if !IsNil(o.Specifications) {
		toSerialize["specifications"] = o.Specifications
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	return toSerialize, nil
}

type NullableTranslationType struct {
	value *TranslationType
	isSet bool
}

func (v NullableTranslationType) Get() *TranslationType {
	return v.value
}

func (v *NullableTranslationType) Set(val *TranslationType) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationType) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationType(val *TranslationType) *NullableTranslationType {
	return &NullableTranslationType{value: val, isSet: true}
}

func (v NullableTranslationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


