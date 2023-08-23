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
exports.TrackItReservationInfoTypeToJSON = exports.TrackItReservationInfoTypeFromJSONTyped = exports.TrackItReservationInfoTypeFromJSON = exports.instanceOfTrackItReservationInfoType = void 0;
const runtime_1 = require("../runtime");
const HousekeepingRoomStatusType_1 = require("./HousekeepingRoomStatusType");
const PMSResStatusType_1 = require("./PMSResStatusType");
const ResGuaranteeType_1 = require("./ResGuaranteeType");
const ResGuestInfoType_1 = require("./ResGuestInfoType");
const TimeSpanType_1 = require("./TimeSpanType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the TrackItReservationInfoType interface.
 */
function instanceOfTrackItReservationInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTrackItReservationInfoType = instanceOfTrackItReservationInfoType;
function TrackItReservationInfoTypeFromJSON(json) {
    return TrackItReservationInfoTypeFromJSONTyped(json, false);
}
exports.TrackItReservationInfoTypeFromJSON = TrackItReservationInfoTypeFromJSON;
function TrackItReservationInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'timeSpan': !(0, runtime_1.exists)(json, 'timeSpan') ? undefined : (0, TimeSpanType_1.TimeSpanTypeFromJSON)(json['timeSpan']),
        'guestInfo': !(0, runtime_1.exists)(json, 'guestInfo') ? undefined : (0, ResGuestInfoType_1.ResGuestInfoTypeFromJSON)(json['guestInfo']),
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : json['roomType'],
        'roomId': !(0, runtime_1.exists)(json, 'roomId') ? undefined : json['roomId'],
        'roomStatus': !(0, runtime_1.exists)(json, 'roomStatus') ? undefined : (0, HousekeepingRoomStatusType_1.HousekeepingRoomStatusTypeFromJSON)(json['roomStatus']),
        'guarantee': !(0, runtime_1.exists)(json, 'guarantee') ? undefined : (0, ResGuaranteeType_1.ResGuaranteeTypeFromJSON)(json['guarantee']),
        'reservationStatus': !(0, runtime_1.exists)(json, 'reservationStatus') ? undefined : (0, PMSResStatusType_1.PMSResStatusTypeFromJSON)(json['reservationStatus']),
        'computedReservationStatus': !(0, runtime_1.exists)(json, 'computedReservationStatus') ? undefined : (0, PMSResStatusType_1.PMSResStatusTypeFromJSON)(json['computedReservationStatus']),
    };
}
exports.TrackItReservationInfoTypeFromJSONTyped = TrackItReservationInfoTypeFromJSONTyped;
function TrackItReservationInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'timeSpan': (0, TimeSpanType_1.TimeSpanTypeToJSON)(value.timeSpan),
        'guestInfo': (0, ResGuestInfoType_1.ResGuestInfoTypeToJSON)(value.guestInfo),
        'roomType': value.roomType,
        'roomId': value.roomId,
        'roomStatus': (0, HousekeepingRoomStatusType_1.HousekeepingRoomStatusTypeToJSON)(value.roomStatus),
        'guarantee': (0, ResGuaranteeType_1.ResGuaranteeTypeToJSON)(value.guarantee),
        'reservationStatus': (0, PMSResStatusType_1.PMSResStatusTypeToJSON)(value.reservationStatus),
        'computedReservationStatus': (0, PMSResStatusType_1.PMSResStatusTypeToJSON)(value.computedReservationStatus),
    };
}
exports.TrackItReservationInfoTypeToJSON = TrackItReservationInfoTypeToJSON;
