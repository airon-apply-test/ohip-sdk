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
import type { ProfileId } from './ProfileId';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Criteria type for an AR Account.
 * @export
 * @interface ARAccountCriteriaType
 */
export interface ARAccountCriteriaType {
    /**
     * Property where the AR Account exists. This is required.
     * @type {string}
     * @memberof ARAccountCriteriaType
     */
    hotelId?: string;
    /**
     *
     * @type {ProfileId}
     * @memberof ARAccountCriteriaType
     */
    profileId?: ProfileId;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ARAccountCriteriaType
     */
    accountId?: UniqueIDType;
    /**
     * Account Number for the AR Account.
     * @type {string}
     * @memberof ARAccountCriteriaType
     */
    accountNo?: string;
    /**
     * Account Name for the AR Account.
     * @type {string}
     * @memberof ARAccountCriteriaType
     */
    accountName?: string;
}
/**
 * Check if a given object implements the ARAccountCriteriaType interface.
 */
export declare function instanceOfARAccountCriteriaType(value: object): boolean;
export declare function ARAccountCriteriaTypeFromJSON(json: any): ARAccountCriteriaType;
export declare function ARAccountCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARAccountCriteriaType;
export declare function ARAccountCriteriaTypeToJSON(value?: ARAccountCriteriaType | null): any;
