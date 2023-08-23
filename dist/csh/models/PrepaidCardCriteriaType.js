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
exports.PrepaidCardCriteriaTypeToJSON = exports.PrepaidCardCriteriaTypeFromJSONTyped = exports.PrepaidCardCriteriaTypeFromJSON = exports.instanceOfPrepaidCardCriteriaType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const PrepaidCardDetailsType_1 = require("./PrepaidCardDetailsType");
const ProfileId_1 = require("./ProfileId");
const ReservationId_1 = require("./ReservationId");
const SaleCriteriaType_1 = require("./SaleCriteriaType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the PrepaidCardCriteriaType interface.
 */
function instanceOfPrepaidCardCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPrepaidCardCriteriaType = instanceOfPrepaidCardCriteriaType;
function PrepaidCardCriteriaTypeFromJSON(json) {
    return PrepaidCardCriteriaTypeFromJSONTyped(json, false);
}
exports.PrepaidCardCriteriaTypeFromJSON = PrepaidCardCriteriaTypeFromJSON;
function PrepaidCardCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, ReservationId_1.ReservationIdFromJSON)(json['reservationId']),
        'name': !(0, runtime_1.exists)(json, 'name') ? undefined : json['name'],
        'firstName': !(0, runtime_1.exists)(json, 'firstName') ? undefined : json['firstName'],
        'profileId': !(0, runtime_1.exists)(json, 'profileId') ? undefined : (0, ProfileId_1.ProfileIdFromJSON)(json['profileId']),
        'cardNo': !(0, runtime_1.exists)(json, 'cardNo') ? undefined : json['cardNo'],
        'cardNumberMasked': !(0, runtime_1.exists)(json, 'cardNumberMasked') ? undefined : json['cardNumberMasked'],
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'pinCode': !(0, runtime_1.exists)(json, 'pinCode') ? undefined : json['pinCode'],
        'interfaceId': !(0, runtime_1.exists)(json, 'interfaceId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['interfaceId']),
        'cardDetails': !(0, runtime_1.exists)(json, 'cardDetails') ? undefined : (0, PrepaidCardDetailsType_1.PrepaidCardDetailsTypeFromJSON)(json['cardDetails']),
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
        'giftCard': !(0, runtime_1.exists)(json, 'giftCard') ? undefined : json['giftCard'],
        'postChargeToRoom': !(0, runtime_1.exists)(json, 'postChargeToRoom') ? undefined : json['postChargeToRoom'],
        'saleCriteria': !(0, runtime_1.exists)(json, 'saleCriteria') ? undefined : (0, SaleCriteriaType_1.SaleCriteriaTypeFromJSON)(json['saleCriteria']),
        'vendorInterfaceID': !(0, runtime_1.exists)(json, 'vendorInterfaceID') ? undefined : json['vendorInterfaceID'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
    };
}
exports.PrepaidCardCriteriaTypeFromJSONTyped = PrepaidCardCriteriaTypeFromJSONTyped;
function PrepaidCardCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'reservationId': (0, ReservationId_1.ReservationIdToJSON)(value.reservationId),
        'name': value.name,
        'firstName': value.firstName,
        'profileId': (0, ProfileId_1.ProfileIdToJSON)(value.profileId),
        'cardNo': value.cardNo,
        'cardNumberMasked': value.cardNumberMasked,
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'pinCode': value.pinCode,
        'interfaceId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.interfaceId),
        'cardDetails': (0, PrepaidCardDetailsType_1.PrepaidCardDetailsTypeToJSON)(value.cardDetails),
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
        'giftCard': value.giftCard,
        'postChargeToRoom': value.postChargeToRoom,
        'saleCriteria': (0, SaleCriteriaType_1.SaleCriteriaTypeToJSON)(value.saleCriteria),
        'vendorInterfaceID': value.vendorInterfaceID,
        'cashierId': value.cashierId,
    };
}
exports.PrepaidCardCriteriaTypeToJSON = PrepaidCardCriteriaTypeToJSON;
