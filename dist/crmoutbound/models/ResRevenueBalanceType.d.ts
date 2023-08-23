/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Revenue and Balance Amount summary for the reservation.
 * @export
 * @interface ResRevenueBalanceType
 */
export interface ResRevenueBalanceType {
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    totalFixedCharge?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    totalPayment?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    roomRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    foodAndBevRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    otherRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    nonRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    totalRevenue?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResRevenueBalanceType
     */
    balance?: CurrencyAmountType;
}
/**
 * Check if a given object implements the ResRevenueBalanceType interface.
 */
export declare function instanceOfResRevenueBalanceType(value: object): boolean;
export declare function ResRevenueBalanceTypeFromJSON(json: any): ResRevenueBalanceType;
export declare function ResRevenueBalanceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResRevenueBalanceType;
export declare function ResRevenueBalanceTypeToJSON(value?: ResRevenueBalanceType | null): any;
