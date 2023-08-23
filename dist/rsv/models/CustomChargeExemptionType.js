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
exports.CustomChargeExemptionTypeToJSON = exports.CustomChargeExemptionTypeFromJSONTyped = exports.CustomChargeExemptionTypeFromJSON = exports.instanceOfCustomChargeExemptionType = void 0;
const runtime_1 = require("../runtime");
const CodeDescriptionType_1 = require("./CodeDescriptionType");
const CustomChargeExemptionDateType_1 = require("./CustomChargeExemptionDateType");
const CustomChargeQuantityType_1 = require("./CustomChargeQuantityType");
const ExcludedDateType_1 = require("./ExcludedDateType");
/**
 * Check if a given object implements the CustomChargeExemptionType interface.
 */
function instanceOfCustomChargeExemptionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCustomChargeExemptionType = instanceOfCustomChargeExemptionType;
function CustomChargeExemptionTypeFromJSON(json) {
    return CustomChargeExemptionTypeFromJSONTyped(json, false);
}
exports.CustomChargeExemptionTypeFromJSON = CustomChargeExemptionTypeFromJSON;
function CustomChargeExemptionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'customChargesExemption': !(0, runtime_1.exists)(json, 'customChargesExemption') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['customChargesExemption']),
        'customChargeQuantity': !(0, runtime_1.exists)(json, 'customChargeQuantity') ? undefined : (0, CustomChargeQuantityType_1.CustomChargeQuantityTypeFromJSON)(json['customChargeQuantity']),
        'customChargeDates': !(0, runtime_1.exists)(json, 'customChargeDates') ? undefined : (json['customChargeDates'].map(CustomChargeExemptionDateType_1.CustomChargeExemptionDateTypeFromJSON)),
        'excludedDates': !(0, runtime_1.exists)(json, 'excludedDates') ? undefined : (json['excludedDates'].map(ExcludedDateType_1.ExcludedDateTypeFromJSON)),
        'percentage': !(0, runtime_1.exists)(json, 'percentage') ? undefined : json['percentage'],
        'propertyExemption': !(0, runtime_1.exists)(json, 'propertyExemption') ? undefined : json['propertyExemption'],
    };
}
exports.CustomChargeExemptionTypeFromJSONTyped = CustomChargeExemptionTypeFromJSONTyped;
function CustomChargeExemptionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'customChargesExemption': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.customChargesExemption),
        'customChargeQuantity': (0, CustomChargeQuantityType_1.CustomChargeQuantityTypeToJSON)(value.customChargeQuantity),
        'customChargeDates': value.customChargeDates === undefined ? undefined : (value.customChargeDates.map(CustomChargeExemptionDateType_1.CustomChargeExemptionDateTypeToJSON)),
        'excludedDates': value.excludedDates === undefined ? undefined : (value.excludedDates.map(ExcludedDateType_1.ExcludedDateTypeToJSON)),
        'percentage': value.percentage,
        'propertyExemption': value.propertyExemption,
    };
}
exports.CustomChargeExemptionTypeToJSON = CustomChargeExemptionTypeToJSON;
