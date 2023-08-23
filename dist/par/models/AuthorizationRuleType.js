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
exports.AuthorizationRuleTypeToJSON = exports.AuthorizationRuleTypeFromJSONTyped = exports.AuthorizationRuleTypeFromJSON = exports.instanceOfAuthorizationRuleType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the AuthorizationRuleType interface.
 */
function instanceOfAuthorizationRuleType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAuthorizationRuleType = instanceOfAuthorizationRuleType;
function AuthorizationRuleTypeFromJSON(json) {
    return AuthorizationRuleTypeFromJSONTyped(json, false);
}
exports.AuthorizationRuleTypeFromJSON = AuthorizationRuleTypeFromJSON;
function AuthorizationRuleTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'percent': !(0, runtime_1.exists)(json, 'percent') ? undefined : json['percent'],
    };
}
exports.AuthorizationRuleTypeFromJSONTyped = AuthorizationRuleTypeFromJSONTyped;
function AuthorizationRuleTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'code': value.code,
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'percent': value.percent,
    };
}
exports.AuthorizationRuleTypeToJSON = AuthorizationRuleTypeToJSON;
