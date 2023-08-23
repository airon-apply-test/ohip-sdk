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
exports.AllocationGridByRoomTypeTypeToJSON = exports.AllocationGridByRoomTypeTypeFromJSONTyped = exports.AllocationGridByRoomTypeTypeFromJSON = exports.instanceOfAllocationGridByRoomTypeType = void 0;
const runtime_1 = require("../runtime");
const AllocationGridByDateRangesType_1 = require("./AllocationGridByDateRangesType");
const SellLimitGridByDateRangeType_1 = require("./SellLimitGridByDateRangeType");
/**
 * Check if a given object implements the AllocationGridByRoomTypeType interface.
 */
function instanceOfAllocationGridByRoomTypeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAllocationGridByRoomTypeType = instanceOfAllocationGridByRoomTypeType;
function AllocationGridByRoomTypeTypeFromJSON(json) {
    return AllocationGridByRoomTypeTypeFromJSONTyped(json, false);
}
exports.AllocationGridByRoomTypeTypeFromJSON = AllocationGridByRoomTypeTypeFromJSON;
function AllocationGridByRoomTypeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'allocationGridDates': !(0, runtime_1.exists)(json, 'allocationGridDates') ? undefined : (json['allocationGridDates'].map(AllocationGridByDateRangesType_1.AllocationGridByDateRangesTypeFromJSON)),
        'sellLimitGridDates': !(0, runtime_1.exists)(json, 'sellLimitGridDates') ? undefined : (json['sellLimitGridDates'].map(SellLimitGridByDateRangeType_1.SellLimitGridByDateRangeTypeFromJSON)),
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : json['roomType'],
    };
}
exports.AllocationGridByRoomTypeTypeFromJSONTyped = AllocationGridByRoomTypeTypeFromJSONTyped;
function AllocationGridByRoomTypeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'allocationGridDates': value.allocationGridDates === undefined ? undefined : (value.allocationGridDates.map(AllocationGridByDateRangesType_1.AllocationGridByDateRangesTypeToJSON)),
        'sellLimitGridDates': value.sellLimitGridDates === undefined ? undefined : (value.sellLimitGridDates.map(SellLimitGridByDateRangeType_1.SellLimitGridByDateRangeTypeToJSON)),
        'roomType': value.roomType,
    };
}
exports.AllocationGridByRoomTypeTypeToJSON = AllocationGridByRoomTypeTypeToJSON;
