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
import type { CurrencyAmountType } from './CurrencyAmountType';
/**
 * Balance information for AR.
 * @export
 * @interface ARBalanceType
 */
export interface ARBalanceType {
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARBalanceType
     */
    debit?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARBalanceType
     */
    credit?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARBalanceType
     */
    total?: CurrencyAmountType;
}
/**
 * Check if a given object implements the ARBalanceType interface.
 */
export declare function instanceOfARBalanceType(value: object): boolean;
export declare function ARBalanceTypeFromJSON(json: any): ARBalanceType;
export declare function ARBalanceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARBalanceType;
export declare function ARBalanceTypeToJSON(value?: ARBalanceType | null): any;
