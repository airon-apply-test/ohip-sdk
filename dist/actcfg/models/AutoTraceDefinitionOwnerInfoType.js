"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity Management API
 * APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.AutoTraceDefinitionOwnerInfoTypeToJSON = exports.AutoTraceDefinitionOwnerInfoTypeFromJSONTyped = exports.AutoTraceDefinitionOwnerInfoTypeFromJSON = exports.instanceOfAutoTraceDefinitionOwnerInfoType = void 0;
const runtime_1 = require("../runtime");
const ActivityOwnerType_1 = require("./ActivityOwnerType");
const AutoTraceOwnerAssignmentType_1 = require("./AutoTraceOwnerAssignmentType");
/**
 * Check if a given object implements the AutoTraceDefinitionOwnerInfoType interface.
 */
function instanceOfAutoTraceDefinitionOwnerInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAutoTraceDefinitionOwnerInfoType = instanceOfAutoTraceDefinitionOwnerInfoType;
function AutoTraceDefinitionOwnerInfoTypeFromJSON(json) {
    return AutoTraceDefinitionOwnerInfoTypeFromJSONTyped(json, false);
}
exports.AutoTraceDefinitionOwnerInfoTypeFromJSON = AutoTraceDefinitionOwnerInfoTypeFromJSON;
function AutoTraceDefinitionOwnerInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'ownerofActivity': !(0, runtime_1.exists)(json, 'ownerofActivity') ? undefined : (0, ActivityOwnerType_1.ActivityOwnerTypeFromJSON)(json['ownerofActivity']),
        'customOwnerCode': !(0, runtime_1.exists)(json, 'customOwnerCode') ? undefined : json['customOwnerCode'],
        'ownerAssignmentExist': !(0, runtime_1.exists)(json, 'ownerAssignmentExist') ? undefined : json['ownerAssignmentExist'],
        'ownerAssignment': !(0, runtime_1.exists)(json, 'ownerAssignment') ? undefined : (json['ownerAssignment'].map(AutoTraceOwnerAssignmentType_1.AutoTraceOwnerAssignmentTypeFromJSON)),
    };
}
exports.AutoTraceDefinitionOwnerInfoTypeFromJSONTyped = AutoTraceDefinitionOwnerInfoTypeFromJSONTyped;
function AutoTraceDefinitionOwnerInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'ownerofActivity': (0, ActivityOwnerType_1.ActivityOwnerTypeToJSON)(value.ownerofActivity),
        'customOwnerCode': value.customOwnerCode,
        'ownerAssignmentExist': value.ownerAssignmentExist,
        'ownerAssignment': value.ownerAssignment === undefined ? undefined : (value.ownerAssignment.map(AutoTraceOwnerAssignmentType_1.AutoTraceOwnerAssignmentTypeToJSON)),
    };
}
exports.AutoTraceDefinitionOwnerInfoTypeToJSON = AutoTraceDefinitionOwnerInfoTypeToJSON;
