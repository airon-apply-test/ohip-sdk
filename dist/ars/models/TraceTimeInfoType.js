"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.TraceTimeInfoTypeToJSON = exports.TraceTimeInfoTypeFromJSONTyped = exports.TraceTimeInfoTypeFromJSON = exports.instanceOfTraceTimeInfoType = void 0;
const runtime_1 = require("../runtime");
const DateTimeSpanType_1 = require("./DateTimeSpanType");
/**
 * Check if a given object implements the TraceTimeInfoType interface.
 */
function instanceOfTraceTimeInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTraceTimeInfoType = instanceOfTraceTimeInfoType;
function TraceTimeInfoTypeFromJSON(json) {
    return TraceTimeInfoTypeFromJSONTyped(json, false);
}
exports.TraceTimeInfoTypeFromJSON = TraceTimeInfoTypeFromJSON;
function TraceTimeInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'dateTimeSpan': !(0, runtime_1.exists)(json, 'dateTimeSpan') ? undefined : (0, DateTimeSpanType_1.DateTimeSpanTypeFromJSON)(json['dateTimeSpan']),
        'traceOn': !(0, runtime_1.exists)(json, 'traceOn') ? undefined : (new Date(json['traceOn'])),
        'traceTime': !(0, runtime_1.exists)(json, 'traceTime') ? undefined : json['traceTime'],
        'enteredBy': !(0, runtime_1.exists)(json, 'enteredBy') ? undefined : json['enteredBy'],
    };
}
exports.TraceTimeInfoTypeFromJSONTyped = TraceTimeInfoTypeFromJSONTyped;
function TraceTimeInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'dateTimeSpan': (0, DateTimeSpanType_1.DateTimeSpanTypeToJSON)(value.dateTimeSpan),
        'traceOn': value.traceOn === undefined ? undefined : (value.traceOn.toISOString()),
        'traceTime': value.traceTime,
        'enteredBy': value.enteredBy,
    };
}
exports.TraceTimeInfoTypeToJSON = TraceTimeInfoTypeToJSON;
