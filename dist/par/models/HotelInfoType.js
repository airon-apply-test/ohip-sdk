"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HotelInfoTypeToJSON = exports.HotelInfoTypeFromJSONTyped = exports.HotelInfoTypeFromJSON = exports.instanceOfHotelInfoType = void 0;
const runtime_1 = require("../runtime");
const AlternateHotelType_1 = require("./AlternateHotelType");
const CommentInfoType_1 = require("./CommentInfoType");
const HotelAttractionType_1 = require("./HotelAttractionType");
const HotelContactType_1 = require("./HotelContactType");
const HotelCorporateInformationsType_1 = require("./HotelCorporateInformationsType");
const HotelEventSpacesType_1 = require("./HotelEventSpacesType");
const HotelInfoTypeAccommodationDetails_1 = require("./HotelInfoTypeAccommodationDetails");
const HotelInfoTypeAddress_1 = require("./HotelInfoTypeAddress");
const HotelInfoTypeCommunication_1 = require("./HotelInfoTypeCommunication");
const HotelInfoTypeGeneralInformation_1 = require("./HotelInfoTypeGeneralInformation");
const HotelInfoTypePrimaryDetails_1 = require("./HotelInfoTypePrimaryDetails");
const HotelInfoTypePropertyControls_1 = require("./HotelInfoTypePropertyControls");
const HotelRateRangeType_1 = require("./HotelRateRangeType");
const HotelRestaurantType_1 = require("./HotelRestaurantType");
const MeetingRoomType_1 = require("./MeetingRoomType");
/**
 * Check if a given object implements the HotelInfoType interface.
 */
function instanceOfHotelInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelInfoType = instanceOfHotelInfoType;
function HotelInfoTypeFromJSON(json) {
    return HotelInfoTypeFromJSONTyped(json, false);
}
exports.HotelInfoTypeFromJSON = HotelInfoTypeFromJSON;
function HotelInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'primaryDetails': !(0, runtime_1.exists)(json, 'primaryDetails') ? undefined : (0, HotelInfoTypePrimaryDetails_1.HotelInfoTypePrimaryDetailsFromJSON)(json['primaryDetails']),
        'generalInformation': !(0, runtime_1.exists)(json, 'generalInformation') ? undefined : (0, HotelInfoTypeGeneralInformation_1.HotelInfoTypeGeneralInformationFromJSON)(json['generalInformation']),
        'accommodationDetails': !(0, runtime_1.exists)(json, 'accommodationDetails') ? undefined : (0, HotelInfoTypeAccommodationDetails_1.HotelInfoTypeAccommodationDetailsFromJSON)(json['accommodationDetails']),
        'propertyControls': !(0, runtime_1.exists)(json, 'propertyControls') ? undefined : (0, HotelInfoTypePropertyControls_1.HotelInfoTypePropertyControlsFromJSON)(json['propertyControls']),
        'communication': !(0, runtime_1.exists)(json, 'communication') ? undefined : (0, HotelInfoTypeCommunication_1.HotelInfoTypeCommunicationFromJSON)(json['communication']),
        'address': !(0, runtime_1.exists)(json, 'address') ? undefined : (0, HotelInfoTypeAddress_1.HotelInfoTypeAddressFromJSON)(json['address']),
        'hotelRestaurants': !(0, runtime_1.exists)(json, 'hotelRestaurants') ? undefined : (json['hotelRestaurants'].map(HotelRestaurantType_1.HotelRestaurantTypeFromJSON)),
        'hotelRateRanges': !(0, runtime_1.exists)(json, 'hotelRateRanges') ? undefined : (json['hotelRateRanges'].map(HotelRateRangeType_1.HotelRateRangeTypeFromJSON)),
        'alternateHotels': !(0, runtime_1.exists)(json, 'alternateHotels') ? undefined : (json['alternateHotels'].map(AlternateHotelType_1.AlternateHotelTypeFromJSON)),
        'hotelContacts': !(0, runtime_1.exists)(json, 'hotelContacts') ? undefined : (json['hotelContacts'].map(HotelContactType_1.HotelContactTypeFromJSON)),
        'hotelEventSpaces': !(0, runtime_1.exists)(json, 'hotelEventSpaces') ? undefined : (0, HotelEventSpacesType_1.HotelEventSpacesTypeFromJSON)(json['hotelEventSpaces']),
        'hotelNotes': !(0, runtime_1.exists)(json, 'hotelNotes') ? undefined : (json['hotelNotes'].map(CommentInfoType_1.CommentInfoTypeFromJSON)),
        'hotelCorporateInformations': !(0, runtime_1.exists)(json, 'hotelCorporateInformations') ? undefined : (0, HotelCorporateInformationsType_1.HotelCorporateInformationsTypeFromJSON)(json['hotelCorporateInformations']),
        'attractions': !(0, runtime_1.exists)(json, 'attractions') ? undefined : (json['attractions'].map(HotelAttractionType_1.HotelAttractionTypeFromJSON)),
        'meetingRooms': !(0, runtime_1.exists)(json, 'meetingRooms') ? undefined : (json['meetingRooms'].map(MeetingRoomType_1.MeetingRoomTypeFromJSON)),
        'chainCode': !(0, runtime_1.exists)(json, 'chainCode') ? undefined : json['chainCode'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'hotelCityCode': !(0, runtime_1.exists)(json, 'hotelCityCode') ? undefined : json['hotelCityCode'],
        'hotelName': !(0, runtime_1.exists)(json, 'hotelName') ? undefined : json['hotelName'],
        'hotelCodeContext': !(0, runtime_1.exists)(json, 'hotelCodeContext') ? undefined : json['hotelCodeContext'],
        'chainName': !(0, runtime_1.exists)(json, 'chainName') ? undefined : json['chainName'],
    };
}
exports.HotelInfoTypeFromJSONTyped = HotelInfoTypeFromJSONTyped;
function HotelInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'primaryDetails': (0, HotelInfoTypePrimaryDetails_1.HotelInfoTypePrimaryDetailsToJSON)(value.primaryDetails),
        'generalInformation': (0, HotelInfoTypeGeneralInformation_1.HotelInfoTypeGeneralInformationToJSON)(value.generalInformation),
        'accommodationDetails': (0, HotelInfoTypeAccommodationDetails_1.HotelInfoTypeAccommodationDetailsToJSON)(value.accommodationDetails),
        'propertyControls': (0, HotelInfoTypePropertyControls_1.HotelInfoTypePropertyControlsToJSON)(value.propertyControls),
        'communication': (0, HotelInfoTypeCommunication_1.HotelInfoTypeCommunicationToJSON)(value.communication),
        'address': (0, HotelInfoTypeAddress_1.HotelInfoTypeAddressToJSON)(value.address),
        'hotelRestaurants': value.hotelRestaurants === undefined ? undefined : (value.hotelRestaurants.map(HotelRestaurantType_1.HotelRestaurantTypeToJSON)),
        'hotelRateRanges': value.hotelRateRanges === undefined ? undefined : (value.hotelRateRanges.map(HotelRateRangeType_1.HotelRateRangeTypeToJSON)),
        'alternateHotels': value.alternateHotels === undefined ? undefined : (value.alternateHotels.map(AlternateHotelType_1.AlternateHotelTypeToJSON)),
        'hotelContacts': value.hotelContacts === undefined ? undefined : (value.hotelContacts.map(HotelContactType_1.HotelContactTypeToJSON)),
        'hotelEventSpaces': (0, HotelEventSpacesType_1.HotelEventSpacesTypeToJSON)(value.hotelEventSpaces),
        'hotelNotes': value.hotelNotes === undefined ? undefined : (value.hotelNotes.map(CommentInfoType_1.CommentInfoTypeToJSON)),
        'hotelCorporateInformations': (0, HotelCorporateInformationsType_1.HotelCorporateInformationsTypeToJSON)(value.hotelCorporateInformations),
        'attractions': value.attractions === undefined ? undefined : (value.attractions.map(HotelAttractionType_1.HotelAttractionTypeToJSON)),
        'meetingRooms': value.meetingRooms === undefined ? undefined : (value.meetingRooms.map(MeetingRoomType_1.MeetingRoomTypeToJSON)),
        'chainCode': value.chainCode,
        'hotelId': value.hotelId,
        'hotelCityCode': value.hotelCityCode,
        'hotelName': value.hotelName,
        'hotelCodeContext': value.hotelCodeContext,
        'chainName': value.chainName,
    };
}
exports.HotelInfoTypeToJSON = HotelInfoTypeToJSON;
