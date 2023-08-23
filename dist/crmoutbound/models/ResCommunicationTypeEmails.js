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
exports.ResCommunicationTypeEmailsToJSON = exports.ResCommunicationTypeEmailsFromJSONTyped = exports.ResCommunicationTypeEmailsFromJSON = exports.instanceOfResCommunicationTypeEmails = void 0;
const runtime_1 = require("../runtime");
const EmailInfoType_1 = require("./EmailInfoType");
/**
 * Check if a given object implements the ResCommunicationTypeEmails interface.
 */
function instanceOfResCommunicationTypeEmails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResCommunicationTypeEmails = instanceOfResCommunicationTypeEmails;
function ResCommunicationTypeEmailsFromJSON(json) {
    return ResCommunicationTypeEmailsFromJSONTyped(json, false);
}
exports.ResCommunicationTypeEmailsFromJSON = ResCommunicationTypeEmailsFromJSON;
function ResCommunicationTypeEmailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'emailInfo': !(0, runtime_1.exists)(json, 'emailInfo') ? undefined : (json['emailInfo'].map(EmailInfoType_1.EmailInfoTypeFromJSON)),
        'allRowsFetched': !(0, runtime_1.exists)(json, 'allRowsFetched') ? undefined : json['allRowsFetched'],
        'totalRows': !(0, runtime_1.exists)(json, 'totalRows') ? undefined : json['totalRows'],
    };
}
exports.ResCommunicationTypeEmailsFromJSONTyped = ResCommunicationTypeEmailsFromJSONTyped;
function ResCommunicationTypeEmailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'emailInfo': value.emailInfo === undefined ? undefined : (value.emailInfo.map(EmailInfoType_1.EmailInfoTypeToJSON)),
        'allRowsFetched': value.allRowsFetched,
        'totalRows': value.totalRows,
    };
}
exports.ResCommunicationTypeEmailsToJSON = ResCommunicationTypeEmailsToJSON;
