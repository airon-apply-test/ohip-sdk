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
exports.DepositMaturityInfoTypeToJSON = exports.DepositMaturityInfoTypeFromJSONTyped = exports.DepositMaturityInfoTypeFromJSON = exports.instanceOfDepositMaturityInfoType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const DepositMaturityType_1 = require("./DepositMaturityType");
/**
 * Check if a given object implements the DepositMaturityInfoType interface.
 */
function instanceOfDepositMaturityInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDepositMaturityInfoType = instanceOfDepositMaturityInfoType;
function DepositMaturityInfoTypeFromJSON(json) {
    return DepositMaturityInfoTypeFromJSONTyped(json, false);
}
exports.DepositMaturityInfoTypeFromJSON = DepositMaturityInfoTypeFromJSON;
function DepositMaturityInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'depositMaturityType': !(0, runtime_1.exists)(json, 'depositMaturityType') ? undefined : (0, DepositMaturityType_1.DepositMaturityTypeFromJSON)(json['depositMaturityType']),
        'totalAmountTransferrable': !(0, runtime_1.exists)(json, 'totalAmountTransferrable') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['totalAmountTransferrable']),
        'totalAmountDue': !(0, runtime_1.exists)(json, 'totalAmountDue') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['totalAmountDue']),
    };
}
exports.DepositMaturityInfoTypeFromJSONTyped = DepositMaturityInfoTypeFromJSONTyped;
function DepositMaturityInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'depositMaturityType': (0, DepositMaturityType_1.DepositMaturityTypeToJSON)(value.depositMaturityType),
        'totalAmountTransferrable': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.totalAmountTransferrable),
        'totalAmountDue': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.totalAmountDue),
    };
}
exports.DepositMaturityInfoTypeToJSON = DepositMaturityInfoTypeToJSON;
