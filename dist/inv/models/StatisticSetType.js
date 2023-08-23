"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.StatisticSetTypeToJSON = exports.StatisticSetTypeFromJSONTyped = exports.StatisticSetTypeFromJSON = exports.instanceOfStatisticSetType = void 0;
const runtime_1 = require("../runtime");
const NumericCategorySummaryType_1 = require("./NumericCategorySummaryType");
const RevenueCategorySummaryType_1 = require("./RevenueCategorySummaryType");
/**
 * Check if a given object implements the StatisticSetType interface.
 */
function instanceOfStatisticSetType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfStatisticSetType = instanceOfStatisticSetType;
function StatisticSetTypeFromJSON(json) {
    return StatisticSetTypeFromJSONTyped(json, false);
}
exports.StatisticSetTypeFromJSON = StatisticSetTypeFromJSON;
function StatisticSetTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'revenue': !(0, runtime_1.exists)(json, 'revenue') ? undefined : (json['revenue'].map(RevenueCategorySummaryType_1.RevenueCategorySummaryTypeFromJSON)),
        'inventory': !(0, runtime_1.exists)(json, 'inventory') ? undefined : (json['inventory'].map(NumericCategorySummaryType_1.NumericCategorySummaryTypeFromJSON)),
        'statisticDate': !(0, runtime_1.exists)(json, 'statisticDate') ? undefined : (new Date(json['statisticDate'])),
        'weekendDate': !(0, runtime_1.exists)(json, 'weekendDate') ? undefined : json['weekendDate'],
    };
}
exports.StatisticSetTypeFromJSONTyped = StatisticSetTypeFromJSONTyped;
function StatisticSetTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'revenue': value.revenue === undefined ? undefined : (value.revenue.map(RevenueCategorySummaryType_1.RevenueCategorySummaryTypeToJSON)),
        'inventory': value.inventory === undefined ? undefined : (value.inventory.map(NumericCategorySummaryType_1.NumericCategorySummaryTypeToJSON)),
        'statisticDate': value.statisticDate === undefined ? undefined : (value.statisticDate.toISOString().substr(0, 10)),
        'weekendDate': value.weekendDate,
    };
}
exports.StatisticSetTypeToJSON = StatisticSetTypeToJSON;
