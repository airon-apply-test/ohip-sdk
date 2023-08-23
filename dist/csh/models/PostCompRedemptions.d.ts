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
import type { PostCompRedemptionsCriteria } from './PostCompRedemptionsCriteria';
import type { WarningType } from './WarningType';
/**
 * Request type of complimentary bucket redemptions posting.
 * @export
 * @interface PostCompRedemptions
 */
export interface PostCompRedemptions {
    /**
     *
     * @type {PostCompRedemptionsCriteria}
     * @memberof PostCompRedemptions
     */
    criteria?: PostCompRedemptionsCriteria;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostCompRedemptions
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostCompRedemptions interface.
 */
export declare function instanceOfPostCompRedemptions(value: object): boolean;
export declare function PostCompRedemptionsFromJSON(json: any): PostCompRedemptions;
export declare function PostCompRedemptionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostCompRedemptions;
export declare function PostCompRedemptionsToJSON(value?: PostCompRedemptions | null): any;
