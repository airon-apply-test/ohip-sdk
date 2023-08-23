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
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { DepositMaturityType } from './DepositMaturityType';
/**
 * Deposit maturity information type.
 * @export
 * @interface DepositMaturityInfoType
 */
export interface DepositMaturityInfoType {
    /**
     *
     * @type {DepositMaturityType}
     * @memberof DepositMaturityInfoType
     */
    depositMaturityType?: DepositMaturityType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof DepositMaturityInfoType
     */
    totalAmountTransferrable?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof DepositMaturityInfoType
     */
    totalAmountDue?: CurrencyAmountType;
}
/**
 * Check if a given object implements the DepositMaturityInfoType interface.
 */
export declare function instanceOfDepositMaturityInfoType(value: object): boolean;
export declare function DepositMaturityInfoTypeFromJSON(json: any): DepositMaturityInfoType;
export declare function DepositMaturityInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DepositMaturityInfoType;
export declare function DepositMaturityInfoTypeToJSON(value?: DepositMaturityInfoType | null): any;
