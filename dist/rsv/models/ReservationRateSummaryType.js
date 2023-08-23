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
exports.ReservationRateSummaryTypeToJSON = exports.ReservationRateSummaryTypeFromJSONTyped = exports.ReservationRateSummaryTypeFromJSON = exports.instanceOfReservationRateSummaryType = void 0;
const runtime_1 = require("../runtime");
const ReservationRateSummaryDetailType_1 = require("./ReservationRateSummaryDetailType");
/**
 * Check if a given object implements the ReservationRateSummaryType interface.
 */
function instanceOfReservationRateSummaryType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationRateSummaryType = instanceOfReservationRateSummaryType;
function ReservationRateSummaryTypeFromJSON(json) {
    return ReservationRateSummaryTypeFromJSONTyped(json, false);
}
exports.ReservationRateSummaryTypeFromJSON = ReservationRateSummaryTypeFromJSON;
function ReservationRateSummaryTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'details': !(0, runtime_1.exists)(json, 'details') ? undefined : (json['details'].map(ReservationRateSummaryDetailType_1.ReservationRateSummaryDetailTypeFromJSON)),
        'gross': !(0, runtime_1.exists)(json, 'gross') ? undefined : json['gross'],
        'net': !(0, runtime_1.exists)(json, 'net') ? undefined : json['net'],
        'fixedCharges': !(0, runtime_1.exists)(json, 'fixedCharges') ? undefined : json['fixedCharges'],
        'deposit': !(0, runtime_1.exists)(json, 'deposit') ? undefined : json['deposit'],
        'totalCostOfStay': !(0, runtime_1.exists)(json, 'totalCostOfStay') ? undefined : json['totalCostOfStay'],
        'outStandingCostOfStay': !(0, runtime_1.exists)(json, 'outStandingCostOfStay') ? undefined : json['outStandingCostOfStay'],
        'guestPay': !(0, runtime_1.exists)(json, 'guestPay') ? undefined : json['guestPay'],
        'routing': !(0, runtime_1.exists)(json, 'routing') ? undefined : json['routing'],
        'currencyCode': !(0, runtime_1.exists)(json, 'currencyCode') ? undefined : json['currencyCode'],
        'start': !(0, runtime_1.exists)(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !(0, runtime_1.exists)(json, 'end') ? undefined : (new Date(json['end'])),
        'hasSuppressedRate': !(0, runtime_1.exists)(json, 'hasSuppressedRate') ? undefined : json['hasSuppressedRate'],
    };
}
exports.ReservationRateSummaryTypeFromJSONTyped = ReservationRateSummaryTypeFromJSONTyped;
function ReservationRateSummaryTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'details': value.details === undefined ? undefined : (value.details.map(ReservationRateSummaryDetailType_1.ReservationRateSummaryDetailTypeToJSON)),
        'gross': value.gross,
        'net': value.net,
        'fixedCharges': value.fixedCharges,
        'deposit': value.deposit,
        'totalCostOfStay': value.totalCostOfStay,
        'outStandingCostOfStay': value.outStandingCostOfStay,
        'guestPay': value.guestPay,
        'routing': value.routing,
        'currencyCode': value.currencyCode,
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0, 10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0, 10)),
        'hasSuppressedRate': value.hasSuppressedRate,
    };
}
exports.ReservationRateSummaryTypeToJSON = ReservationRateSummaryTypeToJSON;
