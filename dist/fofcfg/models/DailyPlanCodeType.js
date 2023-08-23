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
exports.DailyPlanCodeTypeToJSON = exports.DailyPlanCodeTypeFromJSONTyped = exports.DailyPlanCodeTypeFromJSON = exports.instanceOfDailyPlanCodeType = void 0;
const runtime_1 = require("../runtime");
const TranslationTextType2000_1 = require("./TranslationTextType2000");
/**
 * Check if a given object implements the DailyPlanCodeType interface.
 */
function instanceOfDailyPlanCodeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDailyPlanCodeType = instanceOfDailyPlanCodeType;
function DailyPlanCodeTypeFromJSON(json) {
    return DailyPlanCodeTypeFromJSONTyped(json, false);
}
exports.DailyPlanCodeTypeFromJSON = DailyPlanCodeTypeFromJSON;
function DailyPlanCodeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'arrangementId': !(0, runtime_1.exists)(json, 'arrangementId') ? undefined : json['arrangementId'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : (0, TranslationTextType2000_1.TranslationTextType2000FromJSON)(json['description']),
        'taxTypeCode': !(0, runtime_1.exists)(json, 'taxTypeCode') ? undefined : json['taxTypeCode'],
        'revenue': !(0, runtime_1.exists)(json, 'revenue') ? undefined : json['revenue'],
        'transactionCodes': !(0, runtime_1.exists)(json, 'transactionCodes') ? undefined : json['transactionCodes'],
    };
}
exports.DailyPlanCodeTypeFromJSONTyped = DailyPlanCodeTypeFromJSONTyped;
function DailyPlanCodeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'code': value.code,
        'arrangementId': value.arrangementId,
        'description': (0, TranslationTextType2000_1.TranslationTextType2000ToJSON)(value.description),
        'taxTypeCode': value.taxTypeCode,
        'revenue': value.revenue,
        'transactionCodes': value.transactionCodes,
    };
}
exports.DailyPlanCodeTypeToJSON = DailyPlanCodeTypeToJSON;
