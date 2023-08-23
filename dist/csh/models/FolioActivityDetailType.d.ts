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
/**
 * Details of Fiscal Folio Activity made.
 * @export
 * @interface FolioActivityDetailType
 */
export interface FolioActivityDetailType {
    /**
     * Fiscal Bill Number returned from Fiscal Program.
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    fiscalBillNo?: string;
    /**
     * Bill Number.
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    billNo?: string;
    /**
     * Folio Type.
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    folioType?: string;
    /**
     * Queue Name.
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    queueName?: string;
    /**
     * Bill Generation Date.
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    billGenerationDate?: string;
    /**
     * Return Status from Fiscal Program Application (Fiscal Printer).
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    status?: string;
    /**
     * Fiscal folio status for the partner system.
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    partnerFiscalFolioStatus?: string;
    /**
     * Number of fiscal response attempts made for the folio
     * @type {number}
     * @memberof FolioActivityDetailType
     */
    responseAttemptNo?: number;
    /**
     * Return Message from Fiscal Printing Program.
     * @type {string}
     * @memberof FolioActivityDetailType
     */
    messageText?: string;
}
/**
 * Check if a given object implements the FolioActivityDetailType interface.
 */
export declare function instanceOfFolioActivityDetailType(value: object): boolean;
export declare function FolioActivityDetailTypeFromJSON(json: any): FolioActivityDetailType;
export declare function FolioActivityDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FolioActivityDetailType;
export declare function FolioActivityDetailTypeToJSON(value?: FolioActivityDetailType | null): any;
