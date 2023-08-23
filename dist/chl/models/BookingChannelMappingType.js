"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.BookingChannelMappingTypeToJSON = exports.BookingChannelMappingTypeFromJSONTyped = exports.BookingChannelMappingTypeFromJSON = exports.instanceOfBookingChannelMappingType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the BookingChannelMappingType interface.
 */
function instanceOfBookingChannelMappingType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBookingChannelMappingType = instanceOfBookingChannelMappingType;
function BookingChannelMappingTypeFromJSON(json) {
    return BookingChannelMappingTypeFromJSONTyped(json, false);
}
exports.BookingChannelMappingTypeFromJSON = BookingChannelMappingTypeFromJSON;
function BookingChannelMappingTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'bookingChannelCode': !(0, runtime_1.exists)(json, 'bookingChannelCode') ? undefined : json['bookingChannelCode'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'name': !(0, runtime_1.exists)(json, 'name') ? undefined : json['name'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
    };
}
exports.BookingChannelMappingTypeFromJSONTyped = BookingChannelMappingTypeFromJSONTyped;
function BookingChannelMappingTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'bookingChannelCode': value.bookingChannelCode,
        'code': value.code,
        'name': value.name,
        'description': value.description,
    };
}
exports.BookingChannelMappingTypeToJSON = BookingChannelMappingTypeToJSON;
