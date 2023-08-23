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
exports.TransferARCreditCardPaymentsRequestToJSON = exports.TransferARCreditCardPaymentsRequestFromJSONTyped = exports.TransferARCreditCardPaymentsRequestFromJSON = exports.instanceOfTransferARCreditCardPaymentsRequest = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const TransferARCreditCardPaymentsType_1 = require("./TransferARCreditCardPaymentsType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the TransferARCreditCardPaymentsRequest interface.
 */
function instanceOfTransferARCreditCardPaymentsRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTransferARCreditCardPaymentsRequest = instanceOfTransferARCreditCardPaymentsRequest;
function TransferARCreditCardPaymentsRequestFromJSON(json) {
    return TransferARCreditCardPaymentsRequestFromJSONTyped(json, false);
}
exports.TransferARCreditCardPaymentsRequestFromJSON = TransferARCreditCardPaymentsRequestFromJSON;
function TransferARCreditCardPaymentsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'criteria': !(0, runtime_1.exists)(json, 'criteria') ? undefined : (0, TransferARCreditCardPaymentsType_1.TransferARCreditCardPaymentsTypeFromJSON)(json['criteria']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.TransferARCreditCardPaymentsRequestFromJSONTyped = TransferARCreditCardPaymentsRequestFromJSONTyped;
function TransferARCreditCardPaymentsRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'criteria': (0, TransferARCreditCardPaymentsType_1.TransferARCreditCardPaymentsTypeToJSON)(value.criteria),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.TransferARCreditCardPaymentsRequestToJSON = TransferARCreditCardPaymentsRequestToJSON;
