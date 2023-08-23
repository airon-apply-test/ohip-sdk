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
import type { AuthorizeCompRedemptionsRQCompRedemptionType } from './AuthorizeCompRedemptionsRQCompRedemptionType';
/**
 *
 * @export
 * @interface AuthorizeCompRedemptionsRequest
 */
export interface AuthorizeCompRedemptionsRequest {
    /**
     * Collection of Complimentary Redemptions for approval.
     * @type {Array<AuthorizeCompRedemptionsRQCompRedemptionType>}
     * @memberof AuthorizeCompRedemptionsRequest
     */
    compRedemptions?: Array<AuthorizeCompRedemptionsRQCompRedemptionType>;
}
/**
 * Check if a given object implements the AuthorizeCompRedemptionsRequest interface.
 */
export declare function instanceOfAuthorizeCompRedemptionsRequest(value: object): boolean;
export declare function AuthorizeCompRedemptionsRequestFromJSON(json: any): AuthorizeCompRedemptionsRequest;
export declare function AuthorizeCompRedemptionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthorizeCompRedemptionsRequest;
export declare function AuthorizeCompRedemptionsRequestToJSON(value?: AuthorizeCompRedemptionsRequest | null): any;
