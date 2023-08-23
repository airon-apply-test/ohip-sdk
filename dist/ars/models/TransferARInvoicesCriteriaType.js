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
exports.TransferARInvoicesCriteriaTypeToJSON = exports.TransferARInvoicesCriteriaTypeFromJSONTyped = exports.TransferARInvoicesCriteriaTypeFromJSON = exports.instanceOfTransferARInvoicesCriteriaType = void 0;
const runtime_1 = require("../runtime");
const ARAccountCriteriaType_1 = require("./ARAccountCriteriaType");
const ARInvoiceType_1 = require("./ARInvoiceType");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the TransferARInvoicesCriteriaType interface.
 */
function instanceOfTransferARInvoicesCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTransferARInvoicesCriteriaType = instanceOfTransferARInvoicesCriteriaType;
function TransferARInvoicesCriteriaTypeFromJSON(json) {
    return TransferARInvoicesCriteriaTypeFromJSONTyped(json, false);
}
exports.TransferARInvoicesCriteriaTypeFromJSON = TransferARInvoicesCriteriaTypeFromJSON;
function TransferARInvoicesCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'invoicesTransferType': !(0, runtime_1.exists)(json, 'invoicesTransferType') ? undefined : (json['invoicesTransferType'].map(ARInvoiceType_1.ARInvoiceTypeFromJSON)),
        'fromAccount': !(0, runtime_1.exists)(json, 'fromAccount') ? undefined : (0, ARAccountCriteriaType_1.ARAccountCriteriaTypeFromJSON)(json['fromAccount']),
        'toAccount': !(0, runtime_1.exists)(json, 'toAccount') ? undefined : (0, ARAccountCriteriaType_1.ARAccountCriteriaTypeFromJSON)(json['toAccount']),
        'transferAmount': !(0, runtime_1.exists)(json, 'transferAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['transferAmount']),
        'remarks': !(0, runtime_1.exists)(json, 'remarks') ? undefined : json['remarks'],
        'overrideCreditHoldCheck': !(0, runtime_1.exists)(json, 'overrideCreditHoldCheck') ? undefined : json['overrideCreditHoldCheck'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
    };
}
exports.TransferARInvoicesCriteriaTypeFromJSONTyped = TransferARInvoicesCriteriaTypeFromJSONTyped;
function TransferARInvoicesCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'invoicesTransferType': value.invoicesTransferType === undefined ? undefined : (value.invoicesTransferType.map(ARInvoiceType_1.ARInvoiceTypeToJSON)),
        'fromAccount': (0, ARAccountCriteriaType_1.ARAccountCriteriaTypeToJSON)(value.fromAccount),
        'toAccount': (0, ARAccountCriteriaType_1.ARAccountCriteriaTypeToJSON)(value.toAccount),
        'transferAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.transferAmount),
        'remarks': value.remarks,
        'overrideCreditHoldCheck': value.overrideCreditHoldCheck,
        'cashierId': value.cashierId,
    };
}
exports.TransferARInvoicesCriteriaTypeToJSON = TransferARInvoicesCriteriaTypeToJSON;
