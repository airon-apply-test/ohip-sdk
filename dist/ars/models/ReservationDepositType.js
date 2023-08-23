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
exports.ReservationDepositTypeToJSON = exports.ReservationDepositTypeFromJSONTyped = exports.ReservationDepositTypeFromJSON = exports.instanceOfReservationDepositType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ReservationDepositType interface.
 */
function instanceOfReservationDepositType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationDepositType = instanceOfReservationDepositType;
function ReservationDepositTypeFromJSON(json) {
    return ReservationDepositTypeFromJSONTyped(json, false);
}
exports.ReservationDepositTypeFromJSON = ReservationDepositTypeFromJSON;
function ReservationDepositTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'amountRequired': !(0, runtime_1.exists)(json, 'amountRequired') ? undefined : json['amountRequired'],
        'amountPaid': !(0, runtime_1.exists)(json, 'amountPaid') ? undefined : json['amountPaid'],
        'amountOwed': !(0, runtime_1.exists)(json, 'amountOwed') ? undefined : json['amountOwed'],
        'dueDate': !(0, runtime_1.exists)(json, 'dueDate') ? undefined : (new Date(json['dueDate'])),
        'postingDate': !(0, runtime_1.exists)(json, 'postingDate') ? undefined : (new Date(json['postingDate'])),
        'hasPaid': !(0, runtime_1.exists)(json, 'hasPaid') ? undefined : json['hasPaid'],
        'hasOutstanding': !(0, runtime_1.exists)(json, 'hasOutstanding') ? undefined : json['hasOutstanding'],
    };
}
exports.ReservationDepositTypeFromJSONTyped = ReservationDepositTypeFromJSONTyped;
function ReservationDepositTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'amountRequired': value.amountRequired,
        'amountPaid': value.amountPaid,
        'amountOwed': value.amountOwed,
        'dueDate': value.dueDate === undefined ? undefined : (value.dueDate.toISOString().substr(0, 10)),
        'postingDate': value.postingDate === undefined ? undefined : (value.postingDate.toISOString().substr(0, 10)),
        'hasPaid': value.hasPaid,
        'hasOutstanding': value.hasOutstanding,
    };
}
exports.ReservationDepositTypeToJSON = ReservationDepositTypeToJSON;
