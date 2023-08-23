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
exports.VIPTypeToJSON = exports.VIPTypeFromJSONTyped = exports.VIPTypeFromJSON = exports.instanceOfVIPType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the VIPType interface.
 */
function instanceOfVIPType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfVIPType = instanceOfVIPType;
function VIPTypeFromJSON(json) {
    return VIPTypeFromJSONTyped(json, false);
}
exports.VIPTypeFromJSON = VIPTypeFromJSON;
function VIPTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'vipCode': !(0, runtime_1.exists)(json, 'vipCode') ? undefined : json['vipCode'],
        'vipDescription': !(0, runtime_1.exists)(json, 'vipDescription') ? undefined : json['vipDescription'],
    };
}
exports.VIPTypeFromJSONTyped = VIPTypeFromJSONTyped;
function VIPTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'vipCode': value.vipCode,
        'vipDescription': value.vipDescription,
    };
}
exports.VIPTypeToJSON = VIPTypeToJSON;
