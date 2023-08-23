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
exports.AutoTraceGroupConfigTypeToJSON = exports.AutoTraceGroupConfigTypeFromJSONTyped = exports.AutoTraceGroupConfigTypeFromJSON = exports.AutoTraceGroupConfigType = void 0;
/**
 * A collection of supported list of Auto Trace Groups.
 * @export
 */
exports.AutoTraceGroupConfigType = {
    Accounts: 'Accounts',
    Contacts: 'Contacts',
    Blocks: 'Blocks',
    Activities: 'Activities'
};
function AutoTraceGroupConfigTypeFromJSON(json) {
    return AutoTraceGroupConfigTypeFromJSONTyped(json, false);
}
exports.AutoTraceGroupConfigTypeFromJSON = AutoTraceGroupConfigTypeFromJSON;
function AutoTraceGroupConfigTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.AutoTraceGroupConfigTypeFromJSONTyped = AutoTraceGroupConfigTypeFromJSONTyped;
function AutoTraceGroupConfigTypeToJSON(value) {
    return value;
}
exports.AutoTraceGroupConfigTypeToJSON = AutoTraceGroupConfigTypeToJSON;
