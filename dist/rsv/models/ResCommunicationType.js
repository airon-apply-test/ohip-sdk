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
exports.ResCommunicationTypeToJSON = exports.ResCommunicationTypeFromJSONTyped = exports.ResCommunicationTypeFromJSON = exports.instanceOfResCommunicationType = void 0;
const runtime_1 = require("../runtime");
const ResCommunicationTypeEmails_1 = require("./ResCommunicationTypeEmails");
const ResCommunicationTypeTelephones_1 = require("./ResCommunicationTypeTelephones");
/**
 * Check if a given object implements the ResCommunicationType interface.
 */
function instanceOfResCommunicationType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResCommunicationType = instanceOfResCommunicationType;
function ResCommunicationTypeFromJSON(json) {
    return ResCommunicationTypeFromJSONTyped(json, false);
}
exports.ResCommunicationTypeFromJSON = ResCommunicationTypeFromJSON;
function ResCommunicationTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'telephones': !(0, runtime_1.exists)(json, 'telephones') ? undefined : (0, ResCommunicationTypeTelephones_1.ResCommunicationTypeTelephonesFromJSON)(json['telephones']),
        'emails': !(0, runtime_1.exists)(json, 'emails') ? undefined : (0, ResCommunicationTypeEmails_1.ResCommunicationTypeEmailsFromJSON)(json['emails']),
    };
}
exports.ResCommunicationTypeFromJSONTyped = ResCommunicationTypeFromJSONTyped;
function ResCommunicationTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'telephones': (0, ResCommunicationTypeTelephones_1.ResCommunicationTypeTelephonesToJSON)(value.telephones),
        'emails': (0, ResCommunicationTypeEmails_1.ResCommunicationTypeEmailsToJSON)(value.emails),
    };
}
exports.ResCommunicationTypeToJSON = ResCommunicationTypeToJSON;
