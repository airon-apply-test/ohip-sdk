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
 * Specifies Company or Travel Agent profile using IATA or Corp. No.
 * @export
 * @interface ResGuestExternalInfoType
 */
export interface ResGuestExternalInfoType {
    /**
     * Given name, first name or names
     * @type {string}
     * @memberof ResGuestExternalInfoType
     */
    givenName?: string;
    /**
     * Family name, last name.
     * @type {string}
     * @memberof ResGuestExternalInfoType
     */
    surname?: string;
}
/**
 * Check if a given object implements the ResGuestExternalInfoType interface.
 */
export declare function instanceOfResGuestExternalInfoType(value: object): boolean;
export declare function ResGuestExternalInfoTypeFromJSON(json: any): ResGuestExternalInfoType;
export declare function ResGuestExternalInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResGuestExternalInfoType;
export declare function ResGuestExternalInfoTypeToJSON(value?: ResGuestExternalInfoType | null): any;
