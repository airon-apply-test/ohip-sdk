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
exports.ChangeARInvoicesAccountCriteriaTypeToJSON = exports.ChangeARInvoicesAccountCriteriaTypeFromJSONTyped = exports.ChangeARInvoicesAccountCriteriaTypeFromJSON = exports.instanceOfChangeARInvoicesAccountCriteriaType = void 0;
const runtime_1 = require("../runtime");
const ARAccountCriteriaType_1 = require("./ARAccountCriteriaType");
const ARInvoiceType_1 = require("./ARInvoiceType");
/**
 * Check if a given object implements the ChangeARInvoicesAccountCriteriaType interface.
 */
function instanceOfChangeARInvoicesAccountCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChangeARInvoicesAccountCriteriaType = instanceOfChangeARInvoicesAccountCriteriaType;
function ChangeARInvoicesAccountCriteriaTypeFromJSON(json) {
    return ChangeARInvoicesAccountCriteriaTypeFromJSONTyped(json, false);
}
exports.ChangeARInvoicesAccountCriteriaTypeFromJSON = ChangeARInvoicesAccountCriteriaTypeFromJSON;
function ChangeARInvoicesAccountCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'invoices': !(0, runtime_1.exists)(json, 'invoices') ? undefined : (json['invoices'].map(ARInvoiceType_1.ARInvoiceTypeFromJSON)),
        'toAccount': !(0, runtime_1.exists)(json, 'toAccount') ? undefined : (0, ARAccountCriteriaType_1.ARAccountCriteriaTypeFromJSON)(json['toAccount']),
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
    };
}
exports.ChangeARInvoicesAccountCriteriaTypeFromJSONTyped = ChangeARInvoicesAccountCriteriaTypeFromJSONTyped;
function ChangeARInvoicesAccountCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'invoices': value.invoices === undefined ? undefined : (value.invoices.map(ARInvoiceType_1.ARInvoiceTypeToJSON)),
        'toAccount': (0, ARAccountCriteriaType_1.ARAccountCriteriaTypeToJSON)(value.toAccount),
        'cashierId': value.cashierId,
    };
}
exports.ChangeARInvoicesAccountCriteriaTypeToJSON = ChangeARInvoicesAccountCriteriaTypeToJSON;
