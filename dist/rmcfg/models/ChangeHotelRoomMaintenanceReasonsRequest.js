"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ChangeHotelRoomMaintenanceReasonsRequestToJSON = exports.ChangeHotelRoomMaintenanceReasonsRequestFromJSONTyped = exports.ChangeHotelRoomMaintenanceReasonsRequestFromJSON = exports.instanceOfChangeHotelRoomMaintenanceReasonsRequest = void 0;
const runtime_1 = require("../runtime");
const HotelRoomMaintenanceReasonType_1 = require("./HotelRoomMaintenanceReasonType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ChangeHotelRoomMaintenanceReasonsRequest interface.
 */
function instanceOfChangeHotelRoomMaintenanceReasonsRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChangeHotelRoomMaintenanceReasonsRequest = instanceOfChangeHotelRoomMaintenanceReasonsRequest;
function ChangeHotelRoomMaintenanceReasonsRequestFromJSON(json) {
    return ChangeHotelRoomMaintenanceReasonsRequestFromJSONTyped(json, false);
}
exports.ChangeHotelRoomMaintenanceReasonsRequestFromJSON = ChangeHotelRoomMaintenanceReasonsRequestFromJSON;
function ChangeHotelRoomMaintenanceReasonsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelRoomMaintenanceReasons': !(0, runtime_1.exists)(json, 'hotelRoomMaintenanceReasons') ? undefined : (json['hotelRoomMaintenanceReasons'].map(HotelRoomMaintenanceReasonType_1.HotelRoomMaintenanceReasonTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ChangeHotelRoomMaintenanceReasonsRequestFromJSONTyped = ChangeHotelRoomMaintenanceReasonsRequestFromJSONTyped;
function ChangeHotelRoomMaintenanceReasonsRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelRoomMaintenanceReasons': value.hotelRoomMaintenanceReasons === undefined ? undefined : (value.hotelRoomMaintenanceReasons.map(HotelRoomMaintenanceReasonType_1.HotelRoomMaintenanceReasonTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ChangeHotelRoomMaintenanceReasonsRequestToJSON = ChangeHotelRoomMaintenanceReasonsRequestToJSON;
