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
exports.InvoicePostingsToJSON = exports.InvoicePostingsFromJSONTyped = exports.InvoicePostingsFromJSON = exports.instanceOfInvoicePostings = void 0;
const runtime_1 = require("../runtime");
const ARInvoicePostingType_1 = require("./ARInvoicePostingType");
const InstanceLink_1 = require("./InstanceLink");
const TrxInfoType_1 = require("./TrxInfoType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the InvoicePostings interface.
 */
function instanceOfInvoicePostings(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfInvoicePostings = instanceOfInvoicePostings;
function InvoicePostingsFromJSON(json) {
    return InvoicePostingsFromJSONTyped(json, false);
}
exports.InvoicePostingsFromJSON = InvoicePostingsFromJSON;
function InvoicePostingsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'invoicePostingsDetails': !(0, runtime_1.exists)(json, 'invoicePostingsDetails') ? undefined : (json['invoicePostingsDetails'].map(ARInvoicePostingType_1.ARInvoicePostingTypeFromJSON)),
        'trxCodesInfo': !(0, runtime_1.exists)(json, 'trxCodesInfo') ? undefined : (json['trxCodesInfo'].map(TrxInfoType_1.TrxInfoTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.InvoicePostingsFromJSONTyped = InvoicePostingsFromJSONTyped;
function InvoicePostingsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'invoicePostingsDetails': value.invoicePostingsDetails === undefined ? undefined : (value.invoicePostingsDetails.map(ARInvoicePostingType_1.ARInvoicePostingTypeToJSON)),
        'trxCodesInfo': value.trxCodesInfo === undefined ? undefined : (value.trxCodesInfo.map(TrxInfoType_1.TrxInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.InvoicePostingsToJSON = InvoicePostingsToJSON;
