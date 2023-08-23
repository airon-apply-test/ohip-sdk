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
exports.PostReservationByBlockRequestToJSON = exports.PostReservationByBlockRequestFromJSONTyped = exports.PostReservationByBlockRequestFromJSON = exports.instanceOfPostReservationByBlockRequest = void 0;
const runtime_1 = require("../runtime");
const ChannelResvRQInfoType_1 = require("./ChannelResvRQInfoType");
const HotelReservationsType_1 = require("./HotelReservationsType");
const InstanceLink_1 = require("./InstanceLink");
const ReservationInstructionType_1 = require("./ReservationInstructionType");
const ReservationsInstructionsType_1 = require("./ReservationsInstructionsType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostReservationByBlockRequest interface.
 */
function instanceOfPostReservationByBlockRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostReservationByBlockRequest = instanceOfPostReservationByBlockRequest;
function PostReservationByBlockRequestFromJSON(json) {
    return PostReservationByBlockRequestFromJSONTyped(json, false);
}
exports.PostReservationByBlockRequestFromJSON = PostReservationByBlockRequestFromJSON;
function PostReservationByBlockRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservations': !(0, runtime_1.exists)(json, 'reservations') ? undefined : (0, HotelReservationsType_1.HotelReservationsTypeFromJSON)(json['reservations']),
        'fetchInstructions': !(0, runtime_1.exists)(json, 'fetchInstructions') ? undefined : (json['fetchInstructions'].map(ReservationInstructionType_1.ReservationInstructionTypeFromJSON)),
        'reservationsInstructionsType': !(0, runtime_1.exists)(json, 'reservationsInstructionsType') ? undefined : (0, ReservationsInstructionsType_1.ReservationsInstructionsTypeFromJSON)(json['reservationsInstructionsType']),
        'channelInformation': !(0, runtime_1.exists)(json, 'channelInformation') ? undefined : (0, ChannelResvRQInfoType_1.ChannelResvRQInfoTypeFromJSON)(json['channelInformation']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostReservationByBlockRequestFromJSONTyped = PostReservationByBlockRequestFromJSONTyped;
function PostReservationByBlockRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservations': (0, HotelReservationsType_1.HotelReservationsTypeToJSON)(value.reservations),
        'fetchInstructions': value.fetchInstructions === undefined ? undefined : (value.fetchInstructions.map(ReservationInstructionType_1.ReservationInstructionTypeToJSON)),
        'reservationsInstructionsType': (0, ReservationsInstructionsType_1.ReservationsInstructionsTypeToJSON)(value.reservationsInstructionsType),
        'channelInformation': (0, ChannelResvRQInfoType_1.ChannelResvRQInfoTypeToJSON)(value.channelInformation),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostReservationByBlockRequestToJSON = PostReservationByBlockRequestToJSON;
