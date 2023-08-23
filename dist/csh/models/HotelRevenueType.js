"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HotelRevenueTypeToJSON = exports.HotelRevenueTypeFromJSONTyped = exports.HotelRevenueTypeFromJSON = exports.instanceOfHotelRevenueType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the HotelRevenueType interface.
 */
function instanceOfHotelRevenueType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelRevenueType = instanceOfHotelRevenueType;
function HotelRevenueTypeFromJSON(json) {
    return HotelRevenueTypeFromJSONTyped(json, false);
}
exports.HotelRevenueTypeFromJSON = HotelRevenueTypeFromJSON;
function HotelRevenueTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomRevenue': !(0, runtime_1.exists)(json, 'roomRevenue') ? undefined : json['roomRevenue'],
        'houseRevenue': !(0, runtime_1.exists)(json, 'houseRevenue') ? undefined : json['houseRevenue'],
        'compRevenue': !(0, runtime_1.exists)(json, 'compRevenue') ? undefined : json['compRevenue'],
        'occupiedRooms': !(0, runtime_1.exists)(json, 'occupiedRooms') ? undefined : json['occupiedRooms'],
        'houseRooms': !(0, runtime_1.exists)(json, 'houseRooms') ? undefined : json['houseRooms'],
        'compRooms': !(0, runtime_1.exists)(json, 'compRooms') ? undefined : json['compRooms'],
        'zeroRoomsRevenue': !(0, runtime_1.exists)(json, 'zeroRoomsRevenue') ? undefined : json['zeroRoomsRevenue'],
        'currencyCode': !(0, runtime_1.exists)(json, 'currencyCode') ? undefined : json['currencyCode'],
    };
}
exports.HotelRevenueTypeFromJSONTyped = HotelRevenueTypeFromJSONTyped;
function HotelRevenueTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomRevenue': value.roomRevenue,
        'houseRevenue': value.houseRevenue,
        'compRevenue': value.compRevenue,
        'occupiedRooms': value.occupiedRooms,
        'houseRooms': value.houseRooms,
        'compRooms': value.compRooms,
        'zeroRoomsRevenue': value.zeroRoomsRevenue,
        'currencyCode': value.currencyCode,
    };
}
exports.HotelRevenueTypeToJSON = HotelRevenueTypeToJSON;
