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
exports.RelativePositionTypeToJSON = exports.RelativePositionTypeFromJSONTyped = exports.RelativePositionTypeFromJSON = exports.instanceOfRelativePositionType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the RelativePositionType interface.
 */
function instanceOfRelativePositionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRelativePositionType = instanceOfRelativePositionType;
function RelativePositionTypeFromJSON(json) {
    return RelativePositionTypeFromJSONTyped(json, false);
}
exports.RelativePositionTypeFromJSON = RelativePositionTypeFromJSON;
function RelativePositionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'distance': !(0, runtime_1.exists)(json, 'distance') ? undefined : json['distance'],
        'distanceType': !(0, runtime_1.exists)(json, 'distanceType') ? undefined : json['distanceType'],
        'drivingTime': !(0, runtime_1.exists)(json, 'drivingTime') ? undefined : json['drivingTime'],
    };
}
exports.RelativePositionTypeFromJSONTyped = RelativePositionTypeFromJSONTyped;
function RelativePositionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'distance': value.distance,
        'distanceType': value.distanceType,
        'drivingTime': value.drivingTime,
    };
}
exports.RelativePositionTypeToJSON = RelativePositionTypeToJSON;
