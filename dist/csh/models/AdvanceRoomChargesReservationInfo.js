"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdvanceRoomChargesReservationInfoToJSON = exports.AdvanceRoomChargesReservationInfoFromJSONTyped = exports.AdvanceRoomChargesReservationInfoFromJSON = exports.instanceOfAdvanceRoomChargesReservationInfo = void 0;
const runtime_1 = require("../runtime");
const AdvanceRoomChargesInfoType_1 = require("./AdvanceRoomChargesInfoType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the AdvanceRoomChargesReservationInfo interface.
 */
function instanceOfAdvanceRoomChargesReservationInfo(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAdvanceRoomChargesReservationInfo = instanceOfAdvanceRoomChargesReservationInfo;
function AdvanceRoomChargesReservationInfoFromJSON(json) {
    return AdvanceRoomChargesReservationInfoFromJSONTyped(json, false);
}
exports.AdvanceRoomChargesReservationInfoFromJSON = AdvanceRoomChargesReservationInfoFromJSON;
function AdvanceRoomChargesReservationInfoFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationInfo': !(0, runtime_1.exists)(json, 'reservationInfo') ? undefined : (0, AdvanceRoomChargesInfoType_1.AdvanceRoomChargesInfoTypeFromJSON)(json['reservationInfo']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.AdvanceRoomChargesReservationInfoFromJSONTyped = AdvanceRoomChargesReservationInfoFromJSONTyped;
function AdvanceRoomChargesReservationInfoToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationInfo': (0, AdvanceRoomChargesInfoType_1.AdvanceRoomChargesInfoTypeToJSON)(value.reservationInfo),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.AdvanceRoomChargesReservationInfoToJSON = AdvanceRoomChargesReservationInfoToJSON;
