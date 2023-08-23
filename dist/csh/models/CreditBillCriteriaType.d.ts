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
import type { ChargeCriteriaType } from './ChargeCriteriaType';
import type { FiscalServiceType } from './FiscalServiceType';
import type { FolioType } from './FolioType';
import type { NameValueHeaderDetailType } from './NameValueHeaderDetailType';
import type { PaymentCriteriaType } from './PaymentCriteriaType';
/**
 * Criteria for posting the Credit Bill. Includes charges and payments.
 * @export
 * @interface CreditBillCriteriaType
 */
export interface CreditBillCriteriaType {
    /**
     * Property where the charges are to be posted.
     * @type {string}
     * @memberof CreditBillCriteriaType
     */
    hotelId?: string;
    /**
     * Collection of Charges to be posted.
     * @type {Array<ChargeCriteriaType>}
     * @memberof CreditBillCriteriaType
     */
    charges?: Array<ChargeCriteriaType>;
    /**
     * The payment information to be posted.
     * @type {Array<PaymentCriteriaType>}
     * @memberof CreditBillCriteriaType
     */
    payments?: Array<PaymentCriteriaType>;
    /**
     *
     * @type {FiscalServiceType}
     * @memberof CreditBillCriteriaType
     */
    fiscalFolioInfo?: FiscalServiceType;
    /**
     * Date of the Audit. This is used when postings are being created using the Income Audit functionality.
     * @type {Date}
     * @memberof CreditBillCriteriaType
     */
    incomeAuditDate?: Date;
    /**
     * Applicable for Fiscal Terminal. The ID of the terminal where the fiscal device is connected.
     * @type {string}
     * @memberof CreditBillCriteriaType
     */
    fiscalTerminalId?: string;
    /**
     * Custom Folio Name Value Informatoin to be saved
     * @type {Array<NameValueHeaderDetailType>}
     * @memberof CreditBillCriteriaType
     */
    folioNameValue?: Array<NameValueHeaderDetailType>;
    /**
     * Transaction service type which the Folio is being associated.
     * @type {string}
     * @memberof CreditBillCriteriaType
     */
    trxServiceType?: string;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof CreditBillCriteriaType
     */
    cashierId?: number;
    /**
     *
     * @type {FolioType}
     * @memberof CreditBillCriteriaType
     */
    originalFolio?: FolioType;
    /**
     * Flag can be used when a Credit Bill is to be created for full set of transactions in the Original Bill.
     * @type {boolean}
     * @memberof CreditBillCriteriaType
     */
    fullCredit?: boolean;
}
/**
 * Check if a given object implements the CreditBillCriteriaType interface.
 */
export declare function instanceOfCreditBillCriteriaType(value: object): boolean;
export declare function CreditBillCriteriaTypeFromJSON(json: any): CreditBillCriteriaType;
export declare function CreditBillCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreditBillCriteriaType;
export declare function CreditBillCriteriaTypeToJSON(value?: CreditBillCriteriaType | null): any;
