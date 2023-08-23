"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PostCheckInRequestToJSON = exports.PostCheckInRequestFromJSONTyped = exports.PostCheckInRequestFromJSON = exports.instanceOfPostCheckInRequest = void 0;
const runtime_1 = require("../runtime");
const CheckedInReservationInstructionType_1 = require("./CheckedInReservationInstructionType");
const InstanceLink_1 = require("./InstanceLink");
const ReservationReservation_1 = require("./ReservationReservation");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostCheckInRequest interface.
 */
function instanceOfPostCheckInRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostCheckInRequest = instanceOfPostCheckInRequest;
function PostCheckInRequestFromJSON(json) {
    return PostCheckInRequestFromJSONTyped(json, false);
}
exports.PostCheckInRequestFromJSON = PostCheckInRequestFromJSON;
function PostCheckInRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservation': !(0, runtime_1.exists)(json, 'reservation') ? undefined : (0, ReservationReservation_1.ReservationReservationFromJSON)(json['reservation']),
        'fetchReservationInstruction': !(0, runtime_1.exists)(json, 'fetchReservationInstruction') ? undefined : (json['fetchReservationInstruction'].map(CheckedInReservationInstructionType_1.CheckedInReservationInstructionTypeFromJSON)),
        'includeNotifications': !(0, runtime_1.exists)(json, 'includeNotifications') ? undefined : json['includeNotifications'],
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostCheckInRequestFromJSONTyped = PostCheckInRequestFromJSONTyped;
function PostCheckInRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservation': (0, ReservationReservation_1.ReservationReservationToJSON)(value.reservation),
        'fetchReservationInstruction': value.fetchReservationInstruction === undefined ? undefined : (value.fetchReservationInstruction.map(CheckedInReservationInstructionType_1.CheckedInReservationInstructionTypeToJSON)),
        'includeNotifications': value.includeNotifications,
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostCheckInRequestToJSON = PostCheckInRequestToJSON;
