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
exports.ResGuestTypeProfileInfoToJSON = exports.ResGuestTypeProfileInfoFromJSONTyped = exports.ResGuestTypeProfileInfoFromJSON = exports.instanceOfResGuestTypeProfileInfo = void 0;
const runtime_1 = require("../runtime");
const ProfileCashieringDetailType_1 = require("./ProfileCashieringDetailType");
const ProfileType_1 = require("./ProfileType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ResGuestTypeProfileInfo interface.
 */
function instanceOfResGuestTypeProfileInfo(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResGuestTypeProfileInfo = instanceOfResGuestTypeProfileInfo;
function ResGuestTypeProfileInfoFromJSON(json) {
    return ResGuestTypeProfileInfoFromJSONTyped(json, false);
}
exports.ResGuestTypeProfileInfoFromJSON = ResGuestTypeProfileInfoFromJSON;
function ResGuestTypeProfileInfoFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileIdList': !(0, runtime_1.exists)(json, 'profileIdList') ? undefined : (json['profileIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'profile': !(0, runtime_1.exists)(json, 'profile') ? undefined : (0, ProfileType_1.ProfileTypeFromJSON)(json['profile']),
        'profileCashieringDetail': !(0, runtime_1.exists)(json, 'profileCashieringDetail') ? undefined : (0, ProfileCashieringDetailType_1.ProfileCashieringDetailTypeFromJSON)(json['profileCashieringDetail']),
        'registrationCardNo': !(0, runtime_1.exists)(json, 'registrationCardNo') ? undefined : json['registrationCardNo'],
    };
}
exports.ResGuestTypeProfileInfoFromJSONTyped = ResGuestTypeProfileInfoFromJSONTyped;
function ResGuestTypeProfileInfoToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileIdList': value.profileIdList === undefined ? undefined : (value.profileIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'profile': (0, ProfileType_1.ProfileTypeToJSON)(value.profile),
        'profileCashieringDetail': (0, ProfileCashieringDetailType_1.ProfileCashieringDetailTypeToJSON)(value.profileCashieringDetail),
        'registrationCardNo': value.registrationCardNo,
    };
}
exports.ResGuestTypeProfileInfoToJSON = ResGuestTypeProfileInfoToJSON;
