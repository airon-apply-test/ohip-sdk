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
exports.RoomHoldTypeToJSON = exports.RoomHoldTypeFromJSONTyped = exports.RoomHoldTypeFromJSON = exports.instanceOfRoomHoldType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the RoomHoldType interface.
 */
function instanceOfRoomHoldType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomHoldType = instanceOfRoomHoldType;
function RoomHoldTypeFromJSON(json) {
    return RoomHoldTypeFromJSONTyped(json, false);
}
exports.RoomHoldTypeFromJSON = RoomHoldTypeFromJSON;
function RoomHoldTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'holdUntil': !(0, runtime_1.exists)(json, 'holdUntil') ? undefined : (new Date(json['holdUntil'])),
        'holdUser': !(0, runtime_1.exists)(json, 'holdUser') ? undefined : json['holdUser'],
        'comments': !(0, runtime_1.exists)(json, 'comments') ? undefined : json['comments'],
    };
}
exports.RoomHoldTypeFromJSONTyped = RoomHoldTypeFromJSONTyped;
function RoomHoldTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'holdUntil': value.holdUntil === undefined ? undefined : (value.holdUntil.toISOString()),
        'holdUser': value.holdUser,
        'comments': value.comments,
    };
}
exports.RoomHoldTypeToJSON = RoomHoldTypeToJSON;
