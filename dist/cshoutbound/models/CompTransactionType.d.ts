/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Contains the transaction to be submitted to the vendor
 * @export
 * @interface CompTransactionType
 */
export interface CompTransactionType {
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof CompTransactionType
     */
    amount?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof CompTransactionType
     */
    postedAmount?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof CompTransactionType
     */
    debit?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof CompTransactionType
     */
    credit?: CurrencyAmountType;
    /**
     * Family name, last name or Company Name.
     * @type {string}
     * @memberof CompTransactionType
     */
    name?: string;
    /**
     * Given name, first name or names.
     * @type {string}
     * @memberof CompTransactionType
     */
    firstName?: string;
    /**
     * Membership number.
     * @type {string}
     * @memberof CompTransactionType
     */
    membershipId?: string;
    /**
     * Membership number.
     * @type {string}
     * @memberof CompTransactionType
     */
    membershipType?: string;
    /**
     * Approval date of the posting.
     * @type {Date}
     * @memberof CompTransactionType
     */
    approvalDate?: Date;
    /**
     * Approval status of the posting.
     * @type {string}
     * @memberof CompTransactionType
     */
    approvalStatus?: string;
    /**
     * Approval code of the posting.
     * @type {string}
     * @memberof CompTransactionType
     */
    approvalCode?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof CompTransactionType
     */
    authorizerId?: UniqueIDType;
    /**
     *
     * @type {UniqueIDType}
     * @memberof CompTransactionType
     */
    guestNameId?: UniqueIDType;
    /**
     *
     * @type {UniqueIDType}
     * @memberof CompTransactionType
     */
    resvNameId?: UniqueIDType;
    /**
     * Confirmation number associated with the transaction.
     * @type {string}
     * @memberof CompTransactionType
     */
    confirmationNo?: string;
    /**
     * Room number associated with the transaction.
     * @type {string}
     * @memberof CompTransactionType
     */
    roomId?: string;
    /**
     * Comp number associated with the transaction.
     * @type {string}
     * @memberof CompTransactionType
     */
    compId?: string;
    /**
     * Window number where the transaction is posted.
     * @type {number}
     * @memberof CompTransactionType
     */
    folioNo?: number;
    /**
     * Subgroup for the transaction.
     * @type {string}
     * @memberof CompTransactionType
     */
    subGroup?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof CompTransactionType
     */
    transactionAmount?: CurrencyAmountType;
    /**
     * Transaction Number for which request is being submitted.
     * @type {number}
     * @memberof CompTransactionType
     */
    transactionId?: number;
    /**
     * Transaction code.
     * @type {string}
     * @memberof CompTransactionType
     */
    transactionCode?: string;
    /**
     * Transaction description.
     * @type {string}
     * @memberof CompTransactionType
     */
    transactionDescription?: string;
    /**
     * Transaction status.
     * @type {string}
     * @memberof CompTransactionType
     */
    transactionStatus?: string;
    /**
     * Transaction date of the posting.
     * @type {Date}
     * @memberof CompTransactionType
     */
    transactionDate?: Date;
    /**
     * User name.
     * @type {string}
     * @memberof CompTransactionType
     */
    userName?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof CompTransactionType
     */
    externalUserId?: UniqueIDType;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof CompTransactionType
     */
    cashierId?: number;
}
/**
 * Check if a given object implements the CompTransactionType interface.
 */
export declare function instanceOfCompTransactionType(value: object): boolean;
export declare function CompTransactionTypeFromJSON(json: any): CompTransactionType;
export declare function CompTransactionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompTransactionType;
export declare function CompTransactionTypeToJSON(value?: CompTransactionType | null): any;
