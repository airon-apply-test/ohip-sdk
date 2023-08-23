"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.TotalTypeToJSON = exports.TotalTypeFromJSONTyped = exports.TotalTypeFromJSON = exports.instanceOfTotalType = void 0;
const runtime_1 = require("../runtime");
const TaxesType_1 = require("./TaxesType");
/**
 * Check if a given object implements the TotalType interface.
 */
function instanceOfTotalType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTotalType = instanceOfTotalType;
function TotalTypeFromJSON(json) {
    return TotalTypeFromJSONTyped(json, false);
}
exports.TotalTypeFromJSON = TotalTypeFromJSON;
function TotalTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'taxes': !(0, runtime_1.exists)(json, 'taxes') ? undefined : (0, TaxesType_1.TaxesTypeFromJSON)(json['taxes']),
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'amountBeforeTax': !(0, runtime_1.exists)(json, 'amountBeforeTax') ? undefined : json['amountBeforeTax'],
        'amountAfterTax': !(0, runtime_1.exists)(json, 'amountAfterTax') ? undefined : json['amountAfterTax'],
        'currencyCode': !(0, runtime_1.exists)(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !(0, runtime_1.exists)(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !(0, runtime_1.exists)(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'rateOverride': !(0, runtime_1.exists)(json, 'rateOverride') ? undefined : json['rateOverride'],
    };
}
exports.TotalTypeFromJSONTyped = TotalTypeFromJSONTyped;
function TotalTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'taxes': (0, TaxesType_1.TaxesTypeToJSON)(value.taxes),
        'description': value.description,
        'amountBeforeTax': value.amountBeforeTax,
        'amountAfterTax': value.amountAfterTax,
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
        'code': value.code,
        'rateOverride': value.rateOverride,
    };
}
exports.TotalTypeToJSON = TotalTypeToJSON;
