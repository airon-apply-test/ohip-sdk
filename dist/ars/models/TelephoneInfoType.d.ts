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
import type { TelephoneType } from './TelephoneType';
/**
 * Information on a telephone number for the customer.
 * @export
 * @interface TelephoneInfoType
 */
export interface TelephoneInfoType {
    /**
     *
     * @type {TelephoneType}
     * @memberof TelephoneInfoType
     */
    telephone?: TelephoneType;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof TelephoneInfoType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof TelephoneInfoType
     */
    type?: string;
}
/**
 * Check if a given object implements the TelephoneInfoType interface.
 */
export declare function instanceOfTelephoneInfoType(value: object): boolean;
export declare function TelephoneInfoTypeFromJSON(json: any): TelephoneInfoType;
export declare function TelephoneInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TelephoneInfoType;
export declare function TelephoneInfoTypeToJSON(value?: TelephoneInfoType | null): any;
