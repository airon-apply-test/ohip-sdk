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
exports.ECertificateGenerationDetailsTypeToJSON = exports.ECertificateGenerationDetailsTypeFromJSONTyped = exports.ECertificateGenerationDetailsTypeFromJSON = exports.instanceOfECertificateGenerationDetailsType = void 0;
const runtime_1 = require("../runtime");
const ECertificateIssueSourceType_1 = require("./ECertificateIssueSourceType");
const ECertificateIssueType_1 = require("./ECertificateIssueType");
const ProfileId_1 = require("./ProfileId");
/**
 * Check if a given object implements the ECertificateGenerationDetailsType interface.
 */
function instanceOfECertificateGenerationDetailsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfECertificateGenerationDetailsType = instanceOfECertificateGenerationDetailsType;
function ECertificateGenerationDetailsTypeFromJSON(json) {
    return ECertificateGenerationDetailsTypeFromJSONTyped(json, false);
}
exports.ECertificateGenerationDetailsTypeFromJSON = ECertificateGenerationDetailsTypeFromJSON;
function ECertificateGenerationDetailsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileId': !(0, runtime_1.exists)(json, 'profileId') ? undefined : (0, ProfileId_1.ProfileIdFromJSON)(json['profileId']),
        'certificateType': !(0, runtime_1.exists)(json, 'certificateType') ? undefined : json['certificateType'],
        'source': !(0, runtime_1.exists)(json, 'source') ? undefined : (0, ECertificateIssueSourceType_1.ECertificateIssueSourceTypeFromJSON)(json['source']),
        'issueType': !(0, runtime_1.exists)(json, 'issueType') ? undefined : (0, ECertificateIssueType_1.ECertificateIssueTypeFromJSON)(json['issueType']),
        'hotels': !(0, runtime_1.exists)(json, 'hotels') ? undefined : json['hotels'],
    };
}
exports.ECertificateGenerationDetailsTypeFromJSONTyped = ECertificateGenerationDetailsTypeFromJSONTyped;
function ECertificateGenerationDetailsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileId': (0, ProfileId_1.ProfileIdToJSON)(value.profileId),
        'certificateType': value.certificateType,
        'source': (0, ECertificateIssueSourceType_1.ECertificateIssueSourceTypeToJSON)(value.source),
        'issueType': (0, ECertificateIssueType_1.ECertificateIssueTypeToJSON)(value.issueType),
        'hotels': value.hotels,
    };
}
exports.ECertificateGenerationDetailsTypeToJSON = ECertificateGenerationDetailsTypeToJSON;
