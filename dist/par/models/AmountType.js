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
exports.AmountTypeToJSON = exports.AmountTypeFromJSONTyped = exports.AmountTypeFromJSON = exports.instanceOfAmountType = void 0;
const runtime_1 = require("../runtime");
const AdditionalGuestAmountType_1 = require("./AdditionalGuestAmountType");
const DiscountType_1 = require("./DiscountType");
const PointsType_1 = require("./PointsType");
const ShareDistributionInstructionType_1 = require("./ShareDistributionInstructionType");
const TotalType_1 = require("./TotalType");
/**
 * Check if a given object implements the AmountType interface.
 */
function instanceOfAmountType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAmountType = instanceOfAmountType;
function AmountTypeFromJSON(json) {
    return AmountTypeFromJSONTyped(json, false);
}
exports.AmountTypeFromJSON = AmountTypeFromJSON;
function AmountTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'base': !(0, runtime_1.exists)(json, 'base') ? undefined : (0, TotalType_1.TotalTypeFromJSON)(json['base']),
        'additionalGuestAmounts': !(0, runtime_1.exists)(json, 'additionalGuestAmounts') ? undefined : (json['additionalGuestAmounts'].map(AdditionalGuestAmountType_1.AdditionalGuestAmountTypeFromJSON)),
        'discount': !(0, runtime_1.exists)(json, 'discount') ? undefined : (0, DiscountType_1.DiscountTypeFromJSON)(json['discount']),
        'shareRatePercentage': !(0, runtime_1.exists)(json, 'shareRatePercentage') ? undefined : json['shareRatePercentage'],
        'shareDistributionInstruction': !(0, runtime_1.exists)(json, 'shareDistributionInstruction') ? undefined : (0, ShareDistributionInstructionType_1.ShareDistributionInstructionTypeFromJSON)(json['shareDistributionInstruction']),
        'total': !(0, runtime_1.exists)(json, 'total') ? undefined : (0, TotalType_1.TotalTypeFromJSON)(json['total']),
        'requiredPoints': !(0, runtime_1.exists)(json, 'requiredPoints') ? undefined : (0, PointsType_1.PointsTypeFromJSON)(json['requiredPoints']),
        'effectiveRate': !(0, runtime_1.exists)(json, 'effectiveRate') ? undefined : (0, TotalType_1.TotalTypeFromJSON)(json['effectiveRate']),
        'start': !(0, runtime_1.exists)(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !(0, runtime_1.exists)(json, 'end') ? undefined : (new Date(json['end'])),
    };
}
exports.AmountTypeFromJSONTyped = AmountTypeFromJSONTyped;
function AmountTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'base': (0, TotalType_1.TotalTypeToJSON)(value.base),
        'additionalGuestAmounts': value.additionalGuestAmounts === undefined ? undefined : (value.additionalGuestAmounts.map(AdditionalGuestAmountType_1.AdditionalGuestAmountTypeToJSON)),
        'discount': (0, DiscountType_1.DiscountTypeToJSON)(value.discount),
        'shareRatePercentage': value.shareRatePercentage,
        'shareDistributionInstruction': (0, ShareDistributionInstructionType_1.ShareDistributionInstructionTypeToJSON)(value.shareDistributionInstruction),
        'total': (0, TotalType_1.TotalTypeToJSON)(value.total),
        'requiredPoints': (0, PointsType_1.PointsTypeToJSON)(value.requiredPoints),
        'effectiveRate': (0, TotalType_1.TotalTypeToJSON)(value.effectiveRate),
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0, 10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0, 10)),
    };
}
exports.AmountTypeToJSON = AmountTypeToJSON;
