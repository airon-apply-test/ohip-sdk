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
exports.DepositFolioTransactionDetailsToJSON = exports.DepositFolioTransactionDetailsFromJSONTyped = exports.DepositFolioTransactionDetailsFromJSON = exports.instanceOfDepositFolioTransactionDetails = void 0;
const runtime_1 = require("../runtime");
const DepositDetailPostingType_1 = require("./DepositDetailPostingType");
const InstanceLink_1 = require("./InstanceLink");
const TrxInfoType_1 = require("./TrxInfoType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the DepositFolioTransactionDetails interface.
 */
function instanceOfDepositFolioTransactionDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDepositFolioTransactionDetails = instanceOfDepositFolioTransactionDetails;
function DepositFolioTransactionDetailsFromJSON(json) {
    return DepositFolioTransactionDetailsFromJSONTyped(json, false);
}
exports.DepositFolioTransactionDetailsFromJSON = DepositFolioTransactionDetailsFromJSON;
function DepositFolioTransactionDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'transactions': !(0, runtime_1.exists)(json, 'transactions') ? undefined : (json['transactions'].map(DepositDetailPostingType_1.DepositDetailPostingTypeFromJSON)),
        'trxCodesInfo': !(0, runtime_1.exists)(json, 'trxCodesInfo') ? undefined : (json['trxCodesInfo'].map(TrxInfoType_1.TrxInfoTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.DepositFolioTransactionDetailsFromJSONTyped = DepositFolioTransactionDetailsFromJSONTyped;
function DepositFolioTransactionDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'transactions': value.transactions === undefined ? undefined : (value.transactions.map(DepositDetailPostingType_1.DepositDetailPostingTypeToJSON)),
        'trxCodesInfo': value.trxCodesInfo === undefined ? undefined : (value.trxCodesInfo.map(TrxInfoType_1.TrxInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.DepositFolioTransactionDetailsToJSON = DepositFolioTransactionDetailsToJSON;
