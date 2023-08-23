"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ItemInfoTypeToJSON = exports.ItemInfoTypeFromJSONTyped = exports.ItemInfoTypeFromJSON = exports.instanceOfItemInfoType = void 0;
const runtime_1 = require("../runtime");
const TimeSpanType_1 = require("./TimeSpanType");
const TimeWindowType_1 = require("./TimeWindowType");
/**
 * Check if a given object implements the ItemInfoType interface.
 */
function instanceOfItemInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfItemInfoType = instanceOfItemInfoType;
function ItemInfoTypeFromJSON(json) {
    return ItemInfoTypeFromJSONTyped(json, false);
}
exports.ItemInfoTypeFromJSON = ItemInfoTypeFromJSON;
function ItemInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
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
    };
}
exports.ItemInfoTypeFromJSONTyped = ItemInfoTypeFromJSONTyped;
function ItemInfoTypeToJSON(value) {
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
    };
}
exports.ItemInfoTypeToJSON = ItemInfoTypeToJSON;
