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
exports.CashierReportTypeToJSON = exports.CashierReportTypeFromJSONTyped = exports.CashierReportTypeFromJSON = exports.CashierReportType = void 0;
/**
 * Shift drop report created when cashier closure is completed. It is an integral part of closure report and is not to be printed or accessed by itself.
 * @export
 */
exports.CashierReportType = {
    Cash: 'Cash',
    Check: 'Check',
    ForeignCurrency: 'ForeignCurrency',
    CreditCard: 'CreditCard',
    Miscellaneous: 'Miscellaneous',
    ArSettlements: 'ArSettlements',
    DepositTransfers: 'DepositTransfers',
    Shiftdrop: 'Shiftdrop'
};
function CashierReportTypeFromJSON(json) {
    return CashierReportTypeFromJSONTyped(json, false);
}
exports.CashierReportTypeFromJSON = CashierReportTypeFromJSON;
function CashierReportTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.CashierReportTypeFromJSONTyped = CashierReportTypeFromJSONTyped;
function CashierReportTypeToJSON(value) {
    return value;
}
exports.CashierReportTypeToJSON = CashierReportTypeToJSON;
