"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProfileTypeEmailsToJSON = exports.ProfileTypeEmailsFromJSONTyped = exports.ProfileTypeEmailsFromJSON = exports.instanceOfProfileTypeEmails = void 0;
const runtime_1 = require("../runtime");
const EmailInfoType_1 = require("./EmailInfoType");
/**
 * Check if a given object implements the ProfileTypeEmails interface.
 */
function instanceOfProfileTypeEmails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileTypeEmails = instanceOfProfileTypeEmails;
function ProfileTypeEmailsFromJSON(json) {
    return ProfileTypeEmailsFromJSONTyped(json, false);
}
exports.ProfileTypeEmailsFromJSON = ProfileTypeEmailsFromJSON;
function ProfileTypeEmailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'emailInfo': !(0, runtime_1.exists)(json, 'emailInfo') ? undefined : (json['emailInfo'].map(EmailInfoType_1.EmailInfoTypeFromJSON)),
    };
}
exports.ProfileTypeEmailsFromJSONTyped = ProfileTypeEmailsFromJSONTyped;
function ProfileTypeEmailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'emailInfo': value.emailInfo === undefined ? undefined : (value.emailInfo.map(EmailInfoType_1.EmailInfoTypeToJSON)),
    };
}
exports.ProfileTypeEmailsToJSON = ProfileTypeEmailsToJSON;
