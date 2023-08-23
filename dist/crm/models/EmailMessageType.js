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
exports.EmailMessageTypeToJSON = exports.EmailMessageTypeFromJSONTyped = exports.EmailMessageTypeFromJSON = exports.instanceOfEmailMessageType = void 0;
const runtime_1 = require("../runtime");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the EmailMessageType interface.
 */
function instanceOfEmailMessageType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfEmailMessageType = instanceOfEmailMessageType;
function EmailMessageTypeFromJSON(json) {
    return EmailMessageTypeFromJSONTyped(json, false);
}
exports.EmailMessageTypeFromJSON = EmailMessageTypeFromJSON;
function EmailMessageTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'blockId': !(0, runtime_1.exists)(json, 'blockId') ? undefined : json['blockId'],
        'fromAddress': !(0, runtime_1.exists)(json, 'fromAddress') ? undefined : json['fromAddress'],
        'toAddress': !(0, runtime_1.exists)(json, 'toAddress') ? undefined : json['toAddress'],
        'subject': !(0, runtime_1.exists)(json, 'subject') ? undefined : json['subject'],
        'emailBody': !(0, runtime_1.exists)(json, 'emailBody') ? undefined : json['emailBody'],
        'messageId': !(0, runtime_1.exists)(json, 'messageId') ? undefined : json['messageId'],
        'emailReceiveDate': !(0, runtime_1.exists)(json, 'emailReceiveDate') ? undefined : (new Date(json['emailReceiveDate'])),
        'hasAttachment': !(0, runtime_1.exists)(json, 'hasAttachment') ? undefined : json['hasAttachment'],
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.EmailMessageTypeFromJSONTyped = EmailMessageTypeFromJSONTyped;
function EmailMessageTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'blockId': value.blockId,
        'fromAddress': value.fromAddress,
        'toAddress': value.toAddress,
        'subject': value.subject,
        'emailBody': value.emailBody,
        'messageId': value.messageId,
        'emailReceiveDate': value.emailReceiveDate === undefined ? undefined : (value.emailReceiveDate.toISOString()),
        'hasAttachment': value.hasAttachment,
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.EmailMessageTypeToJSON = EmailMessageTypeToJSON;
