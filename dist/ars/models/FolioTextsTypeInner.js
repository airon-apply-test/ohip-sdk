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
exports.FolioTextsTypeInnerToJSON = exports.FolioTextsTypeInnerFromJSONTyped = exports.FolioTextsTypeInnerFromJSON = exports.instanceOfFolioTextsTypeInner = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the FolioTextsTypeInner interface.
 */
function instanceOfFolioTextsTypeInner(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFolioTextsTypeInner = instanceOfFolioTextsTypeInner;
function FolioTextsTypeInnerFromJSON(json) {
    return FolioTextsTypeInnerFromJSONTyped(json, false);
}
exports.FolioTextsTypeInnerFromJSON = FolioTextsTypeInnerFromJSON;
function FolioTextsTypeInnerFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'text': !(0, runtime_1.exists)(json, 'text') ? undefined : json['text'],
        'row': !(0, runtime_1.exists)(json, 'row') ? undefined : json['row'],
    };
}
exports.FolioTextsTypeInnerFromJSONTyped = FolioTextsTypeInnerFromJSONTyped;
function FolioTextsTypeInnerToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'text': value.text,
        'row': value.row,
    };
}
exports.FolioTextsTypeInnerToJSON = FolioTextsTypeInnerToJSON;
