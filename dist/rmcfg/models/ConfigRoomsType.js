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
exports.ConfigRoomsTypeToJSON = exports.ConfigRoomsTypeFromJSONTyped = exports.ConfigRoomsTypeFromJSON = exports.instanceOfConfigRoomsType = void 0;
const runtime_1 = require("../runtime");
const ConfigRoomType_1 = require("./ConfigRoomType");
/**
 * Check if a given object implements the ConfigRoomsType interface.
 */
function instanceOfConfigRoomsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfConfigRoomsType = instanceOfConfigRoomsType;
function ConfigRoomsTypeFromJSON(json) {
    return ConfigRoomsTypeFromJSONTyped(json, false);
}
exports.ConfigRoomsTypeFromJSON = ConfigRoomsTypeFromJSON;
function ConfigRoomsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'room': !(0, runtime_1.exists)(json, 'room') ? undefined : (json['room'].map(ConfigRoomType_1.ConfigRoomTypeFromJSON)),
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
    };
}
exports.ConfigRoomsTypeFromJSONTyped = ConfigRoomsTypeFromJSONTyped;
function ConfigRoomsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'room': value.room === undefined ? undefined : (value.room.map(ConfigRoomType_1.ConfigRoomTypeToJSON)),
        'hotelId': value.hotelId,
    };
}
exports.ConfigRoomsTypeToJSON = ConfigRoomsTypeToJSON;
