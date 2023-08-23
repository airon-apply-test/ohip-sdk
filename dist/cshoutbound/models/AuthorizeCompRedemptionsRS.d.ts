/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AuthorizeCompRedemptionsRSCompRedemptionType } from './AuthorizeCompRedemptionsRSCompRedemptionType';
import type { WarningType } from './WarningType';
/**
 * Response type of Complimentary Redemptions for approval.
 * @export
 * @interface AuthorizeCompRedemptionsRS
 */
export interface AuthorizeCompRedemptionsRS {
    /**
     * Collection of Complimentary Redemption codes and their respective Approval Code.
     * @type {Array<AuthorizeCompRedemptionsRSCompRedemptionType>}
     * @memberof AuthorizeCompRedemptionsRS
     */
    compRedemptions?: Array<AuthorizeCompRedemptionsRSCompRedemptionType>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof AuthorizeCompRedemptionsRS
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the AuthorizeCompRedemptionsRS interface.
 */
export declare function instanceOfAuthorizeCompRedemptionsRS(value: object): boolean;
export declare function AuthorizeCompRedemptionsRSFromJSON(json: any): AuthorizeCompRedemptionsRS;
export declare function AuthorizeCompRedemptionsRSFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthorizeCompRedemptionsRS;
export declare function AuthorizeCompRedemptionsRSToJSON(value?: AuthorizeCompRedemptionsRS | null): any;
