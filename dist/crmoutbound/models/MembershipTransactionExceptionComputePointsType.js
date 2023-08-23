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
exports.MembershipTransactionExceptionComputePointsTypeToJSON = exports.MembershipTransactionExceptionComputePointsTypeFromJSONTyped = exports.MembershipTransactionExceptionComputePointsTypeFromJSON = exports.instanceOfMembershipTransactionExceptionComputePointsType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the MembershipTransactionExceptionComputePointsType interface.
 */
function instanceOfMembershipTransactionExceptionComputePointsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfMembershipTransactionExceptionComputePointsType = instanceOfMembershipTransactionExceptionComputePointsType;
function MembershipTransactionExceptionComputePointsTypeFromJSON(json) {
    return MembershipTransactionExceptionComputePointsTypeFromJSONTyped(json, false);
}
exports.MembershipTransactionExceptionComputePointsTypeFromJSON = MembershipTransactionExceptionComputePointsTypeFromJSON;
function MembershipTransactionExceptionComputePointsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'awardBasePoints': !(0, runtime_1.exists)(json, 'awardBasePoints') ? undefined : json['awardBasePoints'],
        'awardBonusPoints': !(0, runtime_1.exists)(json, 'awardBonusPoints') ? undefined : json['awardBonusPoints'],
        'tierPointsBaseStay': !(0, runtime_1.exists)(json, 'tierPointsBaseStay') ? undefined : json['tierPointsBaseStay'],
        'tierPointsBonusStay': !(0, runtime_1.exists)(json, 'tierPointsBonusStay') ? undefined : json['tierPointsBonusStay'],
        'tierPointsBaseNights': !(0, runtime_1.exists)(json, 'tierPointsBaseNights') ? undefined : json['tierPointsBaseNights'],
        'tierPointsBonusNights': !(0, runtime_1.exists)(json, 'tierPointsBonusNights') ? undefined : json['tierPointsBonusNights'],
        'tierPointsBaseRevenue': !(0, runtime_1.exists)(json, 'tierPointsBaseRevenue') ? undefined : json['tierPointsBaseRevenue'],
        'tierPointsBonusRevenue': !(0, runtime_1.exists)(json, 'tierPointsBonusRevenue') ? undefined : json['tierPointsBonusRevenue'],
    };
}
exports.MembershipTransactionExceptionComputePointsTypeFromJSONTyped = MembershipTransactionExceptionComputePointsTypeFromJSONTyped;
function MembershipTransactionExceptionComputePointsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'awardBasePoints': value.awardBasePoints,
        'awardBonusPoints': value.awardBonusPoints,
        'tierPointsBaseStay': value.tierPointsBaseStay,
        'tierPointsBonusStay': value.tierPointsBonusStay,
        'tierPointsBaseNights': value.tierPointsBaseNights,
        'tierPointsBonusNights': value.tierPointsBonusNights,
        'tierPointsBaseRevenue': value.tierPointsBaseRevenue,
        'tierPointsBonusRevenue': value.tierPointsBonusRevenue,
    };
}
exports.MembershipTransactionExceptionComputePointsTypeToJSON = MembershipTransactionExceptionComputePointsTypeToJSON;
