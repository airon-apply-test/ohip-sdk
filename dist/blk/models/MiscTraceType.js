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
exports.MiscTraceTypeToJSON = exports.MiscTraceTypeFromJSONTyped = exports.MiscTraceTypeFromJSON = exports.instanceOfMiscTraceType = void 0;
const runtime_1 = require("../runtime");
const TraceBlockInfoType_1 = require("./TraceBlockInfoType");
const TraceResvInfoType_1 = require("./TraceResvInfoType");
const TraceType_1 = require("./TraceType");
/**
 * Check if a given object implements the MiscTraceType interface.
 */
function instanceOfMiscTraceType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfMiscTraceType = instanceOfMiscTraceType;
function MiscTraceTypeFromJSON(json) {
    return MiscTraceTypeFromJSONTyped(json, false);
}
exports.MiscTraceTypeFromJSON = MiscTraceTypeFromJSON;
function MiscTraceTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'trace': !(0, runtime_1.exists)(json, 'trace') ? undefined : (0, TraceType_1.TraceTypeFromJSON)(json['trace']),
        'blockInfo': !(0, runtime_1.exists)(json, 'blockInfo') ? undefined : (0, TraceBlockInfoType_1.TraceBlockInfoTypeFromJSON)(json['blockInfo']),
        'reservationInfo': !(0, runtime_1.exists)(json, 'reservationInfo') ? undefined : (0, TraceResvInfoType_1.TraceResvInfoTypeFromJSON)(json['reservationInfo']),
    };
}
exports.MiscTraceTypeFromJSONTyped = MiscTraceTypeFromJSONTyped;
function MiscTraceTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'trace': (0, TraceType_1.TraceTypeToJSON)(value.trace),
        'blockInfo': (0, TraceBlockInfoType_1.TraceBlockInfoTypeToJSON)(value.blockInfo),
        'reservationInfo': (0, TraceResvInfoType_1.TraceResvInfoTypeToJSON)(value.reservationInfo),
    };
}
exports.MiscTraceTypeToJSON = MiscTraceTypeToJSON;
