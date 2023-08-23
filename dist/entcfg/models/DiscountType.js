"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.DiscountTypeToJSON = exports.DiscountTypeFromJSONTyped = exports.DiscountTypeFromJSON = exports.instanceOfDiscountType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the DiscountType interface.
 */
function instanceOfDiscountType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDiscountType = instanceOfDiscountType;
function DiscountTypeFromJSON(json) {
    return DiscountTypeFromJSONTyped(json, false);
}
exports.DiscountTypeFromJSON = DiscountTypeFromJSON;
function DiscountTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'discountReason': !(0, runtime_1.exists)(json, 'discountReason') ? undefined : json['discountReason'],
        'percent': !(0, runtime_1.exists)(json, 'percent') ? undefined : json['percent'],
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : json['amount'],
        'currencyCode': !(0, runtime_1.exists)(json, 'currencyCode') ? undefined : json['currencyCode'],
        'discountCode': !(0, runtime_1.exists)(json, 'discountCode') ? undefined : json['discountCode'],
    };
}
exports.DiscountTypeFromJSONTyped = DiscountTypeFromJSONTyped;
function DiscountTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'discountReason': value.discountReason,
        'percent': value.percent,
        'amount': value.amount,
        'currencyCode': value.currencyCode,
        'discountCode': value.discountCode,
    };
}
exports.DiscountTypeToJSON = DiscountTypeToJSON;
