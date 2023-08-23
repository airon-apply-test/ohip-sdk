"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommissionPaymentMethodsToJSON = exports.CommissionPaymentMethodsFromJSONTyped = exports.CommissionPaymentMethodsFromJSON = exports.CommissionPaymentMethods = void 0;
/**
 * Check
 * @export
 */
exports.CommissionPaymentMethods = {
    Cent: 'Cent',
    Eft: 'Eft',
    Chk: 'Chk'
};
function CommissionPaymentMethodsFromJSON(json) {
    return CommissionPaymentMethodsFromJSONTyped(json, false);
}
exports.CommissionPaymentMethodsFromJSON = CommissionPaymentMethodsFromJSON;
function CommissionPaymentMethodsFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.CommissionPaymentMethodsFromJSONTyped = CommissionPaymentMethodsFromJSONTyped;
function CommissionPaymentMethodsToJSON(value) {
    return value;
}
exports.CommissionPaymentMethodsToJSON = CommissionPaymentMethodsToJSON;
