"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeSpanTypeToJSON = exports.TimeSpanTypeFromJSONTyped = exports.TimeSpanTypeFromJSON = exports.instanceOfTimeSpanType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the TimeSpanType interface.
 */
function instanceOfTimeSpanType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTimeSpanType = instanceOfTimeSpanType;
function TimeSpanTypeFromJSON(json) {
    return TimeSpanTypeFromJSONTyped(json, false);
}
exports.TimeSpanTypeFromJSON = TimeSpanTypeFromJSON;
function TimeSpanTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'startDate': !(0, runtime_1.exists)(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'endDate': !(0, runtime_1.exists)(json, 'endDate') ? undefined : (new Date(json['endDate'])),
    };
}
exports.TimeSpanTypeFromJSONTyped = TimeSpanTypeFromJSONTyped;
function TimeSpanTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString().substr(0, 10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0, 10)),
    };
}
exports.TimeSpanTypeToJSON = TimeSpanTypeToJSON;
