"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FixedChargeDetailTypeToJSON = exports.FixedChargeDetailTypeFromJSONTyped = exports.FixedChargeDetailTypeFromJSON = exports.instanceOfFixedChargeDetailType = void 0;
const runtime_1 = require("../runtime");
const CodeDescriptionType_1 = require("./CodeDescriptionType");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the FixedChargeDetailType interface.
 */
function instanceOfFixedChargeDetailType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFixedChargeDetailType = instanceOfFixedChargeDetailType;
function FixedChargeDetailTypeFromJSON(json) {
    return FixedChargeDetailTypeFromJSONTyped(json, false);
}
exports.FixedChargeDetailTypeFromJSON = FixedChargeDetailTypeFromJSON;
function FixedChargeDetailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'transaction': !(0, runtime_1.exists)(json, 'transaction') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['transaction']),
        'quantity': !(0, runtime_1.exists)(json, 'quantity') ? undefined : json['quantity'],
        'chargeAmount': !(0, runtime_1.exists)(json, 'chargeAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['chargeAmount']),
        'percent': !(0, runtime_1.exists)(json, 'percent') ? undefined : json['percent'],
        'supplement': !(0, runtime_1.exists)(json, 'supplement') ? undefined : json['supplement'],
        'article': !(0, runtime_1.exists)(json, 'article') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['article']),
        'roomNights': !(0, runtime_1.exists)(json, 'roomNights') ? undefined : json['roomNights'],
    };
}
exports.FixedChargeDetailTypeFromJSONTyped = FixedChargeDetailTypeFromJSONTyped;
function FixedChargeDetailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'transaction': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.transaction),
        'quantity': value.quantity,
        'chargeAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.chargeAmount),
        'percent': value.percent,
        'supplement': value.supplement,
        'article': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.article),
        'roomNights': value.roomNights,
    };
}
exports.FixedChargeDetailTypeToJSON = FixedChargeDetailTypeToJSON;
