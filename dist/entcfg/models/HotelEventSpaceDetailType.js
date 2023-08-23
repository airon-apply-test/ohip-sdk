"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HotelEventSpaceDetailTypeToJSON = exports.HotelEventSpaceDetailTypeFromJSONTyped = exports.HotelEventSpaceDetailTypeFromJSON = exports.instanceOfHotelEventSpaceDetailType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the HotelEventSpaceDetailType interface.
 */
function instanceOfHotelEventSpaceDetailType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelEventSpaceDetailType = instanceOfHotelEventSpaceDetailType;
function HotelEventSpaceDetailTypeFromJSON(json) {
    return HotelEventSpaceDetailTypeFromJSONTyped(json, false);
}
exports.HotelEventSpaceDetailTypeFromJSON = HotelEventSpaceDetailTypeFromJSON;
function HotelEventSpaceDetailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'maxCapacity': !(0, runtime_1.exists)(json, 'maxCapacity') ? undefined : json['maxCapacity'],
        'maxOccupancies': !(0, runtime_1.exists)(json, 'maxOccupancies') ? undefined : json['maxOccupancies'],
    };
}
exports.HotelEventSpaceDetailTypeFromJSONTyped = HotelEventSpaceDetailTypeFromJSONTyped;
function HotelEventSpaceDetailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'code': value.code,
        'description': value.description,
        'maxCapacity': value.maxCapacity,
        'maxOccupancies': value.maxOccupancies,
    };
}
exports.HotelEventSpaceDetailTypeToJSON = HotelEventSpaceDetailTypeToJSON;
