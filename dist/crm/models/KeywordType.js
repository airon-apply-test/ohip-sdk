"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.KeywordTypeToJSON = exports.KeywordTypeFromJSONTyped = exports.KeywordTypeFromJSON = exports.instanceOfKeywordType = void 0;
const runtime_1 = require("../runtime");
const KeywordDetailType_1 = require("./KeywordDetailType");
/**
 * Check if a given object implements the KeywordType interface.
 */
function instanceOfKeywordType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfKeywordType = instanceOfKeywordType;
function KeywordTypeFromJSON(json) {
    return KeywordTypeFromJSONTyped(json, false);
}
exports.KeywordTypeFromJSON = KeywordTypeFromJSON;
function KeywordTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'keywordDetail': !(0, runtime_1.exists)(json, 'keywordDetail') ? undefined : (0, KeywordDetailType_1.KeywordDetailTypeFromJSON)(json['keywordDetail']),
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'keyword': !(0, runtime_1.exists)(json, 'keyword') ? undefined : json['keyword'],
    };
}
exports.KeywordTypeFromJSONTyped = KeywordTypeFromJSONTyped;
function KeywordTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'keywordDetail': (0, KeywordDetailType_1.KeywordDetailTypeToJSON)(value.keywordDetail),
        'type': value.type,
        'keyword': value.keyword,
    };
}
exports.KeywordTypeToJSON = KeywordTypeToJSON;
