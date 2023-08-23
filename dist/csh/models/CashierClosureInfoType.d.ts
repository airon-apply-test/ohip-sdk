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
import type { CashierClosureNoType } from './CashierClosureNoType';
import type { CashierReportType } from './CashierReportType';
/**
 * Cashier closure Information.
 * @export
 * @interface CashierClosureInfoType
 */
export interface CashierClosureInfoType {
    /**
     *
     * @type {CashierClosureNoType}
     * @memberof CashierClosureInfoType
     */
    cashierClosureNo?: CashierClosureNoType;
    /**
     * Property this record belongs to.
     * @type {string}
     * @memberof CashierClosureInfoType
     */
    hotelId?: string;
    /**
     * Cashier title of the cashier.
     * @type {string}
     * @memberof CashierClosureInfoType
     */
    cashierTitle?: string;
    /**
     * Cashier user name.
     * @type {string}
     * @memberof CashierClosureInfoType
     */
    cashierUser?: string;
    /**
     * Date and time cashier was opened.
     * @type {Date}
     * @memberof CashierClosureInfoType
     */
    openDate?: Date;
    /**
     * Date and time cashier was closed.
     * @type {Date}
     * @memberof CashierClosureInfoType
     */
    closeDate?: Date;
    /**
     * Business date cashier was closed.
     * @type {Date}
     * @memberof CashierClosureInfoType
     */
    closureBusinessDate?: Date;
    /**
     * Information of the Cashier.
     * @type {Array<CashierReportType>}
     * @memberof CashierClosureInfoType
     */
    reportsList?: Array<CashierReportType>;
}
/**
 * Check if a given object implements the CashierClosureInfoType interface.
 */
export declare function instanceOfCashierClosureInfoType(value: object): boolean;
export declare function CashierClosureInfoTypeFromJSON(json: any): CashierClosureInfoType;
export declare function CashierClosureInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CashierClosureInfoType;
export declare function CashierClosureInfoTypeToJSON(value?: CashierClosureInfoType | null): any;
