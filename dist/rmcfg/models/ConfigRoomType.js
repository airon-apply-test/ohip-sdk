"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConfigRoomTypeToJSON = exports.ConfigRoomTypeFromJSONTyped = exports.ConfigRoomTypeFromJSON = exports.instanceOfConfigRoomType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const HousekeepingCreditsType_1 = require("./HousekeepingCreditsType");
const RatePlanRatingType_1 = require("./RatePlanRatingType");
const RoomComponentType_1 = require("./RoomComponentType");
const RoomFeatureType_1 = require("./RoomFeatureType");
const RoomRoomType_1 = require("./RoomRoomType");
const RoomSectionType_1 = require("./RoomSectionType");
const RoomTypeShortInfoType_1 = require("./RoomTypeShortInfoType");
const TranslationTextType2000_1 = require("./TranslationTextType2000");
/**
 * Check if a given object implements the ConfigRoomType interface.
 */
function instanceOfConfigRoomType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfConfigRoomType = instanceOfConfigRoomType;
function ConfigRoomTypeFromJSON(json) {
    return ConfigRoomTypeFromJSONTyped(json, false);
}
exports.ConfigRoomTypeFromJSON = ConfigRoomTypeFromJSON;
function ConfigRoomTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : (0, RoomTypeShortInfoType_1.RoomTypeShortInfoTypeFromJSON)(json['roomType']),
        'floor': !(0, runtime_1.exists)(json, 'floor') ? undefined : json['floor'],
        'floorDescription': !(0, runtime_1.exists)(json, 'floorDescription') ? undefined : json['floorDescription'],
        'roomFeatures': !(0, runtime_1.exists)(json, 'roomFeatures') ? undefined : (json['roomFeatures'].map(RoomFeatureType_1.RoomFeatureTypeFromJSON)),
        'roomDescription': !(0, runtime_1.exists)(json, 'roomDescription') ? undefined : json['roomDescription'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : (0, TranslationTextType2000_1.TranslationTextType2000FromJSON)(json['description']),
        'smokingPreference': !(0, runtime_1.exists)(json, 'smokingPreference') ? undefined : json['smokingPreference'],
        'smokingPreferenceDescription': !(0, runtime_1.exists)(json, 'smokingPreferenceDescription') ? undefined : json['smokingPreferenceDescription'],
        'building': !(0, runtime_1.exists)(json, 'building') ? undefined : json['building'],
        'roomAssignmentRating': !(0, runtime_1.exists)(json, 'roomAssignmentRating') ? undefined : (0, RatePlanRatingType_1.RatePlanRatingTypeFromJSON)(json['roomAssignmentRating']),
        'accessible': !(0, runtime_1.exists)(json, 'accessible') ? undefined : json['accessible'],
        'roomId': !(0, runtime_1.exists)(json, 'roomId') ? undefined : json['roomId'],
        'meetingRoom': !(0, runtime_1.exists)(json, 'meetingRoom') ? undefined : json['meetingRoom'],
        'roomComponents': !(0, runtime_1.exists)(json, 'roomComponents') ? undefined : (json['roomComponents'].map(RoomComponentType_1.RoomComponentTypeFromJSON)),
        'connectingRooms': !(0, runtime_1.exists)(json, 'connectingRooms') ? undefined : (json['connectingRooms'].map(RoomRoomType_1.RoomRoomTypeFromJSON)),
        'rateCode': !(0, runtime_1.exists)(json, 'rateCode') ? undefined : json['rateCode'],
        'rateAmount': !(0, runtime_1.exists)(json, 'rateAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['rateAmount']),
        'maximumOccupancy': !(0, runtime_1.exists)(json, 'maximumOccupancy') ? undefined : json['maximumOccupancy'],
        'sellSequence': !(0, runtime_1.exists)(json, 'sellSequence') ? undefined : json['sellSequence'],
        'ownerRoom': !(0, runtime_1.exists)(json, 'ownerRoom') ? undefined : json['ownerRoom'],
        'unitGradeCode': !(0, runtime_1.exists)(json, 'unitGradeCode') ? undefined : json['unitGradeCode'],
        'smoking': !(0, runtime_1.exists)(json, 'smoking') ? undefined : json['smoking'],
        'keyCode': !(0, runtime_1.exists)(json, 'keyCode') ? undefined : json['keyCode'],
        'keyOptions': !(0, runtime_1.exists)(json, 'keyOptions') ? undefined : json['keyOptions'],
        'squareUnits': !(0, runtime_1.exists)(json, 'squareUnits') ? undefined : json['squareUnits'],
        'turndownService': !(0, runtime_1.exists)(json, 'turndownService') ? undefined : json['turndownService'],
        'squareUnitsMeasurement': !(0, runtime_1.exists)(json, 'squareUnitsMeasurement') ? undefined : json['squareUnitsMeasurement'],
        'phoneNumber': !(0, runtime_1.exists)(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'housekeepingCredit': !(0, runtime_1.exists)(json, 'housekeepingCredit') ? undefined : (json['housekeepingCredit'].map(HousekeepingCreditsType_1.HousekeepingCreditsTypeFromJSON)),
        'roomSection': !(0, runtime_1.exists)(json, 'roomSection') ? undefined : (0, RoomSectionType_1.RoomSectionTypeFromJSON)(json['roomSection']),
        'noOfBeds': !(0, runtime_1.exists)(json, 'noOfBeds') ? undefined : json['noOfBeds'],
    };
}
exports.ConfigRoomTypeFromJSONTyped = ConfigRoomTypeFromJSONTyped;
function ConfigRoomTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomType': (0, RoomTypeShortInfoType_1.RoomTypeShortInfoTypeToJSON)(value.roomType),
        'floor': value.floor,
        'floorDescription': value.floorDescription,
        'roomFeatures': value.roomFeatures === undefined ? undefined : (value.roomFeatures.map(RoomFeatureType_1.RoomFeatureTypeToJSON)),
        'roomDescription': value.roomDescription,
        'description': (0, TranslationTextType2000_1.TranslationTextType2000ToJSON)(value.description),
        'smokingPreference': value.smokingPreference,
        'smokingPreferenceDescription': value.smokingPreferenceDescription,
        'building': value.building,
        'roomAssignmentRating': (0, RatePlanRatingType_1.RatePlanRatingTypeToJSON)(value.roomAssignmentRating),
        'accessible': value.accessible,
        'roomId': value.roomId,
        'meetingRoom': value.meetingRoom,
        'roomComponents': value.roomComponents === undefined ? undefined : (value.roomComponents.map(RoomComponentType_1.RoomComponentTypeToJSON)),
        'connectingRooms': value.connectingRooms === undefined ? undefined : (value.connectingRooms.map(RoomRoomType_1.RoomRoomTypeToJSON)),
        'rateCode': value.rateCode,
        'rateAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.rateAmount),
        'maximumOccupancy': value.maximumOccupancy,
        'sellSequence': value.sellSequence,
        'ownerRoom': value.ownerRoom,
        'unitGradeCode': value.unitGradeCode,
        'smoking': value.smoking,
        'keyCode': value.keyCode,
        'keyOptions': value.keyOptions,
        'squareUnits': value.squareUnits,
        'turndownService': value.turndownService,
        'squareUnitsMeasurement': value.squareUnitsMeasurement,
        'phoneNumber': value.phoneNumber,
        'housekeepingCredit': value.housekeepingCredit === undefined ? undefined : (value.housekeepingCredit.map(HousekeepingCreditsType_1.HousekeepingCreditsTypeToJSON)),
        'roomSection': (0, RoomSectionType_1.RoomSectionTypeToJSON)(value.roomSection),
        'noOfBeds': value.noOfBeds,
    };
}
exports.ConfigRoomTypeToJSON = ConfigRoomTypeToJSON;
