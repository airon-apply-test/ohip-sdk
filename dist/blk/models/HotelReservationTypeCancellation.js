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
exports.HotelReservationTypeCancellationToJSON = exports.HotelReservationTypeCancellationFromJSONTyped = exports.HotelReservationTypeCancellationFromJSON = exports.instanceOfHotelReservationTypeCancellation = void 0;
const runtime_1 = require("../runtime");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the HotelReservationTypeCancellation interface.
 */
function instanceOfHotelReservationTypeCancellation(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelReservationTypeCancellation = instanceOfHotelReservationTypeCancellation;
function HotelReservationTypeCancellationFromJSON(json) {
    return HotelReservationTypeCancellationFromJSONTyped(json, false);
}
exports.HotelReservationTypeCancellationFromJSON = HotelReservationTypeCancellationFromJSON;
function HotelReservationTypeCancellationFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'cancellationNo': !(0, runtime_1.exists)(json, 'cancellationNo') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['cancellationNo']),
        'date': !(0, runtime_1.exists)(json, 'date') ? undefined : (new Date(json['date'])),
    };
}
exports.HotelReservationTypeCancellationFromJSONTyped = HotelReservationTypeCancellationFromJSONTyped;
function HotelReservationTypeCancellationToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'code': value.code,
        'cancellationNo': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.cancellationNo),
        'date': value.date === undefined ? undefined : (value.date.toISOString().substr(0, 10)),
    };
}
exports.HotelReservationTypeCancellationToJSON = HotelReservationTypeCancellationToJSON;
