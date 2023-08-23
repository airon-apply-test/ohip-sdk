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


/**
 * The different types of Reminder Cycle calculation rules.
 * @export
 */
export const ReminderCycleType = {
    Days: 'Days',
    EndOfMonth: 'EndOfMonth'
} as const;
export type ReminderCycleType = typeof ReminderCycleType[keyof typeof ReminderCycleType];


export function ReminderCycleTypeFromJSON(json: any): ReminderCycleType {
    return ReminderCycleTypeFromJSONTyped(json, false);
}

export function ReminderCycleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReminderCycleType {
    return json as ReminderCycleType;
}

export function ReminderCycleTypeToJSON(value?: ReminderCycleType | null): any {
    return value as any;
}

