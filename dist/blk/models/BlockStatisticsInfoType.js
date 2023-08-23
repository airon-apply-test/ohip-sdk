"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.BlockStatisticsInfoTypeToJSON = exports.BlockStatisticsInfoTypeFromJSONTyped = exports.BlockStatisticsInfoTypeFromJSON = exports.instanceOfBlockStatisticsInfoType = void 0;
const runtime_1 = require("../runtime");
const BlockStatisticsInfoTypeRate_1 = require("./BlockStatisticsInfoTypeRate");
/**
 * Check if a given object implements the BlockStatisticsInfoType interface.
 */
function instanceOfBlockStatisticsInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBlockStatisticsInfoType = instanceOfBlockStatisticsInfoType;
function BlockStatisticsInfoTypeFromJSON(json) {
    return BlockStatisticsInfoTypeFromJSONTyped(json, false);
}
exports.BlockStatisticsInfoTypeFromJSON = BlockStatisticsInfoTypeFromJSON;
function BlockStatisticsInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'inventory': !(0, runtime_1.exists)(json, 'inventory') ? undefined : json['inventory'],
        'rate': !(0, runtime_1.exists)(json, 'rate') ? undefined : (0, BlockStatisticsInfoTypeRate_1.BlockStatisticsInfoTypeRateFromJSON)(json['rate']),
        'persons': !(0, runtime_1.exists)(json, 'persons') ? undefined : json['persons'],
        'revenue': !(0, runtime_1.exists)(json, 'revenue') ? undefined : json['revenue'],
        'averageRate': !(0, runtime_1.exists)(json, 'averageRate') ? undefined : json['averageRate'],
    };
}
exports.BlockStatisticsInfoTypeFromJSONTyped = BlockStatisticsInfoTypeFromJSONTyped;
function BlockStatisticsInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'inventory': value.inventory,
        'rate': (0, BlockStatisticsInfoTypeRate_1.BlockStatisticsInfoTypeRateToJSON)(value.rate),
        'persons': value.persons,
        'revenue': value.revenue,
        'averageRate': value.averageRate,
    };
}
exports.BlockStatisticsInfoTypeToJSON = BlockStatisticsInfoTypeToJSON;
