/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { DepositCancelRevenueType } from './DepositCancelRevenueType';
import type { DepositPolicyType } from './DepositPolicyType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * A deposit policy attached with the reservation.
 * @export
 * @interface ResDepositPolicyType
 */
export interface ResDepositPolicyType {
    /**
     *
     * @type {DepositCancelRevenueType}
     * @memberof ResDepositPolicyType
     */
    revenueType?: DepositCancelRevenueType;
    /**
     *
     * @type {DepositPolicyType}
     * @memberof ResDepositPolicyType
     */
    policy?: DepositPolicyType;
    /**
     * Comments attached with a deposit.
     * @type {string}
     * @memberof ResDepositPolicyType
     */
    comments?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResDepositPolicyType
     */
    amountPaid?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ResDepositPolicyType
     */
    amountDue?: CurrencyAmountType;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ResDepositPolicyType
     */
    policyId?: UniqueIDType;
}
/**
 * Check if a given object implements the ResDepositPolicyType interface.
 */
export declare function instanceOfResDepositPolicyType(value: object): boolean;
export declare function ResDepositPolicyTypeFromJSON(json: any): ResDepositPolicyType;
export declare function ResDepositPolicyTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResDepositPolicyType;
export declare function ResDepositPolicyTypeToJSON(value?: ResDepositPolicyType | null): any;
