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
import type { NameValueBaseSearchType } from './NameValueBaseSearchType';
import type { NameValueDetailType } from './NameValueDetailType';
/**
 * Name value details collected before Folio Settlement
 * @export
 * @interface NameValueHeaderDetailType
 */
export interface NameValueHeaderDetailType {
    /**
     *
     * @type {NameValueBaseSearchType}
     * @memberof NameValueHeaderDetailType
     */
    nameValueHeader?: NameValueBaseSearchType;
    /**
     *
     * @type {NameValueDetailType}
     * @memberof NameValueHeaderDetailType
     */
    nameValueDetails?: NameValueDetailType;
}
/**
 * Check if a given object implements the NameValueHeaderDetailType interface.
 */
export declare function instanceOfNameValueHeaderDetailType(value: object): boolean;
export declare function NameValueHeaderDetailTypeFromJSON(json: any): NameValueHeaderDetailType;
export declare function NameValueHeaderDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NameValueHeaderDetailType;
export declare function NameValueHeaderDetailTypeToJSON(value?: NameValueHeaderDetailType | null): any;
