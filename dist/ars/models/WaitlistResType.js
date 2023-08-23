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
exports.WaitlistResTypeToJSON = exports.WaitlistResTypeFromJSONTyped = exports.WaitlistResTypeFromJSON = exports.instanceOfWaitlistResType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the WaitlistResType interface.
 */
function instanceOfWaitlistResType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfWaitlistResType = instanceOfWaitlistResType;
function WaitlistResTypeFromJSON(json) {
    return WaitlistResTypeFromJSONTyped(json, false);
}
exports.WaitlistResTypeFromJSON = WaitlistResTypeFromJSON;
function WaitlistResTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reasonDescription': !(0, runtime_1.exists)(json, 'reasonDescription') ? undefined : json['reasonDescription'],
        'priorityDescription': !(0, runtime_1.exists)(json, 'priorityDescription') ? undefined : json['priorityDescription'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'reasonCode': !(0, runtime_1.exists)(json, 'reasonCode') ? undefined : json['reasonCode'],
        'priorityCode': !(0, runtime_1.exists)(json, 'priorityCode') ? undefined : json['priorityCode'],
        'telephone': !(0, runtime_1.exists)(json, 'telephone') ? undefined : json['telephone'],
    };
}
exports.WaitlistResTypeFromJSONTyped = WaitlistResTypeFromJSONTyped;
function WaitlistResTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reasonDescription': value.reasonDescription,
        'priorityDescription': value.priorityDescription,
        'description': value.description,
        'reasonCode': value.reasonCode,
        'priorityCode': value.priorityCode,
        'telephone': value.telephone,
    };
}
exports.WaitlistResTypeToJSON = WaitlistResTypeToJSON;
