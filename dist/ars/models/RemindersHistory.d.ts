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
import type { ARReminderHistoryType } from './ARReminderHistoryType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Reminders History.
 * @export
 * @interface RemindersHistory
 */
export interface RemindersHistory {
    /**
     *
     * @type {Array<ARReminderHistoryType>}
     * @memberof RemindersHistory
     */
    aRReminderHistory?: Array<ARReminderHistoryType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof RemindersHistory
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof RemindersHistory
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the RemindersHistory interface.
 */
export declare function instanceOfRemindersHistory(value: object): boolean;
export declare function RemindersHistoryFromJSON(json: any): RemindersHistory;
export declare function RemindersHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): RemindersHistory;
export declare function RemindersHistoryToJSON(value?: RemindersHistory | null): any;
