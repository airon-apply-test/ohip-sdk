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
exports.AdvanceCheckInTypeToJSON = exports.AdvanceCheckInTypeFromJSONTyped = exports.AdvanceCheckInTypeFromJSON = exports.instanceOfAdvanceCheckInType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the AdvanceCheckInType interface.
 */
function instanceOfAdvanceCheckInType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAdvanceCheckInType = instanceOfAdvanceCheckInType;
function AdvanceCheckInTypeFromJSON(json) {
    return AdvanceCheckInTypeFromJSONTyped(json, false);
}
exports.AdvanceCheckInTypeFromJSON = AdvanceCheckInTypeFromJSON;
function AdvanceCheckInTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'advanceCheckedIn': !(0, runtime_1.exists)(json, 'advanceCheckedIn') ? undefined : json['advanceCheckedIn'],
        'expectedReturnTime': !(0, runtime_1.exists)(json, 'expectedReturnTime') ? undefined : (new Date(json['expectedReturnTime'])),
        'eTRComments': !(0, runtime_1.exists)(json, 'eTRComments') ? undefined : json['eTRComments'],
    };
}
exports.AdvanceCheckInTypeFromJSONTyped = AdvanceCheckInTypeFromJSONTyped;
function AdvanceCheckInTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'advanceCheckedIn': value.advanceCheckedIn,
        'expectedReturnTime': value.expectedReturnTime === undefined ? undefined : (value.expectedReturnTime.toISOString()),
        'eTRComments': value.eTRComments,
    };
}
exports.AdvanceCheckInTypeToJSON = AdvanceCheckInTypeToJSON;
