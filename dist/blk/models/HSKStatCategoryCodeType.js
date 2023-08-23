"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HSKStatCategoryCodeTypeToJSON = exports.HSKStatCategoryCodeTypeFromJSONTyped = exports.HSKStatCategoryCodeTypeFromJSON = exports.HSKStatCategoryCodeType = void 0;
/**
 * Supported housekeeping statistical category codes.
 * @export
 */
exports.HSKStatCategoryCodeType = {
    HotelCode: 'HotelCode'
};
function HSKStatCategoryCodeTypeFromJSON(json) {
    return HSKStatCategoryCodeTypeFromJSONTyped(json, false);
}
exports.HSKStatCategoryCodeTypeFromJSON = HSKStatCategoryCodeTypeFromJSON;
function HSKStatCategoryCodeTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.HSKStatCategoryCodeTypeFromJSONTyped = HSKStatCategoryCodeTypeFromJSONTyped;
function HSKStatCategoryCodeTypeToJSON(value) {
    return value;
}
exports.HSKStatCategoryCodeTypeToJSON = HSKStatCategoryCodeTypeToJSON;
