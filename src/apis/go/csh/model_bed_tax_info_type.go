/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BedTaxInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BedTaxInfoType{}

// BedTaxInfoType Bed Trx Info type
type BedTaxInfoType struct {
	// Property associated with bed tax
	HotelId *string `json:"hotelId,omitempty"`
	// Tax Registration Number
	TaxRegistrationNo *int32 `json:"taxRegistrationNo,omitempty"`
	// Arrival Date with time of the reservation
	TaxArrivaldate *string `json:"taxArrivaldate,omitempty"`
	// Arrival Carrier code of the mode of transportation (Airline No. or Railway No., etc.)
	ArrivalCarrierCode *string `json:"arrivalCarrierCode,omitempty"`
	// Departure Date with time of the reservation
	TaxDepartureDate *string `json:"taxDepartureDate,omitempty"`
	// Departure Carrier code of the mode of transportation (Airline No. or Railway No., etc.)
	DepartureCarrierCode *string `json:"departureCarrierCode,omitempty"`
	// Guest name associated with the transaction.
	GuestName *string `json:"guestName,omitempty"`
	ReservationNameId *UniqueIDType `json:"reservationNameId,omitempty"`
	GuestNameId *UniqueIDType `json:"guestNameId,omitempty"`
	// Arrival Date of the reservation
	Arrival *string `json:"arrival,omitempty"`
	// Departure Date with time of the reservation
	Departure *string `json:"departure,omitempty"`
	// Visa Number
	VisaNumber *string `json:"visaNumber,omitempty"`
	// Issue date of the Visa
	VisaIssueDate *string `json:"visaIssueDate,omitempty"`
	// Expiration date of visa
	VisaExpirationDate *string `json:"visaExpirationDate,omitempty"`
	// Total number of tax stays at the property
	TaxNoOfStays *int32 `json:"taxNoOfStays,omitempty"`
	// Profession of the guest
	Profession *string `json:"profession,omitempty"`
	// Passport number
	PassPortStr *string `json:"passPortStr,omitempty"`
	// Passport number
	Passport *string `json:"passport,omitempty"`
	// Nationality of the guest
	Nationality *string `json:"nationality,omitempty"`
	// Age of the guest
	GuestAge *int32 `json:"guestAge,omitempty"`
	// Country of the guest
	Country *string `json:"country,omitempty"`
	// User-defined remark.
	Remark *string `json:"remark,omitempty"`
}

// NewBedTaxInfoType instantiates a new BedTaxInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBedTaxInfoType() *BedTaxInfoType {
	this := BedTaxInfoType{}
	return &this
}

// NewBedTaxInfoTypeWithDefaults instantiates a new BedTaxInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBedTaxInfoTypeWithDefaults() *BedTaxInfoType {
	this := BedTaxInfoType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *BedTaxInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetTaxRegistrationNo returns the TaxRegistrationNo field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetTaxRegistrationNo() int32 {
	if o == nil || IsNil(o.TaxRegistrationNo) {
		var ret int32
		return ret
	}
	return *o.TaxRegistrationNo
}

// GetTaxRegistrationNoOk returns a tuple with the TaxRegistrationNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetTaxRegistrationNoOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxRegistrationNo) {
		return nil, false
	}
	return o.TaxRegistrationNo, true
}

// HasTaxRegistrationNo returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasTaxRegistrationNo() bool {
	if o != nil && !IsNil(o.TaxRegistrationNo) {
		return true
	}

	return false
}

// SetTaxRegistrationNo gets a reference to the given int32 and assigns it to the TaxRegistrationNo field.
func (o *BedTaxInfoType) SetTaxRegistrationNo(v int32) {
	o.TaxRegistrationNo = &v
}

// GetTaxArrivaldate returns the TaxArrivaldate field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetTaxArrivaldate() string {
	if o == nil || IsNil(o.TaxArrivaldate) {
		var ret string
		return ret
	}
	return *o.TaxArrivaldate
}

// GetTaxArrivaldateOk returns a tuple with the TaxArrivaldate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetTaxArrivaldateOk() (*string, bool) {
	if o == nil || IsNil(o.TaxArrivaldate) {
		return nil, false
	}
	return o.TaxArrivaldate, true
}

// HasTaxArrivaldate returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasTaxArrivaldate() bool {
	if o != nil && !IsNil(o.TaxArrivaldate) {
		return true
	}

	return false
}

// SetTaxArrivaldate gets a reference to the given string and assigns it to the TaxArrivaldate field.
func (o *BedTaxInfoType) SetTaxArrivaldate(v string) {
	o.TaxArrivaldate = &v
}

// GetArrivalCarrierCode returns the ArrivalCarrierCode field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetArrivalCarrierCode() string {
	if o == nil || IsNil(o.ArrivalCarrierCode) {
		var ret string
		return ret
	}
	return *o.ArrivalCarrierCode
}

// GetArrivalCarrierCodeOk returns a tuple with the ArrivalCarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetArrivalCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ArrivalCarrierCode) {
		return nil, false
	}
	return o.ArrivalCarrierCode, true
}

// HasArrivalCarrierCode returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasArrivalCarrierCode() bool {
	if o != nil && !IsNil(o.ArrivalCarrierCode) {
		return true
	}

	return false
}

// SetArrivalCarrierCode gets a reference to the given string and assigns it to the ArrivalCarrierCode field.
func (o *BedTaxInfoType) SetArrivalCarrierCode(v string) {
	o.ArrivalCarrierCode = &v
}

// GetTaxDepartureDate returns the TaxDepartureDate field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetTaxDepartureDate() string {
	if o == nil || IsNil(o.TaxDepartureDate) {
		var ret string
		return ret
	}
	return *o.TaxDepartureDate
}

// GetTaxDepartureDateOk returns a tuple with the TaxDepartureDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetTaxDepartureDateOk() (*string, bool) {
	if o == nil || IsNil(o.TaxDepartureDate) {
		return nil, false
	}
	return o.TaxDepartureDate, true
}

// HasTaxDepartureDate returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasTaxDepartureDate() bool {
	if o != nil && !IsNil(o.TaxDepartureDate) {
		return true
	}

	return false
}

// SetTaxDepartureDate gets a reference to the given string and assigns it to the TaxDepartureDate field.
func (o *BedTaxInfoType) SetTaxDepartureDate(v string) {
	o.TaxDepartureDate = &v
}

// GetDepartureCarrierCode returns the DepartureCarrierCode field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetDepartureCarrierCode() string {
	if o == nil || IsNil(o.DepartureCarrierCode) {
		var ret string
		return ret
	}
	return *o.DepartureCarrierCode
}

// GetDepartureCarrierCodeOk returns a tuple with the DepartureCarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetDepartureCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DepartureCarrierCode) {
		return nil, false
	}
	return o.DepartureCarrierCode, true
}

// HasDepartureCarrierCode returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasDepartureCarrierCode() bool {
	if o != nil && !IsNil(o.DepartureCarrierCode) {
		return true
	}

	return false
}

// SetDepartureCarrierCode gets a reference to the given string and assigns it to the DepartureCarrierCode field.
func (o *BedTaxInfoType) SetDepartureCarrierCode(v string) {
	o.DepartureCarrierCode = &v
}

// GetGuestName returns the GuestName field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetGuestName() string {
	if o == nil || IsNil(o.GuestName) {
		var ret string
		return ret
	}
	return *o.GuestName
}

// GetGuestNameOk returns a tuple with the GuestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetGuestNameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestName) {
		return nil, false
	}
	return o.GuestName, true
}

// HasGuestName returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasGuestName() bool {
	if o != nil && !IsNil(o.GuestName) {
		return true
	}

	return false
}

// SetGuestName gets a reference to the given string and assigns it to the GuestName field.
func (o *BedTaxInfoType) SetGuestName(v string) {
	o.GuestName = &v
}

// GetReservationNameId returns the ReservationNameId field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetReservationNameId() UniqueIDType {
	if o == nil || IsNil(o.ReservationNameId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ReservationNameId
}

// GetReservationNameIdOk returns a tuple with the ReservationNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetReservationNameIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationNameId) {
		return nil, false
	}
	return o.ReservationNameId, true
}

// HasReservationNameId returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasReservationNameId() bool {
	if o != nil && !IsNil(o.ReservationNameId) {
		return true
	}

	return false
}

// SetReservationNameId gets a reference to the given UniqueIDType and assigns it to the ReservationNameId field.
func (o *BedTaxInfoType) SetReservationNameId(v UniqueIDType) {
	o.ReservationNameId = &v
}

// GetGuestNameId returns the GuestNameId field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetGuestNameId() UniqueIDType {
	if o == nil || IsNil(o.GuestNameId) {
		var ret UniqueIDType
		return ret
	}
	return *o.GuestNameId
}

// GetGuestNameIdOk returns a tuple with the GuestNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetGuestNameIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.GuestNameId) {
		return nil, false
	}
	return o.GuestNameId, true
}

// HasGuestNameId returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasGuestNameId() bool {
	if o != nil && !IsNil(o.GuestNameId) {
		return true
	}

	return false
}

// SetGuestNameId gets a reference to the given UniqueIDType and assigns it to the GuestNameId field.
func (o *BedTaxInfoType) SetGuestNameId(v UniqueIDType) {
	o.GuestNameId = &v
}

// GetArrival returns the Arrival field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetArrival() string {
	if o == nil || IsNil(o.Arrival) {
		var ret string
		return ret
	}
	return *o.Arrival
}

// GetArrivalOk returns a tuple with the Arrival field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetArrivalOk() (*string, bool) {
	if o == nil || IsNil(o.Arrival) {
		return nil, false
	}
	return o.Arrival, true
}

// HasArrival returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasArrival() bool {
	if o != nil && !IsNil(o.Arrival) {
		return true
	}

	return false
}

// SetArrival gets a reference to the given string and assigns it to the Arrival field.
func (o *BedTaxInfoType) SetArrival(v string) {
	o.Arrival = &v
}

// GetDeparture returns the Departure field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetDeparture() string {
	if o == nil || IsNil(o.Departure) {
		var ret string
		return ret
	}
	return *o.Departure
}

// GetDepartureOk returns a tuple with the Departure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetDepartureOk() (*string, bool) {
	if o == nil || IsNil(o.Departure) {
		return nil, false
	}
	return o.Departure, true
}

// HasDeparture returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasDeparture() bool {
	if o != nil && !IsNil(o.Departure) {
		return true
	}

	return false
}

// SetDeparture gets a reference to the given string and assigns it to the Departure field.
func (o *BedTaxInfoType) SetDeparture(v string) {
	o.Departure = &v
}

// GetVisaNumber returns the VisaNumber field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetVisaNumber() string {
	if o == nil || IsNil(o.VisaNumber) {
		var ret string
		return ret
	}
	return *o.VisaNumber
}

// GetVisaNumberOk returns a tuple with the VisaNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetVisaNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VisaNumber) {
		return nil, false
	}
	return o.VisaNumber, true
}

// HasVisaNumber returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasVisaNumber() bool {
	if o != nil && !IsNil(o.VisaNumber) {
		return true
	}

	return false
}

// SetVisaNumber gets a reference to the given string and assigns it to the VisaNumber field.
func (o *BedTaxInfoType) SetVisaNumber(v string) {
	o.VisaNumber = &v
}

// GetVisaIssueDate returns the VisaIssueDate field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetVisaIssueDate() string {
	if o == nil || IsNil(o.VisaIssueDate) {
		var ret string
		return ret
	}
	return *o.VisaIssueDate
}

// GetVisaIssueDateOk returns a tuple with the VisaIssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetVisaIssueDateOk() (*string, bool) {
	if o == nil || IsNil(o.VisaIssueDate) {
		return nil, false
	}
	return o.VisaIssueDate, true
}

// HasVisaIssueDate returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasVisaIssueDate() bool {
	if o != nil && !IsNil(o.VisaIssueDate) {
		return true
	}

	return false
}

// SetVisaIssueDate gets a reference to the given string and assigns it to the VisaIssueDate field.
func (o *BedTaxInfoType) SetVisaIssueDate(v string) {
	o.VisaIssueDate = &v
}

// GetVisaExpirationDate returns the VisaExpirationDate field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetVisaExpirationDate() string {
	if o == nil || IsNil(o.VisaExpirationDate) {
		var ret string
		return ret
	}
	return *o.VisaExpirationDate
}

// GetVisaExpirationDateOk returns a tuple with the VisaExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetVisaExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.VisaExpirationDate) {
		return nil, false
	}
	return o.VisaExpirationDate, true
}

// HasVisaExpirationDate returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasVisaExpirationDate() bool {
	if o != nil && !IsNil(o.VisaExpirationDate) {
		return true
	}

	return false
}

// SetVisaExpirationDate gets a reference to the given string and assigns it to the VisaExpirationDate field.
func (o *BedTaxInfoType) SetVisaExpirationDate(v string) {
	o.VisaExpirationDate = &v
}

// GetTaxNoOfStays returns the TaxNoOfStays field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetTaxNoOfStays() int32 {
	if o == nil || IsNil(o.TaxNoOfStays) {
		var ret int32
		return ret
	}
	return *o.TaxNoOfStays
}

// GetTaxNoOfStaysOk returns a tuple with the TaxNoOfStays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetTaxNoOfStaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxNoOfStays) {
		return nil, false
	}
	return o.TaxNoOfStays, true
}

// HasTaxNoOfStays returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasTaxNoOfStays() bool {
	if o != nil && !IsNil(o.TaxNoOfStays) {
		return true
	}

	return false
}

// SetTaxNoOfStays gets a reference to the given int32 and assigns it to the TaxNoOfStays field.
func (o *BedTaxInfoType) SetTaxNoOfStays(v int32) {
	o.TaxNoOfStays = &v
}

// GetProfession returns the Profession field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetProfession() string {
	if o == nil || IsNil(o.Profession) {
		var ret string
		return ret
	}
	return *o.Profession
}

// GetProfessionOk returns a tuple with the Profession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetProfessionOk() (*string, bool) {
	if o == nil || IsNil(o.Profession) {
		return nil, false
	}
	return o.Profession, true
}

// HasProfession returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasProfession() bool {
	if o != nil && !IsNil(o.Profession) {
		return true
	}

	return false
}

// SetProfession gets a reference to the given string and assigns it to the Profession field.
func (o *BedTaxInfoType) SetProfession(v string) {
	o.Profession = &v
}

// GetPassPortStr returns the PassPortStr field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetPassPortStr() string {
	if o == nil || IsNil(o.PassPortStr) {
		var ret string
		return ret
	}
	return *o.PassPortStr
}

// GetPassPortStrOk returns a tuple with the PassPortStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetPassPortStrOk() (*string, bool) {
	if o == nil || IsNil(o.PassPortStr) {
		return nil, false
	}
	return o.PassPortStr, true
}

// HasPassPortStr returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasPassPortStr() bool {
	if o != nil && !IsNil(o.PassPortStr) {
		return true
	}

	return false
}

// SetPassPortStr gets a reference to the given string and assigns it to the PassPortStr field.
func (o *BedTaxInfoType) SetPassPortStr(v string) {
	o.PassPortStr = &v
}

// GetPassport returns the Passport field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetPassport() string {
	if o == nil || IsNil(o.Passport) {
		var ret string
		return ret
	}
	return *o.Passport
}

// GetPassportOk returns a tuple with the Passport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetPassportOk() (*string, bool) {
	if o == nil || IsNil(o.Passport) {
		return nil, false
	}
	return o.Passport, true
}

// HasPassport returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasPassport() bool {
	if o != nil && !IsNil(o.Passport) {
		return true
	}

	return false
}

// SetPassport gets a reference to the given string and assigns it to the Passport field.
func (o *BedTaxInfoType) SetPassport(v string) {
	o.Passport = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *BedTaxInfoType) SetNationality(v string) {
	o.Nationality = &v
}

// GetGuestAge returns the GuestAge field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetGuestAge() int32 {
	if o == nil || IsNil(o.GuestAge) {
		var ret int32
		return ret
	}
	return *o.GuestAge
}

// GetGuestAgeOk returns a tuple with the GuestAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetGuestAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.GuestAge) {
		return nil, false
	}
	return o.GuestAge, true
}

// HasGuestAge returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasGuestAge() bool {
	if o != nil && !IsNil(o.GuestAge) {
		return true
	}

	return false
}

// SetGuestAge gets a reference to the given int32 and assigns it to the GuestAge field.
func (o *BedTaxInfoType) SetGuestAge(v int32) {
	o.GuestAge = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *BedTaxInfoType) SetCountry(v string) {
	o.Country = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *BedTaxInfoType) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BedTaxInfoType) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *BedTaxInfoType) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *BedTaxInfoType) SetRemark(v string) {
	o.Remark = &v
}

func (o BedTaxInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BedTaxInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.TaxRegistrationNo) {
		toSerialize["taxRegistrationNo"] = o.TaxRegistrationNo
	}
	if !IsNil(o.TaxArrivaldate) {
		toSerialize["taxArrivaldate"] = o.TaxArrivaldate
	}
	if !IsNil(o.ArrivalCarrierCode) {
		toSerialize["arrivalCarrierCode"] = o.ArrivalCarrierCode
	}
	if !IsNil(o.TaxDepartureDate) {
		toSerialize["taxDepartureDate"] = o.TaxDepartureDate
	}
	if !IsNil(o.DepartureCarrierCode) {
		toSerialize["departureCarrierCode"] = o.DepartureCarrierCode
	}
	if !IsNil(o.GuestName) {
		toSerialize["guestName"] = o.GuestName
	}
	if !IsNil(o.ReservationNameId) {
		toSerialize["reservationNameId"] = o.ReservationNameId
	}
	if !IsNil(o.GuestNameId) {
		toSerialize["guestNameId"] = o.GuestNameId
	}
	if !IsNil(o.Arrival) {
		toSerialize["arrival"] = o.Arrival
	}
	if !IsNil(o.Departure) {
		toSerialize["departure"] = o.Departure
	}
	if !IsNil(o.VisaNumber) {
		toSerialize["visaNumber"] = o.VisaNumber
	}
	if !IsNil(o.VisaIssueDate) {
		toSerialize["visaIssueDate"] = o.VisaIssueDate
	}
	if !IsNil(o.VisaExpirationDate) {
		toSerialize["visaExpirationDate"] = o.VisaExpirationDate
	}
	if !IsNil(o.TaxNoOfStays) {
		toSerialize["taxNoOfStays"] = o.TaxNoOfStays
	}
	if !IsNil(o.Profession) {
		toSerialize["profession"] = o.Profession
	}
	if !IsNil(o.PassPortStr) {
		toSerialize["passPortStr"] = o.PassPortStr
	}
	if !IsNil(o.Passport) {
		toSerialize["passport"] = o.Passport
	}
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.GuestAge) {
		toSerialize["guestAge"] = o.GuestAge
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	return toSerialize, nil
}

type NullableBedTaxInfoType struct {
	value *BedTaxInfoType
	isSet bool
}

func (v NullableBedTaxInfoType) Get() *BedTaxInfoType {
	return v.value
}

func (v *NullableBedTaxInfoType) Set(val *BedTaxInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableBedTaxInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableBedTaxInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBedTaxInfoType(val *BedTaxInfoType) *NullableBedTaxInfoType {
	return &NullableBedTaxInfoType{value: val, isSet: true}
}

func (v NullableBedTaxInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBedTaxInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


