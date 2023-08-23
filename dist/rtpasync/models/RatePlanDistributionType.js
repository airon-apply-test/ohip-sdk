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
exports.RatePlanDistributionTypeToJSON = exports.RatePlanDistributionTypeFromJSONTyped = exports.RatePlanDistributionTypeFromJSON = exports.instanceOfRatePlanDistributionType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the RatePlanDistributionType interface.
 */
function instanceOfRatePlanDistributionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRatePlanDistributionType = instanceOfRatePlanDistributionType;
function RatePlanDistributionTypeFromJSON(json) {
    return RatePlanDistributionTypeFromJSONTyped(json, false);
}
exports.RatePlanDistributionTypeFromJSON = RatePlanDistributionTypeFromJSON;
function RatePlanDistributionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'restrictionUpdate': !(0, runtime_1.exists)(json, 'restrictionUpdate') ? undefined : json['restrictionUpdate'],
        'rateUpdate': !(0, runtime_1.exists)(json, 'rateUpdate') ? undefined : json['rateUpdate'],
        'myFidelioUploadAllowed': !(0, runtime_1.exists)(json, 'myFidelioUploadAllowed') ? undefined : json['myFidelioUploadAllowed'],
        'channelAllowed': !(0, runtime_1.exists)(json, 'channelAllowed') ? undefined : json['channelAllowed'],
    };
}
exports.RatePlanDistributionTypeFromJSONTyped = RatePlanDistributionTypeFromJSONTyped;
function RatePlanDistributionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'restrictionUpdate': value.restrictionUpdate,
        'rateUpdate': value.rateUpdate,
        'myFidelioUploadAllowed': value.myFidelioUploadAllowed,
        'channelAllowed': value.channelAllowed,
    };
}
exports.RatePlanDistributionTypeToJSON = RatePlanDistributionTypeToJSON;
