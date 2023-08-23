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
exports.CCBatchSettlementTypeToJSON = exports.CCBatchSettlementTypeFromJSONTyped = exports.CCBatchSettlementTypeFromJSON = exports.instanceOfCCBatchSettlementType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const ReservationId_1 = require("./ReservationId");
const ReservationPaymentMethodType_1 = require("./ReservationPaymentMethodType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the CCBatchSettlementType interface.
 */
function instanceOfCCBatchSettlementType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCCBatchSettlementType = instanceOfCCBatchSettlementType;
function CCBatchSettlementTypeFromJSON(json) {
    return CCBatchSettlementTypeFromJSONTyped(json, false);
}
exports.CCBatchSettlementTypeFromJSON = CCBatchSettlementTypeFromJSON;
function CCBatchSettlementTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'settlementId': !(0, runtime_1.exists)(json, 'settlementId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['settlementId']),
        'settlementAmount': !(0, runtime_1.exists)(json, 'settlementAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['settlementAmount']),
        'guestName': !(0, runtime_1.exists)(json, 'guestName') ? undefined : json['guestName'],
        'roomId': !(0, runtime_1.exists)(json, 'roomId') ? undefined : json['roomId'],
        'windowNo': !(0, runtime_1.exists)(json, 'windowNo') ? undefined : json['windowNo'],
        'paymentMethod': !(0, runtime_1.exists)(json, 'paymentMethod') ? undefined : (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeFromJSON)(json['paymentMethod']),
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, ReservationId_1.ReservationIdFromJSON)(json['reservationId']),
        'approvalCode': !(0, runtime_1.exists)(json, 'approvalCode') ? undefined : json['approvalCode'],
        'folioNo': !(0, runtime_1.exists)(json, 'folioNo') ? undefined : json['folioNo'],
        'transactionDate': !(0, runtime_1.exists)(json, 'transactionDate') ? undefined : (new Date(json['transactionDate'])),
    };
}
exports.CCBatchSettlementTypeFromJSONTyped = CCBatchSettlementTypeFromJSONTyped;
function CCBatchSettlementTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'settlementId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.settlementId),
        'settlementAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.settlementAmount),
        'guestName': value.guestName,
        'roomId': value.roomId,
        'windowNo': value.windowNo,
        'paymentMethod': (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeToJSON)(value.paymentMethod),
        'reservationId': (0, ReservationId_1.ReservationIdToJSON)(value.reservationId),
        'approvalCode': value.approvalCode,
        'folioNo': value.folioNo,
        'transactionDate': value.transactionDate === undefined ? undefined : (value.transactionDate.toISOString().substr(0, 10)),
    };
}
exports.CCBatchSettlementTypeToJSON = CCBatchSettlementTypeToJSON;
