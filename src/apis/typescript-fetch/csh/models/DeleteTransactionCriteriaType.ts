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
 * Criteria type for deletion of transactions.
 * @export
 * @interface DeleteTransactionCriteriaType
 */
export interface DeleteTransactionCriteriaType {
    /**
     * Property code where the reservation transaction exists.
     * @type {string}
     * @memberof DeleteTransactionCriteriaType
     */
    hotelId?: string;
    /**
     * The unique transaction number of this transaction.
     * @type {Array<number>}
     * @memberof DeleteTransactionCriteriaType
     */
    transactionList?: Array<number>;
    /**
     * The reason code for the deletion.
     * @type {string}
     * @memberof DeleteTransactionCriteriaType
     */
    reasonCode?: string;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof DeleteTransactionCriteriaType
     */
    cashierId?: number;
}

/**
 * Check if a given object implements the DeleteTransactionCriteriaType interface.
 */
export function instanceOfDeleteTransactionCriteriaType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function DeleteTransactionCriteriaTypeFromJSON(json: any): DeleteTransactionCriteriaType {
    return DeleteTransactionCriteriaTypeFromJSONTyped(json, false);
}

export function DeleteTransactionCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteTransactionCriteriaType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'transactionList': !exists(json, 'transactionList') ? undefined : json['transactionList'],
        'reasonCode': !exists(json, 'reasonCode') ? undefined : json['reasonCode'],
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
    };
}

export function DeleteTransactionCriteriaTypeToJSON(value?: DeleteTransactionCriteriaType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'transactionList': value.transactionList,
        'reasonCode': value.reasonCode,
        'cashierId': value.cashierId,
    };
}

