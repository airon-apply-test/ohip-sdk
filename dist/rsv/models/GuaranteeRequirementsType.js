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
exports.GuaranteeRequirementsTypeToJSON = exports.GuaranteeRequirementsTypeFromJSONTyped = exports.GuaranteeRequirementsTypeFromJSON = exports.instanceOfGuaranteeRequirementsType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the GuaranteeRequirementsType interface.
 */
function instanceOfGuaranteeRequirementsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfGuaranteeRequirementsType = instanceOfGuaranteeRequirementsType;
function GuaranteeRequirementsTypeFromJSON(json) {
    return GuaranteeRequirementsTypeFromJSONTyped(json, false);
}
exports.GuaranteeRequirementsTypeFromJSON = GuaranteeRequirementsTypeFromJSON;
function GuaranteeRequirementsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'requirementCode': !(0, runtime_1.exists)(json, 'requirementCode') ? undefined : json['requirementCode'],
        'cancellation': !(0, runtime_1.exists)(json, 'cancellation') ? undefined : json['cancellation'],
        'phoneNumber': !(0, runtime_1.exists)(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'address': !(0, runtime_1.exists)(json, 'address') ? undefined : json['address'],
        'creditCard': !(0, runtime_1.exists)(json, 'creditCard') ? undefined : json['creditCard'],
        'arrivalTime': !(0, runtime_1.exists)(json, 'arrivalTime') ? undefined : json['arrivalTime'],
        'deposit': !(0, runtime_1.exists)(json, 'deposit') ? undefined : json['deposit'],
    };
}
exports.GuaranteeRequirementsTypeFromJSONTyped = GuaranteeRequirementsTypeFromJSONTyped;
function GuaranteeRequirementsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'requirementCode': value.requirementCode,
        'cancellation': value.cancellation,
        'phoneNumber': value.phoneNumber,
        'address': value.address,
        'creditCard': value.creditCard,
        'arrivalTime': value.arrivalTime,
        'deposit': value.deposit,
    };
}
exports.GuaranteeRequirementsTypeToJSON = GuaranteeRequirementsTypeToJSON;
