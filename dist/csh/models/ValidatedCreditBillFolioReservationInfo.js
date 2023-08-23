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
exports.ValidatedCreditBillFolioReservationInfoToJSON = exports.ValidatedCreditBillFolioReservationInfoFromJSONTyped = exports.ValidatedCreditBillFolioReservationInfoFromJSON = exports.instanceOfValidatedCreditBillFolioReservationInfo = void 0;
const runtime_1 = require("../runtime");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ValidatedCreditBillFolioReservationInfo interface.
 */
function instanceOfValidatedCreditBillFolioReservationInfo(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfValidatedCreditBillFolioReservationInfo = instanceOfValidatedCreditBillFolioReservationInfo;
function ValidatedCreditBillFolioReservationInfoFromJSON(json) {
    return ValidatedCreditBillFolioReservationInfoFromJSONTyped(json, false);
}
exports.ValidatedCreditBillFolioReservationInfoFromJSON = ValidatedCreditBillFolioReservationInfoFromJSON;
function ValidatedCreditBillFolioReservationInfoFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : json['reservationId'],
        'room': !(0, runtime_1.exists)(json, 'room') ? undefined : json['room'],
        'nameId': !(0, runtime_1.exists)(json, 'nameId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['nameId']),
        'guestName': !(0, runtime_1.exists)(json, 'guestName') ? undefined : json['guestName'],
        'altName': !(0, runtime_1.exists)(json, 'altName') ? undefined : json['altName'],
    };
}
exports.ValidatedCreditBillFolioReservationInfoFromJSONTyped = ValidatedCreditBillFolioReservationInfoFromJSONTyped;
function ValidatedCreditBillFolioReservationInfoToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationId': value.reservationId,
        'room': value.room,
        'nameId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.nameId),
        'guestName': value.guestName,
        'altName': value.altName,
    };
}
exports.ValidatedCreditBillFolioReservationInfoToJSON = ValidatedCreditBillFolioReservationInfoToJSON;
