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
exports.HotelReservationTraceTypeToJSON = exports.HotelReservationTraceTypeFromJSONTyped = exports.HotelReservationTraceTypeFromJSON = exports.instanceOfHotelReservationTraceType = void 0;
const runtime_1 = require("../runtime");
const ReservationId_1 = require("./ReservationId");
const TraceResolveType_1 = require("./TraceResolveType");
const TraceTimeInfoType_1 = require("./TraceTimeInfoType");
/**
 * Check if a given object implements the HotelReservationTraceType interface.
 */
function instanceOfHotelReservationTraceType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelReservationTraceType = instanceOfHotelReservationTraceType;
function HotelReservationTraceTypeFromJSON(json) {
    return HotelReservationTraceTypeFromJSONTyped(json, false);
}
exports.HotelReservationTraceTypeFromJSON = HotelReservationTraceTypeFromJSON;
function HotelReservationTraceTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'timeInfo': !(0, runtime_1.exists)(json, 'timeInfo') ? undefined : (0, TraceTimeInfoType_1.TraceTimeInfoTypeFromJSON)(json['timeInfo']),
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, ReservationId_1.ReservationIdFromJSON)(json['reservationId']),
        'departmentId': !(0, runtime_1.exists)(json, 'departmentId') ? undefined : json['departmentId'],
        'traceText': !(0, runtime_1.exists)(json, 'traceText') ? undefined : json['traceText'],
        'resolveInfo': !(0, runtime_1.exists)(json, 'resolveInfo') ? undefined : (0, TraceResolveType_1.TraceResolveTypeFromJSON)(json['resolveInfo']),
        'url': !(0, runtime_1.exists)(json, 'url') ? undefined : json['url'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'instance': !(0, runtime_1.exists)(json, 'instance') ? undefined : json['instance'],
        'idContext': !(0, runtime_1.exists)(json, 'idContext') ? undefined : json['idContext'],
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'idExtension': !(0, runtime_1.exists)(json, 'idExtension') ? undefined : json['idExtension'],
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}
exports.HotelReservationTraceTypeFromJSONTyped = HotelReservationTraceTypeFromJSONTyped;
function HotelReservationTraceTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'timeInfo': (0, TraceTimeInfoType_1.TraceTimeInfoTypeToJSON)(value.timeInfo),
        'reservationId': (0, ReservationId_1.ReservationIdToJSON)(value.reservationId),
        'departmentId': value.departmentId,
        'traceText': value.traceText,
        'resolveInfo': (0, TraceResolveType_1.TraceResolveTypeToJSON)(value.resolveInfo),
        'url': value.url,
        'type': value.type,
        'instance': value.instance,
        'idContext': value.idContext,
        'id': value.id,
        'idExtension': value.idExtension,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
    };
}
exports.HotelReservationTraceTypeToJSON = HotelReservationTraceTypeToJSON;
