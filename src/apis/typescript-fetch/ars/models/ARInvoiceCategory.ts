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
 * 
 * @export
 */
export const ARInvoiceCategory = {
    Normal: 'Normal',
    Credit: 'Credit',
    OldBalance: 'OldBalance',
    PasserBy: 'PasserBy'
} as const;
export type ARInvoiceCategory = typeof ARInvoiceCategory[keyof typeof ARInvoiceCategory];


export function ARInvoiceCategoryFromJSON(json: any): ARInvoiceCategory {
    return ARInvoiceCategoryFromJSONTyped(json, false);
}

export function ARInvoiceCategoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARInvoiceCategory {
    return json as ARInvoiceCategory;
}

export function ARInvoiceCategoryToJSON(value?: ARInvoiceCategory | null): any {
    return value as any;
}

