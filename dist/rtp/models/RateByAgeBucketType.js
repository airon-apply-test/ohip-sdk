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
exports.RateByAgeBucketTypeToJSON = exports.RateByAgeBucketTypeFromJSONTyped = exports.RateByAgeBucketTypeFromJSON = exports.instanceOfRateByAgeBucketType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the RateByAgeBucketType interface.
 */
function instanceOfRateByAgeBucketType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRateByAgeBucketType = instanceOfRateByAgeBucketType;
function RateByAgeBucketTypeFromJSON(json) {
    return RateByAgeBucketTypeFromJSONTyped(json, false);
}
exports.RateByAgeBucketTypeFromJSON = RateByAgeBucketTypeFromJSON;
function RateByAgeBucketTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'rateAmount': !(0, runtime_1.exists)(json, 'rateAmount') ? undefined : json['rateAmount'],
        'minimumAge': !(0, runtime_1.exists)(json, 'minimumAge') ? undefined : json['minimumAge'],
        'maximumAge': !(0, runtime_1.exists)(json, 'maximumAge') ? undefined : json['maximumAge'],
    };
}
exports.RateByAgeBucketTypeFromJSONTyped = RateByAgeBucketTypeFromJSONTyped;
function RateByAgeBucketTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'rateAmount': value.rateAmount,
        'minimumAge': value.minimumAge,
        'maximumAge': value.maximumAge,
    };
}
exports.RateByAgeBucketTypeToJSON = RateByAgeBucketTypeToJSON;
