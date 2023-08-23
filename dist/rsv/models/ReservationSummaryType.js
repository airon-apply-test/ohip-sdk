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
exports.ReservationSummaryTypeToJSON = exports.ReservationSummaryTypeFromJSONTyped = exports.ReservationSummaryTypeFromJSON = exports.instanceOfReservationSummaryType = void 0;
const runtime_1 = require("../runtime");
const DateRangeType_1 = require("./DateRangeType");
const ResGuaranteeType_1 = require("./ResGuaranteeType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ReservationSummaryType interface.
 */
function instanceOfReservationSummaryType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationSummaryType = instanceOfReservationSummaryType;
function ReservationSummaryTypeFromJSON(json) {
    return ReservationSummaryTypeFromJSONTyped(json, false);
}
exports.ReservationSummaryTypeFromJSON = ReservationSummaryTypeFromJSON;
function ReservationSummaryTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'guarantee': !(0, runtime_1.exists)(json, 'guarantee') ? undefined : (0, ResGuaranteeType_1.ResGuaranteeTypeFromJSON)(json['guarantee']),
        'guestFirstName': !(0, runtime_1.exists)(json, 'guestFirstName') ? undefined : json['guestFirstName'],
        'guestLastName': !(0, runtime_1.exists)(json, 'guestLastName') ? undefined : json['guestLastName'],
        'stayPeriod': !(0, runtime_1.exists)(json, 'stayPeriod') ? undefined : (0, DateRangeType_1.DateRangeTypeFromJSON)(json['stayPeriod']),
    };
}
exports.ReservationSummaryTypeFromJSONTyped = ReservationSummaryTypeFromJSONTyped;
function ReservationSummaryTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'guarantee': (0, ResGuaranteeType_1.ResGuaranteeTypeToJSON)(value.guarantee),
        'guestFirstName': value.guestFirstName,
        'guestLastName': value.guestLastName,
        'stayPeriod': (0, DateRangeType_1.DateRangeTypeToJSON)(value.stayPeriod),
    };
}
exports.ReservationSummaryTypeToJSON = ReservationSummaryTypeToJSON;
