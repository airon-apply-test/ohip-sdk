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
exports.RoomTypeShortInfoTypeToJSON = exports.RoomTypeShortInfoTypeFromJSONTyped = exports.RoomTypeShortInfoTypeFromJSON = exports.instanceOfRoomTypeShortInfoType = void 0;
const runtime_1 = require("../runtime");
const RatePlanRatingType_1 = require("./RatePlanRatingType");
const RoomFeatureType_1 = require("./RoomFeatureType");
/**
 * Check if a given object implements the RoomTypeShortInfoType interface.
 */
function instanceOfRoomTypeShortInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomTypeShortInfoType = instanceOfRoomTypeShortInfoType;
function RoomTypeShortInfoTypeFromJSON(json) {
    return RoomTypeShortInfoTypeFromJSONTyped(json, false);
}
exports.RoomTypeShortInfoTypeFromJSON = RoomTypeShortInfoTypeFromJSON;
function RoomTypeShortInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'pseudo': !(0, runtime_1.exists)(json, 'pseudo') ? undefined : json['pseudo'],
        'suite': !(0, runtime_1.exists)(json, 'suite') ? undefined : json['suite'],
        'roomClass': !(0, runtime_1.exists)(json, 'roomClass') ? undefined : json['roomClass'],
        'shortDescription': !(0, runtime_1.exists)(json, 'shortDescription') ? undefined : json['shortDescription'],
        'houseKeeping': !(0, runtime_1.exists)(json, 'houseKeeping') ? undefined : json['houseKeeping'],
        'smokingPreference': !(0, runtime_1.exists)(json, 'smokingPreference') ? undefined : json['smokingPreference'],
        'building': !(0, runtime_1.exists)(json, 'building') ? undefined : json['building'],
        'roomAssignmentRating': !(0, runtime_1.exists)(json, 'roomAssignmentRating') ? undefined : (0, RatePlanRatingType_1.RatePlanRatingTypeFromJSON)(json['roomAssignmentRating']),
        'minimumOccupancy': !(0, runtime_1.exists)(json, 'minimumOccupancy') ? undefined : json['minimumOccupancy'],
        'maximumOccupancy': !(0, runtime_1.exists)(json, 'maximumOccupancy') ? undefined : json['maximumOccupancy'],
        'roomFeatures': !(0, runtime_1.exists)(json, 'roomFeatures') ? undefined : (json['roomFeatures'].map(RoomFeatureType_1.RoomFeatureTypeFromJSON)),
        'accessible': !(0, runtime_1.exists)(json, 'accessible') ? undefined : json['accessible'],
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : json['roomType'],
        'meetingRoom': !(0, runtime_1.exists)(json, 'meetingRoom') ? undefined : json['meetingRoom'],
    };
}
exports.RoomTypeShortInfoTypeFromJSONTyped = RoomTypeShortInfoTypeFromJSONTyped;
function RoomTypeShortInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'pseudo': value.pseudo,
        'suite': value.suite,
        'roomClass': value.roomClass,
        'shortDescription': value.shortDescription,
        'houseKeeping': value.houseKeeping,
        'smokingPreference': value.smokingPreference,
        'building': value.building,
        'roomAssignmentRating': (0, RatePlanRatingType_1.RatePlanRatingTypeToJSON)(value.roomAssignmentRating),
        'minimumOccupancy': value.minimumOccupancy,
        'maximumOccupancy': value.maximumOccupancy,
        'roomFeatures': value.roomFeatures === undefined ? undefined : (value.roomFeatures.map(RoomFeatureType_1.RoomFeatureTypeToJSON)),
        'accessible': value.accessible,
        'roomType': value.roomType,
        'meetingRoom': value.meetingRoom,
    };
}
exports.RoomTypeShortInfoTypeToJSON = RoomTypeShortInfoTypeToJSON;
