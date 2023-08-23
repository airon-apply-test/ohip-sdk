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
exports.PostChargesToARRequestToJSON = exports.PostChargesToARRequestFromJSONTyped = exports.PostChargesToARRequestFromJSON = exports.instanceOfPostChargesToARRequest = void 0;
const runtime_1 = require("../runtime");
const ARChargesPostingCriteriaType_1 = require("./ARChargesPostingCriteriaType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostChargesToARRequest interface.
 */
function instanceOfPostChargesToARRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostChargesToARRequest = instanceOfPostChargesToARRequest;
function PostChargesToARRequestFromJSON(json) {
    return PostChargesToARRequestFromJSONTyped(json, false);
}
exports.PostChargesToARRequestFromJSON = PostChargesToARRequestFromJSON;
function PostChargesToARRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'criteria': !(0, runtime_1.exists)(json, 'criteria') ? undefined : (0, ARChargesPostingCriteriaType_1.ARChargesPostingCriteriaTypeFromJSON)(json['criteria']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostChargesToARRequestFromJSONTyped = PostChargesToARRequestFromJSONTyped;
function PostChargesToARRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'criteria': (0, ARChargesPostingCriteriaType_1.ARChargesPostingCriteriaTypeToJSON)(value.criteria),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostChargesToARRequestToJSON = PostChargesToARRequestToJSON;
