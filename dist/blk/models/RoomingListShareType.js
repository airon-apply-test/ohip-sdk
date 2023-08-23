"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RoomingListShareTypeToJSON = exports.RoomingListShareTypeFromJSONTyped = exports.RoomingListShareTypeFromJSON = exports.instanceOfRoomingListShareType = void 0;
const runtime_1 = require("../runtime");
const EffectiveRateType_1 = require("./EffectiveRateType");
const RoomingListShareReservationType_1 = require("./RoomingListShareReservationType");
const TimeSpanType_1 = require("./TimeSpanType");
/**
 * Check if a given object implements the RoomingListShareType interface.
 */
function instanceOfRoomingListShareType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomingListShareType = instanceOfRoomingListShareType;
function RoomingListShareTypeFromJSON(json) {
    return RoomingListShareTypeFromJSONTyped(json, false);
}
exports.RoomingListShareTypeFromJSON = RoomingListShareTypeFromJSON;
function RoomingListShareTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservation': !(0, runtime_1.exists)(json, 'reservation') ? undefined : (json['reservation'].map(RoomingListShareReservationType_1.RoomingListShareReservationTypeFromJSON)),
        'effectiveRates': !(0, runtime_1.exists)(json, 'effectiveRates') ? undefined : (json['effectiveRates'].map(EffectiveRateType_1.EffectiveRateTypeFromJSON)),
        'timeSpan': !(0, runtime_1.exists)(json, 'timeSpan') ? undefined : (0, TimeSpanType_1.TimeSpanTypeFromJSON)(json['timeSpan']),
    };
}
exports.RoomingListShareTypeFromJSONTyped = RoomingListShareTypeFromJSONTyped;
function RoomingListShareTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservation': value.reservation === undefined ? undefined : (value.reservation.map(RoomingListShareReservationType_1.RoomingListShareReservationTypeToJSON)),
        'effectiveRates': value.effectiveRates === undefined ? undefined : (value.effectiveRates.map(EffectiveRateType_1.EffectiveRateTypeToJSON)),
        'timeSpan': (0, TimeSpanType_1.TimeSpanTypeToJSON)(value.timeSpan),
    };
}
exports.RoomingListShareTypeToJSON = RoomingListShareTypeToJSON;
