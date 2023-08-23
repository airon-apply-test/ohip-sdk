"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.BlockAvailabilityRoomTypeInfoToJSON = exports.BlockAvailabilityRoomTypeInfoFromJSONTyped = exports.BlockAvailabilityRoomTypeInfoFromJSON = exports.instanceOfBlockAvailabilityRoomTypeInfo = void 0;
const runtime_1 = require("../runtime");
const PackageElementType_1 = require("./PackageElementType");
const RoomTypeInfoType_1 = require("./RoomTypeInfoType");
/**
 * Check if a given object implements the BlockAvailabilityRoomTypeInfo interface.
 */
function instanceOfBlockAvailabilityRoomTypeInfo(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBlockAvailabilityRoomTypeInfo = instanceOfBlockAvailabilityRoomTypeInfo;
function BlockAvailabilityRoomTypeInfoFromJSON(json) {
    return BlockAvailabilityRoomTypeInfoFromJSONTyped(json, false);
}
exports.BlockAvailabilityRoomTypeInfoFromJSON = BlockAvailabilityRoomTypeInfoFromJSON;
function BlockAvailabilityRoomTypeInfoFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'packages': !(0, runtime_1.exists)(json, 'packages') ? undefined : (json['packages'].map(PackageElementType_1.PackageElementTypeFromJSON)),
        'roomTypeInfo': !(0, runtime_1.exists)(json, 'roomTypeInfo') ? undefined : (0, RoomTypeInfoType_1.RoomTypeInfoTypeFromJSON)(json['roomTypeInfo']),
        'available': !(0, runtime_1.exists)(json, 'available') ? undefined : json['available'],
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : json['roomType'],
    };
}
exports.BlockAvailabilityRoomTypeInfoFromJSONTyped = BlockAvailabilityRoomTypeInfoFromJSONTyped;
function BlockAvailabilityRoomTypeInfoToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'packages': value.packages === undefined ? undefined : (value.packages.map(PackageElementType_1.PackageElementTypeToJSON)),
        'roomTypeInfo': (0, RoomTypeInfoType_1.RoomTypeInfoTypeToJSON)(value.roomTypeInfo),
        'available': value.available,
        'roomType': value.roomType,
    };
}
exports.BlockAvailabilityRoomTypeInfoToJSON = BlockAvailabilityRoomTypeInfoToJSON;
