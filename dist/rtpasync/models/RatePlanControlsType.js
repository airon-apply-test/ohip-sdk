"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * Opera Cloud Rate Plan Asynchronous Service API
 * APIs catering to the Rate Plan asynchronous related functionality in a hotel.  This includes adding/updating daily rates&apos; pricing schedules and best available rates by day or length of stay. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RatePlanControlsTypeToJSON = exports.RatePlanControlsTypeFromJSONTyped = exports.RatePlanControlsTypeFromJSON = exports.instanceOfRatePlanControlsType = void 0;
const runtime_1 = require("../runtime");
const RatePlanSellControlsType_1 = require("./RatePlanSellControlsType");
const RatePlanYieldControlsType_1 = require("./RatePlanYieldControlsType");
/**
 * Check if a given object implements the RatePlanControlsType interface.
 */
function instanceOfRatePlanControlsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRatePlanControlsType = instanceOfRatePlanControlsType;
function RatePlanControlsTypeFromJSON(json) {
    return RatePlanControlsTypeFromJSONTyped(json, false);
}
exports.RatePlanControlsTypeFromJSON = RatePlanControlsTypeFromJSON;
function RatePlanControlsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'sell': !(0, runtime_1.exists)(json, 'sell') ? undefined : (0, RatePlanSellControlsType_1.RatePlanSellControlsTypeFromJSON)(json['sell']),
        '_yield': !(0, runtime_1.exists)(json, 'yield') ? undefined : (0, RatePlanYieldControlsType_1.RatePlanYieldControlsTypeFromJSON)(json['yield']),
    };
}
exports.RatePlanControlsTypeFromJSONTyped = RatePlanControlsTypeFromJSONTyped;
function RatePlanControlsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'sell': (0, RatePlanSellControlsType_1.RatePlanSellControlsTypeToJSON)(value.sell),
        'yield': (0, RatePlanYieldControlsType_1.RatePlanYieldControlsTypeToJSON)(value._yield),
    };
}
exports.RatePlanControlsTypeToJSON = RatePlanControlsTypeToJSON;
