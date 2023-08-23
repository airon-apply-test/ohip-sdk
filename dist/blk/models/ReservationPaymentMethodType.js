"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReservationPaymentMethodTypeToJSON = exports.ReservationPaymentMethodTypeFromJSONTyped = exports.ReservationPaymentMethodTypeFromJSON = exports.instanceOfReservationPaymentMethodType = void 0;
const runtime_1 = require("../runtime");
const AuthorizationRuleType_1 = require("./AuthorizationRuleType");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const ResPaymentCardType_1 = require("./ResPaymentCardType");
const ReservationPaymentMethodTypeEmailFolioInfo_1 = require("./ReservationPaymentMethodTypeEmailFolioInfo");
/**
 * Check if a given object implements the ReservationPaymentMethodType interface.
 */
function instanceOfReservationPaymentMethodType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationPaymentMethodType = instanceOfReservationPaymentMethodType;
function ReservationPaymentMethodTypeFromJSON(json) {
    return ReservationPaymentMethodTypeFromJSONTyped(json, false);
}
exports.ReservationPaymentMethodTypeFromJSON = ReservationPaymentMethodTypeFromJSON;
function ReservationPaymentMethodTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'paymentCard': !(0, runtime_1.exists)(json, 'paymentCard') ? undefined : (0, ResPaymentCardType_1.ResPaymentCardTypeFromJSON)(json['paymentCard']),
        'balance': !(0, runtime_1.exists)(json, 'balance') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['balance']),
        'authorizationRule': !(0, runtime_1.exists)(json, 'authorizationRule') ? undefined : (0, AuthorizationRuleType_1.AuthorizationRuleTypeFromJSON)(json['authorizationRule']),
        'emailFolioInfo': !(0, runtime_1.exists)(json, 'emailFolioInfo') ? undefined : (0, ReservationPaymentMethodTypeEmailFolioInfo_1.ReservationPaymentMethodTypeEmailFolioInfoFromJSON)(json['emailFolioInfo']),
        'paymentMethod': !(0, runtime_1.exists)(json, 'paymentMethod') ? undefined : json['paymentMethod'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'folioView': !(0, runtime_1.exists)(json, 'folioView') ? undefined : json['folioView'],
    };
}
exports.ReservationPaymentMethodTypeFromJSONTyped = ReservationPaymentMethodTypeFromJSONTyped;
function ReservationPaymentMethodTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'paymentCard': (0, ResPaymentCardType_1.ResPaymentCardTypeToJSON)(value.paymentCard),
        'balance': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.balance),
        'authorizationRule': (0, AuthorizationRuleType_1.AuthorizationRuleTypeToJSON)(value.authorizationRule),
        'emailFolioInfo': (0, ReservationPaymentMethodTypeEmailFolioInfo_1.ReservationPaymentMethodTypeEmailFolioInfoToJSON)(value.emailFolioInfo),
        'paymentMethod': value.paymentMethod,
        'description': value.description,
        'folioView': value.folioView,
    };
}
exports.ReservationPaymentMethodTypeToJSON = ReservationPaymentMethodTypeToJSON;
