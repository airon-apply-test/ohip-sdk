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
exports.PasserByCriteriaTypeToJSON = exports.PasserByCriteriaTypeFromJSONTyped = exports.PasserByCriteriaTypeFromJSON = exports.instanceOfPasserByCriteriaType = void 0;
const runtime_1 = require("../runtime");
const ChargeCriteriaType_1 = require("./ChargeCriteriaType");
const FiscalServiceType_1 = require("./FiscalServiceType");
const NameValueHeaderDetailType_1 = require("./NameValueHeaderDetailType");
const PaymentCriteriaType_1 = require("./PaymentCriteriaType");
const ProfileId_1 = require("./ProfileId");
/**
 * Check if a given object implements the PasserByCriteriaType interface.
 */
function instanceOfPasserByCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPasserByCriteriaType = instanceOfPasserByCriteriaType;
function PasserByCriteriaTypeFromJSON(json) {
    return PasserByCriteriaTypeFromJSONTyped(json, false);
}
exports.PasserByCriteriaTypeFromJSON = PasserByCriteriaTypeFromJSON;
function PasserByCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'charges': !(0, runtime_1.exists)(json, 'charges') ? undefined : (json['charges'].map(ChargeCriteriaType_1.ChargeCriteriaTypeFromJSON)),
        'payments': !(0, runtime_1.exists)(json, 'payments') ? undefined : (json['payments'].map(PaymentCriteriaType_1.PaymentCriteriaTypeFromJSON)),
        'fiscalFolioInfo': !(0, runtime_1.exists)(json, 'fiscalFolioInfo') ? undefined : (0, FiscalServiceType_1.FiscalServiceTypeFromJSON)(json['fiscalFolioInfo']),
        'incomeAuditDate': !(0, runtime_1.exists)(json, 'incomeAuditDate') ? undefined : (new Date(json['incomeAuditDate'])),
        'fiscalTerminalId': !(0, runtime_1.exists)(json, 'fiscalTerminalId') ? undefined : json['fiscalTerminalId'],
        'folioNameValue': !(0, runtime_1.exists)(json, 'folioNameValue') ? undefined : (json['folioNameValue'].map(NameValueHeaderDetailType_1.NameValueHeaderDetailTypeFromJSON)),
        'trxServiceType': !(0, runtime_1.exists)(json, 'trxServiceType') ? undefined : json['trxServiceType'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
        'profileId': !(0, runtime_1.exists)(json, 'profileId') ? undefined : (0, ProfileId_1.ProfileIdFromJSON)(json['profileId']),
    };
}
exports.PasserByCriteriaTypeFromJSONTyped = PasserByCriteriaTypeFromJSONTyped;
function PasserByCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'charges': value.charges === undefined ? undefined : (value.charges.map(ChargeCriteriaType_1.ChargeCriteriaTypeToJSON)),
        'payments': value.payments === undefined ? undefined : (value.payments.map(PaymentCriteriaType_1.PaymentCriteriaTypeToJSON)),
        'fiscalFolioInfo': (0, FiscalServiceType_1.FiscalServiceTypeToJSON)(value.fiscalFolioInfo),
        'incomeAuditDate': value.incomeAuditDate === undefined ? undefined : (value.incomeAuditDate.toISOString().substr(0, 10)),
        'fiscalTerminalId': value.fiscalTerminalId,
        'folioNameValue': value.folioNameValue === undefined ? undefined : (value.folioNameValue.map(NameValueHeaderDetailType_1.NameValueHeaderDetailTypeToJSON)),
        'trxServiceType': value.trxServiceType,
        'cashierId': value.cashierId,
        'profileId': (0, ProfileId_1.ProfileIdToJSON)(value.profileId),
    };
}
exports.PasserByCriteriaTypeToJSON = PasserByCriteriaTypeToJSON;
