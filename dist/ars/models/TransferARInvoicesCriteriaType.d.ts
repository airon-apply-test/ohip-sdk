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
import type { ARAccountCriteriaType } from './ARAccountCriteriaType';
import type { ARInvoiceType } from './ARInvoiceType';
import type { CurrencyAmountType } from './CurrencyAmountType';
/**
 * Criteria for transferring invoice(s) to another account.
 * @export
 * @interface TransferARInvoicesCriteriaType
 */
export interface TransferARInvoicesCriteriaType {
    /**
     * AR Invoice(s) to be transferred.
     * @type {Array<ARInvoiceType>}
     * @memberof TransferARInvoicesCriteriaType
     */
    invoicesTransferType?: Array<ARInvoiceType>;
    /**
     *
     * @type {ARAccountCriteriaType}
     * @memberof TransferARInvoicesCriteriaType
     */
    fromAccount?: ARAccountCriteriaType;
    /**
     *
     * @type {ARAccountCriteriaType}
     * @memberof TransferARInvoicesCriteriaType
     */
    toAccount?: ARAccountCriteriaType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof TransferARInvoicesCriteriaType
     */
    transferAmount?: CurrencyAmountType;
    /**
     * User defined Remarks for this transfer
     * @type {string}
     * @memberof TransferARInvoicesCriteriaType
     */
    remarks?: string;
    /**
     *
     * @type {boolean}
     * @memberof TransferARInvoicesCriteriaType
     */
    overrideCreditHoldCheck?: boolean;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof TransferARInvoicesCriteriaType
     */
    cashierId?: number;
}
/**
 * Check if a given object implements the TransferARInvoicesCriteriaType interface.
 */
export declare function instanceOfTransferARInvoicesCriteriaType(value: object): boolean;
export declare function TransferARInvoicesCriteriaTypeFromJSON(json: any): TransferARInvoicesCriteriaType;
export declare function TransferARInvoicesCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransferARInvoicesCriteriaType;
export declare function TransferARInvoicesCriteriaTypeToJSON(value?: TransferARInvoicesCriteriaType | null): any;
