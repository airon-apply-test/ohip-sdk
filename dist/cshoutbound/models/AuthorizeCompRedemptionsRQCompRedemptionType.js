"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.AuthorizeCompRedemptionsRQCompRedemptionTypeToJSON = exports.AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSONTyped = exports.AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSON = exports.instanceOfAuthorizeCompRedemptionsRQCompRedemptionType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the AuthorizeCompRedemptionsRQCompRedemptionType interface.
 */
function instanceOfAuthorizeCompRedemptionsRQCompRedemptionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAuthorizeCompRedemptionsRQCompRedemptionType = instanceOfAuthorizeCompRedemptionsRQCompRedemptionType;
function AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSON(json) {
    return AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSONTyped(json, false);
}
exports.AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSON = AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSON;
function AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'availableAmount': !(0, runtime_1.exists)(json, 'availableAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['availableAmount']),
    };
}
exports.AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSONTyped = AuthorizeCompRedemptionsRQCompRedemptionTypeFromJSONTyped;
function AuthorizeCompRedemptionsRQCompRedemptionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'code': value.code,
        'description': value.description,
        'availableAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.availableAmount),
    };
}
exports.AuthorizeCompRedemptionsRQCompRedemptionTypeToJSON = AuthorizeCompRedemptionsRQCompRedemptionTypeToJSON;
