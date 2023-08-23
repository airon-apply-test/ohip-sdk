"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ResSharedGuestInfoTypeToJSON = exports.ResSharedGuestInfoTypeFromJSONTyped = exports.ResSharedGuestInfoTypeFromJSON = exports.instanceOfResSharedGuestInfoType = void 0;
const runtime_1 = require("../runtime");
const ProfileId_1 = require("./ProfileId");
/**
 * Check if a given object implements the ResSharedGuestInfoType interface.
 */
function instanceOfResSharedGuestInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResSharedGuestInfoType = instanceOfResSharedGuestInfoType;
function ResSharedGuestInfoTypeFromJSON(json) {
    return ResSharedGuestInfoTypeFromJSONTyped(json, false);
}
exports.ResSharedGuestInfoTypeFromJSON = ResSharedGuestInfoTypeFromJSON;
function ResSharedGuestInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileId': !(0, runtime_1.exists)(json, 'profileId') ? undefined : (0, ProfileId_1.ProfileIdFromJSON)(json['profileId']),
        'firstName': !(0, runtime_1.exists)(json, 'firstName') ? undefined : json['firstName'],
        'lastName': !(0, runtime_1.exists)(json, 'lastName') ? undefined : json['lastName'],
        'fullName': !(0, runtime_1.exists)(json, 'fullName') ? undefined : json['fullName'],
    };
}
exports.ResSharedGuestInfoTypeFromJSONTyped = ResSharedGuestInfoTypeFromJSONTyped;
function ResSharedGuestInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileId': (0, ProfileId_1.ProfileIdToJSON)(value.profileId),
        'firstName': value.firstName,
        'lastName': value.lastName,
        'fullName': value.fullName,
    };
}
exports.ResSharedGuestInfoTypeToJSON = ResSharedGuestInfoTypeToJSON;
