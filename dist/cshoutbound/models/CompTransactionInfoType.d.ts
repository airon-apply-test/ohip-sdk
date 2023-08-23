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
/**
 * Transaction information being processed by the vendor.
 * @export
 * @interface CompTransactionInfoType
 */
export interface CompTransactionInfoType {
    /**
     * Transaction Number for which request is being submitted.
     * @type {number}
     * @memberof CompTransactionInfoType
     */
    transactionId?: number;
    /**
     * Transaction code.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    transactionCode?: string;
    /**
     * Transaction description.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    transactionDescription?: string;
    /**
     * Room number associated with the transaction.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    roomId?: string;
    /**
     * Confirmation number associated with the transaction.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    confirmationNo?: string;
    /**
     * Approval date of the posting.
     * @type {Date}
     * @memberof CompTransactionInfoType
     */
    approvalDate?: Date;
    /**
     * Approval status of the posting.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    approvalStatus?: string;
    /**
     * Approval code of the posting.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    approvalCode?: string;
    /**
     * Approval message of the posting coming from the vendor.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    approvalMessage?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof CompTransactionInfoType
     */
    amount?: CurrencyAmountType;
    /**
     * Family name, last name or Company Name.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    name?: string;
    /**
     * Given name, first name or names.
     * @type {string}
     * @memberof CompTransactionInfoType
     */
    firstName?: string;
}
/**
 * Check if a given object implements the CompTransactionInfoType interface.
 */
export declare function instanceOfCompTransactionInfoType(value: object): boolean;
export declare function CompTransactionInfoTypeFromJSON(json: any): CompTransactionInfoType;
export declare function CompTransactionInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompTransactionInfoType;
export declare function CompTransactionInfoTypeToJSON(value?: CompTransactionInfoType | null): any;
