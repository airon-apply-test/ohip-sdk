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
exports.ProfileSubscriptionTypeProfileInfoToJSON = exports.ProfileSubscriptionTypeProfileInfoFromJSONTyped = exports.ProfileSubscriptionTypeProfileInfoFromJSON = exports.instanceOfProfileSubscriptionTypeProfileInfo = void 0;
const runtime_1 = require("../runtime");
const PersonNameTypeType_1 = require("./PersonNameTypeType");
const ProfileTypeType_1 = require("./ProfileTypeType");
/**
 * Check if a given object implements the ProfileSubscriptionTypeProfileInfo interface.
 */
function instanceOfProfileSubscriptionTypeProfileInfo(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileSubscriptionTypeProfileInfo = instanceOfProfileSubscriptionTypeProfileInfo;
function ProfileSubscriptionTypeProfileInfoFromJSON(json) {
    return ProfileSubscriptionTypeProfileInfoFromJSONTyped(json, false);
}
exports.ProfileSubscriptionTypeProfileInfoFromJSON = ProfileSubscriptionTypeProfileInfoFromJSON;
function ProfileSubscriptionTypeProfileInfoFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'name': !(0, runtime_1.exists)(json, 'name') ? undefined : json['name'],
        'fullName': !(0, runtime_1.exists)(json, 'fullName') ? undefined : json['fullName'],
        'namePrefix': !(0, runtime_1.exists)(json, 'namePrefix') ? undefined : json['namePrefix'],
        'givenName': !(0, runtime_1.exists)(json, 'givenName') ? undefined : json['givenName'],
        'middleName': !(0, runtime_1.exists)(json, 'middleName') ? undefined : json['middleName'],
        'nameSuffix': !(0, runtime_1.exists)(json, 'nameSuffix') ? undefined : json['nameSuffix'],
        'nameTitle': !(0, runtime_1.exists)(json, 'nameTitle') ? undefined : json['nameTitle'],
        'nameType': !(0, runtime_1.exists)(json, 'nameType') ? undefined : (0, PersonNameTypeType_1.PersonNameTypeTypeFromJSON)(json['nameType']),
        'profileType': !(0, runtime_1.exists)(json, 'profileType') ? undefined : (0, ProfileTypeType_1.ProfileTypeTypeFromJSON)(json['profileType']),
    };
}
exports.ProfileSubscriptionTypeProfileInfoFromJSONTyped = ProfileSubscriptionTypeProfileInfoFromJSONTyped;
function ProfileSubscriptionTypeProfileInfoToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'name': value.name,
        'fullName': value.fullName,
        'namePrefix': value.namePrefix,
        'givenName': value.givenName,
        'middleName': value.middleName,
        'nameSuffix': value.nameSuffix,
        'nameTitle': value.nameTitle,
        'nameType': (0, PersonNameTypeType_1.PersonNameTypeTypeToJSON)(value.nameType),
        'profileType': (0, ProfileTypeType_1.ProfileTypeTypeToJSON)(value.profileType),
    };
}
exports.ProfileSubscriptionTypeProfileInfoToJSON = ProfileSubscriptionTypeProfileInfoToJSON;
