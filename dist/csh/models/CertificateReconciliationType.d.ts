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
import type { FBAStatusType } from './FBAStatusType';
import type { ReservationId } from './ReservationId';
/**
 * FBA ( Flexible Benefits Awards ) related fields.
 * @export
 * @interface CertificateReconciliationType
 */
export interface CertificateReconciliationType {
    /**
     *
     * @type {FBAStatusType}
     * @memberof CertificateReconciliationType
     */
    status?: FBAStatusType;
    /**
     * Award's FBA monetary values.
     * @type {number}
     * @memberof CertificateReconciliationType
     */
    monetaryValue?: number;
    /**
     * Award's FBA amount.
     * @type {number}
     * @memberof CertificateReconciliationType
     */
    amount?: number;
    /**
     * Award's FBA posted amount.
     * @type {number}
     * @memberof CertificateReconciliationType
     */
    postedAmount?: number;
    /**
     * Award's FBA reimbursed amount.
     * @type {number}
     * @memberof CertificateReconciliationType
     */
    reimbursedAmount?: number;
    /**
     * Date and time of the FBA posting.
     * @type {Date}
     * @memberof CertificateReconciliationType
     */
    postingDateTime?: Date;
    /**
     * Business date of the FBA posting.
     * @type {Date}
     * @memberof CertificateReconciliationType
     */
    postingBusinessDate?: Date;
    /**
     * Date and time of the FBA settlement.
     * @type {Date}
     * @memberof CertificateReconciliationType
     */
    settlementDateTime?: Date;
    /**
     * Business date of the FBA settlement.
     * @type {Date}
     * @memberof CertificateReconciliationType
     */
    settlementBusinessDate?: Date;
    /**
     * Date and time of the FBA reimbursement.
     * @type {Date}
     * @memberof CertificateReconciliationType
     */
    reimbursementDateTime?: Date;
    /**
     * Business date of the FBA reimbursement.
     * @type {Date}
     * @memberof CertificateReconciliationType
     */
    reimbursementBusinessDate?: Date;
    /**
     * Business date of the FBA bill generation.
     * @type {Date}
     * @memberof CertificateReconciliationType
     */
    fbaBillGenDate?: Date;
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof CertificateReconciliationType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof CertificateReconciliationType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof CertificateReconciliationType
     */
    decimalPlaces?: number;
    /**
     * Indicates if this certificate is a Flexible Benefit Award certificate.
     * @type {boolean}
     * @memberof CertificateReconciliationType
     */
    flexibleBenefitAward?: boolean;
    /**
     * Indicates whether FBA has been posted.
     * @type {boolean}
     * @memberof CertificateReconciliationType
     */
    posted?: boolean;
    /**
     * Indicates whether FBA has been settled.
     * @type {boolean}
     * @memberof CertificateReconciliationType
     */
    settled?: boolean;
    /**
     * Indicates whether FBA has been reimbursed.
     * @type {boolean}
     * @memberof CertificateReconciliationType
     */
    reimbursed?: boolean;
    /**
     * Marks if the certificate is eligible for resettlement
     * @type {boolean}
     * @memberof CertificateReconciliationType
     */
    resettleAllowed?: boolean;
    /**
     * Marks if the certificate is eligible for reimbursement.
     * @type {boolean}
     * @memberof CertificateReconciliationType
     */
    reimburseAllowed?: boolean;
    /**
     * Indicates whether the certificate is Orphan or not.
     * @type {boolean}
     * @memberof CertificateReconciliationType
     */
    orphanCertificate?: boolean;
    /**
     * Hotel context for the selected certificate.
     * @type {string}
     * @memberof CertificateReconciliationType
     */
    hotelId?: string;
    /**
     *
     * @type {ReservationId}
     * @memberof CertificateReconciliationType
     */
    reservationId?: ReservationId;
    /**
     *
     * @type {FBAStatusType}
     * @memberof CertificateReconciliationType
     */
    fBAStatus?: FBAStatusType;
    /**
     * The number for the given certificate.
     * @type {string}
     * @memberof CertificateReconciliationType
     */
    certificateNumber?: string;
}
/**
 * Check if a given object implements the CertificateReconciliationType interface.
 */
export declare function instanceOfCertificateReconciliationType(value: object): boolean;
export declare function CertificateReconciliationTypeFromJSON(json: any): CertificateReconciliationType;
export declare function CertificateReconciliationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CertificateReconciliationType;
export declare function CertificateReconciliationTypeToJSON(value?: CertificateReconciliationType | null): any;
