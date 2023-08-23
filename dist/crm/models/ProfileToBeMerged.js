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
exports.ProfileToBeMergedToJSON = exports.ProfileToBeMergedFromJSONTyped = exports.ProfileToBeMergedFromJSON = exports.instanceOfProfileToBeMerged = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const ProfileType_1 = require("./ProfileType");
const UniqueIDType_1 = require("./UniqueIDType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ProfileToBeMerged interface.
 */
function instanceOfProfileToBeMerged(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileToBeMerged = instanceOfProfileToBeMerged;
function ProfileToBeMergedFromJSON(json) {
    return ProfileToBeMergedFromJSONTyped(json, false);
}
exports.ProfileToBeMergedFromJSON = ProfileToBeMergedFromJSON;
function ProfileToBeMergedFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'victimProfileId': !(0, runtime_1.exists)(json, 'victimProfileId') ? undefined : (json['victimProfileId'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'profileDetails': !(0, runtime_1.exists)(json, 'profileDetails') ? undefined : (0, ProfileType_1.ProfileTypeFromJSON)(json['profileDetails']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ProfileToBeMergedFromJSONTyped = ProfileToBeMergedFromJSONTyped;
function ProfileToBeMergedToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'victimProfileId': value.victimProfileId === undefined ? undefined : (value.victimProfileId.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'profileDetails': (0, ProfileType_1.ProfileTypeToJSON)(value.profileDetails),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ProfileToBeMergedToJSON = ProfileToBeMergedToJSON;
