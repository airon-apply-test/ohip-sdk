"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.GuaranteePolicyTypeCancelPenaltyToJSON = exports.GuaranteePolicyTypeCancelPenaltyFromJSONTyped = exports.GuaranteePolicyTypeCancelPenaltyFromJSON = exports.instanceOfGuaranteePolicyTypeCancelPenalty = void 0;
const runtime_1 = require("../runtime");
const OffsetUnitType_1 = require("./OffsetUnitType");
const PolicyAmountPercentType_1 = require("./PolicyAmountPercentType");
const PolicyDeadlineType_1 = require("./PolicyDeadlineType");
/**
 * Check if a given object implements the GuaranteePolicyTypeCancelPenalty interface.
 */
function instanceOfGuaranteePolicyTypeCancelPenalty(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfGuaranteePolicyTypeCancelPenalty = instanceOfGuaranteePolicyTypeCancelPenalty;
function GuaranteePolicyTypeCancelPenaltyFromJSON(json) {
    return GuaranteePolicyTypeCancelPenaltyFromJSONTyped(json, false);
}
exports.GuaranteePolicyTypeCancelPenaltyFromJSON = GuaranteePolicyTypeCancelPenaltyFromJSON;
function GuaranteePolicyTypeCancelPenaltyFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'deadline': !(0, runtime_1.exists)(json, 'deadline') ? undefined : (0, PolicyDeadlineType_1.PolicyDeadlineTypeFromJSON)(json['deadline']),
        'amountPercent': !(0, runtime_1.exists)(json, 'amountPercent') ? undefined : (0, PolicyAmountPercentType_1.PolicyAmountPercentTypeFromJSON)(json['amountPercent']),
        'penaltyDescription': !(0, runtime_1.exists)(json, 'penaltyDescription') ? undefined : json['penaltyDescription'],
        'offsetUnit': !(0, runtime_1.exists)(json, 'offsetUnit') ? undefined : (0, OffsetUnitType_1.OffsetUnitTypeFromJSON)(json['offsetUnit']),
        'formattedRule': !(0, runtime_1.exists)(json, 'formattedRule') ? undefined : json['formattedRule'],
        'policyCode': !(0, runtime_1.exists)(json, 'policyCode') ? undefined : json['policyCode'],
        'manual': !(0, runtime_1.exists)(json, 'manual') ? undefined : json['manual'],
        'nonRefundable': !(0, runtime_1.exists)(json, 'nonRefundable') ? undefined : json['nonRefundable'],
        'effective': !(0, runtime_1.exists)(json, 'effective') ? undefined : json['effective'],
        'estimatedAmount': !(0, runtime_1.exists)(json, 'estimatedAmount') ? undefined : json['estimatedAmount'],
    };
}
exports.GuaranteePolicyTypeCancelPenaltyFromJSONTyped = GuaranteePolicyTypeCancelPenaltyFromJSONTyped;
function GuaranteePolicyTypeCancelPenaltyToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'deadline': (0, PolicyDeadlineType_1.PolicyDeadlineTypeToJSON)(value.deadline),
        'amountPercent': (0, PolicyAmountPercentType_1.PolicyAmountPercentTypeToJSON)(value.amountPercent),
        'penaltyDescription': value.penaltyDescription,
        'offsetUnit': (0, OffsetUnitType_1.OffsetUnitTypeToJSON)(value.offsetUnit),
        'formattedRule': value.formattedRule,
        'policyCode': value.policyCode,
        'manual': value.manual,
        'nonRefundable': value.nonRefundable,
        'effective': value.effective,
        'estimatedAmount': value.estimatedAmount,
    };
}
exports.GuaranteePolicyTypeCancelPenaltyToJSON = GuaranteePolicyTypeCancelPenaltyToJSON;
