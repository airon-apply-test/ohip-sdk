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
exports.AccumulatedBusinessEventStatusTypeToJSON = exports.AccumulatedBusinessEventStatusTypeFromJSONTyped = exports.AccumulatedBusinessEventStatusTypeFromJSON = exports.AccumulatedBusinessEventStatusType = void 0;
/**
 * State of Accumulated Business Event
 * @export
 */
exports.AccumulatedBusinessEventStatusType = {
    New: 'New',
    Ready: 'Ready'
};
function AccumulatedBusinessEventStatusTypeFromJSON(json) {
    return AccumulatedBusinessEventStatusTypeFromJSONTyped(json, false);
}
exports.AccumulatedBusinessEventStatusTypeFromJSON = AccumulatedBusinessEventStatusTypeFromJSON;
function AccumulatedBusinessEventStatusTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.AccumulatedBusinessEventStatusTypeFromJSONTyped = AccumulatedBusinessEventStatusTypeFromJSONTyped;
function AccumulatedBusinessEventStatusTypeToJSON(value) {
    return value;
}
exports.AccumulatedBusinessEventStatusTypeToJSON = AccumulatedBusinessEventStatusTypeToJSON;
