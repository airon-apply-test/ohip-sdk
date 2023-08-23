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
exports.CCAuthorizationInstructionTypeToJSON = exports.CCAuthorizationInstructionTypeFromJSONTyped = exports.CCAuthorizationInstructionTypeFromJSON = exports.instanceOfCCAuthorizationInstructionType = void 0;
const runtime_1 = require("../runtime");
const CCAuthorizationInstructionTypeAuthorizationSetup_1 = require("./CCAuthorizationInstructionTypeAuthorizationSetup");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const ReservationInfoType_1 = require("./ReservationInfoType");
const ReservationPaymentMethodType_1 = require("./ReservationPaymentMethodType");
/**
 * Check if a given object implements the CCAuthorizationInstructionType interface.
 */
function instanceOfCCAuthorizationInstructionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCCAuthorizationInstructionType = instanceOfCCAuthorizationInstructionType;
function CCAuthorizationInstructionTypeFromJSON(json) {
    return CCAuthorizationInstructionTypeFromJSONTyped(json, false);
}
exports.CCAuthorizationInstructionTypeFromJSON = CCAuthorizationInstructionTypeFromJSON;
function CCAuthorizationInstructionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'currentApprovalAmount': !(0, runtime_1.exists)(json, 'currentApprovalAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['currentApprovalAmount']),
        'totalApprovalAmount': !(0, runtime_1.exists)(json, 'totalApprovalAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['totalApprovalAmount']),
        'incidentalAmount': !(0, runtime_1.exists)(json, 'incidentalAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['incidentalAmount']),
        'authorizationSetup': !(0, runtime_1.exists)(json, 'authorizationSetup') ? undefined : (0, CCAuthorizationInstructionTypeAuthorizationSetup_1.CCAuthorizationInstructionTypeAuthorizationSetupFromJSON)(json['authorizationSetup']),
        'reservationDetail': !(0, runtime_1.exists)(json, 'reservationDetail') ? undefined : (0, ReservationInfoType_1.ReservationInfoTypeFromJSON)(json['reservationDetail']),
        'paymentMethodInfo': !(0, runtime_1.exists)(json, 'paymentMethodInfo') ? undefined : (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeFromJSON)(json['paymentMethodInfo']),
        'sequenceNumber': !(0, runtime_1.exists)(json, 'sequenceNumber') ? undefined : json['sequenceNumber'],
        'transactionType': !(0, runtime_1.exists)(json, 'transactionType') ? undefined : json['transactionType'],
        'originalAuthSequence': !(0, runtime_1.exists)(json, 'originalAuthSequence') ? undefined : json['originalAuthSequence'],
        'usageType': !(0, runtime_1.exists)(json, 'usageType') ? undefined : json['usageType'],
        'vendorTranId': !(0, runtime_1.exists)(json, 'vendorTranId') ? undefined : json['vendorTranId'],
        'initialAuthorizationRequired': !(0, runtime_1.exists)(json, 'initialAuthorizationRequired') ? undefined : json['initialAuthorizationRequired'],
        'businessDate': !(0, runtime_1.exists)(json, 'businessDate') ? undefined : (new Date(json['businessDate'])),
    };
}
exports.CCAuthorizationInstructionTypeFromJSONTyped = CCAuthorizationInstructionTypeFromJSONTyped;
function CCAuthorizationInstructionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'currentApprovalAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.currentApprovalAmount),
        'totalApprovalAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.totalApprovalAmount),
        'incidentalAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.incidentalAmount),
        'authorizationSetup': (0, CCAuthorizationInstructionTypeAuthorizationSetup_1.CCAuthorizationInstructionTypeAuthorizationSetupToJSON)(value.authorizationSetup),
        'reservationDetail': (0, ReservationInfoType_1.ReservationInfoTypeToJSON)(value.reservationDetail),
        'paymentMethodInfo': (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeToJSON)(value.paymentMethodInfo),
        'sequenceNumber': value.sequenceNumber,
        'transactionType': value.transactionType,
        'originalAuthSequence': value.originalAuthSequence,
        'usageType': value.usageType,
        'vendorTranId': value.vendorTranId,
        'initialAuthorizationRequired': value.initialAuthorizationRequired,
        'businessDate': value.businessDate === undefined ? undefined : (value.businessDate.toISOString().substr(0, 10)),
    };
}
exports.CCAuthorizationInstructionTypeToJSON = CCAuthorizationInstructionTypeToJSON;
