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
exports.AdvancedDynamicBaseRateTypeToJSON = exports.AdvancedDynamicBaseRateTypeFromJSONTyped = exports.AdvancedDynamicBaseRateTypeFromJSON = exports.instanceOfAdvancedDynamicBaseRateType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the AdvancedDynamicBaseRateType interface.
 */
function instanceOfAdvancedDynamicBaseRateType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAdvancedDynamicBaseRateType = instanceOfAdvancedDynamicBaseRateType;
function AdvancedDynamicBaseRateTypeFromJSON(json) {
    return AdvancedDynamicBaseRateTypeFromJSONTyped(json, false);
}
exports.AdvancedDynamicBaseRateTypeFromJSON = AdvancedDynamicBaseRateTypeFromJSON;
function AdvancedDynamicBaseRateTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'basedOnRatePlan': !(0, runtime_1.exists)(json, 'basedOnRatePlan') ? undefined : json['basedOnRatePlan'],
        'rounding': !(0, runtime_1.exists)(json, 'rounding') ? undefined : json['rounding'],
        'compareWithRateSchedules': !(0, runtime_1.exists)(json, 'compareWithRateSchedules') ? undefined : json['compareWithRateSchedules'],
        'advancedDependentRatePlans': !(0, runtime_1.exists)(json, 'advancedDependentRatePlans') ? undefined : json['advancedDependentRatePlans'],
    };
}
exports.AdvancedDynamicBaseRateTypeFromJSONTyped = AdvancedDynamicBaseRateTypeFromJSONTyped;
function AdvancedDynamicBaseRateTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'basedOnRatePlan': value.basedOnRatePlan,
        'rounding': value.rounding,
        'compareWithRateSchedules': value.compareWithRateSchedules,
        'advancedDependentRatePlans': value.advancedDependentRatePlans,
    };
}
exports.AdvancedDynamicBaseRateTypeToJSON = AdvancedDynamicBaseRateTypeToJSON;
