/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AmountDeterminationType } from './AmountDeterminationType';
/**
 *
 * @export
 * @interface PackageTransactionCodeType
 */
export interface PackageTransactionCodeType {
    /**
     * The description of the code.
     * @type {string}
     * @memberof PackageTransactionCodeType
     */
    description?: string;
    /**
     * Posting / transaction code.
     * @type {string}
     * @memberof PackageTransactionCodeType
     */
    code?: string;
    /**
     *
     * @type {AmountDeterminationType}
     * @memberof PackageTransactionCodeType
     */
    type?: AmountDeterminationType;
}
/**
 * Check if a given object implements the PackageTransactionCodeType interface.
 */
export declare function instanceOfPackageTransactionCodeType(value: object): boolean;
export declare function PackageTransactionCodeTypeFromJSON(json: any): PackageTransactionCodeType;
export declare function PackageTransactionCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageTransactionCodeType;
export declare function PackageTransactionCodeTypeToJSON(value?: PackageTransactionCodeType | null): any;
