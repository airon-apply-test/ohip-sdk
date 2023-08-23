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
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { PrepaidCardTransactionType } from './PrepaidCardTransactionType';
/**
 * Prepaid Card Details.
 * @export
 * @interface PrepaidCardDetailsType
 */
export interface PrepaidCardDetailsType {
    /**
     * Holds fixed charge detail.
     * @type {Date}
     * @memberof PrepaidCardDetailsType
     */
    initialLoadDate?: Date;
    /**
     * Holds fixed charge detail.
     * @type {Date}
     * @memberof PrepaidCardDetailsType
     */
    activateDate?: Date;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof PrepaidCardDetailsType
     */
    initialCreditTotal?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof PrepaidCardDetailsType
     */
    creditTotal?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof PrepaidCardDetailsType
     */
    debitTotal?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof PrepaidCardDetailsType
     */
    reservedTotal?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof PrepaidCardDetailsType
     */
    balanceTotal?: CurrencyAmountType;
    /**
     * Holds fixed charge detail.
     * @type {Date}
     * @memberof PrepaidCardDetailsType
     */
    validUntilDate?: Date;
    /**
     * Holds fixed charge detail.
     * @type {Date}
     * @memberof PrepaidCardDetailsType
     */
    expiredDate?: Date;
    /**
     * Holds fixed charge detail.
     * @type {Array<PrepaidCardTransactionType>}
     * @memberof PrepaidCardDetailsType
     */
    transactions?: Array<PrepaidCardTransactionType>;
}
/**
 * Check if a given object implements the PrepaidCardDetailsType interface.
 */
export declare function instanceOfPrepaidCardDetailsType(value: object): boolean;
export declare function PrepaidCardDetailsTypeFromJSON(json: any): PrepaidCardDetailsType;
export declare function PrepaidCardDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PrepaidCardDetailsType;
export declare function PrepaidCardDetailsTypeToJSON(value?: PrepaidCardDetailsType | null): any;
