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
exports.ResGuestAdditionalInfoTypeToJSON = exports.ResGuestAdditionalInfoTypeFromJSONTyped = exports.ResGuestAdditionalInfoTypeFromJSON = exports.instanceOfResGuestAdditionalInfoType = void 0;
const runtime_1 = require("../runtime");
const GuestLastStayInfoType_1 = require("./GuestLastStayInfoType");
/**
 * Check if a given object implements the ResGuestAdditionalInfoType interface.
 */
function instanceOfResGuestAdditionalInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResGuestAdditionalInfoType = instanceOfResGuestAdditionalInfoType;
function ResGuestAdditionalInfoTypeFromJSON(json) {
    return ResGuestAdditionalInfoTypeFromJSONTyped(json, false);
}
exports.ResGuestAdditionalInfoTypeFromJSON = ResGuestAdditionalInfoTypeFromJSON;
function ResGuestAdditionalInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'portOfEntry': !(0, runtime_1.exists)(json, 'portOfEntry') ? undefined : json['portOfEntry'],
        'dateOfEntry': !(0, runtime_1.exists)(json, 'dateOfEntry') ? undefined : (new Date(json['dateOfEntry'])),
        'nextDestination': !(0, runtime_1.exists)(json, 'nextDestination') ? undefined : json['nextDestination'],
        'preferredRoomType': !(0, runtime_1.exists)(json, 'preferredRoomType') ? undefined : json['preferredRoomType'],
        'lastStay': !(0, runtime_1.exists)(json, 'lastStay') ? undefined : (0, GuestLastStayInfoType_1.GuestLastStayInfoTypeFromJSON)(json['lastStay']),
        'purposeOfStay': !(0, runtime_1.exists)(json, 'purposeOfStay') ? undefined : json['purposeOfStay'],
        'guestClassification': !(0, runtime_1.exists)(json, 'guestClassification') ? undefined : json['guestClassification'],
        'guestStatus': !(0, runtime_1.exists)(json, 'guestStatus') ? undefined : json['guestStatus'],
    };
}
exports.ResGuestAdditionalInfoTypeFromJSONTyped = ResGuestAdditionalInfoTypeFromJSONTyped;
function ResGuestAdditionalInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'portOfEntry': value.portOfEntry,
        'dateOfEntry': value.dateOfEntry === undefined ? undefined : (value.dateOfEntry.toISOString().substr(0, 10)),
        'nextDestination': value.nextDestination,
        'preferredRoomType': value.preferredRoomType,
        'lastStay': (0, GuestLastStayInfoType_1.GuestLastStayInfoTypeToJSON)(value.lastStay),
        'purposeOfStay': value.purposeOfStay,
        'guestClassification': value.guestClassification,
        'guestStatus': value.guestStatus,
    };
}
exports.ResGuestAdditionalInfoTypeToJSON = ResGuestAdditionalInfoTypeToJSON;
