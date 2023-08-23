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
import type { CashierTransactionsType } from './CashierTransactionsType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Response for the fetch cashier's shift transactions.
 * @export
 * @interface CashierTransactionsDetails
 */
export interface CashierTransactionsDetails {
    /**
     *
     * @type {CashierTransactionsType}
     * @memberof CashierTransactionsDetails
     */
    cashierTransactionsInfo?: CashierTransactionsType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof CashierTransactionsDetails
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CashierTransactionsDetails
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the CashierTransactionsDetails interface.
 */
export declare function instanceOfCashierTransactionsDetails(value: object): boolean;
export declare function CashierTransactionsDetailsFromJSON(json: any): CashierTransactionsDetails;
export declare function CashierTransactionsDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): CashierTransactionsDetails;
export declare function CashierTransactionsDetailsToJSON(value?: CashierTransactionsDetails | null): any;
