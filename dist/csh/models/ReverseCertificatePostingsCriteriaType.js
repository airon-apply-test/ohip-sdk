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
exports.ReverseCertificatePostingsCriteriaTypeToJSON = exports.ReverseCertificatePostingsCriteriaTypeFromJSONTyped = exports.ReverseCertificatePostingsCriteriaTypeFromJSON = exports.instanceOfReverseCertificatePostingsCriteriaType = void 0;
const runtime_1 = require("../runtime");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ReverseCertificatePostingsCriteriaType interface.
 */
function instanceOfReverseCertificatePostingsCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReverseCertificatePostingsCriteriaType = instanceOfReverseCertificatePostingsCriteriaType;
function ReverseCertificatePostingsCriteriaTypeFromJSON(json) {
    return ReverseCertificatePostingsCriteriaTypeFromJSONTyped(json, false);
}
exports.ReverseCertificatePostingsCriteriaTypeFromJSON = ReverseCertificatePostingsCriteriaTypeFromJSON;
function ReverseCertificatePostingsCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationNameId': !(0, runtime_1.exists)(json, 'reservationNameId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['reservationNameId']),
        'transactionList': !(0, runtime_1.exists)(json, 'transactionList') ? undefined : json['transactionList'],
        'reasonCode': !(0, runtime_1.exists)(json, 'reasonCode') ? undefined : json['reasonCode'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
    };
}
exports.ReverseCertificatePostingsCriteriaTypeFromJSONTyped = ReverseCertificatePostingsCriteriaTypeFromJSONTyped;
function ReverseCertificatePostingsCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'reservationNameId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.reservationNameId),
        'transactionList': value.transactionList,
        'reasonCode': value.reasonCode,
        'cashierId': value.cashierId,
    };
}
exports.ReverseCertificatePostingsCriteriaTypeToJSON = ReverseCertificatePostingsCriteriaTypeToJSON;
