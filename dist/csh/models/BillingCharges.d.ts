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
import type { BillingChargesCriteriaType } from './BillingChargesCriteriaType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request to post a charge on a reservation's folio.
 * @export
 * @interface BillingCharges
 */
export interface BillingCharges {
    /**
     *
     * @type {BillingChargesCriteriaType}
     * @memberof BillingCharges
     */
    criteria?: BillingChargesCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof BillingCharges
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BillingCharges
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the BillingCharges interface.
 */
export declare function instanceOfBillingCharges(value: object): boolean;
export declare function BillingChargesFromJSON(json: any): BillingCharges;
export declare function BillingChargesFromJSONTyped(json: any, ignoreDiscriminator: boolean): BillingCharges;
export declare function BillingChargesToJSON(value?: BillingCharges | null): any;
