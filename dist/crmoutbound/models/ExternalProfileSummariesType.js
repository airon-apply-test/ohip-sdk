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
exports.ExternalProfileSummariesTypeToJSON = exports.ExternalProfileSummariesTypeFromJSONTyped = exports.ExternalProfileSummariesTypeFromJSON = exports.instanceOfExternalProfileSummariesType = void 0;
const runtime_1 = require("../runtime");
const ExternalProfileSummaryInfoType_1 = require("./ExternalProfileSummaryInfoType");
const ExternalStatusType_1 = require("./ExternalStatusType");
/**
 * Check if a given object implements the ExternalProfileSummariesType interface.
 */
function instanceOfExternalProfileSummariesType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfExternalProfileSummariesType = instanceOfExternalProfileSummariesType;
function ExternalProfileSummariesTypeFromJSON(json) {
    return ExternalProfileSummariesTypeFromJSONTyped(json, false);
}
exports.ExternalProfileSummariesTypeFromJSON = ExternalProfileSummariesTypeFromJSON;
function ExternalProfileSummariesTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileInfo': !(0, runtime_1.exists)(json, 'profileInfo') ? undefined : (json['profileInfo'].map(ExternalProfileSummaryInfoType_1.ExternalProfileSummaryInfoTypeFromJSON)),
        'totalRows': !(0, runtime_1.exists)(json, 'totalRows') ? undefined : json['totalRows'],
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : (0, ExternalStatusType_1.ExternalStatusTypeFromJSON)(json['status']),
        'haltOperation': !(0, runtime_1.exists)(json, 'haltOperation') ? undefined : json['haltOperation'],
    };
}
exports.ExternalProfileSummariesTypeFromJSONTyped = ExternalProfileSummariesTypeFromJSONTyped;
function ExternalProfileSummariesTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileInfo': value.profileInfo === undefined ? undefined : (value.profileInfo.map(ExternalProfileSummaryInfoType_1.ExternalProfileSummaryInfoTypeToJSON)),
        'totalRows': value.totalRows,
        'status': (0, ExternalStatusType_1.ExternalStatusTypeToJSON)(value.status),
        'haltOperation': value.haltOperation,
    };
}
exports.ExternalProfileSummariesTypeToJSON = ExternalProfileSummariesTypeToJSON;
