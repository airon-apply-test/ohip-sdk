"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProfileAccessTypeToJSON = exports.ProfileAccessTypeFromJSONTyped = exports.ProfileAccessTypeFromJSON = exports.instanceOfProfileAccessType = void 0;
const runtime_1 = require("../runtime");
const ProfileSharedLevelType_1 = require("./ProfileSharedLevelType");
/**
 * Check if a given object implements the ProfileAccessType interface.
 */
function instanceOfProfileAccessType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileAccessType = instanceOfProfileAccessType;
function ProfileAccessTypeFromJSON(json) {
    return ProfileAccessTypeFromJSONTyped(json, false);
}
exports.ProfileAccessTypeFromJSON = ProfileAccessTypeFromJSON;
function ProfileAccessTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'chainCode': !(0, runtime_1.exists)(json, 'chainCode') ? undefined : json['chainCode'],
        'croCode': !(0, runtime_1.exists)(json, 'croCode') ? undefined : json['croCode'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'sharedLevel': !(0, runtime_1.exists)(json, 'sharedLevel') ? undefined : (0, ProfileSharedLevelType_1.ProfileSharedLevelTypeFromJSON)(json['sharedLevel']),
    };
}
exports.ProfileAccessTypeFromJSONTyped = ProfileAccessTypeFromJSONTyped;
function ProfileAccessTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'chainCode': value.chainCode,
        'croCode': value.croCode,
        'hotelId': value.hotelId,
        'sharedLevel': (0, ProfileSharedLevelType_1.ProfileSharedLevelTypeToJSON)(value.sharedLevel),
    };
}
exports.ProfileAccessTypeToJSON = ProfileAccessTypeToJSON;
