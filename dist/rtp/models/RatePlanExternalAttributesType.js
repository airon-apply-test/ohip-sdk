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
exports.RatePlanExternalAttributesTypeToJSON = exports.RatePlanExternalAttributesTypeFromJSONTyped = exports.RatePlanExternalAttributesTypeFromJSON = exports.instanceOfRatePlanExternalAttributesType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the RatePlanExternalAttributesType interface.
 */
function instanceOfRatePlanExternalAttributesType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRatePlanExternalAttributesType = instanceOfRatePlanExternalAttributesType;
function RatePlanExternalAttributesTypeFromJSON(json) {
    return RatePlanExternalAttributesTypeFromJSONTyped(json, false);
}
exports.RatePlanExternalAttributesTypeFromJSON = RatePlanExternalAttributesTypeFromJSON;
function RatePlanExternalAttributesTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'marshaRateProgram': !(0, runtime_1.exists)(json, 'marshaRateProgram') ? undefined : json['marshaRateProgram'],
    };
}
exports.RatePlanExternalAttributesTypeFromJSONTyped = RatePlanExternalAttributesTypeFromJSONTyped;
function RatePlanExternalAttributesTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'marshaRateProgram': value.marshaRateProgram,
    };
}
exports.RatePlanExternalAttributesTypeToJSON = RatePlanExternalAttributesTypeToJSON;
