"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.AvailRatePlanInfoTypeRatePlanInnerToJSON = exports.AvailRatePlanInfoTypeRatePlanInnerFromJSONTyped = exports.AvailRatePlanInfoTypeRatePlanInnerFromJSON = exports.instanceOfAvailRatePlanInfoTypeRatePlanInner = void 0;
const runtime_1 = require("../runtime");
const BookingChannelMappingType_1 = require("./BookingChannelMappingType");
const CancelPenaltyType_1 = require("./CancelPenaltyType");
const GuaranteePolicyType_1 = require("./GuaranteePolicyType");
const GuaranteeType_1 = require("./GuaranteeType");
const MealPlanCodeType_1 = require("./MealPlanCodeType");
const ProfileTypeType_1 = require("./ProfileTypeType");
const RatePlanChannelInfoType_1 = require("./RatePlanChannelInfoType");
const RatePlanCommissionType_1 = require("./RatePlanCommissionType");
/**
 * Check if a given object implements the AvailRatePlanInfoTypeRatePlanInner interface.
 */
function instanceOfAvailRatePlanInfoTypeRatePlanInner(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAvailRatePlanInfoTypeRatePlanInner = instanceOfAvailRatePlanInfoTypeRatePlanInner;
function AvailRatePlanInfoTypeRatePlanInnerFromJSON(json) {
    return AvailRatePlanInfoTypeRatePlanInnerFromJSONTyped(json, false);
}
exports.AvailRatePlanInfoTypeRatePlanInnerFromJSON = AvailRatePlanInfoTypeRatePlanInnerFromJSON;
function AvailRatePlanInfoTypeRatePlanInnerFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'shortInfo': !(0, runtime_1.exists)(json, 'shortInfo') ? undefined : json['shortInfo'],
        'longInfo': !(0, runtime_1.exists)(json, 'longInfo') ? undefined : json['longInfo'],
        'webPage': !(0, runtime_1.exists)(json, 'webPage') ? undefined : json['webPage'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'ratePlanChannelInfo': !(0, runtime_1.exists)(json, 'ratePlanChannelInfo') ? undefined : (0, RatePlanChannelInfoType_1.RatePlanChannelInfoTypeFromJSON)(json['ratePlanChannelInfo']),
        'currencyCode': !(0, runtime_1.exists)(json, 'currencyCode') ? undefined : json['currencyCode'],
        'bookingChannelMappings': !(0, runtime_1.exists)(json, 'bookingChannelMappings') ? undefined : (json['bookingChannelMappings'].map(BookingChannelMappingType_1.BookingChannelMappingTypeFromJSON)),
        'ratePlanLevel': !(0, runtime_1.exists)(json, 'ratePlanLevel') ? undefined : json['ratePlanLevel'],
        'guarantee': !(0, runtime_1.exists)(json, 'guarantee') ? undefined : (0, GuaranteeType_1.GuaranteeTypeFromJSON)(json['guarantee']),
        'cancelPenalty': !(0, runtime_1.exists)(json, 'cancelPenalty') ? undefined : (0, CancelPenaltyType_1.CancelPenaltyTypeFromJSON)(json['cancelPenalty']),
        'mealPlans': !(0, runtime_1.exists)(json, 'mealPlans') ? undefined : (json['mealPlans'].map(MealPlanCodeType_1.MealPlanCodeTypeFromJSON)),
        'ratePlanCode': !(0, runtime_1.exists)(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'ratePlanCategory': !(0, runtime_1.exists)(json, 'ratePlanCategory') ? undefined : json['ratePlanCategory'],
        'taxInclusive': !(0, runtime_1.exists)(json, 'taxInclusive') ? undefined : json['taxInclusive'],
        'serviceFeeInclusive': !(0, runtime_1.exists)(json, 'serviceFeeInclusive') ? undefined : json['serviceFeeInclusive'],
        'rateCommission': !(0, runtime_1.exists)(json, 'rateCommission') ? undefined : (0, RatePlanCommissionType_1.RatePlanCommissionTypeFromJSON)(json['rateCommission']),
        'resGuarantees': !(0, runtime_1.exists)(json, 'resGuarantees') ? undefined : (json['resGuarantees'].map(GuaranteePolicyType_1.GuaranteePolicyTypeFromJSON)),
        'negotiatedBy': !(0, runtime_1.exists)(json, 'negotiatedBy') ? undefined : (0, ProfileTypeType_1.ProfileTypeTypeFromJSON)(json['negotiatedBy']),
        'marketCode': !(0, runtime_1.exists)(json, 'marketCode') ? undefined : json['marketCode'],
        'sourceCode': !(0, runtime_1.exists)(json, 'sourceCode') ? undefined : json['sourceCode'],
        'hotelUseOnly': !(0, runtime_1.exists)(json, 'hotelUseOnly') ? undefined : json['hotelUseOnly'],
        'discountAllowed': !(0, runtime_1.exists)(json, 'discountAllowed') ? undefined : json['discountAllowed'],
        'credentialsRequired': !(0, runtime_1.exists)(json, 'credentialsRequired') ? undefined : json['credentialsRequired'],
    };
}
exports.AvailRatePlanInfoTypeRatePlanInnerFromJSONTyped = AvailRatePlanInfoTypeRatePlanInnerFromJSONTyped;
function AvailRatePlanInfoTypeRatePlanInnerToJSON(value) {
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
        'ratePlanChannelInfo': (0, RatePlanChannelInfoType_1.RatePlanChannelInfoTypeToJSON)(value.ratePlanChannelInfo),
        'currencyCode': value.currencyCode,
        'bookingChannelMappings': value.bookingChannelMappings === undefined ? undefined : (value.bookingChannelMappings.map(BookingChannelMappingType_1.BookingChannelMappingTypeToJSON)),
        'ratePlanLevel': value.ratePlanLevel,
        'guarantee': (0, GuaranteeType_1.GuaranteeTypeToJSON)(value.guarantee),
        'cancelPenalty': (0, CancelPenaltyType_1.CancelPenaltyTypeToJSON)(value.cancelPenalty),
        'mealPlans': value.mealPlans === undefined ? undefined : (value.mealPlans.map(MealPlanCodeType_1.MealPlanCodeTypeToJSON)),
        'ratePlanCode': value.ratePlanCode,
        'hotelId': value.hotelId,
        'ratePlanCategory': value.ratePlanCategory,
        'taxInclusive': value.taxInclusive,
        'serviceFeeInclusive': value.serviceFeeInclusive,
        'rateCommission': (0, RatePlanCommissionType_1.RatePlanCommissionTypeToJSON)(value.rateCommission),
        'resGuarantees': value.resGuarantees === undefined ? undefined : (value.resGuarantees.map(GuaranteePolicyType_1.GuaranteePolicyTypeToJSON)),
        'negotiatedBy': (0, ProfileTypeType_1.ProfileTypeTypeToJSON)(value.negotiatedBy),
        'marketCode': value.marketCode,
        'sourceCode': value.sourceCode,
        'hotelUseOnly': value.hotelUseOnly,
        'discountAllowed': value.discountAllowed,
        'credentialsRequired': value.credentialsRequired,
    };
}
exports.AvailRatePlanInfoTypeRatePlanInnerToJSON = AvailRatePlanInfoTypeRatePlanInnerToJSON;
