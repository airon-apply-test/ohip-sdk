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

// checks if the RelationshipInfoSummaryType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelationshipInfoSummaryType{}

// RelationshipInfoSummaryType RelationshipInfoSummaryType contains information about the associations between and among individuals, companies, travel agents, groups, sources, and contact profiles.
type RelationshipInfoSummaryType struct {
	RelationshipProfile *RelationshipProfileSummaryType `json:"relationshipProfile,omitempty"`
	MasterAccountDetails *MasterAccountInfoType `json:"masterAccountDetails,omitempty"`
	// Relationship identifier.
	RelationshipID *string `json:"relationshipID,omitempty"`
	// Indicates the type of relationship the current profile(Source Profile) has with the related profile(Target Profile).
	SourceRelation *string `json:"sourceRelation,omitempty"`
	// Displays the description of relationship the current profile(Source Profile) has with the related profile(Target Profile).
	SourceRelationDescription *string `json:"sourceRelationDescription,omitempty"`
	// Displays the type of relationship the Related profile(Target Profile) has with the current profile(Source Profile).
	TargetRelation *string `json:"targetRelation,omitempty"`
	// Displays the description of the target relation(Target Profile).
	TargetRelationDescription *string `json:"targetRelationDescription,omitempty"`
}

// NewRelationshipInfoSummaryType instantiates a new RelationshipInfoSummaryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipInfoSummaryType() *RelationshipInfoSummaryType {
	this := RelationshipInfoSummaryType{}
	return &this
}

// NewRelationshipInfoSummaryTypeWithDefaults instantiates a new RelationshipInfoSummaryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipInfoSummaryTypeWithDefaults() *RelationshipInfoSummaryType {
	this := RelationshipInfoSummaryType{}
	return &this
}

// GetRelationshipProfile returns the RelationshipProfile field value if set, zero value otherwise.
func (o *RelationshipInfoSummaryType) GetRelationshipProfile() RelationshipProfileSummaryType {
	if o == nil || IsNil(o.RelationshipProfile) {
		var ret RelationshipProfileSummaryType
		return ret
	}
	return *o.RelationshipProfile
}

// GetRelationshipProfileOk returns a tuple with the RelationshipProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoSummaryType) GetRelationshipProfileOk() (*RelationshipProfileSummaryType, bool) {
	if o == nil || IsNil(o.RelationshipProfile) {
		return nil, false
	}
	return o.RelationshipProfile, true
}

// HasRelationshipProfile returns a boolean if a field has been set.
func (o *RelationshipInfoSummaryType) HasRelationshipProfile() bool {
	if o != nil && !IsNil(o.RelationshipProfile) {
		return true
	}

	return false
}

// SetRelationshipProfile gets a reference to the given RelationshipProfileSummaryType and assigns it to the RelationshipProfile field.
func (o *RelationshipInfoSummaryType) SetRelationshipProfile(v RelationshipProfileSummaryType) {
	o.RelationshipProfile = &v
}

// GetMasterAccountDetails returns the MasterAccountDetails field value if set, zero value otherwise.
func (o *RelationshipInfoSummaryType) GetMasterAccountDetails() MasterAccountInfoType {
	if o == nil || IsNil(o.MasterAccountDetails) {
		var ret MasterAccountInfoType
		return ret
	}
	return *o.MasterAccountDetails
}

// GetMasterAccountDetailsOk returns a tuple with the MasterAccountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoSummaryType) GetMasterAccountDetailsOk() (*MasterAccountInfoType, bool) {
	if o == nil || IsNil(o.MasterAccountDetails) {
		return nil, false
	}
	return o.MasterAccountDetails, true
}

// HasMasterAccountDetails returns a boolean if a field has been set.
func (o *RelationshipInfoSummaryType) HasMasterAccountDetails() bool {
	if o != nil && !IsNil(o.MasterAccountDetails) {
		return true
	}

	return false
}

// SetMasterAccountDetails gets a reference to the given MasterAccountInfoType and assigns it to the MasterAccountDetails field.
func (o *RelationshipInfoSummaryType) SetMasterAccountDetails(v MasterAccountInfoType) {
	o.MasterAccountDetails = &v
}

// GetRelationshipID returns the RelationshipID field value if set, zero value otherwise.
func (o *RelationshipInfoSummaryType) GetRelationshipID() string {
	if o == nil || IsNil(o.RelationshipID) {
		var ret string
		return ret
	}
	return *o.RelationshipID
}

// GetRelationshipIDOk returns a tuple with the RelationshipID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoSummaryType) GetRelationshipIDOk() (*string, bool) {
	if o == nil || IsNil(o.RelationshipID) {
		return nil, false
	}
	return o.RelationshipID, true
}

// HasRelationshipID returns a boolean if a field has been set.
func (o *RelationshipInfoSummaryType) HasRelationshipID() bool {
	if o != nil && !IsNil(o.RelationshipID) {
		return true
	}

	return false
}

// SetRelationshipID gets a reference to the given string and assigns it to the RelationshipID field.
func (o *RelationshipInfoSummaryType) SetRelationshipID(v string) {
	o.RelationshipID = &v
}

// GetSourceRelation returns the SourceRelation field value if set, zero value otherwise.
func (o *RelationshipInfoSummaryType) GetSourceRelation() string {
	if o == nil || IsNil(o.SourceRelation) {
		var ret string
		return ret
	}
	return *o.SourceRelation
}

// GetSourceRelationOk returns a tuple with the SourceRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoSummaryType) GetSourceRelationOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRelation) {
		return nil, false
	}
	return o.SourceRelation, true
}

// HasSourceRelation returns a boolean if a field has been set.
func (o *RelationshipInfoSummaryType) HasSourceRelation() bool {
	if o != nil && !IsNil(o.SourceRelation) {
		return true
	}

	return false
}

// SetSourceRelation gets a reference to the given string and assigns it to the SourceRelation field.
func (o *RelationshipInfoSummaryType) SetSourceRelation(v string) {
	o.SourceRelation = &v
}

// GetSourceRelationDescription returns the SourceRelationDescription field value if set, zero value otherwise.
func (o *RelationshipInfoSummaryType) GetSourceRelationDescription() string {
	if o == nil || IsNil(o.SourceRelationDescription) {
		var ret string
		return ret
	}
	return *o.SourceRelationDescription
}

// GetSourceRelationDescriptionOk returns a tuple with the SourceRelationDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoSummaryType) GetSourceRelationDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRelationDescription) {
		return nil, false
	}
	return o.SourceRelationDescription, true
}

// HasSourceRelationDescription returns a boolean if a field has been set.
func (o *RelationshipInfoSummaryType) HasSourceRelationDescription() bool {
	if o != nil && !IsNil(o.SourceRelationDescription) {
		return true
	}

	return false
}

// SetSourceRelationDescription gets a reference to the given string and assigns it to the SourceRelationDescription field.
func (o *RelationshipInfoSummaryType) SetSourceRelationDescription(v string) {
	o.SourceRelationDescription = &v
}

// GetTargetRelation returns the TargetRelation field value if set, zero value otherwise.
func (o *RelationshipInfoSummaryType) GetTargetRelation() string {
	if o == nil || IsNil(o.TargetRelation) {
		var ret string
		return ret
	}
	return *o.TargetRelation
}

// GetTargetRelationOk returns a tuple with the TargetRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoSummaryType) GetTargetRelationOk() (*string, bool) {
	if o == nil || IsNil(o.TargetRelation) {
		return nil, false
	}
	return o.TargetRelation, true
}

// HasTargetRelation returns a boolean if a field has been set.
func (o *RelationshipInfoSummaryType) HasTargetRelation() bool {
	if o != nil && !IsNil(o.TargetRelation) {
		return true
	}

	return false
}

// SetTargetRelation gets a reference to the given string and assigns it to the TargetRelation field.
func (o *RelationshipInfoSummaryType) SetTargetRelation(v string) {
	o.TargetRelation = &v
}

// GetTargetRelationDescription returns the TargetRelationDescription field value if set, zero value otherwise.
func (o *RelationshipInfoSummaryType) GetTargetRelationDescription() string {
	if o == nil || IsNil(o.TargetRelationDescription) {
		var ret string
		return ret
	}
	return *o.TargetRelationDescription
}

// GetTargetRelationDescriptionOk returns a tuple with the TargetRelationDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoSummaryType) GetTargetRelationDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetRelationDescription) {
		return nil, false
	}
	return o.TargetRelationDescription, true
}

// HasTargetRelationDescription returns a boolean if a field has been set.
func (o *RelationshipInfoSummaryType) HasTargetRelationDescription() bool {
	if o != nil && !IsNil(o.TargetRelationDescription) {
		return true
	}

	return false
}

// SetTargetRelationDescription gets a reference to the given string and assigns it to the TargetRelationDescription field.
func (o *RelationshipInfoSummaryType) SetTargetRelationDescription(v string) {
	o.TargetRelationDescription = &v
}

func (o RelationshipInfoSummaryType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelationshipInfoSummaryType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelationshipProfile) {
		toSerialize["relationshipProfile"] = o.RelationshipProfile
	}
	if !IsNil(o.MasterAccountDetails) {
		toSerialize["masterAccountDetails"] = o.MasterAccountDetails
	}
	if !IsNil(o.RelationshipID) {
		toSerialize["relationshipID"] = o.RelationshipID
	}
	if !IsNil(o.SourceRelation) {
		toSerialize["sourceRelation"] = o.SourceRelation
	}
	if !IsNil(o.SourceRelationDescription) {
		toSerialize["sourceRelationDescription"] = o.SourceRelationDescription
	}
	if !IsNil(o.TargetRelation) {
		toSerialize["targetRelation"] = o.TargetRelation
	}
	if !IsNil(o.TargetRelationDescription) {
		toSerialize["targetRelationDescription"] = o.TargetRelationDescription
	}
	return toSerialize, nil
}

type NullableRelationshipInfoSummaryType struct {
	value *RelationshipInfoSummaryType
	isSet bool
}

func (v NullableRelationshipInfoSummaryType) Get() *RelationshipInfoSummaryType {
	return v.value
}

func (v *NullableRelationshipInfoSummaryType) Set(val *RelationshipInfoSummaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipInfoSummaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipInfoSummaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipInfoSummaryType(val *RelationshipInfoSummaryType) *NullableRelationshipInfoSummaryType {
	return &NullableRelationshipInfoSummaryType{value: val, isSet: true}
}

func (v NullableRelationshipInfoSummaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipInfoSummaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


