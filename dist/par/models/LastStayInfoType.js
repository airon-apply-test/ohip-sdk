"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.LastStayInfoTypeToJSON = exports.LastStayInfoTypeFromJSONTyped = exports.LastStayInfoTypeFromJSON = exports.instanceOfLastStayInfoType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the LastStayInfoType interface.
 */
function instanceOfLastStayInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfLastStayInfoType = instanceOfLastStayInfoType;
function LastStayInfoTypeFromJSON(json) {
    return LastStayInfoTypeFromJSONTyped(json, false);
}
exports.LastStayInfoTypeFromJSON = LastStayInfoTypeFromJSON;
function LastStayInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'lastVisit': !(0, runtime_1.exists)(json, 'lastVisit') ? undefined : (new Date(json['lastVisit'])),
        'lastRoom': !(0, runtime_1.exists)(json, 'lastRoom') ? undefined : json['lastRoom'],
        'lastRate': !(0, runtime_1.exists)(json, 'lastRate') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['lastRate']),
        'totalStay': !(0, runtime_1.exists)(json, 'totalStay') ? undefined : json['totalStay'],
    };
}
exports.LastStayInfoTypeFromJSONTyped = LastStayInfoTypeFromJSONTyped;
function LastStayInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'lastVisit': value.lastVisit === undefined ? undefined : (value.lastVisit.toISOString().substr(0, 10)),
        'lastRoom': value.lastRoom,
        'lastRate': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.lastRate),
        'totalStay': value.totalStay,
    };
}
exports.LastStayInfoTypeToJSON = LastStayInfoTypeToJSON;
