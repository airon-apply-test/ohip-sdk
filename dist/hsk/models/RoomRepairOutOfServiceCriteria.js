"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RoomRepairOutOfServiceCriteriaToJSON = exports.RoomRepairOutOfServiceCriteriaFromJSONTyped = exports.RoomRepairOutOfServiceCriteriaFromJSON = exports.instanceOfRoomRepairOutOfServiceCriteria = void 0;
const runtime_1 = require("../runtime");
const RoomRepairOutOfOrderCriteriaRoomRangeList_1 = require("./RoomRepairOutOfOrderCriteriaRoomRangeList");
/**
 * Check if a given object implements the RoomRepairOutOfServiceCriteria interface.
 */
function instanceOfRoomRepairOutOfServiceCriteria(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomRepairOutOfServiceCriteria = instanceOfRoomRepairOutOfServiceCriteria;
function RoomRepairOutOfServiceCriteriaFromJSON(json) {
    return RoomRepairOutOfServiceCriteriaFromJSONTyped(json, false);
}
exports.RoomRepairOutOfServiceCriteriaFromJSON = RoomRepairOutOfServiceCriteriaFromJSON;
function RoomRepairOutOfServiceCriteriaFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomRangeList': !(0, runtime_1.exists)(json, 'roomRangeList') ? undefined : (0, RoomRepairOutOfOrderCriteriaRoomRangeList_1.RoomRepairOutOfOrderCriteriaRoomRangeListFromJSON)(json['roomRangeList']),
    };
}
exports.RoomRepairOutOfServiceCriteriaFromJSONTyped = RoomRepairOutOfServiceCriteriaFromJSONTyped;
function RoomRepairOutOfServiceCriteriaToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomRangeList': (0, RoomRepairOutOfOrderCriteriaRoomRangeList_1.RoomRepairOutOfOrderCriteriaRoomRangeListToJSON)(value.roomRangeList),
    };
}
exports.RoomRepairOutOfServiceCriteriaToJSON = RoomRepairOutOfServiceCriteriaToJSON;
