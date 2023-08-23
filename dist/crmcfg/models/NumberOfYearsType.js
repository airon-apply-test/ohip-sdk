"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.NumberOfYearsTypeToJSON = exports.NumberOfYearsTypeFromJSONTyped = exports.NumberOfYearsTypeFromJSON = exports.NumberOfYearsType = void 0;
/**
 * Five Year period
 * @export
 */
exports.NumberOfYearsType = {
    One: 'One',
    Two: 'Two',
    Three: 'Three',
    Four: 'Four',
    Five: 'Five'
};
function NumberOfYearsTypeFromJSON(json) {
    return NumberOfYearsTypeFromJSONTyped(json, false);
}
exports.NumberOfYearsTypeFromJSON = NumberOfYearsTypeFromJSON;
function NumberOfYearsTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.NumberOfYearsTypeFromJSONTyped = NumberOfYearsTypeFromJSONTyped;
function NumberOfYearsTypeToJSON(value) {
    return value;
}
exports.NumberOfYearsTypeToJSON = NumberOfYearsTypeToJSON;
