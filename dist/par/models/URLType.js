"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.URLTypeToJSON = exports.URLTypeFromJSONTyped = exports.URLTypeFromJSON = exports.instanceOfURLType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the URLType interface.
 */
function instanceOfURLType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfURLType = instanceOfURLType;
function URLTypeFromJSON(json) {
    return URLTypeFromJSONTyped(json, false);
}
exports.URLTypeFromJSON = URLTypeFromJSON;
function URLTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'value': !(0, runtime_1.exists)(json, 'value') ? undefined : json['value'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'typeDescription': !(0, runtime_1.exists)(json, 'typeDescription') ? undefined : json['typeDescription'],
        'primaryInd': !(0, runtime_1.exists)(json, 'primaryInd') ? undefined : json['primaryInd'],
        'orderSequence': !(0, runtime_1.exists)(json, 'orderSequence') ? undefined : json['orderSequence'],
    };
}
exports.URLTypeFromJSONTyped = URLTypeFromJSONTyped;
function URLTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'value': value.value,
        'type': value.type,
        'typeDescription': value.typeDescription,
        'primaryInd': value.primaryInd,
        'orderSequence': value.orderSequence,
    };
}
exports.URLTypeToJSON = URLTypeToJSON;
