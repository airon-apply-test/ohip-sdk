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
exports.ARInvoiceCriteriaTypeAdditionalFilterToJSON = exports.ARInvoiceCriteriaTypeAdditionalFilterFromJSONTyped = exports.ARInvoiceCriteriaTypeAdditionalFilterFromJSON = exports.instanceOfARInvoiceCriteriaTypeAdditionalFilter = void 0;
const runtime_1 = require("../runtime");
const DateRangeType_1 = require("./DateRangeType");
/**
 * Check if a given object implements the ARInvoiceCriteriaTypeAdditionalFilter interface.
 */
function instanceOfARInvoiceCriteriaTypeAdditionalFilter(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfARInvoiceCriteriaTypeAdditionalFilter = instanceOfARInvoiceCriteriaTypeAdditionalFilter;
function ARInvoiceCriteriaTypeAdditionalFilterFromJSON(json) {
    return ARInvoiceCriteriaTypeAdditionalFilterFromJSONTyped(json, false);
}
exports.ARInvoiceCriteriaTypeAdditionalFilterFromJSON = ARInvoiceCriteriaTypeAdditionalFilterFromJSON;
function ARInvoiceCriteriaTypeAdditionalFilterFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
        'transactionCode': !(0, runtime_1.exists)(json, 'transactionCode') ? undefined : json['transactionCode'],
        'transactionSubGroup': !(0, runtime_1.exists)(json, 'transactionSubGroup') ? undefined : json['transactionSubGroup'],
        'dateRange': !(0, runtime_1.exists)(json, 'dateRange') ? undefined : (0, DateRangeType_1.DateRangeTypeFromJSON)(json['dateRange']),
        'referenceWildCard': !(0, runtime_1.exists)(json, 'referenceWildCard') ? undefined : json['referenceWildCard'],
        'supplementWildCard': !(0, runtime_1.exists)(json, 'supplementWildCard') ? undefined : json['supplementWildCard'],
        'checkNumberWildCard': !(0, runtime_1.exists)(json, 'checkNumberWildCard') ? undefined : json['checkNumberWildCard'],
    };
}
exports.ARInvoiceCriteriaTypeAdditionalFilterFromJSONTyped = ARInvoiceCriteriaTypeAdditionalFilterFromJSONTyped;
function ARInvoiceCriteriaTypeAdditionalFilterToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'cashierId': value.cashierId,
        'transactionCode': value.transactionCode,
        'transactionSubGroup': value.transactionSubGroup,
        'dateRange': (0, DateRangeType_1.DateRangeTypeToJSON)(value.dateRange),
        'referenceWildCard': value.referenceWildCard,
        'supplementWildCard': value.supplementWildCard,
        'checkNumberWildCard': value.checkNumberWildCard,
    };
}
exports.ARInvoiceCriteriaTypeAdditionalFilterToJSON = ARInvoiceCriteriaTypeAdditionalFilterToJSON;
