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
exports.BlockNonCompeteTypeToJSON = exports.BlockNonCompeteTypeFromJSONTyped = exports.BlockNonCompeteTypeFromJSON = exports.instanceOfBlockNonCompeteType = void 0;
const runtime_1 = require("../runtime");
const RateProtectionType_1 = require("./RateProtectionType");
/**
 * Check if a given object implements the BlockNonCompeteType interface.
 */
function instanceOfBlockNonCompeteType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBlockNonCompeteType = instanceOfBlockNonCompeteType;
function BlockNonCompeteTypeFromJSON(json) {
    return BlockNonCompeteTypeFromJSONTyped(json, false);
}
exports.BlockNonCompeteTypeFromJSON = BlockNonCompeteTypeFromJSON;
function BlockNonCompeteTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'industry': !(0, runtime_1.exists)(json, 'industry') ? undefined : json['industry'],
        'industryDescription': !(0, runtime_1.exists)(json, 'industryDescription') ? undefined : json['industryDescription'],
        'criteria': !(0, runtime_1.exists)(json, 'criteria') ? undefined : (0, RateProtectionType_1.RateProtectionTypeFromJSON)(json['criteria']),
        'protectedDates': !(0, runtime_1.exists)(json, 'protectedDates') ? undefined : json['protectedDates'],
    };
}
exports.BlockNonCompeteTypeFromJSONTyped = BlockNonCompeteTypeFromJSONTyped;
function BlockNonCompeteTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'industry': value.industry,
        'industryDescription': value.industryDescription,
        'criteria': (0, RateProtectionType_1.RateProtectionTypeToJSON)(value.criteria),
        'protectedDates': value.protectedDates,
    };
}
exports.BlockNonCompeteTypeToJSON = BlockNonCompeteTypeToJSON;
