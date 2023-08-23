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
exports.RoomTypesToBeChangedToJSON = exports.RoomTypesToBeChangedFromJSONTyped = exports.RoomTypesToBeChangedFromJSON = exports.instanceOfRoomTypesToBeChanged = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const RoomTypesToBeChangedRoomType_1 = require("./RoomTypesToBeChangedRoomType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the RoomTypesToBeChanged interface.
 */
function instanceOfRoomTypesToBeChanged(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomTypesToBeChanged = instanceOfRoomTypesToBeChanged;
function RoomTypesToBeChangedFromJSON(json) {
    return RoomTypesToBeChangedFromJSONTyped(json, false);
}
exports.RoomTypesToBeChangedFromJSON = RoomTypesToBeChangedFromJSON;
function RoomTypesToBeChangedFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : (0, RoomTypesToBeChangedRoomType_1.RoomTypesToBeChangedRoomTypeFromJSON)(json['roomType']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.RoomTypesToBeChangedFromJSONTyped = RoomTypesToBeChangedFromJSONTyped;
function RoomTypesToBeChangedToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomType': (0, RoomTypesToBeChangedRoomType_1.RoomTypesToBeChangedRoomTypeToJSON)(value.roomType),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.RoomTypesToBeChangedToJSON = RoomTypesToBeChangedToJSON;
