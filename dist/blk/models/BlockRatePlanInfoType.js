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
exports.BlockRatePlanInfoTypeToJSON = exports.BlockRatePlanInfoTypeFromJSONTyped = exports.BlockRatePlanInfoTypeFromJSON = exports.instanceOfBlockRatePlanInfoType = void 0;
const runtime_1 = require("../runtime");
const CancelPenaltyType_1 = require("./CancelPenaltyType");
const GuaranteeType_1 = require("./GuaranteeType");
const MealPlanCodeType_1 = require("./MealPlanCodeType");
/**
 * Check if a given object implements the BlockRatePlanInfoType interface.
 */
function instanceOfBlockRatePlanInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBlockRatePlanInfoType = instanceOfBlockRatePlanInfoType;
function BlockRatePlanInfoTypeFromJSON(json) {
    return BlockRatePlanInfoTypeFromJSONTyped(json, false);
}
exports.BlockRatePlanInfoTypeFromJSON = BlockRatePlanInfoTypeFromJSON;
function BlockRatePlanInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'shortInfo': !(0, runtime_1.exists)(json, 'shortInfo') ? undefined : json['shortInfo'],
        'longInfo': !(0, runtime_1.exists)(json, 'longInfo') ? undefined : json['longInfo'],
        'webPage': !(0, runtime_1.exists)(json, 'webPage') ? undefined : json['webPage'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'currencyCode': !(0, runtime_1.exists)(json, 'currencyCode') ? undefined : json['currencyCode'],
        'ratePlanLevel': !(0, runtime_1.exists)(json, 'ratePlanLevel') ? undefined : json['ratePlanLevel'],
        'guarantee': !(0, runtime_1.exists)(json, 'guarantee') ? undefined : (0, GuaranteeType_1.GuaranteeTypeFromJSON)(json['guarantee']),
        'cancelPenalty': !(0, runtime_1.exists)(json, 'cancelPenalty') ? undefined : (0, CancelPenaltyType_1.CancelPenaltyTypeFromJSON)(json['cancelPenalty']),
        'mealPlans': !(0, runtime_1.exists)(json, 'mealPlans') ? undefined : (json['mealPlans'].map(MealPlanCodeType_1.MealPlanCodeTypeFromJSON)),
        'ratePlanCode': !(0, runtime_1.exists)(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'ratePlanCategory': !(0, runtime_1.exists)(json, 'ratePlanCategory') ? undefined : json['ratePlanCategory'],
        'taxInclusive': !(0, runtime_1.exists)(json, 'taxInclusive') ? undefined : json['taxInclusive'],
        'serviceFeeInclusive': !(0, runtime_1.exists)(json, 'serviceFeeInclusive') ? undefined : json['serviceFeeInclusive'],
        'primary': !(0, runtime_1.exists)(json, 'primary') ? undefined : json['primary'],
        'showRateAmount': !(0, runtime_1.exists)(json, 'showRateAmount') ? undefined : json['showRateAmount'],
        'marketCode': !(0, runtime_1.exists)(json, 'marketCode') ? undefined : json['marketCode'],
        'sourceCode': !(0, runtime_1.exists)(json, 'sourceCode') ? undefined : json['sourceCode'],
        'negotiated': !(0, runtime_1.exists)(json, 'negotiated') ? undefined : json['negotiated'],
    };
}
exports.BlockRatePlanInfoTypeFromJSONTyped = BlockRatePlanInfoTypeFromJSONTyped;
function BlockRatePlanInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'shortInfo': value.shortInfo,
        'longInfo': value.longInfo,
        'webPage': value.webPage,
        'description': value.description,
        'currencyCode': value.currencyCode,
        'ratePlanLevel': value.ratePlanLevel,
        'guarantee': (0, GuaranteeType_1.GuaranteeTypeToJSON)(value.guarantee),
        'cancelPenalty': (0, CancelPenaltyType_1.CancelPenaltyTypeToJSON)(value.cancelPenalty),
        'mealPlans': value.mealPlans === undefined ? undefined : (value.mealPlans.map(MealPlanCodeType_1.MealPlanCodeTypeToJSON)),
        'ratePlanCode': value.ratePlanCode,
        'hotelId': value.hotelId,
        'ratePlanCategory': value.ratePlanCategory,
        'taxInclusive': value.taxInclusive,
        'serviceFeeInclusive': value.serviceFeeInclusive,
        'primary': value.primary,
        'showRateAmount': value.showRateAmount,
        'marketCode': value.marketCode,
        'sourceCode': value.sourceCode,
        'negotiated': value.negotiated,
    };
}
exports.BlockRatePlanInfoTypeToJSON = BlockRatePlanInfoTypeToJSON;
