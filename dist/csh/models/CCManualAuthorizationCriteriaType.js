"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CCManualAuthorizationCriteriaTypeToJSON = exports.CCManualAuthorizationCriteriaTypeFromJSONTyped = exports.CCManualAuthorizationCriteriaTypeFromJSON = exports.instanceOfCCManualAuthorizationCriteriaType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the CCManualAuthorizationCriteriaType interface.
 */
function instanceOfCCManualAuthorizationCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCCManualAuthorizationCriteriaType = instanceOfCCManualAuthorizationCriteriaType;
function CCManualAuthorizationCriteriaTypeFromJSON(json) {
    return CCManualAuthorizationCriteriaTypeFromJSONTyped(json, false);
}
exports.CCManualAuthorizationCriteriaTypeFromJSON = CCManualAuthorizationCriteriaTypeFromJSON;
function CCManualAuthorizationCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'manualApprovalCode': !(0, runtime_1.exists)(json, 'manualApprovalCode') ? undefined : json['manualApprovalCode'],
        'folioView': !(0, runtime_1.exists)(json, 'folioView') ? undefined : json['folioView'],
    };
}
exports.CCManualAuthorizationCriteriaTypeFromJSONTyped = CCManualAuthorizationCriteriaTypeFromJSONTyped;
function CCManualAuthorizationCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'manualApprovalCode': value.manualApprovalCode,
        'folioView': value.folioView,
    };
}
exports.CCManualAuthorizationCriteriaTypeToJSON = CCManualAuthorizationCriteriaTypeToJSON;
