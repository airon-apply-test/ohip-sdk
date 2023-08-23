"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReservationMethodsDetailsToJSON = exports.ReservationMethodsDetailsFromJSONTyped = exports.ReservationMethodsDetailsFromJSON = exports.instanceOfReservationMethodsDetails = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const ReservationMethodType_1 = require("./ReservationMethodType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ReservationMethodsDetails interface.
 */
function instanceOfReservationMethodsDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationMethodsDetails = instanceOfReservationMethodsDetails;
function ReservationMethodsDetailsFromJSON(json) {
    return ReservationMethodsDetailsFromJSONTyped(json, false);
}
exports.ReservationMethodsDetailsFromJSON = ReservationMethodsDetailsFromJSON;
function ReservationMethodsDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationMethods': !(0, runtime_1.exists)(json, 'reservationMethods') ? undefined : (json['reservationMethods'].map(ReservationMethodType_1.ReservationMethodTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ReservationMethodsDetailsFromJSONTyped = ReservationMethodsDetailsFromJSONTyped;
function ReservationMethodsDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationMethods': value.reservationMethods === undefined ? undefined : (value.reservationMethods.map(ReservationMethodType_1.ReservationMethodTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ReservationMethodsDetailsToJSON = ReservationMethodsDetailsToJSON;
