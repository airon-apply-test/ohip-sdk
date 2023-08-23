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
exports.TransactionCurrencyExchangeInfoTypeToJSON = exports.TransactionCurrencyExchangeInfoTypeFromJSONTyped = exports.TransactionCurrencyExchangeInfoTypeFromJSON = exports.instanceOfTransactionCurrencyExchangeInfoType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the TransactionCurrencyExchangeInfoType interface.
 */
function instanceOfTransactionCurrencyExchangeInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTransactionCurrencyExchangeInfoType = instanceOfTransactionCurrencyExchangeInfoType;
function TransactionCurrencyExchangeInfoTypeFromJSON(json) {
    return TransactionCurrencyExchangeInfoTypeFromJSONTyped(json, false);
}
exports.TransactionCurrencyExchangeInfoTypeFromJSON = TransactionCurrencyExchangeInfoTypeFromJSON;
function TransactionCurrencyExchangeInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'exchangeRate': !(0, runtime_1.exists)(json, 'exchangeRate') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['exchangeRate']),
        'currencyAmount': !(0, runtime_1.exists)(json, 'currencyAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['currencyAmount']),
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'commissionPercent': !(0, runtime_1.exists)(json, 'commissionPercent') ? undefined : json['commissionPercent'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'exchangeDate': !(0, runtime_1.exists)(json, 'exchangeDate') ? undefined : (new Date(json['exchangeDate'])),
    };
}
exports.TransactionCurrencyExchangeInfoTypeFromJSONTyped = TransactionCurrencyExchangeInfoTypeFromJSONTyped;
function TransactionCurrencyExchangeInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'exchangeRate': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.exchangeRate),
        'currencyAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.currencyAmount),
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'commissionPercent': value.commissionPercent,
        'code': value.code,
        'exchangeDate': value.exchangeDate === undefined ? undefined : (value.exchangeDate.toISOString().substr(0, 10)),
    };
}
exports.TransactionCurrencyExchangeInfoTypeToJSON = TransactionCurrencyExchangeInfoTypeToJSON;
