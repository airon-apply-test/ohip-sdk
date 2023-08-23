"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ARPaymentInvoicesTypeInnerToJSON = exports.ARPaymentInvoicesTypeInnerFromJSONTyped = exports.ARPaymentInvoicesTypeInnerFromJSON = exports.instanceOfARPaymentInvoicesTypeInner = void 0;
const runtime_1 = require("../runtime");
const ARInvoiceCategory_1 = require("./ARInvoiceCategory");
const CashierInfoType_1 = require("./CashierInfoType");
const CurrencyAmountType_1 = require("./CurrencyAmountType");
const FiscalResponseType_1 = require("./FiscalResponseType");
const FolioStatusType_1 = require("./FolioStatusType");
const FolioTextsTypeInner_1 = require("./FolioTextsTypeInner");
const MarketCodeInfoType_1 = require("./MarketCodeInfoType");
const ProfileInfoType_1 = require("./ProfileInfoType");
const ResPaymentCardType_1 = require("./ResPaymentCardType");
const ReservationId_1 = require("./ReservationId");
const ReservationInfoType_1 = require("./ReservationInfoType");
const RoomClassCodeInfoType_1 = require("./RoomClassCodeInfoType");
const SourceCodeInfoType_1 = require("./SourceCodeInfoType");
const TrxInfoType_1 = require("./TrxInfoType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ARPaymentInvoicesTypeInner interface.
 */
function instanceOfARPaymentInvoicesTypeInner(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfARPaymentInvoicesTypeInner = instanceOfARPaymentInvoicesTypeInner;
function ARPaymentInvoicesTypeInnerFromJSON(json) {
    return ARPaymentInvoicesTypeInnerFromJSONTyped(json, false);
}
exports.ARPaymentInvoicesTypeInnerFromJSON = ARPaymentInvoicesTypeInnerFromJSON;
function ARPaymentInvoicesTypeInnerFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, ReservationId_1.ReservationIdFromJSON)(json['reservationId']),
        'guestProfileId': !(0, runtime_1.exists)(json, 'guestProfileId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['guestProfileId']),
        'age': !(0, runtime_1.exists)(json, 'age') ? undefined : json['age'],
        'guestName': !(0, runtime_1.exists)(json, 'guestName') ? undefined : json['guestName'],
        'originalAmount': !(0, runtime_1.exists)(json, 'originalAmount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['originalAmount']),
        'amount': !(0, runtime_1.exists)(json, 'amount') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amount']),
        'payments': !(0, runtime_1.exists)(json, 'payments') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['payments']),
        'balance': !(0, runtime_1.exists)(json, 'balance') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['balance']),
        'reference': !(0, runtime_1.exists)(json, 'reference') ? undefined : json['reference'],
        'remark': !(0, runtime_1.exists)(json, 'remark') ? undefined : json['remark'],
        'adjusted': !(0, runtime_1.exists)(json, 'adjusted') ? undefined : json['adjusted'],
        'parentInvoiceNo': !(0, runtime_1.exists)(json, 'parentInvoiceNo') ? undefined : json['parentInvoiceNo'],
        'compressed': !(0, runtime_1.exists)(json, 'compressed') ? undefined : json['compressed'],
        'compressedDate': !(0, runtime_1.exists)(json, 'compressedDate') ? undefined : (new Date(json['compressedDate'])),
        'transferredOut': !(0, runtime_1.exists)(json, 'transferredOut') ? undefined : json['transferredOut'],
        'transferredIn': !(0, runtime_1.exists)(json, 'transferredIn') ? undefined : json['transferredIn'],
        'market': !(0, runtime_1.exists)(json, 'market') ? undefined : (0, MarketCodeInfoType_1.MarketCodeInfoTypeFromJSON)(json['market']),
        'roomClass': !(0, runtime_1.exists)(json, 'roomClass') ? undefined : (0, RoomClassCodeInfoType_1.RoomClassCodeInfoTypeFromJSON)(json['roomClass']),
        'source': !(0, runtime_1.exists)(json, 'source') ? undefined : (0, SourceCodeInfoType_1.SourceCodeInfoTypeFromJSON)(json['source']),
        'reservationInfo': !(0, runtime_1.exists)(json, 'reservationInfo') ? undefined : (0, ReservationInfoType_1.ReservationInfoTypeFromJSON)(json['reservationInfo']),
        'cashierInfo': !(0, runtime_1.exists)(json, 'cashierInfo') ? undefined : (0, CashierInfoType_1.CashierInfoTypeFromJSON)(json['cashierInfo']),
        'transactionInfo': !(0, runtime_1.exists)(json, 'transactionInfo') ? undefined : (0, TrxInfoType_1.TrxInfoTypeFromJSON)(json['transactionInfo']),
        'paymentCard': !(0, runtime_1.exists)(json, 'paymentCard') ? undefined : (0, ResPaymentCardType_1.ResPaymentCardTypeFromJSON)(json['paymentCard']),
        'addresseeInfo': !(0, runtime_1.exists)(json, 'addresseeInfo') ? undefined : (0, ProfileInfoType_1.ProfileInfoTypeFromJSON)(json['addresseeInfo']),
        'allowPartialTransferYn': !(0, runtime_1.exists)(json, 'allowPartialTransferYn') ? undefined : json['allowPartialTransferYn'],
        'printed': !(0, runtime_1.exists)(json, 'printed') ? undefined : json['printed'],
        'printedDate': !(0, runtime_1.exists)(json, 'printedDate') ? undefined : (new Date(json['printedDate'])),
        'storedFolioId': !(0, runtime_1.exists)(json, 'storedFolioId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['storedFolioId']),
        'storedFolioName': !(0, runtime_1.exists)(json, 'storedFolioName') ? undefined : json['storedFolioName'],
        'storedDebitFolioId': !(0, runtime_1.exists)(json, 'storedDebitFolioId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['storedDebitFolioId']),
        'storedDebitFolioName': !(0, runtime_1.exists)(json, 'storedDebitFolioName') ? undefined : json['storedDebitFolioName'],
        'folioTexts': !(0, runtime_1.exists)(json, 'folioTexts') ? undefined : (json['folioTexts'].map(FolioTextsTypeInner_1.FolioTextsTypeInnerFromJSON)),
        'customNumbers': !(0, runtime_1.exists)(json, 'customNumbers') ? undefined : json['customNumbers'],
        'fiscalResponseType': !(0, runtime_1.exists)(json, 'fiscalResponseType') ? undefined : (0, FiscalResponseType_1.FiscalResponseTypeFromJSON)(json['fiscalResponseType']),
        'transactionNo': !(0, runtime_1.exists)(json, 'transactionNo') ? undefined : json['transactionNo'],
        'transactionDate': !(0, runtime_1.exists)(json, 'transactionDate') ? undefined : (new Date(json['transactionDate'])),
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : json['status'],
        'transactionCode': !(0, runtime_1.exists)(json, 'transactionCode') ? undefined : json['transactionCode'],
        'folioNo': !(0, runtime_1.exists)(json, 'folioNo') ? undefined : json['folioNo'],
        'invoiceNo': !(0, runtime_1.exists)(json, 'invoiceNo') ? undefined : json['invoiceNo'],
        'fiscalBillNo': !(0, runtime_1.exists)(json, 'fiscalBillNo') ? undefined : json['fiscalBillNo'],
        'folioTypeName': !(0, runtime_1.exists)(json, 'folioTypeName') ? undefined : json['folioTypeName'],
        'internalFolioWindowID': !(0, runtime_1.exists)(json, 'internalFolioWindowID') ? undefined : json['internalFolioWindowID'],
        'folioDate': !(0, runtime_1.exists)(json, 'folioDate') ? undefined : (new Date(json['folioDate'])),
        'folioStatus': !(0, runtime_1.exists)(json, 'folioStatus') ? undefined : (0, FolioStatusType_1.FolioStatusTypeFromJSON)(json['folioStatus']),
        'folioNoWithPrefix': !(0, runtime_1.exists)(json, 'folioNoWithPrefix') ? undefined : json['folioNoWithPrefix'],
        'invoiceType': !(0, runtime_1.exists)(json, 'invoiceType') ? undefined : (0, ARInvoiceCategory_1.ARInvoiceCategoryFromJSON)(json['invoiceType']),
        'statementNo': !(0, runtime_1.exists)(json, 'statementNo') ? undefined : json['statementNo'],
        'revenueDate': !(0, runtime_1.exists)(json, 'revenueDate') ? undefined : (new Date(json['revenueDate'])),
        'closeDate': !(0, runtime_1.exists)(json, 'closeDate') ? undefined : (new Date(json['closeDate'])),
        'postingDate': !(0, runtime_1.exists)(json, 'postingDate') ? undefined : (new Date(json['postingDate'])),
        'transferDate': !(0, runtime_1.exists)(json, 'transferDate') ? undefined : (new Date(json['transferDate'])),
        'canBeModified': !(0, runtime_1.exists)(json, 'canBeModified') ? undefined : json['canBeModified'],
        'amountToApply': !(0, runtime_1.exists)(json, 'amountToApply') ? undefined : (0, CurrencyAmountType_1.CurrencyAmountTypeFromJSON)(json['amountToApply']),
    };
}
exports.ARPaymentInvoicesTypeInnerFromJSONTyped = ARPaymentInvoicesTypeInnerFromJSONTyped;
function ARPaymentInvoicesTypeInnerToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'reservationId': (0, ReservationId_1.ReservationIdToJSON)(value.reservationId),
        'guestProfileId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.guestProfileId),
        'age': value.age,
        'guestName': value.guestName,
        'originalAmount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.originalAmount),
        'amount': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amount),
        'payments': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.payments),
        'balance': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.balance),
        'reference': value.reference,
        'remark': value.remark,
        'adjusted': value.adjusted,
        'parentInvoiceNo': value.parentInvoiceNo,
        'compressed': value.compressed,
        'compressedDate': value.compressedDate === undefined ? undefined : (value.compressedDate.toISOString().substr(0, 10)),
        'transferredOut': value.transferredOut,
        'transferredIn': value.transferredIn,
        'market': (0, MarketCodeInfoType_1.MarketCodeInfoTypeToJSON)(value.market),
        'roomClass': (0, RoomClassCodeInfoType_1.RoomClassCodeInfoTypeToJSON)(value.roomClass),
        'source': (0, SourceCodeInfoType_1.SourceCodeInfoTypeToJSON)(value.source),
        'reservationInfo': (0, ReservationInfoType_1.ReservationInfoTypeToJSON)(value.reservationInfo),
        'cashierInfo': (0, CashierInfoType_1.CashierInfoTypeToJSON)(value.cashierInfo),
        'transactionInfo': (0, TrxInfoType_1.TrxInfoTypeToJSON)(value.transactionInfo),
        'paymentCard': (0, ResPaymentCardType_1.ResPaymentCardTypeToJSON)(value.paymentCard),
        'addresseeInfo': (0, ProfileInfoType_1.ProfileInfoTypeToJSON)(value.addresseeInfo),
        'allowPartialTransferYn': value.allowPartialTransferYn,
        'printed': value.printed,
        'printedDate': value.printedDate === undefined ? undefined : (value.printedDate.toISOString().substr(0, 10)),
        'storedFolioId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.storedFolioId),
        'storedFolioName': value.storedFolioName,
        'storedDebitFolioId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.storedDebitFolioId),
        'storedDebitFolioName': value.storedDebitFolioName,
        'folioTexts': value.folioTexts === undefined ? undefined : (value.folioTexts.map(FolioTextsTypeInner_1.FolioTextsTypeInnerToJSON)),
        'customNumbers': value.customNumbers,
        'fiscalResponseType': (0, FiscalResponseType_1.FiscalResponseTypeToJSON)(value.fiscalResponseType),
        'transactionNo': value.transactionNo,
        'transactionDate': value.transactionDate === undefined ? undefined : (value.transactionDate.toISOString().substr(0, 10)),
        'status': value.status,
        'transactionCode': value.transactionCode,
        'folioNo': value.folioNo,
        'invoiceNo': value.invoiceNo,
        'fiscalBillNo': value.fiscalBillNo,
        'folioTypeName': value.folioTypeName,
        'internalFolioWindowID': value.internalFolioWindowID,
        'folioDate': value.folioDate === undefined ? undefined : (value.folioDate.toISOString().substr(0, 10)),
        'folioStatus': (0, FolioStatusType_1.FolioStatusTypeToJSON)(value.folioStatus),
        'folioNoWithPrefix': value.folioNoWithPrefix,
        'invoiceType': (0, ARInvoiceCategory_1.ARInvoiceCategoryToJSON)(value.invoiceType),
        'statementNo': value.statementNo,
        'revenueDate': value.revenueDate === undefined ? undefined : (value.revenueDate.toISOString().substr(0, 10)),
        'closeDate': value.closeDate === undefined ? undefined : (value.closeDate.toISOString().substr(0, 10)),
        'postingDate': value.postingDate === undefined ? undefined : (value.postingDate.toISOString().substr(0, 10)),
        'transferDate': value.transferDate === undefined ? undefined : (value.transferDate.toISOString().substr(0, 10)),
        'canBeModified': value.canBeModified,
        'amountToApply': (0, CurrencyAmountType_1.CurrencyAmountTypeToJSON)(value.amountToApply),
    };
}
exports.ARPaymentInvoicesTypeInnerToJSON = ARPaymentInvoicesTypeInnerToJSON;
