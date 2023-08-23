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
exports.ActivityInfoTypeStatusToJSON = exports.ActivityInfoTypeStatusFromJSONTyped = exports.ActivityInfoTypeStatusFromJSON = exports.instanceOfActivityInfoTypeStatus = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ActivityInfoTypeStatus interface.
 */
function instanceOfActivityInfoTypeStatus(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfActivityInfoTypeStatus = instanceOfActivityInfoTypeStatus;
function ActivityInfoTypeStatusFromJSON(json) {
    return ActivityInfoTypeStatusFromJSONTyped(json, false);
}
exports.ActivityInfoTypeStatusFromJSON = ActivityInfoTypeStatusFromJSON;
function ActivityInfoTypeStatusFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'completedBy': !(0, runtime_1.exists)(json, 'completedBy') ? undefined : json['completedBy'],
        'completedOn': !(0, runtime_1.exists)(json, 'completedOn') ? undefined : (new Date(json['completedOn'])),
        'completed': !(0, runtime_1.exists)(json, 'completed') ? undefined : json['completed'],
    };
}
exports.ActivityInfoTypeStatusFromJSONTyped = ActivityInfoTypeStatusFromJSONTyped;
function ActivityInfoTypeStatusToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'completedBy': value.completedBy,
        'completedOn': value.completedOn === undefined ? undefined : (value.completedOn.toISOString()),
        'completed': value.completed,
    };
}
exports.ActivityInfoTypeStatusToJSON = ActivityInfoTypeStatusToJSON;
