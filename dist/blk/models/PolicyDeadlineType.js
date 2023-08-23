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
exports.PolicyDeadlineTypeToJSON = exports.PolicyDeadlineTypeFromJSONTyped = exports.PolicyDeadlineTypeFromJSON = exports.instanceOfPolicyDeadlineType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the PolicyDeadlineType interface.
 */
function instanceOfPolicyDeadlineType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPolicyDeadlineType = instanceOfPolicyDeadlineType;
function PolicyDeadlineTypeFromJSON(json) {
    return PolicyDeadlineTypeFromJSONTyped(json, false);
}
exports.PolicyDeadlineTypeFromJSON = PolicyDeadlineTypeFromJSON;
function PolicyDeadlineTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'absoluteDeadline': !(0, runtime_1.exists)(json, 'absoluteDeadline') ? undefined : (new Date(json['absoluteDeadline'])),
        'offsetFromArrival': !(0, runtime_1.exists)(json, 'offsetFromArrival') ? undefined : json['offsetFromArrival'],
        'offsetDropTime': !(0, runtime_1.exists)(json, 'offsetDropTime') ? undefined : (new Date(json['offsetDropTime'])),
        'offsetFromBookingDate': !(0, runtime_1.exists)(json, 'offsetFromBookingDate') ? undefined : json['offsetFromBookingDate'],
    };
}
exports.PolicyDeadlineTypeFromJSONTyped = PolicyDeadlineTypeFromJSONTyped;
function PolicyDeadlineTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'absoluteDeadline': value.absoluteDeadline === undefined ? undefined : (value.absoluteDeadline.toISOString()),
        'offsetFromArrival': value.offsetFromArrival,
        'offsetDropTime': value.offsetDropTime === undefined ? undefined : (value.offsetDropTime.toISOString()),
        'offsetFromBookingDate': value.offsetFromBookingDate,
    };
}
exports.PolicyDeadlineTypeToJSON = PolicyDeadlineTypeToJSON;
