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
exports.ReversePaymentCriteriaTypeToJSON = exports.ReversePaymentCriteriaTypeFromJSONTyped = exports.ReversePaymentCriteriaTypeFromJSON = exports.instanceOfReversePaymentCriteriaType = void 0;
const runtime_1 = require("../runtime");
const ProfileId_1 = require("./ProfileId");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ReversePaymentCriteriaType interface.
 */
function instanceOfReversePaymentCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReversePaymentCriteriaType = instanceOfReversePaymentCriteriaType;
function ReversePaymentCriteriaTypeFromJSON(json) {
    return ReversePaymentCriteriaTypeFromJSONTyped(json, false);
}
exports.ReversePaymentCriteriaTypeFromJSON = ReversePaymentCriteriaTypeFromJSON;
function ReversePaymentCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'profileId': !(0, runtime_1.exists)(json, 'profileId') ? undefined : (0, ProfileId_1.ProfileIdFromJSON)(json['profileId']),
        'accountId': !(0, runtime_1.exists)(json, 'accountId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['accountId']),
        'accountNo': !(0, runtime_1.exists)(json, 'accountNo') ? undefined : json['accountNo'],
        'accountName': !(0, runtime_1.exists)(json, 'accountName') ? undefined : json['accountName'],
        'transactionNo': !(0, runtime_1.exists)(json, 'transactionNo') ? undefined : json['transactionNo'],
    };
}
exports.ReversePaymentCriteriaTypeFromJSONTyped = ReversePaymentCriteriaTypeFromJSONTyped;
function ReversePaymentCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'profileId': (0, ProfileId_1.ProfileIdToJSON)(value.profileId),
        'accountId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.accountId),
        'accountNo': value.accountNo,
        'accountName': value.accountName,
        'transactionNo': value.transactionNo,
    };
}
exports.ReversePaymentCriteriaTypeToJSON = ReversePaymentCriteriaTypeToJSON;
