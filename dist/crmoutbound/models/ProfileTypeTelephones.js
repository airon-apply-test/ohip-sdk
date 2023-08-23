"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProfileTypeTelephonesToJSON = exports.ProfileTypeTelephonesFromJSONTyped = exports.ProfileTypeTelephonesFromJSON = exports.instanceOfProfileTypeTelephones = void 0;
const runtime_1 = require("../runtime");
const TelephoneInfoType_1 = require("./TelephoneInfoType");
/**
 * Check if a given object implements the ProfileTypeTelephones interface.
 */
function instanceOfProfileTypeTelephones(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileTypeTelephones = instanceOfProfileTypeTelephones;
function ProfileTypeTelephonesFromJSON(json) {
    return ProfileTypeTelephonesFromJSONTyped(json, false);
}
exports.ProfileTypeTelephonesFromJSON = ProfileTypeTelephonesFromJSON;
function ProfileTypeTelephonesFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'telephoneInfo': !(0, runtime_1.exists)(json, 'telephoneInfo') ? undefined : (json['telephoneInfo'].map(TelephoneInfoType_1.TelephoneInfoTypeFromJSON)),
        'totalRows': !(0, runtime_1.exists)(json, 'totalRows') ? undefined : json['totalRows'],
    };
}
exports.ProfileTypeTelephonesFromJSONTyped = ProfileTypeTelephonesFromJSONTyped;
function ProfileTypeTelephonesToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'telephoneInfo': value.telephoneInfo === undefined ? undefined : (value.telephoneInfo.map(TelephoneInfoType_1.TelephoneInfoTypeToJSON)),
        'totalRows': value.totalRows,
    };
}
exports.ProfileTypeTelephonesToJSON = ProfileTypeTelephonesToJSON;
