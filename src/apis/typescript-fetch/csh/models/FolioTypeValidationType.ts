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
 * Profile detail is not configured with Tax ID.
 * @export
 */
export const FolioTypeValidationType = {
    TaxId: 'TaxID'
} as const;
export type FolioTypeValidationType = typeof FolioTypeValidationType[keyof typeof FolioTypeValidationType];


export function FolioTypeValidationTypeFromJSON(json: any): FolioTypeValidationType {
    return FolioTypeValidationTypeFromJSONTyped(json, false);
}

export function FolioTypeValidationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FolioTypeValidationType {
    return json as FolioTypeValidationType;
}

export function FolioTypeValidationTypeToJSON(value?: FolioTypeValidationType | null): any {
    return value as any;
}

