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
exports.ProfileTypeProfileMembershipsToJSON = exports.ProfileTypeProfileMembershipsFromJSONTyped = exports.ProfileTypeProfileMembershipsFromJSON = exports.instanceOfProfileTypeProfileMemberships = void 0;
const runtime_1 = require("../runtime");
const ProfileMembershipType_1 = require("./ProfileMembershipType");
/**
 * Check if a given object implements the ProfileTypeProfileMemberships interface.
 */
function instanceOfProfileTypeProfileMemberships(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileTypeProfileMemberships = instanceOfProfileTypeProfileMemberships;
function ProfileTypeProfileMembershipsFromJSON(json) {
    return ProfileTypeProfileMembershipsFromJSONTyped(json, false);
}
exports.ProfileTypeProfileMembershipsFromJSON = ProfileTypeProfileMembershipsFromJSON;
function ProfileTypeProfileMembershipsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileMembership': !(0, runtime_1.exists)(json, 'profileMembership') ? undefined : (json['profileMembership'].map(ProfileMembershipType_1.ProfileMembershipTypeFromJSON)),
        'hasMore': !(0, runtime_1.exists)(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !(0, runtime_1.exists)(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !(0, runtime_1.exists)(json, 'count') ? undefined : json['count'],
    };
}
exports.ProfileTypeProfileMembershipsFromJSONTyped = ProfileTypeProfileMembershipsFromJSONTyped;
function ProfileTypeProfileMembershipsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileMembership': value.profileMembership === undefined ? undefined : (value.profileMembership.map(ProfileMembershipType_1.ProfileMembershipTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}
exports.ProfileTypeProfileMembershipsToJSON = ProfileTypeProfileMembershipsToJSON;
