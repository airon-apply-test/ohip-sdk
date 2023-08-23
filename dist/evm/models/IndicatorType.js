"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.IndicatorTypeToJSON = exports.IndicatorTypeFromJSONTyped = exports.IndicatorTypeFromJSON = exports.instanceOfIndicatorType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the IndicatorType interface.
 */
function instanceOfIndicatorType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfIndicatorType = instanceOfIndicatorType;
function IndicatorTypeFromJSON(json) {
    return IndicatorTypeFromJSONTyped(json, false);
}
exports.IndicatorTypeFromJSON = IndicatorTypeFromJSON;
function IndicatorTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'indicatorName': !(0, runtime_1.exists)(json, 'indicatorName') ? undefined : json['indicatorName'],
        'count': !(0, runtime_1.exists)(json, 'count') ? undefined : json['count'],
    };
}
exports.IndicatorTypeFromJSONTyped = IndicatorTypeFromJSONTyped;
function IndicatorTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'indicatorName': value.indicatorName,
        'count': value.count,
    };
}
exports.IndicatorTypeToJSON = IndicatorTypeToJSON;
