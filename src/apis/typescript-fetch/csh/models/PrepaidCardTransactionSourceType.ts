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


/**
 * Vendor transaction source type.
 * @export
 */
export const PrepaidCardTransactionSourceType = {
    Opera: 'Opera',
    Vendor: 'Vendor'
} as const;
export type PrepaidCardTransactionSourceType = typeof PrepaidCardTransactionSourceType[keyof typeof PrepaidCardTransactionSourceType];


export function PrepaidCardTransactionSourceTypeFromJSON(json: any): PrepaidCardTransactionSourceType {
    return PrepaidCardTransactionSourceTypeFromJSONTyped(json, false);
}

export function PrepaidCardTransactionSourceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PrepaidCardTransactionSourceType {
    return json as PrepaidCardTransactionSourceType;
}

export function PrepaidCardTransactionSourceTypeToJSON(value?: PrepaidCardTransactionSourceType | null): any {
    return value as any;
}

