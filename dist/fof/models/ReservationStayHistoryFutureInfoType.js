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
exports.ReservationStayHistoryFutureInfoTypeToJSON = exports.ReservationStayHistoryFutureInfoTypeFromJSONTyped = exports.ReservationStayHistoryFutureInfoTypeFromJSON = exports.instanceOfReservationStayHistoryFutureInfoType = void 0;
const runtime_1 = require("../runtime");
const StayFutureListType_1 = require("./StayFutureListType");
const StayHistoryListType_1 = require("./StayHistoryListType");
/**
 * Check if a given object implements the ReservationStayHistoryFutureInfoType interface.
 */
function instanceOfReservationStayHistoryFutureInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationStayHistoryFutureInfoType = instanceOfReservationStayHistoryFutureInfoType;
function ReservationStayHistoryFutureInfoTypeFromJSON(json) {
    return ReservationStayHistoryFutureInfoTypeFromJSONTyped(json, false);
}
exports.ReservationStayHistoryFutureInfoTypeFromJSON = ReservationStayHistoryFutureInfoTypeFromJSON;
function ReservationStayHistoryFutureInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'historyList': !(0, runtime_1.exists)(json, 'historyList') ? undefined : (0, StayHistoryListType_1.StayHistoryListTypeFromJSON)(json['historyList']),
        'futureList': !(0, runtime_1.exists)(json, 'futureList') ? undefined : (0, StayFutureListType_1.StayFutureListTypeFromJSON)(json['futureList']),
    };
}
exports.ReservationStayHistoryFutureInfoTypeFromJSONTyped = ReservationStayHistoryFutureInfoTypeFromJSONTyped;
function ReservationStayHistoryFutureInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'historyList': (0, StayHistoryListType_1.StayHistoryListTypeToJSON)(value.historyList),
        'futureList': (0, StayFutureListType_1.StayFutureListTypeToJSON)(value.futureList),
    };
}
exports.ReservationStayHistoryFutureInfoTypeToJSON = ReservationStayHistoryFutureInfoTypeToJSON;
