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
import type { NameValueHeaderDetailType } from './NameValueHeaderDetailType';
import type { PaymentCriteriaType } from './PaymentCriteriaType';
import type { ProfileId } from './ProfileId';
/**
 * Criteria type for posting charges.
 * @export
 * @interface PasserByCriteriaType
 */
export interface PasserByCriteriaType {
    /**
     * Property where the charges are to be posted.
     * @type {string}
     * @memberof PasserByCriteriaType
     */
    hotelId?: string;
    /**
     * Collection of Charges to be posted.
     * @type {Array<ChargeCriteriaType>}
     * @memberof PasserByCriteriaType
     */
    charges?: Array<ChargeCriteriaType>;
    /**
     * The payment information to be posted.
     * @type {Array<PaymentCriteriaType>}
     * @memberof PasserByCriteriaType
     */
    payments?: Array<PaymentCriteriaType>;
    /**
     *
     * @type {FiscalServiceType}
     * @memberof PasserByCriteriaType
     */
    fiscalFolioInfo?: FiscalServiceType;
    /**
     * Date of the Audit. This is used when postings are being created using the Income Audit functionality.
     * @type {Date}
     * @memberof PasserByCriteriaType
     */
    incomeAuditDate?: Date;
    /**
     * Applicable for Fiscal Terminal. The ID of the terminal where the fiscal device is connected.
     * @type {string}
     * @memberof PasserByCriteriaType
     */
    fiscalTerminalId?: string;
    /**
     * Custom Folio Name Value Informatoin to be saved
     * @type {Array<NameValueHeaderDetailType>}
     * @memberof PasserByCriteriaType
     */
    folioNameValue?: Array<NameValueHeaderDetailType>;
    /**
     * Transaction service type which the Folio is being associated.
     * @type {string}
     * @memberof PasserByCriteriaType
     */
    trxServiceType?: string;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof PasserByCriteriaType
     */
    cashierId?: number;
    /**
     *
     * @type {ProfileId}
     * @memberof PasserByCriteriaType
     */
    profileId?: ProfileId;
}
/**
 * Check if a given object implements the PasserByCriteriaType interface.
 */
export declare function instanceOfPasserByCriteriaType(value: object): boolean;
export declare function PasserByCriteriaTypeFromJSON(json: any): PasserByCriteriaType;
export declare function PasserByCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PasserByCriteriaType;
export declare function PasserByCriteriaTypeToJSON(value?: PasserByCriteriaType | null): any;
