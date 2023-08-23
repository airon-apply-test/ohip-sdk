"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PolicyApplyCriteriaTypeToJSON = exports.PolicyApplyCriteriaTypeFromJSONTyped = exports.PolicyApplyCriteriaTypeFromJSON = exports.instanceOfPolicyApplyCriteriaType = void 0;
const runtime_1 = require("../runtime");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the PolicyApplyCriteriaType interface.
 */
function instanceOfPolicyApplyCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPolicyApplyCriteriaType = instanceOfPolicyApplyCriteriaType;
function PolicyApplyCriteriaTypeFromJSON(json) {
    return PolicyApplyCriteriaTypeFromJSONTyped(json, false);
}
exports.PolicyApplyCriteriaTypeFromJSON = PolicyApplyCriteriaTypeFromJSON;
function PolicyApplyCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'policyId': !(0, runtime_1.exists)(json, 'policyId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['policyId']),
        'transactionNo': !(0, runtime_1.exists)(json, 'transactionNo') ? undefined : json['transactionNo'],
    };
}
exports.PolicyApplyCriteriaTypeFromJSONTyped = PolicyApplyCriteriaTypeFromJSONTyped;
function PolicyApplyCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'policyId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.policyId),
        'transactionNo': value.transactionNo,
    };
}
exports.PolicyApplyCriteriaTypeToJSON = PolicyApplyCriteriaTypeToJSON;
