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
exports.CompleteRoomRepairCriteriaRoomRangeListToJSON = exports.CompleteRoomRepairCriteriaRoomRangeListFromJSONTyped = exports.CompleteRoomRepairCriteriaRoomRangeListFromJSON = exports.instanceOfCompleteRoomRepairCriteriaRoomRangeList = void 0;
const runtime_1 = require("../runtime");
const CompleteRoomRepairCriteriaRoomRangeListRange_1 = require("./CompleteRoomRepairCriteriaRoomRangeListRange");
const CompleteRoomRepairCriteriaRoomRepairByRoomNumberInnerRoomOutOfOrder_1 = require("./CompleteRoomRepairCriteriaRoomRepairByRoomNumberInnerRoomOutOfOrder");
/**
 * Check if a given object implements the CompleteRoomRepairCriteriaRoomRangeList interface.
 */
function instanceOfCompleteRoomRepairCriteriaRoomRangeList(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCompleteRoomRepairCriteriaRoomRangeList = instanceOfCompleteRoomRepairCriteriaRoomRangeList;
function CompleteRoomRepairCriteriaRoomRangeListFromJSON(json) {
    return CompleteRoomRepairCriteriaRoomRangeListFromJSONTyped(json, false);
}
exports.CompleteRoomRepairCriteriaRoomRangeListFromJSON = CompleteRoomRepairCriteriaRoomRangeListFromJSON;
function CompleteRoomRepairCriteriaRoomRangeListFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomIds': !(0, runtime_1.exists)(json, 'roomIds') ? undefined : json['roomIds'],
        'range': !(0, runtime_1.exists)(json, 'range') ? undefined : (0, CompleteRoomRepairCriteriaRoomRangeListRange_1.CompleteRoomRepairCriteriaRoomRangeListRangeFromJSON)(json['range']),
        'roomOutOfOrder': !(0, runtime_1.exists)(json, 'roomOutOfOrder') ? undefined : (0, CompleteRoomRepairCriteriaRoomRepairByRoomNumberInnerRoomOutOfOrder_1.CompleteRoomRepairCriteriaRoomRepairByRoomNumberInnerRoomOutOfOrderFromJSON)(json['roomOutOfOrder']),
    };
}
exports.CompleteRoomRepairCriteriaRoomRangeListFromJSONTyped = CompleteRoomRepairCriteriaRoomRangeListFromJSONTyped;
function CompleteRoomRepairCriteriaRoomRangeListToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomIds': value.roomIds,
        'range': (0, CompleteRoomRepairCriteriaRoomRangeListRange_1.CompleteRoomRepairCriteriaRoomRangeListRangeToJSON)(value.range),
        'roomOutOfOrder': (0, CompleteRoomRepairCriteriaRoomRepairByRoomNumberInnerRoomOutOfOrder_1.CompleteRoomRepairCriteriaRoomRepairByRoomNumberInnerRoomOutOfOrderToJSON)(value.roomOutOfOrder),
    };
}
exports.CompleteRoomRepairCriteriaRoomRangeListToJSON = CompleteRoomRepairCriteriaRoomRangeListToJSON;
