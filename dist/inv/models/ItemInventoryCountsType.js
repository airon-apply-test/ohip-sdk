"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ItemInventoryCountsTypeToJSON = exports.ItemInventoryCountsTypeFromJSONTyped = exports.ItemInventoryCountsTypeFromJSON = exports.instanceOfItemInventoryCountsType = void 0;
const runtime_1 = require("../runtime");
const DailyItemInventoryCountsType_1 = require("./DailyItemInventoryCountsType");
const TimeSpanType_1 = require("./TimeSpanType");
const TimeWindowType_1 = require("./TimeWindowType");
/**
 * Check if a given object implements the ItemInventoryCountsType interface.
 */
function instanceOfItemInventoryCountsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfItemInventoryCountsType = instanceOfItemInventoryCountsType;
function ItemInventoryCountsTypeFromJSON(json) {
    return ItemInventoryCountsTypeFromJSONTyped(json, false);
}
exports.ItemInventoryCountsTypeFromJSON = ItemInventoryCountsTypeFromJSON;
function ItemInventoryCountsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'availabilityPeriod': !(0, runtime_1.exists)(json, 'availabilityPeriod') ? undefined : (0, TimeWindowType_1.TimeWindowTypeFromJSON)(json['availabilityPeriod']),
        'timeSpan': !(0, runtime_1.exists)(json, 'timeSpan') ? undefined : (0, TimeSpanType_1.TimeSpanTypeFromJSON)(json['timeSpan']),
        'quantity': !(0, runtime_1.exists)(json, 'quantity') ? undefined : json['quantity'],
        'itemHoldId': !(0, runtime_1.exists)(json, 'itemHoldId') ? undefined : json['itemHoldId'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'name': !(0, runtime_1.exists)(json, 'name') ? undefined : json['name'],
        'itemPool': !(0, runtime_1.exists)(json, 'itemPool') ? undefined : json['itemPool'],
        'sellSeparate': !(0, runtime_1.exists)(json, 'sellSeparate') ? undefined : json['sellSeparate'],
        'sellInReservation': !(0, runtime_1.exists)(json, 'sellInReservation') ? undefined : json['sellInReservation'],
        'sellInEvent': !(0, runtime_1.exists)(json, 'sellInEvent') ? undefined : json['sellInEvent'],
        'requiredForBooking': !(0, runtime_1.exists)(json, 'requiredForBooking') ? undefined : json['requiredForBooking'],
        'fixedCharge': !(0, runtime_1.exists)(json, 'fixedCharge') ? undefined : json['fixedCharge'],
        'outsideStay': !(0, runtime_1.exists)(json, 'outsideStay') ? undefined : json['outsideStay'],
        'defaultDuration': !(0, runtime_1.exists)(json, 'defaultDuration') ? undefined : json['defaultDuration'],
        'inventories': !(0, runtime_1.exists)(json, 'inventories') ? undefined : (json['inventories'].map(DailyItemInventoryCountsType_1.DailyItemInventoryCountsTypeFromJSON)),
    };
}
exports.ItemInventoryCountsTypeFromJSONTyped = ItemInventoryCountsTypeFromJSONTyped;
function ItemInventoryCountsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'availabilityPeriod': (0, TimeWindowType_1.TimeWindowTypeToJSON)(value.availabilityPeriod),
        'timeSpan': (0, TimeSpanType_1.TimeSpanTypeToJSON)(value.timeSpan),
        'quantity': value.quantity,
        'itemHoldId': value.itemHoldId,
        'code': value.code,
        'name': value.name,
        'itemPool': value.itemPool,
        'sellSeparate': value.sellSeparate,
        'sellInReservation': value.sellInReservation,
        'sellInEvent': value.sellInEvent,
        'requiredForBooking': value.requiredForBooking,
        'fixedCharge': value.fixedCharge,
        'outsideStay': value.outsideStay,
        'defaultDuration': value.defaultDuration,
        'inventories': value.inventories === undefined ? undefined : (value.inventories.map(DailyItemInventoryCountsType_1.DailyItemInventoryCountsTypeToJSON)),
    };
}
exports.ItemInventoryCountsTypeToJSON = ItemInventoryCountsTypeToJSON;
