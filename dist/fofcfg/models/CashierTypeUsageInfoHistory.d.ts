/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Statistics about the use of the cashier.
 * @export
 * @interface CashierTypeUsageInfoHistory
 */
export interface CashierTypeUsageInfoHistory {
    /**
     * The date that the cashier last opened.
     * @type {Date}
     * @memberof CashierTypeUsageInfoHistory
     */
    lastOpened?: Date;
    /**
     * The number of times the cashier has opened their account today.
     * @type {number}
     * @memberof CashierTypeUsageInfoHistory
     */
    timesOpened?: number;
}
/**
 * Check if a given object implements the CashierTypeUsageInfoHistory interface.
 */
export declare function instanceOfCashierTypeUsageInfoHistory(value: object): boolean;
export declare function CashierTypeUsageInfoHistoryFromJSON(json: any): CashierTypeUsageInfoHistory;
export declare function CashierTypeUsageInfoHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): CashierTypeUsageInfoHistory;
export declare function CashierTypeUsageInfoHistoryToJSON(value?: CashierTypeUsageInfoHistory | null): any;
