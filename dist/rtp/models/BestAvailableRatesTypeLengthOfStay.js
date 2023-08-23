"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.BestAvailableRatesTypeLengthOfStayToJSON = exports.BestAvailableRatesTypeLengthOfStayFromJSONTyped = exports.BestAvailableRatesTypeLengthOfStayFromJSON = exports.instanceOfBestAvailableRatesTypeLengthOfStay = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the BestAvailableRatesTypeLengthOfStay interface.
 */
function instanceOfBestAvailableRatesTypeLengthOfStay(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBestAvailableRatesTypeLengthOfStay = instanceOfBestAvailableRatesTypeLengthOfStay;
function BestAvailableRatesTypeLengthOfStayFromJSON(json) {
    return BestAvailableRatesTypeLengthOfStayFromJSONTyped(json, false);
}
exports.BestAvailableRatesTypeLengthOfStayFromJSON = BestAvailableRatesTypeLengthOfStayFromJSON;
function BestAvailableRatesTypeLengthOfStayFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'lOS1': !(0, runtime_1.exists)(json, 'lOS1') ? undefined : json['lOS1'],
        'lOS2': !(0, runtime_1.exists)(json, 'lOS2') ? undefined : json['lOS2'],
        'lOS3': !(0, runtime_1.exists)(json, 'lOS3') ? undefined : json['lOS3'],
        'lOS4': !(0, runtime_1.exists)(json, 'lOS4') ? undefined : json['lOS4'],
        'lOS5': !(0, runtime_1.exists)(json, 'lOS5') ? undefined : json['lOS5'],
        'lOS6': !(0, runtime_1.exists)(json, 'lOS6') ? undefined : json['lOS6'],
        'lOS7': !(0, runtime_1.exists)(json, 'lOS7') ? undefined : json['lOS7'],
        'lOS8': !(0, runtime_1.exists)(json, 'lOS8') ? undefined : json['lOS8'],
    };
}
exports.BestAvailableRatesTypeLengthOfStayFromJSONTyped = BestAvailableRatesTypeLengthOfStayFromJSONTyped;
function BestAvailableRatesTypeLengthOfStayToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'lOS1': value.lOS1,
        'lOS2': value.lOS2,
        'lOS3': value.lOS3,
        'lOS4': value.lOS4,
        'lOS5': value.lOS5,
        'lOS6': value.lOS6,
        'lOS7': value.lOS7,
        'lOS8': value.lOS8,
    };
}
exports.BestAvailableRatesTypeLengthOfStayToJSON = BestAvailableRatesTypeLengthOfStayToJSON;
