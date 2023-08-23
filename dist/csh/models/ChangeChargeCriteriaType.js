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
exports.ChangeChargeCriteriaTypeToJSON = exports.ChangeChargeCriteriaTypeFromJSONTyped = exports.ChangeChargeCriteriaTypeFromJSON = exports.instanceOfChangeChargeCriteriaType = void 0;
const runtime_1 = require("../runtime");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
/**
 * Check if a given object implements the ChangeChargeCriteriaType interface.
 */
function instanceOfChangeChargeCriteriaType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChangeChargeCriteriaType = instanceOfChangeChargeCriteriaType;
function ChangeChargeCriteriaTypeFromJSON(json) {
    return ChangeChargeCriteriaTypeFromJSONTyped(json, false);
}
exports.ChangeChargeCriteriaTypeFromJSON = ChangeChargeCriteriaTypeFromJSON;
function ChangeChargeCriteriaTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'transactionNo': !(0, runtime_1.exists)(json, 'transactionNo') ? undefined : json['transactionNo'],
        'price': !(0, runtime_1.exists)(json, 'price') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['price']),
        'quantity': !(0, runtime_1.exists)(json, 'quantity') ? undefined : json['quantity'],
        'reference': !(0, runtime_1.exists)(json, 'reference') ? undefined : json['reference'],
        'remark': !(0, runtime_1.exists)(json, 'remark') ? undefined : json['remark'],
        'checkNumber': !(0, runtime_1.exists)(json, 'checkNumber') ? undefined : json['checkNumber'],
        'revenueDate': !(0, runtime_1.exists)(json, 'revenueDate') ? undefined : (new Date(json['revenueDate'])),
        'covers': !(0, runtime_1.exists)(json, 'covers') ? undefined : json['covers'],
        'arrangementCode': !(0, runtime_1.exists)(json, 'arrangementCode') ? undefined : json['arrangementCode'],
        'approvalCode': !(0, runtime_1.exists)(json, 'approvalCode') ? undefined : json['approvalCode'],
        'approvalStatus': !(0, runtime_1.exists)(json, 'approvalStatus') ? undefined : json['approvalStatus'],
        'approvalDate': !(0, runtime_1.exists)(json, 'approvalDate') ? undefined : (new Date(json['approvalDate'])),
        'message': !(0, runtime_1.exists)(json, 'message') ? undefined : json['message'],
        'cashierId': !(0, runtime_1.exists)(json, 'cashierId') ? undefined : json['cashierId'],
    };
}
exports.ChangeChargeCriteriaTypeFromJSONTyped = ChangeChargeCriteriaTypeFromJSONTyped;
function ChangeChargeCriteriaTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'transactionNo': value.transactionNo,
        'price': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.price),
        'quantity': value.quantity,
        'reference': value.reference,
        'remark': value.remark,
        'checkNumber': value.checkNumber,
        'revenueDate': value.revenueDate === undefined ? undefined : (value.revenueDate.toISOString().substr(0, 10)),
        'covers': value.covers,
        'arrangementCode': value.arrangementCode,
        'approvalCode': value.approvalCode,
        'approvalStatus': value.approvalStatus,
        'approvalDate': value.approvalDate === undefined ? undefined : (value.approvalDate.toISOString().substr(0, 10)),
        'message': value.message,
        'cashierId': value.cashierId,
    };
}
exports.ChangeChargeCriteriaTypeToJSON = ChangeChargeCriteriaTypeToJSON;
