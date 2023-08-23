"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.GdsNegotiatedInfoTypeToJSON = exports.GdsNegotiatedInfoTypeFromJSONTyped = exports.GdsNegotiatedInfoTypeFromJSON = exports.instanceOfGdsNegotiatedInfoType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the GdsNegotiatedInfoType interface.
 */
function instanceOfGdsNegotiatedInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfGdsNegotiatedInfoType = instanceOfGdsNegotiatedInfoType;
function GdsNegotiatedInfoTypeFromJSON(json) {
    return GdsNegotiatedInfoTypeFromJSONTyped(json, false);
}
exports.GdsNegotiatedInfoTypeFromJSON = GdsNegotiatedInfoTypeFromJSON;
function GdsNegotiatedInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'accessCode': !(0, runtime_1.exists)(json, 'accessCode') ? undefined : json['accessCode'],
        'order': !(0, runtime_1.exists)(json, 'order') ? undefined : json['order'],
        'inactive': !(0, runtime_1.exists)(json, 'inactive') ? undefined : json['inactive'],
        'start': !(0, runtime_1.exists)(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !(0, runtime_1.exists)(json, 'end') ? undefined : (new Date(json['end'])),
    };
}
exports.GdsNegotiatedInfoTypeFromJSONTyped = GdsNegotiatedInfoTypeFromJSONTyped;
function GdsNegotiatedInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'accessCode': value.accessCode,
        'order': value.order,
        'inactive': value.inactive,
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0, 10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0, 10)),
    };
}
exports.GdsNegotiatedInfoTypeToJSON = GdsNegotiatedInfoTypeToJSON;
