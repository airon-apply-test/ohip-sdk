"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommissionCodesSummaryToJSON = exports.CommissionCodesSummaryFromJSONTyped = exports.CommissionCodesSummaryFromJSON = exports.instanceOfCommissionCodesSummary = void 0;
const runtime_1 = require("../runtime");
const CommissionCodeSummaryInfoType_1 = require("./CommissionCodeSummaryInfoType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the CommissionCodesSummary interface.
 */
function instanceOfCommissionCodesSummary(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCommissionCodesSummary = instanceOfCommissionCodesSummary;
function CommissionCodesSummaryFromJSON(json) {
    return CommissionCodesSummaryFromJSONTyped(json, false);
}
exports.CommissionCodesSummaryFromJSON = CommissionCodesSummaryFromJSON;
function CommissionCodesSummaryFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'commissionCodesSummary': !(0, runtime_1.exists)(json, 'commissionCodesSummary') ? undefined : (json['commissionCodesSummary'].map(CommissionCodeSummaryInfoType_1.CommissionCodeSummaryInfoTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.CommissionCodesSummaryFromJSONTyped = CommissionCodesSummaryFromJSONTyped;
function CommissionCodesSummaryToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'commissionCodesSummary': value.commissionCodesSummary === undefined ? undefined : (value.commissionCodesSummary.map(CommissionCodeSummaryInfoType_1.CommissionCodeSummaryInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.CommissionCodesSummaryToJSON = CommissionCodesSummaryToJSON;
