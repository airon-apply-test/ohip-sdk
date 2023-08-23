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
exports.ECertificateInfoTypeToJSON = exports.ECertificateInfoTypeFromJSONTyped = exports.ECertificateInfoTypeFromJSON = exports.instanceOfECertificateInfoType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const ECertificateInfoTypeHotels_1 = require("./ECertificateInfoTypeHotels");
const ECertificateUsageCriteriaType_1 = require("./ECertificateUsageCriteriaType");
/**
 * Check if a given object implements the ECertificateInfoType interface.
 */
function instanceOfECertificateInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfECertificateInfoType = instanceOfECertificateInfoType;
function ECertificateInfoTypeFromJSON(json) {
    return ECertificateInfoTypeFromJSONTyped(json, false);
}
exports.ECertificateInfoTypeFromJSON = ECertificateInfoTypeFromJSON;
function ECertificateInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'certificateType': !(0, runtime_1.exists)(json, 'certificateType') ? undefined : json['certificateType'],
        'membershipType': !(0, runtime_1.exists)(json, 'membershipType') ? undefined : json['membershipType'],
        'awardCode': !(0, runtime_1.exists)(json, 'awardCode') ? undefined : json['awardCode'],
        'promotionCode': !(0, runtime_1.exists)(json, 'promotionCode') ? undefined : json['promotionCode'],
        'voucherBenefitCode': !(0, runtime_1.exists)(json, 'voucherBenefitCode') ? undefined : json['voucherBenefitCode'],
        'hotels': !(0, runtime_1.exists)(json, 'hotels') ? undefined : (0, ECertificateInfoTypeHotels_1.ECertificateInfoTypeHotelsFromJSON)(json['hotels']),
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'longDescription': !(0, runtime_1.exists)(json, 'longDescription') ? undefined : json['longDescription'],
        'label': !(0, runtime_1.exists)(json, 'label') ? undefined : json['label'],
        'maxExtensionAllowed': !(0, runtime_1.exists)(json, 'maxExtensionAllowed') ? undefined : json['maxExtensionAllowed'],
        'usageCriteria': !(0, runtime_1.exists)(json, 'usageCriteria') ? undefined : (0, ECertificateUsageCriteriaType_1.ECertificateUsageCriteriaTypeFromJSON)(json['usageCriteria']),
        'value': !(0, runtime_1.exists)(json, 'value') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['value']),
        'cost': !(0, runtime_1.exists)(json, 'cost') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['cost']),
        'benefitSummary': !(0, runtime_1.exists)(json, 'benefitSummary') ? undefined : json['benefitSummary'],
    };
}
exports.ECertificateInfoTypeFromJSONTyped = ECertificateInfoTypeFromJSONTyped;
function ECertificateInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'certificateType': value.certificateType,
        'membershipType': value.membershipType,
        'awardCode': value.awardCode,
        'promotionCode': value.promotionCode,
        'voucherBenefitCode': value.voucherBenefitCode,
        'hotels': (0, ECertificateInfoTypeHotels_1.ECertificateInfoTypeHotelsToJSON)(value.hotels),
        'description': value.description,
        'longDescription': value.longDescription,
        'label': value.label,
        'maxExtensionAllowed': value.maxExtensionAllowed,
        'usageCriteria': (0, ECertificateUsageCriteriaType_1.ECertificateUsageCriteriaTypeToJSON)(value.usageCriteria),
        'value': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.value),
        'cost': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.cost),
        'benefitSummary': value.benefitSummary,
    };
}
exports.ECertificateInfoTypeToJSON = ECertificateInfoTypeToJSON;
