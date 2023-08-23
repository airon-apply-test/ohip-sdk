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
exports.CashierClosureNoTypeToJSON = exports.CashierClosureNoTypeFromJSONTyped = exports.CashierClosureNoTypeFromJSON = exports.instanceOfCashierClosureNoType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the CashierClosureNoType interface.
 */
function instanceOfCashierClosureNoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCashierClosureNoType = instanceOfCashierClosureNoType;
function CashierClosureNoTypeFromJSON(json) {
    return CashierClosureNoTypeFromJSONTyped(json, false);
}
exports.CashierClosureNoTypeFromJSON = CashierClosureNoTypeFromJSON;
function CashierClosureNoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'shiftDropLocationCode': !(0, runtime_1.exists)(json, 'shiftDropLocationCode') ? undefined : json['shiftDropLocationCode'],
        'bagNumberForOthers': !(0, runtime_1.exists)(json, 'bagNumberForOthers') ? undefined : json['bagNumberForOthers'],
        'bagNumberForCash': !(0, runtime_1.exists)(json, 'bagNumberForCash') ? undefined : json['bagNumberForCash'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
        'closureNo': !(0, runtime_1.exists)(json, 'closureNo') ? undefined : json['closureNo'],
    };
}
exports.CashierClosureNoTypeFromJSONTyped = CashierClosureNoTypeFromJSONTyped;
function CashierClosureNoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'shiftDropLocationCode': value.shiftDropLocationCode,
        'bagNumberForOthers': value.bagNumberForOthers,
        'bagNumberForCash': value.bagNumberForCash,
        'cashierId': value.cashierId,
        'closureNo': value.closureNo,
    };
}
exports.CashierClosureNoTypeToJSON = CashierClosureNoTypeToJSON;
