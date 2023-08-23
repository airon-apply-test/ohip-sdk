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
exports.ExchangeAmountsToJSON = exports.ExchangeAmountsFromJSONTyped = exports.ExchangeAmountsFromJSON = exports.instanceOfExchangeAmounts = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the ExchangeAmounts interface.
 */
function instanceOfExchangeAmounts(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfExchangeAmounts = instanceOfExchangeAmounts;
function ExchangeAmountsFromJSON(json) {
    return ExchangeAmountsFromJSONTyped(json, false);
}
exports.ExchangeAmountsFromJSON = ExchangeAmountsFromJSON;
function ExchangeAmountsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'debitAmount': !(0, runtime_1.exists)(json, 'debitAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['debitAmount']),
        'creditAmount': !(0, runtime_1.exists)(json, 'creditAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['creditAmount']),
    };
}
exports.ExchangeAmountsFromJSONTyped = ExchangeAmountsFromJSONTyped;
function ExchangeAmountsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'debitAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.debitAmount),
        'creditAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.creditAmount),
    };
}
exports.ExchangeAmountsToJSON = ExchangeAmountsToJSON;
