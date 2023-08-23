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
exports.ProcessReservationDepositCriteriaTypeToJSON = exports.ProcessReservationDepositCriteriaTypeFromJSONTyped = exports.ProcessReservationDepositCriteriaTypeFromJSON = exports.instanceOfProcessReservationDepositCriteriaType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const DepositProcessType_1 = require("./DepositProcessType");
const ReservationId_1 = require("./ReservationId");
/**
 * Check if a given object implements the ProcessReservationDepositCriteriaType interface.
 */
function instanceOfProcessReservationDepositCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProcessReservationDepositCriteriaType = instanceOfProcessReservationDepositCriteriaType;
function ProcessReservationDepositCriteriaTypeFromJSON(json) {
    return ProcessReservationDepositCriteriaTypeFromJSONTyped(json, false);
}
exports.ProcessReservationDepositCriteriaTypeFromJSON = ProcessReservationDepositCriteriaTypeFromJSON;
function ProcessReservationDepositCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'depositProcessType': !(0, runtime_1.exists)(json, 'depositProcessType') ? undefined : (0, DepositProcessType_1.DepositProcessTypeFromJSON)(json['depositProcessType']),
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, ReservationId_1.ReservationIdFromJSON)(json['reservationId']),
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'percentage': !(0, runtime_1.exists)(json, 'percentage') ? undefined : json['percentage'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
    };
}
exports.ProcessReservationDepositCriteriaTypeFromJSONTyped = ProcessReservationDepositCriteriaTypeFromJSONTyped;
function ProcessReservationDepositCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'depositProcessType': (0, DepositProcessType_1.DepositProcessTypeToJSON)(value.depositProcessType),
        'reservationId': (0, ReservationId_1.ReservationIdToJSON)(value.reservationId),
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'percentage': value.percentage,
        'cashierId': value.cashierId,
    };
}
exports.ProcessReservationDepositCriteriaTypeToJSON = ProcessReservationDepositCriteriaTypeToJSON;
