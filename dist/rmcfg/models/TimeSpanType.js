"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
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
        'duration': !(0, runtime_1.exists)(json, 'duration') ? undefined : json['duration'],
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
        'duration': value.duration,
    };
}
exports.TimeSpanTypeToJSON = TimeSpanTypeToJSON;
