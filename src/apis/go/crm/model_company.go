/*
OPERA Cloud Customer Relationship Management API

APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Company type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Company{}

// Company Request object for creation of company/agent/group/source profile. This object contains profile details with unique identifiers of a profile. The standard optional Opera Context element is also included.
type Company struct {
	// Unique Id that references an object uniquely in the system.
	CompanyIdList []UniqueIDType `json:"companyIdList,omitempty"`
	// This type contains unique information of external reference.
	ExternalReferences []ExternalReferenceType `json:"externalReferences,omitempty"`
	CompanyDetails *CompanyProfileType `json:"companyDetails,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewCompany instantiates a new Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompany() *Company {
	this := Company{}
	return &this
}

// NewCompanyWithDefaults instantiates a new Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWithDefaults() *Company {
	this := Company{}
	return &this
}

// GetCompanyIdList returns the CompanyIdList field value if set, zero value otherwise.
func (o *Company) GetCompanyIdList() []UniqueIDType {
	if o == nil || IsNil(o.CompanyIdList) {
		var ret []UniqueIDType
		return ret
	}
	return o.CompanyIdList
}

// GetCompanyIdListOk returns a tuple with the CompanyIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCompanyIdListOk() ([]UniqueIDType, bool) {
	if o == nil || IsNil(o.CompanyIdList) {
		return nil, false
	}
	return o.CompanyIdList, true
}

// HasCompanyIdList returns a boolean if a field has been set.
func (o *Company) HasCompanyIdList() bool {
	if o != nil && !IsNil(o.CompanyIdList) {
		return true
	}

	return false
}

// SetCompanyIdList gets a reference to the given []UniqueIDType and assigns it to the CompanyIdList field.
func (o *Company) SetCompanyIdList(v []UniqueIDType) {
	o.CompanyIdList = v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *Company) GetExternalReferences() []ExternalReferenceType {
	if o == nil || IsNil(o.ExternalReferences) {
		var ret []ExternalReferenceType
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetExternalReferencesOk() ([]ExternalReferenceType, bool) {
	if o == nil || IsNil(o.ExternalReferences) {
		return nil, false
	}
	return o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *Company) HasExternalReferences() bool {
	if o != nil && !IsNil(o.ExternalReferences) {
		return true
	}

	return false
}

// SetExternalReferences gets a reference to the given []ExternalReferenceType and assigns it to the ExternalReferences field.
func (o *Company) SetExternalReferences(v []ExternalReferenceType) {
	o.ExternalReferences = v
}

// GetCompanyDetails returns the CompanyDetails field value if set, zero value otherwise.
func (o *Company) GetCompanyDetails() CompanyProfileType {
	if o == nil || IsNil(o.CompanyDetails) {
		var ret CompanyProfileType
		return ret
	}
	return *o.CompanyDetails
}

// GetCompanyDetailsOk returns a tuple with the CompanyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCompanyDetailsOk() (*CompanyProfileType, bool) {
	if o == nil || IsNil(o.CompanyDetails) {
		return nil, false
	}
	return o.CompanyDetails, true
}

// HasCompanyDetails returns a boolean if a field has been set.
func (o *Company) HasCompanyDetails() bool {
	if o != nil && !IsNil(o.CompanyDetails) {
		return true
	}

	return false
}

// SetCompanyDetails gets a reference to the given CompanyProfileType and assigns it to the CompanyDetails field.
func (o *Company) SetCompanyDetails(v CompanyProfileType) {
	o.CompanyDetails = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Company) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Company) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *Company) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Company) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Company) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *Company) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o Company) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Company) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyIdList) {
		toSerialize["companyIdList"] = o.CompanyIdList
	}
	if !IsNil(o.ExternalReferences) {
		toSerialize["externalReferences"] = o.ExternalReferences
	}
	if !IsNil(o.CompanyDetails) {
		toSerialize["companyDetails"] = o.CompanyDetails
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableCompany struct {
	value *Company
	isSet bool
}

func (v NullableCompany) Get() *Company {
	return v.value
}

func (v *NullableCompany) Set(val *Company) {
	v.value = val
	v.isSet = true
}

func (v NullableCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompany(val *Company) *NullableCompany {
	return &NullableCompany{value: val, isSet: true}
}

func (v NullableCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


