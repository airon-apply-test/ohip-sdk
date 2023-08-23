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
exports.TransferARCreditCardPaymentsTypeToJSON = exports.TransferARCreditCardPaymentsTypeFromJSONTyped = exports.TransferARCreditCardPaymentsTypeFromJSON = exports.instanceOfTransferARCreditCardPaymentsType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the TransferARCreditCardPaymentsType interface.
 */
function instanceOfTransferARCreditCardPaymentsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTransferARCreditCardPaymentsType = instanceOfTransferARCreditCardPaymentsType;
function TransferARCreditCardPaymentsTypeFromJSON(json) {
    return TransferARCreditCardPaymentsTypeFromJSONTyped(json, false);
}
exports.TransferARCreditCardPaymentsTypeFromJSON = TransferARCreditCardPaymentsTypeFromJSON;
function TransferARCreditCardPaymentsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'accountId': !(0, runtime_1.exists)(json, 'accountId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['accountId']),
        'totalAmount': !(0, runtime_1.exists)(json, 'totalAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['totalAmount']),
        'reference': !(0, runtime_1.exists)(json, 'reference') ? undefined : json['reference'],
        'transactionNo': !(0, runtime_1.exists)(json, 'transactionNo') ? undefined : json['transactionNo'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
    };
}
exports.TransferARCreditCardPaymentsTypeFromJSONTyped = TransferARCreditCardPaymentsTypeFromJSONTyped;
function TransferARCreditCardPaymentsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'accountId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.accountId),
        'totalAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.totalAmount),
        'reference': value.reference,
        'transactionNo': value.transactionNo,
        'cashierId': value.cashierId,
    };
}
exports.TransferARCreditCardPaymentsTypeToJSON = TransferARCreditCardPaymentsTypeToJSON;
