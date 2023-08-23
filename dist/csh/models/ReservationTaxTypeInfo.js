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
exports.ReservationTaxTypeInfoToJSON = exports.ReservationTaxTypeInfoFromJSONTyped = exports.ReservationTaxTypeInfoFromJSON = exports.instanceOfReservationTaxTypeInfo = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ReservationTaxTypeInfo interface.
 */
function instanceOfReservationTaxTypeInfo(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationTaxTypeInfo = instanceOfReservationTaxTypeInfo;
function ReservationTaxTypeInfoFromJSON(json) {
    return ReservationTaxTypeInfoFromJSONTyped(json, false);
}
exports.ReservationTaxTypeInfoFromJSON = ReservationTaxTypeInfoFromJSON;
function ReservationTaxTypeInfoFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'collectingAgentTax': !(0, runtime_1.exists)(json, 'collectingAgentTax') ? undefined : json['collectingAgentTax'],
        'printAutoAdjust': !(0, runtime_1.exists)(json, 'printAutoAdjust') ? undefined : json['printAutoAdjust'],
        'reportExemptDays': !(0, runtime_1.exists)(json, 'reportExemptDays') ? undefined : json['reportExemptDays'],
        'reportTaxPercentage': !(0, runtime_1.exists)(json, 'reportTaxPercentage') ? undefined : json['reportTaxPercentage'],
        'minimumLengthOfStay': !(0, runtime_1.exists)(json, 'minimumLengthOfStay') ? undefined : json['minimumLengthOfStay'],
        'taxExemptNo': !(0, runtime_1.exists)(json, 'taxExemptNo') ? undefined : json['taxExemptNo'],
    };
}
exports.ReservationTaxTypeInfoFromJSONTyped = ReservationTaxTypeInfoFromJSONTyped;
function ReservationTaxTypeInfoToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'code': value.code,
        'description': value.description,
        'collectingAgentTax': value.collectingAgentTax,
        'printAutoAdjust': value.printAutoAdjust,
        'reportExemptDays': value.reportExemptDays,
        'reportTaxPercentage': value.reportTaxPercentage,
        'minimumLengthOfStay': value.minimumLengthOfStay,
        'taxExemptNo': value.taxExemptNo,
    };
}
exports.ReservationTaxTypeInfoToJSON = ReservationTaxTypeInfoToJSON;
