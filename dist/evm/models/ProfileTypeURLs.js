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
exports.ProfileTypeURLsToJSON = exports.ProfileTypeURLsFromJSONTyped = exports.ProfileTypeURLsFromJSON = exports.instanceOfProfileTypeURLs = void 0;
const runtime_1 = require("../runtime");
const URLInfoType_1 = require("./URLInfoType");
/**
 * Check if a given object implements the ProfileTypeURLs interface.
 */
function instanceOfProfileTypeURLs(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileTypeURLs = instanceOfProfileTypeURLs;
function ProfileTypeURLsFromJSON(json) {
    return ProfileTypeURLsFromJSONTyped(json, false);
}
exports.ProfileTypeURLsFromJSON = ProfileTypeURLsFromJSON;
function ProfileTypeURLsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'uRLInfo': !(0, runtime_1.exists)(json, 'uRLInfo') ? undefined : (json['uRLInfo'].map(URLInfoType_1.URLInfoTypeFromJSON)),
        'hasMore': !(0, runtime_1.exists)(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !(0, runtime_1.exists)(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !(0, runtime_1.exists)(json, 'count') ? undefined : json['count'],
    };
}
exports.ProfileTypeURLsFromJSONTyped = ProfileTypeURLsFromJSONTyped;
function ProfileTypeURLsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'uRLInfo': value.uRLInfo === undefined ? undefined : (value.uRLInfo.map(URLInfoType_1.URLInfoTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}
exports.ProfileTypeURLsToJSON = ProfileTypeURLsToJSON;
