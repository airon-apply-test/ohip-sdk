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
exports.DateRangeTypeToJSON = exports.DateRangeTypeFromJSONTyped = exports.DateRangeTypeFromJSON = exports.instanceOfDateRangeType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the DateRangeType interface.
 */
function instanceOfDateRangeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDateRangeType = instanceOfDateRangeType;
function DateRangeTypeFromJSON(json) {
    return DateRangeTypeFromJSONTyped(json, false);
}
exports.DateRangeTypeFromJSON = DateRangeTypeFromJSON;
function DateRangeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'start': !(0, runtime_1.exists)(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !(0, runtime_1.exists)(json, 'end') ? undefined : (new Date(json['end'])),
    };
}
exports.DateRangeTypeFromJSONTyped = DateRangeTypeFromJSONTyped;
function DateRangeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0, 10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0, 10)),
    };
}
exports.DateRangeTypeToJSON = DateRangeTypeToJSON;
