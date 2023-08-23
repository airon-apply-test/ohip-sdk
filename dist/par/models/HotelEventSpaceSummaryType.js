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
exports.HotelEventSpaceSummaryTypeToJSON = exports.HotelEventSpaceSummaryTypeFromJSONTyped = exports.HotelEventSpaceSummaryTypeFromJSON = exports.instanceOfHotelEventSpaceSummaryType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the HotelEventSpaceSummaryType interface.
 */
function instanceOfHotelEventSpaceSummaryType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelEventSpaceSummaryType = instanceOfHotelEventSpaceSummaryType;
function HotelEventSpaceSummaryTypeFromJSON(json) {
    return HotelEventSpaceSummaryTypeFromJSONTyped(json, false);
}
exports.HotelEventSpaceSummaryTypeFromJSON = HotelEventSpaceSummaryTypeFromJSON;
function HotelEventSpaceSummaryTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'no': !(0, runtime_1.exists)(json, 'no') ? undefined : json['no'],
        'spaceType': !(0, runtime_1.exists)(json, 'spaceType') ? undefined : json['spaceType'],
        'maxCapacity': !(0, runtime_1.exists)(json, 'maxCapacity') ? undefined : json['maxCapacity'],
        'maxOccupancies': !(0, runtime_1.exists)(json, 'maxOccupancies') ? undefined : json['maxOccupancies'],
    };
}
exports.HotelEventSpaceSummaryTypeFromJSONTyped = HotelEventSpaceSummaryTypeFromJSONTyped;
function HotelEventSpaceSummaryTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'no': value.no,
        'spaceType': value.spaceType,
        'maxCapacity': value.maxCapacity,
        'maxOccupancies': value.maxOccupancies,
    };
}
exports.HotelEventSpaceSummaryTypeToJSON = HotelEventSpaceSummaryTypeToJSON;
