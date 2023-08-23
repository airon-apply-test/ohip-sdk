"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RateCategoryTypeRateCategoryDetailsToJSON = exports.RateCategoryTypeRateCategoryDetailsFromJSONTyped = exports.RateCategoryTypeRateCategoryDetailsFromJSON = exports.instanceOfRateCategoryTypeRateCategoryDetails = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the RateCategoryTypeRateCategoryDetails interface.
 */
function instanceOfRateCategoryTypeRateCategoryDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRateCategoryTypeRateCategoryDetails = instanceOfRateCategoryTypeRateCategoryDetails;
function RateCategoryTypeRateCategoryDetailsFromJSON(json) {
    return RateCategoryTypeRateCategoryDetailsFromJSONTyped(json, false);
}
exports.RateCategoryTypeRateCategoryDetailsFromJSON = RateCategoryTypeRateCategoryDetailsFromJSON;
function RateCategoryTypeRateCategoryDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'beginDate': !(0, runtime_1.exists)(json, 'beginDate') ? undefined : (new Date(json['beginDate'])),
        'endDate': !(0, runtime_1.exists)(json, 'endDate') ? undefined : (new Date(json['endDate'])),
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'rateClass': !(0, runtime_1.exists)(json, 'rateClass') ? undefined : json['rateClass'],
        'sellSequence': !(0, runtime_1.exists)(json, 'sellSequence') ? undefined : json['sellSequence'],
    };
}
exports.RateCategoryTypeRateCategoryDetailsFromJSONTyped = RateCategoryTypeRateCategoryDetailsFromJSONTyped;
function RateCategoryTypeRateCategoryDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'beginDate': value.beginDate === undefined ? undefined : (value.beginDate.toISOString().substr(0, 10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0, 10)),
        'description': value.description,
        'rateClass': value.rateClass,
        'sellSequence': value.sellSequence,
    };
}
exports.RateCategoryTypeRateCategoryDetailsToJSON = RateCategoryTypeRateCategoryDetailsToJSON;
