"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Leisure Management API
 * APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.DateTimeSpanTypeToJSON = exports.DateTimeSpanTypeFromJSONTyped = exports.DateTimeSpanTypeFromJSON = exports.instanceOfDateTimeSpanType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the DateTimeSpanType interface.
 */
function instanceOfDateTimeSpanType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDateTimeSpanType = instanceOfDateTimeSpanType;
function DateTimeSpanTypeFromJSON(json) {
    return DateTimeSpanTypeFromJSONTyped(json, false);
}
exports.DateTimeSpanTypeFromJSON = DateTimeSpanTypeFromJSON;
function DateTimeSpanTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'startDateTime': !(0, runtime_1.exists)(json, 'startDateTime') ? undefined : (new Date(json['startDateTime'])),
        'endDateTime': !(0, runtime_1.exists)(json, 'endDateTime') ? undefined : (new Date(json['endDateTime'])),
    };
}
exports.DateTimeSpanTypeFromJSONTyped = DateTimeSpanTypeFromJSONTyped;
function DateTimeSpanTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'startDateTime': value.startDateTime === undefined ? undefined : (value.startDateTime.toISOString()),
        'endDateTime': value.endDateTime === undefined ? undefined : (value.endDateTime.toISOString()),
    };
}
exports.DateTimeSpanTypeToJSON = DateTimeSpanTypeToJSON;
