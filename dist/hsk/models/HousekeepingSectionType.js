"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HousekeepingSectionTypeToJSON = exports.HousekeepingSectionTypeFromJSONTyped = exports.HousekeepingSectionTypeFromJSON = exports.instanceOfHousekeepingSectionType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the HousekeepingSectionType interface.
 */
function instanceOfHousekeepingSectionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHousekeepingSectionType = instanceOfHousekeepingSectionType;
function HousekeepingSectionTypeFromJSON(json) {
    return HousekeepingSectionTypeFromJSONTyped(json, false);
}
exports.HousekeepingSectionTypeFromJSON = HousekeepingSectionTypeFromJSON;
function HousekeepingSectionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'sectionGroup': !(0, runtime_1.exists)(json, 'sectionGroup') ? undefined : json['sectionGroup'],
        'targetCredits': !(0, runtime_1.exists)(json, 'targetCredits') ? undefined : json['targetCredits'],
        'rooms': !(0, runtime_1.exists)(json, 'rooms') ? undefined : json['rooms'],
        'roomCredits': !(0, runtime_1.exists)(json, 'roomCredits') ? undefined : json['roomCredits'],
        'sequence': !(0, runtime_1.exists)(json, 'sequence') ? undefined : json['sequence'],
        'inactive': !(0, runtime_1.exists)(json, 'inactive') ? undefined : json['inactive'],
    };
}
exports.HousekeepingSectionTypeFromJSONTyped = HousekeepingSectionTypeFromJSONTyped;
function HousekeepingSectionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'code': value.code,
        'description': value.description,
        'sectionGroup': value.sectionGroup,
        'targetCredits': value.targetCredits,
        'rooms': value.rooms,
        'roomCredits': value.roomCredits,
        'sequence': value.sequence,
        'inactive': value.inactive,
    };
}
exports.HousekeepingSectionTypeToJSON = HousekeepingSectionTypeToJSON;
