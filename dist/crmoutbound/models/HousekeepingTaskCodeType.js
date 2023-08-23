"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HousekeepingTaskCodeTypeToJSON = exports.HousekeepingTaskCodeTypeFromJSONTyped = exports.HousekeepingTaskCodeTypeFromJSON = exports.instanceOfHousekeepingTaskCodeType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the HousekeepingTaskCodeType interface.
 */
function instanceOfHousekeepingTaskCodeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHousekeepingTaskCodeType = instanceOfHousekeepingTaskCodeType;
function HousekeepingTaskCodeTypeFromJSON(json) {
    return HousekeepingTaskCodeTypeFromJSONTyped(json, false);
}
exports.HousekeepingTaskCodeTypeFromJSON = HousekeepingTaskCodeTypeFromJSON;
function HousekeepingTaskCodeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'instructions': !(0, runtime_1.exists)(json, 'instructions') ? undefined : json['instructions'],
        'facilityDepartureTask': !(0, runtime_1.exists)(json, 'facilityDepartureTask') ? undefined : json['facilityDepartureTask'],
        'linenChange': !(0, runtime_1.exists)(json, 'linenChange') ? undefined : json['linenChange'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'sequence': !(0, runtime_1.exists)(json, 'sequence') ? undefined : json['sequence'],
    };
}
exports.HousekeepingTaskCodeTypeFromJSONTyped = HousekeepingTaskCodeTypeFromJSONTyped;
function HousekeepingTaskCodeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'instructions': value.instructions,
        'facilityDepartureTask': value.facilityDepartureTask,
        'linenChange': value.linenChange,
        'hotelId': value.hotelId,
        'code': value.code,
        'sequence': value.sequence,
    };
}
exports.HousekeepingTaskCodeTypeToJSON = HousekeepingTaskCodeTypeToJSON;
