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
import type { ARTaxCodeType } from './ARTaxCodeType';
import {
    ARTaxCodeTypeFromJSON,
    ARTaxCodeTypeFromJSONTyped,
    ARTaxCodeTypeToJSON,
} from './ARTaxCodeType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';

/**
 * AR Old Balances Single Posting Type.
 * @export
 * @interface AROldBalanceChargeType
 */
export interface AROldBalanceChargeType {
    /**
     * User-defined posting reference.
     * @type {string}
     * @memberof AROldBalanceChargeType
     */
    postingReference?: string;
    /**
     * User-defined Supplement.
     * @type {string}
     * @memberof AROldBalanceChargeType
     */
    postingRemark?: string;
    /**
     * Date of the Posting.
     * @type {Date}
     * @memberof AROldBalanceChargeType
     */
    date?: Date;
    /**
     * The Folio number of this posting, if there was a Folio entered.
     * @type {number}
     * @memberof AROldBalanceChargeType
     */
    folioNo?: number;
    /**
     * The Fiscal Bill number of this posting
     * @type {string}
     * @memberof AROldBalanceChargeType
     */
    fiscalBillNo?: string;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof AROldBalanceChargeType
     */
    amount?: CurrencyAmountType;
    /**
     * Values of atmost 20 Taxes entered.
     * @type {Array<ARTaxCodeType>}
     * @memberof AROldBalanceChargeType
     */
    taxCodes?: Array<ARTaxCodeType>;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof AROldBalanceChargeType
     */
    paid?: CurrencyAmountType;
}

/**
 * Check if a given object implements the AROldBalanceChargeType interface.
 */
export function instanceOfAROldBalanceChargeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AROldBalanceChargeTypeFromJSON(json: any): AROldBalanceChargeType {
    return AROldBalanceChargeTypeFromJSONTyped(json, false);
}

export function AROldBalanceChargeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AROldBalanceChargeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'postingReference': !exists(json, 'postingReference') ? undefined : json['postingReference'],
        'postingRemark': !exists(json, 'postingRemark') ? undefined : json['postingRemark'],
        'date': !exists(json, 'date') ? undefined : (new Date(json['date'])),
        'folioNo': !exists(json, 'folioNo') ? undefined : json['folioNo'],
        'fiscalBillNo': !exists(json, 'fiscalBillNo') ? undefined : json['fiscalBillNo'],
        'amount': !exists(json, 'amount') ? undefined : CurrencyAmountTypeFromJSON(json['amount']),
        'taxCodes': !exists(json, 'taxCodes') ? undefined : ((json['taxCodes'] as Array<any>).map(ARTaxCodeTypeFromJSON)),
        'paid': !exists(json, 'paid') ? undefined : CurrencyAmountTypeFromJSON(json['paid']),
    };
}

export function AROldBalanceChargeTypeToJSON(value?: AROldBalanceChargeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'postingReference': value.postingReference,
        'postingRemark': value.postingRemark,
        'date': value.date === undefined ? undefined : (value.date.toISOString().substr(0,10)),
        'folioNo': value.folioNo,
        'fiscalBillNo': value.fiscalBillNo,
        'amount': CurrencyAmountTypeToJSON(value.amount),
        'taxCodes': value.taxCodes === undefined ? undefined : ((value.taxCodes as Array<any>).map(ARTaxCodeTypeToJSON)),
        'paid': CurrencyAmountTypeToJSON(value.paid),
    };
}

