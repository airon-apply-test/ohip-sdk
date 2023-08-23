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
exports.NameValueHeaderDetailTypeToJSON = exports.NameValueHeaderDetailTypeFromJSONTyped = exports.NameValueHeaderDetailTypeFromJSON = exports.instanceOfNameValueHeaderDetailType = void 0;
const runtime_1 = require("../runtime");
const NameValueBaseSearchType_1 = require("./NameValueBaseSearchType");
const NameValueDetailType_1 = require("./NameValueDetailType");
/**
 * Check if a given object implements the NameValueHeaderDetailType interface.
 */
function instanceOfNameValueHeaderDetailType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfNameValueHeaderDetailType = instanceOfNameValueHeaderDetailType;
function NameValueHeaderDetailTypeFromJSON(json) {
    return NameValueHeaderDetailTypeFromJSONTyped(json, false);
}
exports.NameValueHeaderDetailTypeFromJSON = NameValueHeaderDetailTypeFromJSON;
function NameValueHeaderDetailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'nameValueHeader': !(0, runtime_1.exists)(json, 'nameValueHeader') ? undefined : (0, NameValueBaseSearchType_1.NameValueBaseSearchTypeFromJSON)(json['nameValueHeader']),
        'nameValueDetails': !(0, runtime_1.exists)(json, 'nameValueDetails') ? undefined : (0, NameValueDetailType_1.NameValueDetailTypeFromJSON)(json['nameValueDetails']),
    };
}
exports.NameValueHeaderDetailTypeFromJSONTyped = NameValueHeaderDetailTypeFromJSONTyped;
function NameValueHeaderDetailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'nameValueHeader': (0, NameValueBaseSearchType_1.NameValueBaseSearchTypeToJSON)(value.nameValueHeader),
        'nameValueDetails': (0, NameValueDetailType_1.NameValueDetailTypeToJSON)(value.nameValueDetails),
    };
}
exports.NameValueHeaderDetailTypeToJSON = NameValueHeaderDetailTypeToJSON;
