/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RelationshipInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelationshipInfoType{}

// RelationshipInfoType Relationship Type contains information about the associations between and among individuals, companies, travel agents, groups, sources, and contact profiles.
type RelationshipInfoType struct {
	ChangeRelationship *ChangeRelationsType `json:"changeRelationship,omitempty"`
	RelationshipProfile *RelationshipProfileType `json:"relationshipProfile,omitempty"`
	MasterAccountInfo *MasterAccountInfoType `json:"masterAccountInfo,omitempty"`
	// Relationship identifier.
	Id *string `json:"id,omitempty"`
	// Indicates the type of relationship the current profile(Source Profile) has with the related profile(Target Profile).
	Relation *string `json:"relation,omitempty"`
	// Displays the description of relationship the current profile(Source Profile) has with the related profile(Target Profile).
	RelationDescription *string `json:"relationDescription,omitempty"`
	// Displays the type of relationship the Related profile(Target Profile) has with the current profile(Source Profile).
	TargetRelation *string `json:"targetRelation,omitempty"`
	// Displays the description of the target relation(Target Profile).
	TargetRelationDescription *string `json:"targetRelationDescription,omitempty"`
}

// NewRelationshipInfoType instantiates a new RelationshipInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipInfoType() *RelationshipInfoType {
	this := RelationshipInfoType{}
	return &this
}

// NewRelationshipInfoTypeWithDefaults instantiates a new RelationshipInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipInfoTypeWithDefaults() *RelationshipInfoType {
	this := RelationshipInfoType{}
	return &this
}

// GetChangeRelationship returns the ChangeRelationship field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetChangeRelationship() ChangeRelationsType {
	if o == nil || IsNil(o.ChangeRelationship) {
		var ret ChangeRelationsType
		return ret
	}
	return *o.ChangeRelationship
}

// GetChangeRelationshipOk returns a tuple with the ChangeRelationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetChangeRelationshipOk() (*ChangeRelationsType, bool) {
	if o == nil || IsNil(o.ChangeRelationship) {
		return nil, false
	}
	return o.ChangeRelationship, true
}

// HasChangeRelationship returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasChangeRelationship() bool {
	if o != nil && !IsNil(o.ChangeRelationship) {
		return true
	}

	return false
}

// SetChangeRelationship gets a reference to the given ChangeRelationsType and assigns it to the ChangeRelationship field.
func (o *RelationshipInfoType) SetChangeRelationship(v ChangeRelationsType) {
	o.ChangeRelationship = &v
}

// GetRelationshipProfile returns the RelationshipProfile field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetRelationshipProfile() RelationshipProfileType {
	if o == nil || IsNil(o.RelationshipProfile) {
		var ret RelationshipProfileType
		return ret
	}
	return *o.RelationshipProfile
}

// GetRelationshipProfileOk returns a tuple with the RelationshipProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetRelationshipProfileOk() (*RelationshipProfileType, bool) {
	if o == nil || IsNil(o.RelationshipProfile) {
		return nil, false
	}
	return o.RelationshipProfile, true
}

// HasRelationshipProfile returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasRelationshipProfile() bool {
	if o != nil && !IsNil(o.RelationshipProfile) {
		return true
	}

	return false
}

// SetRelationshipProfile gets a reference to the given RelationshipProfileType and assigns it to the RelationshipProfile field.
func (o *RelationshipInfoType) SetRelationshipProfile(v RelationshipProfileType) {
	o.RelationshipProfile = &v
}

// GetMasterAccountInfo returns the MasterAccountInfo field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetMasterAccountInfo() MasterAccountInfoType {
	if o == nil || IsNil(o.MasterAccountInfo) {
		var ret MasterAccountInfoType
		return ret
	}
	return *o.MasterAccountInfo
}

// GetMasterAccountInfoOk returns a tuple with the MasterAccountInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetMasterAccountInfoOk() (*MasterAccountInfoType, bool) {
	if o == nil || IsNil(o.MasterAccountInfo) {
		return nil, false
	}
	return o.MasterAccountInfo, true
}

// HasMasterAccountInfo returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasMasterAccountInfo() bool {
	if o != nil && !IsNil(o.MasterAccountInfo) {
		return true
	}

	return false
}

// SetMasterAccountInfo gets a reference to the given MasterAccountInfoType and assigns it to the MasterAccountInfo field.
func (o *RelationshipInfoType) SetMasterAccountInfo(v MasterAccountInfoType) {
	o.MasterAccountInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelationshipInfoType) SetId(v string) {
	o.Id = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetRelation() string {
	if o == nil || IsNil(o.Relation) {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetRelationOk() (*string, bool) {
	if o == nil || IsNil(o.Relation) {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasRelation() bool {
	if o != nil && !IsNil(o.Relation) {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *RelationshipInfoType) SetRelation(v string) {
	o.Relation = &v
}

// GetRelationDescription returns the RelationDescription field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetRelationDescription() string {
	if o == nil || IsNil(o.RelationDescription) {
		var ret string
		return ret
	}
	return *o.RelationDescription
}

// GetRelationDescriptionOk returns a tuple with the RelationDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetRelationDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.RelationDescription) {
		return nil, false
	}
	return o.RelationDescription, true
}

// HasRelationDescription returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasRelationDescription() bool {
	if o != nil && !IsNil(o.RelationDescription) {
		return true
	}

	return false
}

// SetRelationDescription gets a reference to the given string and assigns it to the RelationDescription field.
func (o *RelationshipInfoType) SetRelationDescription(v string) {
	o.RelationDescription = &v
}

// GetTargetRelation returns the TargetRelation field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetTargetRelation() string {
	if o == nil || IsNil(o.TargetRelation) {
		var ret string
		return ret
	}
	return *o.TargetRelation
}

// GetTargetRelationOk returns a tuple with the TargetRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetTargetRelationOk() (*string, bool) {
	if o == nil || IsNil(o.TargetRelation) {
		return nil, false
	}
	return o.TargetRelation, true
}

// HasTargetRelation returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasTargetRelation() bool {
	if o != nil && !IsNil(o.TargetRelation) {
		return true
	}

	return false
}

// SetTargetRelation gets a reference to the given string and assigns it to the TargetRelation field.
func (o *RelationshipInfoType) SetTargetRelation(v string) {
	o.TargetRelation = &v
}

// GetTargetRelationDescription returns the TargetRelationDescription field value if set, zero value otherwise.
func (o *RelationshipInfoType) GetTargetRelationDescription() string {
	if o == nil || IsNil(o.TargetRelationDescription) {
		var ret string
		return ret
	}
	return *o.TargetRelationDescription
}

// GetTargetRelationDescriptionOk returns a tuple with the TargetRelationDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipInfoType) GetTargetRelationDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetRelationDescription) {
		return nil, false
	}
	return o.TargetRelationDescription, true
}

// HasTargetRelationDescription returns a boolean if a field has been set.
func (o *RelationshipInfoType) HasTargetRelationDescription() bool {
	if o != nil && !IsNil(o.TargetRelationDescription) {
		return true
	}

	return false
}

// SetTargetRelationDescription gets a reference to the given string and assigns it to the TargetRelationDescription field.
func (o *RelationshipInfoType) SetTargetRelationDescription(v string) {
	o.TargetRelationDescription = &v
}

func (o RelationshipInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelationshipInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangeRelationship) {
		toSerialize["changeRelationship"] = o.ChangeRelationship
	}
	if !IsNil(o.RelationshipProfile) {
		toSerialize["relationshipProfile"] = o.RelationshipProfile
	}
	if !IsNil(o.MasterAccountInfo) {
		toSerialize["masterAccountInfo"] = o.MasterAccountInfo
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Relation) {
		toSerialize["relation"] = o.Relation
	}
	if !IsNil(o.RelationDescription) {
		toSerialize["relationDescription"] = o.RelationDescription
	}
	if !IsNil(o.TargetRelation) {
		toSerialize["targetRelation"] = o.TargetRelation
	}
	if !IsNil(o.TargetRelationDescription) {
		toSerialize["targetRelationDescription"] = o.TargetRelationDescription
	}
	return toSerialize, nil
}

type NullableRelationshipInfoType struct {
	value *RelationshipInfoType
	isSet bool
}

func (v NullableRelationshipInfoType) Get() *RelationshipInfoType {
	return v.value
}

func (v *NullableRelationshipInfoType) Set(val *RelationshipInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipInfoType(val *RelationshipInfoType) *NullableRelationshipInfoType {
	return &NullableRelationshipInfoType{value: val, isSet: true}
}

func (v NullableRelationshipInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


