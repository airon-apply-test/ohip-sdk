"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.AccumulatedBusinessEventTypeToJSON = exports.AccumulatedBusinessEventTypeFromJSONTyped = exports.AccumulatedBusinessEventTypeFromJSON = exports.instanceOfAccumulatedBusinessEventType = void 0;
const runtime_1 = require("../runtime");
const AccumulatedBusinessEventStatusType_1 = require("./AccumulatedBusinessEventStatusType");
/**
 * Check if a given object implements the AccumulatedBusinessEventType interface.
 */
function instanceOfAccumulatedBusinessEventType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAccumulatedBusinessEventType = instanceOfAccumulatedBusinessEventType;
function AccumulatedBusinessEventTypeFromJSON(json) {
    return AccumulatedBusinessEventTypeFromJSONTyped(json, false);
}
exports.AccumulatedBusinessEventTypeFromJSON = AccumulatedBusinessEventTypeFromJSON;
function AccumulatedBusinessEventTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'primaryKey': !(0, runtime_1.exists)(json, 'primaryKey') ? undefined : json['primaryKey'],
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : (0, AccumulatedBusinessEventStatusType_1.AccumulatedBusinessEventStatusTypeFromJSON)(json['status']),
        'createDate': !(0, runtime_1.exists)(json, 'createDate') ? undefined : (new Date(json['createDate'])),
        '_interface': !(0, runtime_1.exists)(json, 'interface') ? undefined : json['interface'],
        'module': !(0, runtime_1.exists)(json, 'module') ? undefined : json['module'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
    };
}
exports.AccumulatedBusinessEventTypeFromJSONTyped = AccumulatedBusinessEventTypeFromJSONTyped;
function AccumulatedBusinessEventTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'primaryKey': value.primaryKey,
        'status': (0, AccumulatedBusinessEventStatusType_1.AccumulatedBusinessEventStatusTypeToJSON)(value.status),
        'createDate': value.createDate === undefined ? undefined : (value.createDate.toISOString()),
        'interface': value._interface,
        'module': value.module,
        'hotelId': value.hotelId,
    };
}
exports.AccumulatedBusinessEventTypeToJSON = AccumulatedBusinessEventTypeToJSON;
