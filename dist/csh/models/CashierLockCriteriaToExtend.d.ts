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
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Operation for extending the lifetime of a lock on a cashier.
 * @export
 * @interface CashierLockCriteriaToExtend
 */
export interface CashierLockCriteriaToExtend {
    /**
     * The number of seconds to add to the lock's validity time.
     * @type {number}
     * @memberof CashierLockCriteriaToExtend
     */
    additionalTimeToLive?: number;
    /**
     * The Cashier Lock Handle to pass along with operation which required cashier to be locked.
     * @type {number}
     * @memberof CashierLockCriteriaToExtend
     */
    cashierLockHandle?: number;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof CashierLockCriteriaToExtend
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CashierLockCriteriaToExtend
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the CashierLockCriteriaToExtend interface.
 */
export declare function instanceOfCashierLockCriteriaToExtend(value: object): boolean;
export declare function CashierLockCriteriaToExtendFromJSON(json: any): CashierLockCriteriaToExtend;
export declare function CashierLockCriteriaToExtendFromJSONTyped(json: any, ignoreDiscriminator: boolean): CashierLockCriteriaToExtend;
export declare function CashierLockCriteriaToExtendToJSON(value?: CashierLockCriteriaToExtend | null): any;
