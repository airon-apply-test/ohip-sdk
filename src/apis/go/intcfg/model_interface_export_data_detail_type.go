/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InterfaceExportDataDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceExportDataDetailType{}

// InterfaceExportDataDetailType This gives information of export detail of a hotel interface.
type InterfaceExportDataDetailType struct {
	// Specifies the table where the event data is stored.
	SourceTableName *string `json:"sourceTableName,omitempty"`
	// Specifies the column name found in the table.
	ColumnName *string `json:"columnName,omitempty"`
	// Specifies the type of data in the column.
	DataType *string `json:"dataType,omitempty"`
	// Specifies the length of data in the column.
	DataLength *int32 `json:"dataLength,omitempty"`
	// Specifies the tag name found in the export.
	XMLTagName *string `json:"xMLTagName,omitempty"`
	// Specifies whether the column is chosen to be included in the export.
	ColumnSelected *bool `json:"columnSelected,omitempty"`
}

// NewInterfaceExportDataDetailType instantiates a new InterfaceExportDataDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceExportDataDetailType() *InterfaceExportDataDetailType {
	this := InterfaceExportDataDetailType{}
	return &this
}

// NewInterfaceExportDataDetailTypeWithDefaults instantiates a new InterfaceExportDataDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceExportDataDetailTypeWithDefaults() *InterfaceExportDataDetailType {
	this := InterfaceExportDataDetailType{}
	return &this
}

// GetSourceTableName returns the SourceTableName field value if set, zero value otherwise.
func (o *InterfaceExportDataDetailType) GetSourceTableName() string {
	if o == nil || IsNil(o.SourceTableName) {
		var ret string
		return ret
	}
	return *o.SourceTableName
}

// GetSourceTableNameOk returns a tuple with the SourceTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceExportDataDetailType) GetSourceTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceTableName) {
		return nil, false
	}
	return o.SourceTableName, true
}

// HasSourceTableName returns a boolean if a field has been set.
func (o *InterfaceExportDataDetailType) HasSourceTableName() bool {
	if o != nil && !IsNil(o.SourceTableName) {
		return true
	}

	return false
}

// SetSourceTableName gets a reference to the given string and assigns it to the SourceTableName field.
func (o *InterfaceExportDataDetailType) SetSourceTableName(v string) {
	o.SourceTableName = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *InterfaceExportDataDetailType) GetColumnName() string {
	if o == nil || IsNil(o.ColumnName) {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceExportDataDetailType) GetColumnNameOk() (*string, bool) {
	if o == nil || IsNil(o.ColumnName) {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *InterfaceExportDataDetailType) HasColumnName() bool {
	if o != nil && !IsNil(o.ColumnName) {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *InterfaceExportDataDetailType) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *InterfaceExportDataDetailType) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceExportDataDetailType) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *InterfaceExportDataDetailType) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *InterfaceExportDataDetailType) SetDataType(v string) {
	o.DataType = &v
}

// GetDataLength returns the DataLength field value if set, zero value otherwise.
func (o *InterfaceExportDataDetailType) GetDataLength() int32 {
	if o == nil || IsNil(o.DataLength) {
		var ret int32
		return ret
	}
	return *o.DataLength
}

// GetDataLengthOk returns a tuple with the DataLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceExportDataDetailType) GetDataLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.DataLength) {
		return nil, false
	}
	return o.DataLength, true
}

// HasDataLength returns a boolean if a field has been set.
func (o *InterfaceExportDataDetailType) HasDataLength() bool {
	if o != nil && !IsNil(o.DataLength) {
		return true
	}

	return false
}

// SetDataLength gets a reference to the given int32 and assigns it to the DataLength field.
func (o *InterfaceExportDataDetailType) SetDataLength(v int32) {
	o.DataLength = &v
}

// GetXMLTagName returns the XMLTagName field value if set, zero value otherwise.
func (o *InterfaceExportDataDetailType) GetXMLTagName() string {
	if o == nil || IsNil(o.XMLTagName) {
		var ret string
		return ret
	}
	return *o.XMLTagName
}

// GetXMLTagNameOk returns a tuple with the XMLTagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceExportDataDetailType) GetXMLTagNameOk() (*string, bool) {
	if o == nil || IsNil(o.XMLTagName) {
		return nil, false
	}
	return o.XMLTagName, true
}

// HasXMLTagName returns a boolean if a field has been set.
func (o *InterfaceExportDataDetailType) HasXMLTagName() bool {
	if o != nil && !IsNil(o.XMLTagName) {
		return true
	}

	return false
}

// SetXMLTagName gets a reference to the given string and assigns it to the XMLTagName field.
func (o *InterfaceExportDataDetailType) SetXMLTagName(v string) {
	o.XMLTagName = &v
}

// GetColumnSelected returns the ColumnSelected field value if set, zero value otherwise.
func (o *InterfaceExportDataDetailType) GetColumnSelected() bool {
	if o == nil || IsNil(o.ColumnSelected) {
		var ret bool
		return ret
	}
	return *o.ColumnSelected
}

// GetColumnSelectedOk returns a tuple with the ColumnSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceExportDataDetailType) GetColumnSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.ColumnSelected) {
		return nil, false
	}
	return o.ColumnSelected, true
}

// HasColumnSelected returns a boolean if a field has been set.
func (o *InterfaceExportDataDetailType) HasColumnSelected() bool {
	if o != nil && !IsNil(o.ColumnSelected) {
		return true
	}

	return false
}

// SetColumnSelected gets a reference to the given bool and assigns it to the ColumnSelected field.
func (o *InterfaceExportDataDetailType) SetColumnSelected(v bool) {
	o.ColumnSelected = &v
}

func (o InterfaceExportDataDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceExportDataDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceTableName) {
		toSerialize["sourceTableName"] = o.SourceTableName
	}
	if !IsNil(o.ColumnName) {
		toSerialize["columnName"] = o.ColumnName
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.DataLength) {
		toSerialize["dataLength"] = o.DataLength
	}
	if !IsNil(o.XMLTagName) {
		toSerialize["xMLTagName"] = o.XMLTagName
	}
	if !IsNil(o.ColumnSelected) {
		toSerialize["columnSelected"] = o.ColumnSelected
	}
	return toSerialize, nil
}

type NullableInterfaceExportDataDetailType struct {
	value *InterfaceExportDataDetailType
	isSet bool
}

func (v NullableInterfaceExportDataDetailType) Get() *InterfaceExportDataDetailType {
	return v.value
}

func (v *NullableInterfaceExportDataDetailType) Set(val *InterfaceExportDataDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceExportDataDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceExportDataDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceExportDataDetailType(val *InterfaceExportDataDetailType) *NullableInterfaceExportDataDetailType {
	return &NullableInterfaceExportDataDetailType{value: val, isSet: true}
}

func (v NullableInterfaceExportDataDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceExportDataDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


