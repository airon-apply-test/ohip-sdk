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
exports.AuthorizerCreditDetailTypeToJSON = exports.AuthorizerCreditDetailTypeFromJSONTyped = exports.AuthorizerCreditDetailTypeFromJSON = exports.instanceOfAuthorizerCreditDetailType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the AuthorizerCreditDetailType interface.
 */
function instanceOfAuthorizerCreditDetailType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAuthorizerCreditDetailType = instanceOfAuthorizerCreditDetailType;
function AuthorizerCreditDetailTypeFromJSON(json) {
    return AuthorizerCreditDetailTypeFromJSONTyped(json, false);
}
exports.AuthorizerCreditDetailTypeFromJSON = AuthorizerCreditDetailTypeFromJSON;
function AuthorizerCreditDetailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'creditLimit': !(0, runtime_1.exists)(json, 'creditLimit') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['creditLimit']),
        'actualAmount': !(0, runtime_1.exists)(json, 'actualAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['actualAmount']),
        'billingInstruction': !(0, runtime_1.exists)(json, 'billingInstruction') ? undefined : json['billingInstruction'],
    };
}
exports.AuthorizerCreditDetailTypeFromJSONTyped = AuthorizerCreditDetailTypeFromJSONTyped;
function AuthorizerCreditDetailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'creditLimit': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.creditLimit),
        'actualAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.actualAmount),
        'billingInstruction': value.billingInstruction,
    };
}
exports.AuthorizerCreditDetailTypeToJSON = AuthorizerCreditDetailTypeToJSON;
