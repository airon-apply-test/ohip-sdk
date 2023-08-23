"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * Opera Cloud Inventory Asynchronous API
 * APIs to cater for Inventory related asynchronous functionality in OPERA Cloud. This includes viewing inventory data along with its revenue and updating inventory&apos;s sell limits for date ranges. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.4.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RoomTypeSellLimitTypeToJSON = exports.RoomTypeSellLimitTypeFromJSONTyped = exports.RoomTypeSellLimitTypeFromJSON = exports.instanceOfRoomTypeSellLimitType = exports.RoomTypeSellLimitTypeFlatOrPercentageEnum = void 0;
const runtime_1 = require("../runtime");
/**
 * @export
 */
exports.RoomTypeSellLimitTypeFlatOrPercentageEnum = {
    Flat: 'Flat',
    Percentage: 'Percentage'
};
/**
 * Check if a given object implements the RoomTypeSellLimitType interface.
 */
function instanceOfRoomTypeSellLimitType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomTypeSellLimitType = instanceOfRoomTypeSellLimitType;
function RoomTypeSellLimitTypeFromJSON(json) {
    return RoomTypeSellLimitTypeFromJSONTyped(json, false);
}
exports.RoomTypeSellLimitTypeFromJSON = RoomTypeSellLimitTypeFromJSON;
function RoomTypeSellLimitTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'date': !(0, runtime_1.exists)(json, 'date') ? undefined : (new Date(json['date'])),
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : json['amount'],
        'flatOrPercentage': !(0, runtime_1.exists)(json, 'flatOrPercentage') ? undefined : json['flatOrPercentage'],
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : json['roomType'],
    };
}
exports.RoomTypeSellLimitTypeFromJSONTyped = RoomTypeSellLimitTypeFromJSONTyped;
function RoomTypeSellLimitTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'date': value.date === undefined ? undefined : (value.date.toISOString().substr(0, 10)),
        'amount': value.amount,
        'flatOrPercentage': value.flatOrPercentage,
        'roomType': value.roomType,
    };
}
exports.RoomTypeSellLimitTypeToJSON = RoomTypeSellLimitTypeToJSON;
