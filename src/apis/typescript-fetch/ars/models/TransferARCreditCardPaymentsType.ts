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

import { exists, mapValues } from '../runtime';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Transfer AR Credit Card payment information.
 * @export
 * @interface TransferARCreditCardPaymentsType
 */
export interface TransferARCreditCardPaymentsType {
    /**
     * Hotel Code.
     * @type {string}
     * @memberof TransferARCreditCardPaymentsType
     */
    hotelId?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof TransferARCreditCardPaymentsType
     */
    accountId?: UniqueIDType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof TransferARCreditCardPaymentsType
     */
    totalAmount?: CurrencyAmountType;
    /**
     * Reference Text for the payment.
     * @type {string}
     * @memberof TransferARCreditCardPaymentsType
     */
    reference?: string;
    /**
     * Transaction number of the payments to be transfered.
     * @type {Array<number>}
     * @memberof TransferARCreditCardPaymentsType
     */
    transactionNo?: Array<number>;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof TransferARCreditCardPaymentsType
     */
    cashierId?: number;
}

/**
 * Check if a given object implements the TransferARCreditCardPaymentsType interface.
 */
export function instanceOfTransferARCreditCardPaymentsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TransferARCreditCardPaymentsTypeFromJSON(json: any): TransferARCreditCardPaymentsType {
    return TransferARCreditCardPaymentsTypeFromJSONTyped(json, false);
}

export function TransferARCreditCardPaymentsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransferARCreditCardPaymentsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'accountId': !exists(json, 'accountId') ? undefined : UniqueIDTypeFromJSON(json['accountId']),
        'totalAmount': !exists(json, 'totalAmount') ? undefined : CurrencyAmountTypeFromJSON(json['totalAmount']),
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
        'transactionNo': !exists(json, 'transactionNo') ? undefined : json['transactionNo'],
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
    };
}

export function TransferARCreditCardPaymentsTypeToJSON(value?: TransferARCreditCardPaymentsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'accountId': UniqueIDTypeToJSON(value.accountId),
        'totalAmount': CurrencyAmountTypeToJSON(value.totalAmount),
        'reference': value.reference,
        'transactionNo': value.transactionNo,
        'cashierId': value.cashierId,
    };
}

