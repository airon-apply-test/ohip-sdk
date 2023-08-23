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
exports.BlockWashScheduleTypeToJSON = exports.BlockWashScheduleTypeFromJSONTyped = exports.BlockWashScheduleTypeFromJSON = exports.instanceOfBlockWashScheduleType = void 0;
const runtime_1 = require("../runtime");
const BlockGridInvType_1 = require("./BlockGridInvType");
/**
 * Check if a given object implements the BlockWashScheduleType interface.
 */
function instanceOfBlockWashScheduleType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBlockWashScheduleType = instanceOfBlockWashScheduleType;
function BlockWashScheduleTypeFromJSON(json) {
    return BlockWashScheduleTypeFromJSONTyped(json, false);
}
exports.BlockWashScheduleTypeFromJSON = BlockWashScheduleTypeFromJSON;
function BlockWashScheduleTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'washDate': !(0, runtime_1.exists)(json, 'washDate') ? undefined : (new Date(json['washDate'])),
        'roomTypes': !(0, runtime_1.exists)(json, 'roomTypes') ? undefined : json['roomTypes'],
        'washByRooms': !(0, runtime_1.exists)(json, 'washByRooms') ? undefined : (0, BlockGridInvType_1.BlockGridInvTypeFromJSON)(json['washByRooms']),
        'washByPercent': !(0, runtime_1.exists)(json, 'washByPercent') ? undefined : json['washByPercent'],
    };
}
exports.BlockWashScheduleTypeFromJSONTyped = BlockWashScheduleTypeFromJSONTyped;
function BlockWashScheduleTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'washDate': value.washDate === undefined ? undefined : (value.washDate.toISOString().substr(0, 10)),
        'roomTypes': value.roomTypes,
        'washByRooms': (0, BlockGridInvType_1.BlockGridInvTypeToJSON)(value.washByRooms),
        'washByPercent': value.washByPercent,
    };
}
exports.BlockWashScheduleTypeToJSON = BlockWashScheduleTypeToJSON;
