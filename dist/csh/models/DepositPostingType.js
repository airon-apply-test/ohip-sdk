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
exports.DepositPostingTypeToJSON = exports.DepositPostingTypeFromJSONTyped = exports.DepositPostingTypeFromJSON = exports.instanceOfDepositPostingType = void 0;
const runtime_1 = require("../runtime");
const CashierInfoType_1 = require("./CashierInfoType");
const CashieringTransactionTypeType_1 = require("./CashieringTransactionTypeType");
const CompPostingsType_1 = require("./CompPostingsType");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const DetailPostingTypeUpdateInfo_1 = require("./DetailPostingTypeUpdateInfo");
const ExchangeAmounts_1 = require("./ExchangeAmounts");
const PostingGroupType_1 = require("./PostingGroupType");
const ReceiptSummaryType_1 = require("./ReceiptSummaryType");
const ResDepositPolicyType_1 = require("./ResDepositPolicyType");
const ReservationPaymentMethodType_1 = require("./ReservationPaymentMethodType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the DepositPostingType interface.
 */
function instanceOfDepositPostingType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfDepositPostingType = instanceOfDepositPostingType;
function DepositPostingTypeFromJSON(json) {
    return DepositPostingTypeFromJSONTyped(json, false);
}
exports.DepositPostingTypeFromJSON = DepositPostingTypeFromJSON;
function DepositPostingTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'referenceTransactionNo': !(0, runtime_1.exists)(json, 'referenceTransactionNo') ? undefined : json['referenceTransactionNo'],
        'remark': !(0, runtime_1.exists)(json, 'remark') ? undefined : json['remark'],
        'reference': !(0, runtime_1.exists)(json, 'reference') ? undefined : json['reference'],
        'checkNo': !(0, runtime_1.exists)(json, 'checkNo') ? undefined : json['checkNo'],
        'checkCount': !(0, runtime_1.exists)(json, 'checkCount') ? undefined : json['checkCount'],
        'postedAmount': !(0, runtime_1.exists)(json, 'postedAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['postedAmount']),
        'price': !(0, runtime_1.exists)(json, 'price') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['price']),
        'quantity': !(0, runtime_1.exists)(json, 'quantity') ? undefined : json['quantity'],
        'transactionType': !(0, runtime_1.exists)(json, 'transactionType') ? undefined : (0, CashieringTransactionTypeType_1.CashieringTransactionTypeTypeFromJSON)(json['transactionType']),
        'creditAmount': !(0, runtime_1.exists)(json, 'creditAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['creditAmount']),
        'debitAmount': !(0, runtime_1.exists)(json, 'debitAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['debitAmount']),
        'articleCode': !(0, runtime_1.exists)(json, 'articleCode') ? undefined : json['articleCode'],
        'cashierInfo': !(0, runtime_1.exists)(json, 'cashierInfo') ? undefined : (0, CashierInfoType_1.CashierInfoTypeFromJSON)(json['cashierInfo']),
        'paidOut': !(0, runtime_1.exists)(json, 'paidOut') ? undefined : json['paidOut'],
        'arrangementCode': !(0, runtime_1.exists)(json, 'arrangementCode') ? undefined : json['arrangementCode'],
        'groupTypeInfo': !(0, runtime_1.exists)(json, 'groupTypeInfo') ? undefined : (0, PostingGroupType_1.PostingGroupTypeFromJSON)(json['groupTypeInfo']),
        'rateCode': !(0, runtime_1.exists)(json, 'rateCode') ? undefined : json['rateCode'],
        'compRedemptionCode': !(0, runtime_1.exists)(json, 'compRedemptionCode') ? undefined : json['compRedemptionCode'],
        'updateInfo': !(0, runtime_1.exists)(json, 'updateInfo') ? undefined : (0, DetailPostingTypeUpdateInfo_1.DetailPostingTypeUpdateInfoFromJSON)(json['updateInfo']),
        'fbaCertificate': !(0, runtime_1.exists)(json, 'fbaCertificate') ? undefined : json['fbaCertificate'],
        'taxInvoiceNo': !(0, runtime_1.exists)(json, 'taxInvoiceNo') ? undefined : json['taxInvoiceNo'],
        'serviceRecovery': !(0, runtime_1.exists)(json, 'serviceRecovery') ? undefined : json['serviceRecovery'],
        'compPostingsInfo': !(0, runtime_1.exists)(json, 'compPostingsInfo') ? undefined : (0, CompPostingsType_1.CompPostingsTypeFromJSON)(json['compPostingsInfo']),
        'financialTransactionIdList': !(0, runtime_1.exists)(json, 'financialTransactionIdList') ? undefined : (json['financialTransactionIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'exchange': !(0, runtime_1.exists)(json, 'exchange') ? undefined : (0, ExchangeAmounts_1.ExchangeAmountsFromJSON)(json['exchange']),
        'approvalCode': !(0, runtime_1.exists)(json, 'approvalCode') ? undefined : json['approvalCode'],
        'approvalStatus': !(0, runtime_1.exists)(json, 'approvalStatus') ? undefined : json['approvalStatus'],
        'stampDuty': !(0, runtime_1.exists)(json, 'stampDuty') ? undefined : json['stampDuty'],
        'customCharge': !(0, runtime_1.exists)(json, 'customCharge') ? undefined : json['customCharge'],
        'transactionNo': !(0, runtime_1.exists)(json, 'transactionNo') ? undefined : json['transactionNo'],
        'transactionDate': !(0, runtime_1.exists)(json, 'transactionDate') ? undefined : (new Date(json['transactionDate'])),
        'transactionCode': !(0, runtime_1.exists)(json, 'transactionCode') ? undefined : json['transactionCode'],
        'transactionDescription': !(0, runtime_1.exists)(json, 'transactionDescription') ? undefined : json['transactionDescription'],
        'transactionAmount': !(0, runtime_1.exists)(json, 'transactionAmount') ? undefined : json['transactionAmount'],
        'postingDate': !(0, runtime_1.exists)(json, 'postingDate') ? undefined : (new Date(json['postingDate'])),
        'revenueDate': !(0, runtime_1.exists)(json, 'revenueDate') ? undefined : (new Date(json['revenueDate'])),
        'receiptNo': !(0, runtime_1.exists)(json, 'receiptNo') ? undefined : json['receiptNo'],
        'roundingDifferenceTrx': !(0, runtime_1.exists)(json, 'roundingDifferenceTrx') ? undefined : json['roundingDifferenceTrx'],
        'commissionable': !(0, runtime_1.exists)(json, 'commissionable') ? undefined : json['commissionable'],
        'reversePaymentTransactionNo': !(0, runtime_1.exists)(json, 'reversePaymentTransactionNo') ? undefined : json['reversePaymentTransactionNo'],
        'canAdjustInvoice': !(0, runtime_1.exists)(json, 'canAdjustInvoice') ? undefined : json['canAdjustInvoice'],
        'depositTransactionId': !(0, runtime_1.exists)(json, 'depositTransactionId') ? undefined : json['depositTransactionId'],
        'depositPolicy': !(0, runtime_1.exists)(json, 'depositPolicy') ? undefined : (0, ResDepositPolicyType_1.ResDepositPolicyTypeFromJSON)(json['depositPolicy']),
        'comments': !(0, runtime_1.exists)(json, 'comments') ? undefined : json['comments'],
        'paymentMethod': !(0, runtime_1.exists)(json, 'paymentMethod') ? undefined : (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeFromJSON)(json['paymentMethod']),
        'receiptSummaryInfo': !(0, runtime_1.exists)(json, 'receiptSummaryInfo') ? undefined : (0, ReceiptSummaryType_1.ReceiptSummaryTypeFromJSON)(json['receiptSummaryInfo']),
        'folioNo': !(0, runtime_1.exists)(json, 'folioNo') ? undefined : json['folioNo'],
        'transferred': !(0, runtime_1.exists)(json, 'transferred') ? undefined : json['transferred'],
        'folioTypeName': !(0, runtime_1.exists)(json, 'folioTypeName') ? undefined : json['folioTypeName'],
    };
}
exports.DepositPostingTypeFromJSONTyped = DepositPostingTypeFromJSONTyped;
function DepositPostingTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'referenceTransactionNo': value.referenceTransactionNo,
        'remark': value.remark,
        'reference': value.reference,
        'checkNo': value.checkNo,
        'checkCount': value.checkCount,
        'postedAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.postedAmount),
        'price': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.price),
        'quantity': value.quantity,
        'transactionType': (0, CashieringTransactionTypeType_1.CashieringTransactionTypeTypeToJSON)(value.transactionType),
        'creditAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.creditAmount),
        'debitAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.debitAmount),
        'articleCode': value.articleCode,
        'cashierInfo': (0, CashierInfoType_1.CashierInfoTypeToJSON)(value.cashierInfo),
        'paidOut': value.paidOut,
        'arrangementCode': value.arrangementCode,
        'groupTypeInfo': (0, PostingGroupType_1.PostingGroupTypeToJSON)(value.groupTypeInfo),
        'rateCode': value.rateCode,
        'compRedemptionCode': value.compRedemptionCode,
        'updateInfo': (0, DetailPostingTypeUpdateInfo_1.DetailPostingTypeUpdateInfoToJSON)(value.updateInfo),
        'fbaCertificate': value.fbaCertificate,
        'taxInvoiceNo': value.taxInvoiceNo,
        'serviceRecovery': value.serviceRecovery,
        'compPostingsInfo': (0, CompPostingsType_1.CompPostingsTypeToJSON)(value.compPostingsInfo),
        'financialTransactionIdList': value.financialTransactionIdList === undefined ? undefined : (value.financialTransactionIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'exchange': (0, ExchangeAmounts_1.ExchangeAmountsToJSON)(value.exchange),
        'approvalCode': value.approvalCode,
        'approvalStatus': value.approvalStatus,
        'stampDuty': value.stampDuty,
        'customCharge': value.customCharge,
        'transactionNo': value.transactionNo,
        'transactionDate': value.transactionDate === undefined ? undefined : (value.transactionDate.toISOString().substr(0, 10)),
        'transactionCode': value.transactionCode,
        'transactionDescription': value.transactionDescription,
        'transactionAmount': value.transactionAmount,
        'postingDate': value.postingDate === undefined ? undefined : (value.postingDate.toISOString().substr(0, 10)),
        'revenueDate': value.revenueDate === undefined ? undefined : (value.revenueDate.toISOString().substr(0, 10)),
        'receiptNo': value.receiptNo,
        'roundingDifferenceTrx': value.roundingDifferenceTrx,
        'commissionable': value.commissionable,
        'reversePaymentTransactionNo': value.reversePaymentTransactionNo,
        'canAdjustInvoice': value.canAdjustInvoice,
        'depositTransactionId': value.depositTransactionId,
        'depositPolicy': (0, ResDepositPolicyType_1.ResDepositPolicyTypeToJSON)(value.depositPolicy),
        'comments': value.comments,
        'paymentMethod': (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeToJSON)(value.paymentMethod),
        'receiptSummaryInfo': (0, ReceiptSummaryType_1.ReceiptSummaryTypeToJSON)(value.receiptSummaryInfo),
        'folioNo': value.folioNo,
        'transferred': value.transferred,
        'folioTypeName': value.folioTypeName,
    };
}
exports.DepositPostingTypeToJSON = DepositPostingTypeToJSON;
