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
 * Simple type for turndown status instructions to be used in requests for fetching housekeeping rooms. Valid status values are Required, Not Required and Completed.
 * @export
 */
export const TurndownStatusType = {
    Required: 'Required',
    NotRequired: 'NotRequired',
    Compeleted: 'Compeleted'
} as const;
export type TurndownStatusType = typeof TurndownStatusType[keyof typeof TurndownStatusType];


export function TurndownStatusTypeFromJSON(json: any): TurndownStatusType {
    return TurndownStatusTypeFromJSONTyped(json, false);
}

export function TurndownStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TurndownStatusType {
    return json as TurndownStatusType;
}

export function TurndownStatusTypeToJSON(value?: TurndownStatusType | null): any {
    return value as any;
}

