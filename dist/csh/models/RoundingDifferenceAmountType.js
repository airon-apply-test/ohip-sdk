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
exports.RoundingDifferenceAmountTypeToJSON = exports.RoundingDifferenceAmountTypeFromJSONTyped = exports.RoundingDifferenceAmountTypeFromJSON = exports.instanceOfRoundingDifferenceAmountType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the RoundingDifferenceAmountType interface.
 */
function instanceOfRoundingDifferenceAmountType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoundingDifferenceAmountType = instanceOfRoundingDifferenceAmountType;
function RoundingDifferenceAmountTypeFromJSON(json) {
    return RoundingDifferenceAmountTypeFromJSONTyped(json, false);
}
exports.RoundingDifferenceAmountTypeFromJSON = RoundingDifferenceAmountTypeFromJSON;
function RoundingDifferenceAmountTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'roundingDifference': !(0, runtime_1.exists)(json, 'roundingDifference') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['roundingDifference']),
    };
}
exports.RoundingDifferenceAmountTypeFromJSONTyped = RoundingDifferenceAmountTypeFromJSONTyped;
function RoundingDifferenceAmountTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'roundingDifference': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.roundingDifference),
    };
}
exports.RoundingDifferenceAmountTypeToJSON = RoundingDifferenceAmountTypeToJSON;
