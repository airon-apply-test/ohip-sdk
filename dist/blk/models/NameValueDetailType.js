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
exports.NameValueDetailTypeToJSON = exports.NameValueDetailTypeFromJSONTyped = exports.NameValueDetailTypeFromJSON = exports.instanceOfNameValueDetailType = void 0;
const runtime_1 = require("../runtime");
const NameValueType_1 = require("./NameValueType");
/**
 * Check if a given object implements the NameValueDetailType interface.
 */
function instanceOfNameValueDetailType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfNameValueDetailType = instanceOfNameValueDetailType;
function NameValueDetailTypeFromJSON(json) {
    return NameValueDetailTypeFromJSONTyped(json, false);
}
exports.NameValueDetailTypeFromJSON = NameValueDetailTypeFromJSON;
function NameValueDetailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'nameValues': !(0, runtime_1.exists)(json, 'nameValues') ? undefined : (json['nameValues'].map(NameValueType_1.NameValueTypeFromJSON)),
    };
}
exports.NameValueDetailTypeFromJSONTyped = NameValueDetailTypeFromJSONTyped;
function NameValueDetailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'nameValues': value.nameValues === undefined ? undefined : (value.nameValues.map(NameValueType_1.NameValueTypeToJSON)),
    };
}
exports.NameValueDetailTypeToJSON = NameValueDetailTypeToJSON;
