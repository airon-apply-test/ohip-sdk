"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RoomKeyInfoTypeToJSON = exports.RoomKeyInfoTypeFromJSONTyped = exports.RoomKeyInfoTypeFromJSON = exports.instanceOfRoomKeyInfoType = void 0;
const runtime_1 = require("../runtime");
const KeyTrackType_1 = require("./KeyTrackType");
const ReservationId_1 = require("./ReservationId");
/**
 * Check if a given object implements the RoomKeyInfoType interface.
 */
function instanceOfRoomKeyInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomKeyInfoType = instanceOfRoomKeyInfoType;
function RoomKeyInfoTypeFromJSON(json) {
    return RoomKeyInfoTypeFromJSONTyped(json, false);
}
exports.RoomKeyInfoTypeFromJSON = RoomKeyInfoTypeFromJSON;
function RoomKeyInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'keyCount': !(0, runtime_1.exists)(json, 'keyCount') ? undefined : json['keyCount'],
        'keyStartDate': !(0, runtime_1.exists)(json, 'keyStartDate') ? undefined : (new Date(json['keyStartDate'])),
        'keyExpiryDate': !(0, runtime_1.exists)(json, 'keyExpiryDate') ? undefined : (new Date(json['keyExpiryDate'])),
        'keyOptions': !(0, runtime_1.exists)(json, 'keyOptions') ? undefined : json['keyOptions'],
        'additionalRooms': !(0, runtime_1.exists)(json, 'additionalRooms') ? undefined : json['additionalRooms'],
        'roomId': !(0, runtime_1.exists)(json, 'roomId') ? undefined : json['roomId'],
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, ReservationId_1.ReservationIdFromJSON)(json['reservationId']),
        'keyTrack': !(0, runtime_1.exists)(json, 'keyTrack') ? undefined : (0, KeyTrackType_1.KeyTrackTypeFromJSON)(json['keyTrack']),
    };
}
exports.RoomKeyInfoTypeFromJSONTyped = RoomKeyInfoTypeFromJSONTyped;
function RoomKeyInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'keyCount': value.keyCount,
        'keyStartDate': value.keyStartDate === undefined ? undefined : (value.keyStartDate.toISOString()),
        'keyExpiryDate': value.keyExpiryDate === undefined ? undefined : (value.keyExpiryDate.toISOString()),
        'keyOptions': value.keyOptions,
        'additionalRooms': value.additionalRooms,
        'roomId': value.roomId,
        'reservationId': (0, ReservationId_1.ReservationIdToJSON)(value.reservationId),
        'keyTrack': (0, KeyTrackType_1.KeyTrackTypeToJSON)(value.keyTrack),
    };
}
exports.RoomKeyInfoTypeToJSON = RoomKeyInfoTypeToJSON;
