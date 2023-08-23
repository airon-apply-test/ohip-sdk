"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReservationFolioWindowTypeToJSON = exports.ReservationFolioWindowTypeFromJSONTyped = exports.ReservationFolioWindowTypeFromJSON = exports.instanceOfReservationFolioWindowType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const PayeeInfoType_1 = require("./PayeeInfoType");
/**
 * Check if a given object implements the ReservationFolioWindowType interface.
 */
function instanceOfReservationFolioWindowType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationFolioWindowType = instanceOfReservationFolioWindowType;
function ReservationFolioWindowTypeFromJSON(json) {
    return ReservationFolioWindowTypeFromJSONTyped(json, false);
}
exports.ReservationFolioWindowTypeFromJSON = ReservationFolioWindowTypeFromJSON;
function ReservationFolioWindowTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'payeeInfo': !(0, runtime_1.exists)(json, 'payeeInfo') ? undefined : (0, PayeeInfoType_1.PayeeInfoTypeFromJSON)(json['payeeInfo']),
        'balance': !(0, runtime_1.exists)(json, 'balance') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['balance']),
        'paymentMethod': !(0, runtime_1.exists)(json, 'paymentMethod') ? undefined : json['paymentMethod'],
        'folioWindowNo': !(0, runtime_1.exists)(json, 'folioWindowNo') ? undefined : json['folioWindowNo'],
    };
}
exports.ReservationFolioWindowTypeFromJSONTyped = ReservationFolioWindowTypeFromJSONTyped;
function ReservationFolioWindowTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'payeeInfo': (0, PayeeInfoType_1.PayeeInfoTypeToJSON)(value.payeeInfo),
        'balance': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.balance),
        'paymentMethod': value.paymentMethod,
        'folioWindowNo': value.folioWindowNo,
    };
}
exports.ReservationFolioWindowTypeToJSON = ReservationFolioWindowTypeToJSON;
