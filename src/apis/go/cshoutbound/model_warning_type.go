/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the WarningType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarningType{}

// WarningType Used when a message has been successfully processed to report any warnings or business errors that occurred.
type WarningType struct {
	// Property Value
	Value *string `json:"value,omitempty"`
	// An abbreviated version of the error in textual format.
	ShortText *string `json:"shortText,omitempty"`
	// If present, this refers to a table of coded values exchanged between applications to identify errors or warnings.
	Code *string `json:"code,omitempty"`
	// If present, this URL refers to an online description of the error that occurred.
	DocURL *string `json:"docURL,omitempty"`
	// If present, recommended values are those enumerated in the ErrorRS, (NotProcessed Incomplete Complete Unknown) however, the data type is designated as string data, recognizing that trading partners may identify additional status conditions not included in the enumeration.
	Status *string `json:"status,omitempty"`
	// If present, this attribute may identify an unknown or misspelled tag that caused an error in processing. It is recommended that the Tag attribute use XPath notation to identify the location of a tag in the event that more than one tag of the same name is present in the document. Alternatively, the tag name alone can be used to identify missing data [Type=ReqFieldMissing].
	Tag *string `json:"tag,omitempty"`
	// If present, this attribute allows for batch processing and the identification of the record that failed amongst a group of records. This value may contain a concatenation of a unique failed transaction ID with specific record(s) associated with that transaction.
	RecordId *string `json:"recordId,omitempty"`
	// The Warning element MUST contain the Type attribute that uses a recommended set of values to indicate the warning type. The validating XSD can expect to accept values that it has NOT been explicitly coded for and process them by using Type =\"Unknown\".
	Type *string `json:"type,omitempty"`
	// Language identification.
	Language *string `json:"language,omitempty"`
	// Reference Place Holder used as an index for this warning.
	Rph *string `json:"rph,omitempty"`
}

// NewWarningType instantiates a new WarningType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarningType() *WarningType {
	this := WarningType{}
	return &this
}

// NewWarningTypeWithDefaults instantiates a new WarningType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningTypeWithDefaults() *WarningType {
	this := WarningType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WarningType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WarningType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WarningType) SetValue(v string) {
	o.Value = &v
}

// GetShortText returns the ShortText field value if set, zero value otherwise.
func (o *WarningType) GetShortText() string {
	if o == nil || IsNil(o.ShortText) {
		var ret string
		return ret
	}
	return *o.ShortText
}

// GetShortTextOk returns a tuple with the ShortText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetShortTextOk() (*string, bool) {
	if o == nil || IsNil(o.ShortText) {
		return nil, false
	}
	return o.ShortText, true
}

// HasShortText returns a boolean if a field has been set.
func (o *WarningType) HasShortText() bool {
	if o != nil && !IsNil(o.ShortText) {
		return true
	}

	return false
}

// SetShortText gets a reference to the given string and assigns it to the ShortText field.
func (o *WarningType) SetShortText(v string) {
	o.ShortText = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *WarningType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *WarningType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *WarningType) SetCode(v string) {
	o.Code = &v
}

// GetDocURL returns the DocURL field value if set, zero value otherwise.
func (o *WarningType) GetDocURL() string {
	if o == nil || IsNil(o.DocURL) {
		var ret string
		return ret
	}
	return *o.DocURL
}

// GetDocURLOk returns a tuple with the DocURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetDocURLOk() (*string, bool) {
	if o == nil || IsNil(o.DocURL) {
		return nil, false
	}
	return o.DocURL, true
}

// HasDocURL returns a boolean if a field has been set.
func (o *WarningType) HasDocURL() bool {
	if o != nil && !IsNil(o.DocURL) {
		return true
	}

	return false
}

// SetDocURL gets a reference to the given string and assigns it to the DocURL field.
func (o *WarningType) SetDocURL(v string) {
	o.DocURL = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WarningType) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WarningType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WarningType) SetStatus(v string) {
	o.Status = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *WarningType) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *WarningType) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *WarningType) SetTag(v string) {
	o.Tag = &v
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
func (o *WarningType) GetRecordId() string {
	if o == nil || IsNil(o.RecordId) {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordId) {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *WarningType) HasRecordId() bool {
	if o != nil && !IsNil(o.RecordId) {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
func (o *WarningType) SetRecordId(v string) {
	o.RecordId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WarningType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WarningType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WarningType) SetType(v string) {
	o.Type = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *WarningType) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *WarningType) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *WarningType) SetLanguage(v string) {
	o.Language = &v
}

// GetRph returns the Rph field value if set, zero value otherwise.
func (o *WarningType) GetRph() string {
	if o == nil || IsNil(o.Rph) {
		var ret string
		return ret
	}
	return *o.Rph
}

// GetRphOk returns a tuple with the Rph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningType) GetRphOk() (*string, bool) {
	if o == nil || IsNil(o.Rph) {
		return nil, false
	}
	return o.Rph, true
}

// HasRph returns a boolean if a field has been set.
func (o *WarningType) HasRph() bool {
	if o != nil && !IsNil(o.Rph) {
		return true
	}

	return false
}

// SetRph gets a reference to the given string and assigns it to the Rph field.
func (o *WarningType) SetRph(v string) {
	o.Rph = &v
}

func (o WarningType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarningType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ShortText) {
		toSerialize["shortText"] = o.ShortText
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.DocURL) {
		toSerialize["docURL"] = o.DocURL
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.RecordId) {
		toSerialize["recordId"] = o.RecordId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Rph) {
		toSerialize["rph"] = o.Rph
	}
	return toSerialize, nil
}

type NullableWarningType struct {
	value *WarningType
	isSet bool
}

func (v NullableWarningType) Get() *WarningType {
	return v.value
}

func (v *NullableWarningType) Set(val *WarningType) {
	v.value = val
	v.isSet = true
}

func (v NullableWarningType) IsSet() bool {
	return v.isSet
}

func (v *NullableWarningType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarningType(val *WarningType) *NullableWarningType {
	return &NullableWarningType{value: val, isSet: true}
}

func (v NullableWarningType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarningType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


