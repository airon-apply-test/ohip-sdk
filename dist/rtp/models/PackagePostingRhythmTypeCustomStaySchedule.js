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
exports.PackagePostingRhythmTypeCustomStayScheduleToJSON = exports.PackagePostingRhythmTypeCustomStayScheduleFromJSONTyped = exports.PackagePostingRhythmTypeCustomStayScheduleFromJSON = exports.instanceOfPackagePostingRhythmTypeCustomStaySchedule = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the PackagePostingRhythmTypeCustomStaySchedule interface.
 */
function instanceOfPackagePostingRhythmTypeCustomStaySchedule(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPackagePostingRhythmTypeCustomStaySchedule = instanceOfPackagePostingRhythmTypeCustomStaySchedule;
function PackagePostingRhythmTypeCustomStayScheduleFromJSON(json) {
    return PackagePostingRhythmTypeCustomStayScheduleFromJSONTyped(json, false);
}
exports.PackagePostingRhythmTypeCustomStayScheduleFromJSON = PackagePostingRhythmTypeCustomStayScheduleFromJSON;
function PackagePostingRhythmTypeCustomStayScheduleFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'night1': !(0, runtime_1.exists)(json, 'night1') ? undefined : json['night1'],
        'night2': !(0, runtime_1.exists)(json, 'night2') ? undefined : json['night2'],
        'night3': !(0, runtime_1.exists)(json, 'night3') ? undefined : json['night3'],
        'night4': !(0, runtime_1.exists)(json, 'night4') ? undefined : json['night4'],
        'night5': !(0, runtime_1.exists)(json, 'night5') ? undefined : json['night5'],
        'night6': !(0, runtime_1.exists)(json, 'night6') ? undefined : json['night6'],
        'night7': !(0, runtime_1.exists)(json, 'night7') ? undefined : json['night7'],
        'night8': !(0, runtime_1.exists)(json, 'night8') ? undefined : json['night8'],
        'night9': !(0, runtime_1.exists)(json, 'night9') ? undefined : json['night9'],
        'night10': !(0, runtime_1.exists)(json, 'night10') ? undefined : json['night10'],
        'night11': !(0, runtime_1.exists)(json, 'night11') ? undefined : json['night11'],
        'night12': !(0, runtime_1.exists)(json, 'night12') ? undefined : json['night12'],
        'night13': !(0, runtime_1.exists)(json, 'night13') ? undefined : json['night13'],
        'night14': !(0, runtime_1.exists)(json, 'night14') ? undefined : json['night14'],
    };
}
exports.PackagePostingRhythmTypeCustomStayScheduleFromJSONTyped = PackagePostingRhythmTypeCustomStayScheduleFromJSONTyped;
function PackagePostingRhythmTypeCustomStayScheduleToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'night1': value.night1,
        'night2': value.night2,
        'night3': value.night3,
        'night4': value.night4,
        'night5': value.night5,
        'night6': value.night6,
        'night7': value.night7,
        'night8': value.night8,
        'night9': value.night9,
        'night10': value.night10,
        'night11': value.night11,
        'night12': value.night12,
        'night13': value.night13,
        'night14': value.night14,
    };
}
exports.PackagePostingRhythmTypeCustomStayScheduleToJSON = PackagePostingRhythmTypeCustomStayScheduleToJSON;
