"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.StatisticsRevenueTypeTypeToJSON = exports.StatisticsRevenueTypeTypeFromJSONTyped = exports.StatisticsRevenueTypeTypeFromJSON = exports.instanceOfStatisticsRevenueTypeType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the StatisticsRevenueTypeType interface.
 */
function instanceOfStatisticsRevenueTypeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfStatisticsRevenueTypeType = instanceOfStatisticsRevenueTypeType;
function StatisticsRevenueTypeTypeFromJSON(json) {
    return StatisticsRevenueTypeTypeFromJSONTyped(json, false);
}
exports.StatisticsRevenueTypeTypeFromJSON = StatisticsRevenueTypeTypeFromJSON;
function StatisticsRevenueTypeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'revenueAmount': !(0, runtime_1.exists)(json, 'revenueAmount') ? undefined : json['revenueAmount'],
        'revenueLabel': !(0, runtime_1.exists)(json, 'revenueLabel') ? undefined : json['revenueLabel'],
    };
}
exports.StatisticsRevenueTypeTypeFromJSONTyped = StatisticsRevenueTypeTypeFromJSONTyped;
function StatisticsRevenueTypeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'revenueAmount': value.revenueAmount,
        'revenueLabel': value.revenueLabel,
    };
}
exports.StatisticsRevenueTypeTypeToJSON = StatisticsRevenueTypeTypeToJSON;
