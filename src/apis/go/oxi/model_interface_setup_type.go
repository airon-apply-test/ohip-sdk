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

// checks if the InterfaceSetupType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceSetupType{}

// InterfaceSetupType Type represents ID one Interface Setup
type InterfaceSetupType struct {
	// ID of the Interface Setup
	InterfaceId *string `json:"interfaceId,omitempty"`
	// Property for which the Interface Setup is defined.
	HotelId *string `json:"hotelId,omitempty"`
	// Corresponding property in the external system
	ExternalHotelCode *string `json:"externalHotelCode,omitempty"`
	// Logical Name of the machine that runs the OXIHUB
	Machine *int32 `json:"machine,omitempty"`
	// Detailed information about the interface.
	Description *string `json:"description,omitempty"`
	// Type of the message. (XML, TPI, AMF, etc.).
	MessageFormat *string `json:"messageFormat,omitempty"`
	// Collection of XMLTypes
	XMLVersions []InterfaceSetupXMLVersionType `json:"xMLVersions,omitempty"`
	// Database ID
	DatabaseId *string `json:"databaseId,omitempty"`
	// Type of the interface (UPLOAD, DOWNLOAD, EXPORT, FTCRS, HOLIDEX, etc.,).
	InterfaceType *string `json:"interfaceType,omitempty"`
	// Indicates if external system is active
	ExternalSystemActivated *bool `json:"externalSystemActivated,omitempty"`
	SystemType *InterfaceSystemType `json:"systemType,omitempty"`
	// Select for all interfaces that send data from Opera to an external system.
	BatchProcessBE *bool `json:"batchProcessBE,omitempty"`
	// Compress all accumulated Business Events if true.
	ZipData *bool `json:"zipData,omitempty"`
	// Indicates external system does not send the full message for changes.
	DeltaMode *bool `json:"deltaMode,omitempty"`
	DeletionIndicator *InterfaceSetupMessageIndicatorType `json:"deletionIndicator,omitempty"`
	KeepingIndicator *InterfaceSetupMessageIndicatorType `json:"keepingIndicator,omitempty"`
	// Indicates if interface is ORS destination.
	OrsDestination *bool `json:"orsDestination,omitempty"`
	// Indicates if XML Versions Negotiable.
	XmlSchemaVersionsNegotiable *bool `json:"xmlSchemaVersionsNegotiable,omitempty"`
	// Indicates to set all XML Versions to Max version.
	NegotiateXmlSchemaVersion *bool `json:"negotiateXmlSchemaVersion,omitempty"`
	// XML Version to be used by all XLM schemas.
	SetAllToVersion *string `json:"setAllToVersion,omitempty"`
}

// NewInterfaceSetupType instantiates a new InterfaceSetupType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceSetupType() *InterfaceSetupType {
	this := InterfaceSetupType{}
	return &this
}

// NewInterfaceSetupTypeWithDefaults instantiates a new InterfaceSetupType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceSetupTypeWithDefaults() *InterfaceSetupType {
	this := InterfaceSetupType{}
	return &this
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *InterfaceSetupType) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *InterfaceSetupType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetExternalHotelCode returns the ExternalHotelCode field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetExternalHotelCode() string {
	if o == nil || IsNil(o.ExternalHotelCode) {
		var ret string
		return ret
	}
	return *o.ExternalHotelCode
}

// GetExternalHotelCodeOk returns a tuple with the ExternalHotelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetExternalHotelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalHotelCode) {
		return nil, false
	}
	return o.ExternalHotelCode, true
}

// HasExternalHotelCode returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasExternalHotelCode() bool {
	if o != nil && !IsNil(o.ExternalHotelCode) {
		return true
	}

	return false
}

// SetExternalHotelCode gets a reference to the given string and assigns it to the ExternalHotelCode field.
func (o *InterfaceSetupType) SetExternalHotelCode(v string) {
	o.ExternalHotelCode = &v
}

// GetMachine returns the Machine field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetMachine() int32 {
	if o == nil || IsNil(o.Machine) {
		var ret int32
		return ret
	}
	return *o.Machine
}

// GetMachineOk returns a tuple with the Machine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetMachineOk() (*int32, bool) {
	if o == nil || IsNil(o.Machine) {
		return nil, false
	}
	return o.Machine, true
}

// HasMachine returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasMachine() bool {
	if o != nil && !IsNil(o.Machine) {
		return true
	}

	return false
}

// SetMachine gets a reference to the given int32 and assigns it to the Machine field.
func (o *InterfaceSetupType) SetMachine(v int32) {
	o.Machine = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterfaceSetupType) SetDescription(v string) {
	o.Description = &v
}

// GetMessageFormat returns the MessageFormat field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetMessageFormat() string {
	if o == nil || IsNil(o.MessageFormat) {
		var ret string
		return ret
	}
	return *o.MessageFormat
}

// GetMessageFormatOk returns a tuple with the MessageFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetMessageFormatOk() (*string, bool) {
	if o == nil || IsNil(o.MessageFormat) {
		return nil, false
	}
	return o.MessageFormat, true
}

// HasMessageFormat returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasMessageFormat() bool {
	if o != nil && !IsNil(o.MessageFormat) {
		return true
	}

	return false
}

// SetMessageFormat gets a reference to the given string and assigns it to the MessageFormat field.
func (o *InterfaceSetupType) SetMessageFormat(v string) {
	o.MessageFormat = &v
}

// GetXMLVersions returns the XMLVersions field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetXMLVersions() []InterfaceSetupXMLVersionType {
	if o == nil || IsNil(o.XMLVersions) {
		var ret []InterfaceSetupXMLVersionType
		return ret
	}
	return o.XMLVersions
}

// GetXMLVersionsOk returns a tuple with the XMLVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetXMLVersionsOk() ([]InterfaceSetupXMLVersionType, bool) {
	if o == nil || IsNil(o.XMLVersions) {
		return nil, false
	}
	return o.XMLVersions, true
}

// HasXMLVersions returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasXMLVersions() bool {
	if o != nil && !IsNil(o.XMLVersions) {
		return true
	}

	return false
}

// SetXMLVersions gets a reference to the given []InterfaceSetupXMLVersionType and assigns it to the XMLVersions field.
func (o *InterfaceSetupType) SetXMLVersions(v []InterfaceSetupXMLVersionType) {
	o.XMLVersions = v
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetDatabaseId() string {
	if o == nil || IsNil(o.DatabaseId) {
		var ret string
		return ret
	}
	return *o.DatabaseId
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetDatabaseIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseId) {
		return nil, false
	}
	return o.DatabaseId, true
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasDatabaseId() bool {
	if o != nil && !IsNil(o.DatabaseId) {
		return true
	}

	return false
}

// SetDatabaseId gets a reference to the given string and assigns it to the DatabaseId field.
func (o *InterfaceSetupType) SetDatabaseId(v string) {
	o.DatabaseId = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetInterfaceType() string {
	if o == nil || IsNil(o.InterfaceType) {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceType) {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasInterfaceType() bool {
	if o != nil && !IsNil(o.InterfaceType) {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *InterfaceSetupType) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetExternalSystemActivated returns the ExternalSystemActivated field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetExternalSystemActivated() bool {
	if o == nil || IsNil(o.ExternalSystemActivated) {
		var ret bool
		return ret
	}
	return *o.ExternalSystemActivated
}

// GetExternalSystemActivatedOk returns a tuple with the ExternalSystemActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetExternalSystemActivatedOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalSystemActivated) {
		return nil, false
	}
	return o.ExternalSystemActivated, true
}

// HasExternalSystemActivated returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasExternalSystemActivated() bool {
	if o != nil && !IsNil(o.ExternalSystemActivated) {
		return true
	}

	return false
}

// SetExternalSystemActivated gets a reference to the given bool and assigns it to the ExternalSystemActivated field.
func (o *InterfaceSetupType) SetExternalSystemActivated(v bool) {
	o.ExternalSystemActivated = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetSystemType() InterfaceSystemType {
	if o == nil || IsNil(o.SystemType) {
		var ret InterfaceSystemType
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetSystemTypeOk() (*InterfaceSystemType, bool) {
	if o == nil || IsNil(o.SystemType) {
		return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasSystemType() bool {
	if o != nil && !IsNil(o.SystemType) {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given InterfaceSystemType and assigns it to the SystemType field.
func (o *InterfaceSetupType) SetSystemType(v InterfaceSystemType) {
	o.SystemType = &v
}

// GetBatchProcessBE returns the BatchProcessBE field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetBatchProcessBE() bool {
	if o == nil || IsNil(o.BatchProcessBE) {
		var ret bool
		return ret
	}
	return *o.BatchProcessBE
}

// GetBatchProcessBEOk returns a tuple with the BatchProcessBE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetBatchProcessBEOk() (*bool, bool) {
	if o == nil || IsNil(o.BatchProcessBE) {
		return nil, false
	}
	return o.BatchProcessBE, true
}

// HasBatchProcessBE returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasBatchProcessBE() bool {
	if o != nil && !IsNil(o.BatchProcessBE) {
		return true
	}

	return false
}

// SetBatchProcessBE gets a reference to the given bool and assigns it to the BatchProcessBE field.
func (o *InterfaceSetupType) SetBatchProcessBE(v bool) {
	o.BatchProcessBE = &v
}

// GetZipData returns the ZipData field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetZipData() bool {
	if o == nil || IsNil(o.ZipData) {
		var ret bool
		return ret
	}
	return *o.ZipData
}

// GetZipDataOk returns a tuple with the ZipData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetZipDataOk() (*bool, bool) {
	if o == nil || IsNil(o.ZipData) {
		return nil, false
	}
	return o.ZipData, true
}

// HasZipData returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasZipData() bool {
	if o != nil && !IsNil(o.ZipData) {
		return true
	}

	return false
}

// SetZipData gets a reference to the given bool and assigns it to the ZipData field.
func (o *InterfaceSetupType) SetZipData(v bool) {
	o.ZipData = &v
}

// GetDeltaMode returns the DeltaMode field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetDeltaMode() bool {
	if o == nil || IsNil(o.DeltaMode) {
		var ret bool
		return ret
	}
	return *o.DeltaMode
}

// GetDeltaModeOk returns a tuple with the DeltaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetDeltaModeOk() (*bool, bool) {
	if o == nil || IsNil(o.DeltaMode) {
		return nil, false
	}
	return o.DeltaMode, true
}

// HasDeltaMode returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasDeltaMode() bool {
	if o != nil && !IsNil(o.DeltaMode) {
		return true
	}

	return false
}

// SetDeltaMode gets a reference to the given bool and assigns it to the DeltaMode field.
func (o *InterfaceSetupType) SetDeltaMode(v bool) {
	o.DeltaMode = &v
}

// GetDeletionIndicator returns the DeletionIndicator field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetDeletionIndicator() InterfaceSetupMessageIndicatorType {
	if o == nil || IsNil(o.DeletionIndicator) {
		var ret InterfaceSetupMessageIndicatorType
		return ret
	}
	return *o.DeletionIndicator
}

// GetDeletionIndicatorOk returns a tuple with the DeletionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetDeletionIndicatorOk() (*InterfaceSetupMessageIndicatorType, bool) {
	if o == nil || IsNil(o.DeletionIndicator) {
		return nil, false
	}
	return o.DeletionIndicator, true
}

// HasDeletionIndicator returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasDeletionIndicator() bool {
	if o != nil && !IsNil(o.DeletionIndicator) {
		return true
	}

	return false
}

// SetDeletionIndicator gets a reference to the given InterfaceSetupMessageIndicatorType and assigns it to the DeletionIndicator field.
func (o *InterfaceSetupType) SetDeletionIndicator(v InterfaceSetupMessageIndicatorType) {
	o.DeletionIndicator = &v
}

// GetKeepingIndicator returns the KeepingIndicator field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetKeepingIndicator() InterfaceSetupMessageIndicatorType {
	if o == nil || IsNil(o.KeepingIndicator) {
		var ret InterfaceSetupMessageIndicatorType
		return ret
	}
	return *o.KeepingIndicator
}

// GetKeepingIndicatorOk returns a tuple with the KeepingIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetKeepingIndicatorOk() (*InterfaceSetupMessageIndicatorType, bool) {
	if o == nil || IsNil(o.KeepingIndicator) {
		return nil, false
	}
	return o.KeepingIndicator, true
}

// HasKeepingIndicator returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasKeepingIndicator() bool {
	if o != nil && !IsNil(o.KeepingIndicator) {
		return true
	}

	return false
}

// SetKeepingIndicator gets a reference to the given InterfaceSetupMessageIndicatorType and assigns it to the KeepingIndicator field.
func (o *InterfaceSetupType) SetKeepingIndicator(v InterfaceSetupMessageIndicatorType) {
	o.KeepingIndicator = &v
}

// GetOrsDestination returns the OrsDestination field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetOrsDestination() bool {
	if o == nil || IsNil(o.OrsDestination) {
		var ret bool
		return ret
	}
	return *o.OrsDestination
}

// GetOrsDestinationOk returns a tuple with the OrsDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetOrsDestinationOk() (*bool, bool) {
	if o == nil || IsNil(o.OrsDestination) {
		return nil, false
	}
	return o.OrsDestination, true
}

// HasOrsDestination returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasOrsDestination() bool {
	if o != nil && !IsNil(o.OrsDestination) {
		return true
	}

	return false
}

// SetOrsDestination gets a reference to the given bool and assigns it to the OrsDestination field.
func (o *InterfaceSetupType) SetOrsDestination(v bool) {
	o.OrsDestination = &v
}

// GetXmlSchemaVersionsNegotiable returns the XmlSchemaVersionsNegotiable field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetXmlSchemaVersionsNegotiable() bool {
	if o == nil || IsNil(o.XmlSchemaVersionsNegotiable) {
		var ret bool
		return ret
	}
	return *o.XmlSchemaVersionsNegotiable
}

// GetXmlSchemaVersionsNegotiableOk returns a tuple with the XmlSchemaVersionsNegotiable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetXmlSchemaVersionsNegotiableOk() (*bool, bool) {
	if o == nil || IsNil(o.XmlSchemaVersionsNegotiable) {
		return nil, false
	}
	return o.XmlSchemaVersionsNegotiable, true
}

// HasXmlSchemaVersionsNegotiable returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasXmlSchemaVersionsNegotiable() bool {
	if o != nil && !IsNil(o.XmlSchemaVersionsNegotiable) {
		return true
	}

	return false
}

// SetXmlSchemaVersionsNegotiable gets a reference to the given bool and assigns it to the XmlSchemaVersionsNegotiable field.
func (o *InterfaceSetupType) SetXmlSchemaVersionsNegotiable(v bool) {
	o.XmlSchemaVersionsNegotiable = &v
}

// GetNegotiateXmlSchemaVersion returns the NegotiateXmlSchemaVersion field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetNegotiateXmlSchemaVersion() bool {
	if o == nil || IsNil(o.NegotiateXmlSchemaVersion) {
		var ret bool
		return ret
	}
	return *o.NegotiateXmlSchemaVersion
}

// GetNegotiateXmlSchemaVersionOk returns a tuple with the NegotiateXmlSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetNegotiateXmlSchemaVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.NegotiateXmlSchemaVersion) {
		return nil, false
	}
	return o.NegotiateXmlSchemaVersion, true
}

// HasNegotiateXmlSchemaVersion returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasNegotiateXmlSchemaVersion() bool {
	if o != nil && !IsNil(o.NegotiateXmlSchemaVersion) {
		return true
	}

	return false
}

// SetNegotiateXmlSchemaVersion gets a reference to the given bool and assigns it to the NegotiateXmlSchemaVersion field.
func (o *InterfaceSetupType) SetNegotiateXmlSchemaVersion(v bool) {
	o.NegotiateXmlSchemaVersion = &v
}

// GetSetAllToVersion returns the SetAllToVersion field value if set, zero value otherwise.
func (o *InterfaceSetupType) GetSetAllToVersion() string {
	if o == nil || IsNil(o.SetAllToVersion) {
		var ret string
		return ret
	}
	return *o.SetAllToVersion
}

// GetSetAllToVersionOk returns a tuple with the SetAllToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceSetupType) GetSetAllToVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SetAllToVersion) {
		return nil, false
	}
	return o.SetAllToVersion, true
}

// HasSetAllToVersion returns a boolean if a field has been set.
func (o *InterfaceSetupType) HasSetAllToVersion() bool {
	if o != nil && !IsNil(o.SetAllToVersion) {
		return true
	}

	return false
}

// SetSetAllToVersion gets a reference to the given string and assigns it to the SetAllToVersion field.
func (o *InterfaceSetupType) SetSetAllToVersion(v string) {
	o.SetAllToVersion = &v
}

func (o InterfaceSetupType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceSetupType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ExternalHotelCode) {
		toSerialize["externalHotelCode"] = o.ExternalHotelCode
	}
	if !IsNil(o.Machine) {
		toSerialize["machine"] = o.Machine
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MessageFormat) {
		toSerialize["messageFormat"] = o.MessageFormat
	}
	if !IsNil(o.XMLVersions) {
		toSerialize["xMLVersions"] = o.XMLVersions
	}
	if !IsNil(o.DatabaseId) {
		toSerialize["databaseId"] = o.DatabaseId
	}
	if !IsNil(o.InterfaceType) {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	if !IsNil(o.ExternalSystemActivated) {
		toSerialize["externalSystemActivated"] = o.ExternalSystemActivated
	}
	if !IsNil(o.SystemType) {
		toSerialize["systemType"] = o.SystemType
	}
	if !IsNil(o.BatchProcessBE) {
		toSerialize["batchProcessBE"] = o.BatchProcessBE
	}
	if !IsNil(o.ZipData) {
		toSerialize["zipData"] = o.ZipData
	}
	if !IsNil(o.DeltaMode) {
		toSerialize["deltaMode"] = o.DeltaMode
	}
	if !IsNil(o.DeletionIndicator) {
		toSerialize["deletionIndicator"] = o.DeletionIndicator
	}
	if !IsNil(o.KeepingIndicator) {
		toSerialize["keepingIndicator"] = o.KeepingIndicator
	}
	if !IsNil(o.OrsDestination) {
		toSerialize["orsDestination"] = o.OrsDestination
	}
	if !IsNil(o.XmlSchemaVersionsNegotiable) {
		toSerialize["xmlSchemaVersionsNegotiable"] = o.XmlSchemaVersionsNegotiable
	}
	if !IsNil(o.NegotiateXmlSchemaVersion) {
		toSerialize["negotiateXmlSchemaVersion"] = o.NegotiateXmlSchemaVersion
	}
	if !IsNil(o.SetAllToVersion) {
		toSerialize["setAllToVersion"] = o.SetAllToVersion
	}
	return toSerialize, nil
}

type NullableInterfaceSetupType struct {
	value *InterfaceSetupType
	isSet bool
}

func (v NullableInterfaceSetupType) Get() *InterfaceSetupType {
	return v.value
}

func (v *NullableInterfaceSetupType) Set(val *InterfaceSetupType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceSetupType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceSetupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceSetupType(val *InterfaceSetupType) *NullableInterfaceSetupType {
	return &NullableInterfaceSetupType{value: val, isSet: true}
}

func (v NullableInterfaceSetupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceSetupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


