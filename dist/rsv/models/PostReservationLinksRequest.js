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
exports.PostReservationLinksRequestToJSON = exports.PostReservationLinksRequestFromJSONTyped = exports.PostReservationLinksRequestFromJSON = exports.instanceOfPostReservationLinksRequest = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const LinkReservationsCriteriaResponseInstruction_1 = require("./LinkReservationsCriteriaResponseInstruction");
const UniqueIDType_1 = require("./UniqueIDType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostReservationLinksRequest interface.
 */
function instanceOfPostReservationLinksRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostReservationLinksRequest = instanceOfPostReservationLinksRequest;
function PostReservationLinksRequestFromJSON(json) {
    return PostReservationLinksRequestFromJSONTyped(json, false);
}
exports.PostReservationLinksRequestFromJSON = PostReservationLinksRequestFromJSON;
function PostReservationLinksRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'linkToReservationId': !(0, runtime_1.exists)(json, 'linkToReservationId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['linkToReservationId']),
        'responseInstruction': !(0, runtime_1.exists)(json, 'responseInstruction') ? undefined : (0, LinkReservationsCriteriaResponseInstruction_1.LinkReservationsCriteriaResponseInstructionFromJSON)(json['responseInstruction']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostReservationLinksRequestFromJSONTyped = PostReservationLinksRequestFromJSONTyped;
function PostReservationLinksRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'linkToReservationId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.linkToReservationId),
        'responseInstruction': (0, LinkReservationsCriteriaResponseInstruction_1.LinkReservationsCriteriaResponseInstructionToJSON)(value.responseInstruction),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostReservationLinksRequestToJSON = PostReservationLinksRequestToJSON;
