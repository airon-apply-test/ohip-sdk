"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FrequencyTypeTypeToJSON = exports.FrequencyTypeTypeFromJSONTyped = exports.FrequencyTypeTypeFromJSON = exports.FrequencyTypeType = void 0;
/**
 * Simple type for storing frequency type.
 * @export
 */
exports.FrequencyTypeType = {
    Everyday: 'Everyday',
    EveryXDays: 'EveryXDays',
    CustomSchedule: 'CustomSchedule',
    DepartureDayOnly: 'DepartureDayOnly',
    ArrivalDayOnly: 'ArrivalDayOnly',
    NotRequired: 'NotRequired',
    SpecificDays: 'SpecificDays'
};
function FrequencyTypeTypeFromJSON(json) {
    return FrequencyTypeTypeFromJSONTyped(json, false);
}
exports.FrequencyTypeTypeFromJSON = FrequencyTypeTypeFromJSON;
function FrequencyTypeTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.FrequencyTypeTypeFromJSONTyped = FrequencyTypeTypeFromJSONTyped;
function FrequencyTypeTypeToJSON(value) {
    return value;
}
exports.FrequencyTypeTypeToJSON = FrequencyTypeTypeToJSON;
