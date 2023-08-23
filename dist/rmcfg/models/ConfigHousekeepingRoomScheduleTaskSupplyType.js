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
exports.ConfigHousekeepingRoomScheduleTaskSupplyTypeToJSON = exports.ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSONTyped = exports.ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSON = exports.instanceOfConfigHousekeepingRoomScheduleTaskSupplyType = void 0;
const runtime_1 = require("../runtime");
const CodeDescriptionType_1 = require("./CodeDescriptionType");
/**
 * Check if a given object implements the ConfigHousekeepingRoomScheduleTaskSupplyType interface.
 */
function instanceOfConfigHousekeepingRoomScheduleTaskSupplyType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfConfigHousekeepingRoomScheduleTaskSupplyType = instanceOfConfigHousekeepingRoomScheduleTaskSupplyType;
function ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSON(json) {
    return ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSONTyped(json, false);
}
exports.ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSON = ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSON;
function ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'quantity': !(0, runtime_1.exists)(json, 'quantity') ? undefined : json['quantity'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'sequence': !(0, runtime_1.exists)(json, 'sequence') ? undefined : json['sequence'],
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['roomType']),
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'housekeepingRoomScheduleCode': !(0, runtime_1.exists)(json, 'housekeepingRoomScheduleCode') ? undefined : json['housekeepingRoomScheduleCode'],
    };
}
exports.ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSONTyped = ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSONTyped;
function ConfigHousekeepingRoomScheduleTaskSupplyTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'quantity': value.quantity,
        'code': value.code,
        'sequence': value.sequence,
        'roomType': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.roomType),
        'hotelId': value.hotelId,
        'housekeepingRoomScheduleCode': value.housekeepingRoomScheduleCode,
    };
}
exports.ConfigHousekeepingRoomScheduleTaskSupplyTypeToJSON = ConfigHousekeepingRoomScheduleTaskSupplyTypeToJSON;
