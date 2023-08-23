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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ChangePaymentCriteriaType
 */
export interface ChangePaymentCriteriaType {
    /**
     * Property code.
     * @type {string}
     * @memberof ChangePaymentCriteriaType
     */
    hotelId?: string;
    /**
     * Transaction number of the payment being corrected.
     * @type {number}
     * @memberof ChangePaymentCriteriaType
     */
    transactionNo?: number;
    /**
     * Corrected user-defined posting reference.
     * @type {string}
     * @memberof ChangePaymentCriteriaType
     */
    reference?: string;
    /**
     * Corrected user-defined posting remark.
     * @type {string}
     * @memberof ChangePaymentCriteriaType
     */
    remark?: string;
    /**
     * Corrected Cheque number.
     * @type {string}
     * @memberof ChangePaymentCriteriaType
     */
    checkNumber?: string;
    /**
     * Corrected POS covers - number of copies of receipts that got printed for that particular receipt.
     * @type {string}
     * @memberof ChangePaymentCriteriaType
     */
    covers?: string;
    /**
     * Tax Service Accounting Codes used.
     * @type {string}
     * @memberof ChangePaymentCriteriaType
     */
    depositTransactionId?: string;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof ChangePaymentCriteriaType
     */
    cashierId?: number;
}

/**
 * Check if a given object implements the ChangePaymentCriteriaType interface.
 */
export function instanceOfChangePaymentCriteriaType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ChangePaymentCriteriaTypeFromJSON(json: any): ChangePaymentCriteriaType {
    return ChangePaymentCriteriaTypeFromJSONTyped(json, false);
}

export function ChangePaymentCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangePaymentCriteriaType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'transactionNo': !exists(json, 'transactionNo') ? undefined : json['transactionNo'],
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
        'remark': !exists(json, 'remark') ? undefined : json['remark'],
        'checkNumber': !exists(json, 'checkNumber') ? undefined : json['checkNumber'],
        'covers': !exists(json, 'covers') ? undefined : json['covers'],
        'depositTransactionId': !exists(json, 'depositTransactionId') ? undefined : json['depositTransactionId'],
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
    };
}

export function ChangePaymentCriteriaTypeToJSON(value?: ChangePaymentCriteriaType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'transactionNo': value.transactionNo,
        'reference': value.reference,
        'remark': value.remark,
        'checkNumber': value.checkNumber,
        'covers': value.covers,
        'depositTransactionId': value.depositTransactionId,
        'cashierId': value.cashierId,
    };
}

