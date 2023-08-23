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
exports.InterfaceRightsStatusTypeToJSON = exports.InterfaceRightsStatusTypeFromJSONTyped = exports.InterfaceRightsStatusTypeFromJSON = exports.instanceOfInterfaceRightsStatusType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the InterfaceRightsStatusType interface.
 */
function instanceOfInterfaceRightsStatusType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfInterfaceRightsStatusType = instanceOfInterfaceRightsStatusType;
function InterfaceRightsStatusTypeFromJSON(json) {
    return InterfaceRightsStatusTypeFromJSONTyped(json, false);
}
exports.InterfaceRightsStatusTypeFromJSON = InterfaceRightsStatusTypeFromJSON;
function InterfaceRightsStatusTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'right': !(0, runtime_1.exists)(json, 'right') ? undefined : json['right'],
        'statusCode': !(0, runtime_1.exists)(json, 'statusCode') ? undefined : json['statusCode'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'category': !(0, runtime_1.exists)(json, 'category') ? undefined : json['category'],
    };
}
exports.InterfaceRightsStatusTypeFromJSONTyped = InterfaceRightsStatusTypeFromJSONTyped;
function InterfaceRightsStatusTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'right': value.right,
        'statusCode': value.statusCode,
        'description': value.description,
        'category': value.category,
    };
}
exports.InterfaceRightsStatusTypeToJSON = InterfaceRightsStatusTypeToJSON;
